package web3

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
)

type BlockNumberArgs struct{}

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

type GetBlockByNumberArgs struct {
	BlockNum      *ethrpc.BlockNumber
	IncludeTxData bool
}

func (n *GetBlockByNumberArgs) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&n.BlockNum, &n.IncludeTxData}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in CallArgs: %d != %d", g, e)
	}
	return nil
}

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

type VersionArgs struct{}
