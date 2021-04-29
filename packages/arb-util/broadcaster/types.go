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
