package evm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
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

	if msg.Kind == message.RetryableType {
		retryable := message.NewRetryableTxFromData(msg.Data)
		txData := arbos.CreateRetryableTicketData(retryable)
		createTicketTx := &types.LegacyTx{
			Nonce:    0,
			GasPrice: big.NewInt(0),
			Gas:      0,
			To:       &arbos.ARB_RETRYABLE_ADDRESS,
			Value:    retryable.Deposit,
			Data:     txData,
		}
		return &ProcessedTx{
			Result: res,
			Tx:     types.NewTx(createTicketTx),
			Kind:   msg.Kind,
		}, nil
	}

	if msg.Kind != message.L2Type && msg.Kind != message.EthDepositTxType {
		return nil, errors.Errorf("result is not a transaction %v", msg.Kind)
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

	tx := ethMsg.AsEthTx()
	if res.ResultCode == ReturnCode && ethMsg.AsEthTx().To() == nil && len(res.ReturnData) == 0 {
		ethMsg, ok := ethMsg.(message.ContractTransaction)
		if ok {
			// If we're in a successful retryable send to the 0 address, it wasn't treated as contract construction
			tx = ethMsg.AsNonConstructorTx()
		}
	}
	return &ProcessedTx{
		Result:    res,
		Tx:        tx,
		Kind:      msg.Kind,
		L2Subtype: &l2Type,
	}, nil
}

func FilterEthTxResults(results []*TxResult) []*ProcessedTx {
	filteredResults := make([]*ProcessedTx, 0, len(results))
	for _, res := range results {
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
