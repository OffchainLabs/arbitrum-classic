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

#include <data_storage/storageresult.hpp>

#include <avm/machinestate/machinestate.hpp>

#include <rocksdb/status.h>
#include <avm_values/bigint.hpp>
#include <avm_values/codepoint.hpp>
#include <data_storage/readwritetransaction.hpp>
#include <utility>

class Transaction;

struct SaveResults;
struct DeleteResults;

class Machine;

DbResult<std::variant<MachineStateKeys, MachineOutput>> getMachineStateKeys(
    const ReadTransaction& transaction,
    uint256_t machineHash);
MachineOutput extractMachineOutput(
    std::vector<unsigned char>::const_iterator& iter);
std::variant<MachineStateKeys, MachineOutput> extractMachineStateKeys(
    const std::vector<unsigned char>& data);
void serializeMachineOutput(const MachineOutput& output_data,
                            std::vector<unsigned char>& state_data_vector);
void serializeMachineStateKeys(const MachineStateKeys& state_data,
                               std::vector<unsigned char>& state_data_vector);
std::pair<rocksdb::Status, std::map<uint64_t, uint64_t>> saveMachineState(
    ReadWriteTransaction& transaction,
    const Machine& machine);
SaveResults saveTestMachine(ReadWriteTransaction& transaction,
                            Machine& machine);
void deleteMachineState(ReadWriteTransaction& transaction,
                        MachineStateKeys& parsed_state);
DeleteResults deleteMachine(ReadWriteTransaction& tx, uint256_t machine_hash);

#endif /* checkpoint_machine_hpp */
