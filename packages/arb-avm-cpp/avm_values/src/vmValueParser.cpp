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

const char* INT_VAL_LABEL = "Int";
const char* TUP_VAL_LABEL = "Tuple";
const char* CP_VAL_LABEL = "CodePoint";
const char* CP_INTERNAL_LABEL = "Internal";
const char* OPCODE_LABEL = "opcode";
const char* IMMEDIATE_LABEL = "immediate";
const char* CODE_LABEL = "code";
const char* STATIC_LABEL = "static_val";

namespace {

value value_from_json(const nlohmann::json& value_json,
                      size_t op_count,
                      const Code& code,
                      TuplePool& pool) {
    if (value_json.contains(INT_VAL_LABEL)) {
        return uint256_t{value_json[INT_VAL_LABEL].get<std::string>()};
    } else if (value_json.contains(TUP_VAL_LABEL)) {
        auto& json_tup = value_json[TUP_VAL_LABEL];
        if (!json_tup.is_array()) {
            throw std::runtime_error("tuple must contain array");
        }
        std::vector<value> values;
        for (auto& json_val : json_tup) {
            values.push_back(value_from_json(json_val, op_count, code, pool));
        }
        return Tuple(std::move(values), &pool);
    } else if (value_json.contains(CP_VAL_LABEL)) {
        auto& cp_json = value_json[CP_VAL_LABEL];
        auto internal_offset = cp_json.at(CP_INTERNAL_LABEL).get<uint64_t>();
        auto ref = [&]() -> CodePointRef {
            if (internal_offset == std::numeric_limits<uint64_t>::max()) {
                return {0, 0, true};
            } else {
                return {0, op_count - internal_offset, false};
            }
        }();
        return CodePointStub(ref, code.at(ref));
    } else {
        throw std::runtime_error("invalid value type");
    }
}

Operation operation_from_json(const nlohmann::json& op_json,
                              size_t op_count,
                              const Code& code,
                              TuplePool& pool) {
    auto opcode = op_json.at(OPCODE_LABEL).get<OpCode>();
    auto& imm = op_json.at(IMMEDIATE_LABEL);
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
        std::ifstream contract_input_stream(contract_filename);
        if (!contract_input_stream.is_open()) {
            throw std::runtime_error("doesn't exist");
        }
        nlohmann::json contract_json;
        contract_input_stream >> contract_json;
        auto& json_code = contract_json.at(CODE_LABEL);
        if (!json_code.is_array()) {
            throw std::runtime_error("expected code to be array");
        }
        auto op_count = json_code.size();
        Code code;
        CodePointStub prev = code.addSegment();
        for (auto it = json_code.rbegin(); it != json_code.rend(); ++it) {
            prev = code.addOperation(
                prev.pc, operation_from_json(*it, op_count, code, pool));
        }
        value static_val = value_from_json(contract_json.at(STATIC_LABEL),
                                           op_count, code, pool);
        return std::make_pair(
            StaticVmValues{std::move(code), std::move(static_val)}, true);
    } catch (std::exception& e) {
        std::cerr << "Failed to load code file " << contract_filename << ": "
                  << e.what() << "\n";
        return std::make_pair(StaticVmValues{}, false);
    }
}
