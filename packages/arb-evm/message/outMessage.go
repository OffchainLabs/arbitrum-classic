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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/pkg/errors"
	"math/big"
)

const sendMessageRootKind = 0

type OutMessage interface {
}

type SendMessageRoot struct {
	BatchNumber *big.Int
	NumInBatch  *big.Int
	OutputRoot  common.Hash
}

func NewOutMessageFromBytes(val []byte) (OutMessage, error) {
	if len(val) < 1 {
		return nil, errors.New("unexpectedly short send")
	}
	kind := val[0]
	switch kind {
	case sendMessageRootKind:
		return newSendMessageRootFromBytes(val[1:])
	default:
		return nil, errors.Errorf("unsupported message kind %v", kind)
	}
}

func newSendMessageRootFromBytes(val []byte) (*SendMessageRoot, error) {
	if len(val) != 96 {
		return nil, errors.New("unexpected send message root data length")
	}
	batchNum := new(big.Int).SetBytes(val[:32])
	numInBatch := new(big.Int).SetBytes(val[32:64])
	var outputRoot common.Hash
	copy(outputRoot[:], val[64:96])
	return &SendMessageRoot{
		BatchNumber: batchNum,
		NumInBatch:  numInBatch,
		OutputRoot:  outputRoot,
	}, nil
}
