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

package ethbridgetest

import (
	"context"
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetest/protocoltester"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetest/valuetester"
	"log"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetest/messagetester"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
)

var privHex = "27e926925fb5903ee038c894d9880f74d3dd6518e23ab5e5651de93327c7dffa"

var auth *bind.TransactOpts
var tester *messagetester.MessageTester
var valueTester *valuetester.ValueTester
var protocolTester *protocoltester.ProtocolTester

//var arbRollupTester *rollup.ArbRollup
var client *ethclient.Client

var addr1 common.Address
var addr2 common.Address
var addr3 common.Address

var errHash = errors.New("ethbridge calculated wrong hash")
var errMsgHash = errors.New("ethbridge calculated wrong message hash")

func TestMain(m *testing.M) {
	addr1[0] = 76
	addr1[19] = 93
	addr2[0] = 43
	addr2[19] = 12
	addr3[0] = 73
	addr3[19] = 85

	var err error

	auth, err = test.SetupAuth(privHex)
	if err != nil {
		log.Fatal(err)
	}
	client, err = ethclient.Dial(test.GetEthUrl())
	if err != nil {
		log.Fatal(err)
	}
	_, tx, deployedMessageTester, err := messagetester.DeployMessageTester(
		auth,
		client,
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = ethbridge.WaitForReceiptWithResults(
		context.Background(),
		client,
		auth.From,
		tx,
		"DeployMessageTester",
	)
	if err != nil {
		log.Fatal(err)
	}

	_, tx, deployedValueTester, err := valuetester.DeployValueTester(
		auth,
		client,
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = ethbridge.WaitForReceiptWithResults(
		context.Background(),
		client,
		auth.From,
		tx,
		"DeployMessageTester",
	)
	if err != nil {
		log.Fatal(err)
	}

	_, valTx, deployedValueTester, err := valuetester.DeployValueTester(
		auth,
		client,
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = ethbridge.WaitForReceiptWithResults(
		context.Background(),
		client,
		auth.From,
		valTx,
		"DeployValueTester",
	)
	if err != nil {
		log.Fatal(err)
	}

	_, protocolTx, deployedProtocolTester, err := protocoltester.DeployProtocolTester(
		auth,
		client,
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = ethbridge.WaitForReceiptWithResults(
		context.Background(),
		client,
		auth.From,
		protocolTx,
		"DeployValueTester",
	)
	if err != nil {
		log.Fatal(err)
	}

	tester = deployedMessageTester
	valueTester = deployedValueTester
	protocolTester = deployedProtocolTester

	code := m.Run()
	os.Exit(code)
}
