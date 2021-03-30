/*
 * Copyright 2020, Offchain Labs, Inc.
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

#ifndef vmValueParser_hpp
#define vmValueParser_hpp

#include <avm_values/code.hpp>
#include <avm_values/tuple.hpp>

#include <nlohmann/json.hpp>

struct LoadedExecutable {
    std::shared_ptr<CodeSegment> code;
    value static_val;

    LoadedExecutable(std::shared_ptr<CodeSegment> code_, value static_val_)
        : code(std::move(code_)), static_val(std::move(static_val_)) {}
};

value simple_value_from_json(const nlohmann::json& value_json);
std::vector<uint8_t> send_from_json(const nlohmann::json& val);

LoadedExecutable loadExecutable(const std::string& executable_filename);

#endif /* vmValueParser_hpp */
