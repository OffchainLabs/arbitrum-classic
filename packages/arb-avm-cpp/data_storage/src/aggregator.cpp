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

#include <bigint_utils.hpp>

#include <rocksdb/status.h>
#include <rocksdb/utilities/transaction_db.h>

constexpr auto log_key = std::array<char, 1>{3};
constexpr auto message_key = std::array<char, 1>{3};
constexpr auto block_key = std::array<char, 1>{3};
constexpr auto initial_block_key = std::array<char, 1>{3};

constexpr auto request_key_prefix = std::array<char, 1>{1};
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

struct BlockSaveData {
    uint256_t hash;
    uint64_t log_count;
    uint64_t message_count;
    uint256_t bloom;
};

std::array<char, 32 * 2 + sizeof(uint64_t) * 2> blockValue(
    const uint256_t& hash,
    uint64_t log_count,
    uint64_t message_count,
    const uint256_t& bloom) {
    std::array<char, 32 * 2 + sizeof(uint64_t) * 2> key;
    auto it = to_big_endian(hash, key.begin());
    it = addUint64ToKey(log_count, it);
    it = addUint64ToKey(message_count, it);
    to_big_endian(bloom, it);
    return key;
}

BlockSaveData processBlockValue(const std::string& value) {
    auto it = value.begin();
    auto hash = from_big_endian(it, it + 32);
    it += 32;
    uint64_t log_count = extractUint64(it);
    uint64_t message_count = extractUint64(it);
    auto bloom = from_big_endian(it, it + 32);
    return {hash, log_count, message_count, bloom};
}

std::array<char, sizeof(uint64_t) * 2> requestValue(
    uint64_t log_index,
    uint64_t evm_start_log_index) {
    std::array<char, sizeof(uint64_t) * 2> key;
    auto it = addUint64ToKey(log_index, key.begin());
    addUint64ToKey(evm_start_log_index, it);
    return key;
}
}  // namespace

template <size_t N, const std::array<char, N>& key>
struct FlatSaver {
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

    std::array<char, N + sizeof(uint64_t)> entryKey(uint64_t index) {
        std::array<char, N + sizeof(uint64_t)> full_key;
        auto it = std::copy(key.begin(), key.end(), full_key.begin());
        addUint64ToKey(index, it);
        return full_key;
    }

    void saveCount(rocksdb::Transaction& tx, uint64_t count) {
        auto value = uint64Value(count);
        auto s = tx.Put(vecToSlice(key), vecToSlice(value));
        if (!s.ok()) {
            throw std::runtime_error("failed to save count");
        }
    }

    template <typename T>
    void saveNext(rocksdb::Transaction& tx, const T& output) {
        uint64_t current_count = count(tx);
        save(tx, output, current_count);
        saveCount(tx, current_count + 1);
    }

    template <typename T>
    void save(rocksdb::Transaction& tx, const T& output, uint64_t index) {
        auto full_key = entryKey(index);
        auto s = tx.Put(vecToSlice(full_key), vecToSlice(output));
        if (!s.ok()) {
            throw std::runtime_error("failed to save");
        }
    }

    std::string load(rocksdb::Transaction& tx, uint64_t index) {
        uint64_t current_count = count(tx);
        if (index >= current_count) {
            std::stringstream ss;
            ss << "invalid index " << index << "/" << current_count;
            throw std::runtime_error(ss.str());
        }
        auto full_key = entryKey(index);
        std::string value;
        auto s = tx.Get(rocksdb::ReadOptions{}, vecToSlice(full_key), &value);
        if (!s.ok()) {
            throw std::runtime_error("failed load value");
        }
        return value;
    }
};

using LogSaver = FlatSaver<log_key.size(), log_key>;
using MessageSaver = FlatSaver<message_key.size(), message_key>;
using BlockSaver = FlatSaver<block_key.size(), block_key>;

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
                                  uint64_t log_index,
                                  uint64_t evm_start_log_index) {
    auto key = requestKey(request_id);
    auto value = requestValue(log_index, evm_start_log_index);
    auto s = data_storage->txn_db->Put(rocksdb::WriteOptions{}, vecToSlice(key),
                                       vecToSlice(value));
    if (!s.ok()) {
        throw std::runtime_error("failed to save request");
    }
}

std::pair<uint64_t, uint64_t> AggregatorStore::getPossibleRequestInfo(
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
    uint64_t evm_start_log_index = extractUint64(it);
    return {log_index, evm_start_log_index};
}

std::pair<uint64_t, uint256_t> AggregatorStore::latestBlock() const {
    auto tx = data_storage->beginTransaction();
    uint64_t block_count = BlockSaver{}.count(*tx);
    if (block_count == 0) {
        throw std::runtime_error("no blocks in db");
    }
    auto block = getBlock(block_count - 1);
    return {block_count - 1, block.hash};
}

uint64_t AggregatorStore::getInitialBlock() const {
    std::string value;
    auto s = data_storage->txn_db->Get(rocksdb::ReadOptions{},
                                       vecToSlice(initial_block_key), &value);
    if (!s.ok()) {
        throw std::runtime_error("couldn't load initial block");
    }
    auto it = value.begin();
    return extractUint64(it);
}

void AggregatorStore::saveBlock(uint64_t height,
                                const uint256_t& hash,
                                const uint256_t& bloom) {
    auto tx = data_storage->beginTransaction();
    auto value = blockValue(hash, logCount(), messageCount(), bloom);

    std::string initial_value;
    auto s = tx->Get(rocksdb::ReadOptions{}, vecToSlice(initial_block_key),
                     &initial_value);
    if (s.IsNotFound()) {
        auto initial_val = uint64Value(height);
        auto s =
            tx->Put(vecToSlice(initial_block_key), vecToSlice(initial_val));
        if (!s.ok()) {
            throw std::runtime_error("couldn't save initial block");
        }
        BlockSaver{}.save(*tx, value, height);
    } else if (!s.ok()) {
        throw std::runtime_error("couldn't load initial block");
    } else {
        auto current_count = BlockSaver{}.count(*tx);
        if (height != current_count) {
            throw std::runtime_error("must save consecutive blocks");
        }
        BlockSaver{}.saveNext(*tx, value);
    }
    commitTx(*tx);
}

BlockData AggregatorStore::getBlock(uint64_t height) const {
    auto tx = data_storage->beginTransaction();
    auto block_value = BlockSaver{}.load(*tx, height);
    auto parsed_value = processBlockValue(block_value);
    auto initial = getInitialBlock();

    uint64_t prev_log_count = 0;
    uint64_t prev_message_count = 0;
    if (height > initial) {
        auto prev_block_value = BlockSaver{}.load(*tx, height - 1);
        auto prev_parsed_value = processBlockValue(prev_block_value);
        prev_log_count = prev_parsed_value.log_count;
        prev_message_count = prev_parsed_value.message_count;
    }

    return {parsed_value.hash,
            prev_log_count,
            parsed_value.log_count - prev_log_count,
            prev_message_count,
            parsed_value.message_count - prev_message_count,
            parsed_value.bloom};
}

void AggregatorStore::restoreBlock(uint64_t height) {
    auto tx = data_storage->beginTransaction();
    auto block_value = BlockSaver{}.load(*tx, height);
    auto parsed_value = processBlockValue(block_value);
    MessageSaver{}.saveCount(*tx, parsed_value.message_count);
    LogSaver{}.saveCount(*tx, parsed_value.log_count);
    BlockSaver{}.saveCount(*tx, height);
    commitTx(*tx);
}
