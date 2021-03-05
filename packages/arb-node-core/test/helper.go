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
	"io/ioutil"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/rs/zerolog/log"
)

var logger = log.With().Caller().Str("component", "test").Logger()

func SimulatedBackend() (*backends.SimulatedBackend, []*ecdsa.PrivateKey) {
	genesisAlloc := make(map[ethcommon.Address]ethcore.GenesisAccount)
	pks := make([]*ecdsa.PrivateKey, 0)
	balance, _ := new(big.Int).SetString("10000000000000000000", 10) // 10 eth in wei
	for i := 0; i < 15; i++ {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			logger.Fatal().Stack().Err(err).Send()
		}
		pks = append(pks, privateKey)

		genesisAlloc[crypto.PubkeyToAddress(privateKey.PublicKey)] = ethcore.GenesisAccount{
			Balance: balance,
		}
	}

	blockGasLimit := uint64(1000000000)
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)
	return client, pks
}

func FailIfError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func PrepareArbCore(t *testing.T, messages []inbox.InboxMessage) (core.ArbCore, func()) {
	tmpDir, err := ioutil.TempDir("", "arbitrum")
	FailIfError(t, err)
	storage, err := cmachine.NewArbStorage(tmpDir)
	if err != nil {
		os.RemoveAll(tmpDir)
	}
	FailIfError(t, err)
	shutdown := func() {
		storage.CloseArbStorage()
		if err := os.RemoveAll(tmpDir); err != nil {
			panic(err)
		}
	}
	returning := false
	defer (func() {
		if !returning {
			shutdown()
		}
	})()

	err = storage.Initialize(arbos.Path())
	FailIfError(t, err)

	arbCore := storage.GetArbCore()
	started := arbCore.StartThread()
	if !started {
		t.Fatal("failed to start thread")
	}

	if len(messages) > 0 {
		_, err = core.DeliverMessagesAndWait(arbCore, messages, common.Hash{}, false)
		FailIfError(t, err)
	}
	for {
		msgCount, err := arbCore.GetMessageCount()
		FailIfError(t, err)
		if arbCore.MachineIdle() && msgCount.Cmp(big.NewInt(int64(len(messages)))) >= 0 {
			break
		}
		<-time.After(time.Millisecond * 200)
	}

	returning = true
	return arbCore, shutdown
}
