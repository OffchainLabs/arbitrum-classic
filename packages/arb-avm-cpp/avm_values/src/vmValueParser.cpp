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

#include <bigint_utils.hpp>

#include <fstream>

std::vector<char> getContractData(const std::string& contract_filename) {
    std::ifstream contract_file(contract_filename,
                                std::ios::in | std::ios::binary);
    if (!contract_file.is_open()) {
        return {};
    }
    return {std::istreambuf_iterator<char>(contract_file),
            std::istreambuf_iterator<char>()};
}

std::pair<StaticVmValues, bool> parseStaticVmValues(
    const std::string& contract_filename,
    TuplePool& pool) {
    auto parseError = [&]() -> std::pair<StaticVmValues, bool> {
        std::cerr << "Failed to parse file " << contract_filename << std::endl;
        return std::make_pair(StaticVmValues{}, false);
    };

    auto contract_data = getContractData(contract_filename);

    if (contract_data.size() < 32) {
        std::cerr << "Failed to open path: " << contract_filename << std::endl;
        return std::make_pair(StaticVmValues{}, false);
    }

    size_t offset = 0;

    uint32_t version;
    if (offset + sizeof(version) > contract_data.size()) {
        return parseError();
    }
    auto it = contract_data.begin() + offset;
    std::copy(it, it + sizeof(version), reinterpret_cast<char*>(&version));
    offset += sizeof(version);
    version = boost::endian::big_to_native(version);

    if (version != CURRENT_AO_VERSION) {
        std::cerr << "incorrect version of .ao file" << std::endl;
        std::cerr << "expected version " << CURRENT_AO_VERSION
                  << " found version " << version << std::endl;
        return std::make_pair(StaticVmValues{}, false);
    }
    uint32_t extentionId = 1;
    while (extentionId != 0) {
        if (offset + sizeof(extentionId) > contract_data.size()) {
            return parseError();
        }
        auto it = contract_data.begin() + offset;
        std::copy(it, it + sizeof(extentionId),
                  reinterpret_cast<char*>(&extentionId));
        offset += sizeof(extentionId);
        extentionId = boost::endian::big_to_native(extentionId);
        if (extentionId > 0) {
            uint32_t extensionLength;
            if (offset + sizeof(extensionLength) > contract_data.size()) {
                return parseError();
            }
            auto it = contract_data.begin() + offset;
            std::copy(it, it + sizeof(extensionLength),
                      reinterpret_cast<char*>(&extensionLength));
            offset += sizeof(extensionLength);
            extensionLength = boost::endian::big_to_native(extensionLength);

            if (offset + extensionLength > contract_data.size()) {
                return parseError();
            }
            offset += extensionLength;
        }
    }

    uint64_t codeCount;
    if (offset + sizeof(codeCount) > contract_data.size()) {
        return parseError();
    }
    it = contract_data.begin() + offset;
    std::copy(it, it + sizeof(codeCount), reinterpret_cast<char*>(&codeCount));
    offset += sizeof(codeCount);
    codeCount = boost::endian::big_to_native(codeCount);

    // TODO: The following code may read beyond the code buffer leading to a
    // crash To fix this we would need to make all of the deserialization
    // functions do bounds checking This may not be too big of a security risk
    // since this would lead the validator to crash on point rather than at an
    // attacker controlled time, but we should definitely fix if possible

    const char* bufptr = contract_data.data() + offset;
    std::vector<Operation> ops;
    for (uint64_t i = 0; i < codeCount; i++) {
        ops.emplace_back(deserializeOperation(bufptr, pool));
    }
    auto staticVal = deserialize_value(bufptr, pool);

    return std::make_pair(StaticVmValues{Code{opsToCodePoints(ops)}, staticVal},
                          true);
}
