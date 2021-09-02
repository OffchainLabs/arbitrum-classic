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

#include <data_storage/arbstorage.hpp>

#include <data_storage/aggregator.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/code.hpp>

#include <avm/machine.hpp>

#include <avm_values/tuple.hpp>
#include <avm_values/vmValueParser.hpp>
#include <utility>

ArbStorage::ArbStorage(const std::string& db_path,
                       const ArbCoreConfig& coreConfig)
    : datastorage(std::make_shared<DataStorage>(db_path)),
      arb_core(std::make_shared<ArbCore>(datastorage, coreConfig)) {}

rocksdb::Status ArbStorage::initialize(const std::string& executable_path) {
    auto executable = loadExecutable(executable_path);
    return initialize(executable);
}

rocksdb::Status ArbStorage::initialize(const LoadedExecutable& executable) {
    return arb_core->initialize(executable);
}

bool ArbStorage::initialized() const {
    return arb_core->initialized();
}

bool ArbStorage::closeArbStorage() {
    arb_core->abortThread();
    auto status = datastorage->closeDb();
    if (!status.ok()) {
        std::cerr << "error closing database: " << status.ToString()
                  << std::endl;
    }
    return status.ok();
}

std::unique_ptr<AggregatorStore> ArbStorage::getAggregatorStore() const {
    return std::make_unique<AggregatorStore>(datastorage);
}

std::shared_ptr<ArbCore> ArbStorage::getArbCore() {
    return arb_core;
}

std::unique_ptr<Machine> ArbStorage::getInitialMachine() {
    auto cursor = arb_core->getExecutionCursor(0);
    if (!cursor.status.ok()) {
        throw std::runtime_error(
            "failed to get initial machine. Database not initialized of "
            "corrupted");
    }
    return arb_core->takeExecutionCursorMachine(*cursor.data);
}

std::unique_ptr<Machine> ArbStorage::getMachine(uint256_t machineHash,
                                                ValueCache& value_cache) const {
    return arb_core->getMachine<Machine>(machineHash, value_cache);
}

DbResult<value> ArbStorage::getValue(uint256_t value_hash,
                                     ValueCache& value_cache) const {
    ReadTransaction tx(datastorage);
    return ::getValue(tx, value_hash, value_cache, false);
}

std::unique_ptr<ReadTransaction> ArbStorage::makeReadTransaction() {
    return std::make_unique<ReadTransaction>(datastorage);
}

std::unique_ptr<ReadWriteTransaction> ArbStorage::makeReadWriteTransaction() {
    return std::make_unique<ReadWriteTransaction>(datastorage);
}
