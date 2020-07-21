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

package challenges

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"testing"
	"time"
)

var testerAddress ethcommon.Address

func TestChallenges(t *testing.T) {
	client, pks := test.SimulatedBackend()

	auths := make([]*bind.TransactOpts, 0)
	for _, pk := range pks {
		auths = append(auths, bind.NewKeyedTransactor(pk))
	}

	factorAddr, err := ethbridge.DeployChallengeFactory(auths[0], client)
	if err != nil {
		t.Fatal(err)
	}

	testerAddress, _, _, err = ethbridgetestcontracts.DeployChallengeTester(auths[0], client, factorAddr)
	if err != nil {
		t.Fatal(err)
	}
	client.Commit()

	go func() {
		t := time.NewTicker(time.Second * 1)
		for range t.C {
			client.Commit()
		}
	}()

	t.Run("Inbox Top Challenge", func(t *testing.T) {
		testInboxTopChallenge(t, client, auths[0], auths[1])
	})
	t.Run("Execution Challenge", func(t *testing.T) {
		testExecutionChallenge(t, client, auths[4], auths[5])
	})
	t.Run("Messages Challenge", func(t *testing.T) {
		testMessagesChallenge(t, client, auths[2], auths[3])
	})
	t.Run("ERC20 Messages Challenge", func(t *testing.T) {
		testMessagesChallengeERC20(t, client, auths[6], auths[7])
	})
	t.Run("ERC721 Messages Challenge", func(t *testing.T) {
		testMessagesChallengeERC721(t, client, auths[8], auths[9])
	})
	t.Run("Transaction Messages Challenge", func(t *testing.T) {
		testMessagesChallengeTrnx(t, client, auths[10], auths[11])
	})
	t.Run("Contract Trans Messages Challenge", func(t *testing.T) {
		testMessagesChallengeContractTrnx(t, client, auths[12], auths[13])
	})
}
