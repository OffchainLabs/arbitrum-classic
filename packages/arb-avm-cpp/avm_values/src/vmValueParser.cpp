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
const char* OPCODE_SUB_LABEL = "AVMOpcode";
const char* IMMEDIATE_LABEL = "immediate";
const char* CODE_LABEL = "code";
const char* STATIC_LABEL = "static_val";

namespace {

uint256_t int_value_from_json(const nlohmann::json& value_json) {
    return uint256_t{"0x" + value_json[INT_VAL_LABEL].get<std::string>()};
}

value value_from_json(const nlohmann::json& value_json,
                      size_t op_count,
                      const CodeSegment& code,
                      TuplePool& pool) {
    if (value_json.contains(INT_VAL_LABEL)) {
        return int_value_from_json(value_json);
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
        uint64_t pc = 0;
        // Special handle python compiler's marker for error code point
        if (internal_offset != std::numeric_limits<uint64_t>::max()) {
            pc = op_count - internal_offset;
        }
        return CodePointStub({code.segmentID(), pc}, code[pc]);
    } else {
        throw std::runtime_error("invalid value type");
    }
}

Operation operation_from_json(const nlohmann::json& op_json,
                              size_t op_count,
                              const CodeSegment& code,
                              TuplePool& pool) {
    auto opcode_json = op_json.at(OPCODE_LABEL);
    if (opcode_json.contains(OPCODE_SUB_LABEL)) {
        opcode_json = opcode_json.at(OPCODE_SUB_LABEL);
    }
    if (!opcode_json.is_number_integer()) {
        std::cerr << "Invalid opcode " << opcode_json << "\n";
    }
    auto opcode = opcode_json.get<OpCode>();
    auto& imm = op_json.at(IMMEDIATE_LABEL);
    if (imm.is_null()) {
        return {opcode};
    }
    return {opcode, value_from_json(imm, op_count, code, pool)};
}
}  // namespace

value simple_value_from_json(const nlohmann::json& value_json,
                             TuplePool& pool) {
    if (value_json.contains(INT_VAL_LABEL)) {
        return int_value_from_json(value_json);
    } else if (value_json.contains(TUP_VAL_LABEL)) {
        auto& json_tup = value_json[TUP_VAL_LABEL];
        if (!json_tup.is_array()) {
            throw std::runtime_error("tuple must contain array");
        }
        std::vector<value> values;
        for (auto& json_val : json_tup) {
            values.push_back(simple_value_from_json(json_val, pool));
        }
        return Tuple(std::move(values), &pool);
    } else {
        throw std::runtime_error("invalid value type");
    }
}

LoadedExecutable loadExecutable(const std::string& executable_filename,
                                TuplePool& pool) {
    std::ifstream executable_input_stream(executable_filename);
    if (!executable_input_stream.is_open()) {
        throw std::runtime_error("doesn't exist");
    }
    nlohmann::json executable_json;
    executable_input_stream >> executable_json;
    auto& json_code = executable_json.at(CODE_LABEL);
    if (!json_code.is_array()) {
        throw std::runtime_error("expected code to be array");
    }
    auto op_count = json_code.size();
    auto segment = std::make_shared<CodeSegment>(0);
    for (auto it = json_code.rbegin(); it != json_code.rend(); ++it) {
        segment->addOperation(
            operation_from_json(*it, op_count, *segment, pool));
    }
    value static_val = value_from_json(executable_json.at(STATIC_LABEL),
                                       op_count, *segment, pool);
    return {std::move(segment), std::move(static_val)};
}
