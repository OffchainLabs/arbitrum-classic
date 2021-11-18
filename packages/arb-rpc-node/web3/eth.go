/*
 * Copyright 2020-2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package web3

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	arbcommon "github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

var gasPriceFactor = big.NewInt(2)
var gasEstimationCushion = 10

const maxGas = 1<<31 - 1

type ServerConfig struct {
	Mode          RpcMode
	MaxCallAVMGas uint64
	DevopsStubs   bool
}

type Server struct {
	srv                   *aggregator.Server
	ganacheMode           bool
	maxAVMGas             uint64
	aggregator            *arbcommon.Address
	sequencerInboxWatcher *ethbridge.SequencerInboxWatcher
}

const DefaultMaxAVMGas = 500000000

var DefaultConfig = ServerConfig{
	Mode:          NormalMode,
	MaxCallAVMGas: DefaultMaxAVMGas,
}

func NewServer(
	srv *aggregator.Server,
	config ServerConfig,
	sequencerInboxWatcher *ethbridge.SequencerInboxWatcher,
) *Server {
	maxGas := config.MaxCallAVMGas
	if maxGas == 0 {
		maxGas = math.MaxUint64
	}
	return &Server{
		srv:                   srv,
		ganacheMode:           config.Mode == GanacheMode,
		maxAVMGas:             maxGas,
		aggregator:            srv.Aggregator(),
		sequencerInboxWatcher: sequencerInboxWatcher,
	}
}

func (s *Server) ChainId() hexutil.Uint64 {
	return hexutil.Uint64(s.srv.ChainId().Uint64())
}

func (s *Server) GasPrice() (*hexutil.Big, error) {
	snap, err := s.srv.PendingSnapshot()
	if err != nil {
		return nil, err
	}
	prices, err := snap.GetPricesInWei()
	if err != nil {
		return nil, err
	}
	return (*hexutil.Big)(new(big.Int).Mul(prices[5], gasPriceFactor)), nil
}

func (s *Server) Accounts() []common.Address {
	return nil
}

func (s *Server) BlockNumber() (hexutil.Uint64, error) {
	blockCount, err := s.srv.GetBlockCount()
	if err != nil {
		return 0, err
	}
	if blockCount == 0 {
		return 0, errors.New("can't get block number because there are no blocks")
	}
	return hexutil.Uint64(blockCount - 1), nil
}

func (s *Server) GetBalance(address *common.Address, blockNum rpc.BlockNumberOrHash) (*hexutil.Big, error) {
	snap, err := s.getSnapshotForNumberOrHash(blockNum)
	if err != nil {
		return nil, err
	}
	balance, err := snap.GetBalance(arbcommon.NewAddressFromEth(*address))
	if err != nil {
		return nil, errors.Wrap(err, "error getting balance")
	}
	return (*hexutil.Big)(balance), nil
}

func (s *Server) GetStorageAt(address *common.Address, key string, blockNum rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	snap, err := s.getSnapshotForNumberOrHash(blockNum)
	if err != nil {
		return nil, err
	}
	index := new(big.Int).SetBytes(common.FromHex(key))
	storageVal, err := snap.GetStorageAt(arbcommon.NewAddressFromEth(*address), index)
	if err != nil {
		return nil, errors.Wrap(err, "error getting storage")
	}
	return math.U256Bytes(storageVal), nil
}

func (s *Server) getTransactionCountInner(ctx context.Context, address *common.Address, blockNum rpc.BlockNumberOrHash, forwardingOnlyMode bool) (hexutil.Uint64, error) {
	account := arbcommon.NewAddressFromEth(*address)

	if blockNum.BlockNumber != nil && *blockNum.BlockNumber == rpc.PendingBlockNumber {
		pending, err := s.srv.PendingTransactionCount(ctx, account)
		if err != nil {
			return 0, err
		}
		if pending != nil {
			return hexutil.Uint64(*pending), nil
		}
	}

	if forwardingOnlyMode {
		return 0, errors.New("only pending transaction count supported in forwarder only mode")
	}

	snap, err := s.getSnapshotForNumberOrHash(blockNum)
	if err != nil {
		return 0, err
	}
	txCount, err := snap.GetTransactionCount(account)
	if err != nil {
		return 0, errors.Wrap(err, "error getting transaction count")
	}

	return hexutil.Uint64(txCount.Uint64()), nil
}

func (s *Server) GetBlockTransactionCountByHash(blockHash common.Hash) (*hexutil.Big, error) {
	info, err := s.srv.BlockInfoByHash(arbcommon.NewHashFromEth(blockHash))
	if err != nil || info == nil {
		return nil, err
	}
	return s.getBlockTransactionCount(info)
}

func (s *Server) GetBlockTransactionCountByNumber(blockNum *rpc.BlockNumber) (*hexutil.Big, error) {
	height, err := s.srv.BlockNum(blockNum)
	if err != nil {
		return nil, err
	}
	info, err := s.srv.BlockInfoByNumber(height)
	if err != nil || info == nil {
		return nil, err
	}
	return s.getBlockTransactionCount(info)
}

func (s *Server) GetCode(address *common.Address, blockNum rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	if *address == arbos.ARB_NODE_INTERFACE_ADDRESS {
		// Fake code to make the contract appear real
		return hexutil.Bytes{1}, nil
	}
	snap, err := s.getSnapshotForNumberOrHash(blockNum)
	if err != nil {
		return nil, err
	}
	code, err := snap.GetCode(arbcommon.NewAddressFromEth(*address))
	if err != nil {
		return nil, errors.Wrap(err, "error getting code")
	}
	return code, nil
}

func (s *Server) Call(callArgs CallTxArgs, blockNum rpc.BlockNumberOrHash, overrides *map[common.Address]EthCallOverride) (hexutil.Bytes, error) {
	if callArgs.To != nil && *callArgs.To == arbos.ARB_NODE_INTERFACE_ADDRESS {
		var data []byte
		if callArgs.Data != nil {
			data = *callArgs.Data
		}
		return HandleNodeInterfaceCall(s, data, blockNum)
	}

	snap, err := s.getSnapshotForNumberOrHash(blockNum)
	if err != nil {
		return nil, err
	}
	if snap.ArbosVersion() >= 42 && (callArgs.GasPrice == nil || callArgs.GasPrice.ToInt().Sign() <= 0) {
		callArgs.GasPrice = (*hexutil.Big)(big.NewInt(1 << 60))
	}

	if overrides != nil {
		for address, override := range *overrides {
			account := arbcommon.NewAddressFromEth(address)
			if override.Nonce != nil {
				err := snap.SetNonce(account, uint64(*override.Nonce))
				if err != nil {
					return nil, err
				}
			}
			if override.Balance != nil {
				err := snap.SetBalance(account, override.Balance.ToInt())
				if err != nil {
					return nil, err
				}
			}
			if override.Code != nil {
				err := snap.SetCode(account, *override.Code)
				if err != nil {
					return nil, err
				}
			}
			if override.State != nil {
				storage := make(map[arbcommon.Hash]arbcommon.Hash)
				for key, val := range *override.State {
					storage[arbcommon.NewHashFromEth(key)] = arbcommon.NewHashFromEth(val)
				}
				err := snap.SetState(account, storage)
				if err != nil {
					return nil, err
				}
			}
			if override.StateDiff != nil {
				for key, val := range *override.StateDiff {
					err := snap.Store(account, arbcommon.NewHashFromEth(key), arbcommon.NewHashFromEth(val))
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}

	from, msg := buildCallMsg(callArgs)

	res, _, err := snap.Call(msg, from, s.maxAVMGas)
	if err != nil {
		return nil, err
	}
	if res.ResultCode != evm.ReturnCode {
		return nil, evm.HandleCallError(res, s.ganacheMode)
	}
	return res.ReturnData, nil
}

func (s *Server) EstimateGas(args CallTxArgs) (hexutil.Uint64, error) {
	if args.To != nil && *args.To == arbos.ARB_NODE_INTERFACE_ADDRESS {
		// Fake gas for call
		return hexutil.Uint64(21000), nil
	}
	blockNum := rpc.PendingBlockNumber
	snap, err := s.getSnapshot(&blockNum)
	if err != nil {
		return 0, err
	}
	if snap.ArbosVersion() >= 42 && (args.GasPrice == nil || args.GasPrice.ToInt().Sign() <= 0) {
		args.GasPrice = (*hexutil.Big)(big.NewInt(1 << 60))
	}
	from, tx := buildTransactionForEstimation(args)
	var agg arbcommon.Address
	if args.Aggregator != nil {
		agg = arbcommon.NewAddressFromEth(*args.Aggregator)
	} else if s.aggregator != nil {
		agg = *s.aggregator
	}
	res, _, err := snap.EstimateGas(tx, agg, from, s.maxAVMGas)
	if err == nil && res.ResultCode != evm.ReturnCode {
		err = evm.HandleCallError(res, s.ganacheMode)
	}
	if err != nil {
		logging := log.Warn()
		if args.Gas != nil {
			logging = logging.Uint64("gaslimit", uint64(*args.Gas))
		}
		if args.GasPrice != nil {
			logging = logging.Str("gasPrice", args.GasPrice.String())
		}
		if args.Value != nil {
			logging = logging.Str("value", args.Value.String())
		}
		if args.To != nil {
			logging = logging.Str("to", args.To.Hex())
		}
		if args.From != nil {
			logging = logging.Str("from", args.From.Hex())
		}
		if args.Data != nil {
			logging = logging.Hex("data", *args.Data)
		}
		logging.Err(err).Msg("error estimating gas")
		return 0, err
	}

	if res.FeeStats.Price.L2Computation.Cmp(big.NewInt(0)) == 0 {
		return hexutil.Uint64(res.GasUsed.Uint64() + 10000), nil
	} else {
		extraCalldataUnits := (len(res.FeeStats.GasUsed().Bytes()) + len(new(big.Int).Mul(res.FeeStats.Price.L2Computation, gasPriceFactor).Bytes()) + gasEstimationCushion) * 16
		// Adjust calldata units used for calldata from gas limit
		res.FeeStats.UnitsUsed.L1Calldata = res.FeeStats.UnitsUsed.L1Calldata.Add(res.FeeStats.UnitsUsed.L1Calldata, big.NewInt(int64(extraCalldataUnits)))
		used := res.FeeStats.TargetGasUsed()
		used = used.Mul(used, big.NewInt(11))
		used = used.Div(used, big.NewInt(10))
		return hexutil.Uint64(used.Uint64() + 100), nil
	}
}

func (s *Server) GetBlockByHash(blockHashRaw hexutil.Bytes, includeTxData bool) (*GetBlockResult, error) {
	var blockHash arbcommon.Hash
	copy(blockHash[:], blockHashRaw)
	info, err := s.srv.BlockInfoByHash(blockHash)
	if err != nil || info == nil {
		return nil, err
	}
	return s.getBlock(info, includeTxData)
}

func (s *Server) GetBlockByNumber(blockNum *rpc.BlockNumber, includeTxData bool) (*GetBlockResult, error) {
	height, err := s.srv.BlockNum(blockNum)
	if err != nil {
		return nil, err
	}
	info, err := s.srv.BlockInfoByNumber(height)
	if err != nil || info == nil {
		return nil, err
	}
	return s.getBlock(info, includeTxData)
}

func (s *Server) getTransactionInfoByHash(txHash hexutil.Bytes) (*evm.TxResult, *machine.BlockInfo, core.InboxState, error) {
	var requestId arbcommon.Hash
	copy(requestId[:], txHash)
	res, inbox, err := s.srv.GetRequestResult(requestId)
	if err != nil || res == nil {
		return nil, nil, core.InboxState{}, err
	}
	info, err := s.srv.BlockInfoByNumber(res.IncomingRequest.L2BlockNumber.Uint64())
	if err != nil || info == nil {
		return nil, nil, core.InboxState{}, err
	}
	return res, info, inbox, nil
}

func (s *Server) GetTransactionByHash(txHash hexutil.Bytes) (*TransactionResult, error) {
	res, info, _, err := s.getTransactionInfoByHash(txHash)
	if err != nil || res == nil {
		return nil, err
	}
	tx, err := evm.GetTransaction(res)
	if err != nil {
		return nil, err
	}

	var blockHash *common.Hash
	if info != nil {
		h := info.Header.Hash()
		blockHash = &h
	}

	return makeTransactionResult(tx, blockHash), nil
}

func (s *Server) GetTransactionByBlockHashAndIndex(blockHash common.Hash, index hexutil.Uint64) (*TransactionResult, error) {
	info, err := s.srv.BlockInfoByHash(arbcommon.NewHashFromEth(blockHash))
	if err != nil || info == nil {
		return nil, err
	}
	return s.getTransactionByBlockAndIndex(info, index)
}

func (s *Server) GetTransactionByBlockNumberAndIndex(blockNum *rpc.BlockNumber, index hexutil.Uint64) (*TransactionResult, error) {
	height, err := s.srv.BlockNum(blockNum)
	if err != nil {
		return nil, err
	}
	info, err := s.srv.BlockInfoByNumber(height)
	if err != nil || info == nil {
		return nil, err
	}

	return s.getTransactionByBlockAndIndex(info, index)
}

func (s *Server) GetTransactionReceipt(ctx context.Context, txHash hexutil.Bytes, opts *ArbGetTxReceiptOpts) (*GetTransactionReceiptResult, error) {
	res, info, inboxState, err := s.getTransactionInfoByHash(txHash)
	if err != nil || res == nil {
		return nil, err
	}

	receipt := res.ToEthReceipt(arbcommon.NewHashFromEth(info.Header.Hash()))

	tx, err := evm.GetTransaction(res)
	if err != nil {
		return nil, err
	}

	var contractAddress *common.Address
	emptyAddress := common.Address{}
	if receipt.ContractAddress != emptyAddress {
		contractAddress = &receipt.ContractAddress
	}

	var l1InboxBatchInfo *L1InboxBatchInfo
	if opts != nil && opts.ReturnL1InboxBatchInfo {
		if s.sequencerInboxWatcher == nil {
			return nil, errors.New("RPC L1 lookups disabled")
		}
		lookup := s.srv.GetLookup()
		seqNum := new(big.Int).Sub(inboxState.Count, big.NewInt(1))
		batch, err := s.sequencerInboxWatcher.LookupBatchContaining(ctx, lookup, seqNum)
		if err != nil {
			return nil, err
		}
		if batch != nil {
			if batch.GetAfterCount().Cmp(inboxState.Count) < 0 {
				return nil, errors.New("retrieved too early sequencer batch")
			}
			expectedTxAcc, expectedBatchAcc, err := lookup.GetInboxAccPair(seqNum, new(big.Int).Sub(batch.GetAfterCount(), big.NewInt(1)))
			if err != nil {
				return nil, err
			}
			if expectedTxAcc != inboxState.Accumulator || expectedBatchAcc != batch.GetAfterAcc() {
				return nil, errors.New("inconsistent sequencer inbox state")
			}
			currentBlockHeight, err := s.sequencerInboxWatcher.CurrentBlockHeight(ctx)
			if err != nil {
				return nil, err
			}
			rawLog := batch.GetRawLog()
			blockNum := new(big.Int).SetUint64(rawLog.BlockNumber)
			confirmations := new(big.Int).Sub(currentBlockHeight, blockNum)
			if confirmations.Sign() >= 0 {
				l1InboxBatchInfo = &L1InboxBatchInfo{
					Confirmations: (*hexutil.Big)(confirmations),
					BlockNumber:   (*hexutil.Big)(blockNum),
					LogAddress:    rawLog.Address,
					LogTopics:     rawLog.Topics,
					LogData:       rawLog.Data,
				}
			}
		}
	}

	return &GetTransactionReceiptResult{
		TransactionHash:   receipt.TxHash,
		TransactionIndex:  hexutil.Uint64(receipt.TransactionIndex),
		BlockHash:         receipt.BlockHash,
		BlockNumber:       (*hexutil.Big)(receipt.BlockNumber),
		From:              res.IncomingRequest.Sender.ToEthAddress(),
		To:                tx.Tx.To(),
		CumulativeGasUsed: hexutil.Uint64(receipt.CumulativeGasUsed),
		GasUsed:           hexutil.Uint64(res.CalcGasUsed().Uint64()),
		ContractAddress:   contractAddress,
		Logs:              receipt.Logs,
		LogsBloom:         receipt.Bloom.Bytes(),
		Status:            hexutil.Uint64(receipt.Status),

		ReturnCode: hexutil.Uint64(res.ResultCode),
		ReturnData: res.ReturnData,
		FeeStats: &FeeStatsResult{
			Prices:    feeSetToFeeSetResult(res.FeeStats.Price),
			UnitsUsed: feeSetToFeeSetResult(res.FeeStats.UnitsUsed),
			Paid:      feeSetToFeeSetResult(res.FeeStats.Paid),
		},
		L1BlockNumber:    (*hexutil.Big)(res.IncomingRequest.L1BlockNumber),
		L1InboxBatchInfo: l1InboxBatchInfo,
	}, nil
}

func feeSetToFeeSetResult(feeset *evm.FeeSet) *FeeSetResult {
	return &FeeSetResult{
		L1Transaction: (*hexutil.Big)(feeset.L1Transaction),
		L1Calldata:    (*hexutil.Big)(feeset.L1Calldata),
		L2Storage:     (*hexutil.Big)(feeset.L2Storage),
		L2Computation: (*hexutil.Big)(feeset.L2Computation),
	}
}

func (s *Server) getBlockTransactionCount(block *machine.BlockInfo) (*hexutil.Big, error) {
	info, err := s.srv.BlockLogFromInfo(block)
	if err != nil || info == nil {
		return nil, err
	}
	return (*hexutil.Big)(info.BlockStats.TxCount), nil
}

func (s *Server) getTransactionByBlockAndIndex(block *machine.BlockInfo, index hexutil.Uint64) (*TransactionResult, error) {
	txRes, err := s.srv.GetTxInBlockAtIndexResults(block, uint64(index))
	if err != nil {
		return nil, err
	}
	tx, err := evm.GetTransaction(txRes)
	if err != nil {
		return nil, err
	}
	blockHash := block.Header.Hash()
	return makeTransactionResult(tx, &blockHash), nil
}

func (s *Server) getBlock(block *machine.BlockInfo, includeTxData bool) (*GetBlockResult, error) {
	l2Block, results, err := s.srv.GetMachineBlockResults(block)
	if err != nil || results == nil {
		return nil, err
	}

	processedTxes := evm.FilterEthTxResults(results)

	var transactions interface{}
	if includeTxData {
		blockHash := block.Header.Hash()
		txResults := make([]*TransactionResult, 0, len(processedTxes))
		for _, res := range processedTxes {
			txResults = append(txResults, makeTransactionResult(res, &blockHash))
		}
		transactions = txResults
	} else {
		txHashes := make([]hexutil.Bytes, 0, len(processedTxes))
		for _, res := range processedTxes {
			txHashes = append(txHashes, res.Result.IncomingRequest.MessageID.Bytes())
		}
		transactions = txHashes
	}

	return makeBlockResult(l2Block, block.Header, transactions), nil
}

func makeBlockResult(blockLog *evm.BlockInfo, header *types.Header, transactions interface{}) *GetBlockResult {
	size := uint64(0)
	uncles := make([]hexutil.Bytes, 0)
	return &GetBlockResult{
		Number:           (*hexutil.Big)(header.Number),
		Hash:             header.Hash().Bytes(),
		ParentHash:       header.ParentHash.Bytes(),
		MixDigest:        header.MixDigest.Bytes(),
		Nonce:            &header.Nonce,
		Sha3Uncles:       header.UncleHash.Bytes(),
		LogsBloom:        header.Bloom.Bytes(),
		TransactionsRoot: header.TxHash.Bytes(),
		StateRoot:        header.Root.Bytes(),
		ReceiptsRoot:     header.ReceiptHash.Bytes(),
		Miner:            header.Coinbase.Bytes(),
		Difficulty:       (*hexutil.Big)(header.Difficulty),
		TotalDifficulty:  (*hexutil.Big)(header.Difficulty),
		ExtraData:        (*hexutil.Bytes)(&header.Extra),
		Size:             (*hexutil.Uint64)(&size),
		GasLimit:         (*hexutil.Uint64)(&header.GasLimit),
		GasUsed:          (*hexutil.Uint64)(&header.GasUsed),
		Timestamp:        (*hexutil.Uint64)(&header.Time),
		Transactions:     transactions,
		Uncles:           &uncles,

		L1BlockNumber: (*hexutil.Big)(blockLog.L1BlockNum),
	}
}

func makeTransactionResult(processedTx *evm.ProcessedTx, blockHash *common.Hash) *TransactionResult {
	tx := processedTx.Tx
	res := processedTx.Result
	vVal, rVal, sVal := tx.RawSignatureValues()
	txIndex := res.TxIndex.Uint64()
	blockNum := res.IncomingRequest.L2BlockNumber

	provenance := res.IncomingRequest.Provenance
	var parentRequestId *common.Hash
	emptyParent := arbcommon.Hash{}
	if provenance.ParentRequestId != emptyParent {
		h := provenance.ParentRequestId.ToEthHash()
		parentRequestId = &h
	}

	var l2Subtype *hexutil.Uint64
	if processedTx.L2Subtype != nil {
		st := hexutil.Uint64(*processedTx.L2Subtype)
		l2Subtype = &st
	}

	return &TransactionResult{
		BlockHash:        blockHash,
		BlockNumber:      (*hexutil.Big)(blockNum),
		From:             res.IncomingRequest.Sender.ToEthAddress(),
		Gas:              hexutil.Uint64(tx.Gas()),
		GasPrice:         (*hexutil.Big)(tx.GasPrice()),
		Hash:             res.IncomingRequest.MessageID.ToEthHash(),
		Input:            tx.Data(),
		Nonce:            hexutil.Uint64(tx.Nonce()),
		To:               tx.To(),
		TransactionIndex: (*hexutil.Uint64)(&txIndex),
		Value:            (*hexutil.Big)(tx.Value()),
		V:                (*hexutil.Big)(vVal),
		R:                (*hexutil.Big)(rVal),
		S:                (*hexutil.Big)(sVal),

		L1SeqNum:        (*hexutil.Big)(provenance.L1SeqNum),
		ParentRequestId: parentRequestId,
		IndexInParent:   (*hexutil.Big)(provenance.IndexInParent),
		ArbType:         hexutil.Uint64(processedTx.Kind),
		ArbSubType:      l2Subtype,
		L1BlockNumber:   (*hexutil.Big)(res.IncomingRequest.L1BlockNumber),
	}
}

func buildTransactionForEstimation(args CallTxArgs) (arbcommon.Address, *types.Transaction) {
	gas := uint64(0)
	if args.Gas != nil {
		gas = uint64(*args.Gas)
	}
	return buildTransactionImpl(args, gas)
}

func buildTransactionImpl(args CallTxArgs, gas uint64) (arbcommon.Address, *types.Transaction) {
	var from arbcommon.Address
	if args.From != nil {
		from = arbcommon.NewAddressFromEth(*args.From)
	}
	gasPrice := big.NewInt(0)
	if args.GasPrice != nil {
		gasPrice = args.GasPrice.ToInt()
	}
	value := big.NewInt(0)
	if args.Value != nil {
		value = args.Value.ToInt()
	}
	var data []byte
	if args.Data != nil {
		data = *args.Data
	}

	return from, types.NewTx(&types.LegacyTx{
		Nonce:    0,
		GasPrice: gasPrice,
		Gas:      gas,
		To:       args.To,
		Value:    value,
		Data:     data,
	})
}

func buildCallMsg(args CallTxArgs) (arbcommon.Address, message.ContractTransaction) {
	gas := uint64(0)
	if args.Gas != nil {
		gas = uint64(*args.Gas)
	}
	if gas == 0 || gas > maxGas {
		gas = maxGas
	}
	from, tx := buildTransactionImpl(args, gas)
	var dest arbcommon.Address
	if tx.To() != nil {
		dest = arbcommon.NewAddressFromEth(*tx.To())
	}
	return from, message.ContractTransaction{
		BasicTx: message.BasicTx{
			MaxGas:      new(big.Int).SetUint64(tx.Gas()),
			GasPriceBid: tx.GasPrice(),
			DestAddress: dest,
			Payment:     tx.Value(),
			Data:        tx.Data(),
		},
	}
}

func (s *Server) getSnapshot(blockNum *rpc.BlockNumber) (*snapshot.Snapshot, error) {
	if blockNum == nil || *blockNum == rpc.PendingBlockNumber {
		pending, err := s.srv.PendingSnapshot()
		if err != nil {
			return nil, err
		}
		if pending != nil {
			return pending, nil
		}
		// If pending isn't available, we can fall back to latest
		latest := rpc.LatestBlockNumber
		blockNum = &latest
	}

	if *blockNum == rpc.LatestBlockNumber {
		snap, err := s.srv.LatestSnapshot()
		if err != nil {
			return nil, err
		}
		if snap == nil {
			return nil, errors.New("couldn't fetch latest snapshot")
		}
		return snap, nil
	}

	snap, err := s.srv.GetSnapshot(uint64(*blockNum))
	if err != nil {
		return nil, err
	}
	if snap == nil {
		return nil, errors.Errorf("unsupported block number %v", uint64(*blockNum))
	}
	return snap, nil
}

func (s *Server) getSnapshotForNumberOrHash(blockNum rpc.BlockNumberOrHash) (*snapshot.Snapshot, error) {
	if blockNum.BlockNumber != nil {
		return s.getSnapshot(blockNum.BlockNumber)
	}
	if blockNum.BlockHash == nil {
		return nil, errors.New("must specify block number or hash")
	}
	var blockHash arbcommon.Hash
	copy(blockHash[:], blockNum.BlockHash[:])
	info, err := s.srv.BlockInfoByHash(blockHash)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, errors.New("block with hash not found")
	}

	snap, err := s.srv.GetSnapshot(info.Header.Number.Uint64())
	if err != nil {
		return nil, err
	}
	if snap == nil {
		return nil, errors.Errorf("unsupported block number %v", info.Header.Number.Uint64())
	}
	return snap, nil
}
