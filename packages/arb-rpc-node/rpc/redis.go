/*
 * Copyright 2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rpc

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/pkg/errors"
)

type lockoutRedis struct {
	client        *redis.Client
	rpc           string
	timeout       time.Duration
	maxLatency    time.Duration
	seqNumTimeout time.Duration
}

const LOCKOUT_KEY string = "lockout.lockout"
const PRIORITIES_KEY string = "lockout.priorities"
const LIVELINESS_KEY_PREFIX string = "lockout.liveliness."
const SEQUENCE_NUMBER_KEY string = "lockout.sequenceNumber"

func newLockoutRedis(config configuration.Lockout) (*lockoutRedis, error) {
	opts, err := redis.ParseURL(config.Redis)
	if err != nil {
		return nil, err
	}
	return &lockoutRedis{
		client:        redis.NewClient(opts),
		rpc:           config.SelfRPCURL,
		timeout:       config.Timeout,
		maxLatency:    config.MaxLatency,
		seqNumTimeout: config.SeqNumTimeout,
	}, nil
}

func withRetry(ctx context.Context, f func() error) {
	backoff := time.Millisecond * 100
	for {
		select {
		case <-ctx.Done():
			logger.Warn().Msg("redis context canceled")
			return
		default:
		}
		err := errors.WithStack(f())
		if err == nil {
			return
		}
		logger.Warn().Err(err).Msg("redis error")
		time.Sleep(backoff)
		if backoff < time.Second*2 {
			backoff *= 2
		}
	}
}

func withTimeout(parentCtx context.Context, timeout time.Time, f func(context.Context) error) {
	if timeout.Before(time.Now()) {
		return
	}
	timedCtx, cancelTimedCtx := context.WithDeadline(parentCtx, timeout)
	withRetry(timedCtx, func() error {
		return f(timedCtx)
	})
	cancelTimedCtx()
}

func (r *lockoutRedis) selectSequencer(ctx context.Context) (targetSequencer string) {
	withRetry(ctx, func() error {
		prioritiesString, err := r.client.Get(ctx, PRIORITIES_KEY).Result()
		if err == redis.Nil {
			return errors.New("sequencer priorities unset")
		}
		if err != nil {
			return err
		}
		priorities := strings.Split(prioritiesString, ",")
		for _, rpc := range priorities {
			err := r.client.Get(ctx, LIVELINESS_KEY_PREFIX+rpc).Err()
			if err == redis.Nil {
				continue
			}
			if err != nil {
				return err
			}
			targetSequencer = rpc
			return nil
		}
		targetSequencer = ""
		return nil
	})
	return
}

func (r *lockoutRedis) acquireGenericLockout(ctx context.Context, key string, value string, timeout time.Duration, new bool) (hasLockUntil time.Time) {
	withRetry(ctx, func() error {
		attemptingLockUntil := time.Now().Add(timeout)
		var created bool
		var err error
		if new {
			created, err = r.client.SetNX(ctx, key, value, timeout).Result()
		} else {
			err = r.client.Set(ctx, key, value, timeout).Err()
			created = true
		}
		if err != nil {
			return err
		}
		if created {
			hasLockUntil = attemptingLockUntil
		}
		return nil
	})
	return
}

// This series of methods reads and then possibly modifies hasLockUntil via a pointer.
// This ensures that the lockout isn't overrun when it is used, and that the new value is updated.

func (r *lockoutRedis) acquireOrUpdateGenericLockout(ctx context.Context, key string, value string, hasLockUntil *time.Time) {
	if hasLockUntil.Before(time.Now()) {
		*hasLockUntil = r.acquireGenericLockout(ctx, key, value, r.timeout, true)
	} else {
		timedCtx, cancelTimedCtx := context.WithDeadline(ctx, *hasLockUntil)
		*hasLockUntil = r.acquireGenericLockout(timedCtx, key, value, r.timeout, false)
		cancelTimedCtx()
	}
	if *hasLockUntil != (time.Time{}) {
		*hasLockUntil = hasLockUntil.Add(-r.maxLatency)
	}
}

func (r *lockoutRedis) releaseGenericLockout(parentCtx context.Context, key string, hasLockUntil *time.Time) {
	timeout := *hasLockUntil
	*hasLockUntil = time.Time{}
	withTimeout(parentCtx, timeout, func(timedCtx context.Context) error {
		return r.client.Del(timedCtx, key).Err()
	})
}

func (r *lockoutRedis) acquireOrUpdateLockout(ctx context.Context, hasLockUntil *time.Time) {
	r.acquireOrUpdateGenericLockout(ctx, LOCKOUT_KEY, r.rpc, hasLockUntil)
}

func (r *lockoutRedis) releaseLockout(ctx context.Context, hasLockUntil *time.Time) {
	r.releaseGenericLockout(ctx, LOCKOUT_KEY, hasLockUntil)
}

func (r *lockoutRedis) acquireOrUpdateLiveliness(ctx context.Context, hasLockUntil *time.Time) {
	r.acquireOrUpdateGenericLockout(ctx, LIVELINESS_KEY_PREFIX+r.rpc, "OK", hasLockUntil)
}

func (r *lockoutRedis) releaseLiveliness(ctx context.Context, hasLockUntil *time.Time) {
	r.releaseGenericLockout(ctx, LIVELINESS_KEY_PREFIX+r.rpc, hasLockUntil)
}

func (r *lockoutRedis) getLockout(ctx context.Context) (rpc string) {
	withRetry(ctx, func() error {
		var err error
		rpc, err = r.client.Get(ctx, LOCKOUT_KEY).Result()
		if err == redis.Nil {
			rpc = ""
			err = nil
		}
		return err
	})
	return
}

func (r *lockoutRedis) getLatestSeqNum(ctx context.Context) (seqNum *big.Int) {
	withRetry(ctx, func() error {
		seqNumString, err := r.client.Get(ctx, SEQUENCE_NUMBER_KEY).Result()
		if err == redis.Nil {
			seqNum = big.NewInt(0)
			return nil
		}
		if err != nil {
			return err
		}
		var ok bool
		seqNum, ok = new(big.Int).SetString(seqNumString, 10)
		if !ok {
			return errors.New("invalid sequence number in redis")
		}
		return nil
	})
	return
}

func (r *lockoutRedis) updateLatestSeqNum(parentCtx context.Context, seqNum *big.Int, hasLockUntil time.Time) {
	withTimeout(parentCtx, hasLockUntil, func(timedCtx context.Context) error {
		return r.client.Set(timedCtx, SEQUENCE_NUMBER_KEY, seqNum.String(), r.seqNumTimeout).Err()
	})
}
