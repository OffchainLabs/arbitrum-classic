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
#include <avm_values/codepoint.hpp>
#include <avm_values/codepointstub.hpp>

CodeSegmentInner::CodeSegmentInner(uint64_t segment_id_)
    : segment_id(segment_id_) {}

CodeSegmentInner::CodeSegmentInner(uint64_t segment_id_,
                                   std::vector<CodePoint> code_)
    : segment_id(segment_id_), code(code_) {}

CodeSegmentInner::~CodeSegmentInner() = default;

LoadedCodeSegment::LoadedCodeSegment(CodeSegment segment_)
    : CodeSegment(std::move(segment_)), guard(inner->mutex) {}

const CodePoint& LoadedCodeSegment::operator[](uint64_t pc) const {
    return inner->code.at(pc);
}

std::vector<CodePoint>::const_iterator LoadedCodeSegment::begin() const {
    return inner->code.begin();
}

std::vector<CodePoint>::const_iterator LoadedCodeSegment::end() const {
    return inner->code.end();
}

size_t LoadedCodeSegment::size() const {
    return inner->code.size();
}

std::atomic<uint64_t> next_segment_id;

CodeSegment CodeSegment::newSegment() {
    return CodeSegment(
        std::make_shared<CodeSegmentInner>(next_segment_id.fetch_add(1)));
}

CodeSegment CodeSegment::restoreCodeSegment(uint64_t segment_id,
                                            std::vector<CodePoint> code) {
    if (segment_id >= next_segment_id.load()) {
        throw new std::runtime_error(
            "Attempted to restore code segment after next segment id");
    }
    return CodeSegment(std::make_shared<CodeSegmentInner>(segment_id, code));
}

void CodeSegment::restoreNextSegmentId(uint64_t next_segment_id_) {
    auto old_val = next_segment_id.exchange(next_segment_id_);
    if (old_val != 0) {
        throw std::runtime_error(
            "Attempted to restore next segment ID post-initialization");
    }
}

CodeSegment CodeSegment::cloneWithSize(uint64_t size) const {
    // Require that the mutex is locked when this is called
    assert(!inner->mutex.try_lock());
    if (size > inner->code.size()) {
        throw new std::runtime_error(
            "Attempted to create code segment reaching past end");
    }
    CodeSegment ret = CodeSegment::newSegment();
    auto& ret_code = ret.inner->code;
    ret_code.erase(ret_code.begin() + size, ret_code.end());
    std::copy(inner->code.begin(), inner->code.begin() + size,
              std::back_inserter(ret_code));
    return ret;
}

CodePointStub CodeSegment::addOperationAt(Operation op, uint64_t pc) {
    std::unique_lock<std::shared_mutex> guard(inner->mutex, std::try_to_lock);
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
    return {CodePointRef(*this, pc), hash(inner->code.back())};
}

LoadedCodeSegment CodeSegment::load() const {
    return LoadedCodeSegment(*this);
}

uint256_t segmentIdToDbHash(uint64_t segment_id) {
    return hash('c' << 24 | 'o' << 16 | 'd' << 8 | 'e', segment_id);
}

uint256_t hash(CodeSegment segment) {
    return segmentIdToDbHash(segment.segmentID());
}
