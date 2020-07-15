package chainobserver

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
	"math/big"
	"testing"
)

var tester *ethbridgetestcontracts.RollupTester
var dummyRollupAddress = common.Address{1}
var auth *bind.TransactOpts

func TestMainSetup(m *testing.T) {
	client, auths := test.SimulatedBackend()
	auth = auths[0]

	_, machineTx, deployedArbRollup, err := ethbridgetestcontracts.DeployRollupTester(
		auth,
		client,
	)
	if err != nil {
		log.Fatal(err)
	}
	client.Commit()
	_, err = ethbridge.WaitForReceiptWithResults(
		context.Background(),
		client,
		auth.From,
		machineTx,
		"deployedMachineTester",
	)
	if err != nil {
		log.Fatal(err)
	}

	tester = deployedArbRollup
}

func TestComputePrevLeaf(t *testing.T) {
	chain, err := setUpChain(dummyRollupAddress, "dummy", contractPath)
	if err != nil {
		t.Fatal(err)
	}

	assertion := chain.PrepareAssertion()

	bridgeHash, _, err := tester.ComputePrevLeaf(
		nil,
		assertion.GetAssertionParams(),
		assertion.BeforeState.InboxCount,
		assertion.Prev.Deadline().Val,
		uint32(assertion.Prev.LinkType()),
		assertion.Params.NumSteps,
		assertion.Params.ImportedMessageCount,
		assertion.Claim.AssertionStub.DidInboxInsn,
		assertion.Claim.AssertionStub.NumGas)
	if err != nil {
		t.Fatal(err)
	}

	if assertion.Prev.Hash().ToEthHash() != bridgeHash {
		t.Error(bridgeHash)
	}
}

func randomAssertion() *protocol.ExecutionAssertion {
	results := make([]*evm.Result, 0, 5)
	messages := make([]value.Value, 0)
	messages = append(messages, message.NewInboxMessage(
		message.Eth{
			Dest:  common.Address{},
			Value: big.NewInt(75),
		},
		common.NewAddressFromEth(auth.From),
		big.NewInt(0),
		message.NewRandomChainTime(),
	).AsValue())
	for i := int32(0); i < 5; i++ {
		stop := evm.NewRandomResult(message.NewRandomEth(), 2)
		results = append(results, stop)
		messages = append(messages, message.NewRandomInboxMessage(message.NewRandomEth()).AsValue())
	}

	return evm.NewRandomEVMAssertion(results, messages)
}

func TestGenerateInvalidMsgLeaf(t *testing.T) {
	chain, err := setUpChain(dummyRollupAddress, "dummy", contractPath)
	if err != nil {
		t.Fatal(err)
	}

	prevNode := chain.NodeGraph.LatestConfirmed()
	assertion := randomAssertion()

	newNode := structures.NewRandomInvalidNodeFromValidPrev(prevNode, assertion, valprotocol.InvalidMessagesChildType, chain.GetChainParams())

	prepared := chain.PrepareAssertion()
	prepared.Assertion = assertion
	prepared.Claim.AssertionStub = valprotocol.NewExecutionAssertionStubFromAssertion(assertion)

	bridgeHash, _, err := tester.ComputePrevLeaf(
		nil,
		prepared.GetAssertionParams(),
		prepared.BeforeState.InboxCount,
		prepared.Prev.Deadline().Val,
		uint32(prepared.Prev.LinkType()),
		prepared.Params.NumSteps,
		prepared.Params.ImportedMessageCount,
		prepared.Claim.AssertionStub.DidInboxInsn,
		prepared.Claim.AssertionStub.NumGas)
	if err != nil {
		t.Fatal(err)
	}

	if newNode.PrevHash().ToEthHash() != bridgeHash {
		t.Error(bridgeHash)
	}

	invalidMsgHash, err := tester.ChildNodeHash(
		nil,
		newNode.PrevHash(),
		newNode.Deadline().Val,
		newNode.NodeDataHash(),
		new(big.Int).SetUint64(uint64(valprotocol.InvalidMessagesChildType)),
		newNode.VMProtoData().Hash())

	if newNode.Hash().ToEthHash() != invalidMsgHash {
		fmt.Println(bridgeHash)
		fmt.Println(newNode.Hash().ToEthHash())
		t.Error(bridgeHash)
	}
}

func TestGenerateInvalidInboxLeaf(t *testing.T) {
	chain, err := setUpChain(dummyRollupAddress, "dummy", contractPath)
	if err != nil {
		t.Fatal(err)
	}

	prevNode := chain.NodeGraph.LatestConfirmed()
	assertion := randomAssertion()
	newNode := structures.NewRandomInvalidNodeFromValidPrev(prevNode, assertion, valprotocol.InvalidInboxTopChildType, chain.GetChainParams())

	prepared := chain.PrepareAssertion()
	prepared.Assertion = assertion
	prepared.Claim.AssertionStub = valprotocol.NewExecutionAssertionStubFromAssertion(assertion)

	bridgeHash, _, err := tester.ComputePrevLeaf(
		nil,
		prepared.GetAssertionParams(),
		prepared.BeforeState.InboxCount,
		prepared.Prev.Deadline().Val,
		uint32(prepared.Prev.LinkType()),
		prepared.Params.NumSteps,
		prepared.Params.ImportedMessageCount,
		prepared.Claim.AssertionStub.DidInboxInsn,
		prepared.Claim.AssertionStub.NumGas)
	if err != nil {
		t.Fatal(err)
	}

	if newNode.PrevHash().ToEthHash() != bridgeHash {
		t.Error(bridgeHash)
	}

	invalidInboxHash, err := tester.ChildNodeHash(
		nil,
		newNode.PrevHash(),
		newNode.Deadline().Val,
		newNode.NodeDataHash(),
		new(big.Int).SetUint64(uint64(valprotocol.InvalidInboxTopChildType)),
		newNode.VMProtoData().Hash())

	if newNode.Hash().ToEthHash() != invalidInboxHash {
		fmt.Println(bridgeHash)
		fmt.Println(newNode.Hash().ToEthHash())
		t.Error(bridgeHash)
	}
}

func TestGenerateInvalidExecutionLeaf(t *testing.T) {
	chain, err := setUpChain(dummyRollupAddress, "dummy", contractPath)
	if err != nil {
		t.Fatal(err)
	}

	prevNode := chain.NodeGraph.LatestConfirmed()
	assertion := randomAssertion()
	newNode := structures.NewRandomInvalidNodeFromValidPrev(prevNode, assertion, valprotocol.InvalidExecutionChildType, chain.GetChainParams())

	prepared := chain.PrepareAssertion()
	prepared.Assertion = assertion
	prepared.Claim.AssertionStub = valprotocol.NewExecutionAssertionStubFromAssertion(assertion)

	bridgeHash, _, err := tester.ComputePrevLeaf(
		nil,
		prepared.GetAssertionParams(),
		prepared.BeforeState.InboxCount,
		prepared.Prev.Deadline().Val,
		uint32(prepared.Prev.LinkType()),
		prepared.Params.NumSteps,
		prepared.Params.ImportedMessageCount,
		prepared.Claim.AssertionStub.DidInboxInsn,
		prepared.Claim.AssertionStub.NumGas)
	if err != nil {
		t.Fatal(err)
	}

	if newNode.PrevHash().ToEthHash() != bridgeHash {
		t.Error(bridgeHash)
	}

	invalidExecutionHash, err := tester.ChildNodeHash(
		nil,
		newNode.PrevHash(),
		newNode.Deadline().Val,
		newNode.NodeDataHash(),
		new(big.Int).SetUint64(uint64(valprotocol.InvalidExecutionChildType)),
		newNode.VMProtoData().Hash())

	if newNode.Hash().ToEthHash() != invalidExecutionHash {
		fmt.Println(bridgeHash)
		fmt.Println(newNode.Hash().ToEthHash())
		t.Error(bridgeHash)
	}
}
