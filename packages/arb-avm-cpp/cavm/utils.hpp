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

#ifndef cavm_utils_h
#define cavm_utils_h

#include "ctypes.h"

#include <avm_values/codepoint.hpp>
#include <avm_values/tuple.hpp>
#include <avm_values/value.hpp>
#include <data_storage/storageresult.hpp>

#include <vector>

#include <cstdlib>

inline unsigned char* returnCharVectorRaw(
    const std::vector<unsigned char>& data) {
    unsigned char* cData =
        reinterpret_cast<unsigned char*>(malloc(data.size()));
    std::copy(data.begin(), data.end(), cData);
    return cData;
}

inline ByteSlice returnCharVector(const std::vector<unsigned char>& data) {
    return {returnCharVectorRaw(data), static_cast<int>(data.size())};
}

inline ByteSliceResult returnDataResult(const DataResults& results) {
    if (!results.status.ok()) {
        return {{}, false};
    }
    return {returnCharVector(results.data), true};
}

inline uint256_t receiveUint256(const void* data) {
    auto data_ptr = reinterpret_cast<const char*>(data);
    return deserializeUint256t(data_ptr);
}

inline void* returnUint256(const uint256_t& val) {
    std::vector<unsigned char> serializedVal;
    marshal_uint256_t(val, serializedVal);
    return returnCharVectorRaw(serializedVal);
}

inline HashResult returnUint256Result(const ValueResult<uint256_t>& val) {
    if (!val.status.ok()) {
        return {{}, false};
    }
    return {returnUint256(val.data), true};
}

inline Uint64Result returnUint64Result(const ValueResult<uint64_t>& val) {
    if (!val.status.ok()) {
        return {{}, false};
    }
    return {val.data, true};
}

inline ByteSlice returnValueResult(const DbResult<value>& res,
                                   const Code& code) {
    if (!res.status.ok()) {
        return {nullptr, 0};
    }

    std::vector<unsigned char> value;
    marshal_value(res.data, value, code);

    auto value_data = reinterpret_cast<unsigned char*>(malloc(value.size()));
    std::copy(value.begin(), value.end(), value_data);

    auto void_data = reinterpret_cast<void*>(value_data);
    return {void_data, static_cast<int>(value.size())};
}

#endif /* cavm_utils_h */
