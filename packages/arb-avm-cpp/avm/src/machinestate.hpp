//
//  machinestate.hpp
//  avm
//
//  Created by Minh Truong on 9/22/19.
//

#ifndef machinestate_hpp
#define machinestate_hpp

#include <stdio.h>
#include <avm/datastack.hpp>
#include <avm/tokenTracker.hpp>
#include <avm/value.hpp>
#include <memory>
#include <vector>

enum class Status { Extensive, Halted, Error };

typedef std::array<uint256_t, 2> TimeBounds;

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

struct AssertionContext {
    uint32_t numSteps;
    TimeBounds timeBounds;
    std::vector<Message> outMessage;
    std::vector<value> logs;

    explicit AssertionContext(const TimeBounds& tb)
        : numSteps{0}, timeBounds(tb) {}
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

#endif /* machinestate_hpp */
