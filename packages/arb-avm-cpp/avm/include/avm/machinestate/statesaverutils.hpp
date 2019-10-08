//
//  statesaverutils.hpp
//  avm
//
//  Created by Minh Truong on 10/7/19.
//

#ifndef statesaverutils_hpp
#define statesaverutils_hpp

#include "avm/machinestate/value/value.hpp"

struct SerializedValue {
    valueTypes type;
    std::string string_value;
};

struct ParsedCheckpointState {
    std::vector<unsigned char> static_val_key;
    std::vector<unsigned char> register_val_key;
    std::vector<unsigned char> datastack_key;
    std::vector<unsigned char> auxstack_key;
    std::vector<unsigned char> inbox_key;
    std::vector<unsigned char> pending_key;
    std::vector<unsigned char> pc_key;
    unsigned char status_char;
    std::vector<unsigned char> blockreason_str;
    std::vector<unsigned char> balancetracker_str;
};

namespace StateSaverUtils {
SerializedValue SerializeValue(const value& val);
CodePoint deserializeCheckpointCodePt(std::vector<unsigned char> val);
uint256_t deserializeCheckpoint256(std::vector<unsigned char> val);
std::vector<std::vector<unsigned char>> parseSerializedTuple(
    std::vector<unsigned char> data_vector);
ParsedCheckpointState parseCheckpointState(
    std::vector<unsigned char> stored_state);
}  // namespace StateSaverUtils

#endif /* statesaverutils_hpp */
