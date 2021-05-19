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
	"fmt"
	"math/big"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	arbcommon "github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

var logger = log.With().Caller().Stack().Str("component", "web3").Logger()
var gasPriceFactor = big.NewInt(2)
var gasEstimationCushion = 10

type Server struct {
	srv         *aggregator.Server
	ganacheMode bool
	maxCallGas  uint64
	maxAVMGas   uint64
	aggregator  *arbcommon.Address
}

func NewServer(
	srv *aggregator.Server,
	ganacheMode bool,
) *Server {
	return &Server{
		srv:         srv,
		ganacheMode: ganacheMode,
		maxCallGas:  1<<31 - 1,
		maxAVMGas:   500000000,
		aggregator:  srv.Aggregator(),
	}
}

func (s *Server) ChainId() hexutil.Uint64 {
	return hexutil.Uint64(message.ChainAddressToID(
		arbcommon.NewAddressFromEth(s.srv.GetChainAddress()),
	).Uint64())
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

func (s *Server) GetBalance(address *common.Address, blockNum *rpc.BlockNumber) (*hexutil.Big, error) {
	snap, err := s.getSnapshot(blockNum)
	if err != nil {
		return nil, err
	}
	balance, err := snap.GetBalance(arbcommon.NewAddressFromEth(*address))
	if err != nil {
		return nil, errors.Wrap(err, "error getting balance")
	}
	return (*hexutil.Big)(balance), nil
}

func (s *Server) GetStorageAt(address *common.Address, index *hexutil.Big, blockNum *rpc.BlockNumber) (*hexutil.Big, error) {
	snap, err := s.getSnapshot(blockNum)
	if err != nil {
		return nil, err
	}
	storageVal, err := snap.GetStorageAt(arbcommon.NewAddressFromEth(*address), (*big.Int)(index))
	if err != nil {
		return nil, errors.Wrap(err, "error getting storage")
	}
	return (*hexutil.Big)(storageVal), nil
}

func (s *Server) GetTransactionCount(ctx context.Context, address *common.Address, blockNum *rpc.BlockNumber) (hexutil.Uint64, error) {
	account := arbcommon.NewAddressFromEth(*address)
	snap, err := s.getSnapshot(blockNum)
	if err != nil {
		return 0, err
	}
	txCount, err := snap.GetTransactionCount(account)
	if err != nil {
		return 0, errors.Wrap(err, "error getting transaction count")
	}

	count := txCount.Uint64()
	if blockNum == nil || *blockNum == rpc.PendingBlockNumber {
		pending := s.srv.PendingTransactionCount(ctx, account)
		if pending != nil {
			if *pending > count {
				count = *pending
			}
		}
	}

	return hexutil.Uint64(count), nil
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

func (s *Server) GetCode(address *common.Address, blockNum *rpc.BlockNumber) (hexutil.Bytes, error) {
	if *address == arbos.ARB_NODE_INTERFACE_ADDRESS {
		// Fake code to make the contract appear real
		return hexutil.Bytes{1}, nil
	}
	snap, err := s.getSnapshot(blockNum)
	if err != nil {
		return nil, err
	}
	code, err := snap.GetCode(arbcommon.NewAddressFromEth(*address))
	if err != nil {
		return nil, errors.Wrap(err, "error getting code")
	}
	return code, nil
}

func (s *Server) SendRawTransaction(ctx context.Context, data hexutil.Bytes) (hexutil.Bytes, error) {
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(data, tx); err != nil {
		return nil, err
	}
	err := s.srv.SendTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}
	return tx.Hash().Bytes(), nil
}

func (s *Server) Call(callArgs CallTxArgs, blockNum *rpc.BlockNumber) (hexutil.Bytes, error) {
	if callArgs.To != nil && *callArgs.To == arbos.ARB_NODE_INTERFACE_ADDRESS {
		var data []byte
		if callArgs.Data != nil {
			data = *callArgs.Data
		}
		return HandleNodeInterfaceCall(s.srv, data)
	}

	snap, err := s.getSnapshot(blockNum)
	if err != nil {
		return nil, err
	}
	from, msg := buildCallMsg(callArgs, s.maxCallGas)

	res, err := snap.Call(msg, from)
	res, err = handleCallResult(res, err, blockNum)
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
	from, tx := buildTransactionForEstimation(args)
	var agg arbcommon.Address
	if args.Aggregator != nil {
		agg = arbcommon.NewAddressFromEth(*args.Aggregator)
	} else if s.aggregator != nil {
		agg = *s.aggregator
	}
	res, err := snap.EstimateGas(tx, agg, from, s.maxAVMGas)
	res, err = handleCallResult(res, err, &blockNum)
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
	if res.ResultCode != evm.ReturnCode {
		return 0, evm.HandleCallError(res, s.ganacheMode)
	}

	if res.FeeStats.Price.L2Computation.Cmp(big.NewInt(0)) == 0 {
		return hexutil.Uint64(res.GasUsed.Uint64() + 10000), nil
	} else {
		extraCalldataUnits := (len(res.CalcGasUsed().Bytes()) + len(new(big.Int).Mul(res.FeeStats.Price.L2Computation, gasPriceFactor).Bytes()) + gasEstimationCushion) * 16
		// Adjust calldata units used for calldata from gas limit
		res.FeeStats.UnitsUsed.L1Calldata = res.FeeStats.UnitsUsed.L1Calldata.Add(res.FeeStats.UnitsUsed.L1Calldata, big.NewInt(int64(extraCalldataUnits)))
		fmt.Println("Estimate", res.CalcGasUsed())
		fmt.Println("stats", res.FeeStats.PayTarget())
		fmt.Println("stats", res.FeeStats.Paid)
		used := new(big.Int).Div(res.FeeStats.Paid.Total(), res.FeeStats.Price.L2Computation)
		return hexutil.Uint64(used.Uint64() + 1000), nil
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

func (s *Server) getTransactionInfoByHash(txHash hexutil.Bytes) (*evm.TxResult, *machine.BlockInfo, error) {
	var requestId arbcommon.Hash
	copy(requestId[:], txHash)
	res, err := s.srv.GetRequestResult(requestId)
	if err != nil || res == nil {
		return nil, nil, err
	}
	info, err := s.srv.BlockInfoByNumber(res.IncomingRequest.L2BlockNumber.Uint64())
	if err != nil || info == nil {
		return nil, nil, err
	}
	return res, info, nil
}

func (s *Server) GetTransactionByHash(txHash hexutil.Bytes) (*TransactionResult, error) {
	res, info, err := s.getTransactionInfoByHash(txHash)
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

func (s *Server) GetTransactionReceipt(txHash hexutil.Bytes) (*GetTransactionReceiptResult, error) {
	res, info, err := s.getTransactionInfoByHash(txHash)
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
		L1BlockNumber: (*hexutil.Big)(res.IncomingRequest.L1BlockNumber),
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

func buildTransactionForCall(args CallTxArgs, maxGas uint64) (arbcommon.Address, *types.Transaction) {
	gas := uint64(0)
	if args.Gas != nil {
		gas = uint64(*args.Gas)
	}
	if gas == 0 || gas > maxGas {
		gas = maxGas
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

func buildCallMsg(args CallTxArgs, maxGas uint64) (arbcommon.Address, message.ContractTransaction) {
	from, tx := buildTransactionForCall(args, maxGas)
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

func handleCallResult(res *evm.TxResult, err error, blockNum *rpc.BlockNumber) (*evm.TxResult, error) {
	if err != nil {
		logMsg := logger.Warn().Err(err)
		if blockNum != nil {
			logMsg = logMsg.Int64("height", blockNum.Int64())
		} else {
			logMsg = logMsg.Str("height", "nil")
		}
		logMsg.Msg("error executing call")
		return nil, err
	}
	log.Debug().
		Uint64("gasused", res.GasUsed.Uint64()).
		Hex("returndata", res.ReturnData).
		Int("resultcode", int(res.ResultCode)).
		Msg("executed call")

	if res.ResultCode != evm.ReturnCode && res.ResultCode != evm.RevertCode {
		return nil, errors.Errorf("failed to execute call with revert code %v", res.ResultCode)
	}
	return res, err
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
