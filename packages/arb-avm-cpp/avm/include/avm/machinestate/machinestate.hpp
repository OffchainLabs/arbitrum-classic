/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

#ifndef machinestate_hpp
#define machinestate_hpp

#include <avm/machinestate/blockreason.hpp>
#include <avm/machinestate/datastack.hpp>
#include <avm/machinestate/status.hpp>
#include <avm_values/value.hpp>
#include <avm_values/vmValueParser.hpp>

#include <memory>
#include <vector>

struct AssertionContext {
    std::vector<Tuple> inbox_messages;
    nonstd::optional<uint256_t> next_block_height;
    size_t inbox_messages_consumed{0};
    uint256_t numSteps{0};
    uint256_t numGas{0};
    bool blockingSideload{false};
    nonstd::optional<value> fake_inbox_peek_value;
    std::vector<std::vector<uint8_t>> sends;
    std::vector<value> logs;
    std::vector<value> debug_prints;

    AssertionContext() = default;

    AssertionContext(std::vector<Tuple> inbox_messages,
                     const nonstd::optional<uint256_t>& min_next_block_height,
                     uint256_t messages_to_skip);

    // popInbox assumes that the number of messages already consumed is less
    // than the number of messages in the inbox
    Tuple popInbox() {
        return std::move(inbox_messages[inbox_messages_consumed++]);
    }

    bool inboxEmpty() const {
        return inbox_messages_consumed == inbox_messages.size();
    }
};

struct MachineState {
    std::shared_ptr<Code> code;
    mutable nonstd::optional<CodeSegmentSnapshot> loaded_segment;
    value registerVal;
    value static_val;
    Datastack stack;
    Datastack auxstack;
    uint256_t arb_gas_remaining;
    Status state = Status::Extensive;
    CodePointRef pc;
    CodePointStub errpc;
    Tuple staged_message;
    AssertionContext context;

    static MachineState loadFromFile(const std::string& executable_filename);

    MachineState();

    MachineState(std::shared_ptr<Code> code_, value static_val);

    MachineState(std::shared_ptr<Code> code_,
                 value register_val_,
                 value static_val,
                 Datastack stack_,
                 Datastack auxstack_,
                 uint256_t arb_gas_remaining_,
                 Status state_,
                 CodePointRef pc_,
                 CodePointStub errpc_,
                 Tuple staged_message_);

    std::vector<unsigned char> marshalBufferProof();
    uint256_t getMachineSize() const;
    std::vector<unsigned char> marshalForProof() const;
    std::vector<unsigned char> marshalState() const;
    BlockReason runOp(OpCode opcode);
    BlockReason runOne();
    uint256_t hash() const;
    BlockReason isBlocked(bool newMessages) const;

    const CodePoint& loadCurrentInstruction() const;
    uint256_t nextGasCost() const;
};

#endif /* machinestate_hpp */
