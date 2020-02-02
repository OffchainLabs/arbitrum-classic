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

#ifndef codepoint_hpp
#define codepoint_hpp

#include <avm_values/opcodes.hpp>
#include <avm_values/value.hpp>
#include <utility>

struct Operation {
    OpCode opcode;
    std::unique_ptr<value> immediate;

    Operation() noexcept { opcode = OpCode::DEFAULT; };
    explicit Operation(OpCode opcode_) noexcept : opcode(opcode_) {}
    Operation(OpCode opcode_, const value& val) noexcept;

    Operation(const Operation&);
    Operation(Operation&&) noexcept;
    auto operator=(const Operation&) -> Operation&;
    auto operator=(Operation&&) noexcept -> Operation&;
    ~Operation();
    void marshal(std::vector<unsigned char>& buf) const;
    void marshalShallow(std::vector<unsigned char>& buf) const;
    void marshalForProof(std::vector<unsigned char>& buf,
                         bool includeVal) const;
};

auto operator==(const Operation& val1, const Operation& val2) -> bool;
auto operator!=(const Operation& val1, const Operation& val2) -> bool;

extern uint64_t pc_default;

struct CodePoint {
    uint64_t pc;
    Operation op;
    uint256_t nextHash;

    CodePoint() {
        pc = pc_default;
        nextHash = 0;
    }
    CodePoint(uint64_t pc_, Operation op_, uint256_t nextHash_)
        : pc(pc_), op(std::move(op_)), nextHash(std::move(nextHash_)) {}
    void marshal(std::vector<unsigned char>& buf) const;
    auto isSet() -> bool {
        return ((op.opcode != static_cast<OpCode>(0)) || (pc != 0) ||
                (nextHash != 0));
    }
};

auto getErrCodePoint() -> CodePoint;

auto hash(const CodePoint& cp) -> uint256_t;

auto operator==(const CodePoint& val1, const CodePoint& val2) -> bool;
auto operator<<(std::ostream& os, const Operation& val) -> std::ostream&;

auto opsToCodePoints(std::vector<Operation>&& ops) -> std::vector<CodePoint>;

#endif /* codepoint_hpp */
