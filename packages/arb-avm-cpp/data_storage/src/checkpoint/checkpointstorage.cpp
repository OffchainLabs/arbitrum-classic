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
#include <avm_values/vmValueParser.hpp>

#include <rocksdb/options.h>
#include <rocksdb/utilities/transaction.h>
#include <rocksdb/utilities/transaction_db.h>

CheckpointStorage::CheckpointStorage(const std::string& db_path,
                                     const std::string& contract_path)
    : datastorage(std::make_unique<DataStorage>(db_path)),
      pool(std::make_shared<TuplePool>()) {
    initial_state = parseInitialVmValues(contract_path, *pool.get());
}

auto CheckpointStorage::closeCheckpointStorage() -> bool {
    auto status = datastorage->closeDb();
    return status.ok();
}

auto CheckpointStorage::getInitialVmValues() const -> InitialVmValues {
    return initial_state;
}

auto CheckpointStorage::getValue(
    const std::vector<unsigned char>& hash_key) const -> GetResults {
    return datastorage->getValue(hash_key);
}

auto CheckpointStorage::makeTransaction() -> std::unique_ptr<Transaction> {
    return datastorage->makeTransaction();
}

auto CheckpointStorage::makeConstTransaction() const
    -> std::unique_ptr<const Transaction> {
    return datastorage->makeTransaction();
}

auto CheckpointStorage::makeKeyValueStore() -> std::unique_ptr<KeyValueStore> {
    return datastorage->makeKeyValueStore();
}
