//
//  statesaverutils.cpp
//  avm
//
//  Created by Minh Truong on 10/7/19.
//

#include "avm/machinestate/statesaverutils.hpp"
#include "avm/machinestate/tokenTracker.hpp"
#include "avm/machinestate/value/codepoint.hpp"
#include "avm/machinestate/value/tuple.hpp"
#include "bigint_utils.hpp"
#include "util.hpp"

#define UINT64_SIZE 8
#define HASH_KEY_LENGTH 33
#define TUP_TUPLE_LENGTH 34
#define TUP_NUM_LENGTH 34
#define TUP_CODEPT_LENGTH 9

namespace StateSaverUtils {

namespace {
// make sure correct
uint64_t deserialize_int64(char*& bufptr) {
    uint64_t ret_value;
    memcpy(&ret_value, bufptr, UINT64_SIZE);
    return ret_value;
}

void marshal_uint64_t(const uint64_t& val, std::vector<unsigned char>& buf) {
    auto big_endian_val = boost::endian::native_to_big(val);
    std::array<unsigned char, 8> tmpbuf;
    memcpy(tmpbuf.data(), &big_endian_val, sizeof(big_endian_val));

    buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
}

std::vector<unsigned char> serializeForCheckpoint(const Tuple& val) {
    std::vector<unsigned char> value_vector;
    auto type_code = (unsigned char)TUPLE;
    value_vector.push_back(type_code);

    auto hash_key = hash(val);
    std::vector<unsigned char> hash_key_vector;
    marshal_uint256_t(hash_key, hash_key_vector);

    value_vector.insert(value_vector.end(), hash_key_vector.begin(),
                        hash_key_vector.end());

    return value_vector;
}

std::vector<unsigned char> serializeForCheckpoint(const uint256_t& val) {
    std::vector<unsigned char> value_vector;
    auto type_code = (unsigned char)NUM;
    value_vector.push_back(type_code);

    std::vector<unsigned char> num_vector;
    marshal_uint256_t(val, num_vector);

    value_vector.insert(value_vector.end(), num_vector.begin(),
                        num_vector.end());

    return value_vector;
}

std::vector<unsigned char> serializeForCheckpoint(const CodePoint& val) {
    std::vector<unsigned char> value_vector;
    auto type_code = (unsigned char)CODEPT;
    value_vector.push_back(type_code);

    std::vector<unsigned char> pc_vector;
    marshal_uint64_t(val.pc, pc_vector);

    value_vector.insert(value_vector.end(), pc_vector.begin(), pc_vector.end());

    return value_vector;
}

struct Serializer {
    SerializedValue operator()(const Tuple& val) const {
        auto value_vector = serializeForCheckpoint(val);
        std::string str_value(value_vector.begin(), value_vector.end());
        SerializedValue serialized_value{TUPLE_TYPE, str_value};

        return serialized_value;
    }

    SerializedValue operator()(const uint256_t& val) const {
        auto value_vector = serializeForCheckpoint(val);
        std::string str_value(value_vector.begin(), value_vector.end());
        SerializedValue serialized_value{NUM_TYPE, str_value};

        return serialized_value;
    }

    SerializedValue operator()(const CodePoint& val) const {
        auto value_vector = serializeForCheckpoint(val);
        std::string str_value(value_vector.begin(), value_vector.end());
        SerializedValue serialized_value{CODEPT_TYPE, str_value};

        return serialized_value;
    }
};
}  // namespace

ParsedCheckpointState parseCheckpointState(
    std::vector<unsigned char> stored_state) {
    auto current_iter = stored_state.begin();
    auto status = (unsigned char)(*current_iter);
    current_iter += 1;

    // blockreason
    auto block_type = (BlockType)*current_iter;
    auto length_of_block_reason = blockreason_type_length[block_type];
    std::vector<unsigned char> blockreason_vector(
        current_iter, current_iter + length_of_block_reason);

    current_iter += length_of_block_reason;

    // balancetracker
    unsigned int balance_tracker_length;
    memcpy(&balance_tracker_length, &(*current_iter),
           sizeof(balance_tracker_length));
    current_iter += sizeof(unsigned int);

    auto token_pair_length = TOKEN_VAL_LENGTH + TOKEN_TYPE_LENGTH;
    auto total_len = token_pair_length * balance_tracker_length;
    std::vector<unsigned char> balance_track_vector(current_iter,
                                                    current_iter + total_len);
    current_iter += total_len;
    auto next_iter = current_iter + HASH_KEY_LENGTH;

    std::vector<unsigned char> static_val(current_iter, next_iter);
    current_iter = next_iter;
    next_iter += HASH_KEY_LENGTH;
    std::vector<unsigned char> register_val(current_iter, next_iter);
    current_iter = next_iter;
    next_iter += HASH_KEY_LENGTH;
    std::vector<unsigned char> datastack(current_iter, next_iter);
    current_iter = next_iter;
    next_iter += HASH_KEY_LENGTH;
    std::vector<unsigned char> auxstack(current_iter, next_iter);
    current_iter = next_iter;
    next_iter += HASH_KEY_LENGTH;
    std::vector<unsigned char> inbox(current_iter, next_iter);
    current_iter = next_iter;
    next_iter += HASH_KEY_LENGTH;
    std::vector<unsigned char> pending(current_iter, next_iter);
    current_iter = next_iter;
    next_iter += HASH_KEY_LENGTH;
    std::vector<unsigned char> pc(current_iter, next_iter);

    return ParsedCheckpointState{static_val,
                                 register_val,
                                 datastack,
                                 auxstack,
                                 inbox,
                                 pending,
                                 pc,
                                 status,
                                 blockreason_vector,
                                 balance_track_vector};
}

// make sure correct
std::vector<std::vector<unsigned char>> parseSerializedTuple(
    std::vector<unsigned char> data_vector) {
    std::vector<std::vector<unsigned char>> return_vector;

    auto it = data_vector.begin() + 1;

    while (it != data_vector.end()) {
        auto value_type = (valueTypes)*it;
        std::vector<unsigned char> current;

        switch (value_type) {
            case TUPLE_TYPE: {
                auto next_it = it + TUP_TUPLE_LENGTH;
                current.insert(current.end(), it, next_it);
                it = next_it;
            }
            case NUM_TYPE: {
                auto next_it = it + TUP_NUM_LENGTH;
                current.insert(current.end(), it, next_it);
                it = next_it;
            }
            case CODEPT_TYPE: {
                auto next_it = it + TUP_CODEPT_LENGTH;
                current.insert(current.end(), it, next_it);
                it = next_it;
            }
        }

        return_vector.push_back(current);
    }

    return return_vector;
}

CodePoint deserializeCheckpointCodePt(std::vector<unsigned char> val) {
    CodePoint code_point;
    auto buff = reinterpret_cast<char*>(&val[1]);
    auto pc_val = deserialize_int64(buff);
    code_point.pc = pc_val;

    return code_point;
}

uint256_t deserializeCheckpoint256(std::vector<unsigned char> val) {
    auto buff = reinterpret_cast<char*>(&val[1]);
    auto num = deserialize_int256(buff);

    return num;
}

SerializedValue SerializeValue(const value& val) {
    return nonstd::visit(Serializer{}, val);
}
}  // namespace StateSaverUtils
