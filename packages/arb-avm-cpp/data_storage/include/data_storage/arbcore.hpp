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

#include <avm/machinethread.hpp>
#include <memory>
#include <thread>
#include <utility>
#include <vector>
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
        ARBCORE_EMPTY,
        ARBCORE_MESSAGES_READY,
        ARBCORE_NEED_OLDER,
        ARBCORE_SUCCESS,
        ARBCORE_ERROR
    } arbcore_status_enum;

   private:
    // Mutex should be used for all database access.
    // Core thread holds mutex only during reorg.
    // Routines accessing database for log entries will need to acquire mutex.
    // No mutex required to access Sends or Messages.
    std::mutex core_mutex;
    std::shared_ptr<DataStorage> data_storage;

    std::unique_ptr<MachineThread> machine;
    std::shared_ptr<Code> code{};
    Checkpoint pending_checkpoint;

    // Core thread communication input/output
    // Core thread will update if and only if set to ARBCORE_MESSAGES_READY
    std::atomic<arbcore_status_enum> delivering_status{ARBCORE_EMPTY};

    // Core thread communication input
    std::atomic<bool> arbcore_abort{false};
    uint256_t delivering_first_sequence_number;
    uint64_t delivering_block_height{0};
    std::vector<std::vector<unsigned char>> delivering_messages;
    uint256_t delivering_previous_inbox_hash;

    // Core thread communication output
    std::string delivering_error_string;

   public:
    ArbCore() = delete;
    explicit ArbCore(std::shared_ptr<DataStorage> data_storage_)
        : data_storage(std::move(data_storage_)) {}
    void operator()();
    void abortThread();
    void deliverMessages(
        const uint256_t& first_sequence_number,
        uint64_t block_height,
        const std::vector<std::vector<unsigned char>>& messages,
        const uint256_t& previous_inbox_hash);

    rocksdb::Status saveAssertion(Transaction& tx,
                                  uint256_t first_message_sequence_number,
                                  const Assertion& assertion);

    rocksdb::Status saveCheckpoint();
    ValueResult<Checkpoint> getCheckpoint(
        const uint256_t& message_sequence_number) const;
    bool isCheckpointsEmpty() const;
    uint256_t maxCheckpointGas();
    ValueResult<Checkpoint> getCheckpointUsingGas(Transaction& tx,
                                                  const uint256_t& total_gas,
                                                  bool after_gas);
    rocksdb::Status reorgToMessageOrBefore(
        Transaction& tx,
        const uint256_t& message_sequence_number);

    std::unique_ptr<Transaction> makeTransaction();
    std::unique_ptr<const Transaction> makeConstTransaction() const;
    void initialize(const LoadedExecutable& executable);
    bool initialized() const;

    template <class T>
    std::unique_ptr<T> getInitialMachine(ValueCache& value_cache);
    template <class T>
    std::unique_ptr<T> getMachine(uint256_t machineHash,
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
    ValueResult<ExecutionCursor*> getExecutionCursor(uint256_t totalGasUsed,
                                                     ValueCache& cache);

   private:
    nonstd::optional<rocksdb::Status> addMessages(
        uint256_t first_sequence_number,
        uint64_t block_height,
        const std::vector<std::vector<unsigned char>>& messages,
        const uint256_t& previous_inbox_hash,
        const uint256_t& final_machine_sequence_number);
    nonstd::optional<MessageEntry> getNextMessage();
    bool deleteMessage(const MessageEntry& entry);
    ValueResult<std::vector<value>> getLogsNoLock(uint256_t index,
                                                  uint256_t count,
                                                  ValueCache& valueCache);
};

nonstd::optional<rocksdb::Status> deleteLogsStartingAt(Transaction& tx,
                                                       uint256_t log_index);

#endif /* arbcore_hpp */
