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

#include <data_storage/arbcore.hpp>

#include "value/utils.hpp"

#include <avm/machinethread.hpp>
#include <data_storage/aggregator.hpp>
#include <data_storage/checkpoint.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/inboxmessage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>
#include <data_storage/value/value.hpp>
#include <data_storage/value/valuecache.hpp>

#include <ethash/keccak.hpp>
#include <set>
#include <vector>

namespace {
constexpr auto initial_machine_hash_key = std::array<char, 1>{-10};
constexpr auto log_inserted_key = std::array<char, 1>{-60};
constexpr auto log_processed_key = std::array<char, 1>{-61};
constexpr auto send_inserted_key = std::array<char, 1>{-62};
constexpr auto send_processed_key = std::array<char, 1>{-63};
constexpr auto message_entry_inserted_key = std::array<char, 1>{-64};
constexpr auto message_entry_processed_key = std::array<char, 1>{-65};

ValueResult<MessageEntry> getMessageEntry(Transaction& tx,
                                          uint256_t message_sequence_number) {
    std::vector<unsigned char> previous_key;
    marshal_uint256_t(message_sequence_number, previous_key);

    auto result = getVectorUsingFamilyAndKey(
        *tx.transaction, tx.datastorage->messageentry_column.get(),
        vecToSlice(previous_key));
    if (!result.status.ok()) {
        return {result.status, {}};
    }

    auto parsed_state =
        extractMessageEntry(message_sequence_number, vecToSlice(result.data));

    return {result.status, parsed_state};
}

}  // namespace

bool ArbCore::messagesEmpty() {
    return delivering_status == ARBCORE_EMPTY;
}

// deliverMessages sends messages to core thread.  Caller needs to verify that
// messagesEmpty() returns true before calling this function.
void ArbCore::deliverMessages(
    const uint256_t& first_sequence_number,
    const uint64_t block_height,
    const std::vector<std::vector<unsigned char>>& messages,
    const uint256_t& previous_inbox_hash) {
    if (delivering_status != ARBCORE_EMPTY) {
        throw std::runtime_error("message_status != ARBCORE_EMPTY");
    }

    delivering_first_sequence_number = first_sequence_number;
    delivering_block_height = block_height;
    delivering_messages = messages;
    delivering_previous_inbox_hash = previous_inbox_hash;

    delivering_status = ARBCORE_MESSAGES_READY;
}

std::unique_ptr<Transaction> ArbCore::makeTransaction() {
    return Transaction::makeTransaction(data_storage);
}

std::unique_ptr<const Transaction> ArbCore::makeConstTransaction() const {
    auto transaction =
        std::unique_ptr<rocksdb::Transaction>(data_storage->beginTransaction());
    return std::make_unique<Transaction>(data_storage, std::move(transaction));
}

void ArbCore::initialize(const LoadedExecutable& executable) {
    auto tx = makeTransaction();
    code->addSegment(executable.code);
    machine = std::make_unique<MachineThread>(
        MachineState{code, executable.static_val});
    auto res = saveMachine(*tx, *machine);
    if (!res.status.ok()) {
        throw std::runtime_error("failed to save initial machine");
    }
    std::vector<unsigned char> value_data;
    marshal_uint256_t(machine->hash(), value_data);
    auto s = tx->transaction->Put(vecToSlice(initial_machine_hash_key),
                                  vecToSlice(value_data));
    if (!s.ok()) {
        throw std::runtime_error(
            "failed to save initial machine values into db");
    }

    s = tx->commit();
    if (!s.ok()) {
        throw std::runtime_error("failed to commit initial machine into db");
    }
}

bool ArbCore::initialized() const {
    auto tx = makeConstTransaction();
    std::string initial_raw;
    auto s = tx->transaction->GetForUpdate(
        rocksdb::ReadOptions(), data_storage->state_column.get(),
        vecToSlice(initial_machine_hash_key), &initial_raw);
    return s.ok();
}

template <class T>
std::unique_ptr<T> ArbCore::getInitialMachine(ValueCache& value_cache) {
    auto tx = makeConstTransaction();
    std::string initial_raw;
    auto s = tx->transaction->GetForUpdate(
        rocksdb::ReadOptions(), data_storage->state_column.get(),
        vecToSlice(initial_machine_hash_key), &initial_raw);
    if (!s.ok()) {
        throw std::runtime_error("failed to load initial val");
    }

    auto machine_hash = intx::be::unsafe::load<uint256_t>(
        reinterpret_cast<const unsigned char*>(initial_raw.data()));
    return getMachine<T>(machine_hash, value_cache);
}

template std::unique_ptr<Machine> ArbCore::getInitialMachine(ValueCache&);
template std::unique_ptr<MachineThread> ArbCore::getInitialMachine(ValueCache&);

template <class T>
std::unique_ptr<T> ArbCore::getMachine(uint256_t machineHash,
                                       ValueCache& value_cache) {
    auto transaction = makeTransaction();
    auto results = getMachineStateKeys(*transaction, machineHash);
    if (!results.status.ok()) {
        throw std::runtime_error("failed to load machine state");
    }

    return getMachineUsingStateKeys<T>(*transaction, results.data, value_cache);
}

template std::unique_ptr<Machine> ArbCore::getMachine(uint256_t, ValueCache&);
template std::unique_ptr<MachineThread> ArbCore::getMachine(uint256_t,
                                                            ValueCache&);

rocksdb::Status ArbCore::saveCheckpoint() {
    auto tx = Transaction::makeTransaction(data_storage);

    auto status =
        saveMachineState(*tx, *machine, pending_checkpoint.machine_state_keys);
    if (!status.ok()) {
        return status;
    }

    // Pull inbox hash from database
    auto existing_message_entry =
        getMessageEntry(*tx, pending_checkpoint.total_messages_read - 1);
    pending_checkpoint.inbox_hash = existing_message_entry.data.inbox_hash;

    std::vector<unsigned char> key;
    marshal_uint256_t(pending_checkpoint.arb_gas_used, key);
    auto key_slice = vecToSlice(key);
    auto serialized_checkpoint = serializeCheckpoint(pending_checkpoint);
    std::string value_str(serialized_checkpoint.begin(),
                          serialized_checkpoint.end());
    auto put_status = tx->transaction->Put(
        tx->datastorage->checkpoint_column.get(), key_slice, value_str);
    if (!put_status.ok()) {
        return put_status;
    }

    auto commit_status = tx->commit();
    if (!commit_status.ok()) {
        return commit_status;
    }

    return rocksdb::Status::OK();
}

rocksdb::Status ArbCore::saveAssertion(Transaction& tx,
                                       uint256_t first_message_sequence_number,
                                       const Assertion& assertion) {
    auto status = saveLogs(tx, assertion.logs);
    if (!status.ok()) {
        return status;
    }

    status = saveSends(tx, assertion.sends);
    if (!status.ok()) {
        return status;
    }

    pending_checkpoint.applyAssertion(first_message_sequence_number, assertion);

    std::vector<unsigned char> processed_key;
    marshal_uint256_t(pending_checkpoint.total_messages_read, processed_key);
    updateMessageEntryProcessedCount(tx, vecToSlice(processed_key));

    return rocksdb::Status::OK();
}

// reorgToMessageOrBefore resets the checkpoint and database entries
// such that machine state is at or before the requested message. cleaning
// up old references as needed.
rocksdb::Status ArbCore::reorgToMessageOrBefore(
    Transaction& tx,
    const uint256_t& message_sequence_number) {
    // Delete each checkpoint until at or below message_sequence_number
    auto it = std::unique_ptr<rocksdb::Iterator>(tx.transaction->GetIterator(
        rocksdb::ReadOptions(), tx.datastorage->checkpoint_column.get()));

    // Find first message to delete
    it->SeekToLast();
    if (!it->status().ok()) {
        return it->status();
    }

    bool good_checkpoint_found = false;
    Checkpoint last_deleted_checkpoint;
    while (it->Valid()) {
        auto keyBuf = it->key().data();
        auto gas_key = deserializeUint256t(keyBuf);

        std::vector<unsigned char> checkpoint_vector(
            it->value().data(), it->value().data() + it->value().size());
        auto checkpoint = extractCheckpoint(gas_key, checkpoint_vector);

        if (checkpoint.total_messages_read - 1 > message_sequence_number) {
            // Obsolete checkpoint, need to delete referenced machine
            last_deleted_checkpoint = checkpoint;
            deleteMachineState(tx, checkpoint.machine_state_keys);

            // Delete checkpoint to make sure it isn't used later
            tx.transaction->Delete(tx.datastorage->log_column.get(), it->key());
        } else {
            // Good checkpoint
            pending_checkpoint = checkpoint;
            good_checkpoint_found = true;
            break;
        }

        it->Next();
    }
    it = nullptr;

    if (!good_checkpoint_found) {
        // Nothing found, start database from scratch.
        // Machine will be reset when it is confirmed to not be running.

        // Clearing checkpoint will ensure everything is reset properly below
        pending_checkpoint = Checkpoint{};
    }

    // Delete logs individually to handle reference counts
    auto optional_status =
        deleteLogsStartingAt(tx, pending_checkpoint.log_count);
    if (optional_status && !optional_status->ok()) {
        return *optional_status;
    }
    std::vector<unsigned char> key;
    marshal_uint256_t(pending_checkpoint.log_count, key);
    auto status = updateLogInsertedCount(tx, vecToSlice(key));
    if (!status.ok()) {
        return status;
    }
    marshal_uint256_t(pending_checkpoint.send_count, key);
    status = updateSendInsertedCount(tx, vecToSlice(key));
    if (!status.ok()) {
        return status;
    }

    marshal_uint256_t(pending_checkpoint.total_messages_read, key);
    status = updateMessageEntryProcessedCount(tx, vecToSlice(key));
    if (!status.ok()) {
        return status;
    }

    return rocksdb::Status::OK();
}

ValueResult<Checkpoint> ArbCore::getCheckpoint(
    const uint256_t& arb_gas_used) const {
    auto tx = Transaction::makeTransaction(data_storage);

    std::vector<unsigned char> key;
    marshal_uint256_t(arb_gas_used, key);

    auto result = getVectorUsingFamilyAndKey(
        *tx->transaction, tx->datastorage->checkpoint_column.get(),
        vecToSlice(key));
    if (!result.status.ok()) {
        return {result.status, {}};
    }

    return {rocksdb::Status::OK(),
            extractCheckpoint(arb_gas_used, result.data)};
}

bool ArbCore::isCheckpointsEmpty() const {
    auto tx = Transaction::makeTransaction(data_storage);
    auto it = std::unique_ptr<rocksdb::Iterator>(tx->transaction->GetIterator(
        rocksdb::ReadOptions(), tx->datastorage->checkpoint_column.get()));
    it->SeekToLast();
    return !it->Valid();
}

uint256_t ArbCore::maxCheckpointGas() {
    auto tx = Transaction::makeTransaction(data_storage);
    auto it = std::unique_ptr<rocksdb::Iterator>(tx->transaction->GetIterator(
        rocksdb::ReadOptions(), tx->datastorage->checkpoint_column.get()));
    it->SeekToLast();
    if (it->Valid()) {
        auto keyBuf = it->key().data();
        return deserializeUint256t(keyBuf);
    } else {
        return 0;
    }
}

// getCheckpointUsingGase returns the checkpoint at or before the specified gas
// if `after_gas` is false. If `after_gas` is true, checkpoint after specified
// gas is returned.
ValueResult<Checkpoint> ArbCore::getCheckpointUsingGas(
    Transaction& tx,
    const uint256_t& total_gas,
    bool after_gas) {
    auto it = std::unique_ptr<rocksdb::Iterator>(tx.transaction->GetIterator(
        rocksdb::ReadOptions(), tx.datastorage->checkpoint_column.get()));
    std::vector<unsigned char> key;
    marshal_uint256_t(total_gas, key);
    auto key_slice = vecToSlice(key);
    it->SeekForPrev(key_slice);
    if (!it->Valid()) {
        if (!it->status().ok()) {
            return {it->status(), {}};
        }
        return {rocksdb::Status::NotFound(), {}};
    }
    if (after_gas) {
        it->Next();
        if (!it->status().ok()) {
            return {it->status(), {}};
        }
        if (!it->Valid()) {
            return {rocksdb::Status::NotFound(), {}};
        }
    }
    if (!it->status().ok()) {
        return {it->status(), {}};
    }

    std::vector<unsigned char> saved_value(
        it->value().data(), it->value().data() + it->value().size());
    auto parsed_state = extractCheckpoint(total_gas, saved_value);
    return {rocksdb::Status::OK(), parsed_state};
}

template <class T>
std::unique_ptr<T> ArbCore::getMachineUsingStateKeys(
    Transaction& transaction,
    MachineStateKeys state_data,
    ValueCache& value_cache) {
    std::set<uint64_t> segment_ids;

    auto static_results = ::getValueImpl(transaction, state_data.static_hash,
                                         segment_ids, value_cache);
    if (!static_results.status.ok()) {
        throw std::runtime_error("failed loaded core machine static");
    }

    auto register_results = ::getValueImpl(
        transaction, state_data.register_hash, segment_ids, value_cache);
    if (!register_results.status.ok()) {
        throw std::runtime_error("failed to load machine register");
    }

    auto stack_results = ::getValueImpl(transaction, state_data.datastack_hash,
                                        segment_ids, value_cache);
    if (!stack_results.status.ok() ||
        !nonstd::holds_alternative<Tuple>(stack_results.data)) {
        throw std::runtime_error("failed to load machine stack");
    }

    auto auxstack_results = ::getValueImpl(
        transaction, state_data.auxstack_hash, segment_ids, value_cache);
    if (!auxstack_results.status.ok() ||
        !nonstd::holds_alternative<Tuple>(auxstack_results.data)) {
        throw std::runtime_error("failed to load machine auxstack");
    }

    auto staged_message_results = ::getValueImpl(
        transaction, state_data.staged_message_hash, segment_ids, value_cache);
    if (!staged_message_results.status.ok()) {
        throw std::runtime_error("failed to load machine saved message");
    }

    segment_ids.insert(state_data.pc.segment);
    segment_ids.insert(state_data.err_pc.pc.segment);

    bool loaded_segment = true;
    while (loaded_segment) {
        loaded_segment = false;
        std::set<uint64_t> next_segment_ids;
        for (auto it = segment_ids.rbegin(); it != segment_ids.rend(); ++it) {
            if (code->containsSegment(*it)) {
                // If the segment is already loaded, no need to restore it
                continue;
            }
            auto segment =
                getCodeSegment(transaction, *it, next_segment_ids, value_cache);
            code->restoreExistingSegment(std::move(segment));
            loaded_segment = true;
        }
        segment_ids = std::move(next_segment_ids);
    }

    auto state =
        MachineState{code,
                     std::move(register_results.data),
                     std::move(static_results.data),
                     Datastack(nonstd::get<Tuple>(stack_results.data)),
                     Datastack(nonstd::get<Tuple>(auxstack_results.data)),
                     state_data.arb_gas_remaining,
                     state_data.status,
                     state_data.pc,
                     state_data.err_pc,
                     std::move(staged_message_results.data.get<Tuple>())};

    return std::make_unique<T>(state);
}

template std::unique_ptr<Machine>
ArbCore::getMachineUsingStateKeys(Transaction&, MachineStateKeys, ValueCache&);
template std::unique_ptr<MachineThread>
ArbCore::getMachineUsingStateKeys(Transaction&, MachineStateKeys, ValueCache&);

void ArbCore::operator()() {
    ValueCache cache;
    std::unique_ptr<std::thread> machine_thread;
    uint256_t first_sequence_number_in_machine;
    uint256_t last_sequence_number_in_machine;

    delivering_error_string.clear();

    while (!arbcore_abort) {
        if (delivering_status == ARBCORE_MESSAGES_READY) {
            // Add messages
            auto add_status = addMessages(
                delivering_first_sequence_number, delivering_block_height,
                delivering_messages, delivering_previous_inbox_hash,
                last_sequence_number_in_machine);
            if (!add_status) {
                // Messages from previous block invalid because of reorg so
                // request older messages
                delivering_status = ARBCORE_NEED_OLDER;
            } else if (!add_status->ok()) {
                delivering_error_string = add_status->ToString();
                delivering_status = ARBCORE_ERROR;
                break;
            } else {
                delivering_status = ARBCORE_SUCCESS;
            }
        }

        // Check machine thread
        auto machine_status = machine->status();
        if (machine_status == MachineThread::MACHINE_ERROR) {
            machine_thread->join();
            machine_thread = nullptr;

            delivering_error_string = machine->get_error_string();
            break;
        }

        if (machine_status == MachineThread::MACHINE_ABORTED) {
            // Reload machine from checkpoint
            machine_thread->join();
            machine_thread = nullptr;

            auto tx = Transaction::makeTransaction(data_storage);

            machine = getMachineUsingStateKeys<MachineThread>(
                *tx, pending_checkpoint.machine_state_keys, cache);
            machine_status = MachineThread::MACHINE_NONE;
        } else if (machine_status == MachineThread::MACHINE_FINISHED) {
            machine_thread->join();
            machine_thread = nullptr;

            auto tx = Transaction::makeTransaction(data_storage);

            auto last_assertion = machine->getAssertion();
            auto messages_processed = messageEntryProcessedCount(*tx);
            if (!messages_processed.status.ok()) {
                delivering_error_string = messages_processed.status.ToString();
                break;
            }
            auto last_sequence_number_consumed =
                first_sequence_number_in_machine +
                last_assertion.inbox_messages_consumed;

            if (last_sequence_number_consumed >= messages_processed.data) {
                // Machine consumed obsolete message, restore from checkpoint
                machine = getMachineUsingStateKeys<MachineThread>(
                    *tx, pending_checkpoint.machine_state_keys, cache);
                machine_status = MachineThread::MACHINE_NONE;
            } else {
                // Save logs and sends
                auto status = saveAssertion(
                    *tx, first_sequence_number_in_machine, last_assertion);
                if (!status.ok()) {
                    delivering_error_string = status.ToString();
                    break;
                }

                // Maybe save checkpoint
                // TODO Decide how often to create checkpoint
                status = saveCheckpoint();

                // TODO Decide how often to clear ValueCache
                // (only when machine thread stopped)
                cache.clear();

                machine->setStatus(MachineThread::MACHINE_NONE);
            }

            auto status = tx->commit();
            if (!status.ok()) {
                delivering_error_string = status.ToString();
                delivering_status = ARBCORE_ERROR;
                break;
            }
        }

        if (machine->status() == MachineThread::MACHINE_NONE) {
            if (machine_thread) {
                machine_thread->join();
                machine_thread = nullptr;
            }
            // Start execution of machine if new message available
            auto tx = Transaction::makeTransaction(data_storage);
            auto messages_count = messageEntryInsertedCount(*tx);
            if (!messages_count.status.ok()) {
                delivering_error_string = messages_count.status.ToString();
                delivering_status = ARBCORE_ERROR;
                break;
            }

            if (messages_count.data < pending_checkpoint.total_messages_read) {
                // Should never happen, means reorg wasn't done properly
                delivering_error_string =
                    "messages_inserted < pending_checkpoint";
                delivering_status = ARBCORE_ERROR;
                break;
            }

            first_sequence_number_in_machine =
                pending_checkpoint.total_messages_read;
            if (messages_count.data > first_sequence_number_in_machine) {
                // New messages to process
                last_sequence_number_in_machine =
                    first_sequence_number_in_machine;
                auto next_message_result =
                    getMessageEntry(*tx, first_sequence_number_in_machine);
                if (!next_message_result.status.ok()) {
                    delivering_error_string =
                        next_message_result.status.ToString();
                    delivering_status = ARBCORE_ERROR;
                    break;
                }
                if (next_message_result.data.sequence_number !=
                    first_sequence_number_in_machine) {
                    delivering_error_string =
                        "sequence number in message different than expected";
                    delivering_status = ARBCORE_ERROR;
                    break;
                }
                std::vector<std::vector<unsigned char>> messages;
                messages.push_back(next_message_result.data.data);

                nonstd::optional<uint256_t> min_next_block_height;
                if (next_message_result.data.last_message_in_block) {
                    min_next_block_height =
                        next_message_result.data.block_height + 1;
                }

                if (!machine->setRunning()) {
                    // Machine is already running, should never happen
                    delivering_error_string = "Machine already running";
                    break;
                }
                machine_thread = std::make_unique<std::thread>(
                    (std::reference_wrapper<MachineThread>(*machine)), 0, false,
                    std::move(messages), std::move(min_next_block_height));
            }
        }
    }

    // Error occurred, make sure machine stops cleanly
    machine->abortThread();
    machine_thread->join();
}

void ArbCore::abortThread() {
    arbcore_abort = true;
}

rocksdb::Status ArbCore::saveLogs(Transaction& tx,
                                  const std::vector<value>& vals) {
    auto log_result = logInsertedCount(tx);
    if (!log_result.status.ok()) {
        return log_result.status;
    }

    auto log_index = log_result.data;
    for (const auto& val : vals) {
        auto value_result = saveValue(tx, val);
        if (!value_result.status.ok()) {
            return value_result.status;
        }

        std::vector<unsigned char> key;
        marshal_uint256_t(log_index, key);
        auto key_slice = vecToSlice(key);

        std::vector<unsigned char> value_hash;
        marshal_uint256_t(hash_value(val), value_hash);
        rocksdb::Slice value_hash_slice(
            reinterpret_cast<const char*>(value_hash.data()),
            value_hash.size());

        auto status = tx.transaction->Put(data_storage->log_column.get(),
                                          key_slice, value_hash_slice);
        if (!status.ok()) {
            return status;
        }
        log_index += 1;
    }

    std::vector<unsigned char> key;
    marshal_uint256_t(log_index, key);
    return updateLogInsertedCount(tx, vecToSlice(key));
}

ValueResult<std::vector<value>> ArbCore::getLogs(uint256_t index,
                                                 uint256_t count,
                                                 ValueCache& valueCache) {
    auto tx = Transaction::makeTransaction(data_storage);

    // Acquire mutex to make sure no reorg happening
    std::lock_guard<std::mutex> lock(core_mutex);

    return getLogsNoLock(index, count, valueCache);
}

ValueResult<std::vector<value>> ArbCore::getLogsNoLock(uint256_t index,
                                                       uint256_t count,
                                                       ValueCache& valueCache) {
    auto tx = Transaction::makeTransaction(data_storage);

    // Acquire mutex to make sure no reorg happening
    std::lock_guard<std::mutex> lock(core_mutex);

    // Check if attempting to get entries past current valid logs
    auto log_count = logInsertedCount(*tx);
    if (!log_count.status.ok()) {
        return {log_count.status, {}};
    }
    auto max_log_index = log_count.data - 1;
    if (index > max_log_index) {
        return {rocksdb::Status::OK(), {}};
    }
    if (index + count > max_log_index) {
        count = max_log_index - index;
    }

    std::vector<unsigned char> key;
    marshal_uint256_t(index, key);

    auto hash_result = getUint256VectorUsingFamilyAndKey(
        *tx->transaction, data_storage->log_column.get(), vecToSlice(key),
        intx::narrow_cast<size_t>(count));
    if (!hash_result.status.ok()) {
        return {hash_result.status, {}};
    }

    std::vector<value> logs;
    for (const auto& hash : hash_result.data) {
        auto val_result = getValue(*tx, hash, valueCache);
        if (!val_result.status.ok()) {
            return {val_result.status, {}};
        }
        logs.push_back(std::move(val_result.data));
    }

    return {rocksdb::Status::OK(), logs};
}

rocksdb::Status ArbCore::saveSends(
    Transaction& tx,
    const std::vector<std::vector<unsigned char>>& sends) {
    auto send_result = sendInsertedCount(tx);
    if (!send_result.status.ok()) {
        return send_result.status;
    }

    auto send_count = send_result.data;
    for (const auto& send : sends) {
        std::vector<unsigned char> key;
        marshal_uint256_t(send_count, key);
        auto key_slice = vecToSlice(key);

        auto status = tx.transaction->Put(tx.datastorage->send_column.get(),
                                          key_slice, vecToSlice(send));
        if (!status.ok()) {
            return status;
        }
        send_count += 1;
    }

    std::vector<unsigned char> key;
    marshal_uint256_t(send_count, key);
    return updateSendInsertedCount(tx, vecToSlice(key));
}

ValueResult<std::vector<uint256_t>> ArbCore::getInboxHashes(
    uint256_t index,
    uint256_t count) const {
    auto tx = Transaction::makeTransaction(data_storage);

    // Check if attempting to get entries past current valid logs
    auto message_count_result = messageEntryInsertedCount(*tx);
    if (!message_count_result.status.ok()) {
        return {message_count_result.status, {}};
    }
    auto max_message_index = message_count_result.data - 1;
    if (index > max_message_index) {
        return {rocksdb::Status::OK(), {}};
    }
    if (index + count > max_message_index) {
        count = max_message_index - index;
    }

    std::vector<unsigned char> key;
    marshal_uint256_t(index, key);
    auto key_slice = vecToSlice(key);

    auto results = getVectorVectorUsingFamilyAndKey(
        *tx->transaction, data_storage->send_column.get(), key_slice,
        intx::narrow_cast<size_t>(count));
    if (!results.status.ok()) {
        return {results.status, {}};
    }

    std::vector<uint256_t> messages;
    messages.reserve(results.data.size());
    for (const auto& data : results.data) {
        auto message_entry = extractMessageEntry(0, vecToSlice(data));

        messages.push_back(message_entry.inbox_hash);
    }

    return {rocksdb::Status::OK(), messages};
}

ValueResult<std::vector<std::vector<unsigned char>>> ArbCore::getMessages(
    uint256_t index,
    uint256_t count) const {
    auto tx = Transaction::makeTransaction(data_storage);

    // Check if attempting to get entries past current valid logs
    auto message_count = messageEntryInsertedCount(*tx);
    if (!message_count.status.ok()) {
        return {message_count.status, {}};
    }
    auto max_message_index = message_count.data - 1;
    if (index > max_message_index) {
        return {rocksdb::Status::OK(), {}};
    }
    if (index + count > max_message_index) {
        count = max_message_index - index;
    }

    std::vector<unsigned char> key;
    marshal_uint256_t(index, key);
    auto key_slice = vecToSlice(key);

    auto results = getVectorVectorUsingFamilyAndKey(
        *tx->transaction, data_storage->send_column.get(), key_slice,
        intx::narrow_cast<size_t>(count));
    if (!results.status.ok()) {
        return {results.status, {}};
    }

    std::vector<std::vector<unsigned char>> messages;
    messages.reserve(results.data.size());
    for (const auto& data : results.data) {
        auto message_entry = extractMessageEntry(0, vecToSlice(data));

        messages.push_back(message_entry.data);
    }

    return {rocksdb::Status::OK(), messages};
}

ValueResult<std::vector<std::vector<unsigned char>>> ArbCore::getSends(
    uint256_t index,
    uint256_t count) const {
    auto tx = Transaction::makeTransaction(data_storage);

    // Check if attempting to get entries past current valid logs
    auto send_count = sendInsertedCount(*tx);
    if (!send_count.status.ok()) {
        return {send_count.status, {}};
    }
    auto max_message_index = send_count.data - 1;
    if (index > max_message_index) {
        return {rocksdb::Status::OK(), {}};
    }
    if (index + count > max_message_index) {
        count = max_message_index - index;
    }

    std::vector<unsigned char> key;
    marshal_uint256_t(index, key);
    auto key_slice = vecToSlice(key);

    return getVectorVectorUsingFamilyAndKey(
        *tx->transaction, data_storage->send_column.get(), key_slice,
        intx::narrow_cast<size_t>(count));
}

ValueResult<uint256_t> ArbCore::getInboxDelta(uint256_t start_index,
                                              uint256_t count) {
    auto hashes_result = getInboxHashes(start_index, count);
    if (!hashes_result.status.ok()) {
        return {hashes_result.status, 0};
    }

    uint256_t combined_hash = 0;
    for (const auto& current_hash : hashes_result.data) {
        combined_hash = hash(combined_hash, current_hash);
    }

    return {rocksdb::Status::OK(), combined_hash};
}

ValueResult<uint256_t> ArbCore::getInboxAcc(uint256_t index) {
    auto hashes_result = getInboxHashes(index, 1);
    if (!hashes_result.status.ok()) {
        return {hashes_result.status, 0};
    }

    return {rocksdb::Status::OK(), hashes_result.data[0]};
}

ValueResult<uint256_t> ArbCore::getSendAcc(uint256_t start_acc_hash,
                                           uint256_t start_index,
                                           uint256_t count) {
    auto sends_result = getSends(start_index, count);
    if (!sends_result.status.ok()) {
        return {sends_result.status, 0};
    }

    auto combined_hash = start_acc_hash;
    for (const auto& send : sends_result.data) {
        combined_hash = hash(combined_hash, hash(send));
    }
    return {rocksdb::Status::OK(), 0};
}

ValueResult<uint256_t> ArbCore::getLogAcc(uint256_t start_acc_hash,
                                          uint256_t start_index,
                                          uint256_t count,
                                          ValueCache& cache) {
    auto sends_result = getLogs(start_index, count, cache);
    if (!sends_result.status.ok()) {
        return {sends_result.status, 0};
    }

    auto combined_hash = start_acc_hash;
    for (const auto& send : sends_result.data) {
        combined_hash = hash(combined_hash, hash_value(send));
    }
    return {rocksdb::Status::OK(), 0};
}

ValueResult<ExecutionCursor*> ArbCore::getExecutionCursor(
    uint256_t totalGasUsed,
    ValueCache& cache) {
    auto tx = Transaction::makeTransaction(data_storage);
    auto first_checkpoint_result =
        getCheckpointUsingGas(*tx, totalGasUsed, false);
    if (!first_checkpoint_result.status.ok()) {
        return {first_checkpoint_result.status, {}};
    }
    auto first_message_sequence_number =
        first_checkpoint_result.data.total_messages_read - 1;

    uint256_t last_message_sequence_number;
    auto second_checkpoint_result =
        getCheckpointUsingGas(*tx, totalGasUsed, true);
    if (second_checkpoint_result.status.IsNotFound()) {
        // No checkpoint after requested gas so get messages up to current
        // machine state.
        last_message_sequence_number =
            pending_checkpoint.total_messages_read - 1;
    } else if (!second_checkpoint_result.status.ok()) {
        return {second_checkpoint_result.status, {}};
    } else {
        last_message_sequence_number =
            second_checkpoint_result.data.total_messages_read - 1;
    }

    std::vector<unsigned char> message_key;
    marshal_uint256_t(first_message_sequence_number, message_key);
    auto message_key_slice = vecToSlice(message_key);

    auto results = getVectorVectorUsingFamilyAndKey(
        *tx->transaction, data_storage->send_column.get(), message_key_slice,
        intx::narrow_cast<size_t>(last_message_sequence_number -
                                  first_message_sequence_number));
    if (!results.status.ok()) {
        return {results.status, {}};
    }

    std::vector<Tuple> messages;
    std::vector<uint256_t> inbox_hashes;
    nonstd::optional<uint256_t> min_next_block_height;
    auto total_size = results.data.size();
    messages.reserve(total_size);
    inbox_hashes.reserve(total_size);
    for (const auto& data : results.data) {
        auto message_entry = extractMessageEntry(0, vecToSlice(data));

        messages.push_back(messageDataToTuple(message_entry.data));
        inbox_hashes.push_back(message_entry.inbox_hash);
        if (messages.size() == total_size &&
            message_entry.last_message_in_block) {
            min_next_block_height = message_entry.block_height + 1;
        }
    }

    auto cursor_machine = getMachineUsingStateKeys<Machine>(
        *tx, first_checkpoint_result.data.machine_state_keys, cache);

    auto cursor =
        new ExecutionCursor{first_checkpoint_result.data, cursor_machine,
                            messages, inbox_hashes, min_next_block_height};

    // Run machine until specified gas is reached
    auto remaining_gas =
        totalGasUsed - first_checkpoint_result.data.arb_gas_used;
    if (remaining_gas > 0) {
        cursor->Advance(remaining_gas, false);
    }

    return {rocksdb::Status::OK(), cursor};
}

void checkMessages() {}

ValueResult<uint256_t> ArbCore::logInsertedCount(Transaction& tx) const {
    return getUint256UsingFamilyAndKey(*tx.transaction,
                                       tx.datastorage->state_column.get(),
                                       vecToSlice(log_inserted_key));
}
rocksdb::Status ArbCore::updateLogInsertedCount(Transaction& tx,
                                                rocksdb::Slice value_slice) {
    return tx.transaction->Put(tx.datastorage->state_column.get(),
                               vecToSlice(log_inserted_key), value_slice);
}

ValueResult<uint256_t> ArbCore::logProcessedCount(Transaction& tx) const {
    return getUint256UsingFamilyAndKey(*tx.transaction,
                                       tx.datastorage->state_column.get(),
                                       vecToSlice(log_processed_key));
}
rocksdb::Status ArbCore::updateLogProcessedCount(Transaction& tx,
                                                 rocksdb::Slice value_slice) {
    return tx.transaction->Put(tx.datastorage->state_column.get(),
                               vecToSlice(log_processed_key), value_slice);
}

ValueResult<uint256_t> ArbCore::sendInsertedCount(Transaction& tx) const {
    return getUint256UsingFamilyAndKey(*tx.transaction,
                                       tx.datastorage->state_column.get(),
                                       vecToSlice(send_inserted_key));
}
rocksdb::Status ArbCore::updateSendInsertedCount(Transaction& tx,
                                                 rocksdb::Slice value_slice) {
    return tx.transaction->Put(tx.datastorage->state_column.get(),
                               vecToSlice(send_inserted_key), value_slice);
}

ValueResult<uint256_t> ArbCore::sendProcessedCount(Transaction& tx) const {
    return getUint256UsingFamilyAndKey(*tx.transaction,
                                       tx.datastorage->state_column.get(),
                                       vecToSlice(send_processed_key));
}
rocksdb::Status ArbCore::updateSendProcessedCount(Transaction& tx,
                                                  rocksdb::Slice value_slice) {
    return tx.transaction->Put(tx.datastorage->state_column.get(),
                               vecToSlice(send_processed_key), value_slice);
}

ValueResult<uint256_t> ArbCore::messageEntryInsertedCount(
    Transaction& tx) const {
    return getUint256UsingFamilyAndKey(*tx.transaction,
                                       tx.datastorage->state_column.get(),
                                       vecToSlice(message_entry_inserted_key));
}
rocksdb::Status ArbCore::updateMessageEntryInsertedCount(
    Transaction& tx,
    rocksdb::Slice value_slice) {
    return tx.transaction->Put(tx.datastorage->state_column.get(),
                               vecToSlice(message_entry_inserted_key),
                               value_slice);
}

ValueResult<uint256_t> ArbCore::messageEntryProcessedCount(
    Transaction& tx) const {
    return getUint256UsingFamilyAndKey(*tx.transaction,
                                       tx.datastorage->state_column.get(),
                                       vecToSlice(message_entry_processed_key));
}
rocksdb::Status ArbCore::updateMessageEntryProcessedCount(
    Transaction& tx,
    rocksdb::Slice value_slice) {
    return tx.transaction->Put(tx.datastorage->state_column.get(),
                               vecToSlice(message_entry_processed_key),
                               value_slice);
}

// addMessages stores all messages from given block into database.
// The last message in the list is flagged as the last message in the block.
// Returns nonstd::nullopt when caller needs to provide messages from earlier
// block.
nonstd::optional<rocksdb::Status> ArbCore::addMessages(
    uint256_t first_sequence_number,
    uint64_t block_height,
    const std::vector<std::vector<unsigned char>>& messages,
    const uint256_t& previous_inbox_hash,
    const uint256_t& final_machine_sequence_number) {
    auto tx = Transaction::makeTransaction(data_storage);

    // Get the last message sequence number that was added to database
    auto message_count_result = messageEntryInsertedCount(*tx);
    if (message_count_result.status.ok()) {
        return message_count_result.status;
    }
    auto last_inserted_sequence_number = message_count_result.data - 1;

    if (first_sequence_number > 0) {
        if (first_sequence_number > last_inserted_sequence_number + 1) {
            // Not allowed to skip message sequence numbers, ask for older
            // messages
            return nonstd::nullopt;
        }

        // Check that previous_inbox_hash matches hash from previous message
        auto previous_sequence_number = first_sequence_number - 1;
        auto previous_result = getMessageEntry(*tx, previous_sequence_number);
        if (!previous_result.status.ok()) {
            return previous_result.status;
        }

        if (previous_result.data.inbox_hash != previous_inbox_hash) {
            // Previous inbox doesn't match which means reorg happened and
            // caller needs to try again with messages from earlier block
            return nonstd::nullopt;
        }
    }

    size_t current_message_index = 0;
    auto current_sequence_number = first_sequence_number;
    auto final_sequence_number = first_sequence_number + messages.size() - 1;

    if (messages.empty()) {
        // No new messages, just need to truncating obsolete messages
        current_sequence_number = first_sequence_number - 1;
    }

    // Skip any valid messages that we already have in database
    auto messages_count = messages.size();
    auto current_previous_inbox_hash = previous_inbox_hash;
    while ((current_sequence_number <= last_inserted_sequence_number) &&
           (current_message_index < messages_count)) {
        auto existing_message_entry =
            getMessageEntry(*tx, current_sequence_number);
        if (!existing_message_entry.status.ok()) {
            return existing_message_entry.status;
        }
        auto current_inbox_hash = hash_inbox(current_previous_inbox_hash,
                                             messages[current_message_index]);
        if (existing_message_entry.data.inbox_hash != current_inbox_hash) {
            // Entry doesn't match because of reorg
            break;
        }

        current_message_index++;
        current_previous_inbox_hash = current_inbox_hash;
        current_sequence_number = first_sequence_number + current_message_index;
    }

    if (current_sequence_number <= last_inserted_sequence_number) {
        // Reorg occurred
        const std::lock_guard<std::mutex> lock(core_mutex);

        if (final_machine_sequence_number >= current_sequence_number) {
            // Machine is running with obsolete messages
            machine->abortThread();
        }

        auto previous_valid_sequence_number = current_sequence_number - 1;
        std::vector<unsigned char> previous_message_count_key;
        marshal_uint256_t(previous_valid_sequence_number + 1,
                          previous_message_count_key);

        // Truncate MessageEntries to last valid message
        updateMessageEntryInsertedCount(*tx,
                                        vecToSlice(previous_message_count_key));

        // Reorg checkpoint and everything else
        auto reorg_status =
            reorgToMessageOrBefore(*tx, previous_valid_sequence_number);
    }

    while (current_message_index < messages_count) {
        // Encode key
        std::vector<unsigned char> key;
        marshal_uint256_t(current_sequence_number, key);

        auto current_inbox_hash = hash_inbox(current_previous_inbox_hash,
                                             messages[current_message_index]);

        // Encode message entry
        auto messageEntry = MessageEntry{
            current_sequence_number, current_inbox_hash, block_height,
            current_sequence_number == final_sequence_number,
            (messages[current_message_index])};
        auto serialized_messageentry = serializeMessageEntry(messageEntry);

        // Save message entry into database
        auto put_status = tx->transaction->Put(
            tx->datastorage->messageentry_column.get(), vecToSlice(key),
            vecToSlice(serialized_messageentry));
        if (!put_status.ok()) {
            return put_status;
        }

        current_message_index++;
        current_previous_inbox_hash = current_inbox_hash;
        current_sequence_number += 1
    }

    std::vector<unsigned char> count_key;
    marshal_uint256_t(current_sequence_number, count_key);
    updateMessageEntryInsertedCount(*tx, vecToSlice(count_key));

    return tx->commit();
}

// deleteLogsStartingAt deletes the given index along with any
// newer logs. Returns nonstd::nullopt if nothing deleted.
nonstd::optional<rocksdb::Status> deleteLogsStartingAt(Transaction& tx,
                                                       uint256_t log_index) {
    auto it = std::unique_ptr<rocksdb::Iterator>(tx.transaction->GetIterator(
        rocksdb::ReadOptions(), tx.datastorage->log_column.get()));

    // Find first message to delete
    std::vector<unsigned char> key;
    marshal_uint256_t(log_index, key);
    it->Seek(vecToSlice(key));
    if (it->status().IsNotFound()) {
        // Nothing to delete
        return nonstd::nullopt;
    }
    if (!it->status().ok()) {
        return it->status();
    }

    while (it->Valid()) {
        // Remove reference to value
        auto value_hash_ptr = reinterpret_cast<const char*>(it->value().data());
        deleteValue(tx, deserializeUint256t(value_hash_ptr));

        it->Next();
    }
    if (!it->status().ok()) {
        return it->status();
    }

    return rocksdb::Status::OK();
}

// getNextMessage returns the next message to handle.
nonstd::optional<MessageEntry> ArbCore::getNextMessage() {
    auto tx = Transaction::makeTransaction(data_storage);
    auto it = std::unique_ptr<rocksdb::Iterator>(tx->transaction->GetIterator(
        rocksdb::ReadOptions(), tx->datastorage->messageentry_column.get()));

    it->SeekToFirst();
    if (!it->Valid()) {
        return nonstd::nullopt;
    }

    auto key = reinterpret_cast<const char*>(it->key().data());
    return extractMessageEntry(deserializeUint256t(key), it->value());
}

// deleteMessage deletes the provided message only if it has not changed in DB
bool ArbCore::deleteMessage(const MessageEntry& entry) {
    auto tx = Transaction::makeTransaction(data_storage);

    std::vector<unsigned char> key;
    marshal_uint256_t(entry.sequence_number, key);
    auto key_slice = vecToSlice(key);
    std::string value;
    auto get_status = tx->transaction->GetForUpdate(
        rocksdb::ReadOptions(), tx->datastorage->messageentry_column.get(),
        key_slice, &value);
    if (!get_status.ok()) {
        std::cerr << "In deleteMessage get: " << get_status.ToString()
                  << std::endl;
        return false;
    }

    auto db_entry =
        extractMessageEntry(entry.sequence_number, rocksdb::Slice(value));
    if (entry != db_entry) {
        // Entry changed, reorg probably occurred
        return false;
    }

    // Delete message entry
    auto delete_status = tx->transaction->Delete(
        tx->datastorage->messageentry_column.get(), key_slice);
    if (!delete_status.ok()) {
        std::cerr << "In deleteMessage delete: " << delete_status.ToString()
                  << std::endl;
        return false;
    }

    auto commit_status = tx->commit();
    if (!commit_status.ok()) {
        std::cerr << "In deleteMessage commit: " << commit_status.ToString()
                  << std::endl;
        return false;
    }

    return true;
}
