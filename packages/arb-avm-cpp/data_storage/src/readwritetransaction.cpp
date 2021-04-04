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

#include "data_storage/readwritetransaction.hpp"

#include <utility>

ReadWriteTransaction::ReadWriteTransaction(std::shared_ptr<DataStorage> store)
    : ReadConsistentTransaction(std::move(store)) {}

rocksdb::Status ReadWriteTransaction::defaultPut(const rocksdb::Slice& key,
                                                 const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage->column_handles[DataStorage::DEFAULT_COLUMN],
        key, value);
}

rocksdb::Status ReadWriteTransaction::statePut(const rocksdb::Slice& key,
                                               const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage->column_handles[DataStorage::STATE_COLUMN],
        key, value);
}

rocksdb::Status ReadWriteTransaction::checkpointPut(
    const rocksdb::Slice& key,
    const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage
            ->column_handles[DataStorage::CHECKPOINT_COLUMN],
        key, value);
}

rocksdb::Status ReadWriteTransaction::messageEntryPut(
    const rocksdb::Slice& key,
    const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage
            ->column_handles[DataStorage::MESSAGEENTRY_COLUMN],
        key, value);
}

rocksdb::Status ReadWriteTransaction::logPut(const rocksdb::Slice& key,
                                             const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage->column_handles[DataStorage::LOG_COLUMN], key,
        value);
}

rocksdb::Status ReadWriteTransaction::sendPut(const rocksdb::Slice& key,
                                              const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage->column_handles[DataStorage::SEND_COLUMN], key,
        value);
}

rocksdb::Status ReadWriteTransaction::sideloadPut(const rocksdb::Slice& key,
                                                  const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage->column_handles[DataStorage::SIDELOAD_COLUMN],
        key, value);
}

rocksdb::Status ReadWriteTransaction::aggregatorPut(
    const rocksdb::Slice& key,
    const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage
            ->column_handles[DataStorage::AGGREGATOR_COLUMN],
        key, value);
}

rocksdb::Status ReadWriteTransaction::refCountedPut(
    const rocksdb::Slice& key,
    const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage
            ->column_handles[DataStorage::REFCOUNTED_COLUMN],
        key, value);
}

rocksdb::Status ReadWriteTransaction::defaultDelete(const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage->column_handles[DataStorage::DEFAULT_COLUMN],
        key);
}

rocksdb::Status ReadWriteTransaction::stateDelete(const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage->column_handles[DataStorage::STATE_COLUMN],
        key);
}
rocksdb::Status ReadWriteTransaction::checkpointDelete(
    const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage
            ->column_handles[DataStorage::CHECKPOINT_COLUMN],
        key);
}
rocksdb::Status ReadWriteTransaction::messageEntryDelete(
    const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage
            ->column_handles[DataStorage::MESSAGEENTRY_COLUMN],
        key);
}
rocksdb::Status ReadWriteTransaction::logDelete(const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage->column_handles[DataStorage::LOG_COLUMN], key);
}
rocksdb::Status ReadWriteTransaction::sendDelete(const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage->column_handles[DataStorage::SEND_COLUMN],
        key);
}
rocksdb::Status ReadWriteTransaction::sideloadDelete(
    const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage->column_handles[DataStorage::SIDELOAD_COLUMN],
        key);
}
rocksdb::Status ReadWriteTransaction::aggregatorDelete(
    const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage
            ->column_handles[DataStorage::AGGREGATOR_COLUMN],
        key);
}
rocksdb::Status ReadWriteTransaction::refCountedDelete(
    const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage
            ->column_handles[DataStorage::REFCOUNTED_COLUMN],
        key);
}
