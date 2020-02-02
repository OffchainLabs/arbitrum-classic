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

#include <avm_values/codepoint.hpp>
#include <avm_values/tuple.hpp>
#include <avm_values/util.hpp>
#include <bigint_utils.hpp>

Operation::Operation(OpCode opcode_, const value& immediate_) noexcept
    : opcode(opcode_), immediate(std::make_unique<value>(immediate_)) {}

Operation::Operation(const Operation& op) {
    opcode = op.opcode;
    if (op.immediate) {
        immediate = std::make_unique<value>(*op.immediate);
    }
}
Operation::Operation(Operation&&) noexcept = default;
auto Operation::operator=(const Operation& cp) -> Operation& {
    opcode = cp.opcode;
    if (cp.immediate) {
        immediate = std::make_unique<value>(*cp.immediate);
    } else {
        immediate.reset();
    }
    return *this;
}

auto Operation::operator=(Operation&&) noexcept -> Operation& = default;
Operation::~Operation() = default;

auto operator==(const Operation& val1, const Operation& val2) -> bool {
    if (val1.opcode != val2.opcode) {
        return false;
    }
    if (!val1.immediate && !val2.immediate) {
        return true;
    }
    if (val1.immediate && val2.immediate) {
        return *val1.immediate == *val2.immediate;
    }
    return false;
}

auto operator!=(const Operation& val1, const Operation& val2) -> bool {
    if (val1.opcode != val2.opcode) {
        return true;
    }
    if ((val1.immediate && !val2.immediate) ||
        (!val1.immediate && val2.immediate)) {
        return true;
    }
    if (val1.immediate && val2.immediate) {
        return *val1.immediate != *val2.immediate;
    }
    return false;
}

void Operation::marshal(std::vector<unsigned char>& buf) const {
    if (immediate) {
        buf.push_back(1);
        buf.push_back((uint8_t)opcode);
        marshal_value(*immediate, buf);
    } else {
        buf.push_back(0);
        buf.push_back((uint8_t)opcode);
    }
}

void Operation::marshalShallow(std::vector<unsigned char>& buf) const {
    if (immediate) {
        buf.push_back(1);
        buf.push_back((uint8_t)opcode);
        ::marshalShallow(*immediate, buf);
    } else {
        buf.push_back(0);
        buf.push_back((uint8_t)opcode);
    }
}

void Operation::marshalForProof(std::vector<unsigned char>& buf,
                                bool includeVal) const {
    if (immediate) {
        buf.push_back(1);
        buf.push_back((uint8_t)opcode);
        if (includeVal) {
            ::marshalShallow(*immediate, buf);
        } else {
            buf.push_back(HASH_ONLY);
            std::array<unsigned char, 32> tmpbuf;
            to_big_endian(::hash(*immediate), tmpbuf.begin());
            buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
        }
    } else {
        buf.push_back(0);
        buf.push_back((uint8_t)opcode);
    }
}

uint64_t pc_default = -1;

auto operator==(const CodePoint& val1, const CodePoint& val2) -> bool {
    if (val1.pc != val2.pc)
        return false;
    else
        return true;
}

void CodePoint::marshal(std::vector<unsigned char>& buf) const {
    buf.push_back(CODEPT);
    uint64_t bepc = boost::endian::native_to_big(pc);
    std::copy(
        static_cast<const char*>(static_cast<const void*>(&bepc)),
        static_cast<const char*>(static_cast<const void*>(&bepc)) + sizeof bepc,
        std::back_inserter(buf));
    op.marshal(buf);
    std::array<unsigned char, 32> val;
    to_big_endian(nextHash, val.begin());
    buf.insert(buf.end(), val.begin(), val.end());
}

auto hash(const CodePoint& cp) -> uint256_t {
    std::array<uint64_t, 4> nextHashInts;
    to_big_endian(cp.nextHash, nextHashInts.begin());
    if (cp.op.immediate) {
        std::array<unsigned char, 66> valData;
        valData[0] = CODEPT;
        valData[1] = static_cast<unsigned char>(cp.op.opcode);
        auto immHash = ::hash(*cp.op.immediate);
        std::array<uint64_t, 4> valHashInts;
        to_big_endian(immHash, valHashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(valHashInts.data()),
                  reinterpret_cast<unsigned char*>(valHashInts.data()) + 32,
                  valData.begin() + 2);
        std::copy(reinterpret_cast<unsigned char*>(nextHashInts.data()),
                  reinterpret_cast<unsigned char*>(nextHashInts.data()) + 32,
                  valData.end() - 32);
        std::array<unsigned char, 32> hashData;
        evm::Keccak_256(valData.data(), valData.size(), hashData.data());
        return from_big_endian(hashData.begin(), hashData.end());
    } else {
        std::array<unsigned char, 34> valData;
        valData[0] = CODEPT;
        valData[1] = static_cast<unsigned char>(cp.op.opcode);
        std::copy(reinterpret_cast<unsigned char*>(nextHashInts.data()),
                  reinterpret_cast<unsigned char*>(nextHashInts.data()) + 32,
                  valData.end() - 32);
        std::array<unsigned char, 32> hashData;
        evm::Keccak_256(valData.data(), valData.size(), hashData.data());
        return from_big_endian(hashData.begin(), hashData.end());
    }
}

auto operator<<(std::ostream& os, const Operation& val) -> std::ostream& {
    if (val.immediate) {
        os << "Immediate(" << InstructionNames.at(val.opcode) << ", "
           << *val.immediate << ")";
    } else {
        os << "Basic(" << InstructionNames.at(val.opcode) << ")";
    }
    return os;
}

auto getErrCodePoint() -> CodePoint {
    return CodePoint(pc_default, Operation(static_cast<OpCode>(0)), 0);
}

auto opsToCodePoints(std::vector<Operation>&& ops) -> std::vector<CodePoint> {
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
