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

#include "datastack.hpp"
#include "tokenTracker.hpp"
#include "value.hpp"

#include <memory>
#include <vector>

typedef std::array<uint256_t, 2> TimeBounds;

enum class Status { Extensive, Blocked, Halted, Error };

struct AssertionContext {
    TimeBounds timeBounds;
    std::vector<Message> outMessage;
    std::vector<value> logs;

    explicit AssertionContext(const TimeBounds& tb) : timeBounds(tb) {}
};

struct Assertion {
    uint64_t stepCount;
    std::vector<Message> outMessages;
    std::vector<value> logs;
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
    Tuple pendingInbox;
    AssertionContext context;
    Tuple inbox;
    BalanceTracker balance;

    MachineState();

    void deserialize(char* data);

    void readInbox(char* newInbox);
    std::vector<unsigned char> marshalForProof();
    bool hasPendingMessages() const;
    void sendOnchainMessage(const Message& msg);
    void deliverOnchainMessages();
    void sendOffchainMessages(const std::vector<Message>& messages);
    void runOp(OpCode opcode);
    uint256_t hash() const;

   private:
    void deliverMessageStack(Tuple&& messages);
};

class Machine {
    MachineState m;

    friend std::ostream& operator<<(std::ostream&, const Machine&);

   public:
    void deserialize(char* data) { m.deserialize(data); }

    Assertion run(uint64_t stepCount,
                  uint64_t timeBoundStart,
                  uint64_t timeBoundEnd);
    int runOne();
    uint256_t hash() const { return m.hash(); }
    std::vector<unsigned char> marshalForProof() { return m.marshalForProof(); }
    bool hasPendingMessages() const { return m.hasPendingMessages(); }

    uint256_t inboxHash() const { return ::hash(m.inbox); }

    void sendOnchainMessage(const Message& msg);
    void deliverOnchainMessages();
    void sendOffchainMessages(const std::vector<Message>& messages);

    TuplePool& getPool() { return *m.pool; }
};

std::ostream& operator<<(std::ostream& os, const MachineState& val);
std::ostream& operator<<(std::ostream& os, const Machine& val);

#endif /* machine_hpp */
