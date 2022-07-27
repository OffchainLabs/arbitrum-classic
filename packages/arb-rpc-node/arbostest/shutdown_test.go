package arbostest

import (
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
	"testing"
)

func TestShutdown(t *testing.T) {
	ethDeposit := makeEthDeposit(sender, big.NewInt(1000))

	tx1 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        arbos.WithdrawEthData(common.RandAddress()),
	}

	messages := []message.Message{
		ethDeposit,
		message.NewSafeL2Message(tx1),
		message.ShutdownMessage{},
	}

	inboxMessages := makeSimpleInbox(t, messages)
	_, results, sends, _ := runBasicAssertion(t, inboxMessages)
	if len(sends) != 1 {
		t.Fatal("expected one send", len(sends))
	}
	txResults := extractTxResults(t, results)
	if len(txResults) != 2 {
		t.Fatal("unexpected log count ", len(txResults), "instead of", 2)
	}
	allResultsSucceeded(t, txResults)

	outMsg, err := message.NewOutMessageFromBytes(sends[0])
	failIfError(t, err)

	sendRoot, ok := outMsg.(*message.SendMessageRoot)
	if !ok {
		t.Fatal("expected send root")
	}
	if sendRoot.NumInBatch.Uint64() != 1 {
		t.Fatal("expected one message in batch")
	}

	var merkleRootRes *evm.MerkleRootResult
	for _, res := range results {
		res, ok := res.(*evm.MerkleRootResult)
		if !ok {
			continue
		}
		if merkleRootRes != nil {
			t.Fatal("expected single merkle root result")
		}
		merkleRootRes = res
	}
	if merkleRootRes == nil {
		t.Fatal("found no merkle root results")
	}
	if merkleRootRes.BatchNumber.Cmp(sendRoot.BatchNumber) != 0 {
		t.Error("batch number didn't match")
	}
	if merkleRootRes.NumInBatch.Cmp(sendRoot.NumInBatch) != 0 {
		t.Error("batch count didn't match")
	}
}
