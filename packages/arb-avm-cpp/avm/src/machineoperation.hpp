//
//  machineoperation.hpp
//  avm
//
//  Created by Minh Truong on 9/22/19.
//

#ifndef machineoperation_hpp
#define machineoperation_hpp

#include <stdio.h>
#include "machinestate.hpp"

class MachineOperation {
   public:
    static void add(MachineState& m);
    static void mul(MachineState& m);
    static void sub(MachineState& m);
    static void div(MachineState& m);
    static void sdiv(MachineState& m);
    static void mod(MachineState& m);
    static void smod(MachineState& m);
    static void addmod(MachineState& m);
    static void mulmod(MachineState& m);
    static void exp(MachineState& m);
    static void lt(MachineState& m);
    static void gt(MachineState& m);
    static void slt(MachineState& m);
    static void sgt(MachineState& m);
    static void eq(MachineState& m);
    static void iszero(MachineState& m);
    static void bitwiseAnd(MachineState& m);
    static void bitwiseOr(MachineState& m);
    static void bitwiseXor(MachineState& m);
    static void bitwiseNot(MachineState& m);
    static void byte(MachineState& m);
    static void signExtend(MachineState& m);
    static void hashOp(MachineState& m);
    static void typeOp(MachineState& m);
    static void pop(MachineState& m);
    static void spush(MachineState& m);
    static void rpush(MachineState& m);
    static void rset(MachineState& m);
    static void jump(MachineState& m);
    static void cjump(MachineState& m);
    static void stackEmpty(MachineState& m);
    static void pcPush(MachineState& m);
    static void auxPush(MachineState& m);
    static void auxPop(MachineState& m);
    static void auxStackEmpty(MachineState& m);
    static void errPush(MachineState& m);
    static void errSet(MachineState& m);
    static void dup0(MachineState& m);
    static void dup1(MachineState& m);
    static void dup2(MachineState& m);
    static void swap1(MachineState& m);
    static void swap2(MachineState& m);
    static void tget(MachineState& m);
    static void tset(MachineState& m);
    static void tlen(MachineState& m);
    static BlockReason breakpoint(MachineState&);
    static void log(MachineState& m);
    static void debug(MachineState& m);
    static BlockReason send(MachineState& m);
    static void nbsend(MachineState& m);
    static void getTime(MachineState& m);
    static BlockReason inboxOp(MachineState& m);
};

#endif /* machineoperation_hpp */
