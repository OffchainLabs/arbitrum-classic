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

// these represent the fields from the Inbox message that we will broadcast
// for now we're just broadcasting the sequence number
type BroadcastInboxMessage struct {
	InboxSeqNum *big.Int `json:"seqnum"`
	// Sender      common.Address `json:"sender"`
	// GasPrice    *big.Int       `json:"gasprice"`
	// Data        []byte         `json:"data"`
}
type BroadcastMessage struct {
	Messages []*BroadcastInboxMessage `json:"messages"`
}
