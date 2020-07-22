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
	"context"
	"github.com/gorilla/rpc/v2"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/aggregator"
)

func GenerateWeb3Server(ctx context.Context, server *aggregator.Server) (*rpc.Server, error) {
	web3Server, err := NewServer(ctx, server)
	if err != nil {
		return nil, err
	}
	s := rpc.NewServer()

	// Register our own Codec
	s.RegisterCodec(NewUpCodec(), "application/json")
	s.RegisterCodec(NewUpCodec(), "application/json;charset=UTF-8")

	err = s.RegisterService(web3Server, "Eth")
	if err != nil {
		panic(err)
	}

	err = s.RegisterService(new(Net), "Net")
	if err != nil {
		panic(err)
	}

	return s, nil
}
