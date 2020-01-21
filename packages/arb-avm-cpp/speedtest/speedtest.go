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

package main

import (
	"errors"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"io/ioutil"
	"log"
	"runtime"
	"strings"
	"time"
)

func runningTimeForAoFile(filePath string) (time.Duration, uint32, error) {
	ckpDir, err := ioutil.TempDir("/tmp", "speedtest-dummy-ckp")
	if err != nil {
		return time.Duration(0), 0, err
	}
	ckp, err := cmachine.NewCheckpoint(ckpDir, filePath)
	if err != nil {
		return time.Duration(0), 0, err
	}
	mach, err := ckp.GetInitialMachine()
	if err != nil {
		return time.Duration(0), 0, err
	}

	runtime.GC()
	oneStepDuration, nsteps := TimeForSteps(mach.Clone(), 1)
	if nsteps != 1 {
		return time.Duration(0), 0, errors.New("one-step run didn't execute one step")
	}
	runtime.GC()
	durationToCompletion, nsteps := TimeForSteps(mach.Clone(), 1<<25)

	diffDuration := durationToCompletion - oneStepDuration
	return diffDuration, nsteps - 1, nil
}

func TimeForSteps(mach machine.Machine, maxSteps int32) (time.Duration, uint32) {
	unusedTimeBounds := protocol.NewTimeBounds(0, 0)
	startTime := time.Now()
	assn := mach.ExecuteAssertion(maxSteps, unusedTimeBounds)
	endTime := time.Now()
	timeDiff := endTime.Sub(startTime)
	return timeDiff, assn.NumSteps
}

func main() {
	_aos := getAos()
	for _, fn := range _aos {
		testAoAndPrint(fn, fn)
	}
}

func testAoAndPrint(filePath, nickname string) {
	fmt.Println("===", nickname)
	dur, nsteps, err := runningTimeForAoFile(filePath)
	if err == nil {
		durPerStep := 1000.0 * float64(nsteps) / float64(dur.Nanoseconds())
		fmt.Println(nickname, dur, nsteps, durPerStep)
	}
}

func getAos() []string {
	ret := []string{}
	fileInfos, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), ".ao") {
			ret = append(ret, fileInfo.Name())
		}
	}
	return ret
}
