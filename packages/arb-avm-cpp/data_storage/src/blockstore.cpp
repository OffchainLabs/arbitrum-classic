//
//  blockstore.cpp
//  avm
//
//  Created by Harry Kalodner on 5/17/20.
//

#include <data_storage/blockstore.hpp>
#include <data_storage/storageresult.hpp>

#include <bigint_utils.hpp>

#include <rocksdb/status.h>
#include <rocksdb/utilities/transaction_db.h>

constexpr auto height_size = 32;
constexpr auto hash_size = 32;

namespace {
std::array<char, 64> toKey(const uint256_t& height, const uint256_t& hash) {
    std::array<char, 64> key;
    to_big_endian(height, key.begin());
    to_big_endian(hash, key.begin() + height_size);
    return key;
}

std::array<char, height_size> toKeyPrefix(const uint256_t& height) {
    std::array<char, height_size> key;
    to_big_endian(height, key.begin());
    return key;
}

uint256_t keyToHeight(const rocksdb::Slice& key) {
    return from_big_endian(key.data(), key.data() + height_size);
}

uint256_t keyToHash(const rocksdb::Slice& key) {
    return from_big_endian(key.data() + height_size,
                           key.data() + height_size + hash_size);
}
}  // namespace

rocksdb::Status BlockStore::putBlock(const uint256_t& height,
                                     const uint256_t& hash,
                                     const std::vector<char>& value) {
    auto key = toKey(height, hash);
    rocksdb::Slice key_slice(key.begin(), key.size());
    rocksdb::Slice value_slice(value.data(), value.size());
    return txn_db->DB::Put(rocksdb::WriteOptions(), blocks_column.get(),
                           key_slice, value_slice);
}

rocksdb::Status BlockStore::deleteBlock(const uint256_t& height,
                                        const uint256_t& hash) {
    auto key = toKey(height, hash);
    rocksdb::Slice key_slice(key.begin(), key.size());
    return txn_db->DB::Delete(rocksdb::WriteOptions(), blocks_column.get(),
                              key_slice);
}

DataResults BlockStore::getBlock(const uint256_t& height,
                                 const uint256_t& hash) const {
    auto key = toKey(height, hash);
    rocksdb::Slice key_slice(key.begin(), key.size());
    std::string value;
    auto status = txn_db->DB::Get(rocksdb::ReadOptions(), blocks_column.get(),
                                  key_slice, &value);
    return {status, {value.begin(), value.end()}};
}

std::vector<uint256_t> BlockStore::blockHashesAtHeight(
    const uint256_t& height) const {
    std::vector<uint256_t> hashes;

    auto prefix = toKeyPrefix(height);
    rocksdb::Slice prefix_slice(prefix.begin(), prefix.size());

    auto it = std::unique_ptr<rocksdb::Iterator>(
        txn_db->NewIterator(rocksdb::ReadOptions(), blocks_column.get()));

    for (it->Seek(prefix_slice);
         it->key().starts_with(prefix_slice) && it->Valid(); it->Next()) {
        hashes.push_back(keyToHash(it->key()));
    }
    return hashes;
}

uint256_t BlockStore::maxHeight() const {
    auto it = std::unique_ptr<rocksdb::Iterator>(
        txn_db->NewIterator(rocksdb::ReadOptions(), blocks_column.get()));
    it->SeekToLast();
    if (it->Valid()) {
        return keyToHeight(it->key());
    } else {
        return 0;
    }
}

uint256_t BlockStore::minHeight() const {
    auto it = std::unique_ptr<rocksdb::Iterator>(
        txn_db->NewIterator(rocksdb::ReadOptions(), blocks_column.get()));
    it->SeekToFirst();
    if (it->Valid()) {
        return keyToHeight(it->key());
    } else {
        return 0;
    }
}

bool BlockStore::isEmpty() const {
    auto it = std::unique_ptr<rocksdb::Iterator>(
        txn_db->NewIterator(rocksdb::ReadOptions(), blocks_column.get()));
    it->SeekToLast();
    return !it->Valid();
}
