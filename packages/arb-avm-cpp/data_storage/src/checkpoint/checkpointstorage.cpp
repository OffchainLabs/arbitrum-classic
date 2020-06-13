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
#include <data_storage/checkpoint/checkpointutils.hpp>
#include <data_storage/checkpoint/machine.hpp>
#include <data_storage/checkpoint/value.hpp>
#include <data_storage/confirmednodestore.hpp>
#include <data_storage/storageresult.hpp>

#include <avm/machine.hpp>

#include <avm_values/codepoint.hpp>
#include <avm_values/tuple.hpp>
#include <avm_values/vmValueParser.hpp>

#include <rocksdb/options.h>
#include <rocksdb/utilities/transaction.h>
#include <rocksdb/utilities/transaction_db.h>

CheckpointStorage::CheckpointStorage(const std::string& db_path,
                                     const std::string& contract_path)
    : datastorage(std::make_shared<DataStorage>(db_path)),
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
    return std::make_unique<KeyValueStore>(datastorage);
}

std::unique_ptr<BlockStore> CheckpointStorage::getBlockStore() const {
    return std::make_unique<BlockStore>(datastorage);
}

std::unique_ptr<ConfirmedNodeStore> CheckpointStorage::getConfirmedNodeStore()
    const {
    return std::make_unique<ConfirmedNodeStore>(datastorage);
}

std::pair<Machine, bool> CheckpointStorage::getMachine(
    uint256_t machineHash) const {
    auto transaction = makeConstTransaction();
    auto results = getMachineState(*transaction, machineHash);

    auto initial_values = getInitialVmValues();
    if (!initial_values.valid_state) {
        return std::make_pair(Machine{}, false);
    }

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
        initial_values.code,
        initial_values.staticVal,
        std::move(register_results.data),
        Datastack(nonstd::get<Tuple>(stack_results.data)),
        Datastack(nonstd::get<Tuple>(auxstack_results.data)),
        state_data.status,
        CodePointRef(state_data.pc),
        CodePointRef(state_data.err_pc)};
    return std::make_pair(std::move(machine_state), true);
}
