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

#ifndef checkpointdeleter_hpp
#define checkpointdeleter_hpp

#include <vector>

#include <avm_values/tuple.hpp>
#include <avm_values/value.hpp>

struct DeleteResults;
class CheckpointStorage;
struct GetResults;
class Transaction;

namespace rocksdb {
class Status;
}

auto deleteCheckpoint(CheckpointStorage& checkpoint_storage,
                      const std::vector<unsigned char>& checkpoint_name)
    -> DeleteResults;

class MachineStateDeleter {
   private:
    std::unique_ptr<Transaction> transaction;
    auto deleteTuple(const std::vector<unsigned char>& hash_key,
                     const GetResults& results) -> DeleteResults;

   public:
    MachineStateDeleter(std::unique_ptr<Transaction> transaction_);
    auto deleteTuple(const std::vector<unsigned char>& hash_key)
        -> DeleteResults;
    auto deleteValue(const std::vector<unsigned char>& hash_key)
        -> DeleteResults;
    auto commitTransaction() -> rocksdb::Status;
    auto rollBackTransaction() -> rocksdb::Status;
};

#endif /* checkpointdeleter_h */
