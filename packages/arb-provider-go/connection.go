package goarbitrum

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"sync"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/validatorserver"
)

var ARB_SYS_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000064")
var ARB_INFO_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000065")

type ArbConnection struct {
	proxy       ValidatorProxy
	vmId        common.Address
	globalInbox arbbridge.GlobalInbox
	sequenceNum *big.Int
}

func Dial(url string, auth *bind.TransactOpts, ethclint *ethclient.Client) (*ArbConnection, error) {
	client := ethbridge.NewEthAuthClient(ethclint, auth)
	proxy := NewValidatorProxyImpl(url)
	vmIdStr, err := proxy.GetVMInfo()
	if err != nil {
		return nil, err
	}
	vmId := common.HexToAddress(vmIdStr)
	rollup, err := client.NewRollupWatcher(vmId)
	if err != nil {
		return nil, err
	}
	inboxAddr, err := rollup.InboxAddress(context.Background())
	if err != nil {
		return nil, err
	}
	globalInbox, err := client.NewGlobalInbox(inboxAddr)
	if err != nil {
		return nil, err
	}
	return &ArbConnection{proxy: proxy, vmId: vmId, globalInbox: globalInbox}, nil
}

func (conn *ArbConnection) getInfoCon() (*ArbInfo, error) {
	return NewArbInfo(ARB_INFO_ADDRESS, conn)
}

func (conn *ArbConnection) getSysCon() (*ArbSys, error) {
	return NewArbSys(ARB_SYS_ADDRESS, conn)
}

// PendingNonceAt retrieves the current pending nonce associated with an account.
func (conn *ArbConnection) getCurrentNonce(
	ctx context.Context,
	account ethcommon.Address,
) (*big.Int, error) {
	if conn.sequenceNum != nil {
		return conn.sequenceNum, nil
	}
	sysConn, err := conn.getSysCon()
	if err != nil {
		return nil, err
	}
	num, err := sysConn.GetTransactionCount(&bind.CallOpts{Pending: true}, account)
	if err != nil {
		return nil, err
	}
	conn.sequenceNum = num
	return num, nil
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
	contract ethcommon.Address,
	blockNumber *big.Int,
) ([]byte, error) {
	infoCon, err := conn.getInfoCon()
	if err != nil {
		return nil, err
	}
	return infoCon.GetCode(&bind.CallOpts{
		BlockNumber: blockNumber,
	}, contract)
}

func (conn *ArbConnection) call(
	ctx context.Context,
	call ethereum.CallMsg,
	blockNumber *big.Int,
) ([]byte, error) {
	retValue, err := conn.proxy.CallMessage(*call.To, call.From, call.Data)
	if err != nil {
		return nil, err
	}

	logVal, err := evm.ProcessLog(retValue, conn.vmId)
	if err != nil {
		return nil, err
	}
	switch logVal := logVal.(type) {
	case evm.Return:
		return logVal.ReturnVal, nil
	case evm.Stop:
		return []byte{}, nil
	case evm.Revert:
		return nil, fmt.Errorf("call reverted with result %v", string(logVal.ReturnVal))
	default:
		return nil, fmt.Errorf("call reverted")
	}
}

// CallContract executes an Ethereum contract call with the specified data as the
// input.
func (conn *ArbConnection) CallContract(
	ctx context.Context,
	call ethereum.CallMsg,
	blockNumber *big.Int,
) ([]byte, error) {
	return conn.call(ctx, call, blockNumber)
}

///////////////////////////////////////////////////////////////////////////////
// Methods of ContractTransactor

// PendingCodeAt returns the code of the given account in the pending state.
func (conn *ArbConnection) PendingCodeAt(
	ctx context.Context,
	account ethcommon.Address,
) ([]byte, error) {
	infoCon, err := conn.getInfoCon()
	if err != nil {
		return nil, err
	}
	return infoCon.GetCode(&bind.CallOpts{
		Pending: true,
	}, account)
}

// PendingCallContract executes an Ethereum contract call against the pending state.
func (conn *ArbConnection) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	return conn.call(ctx, call, nil)
}

// PendingNonceAt retrieves the current pending nonce associated with an account.
func (conn *ArbConnection) PendingNonceAt(
	ctx context.Context,
	account ethcommon.Address,
) (uint64, error) {
	sysConn, err := conn.getSysCon()
	if err != nil {
		return 0, err
	}
	num, err := sysConn.GetTransactionCount(&bind.CallOpts{Pending: true}, account)
	if err != nil {
		return 0, err
	}
	return num.Uint64(), nil
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
	return conn.globalInbox.SendTransactionMessage(ctx, tx.Data(), conn.vmId, common.NewAddressFromEth(*tx.To()), tx.Value(), new(big.Int).SetUint64(tx.Nonce()))
}

///////////////////////////////////////////////////////////////////////////////
// Methods of ContractFilterer

// FilterLogs executes a log filter operation, blocking during execution and
// returning all the results in one batch.
//
// TODO(karalabe): Deprecate when the subscription one can return past data too.
func (conn *ArbConnection) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	var ret []types.Log
	address, topics := _extractAddrTopics(query)
	logInfos, err := conn.proxy.FindLogs(0, math.MaxInt32, address[:], topics)
	if err != nil {
		return nil, err
	}
	for _, logInfo := range logInfos {
		outs, err := _decodeLogInfo(logInfo)
		if err != nil {
			return nil, err
		}
		ok := true
		for i, targetTopic := range topics {
			if targetTopic != outs.Topics[i] {
				ok = false
			}
		}
		if ok {
			ret = append(ret, *outs)
		}
	}
	return ret, nil
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
	logChan          chan<- types.Log
	errChan          chan error
	address          common.Address
	topics           [][32]byte
	unsubOnce        *sync.Once
	closeChan        chan interface{}
	wg               sync.WaitGroup
}

func _extractAddrTopics(query ethereum.FilterQuery) (addr common.Address, topics [][32]byte) {
	if len(query.Addresses) > 1 {
		panic("GoArbitrum: subscription can't handle more than one contract address")
	}
	addr = common.NewAddressFromEth(query.Addresses[0])

	topics = make([][32]byte, len(query.Topics))
	for i, sl := range query.Topics {
		if len(sl) > 1 {
			panic("GoArbitrum: subscription can't handle ORs of topics")
		}
		copy(topics[i][:], sl[0][:])
	}
	return
}

func _decodeLogInfo(ins *validatorserver.LogInfo) (*types.Log, error) {
	outs := &types.Log{}
	addr, err := hexutil.Decode(ins.Address)
	if err != nil {
		log.Println("_decodeLogInfo error 1:", err)
		return nil, err
	}
	copy(outs.Address[:], addr)
	outs.Topics = make([]ethcommon.Hash, len(ins.Topics))
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
		ch,
		make(chan error, 1),
		address,
		topics,
		&sync.Once{},
		make(chan interface{}),
		sync.WaitGroup{},
	}
	sub.wg.Add(1)
	go func() {
		defer sub.wg.Done()
		defer sub.Unsubscribe()
		ticker := time.NewTicker(subscriptionPollingInterval)
		defer ticker.Stop()
		for {
			select {
			case <-sub.closeChan:
				return
			case <-ticker.C:
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
		}
	}()
	return sub
}

// Unsubscribe cancels the sending of events to the data channel
// and closes the error channel.
func (sub *subscription) Unsubscribe() {
	sub.unsubOnce.Do(func() {
		close(sub.closeChan)
		sub.wg.Wait()
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

///////////////////////////////////////////////////////////////////////////////
// Methods of Deploy Backend
// CodeAt is implemented above

// TransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (conn *ArbConnection) TransactionReceipt(ctx context.Context, txHash ethcommon.Hash) (*types.Receipt, error) {
	result, ok, err := conn.proxy.GetMessageResult(txHash.Bytes())
	if err != nil {
		log.Println("TransactionReceipt error:", err)
		return nil, err
	} else if !ok {
		return nil, ethereum.NotFound
	}

	processed, err := evm.ProcessLog(result, conn.vmId)
	if err != nil {
		log.Println("TransactionReceipt ProcessLog error:", err)
		return nil, err
	}

	status := uint64(0)
	var logs []evm.Log
	switch res := processed.(type) {
	case evm.Return:
		status = 1
		logs = res.Logs
	case evm.Stop:
		status = 1
		logs = res.Logs
	default:
		// Transaction unsuccessful
	}

	ethMsg := processed.GetEthMsg()

	var evmLogs []*types.Log
	if logs != nil {
		for i, l := range logs {
			addressBytes := l.Address.ToBytes()

			evmParsedTopics := make([]ethcommon.Hash, len(l.Topics))
			for j, t := range l.Topics {
				evmParsedTopics[j] = ethcommon.BytesToHash(t[:])
			}

			evmLogs = append(evmLogs, &types.Log{
				Address:     ethcommon.BytesToAddress(addressBytes[12:]),
				Topics:      evmParsedTopics,
				Data:        l.Data,
				BlockNumber: ethMsg.BlockNumber.Uint64(),
				TxHash:      txHash,
				TxIndex:     0,
				BlockHash:   txHash,
				Index:       uint(i),
				Removed:     false,
			})
		}
	}

	return &types.Receipt{
		PostState:         []byte{0},
		Status:            status,
		CumulativeGasUsed: 1,
		Bloom:             types.BytesToBloom([]byte{0}),
		Logs:              evmLogs,
		TxHash:            txHash,
		ContractAddress:   ethcommon.BytesToAddress([]byte{0}),
		GasUsed:           1,
		BlockHash:         txHash,
		BlockNumber:       ethMsg.BlockNumber,
		TransactionIndex:  0,
	}, nil
}

func (conn *ArbConnection) TxToMessage(tx *types.Transaction, from common.Address) message.Transaction {
	return message.Transaction{
		Chain:       conn.vmId,
		To:          common.NewAddressFromEth(*tx.To()),
		From:        from,
		SequenceNum: new(big.Int).SetUint64(tx.Nonce()),
		Value:       tx.Value(),
		Data:        tx.Data(),
	}
}

func (conn *ArbConnection) TxHash(tx *types.Transaction, from common.Address) common.Hash {
	return conn.TxToMessage(tx, from).ReceiptHash()
}
