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

MessageEntry extractMessageEntry(uint256_t sequence_number,
                                 const rocksdb::Slice value) {
    // Extract message entry
    auto entry_vector =
        std::vector<unsigned char>{value.data(), value.data() + value.size()};

    return deserializeMessageEntry(sequence_number, entry_vector);
}

Tuple messageDataToTuple(const std::vector<unsigned char>& data) {
    auto ptr = data.data();

    if (data.size() < sizeof(char) + sizeof(uint256_t) * 5) {
        return {};
    }

    uint256_t kind = ptr[0];
    ptr++;
    auto block_number = intx::be::unsafe::load<uint256_t>(ptr);
    ptr += sizeof(uint256_t);

    auto timestamp = intx::be::unsafe::load<uint256_t>(ptr);
    ptr += sizeof(uint256_t);

    auto sender = intx::be::unsafe::load<uint256_t>(ptr);
    ptr += sizeof(uint256_t);

    auto sequence_num = intx::be::unsafe::load<uint256_t>(ptr);
    ptr += sizeof(uint256_t);

    auto buf_size = intx::be::unsafe::load<uint256_t>(ptr);
    ptr += sizeof(uint256_t);

    auto remaining_length = ptr - data.data() + buf_size;
    if (remaining_length > data.size()) {
        buf_size = remaining_length;
    }

    Buffer buf;
    buf = buf.set_many(0, std::vector<uint8_t>(
                              ptr, ptr + intx::narrow_cast<size_t>(buf_size)));

    Tuple message(kind, block_number, timestamp, sender, sequence_num, buf_size,
                  std::move(buf));

    return message;
}

MessageEntry deserializeMessageEntry(
    const uint256_t sequence_number,
    const std::vector<unsigned char>& entry_vector) {
    auto current_iter = entry_vector.begin();

    auto inbox_hash = extractUint256(current_iter);
    auto block_height = extractUint64(current_iter);
    auto last_message_in_block = current_iter[0] == 1;
    current_iter++;
    uint64_t remaining_size = entry_vector.end() - current_iter;
    auto message =
        std::vector<unsigned char>(current_iter, current_iter + remaining_size);

    return MessageEntry{sequence_number, inbox_hash, block_height,
                        last_message_in_block, message};
}

std::vector<unsigned char> serializeMessageEntry(
    const MessageEntry& state_data) {
    std::vector<unsigned char> state_data_vector;

    marshal_uint256_t(state_data.inbox_hash, state_data_vector);
    marshal_uint64_t(state_data.block_height, state_data_vector);
    state_data_vector.push_back(state_data.last_message_in_block ? 1 : 0);
    state_data_vector.insert(state_data_vector.end(), state_data.data.data(),
                             state_data.data.data() + state_data.data.size());

    return state_data_vector;
}

bool operator==(const MessageEntry& lhs, const MessageEntry& rhs) {
    return lhs.sequence_number == rhs.sequence_number &&
           lhs.inbox_hash == rhs.inbox_hash &&
           lhs.block_height == rhs.block_height &&
           lhs.last_message_in_block == rhs.last_message_in_block &&
           lhs.data.size() == rhs.data.size() &&
           memcmp(lhs.data.data(), rhs.data.data(), lhs.data.size()) == 0;
}

bool operator!=(const MessageEntry& lhs, const MessageEntry& rhs) {
    return !(lhs == rhs);
}
