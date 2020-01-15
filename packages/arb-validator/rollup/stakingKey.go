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

package rollup

//type StakingKey struct {
//	address  common.Address
//	client   arbbridge.ArbAuthClient
//	contract arbbridge.ArbRollup
//}
//
//func (staker *StakingKey) placeStake(ctx context.Context, chain *ChainObserver) error {
//	log.Println("Staking", staker.address)
//	chain.RLock()
//	location := chain.knownValidNode
//	proof1 := GeneratePathProof(chain.nodeGraph.latestConfirmed, location)
//	proof2 := GeneratePathProof(location, chain.nodeGraph.getLeaf(location))
//	stakeAmount := chain.nodeGraph.params.StakeRequirement
//	chain.RUnlock()
//
//	return staker.contract.PlaceStake(ctx, stakeAmount, proof1, proof2)
//}
//
//func (staker *StakingKey) initiateChallenge(ctx context.Context, opp *challengeOpportunity) error {
//	return staker.contract.StartChallenge(
//		ctx,
//		opp.asserter,
//		opp.challenger,
//		opp.prevNodeHash,
//		opp.deadlineTicks.Val,
//		opp.asserterNodeType,
//		opp.challengerNodeType,
//		opp.asserterVMProtoHash,
//		opp.challengerVMProtoHash,
//		opp.asserterProof,
//		opp.challengerProof,
//		opp.asserterNodeHash,
//		opp.challengerDataHash,
//		opp.challengerPeriodTicks,
//	)
//}
//
//func (staker *StakingKey) makeAssertion(ctx context.Context, opp *preparedAssertion, proof []common.Hash) error {
//	return staker.contract.MakeAssertion(
//		ctx,
//		opp.prevPrevLeafHash,
//		opp.prevDataHash,
//		opp.prevDeadline,
//		opp.prevChildType,
//		opp.beforeState,
//		opp.params,
//		opp.claim,
//		proof,
//	)
//}
//
//func (staker *StakingKey) challengePendingTop(
//	contractAddress common.Address,
//	startBlockId *structures.BlockId,
//	startLogIndex uint,
//	pendingInbox *structures.PendingInbox,
//) (challenges.ChallengeState, error) {
//	return challenges.ChallengePendingTopClaim(
//		staker.client,
//		contractAddress,
//		startBlockId,
//		startLogIndex,
//		pendingInbox.MessageStack,
//	)
//}
//
//func (staker *StakingKey) challengeMessages(
//	contractAddress common.Address,
//	startBlockId *structures.BlockId,
//	startLogIndex uint,
//	pendingInbox *structures.PendingInbox,
//	conflictNode *Node,
//) (challenges.ChallengeState, error) {
//	return challenges.ChallengeMessagesClaim(
//		staker.client,
//		contractAddress,
//		startBlockId,
//		startLogIndex,
//		pendingInbox.MessageStack,
//		conflictNode.vmProtoData.PendingTop,
//		conflictNode.disputable.AssertionClaim.AfterPendingTop,
//	)
//}
//
//func (staker *StakingKey) challengeExecution(
//	contractAddress common.Address,
//	startBlockId *structures.BlockId,
//	startLogIndex uint,
//	mach machine.Machine,
//	pre *valprotocol.Precondition,
//) (challenges.ChallengeState, error) {
//	return challenges.ChallengeExecutionClaim(
//		staker.client,
//		contractAddress,
//		startBlockId,
//		startLogIndex,
//		pre,
//		mach,
//		false,
//	)
//}
//
//func (staker *StakingKey) defendPendingTop(
//	contractAddress common.Address,
//	startBlockId *structures.BlockId,
//	startLogIndex uint,
//	pendingInbox *structures.PendingInbox,
//	conflictNode *Node,
//) (challenges.ChallengeState, error) {
//	return challenges.DefendPendingTopClaim(
//		staker.client,
//		contractAddress,
//		startBlockId,
//		startLogIndex,
//		pendingInbox.MessageStack,
//		conflictNode.disputable.AssertionClaim.AfterPendingTop,
//		conflictNode.disputable.MaxPendingTop,
//		100,
//	)
//}
//
//func (staker *StakingKey) defendMessages(
//	contractAddress common.Address,
//	startBlockId *structures.BlockId,
//	startLogIndex uint,
//	pendingInbox *structures.PendingInbox,
//	conflictNode *Node,
//) (challenges.ChallengeState, error) {
//	return challenges.DefendMessagesClaim(
//		staker.client,
//		contractAddress,
//		startBlockId,
//		startLogIndex,
//		pendingInbox.MessageStack,
//		conflictNode.vmProtoData.PendingTop,
//		conflictNode.disputable.AssertionClaim.AfterPendingTop,
//		conflictNode.disputable.AssertionClaim.ImportedMessagesSlice,
//		100,
//	)
//}
//
//func (staker *StakingKey) defendExecution(
//	contractAddress common.Address,
//	startBlockId *structures.BlockId,
//	startLogIndex uint,
//	mach machine.Machine,
//	pre *valprotocol.Precondition,
//	numSteps uint32,
//) (challenges.ChallengeState, error) {
//	return challenges.DefendExecutionClaim(
//		staker.client,
//		contractAddress,
//		startBlockId,
//		startLogIndex,
//		pre,
//		mach,
//		numSteps,
//		50,
//	)
//}
