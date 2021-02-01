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

package machine

import (
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type ArbStorage interface {
	DeleteCheckpoint(machineHash common.Hash) bool
	Initialize(contractPath string) error
	Initialized() bool
	CloseArbStorage() bool
	GetInitialMachine() (Machine, error)
	GetMachine(machineHash common.Hash) (Machine, error)
	SaveValue(val value.Value) bool
	GetValue(hashValue common.Hash) (value.Value, error)
	DeleteValue(hashValue common.Hash) bool
	SaveData(key []byte, serializedValue []byte) bool
	GetData(key []byte) ([]byte, error)
	DeleteData(key []byte) bool
}

type ValueNotFoundError struct {
	HashValue common.Hash
}

func (e *ValueNotFoundError) Error() string {
	return fmt.Sprintf("value with hash %s not found", e.HashValue.String())
}

type MachineNotFoundError struct {
	HashValue common.Hash
}

func (e *MachineNotFoundError) Error() string {
	return fmt.Sprintf("machine with hash %s not found", e.HashValue.String())
}

type DataNotFoundError struct {
	Key []byte
}

func (e *DataNotFoundError) Error() string {
	return fmt.Sprintf("data with key 0x%x not found", e.Key)
}
