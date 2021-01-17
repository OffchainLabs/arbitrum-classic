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
#include <data_storage/messageentry.hpp>
#include <data_storage/messagestore.hpp>
#include <data_storage/storageresultfwd.hpp>
#include <data_storage/value/code.hpp>
#include <data_storage/value/valuecache.hpp>

#include <memory>
#include <thread>
#include <vector>

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
    std::unique_ptr<MessageStore> message_store;
    std::shared_ptr<Code> code;
    Checkpoint pending_checkpoint;

    // Thread communication
    std::atomic<message_status_enum> message_status;
    uint256_t first_sequence_number;
    uint64_t block_height;
    std::vector<std::vector<unsigned char>> messages;
    std::vector<uint256_t> inbox_hashes;
    uint256_t previous_inbox_hash;

   public:
    ArbCore() = delete;
    explicit ArbCore(std::shared_ptr<DataStorage> data_storage_)
        : data_storage(std::move(data_storage_)),
          message_store(std::make_unique<MessageStore>(data_storage)),
          code(std::make_shared<Code>(
              getNextSegmentID(*makeConstTransaction()))),
          message_status(MESSAGES_EMPTY),
          block_height(0) {}
    void operator()();
    void stopThreads();
    void deliverMessages(
        const uint256_t& first_sequence_number,
        uint64_t block_height,
        const std::vector<std::vector<unsigned char>>& messages,
        const std::vector<uint256_t>& inbox_hashes,
        const uint256_t& previous_inbox_hash);

    void saveAssertion(uint256_t first_message_sequence_number,
                       const Assertion& assertion);

    void saveCheckpoint();
    ValueResult<Checkpoint> getCheckpoint(
        const uint256_t& message_sequence_number) const;
    bool isCheckpointsEmpty() const;
    uint256_t maxMessageSequenceNumber();
    DbResult<Checkpoint> getCheckpointAtOrBeforeMessage(
        const uint256_t& message_sequence_number);
    uint256_t reorgToMessageOrBefore(const uint256_t& message_sequence_number);

    std::unique_ptr<Transaction> makeTransaction();
    std::unique_ptr<const Transaction> makeConstTransaction() const;
    void initialize(const LoadedExecutable& executable);
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
                  const std::vector<std::vector<unsigned char>>& inbox_messages,
                  const nonstd::optional<uint256_t>& final_block);

    ValueResult<uint256_t> lastLogInserted(rocksdb::Transaction& transaction);
    ValueResult<uint256_t> lastLogProcessed(rocksdb::Transaction& transaction);
    ValueResult<uint256_t> lastSendInserted(rocksdb::Transaction& transaction);
    ValueResult<uint256_t> lastSendProcessed(rocksdb::Transaction& transaction);
    ValueResult<uint256_t> lastMessageEntryInserted(
        rocksdb::Transaction& transaction);
    ValueResult<uint256_t> lastMessageEntryProcessed(
        rocksdb::Transaction& transaction);
    rocksdb::Status updateLastLogInserted(rocksdb::Transaction& transaction,
                                          rocksdb::Slice value_slice);
    rocksdb::Status updateLastLogProcessed(rocksdb::Transaction& transaction,
                                           rocksdb::Slice value_slice);
    rocksdb::Status updateLastSendInserted(rocksdb::Transaction& transaction,
                                           rocksdb::Slice value_slice);
    rocksdb::Status updateLastSendProcessed(rocksdb::Transaction& transaction,
                                            rocksdb::Slice value_slice);
    rocksdb::Status updateLastMessageEntryInserted(
        rocksdb::Transaction& transaction,
        rocksdb::Slice value_slice);
    rocksdb::Status updateLastMessageEntryProcessed(
        rocksdb::Transaction& transaction,
        rocksdb::Slice value_slice);
    rocksdb::Status saveLog(Transaction& tx, const value& val);
    DbResult<value> getLog(uint256_t index, ValueCache& valueCache) const;
    ValueResult<std::vector<unsigned char>> getSend(uint256_t index) const;
    rocksdb::Status saveSend(Transaction& tx,
                             const std::vector<unsigned char>& send);
    rocksdb::Status saveMessage(Transaction& tx,
                                const std::vector<unsigned char>& message);
    ValueResult<std::vector<unsigned char>> getMessage(uint256_t index) const;
};

#endif /* arbcore_hpp */
