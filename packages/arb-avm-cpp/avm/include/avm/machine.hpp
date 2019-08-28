/*
 * Copyright 2019, Offchain Labs, Inc.
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

#ifndef machine_hpp
#define machine_hpp

#include <avm/datastack.hpp>
#include <avm/tokenTracker.hpp>
#include <avm/value.hpp>

#include <memory>
#include <vector>

typedef std::array<uint256_t, 2> TimeBounds;

enum class Status { Extensive, Halted, Error };

struct NotBlocked {};

struct HaltBlocked {};

struct ErrorBlocked {};

struct BreakpointBlocked {};

struct InboxBlocked {
    uint256_t inbox;
};

struct SendBlocked {
    uint256_t currency;
    TokenType tokenType;
};

using BlockReason = nonstd::variant<NotBlocked,
                                    HaltBlocked,
                                    ErrorBlocked,
                                    BreakpointBlocked,
                                    InboxBlocked,
                                    SendBlocked>;

struct AssertionContext {
    uint32_t numSteps;
    TimeBounds timeBounds;
    std::vector<Message> outMessage;
    std::vector<value> logs;

    explicit AssertionContext(const TimeBounds& tb)
        : numSteps{0}, timeBounds(tb) {}
};

struct Assertion {
    uint64_t stepCount;
    std::vector<Message> outMessages;
    std::vector<value> logs;
};

struct MessageStack {
    Tuple messages;
    uint64_t messageCount;
    TuplePool& pool;

    MessageStack(TuplePool& pool_) : pool(pool_) {}

    bool isEmpty() const { return messageCount == 0; }

    void addMessage(const Message& msg) {
        messages =
            Tuple{uint256_t{0}, std::move(messages), msg.toValue(pool), &pool};
        messageCount++;
    }

    void addMessageStack(MessageStack&& stack) {
        if (!stack.isEmpty()) {
            messages = Tuple(uint256_t(1), std::move(messages),
                             std::move(stack.messages), &pool);
            messageCount += stack.messageCount;
        }
    }

    void clear() {
        messages = Tuple{};
        messageCount = 0;
    }
};

struct MachineState {
    std::shared_ptr<TuplePool> pool;
    std::vector<CodePoint> code;
    value staticVal;
    value registerVal;
    datastack stack;
    datastack auxstack;
    Status state = Status::Extensive;
    uint64_t pc = 0;
    CodePoint errpc;
    MessageStack pendingInbox;
    AssertionContext context;
    MessageStack inbox;
    BalanceTracker balance;
    BlockReason blockReason;

    MachineState();

    bool deserialize(char* data);

    void readInbox(char* newInbox);
    std::vector<unsigned char> marshalForProof();
    uint64_t pendingMessageCount() const;
    void sendOnchainMessage(const Message& msg);
    void deliverOnchainMessages();
    void sendOffchainMessages(const std::vector<Message>& messages);
    BlockReason runOp(OpCode opcode);
    uint256_t hash() const;
};

class Machine {
    MachineState m;

    friend std::ostream& operator<<(std::ostream&, const Machine&);
    void runOne();

   public:
    bool deserialize(char* data) { return m.deserialize(data); }

    Assertion run(uint64_t stepCount,
                  uint64_t timeBoundStart,
                  uint64_t timeBoundEnd);
    Status currentStatus() { return m.state; }
    BlockReason lastBlockReason() { return m.blockReason; }
    uint256_t hash() const { return m.hash(); }
    std::vector<unsigned char> marshalForProof() { return m.marshalForProof(); }
    uint64_t pendingMessageCount() const { return m.pendingMessageCount(); }

    bool canSpend(const TokenType& tokType, const uint256_t& amount) const {
        return m.balance.CanSpend(tokType, amount);
    }
    uint256_t inboxHash() const { return ::hash(m.inbox.messages); }

    void sendOnchainMessage(const Message& msg);
    void deliverOnchainMessages();
    void sendOffchainMessages(const std::vector<Message>& messages);

    TuplePool& getPool() { return *m.pool; }
};

std::ostream& operator<<(std::ostream& os, const MachineState& val);
std::ostream& operator<<(std::ostream& os, const Machine& val);

#endif /* machine_hpp */
