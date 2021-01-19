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
#include <iostream>

const std::string INT_VAL_LABEL = "Int";
const std::string TUP_VAL_LABEL = "Tuple";
const std::string CP_VAL_LABEL = "CodePoint";
const std::string BUF_LABEL = "Buffer";
const std::string BUF_ELEM_LABEL = "elem";
const std::string BUF_LEAF_LABEL = "Leaf";
const std::string BUF_NODE_LABEL = "Node";
const std::string CP_INTERNAL_LABEL = "Internal";
const std::string OPCODE_LABEL = "opcode";
const std::string OPCODE_SUB_LABEL = "AVMOpcode";
const std::string IMMEDIATE_LABEL = "immediate";
const std::string CODE_LABEL = "code";
const std::string STATIC_LABEL = "static_val";

namespace {

uint256_t int_value_from_json(const nlohmann::json& value_json) {
    return intx::from_string<uint256_t>(
        "0x" + value_json[INT_VAL_LABEL].get<std::string>());
}

RawBuffer buffer_value_from_json(const nlohmann::json& buffer_json) {
    if (!buffer_json.contains(BUF_ELEM_LABEL)) {
        throw std::runtime_error("buffer must contain elem");
    }
    auto elem_json = buffer_json[BUF_ELEM_LABEL];
    if (elem_json.contains(BUF_LEAF_LABEL)) {
        auto& leaf_data = elem_json[BUF_LEAF_LABEL];
        if (!leaf_data.is_array()) {
            throw std::runtime_error("leaf data must be array");
        }
        auto data = std::make_shared<std::vector<uint8_t>>();
        for (auto& item : leaf_data) {
            data->push_back(item.get<uint8_t>());
        }
        if (data->size() > LEAF_SIZE) {
            auto res = RawBuffer();
            for (uint64_t i = 0; i < data->size(); i++) {
                res = res.set(i, (*data)[i]);
            }
            return res;
        }
        data->resize(LEAF_SIZE, 0);
        return {std::move(data)};
    } else if (elem_json.contains(BUF_NODE_LABEL)) {
        auto& node_data = elem_json[BUF_NODE_LABEL];
        if (!node_data.is_array()) {
            throw std::runtime_error("node data must be array");
        }
        auto nested_elems = node_data[0];
        if (!nested_elems.is_array()) {
            throw std::runtime_error("node data must be array");
        }
        auto level = node_data[1].get<int>();
        auto data = std::make_shared<std::vector<RawBuffer>>();
        for (auto& item : nested_elems) {
            data->push_back(buffer_value_from_json(item));
        }
        return {std::move(data), level};
    } else {
        throw std::runtime_error("unhandled buffer member type");
    }
}

value value_from_json(nlohmann::json full_value_json,
                      size_t op_count,
                      const CodeSegment& code) {
    std::vector<DeserializedValue> values;
    std::vector<std::reference_wrapper<const nlohmann::json>> json_values{
        full_value_json};
    while (!json_values.empty()) {
        auto value_json = std::move(json_values.back());
        json_values.pop_back();

        if (value_json.get().contains(INT_VAL_LABEL)) {
            values.push_back(value{int_value_from_json(value_json)});
        } else if (value_json.get().contains(TUP_VAL_LABEL)) {
            auto& json_tup = value_json.get()[TUP_VAL_LABEL];
            if (!json_tup.is_array() || json_tup.size() > 8) {
                throw std::runtime_error(
                    "tuple must contain array of size less than 9");
            }
            values.push_back(
                TuplePlaceholder{static_cast<uint8_t>(json_tup.size())});
            for (auto it = json_tup.rbegin(); it != json_tup.rend(); ++it) {
                json_values.push_back(*it);
            }
        } else if (value_json.get().contains(CP_VAL_LABEL)) {
            auto& cp_json = value_json.get()[CP_VAL_LABEL];
            auto internal_offset =
                cp_json.at(CP_INTERNAL_LABEL).get<uint64_t>();
            uint64_t pc = 0;
            // Special handle python compiler's marker for error code point
            if (internal_offset != std::numeric_limits<uint64_t>::max()) {
                pc = op_count - internal_offset;
            }
            values.push_back(
                value{CodePointStub({code.segmentID(), pc}, code[pc])});
        } else if (value_json.get().contains(BUF_LABEL)) {
            values.emplace_back(
                Buffer{buffer_value_from_json(value_json.get()[BUF_LABEL])});
        } else {
            throw std::runtime_error("invalid value type");
        }
    }
    return assembleValueFromDeserialized(std::move(values));
}

Operation operation_from_json(const nlohmann::json& op_json,
                              size_t op_count,
                              const CodeSegment& code) {
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
    return {opcode, value_from_json(std::move(imm), op_count, code)};
}
}  // namespace

value simple_value_from_json(const nlohmann::json& full_value_json) {
    std::vector<DeserializedValue> values;
    std::vector<std::reference_wrapper<const nlohmann::json>> json_values{
        full_value_json};
    while (!json_values.empty()) {
        auto value_json = std::move(json_values.back());
        json_values.pop_back();

        if (value_json.get().contains(INT_VAL_LABEL)) {
            values.push_back(value{int_value_from_json(value_json)});
        } else if (value_json.get().contains(TUP_VAL_LABEL)) {
            const auto& json_tup = value_json.get()[TUP_VAL_LABEL];
            if (!json_tup.is_array() || json_tup.size() > 8) {
                throw std::runtime_error(
                    "tuple must contain array of size less than 9");
            }
            values.push_back(
                TuplePlaceholder{static_cast<uint8_t>(json_tup.size())});
            for (auto it = json_tup.rbegin(); it != json_tup.rend(); ++it) {
                json_values.push_back(*it);
            }
        } else if (value_json.get().contains(BUF_LABEL)) {
            values.emplace_back(
                Buffer{buffer_value_from_json(value_json.get()[BUF_LABEL])});
        } else {
            throw std::runtime_error("invalid value type");
        }
    }
    return assembleValueFromDeserialized(std::move(values));
}

LoadedExecutable loadExecutable(const std::string& executable_filename) {
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
        segment->addOperation(operation_from_json(*it, op_count, *segment));
    }
    value static_val = value_from_json(
        std::move(executable_json.at(STATIC_LABEL)), op_count, *segment);
    return {std::move(segment), std::move(static_val)};
}
