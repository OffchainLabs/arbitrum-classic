/*
 * Copyright 2019-2021, Offchain Labs, Inc.
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
struct MachineState;

struct InboxState {
    uint256_t count;
    uint256_t accumulator;

    void addMessage(const MachineMessage& message) {
        accumulator = message.accumulator;
        count += 1;
    }

    bool operator==(const InboxState& other) const;
    bool operator!=(const InboxState& other) const;
};

template <typename T>
struct MachineEmission {
    T val;
    InboxState inbox;
};

struct AssertionContext {
    std::vector<MachineMessage> inbox_messages;

    std::vector<MachineEmission<std::vector<uint8_t>>> sends;
    std::vector<MachineEmission<value>> logs;
    std::vector<MachineEmission<value>> debug_prints;
    std::deque<InboxMessage> sideloads;
    bool stop_on_sideload{false};
    uint256_t max_gas;
    bool go_over_gas{false};
    bool first_instruction{true};

   private:
    size_t inbox_messages_consumed{0};

   public:
    AssertionContext() = default;

    explicit AssertionContext(MachineExecutionConfig config);

    // popInbox assumes that the number of messages already consumed is less
    // than the number of messages in the inbox
    MachineMessage popInbox() {
        return inbox_messages[inbox_messages_consumed++];
    }

    // peekInbox assumes that the number of messages already consumed is less
    // than the number of messages in the inbox
    [[nodiscard]] const MachineMessage& peekInbox() const {
        return inbox_messages[inbox_messages_consumed];
    }

    [[nodiscard]] bool inboxEmpty() const {
        return inbox_messages_consumed >= inbox_messages.size();
    }

    void resetForContinuedRun() {
        sends.clear();
        logs.clear();
        debug_prints.clear();
        first_instruction = true;
    }
};

struct OneStepProof {
    std::vector<unsigned char> standard_proof;
    std::vector<unsigned char> buffer_proof;
};

struct MachineOutput {
    InboxState fully_processed_inbox;
    uint256_t total_steps;
    uint256_t arb_gas_used;
    uint256_t send_acc;
    uint256_t log_acc;
    uint256_t send_count;
    uint256_t log_count;
    uint256_t l1_block_number;
    uint256_t l2_block_number;
    uint256_t last_inbox_timestamp;
    std::optional<uint256_t> last_sideload;

    bool operator==(const MachineOutput& other) const;
    bool operator!=(const MachineOutput& other) const;
};

struct MachineStateKeys {
    MachineOutput output;
    CodePointStub pc;
    uint256_t static_hash;
    uint256_t register_hash;
    uint256_t datastack_hash;
    uint256_t auxstack_hash;
    uint256_t arb_gas_remaining;
    Status state;
    CodePointStub err_pc;

    MachineStateKeys(MachineOutput output_,
                     CodePointStub pc_,
                     uint256_t static_hash_,
                     uint256_t register_hash_,
                     uint256_t datastack_hash_,
                     uint256_t auxstack_hash_,
                     uint256_t arb_gas_remaining_,
                     Status state_,
                     CodePointStub err_pc_)
        : output(output_),
          pc(pc_),
          static_hash(static_hash_),
          register_hash(register_hash_),
          datastack_hash(datastack_hash_),
          auxstack_hash(auxstack_hash_),
          arb_gas_remaining(arb_gas_remaining_),
          state(state_),
          err_pc(err_pc_) {}

    explicit MachineStateKeys(const MachineState& machine);

    [[nodiscard]] uint256_t getTotalMessagesRead() const;
    [[nodiscard]] uint256_t getInboxAcc() const;
    [[nodiscard]] uint256_t machineHash() const;
};

struct MachineState {
    MachineOutput output;

    CodePointRef pc{0, 0};
    std::shared_ptr<Code> code;
    mutable std::optional<CodeSegmentSnapshot> loaded_segment;
    value registerVal;
    value static_val;
    Datastack stack;
    Datastack auxstack;
    uint256_t arb_gas_remaining;
    Status state{Status::Extensive};
    CodePointStub errpc{{0, 0}, getErrCodePoint()};

    AssertionContext context;

    static MachineState loadFromFile(const std::string& executable_filename);

    MachineState();

    MachineState(std::shared_ptr<CoreCode> code_, value static_val);

    MachineState(MachineOutput output_,
                 CodePointRef pc_,
                 std::shared_ptr<Code> code_,
                 value register_val_,
                 value static_val,
                 Datastack stack_,
                 Datastack auxstack_,
                 uint256_t arb_gas_remaining_,
                 Status state_,
                 CodePointStub errpc_);

    uint256_t getMachineSize() const;
    OneStepProof marshalForProof() const;
    std::vector<unsigned char> marshalState() const;
    BlockReason runOp(OpCode opcode);
    BlockReason runOne();
    uint256_t hash() const { return MachineStateKeys(*this).machineHash(); }
    BlockReason isBlocked(bool newMessages) const;

    CodePoint loadCurrentInstruction() const;
    const Operation& loadCurrentOperation() const;
    uint256_t nextGasCost() const;

    uint256_t getTotalMessagesRead() const;

    void addProcessedMessage(const MachineMessage& message);
    void addProcessedSend(std::vector<uint8_t> data);
    void addProcessedLog(value log_val);

   private:
    void marshalBufferProof(OneStepProof& proof) const;
    uint256_t gasCost(const Operation& op) const;
};

#endif /* machinestate_hpp */
