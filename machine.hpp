//
//  machine.hpp
//  AVMtest
//
//  Created by Harry Kalodner on 4/2/19.
//

#ifndef machine_hpp
#define machine_hpp

#include "value.hpp"
#include "datastack.hpp"
#include "code.hpp"

#include <memory>
#include <vector>
#include <stdio.h>

class datastack;

struct Assertion {
    uint64_t stepCount;
};

class Machine {
    std::vector<instr> code;
    value staticVal;
    value registerVal;
    datastack stack;
    datastack auxstack;
    int state;
    uint64_t pc;
    std::unique_ptr<TuplePool> pool;
    
    void runInstruction();
public:
    Machine();
    Machine(char *&srccode);
    
    Assertion run(uint64_t stepCount);
    int runOne();
};
instr deserialize_opcode(uint64_t pc, char *&bufptr, TuplePool &pool);

#endif /* machine_hpp */
