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

#ifndef code_hpp
#define code_hpp

#include <avm_values/codepoint.hpp>

class Code {
    std::vector<CodePoint> code;

   public:
    Code() { code.push_back(getErrCodePoint()); }

    const CodePoint& operator[](const CodePointRef& ref) const {
        if (ref.is_err) {
            return getErrCodePoint();
        } else {
            return code[ref.pc];
        }
    }

    const CodePoint& at(const CodePointRef& ref) const {
        if (ref.is_err) {
            return getErrCodePoint();
        } else {
            return code.at(ref.pc);
        }
    }

    void addOperation(Operation op) {
        uint256_t prev_hash = 0;
        if (code.size() > 0) {
            prev_hash = hash(code.back());
        }
        code.emplace_back(std::move(op), prev_hash);
    }

    CodePointRef initialCodePointRef() const {
        return {0, code.size() - 1, false};
    }

    friend std::ostream& operator<<(std::ostream& os, const Code& code);
};

std::ostream& operator<<(std::ostream& os, const Code& code);

#endif /* code_hpp */
