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

#include <avm_values/bigint.hpp>
#include <avm_values/tuple.hpp>
#include <utility>

using Address = std::array<unsigned char, 20>;

struct InboxMessage {
   public:
    // arb_gas_used not serialized/deserialized because it is part of index
    uint8_t kind{};
    Address sender{};
    uint256_t block_number;
    uint256_t timestamp;
    uint256_t inbox_sequence_number;
    uint256_t gas_price_l1;
    std::vector<unsigned char> data;

    InboxMessage() = default;
    InboxMessage(uint8_t kind,
                 const Address& sender,
                 uint256_t block_number,
                 uint256_t timestamp,
                 uint256_t inbox_sequence_number,
                 uint256_t gas_price_l1,
                 std::vector<unsigned char> data)
        : kind(kind),
          sender(sender),
          block_number(block_number),
          timestamp(timestamp),
          inbox_sequence_number(inbox_sequence_number),
          gas_price_l1(gas_price_l1),
          data(std::move(data)) {}

    static InboxMessage fromTuple(const Tuple& tup);

    [[nodiscard]] uint256_t hash() const;

    [[nodiscard]] Tuple toTuple() const;
    [[nodiscard]] std::vector<unsigned char> serialize() const;
    [[nodiscard]] std::vector<unsigned char> serializeForProof() const;
    void serializeHeader(std::vector<unsigned char>& state_data_vector) const;
    void serializeImpl(std::vector<unsigned char>& state_data_vector) const;
};

struct MachineMessage {
    InboxMessage message;
    uint256_t accumulator;

    MachineMessage() = default;
    MachineMessage(InboxMessage message_, uint256_t accumulator_)
        : message(message_), accumulator(accumulator_) {}

    void serializeImpl(std::vector<unsigned char>& state_data_vector) const;
};

uint256_t hash_raw_message(const std::vector<unsigned char>& stored_state);
uint256_t hash_inbox(const uint256_t& previous_inbox_acc,
                     const std::vector<unsigned char>& stored_state);
InboxMessage extractInboxMessage(
    const std::vector<unsigned char>& stored_state);
InboxMessage extractInboxMessageImpl(
    std::vector<unsigned char>::const_iterator current_iter,
    const std::vector<unsigned char>::const_iterator end);
MachineMessage extractMachineMessageImpl(
    std::vector<unsigned char>::const_iterator current_iter,
    const std::vector<unsigned char>::const_iterator end);
// An efficient version of extractInboxMessage that ignores everything except
// the block number
template <typename Iterator>
uint256_t extractInboxMessageBlockNumber(Iterator& iter) {
    iter++;      // skip kind
    iter += 20;  // skip sender
    return extractUint256(iter);
}

#endif /* data_storage_inboxmessage_hpp */
