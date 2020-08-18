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

package web3

import (
	"github.com/gorilla/rpc/v2"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func GenerateWeb3Server(server *aggregator.Server) (*rpc.Server, error) {
	s := rpc.NewServer()
	// Register our own Codec
	s.RegisterCodec(NewUpCodec(), "application/json")
	s.RegisterCodec(NewUpCodec(), "application/json;charset=UTF-8")

	err := s.RegisterService(NewServer(server), "Eth")
	if err != nil {
		panic(err)
	}

	net := &Net{chainId: message.ChainAddressToID(common.NewAddressFromEth(server.GetChainAddress())).Uint64()}
	err = s.RegisterService(net, "Net")
	if err != nil {
		panic(err)
	}

	web3 := &Web3{}
	err = s.RegisterService(web3, "Web3")
	if err != nil {
		panic(err)
	}

	return s, nil
}
