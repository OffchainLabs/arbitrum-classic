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

#include <data_storage/confirmednodestore.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>

#include <bigint_utils.hpp>

#include <rocksdb/status.h>
#include <rocksdb/utilities/transaction_db.h>

constexpr auto node_key_prefix = std::array<char, 1>{0};
constexpr auto node_hash_key_prefix = std::array<char, 1>{1};
constexpr auto node_height_key_prefix = std::array<char, 1>{2};
constexpr auto height_size = sizeof(uint64_t);
constexpr auto hash_size = 32;
constexpr auto node_key_size = node_key_prefix.size() + height_size + hash_size;
constexpr auto node_hash_key_size = node_hash_key_prefix.size() + hash_size;
constexpr auto node_hash_value_size = height_size;
constexpr auto node_height_key_size =
    node_height_key_prefix.size() + height_size;
constexpr auto node_height_value_size = hash_size;

constexpr auto chain_count_key = std::array<char, 1>{3};

namespace {

template <typename Iterator>
auto addHeightToKey(uint64_t height, Iterator it) {
    auto big_height = boost::endian::native_to_big(height);
    auto big_height_ptr = reinterpret_cast<const char*>(&big_height);
    return std::copy(big_height_ptr, big_height_ptr + sizeof(big_height), it);
}

std::array<char, node_key_size> toNodeKey(uint64_t height,
                                          const uint256_t& hash) {
    std::array<char, node_key_size> key;
    auto it =
        std::copy(node_key_prefix.begin(), node_key_prefix.end(), key.begin());
    it = addHeightToKey(height, it);
    to_big_endian(hash, it);
    return key;
}

std::array<char, node_hash_key_size> toNodeHashKey(const uint256_t& hash) {
    std::array<char, node_hash_key_size> key;
    auto it = std::copy(node_hash_key_prefix.begin(),
                        node_hash_key_prefix.end(), key.begin());
    to_big_endian(hash, it);
    return key;
}

std::array<char, node_hash_value_size> toNodeHashValue(uint64_t height) {
    std::array<char, node_hash_value_size> key;
    addHeightToKey(height, key.begin());
    return key;
}

std::array<char, node_height_key_size> toNodeHeightKey(uint64_t height) {
    std::array<char, node_height_key_size> key;
    auto it = std::copy(node_height_key_prefix.begin(),
                        node_height_key_prefix.end(), key.begin());
    addHeightToKey(height, it);
    return key;
}

std::array<char, node_height_value_size> toNodeHeightValue(
    const uint256_t& hash) {
    std::array<char, node_height_value_size> key;
    to_big_endian(hash, key.begin());
    return key;
}

uint64_t valueToHeight(const std::string& value) {
    uint64_t big_height;
    auto big_height_ptr = reinterpret_cast<char*>(&big_height);
    std::copy(value.data(), value.data() + value.size(), big_height_ptr);
    return boost::endian::big_to_native(big_height);
}

uint256_t valueToHash(const std::string& value) {
    return from_big_endian(value.data(), value.data() + value.size());
}
}  // namespace

rocksdb::Status ConfirmedNodeStore::putNode(uint64_t height,
                                            const uint256_t& hash,
                                            const std::vector<char>& value) {
    auto transaction = std::unique_ptr<rocksdb::Transaction>(
        data_storage->txn_db->BeginTransaction(rocksdb::WriteOptions()));

    auto key = toNodeKey(height, hash);
    rocksdb::Slice node_key_slice(key.begin(), key.size());
    rocksdb::Slice node_value_slice(value.data(), value.size());

    auto s = transaction->Put(data_storage->node_column.get(), node_key_slice,
                              node_value_slice);
    if (!s.ok()) {
        transaction->Rollback();
        return s;
    }

    auto node_hash_key = toNodeHashKey(hash);
    auto node_hash_value = toNodeHashValue(height);
    rocksdb::Slice node_hash_key_slice(node_hash_key.begin(),
                                       node_hash_key.size());
    rocksdb::Slice node_hash_value_slice(node_hash_value.begin(),
                                         node_hash_value.size());

    s = transaction->Put(data_storage->node_column.get(), node_hash_key_slice,
                         node_hash_value_slice);
    if (!s.ok()) {
        transaction->Rollback();
        return s;
    }

    auto node_height_key = toNodeHeightKey(height);
    auto node_height_value = toNodeHeightValue(hash);
    rocksdb::Slice node_height_key_slice(node_height_key.begin(),
                                         node_height_key.size());
    rocksdb::Slice node_height_value_slice(node_height_value.begin(),
                                           node_height_value.size());

    s = transaction->Put(data_storage->node_column.get(), node_height_key_slice,
                         node_height_value_slice);
    if (!s.ok()) {
        transaction->Rollback();
        return s;
    }

    rocksdb::Slice chain_count_key_slice(chain_count_key.begin(),
                                         chain_count_key.size());
    std::string chain_count_value;
    s = transaction->Get(rocksdb::ReadOptions(),
                         data_storage->node_column.get(), chain_count_key_slice,
                         &chain_count_value);

    uint64_t current_max_height = 0;
    if (s.ok()) {
        current_max_height = valueToHeight(chain_count_value);
    }

    if (height > current_max_height) {
        s = transaction->Put(data_storage->node_column.get(),
                             chain_count_key_slice, node_hash_value_slice);
        if (!s.ok()) {
            transaction->Rollback();
            return s;
        }
    }

    return transaction->Commit();
}

DataResults ConfirmedNodeStore::getNode(uint64_t height,
                                        const uint256_t& hash) const {
    auto key = toNodeKey(height, hash);
    rocksdb::Slice key_slice(key.begin(), key.size());
    std::string value;
    auto status = data_storage->txn_db->DB::Get(rocksdb::ReadOptions(),
                                                data_storage->node_column.get(),
                                                key_slice, &value);
    return {status, {value.begin(), value.end()}};
}

ValueResult<uint64_t> ConfirmedNodeStore::getHeight(
    const uint256_t& hash) const {
    auto key = toNodeHashKey(hash);
    rocksdb::Slice key_slice(key.begin(), key.size());
    std::string value;
    auto status = data_storage->txn_db->DB::Get(rocksdb::ReadOptions(),
                                                data_storage->node_column.get(),
                                                key_slice, &value);

    if (!status.ok()) {
        return {status, 0};
    }
    return {status, valueToHeight(value)};
}

ValueResult<uint256_t> ConfirmedNodeStore::getHash(uint64_t height) const {
    auto key = toNodeHeightKey(height);
    rocksdb::Slice key_slice(key.begin(), key.size());
    std::string value;
    auto status = data_storage->txn_db->DB::Get(rocksdb::ReadOptions(),
                                                data_storage->node_column.get(),
                                                key_slice, &value);

    if (!status.ok()) {
        return {status, 0};
    }
    return {status, valueToHash(value)};
}

bool ConfirmedNodeStore::isEmpty() const {
    auto it =
        std::unique_ptr<rocksdb::Iterator>(data_storage->txn_db->NewIterator(
            rocksdb::ReadOptions(), data_storage->node_column.get()));
    it->SeekToLast();
    return !it->Valid();
}

uint64_t ConfirmedNodeStore::maxNodeHeight() const {
    rocksdb::Slice key_slice(chain_count_key.begin(), chain_count_key.size());
    std::string value;
    auto status = data_storage->txn_db->DB::Get(rocksdb::ReadOptions(),
                                                data_storage->node_column.get(),
                                                key_slice, &value);

    if (!status.ok()) {
        return 0;
    }
    return valueToHeight(value);
}
