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

    return {result.status, std::move(parsed_state)};
}

}  // namespace

bool ArbCore::machineIdle() {
    return machine_idle;
}

ArbCore::message_status_enum ArbCore::messagesStatus() {
    auto current_status = message_data_status.load();
    if (current_status != MESSAGES_ERROR && current_status != MESSAGES_READY) {
        message_data_status = MESSAGES_EMPTY;
    }
    return current_status;
}

std::string ArbCore::messagesClearError() {
    if (message_data_status != ArbCore::MESSAGES_ERROR &&
        message_data_status != ArbCore::MESSAGES_NEED_OLDER) {
        return nullptr;
    }

    message_data_status = MESSAGES_EMPTY;
    auto str = core_error_string;
    core_error_string.clear();

    return str;
}

std::string ArbCore::machineClearError() {
    if (!machine_error) {
        return nullptr;
    }

    machine_error = false;
    auto str = machine_error_string;
    machine_error_string.clear();

    return str;
}

bool ArbCore::startThread() {
    abortThread();

    core_thread =
        std::make_unique<std::thread>((std::reference_wrapper<ArbCore>(*this)));

    return true;
}

void ArbCore::abortThread() {
    if (core_thread) {
        arbcore_abort = true;
        core_thread->join();
        core_thread = nullptr;
    }
    arbcore_abort = false;
}

// deliverMessages sends messages to core thread
bool ArbCore::deliverMessages(std::vector<std::vector<unsigned char>>& messages,
                              const uint256_t& previous_inbox_hash,
                              const bool last_block_complete) {
    if (message_data_status != MESSAGES_EMPTY) {
        return false;
    }

    message_data.messages = std::move(messages);
    message_data.previous_inbox_hash = previous_inbox_hash;
    message_data.last_block_complete = last_block_complete;

    message_data_status = MESSAGES_READY;

    return true;
}

std::unique_ptr<Transaction> ArbCore::makeTransaction() {
    return Transaction::makeTransaction(data_storage);
}

std::unique_ptr<const Transaction> ArbCore::makeConstTransaction() const {
    auto transaction =
        std::unique_ptr<rocksdb::Transaction>(data_storage->beginTransaction());
    return std::make_unique<Transaction>(data_storage, std::move(transaction));
}

rocksdb::Status ArbCore::initialize(const LoadedExecutable& executable) {
    auto tx = makeTransaction();

    code = std::make_shared<Code>(0);
    code->addSegment(executable.code);
    machine = std::make_unique<MachineThread>(
        MachineState{code, executable.static_val});

    auto result = getInitialMachineHash(*tx);
    if (result.status.ok() && machine->hash() == result.data) {
        if (machine->hash() != result.data) {
            // Need to delete database and start from scratch
            std::cerr << "Incorrect initial machine in database" << std::endl;
            return rocksdb::Status::Corruption();
        }

        // Use latest existing checkpoint
        ValueCache cache;
        auto status = reorgToMessageOrBefore(*tx, 0, true, cache);
        if (!status.ok()) {
            std::cerr << "Error with initial reorg: " << status.ToString()
                      << std::endl;
            return status;
        }

        // Make sure logs cursors are starting at the correct logs
        for (auto& logs_cursor : logs_cursors) {
            logs_cursor.current_total_count = pending_checkpoint.log_count;
        }
    } else {
        // Need to initialize database from scratch
        auto res = saveMachine(*tx, *machine);
        if (!res.status.ok()) {
            std::cerr << "failed to save initial machine: "
                      << res.status.ToString() << std::endl;
            return res.status;
        }

        std::vector<unsigned char> value_data;
        marshal_uint256_t(machine->hash(), value_data);
        auto s = tx->transaction->Put(data_storage->state_column.get(),
                                      vecToSlice(initial_machine_hash_key),
                                      vecToSlice(value_data));
        if (!s.ok()) {
            std::cerr << "failed to save initial machine values into db: "
                      << res.status.ToString() << std::endl;
            return s;
        }

        s = saveCheckpoint(*tx);
        if (!s.ok()) {
            std::cerr << "failed to save initial checkpoint into db: "
                      << s.ToString() << std::endl;
            return s;
        }

        auto status = updateLogInsertedCount(*tx, 0);
        if (!status.ok()) {
            throw std::runtime_error("failed to initialize log inserted count");
        }
        status = updateSendInsertedCount(*tx, 0);
        if (!status.ok()) {
            throw std::runtime_error("failed to initialize log inserted count");
        }
        status = updateMessageEntryInsertedCount(*tx, 0);
        if (!status.ok()) {
            throw std::runtime_error("failed to initialize log inserted count");
        }
        status = updateMessageEntryProcessedCount(*tx, 0);
        if (!status.ok()) {
            throw std::runtime_error("failed to initialize log inserted count");
        }
    }

    auto s = tx->commit();
    if (!s.ok()) {
        std::cerr << "failed to commit initial state into db: " << s.ToString()
                  << std::endl;
        return s;
    }

    return rocksdb::Status::OK();
}

bool ArbCore::initialized() const {
    auto tx = makeConstTransaction();
    std::string initial_raw;
    auto s = tx->transaction->GetForUpdate(
        rocksdb::ReadOptions(), data_storage->state_column.get(),
        vecToSlice(initial_machine_hash_key), &initial_raw);
    return s.ok();
}

ValueResult<uint256_t> ArbCore::getInitialMachineHash(Transaction& tx) {
    std::string initial_raw;
    auto s = tx.transaction->GetForUpdate(
        rocksdb::ReadOptions(), data_storage->state_column.get(),
        vecToSlice(initial_machine_hash_key), &initial_raw);
    if (!s.ok()) {
        return {s, 0};
    }

    return {rocksdb::Status::OK(),
            intx::be::unsafe::load<uint256_t>(
                reinterpret_cast<const unsigned char*>(initial_raw.data()))};
}

template <class T>
std::unique_ptr<T> ArbCore::getInitialMachineImpl(Transaction& tx,
                                                  ValueCache& value_cache) {
    auto machine_hash = getInitialMachineHash(tx);
    return getMachine<T>(machine_hash.data, value_cache);
}

template std::unique_ptr<Machine> ArbCore::getInitialMachineImpl(Transaction&,
                                                                 ValueCache&);
template std::unique_ptr<MachineThread> ArbCore::getInitialMachineImpl(
    Transaction&,
    ValueCache&);

template <class T>
std::unique_ptr<T> ArbCore::getInitialMachine(ValueCache& value_cache) {
    auto tx = makeTransaction();
    return getInitialMachineImpl<T>(*tx, value_cache);
}

template std::unique_ptr<Machine> ArbCore::getInitialMachine(ValueCache&);
template std::unique_ptr<MachineThread> ArbCore::getInitialMachine(ValueCache&);

template <class T>
std::unique_ptr<T> ArbCore::getMachineImpl(Transaction& tx,
                                           uint256_t machineHash,
                                           ValueCache& value_cache) {
    auto results = getMachineStateKeys(tx, machineHash);
    if (!results.status.ok()) {
        throw std::runtime_error("failed to load machine state");
    }

    return getMachineUsingStateKeys<T>(tx, results.data, value_cache);
}

template std::unique_ptr<Machine> ArbCore::getMachineImpl(Transaction&,
                                                          uint256_t,
                                                          ValueCache&);
template std::unique_ptr<MachineThread> ArbCore::getMachineImpl(Transaction&,
                                                                uint256_t,
                                                                ValueCache&);

template <class T>
std::unique_ptr<T> ArbCore::getMachine(uint256_t machineHash,
                                       ValueCache& value_cache) {
    auto tx = makeTransaction();
    return getMachineImpl<T>(*tx, machineHash, value_cache);
}

template std::unique_ptr<Machine> ArbCore::getMachine(uint256_t, ValueCache&);
template std::unique_ptr<MachineThread> ArbCore::getMachine(uint256_t,
                                                            ValueCache&);

rocksdb::Status ArbCore::saveCheckpoint(Transaction& tx) {
    auto status =
        saveMachineState(tx, *machine, pending_checkpoint.machine_state_keys);
    if (!status.ok()) {
        return status;
    }

    // Pull inbox hash from database
    if (pending_checkpoint.total_messages_read > 0) {
        auto existing_message_entry =
            getMessageEntry(tx, pending_checkpoint.total_messages_read - 1);
        pending_checkpoint.inbox_hash = existing_message_entry.data.inbox_hash;
    } else {
        pending_checkpoint.inbox_hash = 0;
    }

    std::vector<unsigned char> key;
    marshal_uint256_t(pending_checkpoint.arb_gas_used, key);
    auto key_slice = vecToSlice(key);
    auto serialized_checkpoint = serializeCheckpoint(pending_checkpoint);
    std::string value_str(serialized_checkpoint.begin(),
                          serialized_checkpoint.end());
    auto put_status = tx.transaction->Put(
        tx.datastorage->checkpoint_column.get(), key_slice, value_str);
    if (!put_status.ok()) {
        return put_status;
    }

    return rocksdb::Status::OK();
}

rocksdb::Status ArbCore::saveAssertion(Transaction& tx,
                                       const Assertion& assertion) {
    auto status = saveLogs(tx, assertion.logs);
    if (!status.ok()) {
        return status;
    }

    status = saveSends(tx, assertion.sends);
    if (!status.ok()) {
        return status;
    }

    pending_checkpoint.applyAssertion(assertion);

    updateMessageEntryProcessedCount(tx,
                                     pending_checkpoint.total_messages_read);

    return rocksdb::Status::OK();
}

// reorgToMessageOrBefore resets the checkpoint and database entries
// such that machine state is at or before the requested message. cleaning
// up old references as needed.
// If use_latest is true, message_sequence_number is ignored and the latest
// checkpoint is used.
rocksdb::Status ArbCore::reorgToMessageOrBefore(
    Transaction& tx,
    const uint256_t& message_sequence_number,
    const bool use_latest,
    ValueCache& cache) {
    auto it = std::unique_ptr<rocksdb::Iterator>(tx.transaction->GetIterator(
        rocksdb::ReadOptions(), tx.datastorage->checkpoint_column.get()));

    // Find first checkpoint to delete
    it->SeekToLast();
    if (!it->status().ok()) {
        return it->status();
    }

    // Delete each checkpoint until at or below message_sequence_number
    auto good_checkpoint_found = false;
    Checkpoint last_deleted_checkpoint;
    if (use_latest) {
        std::vector<unsigned char> checkpoint_vector(
            it->value().data(), it->value().data() + it->value().size());
        pending_checkpoint = extractCheckpoint(checkpoint_vector);
        good_checkpoint_found = true;
    } else {
        while (it->Valid()) {
            std::vector<unsigned char> checkpoint_vector(
                it->value().data(), it->value().data() + it->value().size());
            auto checkpoint = extractCheckpoint(checkpoint_vector);

            if (message_sequence_number >= checkpoint.total_messages_read - 1) {
                // Good checkpoint
                pending_checkpoint = checkpoint;
                good_checkpoint_found = true;
                break;
            }

            // Obsolete checkpoint, need to delete referenced machine
            last_deleted_checkpoint = checkpoint;
            deleteMachineState(tx, checkpoint.machine_state_keys);

            // Delete checkpoint to make sure it isn't used later
            tx.transaction->Delete(tx.datastorage->checkpoint_column.get(),
                                   it->key());

            it->Prev();
            if (!it->status().ok()) {
                return it->status();
            }
        }
    }
    it = nullptr;

    if (!good_checkpoint_found) {
        // Nothing found, start database from scratch.
        // Machine will be reset when it is confirmed to not be running.

        // Clearing checkpoint will ensure everything is reset properly below
        pending_checkpoint = Checkpoint{};
    }

    // Update log cursors, must be called before logs are deleted
    for (size_t i = 0; i < logs_cursors.size(); i++) {
        auto status =
            handleLogsCursorReorg(tx, i, pending_checkpoint.log_count, cache);
        if (!status.ok()) {
            return status;
        }
    }

    // Delete logs individually to handle reference counts
    auto optional_status =
        deleteLogsStartingAt(tx, pending_checkpoint.log_count);
    if (optional_status && !optional_status->ok()) {
        return *optional_status;
    }

    auto status = updateLogInsertedCount(tx, pending_checkpoint.log_count);
    if (!status.ok()) {
        return status;
    }

    status = updateSendInsertedCount(tx, pending_checkpoint.send_count);
    if (!status.ok()) {
        return status;
    }

    status = updateMessageEntryProcessedCount(
        tx, pending_checkpoint.total_messages_read);
    if (!status.ok()) {
        return status;
    }

    return rocksdb::Status::OK();
}

ValueResult<Checkpoint> ArbCore::getCheckpoint(
    Transaction& tx,
    const uint256_t& arb_gas_used) const {
    std::vector<unsigned char> key;
    marshal_uint256_t(arb_gas_used, key);

    auto result = getVectorUsingFamilyAndKey(
        *tx.transaction, tx.datastorage->checkpoint_column.get(),
        vecToSlice(key));
    if (!result.status.ok()) {
        return {result.status, {}};
    }

    return {rocksdb::Status::OK(), extractCheckpoint(result.data)};
}

bool ArbCore::isCheckpointsEmpty(Transaction& tx) const {
    auto it = std::unique_ptr<rocksdb::Iterator>(tx.transaction->GetIterator(
        rocksdb::ReadOptions(), tx.datastorage->checkpoint_column.get()));
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

// getCheckpointUsingGas returns the checkpoint at or before the specified gas
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
    auto parsed_state = extractCheckpoint(saved_value);
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
                     state_data.total_messages_consumed,
                     std::move(staged_message_results.data)};

    return std::make_unique<T>(state);
}

template std::unique_ptr<Machine>
ArbCore::getMachineUsingStateKeys(Transaction&, MachineStateKeys, ValueCache&);
template std::unique_ptr<MachineThread>
ArbCore::getMachineUsingStateKeys(Transaction&, MachineStateKeys, ValueCache&);

// operator() runs the main thread for ArbCore.  It is responsible for adding
// messages to the queue, starting machine thread when needed and colleting
// results of machine thread.
// This thread will update `delivering_messages` if and only if
// `delivering_messages` is set to MESSAGES_READY
void ArbCore::operator()() {
    ValueCache cache;
    uint256_t first_sequence_number_in_machine;
    uint256_t last_sequence_number_in_machine;

    while (!arbcore_abort) {
        if (message_data_status == MESSAGES_READY) {
            // Reorg might occur while adding messages
            auto add_status = addMessages(
                message_data.messages, message_data.previous_inbox_hash,
                last_sequence_number_in_machine,
                message_data.last_block_complete, cache);
            if (!add_status) {
                // Messages from previous block invalid because of reorg so
                // request older messages
                message_data_status = MESSAGES_NEED_OLDER;
            } else if (!add_status->ok()) {
                core_error_string = add_status->ToString();
                message_data_status = MESSAGES_ERROR;
                std::cerr << "ArbCore inbox processed stopped with error: "
                          << core_error_string << "\n";
                break;
            } else {
                machine_idle = false;
                message_data_status = MESSAGES_SUCCESS;
            }
        }

        // Check machine thread
        if (machine->status() == MachineThread::MACHINE_ERROR) {
            core_error_string = machine->getErrorString();
            std::cerr << "AVM machine stopped with error: " << core_error_string
                      << "\n";
            break;
        }

        if (machine->status() == MachineThread::MACHINE_ABORTED) {
            // Machine was executing obsolete messages so restore machine
            // from last checkpoint
            auto tx = Transaction::makeTransaction(data_storage);

            machine = getMachineUsingStateKeys<MachineThread>(
                *tx, pending_checkpoint.machine_state_keys, cache);
        } else if (machine->status() == MachineThread::MACHINE_SUCCESS) {
            auto tx = Transaction::makeTransaction(data_storage);

            auto last_assertion = machine->getAssertion();
            auto messages_inserted = messageEntryInsertedCountImpl(*tx);
            if (!messages_inserted.status.ok()) {
                core_error_string = messages_inserted.status.ToString();
                std::cerr << "ArbCore message insertion failed: "
                          << core_error_string << "\n";
                break;
            }
            auto total_messages_consumed =
                first_sequence_number_in_machine +
                last_assertion.inbox_messages_consumed;

            if (total_messages_consumed > messages_inserted.data) {
                // Machine consumed obsolete message, restore from checkpoint
                machine = getMachineUsingStateKeys<MachineThread>(
                    *tx, pending_checkpoint.machine_state_keys, cache);
            } else {
                // Save logs and sends
                auto status = saveAssertion(*tx, last_assertion);
                if (!status.ok()) {
                    core_error_string = status.ToString();
                    std::cerr << "ArbCore assertion saving failed: "
                              << core_error_string << "\n";
                    break;
                }

                // Maybe save checkpoint
                // TODO Decide how often to create checkpoint
                status = saveCheckpoint(*tx);

                // TODO Decide how often to clear ValueCache
                // (only clear cache when machine thread stopped)
                cache.clear();
            }

            auto status = tx->commit();
            if (!status.ok()) {
                core_error_string = status.ToString();
                machine_error = true;
                std::cerr << "ArbCore database update failed: "
                          << core_error_string << "\n";
                break;
            }
        }

        if (machine->status() == MachineThread::MACHINE_NONE) {
            // Start execution of machine if new message available
            auto tx = Transaction::makeTransaction(data_storage);
            auto messages_count = messageEntryInsertedCountImpl(*tx);
            if (!messages_count.status.ok()) {
                core_error_string = messages_count.status.ToString();
                machine_error = true;
                std::cerr << "ArbCore message count fetching failed: "
                          << core_error_string << "\n";
                break;
            }

            if (messages_count.data < pending_checkpoint.total_messages_read) {
                // Should never happen, means reorg wasn't done properly
                core_error_string = "messages_inserted < pending_checkpoint";
                machine_error = true;
                std::cerr << "ArbCore reorg error: " << core_error_string
                          << "\n";
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
                    core_error_string = next_message_result.status.ToString();
                    machine_error = true;
                    std::cerr << "ArbCore failed getting message entry: "
                              << core_error_string << "\n";
                    break;
                }
                if (next_message_result.data.sequence_number !=
                    first_sequence_number_in_machine) {
                    core_error_string =
                        "sequence number in message different than expected";
                    machine_error = true;
                    std::cerr << "ArbCore error: " << core_error_string << "\n";
                    break;
                }
                std::vector<std::vector<unsigned char>> messages;
                messages.push_back(next_message_result.data.data);

                auto status = machine->runMachine(
                    0, false, std::move(messages), 0,
                    next_message_result.data.last_message_in_block);
                if (!status) {
                    core_error_string = "Error starting machine thread";
                    machine_error = true;
                    std::cerr << "ArbCore error: " << core_error_string << "\n";
                    break;
                }
            } else {
                // Machine all caught up, no messages to process
                machine_idle = true;
            }
        }

        for (size_t i = 0; i < logs_cursors.size(); i++) {
            if (logs_cursors[i].status == DataCursor::REQUESTED) {
                auto tx = Transaction::makeTransaction(data_storage);
                handleLogsCursorRequested(*tx, i, cache);
            }
        }

        if (!machineIdle() || message_data_status != MESSAGES_READY) {
            // Machine is already running or new messages, so sleep for a short
            // while
            std::this_thread::sleep_for(std::chrono::milliseconds(200));
        }
    }

    // Error occurred, make sure machine stops cleanly
    machine->abortMachine();
}

rocksdb::Status ArbCore::saveLogs(Transaction& tx,
                                  const std::vector<value>& vals) {
    auto log_result = logInsertedCountImpl(tx);
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

    return updateLogInsertedCount(tx, log_index);
}

ValueResult<std::vector<value>> ArbCore::getLogs(uint256_t index,
                                                 uint256_t count,
                                                 ValueCache& valueCache) {
    auto tx = Transaction::makeTransaction(data_storage);

    // Acquire mutex to make sure no reorg happening
    std::lock_guard<std::mutex> lock(core_reorg_mutex);

    return getLogsNoLock(*tx, index, count, valueCache);
}

ValueResult<std::vector<value>> ArbCore::getLogsNoLock(Transaction& tx,
                                                       uint256_t index,
                                                       uint256_t count,
                                                       ValueCache& valueCache) {
    // Check if attempting to get entries past current valid logs
    auto log_count = logInsertedCountImpl(tx);
    if (!log_count.status.ok()) {
        return {log_count.status, {}};
    }
    auto max_log_count = log_count.data;
    if (index >= max_log_count) {
        return {rocksdb::Status::NotFound(), {}};
    }
    if (index + count > max_log_count) {
        count = max_log_count - index;
    }

    std::vector<unsigned char> key;
    marshal_uint256_t(index, key);

    auto hash_result = getUint256VectorUsingFamilyAndKey(
        *tx.transaction, data_storage->log_column.get(), vecToSlice(key),
        intx::narrow_cast<size_t>(count));
    if (!hash_result.status.ok()) {
        return {hash_result.status, {}};
    }

    std::vector<value> logs;
    for (const auto& hash : hash_result.data) {
        auto val_result = getValue(tx, hash, valueCache);
        if (!val_result.status.ok()) {
            return {val_result.status, {}};
        }
        logs.push_back(std::move(val_result.data));
    }

    return {rocksdb::Status::OK(), std::move(logs)};
}

rocksdb::Status ArbCore::saveSends(
    Transaction& tx,
    const std::vector<std::vector<unsigned char>>& sends) {
    auto send_result = sendInsertedCountImpl(tx);
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

    return updateSendInsertedCount(tx, send_count);
}

ValueResult<std::vector<uint256_t>> ArbCore::getInboxHashes(
    uint256_t index,
    uint256_t count) const {
    auto tx = Transaction::makeTransaction(data_storage);

    // Check if attempting to get entries past current valid logs
    auto message_count_result = messageEntryInsertedCountImpl(*tx);
    if (!message_count_result.status.ok()) {
        return {message_count_result.status, {}};
    }
    auto max_message_count = message_count_result.data;
    if (index >= max_message_count) {
        return {rocksdb::Status::NotFound(), {}};
    }
    if (index + count > max_message_count) {
        count = max_message_count - index;
    }

    std::vector<unsigned char> key;
    marshal_uint256_t(index, key);
    auto key_slice = vecToSlice(key);

    auto results = getVectorVectorUsingFamilyAndKey(
        *tx->transaction, data_storage->messageentry_column.get(), key_slice,
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

    return {rocksdb::Status::OK(), std::move(messages)};
}

ValueResult<std::vector<std::vector<unsigned char>>> ArbCore::getMessages(
    uint256_t index,
    uint256_t count) const {
    auto tx = Transaction::makeTransaction(data_storage);

    // Check if attempting to get entries past current valid logs
    auto message_count = messageEntryInsertedCountImpl(*tx);
    if (!message_count.status.ok()) {
        return {message_count.status, {}};
    }
    auto max_message_count = message_count.data;
    if (index >= max_message_count) {
        return {rocksdb::Status::NotFound(), {}};
    }
    if (index + count > max_message_count) {
        count = max_message_count - index;
    }

    std::vector<unsigned char> key;
    marshal_uint256_t(index, key);
    auto key_slice = vecToSlice(key);

    auto results = getVectorVectorUsingFamilyAndKey(
        *tx->transaction, data_storage->messageentry_column.get(), key_slice,
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

    return {rocksdb::Status::OK(), std::move(messages)};
}

ValueResult<std::vector<uint256_t>> ArbCore::getMessageHashes(
    uint256_t index,
    uint256_t count) const {
    auto messages = getMessages(index, count);
    if (!messages.status.ok()) {
        return {messages.status, {}};
    }

    std::vector<uint256_t> hashes;
    hashes.reserve(messages.data.size());
    for (const auto& message : messages.data) {
        auto hash = hash_value(extractInboxMessage(message).toTuple());
        hashes.push_back(hash);
    }

    return {rocksdb::Status::OK(), std::move(hashes)};
}

ValueResult<std::vector<std::vector<unsigned char>>> ArbCore::getSends(
    uint256_t index,
    uint256_t count) const {
    auto tx = Transaction::makeTransaction(data_storage);

    // Check if attempting to get entries past current valid logs
    auto send_count = sendInsertedCountImpl(*tx);
    if (!send_count.status.ok()) {
        return {send_count.status, {}};
    }
    auto max_send_count = send_count.data;
    if (index >= max_send_count) {
        return {rocksdb::Status::NotFound(), {}};
    }
    if (index + count > max_send_count) {
        count = max_send_count - index;
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
    auto hashes_result = getMessageHashes(start_index, count);
    if (!hashes_result.status.ok()) {
        return {hashes_result.status, 0};
    }

    uint256_t combined_hash = 0;
    for (size_t i = 0; i < hashes_result.data.size(); ++i) {
        combined_hash =
            hash(combined_hash,
                 hashes_result.data[hashes_result.data.size() - 1 - i]);
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
    return {rocksdb::Status::OK(), combined_hash};
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
    return {rocksdb::Status::OK(), combined_hash};
}

ValueResult<std::unique_ptr<ExecutionCursor>> ArbCore::getExecutionCursor(
    uint256_t total_gas_used,
    ValueCache& cache) {
    auto tx = Transaction::makeTransaction(data_storage);

    auto execution_cursor = std::make_unique<ExecutionCursor>();

    auto status = getExecutionCursorImpl(*tx, *execution_cursor, total_gas_used,
                                         false, 10, cache);

    return {status, std::move(execution_cursor)};
}

rocksdb::Status ArbCore::advanceExecutionCursor(
    ExecutionCursor& execution_cursor,
    uint256_t max_gas,
    bool go_over_gas,
    ValueCache& cache) {
    auto tx = Transaction::makeTransaction(data_storage);

    return getExecutionCursorImpl(*tx, execution_cursor,
                                  execution_cursor.arb_gas_used + max_gas,
                                  go_over_gas, 10, cache);
}

rocksdb::Status ArbCore::getExecutionCursorImpl(
    Transaction& tx,
    ExecutionCursor& execution_cursor,
    uint256_t total_gas_used,
    bool go_over_gas,
    uint256_t message_group_size,
    ValueCache& cache) {
    if (!execution_cursor.machine) {
        // Existing execution_cursor cannot be used
        execution_cursor.resetCheckpoint();
    }

    auto handle_reorg = true;
    while (handle_reorg) {
        handle_reorg = false;

        auto status =
            executionCursorSetup(tx, execution_cursor, total_gas_used, cache);

        while (true) {
            auto result = executionCursorAddMessages(tx, execution_cursor,
                                                     message_group_size);
            if (!result.status.ok()) {
                return result.status;
            }
            if (!result.data) {
                // Reorg occurred, need to recreate machine
                handle_reorg = true;
                break;
            }

            // Run machine until specified gas is reached
            auto remaining_gas = total_gas_used - execution_cursor.arb_gas_used;
            if (remaining_gas > 0) {
                auto assertion = machine->run(
                    remaining_gas, go_over_gas, execution_cursor.messages,
                    execution_cursor.messages_to_skip, false);
                if (assertion.gasCount == 0) {
                    // Nothing was executed
                    break;
                }
                execution_cursor.messages_to_skip +=
                    assertion.inbox_messages_consumed;
                if (execution_cursor.messages_to_skip > 0) {
                    execution_cursor.inbox_hash =
                        execution_cursor
                            .inbox_hashes[execution_cursor.messages_to_skip -
                                          1];
                }
                execution_cursor.applyAssertion(assertion);
                if (assertion.inbox_messages_consumed !=
                    execution_cursor.messages.size()) {
                    // Not all messages were consumed
                    break;
                }
            } else {
                // Gas reached
                break;
            }
        }
    }

    return rocksdb::Status::OK();
}

rocksdb::Status ArbCore::resolveStagedMessage(Transaction& tx,
                                              value& message,
                                              ValueCache& cache) const {
    if (nonstd::holds_alternative<uint256_t>(message)) {
        auto sequence_number = nonstd::variants::get<uint256_t>(message);
        if (sequence_number >= pending_checkpoint.total_messages_read - 1) {
            // Staged message obsolete
            return rocksdb::Status::NotFound();
        }

        auto message_lookup = getValue(tx, sequence_number, cache);
        if (!message_lookup.status.ok()) {
            // Unable to resolve cursor, no valid message found
            return message_lookup.status;
        }
    }

    return rocksdb::Status::OK();
}

rocksdb::Status ArbCore::executionCursorSetup(Transaction& tx,
                                              ExecutionCursor& execution_cursor,
                                              const uint256_t& total_gas_used,
                                              ValueCache& cache) {
    const std::lock_guard<std::mutex> lock(core_reorg_mutex);
    auto checkpoint_result = getCheckpointUsingGas(tx, total_gas_used, false);
    if (checkpoint_result.status.IsNotFound()) {
        if (!execution_cursor.machine) {
            // Initialize machine to starting state
            auto initial_hash = getInitialMachineHash(tx);
            if (!initial_hash.status.ok()) {
                return initial_hash.status;
            }
            auto result = getMachineStateKeys(tx, initial_hash.data);
            if (!result.status.ok()) {
                return result.status;
            }
            execution_cursor.machine_state_keys = result.data;
            execution_cursor.machine = getMachineUsingStateKeys<Machine>(
                tx, execution_cursor.machine_state_keys, cache);
        }

        // Use execution cursor as is
        return rocksdb::Status::OK();
    } else if (!checkpoint_result.status.ok()) {
        return checkpoint_result.status;
    } else if (execution_cursor.machine &&
               execution_cursor.arb_gas_used >
                   checkpoint_result.data.arb_gas_used) {
        // Execution cursor used more gas than checkpoint so use it if inbox
        // hash valid
        auto result = executionCursorAddMessages(tx, execution_cursor, 0);
        if (!result.status.ok()) {
            return result.status;
        }

        if (result.data) {
            // Execution cursor machine still valid, so use it
            return rocksdb::Status::OK();
        }
    }

    auto staged_message = getValue(
        tx, checkpoint_result.data.machine_state_keys.staged_message_hash,
        cache);
    if (!staged_message.status.ok()) {
        return staged_message.status;
    }

    auto resolve_status = resolveStagedMessage(tx, staged_message.data, cache);
    if (!resolve_status.ok()) {
        // Unable to resolve staged_message, so can't use checkpoint
        return resolve_status;
    }

    // Update execution_cursor with checkpoint
    execution_cursor.resetCheckpoint();
    execution_cursor.setCheckpoint(checkpoint_result.data);
    execution_cursor.machine = getMachineUsingStateKeys<Machine>(
        tx, execution_cursor.machine_state_keys, cache);

    // Replace staged_message with resolved value
    execution_cursor.machine->machine_state.staged_message =
        staged_message.data;

    return rocksdb::Status::OK();
}

ValueResult<bool> ArbCore::executionCursorAddMessages(
    Transaction& tx,
    ExecutionCursor& execution_cursor,
    const uint256_t& orig_message_group_size) {
    const std::lock_guard<std::mutex> lock(core_reorg_mutex);

    auto message_group_size = orig_message_group_size;

    // Check if current machine is obsolete
    if (execution_cursor.total_messages_read > 0) {
        auto stored_result =
            getMessageEntry(tx, execution_cursor.total_messages_read - 1);
        if (!stored_result.status.ok()) {
            return {stored_result.status, false};
        }

        if (execution_cursor.inbox_hash != stored_result.data.inbox_hash) {
            // Obsolete machine, reorg occurred
            return {rocksdb::Status::OK(), false};
        }
    }

    // Delete any pending messages because they may have been affected by reorg
    execution_cursor.first_message_sequence_number +=
        execution_cursor.messages_to_skip;
    execution_cursor.messages.clear();
    execution_cursor.inbox_hashes.clear();
    execution_cursor.messages_to_skip = 0;

    auto current_message_sequence_number =
        execution_cursor.first_message_sequence_number;

    if (current_message_sequence_number >=
        pending_checkpoint.total_messages_read) {
        // Already past core machine, probably reorg
        return {rocksdb::Status::OK(), false};
    }

    if (current_message_sequence_number + message_group_size >=
        pending_checkpoint.total_messages_read) {
        // Don't read past primary machine
        message_group_size = pending_checkpoint.total_messages_read -
                             current_message_sequence_number;
    }

    if (message_group_size == 0) {
        // No messages to read
        return {rocksdb::Status::OK(), true};
    }

    std::vector<unsigned char> message_key;
    marshal_uint256_t(current_message_sequence_number, message_key);
    auto message_key_slice = vecToSlice(message_key);

    auto results = getVectorVectorUsingFamilyAndKey(
        *tx.transaction, data_storage->messageentry_column.get(),
        message_key_slice, intx::narrow_cast<size_t>(message_group_size));
    if (!results.status.ok()) {
        return {results.status, false};
    }

    std::vector<Tuple> messages;
    std::vector<uint256_t> inbox_hashes;
    auto total_size = results.data.size();
    messages.reserve(total_size);
    inbox_hashes.reserve(total_size);
    for (const auto& data : results.data) {
        auto message_entry = extractMessageEntry(0, vecToSlice(data));
        auto inbox_message = extractInboxMessage(message_entry.data);
        messages.push_back(inbox_message.toTuple());
        inbox_hashes.push_back(message_entry.inbox_hash);
    }

    execution_cursor.messages = std::move(messages);
    execution_cursor.inbox_hashes = std::move(inbox_hashes);
    execution_cursor.messages_to_skip = 0;

    return {rocksdb::Status::OK(), true};
}

ValueResult<uint256_t> ArbCore::logInsertedCount() const {
    auto tx = Transaction::makeTransaction(data_storage);

    return logInsertedCountImpl(*tx);
}

ValueResult<uint256_t> ArbCore::logInsertedCountImpl(Transaction& tx) const {
    return getUint256UsingFamilyAndKey(*tx.transaction,
                                       tx.datastorage->state_column.get(),
                                       vecToSlice(log_inserted_key));
}
rocksdb::Status ArbCore::updateLogInsertedCount(Transaction& tx,
                                                const uint256_t& log_index) {
    std::vector<unsigned char> value;
    marshal_uint256_t(log_index, value);

    return tx.transaction->Put(tx.datastorage->state_column.get(),
                               vecToSlice(log_inserted_key), vecToSlice(value));
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

ValueResult<uint256_t> ArbCore::sendInsertedCount() const {
    auto tx = Transaction::makeTransaction(data_storage);

    return sendInsertedCountImpl(*tx);
}

ValueResult<uint256_t> ArbCore::sendInsertedCountImpl(Transaction& tx) const {
    return getUint256UsingFamilyAndKey(*tx.transaction,
                                       tx.datastorage->state_column.get(),
                                       vecToSlice(send_inserted_key));
}

rocksdb::Status ArbCore::updateSendInsertedCount(Transaction& tx,
                                                 const uint256_t& send_index) {
    std::vector<unsigned char> value;
    marshal_uint256_t(send_index, value);

    return tx.transaction->Put(tx.datastorage->state_column.get(),
                               vecToSlice(send_inserted_key),
                               vecToSlice(value));
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

ValueResult<uint256_t> ArbCore::messageEntryInsertedCount() const {
    auto tx = Transaction::makeTransaction(data_storage);

    return messageEntryInsertedCountImpl(*tx);
}

ValueResult<uint256_t> ArbCore::messageEntryInsertedCountImpl(
    Transaction& tx) const {
    return getUint256UsingFamilyAndKey(*tx.transaction,
                                       tx.datastorage->state_column.get(),
                                       vecToSlice(message_entry_inserted_key));
}

rocksdb::Status ArbCore::updateMessageEntryInsertedCount(
    Transaction& tx,
    const uint256_t& message_index) {
    std::vector<unsigned char> value;
    marshal_uint256_t(message_index, value);
    return tx.transaction->Put(tx.datastorage->state_column.get(),
                               vecToSlice(message_entry_inserted_key),
                               vecToSlice(value));
}

ValueResult<uint256_t> ArbCore::messageEntryProcessedCount() const {
    auto tx = Transaction::makeTransaction(data_storage);
    return messageEntryProcessedCountImpl(*tx);
}

ValueResult<uint256_t> ArbCore::messageEntryProcessedCountImpl(
    Transaction& tx) const {
    return getUint256UsingFamilyAndKey(*tx.transaction,
                                       tx.datastorage->state_column.get(),
                                       vecToSlice(message_entry_processed_key));
}

rocksdb::Status ArbCore::updateMessageEntryProcessedCount(
    Transaction& tx,
    const uint256_t& message_index) {
    std::vector<unsigned char> value;
    marshal_uint256_t(message_index, value);
    return tx.transaction->Put(tx.datastorage->state_column.get(),
                               vecToSlice(message_entry_processed_key),
                               vecToSlice(value));
}

// addMessages stores all messages from given block into database.
// The last message in the list is flagged as the last message in the block.
// Returns nonstd::nullopt when caller needs to provide messages from earlier
// block.
nonstd::optional<rocksdb::Status> ArbCore::addMessages(
    const std::vector<std::vector<unsigned char>>& messages,
    const uint256_t& previous_inbox_hash,
    const uint256_t& final_machine_sequence_number,
    const bool last_block_complete,
    ValueCache& cache) {
    auto tx = Transaction::makeTransaction(data_storage);

    // Get the last message sequence number that was added to database
    auto message_count_result = messageEntryInsertedCountImpl(*tx);
    if (!message_count_result.status.ok()) {
        return message_count_result.status;
    }
    nonstd::optional<uint256_t> last_inserted_sequence_number;
    if (message_count_result.data > 0) {
        last_inserted_sequence_number = message_count_result.data - 1;
    }

    auto first_message = extractInboxMessage(messages[0]);

    if (first_message.inbox_sequence_number > 0) {
        if (!last_inserted_sequence_number.has_value() ||
            first_message.inbox_sequence_number >
                *last_inserted_sequence_number + 1) {
            // Not allowed to skip message sequence numbers, ask for older
            // messages
            return nonstd::nullopt;
        }

        // Check that previous_inbox_hash matches hash from previous message
        auto previous_sequence_number = first_message.inbox_sequence_number - 1;
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
    auto current_sequence_number = first_message.inbox_sequence_number;

    if (messages.empty()) {
        // No new messages, just need to truncating obsolete messages
        current_sequence_number = first_message.inbox_sequence_number - 1;
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
        current_sequence_number =
            first_message.inbox_sequence_number + current_message_index;
    }

    if (current_sequence_number <= last_inserted_sequence_number) {
        // Reorg occurred
        const std::lock_guard<std::mutex> lock(core_reorg_mutex);

        if (final_machine_sequence_number >= current_sequence_number) {
            // Machine is running with obsolete messages
            machine->abortMachine();
        }

        auto previous_valid_sequence_number = current_sequence_number - 1;

        // Truncate MessageEntries to last valid message
        updateMessageEntryInsertedCount(*tx,
                                        previous_valid_sequence_number + 1);

        // Reorg checkpoint and everything else
        auto reorg_status = reorgToMessageOrBefore(
            *tx, previous_valid_sequence_number, false, cache);
    }

    InboxMessage next_inbox_message;
    if (current_message_index < messages_count) {
        next_inbox_message =
            extractInboxMessage(messages[current_message_index]);
    }
    while (current_message_index < messages_count) {
        // Encode key
        std::vector<unsigned char> key;
        marshal_uint256_t(current_sequence_number, key);

        auto current_inbox_hash = hash_inbox(current_previous_inbox_hash,
                                             messages[current_message_index]);

        auto current_inbox_message = std::move(next_inbox_message);
        if (current_message_index < messages_count) {
            next_inbox_message =
                extractInboxMessage(messages[current_message_index]);
        } else {
            next_inbox_message = {};
        }

        bool last_message_in_block;
        if (last_block_complete &&
            ((current_message_index == messages_count - 1) ||
             current_inbox_message.block_number !=
                 next_inbox_message.block_number)) {
            last_message_in_block = true;
        } else {
            last_message_in_block = false;
        }

        // Encode message entry
        auto messageEntry = MessageEntry{
            current_sequence_number, current_inbox_hash,
            intx::narrow_cast<uint64_t>(current_inbox_message.block_number),
            last_message_in_block, (messages[current_message_index])};
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
        current_sequence_number += 1;
    }

    updateMessageEntryInsertedCount(*tx, current_sequence_number);

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

bool ArbCore::handleLogsCursorRequested(Transaction& tx,
                                        size_t cursor_index,
                                        ValueCache& cache) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        return false;
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    // Provide requested logs
    logs_cursors[cursor_index].data.clear();
    auto log_inserted_count = logInsertedCountImpl(tx);
    if (!log_inserted_count.status.ok()) {
        logs_cursors[cursor_index].error_string =
            log_inserted_count.status.ToString();
        logs_cursors[cursor_index].status = DataCursor::ERROR;
        return true;
    }

    if (logs_cursors[cursor_index].current_total_count >=
        log_inserted_count.data) {
        // No new data available
        logs_cursors[cursor_index].status = DataCursor::READY;
        return true;
    }
    if (logs_cursors[cursor_index].current_total_count +
            logs_cursors[cursor_index].number_requested >
        log_inserted_count.data) {
        // Too many entries requested
        logs_cursors[cursor_index].number_requested =
            log_inserted_count.data -
            logs_cursors[cursor_index].current_total_count;
    }
    if (logs_cursors[cursor_index].number_requested == 0) {
        logs_cursors[cursor_index].status = DataCursor::READY;
        // No new logs to provide
        return true;
    }
    auto requested_logs =
        getLogs(logs_cursors[cursor_index].current_total_count,
                logs_cursors[cursor_index].number_requested, cache);
    if (!requested_logs.status.ok()) {
        logs_cursors[cursor_index].error_string =
            requested_logs.status.ToString();
        logs_cursors[cursor_index].status = DataCursor::ERROR;
        return true;
    }
    logs_cursors[cursor_index].data.insert(
        logs_cursors[cursor_index].data.end(), requested_logs.data.begin(),
        requested_logs.data.end());
    logs_cursors[cursor_index].status = DataCursor::READY;

    return true;
}

// handleLogsCursorReorg must be called before logs are deleted.
// Note that this function should not update logs_cursors[cursor_index].status
// because it is happening out of line.
// Note that cursor reorg never adds new messages, but might add deleted
// messages.
rocksdb::Status ArbCore::handleLogsCursorReorg(Transaction& tx,
                                               size_t cursor_index,
                                               uint256_t log_count,
                                               ValueCache& cache) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        return rocksdb::Status::InvalidArgument();
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    auto log_inserted_count = logInsertedCountImpl(tx);
    if (!log_inserted_count.status.ok()) {
        std::cerr << "Error getting inserted count in Cursor Reorg: "
                  << log_inserted_count.status.ToString() << "\n";
        return log_inserted_count.status;
    }
    auto log_processed_count = logProcessedCount(tx);
    if (!log_processed_count.status.ok()) {
        std::cerr << "Error getting processed count in Cursor Reorg: "
                  << log_inserted_count.status.ToString() << "\n";
        return log_processed_count.status;
    }

    if (log_count >= log_inserted_count.data) {
        // No reorg needed
        return rocksdb::Status::OK();
    }

    if (log_count < log_processed_count.data) {
        // Need to save logs that will be deleted
        auto logs = getLogsNoLock(tx, log_count - 1,
                                  log_inserted_count.data - log_count, cache);
        if (!logs.status.ok()) {
            std::cerr << "Error getting " << log_count - 1
                      << " logs in Cursor reorg starting at "
                      << log_inserted_count.data - log_count << ": "
                      << log_inserted_count.status.ToString() << "\n";
            return logs.status;
        }
        logs_cursors[cursor_index].deleted_data.insert(
            logs_cursors[cursor_index].deleted_data.end(), logs.data.begin(),
            logs.data.end());
    }

    if (!logs_cursors[cursor_index].data.empty()) {
        if (logs_cursors[cursor_index].current_total_count >= log_count) {
            // Don't save anything
            logs_cursors[cursor_index].data.clear();
        } else if (logs_cursors[cursor_index].current_total_count +
                       logs_cursors[cursor_index].data.size() >
                   log_count) {
            // Only part of the data needs to be removed
            auto offset = intx::narrow_cast<size_t>(
                log_count - logs_cursors[cursor_index].current_total_count);
            logs_cursors[cursor_index].data.erase(
                logs_cursors[cursor_index].data.begin() + offset,
                logs_cursors[cursor_index].data.end());
        }
    }

    return rocksdb::Status::OK();
}

bool ArbCore::logsCursorRequest(size_t cursor_index, uint256_t count) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        return false;
    }

    if (logs_cursors[cursor_index].status != DataCursor::EMPTY) {
        return false;
    }

    logs_cursors[cursor_index].number_requested = count;
    logs_cursors[cursor_index].status = DataCursor::REQUESTED;

    return true;
}

nonstd::optional<std::vector<value>> ArbCore::logsCursorGetLogs(
    size_t cursor_index) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        return nonstd::nullopt;
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    if (logs_cursors[cursor_index].status != DataCursor::READY ||
        !logs_cursors[cursor_index].deleted_data.empty()) {
        return nonstd::nullopt;
    }

    logs_cursors[cursor_index].pending_total_count =
        logs_cursors[cursor_index].current_total_count +
        logs_cursors[cursor_index].data.size();

    std::vector<value> logs{std::move(logs_cursors[cursor_index].data)};
    logs_cursors[cursor_index].data.clear();

    return logs;
}

nonstd::optional<std::vector<value>> ArbCore::logsCursorGetDeletedLogs(
    size_t cursor_index) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        return nonstd::nullopt;
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    if (logs_cursors[cursor_index].status != DataCursor::READY ||
        logs_cursors[cursor_index].deleted_data.empty()) {
        return nonstd::nullopt;
    }

    std::vector<value> logs{std::move(logs_cursors[cursor_index].deleted_data)};
    logs_cursors[cursor_index].data.clear();

    return logs;
}

bool ArbCore::logsCursorConfirmReceived(size_t cursor_index) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        return false;
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    if (logs_cursors[cursor_index].status != DataCursor::READY) {
        logs_cursors[cursor_index].error_string =
            "logsCursorConfirmReceived called at wrong state";
        std::cerr << "logsCursorConfirmReceived called at wrong state: "
                  << logs_cursors[cursor_index].status << "\n";
        logs_cursors[cursor_index].status = DataCursor::ERROR;
        return false;
    }

    if (!logs_cursors[cursor_index].data.empty() ||
        !logs_cursors[cursor_index].deleted_data.empty()) {
        // Still have logs to get
        return false;
    }

    logs_cursors[cursor_index].current_total_count =
        logs_cursors[cursor_index].pending_total_count;
    logs_cursors[cursor_index].status = DataCursor::EMPTY;

    return true;
}

bool ArbCore::logsCursorCheckError(size_t cursor_index) const {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        return false;
    }

    return logs_cursors[cursor_index].status == DataCursor::ERROR;
}

std::string ArbCore::logsCursorClearError(size_t cursor_index) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        return "Invalid logsCursor index";
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    if (logs_cursors[cursor_index].status != DataCursor::ERROR) {
        std::cerr << "logsCursorClearError called when status not ERROR"
                  << std::endl;
        return "logsCursorClearError called when sttaus not ERROR";
    }

    auto str = logs_cursors[cursor_index].error_string;
    logs_cursors[cursor_index].error_string.clear();
    logs_cursors[cursor_index].data.clear();
    logs_cursors[cursor_index].deleted_data.clear();
    logs_cursors[cursor_index].status = DataCursor::EMPTY;

    return str;
}
