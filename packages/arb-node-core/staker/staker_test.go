package staker

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"math/big"
	"testing"
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

	client.Commit()

	val, err := ethbridge.NewValidator(validatorAddress, rollupAddr, client, ethbridge.NewTransactAuth(auth))
	test.FailIfError(t, err)

	lookup := core.NewValidatorLookupMock(mach)

	staker, err := NewStaker(ctx, lookup, client, val, common.NewAddressFromEth(validatorUtilsAddr))
	test.FailIfError(t, err)

	for i := 0; i < 100; i++ {
		client.Commit()
	}

	err = staker.newStake(ctx)
	test.FailIfError(t, err)

	if staker.builder.TransactionCount() == 0 {
		t.Fatal("didn't place stake")
	}

	_, err = val.ExecuteTransactions(ctx, staker.builder)
	test.FailIfError(t, err)

	client.Commit()

	stakerInfo, err := staker.rollup.StakerInfo(ctx, common.NewAddressFromEth(validatorAddress))
	test.FailIfError(t, err)

	if stakerInfo.CurrentChallenge != nil {
		t.Fatal("shouldn't be in challenge")
	}
	if stakerInfo.LatestStakedNode.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("staked on wrong node")
	}
}
