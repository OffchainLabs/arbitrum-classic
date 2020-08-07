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
	"bufio"
	"context"
	"crypto/ecdsa"
	"errors"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type fundRequest struct {
	account common.Address
	amount  *big.Int
}

func fundAccounts(
	ctx context.Context,
	client ethutils.EthClient,
	inboxAddress common.Address,
	rollupAddress common.Address,
	fundingPKs []*ecdsa.PrivateKey,
	requests []fundRequest,
) error {
	auths := make([]*bind.TransactOpts, 0, len(fundingPKs))
	for _, pk := range fundingPKs {
		auth := bind.NewKeyedTransactor(pk)
		auth.GasPrice = big.NewInt(0)
		auths = append(auths, auth)
	}
	remainingRequests := requests
	for len(remainingRequests) > 0 {
		if len(auths) == 0 {
			return errors.New("not enough value to fund account")
		}
		req := remainingRequests[0]
		auth := auths[0]
		currentBalance, err := client.BalanceAt(ctx, auth.From, nil)
		if err != nil {
			return nil
		}
		currentBalance = currentBalance.Sub(currentBalance, big.NewInt(100000))

		globalInbox, err := ethbridge.NewEthAuthClient(client, auth).NewGlobalInbox(inboxAddress, rollupAddress)
		if err != nil {
			return err
		}

		amountToDeposit := req.amount
		if amountToDeposit.Cmp(currentBalance) > 0 {
			amountToDeposit = currentBalance
			req.amount = req.amount.Sub(req.amount, amountToDeposit)
			auths = auths[1:]
		} else {
			remainingRequests = remainingRequests[1:]
		}

		log.Println("Sending", amountToDeposit, "to", req.account)

		if err := globalInbox.DepositEthMessage(
			context.Background(),
			req.account,
			amountToDeposit,
		); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	requests := make([]fundRequest, 0)
	for _, arg := range os.Args[1:] {
		split := strings.Split(arg, ":")
		if len(split) != 2 {
			log.Fatal("Expected funding accounts format [account:amount]...")
		}
		hexAmount, ok := new(big.Int).SetString(split[1], 16)
		if !ok {
			log.Fatal("amount must be in hex")
		}
		requests = append(requests, fundRequest{
			account: common.HexToAddress(split[0]),
			amount:  hexAmount,
		})
	}

	dbPath := "test-aggregator-db"
	if err := os.RemoveAll(dbPath); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := os.RemoveAll(dbPath); err != nil {
			log.Fatal(err)
		}
	}()

	client, pks := test.SimulatedBackend()
	auth := bind.NewKeyedTransactor(pks[0])
	go func() {
		t := time.NewTicker(time.Millisecond * 1000)
		for range t.C {
			client.Commit()
		}
	}()

	config := valprotocol.ChainParams{
		StakeRequirement:        big.NewInt(10),
		GracePeriod:             common.TimeTicks{Val: big.NewInt(13000 * 2)},
		MaxExecutionSteps:       10000000000,
		ArbGasSpeedLimitPerTick: 200000,
	}

	factoryAddr, err := ethbridge.DeployRollupFactory(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	arbAuthClient := ethbridge.NewEthAuthClient(client, auth)

	factory, err := arbAuthClient.NewArbFactory(common.NewAddressFromEth(factoryAddr))
	if err != nil {
		log.Fatal(err)
	}

	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		log.Fatal(err)
	}

	rollupAddress, _, err := factory.CreateRollup(
		context.Background(),
		mach.Hash(),
		config,
		common.Address{},
	)
	if err != nil {
		log.Fatal(err)
	}

	inboxAddress, err := factory.GlobalInboxAddress()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created rollup chain at", rollupAddress.Hex(), "with machine hash", mach.Hash())

	if err := fundAccounts(context.Background(), client, inboxAddress, rollupAddress, pks[1:], requests); err != nil {
		log.Fatal(err)
	}

	log.Println("Funded requested accounts")

	serverLogger := &logger{}

	go func() {
		if err := rpc.LaunchAggregator(
			context.Background(),
			client,
			auth,
			rollupAddress,
			arbos.Path(),
			dbPath,
			"1235",
			"8547",
			utils.RPCFlags{},
			time.Millisecond*500,
			serverLogger,
		); err != nil {
			log.Fatal(err)
		}
	}()

	_, err = bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}

	globalInbox, err := ethbridge.NewEthClient(client).NewGlobalInboxWatcher(inboxAddress, rollupAddress)
	if err != nil {
		log.Fatal(err)
	}

	events, err := globalInbox.GetDeliveredEvents(context.Background(), nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Processed", len(events), "messages")

	msgs := make([]inbox.InboxMessage, 0, len(events))

	//calls := serverLogger.calls
	//prevTimestamp := big.NewInt(0)
	//addCall := func() {
	//	nextCall := calls[0]
	//	callMsg := message.NewInboxMessage(
	//		message.NewL2Message(nextCall.msg),
	//		common.NewAddressFromEth(nextCall.sender),
	//		big.NewInt(int64(len(msgs))),
	//		inbox.ChainTime{
	//			BlockNum:  nextCall.blockId.Height,
	//			Timestamp: prevTimestamp,
	//		},
	//	)
	//	msgs = append(msgs, callMsg)
	//	calls = calls[1:]
	//}

	for _, ev := range events {
		//for len(calls) > 0 && calls[0].blockId.Height.Cmp(ev.BlockId.Height) < 0 {
		//	addCall()
		//}
		//ev.Message.InboxSeqNum = big.NewInt(int64(len(msgs)))
		msgs = append(msgs, ev.Message)

		log.Println("Incoming message", ev.Message)
		//prevTimestamp = ev.Message.ChainTime.Timestamp
	}

	//for len(calls) > 0 {
	//	addCall()
	//}

	//for _, call := range serverLogger.calls {
	//	log.Println("got call", call.msg, "from", call.sender.Hex())
	//}

	//lastCall := serverLogger.calls[len(serverLogger.calls)-1]
	//callMsg := message.NewInboxMessage(
	//	message.NewL2Message(lastCall.msg),
	//	common.NewAddressFromEth(lastCall.sender),
	//	big.NewInt(int64(len(msgs))),
	//	inbox.ChainTime{
	//		BlockNum:  lastCall.blockId.Height,
	//		Timestamp: events[len(events)-1].Message.ChainTime.Timestamp,
	//	},
	//)
	//msgs = append(msgs, callMsg)

	//log.Println("Appended call as inbox message", callMsg)

	assertion, _ := mach.ExecuteAssertion(10000000000000, msgs, 0)
	testVec, err := inbox.TestVectorJSON(msgs, assertion.ParseLogs(), assertion.ParseOutMessages())
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("testVec.json", testVec, 0644); err != nil {
		log.Fatal(err)
	}

	log.Println("Saving test vector")
	log.Println(len(msgs), "inbox messages")
	log.Println(assertion.LogsCount, "logs")
	log.Println(assertion.OutMsgsCount, "sends")

	for _, avmLog := range assertion.ParseLogs() {
		res, err := evm.NewResultFromValue(avmLog)
		if err != nil {
			log.Fatal(err)
		}
		switch res := res.(type) {
		case *evm.TxResult:
			log.Println("Got tx result", res)
		case *evm.BlockInfo:
			log.Println("Got block info", res)
		}
	}
}

type loggedCall struct {
	msg     message.ContractTransaction
	sender  ethcommon.Address
	blockId *common.BlockId
}

type logger struct {
	calls []loggedCall
}

func (l *logger) LogCall(msg message.ContractTransaction, sender ethcommon.Address, blockId *common.BlockId) {
	l.calls = append(l.calls, loggedCall{
		msg:     msg,
		sender:  sender,
		blockId: blockId,
	})
}
