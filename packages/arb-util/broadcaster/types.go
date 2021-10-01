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

package broadcaster

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

type ConfirmedAccumulator struct {
	IsConfirmed bool        `json:"isConfirmed"`
	Accumulator common.Hash `json:"accumulator"`
}

type BroadcastFeedMessage struct {
	FeedItem  SequencerFeedItem `json:"feedItem"`
	Signature []byte            `json:"signature"`
}

type SequencerFeedItem struct {
	BatchItem inbox.SequencerBatchItem `json:"batchItem"`
	PrevAcc   common.Hash              `json:"prevAcc"`
}

type BroadcastMessage struct {
	Version              int                     `json:"version"`
	Messages             []*BroadcastFeedMessage `json:"messages"`
	ConfirmedAccumulator ConfirmedAccumulator    `json:"confirmedAccumulator"`
}
