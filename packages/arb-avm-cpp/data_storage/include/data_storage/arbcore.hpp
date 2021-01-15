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

#ifndef arbcore_hpp
#define arbcore_hpp

#include <avm/machine.hpp>
#include <avm_values/bigint.hpp>
#include <data_storage/checkpoint.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresultfwd.hpp>
#include <data_storage/value/code.hpp>
#include <data_storage/value/valuecache.hpp>

#include <memory>
#include <vector>
#include "messageentry.hpp"

namespace rocksdb {
class TransactionDB;
class Status;
struct Slice;
class ColumnFamilyHandle;
}  // namespace rocksdb

class ArbCore {
   public:
    typedef enum {
        MESSAGES_EMPTY,
        MESSAGES_READY,
        MESSAGES_NEED_OLDER,
        MESSAGES_SUCCESS
    } message_status_enum;

    typedef enum {
        MACHINE_NONE,
        MACHINE_REQUEST_STOP,
        MACHINE_FINISHED
    } machine_status_enum;

   private:
    std::shared_ptr<DataStorage> data_storage;
    std::unique_ptr<Machine> machine;
    std::shared_ptr<Code> code;
    Checkpoint pending_checkpoint;
    std::unique_ptr<std::thread> core_thread;

    // Thread communication
    std::atomic<message_status_enum> message_status;
    std::vector<MessageEntry> messages;

   public:
    ArbCore() = delete;
    explicit ArbCore(std::shared_ptr<DataStorage> data_storage_)
        : data_storage(std::move(data_storage_)),
          code(std::make_shared<Code>(
              getNextSegmentID(*makeConstTransaction()))),
          message_status(MESSAGES_EMPTY) {}
    void operator()();
    void stopThreads();
    void saveCheckpoint();
    void saveAssertion(uint256_t first_message_sequence_number,
                       const Assertion& assertion);
    DbResult<Checkpoint> getCheckpoint(
        const uint256_t& message_sequence_number) const;

    bool isEmpty() const;
    uint256_t maxMessageSequenceNumber();
    DbResult<Checkpoint> getCheckpointAtOrBeforeMessage(
        const uint256_t& message_sequence_number);
    uint256_t reorgToMessageOrBefore(const uint256_t& message_sequence_number);
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
    Assertion run(uint64_t gas_limit,
                  bool hard_gas_limit,
                  uint256_t first_message_sequence_number,
                  const std::vector<rocksdb::Slice>& inbox_messages,
                  nonstd::optional<uint256_t> final_block);

    nonstd::optional<rocksdb::Status> addMessages(
        const uint256_t first_sequence_number,
        const uint64_t block_height,
        const std::vector<rocksdb::Slice>& messages,
        const std::vector<uint256_t>& inbox_hashes,
        const uint256_t& previous_inbox_hash);
    nonstd::optional<MessageEntry> getNextMessage();
    nonstd::optional<MessageEntry> getLastMessage();
    bool deleteMessage(const MessageEntry& entry);
};

nonstd::optional<rocksdb::Status> deleteMessagesStartingAt(
    Transaction& tx,
    uint256_t sequence_number);

rocksdb::Status addMessagesWithoutCheck(
    Transaction& tx,
    const uint256_t first_sequence_number,
    const uint64_t block_height,
    const std::vector<rocksdb::Slice>& messages,
    const std::vector<uint256_t>& inbox_hashes);

#endif /* arbcore_hpp */
