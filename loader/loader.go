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

package loader

import (
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-avm/vm"

	"encoding/binary"
	"io"
	"os"
)

func LoadMachineFromFile(fileName string, warnMode bool) (*vm.Machine, []string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	return LoadMachine(f, warnMode)
}

func LoadInsnsStaticAndLocationsFromFile(fileName string) ([]value.Operation, value.Value, []string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, nil, nil, err
	}
	defer f.Close()

	return LoadInsnsStaticAndLocations(f)
}

func LoadMachine(rd io.Reader, warnMode bool) (*vm.Machine, []string, error) {
	insns, static, locations, err := LoadInsnsStaticAndLocations(rd)
	if err != nil {
		return nil, nil, err
	}
	maxSize := int64(1) << 62
	return vm.NewMachine(insns, static, warnMode, locations, maxSize), locations, nil
}

func LoadInsnsStaticAndLocations(rd io.Reader) ([]value.Operation, value.Value, []string, error) {
	var insnsLen int64
	err := binary.Read(rd, binary.BigEndian, &insnsLen)
	if err != nil {
		return nil, nil, nil, err
	}
	insns := make([]value.Operation, insnsLen)
	for i := int64(0); i < insnsLen; i++ {
		insns[i], err = value.NewOperationFromReader(rd)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	locations := make([]string, insnsLen)
	//for i := int64(0); i < insnsLen; i++ {
	//	locations[i], err = arbutil.ReadString(rd)
	//	if err != nil {
	//		return nil, nil, nil, err
	//	}
	//}

	static, err2 := value.UnmarshalValue(rd)
	if err2 != nil {
		return nil, nil, nil, err2
	}

	return insns, static, locations, nil
}

type Error struct {
	str string
}

func (le Error) Error() string {
	return le.str
}
