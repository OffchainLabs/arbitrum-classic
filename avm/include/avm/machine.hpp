//
//  machine.hpp
//  AVMtest
//
//  Created by Harry Kalodner on 4/2/19.
//

#ifndef machine_hpp
#define machine_hpp

#include "code.hpp"
#include "datastack.hpp"
#include "value.hpp"

#include <memory>
#include <vector>

class datastack;

struct Assertion {
    uint64_t stepCount;
};

struct MachineState {
    std::vector<instr> code;
    value staticVal;
    value registerVal;
    datastack stack;
    datastack auxstack;
    int state;
    uint64_t pc;
    CodePoint errpc;
    value inbox;
    std::unique_ptr<TuplePool> pool;

    MachineState();
    MachineState(std::vector<instr> code);
    MachineState(char*& srccode, char*& inboxdata);

    void runOp(OpCode opcode);
};

class Machine {
    MachineState m;

   public:
    Machine(char*& srccode, char*& inboxdata) : m(srccode, inboxdata) {}

    Assertion run(uint64_t stepCount);
    int runOne();
};
instr deserialize_opcode(uint64_t pc, char*& bufptr, TuplePool& pool);

#endif /* machine_hpp */
