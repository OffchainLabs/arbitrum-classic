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

#include <boost/endian/conversion.hpp>

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
    char* cData = static_cast<char*>(malloc(data.size()));
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

inline ByteSlice* returnCharVectorVectorRaw(
    const std::vector<std::vector<unsigned char>>& data_vec) {
    auto cData =
        static_cast<ByteSlice*>(malloc(data_vec.size() * sizeof(ByteSlice)));
    for (size_t i = 0; i < data_vec.size(); i++) {
        cData[i] = returnCharVector(data_vec[i]);
    }
    return cData;
}

inline ByteSliceArray returnCharVectorVector(
    const std::vector<std::vector<unsigned char>>& data) {
    return {returnCharVectorVectorRaw(data), static_cast<int>(data.size())};
}

inline uint256_t receiveUint256(const void* data) {
    auto data_ptr = reinterpret_cast<const char*>(data);
    return deserializeUint256t(data_ptr);
}

inline std::vector<unsigned char> receiveByteSlice(const ByteSlice& bs) {
    std::vector<unsigned char> data;
    data.resize(bs.length);
    auto ptr = reinterpret_cast<unsigned char*>(bs.data);
    std::copy(ptr, ptr + bs.length, data.data());
    return data;
}

inline std::vector<std::vector<unsigned char>> receiveByteSliceArray(
    const ByteSliceArray& bsa) {
    std::vector<std::vector<unsigned char>> data;
    auto slices = reinterpret_cast<ByteSlice*>(bsa.slices);
    for (int i = 0; i < bsa.count; ++i) {
        data.emplace_back(receiveByteSlice(slices[i]));
    }
    return data;
}

inline void* returnUint256(const uint256_t& val) {
    std::vector<unsigned char> serializedVal;
    marshal_uint256_t(val, serializedVal);
    return returnCharVectorRaw(serializedVal);
}

inline Uint256Result returnUint256Result(const ValueResult<uint256_t>& val) {
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

inline RawAssertion makeRawAssertion(Assertion& assertion,
                                     uint256_t before_send_acc,
                                     uint256_t before_log_acc) {
    std::vector<unsigned char> sendData;
    uint256_t sendAcc = before_send_acc;
    for (const auto& send : assertion.sends) {
        auto big_size =
            boost::endian::native_to_big(static_cast<uint64_t>(send.size()));
        auto big_size_ptr = reinterpret_cast<const char*>(&big_size);
        sendData.insert(sendData.end(), big_size_ptr,
                        big_size_ptr + sizeof(big_size));
        sendData.insert(sendData.end(), send.begin(), send.end());
        sendAcc = hash(sendAcc, hash(send));
    }

    std::vector<unsigned char> logData;
    uint256_t logAcc = before_log_acc;
    for (const auto& log : assertion.logs) {
        marshal_value(log, logData);
        logAcc = hash(logAcc, hash_value(log));
    }

    std::vector<unsigned char> debugPrintData;
    for (const auto& debugPrint : assertion.debugPrints) {
        marshal_value(debugPrint, debugPrintData);
    }

    // TODO extend usage of uint256
    return {intx::narrow_cast<uint64_t>(assertion.inbox_messages_consumed),
            returnCharVector(sendData),
            static_cast<int>(assertion.sends.size()),
            returnUint256(sendAcc),
            returnCharVector(logData),
            static_cast<int>(assertion.logs.size()),
            returnUint256(logAcc),
            returnCharVector(debugPrintData),
            static_cast<int>(assertion.debugPrints.size()),
            intx::narrow_cast<uint64_t>(assertion.stepCount),
            intx::narrow_cast<uint64_t>(assertion.gasCount)};
}

inline RawAssertion makeEmptyAssertion() {
    return {0,
            returnCharVector(std::vector<char>{}),
            0,
            0,
            returnCharVector(std::vector<char>{}),
            0,
            0,
            returnCharVector(std::vector<char>{}),
            0,
            0,
            0};
}

inline Tuple getTuple(void* data) {
    auto charData = reinterpret_cast<const char*>(data);
    return std::get<Tuple>(deserialize_value(charData));
}

inline std::vector<std::vector<unsigned char>> getInboxMessages(void* data) {
    if (data == nullptr) {
        return {};
    }
    auto charData = static_cast<ByteSliceArray*>(data);
    auto slices = static_cast<ByteSlice*>(charData->slices);
    std::vector<std::vector<unsigned char>> messages;
    messages.reserve(charData->count);
    for (int i = 0; i < charData->count; ++i) {
        auto data_ptr = static_cast<unsigned char*>(slices[i].data);
        messages.emplace_back(data_ptr, data_ptr + slices[i].length);
    }
    return messages;
}

#endif /* cavm_utils_hpp */
