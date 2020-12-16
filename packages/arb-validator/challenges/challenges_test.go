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
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"testing"
	"time"
)

var testerAddress ethcommon.Address

func TestChallenges(t *testing.T) {
	clnt, pks := test.SimulatedBackend()
	client := &ethutils.SimulatedEthClient{SimulatedBackend: clnt}

	ctx := context.Background()
	txOpts := make([]*bind.TransactOpts, 0, len(pks))
	authClients := make([]*ethbridge.EthArbAuthClient, 0, len(pks))
	for _, pk := range pks {
		txOpt := bind.NewKeyedTransactor(pk)

		authClient, err := ethbridge.NewEthAuthClient(ctx, client, txOpt)
		if err != nil {
			t.Fatal(err)
		}

		txOpts = append(txOpts, txOpt)
		authClients = append(authClients, authClient)
	}

	factorAddr, _, err := ethbridge.DeployChallengeFactory(ctx, authClients[0], client)
	if err != nil {
		t.Fatal(err)
	}

	testerAddress, _, err = authClients[0].MakeContract(ctx, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
		return ethbridgetestcontracts.DeployChallengeTester(auth, client, factorAddr)
	})
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
		testInboxTopChallenge(t, ctx, client, authClients[0], authClients[1])
	})
	t.Run("Execution Challenge", func(t *testing.T) {
		testExecutionChallenge(t, ctx, client, authClients[4], authClients[5])
	})
}
