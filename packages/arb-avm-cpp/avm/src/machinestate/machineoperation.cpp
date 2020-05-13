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

#include <avm/machinestate/machineoperation.hpp>
#include <avm/machinestate/machinestate.hpp>
#include <avm_values/util.hpp>
#include <bigint_utils.hpp>

#include <secp256k1_recovery.h>
#include <libff/algebra/curves/alt_bn128/alt_bn128_g1.hpp>

namespace machineoperation {

uint256_t& assumeInt(value& val) {
    auto aNum = nonstd::get_if<uint256_t>(&val);
    if (!aNum) {
        throw bad_pop_type{};
    }
    return *aNum;
}

const uint256_t& assumeInt(const value& val) {
    auto aNum = nonstd::get_if<uint256_t>(&val);
    if (!aNum) {
        throw bad_pop_type{};
    }
    return *aNum;
}

uint64_t assumeInt64(uint256_t& val) {
    if (val > std::numeric_limits<uint64_t>::max())
        throw int_out_of_bounds{};

    return static_cast<uint64_t>(val);
}

Tuple& assumeTuple(value& val) {
    auto tup = nonstd::get_if<Tuple>(&val);
    if (!tup) {
        throw bad_pop_type{};
    }
    return *tup;
}

void add(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum + bNum;
    m.stack.popClear();
    ++m.pc;
}

void mul(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum * bNum;
    m.stack.popClear();
    ++m.pc;
}

void sub(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum - bNum;
    m.stack.popClear();
    ++m.pc;
}

void div(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (bNum == 0) {
        m.state = Status::Error;
    } else {
        m.stack[1] = aNum / bNum;
    }
    m.stack.popClear();
    ++m.pc;
}

void sdiv(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    const auto min = (std::numeric_limits<uint256_t>::max() / 2) + 1;

    if (bNum == 0) {
        m.state = Status::Error;
    } else if (aNum == min && bNum == -1) {
        m.stack[1] = aNum;
    } else {
        const auto signA = get_sign(aNum);
        const auto signB = get_sign(bNum);
        if (signA == -1)
            aNum = 0 - aNum;
        if (signB == -1)
            bNum = 0 - bNum;
        m.stack[1] = (aNum / bNum) * signA * signB;
    }
    m.stack.popClear();
    ++m.pc;
}

void mod(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (bNum != 0) {
        m.stack[1] = aNum % bNum;
    } else {
        m.state = Status::Error;
    }
    m.stack.popClear();
    ++m.pc;
}

void smod(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    if (bNum == 0) {
        m.state = Status::Error;
    } else {
        const auto signA = get_sign(aNum);
        const auto signB = get_sign(bNum);
        if (signA == -1)
            aNum = 0 - aNum;
        if (signB == -1)
            bNum = 0 - bNum;
        m.stack[1] = (aNum % bNum) * signA;
    }
    m.stack.popClear();
    ++m.pc;
}

void addmod(MachineState& m) {
    m.stack.prepForMod(3);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    auto& cNum = assumeInt(m.stack[2]);

    if (cNum == 0) {
        m.state = Status::Error;
    } else {
        uint512_t aBig = aNum;
        uint512_t bBig = bNum;
        m.stack[2] = static_cast<uint256_t>((aBig + bBig) % cNum);
    }
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

void mulmod(MachineState& m) {
    m.stack.prepForMod(3);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    auto& cNum = assumeInt(m.stack[2]);

    if (cNum == 0) {
        m.state = Status::Error;
    } else {
        uint512_t aBig = aNum;
        uint512_t bBig = bNum;
        m.stack[2] = static_cast<uint256_t>((aBig * bBig) % cNum);
    }
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

void exp(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    uint64_t bSmall = assumeInt64(bNum);
    m.stack[1] = power(aNum, bSmall);
    m.stack.popClear();
    ++m.pc;
}

void lt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = (aNum < bNum) ? 1 : 0;
    m.stack.popClear();
    ++m.pc;
}

void gt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = (aNum > bNum) ? 1 : 0;
    m.stack.popClear();
    ++m.pc;
}

void slt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (aNum == bNum) {
        m.stack[1] = 0;
    } else {
        uint8_t signA = get_sign(aNum);
        uint8_t signB = get_sign(bNum);

        if (signA != signB) {
            m.stack[1] = signA == 1 ? 0 : 1;
        } else {
            m.stack[1] = aNum < bNum ? 1 : 0;
        }
    }
    m.stack.popClear();
    ++m.pc;
}

void sgt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (aNum == bNum) {
        m.stack[1] = 0;
    } else {
        uint8_t signA = get_sign(aNum);
        uint8_t signB = get_sign(bNum);

        if (signA != signB) {
            m.stack[1] = signA == 1 ? 1 : 0;
        } else {
            m.stack[1] = aNum > bNum ? 1 : 0;
        }
    }
    m.stack.popClear();
    ++m.pc;
}

void eq(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aVal = m.stack[0];
    auto& bVal = m.stack[1];
    m.stack[1] = aVal == bVal ? 1 : 0;
    m.stack.popClear();
    ++m.pc;
}

void iszero(MachineState& m) {
    m.stack.prepForMod(1);
    auto& aNum = assumeInt(m.stack[0]);
    m.stack[0] = aNum.is_zero() ? 1 : 0;
    ++m.pc;
}

void bitwiseAnd(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum & bNum;
    m.stack.popClear();
    ++m.pc;
}

void bitwiseOr(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum | bNum;
    m.stack.popClear();
    ++m.pc;
}

void bitwiseXor(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum ^ bNum;
    m.stack.popClear();
    ++m.pc;
}

void bitwiseNot(MachineState& m) {
    m.stack.prepForMod(1);
    auto& aNum = assumeInt(m.stack[0]);
    m.stack[0] = ~aNum;
    ++m.pc;
}

void byte(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    if (bNum >= 32) {
        m.stack[1] = 0;
    } else {
        const auto shift = 256 - 8 - 8 * shrink<uint8_t>(bNum);
        const auto mask = uint256_t(255) << shift;
        m.stack[1] = (aNum & mask) >> shift;
    }
    m.stack.popClear();
    ++m.pc;
}

void signExtend(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    if (bNum >= 32) {
        m.stack[1] = m.stack[0];
    } else {
        uint256_t t = 248 - 8 * bNum;
        int signBit = bit(aNum, (int)(255 - t));
        uint256_t mask = power(uint256_t(2), uint64_t(255 - t)) - 1;
        if (signBit == 0) {
            m.stack[1] = aNum & mask;
        } else {
            mask ^= -1;
            m.stack[1] = aNum | mask;
        }
    }
    m.stack.popClear();
    ++m.pc;
}

void hashOp(MachineState& m) {
    m.stack.prepForMod(1);
    m.stack[0] = ::hash(m.stack[0]);
    ++m.pc;
}

void typeOp(MachineState& m) {
    m.stack.prepForMod(1);
    if (nonstd::holds_alternative<uint256_t>(m.stack[0]))
        m.stack[0] = NUM;
    else if (nonstd::holds_alternative<CodePoint>(m.stack[0]))
        m.stack[0] = CODEPT;
    else if (nonstd::holds_alternative<Tuple>(m.stack[0]))
        m.stack[0] = TUPLE;
    ++m.pc;
}

void ethhash2Op(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    std::array<unsigned char, 64> inData;
    std::array<uint64_t, 4> anumInts;
    to_big_endian(aNum, anumInts.begin());
    std::copy(reinterpret_cast<unsigned char*>(anumInts.data()),
              reinterpret_cast<unsigned char*>(anumInts.data()) + 32,
              inData.begin());
    std::array<uint64_t, 4> bnumInts;
    to_big_endian(bNum, bnumInts.begin());
    std::copy(reinterpret_cast<unsigned char*>(bnumInts.data()),
              reinterpret_cast<unsigned char*>(bnumInts.data()) + 32,
              inData.begin() + 32);

    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(inData.data(), 64, hashData.data());

    m.stack[1] = from_big_endian(hashData.begin(), hashData.end());

    m.stack.popClear();
    ++m.pc;
}

void pop(MachineState& m) {
    m.stack.popClear();
    ++m.pc;
}

void spush(MachineState& m) {
    value copiedStatic = m.staticVal;
    m.stack.push(std::move(copiedStatic));
    ++m.pc;
}

void rpush(MachineState& m) {
    value copiedRegister = m.registerVal;
    m.stack.push(std::move(copiedRegister));
    ++m.pc;
}

void rset(MachineState& m) {
    m.stack.prepForMod(1);
    m.registerVal = m.stack[0];
    m.stack.popClear();
    ++m.pc;
}

void jump(MachineState& m) {
    m.stack.prepForMod(1);
    auto target = nonstd::get_if<CodePoint>(&m.stack[0]);
    if (target) {
        m.pc = target->pc;
    } else {
        m.state = Status::Error;
    }
    m.stack.popClear();
}

void cjump(MachineState& m) {
    m.stack.prepForMod(2);
    auto target = nonstd::get_if<CodePoint>(&m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (bNum != 0) {
        if (target) {
            m.pc = target->pc;
        } else {
            m.state = Status::Error;
        }
    } else {
        ++m.pc;
    }
    m.stack.popClear();
    m.stack.popClear();
}

void stackEmpty(MachineState& m) {
    if (m.stack.stacksize() == 0) {
        m.stack.push(uint256_t{1});
    } else {
        m.stack.push(uint256_t{0});
    }
    ++m.pc;
}

void pcPush(MachineState& m) {
    m.stack.push(m.code[m.pc]);
    ++m.pc;
}

void auxPush(MachineState& m) {
    m.stack.prepForMod(1);
    m.auxstack.push(std::move(m.stack[0]));
    m.stack.popClear();
    ++m.pc;
}

void auxPop(MachineState& m) {
    m.auxstack.prepForMod(1);
    m.stack.push(std::move(m.auxstack[0]));
    m.auxstack.popClear();
    ++m.pc;
}

void auxStackEmpty(MachineState& m) {
    if (m.auxstack.stacksize() == 0) {
        m.stack.push(uint256_t{1});
    } else {
        m.stack.push(uint256_t{0});
    }
    ++m.pc;
}

void errPush(MachineState& m) {
    m.stack.push(m.errpc);
    ++m.pc;
}

void errSet(MachineState& m) {
    m.stack.prepForMod(1);
    auto codePointVal = nonstd::get_if<CodePoint>(&m.stack[0]);
    if (!codePointVal) {
        m.state = Status::Error;
    } else {
        m.errpc = *codePointVal;
    }
    m.stack.popClear();
    ++m.pc;
}

void dup0(MachineState& m) {
    value valACopy = m.stack[0];
    m.stack.push(std::move(valACopy));
    ++m.pc;
}

void dup1(MachineState& m) {
    value valBCopy = m.stack[1];
    m.stack.push(std::move(valBCopy));
    ++m.pc;
}

void dup2(MachineState& m) {
    value valCCopy = m.stack[2];
    m.stack.push(std::move(valCCopy));
    ++m.pc;
}

void swap1(MachineState& m) {
    m.stack.prepForMod(2);
    std::swap(m.stack[0], m.stack[1]);
    ++m.pc;
}

void swap2(MachineState& m) {
    m.stack.prepForMod(3);
    std::swap(m.stack[0], m.stack[2]);
    ++m.pc;
}

void tget(MachineState& m) {
    m.stack.prepForMod(2);
    auto& bigIndex = assumeInt(m.stack[0]);
    auto index = assumeInt64(bigIndex);
    auto& tup = assumeTuple(m.stack[1]);
    m.stack[1] = tup.get_element(index);
    m.stack.popClear();
    ++m.pc;
}

void tset(MachineState& m) {
    m.stack.prepForMod(3);
    auto& bigIndex = assumeInt(m.stack[0]);
    auto index = assumeInt64(bigIndex);
    auto& tup = assumeTuple(m.stack[1]);
    tup.set_element(index, std::move(m.stack[2]));
    m.stack[2] = std::move(tup);
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

void tlen(MachineState& m) {
    m.stack.prepForMod(1);
    m.stack[0] = assumeTuple(m.stack[0]).tuple_size();
    ++m.pc;
}

void ec_recover(MachineState& m) {
    m.stack.prepForMod(4);
    secp256k1_context* context = secp256k1_context_create(
        SECP256K1_CONTEXT_SIGN | SECP256K1_CONTEXT_VERIFY);
    secp256k1_pubkey pubkey;
    secp256k1_ecdsa_recoverable_signature signature;
    for (int i = 0; i < 2; i++) {
        to_big_endian(assumeInt(m.stack[1 - i]), signature.data + (32 * i));
    }
    signature.data[64] = static_cast<unsigned char>(assumeInt(m.stack[3]));
    unsigned char message[32];
    for (int i = 0; i < 8; i++) {
        for (int j = 0; j < 4; j++) {
            message[4 * i + j] = static_cast<uint32_t>(assumeInt(m.stack[2]) >>
                                                       (248 - 32 * i - 8 * j));
        }
    }
    bool result =
        secp256k1_ecdsa_recover(context, &pubkey, &signature, message);
    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(pubkey.data, 64, hashData.data());
    auto cool = from_big_endian(hashData.begin(), hashData.end());
    m.stack.popClear();
    m.stack.popClear();
    m.stack[0] = cool;
    m.stack.push(uint256_t(result));
}

BlockReason breakpoint(MachineState& m) {
    ++m.pc;
    return BreakpointBlocked{};
}

void log(MachineState& m) {
    m.stack.prepForMod(1);
    m.context.logs.push_back(std::move(m.stack[0]));
    m.stack.popClear();
    ++m.pc;
}

void debug(MachineState& m) {
    Datastack tmpstk;
    std::cout << std::endl;
    std::cout << "full stack - size=" << m.stack.stacksize() << std::endl;
    while (m.stack.stacksize() > 0) {
        std::cout << m.stack[0] << std::endl;
        tmpstk.push(std::move(m.stack[0]));
        m.stack.popClear();
    }
    while (tmpstk.stacksize() > 0) {
        m.stack.push(std::move(tmpstk[0]));
        tmpstk.popClear();
    }
    std::cout << "register val=" << m.registerVal << std::endl << std::endl;
    ++m.pc;
}

BlockReason send(MachineState& m) {
    m.stack.prepForMod(1);
    m.context.outMessage.push_back(std::move(m.stack[0]));
    m.stack.popClear();
    ++m.pc;
    return NotBlocked();
}

void getTime(MachineState& m) {
    Tuple tup(m.pool.get(), 4);
    tup.set_element(0, m.context.timeBounds.lowerBoundBlock);
    tup.set_element(1, m.context.timeBounds.upperBoundBlock);
    tup.set_element(2, m.context.timeBounds.lowerBoundTimestamp);
    tup.set_element(3, m.context.timeBounds.upperBoundTimestamp);
    m.stack.push(std::move(tup));
    ++m.pc;
}

BlockReason inboxOp(MachineState& m) {
    m.stack.prepForMod(1);
    auto& aNum = assumeInt(m.stack[0]);
    if (aNum > m.context.timeBounds.lowerBoundBlock &&
        m.context.inbox.tuple_size() == 0) {
        return InboxBlocked(aNum);
    } else {
        m.stack[0] = std::move(m.context.inbox);
        ++m.pc;
        m.context.executedInbox();
        return NotBlocked{};
    }
}
}  // namespace machineoperation
