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
	"github.com/offchainlabs/arbitrum/packages/arb-util/arblog"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/transactauth"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

var logger = arblog.Logger.With().Str("component", "ethbridge").Logger()

type Bridge struct {
	*DelayedBridgeWatcher
	auth transactauth.TransactAuth
}

func NewBridge(address ethcommon.Address, fromBlock int64, client ethutils.EthClient, auth transactauth.TransactAuth) (*Bridge, error) {
	watcher, err := NewDelayedBridgeWatcher(address, fromBlock, client)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Bridge{
		DelayedBridgeWatcher: watcher,
		auth:                 auth,
	}, nil
}

func (b *Bridge) SendL2MessageFromOrigin() {

}
