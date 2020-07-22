package web3

import (
	"context"
	goarbitrum "github.com/offchainlabs/arbitrum/packages/arb-provider-go"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	arbcommon "github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arboscontracts"
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

func (s *Server) BlockNumber(r *http.Request, args *BlockNumberArgs, reply *string) error {
	block, err := s.srv.GetBlockCount(r.Context())
	if err != nil {
		return err
	}
	*reply = "0x" + new(big.Int).SetUint64(block).Text(16)
	return nil
}

func (s *Server) GetBalance(r *http.Request, args *GetBalanceArgs, reply *string) error {
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

func (s *Server) GetTransactionCount(r *http.Request, args *GetTransactionCountArgs, reply *string) error {
	balance, err := s.sys.GetTransactionCount(
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

//
//func (eth *Server) GetBlockByNumber(r *http.Request, args *GetBlockByNumberArgs, reply *string) error {
//	balance, err := eth.sys.GetTransactionCount(
//		&bind.CallOpts{
//			Pending:     false,
//			From:        common.Address{},
//			BlockNumber: big.NewInt(args.BlockNum.Int64()),
//			Context:     r.Context(),
//		},
//		*args.Address,
//	)
//	if err != nil {
//		return err
//	}
//	*reply = "0x" + balance.Text(16)
//	return nil
//}

func (s *Server) Call(r *http.Request, args *CallArgs, reply *string) error {
	var from common.Address
	if args.CallArgs.From != nil {
		from = *args.CallArgs.From
	}
	gas := uint64(0)
	if args.CallArgs.Gas != nil {
		gas = uint64(*args.CallArgs.Gas)
	}
	gasPrice := big.NewInt(0)
	if args.CallArgs.GasPrice != nil {
		gasPrice = args.CallArgs.GasPrice.ToInt()
	}
	value := big.NewInt(0)
	if args.CallArgs.Value != nil {
		value = args.CallArgs.Value.ToInt()
	}
	var data []byte
	if args.CallArgs.Data != nil {
		data = *args.CallArgs.Data
	}
	ret, err := s.conn.CallContract(
		r.Context(),
		ethereum.CallMsg{
			From:     from,
			To:       args.CallArgs.To,
			Gas:      gas,
			GasPrice: gasPrice,
			Value:    value,
			Data:     data,
		},
		big.NewInt(args.BlockNum.Int64()),
	)
	if err != nil {
		return err
	}
	*reply = hexutil.Encode(ret)
	return nil
}
