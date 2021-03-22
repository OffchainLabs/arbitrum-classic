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

#include <avm_values/codepoint.hpp>
#include <avm_values/codepointstub.hpp>
#include <avm_values/tuple.hpp>
#include <data_storage/value/serialize.hpp>

// Serialization

void serializeValue(const uint256_t& val, std::vector<unsigned char>& bytes) {
    bytes.push_back(NUM);
    marshal_uint256_t(val, bytes);
}
void serializeValue(const CodePointStub& val,
                    std::vector<unsigned char>& bytes) {
    bytes.push_back(CODE_POINT_STUB);
    val.marshal(bytes);
}
void serializeValue(const Tuple& val, std::vector<unsigned char>& bytes) {
    bytes.push_back(TUPLE + val.tuple_size());
    for (uint64_t i = 0; i < val.tuple_size(); i++) {
        auto inner = val.get_element_unsafe(i);
        // Unless the inner value is another tuple, serialize it inline
        // TODO: inline tuples at a given chance
        if (auto tup = std::get_if<Tuple>(&inner)) {
            bytes.push_back(HASH_PRE_IMAGE);
            marshal_uint256_t(hash(*tup), bytes);
        } else {
            serializeValue(inner, bytes);
        }
    }
}
void serializeValue(const HashPreImage&, std::vector<unsigned char>&) {
    throw std::runtime_error("Can't serialize hash preimage in db");
}
void serializeValue(const Buffer& b, std::vector<unsigned char>& bytes) {
    bytes.push_back(BUFFER);
    b.serialize(bytes);
}
void serializeValue(const CodeSegment& val, std::vector<unsigned char>& bytes) {
    bytes.push_back(CODE_SEGMENT);
    marshal_uint64_t(val.segmentID(), bytes);
    auto code = val.load();
    marshal_uint64_t(code.size(), bytes);
    for (const CodePoint& point : code) {
        bytes.push_back(
            static_cast<unsigned char>(point.op.immediate.has_value()));
        bytes.push_back(static_cast<unsigned char>(point.op.opcode));
        marshal_uint256_t(point.nextHash, bytes);
        if (point.op.immediate) {
            serializeValue(*point.op.immediate, bytes);
        }
    }
}

void serializeValue(const value& val, std::vector<unsigned char>& bytes) {
    std::visit([&](const auto& val) { serializeValue(val, bytes); }, val);
}

// Get dependencies, used for both saving and deleting

void getValueDependencies(const uint256_t&, std::vector<value>&) {}
void getValueDependencies(const CodePointStub& val,
                          std::vector<value>& dependencies) {
    dependencies.emplace_back(val.pc.segment);
}
void getValueDependencies(const Tuple& val, std::vector<value>& dependencies) {
    for (uint64_t i = 0; i < val.tuple_size(); i++) {
        auto inner = val.get_element_unsafe(i);
        // TODO: inline tuples at a given chance
        if (auto tup = std::get_if<Tuple>(&inner)) {
            dependencies.emplace_back(*tup);
        } else {
            getValueDependencies(inner, dependencies);
        }
    }
}
void getValueDependencies(const HashPreImage&, std::vector<value>&) {
    throw std::runtime_error("Can't serialize hash preimage in db");
}
void getValueDependencies(const Buffer& b, std::vector<value>& dependencies) {
    for (auto child : b.getDependencies()) {
        dependencies.emplace_back(child);
    }
}
void getValueDependencies(const CodeSegment& val,
                          std::vector<value>& dependencies) {
    auto code = val.load();
    for (const CodePoint& point : code) {
        if (point.op.immediate) {
            getValueDependencies(*point.op.immediate, dependencies);
        }
    }
}

void getValueDependencies(const value& val, std::vector<value>& dependencies) {
    std::visit(
        [&](const auto& val) { getValueDependencies(val, dependencies); }, val);
}
