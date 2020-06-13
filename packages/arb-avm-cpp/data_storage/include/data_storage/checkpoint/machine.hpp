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

#ifndef checkpoint_machine_hpp
#define checkpoint_machine_hpp

#include <avm_values/bigint.hpp>
#include <avm_values/codepoint.hpp>

class Transaction;

template <typename T>
struct DbResult;

struct SaveResults;
struct DeleteResults;

struct MachineStateKeys {
    uint256_t register_hash;
    uint256_t datastack_hash;
    uint256_t auxstack_hash;
    CodePointStub pc;
    CodePointStub err_pc;
    unsigned char status_char;
};

DbResult<MachineStateKeys> getMachineState(const Transaction& transaction,
                                           uint256_t machineHash);
SaveResults saveMachineState(Transaction& transaction,
                             const MachineStateKeys& state_data,
                             uint256_t machineHash);
DeleteResults deleteMachine(Transaction& transaction, uint256_t machine_hash);

#endif /* checkpoint_machine_hpp */
