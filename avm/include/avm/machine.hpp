//
//  machine.hpp
//  AVMtest
//
//  Created by Harry Kalodner on 4/2/19.
//

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
    void deliverMessageStack(value messages);
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
