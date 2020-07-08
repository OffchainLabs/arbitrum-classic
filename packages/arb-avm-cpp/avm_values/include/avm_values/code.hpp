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

class Code;
class Transaction;
struct LoadedExecutable;

// The public interface of CodeSegment is thread safe assuming that the indexes
// used to access the segment are less than or equal to the size of the segment
// when you initially load it
class CodeSegment {
    uint64_t segment;
    std::vector<CodePoint> code;

    friend class Code;
    friend LoadedExecutable loadExecutable(const std::string& contract_filename,
                                           TuplePool& pool);

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

    uint64_t segmentID() const { return segment; }

    const CodePoint& operator[](uint64_t pc) const { return code.at(pc); }

    const CodePoint& at(uint64_t pc) const { return code.at(pc); }

    friend std::ostream& operator<<(std::ostream& os, const CodeSegment& code);
};

struct CodeSegmentSnapshot {
    std::shared_ptr<const CodeSegment> segment;
    uint64_t op_count;
};

class Code {
    mutable std::mutex mutex;
    std::unordered_map<uint64_t, std::shared_ptr<CodeSegment>> segments;
    uint64_t next_segment_num;

   public:
    Code() : next_segment_num(0) { addSegment(); }
    Code(std::shared_ptr<CodeSegment> segment) : next_segment_num(1) {
        assert(segment->segmentID() == 0);
        segments[0] = std::move(segment);
    }

    std::shared_ptr<const CodeSegment> loadCodeSegment(
        uint64_t segment_num) const {
        mutex.lock();
        std::shared_ptr<const CodeSegment> segment = segments.at(segment_num);
        mutex.unlock();
        return segment;
    }

    bool containsSegment(uint64_t segment_id) const {
        const std::lock_guard<std::mutex> lock(mutex);
        return segments.find(segment_id) != segments.end();
    }

    void restoreExistingSegment(std::shared_ptr<CodeSegment> segment) {
        const std::lock_guard<std::mutex> lock(mutex);
        uint64_t segment_id = segment->segmentID();
        if (segments.find(segment->segmentID()) == segments.end()) {
            segments[segment_id] = std::move(segment);
        }
    }

    std::vector<CodeSegmentSnapshot> snapshot() const {
        const std::lock_guard<std::mutex> lock(mutex);
        std::vector<CodeSegmentSnapshot> ret;
        for (const auto& key_val : segments) {
            ret.push_back({key_val.second, key_val.second->size()});
        }
        return ret;
    }

    const CodePoint& loadCodePoint(const CodePointRef& ref) const {
        const std::lock_guard<std::mutex> lock(mutex);
        return (*segments.at(ref.segment))[ref.pc];
    }

    CodePointStub addSegment() {
        const std::lock_guard<std::mutex> lock(mutex);
        uint64_t segment_num = next_segment_num++;
        auto new_segment = std::make_shared<CodeSegment>(segment_num);
        CodePointStub stub{{segment_num, 0}, hash((*new_segment)[0])};
        segments[segment_num] = std::move(new_segment);
        return stub;
    }

    CodePointStub addOperation(const CodePointRef& ref, Operation op) {
        auto& segment = segments[ref.segment];
        auto initial_pc = segment->size() - 1;
        if (ref.pc == initial_pc) {
            // This is the first pc in the segment so we can append directly
            if (segment->size() == segment->capacity() &&
                segment.use_count() > 1) {
                // Segment is full, so we must allocate a new segment
                segments[ref.segment] = segment->clone();
                segment = segments[ref.segment];
            }
            return segment->addOperation(std::move(op));
        } else {
            // This segment was already mutated elsewhere, therefore we must
            // make a copy
            uint64_t new_segment_num = next_segment_num++;
            auto new_segment = segment->getSubset(new_segment_num, ref.pc);
            auto stub = new_segment->addOperation(std::move(op));
            segments[new_segment_num] = std::move(new_segment);
            return stub;
        }
    }

    CodePointRef initialCodePointRef() const {
        const std::lock_guard<std::mutex> lock(mutex);
        return {0, segments.at(0)->size() - 1};
    }

    friend std::ostream& operator<<(std::ostream& os, const Code& code);
};

std::ostream& operator<<(std::ostream& os, const Code& code);

#endif /* code_hpp */
