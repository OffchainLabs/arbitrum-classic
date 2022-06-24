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

const int send_size_limit = 10000;
constexpr int max_ec_pairing_points = 30;
constexpr int ec_pair_gas_cost = 500'000;

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
void shr(MachineState& m);
void sar(MachineState& m);
void hashOp(MachineState& m);
void typeOp(MachineState& m);
void ethhash2Op(MachineState& m);
void keccakF(MachineState& m);
void sha256F(MachineState& m);
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
void ec_add(MachineState& m);
void ec_mul(MachineState& m);
void ec_pairing(MachineState& m);
uint64_t ec_pairing_variable_gas_cost(MachineState& m);
BlockReason breakpoint(MachineState&);
void log(MachineState& m);
void debug(MachineState& m);
void send(MachineState& m);
BlockReason inboxOp(MachineState& m);
void setgas(MachineState& m);
void pushgas(MachineState& m);
void errcodept(MachineState& m);
void pushinsn(MachineState& m);
void pushinsnimm(MachineState& m);
BlockReason sideload(MachineState& m);

void newbuffer(MachineState& m);
void getbuffer8(MachineState& m);
void getbuffer64(MachineState& m);
void getbuffer256(MachineState& m);
void setbuffer8(MachineState& m);
void setbuffer64(MachineState& m);
void setbuffer256(MachineState& m);

namespace internal {
void encodeKeccakState(const Tuple& tup, uint64_t* state);
Tuple decodeKeccakState(const uint64_t* state);

uint256_t sha256_block(const uint256_t& digest_int,
                       std::array<uint8_t, 64>& input_data);
}  // namespace internal
}  // namespace machineoperation

#endif /* machineoperation_hpp */
