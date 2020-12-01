/*
* Copyright 2020, Offchain Labs, Inc.
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

package message

import (
	"bytes"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/pkg/errors"
	"math/big"
	"math/rand"
)

type FunctionTableEntry struct {
	FuncID  [4]byte
	Payable byte
	MaxGas  *big.Int
}

func (fte FunctionTableEntry) Equals(fte2 FunctionTableEntry) bool {
	return fte.FuncID == fte2.FuncID &&
		fte.Payable == fte2.Payable &&
		fte.MaxGas.Cmp(fte2.MaxGas) == 0
}

func NewRandomFunctionTableEntry() FunctionTableEntry {
	var funcID [4]byte
	copy(funcID[:], common.RandBytes(4))
	return FunctionTableEntry{
		FuncID:  funcID,
		Payable: byte(rand.Int() % 2),
		MaxGas:  common.RandBigInt(),
	}
}

type FunctionTable []FunctionTableEntry

func NewFunctionTableFromData(data []byte) (FunctionTable, error) {
	var ft FunctionTable
	r := bytes.NewReader(data)
	length := new(big.Int)
	if err := rlp.Decode(r, length); err != nil {
		return nil, err
	}
	if length.Cmp(big.NewInt(1024)) > 0 {
		return nil, errors.New("function table is too big")
	}
	for i := uint64(0); i < length.Uint64(); i++ {
		var funcId [4]byte
		if n, err := r.Read(funcId[:]); err != nil || n != 4 {
			return nil, errors.New("failed to read func id")
		}
		payable, err := r.ReadByte()
		if err != nil {
			return nil, errors.New("failed to read payable")
		}
		maxGas := new(big.Int)
		if err := rlp.Decode(r, maxGas); err != nil {
			return nil, err
		}
		ft = append(ft, FunctionTableEntry{
			FuncID:  funcId,
			Payable: payable,
			MaxGas:  maxGas,
		})
	}
	return ft, nil
}

func (ft FunctionTable) Encode() ([]byte, error) {
	var data []byte

	addRLPData := func(item interface{}) error {
		encoded, err := rlp.EncodeToBytes(item)
		if err != nil {
			return err
		}
		data = append(data, encoded...)
		return nil
	}

	if err := addRLPData(big.NewInt(int64(len(ft)))); err != nil {
		return nil, err
	}

	for _, row := range ft {
		data = append(data, row.FuncID[:]...)
		data = append(data, row.Payable)
		if err := addRLPData(row.MaxGas); err != nil {
			return nil, err
		}
	}
	return data, nil
}
