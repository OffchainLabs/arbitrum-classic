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
    uint64_t segment;
    std::vector<CodePoint> code;

    CodeSegment(uint64_t segment_, std::vector<CodePoint> code_)
        : segment(segment_), code(std::move(code_)) {}

   public:
    CodeSegment(uint64_t segment_) : segment(segment_) {
        code.push_back(getErrCodePoint());
    }

    size_t size() const { return code.size(); }

    const CodePoint& operator[](uint64_t pc) const { return code[pc]; }

    const CodePoint& at(uint64_t pc) const { return code.at(pc); }

    CodePointStub addOperation(Operation op) {
        uint256_t prev_hash = 0;
        if (code.size() > 0) {
            prev_hash = hash(code.back());
        }
        code.emplace_back(std::move(op), prev_hash);
        return {{segment, code.size() - 1}, hash(code.back())};
    }

    // Return the subset of this code segment starting in the given pc
    CodeSegment getSubset(uint64_t new_segment, uint64_t pc) const {
        return {new_segment,
                std::vector<CodePoint>{code.begin(), code.begin() + pc}};
    }

    friend std::ostream& operator<<(std::ostream& os, const CodeSegment& code);
};

class Code {
    std::vector<std::shared_ptr<CodeSegment>> segments;

   public:

    const CodePoint& loadCodePoint(const CodePointRef& ref) const {
        const std::lock_guard<std::mutex> lock(mutex);
        return (*segments[ref.segment])[ref.pc];
    }

    CodePointStub addSegment() {
        uint64_t segment_num = segments.size();
        auto new_segment = std::make_shared<CodeSegment>(segment_num);
        segments.emplace_back(std::move(new_segment));
        auto& segment = *segments.back();
        return {{segment_num, 0}, hash(segment[0])};
    }

    CodePointStub addOperation(const CodePointRef& ref, Operation op) {
        auto& segment = segments[ref.segment];

        // This is the first pc in the segment so we can append directly
        if (ref.pc == segment.size() - 1) {
            return segment.addOperation(std::move(op));
        } else {
            // This segment was already mutated elsewhere, therefore we must
            // make a copy
            uint64_t new_segment_num = segments.size();
            auto new_segment = segment.getSubset(new_segment_num, ref.pc);
            segments.push_back(new_segment);
            return segments.back().addOperation(std::move(op));
        }
    }

    CodePointRef initialCodePointRef() const {
        const std::lock_guard<std::mutex> lock(mutex);
        return {0, segments[0]->size() - 1};
    }

    friend std::ostream& operator<<(std::ostream& os, const Code& code);
};

std::ostream& operator<<(std::ostream& os, const Code& code);

#endif /* code_hpp */
