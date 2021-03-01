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

void Checkpoint::resetCheckpoint() {
    total_steps = 0;
    arb_gas_used = 0;
    total_messages_read = 0;
    inbox_hash = 0;
    next_sideload_block_number = 0;
    block_height = 0;
    send_count = 0;
    log_count = 0;
    machine_state_keys = MachineStateKeys{};
}

// applyAssertion does not update processed_message_accumulator_hash so it will
// have to be updated by caller.
void Checkpoint::applyAssertion(const Assertion& assertion) {
    total_steps += assertion.stepCount;
    arb_gas_used += assertion.gasCount;
    std::cout << "Apply " << total_messages_read << " "
              << assertion.inbox_messages_consumed << std::endl;
    total_messages_read += assertion.inbox_messages_consumed;
    send_count += assertion.sends.size();
    log_count += assertion.logs.size();
    if (assertion.sideloadBlockNumber) {
        next_sideload_block_number = *assertion.sideloadBlockNumber + 1;
    }
}

Checkpoint extractCheckpoint(const std::vector<unsigned char>& stored_state) {
    auto current_iter = stored_state.begin();

    auto arb_gas_used = extractUint256(current_iter);
    auto total_steps = extractUint256(current_iter);
    auto total_messages_read = extractUint256(current_iter);
    auto processed_message_accumulator_hash = extractUint256(current_iter);
    auto next_sideload_block_number = extractUint256(current_iter);
    auto block_height = extractUint64(current_iter);
    auto send_count = extractUint64(current_iter);
    auto log_count = extractUint64(current_iter);

    auto machineStateKeys = extractMachineStateKeys(current_iter);

    return Checkpoint{total_steps,
                      arb_gas_used,
                      total_messages_read,
                      processed_message_accumulator_hash,
                      next_sideload_block_number,
                      block_height,
                      send_count,
                      log_count,
                      machineStateKeys};
}

std::vector<unsigned char> serializeCheckpoint(const Checkpoint& state_data) {
    std::vector<unsigned char> state_data_vector;

    marshal_uint256_t(state_data.arb_gas_used, state_data_vector);
    marshal_uint256_t(state_data.total_steps, state_data_vector);
    marshal_uint256_t(state_data.total_messages_read, state_data_vector);
    marshal_uint256_t(state_data.inbox_hash, state_data_vector);
    marshal_uint256_t(state_data.next_sideload_block_number, state_data_vector);
    marshal_uint64_t(state_data.block_height, state_data_vector);
    marshal_uint64_t(state_data.send_count, state_data_vector);
    marshal_uint64_t(state_data.log_count, state_data_vector);

    serializeMachineStateKeys(state_data.machine_state_keys, state_data_vector);

    return state_data_vector;
}
