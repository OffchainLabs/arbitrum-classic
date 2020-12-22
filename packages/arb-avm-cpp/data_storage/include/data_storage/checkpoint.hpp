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

#ifndef data_storage_value_checkpoint_hpp
#define data_storage_value_checkpoint_hpp

#include <avm/machinestate/status.hpp>
#include <avm_values/bigint.hpp>
#include <avm_values/codepointstub.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresultfwd.hpp>

struct Checkpoint {
   public:
    uint64_t step_count{};
    uint256_t machine_hash;
    uint64_t messages_read_count{};
    uint256_t inbox_accumulator_hash;
    uint256_t block_hash;
    uint64_t block_height{};
    uint64_t logs_output{};
    uint64_t messages_output{};
    uint256_t arb_gas_used;

    Checkpoint() = default;
    Checkpoint(uint64_t step_count,
               uint256_t machine_hash,
               uint64_t messages_read_count,
               uint256_t inbox_accumulator_hash,
               uint256_t block_hash,
               uint256_t block_height,
               uint64_t logs_output,
               uint64_t messages_output,
               uint256_t arb_gas_used)
        : step_count(step_count),
          machine_hash(machine_hash),
          messages_read_count(messages_read_count),
          inbox_accumulator_hash(inbox_accumulator_hash),
          block_hash(block_hash),
          block_height(block_height),
          logs_output(logs_output),
          messages_output(messages_output),
          arb_gas_used(arb_gas_used) {}
    static DbResult<Checkpoint> getCheckpoint(Transaction& transaction,
                                              const uint64_t& message_number);
    static DbResult<Checkpoint> getKeyCheckpoint(Transaction& transaction,
                                                 rocksdb::Slice key_slice);
    static rocksdb::Status putCheckpoint(Transaction& transaction,
                                         const Checkpoint& checkpoint);
    static rocksdb::Status deleteCheckpoint(Transaction& transaction,
                                            const uint64_t& message_count);
    static DbResult<Checkpoint> atMessageOrPrevious(Transaction& transaction,
                                                    uint64_t message_number);
    static bool isEmpty(Transaction& transaction);
    static uint64_t maxCheckpointMessageNumber(Transaction& transaction);
};

#endif /* data_storage_value_checkpoint_hpp */
