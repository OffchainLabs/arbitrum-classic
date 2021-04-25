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
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
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
