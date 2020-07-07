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
#include <data_storage/value/code.hpp>
#include <data_storage/value/machine.hpp>
#include <data_storage/value/value.hpp>

#include <avm/machine.hpp>

#include <avm_values/codepointstub.hpp>
#include <avm_values/tuple.hpp>
#include <avm_values/vmValueParser.hpp>

#include <rocksdb/options.h>
#include <rocksdb/utilities/transaction.h>
#include <rocksdb/utilities/transaction_db.h>

CheckpointStorage::CheckpointStorage(std::shared_ptr<DataStorage> datastorage_,
                                     LoadedExecutable exec,
                                     std::shared_ptr<TuplePool> pool_)
    : datastorage(std::move(datastorage_)),
      code(std::make_shared<Code>(std::move(exec.code))),
      static_val(std::move(exec.static_val)),
      pool(std::move(pool_)) {}

CheckpointStorage::CheckpointStorage(std::shared_ptr<DataStorage> datastorage_,
                                     const std::string& contract_path,
                                     std::shared_ptr<TuplePool> pool)
    : CheckpointStorage(std::move(datastorage_),
                        loadExecutable(contract_path, *pool),
                        std::move(pool)) {}

CheckpointStorage::CheckpointStorage(const std::string& db_path,
                                     const std::string& contract_path)
    : CheckpointStorage(std::make_shared<DataStorage>(db_path),
                        contract_path,
                        std::make_shared<TuplePool>()) {}

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
    return {MachineState{code, static_val, pool}};
}

std::pair<Machine, bool> CheckpointStorage::getMachine(
    uint256_t machineHash) const {
    std::set<uint64_t> segment_ids;
    auto transaction = makeConstTransaction();
    auto results = getMachineState(*transaction, machineHash);
    if (!results.status.ok()) {
        return std::make_pair(Machine{}, false);
    }

    auto state_data = results.data;

    auto register_results = ::getValueImpl(
        *transaction, state_data.register_hash, pool.get(), segment_ids);
    if (!register_results.status.ok()) {
        return std::make_pair(Machine{}, false);
    }

    auto stack_results = ::getValueImpl(*transaction, state_data.datastack_hash,
                                        pool.get(), segment_ids);
    if (!stack_results.status.ok() ||
        !nonstd::holds_alternative<Tuple>(stack_results.data)) {
        return std::make_pair(Machine{}, false);
    }

    auto auxstack_results = ::getValueImpl(
        *transaction, state_data.auxstack_hash, pool.get(), segment_ids);
    if (!auxstack_results.status.ok() ||
        !nonstd::holds_alternative<Tuple>(auxstack_results.data)) {
        return std::make_pair(Machine{}, false);
    }

    segment_ids.insert(state_data.pc.segment);
    segment_ids.insert(state_data.err_pc.pc.segment);

    // Later segments can reference earlier ones, so load the segments backwards
    for (auto it = segment_ids.rbegin(); it != segment_ids.rend(); ++it) {
        if (code->containsSegment(*it)) {
            // If the segment is already loaded, no need to restore it
            continue;
        }
        auto segment =
            getCodeSegment(*transaction, *it, pool.get(), segment_ids);
        if (!segment) {
            return std::make_pair(Machine{}, false);
        }
        code->restoreExistingSegment(std::move(segment));
    }

    MachineState machine_state{
        pool,
        code,
        std::move(register_results.data),
        static_val,
        Datastack(nonstd::get<Tuple>(stack_results.data)),
        Datastack(nonstd::get<Tuple>(auxstack_results.data)),
        state_data.arb_gas_remaining,
        state_data.status,
        state_data.pc,
        state_data.err_pc};
    return std::make_pair(std::move(machine_state), true);
}

DbResult<value> CheckpointStorage::getValue(uint256_t value_hash) const {
    auto tx = makeConstTransaction();
    return ::getValue(*tx, value_hash, pool.get());
}
