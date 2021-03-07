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

#include <avm/inboxmessage.hpp>
#include <avm/machinestate/blockreason.hpp>
#include <avm/machinestate/datastack.hpp>
#include <avm/machinestate/status.hpp>

#include <avm_values/value.hpp>
#include <avm_values/vmValueParser.hpp>

#include <deque>
#include <memory>
#include <vector>

class MachineExecutionConfig;

struct AssertionContext {
    std::vector<InboxMessage> inbox_messages;
    std::optional<uint256_t> next_block_height;
    size_t inbox_messages_consumed{0};
    size_t messages_to_skip{0};
    uint256_t numSteps{0};
    uint256_t numGas{0};
    std::optional<value> fake_inbox_peek_value;
    std::vector<std::vector<uint8_t>> sends;
    std::vector<value> logs;
    std::vector<value> debug_prints;
    std::deque<InboxMessage> sideloads;
    bool stop_on_sideload{false};
    uint256_t max_gas;
    bool go_over_gas{false};

    AssertionContext() = default;

    explicit AssertionContext(MachineExecutionConfig config);

    // popInbox assumes that the number of messages already consumed is less
    // than the number of messages in the inbox
    InboxMessage popInbox() {
        return inbox_messages[inbox_messages_consumed++];
    }

    // peekInbox assumes that the number of messages already consumed is less
    // than the number of messages in the inbox
    [[nodiscard]] const InboxMessage& peekInbox() const {
        return inbox_messages[inbox_messages_consumed];
    }

    [[nodiscard]] bool inboxEmpty() const {
        return inbox_messages_consumed >= inbox_messages.size();
    }

    void resetForContinuedRun() {
        sends.clear();
        logs.clear();
        debug_prints.clear();
        max_gas -= numGas;
        numGas = 0;
        numSteps = 0;
        messages_to_skip = inbox_messages_consumed;
    }
};

struct OneStepProof {
    std::vector<unsigned char> standard_proof;
    std::vector<unsigned char> buffer_proof;
};

struct MachineState {
    std::shared_ptr<Code> code;
    mutable std::optional<CodeSegmentSnapshot> loaded_segment;
    value registerVal;
    value static_val;
    Datastack stack;
    Datastack auxstack;
    uint256_t arb_gas_remaining;
    Status state{Status::Extensive};
    CodePointRef pc{0, 0};
    CodePointStub errpc{{0, 0}, getErrCodePoint()};
    uint256_t fully_processed_messages;
    uint256_t fully_processed_inbox_accumulator;
    staged_variant staged_message{std::monostate()};
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
                 uint256_t fully_processed_messages_,
                 uint256_t fully_processed_inbox_accumulator_,
                 staged_variant staged_message_);

    uint256_t getMachineSize() const;
    OneStepProof marshalForProof() const;
    std::vector<unsigned char> marshalState() const;
    BlockReason runOp(OpCode opcode);
    BlockReason runOne();
    std::optional<uint256_t> hash() const;
    BlockReason isBlocked(bool newMessages) const;

    const CodePoint& loadCurrentInstruction() const;
    uint256_t nextGasCost() const;

    bool stagedMessageEmpty() const;
    bool stagedMessageUnresolved() const;
    std::optional<uint256_t> getUnresolvedStagedMessageBlockHeight() const;
    std::optional<Tuple> getStagedMessageTuple() const;
    uint256_t getMessagesConsumed() const;

   private:
    void marshalBufferProof(OneStepProof& proof) const;
};

#endif /* machinestate_hpp */
