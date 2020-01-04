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

const char* getContractData(const std::string& contract_filename) {
    std::ifstream myfile;

    struct stat filestatus;
    stat(contract_filename.c_str(), &filestatus);

    char* buf = (char*)malloc(filestatus.st_size);

    myfile.open(contract_filename, std::ios::in);

    if (myfile.is_open()) {
        myfile.read((char*)buf, filestatus.st_size);
        myfile.close();
    }

    return buf;
}

InitialVmValues getInitialVmValues(const std::string& contract_filename,
                                   TuplePool* pool) {
    InitialVmValues initial_state;

    auto bufptr = getContractData(contract_filename);

    uint32_t version;
    memcpy(&version, bufptr, sizeof(version));
    version = __builtin_bswap32(version);
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
            extentionId = __builtin_bswap32(extentionId);
            bufptr += sizeof(extentionId);
            if (extentionId > 0) {
                //            std::cout << "found extention" << std::endl;
            }
        }
        uint64_t codeCount;
        memcpy(&codeCount, bufptr, sizeof(codeCount));
        bufptr += sizeof(codeCount);
        codeCount = boost::endian::big_to_native(codeCount);

        initial_state.code.reserve(codeCount);

        std::vector<Operation> ops;
        for (uint64_t i = 0; i < codeCount; i++) {
            ops.emplace_back(deserializeOperation(bufptr, *pool));
        }
        initial_state.code = opsToCodePoints(ops);
        initial_state.staticVal = deserialize_value(bufptr, *pool);
        initial_state.valid_state = true;

        return initial_state;
    }
}
