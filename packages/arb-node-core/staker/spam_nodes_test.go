/*
 * Copyright 2021, Offchain Labs, Inc.
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

package staker

import (
	"context"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func makeStaker(ctx context.Context, t *testing.T, origAuth *bind.TransactOpts, client *ethutils.SimulatedEthClient, rollupAddr ethcommon.Address, validatorUtilsAddr ethcommon.Address, lookup core.ArbCoreLookup) *Staker {
	privateKey, err := crypto.GenerateKey()
	test.FailIfError(t, err)
	auth := bind.NewKeyedTransactor(privateKey)
	nonce, err := client.PendingNonceAt(ctx, origAuth.From)
	test.FailIfError(t, err)
	transferTx := types.NewTransaction(nonce, auth.From, big.NewInt(10000000), 21000, big.NewInt(0), []byte{})
	transferTx, err = origAuth.Signer(origAuth.From, transferTx)
	test.FailIfError(t, err)
	client.SendTransaction(ctx, transferTx)

	validatorAddress, _, _, err := ethbridgecontracts.DeployValidator(auth, client)
	test.FailIfError(t, err)
	wallet, err := ethbridge.NewValidator(validatorAddress, rollupAddr, client, ethbridge.NewTransactAuth(auth))

	staker, _, err := NewStaker(ctx, lookup, client, wallet, common.NewAddressFromEth(validatorUtilsAddr), MakeNodesStrategy)
	test.FailIfError(t, err)

	return staker
}

type divergenceInfo struct {
	honestNode *big.Int
	faultyNode *big.Int
}

func stakeOnNodes(ctx context.Context, t *testing.T, client *ethutils.SimulatedEthClient, staker *Staker, honestPath []*big.Int) *divergenceInfo {
	var divergence *divergenceInfo
	var lastStakedNode *big.Int
	i := 0
	for rand.Uint32()&7 != 0 {
		err := staker.advanceStake(ctx)
		test.FailIfError(t, err)
		_, err = staker.wallet.ExecuteTransactions(ctx, staker.builder)
		test.FailIfError(t, err)
		for i := 0; i < 100; i++ {
			client.Commit()
		}

		stakerInfo, err := staker.rollup.StakerInfo(ctx, staker.wallet.Address())
		test.FailIfError(t, err)
		if lastStakedNode != nil && stakerInfo.LatestStakedNode.Cmp(lastStakedNode) == 0 {
			break
		}
		lastStakedNode = stakerInfo.LatestStakedNode
		if divergence == nil && i < len(honestPath) && stakerInfo.LatestStakedNode.Cmp(honestPath[i]) != 0 {
			divergence = &divergenceInfo{
				honestNode: honestPath[i],
				faultyNode: stakerInfo.LatestStakedNode,
			}
		}
		i++
	}

	return divergence
}

func TestSpamNodes(t *testing.T) {
	ctx := context.Background()

	mach, err := cmachine.New(arbos.Path())
	test.FailIfError(t, err)

	confirmPeriodBlocks := big.NewInt(100)
	extraChallengeTimeBlocks := big.NewInt(0)
	arbGasSpeedLimitPerBlock := big.NewInt(500)
	baseStake := big.NewInt(1)
	var stakeToken common.Address
	var owner common.Address
	var extraConfig []byte

	clnt, pks := test.SimulatedBackend()
	auth := bind.NewKeyedTransactor(pks[0])
	client := &ethutils.SimulatedEthClient{SimulatedBackend: clnt}

	rollupAddr := deployRollup(
		t,
		auth,
		client,
		mach.Hash(),
		confirmPeriodBlocks,
		extraChallengeTimeBlocks,
		arbGasSpeedLimitPerBlock,
		baseStake,
		stakeToken,
		owner,
		extraConfig,
	)
	validatorUtilsAddr, _, _, err := ethbridgecontracts.DeployValidatorUtils(auth, client)
	test.FailIfError(t, err)
	client.Commit()

	arbCore, shutdown := test.PrepareArbCore(t, []inbox.InboxMessage{})
	defer shutdown()

	honestStaker := makeStaker(ctx, t, auth, client, rollupAddr, validatorUtilsAddr, arbCore)

	reader, err := NewInboxReader(ctx, honestStaker.bridge, arbCore)
	test.FailIfError(t, err)
	reader.Start(ctx)

	for i := 1; i <= 10; i++ {
		msgCount, err := arbCore.GetMessageCount()
		test.FailIfError(t, err)
		logCount, err := arbCore.GetLogCount()
		test.FailIfError(t, err)
		if msgCount.Cmp(big.NewInt(1)) >= 0 && logCount.Cmp(big.NewInt(1)) >= 0 {
			// We've found the inbox message
			reader.Stop()
			break
		}
		if i == 10 {
			t.Fatal("Failed to load initializing message")
		}
		<-time.After(time.Second * 1)
	}

	for i := 0; i < 100; i++ {
		client.Commit()
	}

	// Prepare the stakers in advance so we don't hit exponential stake increases
	faultyStakers := make([]*Staker, 0)
	for i := 0; i < 100; i++ {
		faultConfig := challenge.FaultConfig{}
		if i > 0 && rand.Uint32()&7 != 0 {
			faultConfig.DistortMachineAtGas = big.NewInt(int64(rand.Uint32() & ((1 << 13) - 1)))
		}
		faultyCore := challenge.NewFaultyCore(arbCore, faultConfig)
		staker := makeStaker(ctx, t, auth, client, rollupAddr, validatorUtilsAddr, faultyCore)
		err = staker.newStake(ctx)
		test.FailIfError(t, err)
		_, err = staker.wallet.ExecuteTransactions(ctx, staker.builder)
		test.FailIfError(t, err)
		faultyStakers = append(faultyStakers, staker)
	}
	client.Commit()

	honestPath := make([]*big.Int, 0)
	err = honestStaker.newStake(ctx)
	test.FailIfError(t, err)
	for i := 0; i < 5; i++ {
		err = honestStaker.advanceStake(ctx)
		test.FailIfError(t, err)
		_, err = honestStaker.wallet.ExecuteTransactions(ctx, honestStaker.builder)
		test.FailIfError(t, err)
		for j := 0; j < 100; j++ {
			client.Commit()
		}

		stakerInfo, err := honestStaker.rollup.StakerInfo(ctx, honestStaker.wallet.Address())
		test.FailIfError(t, err)
		if stakerInfo == nil || stakerInfo.LatestStakedNode == nil || (i > 0 && stakerInfo.LatestStakedNode.Cmp(honestPath[i-1]) == 0) {
			t.Fatal("Honest staker didn't progress stake")
		}
		honestPath = append(honestPath, stakerInfo.LatestStakedNode)
	}

	expectedConflicts := make(map[common.Address]divergenceInfo)
	for _, staker := range faultyStakers {
		divergence := stakeOnNodes(ctx, t, client, staker, honestPath)
		if divergence != nil {
			expectedConflicts[staker.wallet.Address()] = *divergence
		}
	}

	otherStaker, ourNode, otherNode, err := honestStaker.validatorUtils.FindStakerConflict(ctx, honestStaker.wallet.Address())
	test.FailIfError(t, err)
	if otherStaker == nil {
		t.Fatal("Failed to find any conflict")
	}

	rawValidatorUtils, err := ethbridgecontracts.NewValidatorUtils(validatorUtilsAddr, client)
	test.FailIfError(t, err)
	i := big.NewInt(0)
	for len(expectedConflicts) > 0 {
		var emptyAddress ethcommon.Address
		var otherStaker ethcommon.Address
		hasMore := true
		for otherStaker == emptyAddress {
			if !hasMore {
				t.Fatal("Failed to find all expected conflicts")
			}
			otherStaker, ourNode, otherNode, hasMore, err = rawValidatorUtils.FindStakerConflict(&bind.CallOpts{Context: ctx}, rollupAddr, honestStaker.wallet.Address().ToEthAddress(), i, big.NewInt(1))
			test.FailIfError(t, err)
			i = i.Add(i, big.NewInt(1))
		}

		expectedDivergence, ok := expectedConflicts[common.NewAddressFromEth(otherStaker)]
		if !ok {
			t.Fatal("FindStakerConflict found unexpected conflicting staker")
		}
		delete(expectedConflicts, common.NewAddressFromEth(otherStaker))
		if expectedDivergence.honestNode.Cmp(ourNode) != 0 {
			t.Fatal("FindStakerConflict returned unexpected node from us")
		}
		if expectedDivergence.faultyNode.Cmp(otherNode) != 0 {
			t.Fatal("FindStakerConflict returned unexpected node from other staker")
		}
	}
}
