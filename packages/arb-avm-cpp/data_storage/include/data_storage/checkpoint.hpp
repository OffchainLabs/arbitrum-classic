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

#ifndef data_storage_checkpoint_hpp
#define data_storage_checkpoint_hpp

#include <avm/machine.hpp>
#include <avm/machinestate/status.hpp>
#include <avm_values/bigint.hpp>
#include <avm_values/codepointstub.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresultfwd.hpp>
#include <data_storage/value/machine.hpp>
#include <utility>

struct Checkpoint {
   public:
    // arb_gas_used not serialized/deserialized because it is part of index
    uint256_t arb_gas_used;

    uint256_t message_sequence_number_processed;
    uint256_t processed_message_accumulator_hash;
    uint64_t reorg_index{};

    uint64_t block_height{};
    uint64_t send_count{};
    uint64_t log_count{};
    MachineStateKeys machine_state_keys{};

    Checkpoint() = default;
    Checkpoint(uint256_t arb_gas_used,
               uint256_t message_sequence_number_processed,
               uint256_t processed_message_accumulator_hash,
               uint64_t reorg_index,
               uint64_t block_height,
               uint64_t send_count,
               uint64_t log_count,
               MachineStateKeys machine_state_keys)
        : arb_gas_used(arb_gas_used),
          message_sequence_number_processed(message_sequence_number_processed),
          processed_message_accumulator_hash(
              processed_message_accumulator_hash),
          reorg_index(reorg_index),
          block_height(block_height),
          send_count(send_count),
          log_count(log_count),
          machine_state_keys(machine_state_keys) {}
};

Checkpoint extractCheckpoint(uint256_t arb_gas_used,
                             const std::vector<unsigned char>& stored_state);

std::vector<unsigned char> serializeCheckpoint(const Checkpoint& state_data);

#endif /* data_storage_checkpoint_hpp */
