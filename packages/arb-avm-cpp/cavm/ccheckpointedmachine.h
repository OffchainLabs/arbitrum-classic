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

#ifndef ccheckpointedmachine_h
#define ccheckpointedmachine_h

#include <stdint.h>
#include "ccheckpointstorage.h"
#include "cmachine.h"

#ifdef __cplusplus
extern "C" {
#endif

CCheckpointedMachine* checkpointedMachineCreate(const char* executable_filename,
                                                const char* db_path);
int initializeCheckpointedMachine(CCheckpointedMachine* cm,
                                  const char* executable_path);
int checkpointedMachineInitialized(CCheckpointedMachine* cm);
int closeCheckpointedMachine(CCheckpointedMachine* cm);
void checkpointedMachineDestroy(CCheckpointedMachine* cm);

RawAssertion checkpointedExecuteAssertion(CCheckpointedMachine* cm,
                                          uint64_t maxSteps,
                                          void* inbox_messages,
                                          uint64_t message_count,
                                          uint64_t wallLimit);

RawAssertion checkpointedExecuteCallServerAssertion(CCheckpointedMachine* cm,
                                                    uint64_t maxSteps,
                                                    void* inbox_messages,
                                                    uint64_t message_count,
                                                    void* fake_inbox_peek_value,
                                                    uint64_t wallLimit);

RawAssertion checkpointedExecuteSideloadedAssertion(CCheckpointedMachine* cm,
                                                    uint64_t maxSteps,
                                                    void* inbox_messages,
                                                    uint64_t message_count,
                                                    void* sideload,
                                                    uint64_t wallLimit);

#ifdef __cplusplus
}
#endif

#endif /* ccheckpointedmachine_h */
