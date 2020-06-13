/*
 * Copyright 2019, Offchain Labs, Inc.
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

#include <iterator>

#include <data_storage/checkpoint/checkpointutils.hpp>

#include <avm_values/codepoint.hpp>
#include <avm_values/tuple.hpp>
#include <avm_values/util.hpp>

#include <bigint_utils.hpp>

constexpr int UINT64_SIZE = 8;
constexpr int HASH_KEY_LENGTH = 32;
constexpr int TUP_TUPLE_LENGTH = 33;
constexpr int TUP_NUM_LENGTH = 33;
constexpr int TUP_CODEPT_LENGTH = 41;

std::unordered_map<int, int> blockreason_type_length = {{0, 1},
                                                        {1, 1},
                                                        {2, 1},
                                                        {3, 1},
                                                        {4, 34}};

namespace checkpoint {

uint64_t deserialize_uint64(const char*& bufptr) {
    uint64_t ret_value;
    memcpy(&ret_value, bufptr, UINT64_SIZE);
    auto val = boost::endian::big_to_native(ret_value);
    bufptr += sizeof(uint64_t);
    return val;
}

void marshal_uint64_t(uint64_t val, std::vector<unsigned char>& buf) {
    auto big_endian_val = boost::endian::native_to_big(val);
    std::array<unsigned char, UINT64_SIZE> tmpbuf;
    memcpy(tmpbuf.data(), &big_endian_val, sizeof(big_endian_val));

    buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
}

void serializeCodePointStub(const CodePointStub& val,
                            std::vector<unsigned char>& value_vector) {
    marshal_uint64_t(val.pc, value_vector);
    marshal_uint256_t(val.hash, value_vector);
}

struct ValueSerializer {
    std::vector<unsigned char> operator()(const Tuple& val) const {
        std::vector<unsigned char> value_vector;
        value_vector.push_back(TUPLE);
        auto hash_key = hash_value(val);
        marshal_uint256_t(hash_key, value_vector);

        return value_vector;
    }

    std::vector<unsigned char> operator()(const uint256_t& val) const {
        std::vector<unsigned char> value_vector;
        value_vector.push_back(NUM);
        marshal_uint256_t(val, value_vector);

        return value_vector;
    }

    std::vector<unsigned char> operator()(const CodePointStub& val) const {
        std::vector<unsigned char> value_vector;
        auto type_code = static_cast<unsigned char>(CODEPT);
        value_vector.push_back(type_code);
        serializeCodePointStub(val, value_vector);
        return value_vector;
    }

    std::vector<unsigned char> operator()(const HashPreImage& val) const {
        std::vector<unsigned char> value_vector;
        auto type_code = static_cast<unsigned char>(HASH_PRE_IMAGE);
        value_vector.push_back(type_code);
        val.marshal(value_vector);

        return value_vector;
    }
};

namespace utils {

std::vector<unsigned char> GetHashKey(const value& val) {
    auto hash_key = hash_value(val);
    std::vector<unsigned char> hash_key_vector;
    marshal_uint256_t(hash_key, hash_key_vector);

    return hash_key_vector;
}

std::vector<std::vector<unsigned char>> parseTuple(
    const std::vector<unsigned char>& data) {
    std::vector<std::vector<unsigned char>> return_vector;

    auto iter = data.begin();
    uint8_t count = *iter - TUPLE;
    ++iter;

    for (uint8_t i = 0; i < count; i++) {
        auto value_type = static_cast<ValueTypes>(*iter);
        std::vector<unsigned char> current;

        switch (value_type) {
            case NUM: {
                auto next_it = iter + TUP_NUM_LENGTH;
                current.insert(current.end(), iter, next_it);
                iter = next_it;
                break;
            }
            case CODEPT: {
                auto next_it = iter + TUP_CODEPT_LENGTH;
                current.insert(current.end(), iter, next_it);
                iter = next_it;
                break;
            }
            case HASH_PRE_IMAGE: {
                throw std::runtime_error("HASH_ONLY item");
            }
            default: {
                uint8_t tup_size = value_type - TUPLE;
                if (tup_size > 8) {
                    throw std::runtime_error(
                        "tried to parse tuple with invalid typecode");
                }
                auto next_it = iter + TUP_TUPLE_LENGTH;
                current.insert(current.end(), iter, next_it);
                iter = next_it;
                break;
            }
        }
        return_vector.push_back(current);
    }
    return return_vector;
}

CodePointStub deserializeCodePointStub(const char*& bufptr) {
    auto pc_val = deserialize_uint64(bufptr);
    auto hash_val = deserializeUint256t(bufptr);
    return {pc_val, hash_val};
}

uint256_t deserializeUint256_t(const std::vector<unsigned char>& val) {
    auto buff = reinterpret_cast<const char*>(&val[1]);
    return deserializeUint256t(buff);
}

std::vector<unsigned char> serializeValue(const value& val) {
    return nonstd::visit(ValueSerializer{}, val);
}

using iterator = std::vector<unsigned char>::const_iterator;

unsigned char extractStatus(iterator& iter) {
    auto status = (unsigned char)(*iter);
    ++iter;

    return status;
}

std::vector<unsigned char> extractHashKey(iterator& iter) {
    auto end_iter = iter + HASH_KEY_LENGTH;
    std::vector<unsigned char> hash_key(iter, end_iter);
    iter = end_iter;

    return hash_key;
}

CodePointStub extractCodePointStub(iterator& iter) {
    auto ptr = reinterpret_cast<const char*>(&*iter);
    auto pc_val = deserialize_uint64(ptr);
    auto hash_val = deserializeUint256t(ptr);
    iter += sizeof(pc_val) + 32;
    return {pc_val, hash_val};
}

MachineStateKeys extractStateKeys(
    const std::vector<unsigned char>& stored_state) {
    auto current_iter = stored_state.begin();

    auto status = extractStatus(current_iter);

    auto register_val = extractHashKey(current_iter);
    auto datastack = extractHashKey(current_iter);
    auto auxstack = extractHashKey(current_iter);
    auto pc = extractCodePointStub(current_iter);
    auto err_pc = extractCodePointStub(current_iter);

    return MachineStateKeys{register_val, datastack, auxstack,
                            pc,           err_pc,    status};
}

std::vector<unsigned char> serializeStateKeys(
    const MachineStateKeys& state_data) {
    std::vector<unsigned char> state_data_vector;
    state_data_vector.push_back(state_data.status_char);

    state_data_vector.insert(state_data_vector.end(),
                             state_data.register_val_key.begin(),
                             state_data.register_val_key.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.datastack_key.begin(),
                             state_data.datastack_key.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.auxstack_key.begin(),
                             state_data.auxstack_key.end());

    serializeCodePointStub(state_data.pc, state_data_vector);
    serializeCodePointStub(state_data.err_pc, state_data_vector);
    return state_data_vector;
}
}  // namespace utils
}  // namespace checkpoint
