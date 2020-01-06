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
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

var dummyAddress common.Address

func TestCreateEmptyChain(t *testing.T) {
	chain, _, err := setUpChain()
	if err != nil {
		t.Fatal(err)
	}

	chainBuf := chain.MarshalToBuf()
	chain2 := chainBuf.Unmarshal(dummyAddress, nil)
	if !chain.Equals(chain2) {
		t.Fail()
	}
}

func TestDoAssertion(t *testing.T) {
	chain, _, err := setUpChain()
	if err != nil {
		t.Fatal(err)
	}

	doAnAssertion(chain, chain.latestConfirmed)
	validTip := chain.latestConfirmed.GetSuccessor(chain.NodeGraph, structures.ValidChildType)
	doAnAssertion(chain, validTip)

	chainBuf := chain.MarshalToBuf()
	chain2 := chainBuf.Unmarshal(dummyAddress, nil)
	if !chain.Equals(chain2) {
		t.Fail()
	}
}

func doAnAssertion(chain *ChainObserver, baseNode *Node) {
	theMachine := baseNode.machine
	timeBounds := &protocol.TimeBoundsBlocks{
		Start: &protocol.TimeBlocksBuf{Val: utils.MarshalInt64ToBigIntBuf(0)},
		End:   &protocol.TimeBlocksBuf{Val: utils.MarshalInt64ToBigIntBuf(1000)},
	}
	execAssertion, numGas := theMachine.ExecuteAssertion(1, timeBounds)
	_ = execAssertion

	assertionParams := &structures.AssertionParams{
		NumSteps:             1,
		TimeBounds:           timeBounds,
		ImportedMessageCount: big.NewInt(0),
	}
	assertionStub := &protocol.ExecutionAssertionStub{
		AfterHash:        value.NewHashBuf(theMachine.Hash()),
		DidInboxInsn:     false,
		NumGas:           uint64(numGas),
		FirstMessageHash: value.NewHashBuf([32]byte{}),
		LastMessageHash:  value.NewHashBuf([32]byte{}),
		FirstLogHash:     value.NewHashBuf([32]byte{}),
		LastLogHash:      value.NewHashBuf([32]byte{}),
	}
	assertionClaim := &structures.AssertionClaim{
		AfterPendingTop:       chain.pendingInbox.GetTopHash(),
		ImportedMessagesSlice: value.NewEmptyTuple().Hash(),
		AssertionStub:         assertionStub,
	}
	disputableNode := structures.NewDisputableNode(
		assertionParams,
		assertionClaim,
		chain.pendingInbox.GetTopHash(),
		big.NewInt(0),
	)
	chain.CreateNodesOnAssert(
		baseNode,
		disputableNode,
		theMachine,
		protocol.NewTimeBlocks(big.NewInt(10)),
	)
}

func TestCreateStakers(t *testing.T) {
	chain, _, err := setUpChain()
	if err != nil {
		t.Fatal(err)
	}

	createSomeStakers(chain)

	chainBuf := chain.MarshalToBuf()
	chain2 := chainBuf.Unmarshal(dummyAddress, nil)
	if !chain.Equals(chain2) {
		t.Fail()
	}
}

func setUpChain() (*ChainObserver, machine.Machine, error) {
	var dummyAddress common.Address
	theMachine, err := loader.LoadMachineFromFile("contract.ao", true, "test")
	if err != nil {
		return nil, nil, err
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
	return chain, theMachine, nil
}

func createSomeStakers(chain *ChainObserver) {
	for i := 0; i < 5; i++ {
		stakerAddress := common.BytesToAddress([]byte{byte(i)})
		chain.CreateStake(stakerAddress, chain.latestConfirmed.hash, structures.TimeFromSeconds(73))
	}
}
