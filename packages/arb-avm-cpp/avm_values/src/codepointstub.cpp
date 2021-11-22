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

#include <avm_values/codepoint.hpp>
#include <avm_values/codepointstub.hpp>
#include <avm_values/value.hpp>

#include <iostream>

CodePointStub::CodePointStub(const CodePointRef& pc_, const CodePoint& cp_)
    : pc(pc_), hash(::hash(cp_)) {}

CodePointStub::CodePointStub(const CodePointRef& pc_, uint256_t hash_)
    : pc(pc_), hash(hash_) {}

std::ostream& operator<<(std::ostream& os, const CodePointRef& cpr) {
    os << "(" << cpr.segment << ", " << cpr.pc << ")";
    return os;
}

void CodePointRef::marshal(std::vector<unsigned char>& buf) const {
    marshal_uint64_t(segment, buf);
    marshal_uint64_t(pc, buf);
}

void CodePointStub::marshal(std::vector<unsigned char>& buf) const {
    pc.marshal(buf);
    marshal_uint256_t(hash, buf);
}

bool CodePointStub::is_error() const {
    return hash == getErrCodePointHash();
}
