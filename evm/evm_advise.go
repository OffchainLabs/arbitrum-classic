package evm

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arb-avm/value"
)

type Result interface {
	IsResult()
}

type Log struct {
}

func (e Log) String() string {
	return ""
}

type Return struct {
	Contract  *big.Int
	FuncCode  *big.Int
	ReturnVal []byte
	Logs      []Log
}

func (e Return) IsResult() {}

func (e Return) String() string {
	var sb strings.Builder
	sb.WriteString("EVMReturn(func: ")
	sb.WriteString(hexutil.Encode(e.FuncCode.Bytes()))
	sb.WriteString(", returnVal: ")
	sb.WriteString(hexutil.Encode(e.ReturnVal))
	sb.WriteString(", logs: [")
	for i, log := range e.Logs {
		sb.WriteString(log.String())
		if i != len(e.Logs)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("])")
	return sb.String()
}

type Revert struct {
	Contract  *big.Int
	FuncCode  *big.Int
	ReturnVal []byte
}

func (e Revert) IsResult() {}

func (e Revert) String() string {
	var sb strings.Builder
	sb.WriteString("EVMRevert(func: ")
	sb.WriteString(hexutil.Encode(e.FuncCode.Bytes()))
	sb.WriteString(", returnVal: ")
	sb.WriteString(hexutil.Encode(e.ReturnVal))
	sb.WriteString(")")
	return sb.String()
}

type Stop struct {
	contract  *big.Int
	funcCode *big.Int
	logs     []Log
}

func (e Stop) IsResult() {}

func (e Stop) String() string {
	var sb strings.Builder
	sb.WriteString("EVMStop(func: ")
	sb.WriteString(hexutil.Encode(e.funcCode.Bytes()))
	sb.WriteString(", logs: [")
	for i, log := range e.logs {
		sb.WriteString(log.String())
		if i != len(e.logs)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("])")
	return sb.String()
}

type FuncCall struct {
	funcID [4]byte
	logs   value.Value
}

const (
	EVM_REVERT_CODE = 0
	EVM_INVALID_CODE = 1
	EVM_RETURN_CODE = 2
	EVM_STOP_CODE = 3
)

// [logs, contract_num, func_code, return_val, return_code]
func ProcessLog(val value.Value) (Result, error) {
	tup, ok := val.(value.TupleValue)
	if !ok {
		return nil, errors.New("advise expected tuple value")
	}
	if tup.Len() != 5 {
		return nil, fmt.Errorf("advise expected tuple of length 5, but recieved %v", tup)
	}
	returnCodeVal, _ := tup.GetByInt64(4)
	returnCode, ok := returnCodeVal.(value.IntValue)
	if !ok {
		return nil, errors.New("return code must be an int")
	}

	addressVal, _ := tup.GetByInt64(1)
	addressInt, ok := addressVal.(value.IntValue)
	if !ok {
		return nil, errors.New("func id must be an int")
	}

	funcCodeVal, _ := tup.GetByInt64(2)
	funcCodeInt, ok := funcCodeVal.(value.IntValue)
	if !ok {
		return nil, errors.New("func id must be an int")
	}

	switch returnCode.BigInt().Uint64() {
	case EVM_RETURN_CODE:
		// EVM Return
		returnVal, _ := tup.GetByInt64(3)
		returnBytes, err := SizedByteArrayToHex(returnVal)
		if err != nil {
			return nil, err
		}
		return Return{addressInt.BigInt(), funcCodeInt.BigInt(), returnBytes, []Log{}}, nil
	case EVM_REVERT_CODE:
		// EVM Revert
		returnVal, _ := tup.GetByInt64(3)
		returnBytes, err := SizedByteArrayToHex(returnVal)
		if err != nil {
			return nil, err
		}
		return Revert{addressInt.BigInt(), funcCodeInt.BigInt(), returnBytes}, nil
	case EVM_STOP_CODE:
		// EVM Stop
		return Stop{addressInt.BigInt(),funcCodeInt.BigInt(), []Log{}}, nil
	case EVM_INVALID_CODE:
		return nil, errors.New("invalid tx")
	default:
		// Unknown type
		return nil, errors.New("unknown return code")
	}
}
