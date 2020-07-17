package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arboscontracts"

	"log"
	"math"
	"math/big"
	"net/http"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	goarbitrum "github.com/offchainlabs/arbitrum/packages/arb-provider-go"
)

// UpCodec creates a CodecRequest to process each request.
type UpCodec struct {
}

// NewUpCodec returns a new UpCodec.
func NewUpCodec() *UpCodec {
	return &UpCodec{}
}

// NewRequest returns a new CodecRequest of type UpCodecRequest.
func (c *UpCodec) NewRequest(r *http.Request) rpc.CodecRequest {
	outerCR := &UpCodecRequest{}
	jsonC := NewCodec()
	innerCR := jsonC.NewRequest(r)

	outerCR.CodecRequest = innerCR.(*CodecRequest)
	return outerCR
}

type UpCodecRequest struct {
	*CodecRequest
}

func (c *UpCodecRequest) Method() (string, error) {
	m, err := c.CodecRequest.Method()
	if len(m) > 1 && err == nil {
		parts := strings.Split(m, "_")
		if len(parts) != 2 {
			return "", errors.New("malformed request")
		}
		service, method := parts[0], parts[1]

		r, n := utf8.DecodeRuneInString(service) // get the first rune, and it's length
		if unicode.IsLower(r) {
			service = string(unicode.ToUpper(r)) + service[n:]
		}
		r, n = utf8.DecodeRuneInString(method)
		if unicode.IsLower(r) {
			method = string(unicode.ToUpper(r)) + method[n:]
		}
		modifiedRequest := service + "." + method
		log.Println("Made request", modifiedRequest)
		return modifiedRequest, nil
	}
	return m, err
}

type Eth struct {
	conn *goarbitrum.ArbConnection
	info *arboscontracts.ArbInfo
	sys  *arboscontracts.ArbSys
}

func NewEth(conn *goarbitrum.ArbConnection) (*Eth, error) {
	info, err := arboscontracts.NewArbInfo(goarbitrum.ARB_INFO_ADDRESS, conn)
	if err != nil {
		return nil, err
	}

	sys, err := arboscontracts.NewArbSys(goarbitrum.ARB_SYS_ADDRESS, conn)
	if err != nil {
		return nil, err
	}
	return &Eth{conn: conn, info: info, sys: sys}, nil
}

type BlockNumberArgs struct{}

func (eth *Eth) BlockNumber(r *http.Request, args *BlockNumberArgs, reply *string) error {
	//blockHeight, err := eth.sys.TimeUpperBound(
	//	&bind.CallOpts{
	//		Pending:     false,
	//		From:        common.Address{},
	//		BlockNumber: nil,
	//		Context:     r.Context(),
	//	},
	//)
	//if err != nil {
	//	return err
	//}
	//*reply = "0x" + blockHeight.Text(16)
	//return nil
	//log.Println("Calling BlockNumber")
	*reply = "0x" + big.NewInt(532).Text(16)
	return nil
}

type GetBalanceArgs struct {
	Address  *common.Address
	BlockNum *ethrpc.BlockNumber
}

func (n *GetBalanceArgs) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&n.Address, &n.BlockNum}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in CallArgs: %d != %d", g, e)
	}
	return nil
}

func (eth *Eth) GetBalance(r *http.Request, args *GetBalanceArgs, reply *string) error {
	balance, err := eth.info.GetBalance(
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

type GetTransactionCountArgs struct {
	Address  *common.Address
	BlockNum *ethrpc.BlockNumber
}

func (n *GetTransactionCountArgs) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&n.Address, &n.BlockNum}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in CallArgs: %d != %d", g, e)
	}
	return nil
}

func (eth *Eth) GetTransactionCount(r *http.Request, args *GetTransactionCountArgs, reply *string) error {
	balance, err := eth.sys.GetTransactionCount(
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

//type GetBlockByNumberArgs struct {
//	BlockNum      *ethrpc.BlockNumber
//	IncludeTxData bool
//}
//
//func (n *GetBlockByNumberArgs) UnmarshalJSON(buf []byte) error {
//	tmp := []interface{}{&n.BlockNum, &n.IncludeTxData}
//	wantLen := len(tmp)
//	if err := json.Unmarshal(buf, &tmp); err != nil {
//		return err
//	}
//	if g, e := len(tmp), wantLen; g != e {
//		return fmt.Errorf("wrong number of fields in CallArgs: %d != %d", g, e)
//	}
//	return nil
//}
//
//func (eth *Eth) GetBlockByNumber(r *http.Request, args *GetBlockByNumberArgs, reply *string) error {
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

type CallTxArgs struct {
	From     *common.Address `json:"from"`
	To       *common.Address `json:"to"`
	Gas      *hexutil.Uint64 `json:"gas"`
	GasPrice *hexutil.Big    `json:"gasPrice"`
	Value    *hexutil.Big    `json:"value"`
	Data     *hexutil.Bytes  `json:"data"`
}

type CallArgs struct {
	CallArgs *CallTxArgs
	BlockNum *ethrpc.BlockNumber
}

func (n *CallArgs) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&n.CallArgs, &n.BlockNum}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in CallArgs: %d != %d", g, e)
	}
	return nil
}

func (eth *Eth) Call(r *http.Request, args *CallArgs, reply *string) error {
	log.Printf("Call %+v %v %v\n", args.CallArgs, hexutil.Encode(args.CallArgs.To.Bytes()), *args.BlockNum)
	var from common.Address
	if args.CallArgs.From != nil {
		from = *args.CallArgs.From
	}
	gas := uint64(math.MaxUint64 / 2)
	if args.CallArgs.Gas != nil {
		gas = uint64(*args.CallArgs.Gas)
	}
	gasPrice := new(big.Int)
	if args.CallArgs.GasPrice != nil {
		gasPrice = args.CallArgs.GasPrice.ToInt()
	}
	value := new(big.Int)
	if args.CallArgs.Value != nil {
		value = args.CallArgs.Value.ToInt()
	}
	var data []byte
	if args.CallArgs.Data != nil {
		data = *args.CallArgs.Data
	}
	ret, err := eth.conn.CallContract(
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
	log.Println("Call response", *reply)
	return nil
}

type Net struct {
}

type VersionArgs struct{}

func (net *Net) Version(r *http.Request, args *VersionArgs, reply *string) error {
	*reply = "123456789"
	return nil
}

func main() {
	s := rpc.NewServer()

	// Register our own Codec
	s.RegisterCodec(NewUpCodec(), "application/json")
	s.RegisterCodec(NewUpCodec(), "application/json;charset=UTF-8")

	clnt, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		panic(err)
	}

	arbclnt, err := goarbitrum.Dial("http://localhost:1235", nil, clnt)
	if err != nil {
		panic(err)
	}

	eth, err := NewEth(arbclnt)
	if err != nil {
		panic(err)
	}

	err = s.RegisterService(eth, "Eth")
	if err != nil {
		panic(err)
	}

	err = s.RegisterService(new(Net), "Net")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.Handle("/", s)
	fmt.Println(http.ListenAndServe(":8545", r))
}
