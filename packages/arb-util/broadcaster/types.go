package broadcaster

import (
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type ConfirmedAccumulator struct {
	IsConfirmed bool        `json:"isConfirmed"`
	Accumulator common.Hash `json:"accumulator"`
}

type BroadcastInboxMessage struct {
	FeedItem  monitor.SequencerFeedItem `json:"feedItem"`
	Signature []byte                    `json:"signature"`
}

type BroadcastMessage struct {
	Version              int                      `json:"version"`
	Messages             []*BroadcastInboxMessage `json:"messages"`
	ConfirmedAccumulator ConfirmedAccumulator     `json:"confirmedAccumulator"`
}
