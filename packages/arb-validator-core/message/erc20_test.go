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
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func generateTestERC20() ERC20 {
	addr1 := common.Address{}
	addr1[0] = 76
	addr1[19] = 93

	addr2 := common.Address{}
	addr2[0] = 43
	addr2[19] = 12

	addr3 := common.Address{}
	addr3[0] = 87
	addr3[19] = 32

	return ERC20{
		To:           addr1,
		From:         addr2,
		TokenAddress: addr3,
		Value:        big.NewInt(89735406),
	}
}

func TestMarshalERC20(t *testing.T) {
	msg := generateTestERC20()

	msg2, err := UnmarshalExecuted(msg.Type(), msg.AsInboxValue(), generateTestChain())
	if err != nil {
		t.Error(err)
	}

	if !msg.Equals(msg2) {
		t.Error("Unmarshalling didn't reverse marshalling", msg, msg2)
	}
}

func TestCheckpointERC20(t *testing.T) {
	msg := generateTestERC20()
	msg2, err := UnmarshalFromCheckpoint(msg.Type(), msg.CheckpointValue())
	if err != nil {
		t.Error(err)
	}

	if !msg.Equals(msg2) {
		t.Error("Unmarshalling didn't reverse marshalling", msg, msg2)
	}
}
