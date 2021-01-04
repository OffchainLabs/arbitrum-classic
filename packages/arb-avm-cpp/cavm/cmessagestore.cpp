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

/*
Uint64Result putMessages(CMessageStore* storage_ptr,
                const uint64_t first_message_id,
                const uint64_t block_id,
                void* inbox_messages,
                void* inbox_hashes,
                const uint64_t message_count,
                const uint
                         ) {
    auto message_store = static_cast<MessageStore*>(storage_ptr);
    auto messages = getInboxMessages(inbox_messages, message_count);

    try {
        return {
            message_store->addMessages(first_message_id, block_id, messages),
true}; } catch (const std::exception& e) { return {0, false};
    }
}
*/
