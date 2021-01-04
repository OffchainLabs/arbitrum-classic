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

#include <data_storage/checkpoint.hpp>

#include "value/referencecount.hpp"
#include "value/utils.hpp"

#include <data_storage/value/machine.hpp>

#include <boost/endian/conversion.hpp>

Checkpoint extractCheckpoint(const std::vector<unsigned char>& stored_state) {
    auto current_iter = stored_state.begin();

    auto step_count = extractUint64(current_iter);
    auto messages_read_count = extractUint64(current_iter);
    auto inbox_accumulator_hash = extractUint256(current_iter);
    auto block_hash = extractUint256(current_iter);
    auto block_height = extractUint64(current_iter);
    auto logs_output = extractUint64(current_iter);
    auto messages_output = extractUint64(current_iter);
    auto arb_gas_used = extractUint256(current_iter);

    auto machineStateKeys = extractMachineStateKeys(current_iter);

    return Checkpoint{
        step_count,      messages_read_count, inbox_accumulator_hash,
        block_hash,      block_height,        logs_output,
        messages_output, arb_gas_used,        machineStateKeys};
}

std::vector<unsigned char> serializeCheckpoint(const Checkpoint& state_data) {
    std::vector<unsigned char> state_data_vector;

    marshal_uint64_t(state_data.step_count, state_data_vector);
    marshal_uint64_t(state_data.messages_read_count, state_data_vector);
    marshal_uint256_t(state_data.inbox_accumulator_hash, state_data_vector);
    marshal_uint256_t(state_data.block_hash, state_data_vector);
    marshal_uint64_t(state_data.block_height, state_data_vector);
    marshal_uint64_t(state_data.logs_output, state_data_vector);
    marshal_uint64_t(state_data.messages_output, state_data_vector);
    marshal_uint256_t(state_data.arb_gas_used, state_data_vector);

    serializeMachineStateKeys(state_data.machine_state_keys, state_data_vector);

    return state_data_vector;
}
