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
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func TestCreateChain(t *testing.T) {
	var dummyAddress common.Address
	theMachine, err := loader.LoadMachineFromFile("contract.ao", true, "test")
	if err != nil {
		t.Fatal(err)
	}
	chain := NewChain(
		dummyAddress,
		theMachine,
		structures.ChainParams{
			StakeRequirement:        big.NewInt(1),
			GracePeriod:             structures.TimeFromSeconds(60 * 60),
			MaxExecutionSteps:       1000000,
			ArbGasSpeedLimitPerTick: 1000,
		},
	)
	chainBuf := chain.MarshalToBuf()
	chain2 := chainBuf.Unmarshal(dummyAddress, nil)
	if !chain.Equals(chain2) {
		t.Fail()
	}
}
