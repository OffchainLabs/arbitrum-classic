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
	"log"
	"math/big"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

func getInsnMultiplier(filePath string) uint64 {
	ll := len(filePath)
	numPopsStr := filePath[ll-4 : ll-3]
	numPops, err := strconv.Atoi(numPopsStr)
	if err != nil {
		log.Fatal(err)
	}
	numPushesStr := filePath[ll-6 : ll-5]
	numPushes, err := strconv.Atoi(numPushesStr)
	if err != nil {
		log.Fatal(err)
	}
	numExtraUnderscores := strings.Count(filePath, "_") - 2
	return uint64(1 + numExtraUnderscores + numPops + numPushes)
}

func runAoFile(b *testing.B, filePath string) {
	insnMultiplier := getInsnMultiplier(filePath)
	ckpDir, err := ioutil.TempDir("/tmp", "speedtest-dummy-ckp")
	if err != nil {
		b.Fatal(err)
	}
	ckp, err := cmachine.NewCheckpoint(ckpDir, filePath)
	if err != nil {
		b.Fatal(err)
	}
	mach, err := ckp.GetInitialMachine()
	if err != nil {
		b.Fatal(err)
	}

	unusedTimeBounds := &protocol.TimeBoundsBlocks{
		Start: common.NewTimeBlocks(big.NewInt(0)),
		End:   common.NewTimeBlocks(big.NewInt(0)),
	}
	b.ResetTimer()
	_, _ = mach.ExecuteAssertion(uint64(b.N)*insnMultiplier, unusedTimeBounds, value.NewEmptyTuple(), time.Hour)
}

func nameFromFn(fn string) string {
	ll := len(fn)
	fnSlices := strings.Split(fn[:ll-7], "/")
	ret := fnSlices[len(fnSlices)-1]
	numPopsStr := fn[ll-4 : ll-3]
	numPops, err := strconv.Atoi(numPopsStr)
	if err != nil {
		log.Fatal(err)
	}
	numPushesStr := fn[ll-6 : ll-5]
	numPushes, err := strconv.Atoi(numPushesStr)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < numPushes; i++ {
		ret = "push_" + ret
	}
	for i := 0; i < numPops; i++ {
		ret += "_pop"
	}
	return ret
}

func BenchmarkInsns(b *testing.B) {
	_aos := getAos()
	for _, fn := range _aos {
		op := fn
		b.Run(nameFromFn(op), func(b *testing.B) {
			runAoFile(b, op)
		})
	}
}

func getAos() []string {
	ret := []string{}
	fileInfos, err := ioutil.ReadDir("./aos/")
	if err != nil {
		log.Fatal(err)
	}
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), ".ao") {
			ret = append(ret, "aos/"+fileInfo.Name())
		}
	}
	return ret
}
