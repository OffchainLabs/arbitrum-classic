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

void deleteArbCore(CArbCore* m);

int arbCoreDeliverMessages(CArbCore* arb_core_ptr,
                           uint64_t first_message_sequence_number,
                           uint64_t block_height,
                           void* inbox_messages,
                           void* previous_inbox_hash_ptr);

ByteSliceArrayResult arbCoreGetSends(CArbCore* arb_core_ptr,
                                     const void* start_index_ptr,
                                     const void* count_ptr);

ByteSliceArrayResult arbCoreGetMessages(CArbCore* arb_core_ptr,
                                        const void* start_index_ptr,
                                        const void* count_ptr);

int arbCoreGetInboxDelta(CArbCore* arb_core_ptr,
                         const void* start_index_ptr,
                         const void* count_ptr,
                         void* ret);

int arbCoreGetInboxAcc(CArbCore* arb_core_ptr,
                       const void* index_ptr,
                       void* ret);
int arbCoreGetSendAcc(CArbCore* arb_core_ptr,
                      const void* start_acc_hash,
                      const void* start_index_ptr,
                      const void* count_ptr,
                      void* ret);
int arbCoreGetLogAcc(CArbCore* arb_core_ptr,
                     const void* start_acc_hash,
                     const void* start_index_ptr,
                     const void* count_ptr,
                     void* ret,
                     CValueCache* cache_ptr);

int arbCoreLogsCursorRequest(CArbCore* arb_core_ptr, const void* count);
ByteSliceArrayResult arbCoreLogsCursorGetLogs(CArbCore* arb_core_ptr);
int arbCoreLogsCursorSetNextIndex(CArbCore* arb_core_ptr, const void* count);
int arbCoreLogsCursorCheckError(CArbCore* arb_core_ptr);
char* arbCoreLogsCursorClearError(CArbCore* arb_core_ptr);

CExecutionCursor* arbCoreGetExecutionCursor(CArbCore* arb_core_ptr,
                                            const void* total_gas_used_ptr,
                                            CValueCache* cache_ptr);

#ifdef __cplusplus
}
#endif

#endif /* carbcore_h */
