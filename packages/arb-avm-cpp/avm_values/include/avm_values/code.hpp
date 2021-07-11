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
#include <map>
#include <memory>
#include <mutex>
#include <unordered_map>
#include <vector>

template <typename T>
class CodeBase;

class CoreCode;
class RunningCode;
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

    friend class CoreCode;
    friend class CodeBase<CoreCode>;

    friend class RunningCode;
    friend class CodeBase<RunningCode>;

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
    uint64_t next_segment_num;
};

struct CopiedSegment {
    std::shared_ptr<CodeSegment> segment;
    CodePointStub stub;
};

class Code {
   public:
    virtual ~Code();

    virtual uint64_t getNextSegmentNum() const = 0;

    virtual CodeSegmentSnapshot loadCodeSegment(uint64_t segment_num) const = 0;

    virtual CodePoint loadCodePoint(const CodePointRef& ref) const = 0;

    virtual CodePointStub addSegment() = 0;

    virtual CodePointStub addOperation(const CodePointRef& ref,
                                       Operation op) = 0;

    virtual std::variant<CodePointStub, CopiedSegment> tryAddOperation(
        const CodePointRef& ref,
        Operation op,
        uint64_t new_segment_num) = 0;

    virtual CodeSnapshot snapshot() const = 0;
};

template <typename T>
class CodeBase {
    const T* getThis() const { return static_cast<const T*>(this); }

    T* getThis() { return static_cast<T*>(this); }

   protected:
    CodeSegmentSnapshot loadCodeSegmentImpl(uint64_t segment_num) const {
        auto& segment = getThis()->getSegment(segment_num);
        return {segment, segment->size(), segment->cached_hashes.size()};
    }

    CodePoint loadCodePointImpl(const CodePointRef& ref) const {
        auto& segment = getThis()->getSegment(ref.segment);
        return segment->loadCodePoint(ref.pc);
    }

    CodePointStub addSegmentImpl() {
        uint64_t segment_num = getThis()->nextSegmentNum();
        auto new_segment = std::make_shared<CodeSegment>(segment_num);
        auto stub = new_segment->initialCodePointStub();
        getThis()->storeSegment(std::move(new_segment));
        return stub;
    }

    bool canAppendOperation(const std::shared_ptr<CodeSegment>& segment,
                            const CodePointRef& ref) {
        auto initial_pc = segment->size() - 1;
        if (ref.pc != initial_pc) {
            return false;
        }
        if (segment.use_count() == 1) {
            // No other code has a reference to this segment. That means we
            // can modify the segment even if it forces a reallocation
            return true;
        }

        if (segment->size() < segment->capacity()) {
            // The segment has extra capacity so we can add to it even if
            // other code has references to it
            return true;
        }

        // Fall back to making a copy as there are other references and no
        // space to add this operation
        return false;
    }

    CodePointStub addOperationImpl(const CodePointRef& ref, Operation op) {
        auto& segment = getThis()->getSegment(ref.segment);
        if (canAppendOperation(segment, ref)) {
            return segment->addOperation(std::move(op));
        }

        uint64_t new_segment_num = getThis()->nextSegmentNum();
        auto new_segment = segment->getSubset(new_segment_num, ref.pc);
        auto stub = new_segment->addOperation(std::move(op));
        getThis()->storeSegment(std::move(new_segment));
        return stub;
    }

    // If ref refers to a position in a code segment that can be appended to
    // without reallocation then just add to that segment, otherwise prepare a
    // new segment with the op added to it after ref
    std::variant<CodePointStub, CopiedSegment> tryAddOperationImpl(
        const CodePointRef& ref,
        Operation op,
        uint64_t new_segment_num) {
        auto& segment = getThis()->getSegment(ref.segment);
        if (canAppendOperation(segment, ref)) {
            return segment->addOperation(std::move(op));
        }
        auto new_segment = segment->getSubset(new_segment_num, ref.pc);
        auto stub = new_segment->addOperation(std::move(op));
        return CopiedSegment{std::move(new_segment), stub};
    }
};

class CoreCode : public CodeBase<CoreCode>, public Code {
    friend CodeBase<CoreCode>;

    mutable std::mutex mutex;
    std::unordered_map<uint64_t, std::shared_ptr<CodeSegment>> segments;
    uint64_t next_segment_num;

    const std::shared_ptr<CodeSegment>& getSegment(uint64_t segment_num) const {
        return segments.at(segment_num);
    }

    uint64_t nextSegmentNum() { return next_segment_num++; }

    void storeSegment(std::shared_ptr<CodeSegment> segment) {
        segments[segment->segmentID()] = std::move(segment);
    }

   public:
    CoreCode() : CoreCode(0) {}
    CoreCode(uint64_t next_segment_num_)
        : next_segment_num(next_segment_num_) {}

    uint64_t getNextSegmentNum() const {
        const std::lock_guard<std::mutex> lock(mutex);
        return next_segment_num;
    }

    CodeSegmentSnapshot loadCodeSegment(uint64_t segment_num) const {
        const std::lock_guard<std::mutex> lock(mutex);
        auto& segment = segments.at(segment_num);
        return {segment, segment->size(), segment->cached_hashes.size()};
    }

    bool containsSegment(uint64_t segment_id) const {
        const std::lock_guard<std::mutex> lock(mutex);
        return segments.find(segment_id) != segments.end();
    }

    void commitChanges(RunningCode& code,
                       const std::map<uint64_t, uint64_t>& segment_counts);

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
        return loadCodePointImpl(ref);
    }

    CodePointStub addSegment() {
        const std::lock_guard<std::mutex> lock(mutex);
        return addSegmentImpl();
    }

    void addSegment(std::shared_ptr<CodeSegment> segment) {
        const std::lock_guard<std::mutex> lock(mutex);
        assert(segment->segmentID() == next_segment_num);
        segments[next_segment_num] = std::move(segment);
        next_segment_num++;
    }

    CodePointStub addOperation(const CodePointRef& ref, Operation op) {
        const std::lock_guard<std::mutex> lock(mutex);
        return addOperationImpl(ref, std::move(op));
    }

    std::variant<CodePointStub, CopiedSegment> tryAddOperation(
        const CodePointRef& ref,
        Operation op,
        uint64_t new_segment_num) {
        const std::lock_guard<std::mutex> lock(mutex);
        return tryAddOperationImpl(ref, op, new_segment_num);
    }

    CodePointRef initialCodePointRef() const {
        const std::lock_guard<std::mutex> lock(mutex);
        return {0, segments.at(0)->size() - 1};
    }
};

class RunningCode : public CodeBase<RunningCode>, public Code {
    friend CodeBase<RunningCode>;

    mutable std::mutex mutex;
    uint64_t first_segment;
    std::vector<std::shared_ptr<CodeSegment>> segment_list;

    std::shared_ptr<Code> parent;

    const std::shared_ptr<CodeSegment>& getSegment(uint64_t segment_num) const {
        return segment_list.at(segment_num - first_segment);
    }

    uint64_t nextSegmentNum() const {
        return first_segment + segment_list.size();
    }

    void storeSegment(std::shared_ptr<CodeSegment> segment) {
        segment_list.push_back(std::move(segment));
    }

   public:
    RunningCode(std::shared_ptr<Code> parent_)
        : first_segment(parent_->getNextSegmentNum()),
          parent(std::move(parent_)) {}

    uint64_t fillInCode(
        std::unordered_map<uint64_t, std::shared_ptr<CodeSegment>>&
            parent_segments,
        const std::map<uint64_t, uint64_t>& segment_counts) const {
        const std::lock_guard<std::mutex> lock(mutex);
        auto it = segment_counts.lower_bound(first_segment);
        auto end = segment_counts.end();
        for (; it != end; ++it) {
            auto inserted = parent_segments.insert(
                std::make_pair(it->first, getSegment(it->first)));
            // Verify that the element didn't exist previously
            assert(inserted.second);
            if (!inserted.second) {
                throw std::runtime_error(
                    "code segment id collision when filling in code");
            }
        }
        return nextSegmentNum();
    }

    std::shared_ptr<Code> getParent() const { return parent; }

    uint64_t getNextSegmentNum() const {
        const std::lock_guard<std::mutex> lock(mutex);
        return first_segment + segment_list.size();
    }

    CodeSnapshot snapshot() const {
        auto snap = parent->snapshot();
        const std::lock_guard<std::mutex> lock(mutex);
        for (const auto& segment : segment_list) {
            snap.segments[segment->segmentID()] = {
                segment, segment->size(), segment->cached_hashes.size()};
        }
        snap.next_segment_num = first_segment + segment_list.size();
        return snap;
    }

    CodeSegmentSnapshot loadCodeSegment(uint64_t segment_num) const {
        const std::lock_guard<std::mutex> lock(mutex);
        if (segment_num < first_segment) {
            return parent->loadCodeSegment(segment_num);
        }
        return loadCodeSegmentImpl(segment_num);
    }

    CodePoint loadCodePoint(const CodePointRef& ref) const {
        const std::lock_guard<std::mutex> lock(mutex);
        if (ref.segment < first_segment) {
            return parent->loadCodePoint(ref);
        }
        return loadCodePointImpl(ref);
    }

    CodePointStub addSegment() {
        const std::lock_guard<std::mutex> lock(mutex);
        return addSegmentImpl();
    }

    CodePointStub addOperation(const CodePointRef& ref, Operation op) {
        const std::lock_guard<std::mutex> lock(mutex);
        if (ref.segment < first_segment) {
            auto add_var = parent->tryAddOperation(ref, std::move(op),
                                                   getNextSegmentNum());
            if (std::holds_alternative<CodePointStub>(add_var)) {
                return std::get<CodePointStub>(add_var);
            } else {
                auto& added = std::get<CopiedSegment>(add_var);
                storeSegment(std::move(added.segment));
                return added.stub;
            }
        }
        return addOperationImpl(ref, std::move(op));
    }

    std::variant<CodePointStub, CopiedSegment> tryAddOperation(
        const CodePointRef& ref,
        Operation op,
        uint64_t new_segment_num) {
        const std::lock_guard<std::mutex> lock(mutex);
        if (ref.segment < first_segment) {
            return parent->tryAddOperation(ref, std::move(op), new_segment_num);
        }
        return tryAddOperationImpl(ref, op, new_segment_num);
    }
};

#endif /* code_hpp */
