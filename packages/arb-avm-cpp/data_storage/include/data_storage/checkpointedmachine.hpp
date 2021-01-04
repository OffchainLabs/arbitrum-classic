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

#ifndef checkpointstore_hpp
#define checkpointstore_hpp

#include <avm/machine.hpp>
#include <avm_values/bigint.hpp>
#include <data_storage/checkpoint.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresultfwd.hpp>
#include <data_storage/value/code.hpp>
#include <data_storage/value/valuecache.hpp>

#include <memory>
#include <vector>

namespace rocksdb {
class TransactionDB;
class Status;
struct Slice;
class ColumnFamilyHandle;
}  // namespace rocksdb

class CheckpointedMachine {
   private:
    std::shared_ptr<DataStorage> data_storage;
    std::unique_ptr<Machine> machine;
    std::shared_ptr<Code> code;
    Checkpoint pending_checkpoint;

   public:
    CheckpointedMachine() = delete;
    explicit CheckpointedMachine(std::shared_ptr<DataStorage> data_storage_)
        : data_storage(std::move(data_storage_)),
          code(std::make_shared<Code>(
              getNextSegmentID(*makeConstTransaction()))) {}
    void saveCheckpoint();
    void saveAssertion(const Assertion& assertion);
    rocksdb::Status deleteCheckpoint(const uint64_t& message_number);
    DbResult<Checkpoint> getCheckpoint(const uint64_t& message_number) const;

    bool isEmpty() const;
    uint64_t maxMessageNumber();
    DbResult<Checkpoint> getCheckpointAtOrBeforeMessage(
        const uint64_t& message_number);
    uint64_t reorgToMessageOrBefore(const uint64_t& message_number);
    std::unique_ptr<Transaction> makeTransaction();
    std::unique_ptr<const Transaction> makeConstTransaction() const;
    void initialize(LoadedExecutable executable);
    bool initialized() const;
    std::unique_ptr<Machine> getInitialMachine(ValueCache& value_cache);
    std::unique_ptr<Machine> getMachine(uint256_t machineHash,
                                        ValueCache& value_cache);
    std::unique_ptr<Machine> getMachineUsingStateKeys(
        Transaction& transaction,
        MachineStateKeys state_data,
        ValueCache& value_cache);
    Assertion run(uint64_t stepCount,
                  std::vector<Tuple> inbox_messages,
                  std::chrono::seconds wallLimit);
    Assertion runCallServer(uint64_t stepCount,
                            std::vector<Tuple> inbox_messages,
                            std::chrono::seconds wallLimit,
                            value fake_inbox_peek_value);
    Assertion runSideloaded(uint64_t stepCount,
                            std::vector<Tuple> inbox_messages,
                            std::chrono::seconds wallLimit,
                            Tuple sideload_value);
};

#endif /* checkpointstore_hpp */
