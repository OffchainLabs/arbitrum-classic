package goarbitrum

import (
	"context"
	"errors"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	errors2 "github.com/pkg/errors"
	"math/big"
	"sync"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

var ARB_SYS_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000064")
var ARB_INFO_ADDRESS = ethcommon.HexToAddress("0x0000000000000000000000000000000000000065")

type ArbConnection struct {
	proxy       ValidatorProxy
	vmId        common.Address
	globalInbox arbbridge.GlobalInbox
	sequenceNum *big.Int
}

func Dial(url string, auth *bind.TransactOpts, ethclint ethutils.EthClient) (*ArbConnection, error) {
	client := ethbridge.NewEthAuthClient(ethclint, auth)
	proxy := NewValidatorProxyImpl(url)
	vmIdStr, err := proxy.GetChainAddress(context.Background())
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
	globalInbox, err := client.NewGlobalInbox(inboxAddr, vmId)
	if err != nil {
		return nil, err
	}
	return &ArbConnection{proxy: proxy, vmId: vmId, globalInbox: globalInbox}, nil
}

func (conn *ArbConnection) getInfoCon() (*arboscontracts.ArbInfo, error) {
	return arboscontracts.NewArbInfo(ARB_INFO_ADDRESS, conn)
}

func (conn *ArbConnection) getSysCon() (*arboscontracts.ArbSys, error) {
	return arboscontracts.NewArbSys(ARB_SYS_ADDRESS, conn)
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
	num, err := sysConn.GetTransactionCount(
		&bind.CallOpts{Context: ctx, Pending: true},
		account,
	)
	if err != nil {
		return nil, err
	}
	conn.sequenceNum = num
	return num, nil
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
	code, err := infoCon.GetCode(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: blockNumber,
	}, contract)
	if err != nil {
		return nil, errors2.Wrap(err, "couldn't get code")
	}
	return code, nil
}

func processCallRet(retValue value.Value) ([]byte, error) {
	logVal, err := evm.NewResultFromValue(retValue)
	if err != nil {
		return nil, err
	}
	if logVal.ResultCode != evm.ReturnCode {
		return nil, fmt.Errorf("call reverted %v", logVal)
	}
	return logVal.ReturnData, nil
}

// CallContract executes an Ethereum contract call with the specified data as the
// input.
func (conn *ArbConnection) CallContract(
	ctx context.Context,
	call ethereum.CallMsg,
	blockNumber *big.Int,
) ([]byte, error) {
	var dest common.Address
	if call.To != nil {
		dest = common.NewAddressFromEth(*call.To)
	}
	gasPriceBid := big.NewInt(0)
	if call.GasPrice != nil {
		gasPriceBid = call.GasPrice
	}
	tx := message.Call{
		MaxGas:      new(big.Int).SetUint64(call.Gas),
		GasPriceBid: gasPriceBid,
		DestAddress: dest,
		Data:        call.Data,
	}
	retValue, err := conn.proxy.CallMessage(ctx, tx, call.From)
	if err != nil {
		return nil, err
	}
	return processCallRet(retValue)
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
	code, err := infoCon.GetCode(&bind.CallOpts{
		Context: ctx,
		Pending: true,
	}, account)
	if err != nil {
		return nil, errors2.Wrap(err, "couldn't get pending code")
	}
	return code, nil
}

// PendingCallContract executes an Ethereum contract call against the pending state.
func (conn *ArbConnection) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	var dest common.Address
	if call.To != nil {
		dest = common.NewAddressFromEth(*call.To)
	}
	gasPriceBid := big.NewInt(0)
	if call.GasPrice != nil {
		gasPriceBid = call.GasPrice
	}
	tx := message.Call{
		MaxGas:      new(big.Int).SetUint64(call.Gas),
		GasPriceBid: gasPriceBid,
		DestAddress: dest,
		Data:        call.Data,
	}
	retValue, err := conn.proxy.PendingCall(ctx, tx, call.From)
	if err != nil {
		return nil, err
	}
	return processCallRet(retValue)
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
	num, err := sysConn.GetTransactionCount(&bind.CallOpts{Context: ctx, Pending: true}, account)
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
func (conn *ArbConnection) EstimateGas(
	ctx context.Context,
	call ethereum.CallMsg,
) (gas uint64, err error) {
	return 100000000, nil
}

// SendTransaction injects the transaction into the pending pool for execution.
func (conn *ArbConnection) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	arbTx := message.NewTransactionFromEthTx(tx)
	_, err := conn.globalInbox.SendL2Message(ctx, conn.vmId, message.L2Message{Msg: arbTx})
	return err
}

///////////////////////////////////////////////////////////////////////////////
// Methods of ContractFilterer

// FilterLogs executes a log filter operation, blocking during execution and
// returning all the results in one batch.

// TODO: Currently FilterLogs does not properly handle reorgs by replaying undone
// logs with the removed flag set
func (conn *ArbConnection) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	logInfos, err := conn.proxy.FindLogs(
		ctx,
		_extractQueryHeight(query.FromBlock),
		_extractQueryHeight(query.ToBlock),
		query.Addresses,
		query.Topics,
	)
	if err != nil {
		return nil, err
	}
	logs := make([]types.Log, 0, len(logInfos))
	for _, l := range logInfos {
		logs = append(logs, *l.ToEVMLog())
	}
	return logs, nil
}

// SubscribeFilterLogs creates a background log filtering operation, returning
// a subscription immediately, which can be used to stream the found events.
func (conn *ArbConnection) SubscribeFilterLogs(
	ctx context.Context,
	query ethereum.FilterQuery,
	ch chan<- types.Log,
) (ethereum.Subscription, error) {
	return newSubscription(ctx, conn, query, ch), nil
}

const subscriptionPollingInterval = 5 * time.Second

type subscription struct {
	logChan   chan<- types.Log
	errChan   chan error
	unsubOnce *sync.Once
	closeChan chan interface{}
	wg        sync.WaitGroup
}

func _extractQueryHeight(val *big.Int) *uint64 {
	var ret *uint64
	if val != nil {
		intVal := val.Uint64()
		ret = &intVal
	}
	return ret
}

func newSubscription(ctx context.Context, conn *ArbConnection, query ethereum.FilterQuery, ch chan<- types.Log) *subscription {
	// We will assume that FromBlock is always non-nil
	if query.FromBlock == nil {
		query.FromBlock = big.NewInt(0)
	}
	sub := &subscription{
		ch,
		make(chan error, 1),
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
			case <-ctx.Done():
				return
			case <-ticker.C:
				logInfos, err := conn.proxy.FindLogs(
					ctx,
					_extractQueryHeight(query.FromBlock),
					_extractQueryHeight(query.ToBlock),
					query.Addresses,
					query.Topics,
				)
				if err != nil {
					sub.errChan <- err
					return
				}
				for _, l := range logInfos {
					sub.logChan <- *l.ToEVMLog()
				}
				// We will always receive logs in order by block and always
				// receive all logs from a given block in a single query.
				// Therefore the next query can start with block height one
				// greater than the last log in the previous query
				if len(logInfos) > 0 {
					query.FromBlock = new(big.Int).Add(logInfos[len(logInfos)-1].Block.Height.AsInt(), big.NewInt(1))
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
	index, startLogIndex, val, err := conn.proxy.GetRequestResult(ctx, common.NewHashFromEth(txHash))
	if err != nil {
		return nil, errors2.Wrap(err, "TransactionReceipt error:")
	}
	if val == nil {
		return nil, ethereum.NotFound
	}
	result, err := evm.NewResultFromValue(val)
	if err != nil {
		return nil, err
	}

	if result.L1Message.MessageID().ToEthHash() != txHash {
		return nil, errors.New("tx hash doesn't match")
	}

	status := uint64(0)
	if result.ResultCode == evm.ReturnCode {
		status = 1
	}

	blockInfo, err := conn.proxy.BlockInfo(ctx, result.L1Message.ChainTime.BlockNum.AsInt().Uint64())
	if err != nil {
		return nil, err
	}
	txIndex := index - blockInfo.StartLog
	var evmLogs []*types.Log
	logIndex := startLogIndex
	for _, l := range result.EVMLogs {
		ethLog := &types.Log{
			Address:     l.Address.ToEthAddress(),
			Topics:      common.NewEthHashesFromHashes(l.Topics),
			Data:        l.Data,
			BlockNumber: result.L1Message.ChainTime.BlockNum.AsInt().Uint64(),
			TxHash:      txHash,
			TxIndex:     uint(txIndex),
			BlockHash:   blockInfo.Hash.ToEthHash(),
			Index:       uint(logIndex),
		}
		logIndex++
		evmLogs = append(evmLogs, ethLog)
	}

	contractAddress := ethcommon.Address{}
	if result.L1Message.Kind == message.L2Type {
		msg, err := message.NewL2MessageFromData(result.L1Message.Data)
		if err == nil {
			if msg, ok := msg.(message.Transaction); ok {
				emptyAddress := common.Address{}
				if msg.DestAddress == emptyAddress {
					copy(contractAddress[:], result.ReturnData[12:])
				}
			}
		}
	}

	return &types.Receipt{
		PostState: []byte{0},
		Status:    status,
		// TODO: Fill in with real value
		CumulativeGasUsed: 1,
		Bloom:             types.BytesToBloom(types.LogsBloom(evmLogs).Bytes()),
		Logs:              evmLogs,
		TxHash:            txHash,
		ContractAddress:   contractAddress,
		GasUsed:           result.GasUsed.Uint64(),
		BlockHash:         blockInfo.Hash.ToEthHash(),
		BlockNumber:       result.L1Message.ChainTime.BlockNum.AsInt(),
		TransactionIndex:  uint(txIndex),
	}, nil
}

func (conn *ArbConnection) TxHash(tx *types.Transaction) (common.Hash, error) {
	from, err := types.HomesteadSigner{}.Sender(tx)
	if err != nil {
		return common.Hash{}, err
	}
	return message.NewTransactionFromEthTx(tx).MessageID(common.NewAddressFromEth(from), conn.vmId), nil
}
