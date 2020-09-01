package web3

import (
	"context"
	"errors"
	errors2 "github.com/pkg/errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
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

func (s *Server) GetCode(ctx context.Context, address *common.Address, blockNum *rpc.BlockNumber) (hexutil.Bytes, error) {
	snap, err := s.getSnapshot(ctx, blockNum)
	if err != nil {
		return nil, err
	}
	code, err := snap.GetCode(arbcommon.NewAddressFromEth(*address))
	if err != nil {
		return nil, errors2.Wrap(err, "error getting code")
	}
	return hexutil.Bytes(code), nil
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

func (s *Server) blockNum(block *rpc.BlockNumber) (uint64, error) {
	if *block == rpc.LatestBlockNumber {
		return s.srv.GetBlockCount(), nil
	} else if *block >= 0 {
		return uint64(*block), nil
	} else {
		return 0, errors.New("unsupported block num")
	}
}

func (s *Server) GetBlockByHash(ctx context.Context, blockHashRaw hexutil.Bytes, includeTxData bool) (*GetBlockResult, error) {
	var blockHash arbcommon.Hash
	copy(blockHash[:], blockHashRaw)

	header, err := s.srv.GetBlockHeaderByHash(ctx, blockHash)
	if err != nil {
		// If we can't get the header, return nil
		return nil, nil
	}
	return s.getBlock(header, includeTxData)
}

func (s *Server) GetBlockByNumber(ctx context.Context, blockNum *rpc.BlockNumber, includeTxData bool) (*GetBlockResult, error) {
	height, err := s.blockNum(blockNum)
	if err != nil {
		return nil, err
	}
	header, err := s.srv.GetBlockHeaderByNumber(ctx, height)
	if err != nil {
		// If we can't get the header, return nil
		return nil, err
	}
	return s.getBlock(header, includeTxData)
}

func (s *Server) getBlock(header *types.Header, includeTxData bool) (*GetBlockResult, error) {
	results, err := s.srv.GetBlockResults(header.Number.Uint64())
	if err != nil {
		return nil, err
	}

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
	}, nil
}

func buildCallMsg(args CallTxArgs) (arbcommon.Address, message.ContractTransaction) {
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
	return from, message.ContractTransaction{
		MaxGas:      new(big.Int).SetUint64(gas),
		GasPriceBid: gasPrice,
		DestAddress: dest,
		Payment:     value,
		Data:        data,
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

func (s *Server) GasPrice() hexutil.Uint64 {
	return 0
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

func (s *Server) GetTransactionReceipt(ctx context.Context, txHash hexutil.Bytes) (*GetTransactionReceiptResult, error) {
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
	return &GetTransactionReceiptResult{
		Status:            receipt.Status,
		CumulativeGasUsed: receipt.CumulativeGasUsed,
		Bloom:             hexutil.Encode(receipt.Bloom.Bytes()),
		Logs:              receipt.Logs,
		TxHash:            receipt.TxHash,
		ContractAddress:   receipt.ContractAddress.Hex(),
		GasUsed:           receipt.GasUsed,
		BlockHash:         receipt.BlockHash,
		BlockNumber:       receipt.BlockNumber,
		TransactionIndex:  receipt.TransactionIndex,
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
	blockNum := hexutil.EncodeBig(res.IncomingRequest.ChainTime.BlockNum.AsInt())
	blockHash := blockInfo.Hash.ToEthHash()
	var to *string
	if tx.To() != nil {
		toStr := tx.To().Hex()
		to = &toStr
	}
	return &TransactionResult{
		BlockHash:        &blockHash,
		BlockNumber:      &blockNum,
		From:             res.IncomingRequest.Sender.ToEthAddress().Hex(),
		Gas:              hexutil.EncodeUint64(tx.Gas()),
		GasPrice:         hexutil.EncodeBig(tx.GasPrice()),
		Hash:             tx.Hash(),
		Input:            hexutil.Encode(tx.Data()),
		Nonce:            hexutil.EncodeUint64(tx.Nonce()),
		To:               to,
		TransactionIndex: &txIndex,
		Value:            hexutil.EncodeBig(tx.Value()),
		V:                hexutil.EncodeBig(vVal),
		R:                hexutil.EncodeBig(rVal),
		S:                hexutil.EncodeBig(sVal),
	}, nil
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
	txRes, err := s.makeTransactionResult(res)
	if err != nil {
		return nil, err
	}
	return txRes, nil
}

func (s *Server) GetLogs(ctx context.Context, args *GetLogsArgs) ([]LogResult, error) {
	var fromHeight *uint64
	if args.FromBlock != nil {
		from, err := s.blockNum(args.FromBlock)
		if err != nil {
			return nil, err
		}
		fromHeight = &from
	}

	var toHeight *uint64
	if args.ToBlock != nil {
		to, err := s.blockNum(args.ToBlock)
		if err != nil {
			return nil, err
		}
		toHeight = &to
	}

	addresses := make([]common.Address, 0, 1)
	if args.Address != nil {
		addresses = args.Address.addresses
	}

	topicGroups := make([][]common.Hash, 0, len(args.Topics))
	for _, topic := range args.Topics {
		topicGroups = append(topicGroups, topic.topics)
	}

	logs, err := s.srv.FindLogs(ctx, fromHeight, toHeight, addresses, topicGroups)
	if err != nil {
		return nil, err
	}
	res := make([]LogResult, 0, len(logs))
	for _, evmLog := range logs {
		logIndex := hexutil.EncodeUint64(evmLog.Index)
		txIndex := hexutil.EncodeUint64(evmLog.TxIndex)
		txHash := evmLog.TxHash.ToEthHash()
		blockHash := evmLog.Block.HeaderHash.ToEthHash()
		blockNum := hexutil.EncodeBig(evmLog.Block.Height.AsInt())
		res = append(res, LogResult{
			Removed:          false,
			LogIndex:         &logIndex,
			TransactionIndex: &txIndex,
			TransactionHash:  &txHash,
			BlockHash:        &blockHash,
			BlockNumber:      &blockNum,
			Address:          evmLog.Address.Hex(),
			Data:             hexutil.Encode(evmLog.Data),
			Topics:           arbcommon.NewEthHashesFromHashes(evmLog.Topics),
		})
	}
	return res, nil
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
