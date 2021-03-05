/*
 * Copyright 2021, Offchain Labs, Inc.
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

package evm

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/pkg/errors"
	"math/big"
)

type SendResult struct {
	BatchNumber *big.Int
	BatchIndex  *big.Int
	Data        []byte
}

func NewSendResultFromValue(tup *value.TupleValue) (*SendResult, error) {
	if tup.Len() != 5 {
		return nil, errors.Errorf("send result expected tuple of length 5, but recieved len %v", tup.Len())
	}

	resultKindVal, _ := tup.GetByInt64(0)
	batchNumberVal, _ := tup.GetByInt64(1)
	batchIndexVal, _ := tup.GetByInt64(2)
	dataSizeVal, _ := tup.GetByInt64(3)
	dataContentsVal, _ := tup.GetByInt64(4)

	resultKindInt, ok := resultKindVal.(value.IntValue)
	if !ok {
		return nil, errors.New("resultKind must be an int")
	}
	if resultKindInt.BigInt().Uint64() != 2 {
		return nil, errors.New("incorrect result kind for send")
	}
	batchNumberInt, ok := batchNumberVal.(value.IntValue)
	if !ok {
		return nil, errors.New("batchNumber must be an int")
	}
	batchIndexInt, ok := batchIndexVal.(value.IntValue)
	if !ok {
		return nil, errors.New("batchIndex must be an int")
	}
	dataSizeInt, ok := dataSizeVal.(value.IntValue)
	if !ok {
		return nil, errors.New("dataSize must be an int")
	}
	dataContentsBuf, ok := dataContentsVal.(*value.Buffer)
	if !ok {
		return nil, errors.New("dataContents must be a buffer")
	}

	data, err := inbox.BufAndLengthToBytes(dataSizeInt.BigInt(), dataContentsBuf)
	if err != nil {
		return nil, err
	}
	return &SendResult{
		BatchNumber: batchNumberInt.BigInt(),
		BatchIndex:  batchIndexInt.BigInt(),
		Data:        data,
	}, nil
}

type SendResultMessage interface {
}

type SendResultMessageType uint8

const (
	WithdrawEthType SendResultMessageType = 0
	SendTxToL1Type  SendResultMessageType = 3
	BuddyResultType SendResultMessageType = 5
)

func NewVirtualSendResultFromData(data []byte) (SendResultMessage, error) {
	if len(data) < 32 {
		return nil, errors.New("send result message must be at least 32 bytes")
	}
	typecode := new(big.Int).SetBytes(data[:32])
	switch SendResultMessageType(typecode.Uint64()) {
	case WithdrawEthType:
		return NewWithdrawEthResultFromData(data)
	case SendTxToL1Type:
		return NewL2ToL1TxResultFromData(data)
	case BuddyResultType:
		return NewBuddyResultFromData(data)
	default:
		return nil, errors.Errorf("unhandled send result message type %v", data[0])
	}
}

type BuddyResult struct {
	Address   common.Address
	Succeeded bool
}

func addressFromBytes(data []byte) common.Address {
	var address common.Address
	copy(address[:], data[12:])
	return address
}

func NewBuddyResultFromData(data []byte) (*BuddyResult, error) {
	if len(data) != 65 {
		return nil, errors.New("unexpected buddy result length")
	}
	typeCode := new(big.Int).SetBytes(data[0:32])
	contract := data[32:64]
	success := data[64]

	if typeCode.Cmp(big.NewInt(int64(BuddyResultType))) != 0 {
		return nil, errors.New("unexpected type code")
	}

	return &BuddyResult{
		Address:   addressFromBytes(contract),
		Succeeded: success == 1,
	}, nil
}

type WithdrawEthResult struct {
	Destination common.Address
	Amount      *big.Int
}

func NewWithdrawEthResultFromData(data []byte) (*WithdrawEthResult, error) {
	if len(data) != 32*3 {
		return nil, errors.New("unexpected withdraw eth result length")
	}
	typeCode := new(big.Int).SetBytes(data[0:32])
	destination := data[32:64]
	amount := new(big.Int).SetBytes(data[64:])

	if typeCode.Cmp(big.NewInt(int64(WithdrawEthType))) != 0 {
		return nil, errors.New("unexpected type code")
	}

	var address common.Address
	copy(address[:], destination[12:])

	return &WithdrawEthResult{
		Destination: address,
		Amount:      amount,
	}, nil
}

type L2ToL1TxResult struct {
	L2Sender  common.Address
	L1Dest    common.Address
	L2Block   *big.Int
	L1Block   *big.Int
	Timestamp *big.Int
	Value     *big.Int
	Calldata  []byte
}

func NewL2ToL1TxResultFromData(data []byte) (*L2ToL1TxResult, error) {
	if len(data) < 224 {
		return nil, errors.New("unexpected L2 to L1 tx result length")
	}
	typeCode := new(big.Int).SetBytes(data[0:32])
	if typeCode.Cmp(big.NewInt(int64(SendTxToL1Type))) != 0 {
		return nil, errors.New("unexpected type code")
	}
	l2Sender := data[32:64]
	l1Dest := data[64:96]
	l2Block := data[96:128]
	l1Block := data[128:160]
	timestamp := data[160:192]
	payment := data[192:224]
	calldata := data[224:]
	return &L2ToL1TxResult{
		L2Sender:  addressFromBytes(l2Sender),
		L1Dest:    addressFromBytes(l1Dest),
		L2Block:   new(big.Int).SetBytes(l2Block),
		L1Block:   new(big.Int).SetBytes(l1Block),
		Timestamp: new(big.Int).SetBytes(timestamp),
		Value:     new(big.Int).SetBytes(payment),
		Calldata:  calldata,
	}, nil
}
