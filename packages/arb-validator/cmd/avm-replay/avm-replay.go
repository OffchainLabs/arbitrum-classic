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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
	"math/big"
	"os"
)

func main() {
	filename := "messages.dat"
	//if err := getMessages(
	//	"https://ropsten.infura.io/v3/5851cb75448e4f8da37c5015006589a6",
	//	common.HexToAddress("0x87a7bAd640eaa9089be09d827952bf9E45Ce8780"),
	//	filename,
	//); err != nil {
	//	panic(err)
	//}

	if err := testMessages(filename, "contract.ao"); err != nil {
		panic(err)
	}
}

func testMessages(filename string, contract string) error {
	messages, err := loadMessages(filename)
	if err != nil {
		return err
	}
	//
	//mach, err := loader.LoadMachineFromFile(contract, false, "cpp")
	//if err != nil {
	//	return err
	//}
	//
	//tb := protocol.NewRandomTimeBounds()

	singleMessages := make([]message.Delivered, 0)
	for _, msg := range messages {
		for _, del := range msg.VMInboxMessages() {
			singleMessages = append(singleMessages, message.Delivered{
				Message:      del.Message,
				DeliveryInfo: del.DeliveryInfo,
			})
		}

	}

	//runMsg := func(msg message.Delivered) error {
	//	vmInbox := structures.NewVMInbox()
	//	vmInbox.DeliverMessage(msg)
	//
	//	_, _ = mach.ExecuteAssertion(
	//		10000000,
	//		tb,
	//		vmInbox.AsValue(),
	//		1000,
	//	)
	//	blocked := mach.IsBlocked(common.NewTimeBlocksInt(0), true)
	//	if blocked != nil {
	//		return fmt.Errorf("machine is blocked: %v", blocked)
	//	}
	//	return nil
	//}

	//for _, msg := range singleMessages[:len(singleMessages)-1] {
	//	if err := runMsg(msg); err != nil {
	//		log.Println(err)
	//		return nil
	//	}
	//}

	lastMessage := singleMessages[len(singleMessages)-1]

	info := lastMessage.DeliveryInfo
	msg := lastMessage.Message.(message.Transaction)
	msg.SequenceNum = big.NewInt(0)

	vmInbox := structures.NewVMInbox()
	vmInbox.DeliverMessage(message.Delivered{
		Message:      msg,
		DeliveryInfo: info,
	})
	val := vmInbox.AsValue()
	log.Println(hexutil.Encode(value.MarshalValueToBytes(val)))

	//log.Println("Delivering crash message", msg)
	//i := 0
	//for {
	//	mach.PrintState()
	//	blocked := mach.IsBlocked(common.NewTimeBlocksInt(0), true)
	//	if blocked != nil {
	//		log.Printf("machine after %v steps is blocked: %v\n", i, blocked)
	//		return nil
	//	}
	//	_, _ = mach.ExecuteAssertion(
	//		1,
	//		tb,
	//		vmInbox.AsValue(),
	//		1000,
	//	)
	//	i++
	//}
	return nil
}

func loadMessages(filename string) ([]message.Delivered, error) {
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

	received := make([]message.Received, 0, len(messageVals))
	for _, val := range messageVals {
		msg, err := message.UnmarshalReceivedFromCheckpoint(val)
		if err != nil {
			return nil, err
		}
		received = append(received, msg)
	}

	log.Println("Got", len(received), "messages")

	inbox := structures.NewInbox()
	for _, msg := range received {
		inbox.DeliverMessage(msg)
	}

	return inbox.GetAllMessages(), nil
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

	received, err := inboxWatcher.GetAllReceived(ctx, blockId.Height.AsInt(), nil)
	if err != nil {
		return err
	}

	log.Println("Got", len(received), "messages")

	values := make([]value.Value, 0, len(received))
	for _, msg := range received {
		values = append(values, msg.CheckpointValue())
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
