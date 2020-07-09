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

#include "bigint_utils.hpp"

#include <rocksdb/options.h>
#include <rocksdb/utilities/transaction.h>
#include <rocksdb/utilities/transaction_db.h>

namespace {
const char* initial_slice_label = "initial";
}

CheckpointStorage::CheckpointStorage(const std::string& db_path)
    : datastorage(std::make_shared<DataStorage>(db_path)),
      code(std::make_shared<Code>(getNextSegmentID(*makeConstTransaction()))),
      pool(std::make_shared<TuplePool>()) {}

void CheckpointStorage::initialize(const std::string& executable_path) {
    auto executable = loadExecutable(executable_path, *pool);
    auto tx = makeTransaction();
    code->addSegment(std::move(executable.code));
    Machine mach{MachineState{code, std::move(executable.static_val), pool}};
    auto res = saveMachine(*tx, mach);
    auto save_exp = std::runtime_error("failed to save");
    if (!res.status.ok()) {
        throw save_exp;
    }
    std::vector<unsigned char> value_data;
    marshal_uint256_t(mach.hash(), value_data);
    rocksdb::Slice value_slice{reinterpret_cast<const char*>(value_data.data()),
                               value_data.size()};
    auto s =
        tx->transaction->Put(rocksdb::Slice(initial_slice_label), value_slice);
    if (!s.ok()) {
        throw save_exp;
    }
    s = tx->commit();
    if (!s.ok()) {
        throw save_exp;
    }
}

bool CheckpointStorage::initialized() const {
    auto tx = makeConstTransaction();
    std::string initial_raw;
    auto s =
        tx->transaction->Get(rocksdb::ReadOptions(),
                             rocksdb::Slice(initial_slice_label), &initial_raw);
    return s.ok();
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
    auto tx = makeConstTransaction();
    std::string initial_raw;
    auto s =
        tx->transaction->Get(rocksdb::ReadOptions(),
                             rocksdb::Slice(initial_slice_label), &initial_raw);
    if (!s.ok()) {
        throw std::runtime_error("failed to load initial val");
    }
    auto machine_hash =
        from_big_endian(initial_raw.data(), initial_raw.data() + UINT256_SIZE);
    return getMachine(machine_hash);
}

Machine CheckpointStorage::getMachine(uint256_t machineHash) const {
    std::set<uint64_t> segment_ids;
    auto transaction = makeConstTransaction();
    auto results = getMachineState(*transaction, machineHash);
    if (!results.status.ok()) {
        throw std::runtime_error("failed to load machine state");
    }

    auto state_data = results.data;

    auto static_results = ::getValueImpl(*transaction, state_data.static_hash,
                                         pool.get(), segment_ids);
    if (!static_results.status.ok()) {
        throw std::runtime_error("failed loaded core machine static");
    }

    auto register_results = ::getValueImpl(
        *transaction, state_data.register_hash, pool.get(), segment_ids);
    if (!register_results.status.ok()) {
        throw std::runtime_error("failed to load machine register");
    }

    auto stack_results = ::getValueImpl(*transaction, state_data.datastack_hash,
                                        pool.get(), segment_ids);
    if (!stack_results.status.ok() ||
        !nonstd::holds_alternative<Tuple>(stack_results.data)) {
        throw std::runtime_error("failed to load machine stack");
    }

    auto auxstack_results = ::getValueImpl(
        *transaction, state_data.auxstack_hash, pool.get(), segment_ids);
    if (!auxstack_results.status.ok() ||
        !nonstd::holds_alternative<Tuple>(auxstack_results.data)) {
        throw std::runtime_error("failed to load machine auxstack");
    }

    segment_ids.insert(state_data.pc.segment);
    segment_ids.insert(state_data.err_pc.pc.segment);

    bool loaded_segment = true;
    while (loaded_segment) {
        loaded_segment = false;
        std::set<uint64_t> next_segment_ids;
        for (auto it = segment_ids.rbegin(); it != segment_ids.rend(); ++it) {
            if (code->containsSegment(*it)) {
                // If the segment is already loaded, no need to restore it
                continue;
            }
            auto segment =
                getCodeSegment(*transaction, *it, pool.get(), next_segment_ids);
            code->restoreExistingSegment(std::move(segment));
            loaded_segment = true;
        }
        segment_ids = std::move(next_segment_ids);
    }

    return MachineState{pool,
                        code,
                        std::move(register_results.data),
                        std::move(static_results.data),
                        Datastack(nonstd::get<Tuple>(stack_results.data)),
                        Datastack(nonstd::get<Tuple>(auxstack_results.data)),
                        state_data.arb_gas_remaining,
                        state_data.status,
                        state_data.pc,
                        state_data.err_pc};
}

DbResult<value> CheckpointStorage::getValue(uint256_t value_hash) const {
    auto tx = makeConstTransaction();
    return ::getValue(*tx, value_hash, pool.get());
}
