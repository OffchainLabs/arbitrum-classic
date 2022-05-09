/*
 * Copyright 2022, Offchain Labs, Inc.
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

#ifndef codesegment_hpp
#define codesegment_hpp

#include <avm_values/codepoint.hpp>

#include <mutex>

class ValueLoader;

struct CodeSegmentDataHandle {
    std::shared_lock<std::shared_mutex> lock;
    const std::vector<Operation>* operations;
    const std::vector<uint256_t>* cached_hashes;
};

class CodeSegment {
   private:
    static_assert(std::is_nothrow_move_constructible<Operation>::value,
                  "Operation should be noexcept MoveConstructible");
    static_assert(std::is_nothrow_move_constructible<uint256_t>::value,
                  "uint256_t should be noexcept MoveConstructible");
    mutable std::shared_mutex mutex;
    mutable std::vector<Operation> operations;
    mutable std::vector<uint256_t> cached_hashes;
    mutable uint256_t prev_hash;
    mutable bool frozen;

    // requires the mutex is held
    void requireLoaded() const {
        if (prev_hash != 0 && operations.size() == 0) {
            throw std::runtime_error(
                "Attempted to perform operation on unloaded code segment");
        }
    }

   public:
    explicit CodeSegment(std::vector<Operation> ops) {
        for (auto it = ops.rbegin(); it != ops.rend(); ++it) {
            addOperation(std::move(*it));
        }
    }

    CodeSegment(std::vector<Operation> operations_,
                std::vector<uint256_t> next_hashes_)
        : operations(std::move(operations_)),
          cached_hashes(std::move(next_hashes_)) {
        prev_hash = ::hash(loadCodePoint(operations.size() - 1));
    }

    void freeze() const {
        std::lock_guard<std::shared_mutex> lock(mutex);
        frozen = true;
    }

    [[nodiscard]] bool isLoaded() const {
        std::shared_lock<std::shared_mutex> lock(mutex);
        return prev_hash == 0 || operations.size() > 0;
    }

    [[nodiscard]] uint64_t size() const {
        std::shared_lock<std::shared_mutex> lock(mutex);
        requireLoaded();
        return operations.size();
    }

    void addOperation(Operation op) const {
        std::lock_guard<std::shared_mutex> lock(mutex);
        requireLoaded();
        CodePoint cp{std::move(op), prev_hash};
        prev_hash = hash(cp);
        operations.push_back(std::move(cp.op));
        if (operations.size() % 10 == 0) {
            cached_hashes.push_back(prev_hash);
        }
    }

    [[nodiscard]] bool tryAddOperation(Operation op, uint64_t pc) const {
        std::unique_lock<std::shared_mutex> lock(mutex, std::try_to_lock);
        if (!lock.owns_lock()) {
            return false;
        }
        requireLoaded();
        if (size() != pc || frozen) {
            return false;
        }
        CodePoint cp{std::move(op), prev_hash};
        prev_hash = hash(cp);
        operations.push_back(std::move(cp.op));
        if (operations.size() % 10 == 0) {
            cached_hashes.push_back(prev_hash);
        }
        return true;
    }

    // Return the subset of this code segment starting in the given pc
    [[nodiscard]] CodeSegment getSubset(uint64_t pc) const {
        std::shared_lock<std::shared_mutex> lock(mutex);
        requireLoaded();
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
        std::shared_lock<std::shared_mutex> lock(mutex);
        requireLoaded();
        uint256_t previous_hash;

        if (pc / 10 > 0) {
            previous_hash = cached_hashes.at(pc / 10 - 1);
        }
        for (uint64_t i = (pc / 10) * 10; i < pc; i++) {
            previous_hash = hash(CodePoint(operations[i], previous_hash));
        }
        return {operations[pc], previous_hash};
    }

    [[nodiscard]] const Operation& loadOperation(uint64_t pc) const {
        std::shared_lock<std::shared_mutex> lock(mutex);
        requireLoaded();
        return operations[pc];
    }

    [[nodiscard]] Operation& loadOperationUnsafe(uint64_t pc) {
        std::unique_lock<std::shared_mutex> lock(mutex);
        requireLoaded();
        return operations[pc];
    }

    [[nodiscard]] uint256_t getHash() const {
        std::shared_lock<std::shared_mutex> lock(mutex);
        return prev_hash;
    }

    void reserve(size_t size) const {
        std::lock_guard<std::shared_mutex> lock(mutex);
        requireLoaded();
        operations.reserve(size);
        cached_hashes.reserve(size / 10);
    }

    boost::intrusive_ptr<CodeSegment> fullClone(uint64_t len) const {
        std::shared_lock<std::shared_mutex> lock(mutex);
        if (len > size()) {
            throw std::runtime_error(
                "Attempted to clone code segment past size");
        }
        std::vector<Operation> operations;
        std::vector<uint256_t> cached_hashes;
        operations.insert(operations.end(), operations.begin(),
                          operations.begin() + len);
        cached_hashes.insert(cached_hashes.end(), cached_hashes.begin(),
                             cached_hashes.begin() + len);
        return new CodeSegment(operations, cached_hashes);
    }

    CodeSegmentDataHandle getDataHandle() const {
        std::shared_lock<std::shared_mutex> lock(mutex);
        return {std::move(lock), &operations, &cached_hashes};
    }

    static boost::intrusive_ptr<CodeSegment> newCodeSegment() {
        return new CodeSegment(std::vector<Operation>(1, getErrOperation()));
    }

    static CodePointStub errCodePt() {
        return CodePointStub{CodePointRef{newCodeSegment(), 1},
                             getErrCodePointHash()};
    }

   private:
    std::atomic<uint64_t> reference_count = 0;

    friend void intrusive_ptr_add_ref(CodeSegment* x) {
        x->reference_count.fetch_add(std::memory_order_relaxed);
    }

    friend void intrusive_ptr_release(CodeSegment* x) {
        if (x->reference_count.fetch_sub(std::memory_order_release) == 0) {
            std::atomic_thread_fence(std::memory_order_acquire);
            delete x;
        }
    }
};

uint256_t hash(const boost::intrusive_ptr<CodeSegment>& segment) {
    return segment->getHash();
}

#endif /* codesegment_hpp */
