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

#include <avm/checkpoint/checkpointutils.hpp>
#include <avm/machinestate/blockreason.hpp>
#include <avm/machinestate/tokenTracker.hpp>
#include <avm/value/codepoint.hpp>
#include <avm/value/tuple.hpp>
#include <bigint_utils.hpp>
#include <util.hpp>

constexpr int UINT64_SIZE = 8;
constexpr int HASH_KEY_LENGTH = 33;
constexpr int TUP_TUPLE_LENGTH = 34;
constexpr int TUP_NUM_LENGTH = 34;
constexpr int TUP_CODEPT_LENGTH = 9;

namespace checkpoint {

uint64_t deserialize_int64(const char*& bufptr) {
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

struct ValueSerializer {
    std::vector<unsigned char> operator()(const Tuple& val) const {
        std::vector<unsigned char> value_vector;
        auto type_code = static_cast<unsigned char>(TUPLE);
        value_vector.push_back(type_code);

        auto hash_key = hash(val);
        marshal_uint256_t(hash_key, value_vector);

        return value_vector;
    }

    std::vector<unsigned char> operator()(const uint256_t& val) const {
        std::vector<unsigned char> value_vector;
        auto type_code = static_cast<unsigned char>(NUM);
        value_vector.push_back(type_code);

        marshal_uint256_t(val, value_vector);

        return value_vector;
    }

    std::vector<unsigned char> operator()(const CodePoint& val) const {
        std::vector<unsigned char> value_vector;
        auto type_code = static_cast<unsigned char>(CODEPT);
        value_vector.push_back(type_code);

        std::vector<unsigned char> pc_vector;
        marshal_uint64_t(val.pc, pc_vector);

        value_vector.insert(value_vector.end(), pc_vector.begin(),
                            pc_vector.end());
        return value_vector;
    }
};

using iterator = std::vector<unsigned char>::const_iterator;

unsigned char extractStatus(iterator& iter) {
    auto status = (unsigned char)(*iter);
    ++iter;

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

namespace utils {

std::vector<std::vector<unsigned char>> parseTuple(
    const std::vector<unsigned char>& data) {
    std::vector<std::vector<unsigned char>> return_vector;

    auto iter = data.begin() + 1;

    while (iter < data.end()) {
        auto value_type = static_cast<ValueTypes>(*iter);
        std::vector<unsigned char> current;

        switch (value_type) {
            case TUPLE: {
                auto next_it = iter + TUP_TUPLE_LENGTH;
                current.insert(current.end(), iter, next_it);
                iter = next_it;
                break;
            }
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
            case HASH_ONLY: {
                throw std::runtime_error("HASH_ONLY item");
            }
        }
        return_vector.push_back(current);
    }
    return return_vector;
}

CodePoint deserializeCodepoint(const std::vector<unsigned char>& val,
                               const std::vector<CodePoint>& code) {
    auto buff = reinterpret_cast<const char*>(&val[1]);
    auto pc_val = deserialize_int64(buff);
    if (pc_val == pc_default) {
        return getErrCodePoint();
    } else {
        return code[pc_val];
    }
}

uint256_t deserializeUint256_t(const std::vector<unsigned char>& val) {
    auto buff = reinterpret_cast<const char*>(&val[2]);
    return deserializeUint256t(buff);
}

std::vector<unsigned char> serializeValue(const value& val) {
    return nonstd::visit(ValueSerializer{}, val);
}

ParsedState parseState(const std::vector<unsigned char>& stored_state) {
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
    auto err_pc = extractHashKey(current_iter);

    return ParsedState{static_val,
                       register_val,
                       datastack,
                       auxstack,
                       inbox,
                       inbox_count,
                       pending,
                       pending_count,
                       pc,
                       err_pc,
                       status,
                       blockreason_vector,
                       balance_track_vector};
}

std::vector<unsigned char> serializeState(const ParsedState& state_data) {
    std::vector<unsigned char> state_data_vector;
    state_data_vector.push_back(state_data.status_char);

    state_data_vector.insert(state_data_vector.end(),
                             state_data.blockreason_str.begin(),
                             state_data.blockreason_str.end());

    unsigned int tracker_length = state_data.balancetracker_str.size();
    std::vector<unsigned char> tracker_len_vector(sizeof(tracker_length));
    memcpy(&tracker_len_vector[0], &tracker_length, sizeof(tracker_length));

    state_data_vector.insert(state_data_vector.end(),
                             tracker_len_vector.begin(),
                             tracker_len_vector.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.balancetracker_str.begin(),
                             state_data.balancetracker_str.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.static_val_key.begin(),
                             state_data.static_val_key.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.register_val_key.begin(),
                             state_data.register_val_key.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.datastack_key.begin(),
                             state_data.datastack_key.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.auxstack_key.begin(),
                             state_data.auxstack_key.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.inbox_key.begin(),
                             state_data.inbox_key.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.inbox_count_key.begin(),
                             state_data.inbox_count_key.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.pending_key.begin(),
                             state_data.pending_key.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.pending_count_key.begin(),
                             state_data.pending_count_key.end());

    state_data_vector.insert(state_data_vector.end(), state_data.pc_key.begin(),
                             state_data.pc_key.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.err_pc_key.begin(),
                             state_data.err_pc_key.end());
    return state_data_vector;
}
}  // namespace utils
namespace storage {

std::tuple<uint32_t, std::vector<unsigned char>> parseCountAndValue(
    const std::string& string_value) {
    if (string_value.empty()) {
        return std::make_tuple(0, std::vector<unsigned char>());
    } else {
        const char* c_string = string_value.c_str();
        uint32_t ref_count;
        memcpy(&ref_count, c_string, sizeof(ref_count));
        std::vector<unsigned char> saved_value(
            string_value.begin() + sizeof(ref_count), string_value.end());

        return std::make_tuple(ref_count, saved_value);
    }
}

std::vector<unsigned char> serializeCountAndValue(
    uint32_t count,
    const std::vector<unsigned char>& value) {
    std::vector<unsigned char> output_vector(sizeof(count));
    memcpy(&output_vector[0], &count, sizeof(count));
    output_vector.insert(output_vector.end(), value.begin(), value.end());

    return output_vector;
}
}  // namespace storage
}  // namespace checkpoint
