package challenge

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/core"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
)

func logInboxAccs(t *testing.T, lookupMock *core.ValidatorLookupMock) {
	for _, inboxAccHash := range lookupMock.InboxAccs {
		t.Log(inboxAccHash)
	}
}

func TestInboxConsistencyChallenge(t *testing.T) {
	ctx := context.Background()

	mach, err := cmachine.New(arbos.Path())
	test.FailIfError(t, err)

	msg1 := inbox.NewRandomInboxMessage()
	msg2 := inbox.NewRandomInboxMessage()

	asserterLookup := core.NewValidatorLookupMock(mach)
	asserterLookup.AddMessage(inbox.NewRandomInboxMessage())
	asserterLookup.AddMessage(msg2)
	asserterLookup.AddMessage(inbox.NewRandomInboxMessage())
	challengerLookup := core.NewValidatorLookupMock(mach)
	challengerLookup.AddMessage(msg1)
	challengerLookup.AddMessage(msg2)
	challengerLookup.AddMessage(inbox.NewRandomInboxMessage())

	logInboxAccs(t, challengerLookup)

	prevState := &core.NodeState{
		ProposedBlock:  big.NewInt(0),
		TotalGasUsed:   big.NewInt(0),
		MachineHash:    mach.Hash(),
		InboxHash:      common.Hash{},
		InboxCount:     big.NewInt(0),
		TotalSendCount: big.NewInt(0),
		TotalLogCount:  big.NewInt(0),
		InboxMaxCount:  big.NewInt(0),
	}

	inboxAcc, err := asserterLookup.GetInboxAcc(big.NewInt(1))
	test.FailIfError(t, err)

	assertionInfo := &core.AssertionInfo{
		InboxDelta: core.CalculateInboxDeltaAcc([]inbox.InboxMessage{msg1}),
		ExecInfo: &core.ExecutionInfo{
			BeforeMachineHash: common.Hash{},
			InboxMessagesRead: big.NewInt(1),
			GasUsed:           big.NewInt(0),
			SendAcc:           common.Hash{},
			SendCount:         big.NewInt(0),
			LogAcc:            common.Hash{},
			LogCount:          big.NewInt(0),
			AfterMachineHash:  common.Hash{},
		},
		AfterInboxHash: inboxAcc,
	}

	assertion := &core.Assertion{
		PrevState:     prevState,
		AssertionInfo: assertionInfo,
	}

	inboxMaxCount := big.NewInt(3)
	inboxTopAcc, err := challengerLookup.GetInboxAcc(inboxMaxCount)
	test.FailIfError(t, err)
	challengedNode := &core.NodeInfo{
		NodeNum: big.NewInt(1),
		BlockProposed: &common.BlockId{
			Height:     common.NewTimeBlocks(common.RandBigInt()),
			HeaderHash: common.RandHash(),
		},
		Assertion:     assertion,
		InboxMaxCount: inboxMaxCount,
		InboxMaxHash:  inboxTopAcc,
	}

	arbGasSpeedLimitPerBlock := big.NewInt(100000)
	challengePeriodBlocks := big.NewInt(100)

	client, asserterAuth, challengerAuth, challengeAddress := initializeChallengeTest(t, challengedNode, arbGasSpeedLimitPerBlock, challengePeriodBlocks)

	challengerChallenge, err := ethbridge.NewChallenge(challengeAddress, client, ethbridge.NewTransactAuth(challengerAuth))
	test.FailIfError(t, err)
	asserterChallenge, err := ethbridge.NewChallenge(challengeAddress, client, ethbridge.NewTransactAuth(asserterAuth))
	test.FailIfError(t, err)

	kind1, err := challengerChallenge.Kind(ctx)
	test.FailIfError(t, err)
	kind2, err := asserterChallenge.Kind(ctx)
	test.FailIfError(t, err)
	if kind1 != kind2 {
		t.Fatal("kind doesn't match")
	}

	challenger := NewChallenger(challengerChallenge, challengerLookup, challengedNode)
	//asserter := NewChallenger(challengerChallenge, lookup, assertion)

	_, err = challenger.handleConflict(ctx)
	test.FailIfError(t, err)
}

func initializeChallengeTest(
	t *testing.T,
	nd *core.NodeInfo,
	arbGasLimitPerBlock *big.Int,
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
		nd.Assertion.InboxConsistencyHash(nd.InboxMaxHash, nd.InboxMaxCount),
		nd.Assertion.InboxDeltaHash(),
		nd.Assertion.ExecutionHash(),
		nd.Assertion.CheckTime(arbGasLimitPerBlock),
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
