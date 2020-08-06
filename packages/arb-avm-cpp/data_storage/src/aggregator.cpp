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
#include <data_storage/storageresult.hpp>

#include <rocksdb/status.h>
#include <rocksdb/utilities/transaction_db.h>

#include <boost/endian/conversion.hpp>

#include <sstream>

constexpr auto log_key = std::array<char, 1>{-50};
constexpr auto message_key = std::array<char, 1>{-51};
constexpr auto block_key = std::array<char, 1>{-52};

constexpr auto request_key_prefix = std::array<char, 1>{-54};
constexpr auto request_key_size = request_key_prefix.size() + 32;

namespace {

void commitTx(rocksdb::Transaction& tx) {
    auto s = tx.Commit();
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

template <typename Iterator>
uint64_t extractUint64(Iterator& it) {
    uint64_t big_height;
    auto big_height_ptr = reinterpret_cast<char*>(&big_height);
    std::copy(it, it + sizeof(big_height), big_height_ptr);
    it += sizeof(uint64_t);
    return boost::endian::big_to_native(big_height);
}

std::array<char, request_key_size> requestKey(const uint256_t& request_id) {
    std::array<char, request_key_size> key;
    auto it = std::copy(request_key_prefix.begin(), request_key_prefix.end(),
                        key.begin());
    to_big_endian(request_id, it);
    return key;
}

std::array<char, sizeof(uint64_t)> uint64Value(uint64_t height) {
    std::array<char, sizeof(uint64_t)> key;
    addUint64ToKey(height, key.begin());
    return key;
}

std::array<char, sizeof(uint64_t)> requestValue(uint64_t log_index) {
    std::array<char, sizeof(uint64_t)> key;
    addUint64ToKey(log_index, key.begin());
    return key;
}
}  // namespace

template <size_t N, const std::array<char, N>& key>
struct EntrySaver {
    std::array<char, N + sizeof(uint64_t)> entryKey(uint64_t index) {
        std::array<char, N + sizeof(uint64_t)> full_key;
        auto it = std::copy(key.begin(), key.end(), full_key.begin());
        addUint64ToKey(index, it);
        return full_key;
    }

    void saveIndex(rocksdb::Transaction& tx, uint64_t count) {
        auto value = uint64Value(count);
        auto s = tx.Put(vecToSlice(key), vecToSlice(value));
        if (!s.ok()) {
            throw std::runtime_error("failed to save count");
        }
    }

    std::string loadEntry(rocksdb::Transaction& tx, uint64_t index) {
        auto full_key = this->entryKey(index);
        std::string value;
        auto s = tx.Get(rocksdb::ReadOptions{}, vecToSlice(full_key), &value);
        if (!s.ok()) {
            throw std::runtime_error("failed load value");
        }
        return value;
    }

    template <typename T>
    void saveEntry(rocksdb::Transaction& tx, const T& output, uint64_t height) {
        auto full_key = this->entryKey(height);
        auto s = tx.Put(vecToSlice(full_key), vecToSlice(output));
        if (!s.ok()) {
            throw std::runtime_error("failed to save");
        }
    }
};

template <size_t N, const std::array<char, N>& key>
struct FlatSaver : private EntrySaver<N, key> {
    uint64_t count(rocksdb::Transaction& tx) {
        std::string value;
        auto s = tx.Get(rocksdb::ReadOptions{}, vecToSlice(key), &value);
        if (s.IsNotFound()) {
            return 0;
        } else if (!s.ok()) {
            throw std::runtime_error("failed to load count");
        }
        auto it = value.begin();
        return extractUint64(it);
    }

    void saveCount(rocksdb::Transaction& tx, uint64_t count) {
        this->saveIndex(tx, count);
    }

    template <typename T>
    void saveNext(rocksdb::Transaction& tx, const T& output) {
        uint64_t current_count = count(tx);
        this->saveEntry(tx, output, current_count);
        saveCount(tx, current_count + 1);
    }

    std::string load(rocksdb::Transaction& tx, uint64_t index) {
        uint64_t current_count = count(tx);
        if (index >= current_count) {
            std::stringstream ss;
            ss << "invalid index " << index << "/" << current_count;
            throw std::runtime_error(ss.str());
        }
        return this->loadEntry(tx, index);
    }
};

template <size_t N, const std::array<char, N>& key>
struct HeightSaver : public EntrySaver<N, key> {
    uint64_t max(rocksdb::Transaction& tx) {
        std::string value;
        auto s = tx.Get(rocksdb::ReadOptions{}, vecToSlice(key), &value);
        if (!s.ok()) {
            throw std::runtime_error("no max saved");
        }
        auto it = value.begin();
        return extractUint64(it);
    }

    void saveMax(rocksdb::Transaction& tx, uint64_t max) {
        this->saveIndex(tx, max);
    }

    template <typename T>
    void save(rocksdb::Transaction& tx, const T& output, uint64_t height) {
        this->saveEntry(tx, output, height);
        saveMax(tx, height);
    }

    std::string load(rocksdb::Transaction& tx, uint64_t index) {
        uint64_t current_max = max(tx);
        if (index > current_max) {
            std::stringstream ss;
            ss << "invalid index " << index << "/" << current_max;
            throw std::runtime_error(ss.str());
        }
        return this->loadEntry(tx, index);
    }
};

using LogSaver = FlatSaver<log_key.size(), log_key>;
using MessageSaver = FlatSaver<message_key.size(), message_key>;
using BlockSaver = HeightSaver<block_key.size(), block_key>;

uint64_t AggregatorStore::logCount() const {
    auto tx = data_storage->beginTransaction();
    return LogSaver{}.count(*tx);
}

void AggregatorStore::saveLog(const std::vector<char>& log) {
    auto tx = data_storage->beginTransaction();
    LogSaver{}.saveNext(*tx, log);
    commitTx(*tx);
}

std::vector<char> AggregatorStore::getLog(uint64_t index) const {
    auto tx = data_storage->beginTransaction();
    auto value = LogSaver{}.load(*tx, index);
    return {value.begin(), value.end()};
}

uint64_t AggregatorStore::messageCount() const {
    auto tx = data_storage->beginTransaction();
    return MessageSaver{}.count(*tx);
}

void AggregatorStore::saveMessage(const std::vector<char>& output) {
    auto tx = data_storage->beginTransaction();
    MessageSaver{}.saveNext(*tx, output);
    commitTx(*tx);
}

std::vector<char> AggregatorStore::getMessage(uint64_t index) const {
    auto tx = data_storage->beginTransaction();
    auto value = MessageSaver{}.load(*tx, index);
    return {value.begin(), value.end()};
}

void AggregatorStore::saveRequest(const uint256_t& request_id,
                                  uint64_t log_index) {
    auto key = requestKey(request_id);
    auto value = requestValue(log_index);
    auto s = data_storage->txn_db->Put(rocksdb::WriteOptions{}, vecToSlice(key),
                                       vecToSlice(value));
    if (!s.ok()) {
        throw std::runtime_error("failed to save request");
    }
}

uint64_t AggregatorStore::getPossibleRequestInfo(
    const uint256_t& request_id) const {
    auto tx = data_storage->beginTransaction();
    auto key = requestKey(request_id);
    std::string request_value;
    auto s = tx->Get(rocksdb::ReadOptions{}, vecToSlice(key), &request_value);
    if (!s.ok()) {
        throw std::runtime_error("couldn't find request");
    }
    auto it = request_value.begin();
    uint64_t log_index = extractUint64(it);
    return log_index;
}

std::pair<uint64_t, std::vector<char>> AggregatorStore::latestBlock() const {
    auto tx = data_storage->beginTransaction();
    uint64_t latest_block = BlockSaver{}.max(*tx);
    return {latest_block, getBlock(latest_block)};
}

void AggregatorStore::saveBlock(uint64_t height,
                                const std::vector<char>& data) {
    auto tx = data_storage->beginTransaction();
    ;
    BlockSaver{}.save(*tx, data, height);
    commitTx(*tx);
}

std::vector<char> AggregatorStore::getBlock(uint64_t height) const {
    auto tx = data_storage->beginTransaction();
    auto block_value = BlockSaver{}.load(*tx, height);
    return {block_value.begin(), block_value.end()};
}

void AggregatorStore::reorg(uint64_t block_height,
                            uint64_t message_count,
                            uint64_t log_count) {
    auto tx = data_storage->beginTransaction();
    MessageSaver{}.saveCount(*tx, message_count);
    LogSaver{}.saveCount(*tx, log_count);
    BlockSaver{}.saveMax(*tx, block_height);
    commitTx(*tx);
}
