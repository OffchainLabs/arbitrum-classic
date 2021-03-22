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

#include <avm_values/code.hpp>
#include <avm_values/codepointstub.hpp>
#include <avm_values/tuple.hpp>

void deserializeNum(std::vector<unsigned char>::const_iterator& bytes,
                    value* result,
                    std::vector<Slot>&) {
    *result = deserializeUint256t(bytes);
}
void deserializeCodePointStub(std::vector<unsigned char>::const_iterator& bytes,
                              value* result,
                              std::vector<Slot>& slots) {
    auto segment_id = deserializeUint64t(bytes);
    auto pc = deserializeUint64t(bytes);
    auto next_hash = deserializeUint256t(bytes);
    CodeSegment segment = CodeSegment::uninitialized();
    *result = CodePointStub({segment, pc}, next_hash);
    slots.emplace_back(
        SlotPointer(&std::get_if<CodePointStub>(result)->pc.segment),
        segmentIdToDbHash(segment_id));
}
void deserializeTuple(std::vector<unsigned char>::const_iterator& bytes,
                      value* result,
                      std::vector<Slot>& slots,
                      size_t size) {
    *result = Tuple::createSizedTuple(size);
    Tuple& tup = std::get<Tuple>(*result);
    for (uint64_t i = 0; i < size; i++) {
        auto ty = *bytes;
        auto ptr = tup.getElementPointer(i);
        tup.markContentsWillChange();
        if (ty == HASH_PRE_IMAGE) {
            bytes++;
            auto hash = deserializeUint256t(bytes);
            slots.emplace_back(SlotPointer(ptr), hash);
        } else {
            deserializeValue(bytes, ptr, slots);
        }
    }
}
void deserializeBuffer(std::vector<unsigned char>::const_iterator& bytes,
                       value* result,
                       std::vector<Slot>& slots) {
    uint8_t depth = *bytes++;
    if (depth == 0) {
        Buffer::LeafData leaf;
        std::copy(bytes, bytes + 32, leaf.begin());
        bytes += 32;
        *result = Buffer(leaf);
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
            slots.emplace_back(SlotPointer(ptr), hash);
        }
        *result = Buffer(left, right);
    }
}
void deserializeCodeSegment(std::vector<unsigned char>::const_iterator& bytes,
                            value* result,
                            std::vector<Slot>& slots) {
    auto segment_id = deserializeUint64t(bytes);
    auto num_code_points = deserializeUint64t(bytes);
    std::vector<CodePoint> code;
    // This reserve is vital to ensure CodePoints (and thus immediates) don't
    // move due to the vector being reallocated to grow
    code.reserve(num_code_points);
    for (uint64_t i = 0; i < num_code_points; i++) {
        bool has_immediate = *bytes++;
        auto op = static_cast<OpCode>(*bytes++);
        auto next_hash = deserializeUint256t(bytes);
        std::optional<value> immediate;
        if (has_immediate) {
            immediate = value();
        }
        code.emplace_back(Operation(op, immediate), next_hash);
        if (has_immediate) {
            deserializeValue(bytes, &*code.back().op.immediate, slots);
        }
    }
    *result = CodeSegment::restoreCodeSegment(segment_id, code);
}

void deserializeValue(std::vector<unsigned char>::const_iterator& bytes,
                      value* result,
                      std::vector<Slot>& slots) {
    auto ty = *bytes++;
    switch (ty) {
        case BUFFER: {
            deserializeBuffer(bytes, result, slots);
            break;
        }
        case NUM: {
            deserializeNum(bytes, result, slots);
            break;
        }
        case CODE_POINT_STUB: {
            deserializeCodePointStub(bytes, result, slots);
            break;
        }
        case HASH_PRE_IMAGE: {
            throw std::runtime_error("Attempted to deserialize HASH_PRE_IMAGE");
            break;
        }
        case CODE_SEGMENT: {
            deserializeCodeSegment(bytes, result, slots);
            break;
        }
        default: {
            auto size = ty - TUPLE;
            if (size > 8) {
                throw std::runtime_error(
                    "attempted to deserialize value with invalid typecode");
            }
            deserializeTuple(bytes, result, slots, size);
            break;
        }
    }
}
