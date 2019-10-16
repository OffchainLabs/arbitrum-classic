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

#include "avm/machinestate/statesaverutils.hpp"
#include <iterator>
#include "avm/machinestate/blockreason.hpp"
#include "avm/machinestate/tokenTracker.hpp"
#include "avm/value/codepoint.hpp"
#include "avm/value/tuple.hpp"
#include "bigint_utils.hpp"
#include "util.hpp"

#define UINT64_SIZE 8
#define HASH_KEY_LENGTH 33
#define TUP_TUPLE_LENGTH 34
#define TUP_NUM_LENGTH 34
#define TUP_CODEPT_LENGTH 9

namespace StateSaverUtils {

namespace {
uint64_t deserialize_int64(char*& bufptr) {
    uint64_t ret_value;
    memcpy(&ret_value, bufptr, UINT64_SIZE);
    return ret_value;
}

void marshal_uint64_t(const uint64_t& val, std::vector<unsigned char>& buf) {
    auto big_endian_val = boost::endian::native_to_big(val);
    std::array<unsigned char, UINT64_SIZE> tmpbuf;
    memcpy(tmpbuf.data(), &big_endian_val, sizeof(big_endian_val));

    buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
}

struct ValueSerializer {
    std::vector<unsigned char> operator()(const Tuple& val) const {
        std::vector<unsigned char> value_vector;
        auto type_code = (unsigned char)TUPLE_TYPE;
        value_vector.push_back(type_code);

        auto hash_key = hash(val);
        std::vector<unsigned char> hash_key_vector;
        marshal_uint256_t(hash_key, hash_key_vector);

        value_vector.insert(value_vector.end(), hash_key_vector.begin(),
                            hash_key_vector.end());

        return value_vector;
    }

    std::vector<unsigned char> operator()(const uint256_t& val) const {
        std::vector<unsigned char> value_vector;
        auto type_code = (unsigned char)NUM_TYPE;
        value_vector.push_back(type_code);

        std::vector<unsigned char> num_vector;
        marshal_uint256_t(val, num_vector);

        value_vector.insert(value_vector.end(), num_vector.begin(),
                            num_vector.end());

        return value_vector;
    }

    std::vector<unsigned char> operator()(const CodePoint& val) const {
        std::vector<unsigned char> value_vector;
        auto type_code = (unsigned char)CODEPT_TYPE;
        value_vector.push_back(type_code);

        std::vector<unsigned char> pc_vector;
        marshal_uint64_t(val.pc, pc_vector);

        value_vector.insert(value_vector.end(), pc_vector.begin(),
                            pc_vector.end());
        return value_vector;
    }
};

using iterator = std::vector<unsigned char>::iterator;

unsigned char extractStatus(iterator& iter) {
    auto status = (unsigned char)(*iter);
    iter += 1;

    return status;
}

std::vector<unsigned char> extractBlockReason(iterator& iter) {
    auto block_type = (BlockType)*iter;
    auto length_of_block_reason = blockreason_type_length[block_type];

    auto end_iter = iter + length_of_block_reason;
    std::vector<unsigned char> blockreason_vector(iter, end_iter);
    iter = end_iter;

    return blockreason_vector;
}

std::vector<unsigned char> extractBalanceTracker(iterator& iter) {
    unsigned int tracker_length;
    memcpy(&tracker_length, &(*iter), sizeof(tracker_length));
    iter += sizeof(tracker_length);

    auto end_iter = iter + tracker_length;
    std::vector<unsigned char> balance_track_vector(iter, end_iter);
    iter = end_iter;

    return balance_track_vector;
}

std::vector<unsigned char> extractHashKey(iterator& iter) {
    auto end_iter = iter + HASH_KEY_LENGTH;
    std::vector<unsigned char> hash_key(iter, end_iter);
    iter = end_iter;

    return hash_key;
}
}  // namespace

ParsedCheckpointState parseCheckpointState(
    std::vector<unsigned char> stored_state) {
    auto current_iter = stored_state.begin();

    auto status = extractStatus(current_iter);
    auto blockreason_vector = extractBlockReason(current_iter);
    auto balance_track_vector = extractBalanceTracker(current_iter);

    auto static_val = extractHashKey(current_iter);
    auto register_val = extractHashKey(current_iter);
    auto datastack = extractHashKey(current_iter);
    auto auxstack = extractHashKey(current_iter);
    auto inbox = extractHashKey(current_iter);
    auto inbox_count = extractHashKey(current_iter);
    auto pending = extractHashKey(current_iter);
    auto pending_count = extractHashKey(current_iter);
    auto pc = extractHashKey(current_iter);

    return ParsedCheckpointState{static_val,
                                 register_val,
                                 datastack,
                                 auxstack,
                                 inbox,
                                 inbox_count,
                                 pending,
                                 pending_count,
                                 pc,
                                 status,
                                 blockreason_vector,
                                 balance_track_vector};
}

std::vector<std::vector<unsigned char>> parseSerializedTuple(
    std::vector<unsigned char> data_vector) {
    std::vector<std::vector<unsigned char>> return_vector;

    auto it = data_vector.begin() + 1;

    while (it < data_vector.end()) {
        auto value_type = (valueTypes)*it;
        std::vector<unsigned char> current;

        switch (value_type) {
            case TUPLE_TYPE: {
                auto next_it = it + TUP_TUPLE_LENGTH;
                current.insert(current.end(), it, next_it);
                it = next_it;
                break;
            }
            case NUM_TYPE: {
                auto next_it = it + TUP_NUM_LENGTH;
                current.insert(current.end(), it, next_it);
                it = next_it;
                break;
            }
            case CODEPT_TYPE: {
                auto next_it = it + TUP_CODEPT_LENGTH;
                current.insert(current.end(), it, next_it);
                it = next_it;
                break;
            }
        }
        return_vector.push_back(current);
    }
    return return_vector;
}

CodePoint deserializeCodepoint(std::vector<unsigned char> val) {
    auto buff = reinterpret_cast<char*>(&val[0]);
    auto pc_val = deserialize_int64(buff);
    return CodePoint(pc_val, Operation(), 0);
}

uint256_t deserializeUint256(std::vector<unsigned char> val) {
    auto buff = reinterpret_cast<char*>(&val[1]);
    return deserialize_int256(buff);
}

std::vector<unsigned char> serializeValue(const value& val) {
    return nonstd::visit(ValueSerializer{}, val);
}
}  // namespace StateSaverUtils
