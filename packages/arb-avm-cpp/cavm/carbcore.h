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

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

void deleteArbCore(CArbCore* m);

int deliverMessages(CArbCore* storage_ptr,
                    const uint64_t first_message_sequence_number,
                    const uint64_t block_height,
                    void* inbox_messages,
                    void* inbox_hashes_ptr,
                    void* previous_inbox_hash_ptr);

ByteSliceResult arbCoreLastSendInserted(CArbCore* arb_core_ptr);
ByteSliceResult arbCoreLastLogInserted(CArbCore* arb_core_ptr);
ByteSliceResult arbCoreInboxMessagesRead(CArbCore* arb_core_ptr);

ByteSliceArrayResult arbCoreGetSends(CArbCore* arb_core_ptr,
                                     const void* start_index_ptr,
                                     const void* count_ptr);

ByteSliceArrayResult arbCoreGetMessages(CArbCore* arb_core_ptr,
                                        const void* start_index_ptr,
                                        const void* count_ptr);

Uint256Result arbCoreGetInboxDelta(CArbCore* arb_core_ptr,
                                   const void* start_index_ptr,
                                   const void* count_ptr);

Uint256Result arbCoreGetInboxAcc(CExecutionCursor* arb_core_ptr,
                                 const void* index);
Uint256Result arbCoreGetSendAcc(CExecutionCursor* arb_core_ptr,
                                const void* start_acc_hash,
                                const void* start_index_ptr,
                                const void* count_ptr);
Uint256Result arbCoreGetLogAcc(CExecutionCursor* arb_core_ptr,
                               const void* start_acc_hash,
                               const void* start_index_ptr,
                               const void* count_ptr);

CExecutionCursor arbCoreGetCursor(CArbCore* arb_core_ptr,
                                  const void* total_gas_used_ptr);

#ifdef __cplusplus
}
#endif

#endif /* carbcore_h */
