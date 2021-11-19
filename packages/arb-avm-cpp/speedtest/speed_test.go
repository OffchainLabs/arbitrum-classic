/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

package speedtest

import (
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

func getInsnMultiplier(b *testing.B, filePath string) uint64 {
	ll := len(filePath)
	numPopsStr := filePath[ll-6 : ll-5]
	numPops, err := strconv.Atoi(numPopsStr)
	if err != nil {
		b.Log("filepath", filePath)
		b.Log("numPops failed string conversion")
		b.Fatal(err)
	}
	numPushesStr := filePath[ll-8 : ll-7]
	numPushes, err := strconv.Atoi(numPushesStr)
	if err != nil {
		b.Log("filepath", filePath)
		b.Log("numPushes failed string conversion")
		b.Fatal(err)
	}
	numExtraUnderscores := strings.Count(filePath, "_") - 2
	return uint64(1 + numExtraUnderscores + numPops + numPushes)
}

func runExecutableFile(b *testing.B, filePath string) {
	insnMultiplier := getInsnMultiplier(b, filePath)
	coreConfig := configuration.DefaultCoreSettingsMaxExecution()
	ckp, err := cmachine.NewArbStorage(b.TempDir(), coreConfig)
	if err != nil {
		b.Fatal(err)
	}
	if err := ckp.Initialize(filePath); err != nil {
		b.Fatal(err)
	}
	core := ckp.GetArbCore()
	cursor, err := core.GetExecutionCursor(big.NewInt(0), true)
	if err != nil {
		b.Fatal(err)
	}
	mach, err := core.TakeMachine(cursor)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	// Last parameter returned is number of steps executed
	_, _, _, err = mach.ExecuteAssertion(uint64(b.N)*insnMultiplier, true, nil)
	if err != nil {
		b.Fatal(err)
	}
}

func nameFromFn(b *testing.B, fn string) string {
	ll := len(fn)
	fnSlices := strings.Split(fn[:ll-7], "/")
	ret := fnSlices[len(fnSlices)-1]
	numPopsStr := fn[ll-6 : ll-5]
	numPops, err := strconv.Atoi(numPopsStr)
	if err != nil {
		b.Log("fn", fn)
		b.Log("numPops failed string conversion")
		b.Fatal(err)
	}
	numPushesStr := fn[ll-8 : ll-7]
	numPushes, err := strconv.Atoi(numPushesStr)
	if err != nil {
		b.Log("fn", fn)
		b.Log("numPushes failed string conversion")
		b.Fatal(err)
	}
	for i := 0; i < numPushes; i++ {
		ret = "push_" + ret
	}
	for i := 0; i < numPops; i++ {
		ret = ret + "_pop"
	}
	return ret
}

func BenchmarkInsns(b *testing.B) {
	executables := getExecutables(b)
	for _, fn := range executables {
		b.Run(nameFromFn(b, fn), func(b *testing.B) {
			runExecutableFile(b, fn)
		})
	}
}

func getExecutables(b *testing.B) []string {
	var ret []string
	fileInfos, err := ioutil.ReadDir("./executables/")
	if err != nil {
		b.Log("Error reading executables directory")
		b.Fatal(err)
	}
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), ".mexe") {
			ret = append(ret, "executables/"+fileInfo.Name())
		}
	}
	return ret
}
