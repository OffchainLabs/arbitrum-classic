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

#ifndef value_hpp
#define value_hpp

#include <avm_values/bigint.hpp>
#include <avm_values/buffer.hpp>
#include <avm_values/codepoint.hpp>
#include <avm_values/opcodes.hpp>
#include <avm_values/tuple.hpp>
#include <avm_values/unloadedvalue.hpp>
#include <avm_values/valuetype.hpp>

#include <variant>

#ifndef __GNUC__
#ifndef __builtin_expect
#define __builtin_expect(x, y) (x)
#endif /* __builtin_expect */
#ifndef __builtin_unreachable
#define __builtin_unreachable()
#endif /* __builtin_unreachable */
#endif /* __GNU_C__ */

constexpr uint8_t value_tagged_bit = uint64_t(1) << 7;
constexpr uint8_t value_unloaded_bit = uint64_t(1) << 6;

constexpr uint8_t value_uint256_tag = value_tagged_bit | 0;
constexpr uint8_t value_tuple_tag = value_tagged_bit | 1;
constexpr uint8_t value_hash_pre_image_tag = value_tagged_bit | 2;
constexpr uint8_t value_buffer_tag = value_tagged_bit | 3;

union TaggedValueContents {
    uint256_t uint256;
    Tuple tuple;
    std::shared_ptr<HashPreImage> hash_pre_image;
    Buffer buffer;
};

struct TaggedValue {
    uint8_t tag;
    TaggedValueContents inner;
};

class Value {
   private:
    union ValueUnion {
        TaggedValue tagged;
        CodePointStub code_point;
        UnloadedValue unloaded;
        ~ValueUnion();
    };

    ValueUnion inner;

   public:
    Value();
    Value(Tuple);
    Value(uint256_t);
    Value(CodePointStub);
    explicit Value(std::shared_ptr<HashPreImage>);
    Value(Buffer);
    Value(UnloadedValue);

    ~Value();
    Value(const Value&);
    Value& operator=(const Value&);
    Value(Value&&);
    Value& operator=(Value&&);

    template <typename T>
    decltype(auto) visit(T visitor) const {
        if (__builtin_expect(!!(inner.tagged.tag & value_tagged_bit), 1))
            [[likely]] {
            switch (inner.tagged.tag) {
                case value_uint256_tag:
                    return visitor(inner.tagged.inner.uint256);
                case value_tuple_tag:
                    return visitor(inner.tagged.inner.tuple);
                case value_hash_pre_image_tag:
                    return visitor(inner.tagged.inner.hash_pre_image);
                case value_buffer_tag:
                    return visitor(inner.tagged.inner.buffer);
                default:
                    assert(0);
                    __builtin_unreachable();
                    throw std::runtime_error("Unknown value tag");
            }
        } else if (inner.tagged.tag & value_unloaded_bit) {
            return visitor(inner.unloaded);
        } else {
            return visitor(inner.code_point);
        }
    }

    template <typename T>
    decltype(auto) visit(T visitor) {
        if (__builtin_expect(!!(inner.tagged.tag & value_tagged_bit), 1))
            [[likely]] {
            switch (inner.tagged.tag) {
                case value_uint256_tag:
                    return visitor(inner.tagged.inner.uint256);
                case value_tuple_tag:
                    return visitor(inner.tagged.inner.tuple);
                case value_hash_pre_image_tag:
                    return visitor(inner.tagged.inner.hash_pre_image);
                case value_buffer_tag:
                    return visitor(inner.tagged.inner.buffer);
                default:
                    assert(0);
                    __builtin_unreachable();
                    throw std::runtime_error("Unknown value tag");
            }
        } else if (inner.tagged.tag & value_unloaded_bit) {
            return visitor(inner.unloaded);
        } else {
            return visitor(inner.code_point);
        }
    }
};

template <typename T>
bool holds_alternative(const Value&);
template <>
bool holds_alternative<Tuple>(const Value&);
template <>
bool holds_alternative<uint256_t>(const Value&);
template <>
bool holds_alternative<CodePointStub>(const Value&);
template <>
bool holds_alternative<std::shared_ptr<HashPreImage>>(const Value&);
template <>
bool holds_alternative<Buffer>(const Value&);
template <>
bool holds_alternative<UnloadedValue>(const Value&);

template <typename T>
T* get_if(Value*);
template <>
Tuple* get_if<Tuple>(Value*);
template <>
uint256_t* get_if<uint256_t>(Value*);
template <>
CodePointStub* get_if<CodePointStub>(Value*);
template <>
std::shared_ptr<HashPreImage>* get_if<std::shared_ptr<HashPreImage>>(Value*);
template <>
Buffer* get_if<Buffer>(Value*);
template <>
UnloadedValue* get_if<UnloadedValue>(Value*);

template <typename T>
const T* get_if(const Value*);
template <>
const Tuple* get_if<Tuple>(const Value*);
template <>
const uint256_t* get_if<uint256_t>(const Value*);
template <>
const CodePointStub* get_if<CodePointStub>(const Value*);
template <>
const std::shared_ptr<HashPreImage>* get_if<std::shared_ptr<HashPreImage>>(
    const Value*);
template <>
const Buffer* get_if<Buffer>(const Value*);
template <>
const UnloadedValue* get_if<UnloadedValue>(const Value*);

template <typename T>
T& get(Value&);
template <>
Tuple& get<Tuple>(Value&);
template <>
uint256_t& get<uint256_t>(Value&);
template <>
CodePointStub& get<CodePointStub>(Value&);
template <>
std::shared_ptr<HashPreImage>& get<std::shared_ptr<HashPreImage>>(Value&);
template <>
Buffer& get<Buffer>(Value&);
template <>
UnloadedValue& get<UnloadedValue>(Value&);

template <typename T>
const T& get(const Value&);
template <>
const Tuple& get<Tuple>(const Value&);
template <>
const uint256_t& get<uint256_t>(const Value&);
template <>
const CodePointStub& get<CodePointStub>(const Value&);
template <>
const std::shared_ptr<HashPreImage>& get<std::shared_ptr<HashPreImage>>(
    const Value&);
template <>
const Buffer& get<Buffer>(const Value&);
template <>
const UnloadedValue& get<UnloadedValue>(const Value&);

template <typename T>
decltype(auto) visit(T visitor, const Value& val) {
    return val.visit(visitor);
}

template <typename T>
decltype(auto) visit(T visitor, Value& val) {
    return val.visit(visitor);
}

struct TuplePlaceholder {
    uint8_t values;
};
using DeserializedValue = std::variant<TuplePlaceholder, Value>;

std::ostream& operator<<(std::ostream& os, const Value& val);
uint256_t hash_value(const Value& value);
bool values_equal(const Value& a, const Value& b);

uint64_t deserialize_uint64_t(const char*& bufptr);
CodePointRef deserializeCodePointRef(const char*& bufptr);
CodePointStub deserializeCodePointStub(const char*& bufptr);
uint256_t deserializeUint256t(const char*& srccode);
Value deserialize_value(const char*& srccode);

void marshal_uint64_t(uint64_t val, std::vector<unsigned char>& buf);

void marshal_value(const Value& val, std::vector<unsigned char>& buf);

class Code;

void marshalForProof(const Value& val,
                     size_t marshal_level,
                     std::vector<unsigned char>& buf,
                     const Code& code);

uint256_t getSize(const Value& val);

Value assembleValueFromDeserialized(std::vector<DeserializedValue> values);

#endif /* value_hpp */
