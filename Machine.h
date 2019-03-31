//
//  Machine.h
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/30/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#ifndef Machine_h
#define Machine_h

typedef struct {
    char buf[32];
} Value;

typedef struct  {
    char InstrType;
    char Instr;
    char ValueType;
    char val[32];
} OpCode;

typedef struct  {
    unsigned long long pcCount;
    OpCode *code;
    char staticValue[32];
//    int state;
//    unsigned long long maxsteps;
} Machine;

//typedef struct {
//    unsigned long long steps_executed;
//    char afterHash[32];
//} Assertion;
typedef unsigned long long  Assertion;

// Creates a new machine object with the given code
#ifdef __cplusplus
extern "C" Machine *Init_machine(void *code, char *staticValue); // C++ compiler sees this
#else
Machine *Init_machine(void *code, char *staticValue);            // C compiler sees this
#endif

// Mutates machine by running it for stepCount steps and returns assertion data from the run
#ifdef __cplusplus
extern "C" Assertion run_machine(Machine *Machine, unsigned long long stepCount);
#else
Assertion run_machine(Machine *Machine, unsigned long long stepCount);
#endif

#endif /* Machine_h */
