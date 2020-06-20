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

func generateTestSingleDeliveredERC20() SingleDelivered {
	return SingleDelivered{
		Message: generateTestERC20(),
		DeliveryInfo: DeliveryInfo{
			ChainTime: ChainTime{
				BlockNum:  common.NewTimeBlocks(big.NewInt(64654)),
				Timestamp: big.NewInt(65435643),
			},
			TxId: big.NewInt(9675),
		},
	}
}

func generateTestDeliveredERC20() Delivered {
	return Delivered{
		Message: generateTestERC20(),
		DeliveryInfo: DeliveryInfo{
			ChainTime: ChainTime{
				BlockNum:  common.NewTimeBlocks(big.NewInt(64654)),
				Timestamp: big.NewInt(65435643),
			},
			TxId: big.NewInt(9675),
		},
	}
}

func TestMarshalSingleDelivered(t *testing.T) {
	msg := generateTestSingleDeliveredERC20()
	inboxVal := msg.AsInboxValue()
	msg2, err := UnmarshalSingleDelivered(inboxVal, generateTestChain())
	if err != nil {
		t.Fatal(err)
	}

	if !msg.Equals(msg2) {
		t.Error("Unmarshalling didn't reverse marshalling", msg, msg2)
	}
}

func TestCheckpoint(t *testing.T) {
	msg := generateTestDeliveredERC20()

	msg2, err := UnmarshalDeliveredFromCheckpoint(msg.CheckpointValue())
	if err != nil {
		t.Error(err)
	}

	if !msg.Equals(msg2) {
		t.Error("Unmarshalling didn't reverse marshalling", msg, msg2)
	}
}
