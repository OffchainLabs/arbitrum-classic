/*
 * Copyright 2021, Offchain Labs, Inc.
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
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	ethcore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/batcher"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
)

var logger = log.With().Caller().Stack().Str("component", "rpc").Logger()

type LockoutBatcher struct {
	// Mutex protects currentBatcher and lockoutExpiresAt
	mutex            sync.RWMutex
	sequencerBatcher *batcher.SequencerBatcher
	core             core.ArbOutputLookup
	inboxReader      *monitor.InboxReader
	redis            *lockoutRedis
	errChan          chan error
	config           configuration.Lockout

	lockoutExpiresAt    time.Time
	livelinessExpiresAt time.Time
	currentSeq          string
	lastLockedSeqNum    *big.Int
	currentBatcher      batcher.TransactionBatcher
}

func SetupLockout(
	ctx context.Context,
	seqBatcher *batcher.SequencerBatcher,
	core core.ArbOutputLookup,
	inboxReader *monitor.InboxReader,
	config configuration.Lockout,
	errChan chan error,
) (*LockoutBatcher, error) {
	redis, err := newLockoutRedis(config)
	if err != nil {
		return nil, err
	}
	newBatcher := &LockoutBatcher{
		sequencerBatcher: seqBatcher,
		currentSeq:       "[starting up]",
		core:             core,
		inboxReader:      inboxReader,
		config:           config,
		redis:            redis,
		errChan:          errChan,
	}
	newBatcher.currentBatcher = newBatcher.getErrorBatcher(errors.New("sequencer lockout manager starting up"))
	newBatcher.sequencerBatcher.LockoutManager = newBatcher
	go newBatcher.lockoutManager(ctx)
	return newBatcher, nil
}

const ACCEPTABLE_SEQ_NUM_GAP int64 = 0

func (b *LockoutBatcher) getErrorBatcher(err error) *errorBatcher {
	return &errorBatcher{
		err:        err,
		aggregator: b.sequencerBatcher.Aggregator(),
	}
}

func (b *LockoutBatcher) lockoutManager(ctx context.Context) {
	holdingMutex := false
	defer (func() {
		if !holdingMutex {
			b.mutex.Lock()
		}
		b.currentBatcher = b.getErrorBatcher(errors.New("sequencer lockout manager starting up"))
		backgroundContext := context.Background()
		b.redis.releaseLockout(backgroundContext, &b.lockoutExpiresAt)
		b.redis.releaseLiveliness(backgroundContext, &b.livelinessExpiresAt)
		b.mutex.Unlock()
		holdingMutex = false
		logger.Debug().Msg("shut down sequencer lockout manager and released locks")
		select {
		case <-ctx.Done():
			break
		default:
			// We aren't shutting down but the lockout manager died
			logger.Error().Msg("sequencer lockout manager died but context isn't shutting down")
			go (func() {
				// We need this goroutine to die so the panic is printed out,
				// but we also want to exit the process as a whole afterwards.
				time.Sleep(time.Second)
				b.errChan <- errors.New("sequencer lockout manager died")
			})()
		}
	})()
	for {
		alive := true
		if !b.hasSequencerLockout() {
			currentSeqNum, err := b.core.GetMessageCount()
			if err != nil {
				logger.Warn().Err(err).Msg("error getting sequence number")
				select {
				case <-ctx.Done():
					return
				case <-time.After(5 * time.Second):
				}
				continue
			}
			if b.lastLockedSeqNum == nil || new(big.Int).Add(currentSeqNum, big.NewInt(ACCEPTABLE_SEQ_NUM_GAP)).Cmp(b.lastLockedSeqNum) < 0 {
				alive = false
				if b.livelinessExpiresAt.After(time.Now()) {
					logger.Warn().Str("ourSeqNum", currentSeqNum.String()).Str("targetSeqNum", b.lastLockedSeqNum.String()).Msg("fell behind sequencer position")
					b.redis.releaseLiveliness(ctx, &b.livelinessExpiresAt)
				}
			}
			b.lastLockedSeqNum = b.redis.getLatestSeqNum(ctx)
		}
		if alive {
			b.redis.acquireOrUpdateLiveliness(ctx, &b.livelinessExpiresAt)
			if b.livelinessExpiresAt.Before(time.Now()) {
				logger.Warn().Str("rpc", b.config.SelfRPCURL).Msg("failed to acquire liveliness lockout, is another sequencer running with this RPC URL?")
			}
		}
		selectedSeq := b.redis.selectSequencer(ctx)
		if selectedSeq == b.config.SelfRPCURL {
			if !holdingMutex {
				b.mutex.Lock()
				holdingMutex = true
			}
			if b.livelinessExpiresAt.After(time.Now()) {
				b.redis.acquireOrUpdateLockout(ctx, &b.lockoutExpiresAt)
			}
			if b.hasSequencerLockout() {
				if b.currentBatcher != b.sequencerBatcher {
					logger.Info().Str("rpc", b.config.SelfRPCURL).Msg("acquired sequencer lockout")
					targetSeqNum := b.redis.getLatestSeqNum(ctx)
					b.lastLockedSeqNum = targetSeqNum
					attemptCatchupUntil := b.lockoutExpiresAt.Add(-b.config.MaxLatency)
					for {
						currentSeqNum, err := b.core.GetMessageCount()
						if err != nil {
							logger.Warn().Err(err).Msg("error getting sequence number")
							select {
							case <-ctx.Done():
								return
							case <-time.After(5 * time.Second):
							}
						}
						if currentSeqNum.Cmp(targetSeqNum) >= 0 {
							logger.
								Info().
								Str("targetSeqNum", targetSeqNum.String()).
								Str("currentSeqNum", currentSeqNum.String()).
								Msg("caught up to previous sequencer position")
							break
						}
						if attemptCatchupUntil.After(time.Now()) {
							logger.
								Warn().
								Str("targetSeqNum", targetSeqNum.String()).
								Str("currentSeqNum", currentSeqNum.String()).
								Msg("failed to catch up to previous sequencer position")
							// There's a limited gap possible here as we checked it previously for liveliness
							// Therefore, we continue as the sequencer regardless, as such a gap is acceptable
							break
						}
						time.Sleep(500 * time.Millisecond)
					}
					if b.hasSequencerLockout() {
						err := b.sequencerBatcher.SequenceDelayedMessages(ctx, true)
						if err != nil {
							logger.Warn().Err(err).Msg("failed to sequence delayed messages after acquiring lockout")
						}
					}
				}
				b.currentBatcher = b.sequencerBatcher
				b.currentSeq = b.config.SelfRPCURL
				b.mutex.Unlock()
				holdingMutex = false
				seqNum, err := b.core.GetMessageCount()
				if err == nil {
					b.redis.updateLatestSeqNum(ctx, seqNum, b.lockoutExpiresAt)
					b.lastLockedSeqNum = seqNum
				} else {
					logger.Warn().Err(err).Msg("error getting sequence number")
				}
			}
		} else if b.currentSeq != selectedSeq {
			if b.currentBatcher == b.sequencerBatcher {
				b.inboxReader.MessageDeliveryMutex.Lock()
			}
			if !holdingMutex {
				b.mutex.Lock()
				holdingMutex = true
			}
			if b.currentBatcher == b.sequencerBatcher {
				logger.Info().Str("newPriority", selectedSeq).Msg("releasing sequencer lockout to make way for new sequencer")
				if b.hasSequencerLockout() {
					seqNum, err := b.core.GetMessageCount()
					if err == nil {
						b.redis.updateLatestSeqNum(ctx, seqNum, b.lockoutExpiresAt)
					} else {
						logger.Warn().Err(err).Msg("error getting sequence number")
					}
					b.redis.releaseLockout(ctx, &b.lockoutExpiresAt)
				}
				b.inboxReader.MessageDeliveryMutex.Unlock()
				b.currentBatcher = nil
			}
			if selectedSeq == "" {
				msg := "no prioritized sequencers online"
				logger.Warn().Msg(msg)
				b.currentBatcher = b.getErrorBatcher(errors.New(msg))
				b.currentSeq = selectedSeq
				b.mutex.Unlock()
				holdingMutex = false
			} else if b.redis.getLockout(ctx) == selectedSeq {
				logger.Info().Str("rpc", selectedSeq).Msg("forwarding to new sequencer")
				var err error
				b.currentBatcher, err = batcher.NewForwarder(ctx, configuration.Forwarder{Target: selectedSeq})
				if err == nil {
					b.currentSeq = selectedSeq
				} else {
					logger.Warn().Err(err).Msg("failed to connect to active sequencer")
					b.currentBatcher = b.getErrorBatcher(err)
				}
				// Note that we don't release the mutex if the selected sequencer doesn't have the lockout
				b.mutex.Unlock()
				holdingMutex = false
			}
		}
		refreshDelay := time.Millisecond * 500
		if b.hasSequencerLockout() {
			firstLockoutExpiresAt := b.lockoutExpiresAt
			if b.livelinessExpiresAt.Before(firstLockoutExpiresAt) {
				firstLockoutExpiresAt = b.livelinessExpiresAt
			}
			lockoutRefresh := time.Until(firstLockoutExpiresAt.Add(-b.config.MaxLatency))
			if lockoutRefresh > refreshDelay {
				refreshDelay = lockoutRefresh
			}
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(refreshDelay):
		}
	}
}

// Does not acquire mutex
func (b *LockoutBatcher) hasSequencerLockout() bool {
	return b.lockoutExpiresAt.After(time.Now())
}

func (b *LockoutBatcher) ShouldSequence() bool {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return b.currentBatcher == b.sequencerBatcher && b.hasSequencerLockout()
}

func (b *LockoutBatcher) getBatcher() batcher.TransactionBatcher {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	if b.currentBatcher == b.sequencerBatcher && !b.hasSequencerLockout() {
		return b.getErrorBatcher(errors.New("sequencer lockout expired"))
	}
	return b.currentBatcher
}

func (b *LockoutBatcher) PendingTransactionCount(ctx context.Context, account common.Address) *uint64 {
	return b.getBatcher().PendingTransactionCount(ctx, account)
}

func (b *LockoutBatcher) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return b.getBatcher().SendTransaction(ctx, tx)
}

func (b *LockoutBatcher) PendingSnapshot() (*snapshot.Snapshot, error) {
	return b.getBatcher().PendingSnapshot()
}

func (b *LockoutBatcher) SubscribeNewTxsEvent(ch chan<- ethcore.NewTxsEvent) event.Subscription {
	return b.sequencerBatcher.SubscribeNewTxsEvent(ch)
}

func (b *LockoutBatcher) Aggregator() *common.Address {
	return b.getBatcher().Aggregator()
}

func (b *LockoutBatcher) Start(ctx context.Context) {
	b.sequencerBatcher.Start(ctx)
}

type errorBatcher struct {
	err        error
	aggregator *common.Address
}

func (b *errorBatcher) PendingTransactionCount(ctx context.Context, account common.Address) *uint64 {
	return nil
}

func (b *errorBatcher) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return b.err
}

func (b *errorBatcher) PendingSnapshot() (*snapshot.Snapshot, error) {
	return nil, b.err
}

func (b *errorBatcher) SubscribeNewTxsEvent(ch chan<- ethcore.NewTxsEvent) event.Subscription {
	return nil
}

func (b *errorBatcher) Aggregator() *common.Address {
	return b.aggregator
}

func (b *errorBatcher) Start(ctx context.Context) {
}
