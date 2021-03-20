/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

#include <avm_values/code.hpp>

CodeSegment CodeSegment::newSegment() {
    return CodeSegment(
        std::make_shared<CodeSegmentInner>(next_segment_id.fetch_add(1)));
}

CodeSegment CodeSegment::cloneWithSize(uint64_t size) const {
    // Require that the mutex is locked when this is called
    assert(!inner->mutex.try_lock());
    if (size > inner->code.size()) {
        throw new std::runtime_error(
            "Attempted to create code segment reaching past end");
    }
    CodeSegment ret = newSegment();
    ret.inner->code.resize(size);
    std::copy(inner->code.begin(), inner->code.begin() + size,
              std::back_inserter(ret.inner->code));
    return ret;
}

CodePointStub CodeSegment::addOperationAt(Operation op, uint64_t pc) {
    std::unique_lock<std::shared_mutex> guard(inner->mutex, std::try_lock_t);
    if (!guard.owns_lock()) {
        // This code segment is being concurrently accessed (probably currently
        // loaded)
        std::shared_lock<std::shared_mutex> guard2(inner->mutex);
        return cloneWithSize(pc).addOperationAt(op, pc);
    } else if (pc < inner->code.size()) {
        return cloneWithSize(pc).addOperationAt(op, pc);
    } else if (pc > inner->code.size()) {
        throw new std::runtime_error(
            "Attempted to insert operation past end of code segment");
    }
    // At this point, pc == inner->code.size(), so we just need to append
    uint256_t prev_hash = 0;
    if (inner->code.size() > 0) {
        prev_hash = hash(inner->code.back());
    }
    inner->code.emplace_back(std::move(op), prev_hash);
    return {{*this, pc}, hash(inner->code.back())};
}
