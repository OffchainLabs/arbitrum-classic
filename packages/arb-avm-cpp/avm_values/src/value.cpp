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
#include <avm_values/codepointstub.hpp>
#include <avm_values/pool.hpp>
#include <avm_values/tuple.hpp>
#include <avm_values/value.hpp>

#include <boost/endian/conversion.hpp>

#include <iomanip>
#include <ostream>

Value::Value() : Value(Tuple()) {}
Value::Value(Tuple tup) : Value(0) {
    inner.tagged.tag = value_tuple_tag;
    inner.tagged.inner.tuple = std::move(tup);
}
Value::Value(uint64_t num) : Value(uint256_t(num)) {}
Value::Value(uint256_t num)
    : inner{TaggedValue{value_num_tag, TaggedValueContents{num}}} {}
Value::Value(CodePointStub code_point) : Value(0) {
    inner.code_point = std::move(code_point);
    assert(!(inner.tagged.tag & value_tagged_bit));
    assert(!(inner.tagged.tag & value_unloaded_bit));
}
Value::Value(std::shared_ptr<HashPreImage> hash_pre_image) : Value(0) {
    inner.tagged.tag = value_hash_pre_image_tag;
    inner.tagged.inner.hash_pre_image = std::move(hash_pre_image);
}
Value::Value(Buffer buffer) : Value(0) {
    inner.tagged.tag = value_buffer_tag;
    inner.tagged.inner.buffer = std::move(buffer);
}
Value::Value(UnloadedValue uv) : Value(0) {
    inner.unloaded = uv;
    assert(!(inner.tagged.tag & value_tagged_bit));
    assert(inner.tagged.tag & value_unloaded_bit);
}

Value::~Value() {
    if (isTagged()) [[likely]] {
        // No need to destruct a uint256. Select for the other tags.
        switch (inner.tagged.tag) {
            case value_tuple_tag:
                inner.tagged.inner.tuple.~Tuple();
                break;
            case value_hash_pre_image_tag:
                inner.tagged.inner.hash_pre_image.~shared_ptr();
                break;
            case value_buffer_tag:
                inner.tagged.inner.buffer.~Buffer();
                break;
        }
    } else if (inner.tagged.tag & value_unloaded_bit) {
        inner.unloaded.~UnloadedValue();
    } else {
        inner.code_point.~CodePointStub();
    }
}

Value::Value(const Value& other) : Value(0) {
    if (other.isTagged()) [[likely]] {
        switch (inner.tagged.tag) {
            case value_num_tag:
                inner.tagged.inner.num = other.inner.tagged.inner.num;
                break;
            case value_tuple_tag:
                inner.tagged.inner.tuple = other.inner.tagged.inner.tuple;
                break;
            case value_hash_pre_image_tag:
                inner.tagged.inner.hash_pre_image =
                    other.inner.tagged.inner.hash_pre_image;
                break;
            case value_buffer_tag:
                inner.tagged.inner.buffer = other.inner.tagged.inner.buffer;
                break;
            default:
                assert(0);
                __builtin_unreachable();
                throw std::runtime_error("Unknown value tag");
        }
    } else if (inner.tagged.tag & value_unloaded_bit) {
        inner.unloaded = other.inner.unloaded;
    } else {
        inner.code_point = other.inner.code_point;
    }
}

Value& Value::operator=(const Value& other) {
    *this = Value(other);
    return *this;
}

Value::Value(Value&& other) : Value(0) {
    if (other.isTagged()) [[likely]] {
        switch (inner.tagged.tag) {
            case value_num_tag:
                inner.tagged.inner.num =
                    std::move(other.inner.tagged.inner.num);
                break;
            case value_tuple_tag:
                inner.tagged.inner.tuple =
                    std::move(other.inner.tagged.inner.tuple);
                break;
            case value_hash_pre_image_tag:
                inner.tagged.inner.hash_pre_image =
                    std::move(other.inner.tagged.inner.hash_pre_image);
                break;
            case value_buffer_tag:
                inner.tagged.inner.buffer =
                    std::move(other.inner.tagged.inner.buffer);
                break;
            default:
                assert(0);
                __builtin_unreachable();
                throw std::runtime_error("Unknown value tag");
        }
    } else if (inner.tagged.tag & value_unloaded_bit) {
        inner.unloaded = std::move(other.inner.unloaded);
    } else {
        inner.code_point = std::move(other.inner.code_point);
    }
}

Value& Value::operator=(Value&& other) {
    if (other.isTagged()) [[likely]] {
        switch (inner.tagged.tag) {
            case value_num_tag:
                std::swap(inner.tagged.inner.num, other.inner.tagged.inner.num);
                break;
            case value_tuple_tag:
                std::swap(inner.tagged.inner.tuple,
                          other.inner.tagged.inner.tuple);
                break;
            case value_hash_pre_image_tag:
                std::swap(inner.tagged.inner.hash_pre_image,
                          other.inner.tagged.inner.hash_pre_image);
                break;
            case value_buffer_tag:
                std::swap(inner.tagged.inner.buffer,
                          other.inner.tagged.inner.buffer);
                break;
            default:
                assert(0);
                __builtin_unreachable();
                throw std::runtime_error("Unknown value tag");
        }
    } else if (inner.tagged.tag & value_unloaded_bit) {
        std::swap(inner.unloaded, other.inner.unloaded);
    } else {
        std::swap(inner.code_point, other.inner.code_point);
    }
    return *this;
}

uint64_t deserialize_uint64_t(const char*& bufptr) {
    auto val = intx::be::unsafe::load<uint64_t>(
        reinterpret_cast<const unsigned char*>(bufptr));
    bufptr += sizeof(val);
    return val;
}

CodePointRef deserializeCodePointRef(const char*& bufptr) {
    uint64_t segment = deserialize_uint64_t(bufptr);
    uint64_t pc = deserialize_uint64_t(bufptr);
    return {segment, pc};
}

CodePointStub deserializeCodePointStub(const char*& bufptr) {
    auto ref = deserializeCodePointRef(bufptr);
    auto hash_val = deserializeUint256t(bufptr);
    return {ref, hash_val};
}

uint256_t deserializeUint256t(const char*& bufptr) {
    auto ret = intx::be::unsafe::load<uint256_t>(
        reinterpret_cast<const unsigned char*>(bufptr));
    bufptr += 32;
    return ret;
}

Value deserialize_value(const char*& bufptr) {
    // Iteratively read all values leaving placeholder for the tuples
    std::vector<DeserializedValue> values;
    uint64_t values_to_read = 1;
    while (values_to_read > 0) {
        uint8_t valType;
        memcpy(&valType, bufptr, sizeof(valType));
        bufptr += sizeof(valType);
        --values_to_read;
        switch (valType) {
            case NUM: {
                values.push_back(Value{deserializeUint256t(bufptr)});
                break;
            }
            case CODEPT: {
                values.push_back(Value{deserializeCodePointStub(bufptr)});
                break;
            }
            default: {
                if (valType >= TUPLE && valType <= TUPLE + 8) {
                    uint8_t tuple_size = valType - TUPLE;
                    values_to_read += tuple_size;
                    values.emplace_back(TuplePlaceholder{tuple_size});
                } else {
                    std::printf("in deserialize_value, unhandled type = %X\n",
                                valType);
                    throw std::runtime_error(
                        "Tried to deserialize unhandled type");
                }
                break;
            }
        }
    }
    return assembleValueFromDeserialized(std::move(values));
}

Value assembleValueFromDeserialized(std::vector<DeserializedValue> values) {
    // Next form the full value out of the interleaved values and placeholders
    size_t total_values_size = values.size();
    for (size_t i = 0; i < total_values_size; ++i) {
        size_t val_pos = total_values_size - 1 - i;
        auto& val = values[val_pos];
        if (!std::holds_alternative<TuplePlaceholder>(val)) {
            continue;
        }
        auto holder = std::get<TuplePlaceholder>(val);
        Tuple tup = Tuple::createSizedTuple(holder.values);
        for (uint8_t j = 0; j < holder.values; ++j) {
            tup.set_element(
                j, std::move(std::get<Value>(values[val_pos + 1 + j])));
        }
        values.erase(values.begin() + val_pos + 1,
                     values.begin() + val_pos + 1 + holder.values);
        values[val_pos] = std::move(tup);
    }
    return std::get<Value>(values.back());
}

void marshal_uint64_t(uint64_t val, std::vector<unsigned char>& buf) {
    uint64_t big_endian_val = boost::endian::native_to_big(val);
    const auto data = reinterpret_cast<const unsigned char*>(&big_endian_val);
    buf.insert(buf.end(), data, data + sizeof(big_endian_val));
}

namespace {
struct Marshaller {
    std::vector<Value>& values;
    std::vector<unsigned char>& buf;

    void operator()(const std::shared_ptr<HashPreImage>& val) const {
        buf.push_back(HASH_PRE_IMAGE);
        val->marshal(buf);
    }

    void operator()(const Tuple& tup) const {
        auto size = tup.tuple_size();
        buf.push_back(TUPLE + size);
        // queue elements in reverse order for serialization
        for (uint64_t i = 0; i < size; i++) {
            values.push_back(tup.get_element(size - 1 - i));
        }
    }

    void operator()(const Buffer& val) const {
        buf.push_back(BUFFER);
        auto data = val.toFlatVector();
        marshal_uint64_t(static_cast<uint64_t>(data.size()), buf);
        buf.insert(buf.end(), data.begin(), data.end());
    }

    void operator()(const uint256_t& val) const {
        buf.push_back(NUM);
        marshal_uint256_t(val, buf);
    }

    void operator()(const CodePointStub& val) const {
        buf.push_back(CODE_POINT_STUB);
        val.marshal(buf);
    }

    void operator()(const UnloadedValue&) const {
        throw std::runtime_error("Cannot marshal unloaded value");
    }
};
}  // namespace

void marshal_value(const Value& full_val, std::vector<unsigned char>& buf) {
    std::vector<Value> values{full_val};
    Marshaller marshaller{values, buf};
    while (!values.empty()) {
        const auto val = std::move(values.back());
        values.pop_back();
        visit(marshaller, val);
    }
}

namespace {
void marshalForProof(const HashPreImage& val,
                     size_t,
                     std::vector<unsigned char>& buf,
                     const Code&) {
    buf.push_back(HASH_PRE_IMAGE);
    val.marshal(buf);
}

void marshalForProof(const std::shared_ptr<HashPreImage>& val,
                     size_t,
                     std::vector<unsigned char>& buf,
                     const Code&) {
    buf.push_back(HASH_PRE_IMAGE);
    val->marshal(buf);
}

size_t childNestLevel(size_t level) {
    if (level > 0) {
        return level - 1;
    } else {
        return 0;
    }
}

void marshalForProof(const Tuple& val,
                     size_t marshal_level,
                     std::vector<unsigned char>& buf,
                     const Code& code) {
    if (marshal_level == 0) {
        marshalForProof(val.getHashPreImage(), marshal_level, buf, code);
    } else {
        buf.push_back(TUPLE + val.tuple_size());
        size_t nested_level = childNestLevel(marshal_level);
        for (uint64_t i = 0; i < val.tuple_size(); i++) {
            auto itemval = val.get_element(i);
            marshalForProof(itemval, nested_level, buf, code);
        }
    }
}

void marshalForProof(const CodePointStub& val,
                     size_t marshal_level,
                     std::vector<unsigned char>& buf,
                     const Code& code) {
    auto cp = code.loadCodePoint(val.pc);
    buf.push_back(CODEPT);
    cp.op.marshalForProof(buf, childNestLevel(marshal_level), code);
    marshal_uint256_t(cp.nextHash, buf);
}

void marshalForProof(const uint256_t& val,
                     size_t,
                     std::vector<unsigned char>& buf,
                     const Code&) {
    buf.push_back(NUM);
    marshal_uint256_t(val, buf);
}

void marshalForProof(const Buffer& val,
                     size_t,
                     std::vector<unsigned char>& buf,
                     const Code&) {
    buf.push_back(BUFFER);
    marshal_uint256_t(val.hash(), buf);
}

}  // namespace

void marshalForProof(const Value& val,
                     size_t marshal_level,
                     std::vector<unsigned char>& buf,
                     const Code& code) {
    return visit(
        [&](const auto& v) {
            return marshalForProof(v, marshal_level, buf, code);
        },
        val);
}

uint256_t hash_value(const Value& value) {
    return visit([](const auto& val) { return hash(val); }, value);
}

bool values_equal(const Value& a, const Value& b) {
    // Fast path: if the values are both ints, compare them directly
    {
        const uint256_t* a_int = get_if<uint256_t>(&a);
        const uint256_t* b_int = get_if<uint256_t>(&b);
        if (a_int && b_int) {
            return *a_int == *b_int;
        }
    }
    // Fast path: if the values are tuples of different sizes, return false
    {
        const Tuple* a_tup = get_if<Tuple>(&a);
        const Tuple* b_tup = get_if<Tuple>(&b);
        if (a_tup && b_tup && a_tup->tuple_size() != b_tup->tuple_size()) {
            return false;
        }
    }
    // Fast path: if the values are of different types, return false
    // Note: ValueTypeVisitor correctly sees through unloaded values
    if (visit(ValueTypeVisitor{}, a) != visit(ValueTypeVisitor{}, b)) {
        return false;
    }
    // Slow path: the preconditions for the fast paths weren't met
    // Check if the hashes are equal
    return hash_value(a) == hash_value(b);
}

struct GetSize {
    uint256_t operator()(const std::shared_ptr<HashPreImage>& val) const {
        return val->getSize();
    }

    uint256_t operator()(const Tuple& val) const { return val.getSize(); }

    uint256_t operator()(const Buffer&) const { return 1; }

    uint256_t operator()(const uint256_t&) const { return 1; }

    uint256_t operator()(const CodePointStub&) const { return 1; }

    uint256_t operator()(const UnloadedValue& val) const {
        return val.value_size();
    }
};

uint256_t getSize(const Value& val) {
    return visit(GetSize{}, val);
}

struct ValuePrinter {
    std::ostream& os;

    std::ostream* operator()(const Buffer& b) const {
        os << "Buffer(";
        if (b.data_length() <= 64) {
            os << "0x";
            std::ios prev_flags(nullptr);
            prev_flags.copyfmt(os);
            for (auto b : b.toFlatVector()) {
                os << std::hex << std::setw(2) << std::setfill('0') << (int)b;
            }
            os.copyfmt(prev_flags);
        } else {
            os << "hash ";
            os << b.hash();
        }
        os << ")";
        return &os;
    }

    std::ostream* operator()(const Tuple& val) const {
        os << val;
        return &os;
    }

    std::ostream* operator()(const uint256_t& val) const {
        os << intx::to_string(val);
        return &os;
    }

    std::ostream* operator()(const std::shared_ptr<HashPreImage>& val) const {
        os << *val;
        return &os;
    }

    std::ostream* operator()(const CodePointStub& val) const {
        //        std::printf("in CodePoint ostream operator\n");
        os << "CodePointStub(" << val.pc.pc << ")";
        return &os;
    }

    std::ostream* operator()(const UnloadedValue& val) const {
        os << "UnloadedValue(type " << val.type() << ", hash " << ::hash(val)
           << ")";
        return &os;
    }
};

std::ostream& operator<<(std::ostream& os, const Value& val) {
    return *visit(ValuePrinter{os}, val);
}
