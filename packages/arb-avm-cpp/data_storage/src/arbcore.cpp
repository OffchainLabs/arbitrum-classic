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

#include <data_storage/arbcore.hpp>

#include "value/utils.hpp"

#include <avm/inboxmessage.hpp>
#include <avm/machinethread.hpp>
#include <data_storage/aggregator.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/readwritetransaction.hpp>
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
constexpr auto logscursor_current_prefix = std::array<char, 1>{-66};

constexpr auto sideload_cache_size = 20;
}  // namespace

ArbCore::ArbCore(std::shared_ptr<DataStorage> data_storage_)
    : data_storage(std::move(data_storage_)),
      code(std::make_shared<Code>(
          getNextSegmentID(*makeReadOnlyTransaction()))) {
    if (logs_cursors.size() > 255) {
        throw std::runtime_error("Too many logscursors");
    }
    for (size_t i = 0; i < logs_cursors.size(); i++) {
        logs_cursors[i].current_total_key.insert(
            logs_cursors[i].current_total_key.end(),
            logscursor_current_prefix.begin(), logscursor_current_prefix.end());
        logs_cursors[i].current_total_key.emplace_back(i);
    }
}

ValueResult<MessageEntry> ArbCore::getMessageEntry(
    const ReadTransaction& tx,
    uint256_t message_sequence_number) const {
    std::vector<unsigned char> previous_key;
    marshal_uint256_t(message_sequence_number, previous_key);

    auto messages_inserted = messageEntryInsertedCountImpl(tx);
    if (!messages_inserted.status.ok()) {
        return {messages_inserted.status, {}};
    }

    if (message_sequence_number > messages_inserted.data) {
        // Don't allow stale entries to be used
        return {rocksdb::Status::NotFound(), {}};
    }

    auto result = tx.messageEntryGetVector(vecToSlice(previous_key));
    if (!result.status.ok()) {
        return {result.status, {}};
    }

    auto parsed_state =
        extractMessageEntry(message_sequence_number, vecToSlice(result.data));

    return {result.status, std::move(parsed_state)};
}

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
        return "";
    }

    message_data_status = MESSAGES_EMPTY;
    auto str = core_error_string;
    core_error_string.clear();

    return str;
}

std::optional<std::string> ArbCore::machineClearError() {
    if (!machine_error) {
        return std::nullopt;
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
bool ArbCore::deliverMessages(
    std::vector<std::vector<unsigned char>> messages,
    const uint256_t& previous_inbox_acc,
    bool last_block_complete,
    const std::optional<uint256_t>& reorg_message_count) {
    if (message_data_status != MESSAGES_EMPTY) {
        return false;
    }

    message_data.messages = std::move(messages);
    message_data.previous_inbox_acc = previous_inbox_acc;
    message_data.last_block_complete = last_block_complete;
    message_data.reorg_message_count = reorg_message_count;

    message_data_status = MESSAGES_READY;

    return true;
}

std::unique_ptr<ReadTransaction> ArbCore::makeReadOnlyTransaction() {
    return ReadTransaction::makeReadOnlyTransaction(data_storage);
}

std::unique_ptr<const ReadTransaction> ArbCore::makeConstReadOnlyTransaction()
    const {
    return ReadTransaction::makeReadOnlyTransaction(data_storage);
}

std::unique_ptr<ReadWriteTransaction> ArbCore::makeReadWriteTransaction() {
    return ReadWriteTransaction::makeReadWriteTransaction(data_storage);
}

rocksdb::Status ArbCore::initialize(const LoadedExecutable& executable) {
    auto tx = makeReadWriteTransaction();

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
    } else {
        // Need to initialize database from scratch
        auto res = saveMachine(*tx, *machine);
        if (!res.status.ok()) {
            std::cerr << "failed to save initial machine: "
                      << res.status.ToString() << std::endl;
            return res.status;
        }

        std::vector<unsigned char> value_data;
        auto machine_hash = machine->hash();
        if (!machine_hash) {
            std::cerr << "failed to compute initial machine hash" << std::endl;
            return rocksdb::Status::Corruption();
        }
        marshal_uint256_t(*machine_hash, value_data);
        auto s = tx->statePut(vecToSlice(initial_machine_hash_key),
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

        for (size_t i = 0; i < logs_cursors.size(); i++) {
            status = logsCursorSaveCurrentTotalCount(*tx, i, 0);
            if (!status.ok()) {
                throw std::runtime_error(
                    "failed to initialize logscursor counts");
            }
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
    auto tx = makeConstReadOnlyTransaction();
    std::string initial_raw;
    auto s = tx->stateGet(vecToSlice(initial_machine_hash_key), &initial_raw);
    return s.ok();
}

ValueResult<uint256_t> ArbCore::getInitialMachineHash(ReadTransaction& tx) {
    std::string initial_raw;
    auto s = tx.stateGet(vecToSlice(initial_machine_hash_key), &initial_raw);
    if (!s.ok()) {
        return {s, 0};
    }

    return {rocksdb::Status::OK(),
            intx::be::unsafe::load<uint256_t>(
                reinterpret_cast<const unsigned char*>(initial_raw.data()))};
}

template <class T>
std::unique_ptr<T> ArbCore::getInitialMachineImpl(ReadTransaction& tx,
                                                  ValueCache& value_cache) {
    auto machine_hash = getInitialMachineHash(tx);
    return getMachine<T>(machine_hash.data, value_cache);
}

template std::unique_ptr<Machine> ArbCore::getInitialMachineImpl(
    ReadTransaction& tx,
    ValueCache& value_cache);
template std::unique_ptr<MachineThread> ArbCore::getInitialMachineImpl(
    ReadTransaction& tx,
    ValueCache& value_cache);

template <class T>
std::unique_ptr<T> ArbCore::getInitialMachine(ValueCache& value_cache) {
    auto tx = makeReadOnlyTransaction();
    tx->enterReadSnapshot();
    return getInitialMachineImpl<T>(*tx, value_cache);
}

template std::unique_ptr<Machine> ArbCore::getInitialMachine(ValueCache&);
template std::unique_ptr<MachineThread> ArbCore::getInitialMachine(ValueCache&);

template <class T>
std::unique_ptr<T> ArbCore::getMachineImpl(ReadTransaction& tx,
                                           uint256_t machineHash,
                                           ValueCache& value_cache) {
    auto results = getMachineStateKeys(tx, machineHash);
    if (std::holds_alternative<rocksdb::Status>(results)) {
        throw std::runtime_error("failed to load machine state");
    }

    return getMachineUsingStateKeys<T>(
        tx, std::get<CountedData<MachineStateKeys>>(results).data, value_cache);
}

template std::unique_ptr<Machine> ArbCore::getMachineImpl(
    ReadTransaction& tx,
    uint256_t machineHash,
    ValueCache& value_cache);
template std::unique_ptr<MachineThread> ArbCore::getMachineImpl(
    ReadTransaction& tx,
    uint256_t machineHash,
    ValueCache& value_cache);

template <class T>
std::unique_ptr<T> ArbCore::getMachine(uint256_t machineHash,
                                       ValueCache& value_cache) {
    auto tx = makeReadOnlyTransaction();
    tx->enterReadSnapshot();
    return getMachineImpl<T>(*tx, machineHash, value_cache);
}

template std::unique_ptr<Machine> ArbCore::getMachine(uint256_t, ValueCache&);
template std::unique_ptr<MachineThread> ArbCore::getMachine(uint256_t,
                                                            ValueCache&);

rocksdb::Status ArbCore::saveCheckpoint(ReadWriteTransaction& tx) {
    auto status = saveMachineState(tx, *machine);
    if (!status.ok()) {
        return status;
    }

    std::vector<unsigned char> key;
    marshal_uint256_t(machine->machine_state.output.arb_gas_used, key);
    auto key_slice = vecToSlice(key);
    std::vector<unsigned char> value_vec;
    serializeMachineStateKeys(MachineStateKeys{machine->machine_state},
                              value_vec);
    auto put_status = tx.checkpointPut(key_slice, vecToSlice(value_vec));
    if (!put_status.ok()) {
        std::cerr << "ArbCore unable to save checkpoint : "
                  << put_status.ToString() << "\n";
        return put_status;
    }

    return rocksdb::Status::OK();
}

rocksdb::Status ArbCore::saveAssertion(ReadWriteTransaction& tx,
                                       const Assertion& assertion,
                                       const uint256_t arb_gas_used) {
    auto status = saveLogs(tx, assertion.logs);
    if (!status.ok()) {
        return status;
    }

    status = saveSends(tx, assertion.sends);
    if (!status.ok()) {
        return status;
    }

    if (assertion.sideloadBlockNumber) {
        status = saveSideloadPosition(tx, *assertion.sideloadBlockNumber,
                                      arb_gas_used);
        if (!status.ok()) {
            return status;
        }
    }

    return rocksdb::Status::OK();
}

// reorgToMessageOrBefore resets the checkpoint and database entries
// such that machine state is at or before the requested message. cleaning
// up old references as needed.
// If use_latest is true, message_sequence_number is ignored and the latest
// checkpoint is used.
rocksdb::Status ArbCore::reorgToMessageOrBefore(
    ReadWriteTransaction& tx,
    const uint256_t& message_sequence_number,
    bool use_latest,
    ValueCache& cache) {
    auto it = tx.checkpointGetIterator();

    // Find first checkpoint to delete
    it->SeekToLast();
    if (!it->status().ok()) {
        return it->status();
    }

    // Delete each checkpoint until at or below message_sequence_number
    auto setup = [&]() -> std::variant<MachineStateKeys, rocksdb::Status> {
        if (use_latest) {
            std::vector<unsigned char> checkpoint_vector(
                it->value().data(), it->value().data() + it->value().size());
            return extractMachineStateKeys(checkpoint_vector.begin(),
                                           checkpoint_vector.end());
        } else {
            while (it->Valid()) {
                std::vector<unsigned char> checkpoint_vector(
                    it->value().data(),
                    it->value().data() + it->value().size());
                auto checkpoint = extractMachineStateKeys(
                    checkpoint_vector.begin(), checkpoint_vector.end());
                if (checkpoint.getTotalMessagesRead() == 0 ||
                    message_sequence_number >=
                        checkpoint.getTotalMessagesRead() - 1) {
                    // Good checkpoint
                    return checkpoint;
                }

                // Obsolete checkpoint, need to delete referenced machine
                deleteMachineState(tx, checkpoint);

                // Delete checkpoint to make sure it isn't used later
                tx.checkpointDelete(it->key());

                it->Prev();
                if (!it->status().ok()) {
                    return it->status();
                }
            }
            return it->status();
        }
    }();

    it = nullptr;
    if (std::holds_alternative<rocksdb::Status>(setup)) {
        return std::get<rocksdb::Status>(setup);
    }
    MachineStateKeys checkpoint = std::get<MachineStateKeys>(std::move(setup));

    uint256_t next_sideload_block_number = 0;
    if (checkpoint.output.last_sideload) {
        next_sideload_block_number = *checkpoint.output.last_sideload + 1;
    }

    auto status = deleteSideloadsStartingAt(tx, next_sideload_block_number);
    if (!status.ok()) {
        return status;
    }

    auto log_inserted_count = logInsertedCountImpl(tx);
    if (!log_inserted_count.status.ok()) {
        std::cerr << "Error getting inserted count in Cursor Reorg: "
                  << log_inserted_count.status.ToString() << "\n";
        return log_inserted_count.status;
    }

    if (checkpoint.output.log_count < log_inserted_count.data) {
        // Update log cursors, must be called before logs are deleted
        for (size_t i = 0; i < logs_cursors.size(); i++) {
            status = handleLogsCursorReorg(tx, i, checkpoint.output.log_count,
                                           cache);
            if (!status.ok()) {
                return status;
            }
        }
    }

    // Delete logs individually to handle reference counts
    auto optional_status =
        deleteLogsStartingAt(tx, checkpoint.output.log_count);
    if (optional_status && !optional_status->ok()) {
        return *optional_status;
    }

    status = updateLogInsertedCount(tx, checkpoint.output.log_count);
    if (!status.ok()) {
        return status;
    }

    status = updateSendInsertedCount(tx, checkpoint.output.send_count);
    if (!status.ok()) {
        return status;
    }

    // Machine was executing obsolete messages so restore machine
    // from last checkpoint
    machine->abortMachine();

    machine = getMachineUsingStateKeys<MachineThread>(tx, checkpoint, cache);

    return rocksdb::Status::OK();
}

std::variant<rocksdb::Status, MachineStateKeys> ArbCore::getCheckpoint(
    ReadTransaction& tx,
    const uint256_t& arb_gas_used) const {
    std::vector<unsigned char> key;
    marshal_uint256_t(arb_gas_used, key);

    auto result = tx.checkpointGetVector(vecToSlice(key));
    if (!result.status.ok()) {
        return result.status;
    }
    return extractMachineStateKeys(result.data.begin(), result.data.end());
}

bool ArbCore::isCheckpointsEmpty(ReadTransaction& tx) const {
    auto it = std::unique_ptr<rocksdb::Iterator>(tx.checkpointGetIterator());
    it->SeekToLast();
    return !it->Valid();
}

uint256_t ArbCore::maxCheckpointGas() {
    auto tx = makeReadOnlyTransaction();
    auto it = tx->checkpointGetIterator();
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
std::variant<rocksdb::Status, MachineStateKeys> ArbCore::getCheckpointUsingGas(
    ReadTransaction& tx,
    const uint256_t& total_gas,
    bool after_gas) {
    auto it = tx.checkpointGetIterator();
    std::vector<unsigned char> key;
    marshal_uint256_t(total_gas, key);
    auto key_slice = vecToSlice(key);
    it->SeekForPrev(key_slice);
    if (!it->Valid()) {
        if (!it->status().ok()) {
            return it->status();
        }
        return rocksdb::Status::NotFound();
    }
    if (after_gas) {
        it->Next();
        if (!it->status().ok()) {
            return it->status();
        }
        if (!it->Valid()) {
            return rocksdb::Status::NotFound();
        }
    }
    if (!it->status().ok()) {
        return it->status();
    }

    std::vector<unsigned char> saved_value(
        it->value().data(), it->value().data() + it->value().size());
    return extractMachineStateKeys(saved_value.begin(), saved_value.end());
}

template <class T>
std::unique_ptr<T> ArbCore::getMachineUsingStateKeys(
    const ReadTransaction& transaction,
    const MachineStateKeys& state_data,
    ValueCache& value_cache) const {
    std::set<uint64_t> segment_ids;

    auto static_results = ::getValueImpl(transaction, state_data.static_hash,
                                         segment_ids, value_cache);

    if (std::holds_alternative<rocksdb::Status>(static_results)) {
        throw std::runtime_error("failed loaded core machine static");
    }

    auto register_results = ::getValueImpl(
        transaction, state_data.register_hash, segment_ids, value_cache);
    if (std::holds_alternative<rocksdb::Status>(register_results)) {
        throw std::runtime_error("failed to load machine register");
    }

    auto stack_results = ::getValueImpl(transaction, state_data.datastack_hash,
                                        segment_ids, value_cache);
    if (std::holds_alternative<rocksdb::Status>(stack_results) ||
        !std::holds_alternative<Tuple>(
            std::get<CountedData<value>>(stack_results).data)) {
        throw std::runtime_error("failed to load machine stack");
    }

    auto auxstack_results = ::getValueImpl(
        transaction, state_data.auxstack_hash, segment_ids, value_cache);
    if (std::holds_alternative<rocksdb::Status>(auxstack_results)) {
        throw std::runtime_error("failed to load machine auxstack");
    }
    if (!std::holds_alternative<Tuple>(
            std::get<CountedData<value>>(auxstack_results).data)) {
        throw std::runtime_error(
            "failed to load machine auxstack because of format error");
    }

    segment_ids.insert(state_data.pc.pc.segment);
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
    };
    auto state = MachineState{
        code,
        std::move(std::get<CountedData<value>>(register_results).data),
        std::move(std::get<CountedData<value>>(static_results).data),
        Datastack(
            std::get<Tuple>(std::get<CountedData<value>>(stack_results).data)),
        Datastack(std::get<Tuple>(
            std::get<CountedData<value>>(auxstack_results).data)),
        state_data.arb_gas_remaining,
        state_data.status,
        state_data.pc.pc,
        state_data.err_pc,
        std::move(state_data.staged_message),
        std::move(state_data.output)};

    return std::make_unique<T>(state);
}

template std::unique_ptr<Machine> ArbCore::getMachineUsingStateKeys(
    const ReadTransaction& transaction,
    const MachineStateKeys& state_data,
    ValueCache& value_cache) const;
template std::unique_ptr<MachineThread> ArbCore::getMachineUsingStateKeys(
    const ReadTransaction& transaction,
    const MachineStateKeys& state_data,
    ValueCache& value_cache) const;

// operator() runs the main thread for ArbCore.  It is responsible for adding
// messages to the queue, starting machine thread when needed and collecting
// results of machine thread.
// This thread will update `delivering_messages` if and only if
// `delivering_messages` is set to MESSAGES_READY
void ArbCore::operator()() {
    ValueCache cache;
    uint256_t message_count_in_machine = 0;
    MachineExecutionConfig execConfig;
    execConfig.stop_on_sideload = true;
    uint256_t max_message_batch_size = 10;

    while (!arbcore_abort) {
        if (message_data_status == MESSAGES_READY) {
            // Reorg might occur while adding messages
            auto add_status = addMessages(
                message_data.messages, message_data.last_block_complete,
                message_data.previous_inbox_acc, message_count_in_machine,
                message_data.reorg_message_count, cache);
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

        if (machine->status() == MachineThread::MACHINE_SUCCESS) {
            auto tx = makeReadWriteTransaction();

            auto last_assertion = machine->nextAssertion();

            // Save logs and sends
            auto status =
                saveAssertion(*tx, last_assertion,
                              machine->machine_state.output.arb_gas_used);
            if (!status.ok()) {
                core_error_string = status.ToString();
                std::cerr << "ArbCore assertion saving failed: "
                          << core_error_string << "\n";
                break;
            }

            // Cache pre-sideload machines
            if (last_assertion.sideloadBlockNumber) {
                auto block = *last_assertion.sideloadBlockNumber;
                std::unique_lock<std::shared_mutex> lock(sideload_cache_mutex);
                sideload_cache[block] = std::make_unique<Machine>(*machine);
                // Remove any sideload_cache entries that are either more
                // than sideload_cache_size blocks old, or in the future
                // (meaning they've been reorg'd out).
                auto it = sideload_cache.begin();
                while (it != sideload_cache.end()) {
                    // Note: we check if block > sideload_cache_size here
                    // to prevent an underflow in the following check.
                    if ((block > sideload_cache_size &&
                         it->first < block - sideload_cache_size) ||
                        it->first > block) {
                        it = sideload_cache.erase(it);
                    } else {
                        it++;
                    }
                }

                // Save checkpoint for every sideload
                status = saveCheckpoint(*tx);
                if (!status.ok()) {
                    core_error_string = status.ToString();
                    std::cerr << "ArbCore checkpoint saving failed: "
                              << core_error_string << "\n";
                    break;
                }

                // TODO Decide how often to clear ValueCache
                // (only clear cache when machine thread stopped)
                cache.clear();

                // Machine was stopped to save sideload, update execConfig
                // and start machine back up where it stopped
                auto machine_success = machine->continueRunningMachine();
                if (!machine_success) {
                    core_error_string = "Error starting machine thread";
                    machine_error = true;
                    std::cerr << "ArbCore error: " << core_error_string << "\n";
                    break;
                }
            }

            status = tx->commit();
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
            auto tx = makeReadOnlyTransaction();
            auto messages_count = messageEntryInsertedCountImpl(*tx);
            if (!messages_count.status.ok()) {
                core_error_string = messages_count.status.ToString();
                machine_error = true;
                std::cerr << "ArbCore message count fetching failed: "
                          << core_error_string << "\n";
                break;
            }
            auto total_messages_read =
                machine->machine_state.getTotalMessagesRead();

            std::vector<std::vector<unsigned char>> messages;
            if (messages_count.data > total_messages_read) {
                // New messages to process
                auto message_batch_size = max_message_batch_size;
                if (message_batch_size > messages_count.data) {
                    message_batch_size = messages_count.data;
                }
                auto next_messages_result = getMessagesImpl(
                    *tx, total_messages_read, message_batch_size);
                if (!next_messages_result.status.ok()) {
                    core_error_string = next_messages_result.status.ToString();
                    machine_error = true;
                    std::cerr << "ArbCore failed getting message entry: "
                              << core_error_string << "\n";
                    break;
                }
                messages.insert(messages.end(),
                                next_messages_result.data.first.begin(),
                                next_messages_result.data.first.end());
                if (next_messages_result.data.second) {
                    execConfig.next_block_height =
                        *next_messages_result.data.second;
                } else {
                    execConfig.next_block_height = std::nullopt;
                }
            } else {
                execConfig.next_block_height = std::nullopt;
            }

            bool resolved_staged = false;
            if (machine->machine_state.stagedMessageUnresolved()) {
                // Resolve staged message if possible.  If message not found,
                // machine will just be blocked
                auto sequence_number =
                    machine->machine_state.getTotalMessagesRead() - 1;
                auto message_lookup = getMessageEntry(*tx, sequence_number);
                if (message_lookup.status.ok()) {
                    auto inbox_message =
                        extractInboxMessage(message_lookup.data.data);
                    machine->machine_state.staged_message = inbox_message;
                    if (messages.empty() &&
                        message_lookup.data.last_message_in_block) {
                        execConfig.next_block_height =
                            message_lookup.data.block_height + 1;
                    }
                }
                if (!message_lookup.status.IsNotFound() &&
                    !message_lookup.status.ok()) {
                    core_error_string = "error resolving staged message";
                    machine_error = true;
                    std::cerr << "ArbCore error: " << core_error_string << ": "
                              << message_lookup.status.ToString() << "\n";
                    break;
                }
                if (message_lookup.status.ok()) {
                    resolved_staged = true;
                }
            }

            if (!messages.empty() || resolved_staged) {
                message_count_in_machine =
                    total_messages_read + messages.size();
                execConfig.setInboxMessagesFromBytes(messages);

                auto status = machine->runMachine(execConfig);
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
                auto tx = makeReadOnlyTransaction();
                handleLogsCursorRequested(*tx, i, cache);
            }
        }

        if (!machineIdle() || message_data_status != MESSAGES_READY) {
            // Machine is already running or new messages, so sleep for a short
            // while
            std::this_thread::sleep_for(std::chrono::milliseconds(5));
        }
    }

    // Error occurred, make sure machine stops cleanly
    machine->abortMachine();
}

rocksdb::Status ArbCore::saveLogs(ReadWriteTransaction& tx,
                                  const std::vector<value>& vals) {
    if (vals.empty()) {
        return rocksdb::Status::OK();
    }
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

        auto status = tx.logPut(key_slice, value_hash_slice);
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
    auto tx = makeReadOnlyTransaction();
    tx->enterReadSnapshot();

    return getLogsNoLock(*tx, index, count, valueCache);
}

ValueResult<std::vector<value>> ArbCore::getLogsNoLock(ReadTransaction& tx,
                                                       uint256_t index,
                                                       uint256_t count,
                                                       ValueCache& valueCache) {
    if (count == 0) {
        return {rocksdb::Status::OK(), {}};
    }

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

    auto hash_result = tx.logGetUint256Vector(vecToSlice(key),
                                              intx::narrow_cast<size_t>(count));
    if (!hash_result.status.ok()) {
        return {hash_result.status, {}};
    }

    std::vector<value> logs;
    for (const auto& hash : hash_result.data) {
        auto val_result = getValue(tx, hash, valueCache);
        if (std::holds_alternative<rocksdb::Status>(val_result)) {
            return {std::get<rocksdb::Status>(val_result), {}};
        }
        logs.push_back(
            std::move(std::get<CountedData<value>>(val_result).data));
    }

    return {rocksdb::Status::OK(), std::move(logs)};
}

rocksdb::Status ArbCore::saveSends(
    ReadWriteTransaction& tx,
    const std::vector<std::vector<unsigned char>>& sends) {
    if (sends.empty()) {
        return rocksdb::Status::OK();
    }
    auto send_result = sendInsertedCountImpl(tx);
    if (!send_result.status.ok()) {
        return send_result.status;
    }

    auto send_count = send_result.data;
    for (const auto& send : sends) {
        std::vector<unsigned char> key;
        marshal_uint256_t(send_count, key);
        auto key_slice = vecToSlice(key);

        auto status = tx.sendPut(key_slice, vecToSlice(send));
        if (!status.ok()) {
            return status;
        }
        send_count += 1;
    }

    return updateSendInsertedCount(tx, send_count);
}

ValueResult<std::vector<std::vector<unsigned char>>> ArbCore::getMessages(
    uint256_t index,
    uint256_t count) const {
    auto tx = makeConstReadOnlyTransaction();

    auto result = getMessagesImpl(*tx, index, count);

    return {result.status, result.data.first};
}

ValueResult<std::pair<std::vector<std::vector<unsigned char>>,
                      std::optional<uint256_t>>>
ArbCore::getMessagesImpl(const ReadTransaction& tx,
                         uint256_t index,
                         uint256_t count) const {
    if (count == 0) {
        return {rocksdb::Status::OK(), {}};
    }

    // Check if attempting to get entries past current valid messages
    auto message_count = messageEntryInsertedCountImpl(tx);
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

    auto results = tx.messageEntryGetVectorVector(
        key_slice, intx::narrow_cast<size_t>(count));
    if (!results.status.ok()) {
        return {results.status, {}};
    }

    std::vector<std::vector<unsigned char>> messages;
    messages.reserve(results.data.size());
    std::optional<uint256_t> next_block_height;
    auto last_index = results.data.size() - 1;
    for (size_t i = 0; i <= last_index; i++) {
        auto message_entry =
            extractMessageEntry(0, vecToSlice(results.data[i]));

        messages.push_back(message_entry.data);
        if (i == last_index && message_entry.last_message_in_block) {
            next_block_height = message_entry.block_height + 1;
        }
    }

    return {rocksdb::Status::OK(), {std::move(messages), next_block_height}};
}

ValueResult<std::vector<std::vector<unsigned char>>> ArbCore::getSends(
    uint256_t index,
    uint256_t count) const {
    auto tx = makeConstReadOnlyTransaction();

    if (count == 0) {
        return {rocksdb::Status::OK(), {}};
    }

    // Check if attempting to get entries past current valid sends
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

    return tx->sendGetVectorVector(key_slice, intx::narrow_cast<size_t>(count));
}

ValueResult<uint256_t> ArbCore::getInboxAcc(uint256_t index) {
    auto tx = makeReadOnlyTransaction();
    tx->enterReadSnapshot();

    auto result = getMessageEntry(*tx, index);
    if (!result.status.ok()) {
        return {result.status, 0};
    }

    return {rocksdb::Status::OK(), result.data.inbox_acc};
}

ValueResult<std::pair<uint256_t, uint256_t>> ArbCore::getInboxAccPair(
    uint256_t index1,
    uint256_t index2) {
    auto tx = makeReadOnlyTransaction();
    tx->enterReadSnapshot();

    auto result1 = getMessageEntry(*tx, index1);
    if (!result1.status.ok()) {
        return {result1.status, {0, 0}};
    }

    auto result2 = getMessageEntry(*tx, index2);
    if (!result2.status.ok()) {
        return {result2.status, {0, 0}};
    }

    return {rocksdb::Status::OK(),
            {result1.data.inbox_acc, result2.data.inbox_acc}};
}

ValueResult<uint256_t> ArbCore::getSendAcc(uint256_t start_acc_hash,
                                           uint256_t start_index,
                                           uint256_t count) const {
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
    auto tx = makeReadOnlyTransaction();
    tx->enterReadSnapshot();

    auto initial_hash = getInitialMachineHash(*tx);
    if (!initial_hash.status.ok()) {
        return {initial_hash.status, nullptr};
    }
    auto result = getMachineStateKeys(*tx, initial_hash.data);
    if (std::holds_alternative<rocksdb::Status>(result)) {
        return {std::get<rocksdb::Status>(result), nullptr};
    }

    auto execution_cursor = std::make_unique<ExecutionCursor>(
        std::get<CountedData<MachineStateKeys>>(result).data);

    auto status = getExecutionCursorImpl(*tx, *execution_cursor, total_gas_used,
                                         false, 10, cache, false);

    return {status, std::move(execution_cursor)};
}

rocksdb::Status ArbCore::advanceExecutionCursor(
    ExecutionCursor& execution_cursor,
    uint256_t max_gas,
    bool go_over_gas,
    ValueCache& cache) {
    auto tx = makeReadOnlyTransaction();

    return getExecutionCursorImpl(
        *tx, execution_cursor,
        execution_cursor.getOutput().arb_gas_used + max_gas, go_over_gas, 10,
        cache, true);
}

MachineState& resolveExecutionVariant(std::unique_ptr<Machine>& mach) {
    return mach->machine_state;
}

MachineStateKeys& resolveExecutionVariant(MachineStateKeys& mach) {
    return mach;
}

std::unique_ptr<Machine>& ArbCore::resolveExecutionCursorMachine(
    const ReadTransaction& tx,
    ExecutionCursor& execution_cursor,
    ValueCache& cache) const {
    if (std::holds_alternative<MachineStateKeys>(execution_cursor.machine)) {
        auto machine_state_keys =
            std::get<MachineStateKeys>(execution_cursor.machine);
        execution_cursor.machine =
            getMachineUsingStateKeys<Machine>(tx, machine_state_keys, cache);
    }
    return std::get<std::unique_ptr<Machine>>(execution_cursor.machine);
}

std::unique_ptr<Machine> ArbCore::takeExecutionCursorMachineImpl(
    const ReadTransaction& tx,
    ExecutionCursor& execution_cursor,
    ValueCache& cache) const {
    auto mach =
        std::move(resolveExecutionCursorMachine(tx, execution_cursor, cache));
    execution_cursor.machine = MachineStateKeys{mach->machine_state};
    return mach;
}

std::unique_ptr<Machine> ArbCore::takeExecutionCursorMachine(
    ExecutionCursor& execution_cursor,
    ValueCache& cache) const {
    auto tx = makeConstReadOnlyTransaction();
    return takeExecutionCursorMachineImpl(*tx, execution_cursor, cache);
}

rocksdb::Status ArbCore::getExecutionCursorImpl(
    ReadTransaction& tx,
    ExecutionCursor& execution_cursor,
    uint256_t total_gas_used,
    bool go_over_gas,
    uint256_t message_group_size,
    ValueCache& cache,
    bool possible_reorg) {
    auto handle_reorg = true;
    while (handle_reorg) {
        handle_reorg = false;

        auto status =
            executionCursorSetup(tx, execution_cursor, total_gas_used);
        if (!status.ok()) {
            return status;
        }

        while (true) {
            auto get_messages_result = executionCursorGetMessages(
                tx, execution_cursor, message_group_size);
            if (!get_messages_result.status.ok()) {
                return get_messages_result.status;
            }
            if (!get_messages_result.data.first) {
                // Reorg occurred, need to recreate machine
                handle_reorg = true;
                if (!possible_reorg) {
                    std::cerr
                        << "Warning: Unexpected execution cursor reorg detected"
                        << std::endl;
                }
                possible_reorg = false;
                break;
            }

            // Run machine until specified gas is reached
            auto remaining_gas =
                total_gas_used - execution_cursor.getOutput().arb_gas_used;
            if (remaining_gas > 0) {
                MachineExecutionConfig execConfig;
                execConfig.max_gas = total_gas_used;
                execConfig.go_over_gas = go_over_gas;
                execConfig.inbox_messages =
                    std::move(get_messages_result.data.second);

                // Resolve staged message if possible.
                // If placeholder message not found, machine will just be
                // blocked
                auto resolve_status = std::visit(
                    [&](auto& machine) {
                        return resolveStagedMessage(
                            tx, resolveExecutionVariant(machine));
                    },
                    execution_cursor.machine);
                if (!resolve_status.IsNotFound() && !resolve_status.ok()) {
                    core_error_string = "error resolving staged message";
                    machine_error = true;
                    std::cerr << "ArbCore error: " << core_error_string << ": "
                              << resolve_status.ToString() << "\n";
                    return resolve_status;
                }
                auto& mach =
                    resolveExecutionCursorMachine(tx, execution_cursor, cache);
                auto assertion = mach->run(execConfig);
                if (assertion.gasCount == 0) {
                    // Nothing was executed
                    break;
                }

                if (total_gas_used <=
                    execution_cursor.getOutput().arb_gas_used) {
                    // Gas reached
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

template <typename T>
rocksdb::Status ArbCore::resolveStagedMessage(const ReadTransaction& tx,
                                              T& machine_state) {
    if (machine_state.stagedMessageUnresolved()) {
        auto sequence_number = machine_state.getTotalMessagesRead();
        auto message_lookup = getMessageEntry(tx, sequence_number - 1);
        if (!message_lookup.status.ok()) {
            // Unable to resolve cursor, no valid message found
            return message_lookup.status;
        }
        machine_state.staged_message =
            extractInboxMessage(message_lookup.data.data);
    }

    return rocksdb::Status::OK();
}

rocksdb::Status ArbCore::executionCursorSetup(ReadTransaction& tx,
                                              ExecutionCursor& execution_cursor,
                                              const uint256_t& total_gas_used,
                                              bool is_for_sideload) {
    auto target_gas_used = total_gas_used;
    while (true) {
        const std::lock_guard<std::mutex> lock(core_reorg_mutex);
        auto checkpoint_result =
            getCheckpointUsingGas(tx, target_gas_used, false);

        if (std::holds_alternative<rocksdb::Status>(checkpoint_result)) {
            auto status = std::get<rocksdb::Status>(checkpoint_result);
            if (!status.IsNotFound()) {
                return status;
            }

            // Initialize machine to starting state
            auto initial_hash = getInitialMachineHash(tx);
            if (!initial_hash.status.ok()) {
                return initial_hash.status;
            }
            auto result = getMachineStateKeys(tx, initial_hash.data);
            if (std::holds_alternative<rocksdb::Status>(result)) {
                return std::get<rocksdb::Status>(result);
            }
            execution_cursor.machine =
                std::get<CountedData<MachineStateKeys>>(result).data;

            // Use execution cursor as is
            return rocksdb::Status::OK();
        }

        auto machine_state_keys = std::get<MachineStateKeys>(checkpoint_result);
        if (execution_cursor.getOutput().arb_gas_used >
            machine_state_keys.output.arb_gas_used) {
            // Execution cursor used more gas than checkpoint so use it if inbox
            // hash valid
            auto result =
                executionCursorGetMessagesNoLock(tx, execution_cursor, 0);
            if (result.status.ok() && result.data.first) {
                // Execution cursor machine still valid, so use it
                return rocksdb::Status::OK();
            }
        }

        if (!is_for_sideload) {
            auto resolve_status = resolveStagedMessage(tx, machine_state_keys);
            if (!resolve_status.ok()) {
                // Unable to resolve staged_message, try earlier checkpoint
                if (machine_state_keys.output.arb_gas_used == 0) {
                    std::cerr << "first checkpoint corrupted" << std::endl;
                    return resolve_status;
                }
                target_gas_used = machine_state_keys.output.arb_gas_used - 1;
                continue;
            }
        }

        execution_cursor.machine = machine_state_keys;
        return rocksdb::Status::OK();
    }
}

ValueResult<std::pair<bool, std::vector<InboxMessage>>>
ArbCore::executionCursorGetMessages(ReadTransaction& tx,
                                    const ExecutionCursor& execution_cursor,
                                    const uint256_t& orig_message_group_size) {
    const std::lock_guard<std::mutex> lock(core_reorg_mutex);

    return executionCursorGetMessagesNoLock(tx, execution_cursor,
                                            orig_message_group_size);
}

ValueResult<std::pair<bool, std::vector<InboxMessage>>>
ArbCore::executionCursorGetMessagesNoLock(
    ReadTransaction& tx,
    const ExecutionCursor& execution_cursor,
    const uint256_t& orig_message_group_size) {
    auto message_group_size = orig_message_group_size;
    std::vector<InboxMessage> messages;

    // Check if current machine is obsolete
    uint256_t totalRead = execution_cursor.getTotalMessagesRead();
    if (totalRead > 0) {
        auto stored_result = getMessageEntry(tx, totalRead - 1);
        auto inboxAcc = execution_cursor.getInboxAcc();
        if (!stored_result.status.ok() || !inboxAcc) {
            // Obsolete machine, reorg occurred
            return {rocksdb::Status::OK(), std::make_pair(false, messages)};
        }
        if (*inboxAcc != stored_result.data.inbox_acc) {
            // Obsolete machine, reorg occurred
            return {rocksdb::Status::OK(), std::make_pair(false, messages)};
        }
    }

    auto current_message_sequence_number =
        execution_cursor.getTotalMessagesRead();

    auto inserted_message_count_result = messageEntryInsertedCountImpl(tx);

    if (current_message_sequence_number + message_group_size >
        inserted_message_count_result.data) {
        // Don't read past primary machine
        message_group_size = inserted_message_count_result.data -
                             current_message_sequence_number;
    }

    if (message_group_size == 0) {
        // No messages to read
        return {rocksdb::Status::OK(), std::make_pair(true, messages)};
    }

    std::vector<unsigned char> message_key;
    marshal_uint256_t(current_message_sequence_number, message_key);
    auto message_key_slice = vecToSlice(message_key);

    auto results = tx.messageEntryGetVectorVector(
        message_key_slice, intx::narrow_cast<size_t>(message_group_size));
    if (!results.status.ok()) {
        return {results.status, std::make_pair(false, messages)};
    }

    auto total_size = results.data.size();
    messages.reserve(total_size);
    for (const auto& data : results.data) {
        auto message_entry = extractMessageEntry(0, vecToSlice(data));
        auto inbox_message = extractInboxMessage(message_entry.data);
        messages.push_back(inbox_message);
    }

    return {rocksdb::Status::OK(), std::make_pair(true, std::move(messages))};
}

ValueResult<uint256_t> ArbCore::logInsertedCount() const {
    auto tx = makeConstReadOnlyTransaction();

    return logInsertedCountImpl(*tx);
}

ValueResult<uint256_t> ArbCore::logInsertedCountImpl(
    const ReadTransaction& tx) const {
    return tx.stateGetUint256(vecToSlice(log_inserted_key));
}
rocksdb::Status ArbCore::updateLogInsertedCount(ReadWriteTransaction& tx,
                                                const uint256_t& log_index) {
    std::vector<unsigned char> value;
    marshal_uint256_t(log_index, value);

    return tx.statePut(vecToSlice(log_inserted_key), vecToSlice(value));
}

ValueResult<uint256_t> ArbCore::logProcessedCount(ReadTransaction& tx) const {
    return tx.stateGetUint256(vecToSlice(log_processed_key));
}
rocksdb::Status ArbCore::updateLogProcessedCount(ReadWriteTransaction& tx,
                                                 rocksdb::Slice value_slice) {
    return tx.statePut(vecToSlice(log_processed_key), value_slice);
}

ValueResult<uint256_t> ArbCore::sendInsertedCount() const {
    auto tx = makeConstReadOnlyTransaction();

    return sendInsertedCountImpl(*tx);
}

ValueResult<uint256_t> ArbCore::sendInsertedCountImpl(
    const ReadTransaction& tx) const {
    return tx.stateGetUint256(vecToSlice(send_inserted_key));
}

rocksdb::Status ArbCore::updateSendInsertedCount(ReadWriteTransaction& tx,
                                                 const uint256_t& send_index) {
    std::vector<unsigned char> value;
    marshal_uint256_t(send_index, value);

    return tx.statePut(vecToSlice(send_inserted_key), vecToSlice(value));
}

ValueResult<uint256_t> ArbCore::sendProcessedCount(ReadTransaction& tx) const {
    return tx.stateGetUint256(vecToSlice(send_processed_key));
}
rocksdb::Status ArbCore::updateSendProcessedCount(ReadWriteTransaction& tx,
                                                  rocksdb::Slice value_slice) {
    return tx.statePut(vecToSlice(send_processed_key), value_slice);
}

ValueResult<uint256_t> ArbCore::messageEntryInsertedCount() const {
    auto tx = makeConstReadOnlyTransaction();

    return messageEntryInsertedCountImpl(*tx);
}

ValueResult<uint256_t> ArbCore::messageEntryInsertedCountImpl(
    const ReadTransaction& tx) const {
    return tx.stateGetUint256(vecToSlice(message_entry_inserted_key));
}

rocksdb::Status ArbCore::updateMessageEntryInsertedCount(
    ReadWriteTransaction& tx,
    const uint256_t& message_index) {
    std::vector<unsigned char> value;
    marshal_uint256_t(message_index, value);
    return tx.statePut(vecToSlice(message_entry_inserted_key),
                       vecToSlice(value));
}

// addMessages stores all messages from given block into database.
// The last message in the list is flagged as the last message in the block.
// Returns std::nullopt when caller needs to provide messages from earlier
// block.
std::optional<rocksdb::Status> ArbCore::addMessages(
    const std::vector<std::vector<unsigned char>>& new_messages,
    bool last_block_complete,
    const uint256_t& prev_inbox_acc,
    const uint256_t& message_count_in_machine,
    const std::optional<uint256_t>& reorg_message_count,
    ValueCache& cache) {
    auto tx = makeReadWriteTransaction();

    auto message_count_result = messageEntryInsertedCountImpl(*tx);
    if (!message_count_result.status.ok()) {
        return message_count_result.status;
    }
    auto existing_message_count = message_count_result.data;

    auto previous_inbox_acc = prev_inbox_acc;

    uint256_t first_sequence_number;
    uint256_t current_sequence_number;
    if (!new_messages.empty()) {
        auto first_message = extractInboxMessage(new_messages[0]);
        first_sequence_number = first_message.inbox_sequence_number;

        if (first_message.inbox_sequence_number > 0) {
            if (first_message.inbox_sequence_number > existing_message_count) {
                // Not allowed to skip message sequence numbers, ask for older
                // messages
                return std::nullopt;
            }

            // Check that previous_inbox_acc matches acc from previous message
            auto previous_sequence_number =
                first_message.inbox_sequence_number - 1;
            auto previous_result =
                getMessageEntry(*tx, previous_sequence_number);
            if (!previous_result.status.ok()) {
                return previous_result.status;
            }

            if (previous_result.data.inbox_acc != previous_inbox_acc) {
                // Previous inbox doesn't match which means reorg happened and
                // caller needs to try again with messages from earlier block
                return std::nullopt;
            }

            current_sequence_number = first_sequence_number;
        }
    } else {
        if (!reorg_message_count) {
            std::cerr << "reorg_sequence_number must be provided if no "
                         "messages provided"
                      << std::endl;
            return std::nullopt;
        }

        if (*reorg_message_count == 0) {
            std::cerr << "cannot reorg past first message right now"
                      << std::endl;
            return std::nullopt;
        }
        current_sequence_number = *reorg_message_count;
        first_sequence_number = current_sequence_number;
    }

    // Skip any valid messages that we already have in database
    auto new_messages_count = new_messages.size();
    size_t new_messages_index = 0;
    while ((current_sequence_number < existing_message_count) &&
           (new_messages_index < new_messages_count)) {
        auto existing_message_entry =
            getMessageEntry(*tx, current_sequence_number);
        if (!existing_message_entry.status.ok()) {
            return existing_message_entry.status;
        }
        auto current_inbox_acc =
            hash_inbox(previous_inbox_acc, new_messages[new_messages_index]);
        if (existing_message_entry.data.inbox_acc != current_inbox_acc) {
            // Entry doesn't match because of reorg
            break;
        }

        if (existing_message_entry.data.last_message_in_block &&
            !(last_block_complete &&
              new_messages_index == new_messages_count - 1)) {
            // existing message was marked as last message in block but
            // new message is not marked as last message, so they should be
            // considered different
            break;
        }

        new_messages_index++;
        previous_inbox_acc = current_inbox_acc;
        current_sequence_number = first_sequence_number + new_messages_index;
    }

    std::optional<uint256_t> previous_valid_sequence_number;
    if (current_sequence_number < existing_message_count) {
        // Reorg occurred
        const std::lock_guard<std::mutex> lock(core_reorg_mutex);

        previous_valid_sequence_number = current_sequence_number - 1;

        // Truncate MessageEntries to last valid message
        updateMessageEntryInsertedCount(*tx, current_sequence_number);

        if (current_sequence_number <= message_count_in_machine - 1) {
            // Reorg checkpoint and everything else
            auto reorg_status = reorgToMessageOrBefore(
                *tx, *previous_valid_sequence_number, false, cache);
            if (!reorg_status.ok()) {
                return reorg_status;
            }
        }
    }

    while (new_messages_index < new_messages_count) {
        // Encode key
        std::vector<unsigned char> key;
        marshal_uint256_t(current_sequence_number, key);

        auto current_inbox_acc =
            hash_inbox(previous_inbox_acc, new_messages[new_messages_index]);

        auto next_inbox_message =
            extractInboxMessage(new_messages[new_messages_index]);
        auto current_inbox_message = std::move(next_inbox_message);
        if (new_messages_index < new_messages_count) {
            next_inbox_message =
                extractInboxMessage(new_messages[new_messages_index]);
        } else {
            next_inbox_message = {};
        }

        bool last_message_in_block;
        if (last_block_complete &&
            ((new_messages_index == new_messages_count - 1) ||
             current_inbox_message.block_number !=
                 next_inbox_message.block_number)) {
            last_message_in_block = true;
        } else {
            last_message_in_block = false;
        }

        // Encode message entry
        auto messageEntry = MessageEntry{
            current_sequence_number, current_inbox_acc,
            intx::narrow_cast<uint64_t>(current_inbox_message.block_number),
            last_message_in_block, (new_messages[new_messages_index])};
        auto serialized_messageentry = serializeMessageEntry(messageEntry);

        // Save message entry into database
        auto put_status = tx->messageEntryPut(
            vecToSlice(key), vecToSlice(serialized_messageentry));
        if (!put_status.ok()) {
            return put_status;
        }

        new_messages_index++;
        previous_inbox_acc = current_inbox_acc;
        current_sequence_number += 1;
    }

    updateMessageEntryInsertedCount(*tx, current_sequence_number);

    return tx->commit();
}

// deleteLogsStartingAt deletes the given index along with any
// newer logs. Returns std::nullopt if nothing deleted.
std::optional<rocksdb::Status> deleteLogsStartingAt(ReadWriteTransaction& tx,
                                                    uint256_t log_index) {
    auto it = tx.logGetIterator();

    // Find first message to delete
    std::vector<unsigned char> key;
    marshal_uint256_t(log_index, key);
    it->Seek(vecToSlice(key));
    if (it->status().IsNotFound()) {
        // Nothing to delete
        return std::nullopt;
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

void ArbCore::handleLogsCursorRequested(ReadTransaction& tx,
                                        size_t cursor_index,
                                        ValueCache& cache) {
    if (cursor_index >= logs_cursors.size()) {
        throw std::runtime_error("Invalid logsCursor index");
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    // Provide requested logs
    logs_cursors[cursor_index].data.clear();
    auto log_inserted_count = logInsertedCountImpl(tx);
    if (!log_inserted_count.status.ok()) {
        logs_cursors[cursor_index].error_string =
            log_inserted_count.status.ToString();
        std::cerr << "logscursor index " << cursor_index
                  << " error getting inserted count: "
                  << log_inserted_count.status.ToString() << std::endl;
        logs_cursors[cursor_index].status = DataCursor::ERROR;
        return;
    }

    auto current_count_result =
        logsCursorGetCurrentTotalCount(tx, cursor_index);
    if (!current_count_result.status.ok()) {
        std::cerr << "Unable to get logs cursor current total count: "
                  << cursor_index << std::endl;
        return;
    }

    if (current_count_result.data == log_inserted_count.data) {
        // Nothing to do
        logs_cursors[cursor_index].status = DataCursor::READY;
        return;
    }
    if (current_count_result.data > log_inserted_count.data) {
        // Error
        std::cerr << "handleLogsCursor current count: "
                  << current_count_result.data << " == "
                  << " log inserted count: " << log_inserted_count.data
                  << std::endl;
        logs_cursors[cursor_index].status = DataCursor::READY;
        return;
    }
    if (current_count_result.data +
            logs_cursors[cursor_index].number_requested >
        log_inserted_count.data) {
        // Too many entries requested
        logs_cursors[cursor_index].number_requested =
            log_inserted_count.data - current_count_result.data;
    }
    if (logs_cursors[cursor_index].number_requested == 0) {
        logs_cursors[cursor_index].status = DataCursor::READY;
        // No new logs to provide
        return;
    }
    auto requested_logs =
        getLogs(current_count_result.data,
                logs_cursors[cursor_index].number_requested, cache);
    if (!requested_logs.status.ok()) {
        logs_cursors[cursor_index].error_string =
            requested_logs.status.ToString();
        logs_cursors[cursor_index].status = DataCursor::ERROR;
        std::cerr << "logscursor index " << cursor_index
                  << " error getting logs: " << requested_logs.status.ToString()
                  << std::endl;
        return;
    }
    logs_cursors[cursor_index].data = std::move(requested_logs.data);
    logs_cursors[cursor_index].status = DataCursor::READY;

    return;
}

// handleLogsCursorReorg must be called before logs are deleted.
// Note that this function should not update logs_cursors[cursor_index].status
// because it is happening out of line.
// Note that cursor reorg never adds new messages, but might add deleted
// messages.
rocksdb::Status ArbCore::handleLogsCursorReorg(ReadWriteTransaction& tx,
                                               size_t cursor_index,
                                               uint256_t log_count,
                                               ValueCache& cache) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        return rocksdb::Status::InvalidArgument();
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    auto current_count_result =
        logsCursorGetCurrentTotalCount(tx, cursor_index);
    if (!current_count_result.status.ok()) {
        std::cerr << "Unable to get logs cursor current total count: "
                  << cursor_index << "\n";
        return current_count_result.status;
    }

    if (current_count_result.data >
        logs_cursors[cursor_index].pending_total_count) {
        logs_cursors[cursor_index].pending_total_count =
            current_count_result.data;
    }

    if (log_count < logs_cursors[cursor_index].pending_total_count) {
        // Need to save logs that will be deleted
        auto logs = getLogsNoLock(
            tx, log_count,
            logs_cursors[cursor_index].pending_total_count - log_count, cache);
        if (!logs.status.ok()) {
            std::cerr << "Error getting "
                      << logs_cursors[cursor_index].pending_total_count -
                             log_count
                      << " logs starting at " << log_count
                      << " in Cursor reorg : " << logs.status.ToString()
                      << "\n";
            return logs.status;
        }
        logs_cursors[cursor_index].deleted_data.insert(
            logs_cursors[cursor_index].deleted_data.end(), logs.data.rbegin(),
            logs.data.rend());

        logs_cursors[cursor_index].pending_total_count = log_count;

        if (current_count_result.data > log_count) {
            auto status =
                logsCursorSaveCurrentTotalCount(tx, cursor_index, log_count);
            if (!status.ok()) {
                std::cerr << "unable to save current total count during reorg"
                          << std::endl;
                return status;
            }
        }
    }

    if (!logs_cursors[cursor_index].data.empty()) {
        if (current_count_result.data >= log_count) {
            // Don't save anything
            logs_cursors[cursor_index].data.clear();
        } else if (current_count_result.data +
                       logs_cursors[cursor_index].data.size() >
                   log_count) {
            // Only part of the data needs to be removed
            auto logs_to_keep = intx::narrow_cast<size_t>(
                log_count - current_count_result.data);
            logs_cursors[cursor_index].data.erase(
                logs_cursors[cursor_index].data.begin() + logs_to_keep,
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

std::optional<std::pair<uint256_t, std::vector<value>>>
ArbCore::logsCursorGetLogs(size_t cursor_index) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        return std::nullopt;
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    if (logs_cursors[cursor_index].status != DataCursor::READY ||
        !logs_cursors[cursor_index].deleted_data.empty()) {
        return std::nullopt;
    }

    auto tx = makeReadOnlyTransaction();
    auto current_count_result =
        logsCursorGetCurrentTotalCount(*tx, cursor_index);
    if (!current_count_result.status.ok()) {
        std::cerr << "Unable to get logs cursor current total count: "
                  << cursor_index << "\n";
        return std::nullopt;
    }

    logs_cursors[cursor_index].pending_total_count =
        current_count_result.data + logs_cursors[cursor_index].data.size();

    std::vector<value> logs;
    logs = std::move(logs_cursors[cursor_index].data);
    logs_cursors[cursor_index].data.clear();

    return {{current_count_result.data, std::move(logs)}};
}

std::optional<std::pair<uint256_t, std::vector<value>>>
ArbCore::logsCursorGetDeletedLogs(size_t cursor_index) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        return std::nullopt;
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    if (logs_cursors[cursor_index].status != DataCursor::READY ||
        logs_cursors[cursor_index].deleted_data.empty()) {
        return std::nullopt;
    }

    std::vector<value> logs;
    logs.swap(logs_cursors[cursor_index].deleted_data);
    logs_cursors[cursor_index].deleted_data.clear();

    auto tx = makeReadOnlyTransaction();
    auto current_count_result =
        logsCursorGetCurrentTotalCount(*tx, cursor_index);
    if (!current_count_result.status.ok()) {
        std::cerr << "Unable to get logs cursor current total count: "
                  << cursor_index << "\n";
        return std::nullopt;
    }

    return {{current_count_result.data, std::move(logs)}};
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

    if (!logs_cursors[cursor_index].data.empty()) {
        // Still have logs to get
        std::cerr << "logs cursor " << cursor_index
                  << " has messages left in cursor when trying to confirm"
                  << std::endl;
        return false;
    }

    if (!logs_cursors[cursor_index].data.empty() ||
        !logs_cursors[cursor_index].deleted_data.empty()) {
        // Still have logs to get
        return false;
    }

    auto tx = makeReadWriteTransaction();
    auto status = logsCursorSaveCurrentTotalCount(
        *tx, cursor_index, logs_cursors[cursor_index].pending_total_count);
    tx->commit();

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

rocksdb::Status ArbCore::logsCursorSaveCurrentTotalCount(
    ReadWriteTransaction& tx,
    size_t cursor_index,
    uint256_t count) {
    std::vector<unsigned char> value_data;
    marshal_uint256_t(count, value_data);
    return tx.statePut(vecToSlice(logs_cursors[cursor_index].current_total_key),
                       vecToSlice(value_data));
}

ValueResult<uint256_t> ArbCore::logsCursorGetCurrentTotalCount(
    ReadTransaction& tx,
    size_t cursor_index) {
    return tx.stateGetUint256(
        vecToSlice(logs_cursors[cursor_index].current_total_key));
}

rocksdb::Status ArbCore::saveSideloadPosition(ReadWriteTransaction& tx,
                                              const uint256_t& block_number,
                                              const uint256_t& arb_gas_used) {
    std::vector<unsigned char> key;
    marshal_uint256_t(block_number, key);
    auto key_slice = vecToSlice(key);

    std::vector<unsigned char> value;
    marshal_uint256_t(arb_gas_used, value);
    auto value_slice = vecToSlice(value);

    return tx.sideloadPut(key_slice, value_slice);
}

ValueResult<uint256_t> ArbCore::getSideloadPosition(
    ReadTransaction& tx,
    const uint256_t& block_number) {
    std::vector<unsigned char> key;
    marshal_uint256_t(block_number, key);
    auto key_slice = vecToSlice(key);

    auto it = tx.sideloadGetIterator();

    it->SeekForPrev(key_slice);

    auto s = it->status();
    if (!s.ok()) {
        return {s, 0};
    }

    auto value_slice = it->value();

    return {s, intx::be::unsafe::load<uint256_t>(
                   reinterpret_cast<const unsigned char*>(value_slice.data()))};
}

rocksdb::Status ArbCore::deleteSideloadsStartingAt(
    ReadWriteTransaction& tx,
    const uint256_t& block_number) {
    // Clear the cache
    {
        std::unique_lock<std::shared_mutex> guard(sideload_cache_mutex);
        auto it = sideload_cache.lower_bound(block_number);
        while (it != sideload_cache.end()) {
            it = sideload_cache.erase(it);
        }
    }

    // Clear the DB
    std::vector<unsigned char> key;
    marshal_uint256_t(block_number, key);
    auto key_slice = vecToSlice(key);

    auto it = tx.sideloadGetIterator();

    it->Seek(key_slice);

    while (it->Valid()) {
        tx.sideloadDelete(it->key());
        it->Next();
    }
    auto s = it->status();
    if (s.IsNotFound()) {
        s = rocksdb::Status::OK();
    }
    return s;
}

ValueResult<std::unique_ptr<Machine>> ArbCore::getMachineForSideload(
    const uint256_t& block_number,
    ValueCache& cache) {
    // Check the cache
    {
        std::shared_lock<std::shared_mutex> lock(sideload_cache_mutex);
        // Look for the first value after the value we want
        auto it = sideload_cache.upper_bound(block_number);
        if (it != sideload_cache.begin()) {
            // Go back a value to find the one we want
            it--;
            return {rocksdb::Status::OK(),
                    std::make_unique<Machine>(*it->second)};
        }
    }
    // Not found in cache, try the DB
    auto tx = makeReadOnlyTransaction();
    tx->enterReadSnapshot();
    auto position_res = getSideloadPosition(*tx, block_number);
    if (!position_res.status.ok()) {
        return {position_res.status, std::unique_ptr<Machine>(nullptr)};
    }

    auto initial_hash = getInitialMachineHash(*tx);
    if (!initial_hash.status.ok()) {
        return {initial_hash.status, nullptr};
    }
    auto result = getMachineStateKeys(*tx, initial_hash.data);
    if (std::holds_alternative<rocksdb::Status>(result)) {
        return {std::get<rocksdb::Status>(result), nullptr};
    }

    auto execution_cursor = std::make_unique<ExecutionCursor>(
        std::get<CountedData<MachineStateKeys>>(result).data);

    auto status =
        executionCursorSetup(*tx, *execution_cursor, position_res.data, true);

    return {status,
            takeExecutionCursorMachineImpl(*tx, *execution_cursor, cache)};
}
