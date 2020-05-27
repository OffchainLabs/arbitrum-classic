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

#include "cnodestore.h"
#include "utils.hpp"

#include <data_storage/nodestore.hpp>
#include <data_storage/storageresult.hpp>

void deleteNodeStore(CNodeStore* m) {
    delete static_cast<NodeStore*>(m);
}

int putNode(CNodeStore* ptr,
            uint64_t height,
            const void* hash_ptr,
            const void* data,
            int data_length) {
    auto hash = receiveUint256(hash_ptr);
    auto data_ptr = reinterpret_cast<const char*>(data);
    auto data_vector = std::vector<char>(data_ptr, data_ptr + data_length);
    return static_cast<NodeStore*>(ptr)
        ->putNode(height, hash, data_vector)
        .ok();
}

ByteSliceResult getNode(CNodeStore* ptr,
                        uint64_t height,
                        const void* hash_ptr) {
    auto hash = receiveUint256(hash_ptr);
    return returnDataResult(
        static_cast<NodeStore*>(ptr)->getNode(height, hash));
}

Uint64Result getNodeHeight(CNodeStore* ptr, const void* hash_ptr) {
    auto hash = receiveUint256(hash_ptr);
    return returnUint64Result(static_cast<NodeStore*>(ptr)->getHeight(hash));
}

HashResult getNodeHash(CNodeStore* ptr, uint64_t height) {
    return returnUint256Result(static_cast<NodeStore*>(ptr)->getHash(height));
}

uint64_t longestNodeChainCount(CNodeStore* ptr) {
    return static_cast<NodeStore*>(ptr)->longestChainCount();
}
