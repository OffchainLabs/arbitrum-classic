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

package dev

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

func setupTest(t *testing.T) (
	common.Address,
	*bind.TransactOpts,
	*bind.TransactOpts,
	*aggregator.Server,
	*Backend,
	func(),
) {
	config := protocol.ChainParams{
		StakeRequirement:          big.NewInt(10),
		StakeToken:                common.Address{},
		GracePeriod:               common.NewTimeBlocksInt(3),
		MaxExecutionSteps:         10000000000,
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}

	backend, _, srv, cancelDevNode := NewTestDevNode(t, *arbosfile, config, common.RandAddress(), nil)

	privkey, err := crypto.GenerateKey()
	test.FailIfError(t, err)
	otherAuth := bind.NewKeyedTransactor(privkey)

	privkey2, err := crypto.GenerateKey()
	test.FailIfError(t, err)
	beneficiaryAuth := bind.NewKeyedTransactor(privkey2)

	sender := common.RandAddress()

	deposit := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(1000000),
				GasPriceBid: big.NewInt(0),
				DestAddress: common.NewAddressFromEth(otherAuth.From),
				Payment:     big.NewInt(100),
				Data:        nil,
			},
		}),
	}
	_, err = backend.AddInboxMessage(deposit, common.RandAddress())
	test.FailIfError(t, err)

	return sender, beneficiaryAuth, otherAuth, srv, backend, cancelDevNode
}

func setupTicket(t *testing.T, backend *Backend, sender, destination common.Address, data []byte, beneficiary common.Address) (message.RetryableTx, common.Hash) {
	retryableTx := message.RetryableTx{
		Destination:       destination,
		Value:             big.NewInt(20),
		Deposit:           big.NewInt(100),
		MaxSubmissionCost: big.NewInt(30),
		CreditBack:        common.RandAddress(),
		Beneficiary:       beneficiary,
		MaxGas:            big.NewInt(0),
		GasPriceBid:       big.NewInt(0),
		Data:              data,
	}

	requestId, err := backend.AddInboxMessage(retryableTx, sender)
	test.FailIfError(t, err)

	return retryableTx, requestId
}

func TestRetryableRedeem(t *testing.T) {
	sender, beneficiaryAuth, otherAuth, srv, backend, closeFunc := setupTest(t)
	defer closeFunc()

	client := web3.NewEthClient(srv, true)
	retryable, err := arboscontracts.NewArbRetryableTx(arbos.ARB_RETRYABLE_ADDRESS, client)
	test.FailIfError(t, err)

	simpleABI, err := abi.JSON(strings.NewReader(arbostestcontracts.SimpleABI))
	test.FailIfError(t, err)

	dest, _, _, err := arbostestcontracts.DeploySimple(otherAuth, client)
	test.FailIfError(t, err)

	retryableTx, requestId := setupTicket(t, backend, sender, common.NewAddressFromEth(dest), simpleABI.Methods["exists"].ID, common.NewAddressFromEth(beneficiaryAuth.From))
	ticketId := hashing.SoliditySHA3(hashing.Bytes32(requestId), hashing.Uint256(big.NewInt(0)))

	redeemReceipt, err := client.TransactionReceipt(context.Background(), requestId.ToEthHash())
	test.FailIfError(t, err)

	if redeemReceipt == nil || redeemReceipt.Status != 1 {
		t.Fatal("retryable tx failed")
	}

	finalReceipt, err := client.TransactionReceipt(context.Background(), ticketId.ToEthHash())
	test.FailIfError(t, err)

	if finalReceipt != nil {
		t.Fatal("shouldn't have receipt yet")
	}

	creationBlock, err := backend.db.GetBlockWithHash(common.NewHashFromEth(redeemReceipt.BlockHash))
	test.FailIfError(t, err)

	lifetime, err := retryable.GetLifetime(&bind.CallOpts{})
	test.FailIfError(t, err)

	timeout, err := retryable.GetTimeout(&bind.CallOpts{}, ticketId)
	test.FailIfError(t, err)

	if timeout.Uint64() != creationBlock.Header.Time+lifetime.Uint64() {
		t.Fatal("wrong timeout")
	}

	price, nextUpdateTimestamp, err := retryable.GetKeepalivePrice(&bind.CallOpts{}, ticketId)
	test.FailIfError(t, err)
	if price.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("wrong price")
	}
	t.Log("nextUpdateTimestamp", nextUpdateTimestamp)

	beneficiary, err := retryable.GetBeneficiary(&bind.CallOpts{}, ticketId)
	test.FailIfError(t, err)
	if beneficiary != retryableTx.Beneficiary.ToEthAddress() {
		t.Fatal("wrong beneficiary")
	}

	correctSenderBalance := new(big.Int).Sub(retryableTx.Deposit, retryableTx.Value)
	correctSenderBalance = correctSenderBalance.Sub(correctSenderBalance, retryableTx.MaxSubmissionCost)

	balanceCheck(t, srv, sender, retryableTx, correctSenderBalance, big.NewInt(0), retryableTx.MaxSubmissionCost, big.NewInt(0))

	_, err = retryable.Keepalive(otherAuth, ticketId)
	test.FailIfError(t, err)

	newTimeout, err := retryable.GetTimeout(&bind.CallOpts{}, ticketId)
	test.FailIfError(t, err)

	if newTimeout.Uint64() != creationBlock.Header.Time+lifetime.Uint64()*2 {
		t.Fatal("wrong timeout")
	}

	tx, err := retryable.Redeem(otherAuth, ticketId)
	test.FailIfError(t, err)

	redeemReceipt, err = client.TransactionReceipt(context.Background(), tx.Hash())
	test.FailIfError(t, err)

	if redeemReceipt == nil || redeemReceipt.Status != 1 {
		t.Fatal("redeem tx failed")
	}

	if len(redeemReceipt.Logs) == 0 {
		t.Fatal("should have at least one log")
	}

	if redeemReceipt.Logs[len(redeemReceipt.Logs)-1].Topics[0] != arbos.RetryRedeemedEvent.ID {
		t.Fatal("wrong log topic")
	}

	finalReceipt, err = client.TransactionReceipt(context.Background(), ticketId.ToEthHash())
	test.FailIfError(t, err)

	if finalReceipt == nil || finalReceipt.Status != 1 {
		t.Fatal("final tx failed")
	}

	redeemRequest, err := backend.db.GetRequest(common.NewHashFromEth(tx.Hash()))
	test.FailIfError(t, err)

	if len(redeemRequest.ReturnData) != 0 {
		t.Error("expected redeem to have no return data")
	}

	balanceCheck(t, srv, sender, retryableTx, correctSenderBalance, big.NewInt(0), retryableTx.MaxSubmissionCost, retryableTx.Value)

	var txLogs []*types.Log
	if arbosVersion < 6 {
		txLogs = redeemReceipt.Logs[:len(redeemReceipt.Logs)-1]
	} else {
		txLogs = finalReceipt.Logs

		finalRequest, err := backend.db.GetRequest(ticketId)
		test.FailIfError(t, err)

		if len(finalRequest.ReturnData) != 32 {
			t.Error("expected final tx to have 32 bytes of return data but got", len(finalRequest.ReturnData))
		} else {
			ret := new(big.Int).SetBytes(finalRequest.ReturnData)
			if ret.Cmp(big.NewInt(10)) != 0 {
				t.Error("incorrect return data")
			}
		}
	}
	if len(txLogs) != 1 {
		t.Fatal("wrong log count", len(txLogs))
	}
	if txLogs[0].Topics[0] != simpleABI.Events["TestEvent"].ID {
		t.Fatal("wrong event topic")
	}
}

func TestRetryableCancel(t *testing.T) {
	sender, beneficiaryAuth, otherAuth, srv, backend, closeFunc := setupTest(t)
	defer closeFunc()
	retryableTx, requestId := setupTicket(t, backend, sender, common.RandAddress(), nil, common.NewAddressFromEth(beneficiaryAuth.From))
	ticketId := hashing.SoliditySHA3(hashing.Bytes32(requestId), hashing.Uint256(big.NewInt(0)))

	client := web3.NewEthClient(srv, true)
	retryable, err := arboscontracts.NewArbRetryableTx(arbos.ARB_RETRYABLE_ADDRESS, client)
	test.FailIfError(t, err)

	_, err = retryable.Cancel(otherAuth, ticketId)
	if err == nil {
		t.Fatal("cancel should fail from non beneficiary")
	}

	tx, err := retryable.Cancel(beneficiaryAuth, ticketId)
	test.FailIfError(t, err)

	txReceipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	test.FailIfError(t, err)

	if txReceipt == nil || txReceipt.Status != 1 {
		t.Fatal("cancel tx failed")
	}

	if len(txReceipt.Logs) != 1 {
		t.Fatal("wrong log count")
	}

	if txReceipt.Logs[0].Topics[0] != arbos.RetryCanceledEvent.ID {
		t.Fatal("wrong log topic")
	}

	correctSenderBalance := new(big.Int).Sub(retryableTx.Deposit, retryableTx.Value)
	correctSenderBalance = correctSenderBalance.Sub(correctSenderBalance, retryableTx.MaxSubmissionCost)

	correctBeneficiaryValue := retryableTx.Value

	balanceCheck(t, srv, sender, retryableTx, correctSenderBalance, correctBeneficiaryValue, retryableTx.MaxSubmissionCost, big.NewInt(0))
}

func TestRetryableTimeout(t *testing.T) {
	sender, beneficiaryAuth, _, srv, backend, closeFunc := setupTest(t)
	defer closeFunc()
	retryableTx, requestId := setupTicket(t, backend, sender, common.RandAddress(), nil, common.NewAddressFromEth(beneficiaryAuth.From))
	ticketId := hashing.SoliditySHA3(hashing.Bytes32(requestId), hashing.Uint256(big.NewInt(0)))

	client := web3.NewEthClient(srv, true)
	retryable, err := arboscontracts.NewArbRetryableTx(arbos.ARB_RETRYABLE_ADDRESS, client)
	test.FailIfError(t, err)

	_, err = retryable.GetBeneficiary(&bind.CallOpts{}, ticketId)
	test.FailIfError(t, err)

	timeout, err := retryable.GetTimeout(&bind.CallOpts{}, ticketId)
	test.FailIfError(t, err)

	lifetime, err := retryable.GetLifetime(&bind.CallOpts{})
	test.FailIfError(t, err)
	backend.l1Emulator.IncreaseTime(lifetime.Int64() * 10)

	_, err = backend.AddInboxMessage(message.NewSafeL2Message(message.HeartbeatMessage{}), common.RandAddress())
	test.FailIfError(t, err)

	latest := rpc.LatestBlockNumber
	l2BlockNum, err := srv.BlockNum(&latest)
	test.FailIfError(t, err)
	l2Block, err := srv.BlockInfoByNumber(l2BlockNum)
	test.FailIfError(t, err)
	if timeout.Uint64() >= l2Block.Header.Time {
		t.Fatal("should've moved forward more", l2Block.Header.Time, timeout.Uint64())
	}

	retryableTx2 := message.RetryableTx{
		Destination:       common.RandAddress(),
		Value:             big.NewInt(5),
		Deposit:           big.NewInt(20),
		MaxSubmissionCost: big.NewInt(10),
		CreditBack:        common.RandAddress(),
		Beneficiary:       common.RandAddress(),
		MaxGas:            big.NewInt(0),
		GasPriceBid:       big.NewInt(0),
		Data:              nil,
	}

	//// Send and cancel retryable to trigger pruning
	otherRequest, err := backend.AddInboxMessage(retryableTx2, common.RandAddress())
	test.FailIfError(t, err)

	txReceipt, err := client.TransactionReceipt(context.Background(), otherRequest.ToEthHash())
	test.FailIfError(t, err)

	if txReceipt == nil {
		t.Fatal("other retryable tx doesn't exist")
	}

	if txReceipt.Status != 1 {
		t.Fatal("other retryable tx failed")
	}

	otherTicket := hashing.SoliditySHA3(hashing.Bytes32(otherRequest), hashing.Uint256(big.NewInt(0)))
	_, err = retryable.GetBeneficiary(&bind.CallOpts{}, otherTicket)
	test.FailIfError(t, err)

	_, err = retryable.GetBeneficiary(&bind.CallOpts{}, ticketId)
	if err == nil {
		t.Fatal("should revert after timeout")
	}

	correctSenderBalance := new(big.Int).Sub(retryableTx.Deposit, retryableTx.Value)
	correctSenderBalance = correctSenderBalance.Sub(correctSenderBalance, retryableTx.MaxSubmissionCost)

	correctBeneficiaryValue := retryableTx.Value
	balanceCheck(t, srv, sender, retryableTx, correctSenderBalance, correctBeneficiaryValue, retryableTx.MaxSubmissionCost, big.NewInt(0))
}

func balanceCheck(
	t *testing.T,
	srv *aggregator.Server,
	sender common.Address,
	retryableTx message.RetryableTx,
	correctSenderBalance,
	correctBeneficiaryBalance,
	correctCreditBackBalance,
	correctDestinationBalance *big.Int,
) {
	t.Helper()
	snap, err := srv.PendingSnapshot()
	test.FailIfError(t, err)

	senderBalance, err := snap.GetBalance(sender)
	test.FailIfError(t, err)

	if senderBalance.Cmp(correctSenderBalance) != 0 {
		t.Error("unexpected sender balance", senderBalance, "instead of", correctSenderBalance)
	}

	beneficiaryBalance, err := snap.GetBalance(retryableTx.Beneficiary)
	test.FailIfError(t, err)

	if beneficiaryBalance.Cmp(correctBeneficiaryBalance) != 0 {
		t.Error("unexpected beneficiary balance", beneficiaryBalance, "instead of", correctBeneficiaryBalance)
	}

	creditBackBalance, err := snap.GetBalance(retryableTx.CreditBack)
	test.FailIfError(t, err)
	if creditBackBalance.Cmp(correctCreditBackBalance) != 0 {
		t.Error("unexpected credit back balance", creditBackBalance, "instead of", correctCreditBackBalance)
	}

	destinationBalance, err := snap.GetBalance(retryableTx.Destination)
	test.FailIfError(t, err)

	if destinationBalance.Cmp(correctDestinationBalance) != 0 {
		t.Error("unexpected destination balance")
	}
}
func TestRetryableReverted(t *testing.T) {
	sender, beneficiaryAuth, otherAuth, srv, backend, closeFunc := setupTest(t)
	defer closeFunc()

	client := web3.NewEthClient(srv, true)

	simpleABI, err := abi.JSON(strings.NewReader(arbostestcontracts.SimpleABI))
	test.FailIfError(t, err)

	dest, _, _, err := arbostestcontracts.DeploySimple(otherAuth, client)
	test.FailIfError(t, err)

	retryableTx := message.RetryableTx{
		Destination:       common.NewAddressFromEth(dest),
		Value:             big.NewInt(20),
		Deposit:           big.NewInt(100),
		MaxSubmissionCost: big.NewInt(30),
		CreditBack:        common.RandAddress(),
		Beneficiary:       common.NewAddressFromEth(beneficiaryAuth.From),
		MaxGas:            big.NewInt(0),
		GasPriceBid:       big.NewInt(0),
		Data:              simpleABI.Methods["reverts"].ID,
	}

	requestId, err := backend.AddInboxMessage(retryableTx, sender)
	test.FailIfError(t, err)

	ticketId := hashing.SoliditySHA3(hashing.Bytes32(requestId), hashing.Uint256(big.NewInt(0)))

	retryable, err := arboscontracts.NewArbRetryableTx(arbos.ARB_RETRYABLE_ADDRESS, client)
	test.FailIfError(t, err)

	_, err = retryable.Redeem(otherAuth, ticketId)
	if err == nil {
		t.Fatal("expected error from redeem")
	}

	if arbosVersion >= 9 && err.Error() != "failed to estimate gas needed: execution reverted: this is a test" {
		t.Error("wrong error message from redeem", err)
	}
	balanceCheck(t, srv, sender, retryableTx, big.NewInt(50), big.NewInt(0), big.NewInt(30), big.NewInt(0))
}

func TestRetryableWithReturnData(t *testing.T) {
	sender, beneficiaryAuth, otherAuth, srv, backend, closeFunc := setupTest(t)
	defer closeFunc()

	client := web3.NewEthClient(srv, true)

	simpleABI, err := abi.JSON(strings.NewReader(arbostestcontracts.SimpleABI))
	test.FailIfError(t, err)

	dest, _, _, err := arbostestcontracts.DeploySimple(otherAuth, client)
	test.FailIfError(t, err)

	retryableTx := message.RetryableTx{
		Destination:       common.NewAddressFromEth(dest),
		Value:             big.NewInt(20),
		Deposit:           big.NewInt(100),
		MaxSubmissionCost: big.NewInt(30),
		CreditBack:        common.RandAddress(),
		Beneficiary:       common.NewAddressFromEth(beneficiaryAuth.From),
		MaxGas:            big.NewInt(0),
		GasPriceBid:       big.NewInt(0),
		Data:              simpleABI.Methods["exists"].ID,
	}

	requestId, err := backend.AddInboxMessage(retryableTx, sender)
	test.FailIfError(t, err)

	ticketId := hashing.SoliditySHA3(hashing.Bytes32(requestId), hashing.Uint256(big.NewInt(0)))

	retryable, err := arboscontracts.NewArbRetryableTx(arbos.ARB_RETRYABLE_ADDRESS, client)
	test.FailIfError(t, err)

	tx, err := retryable.Redeem(otherAuth, ticketId)
	test.FailIfError(t, err)

	res, err := backend.db.GetRequest(ticketId)
	test.FailIfError(t, err)

	res2, err := backend.db.GetRequest(common.NewHashFromEth(tx.Hash()))
	test.FailIfError(t, err)

	if len(res.ReturnData) != 32 {
		t.Fatal("expected 32 byte of return data")
	}
	if new(big.Int).SetBytes(res.ReturnData).Cmp(big.NewInt(10)) != 0 {
		t.Error("wrong return value")
	}

	if len(res2.ReturnData) != 0 {
		t.Fatal("expected no return data")
	}
}
