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

#include <avm_values/vmValueParser.hpp>

#include <nlohmann/json.hpp>

#include <fstream>

namespace {

value value_from_json(const nlohmann::json& j,
                      size_t op_count,
                      const Code& code,
                      TuplePool& pool) {
    if (j.contains("Int")) {
        return uint256_t{j["Int"].get<std::string>()};
    } else if (j.contains("Tuple")) {
        auto& json_tup = j["Tuple"];
        if (!json_tup.is_array()) {
            throw std::runtime_error("tuple must contain array");
        }
        std::vector<value> values;
        for (auto& json_val : json_tup) {
            values.push_back(value_from_json(json_val, op_count, code, pool));
        }
        return Tuple(std::move(values), &pool);
    } else if (j.contains("CodePoint")) {
        auto& cp_json = j["CodePoint"];
        auto internal_offset = cp_json.at("Internal").get<uint64_t>();
        CodePointRef ref;
        if (internal_offset == std::numeric_limits<uint64_t>::max()) {
            ref = {0, true};
        } else {
            ref = {op_count - internal_offset, false};
        }
        return CodePointStub(ref, code.at(ref));
    } else {
        throw std::runtime_error("invalid value type");
    }
}

Operation operation_from_json(const nlohmann::json& j,
                              size_t op_count,
                              const Code& code,
                              TuplePool& pool) {
    auto opcode = j.at("opcode").get<OpCode>();
    auto& imm = j.at("immediate");
    if (imm.is_null()) {
        return {opcode};
    }
    return {opcode, value_from_json(imm, op_count, code, pool)};
}
}  // namespace

std::pair<StaticVmValues, bool> parseStaticVmValues(
    const std::string& contract_filename,
    TuplePool& pool) {
    try {
        std::ifstream i(contract_filename);
        if (!i.is_open()) {
            throw std::runtime_error("doesn't exist");
        }
        nlohmann::json j;
        i >> j;
        auto& json_code = j.at("code");
        if (!json_code.is_array()) {
            throw std::runtime_error("expected code to be array");
        }
        auto op_count = json_code.size();
        Code code;
        for (auto it = json_code.rbegin(); it != json_code.rend(); ++it) {
            code.addOperation(operation_from_json(*it, op_count, code, pool));
        }
        value static_val =
            value_from_json(j.at("static_val"), op_count, code, pool);
        return std::make_pair(
            StaticVmValues{std::move(code), std::move(static_val)}, true);
    } catch (std::exception& e) {
        std::cerr << "Failed to load code file " << contract_filename << ": "
                  << e.what() << "\n";
        return std::make_pair(StaticVmValues{}, false);
    }
}
