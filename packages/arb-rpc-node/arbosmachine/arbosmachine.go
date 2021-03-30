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

package arbosmachine

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger = log.With().Stack().Str("component", "arbosmachine").Logger()

type Machine struct {
	machine.Machine
}

func New(mach machine.Machine) *Machine {
	return &Machine{Machine: mach}
}

func (m *Machine) Clone() machine.Machine {
	return &Machine{Machine: m.Machine.Clone()}
}

type callFrameAddresses struct {
	executor common.Address
	code     common.Address
}

func (cf callFrameAddresses) MarshalZerologObject(e *zerolog.Event) {
	e.Str("executor", cf.executor.Hex()).Str("code", cf.code.Hex())
}

func (cf callFrameAddresses) same() bool {
	return cf.executor == cf.code
}

func getAddresses(val value.Value) (callFrameAddresses, bool) {
	tup, ok := val.(*value.TupleValue)
	if !ok || tup.Len() != 2 {
		return callFrameAddresses{}, false
	}
	executorAddressVal, _ := tup.GetByInt64(0)
	codeAddressVal, _ := tup.GetByInt64(1)

	executorAddressInt, ok := executorAddressVal.(value.IntValue)
	if !ok {
		return callFrameAddresses{}, false
	}

	codeAddressInt, ok := codeAddressVal.(value.IntValue)
	if !ok {
		return callFrameAddresses{}, false
	}
	cf := callFrameAddresses{
		executor: inbox.NewAddressFromInt(executorAddressInt),
		code:     inbox.NewAddressFromInt(codeAddressInt),
	}
	return cf, true
}

func generateLog(txID, currentFrame, parentFrame value.Value, kind string) (*zerolog.Event, bool) {
	txIDInt, ok := txID.(value.IntValue)
	if !ok {
		return nil, false
	}
	txIDBytes := txIDInt.ToBytes()

	current, ok := getAddresses(currentFrame)
	if !ok {
		return nil, false
	}

	parent, ok := getAddresses(parentFrame)
	if !ok {
		return nil, false
	}
	logger := logger.Warn().
		Hex("tx_id", txIDBytes[:]).
		Str("kind", kind)
	if current.same() {
		logger = logger.Str("current", current.code.Hex())
	} else {
		logger = logger.Object("current", current)
	}

	if parent.same() {
		emptyAddress := common.Address{}
		if parent.code != emptyAddress {
			logger = logger.Str("parent", parent.code.Hex())
		}
	} else {
		logger = logger.Object("parent", parent)
	}
	return logger, true
}

func decodeEVMCallError(errorCode uint64) string {
	switch errorCode {
	case 0:
		return "application code error"
	case 1:
		return "Failed to transfer eth balance to contract: insufficient balance or unknown error"
	case 2:
		return "Can't pay for gas to contract"
	case 3:
		return "Failed to transfer eth balance to EOA"
	case 4:
		return "Can't pay for gas in constructor"
	case 5:
		return "Somehow the constructor didn't have storage"
	case 7:
		return "Should never reach end of call entry function"
	case 8:
		return "Should never reach end of call return function"
	case 9:
		return "Should never reach end of call return function"
	case 10:
		return "Called evmCallStack_getTopFrameMemoryOrDie while not in global stack frame"
	case 11:
		return "EVM code tried to jump to a forbidden EVM jump destination"
	case 12:
		return "Shouldn't reach at end of evmOp_getjumpaddr function"
	case 14:
		return "Called evmCallStack_queueMessage while not in global stack frame"
	case 15:
		return "Can't pay for gas in constructor"
	case 17:
		return "Called arbAddressTable_txcall while not in EVM tx"
	case 18:
		return "Called arbBLS_txcall or arbFunctionTable_txcall while not in EVM tx"
	case 19:
		return "Called arbosTest_txcall while not in EVM tx or error in snapshotAuxStack or restoreAuxStackAndCall"
	case 20:
		return "Chain hasn't been initialized"
	case 21:
		return "Called arbsys_txcall while not in EVM tx"
	case 22:
		return "Called arbowner_txcall while not in EVM tx"
	case 23:
		return "generateCodeForEvmSegment pushN without data"
	default:
		return "unknown"
	}
}

func handleDebugPrint(d value.Value) bool {
	tup, ok := d.(*value.TupleValue)
	if !ok || tup.Len() == 0 {
		return false
	}
	// Tuple already checked to be at least size 1
	debugPrintType, _ := tup.GetByInt64(0)
	debugPrintTypeInt, ok := debugPrintType.(value.IntValue)
	if !ok {
		return false
	}
	typ := debugPrintTypeInt.BigInt().Uint64()
	if typ == 664 {
		if tup.Len() != 2 {
			return false
		}
		subCodeVal, _ := tup.GetByInt64(1)
		subCodeInt, ok := subCodeVal.(value.IntValue)
		if !ok {
			return false
		}
		errorCode := subCodeInt.BigInt().Uint64()
		errorStr := decodeEVMCallError(errorCode)
		logger.Warn().
			Str("description", errorStr).
			Str("kind", "evm_call_error").
			Uint64("error_code", errorCode).
			Str("description", errorStr).
			Msg("debugprint")
		return true
	} else if typ == 666 && tup.Len() == 1 {
		logger.Error().Str("kind", "error_in_error_handler").Msg("debugprint")
		return true
	} else if typ == 665 || typ == 666 {
		var kind string
		if typ == 665 {
			kind = "out_of_gas"
		} else {
			kind = "hit_error_handler"
		}

		if tup.Len() != 4 {
			return false
		}
		txID, _ := tup.GetByInt64(1)
		currentFrame, _ := tup.GetByInt64(2)
		parentFrame, _ := tup.GetByInt64(3)

		logger, ok := generateLog(txID, currentFrame, parentFrame, kind)
		if !ok {
			return false
		}
		logger.Msg("debugprint")
		return true
	} else if typ == 10000 {
		if tup.Len() != 6 {
			return false
		}
		evmPC, _ := tup.GetByInt64(2)
		txID, _ := tup.GetByInt64(3)
		currentFrame, _ := tup.GetByInt64(4)
		parentFrame, _ := tup.GetByInt64(5)

		evmPCInt, ok := evmPC.(value.IntValue)
		if !ok {
			return false
		}

		logger, ok := generateLog(txID, currentFrame, parentFrame, "evm_revert")
		if !ok {
			return false
		}
		logger.Uint64("pc", evmPCInt.BigInt().Uint64()).Msg("debugprint")
		return true
	} else {
		return false
	}
}

func handleDebugPrints(debugPrints []value.Value) {
	for _, d := range debugPrints {
		if !handleDebugPrint(d) {
			logger.Debug().Str("raw", d.String()).Msg("debugprint")
		}
	}
}

func (m *Machine) ExecuteAssertion(
	maxGas uint64,
	goOverGas bool,
	messages []inbox.InboxMessage,
	finalMessageOfBlock bool,
) (*protocol.ExecutionAssertion, []value.Value, uint64) {
	assertion, debugPrints, numSteps := m.Machine.ExecuteAssertion(maxGas, goOverGas, messages, finalMessageOfBlock)
	handleDebugPrints(debugPrints)
	return assertion, debugPrints, numSteps
}
