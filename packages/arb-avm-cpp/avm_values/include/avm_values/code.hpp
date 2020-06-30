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

class CodeSegment {
    std::vector<CodePoint> code;

    CodeSegment(std::vector<CodePoint> code_) : code(std::move(code_)) {}

   public:
    CodeSegment() { code.push_back(getErrCodePoint()); }

    size_t size() const { return code.size(); }

    const CodePoint& operator[](uint64_t pc) const { return code[pc]; }

    const CodePoint& at(uint64_t pc) const { return code.at(pc); }

    void addOperation(Operation op) {
        uint256_t prev_hash = 0;
        if (code.size() > 0) {
            prev_hash = hash(code.back());
        }
        code.emplace_back(std::move(op), prev_hash);
    }

    // Return the subset of this code segment starting in the given pc
    CodeSegment getSubset(uint64_t pc) const {
        return {std::vector<CodePoint>{code.begin(), code.begin() + pc}};
    }

    friend std::ostream& operator<<(std::ostream& os, const CodeSegment& code);
};

class Code {
    std::vector<CodeSegment> segments;

   public:
    const CodePoint& operator[](const CodePointRef& ref) const {
        if (ref.is_err) {
            return getErrCodePoint();
        } else {
            return segments[ref.segment][ref.pc];
        }
    }

    const CodePoint& at(const CodePointRef& ref) const {
        if (ref.is_err) {
            return getErrCodePoint();
        } else {
            return segments.at(ref.segment).at(ref.pc);
        }
    }

    CodePointRef addSegment() {
        segments.emplace_back();
        return {segments.size() - 1, 0, false};
    }

    CodePointRef addOperation(const CodePointRef& ref, Operation op) {
        auto& segment = segments[ref.segment];

        // This is the first pc in the segment so we can append directly
        if (ref.pc == segment.size() - 1) {
            segment.addOperation(std::move(op));
            return {ref.segment, ref.pc + 1, false};
        } else {
            // This segment was already mutated elsewhere, therefore we must
            // make a copy
            auto new_segment = segment.getSubset(ref.pc);
            new_segment.addOperation(std::move(op));
            segments.push_back(new_segment);
            return {segments.size() - 1, ref.pc + 1, false};
        }
    }

    CodePointRef initialCodePointRef() const {
        return {0, segments[0].size() - 1, false};
    }

    friend std::ostream& operator<<(std::ostream& os, const Code& code);
};

std::ostream& operator<<(std::ostream& os, const Code& code);

#endif /* code_hpp */
