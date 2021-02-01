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
#include <data_storage/storageresultfwd.hpp>
#include <data_storage/value/code.hpp>
#include <data_storage/value/valuecache.hpp>
#include <utility>

#include <avm/machinethread.hpp>
#include <memory>
#include <queue>
#include <thread>
#include <utility>
#include <vector>
#include "datacursor.hpp"
#include "executioncursor.hpp"

namespace rocksdb {
class TransactionDB;
class Status;
struct Slice;
class ColumnFamilyHandle;
}  // namespace rocksdb

class ArbCore {
   public:
    typedef enum {
        MESSAGES_EMPTY,       // Out: Ready to receive messages
        MESSAGES_READY,       // In:  Messages in vector
        MESSAGES_SUCCESS,     // In:  Messages processed successfully
        MESSAGES_NEED_OLDER,  // Out: Last message invalid, need older messages
        MESSAGES_ERROR        // Out: Error processing messages
    } messages_status_enum;

   private:
    std::unique_ptr<std::thread> core_thread;

    // Core thread holds mutex only during reorg.
    // Routines accessing database for log entries will need to acquire mutex
    // because obsolete log entries have `Value` references removed causing
    // reference counts to be decremented and possibly deleted.
    // No mutex required to access Sends or Messages because obsolete entries
    // are not deleted.
    std::mutex core_reorg_mutex;
    std::shared_ptr<DataStorage> data_storage;

    std::unique_ptr<MachineThread> machine;
    std::shared_ptr<Code> code{};
    Checkpoint pending_checkpoint;

    // Core thread inbox input/output. Core thread will update if and only if
    // set to MESSAGES_READY or MESSAGES_SUCCESS
    std::atomic<messages_status_enum> delivering_inbox_status{MESSAGES_EMPTY};

    // Core thread inbox input
    std::atomic<bool> arbcore_abort{false};
    std::vector<std::vector<unsigned char>> delivering_inbox_messages;
    uint256_t delivering_previous_inbox_hash;

    // Core thread inbox output
    std::string delivering_inbox_error_string;

    // Core thread logs output
    DataCursor logs_cursor;

   public:
    ArbCore() = delete;
    explicit ArbCore(std::shared_ptr<DataStorage> data_storage_)
        : data_storage(std::move(data_storage_)),
          code(std::make_shared<Code>(
              getNextSegmentID(*makeConstTransaction()))) {}
    void operator()();
    bool startThread();
    void abortThread();
    void deliverMessages(
        const std::vector<std::vector<unsigned char>>& messages,
        const uint256_t& previous_inbox_hash);

    rocksdb::Status saveAssertion(Transaction& tx, const Assertion& assertion);

    rocksdb::Status saveCheckpoint(Transaction& tx);
    ValueResult<Checkpoint> getCheckpoint(
        Transaction& tx,
        const uint256_t& message_sequence_number) const;
    bool isCheckpointsEmpty() const;
    uint256_t maxCheckpointGas();
    ValueResult<Checkpoint> getCheckpointUsingGas(Transaction& tx,
                                                  const uint256_t& total_gas,
                                                  bool after_gas);
    rocksdb::Status reorgToMessageOrBefore(
        Transaction& tx,
        const uint256_t& message_sequence_number,
        ValueCache& cache);

    std::unique_ptr<Transaction> makeTransaction();
    std::unique_ptr<const Transaction> makeConstTransaction() const;
    void initialize(const LoadedExecutable& executable);
    bool initialized() const;

    uint256_t getInitialMachineHash(Transaction& tx);

    template <class T>
    std::unique_ptr<T> getInitialMachine(ValueCache& value_cache);

    template <class T>
    std::unique_ptr<T> getInitialMachineImpl(Transaction& tx,
                                             ValueCache& value_cache);

    template <class T>
    std::unique_ptr<T> getMachine(uint256_t machineHash,
                                  ValueCache& value_cache);
    template <class T>
    std::unique_ptr<T> getMachineImpl(Transaction& tx,
                                      uint256_t machineHash,
                                      ValueCache& value_cache);
    template <class T>
    std::unique_ptr<T> getMachineUsingStateKeys(Transaction& transaction,
                                                MachineStateKeys state_data,
                                                ValueCache& value_cache);

    ValueResult<uint256_t> logInsertedCount(Transaction& tx) const;
    rocksdb::Status updateLogInsertedCount(Transaction& tx,
                                           rocksdb::Slice value_slice);
    ValueResult<uint256_t> logProcessedCount(Transaction& tx) const;
    rocksdb::Status updateLogProcessedCount(Transaction& tx,
                                            rocksdb::Slice value_slice);
    ValueResult<uint256_t> sendInsertedCount(Transaction& tx) const;
    rocksdb::Status updateSendInsertedCount(Transaction& tx,
                                            rocksdb::Slice value_slice);
    ValueResult<uint256_t> sendProcessedCount(Transaction& tx) const;
    rocksdb::Status updateSendProcessedCount(Transaction& tx,
                                             rocksdb::Slice value_slice);
    ValueResult<uint256_t> messageEntryInsertedCount(Transaction& tx) const;
    rocksdb::Status updateMessageEntryInsertedCount(Transaction& tx,
                                                    rocksdb::Slice value_slice);
    ValueResult<uint256_t> messageEntryProcessedCount(Transaction& tx) const;
    rocksdb::Status updateMessageEntryProcessedCount(
        Transaction& tx,
        rocksdb::Slice value_slice);

    rocksdb::Status saveLogs(Transaction& tx, const std::vector<value>& val);
    ValueResult<std::vector<value>> getLogs(uint256_t index,
                                            uint256_t count,
                                            ValueCache& valueCache);
    ValueResult<std::vector<std::vector<unsigned char>>> getSends(
        uint256_t index,
        uint256_t count) const;
    ValueResult<std::vector<uint256_t>> getInboxHashes(uint256_t index,
                                                       uint256_t count) const;
    ValueResult<std::vector<std::vector<unsigned char>>> getMessages(
        uint256_t index,
        uint256_t count) const;
    rocksdb::Status saveSends(
        Transaction& tx,
        const std::vector<std::vector<unsigned char>>& send);
    bool messagesEmpty();
    ValueResult<uint256_t> getInboxDelta(uint256_t start_index,
                                         uint256_t count);
    ValueResult<uint256_t> getInboxAcc(uint256_t index);
    ValueResult<uint256_t> getSendAcc(uint256_t start_acc_hash,
                                      uint256_t start_index,
                                      uint256_t count);
    ValueResult<uint256_t> getLogAcc(uint256_t start_acc_hash,
                                     uint256_t start_index,
                                     uint256_t count,
                                     ValueCache& cache);
    ValueResult<std::unique_ptr<ExecutionCursor>> getExecutionCursor(
        uint256_t total_gas_used,
        ValueCache& cache);
    rocksdb::Status Advance(ExecutionCursor& execution_cursor,
                            uint256_t max_gas,
                            bool go_over_gas,
                            ValueCache& cache);
    rocksdb::Status getExecutionCursorImpl(Transaction& tx,
                                           ExecutionCursor& execution_cursor,
                                           uint256_t total_gas_used,
                                           bool go_over_gas,
                                           uint256_t message_group_size,
                                           ValueCache& cache);

    bool logsCursorRequest(uint256_t count);
    nonstd::optional<std::vector<value>> logsCursorGetLogs();
    bool logsCursorSetNextIndex(uint256_t next_index);
    bool logsCursorCheckError() const;
    std::string logsCursorClearError();

   private:
    nonstd::optional<rocksdb::Status> addMessages(
        const std::vector<std::vector<unsigned char>>& messages,
        const uint256_t& previous_inbox_hash,
        const uint256_t& final_machine_sequence_number,
        ValueCache& cache);
    nonstd::optional<MessageEntry> getNextMessage();
    bool deleteMessage(const MessageEntry& entry);
    ValueResult<std::vector<value>> getLogsNoLock(Transaction& tx,
                                                  uint256_t index,
                                                  uint256_t count,
                                                  ValueCache& valueCache);

    void handleLogsCursorRequested(Transaction& tx, ValueCache& cache);
    void handleLogsCursorProcessed(Transaction& tx);
    rocksdb::Status handleLogsCursorReorg(Transaction& tx,
                                          uint256_t log_count,
                                          ValueCache& cache);
    ValueResult<bool> executionCursorAddMessages(
        Transaction& tx,
        ExecutionCursor& execution_cursor,
        const uint256_t& orig_message_group_size);
    rocksdb::Status executionCursorSetup(Transaction& tx,
                                         ExecutionCursor& execution_cursor,
                                         const uint256_t& total_gas_used,
                                         ValueCache& cache);
};

nonstd::optional<rocksdb::Status> deleteLogsStartingAt(Transaction& tx,
                                                       uint256_t log_index);

#endif /* arbcore_hpp */
