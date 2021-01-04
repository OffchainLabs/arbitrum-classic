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
#include <data_storage/blockstore.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/code.hpp>

#include <avm/machine.hpp>

#include <avm_values/tuple.hpp>
#include <avm_values/vmValueParser.hpp>
#include <utility>

ArbStorage::ArbStorage(const std::string& db_path)
    : datastorage(std::make_shared<DataStorage>(db_path)),
      cmach(std::make_shared<CheckpointedMachine>(datastorage)) {}

void ArbStorage::initialize(const std::string& executable_path) {
    auto executable = loadExecutable(executable_path);
    initialize(std::move(executable));
}

void ArbStorage::initialize(LoadedExecutable executable) {
    cmach->initialize(std::move(executable));
}

bool ArbStorage::initialized() const {
    return cmach->initialized();
}

bool ArbStorage::closeArbStorage() {
    auto status = datastorage->closeDb();
    return status.ok();
}

std::unique_ptr<KeyValueStore> ArbStorage::makeKeyValueStore() {
    return std::make_unique<KeyValueStore>(datastorage);
}

std::unique_ptr<BlockStore> ArbStorage::getBlockStore() const {
    return std::make_unique<BlockStore>(datastorage);
}

std::unique_ptr<AggregatorStore> ArbStorage::getAggregatorStore() const {
    return std::make_unique<AggregatorStore>(datastorage);
}

std::unique_ptr<MessageStore> ArbStorage::getMessageStore() const {
    return std::make_unique<MessageStore>(datastorage);
}

std::shared_ptr<CheckpointedMachine> ArbStorage::getCheckpointedMachine() {
    return cmach;
}

std::unique_ptr<Machine> ArbStorage::getInitialMachine(
    ValueCache& value_cache) const {
    return cmach->getInitialMachine(value_cache);
}

std::unique_ptr<Machine> ArbStorage::getMachine(uint256_t machineHash,
                                                ValueCache& value_cache) const {
    return cmach->getMachine(machineHash, value_cache);
}

DbResult<value> ArbStorage::getValue(uint256_t value_hash,
                                     ValueCache& value_cache) const {
    auto tx = cmach->makeConstTransaction();
    return ::getValue(*tx, value_hash, value_cache);
}

std::unique_ptr<Transaction> ArbStorage::makeTransaction() {
    return cmach->makeTransaction();
}

std::unique_ptr<const Transaction> ArbStorage::makeConstTransaction() const {
    return cmach->makeConstTransaction();
}
