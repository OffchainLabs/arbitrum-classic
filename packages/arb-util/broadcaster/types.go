package broadcaster

import (
	"math/big"
)

type Object map[string]interface{}

type Request struct {
	ID     int    `json:"id"`
	Method string `json:"method"`
	Params Object `json:"params"`
}
type PongResponse struct {
	Time string `json:"time"`
}

type BroadcastInboxMessage struct {
	BeforeAccumulator *big.Int `json:"beforeAccumulator"`
	InboxMessage      []byte   `json:"inboxMessage"`
	Signature         *big.Int `json:"signature"`
	SeqNum            *big.Int `json:"seqnum"`
}
type BroadcastMessage struct {
	Messages []*BroadcastInboxMessage `json:"messages"`
}
