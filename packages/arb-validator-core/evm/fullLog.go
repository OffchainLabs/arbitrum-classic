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
	"strconv"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type FullLog struct {
	Log
	TxIndex    uint64
	TxHash     common.Hash
	NodeHeight uint64
	NodeHash   common.Hash
	Index      uint64
	Removed    bool
}

func (l FullLog) Equals(o FullLog) bool {
	return l.Log.Equals(o.Log) &&
		l.TxIndex == o.TxIndex &&
		l.TxHash == o.TxHash &&
		l.NodeHeight == o.NodeHeight &&
		l.NodeHash == o.NodeHash &&
		l.Index == o.Index &&
		l.Removed == o.Removed
}

func (l FullLog) ToEVMLog() *types.Log {
	evmParsedTopics := make([]ethcommon.Hash, len(l.Topics))
	for j, t := range l.Topics {
		evmParsedTopics[j] = ethcommon.BytesToHash(t[:])
	}

	return &types.Log{
		Address:     l.Address.ToEthAddress(),
		Topics:      evmParsedTopics,
		Data:        l.Data,
		BlockNumber: l.NodeHeight,
		TxHash:      l.TxHash.ToEthHash(),
		TxIndex:     uint(l.TxIndex),
		BlockHash:   l.NodeHash.ToEthHash(),
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
		BlockHash:        l.NodeHash.String(),
		BlockNumber:      "0x" + strconv.FormatUint(l.NodeHeight, 16),
		Data:             hexutil.Encode(l.Data),
		LogIndex:         "0x" + strconv.FormatUint(0, 16),
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
	ret.NodeHeight, err = hexutil.DecodeUint64(x.BlockNumber)
	if err != nil {
		return ret, err
	}
	ret.NodeHash = common.NewHashFromEth(ethcommon.HexToHash(x.BlockHash))
	ret.Index, err = hexutil.DecodeUint64(x.Index)
	if err != nil {
		return ret, err
	}
	ret.Removed = false
	return ret, nil
}
