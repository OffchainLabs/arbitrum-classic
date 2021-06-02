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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/prometheus/client_golang/prometheus"
)

type Web3 struct {
	counter *prometheus.CounterVec
}

func (web3 *Web3) ClientVersion() string {
	return "arb-rpc-node/v0.8.0"
}

func (web3 *Web3) Sha3(data hexutil.Bytes) hexutil.Bytes {
	web3.counter.WithLabelValues("web3_sha3", "true")
	return crypto.Keccak256(data)
}
