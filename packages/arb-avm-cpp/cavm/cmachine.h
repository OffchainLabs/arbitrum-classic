/*
 * Copyright 2019-2021, Offchain Labs, Inc.
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

#ifndef cmachine_h
#define cmachine_h

#include <stdint.h>
#include "carbstorage.h"

#ifdef __cplusplus
extern "C" {
#endif

enum CBlockType {
    BLOCK_TYPE_NOT_BLOCKED = 0,
    BLOCK_TYPE_HALT = 1,
    BLOCK_TYPE_ERROR = 2,
    BLOCK_TYPE_BREAKPOINT = 3,
    BLOCK_TYPE_INBOX = 4,
    BLOCK_TYPE_SEND = 5,
    BLOCK_TYPE_SIDELOAD = 6
};

typedef enum {
    STATUS_EXTENSIVE = 0,
    STATUS_ERROR_STOP = 1,
    STATUS_HALT = 2,
    // STATE_UNKNOWN is used if the underlying MachineState state isn't a valid
    // value
    STATE_UNKNOWN = 3
} CStatus;

typedef struct {
    enum CBlockType blockType;
    ByteSlice val;
} CBlockReason;

typedef struct {
    ByteSlice standard_proof;
    ByteSlice buffer_proof;
} COneStepProof;

CMachine* machineCreate(const char* filename);
void machineDestroy(CMachine* m);
void machineAbort(CMachine* m);

// Ret must have 32 bytes of storage allocated for returned hash
int machineHash(CMachine* m, void* ret);
CMachine* machineClone(CMachine* m);

// Ret must have 32 bytes of storage allocated for returned hash
CStatus machineCurrentStatus(CMachine* m);
CBlockReason machineIsBlocked(CMachine* m, int newMessages);

RawAssertionResult executeAssertion(CMachine* m,
                                    const CMachineExecutionConfig* c);

COneStepProof machineMarshallForProof(CMachine* m);

ByteSlice machineMarshallState(CMachine* m);

char* machineInfo(CMachine* m);

void machineCodePointHash(CMachine* m, void*);

CMachineExecutionConfig* machineExecutionConfigCreate();
void machineExecutionConfigDestroy(CMachineExecutionConfig* m);
void* machineExecutionConfigClone(CMachineExecutionConfig* c);
void machineExecutionConfigSetMaxGas(CMachineExecutionConfig* c,
                                     uint64_t max_gas,
                                     int go_over_gas);
void machineExecutionConfigSetInboxMessages(CMachineExecutionConfig* c,
                                            ByteSliceArray bytes);
void machineExecutionConfigSetSideloads(CMachineExecutionConfig* c,
                                        ByteSliceArray bytes);
void machineExecutionConfigSetStopOnSideload(CMachineExecutionConfig* c,
                                             int stop_on_sideload);
void machineExecutionConfigSetStopOnBreakpoint(CMachineExecutionConfig* c,
                                               int stop_on_breakpoint);

#ifdef __cplusplus
}
#endif

#endif /* cmachine_h */
