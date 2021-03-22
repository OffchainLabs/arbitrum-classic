/*
 * Copyright 2021, Offchain Labs, Inc.
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

#include <data_storage/value/deserialize.hpp>

#include <avm_values/codepointstub.hpp>
#include <avm_values/tuple.hpp>

uint256_t deserializeNum(std::vector<unsigned char>::const_iterator& bytes,
                         std::vector<Slot>& slots) {
    return deserializeUint256t(bytes);
}
CodePointStub deserializeCodePointStub(
    std::vector<unsigned char>::const_iterator& bytes,
    std::vector<Slot>& slots) {
    auto segment_id = deserializeUint64t(bytes);
    auto pc = deserializeUint64t(bytes);
    auto next_hash = deserializeUint256t(bytes);
    CodeSegment segment = CodeSegment::uninitialized();
    slots.emplace_back(&segment, segmentIdToDbHash(segment_id));
    return {{segment, pc}, next_hash};
}
Tuple deserializeTuple(std::vector<unsigned char>::const_iterator& bytes,
                       std::vector<Slot>& slots,
                       size_t size) {
    Tuple ret = Tuple::createSizedTuple(size);
    for (uint64_t i = 0; i < size; i++) {
        auto ty = *bytes;
        if (ty == HASH_PRE_IMAGE) {
            bytes++;
            auto hash = deserializeUint256t(bytes);
            slots.emplace_back(ret.getElementPointer(i), hash);
            ret.markContentsWillChange();
        } else {
            ret.unsafe_set_element(i, deserializeValue(bytes, slots));
        }
    }
}
Buffer deserializeBuffer(std::vector<unsigned char>::const_iterator& bytes,
                         std::vector<Slot>& slots) {
    uint8_t depth = *bytes++;
    if (depth == 0) {
        Buffer::LeafData leaf;
        std::copy(bytes, bytes + 32, leaf.begin());
        bytes += 32;
        return Buffer(leaf);
    } else {
        auto left = std::make_shared<Buffer>();
        auto right = std::make_shared<Buffer>();
        for (uint8_t i = 0; i < 2; i++) {
            if (*bytes++ != HASH_PRE_IMAGE) {
                throw std::runtime_error(
                    "TODO: implement inline buffer deserialization");
            }
            auto hash = deserializeUint256t(bytes);
            auto ptr = (i == 0) ? left.get() : right.get();
            slots.emplace_back(ptr, hash);
        }
        return Buffer(left, right);
    }
}
CodeSegment deserializeCodeSegment(
    std::vector<unsigned char>::const_iterator& bytes,
    std::vector<Slot>& slots) {
    auto segment_id = deserializeUint64t(bytes);
    auto num_code_points = deserializeUint64t(bytes);
    std::vector<CodePoint> code;
    code.reserve(num_code_points);
    for (const CodePoint& point : code) {
        bool has_immediate = *bytes++;
        auto op = static_cast<OpCode>(*bytes++);
        auto next_hash = deserializeUint256t(bytes);
        std::optional<value> immediate;
        if (has_immediate) {
            immediate = deserializeValue(bytes, slots);
        }
        code.emplace_back(op, immediate, next_hash);
    }
    return CodeSegment::restoreCodeSegment(segment_id, code);
}

value deserializeValue(std::vector<unsigned char>::const_iterator& bytes,
                       std::vector<Slot>& slots) {
    auto ty = *bytes++;
    switch (ty) {
        case BUFFER: {
            return deserializeBuffer(bytes, slots);
        }
        case NUM: {
            return deserializeNum(bytes, slots);
        }
        case CODE_POINT_STUB: {
            return deserializeCodePointStub(bytes, slots);
        }
        case HASH_PRE_IMAGE: {
            throw std::runtime_error("Attempted to deserialize HASH_PRE_IMAGE");
        }
        case CODE_SEGMENT: {
            return deserializeCodeSegment(bytes, slots);
        }
        default: {
            auto size = ty - TUPLE;
            if (size > 8) {
                throw std::runtime_error(
                    "attempted to deserialize value with invalid typecode");
            }
            return deserializeTuple(bytes, slots, size);
        }
    }
}
