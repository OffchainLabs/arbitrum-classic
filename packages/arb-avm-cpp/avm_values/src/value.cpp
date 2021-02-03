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

#include <ostream>

uint64_t deserialize_uint64_t(const char*& bufptr) {
    uint64_t val = intx::be::unsafe::load<uint64_t>(
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

value deserialize_value(const char*& bufptr) {
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
                values.push_back(value{deserializeUint256t(bufptr)});
                break;
            }
            case CODEPT: {
                values.push_back(value{deserializeCodePointStub(bufptr)});
                break;
            }
            default: {
                if (valType >= TUPLE && valType <= TUPLE + 8) {
                    uint8_t tuple_size = valType - TUPLE;
                    values_to_read += tuple_size;
                    values.push_back(TuplePlaceholder{tuple_size});
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

value assembleValueFromDeserialized(std::vector<DeserializedValue> values) {
    // Next form the full value out of the interleaved values and placeholders
    size_t total_values_size = values.size();
    for (size_t i = 0; i < total_values_size; ++i) {
        size_t val_pos = total_values_size - 1 - i;
        auto& val = values[val_pos];
        if (!nonstd::holds_alternative<TuplePlaceholder>(val)) {
            continue;
        }
        auto holder = val.get<TuplePlaceholder>();
        Tuple tup(holder.values);
        for (uint8_t j = 0; j < holder.values; ++j) {
            tup.set_element(j, std::move(values[val_pos + 1 + j].get<value>()));
        }
        values.erase(values.begin() + val_pos + 1,
                     values.begin() + val_pos + 1 + holder.values);
        values[val_pos] = std::move(tup);
    }
    return values.back().get<value>();
}

void marshal_uint64_t(uint64_t val, std::vector<unsigned char>& buf) {
    uint64_t big_endian_val = boost::endian::native_to_big(val);
    const unsigned char* data =
        reinterpret_cast<const unsigned char*>(&big_endian_val);
    buf.insert(buf.end(), data, data + sizeof(big_endian_val));
}

namespace {
struct Marshaller {
    std::vector<value>& values;
    std::vector<unsigned char>& buf;

    void operator()(const HashPreImage& val) const {
        buf.push_back(HASH_PRE_IMAGE);
        val.marshal(buf);
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
};
}  // namespace

void marshal_value(const value& full_val, std::vector<unsigned char>& buf) {
    std::vector<value> values{full_val};
    Marshaller marshaller{values, buf};
    while (!values.empty()) {
        const auto val = std::move(values.back());
        values.pop_back();
        nonstd::visit(marshaller, val);
    }
}

namespace {
void marshalForProof(const HashPreImage& val,
                     MarshalLevel,
                     std::vector<unsigned char>& buf,
                     const Code&) {
    buf.push_back(HASH_PRE_IMAGE);
    val.marshal(buf);
}

MarshalLevel childNestLevel(MarshalLevel level) {
    if (level == MarshalLevel::FULL) {
        return MarshalLevel::FULL;
    } else {
        return MarshalLevel::STUB;
    }
}

void marshalForProof(const Tuple& val,
                     MarshalLevel marshal_level,
                     std::vector<unsigned char>& buf,
                     const Code& code) {
    if (marshal_level == MarshalLevel::STUB) {
        marshalForProof(val.getHashPreImage(), marshal_level, buf, code);
    } else {
        buf.push_back(TUPLE + val.tuple_size());
        MarshalLevel nested_level = childNestLevel(marshal_level);
        for (uint64_t i = 0; i < val.tuple_size(); i++) {
            auto itemval = val.get_element(i);
            marshalForProof(itemval, nested_level, buf, code);
        }
    }
}

void marshalForProof(const CodePointStub& val,
                     MarshalLevel marshal_level,
                     std::vector<unsigned char>& buf,
                     const Code& code) {
    auto& cp = code.loadCodePoint(val.pc);
    buf.push_back(CODEPT);
    cp.op.marshalForProof(buf, childNestLevel(marshal_level), code);
    marshal_uint256_t(cp.nextHash, buf);
}

void marshalForProof(const uint256_t& val,
                     MarshalLevel,
                     std::vector<unsigned char>& buf,
                     const Code&) {
    buf.push_back(NUM);
    marshal_uint256_t(val, buf);
}

void marshalForProof(const Buffer& val,
                     MarshalLevel,
                     std::vector<unsigned char>& buf,
                     const Code&) {
    buf.push_back(BUFFER);
    marshal_uint256_t(val.hash(), buf);
}

}  // namespace

void marshalForProof(const value& val,
                     MarshalLevel marshal_level,
                     std::vector<unsigned char>& buf,
                     const Code& code) {
    return nonstd::visit(
        [&](const auto& v) {
            return marshalForProof(v, marshal_level, buf, code);
        },
        val);
}

uint256_t hash_value(const value& value) {
    return nonstd::visit([](const auto& val) { return hash(val); }, value);
}

struct GetSize {
    uint256_t operator()(const HashPreImage& val) const {
        return val.getSize();
    }

    uint256_t operator()(const Tuple& val) const { return val.getSize(); }

    uint256_t operator()(const Buffer&) const { return 1; }

    uint256_t operator()(const uint256_t&) const { return 1; }

    uint256_t operator()(const CodePointStub&) const { return 1; }
};

uint256_t getSize(const value& val) {
    return nonstd::visit(GetSize{}, val);
}

struct ValuePrinter {
    std::ostream& os;

    std::ostream* operator()(const Buffer& b) const {
        os << "Buffer(" << hash(b) << ")";
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

    std::ostream* operator()(const HashPreImage& val) const {
        os << val;
        return &os;
    }

    std::ostream* operator()(const CodePointStub& val) const {
        //        std::printf("in CodePoint ostream operator\n");
        os << "CodePointStub(" << val.pc.pc << ")";
        return &os;
    }
};

std::ostream& operator<<(std::ostream& os, const value& val) {
    return *nonstd::visit(ValuePrinter{os}, val);
}
