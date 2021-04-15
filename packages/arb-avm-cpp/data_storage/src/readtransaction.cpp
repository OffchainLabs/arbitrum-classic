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
#include <data_storage/value/utils.hpp>

ReadTransaction::ReadTransaction(std::shared_ptr<DataStorage> store)
    : transaction(Transaction::makeTransaction(std::move(store))) {}

rocksdb::Status ReadTransaction::defaultGet(const rocksdb::Slice& key,
                                            std::string* value) const {
    return transaction->transaction->Get(
        read_options,
        transaction->datastorage->column_handles[DataStorage::DEFAULT_COLUMN],
        key, value);
}
rocksdb::Status ReadTransaction::stateGet(const rocksdb::Slice& key,
                                          std::string* value) const {
    return transaction->transaction->Get(
        read_options,
        transaction->datastorage->column_handles[DataStorage::STATE_COLUMN],
        key, value);
}
rocksdb::Status ReadTransaction::checkpointGet(const rocksdb::Slice& key,
                                               std::string* value) const {
    return transaction->transaction->Get(
        read_options,
        transaction->datastorage
            ->column_handles[DataStorage::CHECKPOINT_COLUMN],
        key, value);
}
rocksdb::Status ReadTransaction::logGet(const rocksdb::Slice& key,
                                        std::string* value) const {
    return transaction->transaction->Get(
        read_options,
        transaction->datastorage->column_handles[DataStorage::LOG_COLUMN], key,
        value);
}
rocksdb::Status ReadTransaction::sendGet(const rocksdb::Slice& key,
                                         std::string* value) const {
    return transaction->transaction->Get(
        read_options,
        transaction->datastorage->column_handles[DataStorage::SEND_COLUMN], key,
        value);
}
rocksdb::Status ReadTransaction::sideloadGet(const rocksdb::Slice& key,
                                             std::string* value) const {
    return transaction->transaction->Get(
        read_options,
        transaction->datastorage->column_handles[DataStorage::SIDELOAD_COLUMN],
        key, value);
}

rocksdb::Status ReadTransaction::aggregatorGet(const rocksdb::Slice& key,
                                               std::string* value) const {
    return transaction->transaction->Get(
        read_options,
        transaction->datastorage
            ->column_handles[DataStorage::AGGREGATOR_COLUMN],
        key, value);
}

rocksdb::Status ReadTransaction::refCountedGet(const rocksdb::Slice& key,
                                               std::string* value) const {
    return transaction->transaction->Get(
        read_options,
        transaction->datastorage
            ->column_handles[DataStorage::REFCOUNTED_COLUMN],
        key, value);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::defaultGetIterator() const {
    auto it = transaction->transaction->GetIterator(
        read_options,
        transaction->datastorage->column_handles[DataStorage::DEFAULT_COLUMN]);
    return std::unique_ptr<rocksdb::Iterator>(it);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::stateGetIterator() const {
    auto it = transaction->transaction->GetIterator(
        read_options,
        transaction->datastorage->column_handles[DataStorage::STATE_COLUMN]);
    return std::unique_ptr<rocksdb::Iterator>(it);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::checkpointGetIterator()
    const {
    auto it = transaction->transaction->GetIterator(
        read_options, transaction->datastorage
                          ->column_handles[DataStorage::CHECKPOINT_COLUMN]);
    return std::unique_ptr<rocksdb::Iterator>(it);
}

std::unique_ptr<rocksdb::Iterator>
ReadTransaction::sequencerBatchItemGetIterator(
    rocksdb::Slice* lower_bound,
    rocksdb::Slice* upper_bound) const {
    auto read_opts = read_options;
    read_opts.iterate_lower_bound = lower_bound;
    read_opts.iterate_upper_bound = upper_bound;
    auto it = transaction->transaction->GetIterator(
        read_options,
        transaction->datastorage
            ->column_handles[DataStorage::SEQUENCERBATCHITEM_COLUMN]);
    return std::unique_ptr<rocksdb::Iterator>(it);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::delayedMessageGetIterator(
    rocksdb::Slice* lower_bound,
    rocksdb::Slice* upper_bound) const {
    auto read_opts = read_options;
    read_opts.iterate_lower_bound = lower_bound;
    read_opts.iterate_upper_bound = upper_bound;
    auto it = transaction->transaction->GetIterator(
        read_options, transaction->datastorage
                          ->column_handles[DataStorage::DELAYEDMESSAGE_COLUMN]);
    return std::unique_ptr<rocksdb::Iterator>(it);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::logGetIterator() const {
    auto it = transaction->transaction->GetIterator(
        read_options,
        transaction->datastorage->column_handles[DataStorage::LOG_COLUMN]);
    return std::unique_ptr<rocksdb::Iterator>(it);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::sendGetIterator() const {
    auto it = transaction->transaction->GetIterator(
        read_options,
        transaction->datastorage->column_handles[DataStorage::SEND_COLUMN]);
    return std::unique_ptr<rocksdb::Iterator>(it);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::sideloadGetIterator()
    const {
    auto it = transaction->transaction->GetIterator(
        read_options,
        transaction->datastorage->column_handles[DataStorage::SIDELOAD_COLUMN]);
    return std::unique_ptr<rocksdb::Iterator>(it);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::aggregatorGetIterator()
    const {
    auto it = transaction->transaction->GetIterator(
        read_options, transaction->datastorage
                          ->column_handles[DataStorage::AGGREGATOR_COLUMN]);
    return std::unique_ptr<rocksdb::Iterator>(it);
}

std::unique_ptr<rocksdb::Iterator> ReadTransaction::refCountedGetIterator()
    const {
    auto it = transaction->transaction->GetIterator(
        read_options, transaction->datastorage
                          ->column_handles[DataStorage::REFCOUNTED_COLUMN]);
    return std::unique_ptr<rocksdb::Iterator>(it);
}

ValueResult<uint256_t> ReadTransaction::defaultGetUint256(
    const rocksdb::Slice key_slice) const {
    return getUint256UsingFamilyAndKey(
        transaction->datastorage->column_handles[DataStorage::DEFAULT_COLUMN],
        key_slice);
}

ValueResult<uint256_t> ReadTransaction::stateGetUint256(
    const rocksdb::Slice key_slice) const {
    return getUint256UsingFamilyAndKey(
        transaction->datastorage->column_handles[DataStorage::STATE_COLUMN],
        key_slice);
}

ValueResult<uint256_t> ReadTransaction::checkpointGetUint256(
    const rocksdb::Slice key_slice) const {
    return getUint256UsingFamilyAndKey(
        transaction->datastorage
            ->column_handles[DataStorage::CHECKPOINT_COLUMN],
        key_slice);
}

ValueResult<uint256_t> ReadTransaction::logGetUint256(
    const rocksdb::Slice key_slice) const {
    return getUint256UsingFamilyAndKey(
        transaction->datastorage->column_handles[DataStorage::LOG_COLUMN],
        key_slice);
}

ValueResult<uint256_t> ReadTransaction::sendGetUint256(
    const rocksdb::Slice key_slice) const {
    return getUint256UsingFamilyAndKey(
        transaction->datastorage->column_handles[DataStorage::SEND_COLUMN],
        key_slice);
}

ValueResult<uint256_t> ReadTransaction::sideloadGetUint256(
    const rocksdb::Slice key_slice) const {
    return getUint256UsingFamilyAndKey(
        transaction->datastorage->column_handles[DataStorage::SIDELOAD_COLUMN],
        key_slice);
}

ValueResult<uint256_t> ReadTransaction::aggregatorGetUint256(
    const rocksdb::Slice key_slice) const {
    return getUint256UsingFamilyAndKey(
        transaction->datastorage
            ->column_handles[DataStorage::AGGREGATOR_COLUMN],
        key_slice);
}

ValueResult<uint256_t> ReadTransaction::refCountedGetUint256(
    const rocksdb::Slice key_slice) const {
    return getUint256UsingFamilyAndKey(
        transaction->datastorage
            ->column_handles[DataStorage::REFCOUNTED_COLUMN],
        key_slice);
}

ValueResult<std::vector<std::vector<unsigned char>>>
ReadTransaction::sequencerBatchItemGetVectorVector(
    const rocksdb::Slice first_key_slice,
    size_t count) const {
    return getVectorVectorUsingFamilyAndKey(
        transaction->datastorage
            ->column_handles[DataStorage::SEQUENCERBATCHITEM_COLUMN],
        first_key_slice, count);
}

ValueResult<std::vector<std::vector<unsigned char>>>
ReadTransaction::sendGetVectorVector(const rocksdb::Slice first_key_slice,
                                     size_t count) const {
    return getVectorVectorUsingFamilyAndKey(
        transaction->datastorage->column_handles[DataStorage::SEND_COLUMN],
        first_key_slice, count);
}

ValueResult<std::vector<unsigned char>>
ReadTransaction::sequencerBatchItemGetVector(
    const rocksdb::Slice first_key_slice) const {
    return getVectorUsingFamilyAndKey(
        transaction->datastorage
            ->column_handles[DataStorage::SEQUENCERBATCHITEM_COLUMN],
        first_key_slice);
}

ValueResult<std::vector<unsigned char>>
ReadTransaction::delayedMessageGetVector(
    const rocksdb::Slice first_key_slice) const {
    return getVectorUsingFamilyAndKey(
        transaction->datastorage
            ->column_handles[DataStorage::DELAYEDMESSAGE_COLUMN],
        first_key_slice);
}

ValueResult<std::vector<unsigned char>> ReadTransaction::checkpointGetVector(
    const rocksdb::Slice first_key_slice) const {
    return getVectorUsingFamilyAndKey(
        transaction->datastorage
            ->column_handles[DataStorage::CHECKPOINT_COLUMN],
        first_key_slice);
}

ValueResult<std::vector<uint256_t>> ReadTransaction::logGetUint256Vector(
    const rocksdb::Slice first_key_slice,
    size_t count) const {
    return getUint256VectorUsingFamilyAndKey(
        transaction->datastorage->column_handles[DataStorage::LOG_COLUMN],
        first_key_slice, count);
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
