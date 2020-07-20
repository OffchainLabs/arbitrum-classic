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
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
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
	Block            *common.BlockId
	Proof            *AVMLogProof
}

func (tx *TxInfo) Equals(o *TxInfo) bool {
	return tx.TransactionIndex == o.TransactionIndex &&
		tx.TransactionHash == o.TransactionHash &&
		value.Eq(tx.RawVal, o.RawVal) &&
		tx.StartLogIndex == o.StartLogIndex &&
		tx.Block.Equals(o.Block) &&
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
		BlockHash:     tx.Block.HeaderHash.String(),
		BlockHeight:   tx.Block.Height.AsInt().Uint64(),
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
		Block: &common.BlockId{
			Height:     common.NewTimeBlocks(new(big.Int).SetUint64(x.BlockHeight)),
			HeaderHash: common.NewHashFromEth(ethcommon.HexToHash(x.BlockHash)),
		},
		Proof: x.Proof,
	}, nil
}

func (tx *TxInfo) ToEthReceipt() (*types.Receipt, error) {
	result, err := NewResultFromValue(tx.RawVal)
	if err != nil {
		log.Println("TransactionReceipt NewResultFromValue error:", err)
		return nil, err
	}

	status := uint64(0)
	if result.ResultCode == ReturnCode {
		status = 1
	}

	var evmLogs []*types.Log
	logIndex := tx.StartLogIndex
	for _, l := range result.EVMLogs {
		l := FullLog{
			Log:     l,
			TxIndex: tx.TransactionIndex,
			TxHash:  tx.TransactionHash,
			Block:   tx.Block,
			Index:   logIndex,
		}.ToEVMLog()

		evmLogs = append(evmLogs, l)
		logIndex++
	}

	contractAddress := ethcommon.Address{}
	if result.L1Message.Kind == message.L2Type {
		msg, err := message.NewL2MessageFromData(result.L1Message.Data)
		if err == nil {
			if msg, ok := msg.(message.Transaction); ok {
				emptyAddress := common.Address{}
				if msg.DestAddress == emptyAddress {
					copy(contractAddress[:], result.ReturnData[12:])
				}
			}
		}
	}

	return &types.Receipt{
		PostState:         []byte{0},
		Status:            status,
		CumulativeGasUsed: 1,
		Bloom:             types.BytesToBloom(types.LogsBloom(evmLogs).Bytes()),
		Logs:              evmLogs,
		TxHash:            tx.TransactionHash.ToEthHash(),
		ContractAddress:   contractAddress,
		GasUsed:           result.GasUsed.Uint64(),
		BlockHash:         tx.Block.HeaderHash.ToEthHash(),
		BlockNumber:       tx.Block.Height.AsInt(),
		TransactionIndex:  uint(tx.TransactionIndex),
	}, nil
}
