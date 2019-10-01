//
//  messagestack.hpp
//  avm
//
//  Created by Minh Truong on 9/23/19.
//

#ifndef messagestack_hpp
#define messagestack_hpp

#include <avm/tuple.hpp>
#include <avm/value.hpp>
#include "avm/machinestatesaver.hpp"
#include "avm/tokenTracker.hpp"

struct MessageStack {
    Tuple messages;
    uint64_t messageCount;
    TuplePool* pool;

    MessageStack(TuplePool* pool_) : pool(pool_) {}

    bool isEmpty() const { return messageCount == 0; }

    void addMessage(const Message& msg) {
        messages =
            Tuple{uint256_t{0}, std::move(messages), msg.toValue(*pool), pool};
        messageCount++;
    }

    void addMessageStack(MessageStack&& stack) {
        if (!stack.isEmpty()) {
            messages = Tuple(uint256_t(1), std::move(messages),
                             std::move(stack.messages), pool);
            messageCount += stack.messageCount;
        }
    }

    void clear() {
        messages = Tuple{};
        messageCount = 0;
    }

    GetResults CheckpointState(MachineStateSaver msSaver) {
        return msSaver.SaveValue(messages);
    }
};

#endif /* messagestack_hpp */
