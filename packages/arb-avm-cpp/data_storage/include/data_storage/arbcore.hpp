/*
 * Copyright 2020-2021, Offchain Labs, Inc.
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
#include <data_storage/datacursor.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/executioncursor.hpp>
#include <data_storage/messageentry.hpp>
#include <data_storage/readsnapshottransaction.hpp>
#include <data_storage/storageresultfwd.hpp>
#include <data_storage/value/code.hpp>
#include <data_storage/value/valuecache.hpp>
#include <utility>

#include <avm/machinethread.hpp>
#include <map>
#include <memory>
#include <queue>
#include <shared_mutex>
#include <thread>
#include <utility>
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
        MESSAGES_EMPTY,       // Out: Ready to receive messages
        MESSAGES_READY,       // In:  Messages in vector
        MESSAGES_SUCCESS,     // Out:  Messages processed successfully
        MESSAGES_NEED_OLDER,  // Out: Last message invalid, need older messages
        MESSAGES_ERROR        // Out: Error processing messages
    } message_status_enum;

    struct logscursor_logs {
        uint256_t first_log_index;
        std::vector<value> logs;
        std::vector<value> deleted_logs;
    };

   private:
    struct message_data_struct {
        std::vector<std::vector<unsigned char>> messages;
        uint256_t previous_inbox_acc;
        bool last_block_complete{false};
        std::optional<uint256_t> reorg_message_count;
    };

   private:
    std::unique_ptr<std::thread> core_thread;

    // Core thread input
    std::atomic<bool> arbcore_abort{false};

    // Core thread input
    std::atomic<bool> save_checkpoint{false};
    rocksdb::Status save_checkpoint_status;

    // Core thread input for cleaning up old checkpoints
    // Delete_checkpoints_before message should be set to non-zero value
    // after other parameters are set.  Main thread will set
    // delete_checkpoints_before_message to zero once cleanup has finished.
    std::atomic<uint256_t> delete_checkpoints_before_message{0};
    // If save_checkpoint_message_interval is zero, all checkpoints after
    // delete_checkpoints_before_message are deleted.  If non-zero, only
    // the last checkpoint for each message interval is saved.
    uint256_t save_checkpoint_message_interval{0};
    // Do not delete any checkpoints after ignore_checkpoints_after_message
    uint256_t ignore_checkpoints_after_message{0};

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

    // Cache a machine ready to sideload view transactions just after recent
    // blocks
    std::shared_mutex sideload_cache_mutex;
    std::map<uint256_t, std::unique_ptr<Machine>> sideload_cache;

    // Core thread inbox status input/output. Core thread will update if and
    // only if set to MESSAGES_READY
    std::atomic<message_status_enum> message_data_status{MESSAGES_EMPTY};

    // Core thread inbox input
    message_data_struct message_data;

    // Core thread inbox output
    std::string core_error_string;

    // Core thread logs output
    std::vector<DataCursor> logs_cursors{1};

    // Core thread machine state output
    std::atomic<bool> machine_idle{false};
    std::atomic<bool> machine_error{false};
    std::string machine_error_string;

    MachineOutput last_machine_output;
    std::shared_mutex last_machine_output_mutex;

   public:
    ArbCore() = delete;
    explicit ArbCore(std::shared_ptr<DataStorage> data_storage_);

    ~ArbCore() { abortThread(); }
    rocksdb::Status initialize(const LoadedExecutable& executable);
    bool initialized() const;
    void operator()();

   public:
    // Public Thread interaction
    bool startThread();
    void abortThread();

   private:
    // Private database interaction
    rocksdb::Status saveAssertion(ReadWriteTransaction& tx,
                                  const Assertion& assertion,
                                  uint256_t arb_gas_used);
    std::variant<rocksdb::Status, MachineStateKeys> getCheckpoint(
        ReadTransaction& tx,
        const uint256_t& arb_gas_used) const;
    std::variant<rocksdb::Status, MachineStateKeys> getCheckpointUsingGas(
        ReadTransaction& tx,
        const uint256_t& total_gas,
        bool after_gas);
    rocksdb::Status reorgToMessageOrBefore(
        const uint256_t& message_sequence_number,
        bool use_latest,
        ValueCache& cache);
    template <class T>
    std::unique_ptr<T> getMachineUsingStateKeys(
        const ReadTransaction& transaction,
        const MachineStateKeys& state_data,
        ValueCache& value_cache) const;

   public:
    // To be deprecated, use checkpoints instead
    template <class T>
    std::unique_ptr<T> getMachine(uint256_t machineHash,
                                  ValueCache& value_cache);

   private:
    template <class T>
    std::unique_ptr<T> getMachineImpl(ReadTransaction& tx,
                                      uint256_t machineHash,
                                      ValueCache& value_cache);
    rocksdb::Status saveCheckpoint(ReadWriteTransaction& tx);

   public:
    // Useful for unit tests
    // Do not call triggerSaveCheckpoint from multiple threads at the same time
    rocksdb::Status triggerSaveCheckpoint();
    bool isCheckpointsEmpty(ReadTransaction& tx) const;
    uint256_t maxCheckpointGas();

   public:
    // Managing machine state
    bool machineIdle();
    std::optional<std::string> machineClearError();
    uint256_t machineMessagesRead();

   public:
    // Sending messages to core thread
    bool deliverMessages(std::vector<std::vector<unsigned char>> messages,
                         const uint256_t& previous_inbox_acc,
                         bool last_block_complete,
                         const std::optional<uint256_t>& reorg_height);
    message_status_enum messagesStatus();
    std::string messagesClearError();

   public:
    // Logs Cursor interaction
    bool logsCursorRequest(size_t cursor_index, uint256_t count);
    ValueResult<logscursor_logs> logsCursorGetLogs(size_t cursor_index);
    bool logsCursorCheckError(size_t cursor_index) const;
    std::string logsCursorClearError(size_t cursor_index);
    bool logsCursorConfirmReceived(size_t cursor_index);
    ValueResult<uint256_t> logsCursorPosition(size_t cursor_index) const;

   private:
    // Logs cursor internal functions
    void handleLogsCursorRequested(ReadTransaction& tx,
                                   size_t cursor_index,
                                   ValueCache& cache);
    rocksdb::Status handleLogsCursorReorg(size_t cursor_index,
                                          uint256_t log_count,
                                          ValueCache& cache);

   public:
    // Execution Cursor interaction
    ValueResult<std::unique_ptr<ExecutionCursor>> getExecutionCursor(
        uint256_t total_gas_used,
        ValueCache& cache);
    rocksdb::Status advanceExecutionCursor(ExecutionCursor& execution_cursor,
                                           uint256_t max_gas,
                                           bool go_over_gas,
                                           ValueCache& cache);

    std::unique_ptr<Machine> takeExecutionCursorMachine(
        ExecutionCursor& execution_cursor,
        ValueCache& cache) const;

   private:
    // Execution cursor internal functions
    rocksdb::Status advanceExecutionCursorImpl(
        ExecutionCursor& execution_cursor,
        uint256_t total_gas_used,
        bool go_over_gas,
        uint256_t message_group_size,
        ValueCache& cache);

    std::unique_ptr<Machine>& resolveExecutionCursorMachine(
        const ReadTransaction& tx,
        ExecutionCursor& execution_cursor,
        ValueCache& cache) const;
    std::unique_ptr<Machine> takeExecutionCursorMachineImpl(
        const ReadTransaction& tx,
        ExecutionCursor& execution_cursor,
        ValueCache& cache) const;

   public:
    ValueResult<uint256_t> logInsertedCount() const;
    ValueResult<uint256_t> sendInsertedCount() const;
    ValueResult<uint256_t> messageEntryInsertedCount() const;
    ValueResult<std::vector<value>> getLogs(uint256_t index,
                                            uint256_t count,
                                            ValueCache& valueCache);
    ValueResult<std::vector<std::vector<unsigned char>>> getSends(
        uint256_t index,
        uint256_t count) const;

    ValueResult<std::vector<std::vector<unsigned char>>> getMessages(
        uint256_t index,
        uint256_t count) const;
    ValueResult<uint256_t> getInboxAcc(uint256_t index);
    ValueResult<std::pair<uint256_t, uint256_t>> getInboxAccPair(
        uint256_t index1,
        uint256_t index2);

   private:
    ValueResult<std::pair<std::vector<std::vector<unsigned char>>,
                          std::optional<uint256_t>>>
    getMessagesImpl(const ReadTransaction& tx,
                    uint256_t index,
                    uint256_t count) const;
    template <typename T>
    rocksdb::Status resolveStagedMessage(const ReadTransaction& tx,
                                         T& machine_state);
    // Private database interaction
    ValueResult<MessageEntry> getMessageEntry(
        const ReadTransaction& tx,
        uint256_t message_sequence_number) const;
    ValueResult<uint256_t> logInsertedCountImpl(
        const ReadTransaction& tx) const;

    ValueResult<uint256_t> logProcessedCount(ReadTransaction& tx) const;
    rocksdb::Status updateLogProcessedCount(ReadWriteTransaction& tx,
                                            rocksdb::Slice value_slice);
    ValueResult<uint256_t> sendInsertedCountImpl(
        const ReadTransaction& tx) const;

    ValueResult<uint256_t> sendProcessedCount(ReadTransaction& tx) const;
    rocksdb::Status updateSendProcessedCount(ReadWriteTransaction& tx,
                                             rocksdb::Slice value_slice);
    ValueResult<uint256_t> messageEntryInsertedCountImpl(
        const ReadTransaction& tx) const;

    rocksdb::Status saveLogs(ReadWriteTransaction& tx,
                             const std::vector<value>& val);
    rocksdb::Status saveSends(
        ReadWriteTransaction& tx,
        const std::vector<std::vector<unsigned char>>& sends);

   private:
    std::optional<rocksdb::Status> addMessages(
        const std::vector<std::vector<unsigned char>>& new_messages,
        bool last_block_complete,
        const uint256_t& prev_inbox_acc,
        const std::optional<uint256_t>& reorg_message_count,
        ValueCache& cache);
    ValueResult<std::vector<value>> getLogsNoLock(ReadTransaction& tx,
                                                  uint256_t index,
                                                  uint256_t count,
                                                  ValueCache& valueCache);

    bool isValid(ReadTransaction& tx,
                 const InboxState& fully_processed_inbox,
                 const staged_variant& staged_message);

    ValueResult<std::pair<bool, std::vector<InboxMessage>>>
    executionCursorGetMessages(ReadTransaction& tx,
                               const ExecutionCursor& execution_cursor,
                               const uint256_t& orig_message_group_size);
    ValueResult<std::pair<bool, std::vector<InboxMessage>>>
    executionCursorGetMessagesNoLock(ReadTransaction& tx,
                                     const ExecutionCursor& execution_cursor,
                                     const uint256_t& orig_message_group_size);

    std::variant<rocksdb::Status, MachineStateKeys> getClosestExecutionMachine(
        ReadTransaction& tx,
        const uint256_t& total_gas_used,
        bool is_for_sideload = false);

    rocksdb::Status updateLogInsertedCount(ReadWriteTransaction& tx,
                                           const uint256_t& log_index);
    rocksdb::Status updateSendInsertedCount(ReadWriteTransaction& tx,
                                            const uint256_t& send_index);
    rocksdb::Status updateMessageEntryInsertedCount(
        ReadWriteTransaction& tx,
        const uint256_t& message_index);

   public:
    // Public sideload interaction
    ValueResult<std::unique_ptr<Machine>> getMachineForSideload(
        const uint256_t& block_number,
        ValueCache& cache);

    ValueResult<uint256_t> getSideloadPosition(ReadTransaction& tx,
                                               const uint256_t& block_number);

   private:
    // Private sideload interaction
    rocksdb::Status saveSideloadPosition(ReadWriteTransaction& tx,
                                         const uint256_t& block_number,
                                         const uint256_t& arb_gas_used);

    rocksdb::Status deleteSideloadsStartingAt(ReadWriteTransaction& tx,
                                              const uint256_t& block_number);
    rocksdb::Status logsCursorSaveCurrentTotalCount(ReadWriteTransaction& tx,
                                                    size_t cursor_index,
                                                    uint256_t count);
    ValueResult<uint256_t> logsCursorGetCurrentTotalCount(
        const ReadTransaction& tx,
        size_t cursor_index) const;
};

std::optional<rocksdb::Status> deleteLogsStartingAt(ReadWriteTransaction& tx,
                                                    uint256_t log_index);

#endif /* arbcore_hpp */
