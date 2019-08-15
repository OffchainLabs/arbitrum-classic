package goarbitrum

import (
	"context"
	"errors"
	"log"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/coordinator"
)

type ArbConnection struct {
	proxy      ValidatorProxy
	vmId       []byte
	privateKey []byte
	hexPubkey  string
}

func Dial(url string, privateKey []byte, hexPubkey string) (*ArbConnection, error) {
	proxy := NewValidatorProxyImpl(url)
	vmIdStr, err := proxy.GetVMInfo()
	if err != nil {
		return nil, err
	}
	vmId, err := hexutil.Decode(vmIdStr)
	if err != nil {
		return nil, err
	}
	return &ArbConnection{proxy, vmId, append([]byte{}, privateKey...), hexPubkey}, nil
}

func _nyiError(funcname string) error {
	return errors.New("goarbitrum error: " + funcname + " not yet implemented")
}

///////////////////////////////////////////////////////////////////////////////
// Methods of ContractCaller

// CodeAt returns the code of the given account. This is needed to differentiate
// between contract internal errors and the local chain being out of sync.
func (conn *ArbConnection) CodeAt(
	ctx context.Context,
	contract common.Address,
	blockNumber *big.Int,
) ([]byte, error) {
	return nil, _nyiError("CodeAt")
}

// CallContract executes an Ethereum contract call with the specified data as the
// input.
func (conn *ArbConnection) CallContract(
	ctx context.Context,
	call ethereum.CallMsg,
	blockNumber *big.Int,
) ([]byte, error) {
	dataValue, err := evm.BytesToSizedByteArray(call.Data)
	if err != nil {
		return nil, err
	}
	destAddrValue := value.NewIntValue(new(big.Int).SetBytes(call.To[:]))
	seqNumValue := value.NewIntValue(new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(2)))
	arbCallValue, err := value.NewTupleFromSlice([]value.Value{dataValue, destAddrValue, seqNumValue})
	if err != nil {
		panic("Unexpected error building arbCallValue")
	}
	retValue, err := conn.proxy.CallMessage(arbCallValue, call.From)
	if err != nil {
		return nil, err
	}

	logVal, err := evm.ProcessLog(retValue)
	if err != nil {
		return nil, err
	}
	switch logVal := logVal.(type) {
	case evm.Return:
		return logVal.ReturnVal, nil
	case evm.Stop:
		return []byte{}, nil
	default:
		return nil, errors.New("call reverted")
	}
}

///////////////////////////////////////////////////////////////////////////////
// Methods of ContractTransactor

// PendingCodeAt returns the code of the given account in the pending state.
func (conn *ArbConnection) PendingCodeAt(
	ctx context.Context,
	account common.Address,
) ([]byte, error) {
	return nil, _nyiError("PendingCodeAt")
}

// PendingNonceAt retrieves the current pending nonce associated with an account.
func (conn *ArbConnection) PendingNonceAt(
	ctx context.Context,
	account common.Address,
) (uint64, error) {
	return 0, nil
}

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
func (conn *ArbConnection) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}

// EstimateGas tries to estimate the gas needed to execute a specific
// transaction based on the current pending state of the backend blockchain.
// There is no guarantee that this is the true gas limit requirement as other
// transactions may be added or removed by miners, but it should provide a basis
// for setting a reasonable default.

// EstimateGas tries to estimate the gas needed to execute a specific
// transaction based on the current pending state of the backend blockchain.
// There is no guarantee that this is the true gas limit requirement as other
// transactions may be added or removed by miners, but it should provide a basis
// for setting a reasonable default.
func (conn *ArbConnection) EstimateGas(
	ctx context.Context,
	call ethereum.CallMsg,
) (gas uint64, err error) {
	return 100000, nil
}

// SendTransaction injects the transaction into the pending pool for execution.
func (conn *ArbConnection) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	dataValue, err := evm.BytesToSizedByteArray(tx.Data())
	if err != nil {
		log.Println("Error converting to SizedByteArray")
		return err
	}
	destAddrValue := value.NewIntValue(new(big.Int).SetBytes(tx.To()[:]))
	seqNumValue := value.NewIntValue(new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(2)))

	arbCallValue, err := value.NewTupleFromSlice([]value.Value{dataValue, destAddrValue, seqNumValue})
	if err != nil {
		panic("Unexpected error building arbCallValue")
	}

	tokenType := [21]byte{}
	messageHash := solsha3.SoliditySHA3(
		solsha3.Bytes32(conn.vmId),
		solsha3.Bytes32(arbCallValue.Hash()),
		solsha3.Uint256(big.NewInt(0)), // amount
		tokenType[:],
	)
	signedMsg := solsha3.SoliditySHA3WithPrefix(solsha3.Bytes32(messageHash))
	sig, err := secp256k1.Sign(signedMsg, conn.privateKey)
	if err != nil {
		return err
	}

	txHash, err := conn.proxy.SendMessage(arbCallValue, conn.hexPubkey, sig)
	if err != nil {
		log.Println("SendTransaction: error returned from proxy.SendMessage:", err)
		return err
	}

	return func() error {
		for {
			resultVal, ok, err := conn.proxy.GetMessageResult(txHash)
			if err != nil {
				log.Println("GetMessageResult error:", err)
				return err
			}
			if !ok {
				time.Sleep(2 * time.Second)
			} else {
				result, err := evm.ProcessLog(resultVal)
				if err != nil {
					log.Println("GetMessageResultLog error:", err)
					return err
				}
				switch res := result.(type) {
				case evm.Revert:
					log.Println("call reverted:", string(res.ReturnVal))
				default:
					// do nothing
				}
				return nil
			}
		}
	}()
}

///////////////////////////////////////////////////////////////////////////////
// Methods of ContractFilterer

// FilterLogs executes a log filter operation, blocking during execution and
// returning all the results in one batch.
//
// TODO(karalabe): Deprecate when the subscription one can return past data too.
func (conn *ArbConnection) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return nil, _nyiError("FilterLogs")
}

// SubscribeFilterLogs creates a background log filtering operation, returning
// a subscription immediately, which can be used to stream the found events.
func (conn *ArbConnection) SubscribeFilterLogs(
	ctx context.Context,
	query ethereum.FilterQuery,
	ch chan<- types.Log,
) (ethereum.Subscription, error) {
	return newSubscription(conn, query, ch), nil
}

const subscriptionPollingInterval = 5 * time.Second

type subscription struct {
	proxy            ValidatorProxy
	firstBlockUnseen uint64
	active           bool
	logChan          chan<- types.Log
	errChan          chan error
	address          common.Address
	topics           [][32]byte
	unsubOnce        *sync.Once
}

func _extractAddrTopics(query ethereum.FilterQuery) (addr common.Address, topics [][32]byte) {
	if len(query.Addresses) > 1 {
		panic("GoArbitrum: subscription can't handle more than one contract address")
	}
	addr = query.Addresses[0]

	topics = make([][32]byte, len(query.Topics))
	for i, sl := range query.Topics {
		if len(sl) > 1 {
			panic("GoArbitrum: subscription can't handle ORs of topics")
		}
		copy(topics[i][:], sl[0][:])
	}
	return
}

func _decodeLogInfo(ins *coordinator.LogInfo) (*types.Log, error) {
	outs := &types.Log{}
	addr, err := hexutil.Decode(ins.Address)
	if err != nil {
		log.Println("_decodeLogInfo error 1:", err)
		return nil, err
	}
	copy(outs.Address[:], addr)
	outs.Topics = make([]common.Hash, len(ins.Topics))
	for i, top := range ins.Topics {
		decodedTopic, err := hexutil.Decode(top)
		if err != nil {
			log.Println("_decodeLogInfo error 2:", err)
			return nil, err
		}
		copy(outs.Topics[i][:], decodedTopic)
	}
	outs.Data, err = hexutil.Decode(ins.Data)
	if err != nil {
		log.Println("_decodeLogInfo error 3:", err)
		return nil, err
	}
	outs.BlockNumber, err = hexutil.DecodeUint64(ins.BlockNumber)
	if err != nil {
		log.Println("_decodeLogInfo error 4:", err)
		return nil, err
	}
	hh, err := hexutil.Decode(ins.Address)
	if err != nil {
		log.Println("_decodeLogInfo error 5:", err)
		return nil, err
	}
	copy(outs.TxHash[:], hh)
	txi64, err := hexutil.DecodeUint64(ins.TransactionIndex)
	if err != nil {
		log.Println("_decodeLogInfo error 6:", err)
		log.Println("value was", ins.TransactionIndex)
		return nil, err
	}
	outs.TxIndex = uint(txi64)
	hh, err = hexutil.Decode(ins.BlockHash)
	if err != nil {
		log.Println("_decodeLogInfo error 7:", err)
		return nil, err
	}
	copy(outs.BlockHash[:], hh)
	iui, err := hexutil.DecodeUint64(ins.LogIndex)
	if err != nil {
		log.Println("_decodeLogInfo error 8:", err)
		return nil, err
	}
	outs.Index = uint(iui)
	outs.Removed = false
	return outs, nil
}

func newSubscription(conn *ArbConnection, query ethereum.FilterQuery, ch chan<- types.Log) *subscription {
	address, topics := _extractAddrTopics(query)
	sub := &subscription{
		conn.proxy,
		0,
		true,
		ch,
		make(chan error, 1),
		address,
		topics,
		&sync.Once{},
	}
	go func() {
		defer sub.Unsubscribe()
		for {
			time.Sleep(subscriptionPollingInterval)
			if !sub.active {
				return
			}
			logInfos, err := sub.proxy.FindLogs(int64(sub.firstBlockUnseen), math.MaxInt32, sub.address[:], sub.topics)
			if err != nil {
				sub.errChan <- err
				return
			}
			for _, logInfo := range logInfos {
				outs, err := _decodeLogInfo(logInfo)
				if err != nil {
					sub.errChan <- err
					return
				}
				ok := true
				for i, targetTopic := range topics {
					if targetTopic != outs.Topics[i] {
						ok = false
					}
				}
				if outs.BlockNumber < sub.firstBlockUnseen {
					ok = false
				}
				if ok {
					sub.logChan <- *outs
					if sub.firstBlockUnseen <= outs.BlockNumber {
						sub.firstBlockUnseen = outs.BlockNumber + 1
					}
				}
			}
		}
	}()
	return sub
}

// Unsubscribe cancels the sending of events to the data channel
// and closes the error channel.
func (sub *subscription) Unsubscribe() {
	sub.unsubOnce.Do(func() {
		sub.active = false
		close(sub.errChan)
	})
}

// Err returns the subscription error channel. The error channel receives
// a value if there is an issue with the subscription (e.g. the network connection
// delivering the events has been closed). Only one value will ever be sent.
// The error channel is closed by Unsubscribe.
func (sub *subscription) Err() <-chan error {
	return sub.errChan
}
