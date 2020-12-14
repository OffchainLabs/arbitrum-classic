/*
 * Copyright 2020, Offchain Labs, Inc.
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

#ifndef data_storage_checkpointedmachine_hpp
#define data_storage_checkpointedmachine_hpp

#include "avm/machine.hpp"
#include "avm_values/bigint.hpp"
#include "data_storage/aggregator.hpp"
#include "data_storage/checkpointstorage.hpp"

#include "nonstd/optional.hpp"
#include "rocksdb/utilities/transaction.h"

class DataStorage;

class CheckpointedMachine {
    Machine mach;
    CheckpointStorage db;

   public:
    CheckpointedMachine(MachineState machine_state_, const std::string& db_path)
        : mach(std::move(machine_state_)), db(db_path) {}

    CheckpointedMachine(std::shared_ptr<Code> code,
                        value static_val,
                        const std::string& db_path)
        : mach(std::move(code), std::move(static_val)), db(db_path) {}

    static CheckpointedMachine loadFromFile(
        const std::string& executable_filename,
        const std::string& db_path) {
        return {MachineState::loadFromFile(executable_filename), db_path};
    }

    void initialize(LoadedExecutable executable);
    void initialize(const std::string& executable_path);
    bool initialized() const;
    bool closeCheckpointedMachine();

    std::unique_ptr<AggregatorStore> getAggregatorStore() const;

    Assertion runSideloaded(uint64_t stepCount,
                            std::vector<Tuple> inbox_messages,
                            std::chrono::seconds wallLimit,
                            Tuple sideload);

    Assertion run(uint64_t stepCount,
                  std::vector<Tuple> inbox_messages,
                  std::chrono::seconds wallLimit);

    Assertion runCallServer(uint64_t stepCount,
                            std::vector<Tuple> inbox_messages,
                            std::chrono::seconds wallLimit,
                            value fake_inbox_peek_value);
};

#endif /* data_storage_checkpointedmachine_hpp */
