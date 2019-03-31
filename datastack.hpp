//
//  datastack.hpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/30/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#ifndef datastack_hpp
#define datastack_hpp

#include <stdio.h>
#include <stack>
#include <unordered_map>
#include "uint256_t.h"
#include "pool.hpp"
#include "value.hpp"

enum opcodes{
    HALT=0x00,
    ADD=0x01,
    MUL=0x02,
    POP=0x30,
    RPUSH=0x32,
    RSET=0x33,
    JUMP=0x35,
    PCPUSH=0x38,
    NOP=0x3C,
    TGET=0x50,
    TSET=0x51,
    PRTTOP=0xFF,
    PRTSTK=0xFE
};

enum states{EXTENSIVE, HALTED, ERROR};


class datastack{
public:
    stack<value*> basedatastack;
    unordered_map<uint64_t, uint64_t> pcmap;
    value *A;
    value *B;
    value *C;
    datastack *rest;
    unsigned int size;
    
    void push(value &newdata);
    void push(uint256_t val, int type);
    void pop();
    void popNoDel();
    value *top();
    uint64_t stacksize();
    
    void rpush(value &val);
    void rset(value &val);
    void pcpush(uint64_t i, uint64_t j);
    pair<int,uint64_t> jmp();
    int tget();
    int tset();
    int add();
    int mul();
};

#endif /* datastack_hpp */
