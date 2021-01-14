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

#include "utils.hpp"

#include <data_storage/messagestore.hpp>

void deleteMessageStore(CMessageStore* m) {
    delete static_cast<MessageStore*>(m);
}

int putMessages(CMessageStore* storage_ptr,
                const uint64_t first_message_sequence_number,
                const uint64_t block_height,
                void* inbox_messages,
                void* inbox_hashes_ptr,
                void* previous_inbox_hash_ptr,
                const uint) {
    auto message_store = static_cast<MessageStore*>(storage_ptr);
    auto messages = getInboxMessages(inbox_messages);
    auto inbox_hashes = receiveUint256Vector(inbox_hashes_ptr, messages.size());
    auto previous_inbox_hash = receiveUint256(previous_inbox_hash_ptr);

    try {
        auto status = message_store->addMessages(
            first_message_sequence_number, block_height, messages, inbox_hashes,
            previous_inbox_hash);
        if (status) {
            return status->ok();
        }

        // Caller needs to send messages from earlier block
        return false;
    } catch (const std::exception& e) {
        // TODO: Return error message
        return false;
    }
}
