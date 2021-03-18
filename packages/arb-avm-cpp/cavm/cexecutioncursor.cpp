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

int executionCursorMachineHash(CExecutionCursor* execution_cursor_ptr,
                               void* ret) {
    auto executionCursor = static_cast<ExecutionCursor*>(execution_cursor_ptr);
    try {
        auto index_result = executionCursor->machineHash();
        if (!index_result) {
            // Unable to compute machine hash
            std::cerr << "unable to create machine hash for execution cursor"
                      << std::endl;
            return false;
        }
        std::array<unsigned char, 32> val{};
        to_big_endian(*index_result, val.begin());
        std::copy(val.begin(), val.end(), reinterpret_cast<char*>(ret));

        return true;
    } catch (const std::exception& e) {
        return false;
    }
}

int executionCursorInboxAcc(CExecutionCursor* execution_cursor_ptr, void* ret) {
    auto executionCursor = static_cast<ExecutionCursor*>(execution_cursor_ptr);
    try {
        auto acc = executionCursor->getInboxAcc();
        if (!acc.has_value()) {
            return false;
        }
        std::array<unsigned char, 32> val{};
        to_big_endian(*acc, val.begin());
        std::copy(val.begin(), val.end(), reinterpret_cast<char*>(ret));

        return true;
    } catch (const std::exception& e) {
        return false;
    }
}

int executionCursorSendAcc(CExecutionCursor* execution_cursor_ptr, void* ret) {
    auto executionCursor = static_cast<ExecutionCursor*>(execution_cursor_ptr);
    try {
        std::array<unsigned char, 32> val{};
        to_big_endian(executionCursor->getOutput().send_acc, val.begin());
        std::copy(val.begin(), val.end(), reinterpret_cast<char*>(ret));

        return true;
    } catch (const std::exception& e) {
        return false;
    }
}

int executionCursorLogAcc(CExecutionCursor* execution_cursor_ptr, void* ret) {
    auto executionCursor = static_cast<ExecutionCursor*>(execution_cursor_ptr);
    try {
        std::array<unsigned char, 32> val{};
        to_big_endian(executionCursor->getOutput().log_acc, val.begin());
        std::copy(val.begin(), val.end(), reinterpret_cast<char*>(ret));

        return true;
    } catch (const std::exception& e) {
        return false;
    }
}

Uint256Result executionCursorTotalMessagesRead(
    CExecutionCursor* execution_cursor_ptr) {
    try {
        auto index_result = static_cast<ExecutionCursor*>(execution_cursor_ptr)
                                ->getTotalMessagesRead();
        return {returnUint256(index_result), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

Uint256Result executionCursorTotalGasConsumed(
    CExecutionCursor* execution_cursor_ptr) {
    try {
        auto index_result = static_cast<ExecutionCursor*>(execution_cursor_ptr)
                                ->getOutput()
                                .arb_gas_used;
        return {returnUint256(index_result), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

Uint256Result executionCursorTotalSteps(
    CExecutionCursor* execution_cursor_ptr) {
    try {
        auto index_result = static_cast<ExecutionCursor*>(execution_cursor_ptr)
                                ->getOutput()
                                .total_steps;
        return {returnUint256(index_result), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

Uint256Result executionCursorTotalSendCount(
    CExecutionCursor* execution_cursor_ptr) {
    try {
        auto index_result = static_cast<ExecutionCursor*>(execution_cursor_ptr)
                                ->getOutput()
                                .send_count;
        return {returnUint256(index_result), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

Uint256Result executionCursorTotalLogCount(
    CExecutionCursor* execution_cursor_ptr) {
    try {
        auto index_result = static_cast<ExecutionCursor*>(execution_cursor_ptr)
                                ->getOutput()
                                .log_count;
        return {returnUint256(index_result), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}
