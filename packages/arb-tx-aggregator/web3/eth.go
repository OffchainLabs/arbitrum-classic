package web3

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rlp"
	errors2 "github.com/pkg/errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	arbcommon "github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type Server struct {
	srv *aggregator.Server
}

func NewServer(
	srv *aggregator.Server,
) *Server {
	return &Server{srv: srv}
}

func (s *Server) ChainId() hexutil.Uint64 {
	return hexutil.Uint64(message.ChainAddressToID(
		arbcommon.NewAddressFromEth(s.srv.GetChainAddress()),
	).Uint64())
}

func (s *Server) GasPrice() hexutil.Uint64 {
	return 0
}

func (s *Server) Accounts() []common.Address {
	return nil
}

func (s *Server) BlockNumber() hexutil.Uint64 {
	return hexutil.Uint64(s.srv.GetBlockCount())
}

func (s *Server) GetBalance(ctx context.Context, address *common.Address, blockNum *rpc.BlockNumber) (*hexutil.Big, error) {
	snap, err := s.getSnapshot(ctx, blockNum)
	if err != nil {
		return nil, err
	}
	balance, err := snap.GetBalance(arbcommon.NewAddressFromEth(*address))
	if err != nil {
		return nil, errors2.Wrap(err, "error getting balance")
	}
	return (*hexutil.Big)(balance), nil
}

func (s *Server) GetStorageAt(ctx context.Context, address *common.Address, index *hexutil.Big, blockNum *rpc.BlockNumber) (*hexutil.Big, error) {
	snap, err := s.getSnapshot(ctx, blockNum)
	if err != nil {
		return nil, err
	}
	storageVal, err := snap.GetStorageAt(arbcommon.NewAddressFromEth(*address), (*big.Int)(index))
	if err != nil {
		return nil, errors2.Wrap(err, "error getting storage")
	}
	return (*hexutil.Big)(storageVal), nil
}

func (s *Server) GetTransactionCount(ctx context.Context, address *common.Address, blockNum *rpc.BlockNumber) (hexutil.Uint64, error) {
	account := arbcommon.NewAddressFromEth(*address)
	if blockNum == nil || *blockNum == rpc.PendingBlockNumber {
		count := s.srv.PendingTransactionCount(account)
		if count != nil {
			return hexutil.Uint64(*count), nil
		}
	}
	snap, err := s.getSnapshot(ctx, blockNum)
	if err != nil {
		return 0, err
	}
	txCount, err := snap.GetTransactionCount(account)
	if err != nil {
		return 0, errors2.Wrap(err, "error getting transaction count")
	}
	return hexutil.Uint64(txCount.Uint64()), nil
}

func (s *Server) GetBlockTransactionCountByHash(ctx context.Context, blockHash common.Hash) (*hexutil.Big, error) {
	header, err := s.srv.Client.HeaderByHash(ctx, blockHash)
	if err != nil {
		// If we can't get the header, return nil
		return nil, nil
	}
	return s.getBlockTransactionCount(header.Number.Uint64())
}

func (s *Server) GetBlockTransactionCountByNumber(blockNum *rpc.BlockNumber) (*hexutil.Big, error) {
	height, err := s.blockNum(blockNum)
	if err != nil {
		return nil, err
	}
	return s.getBlockTransactionCount(height)
}

func (s *Server) GetCode(ctx context.Context, address *common.Address, blockNum *rpc.BlockNumber) (hexutil.Bytes, error) {
	snap, err := s.getSnapshot(ctx, blockNum)
	if err != nil {
		return nil, err
	}
	code, err := snap.GetCode(arbcommon.NewAddressFromEth(*address))
	if err != nil {
		return nil, errors2.Wrap(err, "error getting code")
	}
	return code, nil
}

func (s *Server) SendRawTransaction(data hexutil.Bytes) (hexutil.Bytes, error) {
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(data, tx); err != nil {
		return nil, err
	}
	txHash, err := s.srv.SendTransaction(tx)
	if err != nil {
		return nil, err
	}
	return txHash[:], nil
}

func (s *Server) Call(ctx context.Context, callArgs CallTxArgs, blockNum *rpc.BlockNumber) (hexutil.Bytes, error) {
	res, err := s.executeCall(ctx, callArgs, blockNum)
	if err != nil {
		return nil, err
	}
	return res.ReturnData, nil
}

func (s *Server) EstimateGas(ctx context.Context, args CallTxArgs) (hexutil.Uint64, error) {
	blockNum := rpc.PendingBlockNumber
	res, err := s.executeCall(ctx, args, &blockNum)
	if err != nil {
		return 0, err
	}
	return hexutil.Uint64(res.GasUsed.Uint64() + 1000000), nil
}

func (s *Server) GetBlockByHash(ctx context.Context, blockHashRaw hexutil.Bytes, includeTxData bool) (*GetBlockResult, error) {
	var blockHash common.Hash
	copy(blockHash[:], blockHashRaw)

	header, err := s.srv.Client.HeaderByHash(ctx, blockHash)
	if err != nil {
		// If we can't get the header, return nil
		return nil, nil
	}
	return s.getBlock(header, blockHash, includeTxData)
}

func (s *Server) GetBlockByNumber(ctx context.Context, blockNum *rpc.BlockNumber, includeTxData bool) (*GetBlockResult, error) {
	height, err := s.blockNum(blockNum)
	if err != nil {
		return nil, err
	}
	header, err := s.srv.Client.HeaderByNumber(ctx, new(big.Int).SetUint64(height))
	if err != nil {
		// If we can't get the header, return nil
		return nil, err
	}
	l1BlockInfo, err := s.srv.Client.BlockInfoByNumber(ctx, new(big.Int).SetUint64(height))
	if err != nil {
		return nil, err
	}

	return s.getBlock(header, l1BlockInfo.Hash, includeTxData)
}

func (s *Server) GetTransactionByHash(txHash hexutil.Bytes) (*TransactionResult, error) {
	var requestId arbcommon.Hash
	copy(requestId[:], txHash)
	val, err := s.srv.GetRequestResult(requestId)
	if err != nil {
		return nil, err
	}
	res, err := evm.NewTxResultFromValue(val)
	if err != nil {
		return nil, err
	}
	return s.makeTransactionResult(res)
}

func (s *Server) GetTransactionByBlockHashAndIndex(ctx context.Context, blockHash common.Hash, index hexutil.Uint64) (*TransactionResult, error) {
	header, err := s.srv.Client.HeaderByHash(ctx, blockHash)
	if err != nil {
		// If we can't get the header, return nil
		return nil, nil
	}

	return s.getTransactionByBlockAndIndex(header.Number.Uint64(), index)
}

func (s *Server) GetTransactionByBlockNumberAndIndex(blockNum *rpc.BlockNumber, index hexutil.Uint64) (*TransactionResult, error) {
	height, err := s.blockNum(blockNum)
	if err != nil {
		return nil, err
	}

	return s.getTransactionByBlockAndIndex(height, index)
}

func (s *Server) GetTransactionReceipt(txHash hexutil.Bytes) (*GetTransactionReceiptResult, error) {
	var requestId arbcommon.Hash
	copy(requestId[:], txHash)
	val, err := s.srv.GetRequestResult(requestId)
	if val == nil || err != nil {
		return nil, nil
	}
	result, err := evm.NewTxResultFromValue(val)
	if err != nil {
		return nil, err
	}

	blockInfo, err := s.srv.BlockInfo(result.IncomingRequest.ChainTime.BlockNum.AsInt().Uint64())
	if err != nil {
		return nil, err
	}
	if blockInfo == nil {
		return nil, nil
	}

	receipt := result.ToEthReceipt(blockInfo.Hash)

	var contractAddress *common.Address
	emptyAddress := common.Address{}
	if receipt.ContractAddress != emptyAddress {
		contractAddress = &receipt.ContractAddress
	}

	provenance := result.IncomingRequest.Provenance
	var parentRequestId *common.Hash
	emptyParent := arbcommon.Hash{}
	if provenance.ParentRequestId != emptyParent {
		h := provenance.ParentRequestId.ToEthHash()
		parentRequestId = &h
	}

	return &GetTransactionReceiptResult{
		Status:            hexutil.Uint64(receipt.Status),
		CumulativeGasUsed: hexutil.Uint64(receipt.CumulativeGasUsed),
		Bloom:             receipt.Bloom.Bytes(),
		Logs:              receipt.Logs,
		TxHash:            receipt.TxHash,
		ContractAddress:   contractAddress,
		GasUsed:           hexutil.Uint64(receipt.GasUsed),
		BlockHash:         receipt.BlockHash,
		BlockNumber:       (*hexutil.Big)(receipt.BlockNumber),
		TransactionIndex:  hexutil.Uint64(receipt.TransactionIndex),
		ReturnCode:        hexutil.Uint64(result.ResultCode),
		L1SeqNum:          (*hexutil.Big)(provenance.L1SeqNum),
		ParentRequestId:   parentRequestId,
		IndexInParent:     (*hexutil.Big)(provenance.IndexInParent),
	}, nil
}

func (s *Server) GetLogs(ctx context.Context, args filters.FilterCriteria) ([]*types.Log, error) {
	var fromHeight *uint64

	if args.FromBlock != nil {
		fromRaw := args.FromBlock.Int64()
		from, err := s.blockNum((*rpc.BlockNumber)(&fromRaw))
		if err != nil {
			return nil, err
		}
		fromHeight = &from
	}

	var toHeight *uint64
	if args.ToBlock != nil {
		toRaw := args.FromBlock.Int64()
		to, err := s.blockNum((*rpc.BlockNumber)(&toRaw))
		if err != nil {
			return nil, err
		}
		toHeight = &to
	}

	logs, err := s.srv.FindLogs(ctx, fromHeight, toHeight, args.Addresses, args.Topics)
	if err != nil {
		return nil, err
	}
	res := make([]*types.Log, 0, len(logs))
	for _, evmLog := range logs {
		res = append(res, evmLog.ToEVMLog())
	}
	return res, nil
}

func (s *Server) getBlockTransactionCount(height uint64) (*hexutil.Big, error) {
	block, err := s.srv.BlockInfo(height)
	if err != nil {
		return nil, err
	}

	blockInfo, err := s.srv.GetBlockInfo(block)
	if err != nil {
		return nil, err
	}
	return (*hexutil.Big)(blockInfo.BlockStats.TxCount), nil
}

func (s *Server) getTransactionByBlockAndIndex(height uint64, index hexutil.Uint64) (*TransactionResult, error) {
	block, err := s.srv.BlockInfo(height)
	if err != nil {
		// If we can't get the header, return nil
		return nil, nil
	}

	blockInfo, err := s.srv.GetBlockInfo(block)
	if err != nil {
		return nil, err
	}

	txRes, err := s.srv.GetTxInBlockAtIndexResults(blockInfo, uint64(index))
	if err != nil {
		return nil, err
	}
	return s.makeTransactionResult(txRes)
}

func (s *Server) getBlock(header *types.Header, blockHash common.Hash, includeTxData bool) (*GetBlockResult, error) {
	block, err := s.srv.BlockInfo(header.Number.Uint64())
	if err != nil {
		return nil, err
	}

	blockInfo, err := s.srv.GetBlockInfo(block)
	if err != nil {
		return nil, err
	}

	results, err := s.srv.GetBlockResults(blockInfo)
	if err != nil {
		return nil, err
	}

	bloom, gasLimit, gasUsed := aggregator.GetBlockFields(block, blockInfo)

	var transactions interface{}
	if includeTxData {
		txes := make([]*TransactionResult, 0, len(results))
		for _, res := range results {
			txRes, err := s.makeTransactionResult(res)
			if err != nil {
				return nil, err
			}
			txes = append(txes, txRes)
		}
		transactions = txes
	} else {
		txes := make([]hexutil.Bytes, 0, len(results))
		for _, res := range results {
			txes = append(txes, res.IncomingRequest.MessageID.Bytes())
		}
		transactions = txes
	}
	size := uint64(0)
	uncles := make([]hexutil.Bytes, 0)
	return &GetBlockResult{
		Number:           (*hexutil.Big)(header.Number),
		Hash:             blockHash.Bytes(),
		ParentHash:       header.ParentHash.Bytes(),
		MixDigest:        header.MixDigest.Bytes(),
		Nonce:            &header.Nonce,
		Sha3Uncles:       header.UncleHash.Bytes(),
		LogsBloom:        bloom.Bytes(),
		TransactionsRoot: header.TxHash.Bytes(),
		StateRoot:        header.Root.Bytes(),
		ReceiptsRoot:     header.ReceiptHash.Bytes(),
		Miner:            header.Coinbase.Bytes(),
		Difficulty:       (*hexutil.Big)(header.Difficulty),
		TotalDifficulty:  (*hexutil.Big)(header.Difficulty),
		ExtraData:        (*hexutil.Bytes)(&header.Extra),
		Size:             (*hexutil.Uint64)(&size),
		GasLimit:         (*hexutil.Uint64)(&gasLimit),
		GasUsed:          (*hexutil.Uint64)(&gasUsed),
		Timestamp:        (*hexutil.Uint64)(&header.Time),
		Transactions:     transactions,
		Uncles:           &uncles,
	}, nil
}

func (s *Server) makeTransactionResult(res *evm.TxResult) (*TransactionResult, error) {
	tx, err := aggregator.GetTransaction(res.IncomingRequest)
	if err != nil {
		return nil, err
	}
	blockInfo, err := s.srv.BlockInfo(res.IncomingRequest.ChainTime.BlockNum.AsInt().Uint64())
	if err != nil {
		return nil, err
	}
	vVal, rVal, sVal := tx.RawSignatureValues()
	txIndex := res.TxIndex.Uint64()
	blockNum := res.IncomingRequest.ChainTime.BlockNum.AsInt()
	blockHash := blockInfo.Hash.ToEthHash()
	return &TransactionResult{
		BlockHash:        &blockHash,
		BlockNumber:      (*hexutil.Big)(blockNum),
		From:             res.IncomingRequest.Sender.ToEthAddress(),
		Gas:              hexutil.Uint64(tx.Gas()),
		GasPrice:         (*hexutil.Big)(tx.GasPrice()),
		Hash:             tx.Hash(),
		Input:            tx.Data(),
		Nonce:            hexutil.Uint64(tx.Nonce()),
		To:               tx.To(),
		TransactionIndex: (*hexutil.Uint64)(&txIndex),
		Value:            (*hexutil.Big)(tx.Value()),
		V:                (*hexutil.Big)(vVal),
		R:                math.U256Bytes(rVal),
		S:                math.U256Bytes(sVal),
	}, nil
}

func buildCallMsg(args CallTxArgs) (arbcommon.Address, message.Call) {
	var from arbcommon.Address
	if args.From != nil {
		from = arbcommon.NewAddressFromEth(*args.From)
	}
	gas := uint64(0)
	if args.Gas != nil {
		gas = uint64(*args.Gas)
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

	var dest arbcommon.Address
	if args.To != nil {
		dest = arbcommon.NewAddressFromEth(*args.To)
	}
	return from, message.Call{
		BasicTx: message.BasicTx{
			MaxGas:      new(big.Int).SetUint64(gas),
			GasPriceBid: gasPrice,
			DestAddress: dest,
			Payment:     value,
			Data:        data,
		},
	}
}

func (s *Server) executeCall(ctx context.Context, args CallTxArgs, blockNum *rpc.BlockNumber) (*evm.TxResult, error) {
	snap, err := s.getSnapshot(ctx, blockNum)
	if err != nil {
		return nil, err
	}
	from, msg := buildCallMsg(args)
	msg = s.srv.AdjustGas(msg)
	return snap.Call(msg, from)
}

func (s *Server) getSnapshot(ctx context.Context, blockNum *rpc.BlockNumber) (*snapshot.Snapshot, error) {
	if blockNum == nil || *blockNum == rpc.PendingBlockNumber {
		return s.srv.PendingSnapshot(), nil
	}

	if *blockNum == rpc.LatestBlockNumber {
		return s.srv.LatestSnapshot(), nil
	}

	snap, err := s.srv.GetSnapshot(ctx, uint64(*blockNum))
	if err != nil {
		return nil, err
	}
	if snap != nil {
		return snap, nil
	}

	return nil, errors.New("unsupported block number")
}

func (s *Server) blockNum(block *rpc.BlockNumber) (uint64, error) {
	if *block == rpc.LatestBlockNumber {
		return s.srv.GetBlockCount(), nil
	} else if *block >= 0 {
		return uint64(*block), nil
	} else {
		return 0, errors.New("unsupported block num")
	}
}
