/*
 * Copyright 2019, Offchain Labs, Inc.
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
#include <data_storage/checkpoint/checkpointstorage.hpp>
#include <data_storage/storageresult.hpp>

#include <avm_values/codepoint.hpp>
#include <avm_values/tuple.hpp>

#include <rocksdb/options.h>
#include <rocksdb/utilities/transaction.h>
#include <rocksdb/utilities/transaction_db.h>

CheckpointStorage::CheckpointStorage(const std::string& db_path,
                                     const InitialVmValues& initial_state_)
    : initial_state(initial_state_) {
    auto db = new DataStorage(db_path);
    datastorage = std::unique_ptr<DataStorage>(db);
}

InitialVmValues CheckpointStorage::getInitialVmValues() const {
    return initial_state;
}

GetResults CheckpointStorage::getValue(
    const std::vector<unsigned char>& hash_key) const {
    return datastorage->getValue(hash_key);
}

std::unique_ptr<Transaction> CheckpointStorage::makeTransaction() {
    return datastorage->makeTransaction();
}

std::unique_ptr<const Transaction> CheckpointStorage::makeConstTransaction()
    const {
    return datastorage->makeTransaction();
}
