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
	"bytes"
	"crypto/ecdsa"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

var preLondon = false

func SimulatedBackend(t *testing.T) (*backends.SimulatedBackend, []*bind.TransactOpts) {
	genesisAlloc := make(map[ethcommon.Address]ethcore.GenesisAccount)
	auths := make([]*bind.TransactOpts, 0)
	balance, _ := new(big.Int).SetString("10000000000000000000", 10) // 10 eth in wei
	for i := 0; i < 15; i++ {
		// Intentionally use weak randomness since this is for testing and we
		// want to be able to reproduce
		randBytes := make([]byte, 64)
		_, err := rand.Read(randBytes)
		FailIfError(t, err)
		reader := bytes.NewReader(randBytes)
		privateKey, err := ecdsa.GenerateKey(crypto.S256(), reader)
		FailIfError(t, err)
		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
		FailIfError(t, err)
		auths = append(auths, auth)
		genesisAlloc[crypto.PubkeyToAddress(privateKey.PublicKey)] = ethcore.GenesisAccount{
			Balance: balance,
		}
	}

	blockGasLimit := uint64(1000000000)
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)
	if preLondon {
		for _, auth := range auths {
			auth.GasPrice = big.NewInt(0)
		}
		client.Blockchain().Config().LondonBlock.SetUint64(10000000000)
	}
	return client, auths
}

func FailIfError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		type stackTracer interface {
			StackTrace() errors.StackTrace
		}
		sterr, ok := err.(stackTracer)
		if ok {
			t.Log("stack trace for error", sterr.StackTrace())
		}
		t.Fatal(err)
	}
}
