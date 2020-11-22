package evm

import (
	"errors"
	"github.com/ethereum/go-ethereum/core/types"
	"log"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

type ProcessedTx struct {
	Result    *TxResult
	Tx        *types.Transaction
	Kind      inbox.Type
	L2Subtype *message.L2SubType
}

func GetTransaction(res *TxResult) (*ProcessedTx, error) {
	msg := res.IncomingRequest
	// Special handling for buddy deploy
	if msg.Kind == message.L2BuddyDeploy {
		buddyDeployMessage := message.NewBuddyDeploymentFromData(msg.Data)
		return &ProcessedTx{
			Result: res,
			Tx:     buddyDeployMessage.AsEthTx(),
			Kind:   msg.Kind,
		}, nil
	}

	if msg.Kind != message.L2Type {
		return nil, errors.New("result is not a transaction")
	}
	l2msg, err := message.L2Message{Data: msg.Data}.AbstractMessage()
	if err != nil {
		return nil, err
	}
	ethMsg, ok := l2msg.(message.EthConvertable)
	if !ok {
		return nil, errors.New("message not convertible to receipt")
	}
	l2Type := l2msg.L2Type()
	return &ProcessedTx{
		Result:    res,
		Tx:        ethMsg.AsEthTx(),
		Kind:      msg.Kind,
		L2Subtype: &l2Type,
	}, nil
}

func FilterEthTxResults(results []*TxResult) []*ProcessedTx {
	filteredResults := make([]*ProcessedTx, 0, len(results))
	for _, res := range results {
		kind := res.IncomingRequest.Kind
		// Ignore other message types
		if kind != message.L2Type && kind != message.L2BuddyDeploy {
			continue
		}
		processed, err := GetTransaction(res)
		if err != nil {
			log.Println("Couldn't return transaction for request", res.IncomingRequest.MessageID)
			continue
		}
		filteredResults = append(filteredResults, processed)
	}
	return filteredResults
}
