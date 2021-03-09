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
    : ReadOnlyTransaction(std::move(store)) {}

std::unique_ptr<ReadWriteTransaction>
ReadWriteTransaction::makeReadWriteTransaction(
    std::shared_ptr<DataStorage> store) {
    return std::make_unique<ReadWriteTransaction>(std::move(store));
}

rocksdb::Status ReadWriteTransaction::defaultPut(const rocksdb::Slice& key,
                                                 const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage->default_column.get(), key, value);
}

rocksdb::Status ReadWriteTransaction::statePut(const rocksdb::Slice& key,
                                               const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage->state_column.get(), key, value);
}

rocksdb::Status ReadWriteTransaction::checkpointPut(
    const rocksdb::Slice& key,
    const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage->checkpoint_column.get(), key, value);
}

rocksdb::Status ReadWriteTransaction::messageEntryPut(
    const rocksdb::Slice& key,
    const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage->messageentry_column.get(), key, value);
}

rocksdb::Status ReadWriteTransaction::logPut(const rocksdb::Slice& key,
                                             const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage->log_column.get(), key, value);
}

rocksdb::Status ReadWriteTransaction::sendPut(const rocksdb::Slice& key,
                                              const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage->send_column.get(), key, value);
}

rocksdb::Status ReadWriteTransaction::sideloadPut(const rocksdb::Slice& key,
                                                  const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage->sideload_column.get(), key, value);
}

rocksdb::Status ReadWriteTransaction::aggregatorPut(
    const rocksdb::Slice& key,
    const rocksdb::Slice& value) {
    return transaction->transaction->Put(
        transaction->datastorage->aggregator_column.get(), key, value);
}

rocksdb::Status ReadWriteTransaction::defaultDelete(const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage->default_column.get(), key);
}

rocksdb::Status ReadWriteTransaction::stateDelete(const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage->state_column.get(), key);
}
rocksdb::Status ReadWriteTransaction::checkpointDelete(
    const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage->checkpoint_column.get(), key);
}
rocksdb::Status ReadWriteTransaction::messageEntryDelete(
    const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage->messageentry_column.get(), key);
}
rocksdb::Status ReadWriteTransaction::logDelete(const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage->log_column.get(), key);
}
rocksdb::Status ReadWriteTransaction::sendDelete(const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage->send_column.get(), key);
}
rocksdb::Status ReadWriteTransaction::sideloadDelete(
    const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage->sideload_column.get(), key);
}
rocksdb::Status ReadWriteTransaction::aggregatorDelete(
    const rocksdb::Slice& key) {
    return transaction->transaction->Delete(
        transaction->datastorage->aggregator_column.get(), key);
}
