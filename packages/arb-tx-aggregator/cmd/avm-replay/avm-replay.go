/*
 * Copyright 2020, Offchain Labs, Inc.
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

package main

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/l2message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
)

func main() {
	filename := "messages.dat"
	//if err := getMessages(
	//	"https://ropsten.infura.io/v3/5851cb75448e4f8da37c5015006589a6",
	//	common.HexToAddress("0xEb57E0FB729E892E84259B4e6Dc08442Aa6E9Ee4"),
	//	filename,
	//); err != nil {
	//	panic(err)
	//}

	//if err := getMessages(
	//	"http://localhost:7545",
	//	common.HexToAddress("0xc68DCee7b8cA57F41D1A417103CB65836E99e013"),
	//	filename,
	//); err != nil {
	//	panic(err)
	//}

	if err := testMessages(filename, arbos.Path()); err != nil {
		panic(err)
	}
}

func toEth(val *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(val.String())
	return new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
}

func runMessage(mach machine.Machine, msg message.InboxMessage) (*evm.Result, error) {
	inbox := value.NewEmptyTuple()
	inbox = value.NewTuple2(inbox, msg.AsValue())
	assertion, _ := mach.ExecuteAssertion(
		100000,
		inbox,
		1000,
	)
	//log.Println("ran assertion")
	logs := assertion.ParseLogs()
	if len(logs) != 1 {
		log.Fatal("returned incorrect log count")
	}
	evmResult, err := evm.NewResultFromValue(logs[0])
	if err != nil {
		return nil, err
	}
	return evmResult, nil
}

var chain = common.HexToAddress("0xc68DCee7b8cA57F41D1A417103CB65836E99e013")

func printL2Message(tx []byte) error {
	msg, err := l2message.NewL2MessageFromData(tx)
	if err != nil {
		return err
	}
	switch msg := msg.(type) {
	case l2message.TransactionBatch:
		for _, tx := range msg.Transactions {
			if err := printL2Message(tx); err != nil {
				return err
			}
		}
	case l2message.SignedTransaction:
		ethTx, err := msg.AsEthTx(chain)
		if err != nil {
			return err
		}

		//sender, err := types.NewEIP155Signer(l2message.ChainAddressToID(chain)).Sender(ethTx)
		//if err != nil {
		//	return err
		//}

		log.Println("SignedTransaction", ethTx.Hash().Hex()) // , "from", sender.Hex()
		log.Println("tx:", l2message.NewSignedTransactionFromEth(ethTx))
		//log.Println(msg)
	default:
		log.Printf("Input: %T\n", msg)
	}
	return nil
}

func testMessages(filename string, contract string) error {
	messages, err := loadMessages(filename)
	if err != nil {
		return err
	}

	mach, err := cmachine.New(contract)
	if err != nil {
		return err
	}

	//for _, msg := range messages {
	//	inbox := value.NewEmptyTuple()
	//	inbox = value.NewTuple2(inbox, msg.AsValue())
	//	assertion, _ := mach.ExecuteAssertion(100000000000, inbox, 0)
	//	log.Println("Ran assertion", assertion.NumGas)
	//}

	inbox := value.NewEmptyTuple()
	for _, msg := range messages {
		nested, err := msg.NestedMessage()
		if err != nil {
			return err
		}
		if tx, ok := nested.(message.L2Message); ok {
			if err := printL2Message(tx.Data); err != nil {
				return err
			}
		} else {
			log.Printf("Input %T: %v from %v\n", nested, nested, msg.Sender)
		}
		inbox = value.NewTuple2(inbox, msg.AsValue())
	}
	assertion, steps := mach.ExecuteAssertion(100000000000, inbox, 0)
	log.Println("Ran for", steps, assertion.NumGas)
	//testData, err := value.TestVectorJSON(inbox, assertion.ParseLogs(), assertion.ParseOutMessages())
	//if err != nil {
	//	return err
	//}
	//log.Println(string(testData))
	logs := assertion.ParseLogs()
	log.Println("Had logs", len(logs))
	for _, avmLog := range logs {
		res, err := evm.NewResultFromValue(avmLog)
		if err != nil {
			return err
		}
		log.Println("Got res", res.ResultCode, res.GasUsed, res.L1Message.Sender, res.L1Message.MessageID())
		log.Println("Res had logs", res.EVMLogs)
	}
	return nil
}

func loadMessages(filename string) ([]message.InboxMessage, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	messagesStackVal, err := value.UnmarshalValue(f)
	if err != nil {
		return nil, err
	}

	messageVals, err := message.StackValueToList(messagesStackVal)
	if err != nil {
		return nil, err
	}

	messages := make([]message.InboxMessage, 0, len(messageVals))
	for _, val := range messageVals {
		msg, err := message.NewInboxMessageFromValue(val)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	log.Println("Got", len(messages), "messages")
	return messages, nil
}

func getMessages(ethURL string, rollupAddress common.Address, filename string) error {
	ctx := context.Background()

	ethclint, err := ethclient.Dial(ethURL)
	if err != nil {
		return err
	}

	client := ethbridge.NewEthClient(ethclint)
	rollupWatcher, err := client.NewRollupWatcher(rollupAddress)
	if err != nil {
		return err
	}

	inboxAddress, err := rollupWatcher.InboxAddress(ctx)
	if err != nil {
		return err
	}

	inboxWatcher, err := client.NewGlobalInboxWatcher(inboxAddress, rollupAddress)
	if err != nil {
		return err
	}

	_, blockId, _, err := rollupWatcher.GetCreationInfo(ctx)
	if err != nil {
		return err
	}

	events, err := inboxWatcher.GetDeliveredEvents(ctx, blockId.Height.AsInt(), nil)
	if err != nil {
		return err
	}

	log.Println("Got", len(events), "messages")

	values := make([]value.Value, 0, len(events))
	for _, ev := range events {
		values = append(values, ev.Message.AsValue())
	}

	messagesStackVal := message.ListToStackValue(values)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if err := value.MarshalValue(messagesStackVal, f); err != nil {
		return err
	}
	return nil
}
