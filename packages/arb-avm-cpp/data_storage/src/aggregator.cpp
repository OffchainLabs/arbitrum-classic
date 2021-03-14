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

#include <data_storage/aggregator.hpp>

#include "value/utils.hpp"

#include <data_storage/datastorage.hpp>
#include <data_storage/readwritetransaction.hpp>
#include <data_storage/storageresult.hpp>

#include <rocksdb/status.h>

#include <boost/endian/conversion.hpp>

#include <sstream>

constexpr auto logs_processed_key = std::array<char, 1>{-51};

constexpr auto block_key = std::array<char, 1>{-52};

constexpr auto request_key_prefix = std::array<char, 1>{-54};
constexpr auto request_key_size = request_key_prefix.size() + 32;

constexpr auto block_hash_key_prefix = std::array<char, 1>{-55};
constexpr auto block_hash_key_size = block_hash_key_prefix.size() + 32;

constexpr auto message_batch_key_prefix = std::array<char, 1>{-56};
constexpr auto message_batch_key_size = message_batch_key_prefix.size() + 32;

namespace {

void commitTx(ReadWriteTransaction& tx) {
    auto s = tx.commit();
    if (!s.ok()) {
        throw std::runtime_error("failed to commit tx");
    }
}

template <typename Iterator>
auto addUint64ToKey(uint64_t height, Iterator it) {
    auto big_height = boost::endian::native_to_big(height);
    auto big_height_ptr = reinterpret_cast<const char*>(&big_height);
    return std::copy(big_height_ptr, big_height_ptr + sizeof(big_height), it);
}

std::array<char, request_key_size> requestKey(const uint256_t& request_id) {
    std::array<char, request_key_size> key{};
    auto it = std::copy(request_key_prefix.begin(), request_key_prefix.end(),
                        key.begin());
    to_big_endian(request_id, it);
    return key;
}

std::array<char, message_batch_key_size> messageBatchKey(
    const uint256_t& batch_number) {
    std::array<char, message_batch_key_size> key{};
    auto it = std::copy(message_batch_key_prefix.begin(),
                        message_batch_key_prefix.end(), key.begin());
    to_big_endian(batch_number, it);
    return key;
}

std::array<char, sizeof(uint64_t)> uint64Value(uint64_t height) {
    std::array<char, sizeof(uint64_t)> key{};
    addUint64ToKey(height, key.begin());
    return key;
}

std::array<char, sizeof(uint64_t)> requestValue(uint64_t log_index) {
    return uint64Value(log_index);
}

std::array<char, request_key_size> blockHashKey(const uint256_t& block_hash) {
    std::array<char, block_hash_key_size> key{};
    auto it = std::copy(block_hash_key_prefix.begin(),
                        block_hash_key_prefix.end(), key.begin());
    to_big_endian(block_hash, it);
    return key;
}

std::array<char, sizeof(uint64_t)> blockHashValue(uint64_t block_height) {
    return uint64Value(block_height);
}

void saveBlockCount(ReadWriteTransaction& tx, uint64_t max) {
    auto value = uint64Value(max);
    auto s = tx.aggregatorPut(vecToSlice(block_key), vecToSlice(value));
    if (!s.ok()) {
        throw std::runtime_error("failed to save count");
    }
}

uint64_t blockCountImpl(const ReadTransaction& tx) {
    std::string value;
    auto s = tx.aggregatorGet(vecToSlice(block_key), &value);
    if (!s.ok()) {
        throw std::runtime_error("no block count saved");
    }
    auto it = value.begin();
    return extractUint64(it);
}

void updateLogsProcessedCountImpl(ReadWriteTransaction& tx,
                                  const uint256_t& count) {
    std::vector<unsigned char> value;
    marshal_uint256_t(count, value);
    auto s =
        tx.aggregatorPut(vecToSlice(logs_processed_key), vecToSlice(value));
    if (!s.ok()) {
        throw std::runtime_error("filed to save processed count");
    }
}
}  // namespace

std::array<char, block_key.size() + sizeof(uint64_t)> blockEntryKey(
    uint64_t index) {
    std::array<char, block_key.size() + sizeof(uint64_t)> full_key{};
    auto it = std::copy(block_key.begin(), block_key.end(), full_key.begin());
    addUint64ToKey(index, it);
    return full_key;
}

namespace {
template <typename Key>
std::optional<uint64_t> returnIndex(const ReadTransaction& tx, const Key& key) {
    std::string request_value;
    auto s = tx.aggregatorGet(vecToSlice(key), &request_value);
    if (s.IsNotFound()) {
        return std::nullopt;
    }
    if (!s.ok()) {
        throw std::runtime_error("failed to load index");
    }
    auto it = request_value.begin();
    return extractUint64(it);
}
}  // namespace

AggregatorStore::AggregatorStore(std::shared_ptr<DataStorage> data_storage_)
    : data_storage(std::move(data_storage_)) {
    ReadWriteTransaction tx(data_storage);
    std::string value;
    auto s = tx.aggregatorGet(vecToSlice(block_key), &value);
    if (s.IsNotFound()) {
        saveBlockCount(tx, 0);
        updateLogsProcessedCountImpl(tx, 0);
    }
    commitTx(tx);
}

std::optional<uint64_t> AggregatorStore::getPossibleRequestInfo(
    const uint256_t& request_id) const {
    ReadTransaction tx(data_storage);
    return returnIndex(tx, requestKey(request_id));
}

std::optional<uint64_t> AggregatorStore::getPossibleBlock(
    const uint256_t& block_hash) const {
    ReadTransaction tx(data_storage);
    return returnIndex(tx, blockHashKey(block_hash));
}

uint64_t AggregatorStore::blockCount() const {
    ReadTransaction tx(data_storage);
    return blockCountImpl(tx);
}

void AggregatorStore::saveMessageBatch(const uint256_t& batchNum,
                                       const uint64_t& logIndex) {
    ReadWriteTransaction tx(data_storage);
    auto full_key = messageBatchKey(batchNum);
    auto index_value = uint64Value(logIndex);

    auto status =
        tx.aggregatorPut(vecToSlice(full_key), vecToSlice(index_value));
    if (!status.ok()) {
        throw std::runtime_error("failed to save");
    }

    commitTx(tx);
}

std::optional<uint64_t> AggregatorStore::getMessageBatch(
    const uint256_t& batchNum) {
    ReadTransaction tx(data_storage);
    return returnIndex(tx, messageBatchKey(batchNum));
}

void AggregatorStore::saveBlock(uint64_t height,
                                const uint256_t& block_hash,
                                const std::vector<uint256_t>& requests,
                                const uint64_t* log_indexes,
                                const std::vector<char>& data) {
    ReadWriteTransaction tx(data_storage);
    auto block_hash_key = blockHashKey(block_hash);
    auto block_value = blockHashValue(height);
    auto s =
        tx.aggregatorPut(vecToSlice(block_hash_key), vecToSlice(block_value));
    if (!s.ok()) {
        throw std::runtime_error("failed to save block hash");
    }

    for (size_t i = 0; i < requests.size(); i++) {
        auto request_key = requestKey(requests[i]);
        auto request_value = requestValue(log_indexes[i]);
        s = tx.aggregatorPut(vecToSlice(request_key),
                             vecToSlice(request_value));
        if (!s.ok()) {
            throw std::runtime_error("failed to save request");
        }
    }

    uint64_t current_count = blockCountImpl(tx);
    if (height > current_count) {
        std::stringstream ss;
        ss << "tried to save block with unexpected height: got" << height
           << " but expected " << current_count;
        throw std::runtime_error(ss.str());
    }
    auto full_key = blockEntryKey(height);
    s = tx.aggregatorPut(vecToSlice(full_key), vecToSlice(data));
    if (!s.ok()) {
        throw std::runtime_error("failed to save");
    }
    saveBlockCount(tx, height + 1);
    commitTx(tx);
}

std::vector<char> AggregatorStore::getBlock(uint64_t height) const {
    ReadSnapshotTransaction tx(data_storage);
    uint64_t current_count = blockCountImpl(tx);
    if (height >= current_count) {
        std::stringstream ss;
        ss << "invalid index " << height << " with count " << current_count;
        throw std::runtime_error(ss.str());
    }
    auto full_key = blockEntryKey(height);
    std::string value;
    auto s = tx.aggregatorGet(vecToSlice(full_key), &value);
    if (!s.ok()) {
        throw std::runtime_error("failed load value");
    }
    return {value.begin(), value.end()};
}

void AggregatorStore::reorg(uint64_t block_height) {
    ReadWriteTransaction tx(data_storage);
    saveBlockCount(tx, block_height);
    commitTx(tx);
}

ValueResult<uint256_t> AggregatorStore::logsProcessedCount() const {
    ReadTransaction tx(data_storage);
    return tx.aggregatorGetUint256(vecToSlice(logs_processed_key));
}

void AggregatorStore::updateLogsProcessedCount(const uint256_t& count) {
    ReadWriteTransaction tx(data_storage);
    updateLogsProcessedCountImpl(tx, count);
    commitTx(tx);
}
