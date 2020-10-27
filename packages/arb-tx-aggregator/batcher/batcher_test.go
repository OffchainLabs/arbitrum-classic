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

package batcher

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"testing"
	"time"
)

func TestBatcher(t *testing.T) {
	client, pks := test.SimulatedBackend()
	l1Client := &ethutils.SimulatedEthClient{SimulatedBackend: client}
	auth := bind.NewKeyedTransactor(pks[0])
	inbox, _, _, err := ethbridgecontracts.DeployGlobalInbox(auth, client)
	if err != nil {
		t.Fatal(err)
	}
	client.Commit()

	ethAuth := ethbridge.NewEthAuthClient(l1Client, auth)
	chain := common.RandAddress()

	globalInbox, err := ethAuth.NewGlobalInbox(common.NewAddressFromEth(inbox), chain)
	if err != nil {
		t.Fatal(err)
	}

	batcher := NewBatcher(
		context.Background(),
		nil,
		chain,
		l1Client,
		globalInbox,
		time.Second,
		false,
	)
	t.Log(batcher)
}
