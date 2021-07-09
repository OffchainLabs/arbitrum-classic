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

package ethbridge

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func TestRetryable(t *testing.T) {
	clnt, pks := test.SimulatedBackend(t)
	auth, err := bind.NewKeyedTransactorWithChainID(pks[0], big.NewInt(1337))
	test.FailIfError(t, err)
	bridgeAddress, _, bridge, err := ethbridgecontracts.DeployBridge(auth, clnt)
	test.FailIfError(t, err)
	inboxAddress, _, inbox, err := ethbridgecontracts.DeployInbox(auth, clnt)
	test.FailIfError(t, err)
	clnt.Commit()

	_, err = bridge.Initialize(auth)
	test.FailIfError(t, err)
	_, err = inbox.Initialize(auth, bridgeAddress, ethcommon.Address{})
	test.FailIfError(t, err)
	_, err = bridge.SetInbox(auth, inboxAddress, true)
	test.FailIfError(t, err)
	arbTx := message.RetryableTx{
		Destination:       common.RandAddress(),
		Value:             common.RandBigInt(),
		Deposit:           big.NewInt(1000),
		MaxSubmissionCost: common.RandBigInt(),
		CreditBack:        common.RandAddress(),
		Beneficiary:       common.RandAddress(),
		MaxGas:            common.RandBigInt(),
		GasPriceBid:       common.RandBigInt(),
		Data:              common.RandBytes(100),
	}
	tx, err := inbox.CreateRetryableTicket(
		&bind.TransactOpts{
			From:   auth.From,
			Signer: auth.Signer,
			Value:  arbTx.Deposit,
		},
		arbTx.Destination.ToEthAddress(),
		arbTx.Value,
		arbTx.MaxSubmissionCost,
		arbTx.CreditBack.ToEthAddress(),
		arbTx.Beneficiary.ToEthAddress(),
		arbTx.MaxGas,
		arbTx.GasPriceBid,
		arbTx.Data,
	)
	test.FailIfError(t, err)

	clnt.Commit()

	receipt, err := clnt.TransactionReceipt(context.Background(), tx.Hash())
	test.FailIfError(t, err)

	if len(receipt.Logs) != 2 {
		t.Fatal("wrong receipt count")
	}
	if receipt.Logs[1].Topics[0] != inboxMessageDeliveredID {
		t.Fatal("wrong topic")
	}
	ev, err := inbox.ParseInboxMessageDelivered(*receipt.Logs[1])
	test.FailIfError(t, err)

	parsedArbTx := message.NewRetryableTxFromData(ev.Data)
	if !parsedArbTx.Equals(arbTx) {
		t.Log(parsedArbTx)
		t.Log(arbTx)
		t.Error("event data not equal")
	}
}
