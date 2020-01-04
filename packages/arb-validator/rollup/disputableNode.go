/*
* Copyright 2019-2020, Offchain Labs, Inc.
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

package rollup

import (
	"math/big"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
	protocol "github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

type DisputableNode struct {
	timeBounds            [2]RollupTime
	importedMessageCount  *big.Int
	afterPendingTop       [32]byte
	importedMessagesSlice [32]byte
	assertionStub         *protocol.AssertionStub
	hash                  [32]byte
}

func (dn *DisputableNode) MarshalToBuf() *DisputableNodeBuf {
	return &DisputableNodeBuf{
		TimeLowerBound:        dn.timeBounds[0].MarshalToBuf(),
		TimeUpperBound:        dn.timeBounds[1].MarshalToBuf(),
		AfterPendingTop:       marshalHash(dn.afterPendingTop),
		ImportedMessagesSlice: marshalHash(dn.importedMessagesSlice),
		ImportedMessageCount:  marshalBigInt(dn.importedMessageCount),
		AssertionStub:         dn.assertionStub,
	}
}

func (buf *DisputableNodeBuf) Unmarshal() *DisputableNode {
	ret := &DisputableNode{
		timeBounds:            [2]RollupTime{buf.TimeLowerBound.Unmarshal(), buf.TimeUpperBound.Unmarshal()},
		afterPendingTop:       unmarshalHash(buf.AfterPendingTop),
		importedMessagesSlice: unmarshalHash(buf.ImportedMessagesSlice),
		importedMessageCount:  unmarshalBigInt(buf.ImportedMessageCount),
		assertionStub:         buf.AssertionStub,
	}
	ret.hash = ret._hash()
	return ret
}

func (dn *DisputableNode) _hash() [32]byte {
	var ret [32]byte
	retSlice := solsha3.SoliditySHA3(
		solsha3.Bytes32(unmarshalHash(dn.assertionStub.AfterHash)),
		solsha3.Bool(dn.assertionStub.DidInboxInsn),
		solsha3.Uint32(dn.assertionStub.NumSteps),
		solsha3.Uint64(dn.assertionStub.NumGas),
		solsha3.Bytes32(unmarshalHash(dn.assertionStub.FirstMessageHash)),
		solsha3.Bytes32(unmarshalHash(dn.assertionStub.LastMessageHash)),
		solsha3.Bytes32(unmarshalHash(dn.assertionStub.FirstLogHash)),
		solsha3.Bytes32(unmarshalHash(dn.assertionStub.LastLogHash)),
	)
	copy(ret[:], retSlice)
	return ret
}
