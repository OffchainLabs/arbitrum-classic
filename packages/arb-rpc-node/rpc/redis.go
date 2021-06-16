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
	"github.com/pkg/errors"
)

type lockoutRedis struct {
	client *redis.Client
}

func (r *lockoutRedis) withRetry(ctx context.Context, f func() error) {
	backoff := time.Millisecond * 100
	for {
		err := f()
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

func (r *lockoutRedis) selectSequencer(ctx context.Context) (targetSequencer string) {
	r.withRetry(ctx, func() error {
		prioritiesString, err := r.client.Get(ctx, "lockout.priorities").Result()
		if err != nil {
			return err
		}
		priorities := strings.Split(prioritiesString, ",")
		for _, hostname := range priorities {
			err := r.client.Get(ctx, "lockout.liveliness."+hostname).Err()
			if err == redis.Nil {
				continue
			}
			if err != nil {
				return err
			}
			targetSequencer = hostname
			return nil
		}
		return errors.New("no prioritized sequencer online")
	})
	return
}

func (r *lockoutRedis) updateLiveliness(ctx context.Context, hostname string, timeout time.Duration) {
	r.withRetry(ctx, func() error {
		return r.client.Set(ctx, "lockout.liveliness."+hostname, "OK", timeout).Err()
	})
}

func (r *lockoutRedis) removeLiveliness(ctx context.Context, hostname string) {
	r.withRetry(ctx, func() error {
		return r.client.Del(ctx, "lockout.liveliness."+hostname).Err()
	})
}

func (r *lockoutRedis) acquireLockout(ctx context.Context, hostname string, timeout time.Duration) (hasLockUntil time.Time) {
	r.withRetry(ctx, func() error {
		hasLockUntil = time.Now().Add(timeout)
		created, err := r.client.SetNX(ctx, "lockout.lockout", hostname, timeout).Result()
		if err != nil {
			return err
		}
		if !created {
			hasLockUntil = time.Time{}
		}
		return nil
	})
	return
}

func (r *lockoutRedis) releaseLockoutNoRetry(ctx context.Context, hostname string) error {
	return r.client.Del(ctx, "lockout.lockout").Err()
}

func (r *lockoutRedis) getLockout(ctx context.Context) (hostname string) {
	r.withRetry(ctx, func() error {
		var err error
		hostname, err = r.client.Get(ctx, "lockout.lockout").Result()
		return err
	})
	return
}

func (r *lockoutRedis) getLatestSeqNum(ctx context.Context) (seqNum *big.Int) {
	r.withRetry(ctx, func() error {
		seqNumString, err := r.client.Get(ctx, "lockout.sequenceNumber").Result()
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

func (r *lockoutRedis) updateLatestSeqNum(ctx context.Context, seqNum *big.Int) {
	r.withRetry(ctx, func() error {
		return r.client.Set(ctx, "lockout.sequenceNumber", seqNum.String(), 0).Err()
	})
}
