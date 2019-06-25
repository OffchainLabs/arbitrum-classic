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

#include <memory>
#include <vector>

class datastack;

struct Assertion {
    uint64_t stepCount;
};

struct MachineState {
    std::unique_ptr<TuplePool> pool;
    std::vector<CodePoint> code;
    value staticVal;
    value registerVal;
    datastack stack;
    datastack auxstack;
    int state;
    uint64_t pc;
    CodePoint errpc;
    value inbox;

    MachineState();
    MachineState(std::vector<CodePoint> code);
    MachineState(char*& srccode, char*& inboxdata);

    void runOp(OpCode opcode);
    uint256_t hash() const;
};

class Machine {
    MachineState m;

    friend std::ostream& operator<<(std::ostream&, const Machine&);

   public:
    Machine(char*& srccode, char*& inboxdata) : m(srccode, inboxdata) {}

    Assertion run(uint64_t stepCount);
    int runOne();
    uint256_t hash() const { return m.hash(); }
};

std::ostream& operator<<(std::ostream& os, const MachineState& val);
std::ostream& operator<<(std::ostream& os, const Machine& val);

#endif /* machine_hpp */
