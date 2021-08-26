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
	"bytes"
	"context"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func TestL2ToL1Tx(t *testing.T) {
	config := protocol.ChainParams{
		GracePeriod:               common.NewTimeBlocksInt(3),
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}

	backend, db, srv, cancelDevNode := NewTestDevNode(t, *arbosfile, config, common.RandAddress(), nil, false)
	defer cancelDevNode()

	client := web3.NewEthClient(srv, true)
	arbSys, err := arboscontracts.NewArbSys(arbos.ARB_SYS_ADDRESS, client)
	if err != nil {
		t.Fatal(err)
	}
	privkey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privkey)

	clnt, auths := test.SimulatedBackend(t)
	ethAuth := auths[0]
	deposit := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(1000000),
				GasPriceBid: big.NewInt(0),
				DestAddress: common.NewAddressFromEth(auth.From),
				Payment:     big.NewInt(100),
				Data:        nil,
			},
		}),
	}
	if _, err := backend.AddInboxMessage(deposit, common.RandAddress()); err != nil {
		t.Fatal(err)
	}

	latest, err := backend.db.LatestBlock()
	test.FailIfError(t, err)
	blockInfo, _, err := backend.db.GetBlockResults(latest)
	test.FailIfError(t, err)
	if blockInfo == nil {
		t.Fatal("should get block info")
	}

	rand.Seed(534523435)

	withdrawAmount := big.NewInt(1)

	l2SendLogs := make([]*arboscontracts.ArbSysL2ToL1Transaction, 0)
	l1Dests := make([]common.Address, 0)
	for i := 0; i < 12; i++ {
		dest := common.RandAddress()
		l1Dests = append(l1Dests, dest)
		t.Log("Send tx to L1", dest.Hex())
		tx, err := arbSys.SendTxToL1(&bind.TransactOpts{
			From:     auth.From,
			Nonce:    auth.Nonce,
			Signer:   auth.Signer,
			Value:    withdrawAmount,
			GasPrice: nil,
			GasLimit: 0,
			Context:  nil,
		}, dest.ToEthAddress(), nil)
		if err != nil {
			t.Fatal(err)
		}
		arbRes, err := backend.db.GetRequest(common.NewHashFromEth(tx.Hash()))
		test.FailIfError(t, err)
		if len(arbRes.ReturnData) != 32 {
			t.Fatal("expected return data")
		}
		l2SendNum := new(big.Int).SetBytes(arbRes.ReturnData)
		if l2SendNum.Cmp(big.NewInt(int64(i))) != 0 {
			t.Fatal("unexpected l2 send num", l2SendNum, "instead of", i)
		}
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		test.FailIfError(t, err)
		if len(receipt.Logs) != 1 {
			t.Fatal("unexpected log count")
		}
		sendLog := receipt.Logs[0]
		if sendLog.Topics[0] != arbos.L2ToL1TransactionID {
			t.Fatal("unexpected topic", sendLog.Topics[0], arbos.L2ToL1TransactionID)
		}
		parsedEv, err := arbSys.ParseL2ToL1Transaction(*sendLog)
		if err != nil {
			t.Fatal(err)
		}
		if parsedEv.UniqueId.Cmp(l2SendNum) != 0 {
			t.Fatal("incorrect unique id")
		}
		l2SendLogs = append(l2SendLogs, parsedEv)
		if i%8 == 0 {
			// ArbOS spaces out sends every 1800 seconds by default, so advance one send
			backend.l1Emulator.IncreaseTime(1800)
		}
	}

	batches := make([]*evm.MerkleRootResult, 0)
	for i := 0; i < 3; i++ {
		batch, err := db.GetMessageBatch(big.NewInt(int64(i)))
		if err != nil {
			t.Fatal(err)
		}
		if batch == nil {
			t.Fatal("message batch not found")
		}
		if batch.BatchNumber.Cmp(big.NewInt(int64(i))) != 0 {
			t.Fatal("wrong batch num")
		}
		batches = append(batches, batch)
	}

	bridgeAddress, _, bridge, err := ethbridgecontracts.DeployBridge(ethAuth, clnt)
	test.FailIfError(t, err)
	outboxAddress, _, outbox, err := ethbridgecontracts.DeployOutbox(ethAuth, clnt)
	test.FailIfError(t, err)
	inboxAddress, _, inbox, err := ethbridgecontracts.DeployInbox(ethAuth, clnt)
	test.FailIfError(t, err)
	clnt.Commit()

	_, err = bridge.Initialize(ethAuth)
	test.FailIfError(t, err)
	_, err = outbox.Initialize(ethAuth, ethAuth.From, bridgeAddress)
	test.FailIfError(t, err)
	_, err = inbox.Initialize(ethAuth, bridgeAddress, ethcommon.Address{})
	test.FailIfError(t, err)
	clnt.Commit()

	_, err = bridge.SetOutbox(ethAuth, outboxAddress, true)
	test.FailIfError(t, err)
	_, err = bridge.SetInbox(ethAuth, inboxAddress, true)
	test.FailIfError(t, err)
	clnt.Commit()

	bridgeDeposit := big.NewInt(100000000)
	_, err = inbox.DepositEth(&bind.TransactOpts{
		From:     ethAuth.From,
		Nonce:    ethAuth.Nonce,
		Signer:   ethAuth.Signer,
		Value:    bridgeDeposit,
		GasPrice: nil,
		GasLimit: 0,
		Context:  nil,
	}, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	clnt.Commit()

	beforeBridgeBalance, err := clnt.BalanceAt(context.Background(), bridgeAddress, nil)
	if err != nil {
		t.Fatal(err)
	}
	if beforeBridgeBalance.Cmp(bridgeDeposit) != 0 {
		t.Fatal("bridge didn't receive balance")
	}
	sendCount, err := backend.arbcore.GetSendCount()
	if err != nil {
		t.Fatal(err)
	}
	sends, err := backend.arbcore.GetSends(big.NewInt(0), sendCount)
	if err != nil {
		t.Fatal(err)
	}

	var sendsData []byte
	var sendLengths []*big.Int
	for _, send := range sends {
		sendsData = append(sendsData, send...)
		sendLengths = append(sendLengths, big.NewInt(int64(len(send))))
	}
	_, err = outbox.ProcessOutgoingMessages(ethAuth, sendsData, sendLengths)
	if err != nil {
		t.Fatal(err)
	}
	clnt.Commit()

	for i, batch := range batches {
		outboxEntryAddress, err := outbox.Outboxes(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			t.Fatal(err)
		}
		outboxEntry, err := ethbridgecontracts.NewOutboxEntry(outboxEntryAddress, clnt)
		if err != nil {
			t.Fatal(err)
		}
		root, err := outboxEntry.Root(&bind.CallOpts{})
		if err != nil {
			t.Fatal(err)
		}
		if root != batch.Tree.Hash() {
			t.Fatal("wrong root")
		}
		numRemaining, err := outboxEntry.NumRemaining(&bind.CallOpts{})
		if err != nil {
			t.Fatal(err)
		}
		if numRemaining.Cmp(batch.NumInBatch) != 0 {
			t.Fatal("wrong num remaining")
		}
	}

	nodeInterface, err := arboscontracts.NewNodeInterface(arbos.ARB_NODE_INTERFACE_ADDRESS, client)
	if err != nil {
		t.Fatal(err)
	}

	totalEntries := 0
	for i, batch := range batches {
		entries := batch.Tree.Entries()
		for j, entry := range entries {
			send, err := evm.NewVirtualSendResultFromData(entry)
			if err != nil {
				t.Fatal(err)
			}
			res, ok := send.(*evm.L2ToL1TxResult)
			if !ok {
				t.Fatal("expected l2 to l1 result")
			}

			// Verify that the L2 log emitted with this event was correct
			l2SendLog := l2SendLogs[totalEntries]
			if l2SendLog.BatchNumber.Cmp(big.NewInt(int64(i))) != 0 {
				t.Fatal("wrong batch num")
			}
			if l2SendLog.IndexInBatch.Cmp(big.NewInt(int64(j))) != 0 {
				t.Fatal("wrong item num")
			}
			if l2SendLog.Caller != res.L2Sender.ToEthAddress() {
				t.Fatal("wrong l2 sender")
			}
			if l2SendLog.Destination != res.L1Dest.ToEthAddress() {
				t.Fatal("wrong l1 dest")
			}
			if l2SendLog.ArbBlockNum.Cmp(res.L2Block) != 0 {
				t.Fatal("wrong l2 block")
			}
			if l2SendLog.EthBlockNum.Cmp(res.L1Block) != 0 {
				t.Fatal("wrong l1 block")
			}
			if l2SendLog.Timestamp.Cmp(res.Timestamp) != 0 {
				t.Fatal("wrong timestamp")
			}
			if l2SendLog.Callvalue.Cmp(res.Value) != 0 {
				t.Fatal("wrong amount")
			}
			if !bytes.Equal(l2SendLog.Data, res.Calldata) {
				t.Fatal("wrong calldata")
			}

			batchNum := big.NewInt(int64(i))
			index := uint64(j)
			msgData, err := nodeInterface.LookupMessageBatchProof(&bind.CallOpts{}, batchNum, index)
			if err != nil {
				t.Fatal(err)
			}

			if msgData.L2Sender != res.L2Sender.ToEthAddress() {
				t.Fatal("wrong l2 sender")
			}
			if msgData.L1Dest != res.L1Dest.ToEthAddress() {
				t.Fatal("wrong l1 dest")
			}
			if msgData.L2Block.Cmp(res.L2Block) != 0 {
				t.Fatal("wrong l2 block")
			}
			if msgData.L1Block.Cmp(res.L1Block) != 0 {
				t.Fatal("wrong l1 block")
			}
			if msgData.Timestamp.Cmp(res.Timestamp) != 0 {
				t.Fatal("wrong timestamp")
			}
			if msgData.Amount.Cmp(res.Value) != 0 {
				t.Fatal("wrong amount")
			}
			if !bytes.Equal(msgData.CalldataForL1, res.Calldata) {
				t.Fatal("wrong calldata")
			}
			t.Log("Execute", msgData.L1Dest.Hex(), msgData.Amount, hexutil.Encode(msgData.CalldataForL1))
			tx, err := outbox.ExecuteTransaction(
				ethAuth,
				batchNum,
				msgData.Proof,
				msgData.Path,
				msgData.L2Sender,
				msgData.L1Dest,
				msgData.L2Block,
				msgData.L1Block,
				msgData.Timestamp,
				msgData.Amount,
				msgData.CalldataForL1,
			)
			if err != nil {
				t.Fatal(err)
			}
			clnt.Commit()
			receipt, err := clnt.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				t.Fatal(err)
			}
			if receipt.Status == types.ReceiptStatusFailed {
				t.Fatal("transaction failed")
			}
			totalEntries++
		}
	}

	for _, dest := range l1Dests[:totalEntries] {
		code, err := clnt.CodeAt(context.Background(), dest.ToEthAddress(), nil)
		if err != nil {
			t.Fatal(err)
		}
		if len(code) != 0 {
			t.Fatal("should have no code")
		}

		balance, err := clnt.BalanceAt(context.Background(), dest.ToEthAddress(), nil)
		if err != nil {
			t.Fatal(err)
		}
		if balance.Cmp(withdrawAmount) != 0 {
			t.Fatal("wrong balance after", balance)
		}
	}
}
