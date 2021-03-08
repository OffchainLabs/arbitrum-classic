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

package cmachine

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lcavm -lavm -ldata_storage -lavm_values -lstdc++ -lm -lrocksdb -ldl
#include "../cavm/cexecutioncursor.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"math/big"
	"runtime"
	"unsafe"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/pkg/errors"
)

type ExecutionCursor struct {
	c                 unsafe.Pointer
	machineHash       common.Hash
	totalMessagesRead *big.Int
	inboxAcc          common.Hash
	sendAcc           common.Hash
	logAcc            common.Hash
	totalGasConsumed  *big.Int
	totalSendCount    *big.Int
	totalLogCount     *big.Int
	totalSteps        *big.Int
}

func deleteExecutionCursor(ac *ExecutionCursor) {
	C.deleteExecutionCursor(ac.c)
}

func NewExecutionCursor(c unsafe.Pointer) (*ExecutionCursor, error) {
	ec := &ExecutionCursor{c: c}
	runtime.SetFinalizer(ec, deleteExecutionCursor)

	err := ec.updateValues()
	if err != nil {
		return nil, err
	}

	return ec, nil
}

func (ec *ExecutionCursor) Clone() core.ExecutionCursor {
	return &ExecutionCursor{
		c:                 C.executionCursorClone(ec.c),
		machineHash:       ec.machineHash,
		totalMessagesRead: ec.totalMessagesRead,
		inboxAcc:          ec.inboxAcc,
		sendAcc:           ec.sendAcc,
		logAcc:            ec.logAcc,
		totalGasConsumed:  ec.totalGasConsumed,
		totalSendCount:    ec.totalSendCount,
		totalLogCount:     ec.totalLogCount,
		totalSteps:        ec.totalSteps,
	}
}

func (ec *ExecutionCursor) TakeMachine() (machine.Machine, error) {
	cMachine := C.executionCursorTakeMachine(ec.c)
	if cMachine == nil {
		return nil, errors.Errorf("error taking machine from execution cursor")
	}
	ret := &Machine{cMachine}

	runtime.SetFinalizer(ret, cdestroyVM)
	return ret, nil
}

func (ec *ExecutionCursor) updateValues() error {
	status := C.executionCursorMachineHash(ec.c, unsafe.Pointer(&ec.machineHash[0]))
	if status == 0 {
		return errors.New("failed to load machine hash")
	}

	status = C.executionCursorInboxAcc(ec.c, unsafe.Pointer(&ec.inboxAcc[0]))
	if status == 0 {
		return errors.New("failed to load inbox acc")
	}

	status = C.executionCursorSendAcc(ec.c, unsafe.Pointer(&ec.sendAcc[0]))
	if status == 0 {
		return errors.New("failed to load send acc")
	}

	status = C.executionCursorLogAcc(ec.c, unsafe.Pointer(&ec.logAcc[0]))
	if status == 0 {
		return errors.New("failed to load log acc")
	}

	result := C.executionCursorTotalMessagesRead(ec.c)
	if result.found == 0 {
		return errors.New("failed to get TotalMessagesRead")
	}
	ec.totalMessagesRead = receiveBigInt(result.value)

	result = C.executionCursorTotalGasConsumed(ec.c)
	if result.found == 0 {
		return errors.New("failed to get TotalGasConsumed")
	}
	ec.totalGasConsumed = receiveBigInt(result.value)

	result = C.executionCursorTotalSendCount(ec.c)
	if result.found == 0 {
		return errors.New("failed to get TotalSendCount")
	}
	ec.totalSendCount = receiveBigInt(result.value)

	result = C.executionCursorTotalLogCount(ec.c)
	if result.found == 0 {
		return errors.New("failed to get NextInboxMessageIndex")
	}
	ec.totalLogCount = receiveBigInt(result.value)

	result = C.executionCursorTotalSteps(ec.c)
	if result.found == 0 {
		return errors.New("failed to get TotalSteps")
	}
	ec.totalSteps = receiveBigInt(result.value)

	return nil
}

func (ec *ExecutionCursor) MachineHash() common.Hash {
	return ec.machineHash
}

func (ec *ExecutionCursor) InboxAcc() common.Hash {
	return ec.inboxAcc
}

func (ec *ExecutionCursor) SendAcc() common.Hash {
	return ec.sendAcc
}

func (ec *ExecutionCursor) LogAcc() common.Hash {
	return ec.logAcc
}

func (ec *ExecutionCursor) TotalMessagesRead() *big.Int {
	return ec.totalMessagesRead
}

func (ec *ExecutionCursor) TotalGasConsumed() *big.Int {
	return ec.totalGasConsumed
}

func (ec *ExecutionCursor) TotalSendCount() *big.Int {
	return ec.totalSendCount
}

func (ec *ExecutionCursor) TotalLogCount() *big.Int {
	return ec.totalLogCount
}

func (ec *ExecutionCursor) TotalSteps() *big.Int {
	return ec.totalSteps
}
