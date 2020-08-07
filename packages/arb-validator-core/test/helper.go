/*
 * Copyright 2019, Offchain Labs, Inc.
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

package test

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

func SimulatedBackend() (*backends.SimulatedBackend, []*ecdsa.PrivateKey) {
	genesisAlloc := make(map[ethcommon.Address]core.GenesisAccount)
	pks := make([]*ecdsa.PrivateKey, 0)
	balance, _ := new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)
	for i := 0; i < 100; i++ {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
		pks = append(pks, privateKey)

		genesisAlloc[crypto.PubkeyToAddress(privateKey.PublicKey)] = core.GenesisAccount{
			Balance: balance,
		}
	}

	blockGasLimit := uint64(1300000000)
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)
	return client, pks
}
