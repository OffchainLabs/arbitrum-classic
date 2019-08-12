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

package goloader

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-go/vm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type RawExtension struct {
	id   uint32
	data []byte
}

type Error struct {
	str string
}

func (le Error) Error() string {
	return le.str
}

func LoadMachineFromFile(fileName string, warnMode bool) (*vm.Machine, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return LoadMachine(f, warnMode)
}

const CurrentAOVersion uint32 = 1

func LoadMachine(rd io.Reader, warnMode bool) (*vm.Machine, error) {
	var aoVersion uint32
	err := binary.Read(rd, binary.BigEndian, &aoVersion)
	if err != nil {
		return nil, err
	}

	if aoVersion != CurrentAOVersion {
		return nil, fmt.Errorf("AO file has unsupported version %v", aoVersion)
	}

	extensions := make([]RawExtension, 0)
	var extensionID uint32 = 1
	for extensionID != 0 {
		err := binary.Read(rd, binary.BigEndian, &extensionID)
		if err != nil {
			return nil, err
		}
		if extensionID > 0 {
			var extensionLength uint32
			err := binary.Read(rd, binary.BigEndian, &extensionLength)
			if err != nil {
				return nil, err
			}
			extensionData := make([]byte, extensionLength)
			_, err = rd.Read(extensionData)
			if err != nil {
				return nil, err
			}
			extensions = append(extensions, RawExtension{
				id:   extensionID,
				data: extensionData,
			})
		}
	}

	var insnsLen uint64
	err = binary.Read(rd, binary.BigEndian, &insnsLen)
	if err != nil {
		return nil, err
	}
	insns := make([]value.Operation, insnsLen)
	for i := uint64(0); i < insnsLen; i++ {
		insns[i], err = value.NewOperationFromReader(rd)
		if err != nil {
			return nil, err
		}
	}

	static, err2 := value.UnmarshalValue(rd)
	if err2 != nil {
		return nil, err2
	}

	maxSize := int64(1) << 62
	return vm.NewMachine(insns, static, warnMode, maxSize), nil
}
