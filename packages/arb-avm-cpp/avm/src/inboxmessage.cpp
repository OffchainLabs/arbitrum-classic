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

#include <avm/inboxmessage.hpp>

//#include "value/utils.hpp"

#include <ethash/keccak.hpp>

uint256_t InboxMessage::hash(const uint256_t& previous_inbox_hash) const {
    std::vector<unsigned char> inbox_vector;

    inbox_vector.push_back(kind);
    inbox_vector.insert(inbox_vector.end(), sender.begin(), sender.end());
    marshal_uint256_t(block_number, inbox_vector);
    marshal_uint256_t(timestamp, inbox_vector);
    marshal_uint256_t(inbox_sequence_number, inbox_vector);
    auto data_hash = ::hash(data);
    marshal_uint256_t(data_hash, inbox_vector);

    auto message_hash = ::hash(inbox_vector);

    return ::hash(previous_inbox_hash, message_hash);
}

uint256_t hash_raw_message(const std::vector<unsigned char>& stored_state) {
    constexpr auto message_fixed_size = 1 + 20 + 32 * 3;

    // Calculate hash of variable length data
    std::vector<unsigned char> variable_data{
        stored_state.begin() + message_fixed_size, stored_state.end()};
    auto variable_hash = hash(variable_data);

    std::vector<unsigned char> fixed_data{
        stored_state.begin(), stored_state.begin() + message_fixed_size};
    marshal_uint256_t(variable_hash, fixed_data);

    return hash(fixed_data);
}

uint256_t hash_inbox(const uint256_t& previous_inbox_hash,
                     const std::vector<unsigned char>& stored_state) {
    return hash(previous_inbox_hash, hash_raw_message(stored_state));
    ;
}

namespace {
template <typename Iterator>
uint256_t extractUint256(Iterator& iter) {
    auto ptr = reinterpret_cast<const char*>(&*iter);
    auto int_val = deserializeUint256t(ptr);
    iter += 32;
    return int_val;
}
}  // namespace

InboxMessage extractInboxMessage(
    const std::vector<unsigned char>& stored_state) {
    auto current_iter = stored_state.begin();

    auto kind = *reinterpret_cast<const uint8_t*>(&*current_iter);
    current_iter++;
    Address sender;
    std::copy(current_iter, current_iter + sender.size(), sender.begin());
    current_iter += sender.size();
    auto block_number = extractUint256(current_iter);
    auto timestamp = extractUint256(current_iter);
    auto inbox_sequence_number = extractUint256(current_iter);

    std::vector<unsigned char> data;
    data.insert(data.end(), current_iter, stored_state.end());

    return InboxMessage{
        kind, sender, block_number, timestamp, inbox_sequence_number, data};
}

std::vector<unsigned char> InboxMessage::serialize() const {
    std::vector<unsigned char> state_data_vector;
    state_data_vector.push_back(kind);
    state_data_vector.insert(state_data_vector.end(), sender.begin(),
                             sender.end());
    marshal_uint256_t(block_number, state_data_vector);
    marshal_uint256_t(timestamp, state_data_vector);
    marshal_uint256_t(inbox_sequence_number, state_data_vector);
    state_data_vector.insert(state_data_vector.end(), data.begin(), data.end());
    return state_data_vector;
}

std::vector<unsigned char> InboxMessage::serializeForProof() const {
    std::vector<unsigned char> state_data_vector;
    state_data_vector.push_back(kind);
    state_data_vector.insert(state_data_vector.end(), sender.begin(),
                             sender.end());
    marshal_uint256_t(block_number, state_data_vector);
    marshal_uint256_t(timestamp, state_data_vector);
    marshal_uint256_t(inbox_sequence_number, state_data_vector);
    uint256_t proofLength = state_data_vector.size();
    marshal_uint256_t(proofLength, state_data_vector);
    state_data_vector.insert(state_data_vector.end(), data.begin(), data.end());
    return state_data_vector;
}

Tuple InboxMessage::toTuple() {
    uint8_t raw_sender[32];
    std::fill_n(&raw_sender[0], 12, 0);
    std::copy(sender.begin(), sender.end(), &raw_sender[12]);
    return {uint256_t{kind},
            block_number,
            timestamp,
            intx::be::load<uint256_t>(raw_sender),
            inbox_sequence_number,
            uint256_t{data.size()},
            Buffer{data}};
}

InboxMessage InboxMessage::fromTuple(const Tuple& tup) {
    if (tup.tuple_size() != 7) {
        throw std::runtime_error("wrong tup size");
    }
    auto kind =
        intx::narrow_cast<uint8_t>(tup.get_element_unsafe(0).get<uint256_t>());
    auto block_number = tup.get_element_unsafe(1).get<uint256_t>();
    auto timestamp = tup.get_element_unsafe(2).get<uint256_t>();
    auto sender_int = tup.get_element_unsafe(3).get<uint256_t>();
    auto inbox_sequence_number = tup.get_element_unsafe(4).get<uint256_t>();
    auto data_size =
        intx::narrow_cast<uint64_t>(tup.get_element_unsafe(5).get<uint256_t>());
    auto data_buf = tup.get_element_unsafe(6).get<Buffer>();

    uint8_t raw_sender[32];
    intx::be::store(raw_sender, sender_int);

    Address sender;
    std::copy(&raw_sender[12], &raw_sender[32], sender.begin());

    std::vector<unsigned char> data;
    data.reserve(data_size);
    for (uint64_t i = 0; i < data_size; i++) {
        data.push_back(data_buf.get(i));
    }
    return InboxMessage{
        kind, sender, block_number, timestamp, inbox_sequence_number, data};
}
