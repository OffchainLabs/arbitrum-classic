//
//  cblockstore.cpp
//  avm
//
//  Created by Harry Kalodner on 5/27/20.
//

#include "cblockstore.h"

#include "utils.hpp"

#include <data_storage/blockstore.hpp>
#include <data_storage/checkpoint/checkpointstorage.hpp>

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
    unsigned char* hashesData = (unsigned char*)malloc(serializedHashes.size());
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
