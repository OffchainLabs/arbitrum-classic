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

#ifndef checkpointutils_hpp
#define checkpointutils_hpp

#include <avm_values/codepoint.hpp>
#include <avm_values/tuple.hpp>

struct MachineStateKeys {
    std::vector<unsigned char> register_val_key;
    std::vector<unsigned char> datastack_key;
    std::vector<unsigned char> auxstack_key;
    std::vector<unsigned char> pc_key;
    std::vector<unsigned char> err_pc_key;
    unsigned char status_char;
};

extern std::unordered_map<int, int> blockreason_type_length;

namespace checkpoint {
namespace utils {
auto serializeValue(const value& val) -> std::vector<unsigned char>;
auto deserializeCodepoint(const std::vector<unsigned char>& val,
                          const std::vector<CodePoint>& code) -> CodePoint;
auto deserializeUint256_t(const std::vector<unsigned char>& val) -> uint256_t;
auto parseTuple(const std::vector<unsigned char>& data)
    -> std::vector<std::vector<unsigned char>>;
auto extractStateKeys(const std::vector<unsigned char>& stored_state)
    -> MachineStateKeys;
auto serializeStateKeys(const MachineStateKeys& state_data)
    -> std::vector<unsigned char>;
}  // namespace utils
}  // namespace checkpoint

#endif /* checkpointutils_hpp */
