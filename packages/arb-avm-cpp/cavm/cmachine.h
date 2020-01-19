/*
 * Copyright 2019, Offchain Labs, Inc.
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

#ifndef Machine_h
#define Machine_h

#include <stdint.h>
#include "ccheckpointstorage.h"

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
};

typedef enum {
    STATUS_EXTENSIVE = 0,
    STATUS_ERROR_STOP = 1,
    STATUS_HALT = 2,
} CStatus;

typedef struct {
    enum CBlockType blockType;
    ByteSlice val;
} CBlockReason;

typedef struct {
    unsigned char* outMessageData;
    int outMessageLength;
    int outMessageCount;
    unsigned char* logData;
    int logLength;
    int logCount;
    uint64_t numSteps;
    uint64_t numGas;
    int didInboxInsn;
} RawAssertion;

CMachine* machineCreate(const char* filename);
void machineDestroy(CMachine* m);

// Ret must have 32 bytes of storage allocated for returned hash
void machineHash(CMachine* m, void* ret);
CMachine* machineClone(CMachine* m);

// Ret must have 32 bytes of storage allocated for returned hash
CStatus machineCurrentStatus(CMachine* m);
CBlockReason machineLastBlockReason(CMachine* m);

RawAssertion machineExecuteAssertion(CMachine* m,
                                     uint64_t maxSteps,
                                     void* timeboundStart,
                                     void* timeboundEnd,
                                     void* inbox);

ByteSlice machineMarshallForProof(CMachine* m);

void machinePrint(CMachine* m);

int checkpointMachine(CMachine* m, CCheckpointStorage* storage);
int restoreMachine(CMachine* m,
                   CCheckpointStorage* storage,
                   const void* machine_hash);

#ifdef __cplusplus
}
#endif

#endif /* Machine_h */
