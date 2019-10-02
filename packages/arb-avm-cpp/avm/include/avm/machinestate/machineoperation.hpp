/*
 * Copyright 2019, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
