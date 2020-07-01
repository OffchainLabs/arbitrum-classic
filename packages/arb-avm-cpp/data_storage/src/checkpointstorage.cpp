/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

#include <data_storage/checkpointstorage.hpp>

#include <data_storage/blockstore.hpp>
#include <data_storage/confirmednodestore.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>
#include <data_storage/value/value.hpp>

#include <avm/machine.hpp>

#include <avm_values/codepointstub.hpp>
#include <avm_values/tuple.hpp>
#include <avm_values/vmValueParser.hpp>

#include <rocksdb/options.h>
#include <rocksdb/utilities/transaction.h>
#include <rocksdb/utilities/transaction_db.h>

CheckpointStorage::CheckpointStorage(const std::string& db_path,
                                     const std::string& contract_path)
    : datastorage(std::make_shared<DataStorage>(db_path)),
      pool(std::make_shared<TuplePool>()) {
    auto ret = parseStaticVmValues(contract_path, *pool.get());
    if (!ret.second) {
        throw std::runtime_error("invalid initial values");
    }
    initial_state = std::make_shared<StaticVmValues>(std::move(ret.first));
}

bool CheckpointStorage::closeCheckpointStorage() {
    auto status = datastorage->closeDb();
    return status.ok();
}

std::unique_ptr<Transaction> CheckpointStorage::makeTransaction() {
    rocksdb::WriteOptions writeOptions;
    auto transaction = std::unique_ptr<rocksdb::Transaction>(
        datastorage->txn_db->BeginTransaction(writeOptions));
    return std::make_unique<Transaction>(datastorage, std::move(transaction));
}

std::unique_ptr<const Transaction> CheckpointStorage::makeConstTransaction()
    const {
    rocksdb::WriteOptions writeOptions;
    auto transaction = std::unique_ptr<rocksdb::Transaction>(
        datastorage->txn_db->BeginTransaction(writeOptions));
    return std::make_unique<Transaction>(datastorage, std::move(transaction));
}

std::unique_ptr<KeyValueStore> CheckpointStorage::makeKeyValueStore() {
    return std::make_unique<KeyValueStore>(datastorage);
}

std::unique_ptr<BlockStore> CheckpointStorage::getBlockStore() const {
    return std::make_unique<BlockStore>(datastorage);
}

std::unique_ptr<ConfirmedNodeStore> CheckpointStorage::getConfirmedNodeStore()
    const {
    return std::make_unique<ConfirmedNodeStore>(datastorage);
}

Machine CheckpointStorage::getInitialMachine() const {
    return {MachineState{initial_state, pool}};
}

std::pair<Machine, bool> CheckpointStorage::getMachine(
    uint256_t machineHash) const {
    auto transaction = makeConstTransaction();
    auto results = getMachineState(*transaction, machineHash);
    if (!results.status.ok()) {
        return std::make_pair(Machine{}, false);
    }

    auto state_data = results.data;

    auto register_results =
        ::getValue(*transaction, state_data.register_hash, pool.get());
    if (!register_results.status.ok()) {
        return std::make_pair(Machine{}, false);
    }

    auto stack_results =
        ::getValue(*transaction, state_data.datastack_hash, pool.get());
    if (!stack_results.status.ok() ||
        !nonstd::holds_alternative<Tuple>(stack_results.data)) {
        return std::make_pair(Machine{}, false);
    }

    auto auxstack_results =
        ::getValue(*transaction, state_data.auxstack_hash, pool.get());
    if (!auxstack_results.status.ok() ||
        !nonstd::holds_alternative<Tuple>(auxstack_results.data)) {
        return std::make_pair(Machine{}, false);
    }

    MachineState machine_state{
        pool,
        initial_state,
        std::move(register_results.data),
        Datastack(nonstd::get<Tuple>(stack_results.data)),
        Datastack(nonstd::get<Tuple>(auxstack_results.data)),
        state_data.status,
        state_data.pc,
        state_data.err_pc};
    return std::make_pair(std::move(machine_state), true);
}

DbResult<value> CheckpointStorage::getValue(uint256_t value_hash) const {
    auto tx = makeConstTransaction();
    return ::getValue(*tx, value_hash, pool.get());
}
