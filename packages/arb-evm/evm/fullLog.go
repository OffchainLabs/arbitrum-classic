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
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strconv"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type FullLog struct {
	Log
	TxIndex uint64
	TxHash  common.Hash
	Index   uint64
	Block   *common.BlockId
	Removed bool
}

func (l FullLog) Equals(o FullLog) bool {
	return l.Log.Equals(o.Log) &&
		l.TxIndex == o.TxIndex &&
		l.TxHash == o.TxHash &&
		l.Block.Equals(o.Block) &&
		l.Index == o.Index &&
		l.Removed == o.Removed
}

func (l FullLog) ToEVMLog() *types.Log {
	return &types.Log{
		Address:     l.Address.ToEthAddress(),
		Topics:      common.NewEthHashesFromHashes(l.Topics),
		Data:        l.Data,
		BlockNumber: l.Block.Height.AsInt().Uint64(),
		TxHash:      l.TxHash.ToEthHash(),
		TxIndex:     uint(l.TxIndex),
		BlockHash:   l.Block.HeaderHash.ToEthHash(),
		Index:       uint(l.Index),
		Removed:     l.Removed,
	}
}

func (l FullLog) Marshal() *FullLogBuf {
	topicStrings := make([]string, 0, len(l.Topics))
	for _, t := range l.Topics {
		topicStrings = append(topicStrings, t.String())
	}
	return &FullLogBuf{
		Address:          l.Address.Hex(),
		BlockHash:        l.Block.HeaderHash.String(),
		BlockHeight:      l.Block.Height.AsInt().Uint64(),
		Data:             hexutil.Encode(l.Data),
		Topics:           topicStrings,
		TransactionIndex: "0x" + strconv.FormatUint(l.TxIndex, 16),
		TransactionHash:  l.TxHash.String(),
		Index:            "0x" + strconv.FormatUint(l.Index, 16),
	}
}

func (x *FullLogBuf) Unmarshal() (FullLog, error) {
	ret := FullLog{}
	ret.Address = common.HexToAddress(x.Address)
	ret.Topics = make([]common.Hash, 0, len(x.Topics))
	for _, top := range x.Topics {
		ret.Topics = append(ret.Topics, common.NewHashFromEth(ethcommon.HexToHash(top)))
	}
	var err error
	ret.Data, err = hexutil.Decode(x.Data)
	if err != nil {
		return ret, err
	}
	ret.TxIndex, err = hexutil.DecodeUint64(x.TransactionIndex)
	if err != nil {
		return ret, err
	}
	ret.TxHash = common.NewHashFromEth(ethcommon.HexToHash(x.TransactionHash))
	ret.Block = &common.BlockId{
		Height:     common.NewTimeBlocks(new(big.Int).SetUint64(x.BlockHeight)),
		HeaderHash: common.NewHashFromEth(ethcommon.HexToHash(x.BlockHash)),
	}
	ret.Index, err = hexutil.DecodeUint64(x.Index)
	if err != nil {
		return ret, err
	}
	ret.Removed = false
	return ret, nil
}
