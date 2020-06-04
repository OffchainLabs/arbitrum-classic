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

#include <data_storage/blockstore.hpp>
#include <data_storage/checkpoint/checkpointstorage.hpp>
#include <data_storage/nodestore.hpp>
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

bool CheckpointStorage::closeCheckpointStorage() {
    auto status = datastorage->closeDb();
    return status.ok();
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

std::unique_ptr<KeyValueStore> CheckpointStorage::makeKeyValueStore() {
    return datastorage->makeKeyValueStore();
}

std::unique_ptr<BlockStore> CheckpointStorage::getBlockStore() const {
    return datastorage->getBlockStore();
}

std::unique_ptr<NodeStore> CheckpointStorage::getNodeStore() const {
    return datastorage->getNodeStore();
}
