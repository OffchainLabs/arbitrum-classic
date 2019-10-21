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

#ifndef statesaverutils_hpp
#define statesaverutils_hpp

#include "avm/value/value.hpp"

struct ParsedState {
    std::vector<unsigned char> static_val_key;
    std::vector<unsigned char> register_val_key;
    std::vector<unsigned char> datastack_key;
    std::vector<unsigned char> auxstack_key;
    std::vector<unsigned char> inbox_key;
    std::vector<unsigned char> inbox_count_key;
    std::vector<unsigned char> pending_key;
    std::vector<unsigned char> pending_count_key;
    std::vector<unsigned char> pc_key;
    unsigned char status_char;
    std::vector<unsigned char> blockreason_str;
    std::vector<unsigned char> balancetracker_str;
};

namespace Checkpoint {
namespace Utils {
std::vector<unsigned char> serializeValue(const value& val);
CodePoint deserializeCodepoint(std::vector<unsigned char>& val);
uint256_t deserializeUint256(std::vector<unsigned char>& val);
std::vector<std::vector<unsigned char>> parseSerializedTuple(
    std::vector<unsigned char> data_vector);
ParsedState parseState(std::vector<unsigned char> stored_state);
std::vector<unsigned char> serializeState(ParsedState state_data);
}  // namespace Utils
}  // namespace Checkpoint

#endif /* statesaverutils_hpp */
