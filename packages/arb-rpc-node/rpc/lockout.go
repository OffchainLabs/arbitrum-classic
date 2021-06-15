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
	"os"
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
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcastclient"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
)

var logger = log.With().Caller().Stack().Str("component", "rpc").Logger()

type LockoutBatcher struct {
	mutex            sync.RWMutex
	sequencerBatcher *batcher.SequencerBatcher
	lockoutExpiresAt time.Time
	currentSeq       string
	core             core.ArbOutputLookup
	inboxReader      *monitor.InboxReader
	inboxFeed        chan broadcaster.BroadcastFeedMessage
	broadcastClient  *broadcastclient.BroadcastClient

	currentBatcher batcher.TransactionBatcher
}

func SetupLockout(
	ctx context.Context,
	seqBatcher batcher.TransactionBatcher,
	redisURL string,
	inboxFeed chan broadcaster.BroadcastFeedMessage,
) (*LockoutBatcher, error) {
	newBatcher := &LockoutBatcher{
		sequencerBatcher: seqBatcher.(*batcher.SequencerBatcher),
		currentBatcher: &errorBatcher{
			err: errors.New("sequencer lockout manager starting up"),
		},
		inboxFeed: inboxFeed,
	}
	newBatcher.sequencerBatcher.LockoutManager = newBatcher
	go newBatcher.lockoutManager(ctx)
	return newBatcher, nil
}

const BROADCAST_FEED_PREFIX string = "ws://"
const BROADCAST_FEED_POSTFIX string = ":9545/"
const RPC_URL_PREFIX string = "http://"
const RPC_URL_POSTFIX string = ":8545/rpc"

func (b *LockoutBatcher) lockoutManager(ctx context.Context) {
	holdingMutex := false
	defer (func() {
		if !holdingMutex {
			b.mutex.Lock()
		}
		b.currentBatcher = &errorBatcher{
			err: errors.New("sequencer lockout manager shutting down"),
		}
		b.mutex.Unlock()
		if ctx.Err() == nil {
			// We aren't shutting down but the lockout manager died
			logger.Error().Msg("lockout manager died but context isn't shutting down")
			os.Exit(1)
		}
	})()
	for {
		b.redis.updateLiveliness(ctx)
		otherSeq := b.redis.selectSequencer(ctx)
		if otherSeq == "" {
			b.mutex.Lock()
			holdingMutex = true
			// Leave a margin of 10 seconds between our expected and real expiry
			b.lockoutExpiresAt = b.redis.acquireLockout(ctx).Add(time.Second * -10)
			if b.currentSeq != "" {
				targetSeqNum := b.redis.getLatestSeqNum(ctx)
				for b.lockoutExpiresAt.Before(time.Now()) {
					currentSeqNum, err := b.core.GetMessageCount()
					if err != nil {
						logger.Warn().Err(err).Msg("error getting sequence number")
						time.Sleep(5 * time.Second)
						continue
					}
					if currentSeqNum.Cmp(targetSeqNum) >= 0 {
						break
					}
					time.Sleep(500 * time.Millisecond)
				}
			}
			b.currentBatcher = b.sequencerBatcher
			b.currentSeq = ""
			b.mutex.Unlock()
			holdingMutex = false
		} else if b.currentSeq != otherSeq {
			if b.currentBatcher == b.sequencerBatcher {
				b.inboxReader.MessageDeliveryMutex.Lock()
			}
			b.mutex.Lock()
			holdingMutex = true
			if b.currentBatcher == b.sequencerBatcher {
				if b.lockoutExpiresAt.After(time.Now()) {
					seqNum, err := b.core.GetMessageCount()
					if err == nil {
						b.redis.updateLatestSeqNum(ctx, seqNum)
					} else {
						logger.Warn().Err(err).Msg("error getting sequence number")
					}
					b.redis.releaseLockout(ctx)
					b.lockoutExpiresAt = time.Time{}
				}
				b.inboxReader.MessageDeliveryMutex.Unlock()
			}
			var err error
			b.currentBatcher, err = batcher.NewForwarder(ctx, RPC_URL_PREFIX+otherSeq+RPC_URL_POSTFIX)
			if err == nil {
				if b.broadcastClient != nil {
					b.broadcastClient.Close()
				}
				b.broadcastClient = broadcastclient.NewBroadcastClient(BROADCAST_FEED_PREFIX+otherSeq+BROADCAST_FEED_POSTFIX, nil, 20*time.Second)
				err = b.broadcastClient.ConnectWithChannel(ctx, b.inboxFeed)
				if err == nil {
					b.currentSeq = otherSeq
				} else {
					logger.Warn().Err(err).
						Msg("failed connect to sequencer broadcast")
				}
			} else {
				logger.Warn().Err(err).Msg("failed to connect to current sequencer")
				b.currentBatcher = &errorBatcher{err: err}
			}
			b.mutex.Unlock()
			holdingMutex = false
		}
	}
}

func (b *LockoutBatcher) ShouldSequence() bool {
	b.mutex.RLock()
	b.mutex.RUnlock()
	return b.currentBatcher == b.sequencerBatcher && b.lockoutExpiresAt.Before(time.Now())
}

func (b *LockoutBatcher) getBatcher() batcher.TransactionBatcher {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	if b.currentBatcher == b.sequencerBatcher && b.lockoutExpiresAt.After(time.Now()) {
		return &errorBatcher{
			err: errors.New("sequencer lockout expired"),
		}
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
	return b.getBatcher().SubscribeNewTxsEvent(ch)
}

func (b *LockoutBatcher) Aggregator() *common.Address {
	return b.getBatcher().Aggregator()
}

func (b *LockoutBatcher) Start(ctx context.Context) {
	b.sequencerBatcher.Start(ctx)
}

type errorBatcher struct {
	err error
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
	return nil
}

func (b *errorBatcher) Start(ctx context.Context) {
}
