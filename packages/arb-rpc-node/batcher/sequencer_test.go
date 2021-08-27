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

package batcher

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/rs/zerolog"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
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
) (ethcommon.Address, ethcommon.Address, *big.Int) {
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

	return createEv.RollupAddress, createEv.Inbox, receipt.BlockNumber
}

func generateTxs(t *testing.T, totalCount int, dataSizePerTx int, chainId *big.Int) []*types.Transaction {
	rand.Seed(4537345)
	signer := types.NewEIP155Signer(chainId)
	randomKeys := make([]*ecdsa.PrivateKey, 0, 10)
	for i := 0; i < 10; i++ {
		pk, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		randomKeys = append(randomKeys, pk)
	}
	txCounts := make(map[ethcommon.Address]uint64)
	var txes []*types.Transaction
	for i := 0; i < totalCount; i++ {
		pk := randomKeys[rand.Intn(len(randomKeys))]
		sender := crypto.PubkeyToAddress(pk.PublicKey)
		txData := make([]byte, dataSizePerTx)
		rand.Read(txData[:])
		tx := types.NewTransaction(txCounts[sender], ethcommon.Address{6}, big.NewInt(0), 1000, big.NewInt(10), txData)
		signedTx, err := types.SignTx(tx, signer, pk)
		if err != nil {
			t.Fatal(err)
		}
		txes = append(txes, signedTx)
		txCounts[sender]++
	}
	return txes
}

func TestSequencerBatcher(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	defer zerolog.SetGlobalLevel(zerolog.InfoLevel)

	arbosPath, err := arbos.Path()
	test.FailIfError(t, err)

	mach, err := cmachine.New(arbosPath)
	test.FailIfError(t, err)

	hash := mach.Hash()
	confirmPeriodBlocks := big.NewInt(100)
	extraChallengeTimeBlocks := big.NewInt(0)
	arbGasSpeedLimitPerBlock := big.NewInt(100000)
	baseStake := big.NewInt(100)
	var stakeToken common.Address
	var owner common.Address
	sequencerDelayBlocks := big.NewInt(200)
	sequencerDelaySeconds := big.NewInt(3000)

	l2ChainId := common.RandBigInt()

	chainIdConfig := message.ChainIDConfig{ChainId: l2ChainId}
	init, err := message.NewInitMessage(protocol.ChainParams{}, owner, []message.ChainConfigOption{chainIdConfig})
	test.FailIfError(t, err)
	extraConfig := init.ExtraConfig

	clnt, auths := test.SimulatedBackend(t)
	auth := auths[0]
	sequencer := common.NewAddressFromEth(auth.From)
	client := &ethutils.SimulatedEthClient{SimulatedBackend: clnt}

	rollupAddr, delayedInboxAddr, rollupBlock := deployRollup(
		t,
		auth,
		client,
		hash,
		confirmPeriodBlocks,
		extraChallengeTimeBlocks,
		arbGasSpeedLimitPerBlock,
		baseStake,
		stakeToken,
		owner,
		sequencer,
		sequencerDelayBlocks,
		sequencerDelaySeconds,
		extraConfig,
	)

	bridgeUtilsAddr, _, _, err := ethbridgecontracts.DeployBridgeUtils(auth, client)
	test.FailIfError(t, err)

	seqMon, shutdown := monitor.PrepareArbCore(t)
	defer shutdown()

	otherMon, shutdown2 := monitor.PrepareArbCore(t)
	defer shutdown2()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rollup, err := ethbridge.NewRollupWatcher(rollupAddr, rollupBlock.Int64(), client, bind.CallOpts{})
	test.FailIfError(t, err)

	transactAuth, err := ethbridge.NewTransactAuth(ctx, client, auth)
	test.FailIfError(t, err)

	delayedInbox, err := ethbridge.NewStandardInbox(delayedInboxAddr, client, transactAuth)
	test.FailIfError(t, err)

	seqInboxAddr, err := rollup.SequencerBridge(ctx)
	test.FailIfError(t, err)

	seqInbox, err := ethbridgecontracts.NewSequencerInbox(seqInboxAddr.ToEthAddress(), client)
	test.FailIfError(t, err)

	dummySequencerFeed := make(chan broadcaster.BroadcastFeedMessage)
	dummyDataSigner := func([]byte) ([]byte, error) { return make([]byte, 0), nil }

	for i := 0; i < 5; i++ {
		client.Commit()
	}
	time.Sleep(time.Second)

	_, err = seqMon.StartInboxReader(ctx, client, common.NewAddressFromEth(rollupAddr), rollupBlock.Int64(), common.NewAddressFromEth(bridgeUtilsAddr), nil, dummySequencerFeed)
	test.FailIfError(t, err)

	_, err = otherMon.StartInboxReader(ctx, client, common.NewAddressFromEth(rollupAddr), rollupBlock.Int64(), common.NewAddressFromEth(bridgeUtilsAddr), nil, dummySequencerFeed)
	test.FailIfError(t, err)

	batcher, err := NewSequencerBatcher(
		ctx,
		seqMon.Core,
		l2ChainId,
		seqMon.Reader,
		client,
		configuration.Sequencer{
			CreateBatchBlockInterval:   40,
			DelayedMessagesTargetDelay: 1,
		},
		seqInbox,
		auth,
		dummyDataSigner,
		nil,
	)
	test.FailIfError(t, err)
	batcher.logBatchGasCosts = true
	batcher.chainTimeCheckInterval = time.Millisecond * 10
	batcher.updateTimestampInterval = big.NewInt(1)
	batcher.sequenceDelayedMessagesInterval = big.NewInt(1)
	go batcher.Start(ctx)
	client.Commit()
	attempts := 0
	for {
		client.Commit()
		totalDelayedCount, err := seqMon.Core.GetTotalDelayedMessagesSequenced()
		test.FailIfError(t, err)
		if totalDelayedCount.Cmp(big.NewInt(0)) != 0 {
			break
		}
		time.Sleep(500 * time.Millisecond)
		attempts++

		if attempts == 20 {
			t.Fatal("sequencer didn't sequence initial message")
		}
	}
	attempts = 0
	for {
		msgCount, err := seqInbox.MessageCount(&bind.CallOpts{Context: ctx})
		test.FailIfError(t, err)

		if msgCount.Sign() > 0 {
			break
		}
		client.Commit()
		time.Sleep(20 * time.Millisecond)
		attempts++

		if attempts == 100 {
			t.Fatal("sequencer didn't create initial batch")
		}
	}

	txs := generateTxs(t, 10, 10, l2ChainId)
	totalDelayedCount := big.NewInt(1)
	for i, tx := range txs {
		if err := batcher.SendTransaction(ctx, tx); err != nil {
			t.Fatal(err)
		}
		delayedCount := 0
		if i%4 == 0 {
			delayedCount = 4
		} else if i%2 == 0 {
			delayedCount = 2
		}
		totalDelayedCount.Add(totalDelayedCount, big.NewInt(int64(delayedCount)))
		for i := 0; i < delayedCount; i++ {
			_, err = delayedInbox.SendL2MessageFromOrigin(ctx, []byte{})
			test.FailIfError(t, err)
		}
		client.Commit()
		<-time.After(time.Millisecond * 500)
	}
	for i := 0; i < 50; i++ {
		client.Commit()
	}

	timeout := time.Now().Add(time.Second * 20)
	for {
		delayedMsgCount, err := seqInbox.TotalDelayedMessagesRead(&bind.CallOpts{Context: ctx})
		test.FailIfError(t, err)
		if delayedMsgCount.Cmp(totalDelayedCount) >= 0 {
			break
		}
		if time.Now().After(timeout) {
			t.Fatal("Exceeded delayed message sequencing timeout")
		}

		time.Sleep(time.Second)
		client.Commit()
	}

	msgCount1, err := seqMon.Core.GetMessageCount()
	test.FailIfError(t, err)
	if msgCount1.Cmp(big.NewInt(30)) < 0 {
		t.Error("Not enough messages, only got", msgCount1.String())
	}

	timeout = time.Now().Add(time.Second * 30)
	for {
		msgCount2, err := otherMon.Core.GetMessageCount()
		test.FailIfError(t, err)
		t.Logf("%v/%v", msgCount2.String(), msgCount1.String())
		if msgCount2.Cmp(msgCount1) == 0 {
			break
		}
		time.Sleep(time.Second)
		client.Commit()

		if time.Now().After(timeout) {
			t.Fatal("Exceeded message delivery timeout")
		}
	}
	lastMsgIndex := new(big.Int).Sub(msgCount1, big.NewInt(1))
	seqMonAcc, err := seqMon.Core.GetInboxAcc(lastMsgIndex)
	test.FailIfError(t, err)
	otherMonAcc, err := otherMon.Core.GetInboxAcc(lastMsgIndex)
	test.FailIfError(t, err)
	if seqMonAcc != otherMonAcc {
		t.Fatal("accumulators differ between monitors")
	}
}
