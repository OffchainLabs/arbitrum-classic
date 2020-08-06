/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

#include <ethash/keccak.h>
#include <secp256k1_recovery.h>
#include <ethash/keccak.hpp>
#include <libff/algebra/curves/alt_bn128/alt_bn128_g1.hpp>

// Many opcode implementations were inspired from the Apache 2.0 licensed EVM
// implementation https://github.com/ethereum/evmone

using namespace intx;

namespace {
template <typename T>
static T shrink(uint256_t i) {
    return static_cast<T>(i & std::numeric_limits<T>::max());
}
}  // namespace

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
    if (val > std::numeric_limits<uint64_t>::max()) {
        throw int_out_of_bounds{};
    }

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
    if (bNum == 0) {
        m.state = Status::Error;
    } else {
        m.stack[1] = sdivrem(aNum, bNum).quot;
    }
    m.stack.popClear();
    ++m.pc;
}

void mod(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (bNum == 0) {
        m.state = Status::Error;
    } else {
        m.stack[1] = aNum % bNum;
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
        m.stack[1] = sdivrem(aNum, bNum).rem;
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
        m.stack[2] = addmod(aNum, bNum, cNum);
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
        m.stack[2] = mulmod(aNum, bNum, cNum);
    }
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

void exp(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = exp(aNum, bNum);
    m.stack.popClear();
    ++m.pc;
}

void signExtend(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    if (aNum < 31) {
        auto sign_bit = 8 * narrow_cast<uint8_t>(aNum) + 7;
        auto sign_mask = uint256_t{1} << sign_bit;
        auto value_mask = sign_mask - 1;
        auto is_neg = (bNum & sign_mask) != 0;
        m.stack[1] = is_neg ? bNum | ~value_mask : bNum & value_mask;
    }
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
    m.stack[0] = aNum == 0 ? 1 : 0;
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

    if (aNum >= 32) {
        m.stack[1] = 0;
    } else {
        const auto shift = 256 - 8 - 8 * shrink<uint8_t>(aNum);
        const auto mask = uint256_t(255) << shift;
        m.stack[1] = (bNum & mask) >> shift;
    }
    m.stack.popClear();
    ++m.pc;
}

void shl(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = bNum << aNum;
    m.stack.popClear();
    ++m.pc;
}

void shr(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = bNum >> aNum;
    m.stack.popClear();
    ++m.pc;
}

void sar(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    if ((bNum & (uint256_t{1} << 255)) == 0) {
        shr(m);
        return;
    }

    if (aNum >= 256) {
        m.stack[1] = ~uint256_t{0};
    } else {
        m.stack[1] = (bNum >> aNum) | (~uint256_t{0} << (256 - aNum));
    }
    m.stack.popClear();
    ++m.pc;
}

void hashOp(MachineState& m) {
    m.stack.prepForMod(1);
    m.stack[0] = ::hash_value(m.stack[0]);
    ++m.pc;
}

void typeOp(MachineState& m) {
    m.stack.prepForMod(1);
    if (nonstd::holds_alternative<uint256_t>(m.stack[0]))
        m.stack[0] = NUM;
    else if (nonstd::holds_alternative<CodePointStub>(m.stack[0]))
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
    auto it = to_big_endian(aNum, inData.begin());
    to_big_endian(bNum, it);

    auto hash_val = ethash::keccak256(inData.data(), inData.size());
    m.stack[1] = be::load<uint256_t>(hash_val);

    m.stack.popClear();
    ++m.pc;
}

namespace internal {
void encodeKeccakState(const Tuple& tup, uint64_t* state) {
    if (tup.tuple_size() != 7) {
        throw bad_pop_type{};
    }

    for (uint64_t i = 0; i < 6; ++i) {
        intx::be::unsafe::store(reinterpret_cast<uint8_t*>(&state[i * 4]),
                                bswap(assumeInt(tup.get_element_unsafe(i))));
    }
    // Handle last val separately
    state[24] = static_cast<uint64_t>(assumeInt(tup.get_element_unsafe(6)));
}

Tuple decodeKeccakState(const uint64_t* state, TuplePool* pool) {
    return Tuple(bswap(intx::be::unsafe::load<uint256_t>(
                     reinterpret_cast<const uint8_t*>(&state[0]))),
                 bswap(intx::be::unsafe::load<uint256_t>(
                     reinterpret_cast<const uint8_t*>(&state[4]))),
                 bswap(intx::be::unsafe::load<uint256_t>(
                     reinterpret_cast<const uint8_t*>(&state[8]))),
                 bswap(intx::be::unsafe::load<uint256_t>(
                     reinterpret_cast<const uint8_t*>(&state[12]))),
                 bswap(intx::be::unsafe::load<uint256_t>(
                     reinterpret_cast<const uint8_t*>(&state[16]))),
                 bswap(intx::be::unsafe::load<uint256_t>(
                     reinterpret_cast<const uint8_t*>(&state[20]))),
                 uint256_t{state[24]}, pool);
}
}  // namespace internal

void keccakF(MachineState& m) {
    m.stack.prepForMod(1);
    auto& tup = assumeTuple(m.stack[0]);
    uint64_t state[25];

    internal::encodeKeccakState(tup, state);

    ethash_keccakf1600(state);

    m.stack[0] = internal::decodeKeccakState(state, m.pool.get());
    ++m.pc;
}

void pop(MachineState& m) {
    m.stack.popClear();
    ++m.pc;
}

void spush(MachineState& m) {
    value copiedStatic = m.static_val;
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
    auto target = nonstd::get_if<CodePointStub>(&m.stack[0]);
    if (target) {
        m.pc = target->pc;
    } else {
        m.state = Status::Error;
    }
    m.stack.popClear();
}

void cjump(MachineState& m) {
    m.stack.prepForMod(2);
    auto target = nonstd::get_if<CodePointStub>(&m.stack[0]);
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
    m.stack.push(CodePointStub{m.pc, m.loadCurrentInstruction()});
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
    auto codePointVal = nonstd::get_if<CodePointStub>(&m.stack[0]);
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

void xget(MachineState& m) {
    m.stack.prepForMod(1);
    auto& bigIndex = assumeInt(m.stack[0]);
    auto index = assumeInt64(bigIndex);
    auto& tup = assumeTuple(m.auxstack[0]);
    m.stack[0] = tup.get_element(index);
    ++m.pc;
}

void xset(MachineState& m) {
    m.stack.prepForMod(2);
    m.auxstack.prepForMod(1);
    auto& bigIndex = assumeInt(m.stack[0]);
    auto index = assumeInt64(bigIndex);
    auto& tup = assumeTuple(m.auxstack[0]);
    tup.set_element(index, std::move(m.stack[1]));
    m.auxstack[0] = std::move(tup);
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

void tlen(MachineState& m) {
    m.stack.prepForMod(1);
    m.stack[0] = assumeTuple(m.stack[0]).tuple_size();
    ++m.pc;
}

namespace {
uint256_t parseSignature(MachineState& m) {
    auto recovery_int = assumeInt(m.stack[2]);
    if (recovery_int != 0 && recovery_int != 1) {
        return 0;
    }
    std::array<unsigned char, 64> sig_raw;
    auto it = to_big_endian(assumeInt(m.stack[0]), sig_raw.begin());
    to_big_endian(assumeInt(m.stack[1]), it);

    auto message = be::store<ethash::hash256>(assumeInt(m.stack[3]));

    static secp256k1_context* context = secp256k1_context_create(
        SECP256K1_CONTEXT_SIGN | SECP256K1_CONTEXT_VERIFY);

    secp256k1_ecdsa_recoverable_signature sig;
    int parsed_sig = secp256k1_ecdsa_recoverable_signature_parse_compact(
        context, &sig, sig_raw.data(), static_cast<int>(recovery_int));
    if (!parsed_sig) {
        return 0;
    }

    secp256k1_pubkey pubkey;
    if (!secp256k1_ecdsa_recover(context, &pubkey, &sig, message.bytes)) {
        return 0;
    }

    std::array<unsigned char, 65> pubkey_raw;
    size_t output_length = pubkey_raw.size();
    int serialized_pubkey = secp256k1_ec_pubkey_serialize(
        context, pubkey_raw.data(), &output_length, &pubkey,
        SECP256K1_EC_UNCOMPRESSED);
    if (!serialized_pubkey) {
        return 0;
    }
    // Skip header byte
    auto hash_val = ethash::keccak256(pubkey_raw.data() + 1, 64);
    std::fill(&hash_val.bytes[0], &hash_val.bytes[12], 0);
    return be::load<uint256_t>(hash_val);
}
}  // namespace

void ec_recover(MachineState& m) {
    m.stack.prepForMod(4);

    m.stack[3] = parseSignature(m);
    m.stack.popClear();
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
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
    m.stack.prepForMod(1);
    auto val = m.stack.pop();
    std::cout << "debugprint " << val << std::endl;
    ++m.pc;
}

bool send(MachineState& m) {
    m.stack.prepForMod(1);

    auto val_size = getSize(m.stack[0]);
    bool success;

    if (val_size > send_size_limit) {
        success = false;
    } else {
        m.context.outMessage.push_back(std::move(m.stack[0]));
        m.stack.popClear();
        ++m.pc;

        success = true;
    }

    return success;
}

BlockReason inboxPeekOp(MachineState& m) {
    m.stack.prepForMod(1);
    bool has_staged_message = m.staged_message != Tuple{};
    if (!has_staged_message && m.context.inboxEmpty()) {
        if (!m.context.fake_inbox_peek_value) {
            return InboxBlocked();
        }

        // When fake_inbox_peek_value is set we're in callserver mode. Use
        // that value as the message value
        m.stack[0] = m.stack[0] == *m.context.fake_inbox_peek_value ? 1 : 0;
        ++m.pc;
        return NotBlocked{};
    }
    if (!has_staged_message) {
        m.staged_message = m.context.popInbox();
    }
    m.stack[0] = m.stack[0] == m.staged_message.get_element(0) ? 1 : 0;
    ++m.pc;
    return NotBlocked{};
}

BlockReason inboxOp(MachineState& m) {
    bool has_staged_message = m.staged_message != Tuple{};
    if (!has_staged_message && m.context.inboxEmpty()) {
        return InboxBlocked();
    }
    if (has_staged_message) {
        m.stack.push(m.staged_message);
        m.staged_message = Tuple();
    } else {
        m.stack.push(m.context.popInbox());
    }
    ++m.pc;
    return NotBlocked{};
}

void setgas(MachineState& m) {
    m.stack.prepForMod(1);
    auto& aNum = assumeInt(m.stack[0]);
    m.arb_gas_remaining = aNum;
    m.stack.popClear();
    ++m.pc;
}

void pushgas(MachineState& m) {
    auto& gas = m.arb_gas_remaining;
    m.stack.push(gas);
    ++m.pc;
}

void errcodept(MachineState& m) {
    m.stack.push(m.code->addSegment());
    ++m.pc;
}

void pushinsn(MachineState& m) {
    m.stack.prepForMod(2);
    auto target = nonstd::get_if<CodePointStub>(&m.stack[1]);
    if (!target) {
        m.state = Status::Error;
        return;
    }
    auto& op_int = assumeInt(m.stack[0]);
    auto op = static_cast<uint8_t>(op_int);
    m.stack[1] = m.code->addOperation(target->pc, {static_cast<OpCode>(op)});
    m.stack.popClear();
    ++m.pc;
}

void pushinsnimm(MachineState& m) {
    m.stack.prepForMod(3);
    auto target = nonstd::get_if<CodePointStub>(&m.stack[2]);
    if (!target) {
        m.state = Status::Error;
        return;
    }
    auto& op_int = assumeInt(m.stack[0]);
    auto op = static_cast<uint8_t>(op_int);
    m.stack[2] = m.code->addOperation(
        target->pc, {static_cast<OpCode>(op), std::move(m.stack[1])});
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

BlockReason sideload(MachineState& m) {
    if (m.context.sideload_value.tuple_size() != 0) {
        m.stack.push(m.context.sideload_value);
        m.context.sideload_value = Tuple{};
    } else {
        if (m.context.numSteps != 0 && m.context.blockingSideload) {
            return SideloadBlocked{};
        }
        m.stack.push(Tuple{});
    }
    ++m.pc;
    return NotBlocked{};
}

}  // namespace machineoperation
