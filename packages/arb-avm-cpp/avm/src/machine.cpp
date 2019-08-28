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

#include "avm/machine.hpp"

#include "avm/opcodes.hpp"

#include "bigint_utils.hpp"
#include "util.hpp"

#include <iostream>

namespace {

std::vector<CodePoint> opsToCodePoints(const std::vector<Operation>& ops) {
    std::vector<CodePoint> cps;
    cps.reserve(ops.size());
    uint64_t pc = 0;
    for (auto& op : ops) {
        cps.emplace_back(pc, std::move(op), 0);
        pc++;
    }
    for (uint64_t i = 0; i < cps.size() - 1; i++) {
        cps[cps.size() - 2 - i].nextHash = hash(cps[cps.size() - 1 - i]);
    }
    return cps;
}
}  // namespace

class bad_pop_type : public std::exception {
   public:
    virtual const char* what() const noexcept override {
        return "bad_variant_access";
    }
};

class int_out_of_bounds : public std::exception {
   public:
    virtual const char* what() const noexcept override {
        return "int_out_of_bounds";
    }
};

MachineState::MachineState()
    : pool(std::make_unique<TuplePool>()),
      pendingInbox(*pool.get()),
      context({0, 0}),
      inbox(*pool.get()) {}

uint256_t MachineState::hash() const {
    if (state == Status::Halted)
        return 0;
    if (state == Status::Error)
        return 1;

    std::array<unsigned char, 32 * 6> data;
    auto oit = data.begin();
    {
        auto val = ::hash(code[pc]);
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }
    {
        auto val = stack.hash();
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }
    {
        auto val = auxstack.hash();
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }
    {
        auto val = ::hash(registerVal);
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }
    {
        auto val = ::hash(staticVal);
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }
    {
        auto val = ::hash(errpc);
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }

    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(data.data(), 32 * 6, hashData.data());
    return from_big_endian(hashData.begin(), hashData.end());
}

std::ostream& operator<<(std::ostream& os, const MachineState& val) {
    os << "codePointHash " << to_hex_str(hash(val.code[val.pc])) << "\n";
    os << "stackHash " << to_hex_str(val.stack.hash()) << "\n";
    os << "auxStackHash " << to_hex_str(val.auxstack.hash()) << "\n";
    os << "registerHash " << to_hex_str(hash(val.registerVal)) << "\n";
    os << "staticHash " << to_hex_str(hash(val.staticVal)) << "\n";
    os << "errHandlerHash " << to_hex_str(hash(val.errpc)) << "\n";
    return os;
}

std::ostream& operator<<(std::ostream& os, const Machine& val) {
    os << val.m;
    return os;
}

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

bool MachineState::deserialize(char* bufptr) {
    uint32_t version;
    memcpy(&version, bufptr, sizeof(version));
    version = __builtin_bswap32(version);
    bufptr += sizeof(version);

    if (version != CURRENT_AO_VERSION) {
        std::cerr << "incorrect version of .ao file" << std::endl;
        std::cerr << "expected version " << CURRENT_AO_VERSION
                  << " found version " << version << std::endl;
        return false;
    }

    uint32_t extentionId = 1;
    while (extentionId != 0) {
        memcpy(&extentionId, bufptr, sizeof(extentionId));
        extentionId = __builtin_bswap32(extentionId);
        bufptr += sizeof(extentionId);
        if (extentionId > 0) {
            //            std::cout << "found extention" << std::endl;
        }
    }
    uint64_t codeCount;
    memcpy(&codeCount, bufptr, sizeof(codeCount));
    bufptr += sizeof(codeCount);
    codeCount = boost::endian::big_to_native(codeCount);
    code.reserve(codeCount);

    std::vector<Operation> ops;
    for (uint64_t i = 0; i < codeCount; i++) {
        ops.emplace_back(deserializeOperation(bufptr, *pool));
    }
    code = opsToCodePoints(ops);

    staticVal = deserialize_value(bufptr, *pool);
    pc = 0;
    return true;
}

uint64_t MachineState::pendingMessageCount() const {
    return pendingInbox.messageCount;
}

void MachineState::sendOnchainMessage(const Message& msg) {
    pendingInbox.addMessage(msg);
    balance.add(msg.token, msg.currency);
}

void MachineState::sendOffchainMessages(const std::vector<Message>& messages) {
    MessageStack messageStack(*pool.get());
    for (const auto& message : messages) {
        messageStack.addMessage(message);
    }
    inbox.addMessageStack(std::move(messageStack));
}

void MachineState::deliverOnchainMessages() {
    inbox.addMessageStack(std::move(pendingInbox));
    pendingInbox.clear();
}

void uint256_t_to_buf(uint256_t val, std::vector<unsigned char>& buf) {
    std::vector<unsigned char> tmpbuf;
    tmpbuf.resize(32);
    to_big_endian(val, tmpbuf.begin());
    buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
}

std::vector<unsigned char> MachineState::marshalForProof() {
    std::vector<unsigned char> buf;
    std::vector<bool> stackPops = InstructionStackPops.at(code[pc].op.opcode);
    if (code[pc].op.immediate) {
        stackPops.erase(stackPops.begin());
    }
    std::vector<bool> auxStackPops =
        InstructionAuxStackPops.at(code[pc].op.opcode);
    std::vector<value> stackVals;
    uint256_t baseStackHash = stack.SolidityProofValue(stackPops, stackVals);
    std::vector<value> auxStackVals;
    uint256_t baseAuxStackHash =
        auxstack.SolidityProofValue(auxStackPops, auxStackVals);
    uint256_t registerHash = ::hash(registerVal);
    uint256_t staticHash = ::hash(staticVal);
    uint256_t errHandlerHash = ::hash(errpc);
    uint256_t_to_buf(code[pc].nextHash, buf);
    uint256_t_to_buf(baseStackHash, buf);
    uint256_t_to_buf(baseAuxStackHash, buf);
    uint256_t_to_buf(registerHash, buf);
    uint256_t_to_buf(staticHash, buf);
    uint256_t_to_buf(errHandlerHash, buf);
    code[pc].op.marshal(buf);
    for (auto const& stackval : stackVals) {
        marshal_value(stackval, buf);
    }
    for (auto const& auxstackval : auxStackVals) {
        marshal_value(auxstackval, buf);
    }
    return buf;
}

void Machine::sendOnchainMessage(const Message& msg) {
    m.sendOnchainMessage(msg);
}

void Machine::deliverOnchainMessages() {
    m.deliverOnchainMessages();
}

void Machine::sendOffchainMessages(const std::vector<Message>& messages) {
    m.sendOffchainMessages(messages);
}

Assertion Machine::run(uint64_t stepCount,
                       uint64_t timeBoundStart,
                       uint64_t timeBoundEnd) {
    m.context = AssertionContext{TimeBounds{{timeBoundStart, timeBoundEnd}}};
    m.blockReason = NotBlocked{};
    while (m.context.numSteps < stepCount) {
        runOne();
        if (!nonstd::get_if<NotBlocked>(&m.blockReason)) {
            break;
        }
    }
    return {m.context.numSteps, std::move(m.context.outMessage),
            std::move(m.context.logs)};
}

bool isErrorCodePoint(const CodePoint& cp) {
    return cp.nextHash == 0 && cp.op == Operation{static_cast<OpCode>(0)};
}

void Machine::runOne() {
    if (m.state == Status::Error) {
        m.blockReason = ErrorBlocked();
        return;
    }

    if (m.state == Status::Halted) {
        m.blockReason = HaltBlocked();
        return;
    }

    m.context.numSteps++;

    auto& instruction = m.code[m.pc];

    auto startStackSize = m.stack.stacksize();

    if (!isValidOpcode(instruction.op.opcode)) {
        m.state = Status::Error;
    } else {
        if (instruction.op.immediate) {
            auto imm = *instruction.op.immediate;
            m.stack.push(std::move(imm));
        }

        try {
            m.blockReason = m.runOp(instruction.op.opcode);
        } catch (const bad_pop_type& e) {
            m.state = Status::Error;
        } catch (const bad_tuple_index& e) {
            m.state = Status::Error;
        }
    }

    if (m.state != Status::Error) {
        return;
    }

    // Clear stack to base for instruction
    auto stackItems = InstructionStackPops.at(instruction.op.opcode).size();
    while (m.stack.stacksize() > 0 &&
           startStackSize - m.stack.stacksize() < stackItems) {
        m.stack.popClear();
    }

    if (!isErrorCodePoint(m.errpc)) {
        m.pc = m.errpc.pc;
        m.state = Status::Extensive;
    }

    return;
}

template <typename T>
static T shrink(uint256_t i) {
    return static_cast<T>(i & std::numeric_limits<T>::max());
}

namespace {
static void add(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum + bNum;
    m.stack.popClear();
    ++m.pc;
}

static void mul(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum * bNum;
    m.stack.popClear();
    ++m.pc;
}

static void sub(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum - bNum;
    m.stack.popClear();
    ++m.pc;
}

static void div(MachineState& m) {
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

static void sdiv(MachineState& m) {
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

static void mod(MachineState& m) {
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

static void smod(MachineState& m) {
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

static void addmod(MachineState& m) {
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

static void mulmod(MachineState& m) {
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

static void exp(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    uint64_t bSmall = assumeInt64(bNum);
    m.stack[1] = power(aNum, bSmall);
    m.stack.popClear();
    ++m.pc;
}

static void lt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = (aNum < bNum) ? 1 : 0;
    m.stack.popClear();
    ++m.pc;
}

static void gt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = (aNum > bNum) ? 1 : 0;
    m.stack.popClear();
    ++m.pc;
}

static void slt(MachineState& m) {
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

static void sgt(MachineState& m) {
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

static void eq(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aVal = m.stack[0];
    auto& bVal = m.stack[1];
    m.stack[1] = aVal == bVal ? 1 : 0;
    m.stack.popClear();
    ++m.pc;
}

static void iszero(MachineState& m) {
    m.stack.prepForMod(1);
    auto& aNum = assumeInt(m.stack[0]);
    m.stack[0] = aNum.is_zero() ? 1 : 0;
    ++m.pc;
}

static void bitwiseAnd(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum & bNum;
    m.stack.popClear();
    ++m.pc;
}

static void bitwiseOr(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum | bNum;
    m.stack.popClear();
    ++m.pc;
}

static void bitwiseXor(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum ^ bNum;
    m.stack.popClear();
    ++m.pc;
}

static void bitwiseNot(MachineState& m) {
    m.stack.prepForMod(1);
    auto& aNum = assumeInt(m.stack[0]);
    m.stack[0] = ~aNum;
    ++m.pc;
}

static void byte(MachineState& m) {
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

static void signExtend(MachineState& m) {
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

static void hashOp(MachineState& m) {
    m.stack.prepForMod(1);
    m.stack[0] = ::hash(m.stack[0]);
    ++m.pc;
}

static void typeOp(MachineState& m) {
    m.stack.prepForMod(1);
    if (nonstd::holds_alternative<uint256_t>(m.stack[0]))
        m.stack[0] = NUM;
    else if (nonstd::holds_alternative<CodePoint>(m.stack[0]))
        m.stack[0] = CODEPT;
    else if (nonstd::holds_alternative<Tuple>(m.stack[0]))
        m.stack[0] = TUPLE;
    ++m.pc;
}

static void pop(MachineState& m) {
    m.stack.popClear();
    ++m.pc;
}

static void spush(MachineState& m) {
    value copiedStatic = m.staticVal;
    m.stack.push(std::move(copiedStatic));
    ++m.pc;
}

static void rpush(MachineState& m) {
    value copiedRegister = m.registerVal;
    m.stack.push(std::move(copiedRegister));
    ++m.pc;
}

static void rset(MachineState& m) {
    m.stack.prepForMod(1);
    m.registerVal = m.stack[0];
    m.stack.popClear();
    ++m.pc;
}

static void jump(MachineState& m) {
    m.stack.prepForMod(1);
    auto target = nonstd::get_if<CodePoint>(&m.stack[0]);
    if (target) {
        m.pc = target->pc;
    } else {
        m.state = Status::Error;
    }
    m.stack.popClear();
}

static void cjump(MachineState& m) {
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

static void stackEmpty(MachineState& m) {
    if (m.stack.stacksize() == 0) {
        m.stack.push(uint256_t{1});
    } else {
        m.stack.push(uint256_t{0});
    }
    ++m.pc;
}

static void pcPush(MachineState& m) {
    m.stack.push(m.code[m.pc]);
    ++m.pc;
}

static void auxPush(MachineState& m) {
    m.stack.prepForMod(1);
    m.auxstack.push(std::move(m.stack[0]));
    m.stack.popClear();
    ++m.pc;
}

static void auxPop(MachineState& m) {
    m.auxstack.prepForMod(1);
    m.stack.push(std::move(m.auxstack[0]));
    m.auxstack.popClear();
    ++m.pc;
}

static void auxStackEmpty(MachineState& m) {
    if (m.auxstack.stacksize() == 0) {
        m.stack.push(uint256_t{1});
    } else {
        m.stack.push(uint256_t{0});
    }
    ++m.pc;
}

static void errPush(MachineState& m) {
    m.stack.push(m.errpc);
    ++m.pc;
}

static void errSet(MachineState& m) {
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

static void dup0(MachineState& m) {
    value valACopy = m.stack[0];
    m.stack.push(std::move(valACopy));
    ++m.pc;
}

static void dup1(MachineState& m) {
    value valBCopy = m.stack[1];
    m.stack.push(std::move(valBCopy));
    ++m.pc;
}

static void dup2(MachineState& m) {
    value valCCopy = m.stack[2];
    m.stack.push(std::move(valCCopy));
    ++m.pc;
}

static void swap1(MachineState& m) {
    m.stack.prepForMod(2);
    value temp = m.stack[0];
    m.stack[0] = m.stack[1];
    m.stack[1] = temp;
    ++m.pc;
}

static void swap2(MachineState& m) {
    m.stack.prepForMod(3);
    value temp = m.stack[0];
    m.stack[0] = m.stack[2];
    m.stack[2] = temp;
    ++m.pc;
}

static void tget(MachineState& m) {
    m.stack.prepForMod(2);
    auto& index = assumeInt(m.stack[0]);
    auto& tup = assumeTuple(m.stack[1]);
    m.stack[1] = tup.get_element(static_cast<uint32_t>(index));
    m.stack.popClear();
    ++m.pc;
}

static void tset(MachineState& m) {
    m.stack.prepForMod(3);
    auto& index = assumeInt(m.stack[0]);
    auto& tup = assumeTuple(m.stack[1]);
    tup.set_element(static_cast<uint32_t>(index), std::move(m.stack[2]));
    m.stack[2] = std::move(tup);
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

static void tlen(MachineState& m) {
    m.stack.prepForMod(1);
    m.stack[0] = assumeTuple(m.stack[0]).tuple_size();
    ++m.pc;
}

static BlockReason breakpoint(MachineState&) {
    return BreakpointBlocked{};
}

static void log(MachineState& m) {
    m.stack.prepForMod(1);
    m.context.logs.push_back(std::move(m.stack[0]));
    m.stack.popClear();
    ++m.pc;
}

static void debug(MachineState& m) {
    datastack tmpstk;
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

static BlockReason send(MachineState& m) {
    m.stack.prepForMod(1);
    Message outMsg;
    auto success = outMsg.deserialize(m.stack[0]);
    if (!success) {
        m.state = Status::Error;
        return NotBlocked();
    }
    if (!m.balance.Spend(outMsg.token, outMsg.currency)) {
        return SendBlocked{outMsg.currency, outMsg.token};
    } else {
        m.stack.popClear();
        m.context.outMessage.push_back(outMsg);
        ++m.pc;
        return NotBlocked();
    }
}

static void nbsend(MachineState& m) {
    m.stack.prepForMod(1);

    Message outMsg;
    auto success = outMsg.deserialize(m.stack[0]);
    if (!success) {
        m.state = Status::Error;
        return;
    }

    bool spent = m.balance.Spend(outMsg.token, outMsg.currency);
    if (!spent) {
        m.stack[0] = 0;
    } else {
        m.context.outMessage.push_back(outMsg);
        m.stack[0] = 1;
    }
    ++m.pc;
}

static void getTime(MachineState& m) {
    Tuple tup(m.pool.get(), 2);
    tup.set_element(0, m.context.timeBounds[0]);
    tup.set_element(1, m.context.timeBounds[1]);
    m.stack.push(std::move(tup));
    ++m.pc;
}

static BlockReason inboxOp(MachineState& m) {
    m.stack.prepForMod(1);
    auto stackTop = nonstd::get_if<Tuple>(&m.stack[0]);
    if (stackTop && m.inbox.messages == *stackTop) {
        return InboxBlocked{hash(m.inbox.messages)};
    } else {
        value inboxCopy = m.inbox.messages;
        m.stack[0] = std::move(inboxCopy);
        ++m.pc;
        return NotBlocked{};
    }
}
}  // namespace

BlockReason MachineState::runOp(OpCode opcode) {
    switch (opcode) {
        /**************************/
        /*  Arithmetic Operations */
        /**************************/
        case OpCode::ADD:
            add(*this);
            break;
        case OpCode::MUL:
            mul(*this);
            break;
        case OpCode::SUB:
            sub(*this);
            break;
        case OpCode::DIV:
            div(*this);
            break;
        case OpCode::SDIV:
            sdiv(*this);
            break;
        case OpCode::MOD:
            mod(*this);
            break;
        case OpCode::SMOD:
            smod(*this);
            break;
        case OpCode::ADDMOD:
            addmod(*this);
            break;
        case OpCode::MULMOD:
            mulmod(*this);
            break;
        case OpCode::EXP:
            exp(*this);
            break;
        /******************************************/
        /*  Comparison & Bitwise Logic Operations */
        /******************************************/
        case OpCode::LT:
            lt(*this);
            break;
        case OpCode::GT:
            gt(*this);
            break;
        case OpCode::SLT:
            slt(*this);
            break;
        case OpCode::SGT:
            sgt(*this);
            break;
        case OpCode::EQ:
            eq(*this);
            break;
        case OpCode::ISZERO:
            iszero(*this);
            break;
        case OpCode::BITWISE_AND:
            bitwiseAnd(*this);
            break;
        case OpCode::BITWISE_OR:
            bitwiseOr(*this);
            break;
        case OpCode::BITWISE_XOR:
            bitwiseXor(*this);
            break;
        case OpCode::BITWISE_NOT:
            bitwiseNot(*this);
            break;
        case OpCode::BYTE:
            byte(*this);
            break;
        case OpCode::SIGNEXTEND:
            signExtend(*this);
            break;

        /***********************/
        /*  Hashing Operations */
        /***********************/
        case OpCode::HASH:
            hashOp(*this);
            break;

        case OpCode::TYPE:
            typeOp(*this);
            break;

        /***********************************************/
        /*  Stack, Memory, Storage and Flow Operations */
        /***********************************************/
        case OpCode::POP:
            pop(*this);
            break;
        case OpCode::SPUSH:
            spush(*this);
            break;
        case OpCode::RPUSH:
            rpush(*this);
            break;
        case OpCode::RSET:
            rset(*this);
            break;
        case OpCode::JUMP:
            jump(*this);
            break;
        case OpCode::CJUMP:
            cjump(*this);
            break;
        case OpCode::STACKEMPTY:
            stackEmpty(*this);
            break;
        case OpCode::PCPUSH:
            pcPush(*this);
            break;
        case OpCode::AUXPUSH:
            auxPush(*this);
            break;
        case OpCode::AUXPOP:
            auxPop(*this);
            break;
        case OpCode::AUXSTACKEMPTY:
            auxStackEmpty(*this);
            break;
        case OpCode::NOP:
            ++pc;
            break;
        case OpCode::ERRPUSH:
            errPush(*this);
            break;
        case OpCode::ERRSET:
            errSet(*this);
            break;
            /****************************************/
            /*  Duplication and Exchange Operations */
            /****************************************/
        case OpCode::DUP0:
            dup0(*this);
            break;
        case OpCode::DUP1:
            dup1(*this);
            break;
        case OpCode::DUP2:
            dup2(*this);
            break;
        case OpCode::SWAP1:
            swap1(*this);
            break;
        case OpCode::SWAP2:
            swap2(*this);
            break;
            /*********************/
            /*  Tuple Operations */
            /*********************/
        case OpCode::TGET:
            tget(*this);
            break;
        case OpCode::TSET:
            tset(*this);
            break;
        case OpCode::TLEN:
            tlen(*this);
            break;
            /***********************/
            /*  Logging Operations */
            /***********************/
        case OpCode::BREAKPOINT:
            return breakpoint(*this);
        case OpCode::LOG:
            log(*this);
            break;
        case OpCode::DEBUG:
            debug(*this);
            break;
            /**********************/
            /*  System Operations */
            /**********************/
        case OpCode::SEND:
            return send(*this);
        case OpCode::NBSEND:
            nbsend(*this);
            break;
        case OpCode::GETTIME:
            getTime(*this);
            break;
        case OpCode::INBOX:
            return inboxOp(*this);
        case OpCode::ERROR:
            state = Status::Error;
            break;
        case OpCode::HALT:
            state = Status::Halted;
            break;
        default:
            std::cerr << "Unhandled opcode <" << InstructionNames.at(opcode)
                      << ">" << std::hex << static_cast<int>(opcode);
            state = Status::Error;
    }
    return NotBlocked{};
}
