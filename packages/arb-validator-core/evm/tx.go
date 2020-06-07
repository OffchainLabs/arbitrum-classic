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
)

type AVMLogProof struct {
	LogsPreHash   string
	LogsPostHash  string
	LogsValHashes []string
	OnChainTxHash common.Hash
}

func (p *AVMLogProof) Equals(o *AVMLogProof) bool {
	if len(p.LogsValHashes) != len(o.LogsValHashes) {
		return false
	}
	for i, a := range p.LogsValHashes {
		if a != o.LogsValHashes[i] {
			return false
		}
	}
	return p.LogsPreHash == o.LogsPreHash &&
		p.LogsPostHash == o.LogsPostHash &&
		p.OnChainTxHash == o.OnChainTxHash
}

func (p *AVMLogProof) Marshal() *AVMLogProofBuf {
	return &AVMLogProofBuf{
		LogPreHash:    p.LogsPreHash,
		LogPostHash:   p.LogsPostHash,
		LogValHashes:  p.LogsValHashes,
		OnChainTxHash: p.OnChainTxHash.String(),
	}
}

func (x *AVMLogProofBuf) Unmarshal() *AVMLogProof {
	return &AVMLogProof{
		LogsPreHash:   x.LogPreHash,
		LogsPostHash:  x.LogPostHash,
		LogsValHashes: x.LogValHashes,
		OnChainTxHash: common.NewHashFromEth(ethcommon.HexToHash(x.OnChainTxHash)),
	}
}

type TxInfo struct {
	Found            bool
	NodeHeight       uint64
	NodeHash         common.Hash
	TransactionIndex uint64
	TransactionHash  common.Hash
	RawVal           value.Value
	Proof            *AVMLogProof
}

func (tx TxInfo) Equals(o TxInfo) bool {
	return tx.Found == o.Found &&
		tx.NodeHeight == o.NodeHeight &&
		tx.NodeHash == o.NodeHash &&
		tx.TransactionIndex == o.TransactionIndex &&
		tx.TransactionHash == o.TransactionHash &&
		value.Eq(tx.RawVal, o.RawVal) &&
		tx.Proof.Equals(o.Proof)
}

func (tx TxInfo) Marshal() *TxInfoBuf {
	if !tx.Found {
		return &TxInfoBuf{
			Found: false,
		}
	}
	var buf bytes.Buffer
	_ = value.MarshalValue(tx.RawVal, &buf) // error can only occur from writes and bytes.Buffer is safe

	var proof *AVMLogProofBuf
	if tx.Proof != nil {
		proof = tx.Proof.Marshal()
	}
	return &TxInfoBuf{
		Found:      true,
		RawVal:     hexutil.Encode(buf.Bytes()),
		Proof:      proof,
		TxHash:     tx.TransactionHash.String(),
		TxIndex:    tx.TransactionIndex,
		NodeHash:   tx.NodeHash.String(),
		NodeHeight: tx.NodeHeight,
	}
}

func (x *TxInfoBuf) Unmarshal() (TxInfo, error) {
	if !x.Found {
		return TxInfo{Found: false}, nil
	}
	buf, err := hexutil.Decode(x.RawVal)
	if err != nil {
		return TxInfo{}, errors.Wrap(err, "GetMessageResult error")
	}
	val, err := value.UnmarshalValue(bytes.NewReader(buf))
	if err != nil {
		return TxInfo{}, errors.Wrap(err, "ValProxy.GetMessageResult: UnmarshalValue returned error")
	}

	var proof *AVMLogProof
	if x.Proof != nil {
		proof = x.Proof.Unmarshal()
	}

	return TxInfo{
		Found:            x.Found,
		NodeHeight:       x.NodeHeight,
		NodeHash:         common.NewHashFromEth(ethcommon.HexToHash(x.NodeHash)),
		TransactionIndex: x.TxIndex,
		TransactionHash:  common.NewHashFromEth(ethcommon.HexToHash(x.TxHash)),
		RawVal:           val,
		Proof:            proof,
	}, nil
}

func (tx TxInfo) ToEthReceipt(chain common.Address) (*types.Receipt, error) {
	processed, err := ProcessLog(tx.RawVal, chain)
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
	for _, l := range processed.GetLogs() {
		evmParsedTopics := make([]ethcommon.Hash, len(l.Topics))
		for j, t := range l.Topics {
			evmParsedTopics[j] = ethcommon.BytesToHash(t[:])
		}

		l := FullLog{
			Log:        l,
			TxIndex:    tx.TransactionIndex,
			TxHash:     tx.TransactionHash,
			NodeHeight: tx.NodeHeight,
			NodeHash:   tx.NodeHash,
		}.ToEVMLog()

		evmLogs = append(evmLogs, l)
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
		BlockHash:         tx.NodeHash.ToEthHash(),
		BlockNumber:       new(big.Int).SetUint64(tx.NodeHeight),
		TransactionIndex:  uint(tx.TransactionIndex),
	}, nil
}
