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
#include <data_storage/value/utils.hpp>

#include "value/referencecount.hpp"

uint256_t SequencerBatchItem::computeAccumulator(uint256_t prev_acc,
                                                 uint256_t prev_delayed_count,
                                                 uint256_t delayed_acc) {
    std::vector<unsigned char> data;
    if (total_delayed_count > prev_delayed_count) {
        assert(!sequencer_message);
        std::string prefix = "Delayed messages:";
        data.insert(data.end(), prefix.begin(), prefix.end());
        marshal_uint256_t(prev_acc, data);
        marshal_uint256_t(last_sequence_number + 1 -
                              (total_delayed_count - prev_delayed_count),
                          data);
        marshal_uint256_t(prev_delayed_count, data);
        marshal_uint256_t(total_delayed_count, data);
        marshal_uint256_t(delayed_acc, data);
    } else {
        assert(sequencer_message);
        assert(total_delayed_count == prev_delayed_count);
        std::string prefix = "Sequencer message:";
        data.insert(data.end(), prefix.begin(), prefix.end());
        marshal_uint256_t(prev_acc, data);
        marshal_uint256_t(prev_delayed_count, data);
        marshal_uint256_t(extractInboxMessage(*sequencer_message).hash(), data);
    }
    return ::hash(data);
}

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
