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

// applyAssertion does not update processed_message_accumulator_hash so it will
// have to be updated by caller.
void Checkpoint::applyAssertion(const uint256_t& first_message_sequence_number,
                                const Assertion& assertion) {
    arb_gas_used += assertion.gasCount;
    message_sequence_number_processed =
        first_message_sequence_number + assertion.inbox_messages_consumed - 1;
    send_count += assertion.sends.size();
    log_count += assertion.logs.size();
}

Checkpoint extractCheckpoint(const uint256_t arb_gas_used,
                             const std::vector<unsigned char>& stored_state) {
    auto current_iter = stored_state.begin();

    auto message_sequence_number_processed = extractUint256(current_iter);
    auto processed_message_accumulator_hash = extractUint256(current_iter);
    auto block_height = extractUint64(current_iter);
    auto send_count = extractUint64(current_iter);
    auto log_count = extractUint64(current_iter);

    auto machineStateKeys = extractMachineStateKeys(current_iter);

    return Checkpoint{arb_gas_used,
                      message_sequence_number_processed,
                      processed_message_accumulator_hash,
                      block_height,
                      send_count,
                      log_count,
                      machineStateKeys};
}

std::vector<unsigned char> serializeCheckpoint(const Checkpoint& state_data) {
    std::vector<unsigned char> state_data_vector;

    marshal_uint256_t(state_data.message_sequence_number_processed,
                      state_data_vector);
    marshal_uint256_t(state_data.inbox_hash, state_data_vector);
    marshal_uint64_t(state_data.block_height, state_data_vector);
    marshal_uint64_t(state_data.send_count, state_data_vector);
    marshal_uint64_t(state_data.log_count, state_data_vector);

    serializeMachineStateKeys(state_data.machine_state_keys, state_data_vector);

    return state_data_vector;
}
