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

package evm

import (
	"bytes"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/pkg/errors"
	"log"
	"math/big"
	"math/rand"
)

func (x *AVMLogProof) Equals(o *AVMLogProof) bool {
	if len(x.LogValHashes) != len(o.LogValHashes) {
		return false
	}
	for i, a := range x.LogValHashes {
		if a != o.LogValHashes[i] {
			return false
		}
	}
	return x.LogPreHash == o.LogPreHash &&
		x.LogPostHash == o.LogPostHash
}

func NewRandomNodeLocation() *NodeLocation {
	return &NodeLocation{
		NodeHash:   common.RandHash().String(),
		NodeHeight: rand.Uint64(),
		L1TxHash:   common.RandHash().String(),
	}
}

func (nl *NodeLocation) Equals(o *NodeLocation) bool {
	if nl == nil && o == nil {
		return true
	}
	if nl == nil || o == nil {
		return false
	}
	return nl.NodeHeight == o.NodeHeight &&
		nl.NodeHash == o.NodeHash &&
		nl.L1TxHash == o.L1TxHash
}

func (nl *NodeLocation) NodeHashVal() common.Hash {
	return common.NewHashFromEth(ethcommon.HexToHash(nl.NodeHash))
}

type TxInfo struct {
	TransactionIndex uint64
	TransactionHash  common.Hash
	RawVal           value.Value
	StartLogIndex    uint64
	Location         *NodeLocation
	Proof            *AVMLogProof
}

func (tx *TxInfo) Equals(o *TxInfo) bool {
	return tx.TransactionIndex == o.TransactionIndex &&
		tx.TransactionHash == o.TransactionHash &&
		value.Eq(tx.RawVal, o.RawVal) &&
		tx.StartLogIndex == o.StartLogIndex &&
		tx.Location.Equals(o.Location) &&
		tx.Proof.Equals(o.Proof)
}

func (tx *TxInfo) Marshal() *TxInfoBuf {
	if tx == nil {
		return &TxInfoBuf{
			Found: false,
		}
	}
	var buf bytes.Buffer
	_ = value.MarshalValue(tx.RawVal, &buf) // error can only occur from writes and bytes.Buffer is safe

	return &TxInfoBuf{
		Found:         true,
		RawVal:        hexutil.Encode(buf.Bytes()),
		TxHash:        tx.TransactionHash.String(),
		TxIndex:       tx.TransactionIndex,
		StartLogIndex: tx.StartLogIndex,
		Location:      tx.Location,
		Proof:         tx.Proof,
	}
}

func (x *TxInfoBuf) Unmarshal() (*TxInfo, error) {
	if x == nil || !x.Found {
		return nil, nil
	}
	if !x.Found {
		return nil, nil
	}
	buf, err := hexutil.Decode(x.RawVal)
	if err != nil {
		return nil, errors.Wrap(err, "GetMessageResult error")
	}
	val, err := value.UnmarshalValue(bytes.NewReader(buf))
	if err != nil {
		return nil, errors.Wrap(err, "ValProxy.GetMessageResult: UnmarshalValue returned error")
	}

	return &TxInfo{
		TransactionIndex: x.TxIndex,
		TransactionHash:  common.NewHashFromEth(ethcommon.HexToHash(x.TxHash)),
		RawVal:           val,
		StartLogIndex:    x.StartLogIndex,
		Location:         x.Location,
		Proof:            x.Proof,
	}, nil
}

func (tx *TxInfo) ToEthReceipt() (*types.Receipt, error) {
	processed, err := ProcessLog(tx.RawVal)
	if err != nil {
		log.Println("TransactionReceipt ProcessLog error:", err)
		return nil, err
	}

	status := uint64(0)
	switch processed.(type) {
	case Return:
		status = 1
	case Stop:
		status = 1
	default:
		// Transaction unsuccessful
	}

	var evmLogs []*types.Log
	logIndex := tx.StartLogIndex
	for _, l := range processed.GetLogs() {
		evmParsedTopics := make([]ethcommon.Hash, len(l.Topics))
		for j, t := range l.Topics {
			evmParsedTopics[j] = ethcommon.BytesToHash(t[:])
		}

		l := FullLog{
			Log:      l,
			TxIndex:  tx.TransactionIndex,
			TxHash:   tx.TransactionHash,
			Location: tx.Location,
			Index:    logIndex,
		}.ToEVMLog()

		evmLogs = append(evmLogs, l)
		logIndex++
	}

	var blockHash ethcommon.Hash
	var blockNumber *big.Int
	if tx.Location != nil {
		location := tx.Location
		blockHash = ethcommon.HexToHash(location.NodeHash)
		blockNumber = new(big.Int).SetUint64(location.NodeHeight)
	}

	return &types.Receipt{
		PostState:         []byte{0},
		Status:            status,
		CumulativeGasUsed: 1,
		Bloom:             types.BytesToBloom(types.LogsBloom(evmLogs).Bytes()),
		Logs:              evmLogs,
		TxHash:            tx.TransactionHash.ToEthHash(),
		ContractAddress:   ethcommon.Address{},
		GasUsed:           1,
		BlockHash:         blockHash,
		BlockNumber:       blockNumber,
		TransactionIndex:  uint(tx.TransactionIndex),
	}, nil
}
