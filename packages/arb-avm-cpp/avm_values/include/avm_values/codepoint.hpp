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

#include <avm_values/opcodes.hpp>
#include <avm_values/value.hpp>

struct Operation {
    OpCode opcode;
    std::unique_ptr<value> immediate;

    Operation() { opcode = OpCode::DEFAULT; };
    Operation(OpCode opcode_) : opcode(opcode_) {}
    Operation(OpCode opcode_, value val);

    Operation(const Operation&);
    Operation(Operation&&);
    Operation& operator=(const Operation&);
    Operation& operator=(Operation&&);
    ~Operation();
    void marshal(std::vector<unsigned char>& buf, const Code& code) const;
    void marshalForProof(std::vector<unsigned char>& buf,
                         bool includeVal,
                         const Code& code) const;
};

bool operator==(const Operation& val1, const Operation& val2);
bool operator!=(const Operation& val1, const Operation& val2);

std::ostream& operator<<(std::ostream& os, const Operation& val);

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
        : pc(pc_), op(op_), nextHash(nextHash_) {}
    void marshal(std::vector<unsigned char>& buf, const Code& code) const;
    bool isSet() {
        return ((op.opcode != static_cast<OpCode>(0)) || (pc != 0) ||
                (nextHash != 0));
    }
    int getSize() const { return 1; }

    bool isError() const {
        return nextHash == 0 && op == Operation{static_cast<OpCode>(0)};
    }
};

std::ostream& operator<<(std::ostream& os, const CodePoint& val);

bool operator==(const CodePoint& val1, const CodePoint& val2);

uint256_t hash(const CodePoint& cp);

const CodePoint& getErrCodePoint();

struct CodePointStub {
    uint64_t pc;
    uint256_t hash;

    CodePointStub() = default;
    CodePointStub(const CodePoint& cp) : pc(cp.pc), hash(::hash(cp)) {}
    CodePointStub(uint64_t pc_, uint256_t hash_) : pc(pc_), hash(hash_) {}

    friend bool operator==(const CodePointStub& val1,
                           const CodePointStub& val2) {
        return val1.pc == val2.pc && val1.hash == val2.hash;
    }
};

inline uint256_t hash(const CodePointStub& cp) {
    return cp.hash;
}

struct CodePointRef {
    uint64_t pc;
    bool is_err;

    CodePointRef() = default;
    CodePointRef(uint64_t pc_, bool is_err_) : pc(pc_), is_err(is_err_) {}
    CodePointRef(const CodePointStub& stub)
        : pc(stub.pc), is_err(hash(stub) == hash(getErrCodePoint())) {}

    CodePointRef& operator=(uint64_t pc_) {
        pc = pc_;
        is_err = false;
        return *this;
    }

    CodePointRef& operator++() {
        ++pc;
        return *this;
    }

    CodePointRef operator+(uint64_t i) { return {pc + i, is_err}; }

    bool operator==(uint64_t val) const { return pc == val && is_err == false; }

    friend bool operator==(CodePointRef val1, CodePointRef val2) {
        if (!val1.is_err && !val2.is_err && val1.pc == val2.pc) {
            return true;
        }
        return val1.is_err && val2.is_err;
    }
};

std::vector<CodePoint> opsToCodePoints(const std::vector<Operation>& ops);

class Code {
    std::vector<CodePoint> code;

   public:
    Code() = default;
    Code(std::vector<CodePoint> code_);

    const CodePoint& operator[](const CodePointStub& ref) const {
        const auto& err_codepoint = getErrCodePoint();
        if (ref.hash == hash(err_codepoint)) {
            return err_codepoint;
        } else {
            return code[ref.pc];
        }
    }

    const CodePoint& operator[](CodePointRef ref) const {
        if (ref.is_err) {
            return getErrCodePoint();
        } else {
            return code[ref.pc];
        }
    }

    const CodePoint& operator[](uint64_t pos) const { return code[pos]; }

    friend std::ostream& operator<<(std::ostream& os, const Code& code);
};

std::ostream& operator<<(std::ostream& os, const Code& code);

#endif /* codepoint_hpp */
