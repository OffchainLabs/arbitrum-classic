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

#include <avm/machinestate/blockreason.hpp>

struct MachineState;
class Tuple;
class TuplePool;

const int send_size_limit = 10000;

namespace machineoperation {
void add(MachineState& m);
void mul(MachineState& m);
void sub(MachineState& m);
void div(MachineState& m);
void sdiv(MachineState& m);
void mod(MachineState& m);
void smod(MachineState& m);
void addmod(MachineState& m);
void mulmod(MachineState& m);
void exp(MachineState& m);
void signExtend(MachineState& m);
void lt(MachineState& m);
void gt(MachineState& m);
void slt(MachineState& m);
void sgt(MachineState& m);
void eq(MachineState& m);
void iszero(MachineState& m);
void bitwiseAnd(MachineState& m);
void bitwiseOr(MachineState& m);
void bitwiseXor(MachineState& m);
void bitwiseNot(MachineState& m);
void byte(MachineState& m);
void shl(MachineState& m);
void shr(MachineState& m) void sar(MachineState& m) void hashOp(
    MachineState& m);
void typeOp(MachineState& m);
void ethhash2Op(MachineState& m);
void keccakF(MachineState& m);
void pop(MachineState& m);
void spush(MachineState& m);
void rpush(MachineState& m);
void rset(MachineState& m);
void jump(MachineState& m);
void cjump(MachineState& m);
void stackEmpty(MachineState& m);
void pcPush(MachineState& m);
void auxPush(MachineState& m);
void auxPop(MachineState& m);
void auxStackEmpty(MachineState& m);
void errPush(MachineState& m);
void errSet(MachineState& m);
void dup0(MachineState& m);
void dup1(MachineState& m);
void dup2(MachineState& m);
void swap1(MachineState& m);
void swap2(MachineState& m);
void tget(MachineState& m);
void tset(MachineState& m);
void tlen(MachineState& m);
void xget(MachineState& m);
void xset(MachineState& m);
void ec_recover(MachineState& m);
BlockReason breakpoint(MachineState&);
void log(MachineState& m);
void debug(MachineState& m);
bool send(MachineState& m);
BlockReason inboxOp(MachineState& m);
void setgas(MachineState& m);
void pushgas(MachineState& m);
void errcodept(MachineState& m);
void pushinsn(MachineState& m);
void pushinsnimm(MachineState& m);
BlockReason sideload(MachineState& m);

namespace internal {
void encodeKeccakState(const Tuple& tup, uint64_t* state);
Tuple decodeKeccakState(const uint64_t* state, TuplePool* pool);
}  // namespace internal
}  // namespace machineoperation

#endif /* machineoperation_hpp */
