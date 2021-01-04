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

#ifndef data_storage_messageentry_hpp
#define data_storage_messageentry_hpp

#include <avm/machine.hpp>
#include <avm/machinestate/status.hpp>
#include <avm_values/bigint.hpp>
#include <avm_values/codepointstub.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresultfwd.hpp>
#include <data_storage/value/machine.hpp>
#include <utility>

struct MessageEntry {
    // sequence_number not serialized/deserialized because it is part of index
    uint256_t sequence_number{};

    uint256_t message_hash;
    uint256_t inbox_hash;
    uint64_t block_height;
    bool last_message_in_block{};

    MessageEntry() = delete;
    MessageEntry(uint256_t sequence_number,
                 uint256_t message_hash,
                 uint256_t inbox_hash,
                 uint64_t block_height,
                 bool last_message_in_block)
        : sequence_number(sequence_number),
          message_hash(message_hash),
          inbox_hash(inbox_hash),
          block_height(block_height),
          last_message_in_block(last_message_in_block) {}
};

MessageEntry extractMessageEntry(rocksdb::Slice key, rocksdb::Slice value);

MessageEntry deserializeMessageEntry(
    uint256_t sequence_number,
    std::vector<const unsigned char> entry_vector);
std::vector<unsigned char> serializeMessageEntry(
    const MessageEntry& state_data);

bool operator==(const MessageEntry& lhs, const MessageEntry& rhs);
bool operator!=(const MessageEntry& lhs, const MessageEntry& rhs);

#endif /* data_storage_messageentry_hpp */
