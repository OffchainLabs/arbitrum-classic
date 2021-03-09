/*
 * Copyright 2021, Offchain Labs, Inc.
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

#include "data_storage/readtransaction.hpp"
#include <data_storage/storageresult.hpp>
#include "value/utils.hpp"

ReadTransaction::ReadTransaction(std::shared_ptr<DataStorage> store)
    : transaction(Transaction::makeTransaction(std::move(store))) {}

ReadTransaction::~ReadTransaction() {
    if (read_options.snapshot != nullptr) {
        transaction->datastorage->txn_db->ReleaseSnapshot(
            read_options.snapshot);
    }
}

std::unique_ptr<ReadTransaction> ReadTransaction::makeReadOnlyTransaction(
    std::shared_ptr<DataStorage> store) {
    return std::make_unique<ReadTransaction>(std::move(store));
}

void ReadTransaction::enterReadSnapshot() {
    if (read_options.snapshot == nullptr) {
        read_options.snapshot = transaction->datastorage->txn_db->GetSnapshot();
    }
}

void ReadTransaction::exitReadSnapshot() {
    if (read_options.snapshot != nullptr) {
        transaction->datastorage->txn_db->ReleaseSnapshot(
            read_options.snapshot);
        read_options.snapshot = nullptr;
    }
}

rocksdb::Status ReadTransaction::defaultGet(const rocksdb::Slice& key,
                                            std::string* value) const {
    return transaction->transaction->Get(
        read_options, transaction->datastorage->default_column.get(), key,
        value);
}
rocksdb::Status ReadTransaction::stateGet(const rocksdb::Slice& key,
                                          std::string* value) const {
    return transaction->transaction->Get(
        read_options, transaction->datastorage->state_column.get(), key, value);
}
rocksdb::Status ReadTransaction::checkpointGet(const rocksdb::Slice& key,
                                               std::string* value) const {
    return transaction->transaction->Get(
        read_options, transaction->datastorage->checkpoint_column.get(), key,
        value);
}
rocksdb::Status ReadTransaction::messageEntryGet(const rocksdb::Slice& key,
                                                 std::string* value) const {
    return transaction->transaction->Get(
        read_options, transaction->datastorage->messageentry_column.get(), key,
        value);
}
rocksdb::Status ReadTransaction::logGet(const rocksdb::Slice& key,
                                        std::string* value) const {
    return transaction->transaction->Get(
        read_options, transaction->datastorage->log_column.get(), key, value);
}
rocksdb::Status ReadTransaction::sendGet(const rocksdb::Slice& key,
                                         std::string* value) const {
    return transaction->transaction->Get(
        read_options, transaction->datastorage->send_column.get(), key, value);
}
rocksdb::Status ReadTransaction::sideloadGet(const rocksdb::Slice& key,
                                             std::string* value) const {
    return transaction->transaction->Get(
        read_options, transaction->datastorage->sideload_column.get(), key,
        value);
}

rocksdb::Status ReadTransaction::aggregatorGet(const rocksdb::Slice& key,
                                               std::string* value) const {
    return transaction->transaction->Get(
        read_options, transaction->datastorage->aggregator_column.get(), key,
        value);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::defaultGetIterator() const {
    auto it = transaction->transaction->GetIterator(
        read_options, transaction->datastorage->default_column.get());
    return std::unique_ptr<rocksdb::Iterator>(it);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::stateGetIterator() const {
    auto it = transaction->transaction->GetIterator(
        read_options, transaction->datastorage->state_column.get());
    return std::unique_ptr<rocksdb::Iterator>(it);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::checkpointGetIterator()
    const {
    auto it = transaction->transaction->GetIterator(
        read_options, transaction->datastorage->checkpoint_column.get());
    return std::unique_ptr<rocksdb::Iterator>(it);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::messageEntryGetIterator()
    const {
    auto it = transaction->transaction->GetIterator(
        read_options, transaction->datastorage->messageentry_column.get());
    return std::unique_ptr<rocksdb::Iterator>(it);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::logGetIterator() const {
    auto it = transaction->transaction->GetIterator(
        read_options, transaction->datastorage->log_column.get());
    return std::unique_ptr<rocksdb::Iterator>(it);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::sendGetIterator() const {
    auto it = transaction->transaction->GetIterator(
        read_options, transaction->datastorage->send_column.get());
    return std::unique_ptr<rocksdb::Iterator>(it);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::sideloadGetIterator()
    const {
    auto it = transaction->transaction->GetIterator(
        read_options, transaction->datastorage->sideload_column.get());
    return std::unique_ptr<rocksdb::Iterator>(it);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::aggregatorGetIterator()
    const {
    auto it = transaction->transaction->GetIterator(
        read_options, transaction->datastorage->aggregator_column.get());
    return std::unique_ptr<rocksdb::Iterator>(it);
}

ValueResult<uint256_t> ReadTransaction::defaultGetUint256(
    const rocksdb::Slice key_slice) const {
    return getUint256UsingFamilyAndKey(
        transaction->datastorage->default_column.get(), key_slice);
}

ValueResult<uint256_t> ReadTransaction::stateGetUint256(
    const rocksdb::Slice key_slice) const {
    return getUint256UsingFamilyAndKey(
        transaction->datastorage->state_column.get(), key_slice);
}

ValueResult<uint256_t> ReadTransaction::checkpointGetUint256(
    const rocksdb::Slice key_slice) const {
    return getUint256UsingFamilyAndKey(
        transaction->datastorage->checkpoint_column.get(), key_slice);
}

ValueResult<uint256_t> ReadTransaction::messageEntryGetUint256(
    const rocksdb::Slice key_slice) const {
    return getUint256UsingFamilyAndKey(
        transaction->datastorage->messageentry_column.get(), key_slice);
}

ValueResult<uint256_t> ReadTransaction::logGetUint256(
    const rocksdb::Slice key_slice) const {
    return getUint256UsingFamilyAndKey(
        transaction->datastorage->log_column.get(), key_slice);
}

ValueResult<uint256_t> ReadTransaction::sendGetUint256(
    const rocksdb::Slice key_slice) const {
    return getUint256UsingFamilyAndKey(
        transaction->datastorage->send_column.get(), key_slice);
}

ValueResult<uint256_t> ReadTransaction::sideloadGetUint256(
    const rocksdb::Slice key_slice) const {
    return getUint256UsingFamilyAndKey(
        transaction->datastorage->sideload_column.get(), key_slice);
}

ValueResult<uint256_t> ReadTransaction::aggregatorGetUint256(
    const rocksdb::Slice key_slice) const {
    return getUint256UsingFamilyAndKey(
        transaction->datastorage->aggregator_column.get(), key_slice);
}

ValueResult<std::vector<std::vector<unsigned char>>>
ReadTransaction::messageEntryGetVectorVector(
    const rocksdb::Slice first_key_slice,
    size_t count) const {
    return getVectorVectorUsingFamilyAndKey(
        transaction->datastorage->messageentry_column.get(), first_key_slice,
        count);
}

ValueResult<std::vector<std::vector<unsigned char>>>
ReadTransaction::sendGetVectorVector(const rocksdb::Slice first_key_slice,
                                     size_t count) const {
    return getVectorVectorUsingFamilyAndKey(
        transaction->datastorage->send_column.get(), first_key_slice, count);
}

ValueResult<std::vector<unsigned char>> ReadTransaction::messageEntryGetVector(
    const rocksdb::Slice first_key_slice) const {
    return getVectorUsingFamilyAndKey(
        transaction->datastorage->messageentry_column.get(), first_key_slice);
}

ValueResult<std::vector<unsigned char>> ReadTransaction::checkpointGetVector(
    const rocksdb::Slice first_key_slice) const {
    return getVectorUsingFamilyAndKey(
        transaction->datastorage->checkpoint_column.get(), first_key_slice);
}

ValueResult<std::vector<uint256_t>> ReadTransaction::logGetUint256Vector(
    const rocksdb::Slice first_key_slice,
    size_t count) const {
    return getUint256VectorUsingFamilyAndKey(
        transaction->datastorage->log_column.get(), first_key_slice, count);
}

ValueResult<std::vector<std::vector<unsigned char>>>
ReadTransaction::getVectorVectorUsingFamilyAndKey(
    rocksdb::ColumnFamilyHandle* family,
    const rocksdb::Slice first_key_slice,
    const size_t count) const {
    auto it = std::unique_ptr<rocksdb::Iterator>(
        transaction->transaction->GetIterator(read_options, family));

    // Find first message
    it->Seek(vecToSlice(first_key_slice));
    if (!it->status().ok()) {
        return {it->status(), {}};
    }

    std::vector<std::vector<unsigned char>> vectors;
    for (size_t i = 0; i < count; i++) {
        if (!it->Valid()) {
            if (!it->status().ok()) {
                return {it->status(), {}};
            }
            return {rocksdb::Status::NotFound(), {}};
        }
        vectors.emplace_back(it->value().data(),
                             it->value().data() + it->value().size());

        it->Next();
    }

    return {rocksdb::Status::OK(), std::move(vectors)};
}

ValueResult<std::vector<unsigned char>>
ReadTransaction::getVectorUsingFamilyAndKey(
    rocksdb::ColumnFamilyHandle* family,
    const rocksdb::Slice key_slice) const {
    std::string returned_value;

    auto status = transaction->transaction->Get(read_options, family, key_slice,
                                                &returned_value);
    if (!status.ok()) {
        return {status, {}};
    }

    std::vector<unsigned char> saved_value(returned_value.begin(),
                                           returned_value.end());

    return {status, std::move(saved_value)};
}

ValueResult<std::vector<uint256_t>>
ReadTransaction::getUint256VectorUsingFamilyAndKey(
    rocksdb::ColumnFamilyHandle* family,
    const rocksdb::Slice first_key_slice,
    const size_t count) const {
    auto it = std::unique_ptr<rocksdb::Iterator>(
        transaction->transaction->GetIterator(read_options, family));

    // Find first message
    it->Seek(vecToSlice(first_key_slice));
    if (!it->status().ok()) {
        return {it->status(), {}};
    }

    std::vector<uint256_t> vectors;
    for (size_t i = 0; i < count; i++) {
        if (!it->Valid()) {
            if (!it->status().ok()) {
                return {it->status(), {}};
            }
            return {rocksdb::Status::NotFound(), {}};
        }

        auto data = reinterpret_cast<const char*>(it->value().data());
        vectors.push_back(deserializeUint256t(data));

        it->Next();
    }

    return {rocksdb::Status::OK(), std::move(vectors)};
}

ValueResult<uint256_t> ReadTransaction::getUint256UsingFamilyAndKey(
    rocksdb::ColumnFamilyHandle* family,
    const rocksdb::Slice key_slice) const {
    auto result = getVectorUsingFamilyAndKey(family, key_slice);
    if (!result.status.ok()) {
        return {result.status, {}};
    }

    auto data = reinterpret_cast<const char*>(result.data.data());
    return {result.status, deserializeUint256t(data)};
}
