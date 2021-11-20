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

#ifndef codepoint_hpp
#define codepoint_hpp

#include <avm_values/buffer.hpp>
#include <avm_values/opcodes.hpp>
#include <avm_values/tuple.hpp>
#include <avm_values/value.hpp>

#include <optional>

struct Operation {
    OpCode opcode;
    std::unique_ptr<Value> immediate;

    explicit Operation(OpCode opcode_) : opcode(opcode_) {}
    Operation(OpCode opcode_, Value val);

    Operation(const Operation& op) : opcode(op.opcode) {
        if (op.immediate) {
            immediate = std::make_unique<Value>(*op.immediate);
        }
    }

    Operation(Operation&& op) = default;

    void marshalForProof(std::vector<unsigned char>& buf,
                         size_t marshal_level,
                         const Code& code) const;
};

bool operator==(const Operation& val1, const Operation& val2);
bool operator!=(const Operation& val1, const Operation& val2);

std::ostream& operator<<(std::ostream& os, const Operation& val);

struct CodePoint {
    Operation op;
    uint256_t nextHash;

    CodePoint(Operation op_, uint256_t nextHash_)
        : op(std::move(op_)), nextHash(nextHash_) {}

    [[nodiscard]] bool isError() const {
        return nextHash == 0 && op == Operation{static_cast<OpCode>(0)};
    }
};

std::ostream& operator<<(std::ostream& os, const CodePoint& val);

bool operator==(const CodePoint& val1, const CodePoint& val2);

uint256_t hash(const CodePoint& cp);

const Operation& getErrOperation();
const CodePoint& getErrCodePoint();
const uint256_t& getErrCodePointHash();

#endif /* codepoint_hpp */
