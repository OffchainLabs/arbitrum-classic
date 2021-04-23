package broadcaster

import (
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
)

type BroadcastInboxMessage struct {
	FeedItem  monitor.SequencerFeedItem `json:"feedItem"`
	Signature *big.Int                  `json:"signature"`
}
type BroadcastMessage struct {
	Messages []*BroadcastInboxMessage `json:"messages"`
}
