package web3

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	goarbitrum "github.com/offchainlabs/arbitrum/packages/arb-provider-go"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	arbcommon "github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arboscontracts"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Server struct {
	srv  *aggregator.Server
	conn *goarbitrum.ArbConnection
	info *arboscontracts.ArbInfo
	sys  *arboscontracts.ArbSys
}

func NewServer(
	ctx context.Context,
	srv *aggregator.Server,
) (*Server, error) {
	chainAddress, err := srv.GetChainAddress(ctx)
	if err != nil {
		return nil, err
	}
	conn := goarbitrum.NewArbConnection(srv, nil, arbcommon.NewAddressFromEth(chainAddress))
	info, err := arboscontracts.NewArbInfo(arbos.ARB_INFO_ADDRESS, conn)
	if err != nil {
		return nil, err
	}
	sys, err := arboscontracts.NewArbSys(arbos.ARB_SYS_ADDRESS, conn)
	if err != nil {
		return nil, err
	}
	return &Server{srv: srv, conn: conn, info: info, sys: sys}, nil
}

func (s *Server) BlockNumber(r *http.Request, _ *BlockNumberArgs, reply *string) error {
	block, err := s.srv.GetBlockCount(r.Context())
	if err != nil {
		return err
	}
	*reply = "0x" + new(big.Int).SetUint64(block).Text(16)
	return nil
}

func (s *Server) GetBalance(r *http.Request, args *AccountInfoArgs, reply *string) error {
	balance, err := s.info.GetBalance(
		&bind.CallOpts{
			Pending:     false,
			From:        common.Address{},
			BlockNumber: big.NewInt(args.BlockNum.Int64()),
			Context:     r.Context(),
		},
		*args.Address,
	)
	if err != nil {
		return err
	}
	*reply = "0x" + balance.Text(16)
	return nil
}

func makeCallOpts(ctx context.Context, num rpc.BlockNumber, from common.Address) *bind.CallOpts {
	pending := false
	var blockNum *big.Int
	if num == rpc.PendingBlockNumber {
		pending = true
	} else if num != rpc.LatestBlockNumber {
		blockNum = big.NewInt(num.Int64())
	}
	return &bind.CallOpts{
		Pending:     pending,
		From:        from,
		BlockNumber: blockNum,
		Context:     ctx,
	}
}

func (s *Server) GetTransactionCount(r *http.Request, args *AccountInfoArgs, reply *string) error {
	txCount, err := s.sys.GetTransactionCount(
		makeCallOpts(r.Context(), args.BlockNum, common.Address{}),
		*args.Address,
	)
	if err != nil {
		return err
	}
	*reply = "0x" + txCount.Text(16)
	return nil
}

func (s *Server) GetCode(r *http.Request, args *AccountInfoArgs, reply *string) error {
	height, err := s.blockNum(r.Context(), &args.BlockNum)
	if err != nil {
		return err
	}
	code, err := s.conn.CodeAt(r.Context(), *args.Address, new(big.Int).SetUint64(height))
	if err != nil {
		return err
	}
	*reply = hexutil.Encode(code)
	return nil
}

func (s *Server) blockNum(ctx context.Context, block *rpc.BlockNumber) (uint64, error) {
	if *block == rpc.LatestBlockNumber {
		return s.srv.GetBlockCount(ctx)
	} else if *block >= 0 {
		return uint64(*block), nil
	} else {
		return 0, errors.New("unsupported block num")
	}
}

func (s *Server) GetBlockByNumber(r *http.Request, args *GetBlockByNumberArgs, reply *GetBlockResult) error {
	height, err := s.blockNum(r.Context(), args.BlockNum)
	if err != nil {
		return err
	}
	header, err := s.srv.GetBlockHeader(r.Context(), height)
	if err != nil {
		return err
	}
	reply.Header = *header
	results, err := s.srv.GetBlockResults(height)
	if err != nil {
		return err
	}

	if args.IncludeTxData {
		txes := make([]*TransactionResult, 0, len(results))
		for _, res := range results {
			txRes, err := s.makeTransactionResult(r.Context(), res)
			if err != nil {
				return err
			}
			txes = append(txes, txRes)
		}
		reply.Transactions = txes
	} else {
		txes := make([]hexutil.Bytes, 0, len(results))
		for _, res := range results {
			txes = append(txes, res.L1Message.MessageID().Bytes())
		}
		reply.Transactions = txes
	}
	return nil
}

func buildCallMsg(args *CallTxArgs) ethereum.CallMsg {
	var from common.Address
	if args.From != nil {
		from = *args.From
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
	return ethereum.CallMsg{
		From:     from,
		To:       args.To,
		Gas:      gas,
		GasPrice: gasPrice,
		Value:    value,
		Data:     data,
	}
}

func (s *Server) Call(r *http.Request, args *CallArgs, reply *string) error {
	ret, err := s.conn.CallContract(
		r.Context(),
		buildCallMsg(args.CallArgs),
		big.NewInt(args.BlockNum.Int64()),
	)
	if err != nil {
		return err
	}
	*reply = hexutil.Encode(ret)
	return nil
}

func (s *Server) EstimateGas(r *http.Request, args *CallTxArgs, reply *string) error {
	ret, err := s.conn.EstimateGas(
		r.Context(),
		buildCallMsg(args),
	)
	if err != nil {
		return err
	}
	*reply = hexutil.EncodeUint64(ret)
	return nil
}

func (s *Server) GasPrice(_ *http.Request, _ *EmptyArgs, reply *string) error {
	*reply = "0x" + big.NewInt(0).Text(16)
	return nil
}
func (s *Server) SendRawTransaction(r *http.Request, args *SendTransactionArgs, reply *string) error {
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(*args.Data, tx); err != nil {
		return err
	}
	txHash, err := s.srv.SendTransaction(r.Context(), tx)
	if err != nil {
		return err
	}
	*reply = txHash.String()
	return nil
}

func (s *Server) GetTransactionReceipt(r *http.Request, args *GetTransactionReceiptArgs, reply **GetTransactionReceiptResult) error {
	var requestId common.Hash
	copy(requestId[:], *args.Data)
	receipt, err := s.conn.TransactionReceipt(r.Context(), requestId)
	if err == ethereum.NotFound {
		*reply = nil
		return nil
	}
	if err != nil {
		return err
	}
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
	data, _ := json.Marshal(reply)
	log.Println("GetTransactionReceipt", string(data))
	return nil
}

func (s *Server) makeTransactionResult(ctx context.Context, res *evm.TxResult) (*TransactionResult, error) {
	chain, err := s.srv.GetChainAddress(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := aggregator.GetTransaction(res.L1Message, arbcommon.NewAddressFromEth(chain))
	if err != nil {
		return nil, err
	}
	blockInfo, err := s.srv.BlockInfo(ctx, res.L1Message.ChainTime.BlockNum.AsInt().Uint64())
	if err != nil {
		return nil, err
	}
	vVal, rVal, sVal := tx.RawSignatureValues()
	txIndex := res.TxIndex.Uint64()
	blockNum := hexutil.EncodeBig(res.L1Message.ChainTime.BlockNum.AsInt())
	blockHash := blockInfo.Hash.ToEthHash()
	var to *string
	if tx.To() != nil {
		toStr := tx.To().Hex()
		to = &toStr
	}
	return &TransactionResult{
		BlockHash:        &blockHash,
		BlockNumber:      &blockNum,
		From:             res.L1Message.Sender.ToEthAddress().Hex(),
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

func (s *Server) GetTransactionByHash(r *http.Request, args *GetTransactionReceiptArgs, reply **TransactionResult) error {
	var requestId arbcommon.Hash
	copy(requestId[:], *args.Data)
	val, err := s.srv.GetRequestResult(r.Context(), requestId)
	if err != nil {
		return err
	}
	res, err := evm.NewTxResultFromValue(val)
	if err != nil {
		return err
	}
	txRes, err := s.makeTransactionResult(r.Context(), res)
	if err != nil {
		return err
	}
	*reply = txRes
	return nil
}

func (s *Server) GetLogs(r *http.Request, args *GetLogsArgs, reply *[]LogResult) error {
	var fromHeight *uint64
	if args.FromBlock != nil {
		from, err := s.blockNum(r.Context(), args.FromBlock)
		if err != nil {
			return err
		}
		fromHeight = &from
	}

	var toHeight *uint64
	if args.ToBlock != nil {
		to, err := s.blockNum(r.Context(), args.ToBlock)
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
