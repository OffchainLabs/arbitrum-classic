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
void* arbCoreMachineMessagesRead(CArbCore* arbcore_ptr);
int arbCoreMessagesStatus(CArbCore* arbcore_ptr);
char* arbCoreMessagesClearError(CArbCore* arbcore_ptr);

int arbCoreDeliverMessages(CArbCore* arbcore_ptr,
                           void* previous_message_count_ptr,
                           void* previous_inbox_acc_ptr,
                           ByteSliceArray sequencer_batch_items_slice,
                           ByteSliceArray delayed_messages_slice,
                           void* reorg_message_count_ptr);

Uint256Result arbCoreGetLogCount(CArbCore* arbcore_ptr);

ByteSliceArrayResult arbCoreGetLogs(CArbCore* arbcore_ptr,
                                    const void* start_index_ptr,
                                    const void* count_ptr);

Uint256Result arbCoreGetSendCount(CArbCore* arbcore_ptr);

ByteSliceArrayResult arbCoreGetSends(CArbCore* arbcore_ptr,
                                     const void* start_index_ptr,
                                     const void* count_ptr);

Uint256Result arbCoreGetMessageCount(CArbCore* arbcore_ptr);

Uint256Result arbCoreGetTotalDelayedMessagesSequenced(CArbCore* arbcore_ptr);

ByteSliceArrayResult arbCoreGetMessages(CArbCore* arbcore_ptr,
                                        const void* start_index_ptr,
                                        const void* count_ptr);

ByteSliceArrayResult arbCoreGetSequencerBatchItems(CArbCore* arbcore_ptr,
                                                   const void* start_index_ptr);

Uint256Result arbCoreGetSequencerBlockNumberAt(CArbCore* arbcore_ptr,
                                               const void* seq_num_ptr);

ByteSliceResult arbCoreGenInboxProof(CArbCore* arbcore_ptr,
                                     const void* seq_num_ptr,
                                     const void* batch_index_ptr,
                                     const void* batch_end_count_ptr);

int arbCoreGetInboxAcc(CArbCore* arbcore_ptr, const void* index_ptr, void* ret);
int arbCoreGetDelayedInboxAcc(CArbCore* arbcore_ptr,
                              const void* index_ptr,
                              void* ret);
int arbCoreGetInboxAccPair(CArbCore* arbcore_ptr,
                           const void* index1_ptr,
                           const void* index2_ptr,
                           void* ret1,
                           void* ret2);
int arbCoreCountMatchingBatchAccs(CArbCore* arbcore_ptr, ByteSlice data);
Uint256Result arbCoreGetDelayedMessagesToSequence(CArbCore* arbcore_ptr,
                                                  const void* max_block_number);

Uint256Result arbCoreLogsCursorGetPosition(CArbCore* arbcore_ptr,
                                           const void* index_ptr);
int arbCoreLogsCursorRequest(CArbCore* arbcore_ptr,
                             const void* cursor_index,
                             const void* count);
IndexedDoubleByteSliceArrayResult arbCoreLogsCursorGetLogs(
    CArbCore* arbcore_ptr,
    const void* index_ptr);
int arbCoreLogsCursorConfirmReceived(CArbCore* arbcore_ptr,
                                     const void* cursor_index);
int arbCoreLogsCursorCheckError(CArbCore* arbcore_ptr,
                                const void* cursor_index);
char* arbCoreLogsCursorClearError(CArbCore* arbcore_ptr,
                                  const void* cursor_index);

CExecutionCursor* arbCoreGetExecutionCursor(CArbCore* arbcore_ptr,
                                            const void* total_gas_used_ptr);
int arbCoreAdvanceExecutionCursor(CArbCore* arbcore_ptr,
                                  CExecutionCursor* execution_cursor_ptr,
                                  const void* max_gas_ptr,
                                  int go_over_gas);
CMachine* arbCoreGetLastMachine(CArbCore* arbcore_ptr);
Uint256Result arbCoreGetLastMachineTotalGas(CArbCore* arbcore_ptr);
CMachine* arbCoreTakeMachine(CArbCore* arbcore_ptr,
                             CExecutionCursor* execution_cursor_ptr);
CMachine* arbCoreGetMachineForSideload(CArbCore* arbcore_ptr,
                                       uint64_t block_number);

#ifdef __cplusplus
}
#endif

#endif /* carbcore_h */
