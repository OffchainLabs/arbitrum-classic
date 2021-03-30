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

package gotest

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

func OpCodeTestFiles() []string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		err := errors.New("failed to get filename")
		panic(err)
	}
	testCaseDir := filepath.Join(filepath.Dir(filename), "../tests/machine-cases")
	files, err := ioutil.ReadDir(testCaseDir)
	if err != nil {
		log.Fatal(err)
	}
	filenames := make([]string, 0, len(files))
	for _, file := range files {
		if file.Name() != "inbox.mexe" {
			filenames = append(filenames, filepath.Join(testCaseDir, file.Name()))
		}
	}

	return filenames
}

func ArbOSTestFiles() []string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		err := errors.New("failed to get filename")
		panic(err)
	}
	testCaseDir := filepath.Join(filepath.Dir(filename), "../../arb-os/replayTests")
	files, err := ioutil.ReadDir(testCaseDir)
	if err != nil {
		log.Fatal(err)
	}
	filenames := make([]string, 0, len(files))
	extensions := []string{".aoslog", ".json"}
	for _, file := range files {
		name := file.Name()
		for _, extension := range extensions {
			if len(name) > len(extension) && name[len(name)-len(extension):] == extension {
				filenames = append(filenames, filepath.Join(testCaseDir, file.Name()))
				break
			}
		}
	}

	return filenames
}
