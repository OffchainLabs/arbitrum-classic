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

#include <iterator>

#include "utils.hpp"

#include <avm_values/tuple.hpp>

#include <boost/endian/conversion.hpp>

constexpr int UINT64_SIZE = 8;

std::unordered_map<int, int> blockreason_type_length = {{0, 1},
                                                        {1, 1},
                                                        {2, 1},
                                                        {3, 1},
                                                        {4, 34}};

namespace checkpoint {

namespace utils {

uint64_t deserialize_uint64(const char*& bufptr) {
    uint64_t ret_value;
    memcpy(&ret_value, bufptr, UINT64_SIZE);
    auto val = boost::endian::big_to_native(ret_value);
    bufptr += sizeof(uint64_t);
    return val;
}

}  // namespace utils
}  // namespace checkpoint
