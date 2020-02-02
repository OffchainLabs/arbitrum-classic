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
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/checkpointing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

var dummyAddress common.Address

var dummyRollupAddress1 = common.Address{1}
var dummyRollupAddress2 = common.Address{2}
var dummyRollupAddress3 = common.Address{3}
var dummyRollupAddress4 = common.Address{4}

var contractPath = "../contract.ao"

func TestCreateEmptyChain(t *testing.T) {
	testCreateEmptyChain(dummyRollupAddress1, "dummy", contractPath, t)
	testCreateEmptyChain(dummyRollupAddress1, "fresh_rocksdb", contractPath, t)
}

func testCreateEmptyChain(rollupAddress common.Address, checkpointType string, contractPath string, t *testing.T) {
	chain, err := setUpChain(rollupAddress, checkpointType, contractPath)
	if err != nil {
		t.Fatal(err)
	}
	if chain.nodeGraph.leaves.NumLeaves() != 1 {
		t.Fatal("unexpected leaf count")
	}
	tryMarshalUnmarshal(chain, t)
}

func tryMarshalUnmarshal(chain *ChainObserver, t *testing.T) {
	ctx := structures.NewCheckpointContextImpl()
	chainBuf := chain.marshalForCheckpoint(ctx)
	chain2, err := chainBuf.UnmarshalFromCheckpoint(context.TODO(), ctx, nil)
	if err != nil {
		t.Error(err)
	}
	if !chain.equals(chain2) {
		t.Fail()
	}
}

func TestDoAssertion(t *testing.T) {
	testDoAssertion(dummyRollupAddress2, "dummy", contractPath, t)
	testDoAssertion(dummyRollupAddress2, "fresh_rocksdb", contractPath, t)
}

func testDoAssertion(dummyRollupAddress common.Address, checkpointType string, contractPath string, t *testing.T) {
	chain, err := setUpChain(dummyRollupAddress, checkpointType, contractPath)
	if err != nil {
		t.Fatal(err)
	}

	doAnAssertion(chain, chain.nodeGraph.latestConfirmed, t)
	validTip := chain.nodeGraph.latestConfirmed.GetSuccessor(chain.nodeGraph.NodeGraph, structures.ValidChildType)
	doAnAssertion(chain, validTip, t)
	if chain.nodeGraph.leaves.NumLeaves() != 7 {
		t.Fatal("unexpected leaf count")
	}

	tryMarshalUnmarshal(chain, t)
}

func TestChallenge(t *testing.T) {
	testChallenge(dummyRollupAddress3, "dummy", contractPath, t)
	testChallenge(dummyRollupAddress3, "fresh_rocksdb", contractPath, t)
}

func testChallenge(dummyRollupAddress common.Address, checkpointType string, contractPath string, t *testing.T) {
	chain, err := setUpChain(dummyRollupAddress, checkpointType, contractPath)
	if err != nil {
		t.Fatal(err)
	}

	doAnAssertion(chain, chain.nodeGraph.latestConfirmed, t)
	staker1addr := common.Address{1}
	staker2addr := common.Address{2}
	contractAddr := common.Address{3}
	validTip := chain.nodeGraph.latestConfirmed.GetSuccessor(chain.nodeGraph.NodeGraph, structures.ValidChildType)
	tip2 := chain.nodeGraph.latestConfirmed.GetSuccessor(chain.nodeGraph.NodeGraph, structures.InvalidMessagesChildType)
	n1, _, childType, err := GetConflictAncestor(validTip, tip2)
	if err != nil {
		t.Fatal(err)
	}
	confNode := n1.prev
	if !confNode.Equals(chain.nodeGraph.latestConfirmed) {
		t.Fatal("unexpected value for conflict ancestor")
	}
	if childType != structures.InvalidMessagesChildType {
		t.Fatal("unexpected value for conflict type")
	}

	createOneStaker(chain, staker1addr, validTip.hash)
	createOneStaker(chain, staker2addr, tip2.hash)

	chain.nodeGraph.NewChallenge(contractAddr, staker1addr, staker2addr, structures.InvalidMessagesChildType)

	tryMarshalUnmarshal(chain, t)

	chain.nodeGraph.ChallengeResolved(contractAddr, staker1addr, staker2addr)

	tryMarshalUnmarshal(chain, t)
}

func doAnAssertion(chain *ChainObserver, baseNode *Node, t *testing.T) {
	theMachine := baseNode.machine
	timeBounds := &protocol.TimeBoundsBlocks{
		Start: common.NewTimeBlocks(big.NewInt(0)),
		End:   common.NewTimeBlocks(big.NewInt(1000)),
	}
	execAssertion, numGas := theMachine.ExecuteAssertion(1, timeBounds, value.NewEmptyTuple(), time.Hour)
	_ = execAssertion

	assertionParams := &structures.AssertionParams{
		NumSteps:             1,
		TimeBounds:           timeBounds,
		ImportedMessageCount: big.NewInt(0),
	}
	assertionStub := &valprotocol.ExecutionAssertionStub{
		AfterHash:        theMachine.Hash(),
		DidInboxInsn:     false,
		NumGas:           numGas,
		FirstMessageHash: common.Hash{},
		LastMessageHash:  common.Hash{},
		FirstLogHash:     common.Hash{},
		LastLogHash:      common.Hash{},
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

	err := chain.nodeGraph.createNodesOnAssert(
		baseNode,
		disputableNode,
		common.NewTimeBlocks(big.NewInt(10)),
		common.Hash{},
	)
	if err != nil {
		t.Error(err)
	}
	chain.nodeGraph.nodeFromHash[baseNode.successorHashes[3]].machine = theMachine
}

func TestCreateStakers(t *testing.T) {
	testCreateStakers(dummyRollupAddress4, "dummy", contractPath, t)
	testCreateStakers(dummyRollupAddress4, "fresh_rocksdb", contractPath, t)
}

func testCreateStakers(dummyRollupAddress common.Address, checkpointType string, contractPath string, t *testing.T) {
	chain, err := setUpChain(dummyRollupAddress, checkpointType, contractPath)
	if err != nil {
		t.Fatal(err)
	}

	createSomeStakers(chain)
	if checkpointType != "dummy" {
		tryMarshalUnmarshal(chain, t)
	}
}

func setUpChain(rollupAddress common.Address, checkpointType string, contractPath string) (*ChainObserver, error) {
	var checkpointer checkpointing.RollupCheckpointer
	switch checkpointType {
	case "dummy":
		checkpointFac := checkpointing.NewDummyCheckpointerFactory(contractPath)
		checkpointer = checkpointFac.New(context.TODO())
	case "fresh_rocksdb":
		checkpointFac := checkpointing.NewRollupCheckpointerImplFactory(rollupAddress, contractPath, "", big.NewInt(1000000), true)
		checkpointer = checkpointFac.New(context.TODO())
	}
	chain, err := NewChain(
		dummyAddress,
		checkpointer,
		structures.ChainParams{
			StakeRequirement:        big.NewInt(1),
			GracePeriod:             common.TimeFromSeconds(60 * 60),
			MaxExecutionSteps:       1000000,
			MaxTimeBoundsWidth:      20,
			ArbGasSpeedLimitPerTick: 1000,
		},
		false,
		&structures.BlockID{
			Height:     common.NewTimeBlocks(big.NewInt(10)),
			HeaderHash: common.Hash{},
		},
	)
	if err != nil {
		return nil, err
	}
	chain.Start(context.Background())
	return chain, nil
}

func createSomeStakers(chain *ChainObserver) {
	for i := byte(0); i < 5; i++ {
		createOneStaker(chain, common.Address{i}, chain.nodeGraph.latestConfirmed.hash)
	}
}

func createOneStaker(chain *ChainObserver, stakerAddr common.Address, nodeHash common.Hash) {
	chain.createStake(context.Background(), arbbridge.StakeCreatedEvent{
		ChainInfo: arbbridge.ChainInfo{
			BlockID: &structures.BlockID{
				Height:     common.NewTimeBlocks(big.NewInt(73)),
				HeaderHash: common.Hash{},
			},
			LogIndex: 0,
			TxHash:   [32]byte{},
		},
		Staker:   stakerAddr,
		NodeHash: nodeHash,
	})
}
