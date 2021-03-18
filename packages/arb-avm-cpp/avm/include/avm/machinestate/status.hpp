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

#ifndef avm_status_hpp
#define avm_status_hpp

#include <avm/inboxmessage.hpp>
#include <variant>

typedef std::variant<std::monostate, InboxMessage, uint256_t> staged_variant;

enum class Status { Extensive, Halted, Error };

#endif /* avm_status_hpp */
