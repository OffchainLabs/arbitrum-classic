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

#include <memory>
#include <mutex>
#include <unordered_map>
#include <vector>

class Code;
class Transaction;
struct LoadedExecutable;

// The public interface of CodeSegment is thread safe assuming that the indexes
// used to access the segment are less than or equal to the size of the segment
// when you initially load it
class CodeSegment {
    uint64_t segment_id;
    std::vector<CodePoint> code;

    friend class Code;
    friend LoadedExecutable loadExecutable(
        const std::string& executable_filename);

    size_t capacity() const { return code.capacity(); }

    size_t size() const { return code.size(); }

    CodePointStub addOperation(Operation op) {
        uint256_t prev_hash = 0;
        if (code.size() > 0) {
            prev_hash = hash(code.back());
        }
        code.emplace_back(std::move(op), prev_hash);
        return codePointStub(code.size() - 1);
    }

    // Return the subset of this code segment starting in the given pc
    std::shared_ptr<CodeSegment> getSubset(uint64_t new_segment_id,
                                           uint64_t pc) const {
        return std::make_shared<CodeSegment>(
            new_segment_id,
            std::vector<CodePoint>{code.begin(), code.begin() + pc});
    }

   public:
    CodeSegment(uint64_t segment_id_) : segment_id(segment_id_) {
        code.push_back(getErrCodePoint());
    }
    CodeSegment(uint64_t segment_id_, std::vector<CodePoint> code_)
        : segment_id(segment_id_), code(std::move(code_)) {}

    CodeSegment(uint64_t segment_id_, std::vector<Operation> ops)
        : CodeSegment(segment_id_) {
        for (auto it = ops.rbegin(); it != ops.rend(); ++it) {
            addOperation(std::move(*it));
        }
    }

    uint64_t segmentID() const { return segment_id; }

    const CodePoint& operator[](uint64_t pc) const { return code.at(pc); }

    const CodePoint& at(uint64_t pc) const { return code.at(pc); }

    friend std::ostream& operator<<(std::ostream& os, const CodeSegment& code);

    CodePointStub codePointStub(uint64_t pc) const {
        return {{segment_id, pc}, hash(code.at(pc))};
    }
};

struct CodeSegmentSnapshot {
    std::shared_ptr<const CodeSegment> segment;
    uint64_t op_count;
};

struct CodeSnapshot {
    std::unordered_map<uint64_t, CodeSegmentSnapshot> segments;
    uint64_t op_count;
};

class Code {
    mutable std::mutex mutex;
    std::unordered_map<uint64_t, std::shared_ptr<CodeSegment>> segments;
    uint64_t next_segment_num;

   public:
    Code() : Code(0) {}
    Code(uint64_t next_segment_num_) : next_segment_num(next_segment_num_) {}

    CodeSegmentSnapshot loadCodeSegment(uint64_t segment_num) const {
        const std::lock_guard<std::mutex> lock(mutex);
        auto& segment = segments.at(segment_num);
        return {segment, segment->size()};
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

    CodeSnapshot snapshot() const {
        const std::lock_guard<std::mutex> lock(mutex);
        std::unordered_map<uint64_t, CodeSegmentSnapshot> copied_segments;
        for (const auto& key_val : segments) {
            copied_segments[key_val.first] = {key_val.second,
                                              key_val.second->size()};
        }
        return {std::move(copied_segments), next_segment_num};
    }

    const CodePoint& loadCodePoint(const CodePointRef& ref) const {
        const std::lock_guard<std::mutex> lock(mutex);
        return (*segments.at(ref.segment))[ref.pc];
    }

    CodePointStub addSegment() {
        const std::lock_guard<std::mutex> lock(mutex);
        uint64_t segment_num = next_segment_num++;
        auto new_segment = std::make_shared<CodeSegment>(segment_num);
        auto stub = new_segment->codePointStub(0);
        segments[segment_num] = std::move(new_segment);
        return stub;
    }

    void addSegment(std::shared_ptr<CodeSegment> segment) {
        const std::lock_guard<std::mutex> lock(mutex);
        assert(segment->segmentID() == next_segment_num);
        segments[next_segment_num] = std::move(segment);
        next_segment_num++;
    }

    CodePointStub addOperation(const CodePointRef& ref, Operation op) {
        const std::lock_guard<std::mutex> lock(mutex);
        auto& segment = segments[ref.segment];
        auto initial_pc = segment->size() - 1;
        if (ref.pc == initial_pc) {
            if (segment.use_count() == 1) {
                // No other code has a reference to this segment. That means we
                // can modify the segment even if it forces a reallocation
                return segment->addOperation(std::move(op));
            }

            if (segment->size() < segment->capacity()) {
                // The segment has extra capacity so we can add to it even if
                // other code has references to it
                return segment->addOperation(std::move(op));
            }

            // Fall back to making a copy as there are other references and no
            // space to add this operation
        }
        // This segment was already mutated elsewhere, therefore we must
        // make a copy
        uint64_t new_segment_num = next_segment_num++;
        auto new_segment = segment->getSubset(new_segment_num, ref.pc);
        auto stub = new_segment->addOperation(std::move(op));
        segments[new_segment_num] = std::move(new_segment);
        return stub;
    }

    CodePointRef initialCodePointRef() const {
        const std::lock_guard<std::mutex> lock(mutex);
        return {0, segments.at(0)->size() - 1};
    }
};

#endif /* code_hpp */
