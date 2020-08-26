package web3

import (
	"context"
	"errors"
	errors2 "github.com/pkg/errors"
	"log"
	"math/big"
	"net/http"

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

func (s *Server) BlockNumber(_ *http.Request, _ *BlockNumberArgs, reply *string) error {
	block := s.srv.GetBlockCount()
	*reply = "0x" + new(big.Int).SetUint64(block).Text(16)
	return nil
}

func (s *Server) GetBalance(_ *http.Request, args *AccountInfoArgs, reply *string) error {
	snap, err := s.getSnapshot(args.BlockNum)
	if err != nil {
		return err
	}
	balance, err := snap.GetBalance(arbcommon.NewAddressFromEth(*args.Address))
	if err != nil {
		return errors2.Wrap(err, "error getting balance")
	}
	*reply = hexutil.EncodeBig(balance)
	return nil
}

func (s *Server) GetTransactionCount(_ *http.Request, args *AccountInfoArgs, reply *string) error {
	snap, err := s.getSnapshot(args.BlockNum)
	if err != nil {
		return err
	}
	txCount, err := snap.GetTransactionCount(arbcommon.NewAddressFromEth(*args.Address))
	if err != nil {
		return errors2.Wrap(err, "error getting transaction count")
	}
	*reply = hexutil.EncodeBig(txCount)
	return nil
}

func (s *Server) GetCode(_ *http.Request, args *AccountInfoArgs, reply *string) error {
	snap, err := s.getSnapshot(args.BlockNum)
	if err != nil {
		return err
	}
	code, err := snap.GetCode(arbcommon.NewAddressFromEth(*args.Address))
	if err != nil {
		return errors2.Wrap(err, "error getting code")
	}
	*reply = hexutil.Encode(code)
	return nil
}

func (s *Server) GetStorageAt(_ *http.Request, args *GetStorageAtArgs, reply *string) error {
	snap, err := s.getSnapshot(args.BlockNum)
	if err != nil {
		return err
	}
	storageVal, err := snap.GetStorageAt(arbcommon.NewAddressFromEth(*args.Address), (*big.Int)(args.Index))
	if err != nil {
		return errors2.Wrap(err, "error getting storage")
	}
	log.Println("Storage val", storageVal)
	*reply = hexutil.EncodeBig(storageVal)
	return nil
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

func (s *Server) GetBlockByHash(r *http.Request, args *GetBlockByHashArgs, reply **GetBlockResult) error {
	var blockHash arbcommon.Hash
	copy(blockHash[:], args.BlockHash)

	header, err := s.srv.GetBlockHeaderByHash(r.Context(), blockHash)
	if err != nil {
		return err
	}
	return s.getBlock(r.Context(), header, args.IncludeTxData, reply)
}

func (s *Server) GetBlockByNumber(r *http.Request, args *GetBlockByNumberArgs, reply **GetBlockResult) error {
	height, err := s.blockNum(args.BlockNum)
	if err != nil {
		return err
	}
	header, err := s.srv.GetBlockHeaderByNumber(r.Context(), height)
	if err != nil {
		return err
	}
	return s.getBlock(r.Context(), header, args.IncludeTxData, reply)
}

func (s *Server) getBlock(ctx context.Context, header *types.Header, includeTxData bool, reply **GetBlockResult) error {
	results, err := s.srv.GetBlockResults(header.Number.Uint64())
	if err != nil {
		return err
	}

	var transactions interface{}
	if includeTxData {
		txes := make([]*TransactionResult, 0, len(results))
		for _, res := range results {
			txRes, err := s.makeTransactionResult(res)
			if err != nil {
				return err
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
	*reply = &GetBlockResult{
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
	}
	return nil
}

func buildCallMsg(args *CallTxArgs) (arbcommon.Address, message.ContractTransaction) {
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

func (s *Server) executeCall(args *CallTxArgs, blockNum *rpc.BlockNumber) (*evm.TxResult, error) {
	snap, err := s.getSnapshot(blockNum)
	if err != nil {
		return nil, err
	}
	from, msg := buildCallMsg(args)
	msg = s.srv.AdjustGas(msg)
	return snap.Call(msg, from)
}

func (s *Server) Call(_ *http.Request, args *CallArgs, reply *string) error {
	res, err := s.executeCall(args.CallArgs, args.BlockNum)
	if err != nil {
		return err
	}
	*reply = hexutil.Encode(res.ReturnData)
	return nil
}

func (s *Server) EstimateGas(_ *http.Request, args *CallTxArgs, reply *string) error {
	blockNum := rpc.PendingBlockNumber
	res, err := s.executeCall(args, &blockNum)
	if err != nil {
		return err
	}
	*reply = hexutil.EncodeUint64(res.GasUsed.Uint64() + 1000000)
	return nil
}

func (s *Server) GasPrice(_ *http.Request, _ *EmptyArgs, reply *string) error {
	*reply = "0x" + big.NewInt(0).Text(16)
	return nil
}

func (s *Server) SendRawTransaction(_ *http.Request, args *SendTransactionArgs, reply *string) error {
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(*args.Data, tx); err != nil {
		return err
	}
	txHash, err := s.srv.SendTransaction(tx)
	if err != nil {
		return err
	}
	*reply = txHash.String()
	return nil
}

func (s *Server) GetTransactionReceipt(_ *http.Request, args *GetTransactionReceiptArgs, reply **GetTransactionReceiptResult) error {
	var requestId arbcommon.Hash
	copy(requestId[:], *args.Data)
	val, err := s.srv.GetRequestResult(requestId)
	if val == nil || err != nil {
		*reply = nil
		return nil
	}
	result, err := evm.NewTxResultFromValue(val)
	if err != nil {
		return err
	}

	blockInfo, err := s.srv.BlockInfo(result.IncomingRequest.ChainTime.BlockNum.AsInt().Uint64())
	if err != nil {
		return err
	}
	receipt := result.ToEthReceipt(blockInfo.Hash)
	*reply = &GetTransactionReceiptResult{
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
	}
	return nil
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

func (s *Server) GetTransactionByHash(_ *http.Request, args *GetTransactionReceiptArgs, reply **TransactionResult) error {
	var requestId arbcommon.Hash
	copy(requestId[:], *args.Data)
	val, err := s.srv.GetRequestResult(requestId)
	if err != nil {
		return err
	}
	res, err := evm.NewTxResultFromValue(val)
	if err != nil {
		return err
	}
	txRes, err := s.makeTransactionResult(res)
	if err != nil {
		return err
	}
	*reply = txRes
	return nil
}

func (s *Server) GetLogs(r *http.Request, args *GetLogsArgs, reply *[]LogResult) error {
	var fromHeight *uint64
	if args.FromBlock != nil {
		from, err := s.blockNum(args.FromBlock)
		if err != nil {
			return err
		}
		fromHeight = &from
	}

	var toHeight *uint64
	if args.ToBlock != nil {
		to, err := s.blockNum(args.ToBlock)
		if err != nil {
			return err
		}
		toHeight = &to
	}

	addresses := make([]common.Address, 0, 1)
	if args.Address != nil {
		addresses = append(addresses, *args.Address)
	}

	topicGroups := make([][]common.Hash, 0, len(args.Topics))
	for _, topic := range args.Topics {
		topicGroups = append(topicGroups, []common.Hash{topic})
	}

	logs, err := s.srv.FindLogs(r.Context(), fromHeight, toHeight, addresses, topicGroups)
	if err != nil {
		return err
	}
	*reply = make([]LogResult, 0, len(logs))
	for _, evmLog := range logs {
		logIndex := hexutil.EncodeUint64(evmLog.Index)
		txIndex := hexutil.EncodeUint64(evmLog.TxIndex)
		txHash := evmLog.TxHash.ToEthHash()
		blockHash := evmLog.Block.HeaderHash.ToEthHash()
		blockNum := hexutil.EncodeBig(evmLog.Block.Height.AsInt())
		*reply = append(*reply, LogResult{
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
	return nil
}

func (s *Server) getSnapshot(blockNum *rpc.BlockNumber) (*snapshot.Snapshot, error) {
	currentCount := s.srv.GetBlockCount()
	if blockNum == nil || *blockNum == rpc.PendingBlockNumber {
		return s.srv.PendingSnapshot(), nil
	}

	if *blockNum == rpc.LatestBlockNumber || blockNum.Int64() == int64(currentCount) {
		return s.srv.LatestSnapshot(), nil
	}

	return nil, errors.New("unsupported block number")
}
