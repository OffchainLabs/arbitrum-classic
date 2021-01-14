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

#ifndef data_storage_inboxmessage_hpp
#define data_storage_inboxmessage_hpp

#include <avm/machine.hpp>
#include <avm/machinestate/status.hpp>
#include <avm_values/bigint.hpp>
#include <avm_values/codepointstub.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresultfwd.hpp>
#include <data_storage/value/machine.hpp>
#include <utility>

struct InboxMessage {
   public:
    // arb_gas_used not serialized/deserialized because it is part of index
    uint8_t kind;
    uint256_t sender;
    uint256_t block_number;
    uint256_t timestamp;
    uint256_t inbox_sequence_number;
    std::vector<unsigned char> data;

    InboxMessage() = delete;
    InboxMessage(uint8_t kind,
                 uint256_t sender,
                 uint256_t block_number,
                 uint256_t timestamp,
                 uint256_t inbox_sequence_number,
                 std::vector<unsigned char> data)
        : kind(kind),
          sender(sender),
          block_number(block_number),
          timestamp(timestamp),
          inbox_sequence_number(inbox_sequence_number),
          data(std::move(data)) {}
};

InboxMessage extractInboxMessage(
    const std::vector<unsigned char>& stored_state);
std::vector<InboxMessage> extractInboxMessages(
    const std::vector<rocksdb::Slice> slices);

std::vector<unsigned char> serializeInboxMessage(
    const InboxMessage& state_data);

#endif /* data_storage_inboxmessage_hpp */
