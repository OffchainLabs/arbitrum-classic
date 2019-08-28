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

#include <avm/value.hpp>

struct Operation {
    OpCode opcode;
    std::unique_ptr<value> immediate;

    Operation() = default;
    Operation(OpCode opcode_) : opcode(opcode_) {}
    Operation(OpCode opcode_, value val);

    Operation(const Operation&);
    Operation(Operation&&);
    Operation& operator=(const Operation&);
    Operation& operator=(Operation&&);
    ~Operation();
    void marshal(std::vector<unsigned char>& buf) const;
};

bool operator==(const Operation& val1, const Operation& val2);
bool operator!=(const Operation& val1, const Operation& val2);

struct CodePoint {
    uint64_t pc;
    Operation op;
    uint256_t nextHash;

    CodePoint() {}
    CodePoint(uint64_t pc_, Operation op_, uint256_t nextHash_)
        : pc(pc_), op(op_), nextHash(nextHash_) {}
    void marshal(std::vector<unsigned char>& buf) const;
    bool isSet() {
        return ((op.opcode != static_cast<OpCode>(0)) || (pc != 0) ||
                (nextHash != 0));
    }
};

uint256_t hash(const CodePoint& cp);

bool operator==(const CodePoint& val1, const CodePoint& val2);
std::ostream& operator<<(std::ostream& os, const Operation& val);

#endif /* codepoint_hpp */
