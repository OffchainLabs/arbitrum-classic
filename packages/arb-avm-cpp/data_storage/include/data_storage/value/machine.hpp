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

#include <avm/machinestate/status.hpp>

#include <rocksdb/status.h>
#include <avm_values/bigint.hpp>
#include <avm_values/codepoint.hpp>
#include <utility>

class Transaction;

template <typename T>
struct DbResult;

struct SaveResults;
struct DeleteResults;

class Machine;

struct MachineStateKeys {
    uint256_t static_hash;
    uint256_t register_hash;
    uint256_t datastack_hash;
    uint256_t auxstack_hash;
    uint256_t arb_gas_remaining;
    CodePointRef pc{0, 0};
    CodePointStub err_pc{{0, 0}, getErrCodePoint()};
    uint256_t fully_processed_messages;
    uint256_t fully_processed_inbox_accumulator;
    staged_variant staged_message;
    Status status{};

    MachineStateKeys() = default;
    MachineStateKeys(uint256_t static_hash_,
                     uint256_t register_hash_,
                     uint256_t datastack_hash_,
                     uint256_t auxstack_hash_,
                     uint256_t arb_gas_remaining_,
                     CodePointRef pc_,
                     CodePointStub err_pc_,
                     uint256_t fully_processed_messages_,
                     uint256_t fully_processed_inbox_accumulator_,
                     staged_variant staged_message_,
                     Status status_)
        : static_hash(static_hash_),
          register_hash(register_hash_),
          datastack_hash(datastack_hash_),
          auxstack_hash(auxstack_hash_),
          arb_gas_remaining(arb_gas_remaining_),
          pc(pc_),
          err_pc(err_pc_),
          fully_processed_messages(fully_processed_messages_),
          fully_processed_inbox_accumulator(fully_processed_inbox_accumulator_),
          staged_message(std::move(staged_message_)),
          status(status_) {}
};

DbResult<MachineStateKeys> getMachineStateKeys(const Transaction& transaction,
                                               uint256_t machineHash);
MachineStateKeys extractMachineStateKeys(
    std::vector<unsigned char>::const_iterator iter,
    const std::vector<unsigned char>::const_iterator end);
void serializeMachineStateKeys(const MachineStateKeys& state_data,
                               std::vector<unsigned char>& state_data_vector);
rocksdb::Status saveMachineState(Transaction& transaction,
                                 const Machine& machine,
                                 MachineStateKeys& machine_state_keys);
SaveResults saveMachine(Transaction& transaction, const Machine& machine);
void deleteMachineState(Transaction& transaction,
                        MachineStateKeys& parsed_state);
DeleteResults deleteMachine(Transaction& transaction, uint256_t machine_hash);

#endif /* checkpoint_machine_hpp */
