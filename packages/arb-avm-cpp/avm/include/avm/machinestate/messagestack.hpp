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

#include <avm_values/tuple.hpp>
#include <data_storage/storageresult.hpp>

class MachineStateSaver;
class MachineStateFetcher;

struct MessageStackSaveResults {
    SaveResults msgs_tuple_results;
};

struct MessageStack {
    Tuple messages;
    TuplePool* pool;

    MessageStack(TuplePool* pool_) : pool(pool_) {}

    bool isEmpty() const { return messages == Tuple{}; }

    void addMessages(Tuple&& new_messages) {
        if (new_messages != Tuple()) {
            messages = Tuple(uint256_t(1), std::move(messages),
                             std::move(new_messages), pool);
        }
    }

    void clear() { messages = Tuple{}; }

    MessageStackSaveResults checkpointState(MachineStateSaver& msSaver);

    bool initializeMessageStack(const MachineStateFetcher& fetcher,
                                const std::vector<unsigned char>& msgs_key);
};

#endif /* messagestack_hpp */
