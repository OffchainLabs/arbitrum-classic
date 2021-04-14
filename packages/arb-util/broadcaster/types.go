package broadcaster

import (
	"math/big"
)

// Object represents generic message parameters.
// In real-world application it is better to avoid such types for better
// performance.
type Object map[string]interface{}

type Request struct {
	ID     int    `json:"id"`
	Method string `json:"method"`
	Params Object `json:"params"`
}

type Response struct {
	ID     int    `json:"id"`
	Result Object `json:"result"`
}

type Error struct {
	ID    int    `json:"id"`
	Error Object `json:"error"`
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
