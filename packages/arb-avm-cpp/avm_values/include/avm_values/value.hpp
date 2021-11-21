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

constexpr uint64_t value_tagged_bit = uint64_t(1) << 63;
constexpr uint64_t value_unloaded_bit = unloaded_value_fixed_bit;  // 1 << 62
constexpr uint64_t value_tag_flag_bits = value_tagged_bit | value_unloaded_bit;

constexpr uint64_t value_num_tag = value_tagged_bit | 0;
constexpr uint64_t value_tuple_tag = value_tagged_bit | 1;
constexpr uint64_t value_hash_pre_image_tag = value_tagged_bit | 2;
constexpr uint64_t value_buffer_tag = value_tagged_bit | 3;

union TaggedValueContents {
    uint256_t num;
    Tuple tuple;
    std::shared_ptr<HashPreImage> hash_pre_image;
    Buffer buffer;
    ~TaggedValueContents() {}
};

struct TaggedValue {
    uint64_t tag;
    TaggedValueContents inner;
    ~TaggedValue() {}
};

class Value {
    template <typename T>
    friend bool holds_alternative(const Value&);

    template <typename T>
    friend T* get_if(Value*);
    template <typename T>
    friend const T* get_if(const Value*);

    template <typename T>
    friend T& get(Value&);
    template <typename T>
    friend const T& get(const Value&);

    template <typename T>
    friend decltype(auto) visit(T, const Value&);
    template <typename T>
    friend decltype(auto) visit(T, Value&);

   private:
    union ValueUnion {
        TaggedValue tagged;
        CodePointStub code_point;
        UnloadedValue unloaded;
        ~ValueUnion() {}
    };

    ValueUnion inner;

    inline bool isTagged() const {
        return __builtin_expect(!!(inner.tagged.tag & value_tagged_bit), 1);
    }

   public:
    Value();
    Value(Tuple);
    Value(uint64_t);
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
    decltype(auto) visit(T visitor) const {}

    template <typename T>
    decltype(auto) visit(T visitor) {
        if (isTagged()) [[likely]] {
            switch (inner.tagged.tag) {
                case value_num_tag:
                    return visitor(inner.tagged.inner.num);
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

// Make sure we notice if we increase the size of Value
static_assert(sizeof(Value) == 48);

template <typename T>
bool holds_alternative(const Value&);
template <>
inline bool holds_alternative<Tuple>(const Value& val) {
    return val.inner.tagged.tag == value_tuple_tag;
}
template <>
inline bool holds_alternative<uint256_t>(const Value& val) {
    return val.inner.tagged.tag == value_num_tag;
}
template <>
inline bool holds_alternative<CodePointStub>(const Value& val) {
    return (val.inner.tagged.tag & value_tag_flag_bits) == 0;
}
template <>
inline bool holds_alternative<std::shared_ptr<HashPreImage>>(const Value& val) {
    return val.inner.tagged.tag == value_hash_pre_image_tag;
}
template <>
inline bool holds_alternative<Buffer>(const Value& val) {
    return val.inner.tagged.tag == value_buffer_tag;
}
template <>
inline bool holds_alternative<UnloadedValue>(const Value& val) {
    return (val.inner.tagged.tag & value_tag_flag_bits) == value_unloaded_bit;
}

template <typename T>
T& get(Value&);
template <>
inline Tuple& get<Tuple>(Value& val) {
    if (__builtin_expect(holds_alternative<Tuple>(val), 1)) [[likely]] {
        return val.inner.tagged.inner.tuple;
    } else {
        throw std::runtime_error(
            std::string("Attempted to get<Tuple> a Value with tag ") +
            std::to_string(val.inner.tagged.tag));
    }
}
template <>
inline uint256_t& get<uint256_t>(Value& val) {
    if (__builtin_expect(holds_alternative<uint256_t>(val), 1)) [[likely]] {
        return val.inner.tagged.inner.num;
    } else {
        throw std::runtime_error(
            std::string("Attempted to get<uint256_t> a Value with tag ") +
            std::to_string(val.inner.tagged.tag));
    }
}
template <>
inline CodePointStub& get<CodePointStub>(Value& val) {
    if (__builtin_expect(holds_alternative<CodePointStub>(val), 1)) [[likely]] {
        return val.inner.code_point;
    } else {
        throw std::runtime_error(
            std::string("Attempted to get<CodePointStub> a Value with tag ") +
            std::to_string(val.inner.tagged.tag));
    }
}
template <>
inline std::shared_ptr<HashPreImage>& get<std::shared_ptr<HashPreImage>>(
    Value& val) {
    if (__builtin_expect(holds_alternative<std::shared_ptr<HashPreImage>>(val),
                         1)) [[likely]] {
        return val.inner.tagged.inner.hash_pre_image;
    } else {
        throw std::runtime_error(
            std::string("Attempted to get<std::shared_ptr<HashPreImage>> a "
                        "Value with tag ") +
            std::to_string(val.inner.tagged.tag));
    }
}
template <>
inline Buffer& get<Buffer>(Value& val) {
    if (__builtin_expect(holds_alternative<Buffer>(val), 1)) [[likely]] {
        return val.inner.tagged.inner.buffer;
    } else {
        throw std::runtime_error(
            std::string("Attempted to get<Buffer> a Value with tag ") +
            std::to_string(val.inner.tagged.tag));
    }
}
template <>
inline UnloadedValue& get<UnloadedValue>(Value& val) {
    if (__builtin_expect(holds_alternative<UnloadedValue>(val), 1)) [[likely]] {
        return val.inner.unloaded;
    } else {
        throw std::runtime_error(
            std::string("Attempted to get<UnloadedValue> a Value with tag ") +
            std::to_string(val.inner.tagged.tag));
    }
}

template <typename T>
const T& get(const Value&);
template <>
inline const Tuple& get<Tuple>(const Value& val) {
    if (__builtin_expect(holds_alternative<Tuple>(val), 1)) [[likely]] {
        return val.inner.tagged.inner.tuple;
    } else {
        throw std::runtime_error(
            std::string("Attempted to get<Tuple> a Value with tag ") +
            std::to_string(val.inner.tagged.tag));
    }
}
template <>
inline const uint256_t& get<uint256_t>(const Value& val) {
    if (__builtin_expect(holds_alternative<uint256_t>(val), 1)) [[likely]] {
        return val.inner.tagged.inner.num;
    } else {
        throw std::runtime_error(
            std::string("Attempted to get<uint256_t> a Value with tag ") +
            std::to_string(val.inner.tagged.tag));
    }
}
template <>
inline const CodePointStub& get<CodePointStub>(const Value& val) {
    if (__builtin_expect(holds_alternative<CodePointStub>(val), 1)) [[likely]] {
        return val.inner.code_point;
    } else {
        throw std::runtime_error(
            std::string("Attempted to get<CodePointStub> a Value with tag ") +
            std::to_string(val.inner.tagged.tag));
    }
}
template <>
inline const std::shared_ptr<HashPreImage>& get<std::shared_ptr<HashPreImage>>(
    const Value& val) {
    if (__builtin_expect(holds_alternative<std::shared_ptr<HashPreImage>>(val),
                         1)) [[likely]] {
        return val.inner.tagged.inner.hash_pre_image;
    } else {
        throw std::runtime_error(
            std::string("Attempted to get<std::shared_ptr<HashPreImage>> a "
                        "Value with tag ") +
            std::to_string(val.inner.tagged.tag));
    }
}
template <>
inline const Buffer& get<Buffer>(const Value& val) {
    if (__builtin_expect(holds_alternative<Buffer>(val), 1)) [[likely]] {
        return val.inner.tagged.inner.buffer;
    } else {
        throw std::runtime_error(
            std::string("Attempted to get<Buffer> a Value with tag ") +
            std::to_string(val.inner.tagged.tag));
    }
}
template <>
inline const UnloadedValue& get<UnloadedValue>(const Value& val) {
    if (__builtin_expect(holds_alternative<UnloadedValue>(val), 1)) [[likely]] {
        return val.inner.unloaded;
    } else {
        throw std::runtime_error(
            std::string("Attempted to get<UnloadedValue> a Value with tag ") +
            std::to_string(val.inner.tagged.tag));
    }
}

template <typename T>
inline T* get_if(Value* val) {
    if (val && holds_alternative<T>(*val)) {
        return &get<T>(*val);
    } else {
        return nullptr;
    }
}

template <typename T>
inline const T* get_if(const Value* val) {
    if (val && holds_alternative<T>(*val)) {
        return &get<T>(*val);
    } else {
        return nullptr;
    }
}

template <typename T>
decltype(auto) visit(T visitor, const Value& val) {
    if (val.isTagged()) [[likely]] {
        switch (val.inner.tagged.tag) {
            case value_num_tag:
                return visitor(val.inner.tagged.inner.num);
            case value_tuple_tag:
                return visitor(val.inner.tagged.inner.tuple);
            case value_hash_pre_image_tag:
                return visitor(val.inner.tagged.inner.hash_pre_image);
            case value_buffer_tag:
                return visitor(val.inner.tagged.inner.buffer);
            default:
                assert(0);
                __builtin_unreachable();
                throw std::runtime_error("Unknown value tag");
        }
    } else if (val.inner.tagged.tag & value_unloaded_bit) {
        return visitor(val.inner.unloaded);
    } else {
        return visitor(val.inner.code_point);
    }
}

template <typename T>
decltype(auto) visit(T visitor, Value& val) {
    if (val.isTagged()) [[likely]] {
        switch (val.inner.tagged.tag) {
            case value_num_tag:
                return visitor(val.inner.tagged.inner.num);
            case value_tuple_tag:
                return visitor(val.inner.tagged.inner.tuple);
            case value_hash_pre_image_tag:
                return visitor(val.inner.tagged.inner.hash_pre_image);
            case value_buffer_tag:
                return visitor(val.inner.tagged.inner.buffer);
            default:
                assert(0);
                __builtin_unreachable();
                throw std::runtime_error("Unknown value tag");
        }
    } else if (val.inner.tagged.tag & value_unloaded_bit) {
        return visitor(val.inner.unloaded);
    } else {
        return visitor(val.inner.code_point);
    }
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
