//
//  code.hpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/30/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#ifndef code_hpp
#define code_hpp

#include "uint256_t.h"
#include "value.hpp"

//class codept{
//public:
//    int pc;
//    int opcode;
//    uint256_t nextCodeHash;
//};

class instr{
private:
    value *immediate;
    
public:
    unsigned long long pc;
    char opcode;
    uint256_t codeHash;

    value *getimmediate();
    ~instr();
    instr(const instr &i);
    instr(unsigned long long p, char o, value* i);
    //    codept label;
};

#endif /* code_hpp */
