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

#include <avm_values/codepoint.hpp>

#include <avm_values/tuple.hpp>

#include <ethash/keccak.hpp>

#include <iostream>

Operation::Operation(OpCode opcode_, value immediate_)
    : opcode(opcode_), immediate(std::make_unique<value>(immediate_)) {}

void Operation::marshalForProof(std::vector<unsigned char>& buf,
                                MarshalLevel marshal_level,
                                const Code& code) const {
    if (immediate) {
        buf.push_back(1);
        buf.push_back(static_cast<uint8_t>(opcode));
        ::marshalForProof(*immediate, marshal_level, buf, code);
    } else {
        buf.push_back(0);
        buf.push_back(static_cast<uint8_t>(opcode));
    }
}

bool operator==(const Operation& val1, const Operation& val2) {
    if (val1.opcode != val2.opcode) {
        return false;
    }
    if (!val1.immediate && !val2.immediate) {
        return true;
    }
    if (!val1.immediate || !val2.immediate) {
        return false;
    }
    return *val1.immediate == *val2.immediate;
}

bool operator!=(const Operation& val1, const Operation& val2) {
    if (val1.opcode != val2.opcode) {
        return true;
    }
    if (!val1.immediate && !val2.immediate) {
        return false;
    }
    if (!val1.immediate || !val2.immediate) {
        return true;
    }
    return *val1.immediate != *val2.immediate;
}

uint64_t pc_default = -1;

bool operator==(const CodePoint& val1, const CodePoint& val2) {
    if (hash(val1) != hash(val2)) {
        return false;
    } else {
        return true;
    }
}

uint256_t hash(const CodePoint& cp) {
    if (cp.op.immediate) {
        std::array<unsigned char, 66> valData;
        valData[0] = CODEPT;
        valData[1] = static_cast<unsigned char>(cp.op.opcode);
        auto immHash = hash_value(*cp.op.immediate);
        auto it = valData.begin() + 2;
        it = to_big_endian(immHash, it);
        to_big_endian(cp.nextHash, it);
        auto hash_val = ethash::keccak256(valData.data(), valData.size());
        return intx::be::load<uint256_t>(hash_val);
    } else {
        std::array<unsigned char, 34> valData;
        valData[0] = CODEPT;
        valData[1] = static_cast<unsigned char>(cp.op.opcode);
        to_big_endian(cp.nextHash, valData.begin() + 2);
        auto hash_val = ethash::keccak256(valData.data(), valData.size());
        return intx::be::load<uint256_t>(hash_val);
    }
}

std::ostream& operator<<(std::ostream& os, const Operation& val) {
    auto opcode_name = [&]() -> std::string {
        auto it = InstructionNames.find(val.opcode);
        if (it != InstructionNames.end()) {
            return it->second;
        } else {
            return "InvalidOpcode";
        }
    }();

    if (val.immediate) {
        os << "Immediate(" << opcode_name << ", " << *val.immediate << ")";
    } else {
        os << "Basic(" << opcode_name << ")";
    }
    return os;
}

std::ostream& operator<<(std::ostream& os, const CodePoint& val) {
    os << "CodePoint(" << val.op << ", " << intx::to_string(val.nextHash, 16)
       << ")";
    return os;
}

const Operation& getErrOperation() {
    Operation static errop(static_cast<OpCode>(0));
    return errop;
}

const CodePoint& getErrCodePoint() {
    CodePoint static errcp({static_cast<OpCode>(0)}, 0);
    return errcp;
}

const uint256_t& getErrCodePointHash() {
    uint256_t static errpc_hash = hash(getErrCodePoint());
    return errpc_hash;
}
