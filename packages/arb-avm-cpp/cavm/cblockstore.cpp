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

#include "cblockstore.h"

#include "utils.hpp"

#include <data_storage/arbstorage.hpp>
#include <data_storage/blockstore.hpp>

void deleteBlockStore(CBlockStore* m) {
    delete static_cast<BlockStore*>(m);
}

int putBlock(CBlockStore* storage_ptr,
             const void* height,
             const void* hash,
             const void* data,
             int data_length) {
    auto block_store = static_cast<BlockStore*>(storage_ptr);
    auto height_int = receiveUint256(height);
    auto hash_int = receiveUint256(hash);
    auto data_ptr = reinterpret_cast<const char*>(data);
    auto data_vector = std::vector<char>(data_ptr, data_ptr + data_length);
    return block_store->putBlock(height_int, hash_int, data_vector).ok();
}

int deleteBlock(CBlockStore* storage_ptr,
                const void* height,
                const void* hash) {
    auto block_store = static_cast<BlockStore*>(storage_ptr);
    auto height_int = receiveUint256(height);
    auto hash_int = receiveUint256(hash);
    return block_store->deleteBlock(height_int, hash_int).ok();
}

ByteSliceResult getBlock(const CBlockStore* storage_ptr,
                         const void* height,
                         const void* hash) {
    auto block_store = static_cast<const BlockStore*>(storage_ptr);
    auto height_int = receiveUint256(height);
    auto hash_int = receiveUint256(hash);
    return returnDataResult(block_store->getBlock(height_int, hash_int));
}

HashList blockHashesAtHeight(const CBlockStore* storage_ptr,
                             const void* height) {
    auto block_store = static_cast<const BlockStore*>(storage_ptr);
    auto height_int = receiveUint256(height);
    auto hashes = block_store->blockHashesAtHeight(height_int);
    std::vector<unsigned char> serializedHashes;
    for (const auto& hash : hashes) {
        marshal_uint256_t(hash, serializedHashes);
    }
    unsigned char* hashesData =
        reinterpret_cast<unsigned char*>(malloc(serializedHashes.size()));
    std::copy(serializedHashes.begin(), serializedHashes.end(), hashesData);
    return {hashesData, static_cast<int>(hashes.size())};
}

int isBlockStoreEmpty(const CBlockStore* storage_ptr) {
    auto block_store = static_cast<const BlockStore*>(storage_ptr);
    return block_store->isEmpty();
}

void* maxBlockStoreHeight(const CBlockStore* storage_ptr) {
    auto block_store = static_cast<const BlockStore*>(storage_ptr);
    auto height = block_store->maxHeight();
    return returnUint256(height);
}

void* minBlockStoreHeight(const CBlockStore* storage_ptr) {
    auto block_store = static_cast<const BlockStore*>(storage_ptr);
    auto height = block_store->minHeight();
    return returnUint256(height);
}
