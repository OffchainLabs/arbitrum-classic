package staker

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
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

	_, _, rollupCreator, err := ethbridgetestcontracts.DeployRollupCreatorNoProxy(auth, client)
	test.FailIfError(t, err)
	client.Commit()

	_, err = rollupCreator.SetTemplates(auth, rollupAddr, challengeFactoryAddr, nodeFactoryAddr)
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

func TestStaker(t *testing.T) {
	ctx := context.Background()

	mach, err := cmachine.New(arbos.Path())
	test.FailIfError(t, err)

	confirmPeriodBlocks := big.NewInt(100)
	extraChallengeTimeBlocks := big.NewInt(0)
	arbGasSpeedLimitPerBlock := big.NewInt(1000000000)
	baseStake := big.NewInt(100)
	var stakeToken common.Address
	var owner common.Address
	var extraConfig []byte

	clnt, pks := test.SimulatedBackend()
	auth := bind.NewKeyedTransactor(pks[0])
	auth2 := bind.NewKeyedTransactor(pks[0])
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

	validatorAddress, _, _, err := ethbridgecontracts.DeployValidator(auth, client)
	test.FailIfError(t, err)

	validatorAddress2, _, _, err := ethbridgecontracts.DeployValidator(auth2, client)
	test.FailIfError(t, err)

	client.Commit()

	val, err := ethbridge.NewValidator(validatorAddress, rollupAddr, client, ethbridge.NewTransactAuth(auth))
	test.FailIfError(t, err)

	val2, err := ethbridge.NewValidator(validatorAddress2, rollupAddr, client, ethbridge.NewTransactAuth(auth2))
	test.FailIfError(t, err)

	core, shutdown := challenge.PrepareTestArbCore(t, []inbox.InboxMessage{challenge.MakeTestInitMsg()})
	defer shutdown()

	staker, err := NewStaker(ctx, core, client, val, common.NewAddressFromEth(validatorUtilsAddr), MakeNodesStrategy)
	test.FailIfError(t, err)

	faultyCore := challenge.NewFaultyCore(core, challenge.FaultConfig{DistortMachineAtGas: big.NewInt(10000)})

	faultyStaker, err := NewStaker(ctx, faultyCore, client, val2, common.NewAddressFromEth(validatorUtilsAddr), MakeNodesStrategy)
	test.FailIfError(t, err)

	bridgeAddr, err := staker.rollup.Bridge(ctx)
	test.FailIfError(t, err)
	bridge, err := ethbridge.NewBridgeWatcher(bridgeAddr.ToEthAddress(), client)
	test.FailIfError(t, err)
	_, err = NewInboxReader(ctx, bridge, core)
	test.FailIfError(t, err)

	for i := 1; i <= 30; i++ {
		msgCount, err := core.GetMessageCount()
		test.FailIfError(t, err)
		if msgCount.Cmp(big.NewInt(1)) >= 0 {
			// We've found the inbox message
			break
		}
		if i == 100 {
			t.Fatal("Failed to load initializing message")
		}
		<-time.After(time.Second * 1)
	}

	for i := 0; i < 100; i++ {
		if (i % 2) == 0 {
			_, err := staker.Act(ctx)
			test.FailIfError(t, err)
		} else {
			_, err = faultyStaker.Act(ctx)
			test.FailIfError(t, err)
		}
		client.Commit()
	}

	stakerInfo, err := staker.rollup.StakerInfo(ctx, common.NewAddressFromEth(validatorAddress))
	test.FailIfError(t, err)

	if stakerInfo == nil {
		t.Fatal("Staker isn't staked")
	}

	if stakerInfo.CurrentChallenge != nil || stakerInfo.AmountStaked.Cmp(big.NewInt(0)) == 0 {
		t.Fatal("Staker didn't resolve challenge")
	}

	if stakerInfo.LatestStakedNode.Cmp(big.NewInt(0)) == 0 {
		t.Fatal("Staker didn't stake on node")
	}

	latestConfirmed, err := staker.rollup.LatestConfirmedNode(ctx)
	test.FailIfError(t, err)
	if latestConfirmed.Cmp(stakerInfo.LatestStakedNode) != 0 {
		t.Fatal("Staked node remains unconfirmed")
	}

	faultyStakerInfo, err := staker.rollup.StakerInfo(ctx, common.NewAddressFromEth(validatorAddress2))
	test.FailIfError(t, err)

	if faultyStakerInfo.AmountStaked.Cmp(big.NewInt(0)) > 0 {
		t.Fatal("Faulty staker still has stake")
	}
}
