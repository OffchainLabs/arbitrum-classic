package dev

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func TestTopLevelCall(t *testing.T) {
	skipBelowVersion(t, 30)
	config := protocol.ChainParams{
		GracePeriod:               common.NewTimeBlocksInt(3),
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}
	senderKey, err := crypto.GenerateKey()
	test.FailIfError(t, err)

	upgraderAuth, upgraderAccount := OwnerAuthPair(t, nil)

	backend, _, srv, cancelDevNode := NewSimpleTestDevNode(t, config, upgraderAccount)
	defer cancelDevNode()

	senderAuth, err := bind.NewKeyedTransactorWithChainID(senderKey, backend.chainID)
	test.FailIfError(t, err)

	client := web3.NewEthClient(srv, true)

	conAddr, _, con, err := arbostestcontracts.DeployTopLevel(senderAuth, client)
	test.FailIfError(t, err)
	topABI, err := abi.JSON(strings.NewReader(arbostestcontracts.TopLevelABI))
	test.FailIfError(t, err)

	checkTopLevel := func(tx ethcommon.Hash) bool {
		receipt, err := client.TransactionReceipt(context.Background(), tx)
		test.FailIfError(t, err)
		if receipt == nil {
			t.Fatal("expected receipt")
		}
		if len(receipt.Logs) != 1 {
			t.Fatal("unexpected log count")
		}
		if len(receipt.Logs[0].Topics) != 2 {
			t.Fatal("unexpected topic count")
		}
		isTopLevel := new(big.Int).SetBytes(receipt.Logs[0].Topics[1][:])
		if isTopLevel.Cmp(big.NewInt(1)) == 0 {
			return true
		} else if isTopLevel.Cmp(big.NewInt(0)) == 0 {
			return false
		}
		t.Fatal("bad event value")
		return false
	}

	// Basic Tx

	tx, err := con.IsTopLevel(senderAuth)
	test.FailIfError(t, err)
	if !checkTopLevel(tx.Hash()) {
		t.Error("expected top level")
	}

	tx, err = con.NestedNotTop(senderAuth)
	test.FailIfError(t, err)
	if checkTopLevel(tx.Hash()) {
		t.Error("expected not top level")
	}

	createTickets := func(maxGas, gasBig *big.Int) (common.Hash, common.Hash) {
		retryableTx := message.RetryableTx{
			Destination:       common.NewAddressFromEth(conAddr),
			Value:             big.NewInt(0),
			Deposit:           big.NewInt(1000000000),
			MaxSubmissionCost: big.NewInt(30),
			CreditBack:        common.RandAddress(),
			Beneficiary:       common.RandAddress(),
			MaxGas:            maxGas,
			GasPriceBid:       gasBig,
			Data:              topABI.Methods["isTopLevel"].ID,
		}
		requestId1, err := backend.AddInboxMessage(retryableTx, common.RandAddress())
		test.FailIfError(t, err)
		ticketId1 := hashing.SoliditySHA3(hashing.Bytes32(requestId1), hashing.Uint256(big.NewInt(0)))

		retryableTx2 := message.RetryableTx{
			Destination:       common.NewAddressFromEth(conAddr),
			Value:             big.NewInt(0),
			Deposit:           big.NewInt(1000000000),
			MaxSubmissionCost: big.NewInt(30),
			CreditBack:        common.RandAddress(),
			Beneficiary:       common.RandAddress(),
			MaxGas:            maxGas,
			GasPriceBid:       gasBig,
			Data:              topABI.Methods["nestedNotTop"].ID,
		}
		requestId2, err := backend.AddInboxMessage(retryableTx2, common.RandAddress())
		test.FailIfError(t, err)
		ticketId2 := hashing.SoliditySHA3(hashing.Bytes32(requestId2), hashing.Uint256(big.NewInt(0)))
		return ticketId1, ticketId2
	}

	if doUpgrade {
		UpgradeTestDevNode(t, backend, srv, upgraderAuth)
	}

	// Immediate Redeem
	ticketId1, ticketId2 := createTickets(big.NewInt(1000000), big.NewInt(10))
	if !checkTopLevel(ticketId1.ToEthHash()) {
		t.Error("expected top level")
	}
	if checkTopLevel(ticketId2.ToEthHash()) {
		t.Error("expected not top level")
	}

	retryable, err := arboscontracts.NewArbRetryableTx(arbos.ARB_RETRYABLE_ADDRESS, client)
	test.FailIfError(t, err)

	// Delayed Redeem
	ticketId3, ticketId4 := createTickets(big.NewInt(0), big.NewInt(0))
	_, err = retryable.Redeem(senderAuth, ticketId3)
	test.FailIfError(t, err)
	_, err = retryable.Redeem(senderAuth, ticketId4)
	test.FailIfError(t, err)
	if !checkTopLevel(ticketId3.ToEthHash()) {
		t.Error("expected top level")
	}
	if checkTopLevel(ticketId4.ToEthHash()) {
		t.Error("expected not top level")
	}
}
