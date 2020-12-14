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

#ifndef cavm_utils_hpp
#define cavm_utils_hpp

#include "ctypes.h"

#include <avm/machine.hpp>
#include <avm_values/codepointstub.hpp>
#include <avm_values/tuple.hpp>
#include <avm_values/value.hpp>
#include <data_storage/storageresult.hpp>

#include <vector>

#include <cstdlib>

inline unsigned char* returnCharVectorRaw(
    const std::vector<unsigned char>& data) {
    auto cData = reinterpret_cast<unsigned char*>(malloc(data.size()));
    std::copy(data.begin(), data.end(), cData);
    return cData;
}

inline ByteSlice returnCharVector(const std::vector<unsigned char>& data) {
    return {returnCharVectorRaw(data), static_cast<int>(data.size())};
}

inline char* returnCharVectorRaw(const std::vector<char>& data) {
    char* cData = reinterpret_cast<char*>(malloc(data.size()));
    std::copy(data.begin(), data.end(), cData);
    return cData;
}

inline ByteSlice returnCharVector(const std::vector<char>& data) {
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

inline ByteSlice returnValueResult(const DbResult<value>& res) {
    if (!res.status.ok()) {
        return {nullptr, 0};
    }

    std::vector<unsigned char> value;
    marshal_value(res.data, value);

    auto value_data = reinterpret_cast<unsigned char*>(malloc(value.size()));
    std::copy(value.begin(), value.end(), value_data);

    auto void_data = reinterpret_cast<void*>(value_data);
    return {void_data, static_cast<int>(value.size())};
}

inline RawAssertion makeRawAssertion(Assertion& assertion) {
    std::vector<unsigned char> outMsgData;
    for (const auto& outMsg : assertion.outMessages) {
        marshal_value(outMsg, outMsgData);
    }
    std::vector<unsigned char> logData;
    for (const auto& log : assertion.logs) {
        marshal_value(log, logData);
    }

    std::vector<unsigned char> debugPrintData;
    for (const auto& debugPrint : assertion.debugPrints) {
        marshal_value(debugPrint, debugPrintData);
    }

    return {assertion.inbox_messages_consumed,
            returnCharVector(outMsgData),
            static_cast<int>(assertion.outMessages.size()),
            returnCharVector(logData),
            static_cast<int>(assertion.logs.size()),
            returnCharVector(debugPrintData),
            static_cast<int>(assertion.debugPrints.size()),
            assertion.stepCount,
            assertion.gasCount};
}

inline RawAssertion makeEmptyAssertion() {
    return {0, returnCharVector(std::vector<char>{}),
            0, returnCharVector(std::vector<char>{}),
            0, returnCharVector(std::vector<char>{}),
            0, 0,
            0};
}

inline Tuple getTuple(void* data) {
    auto charData = reinterpret_cast<const char*>(data);
    return nonstd::get<Tuple>(deserialize_value(charData));
}

inline std::vector<Tuple> getInboxMessages(void* data, uint64_t message_count) {
    auto charData = reinterpret_cast<const char*>(data);
    std::vector<Tuple> messages;
    messages.reserve(message_count);
    for (uint64_t i = 0; i < message_count; ++i) {
        messages.push_back(deserialize_value(charData).get<Tuple>());
    }
    return messages;
}

#endif /* cavm_utils_hpp */
