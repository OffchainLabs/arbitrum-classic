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

#include "cconfirmednodestore.h"
#include "utils.hpp"

#include <data_storage/confirmednodestore.hpp>
#include <data_storage/storageresult.hpp>

void deleteConfirmedNodeStore(CConfirmedNodeStore* m) {
    delete static_cast<ConfirmedNodeStore*>(m);
}

int putNode(CConfirmedNodeStore* ptr,
            uint64_t height,
            const void* hash_ptr,
            const void* data,
            int data_length) {
    auto hash = receiveUint256(hash_ptr);
    auto data_ptr = reinterpret_cast<const char*>(data);
    auto data_vector = std::vector<char>(data_ptr, data_ptr + data_length);
    return static_cast<ConfirmedNodeStore*>(ptr)
        ->putNode(height, hash, data_vector)
        .ok();
}

ByteSliceResult getNode(CConfirmedNodeStore* ptr,
                        uint64_t height,
                        const void* hash_ptr) {
    auto hash = receiveUint256(hash_ptr);
    return returnDataResult(
        static_cast<ConfirmedNodeStore*>(ptr)->getNode(height, hash));
}

Uint64Result getNodeHeight(CConfirmedNodeStore* ptr, const void* hash_ptr) {
    auto hash = receiveUint256(hash_ptr);
    return returnUint64Result(
        static_cast<ConfirmedNodeStore*>(ptr)->getHeight(hash));
}

HashResult getNodeHash(CConfirmedNodeStore* ptr, uint64_t height) {
    return returnUint256Result(
        static_cast<ConfirmedNodeStore*>(ptr)->getHash(height));
}

int isNodeStoreEmpty(CConfirmedNodeStore* ptr) {
    return static_cast<ConfirmedNodeStore*>(ptr)->isEmpty();
}

uint64_t maxNodeHeight(CConfirmedNodeStore* ptr) {
    return static_cast<ConfirmedNodeStore*>(ptr)->maxNodeHeight();
}
