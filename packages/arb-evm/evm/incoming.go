package evm

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/pkg/errors"
)

type ProcessedTx struct {
	Result    *TxResult
	Tx        *types.Transaction
	Kind      inbox.Type
	L2Subtype *message.L2SubType
}

func GetTransaction(res *TxResult) (*ProcessedTx, error) {
	msg := res.IncomingRequest
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
		if kind != message.L2Type {
			continue
		}
		processed, err := GetTransaction(res)
		if err != nil {
			logger.Info().
				Stack().
				Err(err).
				Hex("request", res.IncomingRequest.MessageID.Bytes()).
				Msg("Couldn't return transaction for request")
			continue
		}
		filteredResults = append(filteredResults, processed)
	}
	return filteredResults
}
