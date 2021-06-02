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
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/nodehealth"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func deployRollup(
	t *testing.T,
	auth *bind.TransactOpts,
	client *ethutils.SimulatedEthClient,
	machineHash [32]byte,
	confirmPeriodBlocks *big.Int,
	extraChallengeTimeBlocks *big.Int,
	arbGasSpeedLimitPerBlock *big.Int,
	baseStake *big.Int,
	stakeToken common.Address,
	owner common.Address,
	sequencer common.Address,
	sequencerDelayBlocks *big.Int,
	sequencerDelaySeconds *big.Int,
	extraConfig []byte,
) ethcommon.Address {
	osp1Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProof(auth, client)
	test.FailIfError(t, err)
	osp2Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProof2(auth, client)
	test.FailIfError(t, err)
	osp3Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProofHash(auth, client)
	test.FailIfError(t, err)
	challengeFactoryAddr, _, _, err := ethbridgetestcontracts.DeployChallengeFactory(auth, client, []ethcommon.Address{osp1Addr, osp2Addr, osp3Addr})
	test.FailIfError(t, err)

	_, tx, rollupCreator, err := ethbridgetestcontracts.DeployRollupCreatorNoProxy(
		auth,
		client,
		challengeFactoryAddr,
		machineHash,
		confirmPeriodBlocks,
		extraChallengeTimeBlocks,
		arbGasSpeedLimitPerBlock,
		baseStake,
		stakeToken.ToEthAddress(),
		owner.ToEthAddress(),
		sequencer.ToEthAddress(),
		sequencerDelayBlocks,
		sequencerDelaySeconds,
		extraConfig,
	)
	test.FailIfError(t, err)
	client.Commit()

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	test.FailIfError(t, err)
	createEv, err := rollupCreator.ParseRollupCreated(*receipt.Logs[len(receipt.Logs)-1])
	test.FailIfError(t, err)

	return createEv.RollupAddress
}

type ExpectedChallengeEnd uint8

const (
	NoChallenge ExpectedChallengeEnd = iota
	OneStepProof
	Timeout
)

func requireChallengeLogs(ctx context.Context, t *testing.T, client ethutils.EthClient, challenge *common.Address, topics []string) {
	if challenge == nil {
		t.Fatal("Expected challenge but found none")
	}
	topicHashes := make([]ethcommon.Hash, 0, len(topics))
	for _, topic := range topics {
		hash := hashing.SoliditySHA3([]byte(topic))
		topicHashes = append(topicHashes, hash.ToEthHash())
	}
	query := ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(0),
		ToBlock:   nil,
		Addresses: []ethcommon.Address{challenge.ToEthAddress()},
		Topics:    [][]ethcommon.Hash{topicHashes},
	}
	logs, err := client.FilterLogs(ctx, query)
	test.FailIfError(t, err)
	if len(logs) == 0 {
		t.Fatal("Challenge ended in unexpected manner")
	}
}

func runStakersTest(t *testing.T, faultConfig challenge.FaultConfig, maxGasPerNode *big.Int, expectedEnd ExpectedChallengeEnd) {
	ctx := context.Background()

	arbosPath, err := arbos.Path()
	test.FailIfError(t, err)

	mach, err := cmachine.New(arbosPath)
	test.FailIfError(t, err)

	hash := mach.Hash()
	confirmPeriodBlocks := big.NewInt(100)
	extraChallengeTimeBlocks := big.NewInt(0)
	arbGasSpeedLimitPerBlock := maxGasPerNode
	baseStake := big.NewInt(100)
	var stakeToken common.Address
	sequencerDelayBlocks := big.NewInt(60)
	sequencerDelaySeconds := big.NewInt(900)
	var extraConfig []byte

	clnt, pks := test.SimulatedBackend(t)
	auth := bind.NewKeyedTransactor(pks[0])
	auth2 := bind.NewKeyedTransactor(pks[1])
	seqAuth := bind.NewKeyedTransactor(pks[2])
	ownerAuth := bind.NewKeyedTransactor(pks[2])
	sequencer := common.NewAddressFromEth(seqAuth.From)
	client := &ethutils.SimulatedEthClient{SimulatedBackend: clnt}

	rollupAddr := deployRollup(
		t,
		auth,
		client,
		hash,
		confirmPeriodBlocks,
		extraChallengeTimeBlocks,
		arbGasSpeedLimitPerBlock,
		baseStake,
		stakeToken,
		common.NewAddressFromEth(ownerAuth.From),
		sequencer,
		sequencerDelayBlocks,
		sequencerDelaySeconds,
		extraConfig,
	)

	bridgeUtilsAddr, _, _, err := ethbridgecontracts.DeployBridgeUtils(auth, client)
	test.FailIfError(t, err)

	validatorUtilsAddr, _, _, err := ethbridgecontracts.DeployValidatorUtils(auth, client)
	test.FailIfError(t, err)

	validatorWalletFactory, _, _, err := ethbridgecontracts.DeployValidatorWalletCreator(auth, client)
	test.FailIfError(t, err)

	valAuth, err := ethbridge.NewTransactAuth(ctx, client, auth, "")
	test.FailIfError(t, err)
	val2Auth, err := ethbridge.NewTransactAuth(ctx, client, auth2, "")
	test.FailIfError(t, err)

	validatorAddress, err := ethbridge.CreateValidatorWallet(ctx, validatorWalletFactory, valAuth, client)
	test.FailIfError(t, err)

	// Should lookup WalletCreated event
	checkValidatorAddress, err := ethbridge.CreateValidatorWallet(ctx, validatorWalletFactory, valAuth, client)
	test.FailIfError(t, err)
	if validatorAddress != checkValidatorAddress {
		t.Error("CreateValidatorWallet didn't reuse existing wallet")
	}

	validatorAddress2, err := ethbridge.CreateValidatorWallet(ctx, validatorWalletFactory, val2Auth, client)
	test.FailIfError(t, err)
	if validatorAddress == validatorAddress2 {
		t.Error("CreateValidatorWallet reused existing wallet for different address")
	}

	client.Commit()

	rollupAdmin, err := ethbridgecontracts.NewRollupAdminFacet(rollupAddr, client)
	test.FailIfError(t, err)
	_, err = rollupAdmin.SetValidator(ownerAuth, []ethcommon.Address{validatorAddress, validatorAddress2}, []bool{true, true})
	test.FailIfError(t, err)
	client.Commit()

	mon, shutdown := monitor.PrepareArbCore(t)
	defer shutdown()

	val, err := ethbridge.NewValidator(validatorAddress, rollupAddr, client, valAuth)
	test.FailIfError(t, err)
	val2, err := ethbridge.NewValidator(validatorAddress2, rollupAddr, client, val2Auth)
	test.FailIfError(t, err)

	staker, _, err := NewStaker(ctx, mon.Core, client, val, common.NewAddressFromEth(validatorUtilsAddr), MakeNodesStrategy)
	test.FailIfError(t, err)

	staker.Validator.GasThreshold = big.NewInt(0)

	seqInboxAddr, err := staker.rollup.SequencerBridge(ctx)
	test.FailIfError(t, err)

	seqInbox, err := ethbridgecontracts.NewSequencerInbox(seqInboxAddr.ToEthAddress(), client)
	test.FailIfError(t, err)

	delayedBridgeAddr, err := staker.rollup.DelayedBridge(ctx)
	test.FailIfError(t, err)

	delayedBridge, err := ethbridgecontracts.NewBridge(delayedBridgeAddr.ToEthAddress(), client)
	test.FailIfError(t, err)

	delayedAcc, err := delayedBridge.InboxAccs(&bind.CallOpts{Context: ctx}, big.NewInt(0))
	test.FailIfError(t, err)
	batchItem := inbox.NewDelayedItem(big.NewInt(0), big.NewInt(1), common.Hash{}, big.NewInt(0), delayedAcc)

	latestHeader, err := client.HeaderByNumber(ctx, nil)
	test.FailIfError(t, err)
	currentBlockNumber := latestHeader.Number
	currentTimestamp := big.NewInt(int64(latestHeader.Time))

	endOfBlockMessage := message.NewInboxMessage(
		message.EndBlockMessage{},
		common.Address{},
		big.NewInt(1),
		big.NewInt(0),
		inbox.ChainTime{
			BlockNum:  common.NewTimeBlocks(currentBlockNumber),
			Timestamp: currentTimestamp,
		},
	)

	endBlockBatchItem := inbox.NewSequencerItem(big.NewInt(1), endOfBlockMessage, batchItem.Accumulator)
	delayedAccInt := new(big.Int).SetBytes(delayedAcc[:])
	metadata := []*big.Int{big.NewInt(0), currentBlockNumber, currentTimestamp, big.NewInt(1), delayedAccInt}
	_, err = seqInbox.AddSequencerL2BatchFromOrigin(seqAuth, []byte{}, []*big.Int{}, metadata, endBlockBatchItem.Accumulator)
	test.FailIfError(t, err)
	for i := 0; i < 5; i++ {
		client.Commit()
	}

	faultyCore := challenge.NewFaultyCore(mon.Core, faultConfig)

	faultyStaker, _, err := NewStaker(ctx, faultyCore, client, val2, common.NewAddressFromEth(validatorUtilsAddr), MakeNodesStrategy)
	test.FailIfError(t, err)

	faultyStaker.Validator.GasThreshold = big.NewInt(0)

	registry := prometheus.NewRegistry()
	const largeChannelBuffer = 200
	healthChan := make(chan nodehealth.Log, largeChannelBuffer)
	healthChan <- nodehealth.Log{Config: true, Var: "disablePrimaryCheck", ValBool: false}
	healthChan <- nodehealth.Log{Config: true, Var: "disableOpenEthereumCheck", ValBool: true}
	healthChan <- nodehealth.Log{Config: true, Var: "healthcheckMetrics", ValBool: false}
	healthChan <- nodehealth.Log{Config: true, Var: "healthcheckRPC", ValStr: "0.0.0.0:8080"}
	nodehealth.Init(healthChan)

	go func() {
		err := nodehealth.StartNodeHealthCheck(ctx, healthChan, registry, registry)
		test.FailIfError(t, err)
	}()

	// Make a dummy feed for now
	var sequencerFeed chan broadcaster.BroadcastFeedMessage

	_, err = mon.StartInboxReader(ctx, client, common.NewAddressFromEth(rollupAddr), common.NewAddressFromEth(bridgeUtilsAddr), healthChan, sequencerFeed)
	test.FailIfError(t, err)

	for i := 1; i <= 10; i++ {
		msgCount, err := mon.Core.GetMessageCount()
		test.FailIfError(t, err)
		logCount, err := mon.Core.GetLogCount()
		test.FailIfError(t, err)
		if msgCount.Cmp(big.NewInt(1)) >= 0 && logCount.Cmp(big.NewInt(1)) >= 0 {
			// We've found the inbox message
			break
		}
		if i == 10 {
			t.Fatal("Failed to load initializing message")
		}
		<-time.After(time.Second * 1)
	}

	faultsExist := faultConfig != challenge.FaultConfig{}
	t.Log("faultsExist:", faultsExist)

	var targetNode *big.Int
	if faultsExist {
		targetNode = big.NewInt(1)
	} else {
		targetNode = big.NewInt(3)
	}

	var lastChallenge *common.Address
	faultyStakerAlive := false
	faultyStakerDead := false

	stakerMadeFirstMove := false
	for i := 400; i >= 0; i-- {
		if (i % 2) == 0 {
			fmt.Println("Honest staker acting")
			tx, err := staker.Act(ctx)
			test.FailIfError(t, err)
			if tx != nil {
				stakerMadeFirstMove = true
			}
		} else if (!faultyStakerAlive || !faultyStakerDead) && stakerMadeFirstMove {
			fmt.Println("Malicious staker acting")
			_, err = faultyStaker.Act(ctx)
			if err != nil {
				errString := err.Error()
				if faultsExist && (strings.Contains(errString, "WRONG_END") || strings.Contains(errString, "BIS_DEADLINE")) && expectedEnd == Timeout {
					faultyStakerAlive = true
					faultyStakerDead = true
				} else {
					test.FailIfError(t, err)
				}
			}
		}
		client.Commit()
		client.Commit()

		faultyStakerInfo, err := staker.rollup.StakerInfo(ctx, common.NewAddressFromEth(validatorAddress2))
		test.FailIfError(t, err)
		if faultyStakerInfo == nil {
			faultyStakerDead = true
		} else {
			faultyStakerAlive = true
			faultyStakerDead = false
			if faultyStakerInfo.CurrentChallenge != nil {
				lastChallenge = faultyStakerInfo.CurrentChallenge
			}
		}

		latestConfirmed, err := staker.rollup.LatestConfirmedNode(ctx)
		test.FailIfError(t, err)
		if latestConfirmed.Cmp(targetNode) >= 0 {
			break
		} else if i == 0 {
			t.Fatal("Node not confirmed")
		}
	}

	switch expectedEnd {
	case NoChallenge:
		if lastChallenge != nil {
			t.Fatal("Unexpected challenge")
		}
	case Timeout:
		requireChallengeLogs(ctx, t, client, lastChallenge, []string{"AsserterTimedOut()", "ChallengerTimedOut()"})
	case OneStepProof:
		requireChallengeLogs(ctx, t, client, lastChallenge, []string{"OneStepProofCompleted()"})
	}

	stakerInfo, err := staker.rollup.StakerInfo(ctx, common.NewAddressFromEth(validatorAddress))
	test.FailIfError(t, err)

	if stakerInfo == nil {
		t.Fatal("Staker isn't staked")
	}

	if stakerInfo.CurrentChallenge != nil {
		t.Fatal("Staker remained in challenge")
	}

	if stakerInfo.LatestStakedNode.Cmp(big.NewInt(0)) == 0 {
		t.Fatal("Staker didn't stake on node")
	}

	faultyStakerInfo, err := staker.rollup.StakerInfo(ctx, common.NewAddressFromEth(validatorAddress2))
	test.FailIfError(t, err)

	if faultsExist {
		if faultyStakerInfo != nil {
			t.Fatal("Faulty staker is still staked")
		}
	} else {
		if faultyStakerInfo == nil {
			t.Fatal("Other staker lost stake")
		}
	}
}

func calculateGasToFirstInbox(t *testing.T) *big.Int {
	mon, shutdown := monitor.PrepareArbCore(t)
	defer shutdown()
	cursor, err := mon.Core.GetExecutionCursor(big.NewInt(100000000))
	test.FailIfError(t, err)
	inboxGas := new(big.Int).Add(cursor.TotalGasConsumed(), big.NewInt(1))
	t.Logf("Found first inbox instruction starting at %v", inboxGas)
	return inboxGas
}

func TestChallengeToOSP(t *testing.T) {
	runStakersTest(t, challenge.FaultConfig{DistortMachineAtGas: big.NewInt(1)}, big.NewInt(390), OneStepProof)
}

func TestChallengeToInboxOSP(t *testing.T) {
	inboxGas := calculateGasToFirstInbox(t)
	runStakersTest(t, challenge.FaultConfig{DistortMachineAtGas: inboxGas}, new(big.Int).Add(inboxGas, big.NewInt(10000)), OneStepProof)
}

func TestChallengeTimeout(t *testing.T) {
	runStakersTest(t, challenge.FaultConfig{DistortMachineAtGas: big.NewInt(1)}, big.NewInt(2), Timeout)
}

func TestStakersCooperative(t *testing.T) {
	runStakersTest(t, challenge.FaultConfig{}, big.NewInt(25000), NoChallenge)
}
