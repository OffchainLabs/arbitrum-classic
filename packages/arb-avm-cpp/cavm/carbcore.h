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
#ifndef carbcore_h
#define carbcore_h

#include "ctypes.h"

#ifdef __cplusplus
extern "C" {
#endif

int arbCoreStartThread(CArbCore* arbcore_ptr);
void arbCoreAbortThread(CArbCore* arbcore_ptr);
int arbCoreMachineIdle(CArbCore* arbcore_ptr);
int arbCoreMessagesStatus(CArbCore* arbcore_ptr);
char* arbCoreMessagesClearError(CArbCore* arbcore_ptr);

int arbCoreDeliverMessages(CArbCore* arbcore_ptr,
                           ByteSliceArray inbox_messages,
                           void* previous_inbox_hash_ptr,
                           int last_block_complete);

Uint256Result arbCoreGetLogCount(CArbCore* arbcore_ptr);

ByteSliceArrayResult arbCoreGetLogs(CArbCore* arbcore_ptr,
                                    const void* start_index_ptr,
                                    const void* count_ptr);

Uint256Result arbCoreGetSendCount(CArbCore* arbcore_ptr);

ByteSliceArrayResult arbCoreGetSends(CArbCore* arbcore_ptr,
                                     const void* start_index_ptr,
                                     const void* count_ptr);

Uint256Result arbCoreGetMessageCount(CArbCore* arbcore_ptr);

ByteSliceArrayResult arbCoreGetMessages(CArbCore* arbcore_ptr,
                                        const void* start_index_ptr,
                                        const void* count_ptr);

HashList arbCoreGetMessageHashes(CArbCore* arbcore_ptr,
                                 const void* start_index_ptr,
                                 const void* count_ptr);

int arbCoreGetInboxDelta(CArbCore* arbcore_ptr,
                         const void* start_index_ptr,
                         const void* count_ptr,
                         void* ret);

int arbCoreGetInboxAcc(CArbCore* arbcore_ptr, const void* index_ptr, void* ret);
int arbCoreGetSendAcc(CArbCore* arbcore_ptr,
                      const void* start_acc_hash,
                      const void* start_index_ptr,
                      const void* count_ptr,
                      void* ret);
int arbCoreGetLogAcc(CArbCore* arbcore_ptr,
                     const void* start_acc_hash,
                     const void* start_index_ptr,
                     const void* count_ptr,
                     void* ret);

int arbCoreLogsCursorRequest(CArbCore* arbcore_ptr, const void* count);
ByteSliceArrayResult arbCoreLogsCursorGetLogs(CArbCore* arbcore_ptr);
ByteSliceArrayResult arbCoreLogsCursorGetDeletedLogs(CArbCore* arbcore_ptr);
int arbCoreLogsCursorSetConfirmedCount(CArbCore* arbcore_ptr,
                                       const void* count_ptr);
int arbCoreLogsCursorCheckError(CArbCore* arbcore_ptr);
char* arbCoreLogsCursorClearError(CArbCore* arbcore_ptr);

CExecutionCursor* arbCoreGetExecutionCursor(CArbCore* arbcore_ptr,
                                            const void* total_gas_used_ptr);
int arbCoreAdvanceExecutionCursor(CArbCore* arbcore_ptr,
                                  CExecutionCursor* execution_cursor_ptr,
                                  const void* max_gas_ptr,
                                  int go_over_gas);

#ifdef __cplusplus
}
#endif

#endif /* carbcore_h */
