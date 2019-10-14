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

#ifndef messagestack_hpp
#define messagestack_hpp

#include "machinestatesaver.hpp"
#include "tokenTracker.hpp"
#include "value/tuple.hpp"
#include "value/value.hpp"

struct MessageStackSaveResults {
    SaveResults msgs_tuple_results;
    SaveResults msg_count_results;
};

struct MessageStack {
    Tuple messages;
    uint64_t messageCount;
    TuplePool* pool;

    MessageStack(TuplePool* pool_) : pool(pool_) { messageCount = 0; }

    MessageStack(TuplePool* pool_, Tuple tuple, uint256_t message_count)
        : pool(pool_) {
        messages = tuple;
        messageCount = (uint64_t)message_count;
    }

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

    MessageStackSaveResults checkpointState(MachineStateSaver msSaver) {
        auto saved_msgs = msSaver.SaveTuple(messages);
        auto saved_msg_count = msSaver.SaveValue((uint256_t)messageCount);

        return MessageStackSaveResults{saved_msgs, saved_msg_count};
    }
};

#endif /* messagestack_hpp */
