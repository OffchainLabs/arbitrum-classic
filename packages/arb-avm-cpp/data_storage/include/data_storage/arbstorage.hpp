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

#ifndef arbstorage_hpp
#define arbstorage_hpp

#include "avm_values/vmValueParser.hpp"
#include "data_storage/arbcore.hpp"
#include "data_storage/datastorage.hpp"
#include "data_storage/value/value.hpp"

#include <memory>
#include <string>
#include <vector>

struct GetResults;
class Machine;
class BlockStore;
class AggregatorStore;

namespace rocksdb {
class TransactionDB;
}

class ArbStorage {
    std::shared_ptr<DataStorage> datastorage;
    std::shared_ptr<ArbCore> cmach;

   public:
    explicit ArbStorage(const std::string& db_path);
    bool closeArbStorage();
    void initialize(LoadedExecutable executable);
    void initialize(const std::string& executable_path);
    bool initialized() const;

    std::unique_ptr<Transaction> makeTransaction();
    std::unique_ptr<const Transaction> makeConstTransaction() const;
    std::unique_ptr<KeyValueStore> makeKeyValueStore();
    std::unique_ptr<BlockStore> getBlockStore() const;
    std::unique_ptr<AggregatorStore> getAggregatorStore() const;
    std::unique_ptr<MessageStore> getMessageStore() const;
    std::shared_ptr<ArbCore> getArbCore();

    std::unique_ptr<Machine> getInitialMachine(ValueCache& value_cache) const;
    std::unique_ptr<Machine> getMachine(uint256_t machineHash,
                                        ValueCache& value_cache) const;
    DbResult<value> getValue(uint256_t value_hash,
                             ValueCache& value_cache) const;
};

#endif /* arbstorage_hpp */
