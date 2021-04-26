package broadcaster

import (
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
)

type BroadcastInboxMessage struct {
	FeedItem  monitor.SequencerFeedItem `json:"feedItem"`
	Signature []byte                    `json:"signature"`
}
type BroadcastMessage struct {
	Messages []*BroadcastInboxMessage `json:"messages"`
}
