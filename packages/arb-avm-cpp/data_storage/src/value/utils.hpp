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

#ifndef checkpointutils_hpp
#define checkpointutils_hpp

#include <avm_values/codepointstub.hpp>
#include <avm_values/tuple.hpp>

extern std::unordered_map<int, int> blockreason_type_length;

namespace checkpoint {
namespace utils {
uint64_t deserialize_uint64(const char*& bufptr);
}  // namespace utils
}  // namespace checkpoint

#endif /* checkpointutils_hpp */
