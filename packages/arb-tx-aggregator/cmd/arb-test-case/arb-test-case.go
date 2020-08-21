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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
)

func main() {
	//if err := getMessages(
	//	"https://ropsten.infura.io/v3/5851cb75448e4f8da37c5015006589a6",
	//	common.HexToAddress("0xEb57E0FB729E892E84259B4e6Dc08442Aa6E9Ee4"),
	//	filename,
	//); err != nil {
	//	panic(err)
	//}

	if err := generateTestCase(
		"http://localhost:7545",
		common.HexToAddress("0xc68DCee7b8cA57F41D1A417103CB65836E99e013"),
		arbos.Path(),
	); err != nil {
		panic(err)
	}
}

func generateTestCase(ethURL string, rollupAddress common.Address, contract string) error {
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

	_, eventId, _, _, err := rollupWatcher.GetCreationInfo(ctx)
	if err != nil {
		return err
	}

	events, err := inboxWatcher.GetDeliveredEvents(ctx, eventId.BlockId.Height.AsInt(), nil)
	if err != nil {
		return err
	}

	messages := make([]inbox.InboxMessage, 0, len(events))
	for i, ev := range events {
		msg, err := message.NestedMessage(ev.Message)
		if err != nil {
			return err
		}
		log.Println("Message", msg, "from", ev.Message.Sender.Hex())
		messages = append(messages, ev.Message)
		if i == len(events)-1 {
			l2 := msg.(message.L2Message)
			ab, err := l2.AbstractMessage()
			if err != nil {
				return err
			}
			tx := ab.(message.TransactionBatch).Transactions[0]
			inner, err := message.L2Message{Data: tx}.AbstractMessage()
			if err != nil {
				return err
			}
			signer := types.NewEIP155Signer(message.ChainAddressToID(rollupAddress))
			ethTx := inner.(message.SignedTransaction).Tx
			sender, err := types.Sender(signer, ethTx)
			if err != nil {
				return err
			}
			log.Println("ethTx", ethTx)
			log.Println("sender", sender.Hex())
		}
	}

	mach, err := cmachine.New(contract)
	if err != nil {
		return err
	}

	assertion, _ := mach.ExecuteAssertion(
		1000000000000,
		messages,
		0,
	)

	for _, lg := range assertion.ParseLogs() {
		res, err := evm.NewTxResultFromValue(lg)
		if err != nil {
			return err
		}
		log.Println("result", res)
		log.Println("sender", res.IncomingRequest.Sender.Hex())

	}

	//data, err := inbox.TestVectorJSON(messages, assertion.ParseLogs(), assertion.ParseOutMessages())
	//if err != nil {
	//	return err
	//}
	//log.Println(string(data))
	return nil
}
