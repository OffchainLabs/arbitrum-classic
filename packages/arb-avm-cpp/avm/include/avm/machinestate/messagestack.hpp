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

#include "avm/value/tuple.hpp"
#include "avm/value/value.hpp"
#include "machinestatesaver.hpp"
#include "tokenTracker.hpp"

struct MessageStackSaveResults {
    SaveResults msgs_tuple_results;
    SaveResults msg_count_results;
};

struct MessageStackGetResults {
    GetResults msgs_tuple_results;
    GetResults msg_count_results;
};

struct MessageStack {
    Tuple messages;
    uint64_t messageCount;
    TuplePool* pool;

    MessageStack(TuplePool* pool_) : pool(pool_) { messageCount = 0; }

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

    MessageStackSaveResults checkpointState(MachineStateSaver& msSaver) {
        auto saved_msgs = msSaver.saveTuple(messages);
        auto converted_num = (uint256_t)messageCount;
        auto saved_msg_count = msSaver.saveValue(converted_num);

        return MessageStackSaveResults{saved_msgs, saved_msg_count};
    }

    bool initializeMessageStack(MachineStateSaver& msSaver,
                                std::vector<unsigned char> msgs_key,
                                std::vector<unsigned char> count_key) {
        auto msgs_res = msSaver.getTuple(msgs_key);
        auto count_res = msSaver.getInt256(count_key);

        if (msgs_res.status.ok() && count_res.status.ok()) {
            messages = msgs_res.data;
            messageCount = (uint64_t)count_res.data;
            return true;
        } else {
            return false;
        }
    }
};

#endif /* messagestack_hpp */
