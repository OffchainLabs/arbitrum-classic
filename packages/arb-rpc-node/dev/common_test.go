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

package dev

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

var arbosfile *string
var arbosVersion int

type ArbOSExec struct {
	Version *int `json:"arbos_version"`
}

func TestMain(m *testing.M) {
	arbosPath, err := arbos.Path()
	if err != nil {
		panic(err)
	}

	arbosfile = flag.String("arbos", arbosPath, "version of arbos to run tests against")
	flag.Parse()

	fileData, err := ioutil.ReadFile(*arbosfile)
	if err != nil {
		panic(err)
	}
	var arbosExec ArbOSExec
	if err := json.Unmarshal(fileData, &arbosExec); err != nil {
		panic(err)
	}
	if arbosExec.Version != nil {
		arbosVersion = *arbosExec.Version
	} else {
		arbosVersion = 1
	}
	os.Exit(m.Run())
}

func skipBelowVersion(t *testing.T, ver int) {
	t.Helper()
	if arbosVersion < ver {
		t.Skipf("Skipping test because version %v too below supported version %v", arbosVersion, ver)
	}
}

func NewTestDevNode(
	t *testing.T,
	arbosPath string,
	params protocol.ChainParams,
	owner common.Address,
	config []message.ChainConfigOption,
) (*Backend, *txdb.TxDB, *aggregator.Server, func()) {
	tmpDir, err := ioutil.TempDir(".", "arbitrum")
	test.FailIfError(t, err)
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Fatal(err)
		}
	}()

	backend, db, rollupAddress, cancelDevNode, txDBErrChan, err := NewDevNode(
		context.Background(),
		tmpDir,
		arbosPath,
		params,
		owner,
		config,
	)
	test.FailIfError(t, err)
	go func() {
		if err := <-txDBErrChan; err != nil {
			t.Error(err)
		}
	}()
	closeFunc := func() {
		cancelDevNode()
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Fatal(err)
		}
	}
	srv := aggregator.NewServer(backend, rollupAddress, db)
	return backend, db, srv, closeFunc
}
