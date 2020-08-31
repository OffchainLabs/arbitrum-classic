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
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
)

func getInsnMultiplier(filePath string) uint64 {
	ll := len(filePath)
	numPopsStr := filePath[ll-6 : ll-5]
	numPops, err := strconv.Atoi(numPopsStr)
	if err != nil {
		log.Fatal(filePath, " ", err)
	}
	numPushesStr := filePath[ll-8 : ll-7]
	numPushes, err := strconv.Atoi(numPushesStr)
	if err != nil {
		log.Fatal(filePath, " ", err)
	}
	numExtraUnderscores := strings.Count(filePath, "_") - 2
	return uint64(1 + numExtraUnderscores + numPops + numPushes)
}

func runExecutableFile(b *testing.B, filePath string) {
	insnMultiplier := getInsnMultiplier(filePath)
	ckpDir, err := ioutil.TempDir("/tmp", "speedtest-dummy-ckp")
	if err != nil {
		b.Fatal(err)
	}
	ckp, err := cmachine.NewCheckpoint(ckpDir)
	if err != nil {
		b.Fatal(err)
	}
	if err := ckp.Initialize(filePath); err != nil {
		b.Fatal(err)
	}
	mach, err := ckp.GetInitialMachine()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	_, _ = mach.ExecuteAssertion(uint64(b.N)*insnMultiplier, nil, time.Hour)
}

func nameFromFn(fn string) string {
	ll := len(fn)
	fnSlices := strings.Split(fn[:ll-7], "/")
	ret := fnSlices[len(fnSlices)-1]
	numPopsStr := fn[ll-6 : ll-5]
	numPops, err := strconv.Atoi(numPopsStr)
	if err != nil {
		log.Fatal(fn, " ", err)
	}
	numPushesStr := fn[ll-8 : ll-7]
	numPushes, err := strconv.Atoi(numPushesStr)
	if err != nil {
		log.Fatal(fn, " ", err)
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
	executables := getExecutables()
	for _, fn := range executables {
		b.Run(nameFromFn(fn), func(b *testing.B) {
			runExecutableFile(b, fn)
		})
	}
}

func getExecutables() []string {
	ret := []string{}
	fileInfos, err := ioutil.ReadDir("./executables/")
	if err != nil {
		log.Fatal(err)
	}
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), ".mexe") {
			ret = append(ret, "executables/"+fileInfo.Name())
		}
	}
	return ret
}
