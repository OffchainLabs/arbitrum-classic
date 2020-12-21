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
#include <data_storage/checkpoint.hpp>
#include <data_storage/checkpointstore.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>

constexpr auto message_number_size = 32;

namespace {
std::array<char, 64> toKey(const uint256_t& height, const uint256_t& hash) {
    std::array<char, 64> key{};
    auto it = key.begin();
    it = to_big_endian(height, it);
    to_big_endian(hash, it);
    return key;
}

std::array<char, message_number_size> toKeyPrefix(const uint256_t& height) {
    std::array<char, message_number_size> key{};
    to_big_endian(height, key.begin());
    return key;
}

uint256_t keyToHeight(const rocksdb::Slice& key) {
    return intx::be::unsafe::load<uint256_t>(
        reinterpret_cast<const unsigned char*>(key.data()));
}

uint256_t keyToHash(const rocksdb::Slice& key) {
    return intx::be::unsafe::load<uint256_t>(
        reinterpret_cast<const unsigned char*>(key.data() +
                                               message_number_size));
}
}  // namespace

void CheckpointStore::saveCheckpoint(const Checkpoint& checkpoint,
                                     Machine& machine) {
    auto tx = Transaction::makeTransaction(data_storage);

    auto machine_result = saveMachine(*tx, machine);
    if (!machine_result.status.ok()) {
        throw std::runtime_error("error saving machine:" +
                                 machine_result.status.ToString());
    }

    auto checkpoint_result = putCheckpoint(*tx, checkpoint);
    if (!checkpoint_result.ok()) {
        throw std::runtime_error("error saving machine: " +
                                 checkpoint_result.ToString());
    }

    auto status = tx->commit();
    if (!status.ok()) {
        throw std::runtime_error("error saving assertion: " +
                                 status.ToString());
    }
}

void CheckpointStore::saveAssertion(const Assertion& assertion) {
    auto tx = Transaction::makeTransaction(data_storage);

    for (const auto& log : assertion.logs) {
        std::vector<unsigned char> logData;
        marshal_value(log, logData);
        AggregatorStore::saveLog(*tx->transaction, logData);
    }

    for (const auto& msg : assertion.outMessages) {
        std::vector<unsigned char> msgData;
        marshal_value(msg, msgData);
        AggregatorStore::saveMessage(*tx->transaction, msgData);
    }

    auto status = tx->commit();
    if (!status.ok()) {
        throw std::runtime_error("error saving assertion: " +
                                 status.ToString());
    }
}

rocksdb::Status CheckpointStore::deleteCheckpoint(const uint256_t& height,
                                                  const uint256_t& hash) {
    auto key = toKey(height, hash);
    rocksdb::Slice key_slice(key.begin(), key.size());
    return data_storage->txn_db->DB::Delete(
        rocksdb::WriteOptions(), data_storage->blocks_column.get(), key_slice);
}

DataResults CheckpointStore::getCheckpoint(const uint256_t& height,
                                           const uint256_t& hash) const {
    auto key = toKey(height, hash);
    rocksdb::Slice key_slice(key.begin(), key.size());
    std::string value;
    auto status = data_storage->txn_db->DB::Get(
        rocksdb::ReadOptions(), data_storage->blocks_column.get(), key_slice,
        &value);
    return {status, {value.begin(), value.end()}};
}

std::vector<uint256_t> CheckpointStore::blockHashesAtHeight(
    const uint256_t& height) const {
    std::vector<uint256_t> hashes;

    auto prefix = toKeyPrefix(height);
    rocksdb::Slice prefix_slice(prefix.begin(), prefix.size());

    auto it =
        std::unique_ptr<rocksdb::Iterator>(data_storage->txn_db->NewIterator(
            rocksdb::ReadOptions(), data_storage->blocks_column.get()));

    for (it->Seek(prefix_slice);
         it->key().starts_with(prefix_slice) && it->Valid(); it->Next()) {
        hashes.push_back(keyToHash(it->key()));
    }
    return hashes;
}

uint256_t CheckpointStore::maxHeight() const {
    auto it =
        std::unique_ptr<rocksdb::Iterator>(data_storage->txn_db->NewIterator(
            rocksdb::ReadOptions(), data_storage->blocks_column.get()));
    it->SeekToLast();
    if (it->Valid()) {
        return keyToHeight(it->key());
    } else {
        return 0;
    }
}

uint256_t CheckpointStore::minHeight() const {
    auto it =
        std::unique_ptr<rocksdb::Iterator>(data_storage->txn_db->NewIterator(
            rocksdb::ReadOptions(), data_storage->blocks_column.get()));
    it->SeekToFirst();
    if (it->Valid()) {
        return keyToHeight(it->key());
    } else {
        return 0;
    }
}

bool CheckpointStore::isEmpty() const {
    auto it =
        std::unique_ptr<rocksdb::Iterator>(data_storage->txn_db->NewIterator(
            rocksdb::ReadOptions(), data_storage->blocks_column.get()));
    it->SeekToLast();
    return !it->Valid();
}
