package evm

import (
	"errors"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

type ProcessedTx struct {
	Tx        *types.Transaction
	Kind      inbox.Type
	L2Subtype *message.L2SubType
}

func GetTransaction(msg IncomingRequest) (*ProcessedTx, error) {
	// Special handling for buddy deploy
	if msg.Kind == message.L2BuddyDeploy {
		buddyDeployMessage := message.NewBuddyDeploymentFromData(msg.Data)
		return &ProcessedTx{
			Tx:   buddyDeployMessage.AsEthTx(),
			Kind: msg.Kind,
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
		Tx:        ethMsg.AsEthTx(),
		Kind:      msg.Kind,
		L2Subtype: &l2Type,
	}, nil
}
