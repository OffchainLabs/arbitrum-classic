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

#include <cassert>
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
    static_assert(std::is_nothrow_move_constructible<Operation>::value,
                  "Operation should be noexcept MoveConstructible");
    static_assert(std::is_nothrow_move_constructible<uint256_t>::value,
                  "uint256_t should be noexcept MoveConstructible");
    std::vector<Operation> operations;
    std::vector<uint256_t> cached_hashes;
    uint256_t prev_hash;

    friend class Code;
    friend class CoreCode;

    friend LoadedExecutable loadExecutable(
        const std::string& executable_filename);

    size_t capacity() const { return operations.capacity(); }

    size_t size() const { return operations.size(); }

    CodePointStub addOperation(Operation op) {
        CodePoint cp{std::move(op), prev_hash};
        prev_hash = hash(cp);
        operations.push_back(std::move(cp.op));
        if (operations.size() % 10 == 0) {
            cached_hashes.push_back(prev_hash);
        }
        return {{segment_id, operations.size() - 1}, prev_hash};
    }

    // Return the subset of this code segment starting in the given pc
    std::shared_ptr<CodeSegment> getSubset(uint64_t new_segment_id,
                                           uint64_t pc) const {
        auto ops = static_cast<std::vector<CodePoint>::difference_type>(pc);
        std::vector<Operation> copy_copy;
        auto code_copy = std::vector<Operation>{operations.begin(),
                                                operations.begin() + ops + 1};
        auto hashes_copy = std::vector<uint256_t>{
            cached_hashes.begin(), cached_hashes.begin() + (ops + 1) / 10};
        // Make endpoint pc + 1 since pc should be included in segment
        return std::make_shared<CodeSegment>(
            new_segment_id, std::move(code_copy), std::move(hashes_copy));
    }

   public:
    CodeSegment(uint64_t segment_id_) : segment_id(segment_id_) {
        addOperation(getErrOperation());
    }

    CodeSegment(uint64_t segment_id_,
                std::vector<Operation> operations_,
                std::vector<uint256_t> next_hashes_)
        : segment_id(segment_id_),
          operations(std::move(operations_)),
          cached_hashes(std::move(next_hashes_)) {
        prev_hash = ::hash(loadCodePoint(operations.size() - 1));
    }

    CodeSegment(uint64_t segment_id_, std::vector<Operation> ops)
        : CodeSegment(segment_id_) {
        for (auto it = ops.rbegin(); it != ops.rend(); ++it) {
            addOperation(std::move(*it));
        }
    }

    uint64_t segmentID() const { return segment_id; }

    CodePoint loadCodePoint(uint64_t pc) const {
        uint256_t prev_hash;

        if (pc / 10 > 0) {
            prev_hash = cached_hashes.at(pc / 10 - 1);
        }
        for (uint64_t i = (pc / 10) * 10; i < pc; i++) {
            prev_hash = hash(CodePoint(operations[i], prev_hash));
        }
        return {operations[pc], prev_hash};
    }

    const Operation& loadOperation(uint64_t pc) const { return operations[pc]; }
    const uint256_t& loadCachedHash(uint64_t i) const {
        return cached_hashes[i];
    }

    //    const CodePoint& operator[](uint64_t pc) const { return code.at(pc); }
    //
    //    const CodePoint& at(uint64_t pc) const { return code.at(pc); }

    friend std::ostream& operator<<(std::ostream& os, const CodeSegment& code);

    CodePointStub initialCodePointStub() const {
        return {{segment_id, 0}, getErrCodePointHash()};
    }

    void reserve(size_t size) {
        operations.reserve(size);
        cached_hashes.reserve(size / 10);
    }
};

struct CodeSegmentSnapshot {
    std::shared_ptr<const CodeSegment> segment;
    uint64_t op_count;
    uint64_t cached_hash_count;
};

struct CodeSnapshot {
    std::unordered_map<uint64_t, CodeSegmentSnapshot> segments;
    uint64_t op_count;
};

class Code {
   public:
    virtual ~Code() = default;
    virtual CodeSegmentSnapshot loadCodeSegment(uint64_t segment_num) const;

    virtual CodePoint loadCodePoint(const CodePointRef& ref) const;

    virtual CodePointStub addSegment();

    virtual CodePointStub addOperation(const CodePointRef& ref, Operation op);
};

class CoreCode : public Code {
    mutable std::mutex mutex;
    std::unordered_map<uint64_t, std::shared_ptr<CodeSegment>> segments;
    uint64_t next_segment_num;

   public:
    CoreCode() : CoreCode(0) {}
    CoreCode(uint64_t next_segment_num_)
        : next_segment_num(next_segment_num_) {}

    CodeSegmentSnapshot loadCodeSegment(uint64_t segment_num) const {
        const std::lock_guard<std::mutex> lock(mutex);
        auto& segment = segments.at(segment_num);
        return {segment, segment->size(), segment->cached_hashes.size()};
    }

    bool containsSegment(uint64_t segment_id) const {
        const std::lock_guard<std::mutex> lock(mutex);
        return segments.find(segment_id) != segments.end();
    }

    void restoreExistingSegment(std::shared_ptr<CodeSegment> segment) {
        const std::lock_guard<std::mutex> lock(mutex);
        uint64_t segment_id = segment->segmentID();
        if (segment_id >= next_segment_num) {
            throw std::runtime_error("code segment loaded incorrectly");
        }
        if (segments.find(segment->segmentID()) == segments.end()) {
            segments[segment_id] = std::move(segment);
        }
    }

    CodeSnapshot snapshot() const {
        const std::lock_guard<std::mutex> lock(mutex);
        std::unordered_map<uint64_t, CodeSegmentSnapshot> copied_segments;
        for (const auto& key_val : segments) {
            copied_segments[key_val.first] = {
                key_val.second, key_val.second->size(),
                key_val.second->cached_hashes.size()};
        }
        return {std::move(copied_segments), next_segment_num};
    }

    CodePoint loadCodePoint(const CodePointRef& ref) const {
        const std::lock_guard<std::mutex> lock(mutex);
        return segments.at(ref.segment)->loadCodePoint(ref.pc);
    }

    CodePointStub addSegment() {
        const std::lock_guard<std::mutex> lock(mutex);
        uint64_t segment_num = next_segment_num++;
        auto new_segment = std::make_shared<CodeSegment>(segment_num);
        auto stub = new_segment->initialCodePointStub();
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
