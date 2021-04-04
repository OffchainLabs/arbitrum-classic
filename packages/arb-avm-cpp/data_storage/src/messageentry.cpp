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

#include <data_storage/messageentry.hpp>

#include "value/referencecount.hpp"
#include "value/utils.hpp"

std::vector<unsigned char> serializeSequencerBatchItem(
    const SequencerBatchItem& item) {
    std::vector<unsigned char> bytes;

    marshal_uint256_t(item.accumulator, bytes);
    marshal_uint256_t(item.total_delayed_count, bytes);
    if (item.sequencer_message) {
        bytes.insert(bytes.end(), item.sequencer_message->begin(),
                     item.sequencer_message->end());
    }

    return bytes;
}

std::vector<unsigned char> serializeDelayedMessage(const DelayedMessage& item) {
    std::vector<unsigned char> bytes;

    marshal_uint256_t(item.delayed_accumulator, bytes);
    bytes.insert(bytes.end(), item.message.begin(), item.message.end());

    return bytes;
}
