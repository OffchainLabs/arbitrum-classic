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
	"crypto/ecdsa"
	"time"

	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/metrics"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
)

func GenerateWeb3Server(server *aggregator.Server, privateKeys []*ecdsa.PrivateKey, ganacheMode bool, plugins map[string]interface{}, metricsConfig *metrics.MetricsConfig) (*rpc.Server, error) {
	s := rpc.NewServer()

	ethServer := NewServer(server, ganacheMode, metricsConfig)

	if err := s.RegisterName("eth", ethServer); err != nil {
		return nil, err
	}

	if err := s.RegisterName("eth", filters.NewPublicFilterAPI(server, false, 2*time.Minute)); err != nil {
		return nil, err
	}

	if err := s.RegisterName("eth", NewAccounts(ethServer, privateKeys, metricsConfig)); err != nil {
		return nil, err
	}

	if err := s.RegisterName("arb", &Arb{srv: server, counter: metricsConfig.MethodCallCounter}); err != nil {
		return nil, err
	}

	if err := s.RegisterName("personal", NewPersonalAccounts(privateKeys, metricsConfig)); err != nil {
		return nil, err
	}

	net := &Net{chainId: server.ChainId().Uint64(), counter: metricsConfig.MethodCallCounter}
	if err := s.RegisterName("net", net); err != nil {
		return nil, err
	}

	if err := s.RegisterName("web3", &Web3{counter: metricsConfig.MethodCallCounter}); err != nil {
		return nil, err
	}

	for name, val := range plugins {
		if err := s.RegisterName(name, val); err != nil {
			return nil, err
		}
	}

	return s, nil
}
