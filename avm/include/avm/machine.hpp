//
//  machine.hpp
//  AVMtest
//
//  Created by Harry Kalodner on 4/2/19.
//

#ifndef machine_hpp
#define machine_hpp

#include "datastack.hpp"
#include "value.hpp"
#include "tokenTracker.hpp"

#include <memory>
#include <vector>

class datastack;

typedef uint256_t TimeBounds[2];

struct Precondition {
    uint256_t BeforeHash;
    TimeBounds timeBounds;
    BalanceTracker beforeBalance;
    value beforeInbox;
};

struct AssertionContext {
    uint64_t stepCount;
    Precondition precondition;
    value beforeInbox;
    BalanceTracker afterBalance;
    std::vector<Message> outMessage;
    std::vector<value> logs;
};

struct Assertion {
    uint64_t stepCount;
};

struct MsgQueue {
    value msg;
    BalanceTracker balance;
};

struct Inbox {
    Tuple *accepted;
    MsgQueue pending;
};

struct MachineState {
    std::shared_ptr<TuplePool> pool;
    std::vector<CodePoint> code;
    value staticVal;
    value registerVal;
    datastack stack;
    datastack auxstack;
    int state;
    uint64_t pc;
    CodePoint errpc;
    value pendingInbox;
    value inbox;
    AssertionContext context;
		
    MachineState();
    MachineState(std::vector<CodePoint> code);
    MachineState(char*& srccode, char*& inboxdata, int inbox_sz);
    MachineState(MachineState const &m);

    void readInbox(char *newInbox);
    void addInboxMessage(char*& newMsg);
    void addInboxMessage(value msg);
    void addInboxMessage(Message &msg);
    void deliverMessages();
    void setTimebounds(uint64_t timeBoundStart, uint64_t timeBoundEnd);
    void runOp(OpCode opcode);
    uint256_t hash() const;
};

class Machine {
    MachineState m;

    friend std::ostream& operator<<(std::ostream&, const Machine&);

   public:
    Machine(char*& srccode, char*& inboxdata, int inbox_sz) : m(srccode, inboxdata, inbox_sz) {}
    Machine(){}
    Machine(Machine const &msrc) : m(msrc.m){}
    Assertion run(uint64_t stepCount, uint64_t timeBoundStart, uint64_t timeBoundEnd);
    Assertion runUntilStop(uint64_t timeBoundStart, uint64_t timeBoundEnd);
    int runOne();
    uint256_t hash() const { return m.hash(); }
    void addInboxMessage(char *msg);
    void deliverMessages();
    void setTimebounds(uint64_t timeboundStart, uint64_t timeboundEnd) {m.setTimebounds(timeboundStart, timeboundEnd);};
};

std::ostream& operator<<(std::ostream& os, const MachineState& val);
std::ostream& operator<<(std::ostream& os, const Machine& val);

#endif /* machine_hpp */
