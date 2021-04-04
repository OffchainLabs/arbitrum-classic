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

struct SequencerBatchItem {
    // last_sequence_number is the key in the DB; not serialized in the value
    uint256_t last_sequence_number;

    uint256_t accumulator;
    uint256_t total_delayed_count;
    std::optional<std::vector<unsigned char>> sequencer_message;
};

template <typename Iterator>
SequencerBatchItem deserializeSequencerBatchItem(uint256_t last_sequence_number,
                                                 Iterator& entry_vector) {
    auto accumulator = extractUint256(current_iter);
    auto total_delayed_count = extractUint256(current_iter);
    std::optional<std::vector<unsigned char>> sequencer_message;
    if (current_iter != entry_vector.end()) {
        sequencer_message = std::vector(current_iter, entry_vector.end());
    }
    return {last_sequence_number, accumulator, total_delayed_count,
            sequencer_message};
}

std::vector<unsigned char> serializeSequencerBatchItem(
    const SequencerBatchItem& item);

struct DelayedMessage {
    // delayed_sequence_number is the key in the DB; not serialized in the value
    uint256_t delayed_sequence_number;

    uint256_t delayed_accumulator;
    std::vector<unsigned char> message;
};

template <typename Iterator>
DelayedMessage deserializeDelayedMessage(uint256_t delayed_sequence_number,
                                         Iterator& current_iter) {
    auto delayed_accumulator = extractUint256(current_iter);
    std::vector<unsigned char> message(current_iter, entry_vector.end());
    return {delayed_sequence_number, delayed_accumulator, message};
}

std::vector<unsigned char> serializeDelayedMessage(const DelayedMessage& item);

#endif /* data_storage_messageentry_hpp */
