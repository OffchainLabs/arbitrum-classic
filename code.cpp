//
//  code.cpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/30/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#include "code.hpp"

value *instr::getimmediate(){
    if (immediate!=NULL){
        immediate->refcount++;
    }
    return immediate;
}

instr::~instr(){
    if (immediate!=NULL && immediate->type!=TUPLE){
        delete immediate;
    }
}

instr::instr(const instr &i){
    pc=i.pc;
    opcode = i.opcode;
    if (i.immediate!=NULL){
        immediate=new value(*(i.immediate));
    } else {
        immediate=NULL;
    }
//        uint256_t codeHash;
    
}

instr::instr(unsigned long long p, char o, value* i){
    pc=p;
    opcode=o;
    if (i!=NULL){
        immediate=new value(*i);
    } else {
        immediate=NULL;
    }
}
