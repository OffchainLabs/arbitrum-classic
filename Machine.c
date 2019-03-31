//
//  Machine.c
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/30/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "Machine.h"

// Creates a new machine object with the given code
Machine *Init_machine(void *srccode, char *staticValue){
    Machine *ret=(Machine *)malloc( sizeof(Machine) );

    char *bufptr=srccode;
    memcpy(&ret->pcCount, bufptr, sizeof(unsigned long long));
    bufptr+=sizeof(unsigned long long);
    ret->code = malloc(ret->pcCount*sizeof(OpCode));
    for (unsigned long long i=0; i<ret->pcCount; i++){
        memcpy(&ret->code[i].InstrType, bufptr, sizeof(ret->code[i].InstrType));
        bufptr+=sizeof(ret->code[i].InstrType);
        memcpy(&ret->code[i].Instr, bufptr, sizeof(ret->code[i].Instr));
        bufptr+=sizeof(ret->code[i].Instr);
        memcpy(&ret->code[i].ValueType, bufptr, sizeof(ret->code[i].ValueType));
        bufptr+=sizeof(ret->code[i].ValueType);
        if (ret->code[i].InstrType==0x01){
            memcpy(ret->code[i].val, bufptr, 32);
            bufptr+=32;
        } else {
            memset(ret->code[i].val, 0, 32);
        }
        
    }
    memset(&ret->staticValue, 0, 32);

    return ret;
}

