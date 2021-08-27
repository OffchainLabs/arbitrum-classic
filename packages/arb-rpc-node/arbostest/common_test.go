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

package arbostest

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"math/big"
	"os"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var (
	chainId      = big.NewInt(764575)
	owner        = common.HexToAddress("0xcd3CFd7829e7d49e1847eA37fc4057537ee5e72f")
	chain        = common.HexToAddress("0x037c4d7bbb0407d1e2c64981855ad8681d0d86d1")
	sender       = common.HexToAddress("0xe91e00167939cb6694d2c422acd208a007293948")
	connAddress1 = common.HexToAddress("0x2aad3e8302f74e0818b7bcd10c2c050526707755")
	connAddress2 = common.HexToAddress("0x016cb751543d1cca5dd02976ac8dbdc0ecaacafd")
)

var arbosfile *string
var arbosVersion int
var doUpgrade bool

type ArbOSExec struct {
	Version *int `json:"arbos_version"`
}

func TestMain(m *testing.M) {

	doUpgrade = flag.Bool("upgrade", false, "Test against an upgraded ArbOS. Overrides 'arbos' flag.");
	
	arbosPath, err := arbos.Path(doUpgrade)
	if err != nil {
		panic(err)
	}

	if doUpgrade {
		arbosfile = arbosPath
	} else {
		arbosfile = flag.String("arbos", arbosPath, "Version of arbos to run tests against")
	}
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
