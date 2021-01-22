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

#include "cexecutioncursor.h"

#include "utils.hpp"

#include <data_storage/executioncursor.hpp>

void deleteExecutionCursor(CExecutionCursor* m) {
    delete static_cast<ExecutionCursor*>(m);
}

CExecutionCursor* executionCursorClone(CExecutionCursor* execution_cursor_ptr) {
    return static_cast<void*>(new ExecutionCursor(
        *static_cast<ExecutionCursor*>(execution_cursor_ptr)));
}

Uint256Result executionCursorMachineHash(
    CExecutionCursor* execution_cursor_ptr) {
    try {
        auto index_result =
            static_cast<ExecutionCursor*>(execution_cursor_ptr)->machineHash();
        return {returnUint256(index_result), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

Uint256Result executionCursorNextInboxMessageIndex(
    CExecutionCursor* execution_cursor_ptr) {
    try {
        auto index_result = static_cast<ExecutionCursor*>(execution_cursor_ptr)
                                ->message_sequence_number_processed +
                            1;
        return {returnUint256(index_result), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

Uint256Result executionCursorInboxHash(CExecutionCursor* execution_cursor_ptr) {
    try {
        auto index_result =
            static_cast<ExecutionCursor*>(execution_cursor_ptr)->inbox_hash;
        return {returnUint256(index_result), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

Uint256Result executionCursorTotalGasConsumed(
    CExecutionCursor* execution_cursor_ptr) {
    try {
        auto index_result =
            static_cast<ExecutionCursor*>(execution_cursor_ptr)->arb_gas_used;
        return {returnUint256(index_result), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}
