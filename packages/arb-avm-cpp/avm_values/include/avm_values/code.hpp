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
#include <avm_values/value.hpp>

#include <cassert>
#include <map>
#include <memory>
#include <mutex>
#include <shared_mutex>
#include <unordered_map>
#include <vector>

class RunningCode;

struct CodeSegmentData {
    static_assert(std::is_nothrow_move_constructible<Operation>::value,
                  "Operation should be noexcept MoveConstructible");
    static_assert(std::is_nothrow_move_constructible<uint256_t>::value,
                  "uint256_t should be noexcept MoveConstructible");
    std::vector<Operation> operations;
    std::vector<uint256_t> cached_hashes;
    uint256_t prev_hash;

    CodeSegmentData() = default;

    explicit CodeSegmentData(std::vector<Operation> ops) {
        for (auto it = ops.rbegin(); it != ops.rend(); ++it) {
            addOperation(std::move(*it));
        }
    }

    CodeSegmentData(std::vector<Operation> operations_,
                    std::vector<uint256_t> next_hashes_)
        : operations(std::move(operations_)),
          cached_hashes(std::move(next_hashes_)) {
        prev_hash = ::hash(loadCodePoint(operations.size() - 1));
    }

    void addOperation(Operation op) {
        CodePoint cp{std::move(op), prev_hash};
        prev_hash = hash(cp);
        operations.push_back(std::move(cp.op));
        if (operations.size() % 10 == 0) {
            cached_hashes.push_back(prev_hash);
        }
    }

    // Return the subset of this code segment starting in the given pc
    [[nodiscard]] CodeSegmentData getSubset(uint64_t pc) const {
        auto ops = static_cast<std::vector<CodePoint>::difference_type>(pc);
        std::vector<Operation> copy_copy;
        auto code_copy = std::vector<Operation>{operations.begin(),
                                                operations.begin() + ops + 1};
        auto hashes_copy = std::vector<uint256_t>{
            cached_hashes.begin(), cached_hashes.begin() + (ops + 1) / 10};
        // Make endpoint pc + 1 since pc should be included in segment
        return {std::move(code_copy), std::move(hashes_copy)};
    }

    [[nodiscard]] CodePoint loadCodePoint(uint64_t pc) const {
        uint256_t previous_hash;

        if (pc / 10 > 0) {
            previous_hash = cached_hashes.at(pc / 10 - 1);
        }
        for (uint64_t i = (pc / 10) * 10; i < pc; i++) {
            previous_hash = hash(CodePoint(operations[i], previous_hash));
        }
        return {operations[pc], previous_hash};
    }

    void reserve(size_t size) {
        operations.reserve(size);
        cached_hashes.reserve(size / 10);
    }
};

class UnsafeCodeSegment {
   public:
    uint64_t segment_id;
    CodeSegmentData data;

    explicit UnsafeCodeSegment(uint64_t segment_id_) : segment_id(segment_id_) {
        addOperation(getErrOperation());
    }

    UnsafeCodeSegment(uint64_t segment_id_, CodeSegmentData data_)
        : segment_id(segment_id_), data(std::move(data_)) {}

    UnsafeCodeSegment(uint64_t segment_id_, std::vector<Operation> ops)
        : segment_id(segment_id_), data(std::move(ops)) {}

    [[nodiscard]] uint64_t segmentID() const { return segment_id; }

    [[nodiscard]] CodePoint loadCodePoint(uint64_t pc) const {
        return data.loadCodePoint(pc);
    }

    [[nodiscard]] const Operation& loadOperation(uint64_t pc) const {
        return data.operations[pc];
    }
    [[nodiscard]] const uint256_t& loadCachedHash(uint64_t i) const {
        return data.cached_hashes[i];
    }

    [[nodiscard]] CodePointStub initialCodePointStub() const {
        return {{segment_id, 0}, getErrCodePointHash()};
    }

    void reserve(size_t size) { data.reserve(size); }

    [[nodiscard]] size_t capacity() const { return data.operations.capacity(); }

    [[nodiscard]] size_t size() const { return data.operations.size(); }

    CodePointStub addOperation(Operation op) {
        data.addOperation(std::move(op));
        return {{segment_id, data.operations.size() - 1}, data.prev_hash};
    }

    [[nodiscard]] CodePointStub lastCodePointStubAdded() const {
        return {{segment_id, data.operations.size() - 1}, data.prev_hash};
    }

    // Return the subset of this code segment starting in the given pc
    [[nodiscard]] std::shared_ptr<UnsafeCodeSegment> getSubset(
        uint64_t new_segment_id,
        uint64_t pc) const {
        return std::make_shared<UnsafeCodeSegment>(new_segment_id,
                                                   data.getSubset(pc));
    }
};

struct CodeSegmentSnapshot {
   private:
    std::shared_ptr<const UnsafeCodeSegment> segment;

   public:
    uint64_t op_count{};
    uint64_t cached_hash_count{};

    CodeSegmentSnapshot() = default;
    CodeSegmentSnapshot(std::shared_ptr<const UnsafeCodeSegment> segment_,
                        uint64_t op_count_,
                        uint64_t cached_hash_count_)
        : segment(std::move(segment_)),
          op_count(op_count_),
          cached_hash_count(cached_hash_count_) {}

    [[nodiscard]] uint64_t segmentID() const { return segment->segmentID(); }

    [[nodiscard]] CodePoint loadCodePoint(uint64_t pc) const {
        return segment->loadCodePoint(pc);
    }

    [[nodiscard]] const Operation& loadOperation(uint64_t pc) const {
        return segment->loadOperation(pc);
    }

    [[nodiscard]] const uint256_t& loadCachedHash(uint64_t i) const {
        return segment->loadCachedHash(i);
    }
};

struct CodeSnapshot {
    std::unordered_map<uint64_t, CodeSegmentSnapshot> segments;
    uint64_t next_segment_num;
};

struct SegmentsAndLock {
    std::unordered_map<uint64_t, std::shared_ptr<UnsafeCodeSegment>>* segments;
    uint64_t* next_segment_num;
    std::unique_lock<std::shared_mutex> lock;
};

class Code {
   public:
    virtual ~Code() = default;

    // Returns mutable segments and a lock for the root CoreCode.
    [[nodiscard]] virtual SegmentsAndLock getRootSegments() const = 0;

    [[nodiscard]] virtual uint64_t initialSegmentForChildCode() const = 0;

    // Warning: attempting to access a segment past the last created segment
    // from this code is undefined behavior and may result in an unexpected
    // segment being returned.
    [[nodiscard]] virtual CodeSegmentSnapshot loadCodeSegment(
        uint64_t segment_num) const = 0;

    // Warning: attempting to access a segment past the last created segment
    // from this code is undefined behavior and may result in an unexpected
    // segment being returned.
    [[nodiscard]] virtual CodePoint loadCodePoint(
        const CodePointRef& ref) const = 0;

    virtual CodePointStub addSegment() = 0;

    virtual CodePointStub addOperation(const CodePointRef& ref,
                                       Operation op) = 0;

    virtual std::variant<CodePointStub, CodeSegmentData> tryAddOperation(
        const CodePointRef& ref,
        Operation op) = 0;

    [[nodiscard]] virtual CodeSnapshot snapshot() const = 0;

    // Warning: attempting to access a segment past the last created segment
    // from this code is undefined behavior and may result in this unexpectedly
    // returning true.
    [[nodiscard]] virtual bool containsSegment(uint64_t segment_id) const = 0;

    virtual void commitCodeToCore(
        const std::map<uint64_t, uint64_t>& segment_counts) const = 0;

    // Removes any segments colliding with the CoreCode,
    // which should only happen after a reorg.
    virtual void cleanupAfterReorg() = 0;
};

template <typename T>
class CodeBase {
   protected:
    std::unique_ptr<T> impl;

   public:
    template <typename... Args>
    explicit CodeBase(Args&&... args)
        : impl(std::make_unique<T>(std::forward<Args>(args)...)) {}

   protected:
    [[nodiscard]] CodeSegmentSnapshot loadCodeSegmentImpl(
        uint64_t segment_num) const {
        auto& segment = impl->getSegment(segment_num);
        return {segment, segment->size(), segment->data.cached_hashes.size()};
    }

    [[nodiscard]] CodePoint loadCodePointImpl(const CodePointRef& ref) const {
        auto& segment = impl->getSegment(ref.segment);
        return segment->loadCodePoint(ref.pc);
    }

    CodePointStub addSegmentImpl() {
        uint64_t segment_num = impl->nextSegmentNum();
        if (segment_num >= (uint64_t(1) << 62)) {
            throw std::runtime_error("Exceeded limit of 2^62 segments");
        }
        auto new_segment = std::make_shared<UnsafeCodeSegment>(segment_num);
        auto stub = new_segment->initialCodePointStub();
        impl->storeSegment(std::move(new_segment));
        return stub;
    }

    bool canAppendOperation(const std::shared_ptr<UnsafeCodeSegment>& segment,
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
        auto& segment = impl->getSegment(ref.segment);
        if (canAppendOperation(segment, ref)) {
            return segment->addOperation(std::move(op));
        }

        uint64_t new_segment_num = impl->nextSegmentNum();
        auto new_segment = segment->getSubset(new_segment_num, ref.pc);
        auto stub = new_segment->addOperation(std::move(op));
        impl->storeSegment(std::move(new_segment));
        return stub;
    }

    // If ref refers to a position in a code segment that can be appended to
    // without reallocation then just add to that segment, otherwise prepare a
    // new segment with the op added to it after ref
    std::variant<CodePointStub, CodeSegmentData> tryAddOperationImpl(
        const CodePointRef& ref,
        Operation op) {
        auto& segment = impl->getSegment(ref.segment);
        if (canAppendOperation(segment, ref)) {
            return segment->addOperation(std::move(op));
        }
        auto new_segment = segment->data.getSubset(ref.pc);
        new_segment.addOperation(std::move(op));
        return new_segment;
    }
};

struct CoreCodeImpl {
    std::unordered_map<uint64_t, std::shared_ptr<UnsafeCodeSegment>> segments;
    uint64_t next_segment_num;

    const std::shared_ptr<UnsafeCodeSegment>& getSegment(
        uint64_t segment_num) const {
        return segments.at(segment_num);
    }

    // Warning: increments next_segment_num, unlike RunningCode's nextSegmentNum
    uint64_t nextSegmentNum() { return next_segment_num++; }

    void storeSegment(std::shared_ptr<UnsafeCodeSegment> segment) {
        segments[segment->segmentID()] = std::move(segment);
    }

    CoreCodeImpl() : CoreCodeImpl(0) {}
    explicit CoreCodeImpl(uint64_t next_segment_num_)
        : next_segment_num(next_segment_num_) {}
};

class CoreCode : public CodeBase<CoreCodeImpl>, public Code {
    mutable std::shared_mutex mutex;

   public:
    CoreCode() : CodeBase<CoreCodeImpl>(0) {}
    explicit CoreCode(uint64_t next_segment_num_)
        : CodeBase<CoreCodeImpl>(next_segment_num_) {}

    SegmentsAndLock getRootSegments() const override {
        auto lock = std::unique_lock<std::shared_mutex>(mutex);
        return SegmentsAndLock{&impl->segments, &impl->next_segment_num,
                               std::move(lock)};
    }

    uint64_t initialSegmentForChildCode() const override {
        const std::shared_lock<std::shared_mutex> lock(mutex);
        return impl->next_segment_num;
    }

    CodeSegmentSnapshot loadCodeSegment(uint64_t segment_num) const override {
        const std::shared_lock<std::shared_mutex> lock(mutex);
        auto& segment = impl->segments.at(segment_num);
        return {segment, segment->size(), segment->data.cached_hashes.size()};
    }

    bool containsSegment(uint64_t segment_id) const override {
        const std::shared_lock<std::shared_mutex> lock(mutex);
        return impl->segments.find(segment_id) != impl->segments.end();
    }

    void restoreExistingSegment(std::shared_ptr<UnsafeCodeSegment> segment) {
        const std::unique_lock<std::shared_mutex> lock(mutex);
        uint64_t segment_id = segment->segmentID();
        if (segment_id >= impl->next_segment_num) {
            throw std::runtime_error("code segment loaded incorrectly");
        }
        if (impl->segments.find(segment->segmentID()) == impl->segments.end()) {
            impl->segments[segment_id] = std::move(segment);
        }
    }

    CodeSnapshot snapshot() const override {
        const std::shared_lock<std::shared_mutex> lock(mutex);
        std::unordered_map<uint64_t, CodeSegmentSnapshot> copied_segments;
        for (const auto& key_val : impl->segments) {
            copied_segments[key_val.first] = {
                key_val.second, key_val.second->size(),
                key_val.second->data.cached_hashes.size()};
        }
        return {std::move(copied_segments), impl->next_segment_num};
    }

    CodePoint loadCodePoint(const CodePointRef& ref) const override {
        const std::shared_lock<std::shared_mutex> lock(mutex);
        return loadCodePointImpl(ref);
    }

    CodePointStub addSegment() override {
        const std::unique_lock<std::shared_mutex> lock(mutex);
        return addSegmentImpl();
    }

    void addSegment(std::shared_ptr<UnsafeCodeSegment> segment) {
        const std::unique_lock<std::shared_mutex> lock(mutex);
        assert(segment->segmentID() == impl->next_segment_num);
        impl->segments[impl->next_segment_num] = std::move(segment);
        impl->next_segment_num++;
    }

    CodePointStub addOperation(const CodePointRef& ref, Operation op) override {
        const std::unique_lock<std::shared_mutex> lock(mutex);
        return addOperationImpl(ref, std::move(op));
    }

    std::variant<CodePointStub, CodeSegmentData> tryAddOperation(
        const CodePointRef& ref,
        Operation op) override {
        const std::unique_lock<std::shared_mutex> lock(mutex);
        return tryAddOperationImpl(ref, std::move(op));
    }

    CodePointRef initialCodePointRef() const {
        const std::shared_lock<std::shared_mutex> lock(mutex);
        return {0, impl->segments.at(0)->size() - 1};
    }

    void commitCodeToCore(const std::map<uint64_t, uint64_t>&) const override {}

    void cleanupAfterReorg() override {}
};

struct RunningCodeImpl {
    uint64_t first_segment;
    std::vector<std::shared_ptr<UnsafeCodeSegment>> segment_list;

    explicit RunningCodeImpl(uint64_t first_segment_)
        : first_segment(first_segment_) {}

    [[nodiscard]] const std::shared_ptr<UnsafeCodeSegment>& getSegment(
        uint64_t segment_num) const {
        return segment_list.at(segment_num - first_segment);
    }

    // Warning: doesn't increment, unlike CoreCode's nextSegmentNum
    [[nodiscard]] uint64_t nextSegmentNum() const {
        return first_segment + segment_list.size();
    }

    void storeSegment(std::shared_ptr<UnsafeCodeSegment> segment) {
        segment_list.push_back(std::move(segment));
    }
};

class RunningCode : public CodeBase<RunningCodeImpl>, public Code {
    mutable std::shared_mutex mutex;

    std::shared_ptr<Code> parent;

    // Requires the mutex is held.
    // Gives no guarantee that parent actually contains segment_id,
    // only that the parent should be asked for the segment.
    bool segmentInParent(uint64_t segment_id) const {
        assert(!mutex.try_lock() && "mutex not held in segmentInParent call");
        return segment_id < impl->first_segment ||
               segment_id >= impl->nextSegmentNum() ||
               impl->getSegment(segment_id) == nullptr;
    }

   public:
    explicit RunningCode(std::shared_ptr<Code> parent_, uint64_t initialSegment)
        : CodeBase<RunningCodeImpl>(initialSegment),
          parent(std::move(parent_)) {}

    explicit RunningCode(std::shared_ptr<Code> parent_)
        : RunningCode(parent_, parent_->initialSegmentForChildCode()) {}

    SegmentsAndLock getRootSegments() const override {
        return parent->getRootSegments();
    }

    void commitCodeToCore(
        const std::map<uint64_t, uint64_t>& segment_counts) const override {
        parent->commitCodeToCore(segment_counts);
        const std::shared_lock<std::shared_mutex> lock(mutex);
        auto root_segments = parent->getRootSegments();
        auto it = segment_counts.lower_bound(impl->first_segment);
        auto end = segment_counts.lower_bound(impl->nextSegmentNum());
        for (; it != end; ++it) {
            auto segment = impl->getSegment(it->first);
            if (segment == nullptr) {
                continue;
            }
            auto inserted = root_segments.segments->insert(
                std::make_pair(it->first, segment));
            // Verify that the element didn't exist previously
            assert(inserted.second);
            if (!inserted.second) {
                throw std::runtime_error(
                    "code segment id collision when filling in code");
            }
        }
        if (impl->nextSegmentNum() > *root_segments.next_segment_num) {
            *root_segments.next_segment_num = impl->nextSegmentNum();
        }
    }

    const std::shared_ptr<Code>& getParent() const { return parent; }

    uint64_t initialSegmentForChildCode() const override {
        const std::shared_lock<std::shared_mutex> lock(mutex);
        return impl->first_segment + impl->segment_list.size();
    }

    CodeSnapshot snapshot() const override {
        auto snap = parent->snapshot();
        const std::shared_lock<std::shared_mutex> lock(mutex);
        for (const auto& segment : impl->segment_list) {
            if (segment == nullptr) {
                continue;
            }
            snap.segments[segment->segmentID()] = {
                segment, segment->size(), segment->data.cached_hashes.size()};
        }
        snap.next_segment_num = impl->first_segment + impl->segment_list.size();
        return snap;
    }

    CodeSegmentSnapshot loadCodeSegment(uint64_t segment_num) const override {
        std::shared_lock<std::shared_mutex> lock(mutex);
        if (segmentInParent(segment_num)) {
            lock.unlock();
            return parent->loadCodeSegment(segment_num);
        }
        return loadCodeSegmentImpl(segment_num);
    }

    CodePoint loadCodePoint(const CodePointRef& ref) const override {
        std::shared_lock<std::shared_mutex> lock(mutex);
        if (segmentInParent(ref.segment)) {
            lock.unlock();
            return parent->loadCodePoint(ref);
        }
        return loadCodePointImpl(ref);
    }

    CodePointStub addSegment() override {
        const std::unique_lock<std::shared_mutex> lock(mutex);
        return addSegmentImpl();
    }

    CodePointStub addOperation(const CodePointRef& ref, Operation op) override {
        const std::unique_lock<std::shared_mutex> lock(mutex);
        if (segmentInParent(ref.segment)) {
            // We don't unlock here as we need the lock if
            // parent->tryAddOperation fails
            auto add_var = parent->tryAddOperation(ref, std::move(op));
            if (holds_alternative<CodePointStub>(add_var)) {
                return get<CodePointStub>(add_var);
            } else {
                auto& added = std::get<CodeSegmentData>(add_var);
                auto new_segment = std::make_shared<UnsafeCodeSegment>(
                    impl->nextSegmentNum(), std::move(added));
                auto stub = new_segment->lastCodePointStubAdded();
                impl->storeSegment(std::move(new_segment));
                return stub;
            }
        }
        return addOperationImpl(ref, std::move(op));
    }

    std::variant<CodePointStub, CodeSegmentData> tryAddOperation(
        const CodePointRef& ref,
        Operation op) override {
        std::shared_lock<std::shared_mutex> lock(mutex);
        if (segmentInParent(ref.segment)) {
            lock.unlock();
            return parent->tryAddOperation(ref, std::move(op));
        }
        return tryAddOperationImpl(ref, std::move(op));
    }

    bool containsSegment(uint64_t segment_id) const override {
        std::shared_lock<std::shared_mutex> lock(mutex);
        if (segmentInParent(segment_id)) {
            lock.unlock();
            return parent->containsSegment(segment_id);
        }
        return segment_id < impl->nextSegmentNum();
    }

    void cleanupAfterReorg() override {
        parent->cleanupAfterReorg();
        const std::shared_lock<std::shared_mutex> lock(mutex);
        auto root_segments = parent->getRootSegments();
        for (size_t i = 0; i < impl->segment_list.size(); i++) {
            auto segment_id = impl->first_segment + i;
            if (root_segments.segments->find(segment_id) !=
                root_segments.segments->end()) {
                impl->segment_list[i] = nullptr;
            }
        }
    }
};

#endif /* code_hpp */
