package challenge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
)

func TestInboxConsistencyChallenge(t *testing.T) {
	ctx := context.Background()
	client, asserter, challenger, challengeAddress := initializeChallengeTest(
		t,
		[32]byte{},
		[32]byte{},
		[32]byte{},
		big.NewInt(0),
		big.NewInt(0),
	)

	challengerChallenge, err := ethbridge.NewChallenge(challengeAddress, client, ethbridge.NewTransactAuth(challenger))
	test.FailIfError(t, err)
	asserterChallenge, err := ethbridge.NewChallenge(challengeAddress, client, ethbridge.NewTransactAuth(asserter))
	test.FailIfError(t, err)

	kind1, err := challengerChallenge.Kind(ctx)
	test.FailIfError(t, err)
	kind2, err := asserterChallenge.Kind(ctx)
	test.FailIfError(t, err)
	if kind1 != kind2 {
		t.Fatal("kind doesn't match")
	}
}

func initializeChallengeTest(
	t *testing.T,
	inboxConsistencyHash [32]byte,
	inboxDeltaHash [32]byte,
	executionHash [32]byte,
	executionCheckTimeBlocks *big.Int,
	challengePeriodBlocks *big.Int,
) (*ethutils.SimulatedEthClient, *bind.TransactOpts, *bind.TransactOpts, ethcommon.Address) {
	clnt, pks := test.SimulatedBackend()
	deployer := bind.NewKeyedTransactor(pks[0])
	asserter := bind.NewKeyedTransactor(pks[1])
	challenger := bind.NewKeyedTransactor(pks[2])
	client := &ethutils.SimulatedEthClient{SimulatedBackend: clnt}
	osp1Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProof(deployer, client)
	test.FailIfError(t, err)
	osp2Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProof2(deployer, client)
	test.FailIfError(t, err)
	_, _, tester, err := ethbridgetestcontracts.DeployChallengeTester(deployer, client, osp1Addr, osp2Addr)
	test.FailIfError(t, err)
	_, err = tester.StartChallenge(
		deployer,
		inboxConsistencyHash,
		inboxDeltaHash,
		executionHash,
		executionCheckTimeBlocks,
		asserter.From,
		challenger.From,
		challengePeriodBlocks,
	)
	test.FailIfError(t, err)
	client.Commit()
	challengeAddress, err := tester.Challenge(&bind.CallOpts{})
	test.FailIfError(t, err)
	return client, asserter, challenger, challengeAddress
}
