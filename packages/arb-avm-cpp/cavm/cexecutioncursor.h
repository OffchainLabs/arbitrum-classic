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
#ifndef cexecutioncursor_h
#define cexecutioncursor_h

#include "ctypes.h"

#include <cstdint>

#ifdef __cplusplus
extern "C" {
#endif

void deleteExecutionCursor(CExecutionCursor* m);

CExecutionCursor* executionCursorClone(CExecutionCursor* execution_cursor_ptr);

int executionCursorMachineHash(CExecutionCursor* execution_cursor_ptr,
                               void* ret);
int executionCursorInboxAcc(CExecutionCursor* execution_cursor_ptr, void* ret);
int executionCursorSendAcc(CExecutionCursor* execution_cursor_ptr, void* ret);
int executionCursorLogAcc(CExecutionCursor* execution_cursor_ptr, void* ret);
Uint256Result executionCursorTotalMessagesRead(
    CExecutionCursor* execution_cursor_ptr);
Uint256Result executionCursorTotalSteps(CExecutionCursor* execution_cursor_ptr);
Uint256Result executionCursorTotalGasConsumed(
    CExecutionCursor* execution_cursor_ptr);
Uint256Result executionCursorTotalSendCount(
    CExecutionCursor* execution_cursor_ptr);
Uint256Result executionCursorTotalLogCount(
    CExecutionCursor* execution_cursor_ptr);

#ifdef __cplusplus
}
#endif

#endif /* cexecutioncursor_h */
