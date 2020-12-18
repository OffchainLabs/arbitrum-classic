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
#include "data_storage/datastorage.hpp"

#include "nonstd/optional.hpp"
#include "rocksdb/utilities/transaction.h"

class CheckpointedMachine {
    std::unique_ptr<Machine> mach;
    std::shared_ptr<DataStorage> storage;

   public:
    CheckpointedMachine(std::unique_ptr<Machine> mach,
                        std::shared_ptr<DataStorage> storage)
        : mach(std::move(mach)), storage(std::move(storage)) {}

    Assertion run(uint64_t stepCount,
                  std::vector<Tuple> inbox_messages,
                  std::chrono::seconds wallLimit);

    Assertion runSideloaded(uint64_t stepCount,
                            std::vector<Tuple> inbox_messages,
                            std::chrono::seconds wallLimit,
                            Tuple sideload);

    Assertion runCallServer(uint64_t stepCount,
                            std::vector<Tuple> inbox_messages,
                            std::chrono::seconds wallLimit,
                            value fake_inbox_peek_value);
};

#endif /* data_storage_checkpointedmachine_hpp */
