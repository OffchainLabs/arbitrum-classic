/*
 * Copyright 2020, Offchain Labs, Inc.
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

#ifndef codepointstub_hpp
#define codepointstub_hpp

#include <avm_values/bigint.hpp>
#include <avm_values/opcodes.hpp>

#include <boost/smart_ptr/intrusive_ptr.hpp>

#include <cassert>
#include <shared_mutex>

struct CodePoint;
class Value;
class CodeSegment;

struct CodePointRef {
    // Must go first for Value union repr
    uint64_t pc;
    // either a CodeSegment or an unloaded value of such
    boost::intrusive_ptr<CodeSegment> root;

    CodePointRef(boost::intrusive_ptr<CodeSegment> root_, uint64_t pc_)
        : pc(pc_), root(root_) {
        assert(pc < (uint64_t(1) << 62));
    }

    CodePointRef& operator++() {
        --pc;
        return *this;
    }

    CodePointRef operator+(uint64_t i) const { return {root, pc - i}; }
    CodePointRef operator-(uint64_t i) const { return {root, pc + i}; }
};

std::ostream& operator<<(std::ostream& os, const CodePointRef& code);

struct CodePointStub {
    CodePointRef pc;
    uint256_t hash;

    CodePointStub(const CodePointRef& pc, const CodePoint& cp);
    CodePointStub(const CodePointRef& pc_, uint256_t hash_);

    friend bool operator==(const CodePointStub& val1,
                           const CodePointStub& val2) {
        return val1.hash == val2.hash;
    }

    friend bool operator!=(const CodePointStub& val1,
                           const CodePointStub& val2) {
        return val1.hash != val2.hash;
    }

    [[nodiscard]] bool is_error() const;
};

inline uint256_t hash(const CodePointStub& cp) {
    return cp.hash;
}

#endif /* codepointstub_hpp */
