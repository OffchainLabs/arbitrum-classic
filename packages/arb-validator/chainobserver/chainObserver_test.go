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

package chainobserver

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/nodegraph"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"math/big"
	"testing"
)

var dummyRollupAddress1 = common.Address{1}
var dummyRollupAddress2 = common.Address{2}
var dummyRollupAddress3 = common.Address{3}
var dummyRollupAddress4 = common.Address{4}

func TestCreateEmptyChain(t *testing.T) {
	testCreateEmptyChain(dummyRollupAddress1, "dummy", contractPath, t)
	testCreateEmptyChain(dummyRollupAddress1, "fresh_rocksdb", contractPath, t)
}

func testCreateEmptyChain(rollupAddress common.Address, checkpointType string, contractPath string, t *testing.T) {
	chain, err := setUpChain(rollupAddress, checkpointType, contractPath)
	if err != nil {
		t.Fatal(err)
	}
	if chain.NodeGraph.Leaves().NumLeaves() != 1 {
		t.Fatal("unexpected leaf count", chain.NodeGraph.Leaves().NumLeaves())
	}
	tryMarshalUnmarshal(chain, t)
}

func tryMarshalUnmarshal(chain *ChainObserver, t *testing.T) {
	ctx := ckptcontext.NewCheckpointContext()
	chainBuf := chain.marshalForCheckpoint(ctx)
	chain2, err := chainBuf.unmarshalFromCheckpoint(ctx, nil)
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
	if err := doAnAssertion(chain, chain.NodeGraph.LatestConfirmed()); err != nil {
		t.Fatal(err)
	}
	validTip := chain.NodeGraph.NodeGraph.GetSuccessor(chain.NodeGraph.LatestConfirmed(), valprotocol.ValidChildType)
	if err := doAnAssertion(chain, validTip); err != nil {
		t.Fatal(err)
	}
	if chain.NodeGraph.Leaves().NumLeaves() != 5 {
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

	if err := doAnAssertion(chain, chain.NodeGraph.LatestConfirmed()); err != nil {
		t.Fatal(err)
	}
	staker1addr := common.Address{1}
	staker2addr := common.Address{2}
	contractAddr := common.Address{3}
	validTip := chain.NodeGraph.NodeGraph.GetSuccessor(chain.NodeGraph.LatestConfirmed(), valprotocol.ValidChildType)
	tip2 := chain.NodeGraph.NodeGraph.GetSuccessor(chain.NodeGraph.LatestConfirmed(), valprotocol.InvalidInboxTopChildType)
	n1, _, childType, err := nodegraph.GetConflictAncestor(validTip, tip2)
	if err != nil {
		t.Fatal(err)
	}
	confNode := n1.Prev()
	if !confNode.Equals(chain.NodeGraph.LatestConfirmed()) {
		t.Fatal("unexpected value for conflict ancestor")
	}
	if childType != valprotocol.InvalidInboxTopChildType {
		t.Fatal("unexpected value for conflict type")
	}

	createOneStaker(chain, staker1addr, validTip.Hash())
	createOneStaker(chain, staker2addr, tip2.Hash())
	challenge := nodegraph.NewChallenge(
		&common.BlockId{
			Height:     common.NewTimeBlocksInt(0),
			HeaderHash: common.Hash{},
		},
		0,
		staker1addr,
		staker2addr,
		contractAddr,
		confNode)
	chain.NodeGraph.NewChallenge(challenge)

	tryMarshalUnmarshal(chain, t)

	chain.NodeGraph.ChallengeResolved(contractAddr, staker1addr, staker2addr)

	tryMarshalUnmarshal(chain, t)
}

func doAnAssertion(chain *ChainObserver, baseNode *structures.Node) error {
	theMachine := baseNode.Machine()
	var messages []inbox.InboxMessage
	execAssertion, _, numSteps := theMachine.ExecuteAssertion(1, true, messages, true)

	assertionParams := &valprotocol.AssertionParams{
		NumSteps:             numSteps,
		ImportedMessageCount: big.NewInt(0),
	}
	assertionStub := structures.NewExecutionAssertionStubFromWholeAssertion(execAssertion, baseNode.VMProtoData().InboxTop, chain.Inbox.MessageStack)
	disputableNode := valprotocol.NewDisputableNode(
		assertionParams,
		assertionStub,
		chain.Inbox.GetTopHash(),
		big.NewInt(0),
	)

	chain.NodeGraph.CreateNodesOnAssert(
		baseNode,
		disputableNode,
		common.NewTimeBlocks(big.NewInt(10)),
	)

	nextValid := chain.NodeGraph.GetSuccessor(baseNode, valprotocol.ValidChildType)

	if err := nextValid.UpdateValidOpinion(theMachine, execAssertion); err != nil {
		return err
	}
	return nil
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

func createSomeStakers(chain *ChainObserver) {
	for i := byte(0); i < 5; i++ {
		createOneStaker(chain, common.Address{i}, chain.NodeGraph.LatestConfirmed().Hash())
	}
}

func createOneStaker(chain *ChainObserver, stakerAddr common.Address, nodeHash common.Hash) {
	chain.createStake(context.Background(), arbbridge.StakeCreatedEvent{
		ChainInfo: arbbridge.ChainInfo{
			BlockId: &common.BlockId{
				Height:     common.NewTimeBlocks(big.NewInt(73)),
				HeaderHash: common.Hash{},
			},
			LogIndex: 0,
		},
		Staker:   stakerAddr,
		NodeHash: nodeHash,
	})
}
