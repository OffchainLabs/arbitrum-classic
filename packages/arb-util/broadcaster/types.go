package broadcaster

import (
	"math/big"
)

type BroadcastInboxMessage struct {
	BeforeAccumulator *big.Int `json:"beforeAccumulator"`
	InboxMessage      []byte   `json:"inboxMessage"`
	Signature         *big.Int `json:"signature"`
	SeqNum            *big.Int `json:"seqnum"`
}
type BroadcastMessage struct {
	Messages []*BroadcastInboxMessage `json:"messages"`
}
