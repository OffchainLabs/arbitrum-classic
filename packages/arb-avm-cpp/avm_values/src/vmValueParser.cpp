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

#include <sys/stat.h>
#include <fstream>

auto getContractData(const std::string& contract_filename)
    -> std::vector<char> {
    std::ifstream myfile;

    struct stat filestatus;
    stat(contract_filename.c_str(), &filestatus);

    std::vector<char> data;
    data.resize(filestatus.st_size);

    myfile.open(contract_filename, std::ios::in);

    if (myfile.is_open()) {
        myfile.read(data.data(), filestatus.st_size);
        myfile.close();
    }

    return data;
}

auto parseInitialVmValues(const std::string& contract_filename, TuplePool& pool)
    -> InitialVmValues {
    InitialVmValues initial_state;

    auto data = getContractData(contract_filename);
    auto bufptr = const_cast<const char*>(data.data());

    uint32_t version;
    memcpy(&version, bufptr, sizeof(version));
    version = boost::endian::big_to_native(version);
    bufptr += sizeof(version);

    if (version != CURRENT_AO_VERSION) {
        std::cerr << "incorrect version of .ao file" << std::endl;
        std::cerr << "expected version " << CURRENT_AO_VERSION
                  << " found version " << version << std::endl;
        initial_state.valid_state = false;
        return initial_state;
    } else {
        uint32_t extentionId = 1;
        while (extentionId != 0) {
            memcpy(&extentionId, bufptr, sizeof(extentionId));
            extentionId = boost::endian::big_to_native(extentionId);
            bufptr += sizeof(extentionId);
            if (extentionId > 0) {
                uint32_t extensionLength;
                memcpy(&extensionLength, bufptr, sizeof(extensionLength));
                extensionLength = boost::endian::big_to_native(extensionLength);
                bufptr += sizeof(extensionLength) + extensionLength;
            }
        }
        uint64_t codeCount;
        memcpy(&codeCount, bufptr, sizeof(codeCount));
        bufptr += sizeof(codeCount);
        codeCount = boost::endian::big_to_native(codeCount);

        initial_state.code.reserve(codeCount);

        std::vector<Operation> ops;
        for (uint64_t i = 0; i < codeCount; i++) {
            ops.emplace_back(deserializeOperation(bufptr, pool));
        }
        initial_state.code = opsToCodePoints(std::move(ops));
        initial_state.staticVal = deserialize_value(bufptr, pool);
        initial_state.valid_state = true;

        return initial_state;
    }
}
