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
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
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
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
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
	nodeFactoryAddr, _, _, err := ethbridgetestcontracts.DeployNodeFactory(auth, client)
	test.FailIfError(t, err)

	rollupAddr, _, _, err := ethbridgecontracts.DeployRollup(auth, client)
	test.FailIfError(t, err)

	bridgeCreatorAddr, _, _, err := ethbridgecontracts.DeployBridgeCreator(auth, client)
	test.FailIfError(t, err)

	_, _, rollupCreator, err := ethbridgetestcontracts.DeployRollupCreatorNoProxy(auth, client)
	test.FailIfError(t, err)
	client.Commit()

	_, err = rollupCreator.SetTemplates(auth, bridgeCreatorAddr, rollupAddr, challengeFactoryAddr, nodeFactoryAddr)
	test.FailIfError(t, err)
	client.Commit()

	tx, err := rollupCreator.CreateRollupNoProxy(
		auth,
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
	sequencerDelayBlocks := big.NewInt(60)
	sequencerDelaySeconds := big.NewInt(900)
	var extraConfig []byte

	clnt, pks := test.SimulatedBackend(t)
	auth := bind.NewKeyedTransactor(pks[0])
	sequencer := common.NewAddressFromEth(auth.From)
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
		owner,
		sequencer,
		sequencerDelayBlocks,
		sequencerDelaySeconds,
		extraConfig,
	)

	seqMon, shutdown := monitor.PrepareArbCore(t)
	defer shutdown()

	otherMon, shutdown2 := monitor.PrepareArbCore(t)
	defer shutdown2()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rollup, err := ethbridge.NewRollupWatcher(rollupAddr, client)
	test.FailIfError(t, err)

	seqInboxAddr, err := rollup.SequencerBridge(ctx)
	test.FailIfError(t, err)

	seqInbox, err := ethbridgecontracts.NewSequencerInbox(seqInboxAddr.ToEthAddress(), client)
	test.FailIfError(t, err)

	dummySequencerFeed := make(chan broadcaster.BroadcastFeedMessage)
	dummyDataSigner := func([]byte) ([]byte, error) { return make([]byte, 0), nil }

	_, err = seqMon.StartInboxReader(ctx, client, common.NewAddressFromEth(rollupAddr), nil, dummySequencerFeed)
	test.FailIfError(t, err)

	_, err = otherMon.StartInboxReader(ctx, client, common.NewAddressFromEth(rollupAddr), nil, dummySequencerFeed)
	test.FailIfError(t, err)

	client.Commit()
	time.Sleep(time.Second)

	batcher, err := NewSequencerBatcher(
		ctx,
		seqMon.Core,
		message.ChainAddressToID(common.NewAddressFromEth(rollupAddr)),
		seqMon.Reader,
		client,
		big.NewInt(1),
		seqInbox,
		auth,
        dummyDataSigner,
	)
	test.FailIfError(t, err)
	batcher.logBatchGasCosts = true
	batcher.chainTimeCheckInterval = time.Millisecond * 10
	go batcher.Start(ctx)
	client.Commit()
	time.Sleep(time.Second)
	client.Commit()

	for _, totalCount := range []int{1, 10, 100} {
		for _, dataSizePerTx := range []int{1, 10, 100} {
			txs := generateTxs(t, totalCount, dataSizePerTx, message.ChainAddressToID(common.NewAddressFromEth(rollupAddr)))
			for _, tx := range txs {
				if err := batcher.SendTransaction(ctx, tx); err != nil {
					t.Fatal(err)
				}
			}
			client.Commit()
			<-time.After(time.Second)
			client.Commit()
		}
	}

	msgCount1, err := seqMon.Core.GetMessageCount()
	test.FailIfError(t, err)
	if msgCount1.Cmp(big.NewInt(668)) < 0 {
		t.Error("Not enough messages, only got", msgCount1.String())
	}

	timeout := time.Now().Add(time.Second * 5)
	for {
		msgCount2, err := otherMon.Core.GetMessageCount()
		test.FailIfError(t, err)
		t.Logf("%v/%v", msgCount2.String(), msgCount1.String())
		if msgCount2.Cmp(msgCount1) == 0 {
			break
		}
		time.Sleep(time.Millisecond * 100)
		client.Commit()

		if time.Now().After(timeout) {
			t.Fatal("Exceeded message delivery timeout")
		}
	}
}
