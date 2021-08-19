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

#include "value/corevalueloader.hpp"

#include <avm/inboxmessage.hpp>
#include <avm/machinethread.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/readsnapshottransaction.hpp>
#include <data_storage/readwritetransaction.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>
#include <data_storage/value/utils.hpp>
#include <data_storage/value/value.hpp>
#include <data_storage/value/valuecache.hpp>

#include <ethash/keccak.hpp>
#include <iomanip>
#include <set>
#include <sstream>
#include <vector>

#ifdef __linux__
#include <sys/prctl.h>
#endif

namespace {
constexpr uint256_t arbcore_schema_version = 2;
constexpr auto log_inserted_key = std::array<char, 1>{-60};
constexpr auto log_processed_key = std::array<char, 1>{-61};
constexpr auto send_inserted_key = std::array<char, 1>{-62};
constexpr auto send_processed_key = std::array<char, 1>{-63};
constexpr auto schema_version_key = std::array<char, 1>{-64};
constexpr auto logscursor_current_prefix = std::array<char, 1>{-120};
}  // namespace

ArbCore::ArbCore(std::shared_ptr<DataStorage> data_storage_,
                 const ArbCoreConfig& coreConfig_)
    : coreConfig(coreConfig_),
      data_storage(std::move(data_storage_)),
      core_code(std::make_shared<CoreCode>(getNextSegmentID(data_storage))),
      timed_sideload_cache(coreConfig.timed_cache_expiration_seconds) {
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
    if (message_data_status != ArbCore::MESSAGES_ERROR) {
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
    if (core_thread != nullptr) {
        return false;
    }

    core_thread =
        std::make_unique<std::thread>(std::reference_wrapper<ArbCore>(*this));

    return true;
}

void ArbCore::abortThread() {
    std::cerr << "Aborting main ArbCore thread" << std::endl;
    if (core_thread) {
        arbcore_abort = true;
        core_thread->join();
        core_thread = nullptr;
    }
    arbcore_abort = false;
}

// deliverMessages sends messages to core thread
bool ArbCore::deliverMessages(
    const uint256_t& previous_message_count,
    const uint256_t& previous_batch_acc,
    std::vector<std::vector<unsigned char>> sequencer_batch_items,
    std::vector<std::vector<unsigned char>> delayed_messages,
    const std::optional<uint256_t>& reorg_batch_items) {
    if (message_data_status != MESSAGES_EMPTY) {
        return false;
    }

    message_data.previous_message_count = previous_message_count;
    message_data.previous_batch_acc = previous_batch_acc;
    message_data.sequencer_batch_items = std::move(sequencer_batch_items);
    message_data.delayed_messages = std::move(delayed_messages);
    message_data.reorg_batch_items = reorg_batch_items;

    message_data_status = MESSAGES_READY;

    return true;
}

ValueLoader ArbCore::makeValueLoader() const {
    return ValueLoader(std::make_unique<CoreValueLoader>(
        data_storage, core_code, ValueCache{1, 0}));
}

rocksdb::Status ArbCore::initialize(const LoadedExecutable& executable) {
    // Use latest existing checkpoint
    ValueCache cache{1, 0};

    auto status = reorgToTimestampOrBefore(
        timed_sideload_cache.expiredTimestamp(), true, cache);
    if (status.ok()) {
        // Database already initialized
        return status;
    }
    if (!status.IsNotFound()) {
        std::cerr << "Error with initial reorg: " << status.ToString()
                  << std::endl;
        return status;
    }

    // Need to initialize database from scratch
    core_code->addSegment(executable.code);
    core_machine = std::make_unique<MachineThread>(
        MachineState{core_code, executable.static_val});
    core_machine->machine_state.value_loader = makeValueLoader();
    core_machine->machine_state.code = std::make_shared<RunningCode>(core_code);

    last_machine = std::make_unique<Machine>(*core_machine);

    ReadWriteTransaction tx(data_storage);

    status = updateSchemaVersion(tx, arbcore_schema_version);
    if (!status.ok()) {
        std::cerr << "failed to save schema version into db: "
                  << status.ToString() << std::endl;
        return status;
    }

    status = saveCheckpoint(tx);
    if (!status.ok()) {
        std::cerr << "failed to save initial checkpoint into db: "
                  << status.ToString() << std::endl;
        return status;
    }

    status = updateLogInsertedCount(tx, 0);
    if (!status.ok()) {
        throw std::runtime_error("failed to initialize log inserted count");
    }
    status = updateSendInsertedCount(tx, 0);
    if (!status.ok()) {
        throw std::runtime_error("failed to initialize log inserted count");
    }

    for (size_t i = 0; i < logs_cursors.size(); i++) {
        status = logsCursorSaveCurrentTotalCount(tx, i, 0);
        if (!status.ok()) {
            throw std::runtime_error("failed to initialize logscursor counts");
        }
    }

    status = tx.commit();
    if (!status.ok()) {
        std::cerr << "failed to commit initial state into db: "
                  << status.ToString() << std::endl;
        return status;
    }

    return rocksdb::Status::OK();
}

bool ArbCore::initialized() const {
    ReadTransaction tx(data_storage);
    std::vector<unsigned char> key;
    marshal_uint256_t(0, key);
    return tx.checkpointGetVector(vecToSlice(key)).status.ok();
}

template <class T>
std::unique_ptr<T> ArbCore::getMachineImpl(ReadTransaction& tx,
                                           uint256_t machineHash,
                                           ValueCache& value_cache) {
    auto results = getMachineStateKeys(tx, machineHash);
    if (std::holds_alternative<rocksdb::Status>(results)) {
        throw std::runtime_error("failed to load machine state");
    }

    auto res =
        std::get<CountedData<std::variant<MachineStateKeys, MachineOutput>>>(
            results)
            .data;
    if (std::holds_alternative<MachineStateKeys>(res)) {
        return getMachineUsingStateKeys<T>(tx, std::get<MachineStateKeys>(res),
                                           value_cache);
    }

    // Machine not found
    return nullptr;
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
    ReadSnapshotTransaction tx(data_storage);
    return getMachineImpl<T>(tx, machineHash, value_cache);
}

template std::unique_ptr<Machine> ArbCore::getMachine(uint256_t, ValueCache&);
template std::unique_ptr<MachineThread> ArbCore::getMachine(uint256_t,
                                                            ValueCache&);

// triggerSaveCheckpoint is meant for unit tests and should not be called from
// multiple threads at the same time.
rocksdb::Status ArbCore::triggerSaveCheckpoint() {
    save_checkpoint = true;
    std::cerr << "Triggering checkpoint save" << std::endl;
    while (save_checkpoint) {
        // Wait until snapshot has been saved
        std::this_thread::sleep_for(std::chrono::milliseconds(10));
    }
    std::cerr << "Checkpoint saved" << std::endl;

    return save_checkpoint_status;
}

rocksdb::Status ArbCore::saveCheckpoint(ReadWriteTransaction& tx) {
    auto& state = core_machine->machine_state;
    if (!isValid(tx, state.output.fully_processed_inbox)) {
        std::cerr << "Attempted to save invalid checkpoint at gas "
                  << state.output.arb_gas_used << std::endl;
        assert(false);
        return rocksdb::Status::OK();
    }

    auto save_res = saveMachineState(tx, *core_machine);
    if (!save_res.first.ok()) {
        return save_res.first;
    }

    auto machine_code =
        dynamic_cast<RunningCode*>(core_machine->machine_state.code.get());
    assert(machine_code != nullptr);
    core_code->commitChanges(*machine_code, save_res.second);
    core_machine->machine_state.code = std::make_shared<RunningCode>(core_code);

    std::vector<unsigned char> key;
    marshal_uint256_t(state.output.arb_gas_used, key);
    auto key_slice = vecToSlice(key);
    std::vector<unsigned char> value_vec;
    serializeMachineStateKeys(MachineStateKeys{state}, value_vec);
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

    if (assertion.sideload_block_number) {
        status = saveSideloadPosition(tx, *assertion.sideload_block_number,
                                      arb_gas_used);
        if (!status.ok()) {
            return status;
        }
    }

    return rocksdb::Status::OK();
}

rocksdb::Status ArbCore::reorgToMessageCountOrBefore(
    const uint256_t& message_count,
    bool initial_start,
    ValueCache& cache) {
    if (initial_start) {
        std::cerr << "Reloading chain starting with message " << message_count
                  << "\n";
    } else {
        std::cerr << "Reorg'ing chain to message " << message_count << "\n";
    }

    return reorgCheckpoints(
        [&](const MachineOutput& output) {
            return message_count >= output.fully_processed_inbox.count;
        },
        initial_start, cache);
}

rocksdb::Status ArbCore::reorgToTimestampOrBefore(const uint256_t& timestamp,
                                                  bool initial_start,
                                                  ValueCache& cache) {
    if (initial_start) {
        std::cerr << "Reloading chain starting with timestamp " << timestamp
                  << "\n";
    } else {
        std::cerr << "Reorg'ing chain to timestamp " << timestamp << "\n";
    }

    return reorgCheckpoints(
        [&](const MachineOutput& output) {
            return timestamp >= output.last_inbox_timestamp;
        },
        initial_start, cache);
}

// reorgCheckpoints resets the checkpoint and database entries
// such that machine state is at or before the requested message. cleaning
// up old references as needed.
// If initial_start is true the various caching data structures are seeded.
rocksdb::Status ArbCore::reorgCheckpoints(
    const std::function<bool(const MachineOutput&)>& check_output,
    bool initial_start,
    ValueCache& cache) {
    std::variant<std::unique_ptr<MachineThread>, rocksdb::Status> setup =
        rocksdb::Status::OK();

    if (initial_start) {
        std::cerr << "Reloading cache" << std::endl;
    } else {
        std::cerr << "Reorganizing" << std::endl;
    }

    // Save selected machine output so we know how long to execute machine to
    // match selected checkpoint.  This will be necessary if selected checkpoint
    // does not include machine or if entry from old_checkpoint_cache is behind
    // the selected checkpoint
    std::optional<MachineOutput> selected_machine_output;

    {
        ReadWriteTransaction tx(data_storage);

        // Delete each old cached machine until at or below
        // message_sequence_number
        {
            std::unique_lock<std::shared_mutex> guard(old_machine_cache_mutex);
            auto old_machine_it = old_machine_cache.end();
            while (old_machine_it != old_machine_cache.begin()) {
                old_machine_it--;
                auto& inbox = old_machine_it->second->machine_state.output
                                  .fully_processed_inbox;
                if (check_output(
                        old_machine_it->second->machine_state.output)) {
                    if (isValid(tx, inbox)) {
                        setup = std::make_unique<MachineThread>(
                            old_machine_it->second->machine_state);
                        break;
                    }

                    std::cerr
                        << "Error: Invalid cached old machine found at gas: "
                        << old_machine_it->second->machine_state.output
                               .arb_gas_used
                        << std::endl;
                    assert(false);
                }
                old_machine_it = old_machine_cache.erase(old_machine_it);
            }
        }

        auto checkpoint_it = tx.checkpointGetIterator();

        // Find first checkpoint to delete
        checkpoint_it->SeekToLast();
        if (!checkpoint_it->status().ok()) {
            std::cerr << "Error: SeekToLast failed during reorg: "
                      << checkpoint_it->status().ToString() << std::endl;
            return checkpoint_it->status();
        }

        if (!checkpoint_it->Valid()) {
            return rocksdb::Status::NotFound();
        }

        // At this point we know database has already been initialized, so
        // check schema version
        auto schema_result = schemaVersion(tx);
        if (!schema_result.status.ok()) {
            std::cerr << "Error getting schema version: "
                      << schema_result.status.ToString()
                      << ", delete database and try again" << std::endl;
            return rocksdb::Status::Corruption();
        }
        if (schema_result.data != arbcore_schema_version) {
            std::cerr << "Database version " << schema_result.data
                      << " does not match expected version "
                      << arbcore_schema_version
                      << ", delete database and try again" << std::endl;
            return rocksdb::Status::Corruption();
        }

        // Delete each checkpoint until at or below message_sequence_number
        while (checkpoint_it->Valid()) {
            std::vector<unsigned char> checkpoint_vector(
                checkpoint_it->value().data(),
                checkpoint_it->value().data() + checkpoint_it->value().size());
            auto checkpoint_variant =
                extractMachineStateKeys(checkpoint_vector);
            if (std::holds_alternative<MachineOutput>(checkpoint_variant)) {
                auto machine_output =
                    std::get<MachineOutput>(checkpoint_variant);
                if (check_output(machine_output)) {
                    if (!selected_machine_output.has_value()) {
                        // Save selected output to know how much machine needs
                        // to be executed if it behind
                        selected_machine_output = machine_output;
                    }

                    if (std::holds_alternative<std::unique_ptr<MachineThread>>(
                            setup)) {
                        // Checkpoint without machine, but machine restored
                        // from old_machine_cache
                        break;
                    }

                    // Need to get machine from older checkpoint without
                    // full reorg because current checkpoint does not
                    // have a valid machine
                    checkpoint_it->Prev();
                    continue;
                }
            }

            auto checkpoint = std::get<MachineStateKeys>(checkpoint_variant);
            if (check_output(checkpoint.output)) {
                if (isValid(tx, checkpoint.output.fully_processed_inbox)) {
                    // Good checkpoint
                    try {
                        if (std::holds_alternative<rocksdb::Status>(setup)) {
                            setup = getMachineUsingStateKeys<MachineThread>(
                                tx, checkpoint, cache);
                        }

                        // Machine loaded from database or from
                        // old_machine_cache
                        break;
                    } catch (const std::exception& e) {
                        std::cerr << "Error loading machine from checkpoint: "
                                  << e.what() << std::endl;
                        assert(false);
                    }
                } else {
                    std::cerr << "Unexpectedly invalid checkpoint inbox at "
                                 "message count "
                              << checkpoint.output.fully_processed_inbox.count
                              << std::endl;
                    assert(false);
                }

                // Obsolete checkpoint, need to delete referenced machine
                deleteMachineState(tx, checkpoint);
            }

            // Delete checkpoint to make sure it isn't used later
            tx.checkpointDelete(checkpoint_it->key());

            checkpoint_it->Prev();
        }
        if (!checkpoint_it->Valid()) {
            setup = checkpoint_it->status();
        }

        checkpoint_it = nullptr;
        if (std::holds_alternative<rocksdb::Status>(setup)) {
            return std::get<rocksdb::Status>(setup);
        }

        auto status = tx.commit();
        if (!status.ok()) {
            std::cerr << "Error: unable to commit during reorg"
                      << status.ToString() << std::endl;
            return status;
        }
    }
    // Machine was executing obsolete messages so restore machine
    // from last checkpoint
    if (core_machine != nullptr) {
        core_machine->abortMachine();
    }

    core_machine = std::get<std::unique_ptr<MachineThread>>(std::move(setup));
    auto& output = core_machine->machine_state.output;

    while (selected_machine_output.has_value() &&
           (core_machine->machine_state.output.arb_gas_used <
            selected_machine_output->arb_gas_used)) {
        // Need to run machine until caught up with current
        // checkpoint
        MachineExecutionConfig execConfig;
        execConfig.stop_on_sideload = initial_start;

        // Add messages and run machine
        auto success = runMachineWithMessages(execConfig,
                                              coreConfig.message_process_count);
        if (!success) {
            std::cerr << "runMachineWithMessages failed" << core_error_string
                      << "\n";
            return rocksdb::Status::Aborted();
        }

        if (core_machine->status() == MachineThread::MACHINE_ERROR) {
            core_error_string = core_machine->getErrorString();
            std::cerr << "AVM machine stopped with error: " << core_error_string
                      << "\n";
            return rocksdb::Status::Aborted();
        }

        while (core_machine->nextAssertion().sideload_block_number) {
            timed_sideload_cache.add(std::make_unique<Machine>(*core_machine));

            if (core_machine->machine_state.output.arb_gas_used >=
                selected_machine_output->arb_gas_used) {
                break;
            }

            // Machine was stopped to save sideload,
            // start machine back up where it stopped
            auto machine_success = core_machine->continueRunningMachine();
            if (!machine_success) {
                core_error_string = "Error continuing machine thread";
                machine_error = true;
                std::cerr << "Error catching up: " << core_error_string << "\n";
                return rocksdb::Status::Aborted();
            }
        }
    }

    if (selected_machine_output.has_value() &&
        core_machine->machine_state.output != selected_machine_output) {
        // Machine in unexpected state, data corruption might have occurred
        std::cerr << "Error catching up: machine in unexpected state"
                  << "\n";
        return rocksdb::Status::Aborted();
    }

    auto log_inserted_count = logInsertedCount();
    if (!log_inserted_count.status.ok()) {
        std::cerr << "Error getting inserted count in Cursor Reorg: "
                  << log_inserted_count.status.ToString() << "\n";
        return log_inserted_count.status;
    }

    if (output.log_count < log_inserted_count.data) {
        // Update log cursors, must be called before logs are deleted
        for (size_t i = 0; i < logs_cursors.size(); i++) {
            auto status = handleLogsCursorReorg(i, output.log_count, cache);
            if (!status.ok()) {
                return status;
            }
        }
    }

    uint256_t next_sideload_block_number = 0;
    if (output.last_sideload) {
        next_sideload_block_number = *output.last_sideload + 1;
    }

    ReadWriteTransaction tx(data_storage);
    auto status = deleteSideloadsStartingAt(tx, next_sideload_block_number);
    if (!status.ok()) {
        return status;
    }

    // Delete logs individually to handle reference counts
    auto optional_status = deleteLogsStartingAt(tx, output.log_count);
    if (optional_status && !optional_status->ok()) {
        return *optional_status;
    }

    status = updateLogInsertedCount(tx, output.log_count);
    if (!status.ok()) {
        return status;
    }

    status = updateSendInsertedCount(tx, output.send_count);
    if (!status.ok()) {
        return status;
    }

    // Update last machine
    {
        std::unique_lock<std::shared_mutex> guard(last_machine_mutex);
        last_machine = std::make_unique<Machine>(*core_machine);
    }

    // Checkpoint was saved at sideload, attempt to continue running
    core_machine->continueRunningMachine();

    return tx.commit();
}

bool ArbCore::isCheckpointsEmpty(ReadTransaction& tx) const {
    auto it = std::unique_ptr<rocksdb::Iterator>(tx.checkpointGetIterator());
    it->SeekToLast();
    return !it->Valid();
}

uint256_t ArbCore::maxCheckpointGas() {
    ReadTransaction tx(data_storage);
    auto it = tx.checkpointGetIterator();
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
    const uint256_t& total_gas) {
    auto it = tx.checkpointGetIterator();
    std::vector<unsigned char> key;
    marshal_uint256_t(total_gas, key);
    auto key_slice = vecToSlice(key);
    it->SeekForPrev(key_slice);
    while (it->Valid()) {
        if (!it->status().ok()) {
            return it->status();
        }

        std::vector<unsigned char> saved_value(
            it->value().data(), it->value().data() + it->value().size());
        auto variantkeys = extractMachineStateKeys(saved_value);

        if (std::holds_alternative<MachineStateKeys>(variantkeys)) {
            // Found checkpoint with machine
            return std::get<MachineStateKeys>(variantkeys);
        }

        // Checkpoint did not contain machine
        it->Prev();
    }

    if (!it->status().ok()) {
        return it->status();
    }
    return rocksdb::Status::NotFound();
}

template <class T>
std::unique_ptr<T> ArbCore::getMachineUsingStateKeys(
    const ReadTransaction& transaction,
    const MachineStateKeys& state_data,
    ValueCache& value_cache) const {
    std::set<uint64_t> segment_ids;

    auto static_results = ::getValueImpl(transaction, state_data.static_hash,
                                         segment_ids, value_cache, false);

    if (std::holds_alternative<rocksdb::Status>(static_results)) {
        std::stringstream ss;
        ss << "failed loaded core machine static: "
           << std::get<rocksdb::Status>(static_results).ToString();
        std::cerr << "getValueImpl error: " << ss.str() << std::endl;
        throw std::runtime_error(ss.str());
    }

    auto register_results =
        ::getValueImpl(transaction, state_data.register_hash, segment_ids,
                       value_cache, ENABLE_LAZY_LOADING);
    if (std::holds_alternative<rocksdb::Status>(register_results)) {
        std::stringstream ss;
        ss << "failed loaded core machine register with hash "
           << state_data.register_hash << ": "
           << std::get<rocksdb::Status>(register_results).ToString();
        std::cerr << "getValueImpl error: " << ss.str() << std::endl;
        throw std::runtime_error(ss.str());
    }

    auto stack_results = ::getValueImpl(transaction, state_data.datastack_hash,
                                        segment_ids, value_cache, false);
    if (std::holds_alternative<rocksdb::Status>(stack_results) ||
        !std::holds_alternative<Tuple>(
            std::get<CountedData<value>>(stack_results).data)) {
        std::cerr << "failed to load machine stack" << std::endl;
        throw std::runtime_error("failed to load machine stack");
    }

    auto auxstack_results = ::getValueImpl(
        transaction, state_data.auxstack_hash, segment_ids, value_cache, false);
    if (std::holds_alternative<rocksdb::Status>(auxstack_results)) {
        std::cerr << "failed to load machine auxstack" << std::endl;
        throw std::runtime_error("failed to load machine auxstack");
    }
    if (!std::holds_alternative<Tuple>(
            std::get<CountedData<value>>(auxstack_results).data)) {
        std::cerr << "failed to load machine auxstack because of format error"
                  << std::endl;
        throw std::runtime_error(
            "failed to load machine auxstack because of format error");
    }

    segment_ids.insert(state_data.pc.pc.segment);
    segment_ids.insert(state_data.err_pc.pc.segment);

    restoreCodeSegments(transaction, core_code, value_cache, segment_ids);

    auto state = MachineState{
        state_data.output,
        state_data.pc.pc,
        std::make_shared<RunningCode>(core_code),
        makeValueLoader(),
        std::move(std::get<CountedData<value>>(register_results).data),
        std::move(std::get<CountedData<value>>(static_results).data),
        Datastack(
            std::get<Tuple>(std::get<CountedData<value>>(stack_results).data)),
        Datastack(std::get<Tuple>(
            std::get<CountedData<value>>(auxstack_results).data)),
        state_data.arb_gas_remaining,
        state_data.state,
        state_data.err_pc};

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

constexpr uint256_t old_machine_cache_interval = 1'000'000;
constexpr size_t old_machine_cache_max_size = 20;

// operator() runs the main thread for ArbCore.  It is responsible for adding
// messages to the queue, starting machine thread when needed and collecting
// results of machine thread.
// This thread will update `delivering_messages` if and only if
// `delivering_messages` is set to MESSAGES_READY
void ArbCore::operator()() {
#ifdef __linux__
    prctl(PR_SET_NAME, "ArbCore", 0, 0, 0);
#endif
    ValueCache cache{5, 0};
    MachineExecutionConfig execConfig;
    execConfig.stop_on_sideload = true;

    uint256_t next_checkpoint_gas =
        maxCheckpointGas() + coreConfig.min_gas_checkpoint_frequency;
    while (!arbcore_abort) {
        bool isMachineValid;
        {
            ReadTransaction tx(data_storage);
            isMachineValid = isValid(tx, core_machine->getReorgData());
        }
        if (!isMachineValid) {
            std::cerr
                << "Core thread operating on invalid machine. Rolling back."
                << std::endl;
            assert(false);
            auto status = reorgToMessageCountOrBefore(0, false, cache);
            if (!status.ok()) {
                std::cerr << "Error in core thread calling "
                             "reorgCheckpoints: "
                          << status.ToString() << std::endl;
            }
            next_checkpoint_gas = coreConfig.min_gas_checkpoint_frequency;
        }
        if (message_data_status == MESSAGES_READY) {
            // Reorg might occur while adding messages
            try {
                auto add_status = addMessages(message_data, cache);
                if (!add_status.status.ok()) {
                    core_error_string = add_status.status.ToString();
                    message_data_status = MESSAGES_ERROR;
                    std::cerr
                        << "ArbCore addMessages error: " << core_error_string
                        << "\n";
                } else {
                    machine_idle = false;
                    message_data_status = MESSAGES_SUCCESS;
                    if (add_status.data.has_value()) {
                        next_checkpoint_gas =
                            add_status.data.value() +
                            coreConfig.min_gas_checkpoint_frequency;
                    }
                }
            } catch (const std::exception& e) {
                core_error_string = e.what();
                message_data_status = MESSAGES_ERROR;
                std::cerr << "ArbCore addMessages exception: "
                          << core_error_string << "\n";
            }
        }

        // Check machine thread
        if (core_machine->status() == MachineThread::MACHINE_ERROR) {
            core_error_string = core_machine->getErrorString();
            std::cerr << "AVM machine stopped with error: " << core_error_string
                      << "\n";
            break;
        }

        if (core_machine->status() == MachineThread::MACHINE_SUCCESS) {
            ReadWriteTransaction tx(data_storage);

            auto last_assertion = core_machine->nextAssertion();

            // Save last machine output
            {
                std::unique_lock<std::shared_mutex> guard(last_machine_mutex);
                last_machine = std::make_unique<Machine>(*core_machine);
            }

            if (core_machine->machine_state.output.arb_gas_used >
                last_old_machine_cache_gas + old_machine_cache_interval) {
                std::unique_lock<std::shared_mutex> guard(
                    old_machine_cache_mutex);
                if (old_machine_cache.size() > old_machine_cache_max_size) {
                    old_machine_cache.erase(old_machine_cache.begin());
                }
                old_machine_cache[core_machine->machine_state.output
                                      .arb_gas_used] =
                    std::make_unique<Machine>(*core_machine);
            }

            // Save logs and sends
            auto status =
                saveAssertion(tx, last_assertion,
                              core_machine->machine_state.output.arb_gas_used);
            if (!status.ok()) {
                core_error_string = status.ToString();
                std::cerr << "ArbCore assertion saving failed: "
                          << core_error_string << "\n";
                break;
            }

            // Cache pre-sideload machines
            if (last_assertion.sideload_block_number) {
                {
                    timed_sideload_cache.add(
                        std::make_unique<Machine>(*core_machine));

                    auto block = *last_assertion.sideload_block_number;
                    std::unique_lock<std::shared_mutex> lock(
                        lru_sideload_cache_mutex);
                    lru_sideload_cache[block] =
                        std::make_unique<Machine>(*core_machine);
                    // Remove any sideload_cache entries that are either more
                    // than sideload_cache_size blocks old, or in the future
                    // (meaning they've been reorg'd out).
                    auto it = lru_sideload_cache.begin();
                    while (it != lru_sideload_cache.end()) {
                        // Note: we check if block > sideload_cache_size here
                        // to prevent an underflow in the following check.
                        if ((block > coreConfig.lru_sideload_cache_size &&
                             it->first <
                                 block - coreConfig.lru_sideload_cache_size) ||
                            it->first > block) {
                            it = lru_sideload_cache.erase(it);
                        } else {
                            it++;
                        }
                    }
                }

                if (core_machine->machine_state.output.arb_gas_used >=
                    next_checkpoint_gas) {
                    // Save checkpoint after min_gas_checkpoint_frequency gas
                    // used
                    status = saveCheckpoint(tx);
                    if (!status.ok()) {
                        core_error_string = status.ToString();
                        std::cerr << "ArbCore checkpoint saving failed: "
                                  << core_error_string << "\n";
                        break;
                    }
                    next_checkpoint_gas =
                        core_machine->machine_state.output.arb_gas_used +
                        coreConfig.min_gas_checkpoint_frequency;
                    // Clear oldest cache and start populating next cache
                    std::cout
                        << "Last checkpoint gas used: "
                        << core_machine->machine_state.output.arb_gas_used
                        << ", L1 block: "
                        << core_machine->machine_state.output.l1_block_number
                        << ", L2 block: "
                        << *last_assertion.sideload_block_number << std::endl;
                    cache.nextCache();
                }

                // Machine was stopped to save sideload, update execConfig
                // and start machine back up where it stopped
                auto machine_success = core_machine->continueRunningMachine();
                if (!machine_success) {
                    core_error_string = "Error starting machine thread";
                    machine_error = true;
                    std::cerr << "ArbCore error: " << core_error_string << "\n";
                    break;
                }
            }

            status = tx.commit();
            if (!status.ok()) {
                core_error_string = status.ToString();
                machine_error = true;
                std::cerr << "ArbCore database update failed: "
                          << core_error_string << "\n";
                break;
            }
        }

        if (core_machine->status() == MachineThread::MACHINE_ABORTED) {
            // Just reset status so machine can be restarted
            core_machine->clearError();
        }

        if (core_machine->status() == MachineThread::MACHINE_NONE) {
            // Start execution of machine if new message available
            auto success = runMachineWithMessages(
                execConfig, coreConfig.message_process_count);
            if (!success) {
                break;
            }
        }

        for (size_t i = 0; i < logs_cursors.size(); i++) {
            if (logs_cursors[i].status == DataCursor::REQUESTED) {
                ReadTransaction tx(data_storage);
                handleLogsCursorRequested(tx, i, cache);
            }
        }

        if (save_checkpoint) {
            ReadWriteTransaction tx(data_storage);
            save_checkpoint_status = saveCheckpoint(tx);
            tx.commit();
            save_checkpoint = false;
        }

        if (!machineIdle() || message_data_status != MESSAGES_READY) {
            // Machine is already running or no new messages, so sleep for a
            // short while
            std::this_thread::sleep_for(std::chrono::milliseconds(5));
        }
    }

    std::cerr << "Exiting main ArbCore thread" << std::endl;

    // Error occurred, make sure machine stops cleanly
    core_machine->abortMachine();
}

bool ArbCore::runMachineWithMessages(MachineExecutionConfig& execConfig,
                                     size_t max_message_batch_size) {
    ReadSnapshotTransaction tx(data_storage);
    auto messages_result = readNextMessages(
        tx, core_machine->machine_state.output.fully_processed_inbox,
        max_message_batch_size);
    if (!messages_result.status.ok()) {
        core_error_string = messages_result.status.ToString();
        machine_error = true;
        std::cerr << "ArbCore failed getting message entry: "
                  << core_error_string << "\n";
        return false;
    }

    if (!messages_result.data.empty()) {
        execConfig.inbox_messages = messages_result.data;

        auto success = core_machine->runMachine(execConfig);
        if (!success) {
            core_error_string = "Error starting machine thread";
            machine_error = true;
            std::cerr << "ArbCore error: " << core_error_string << "\n";
            return false;
        }

        if (delete_checkpoints_before_message != uint256_t(0)) {
            /*
            deleteOldCheckpoints(delete_checkpoints_before_message,
                                 save_checkpoint_message_interval,
                                 ignore_checkpoints_after_message);
            */
            ignore_checkpoints_after_message = 0;
            save_checkpoint_message_interval = 0;
            delete_checkpoints_before_message = 0;
        }
    } else {
        // Machine all caught up, no messages to process
        machine_idle = true;
    }

    return true;
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
        auto value_result = saveValue(tx, val, *core_code);
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
    ReadSnapshotTransaction tx(data_storage);

    return getLogsNoLock(tx, index, count, valueCache);
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
        return {rocksdb::Status::OK(), {}};
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
        auto val_result = getValue(tx, hash, valueCache, false);
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

ValueResult<std::vector<std::vector<unsigned char>>>
ArbCore::getSequencerBatchItems(uint256_t index) const {
    ReadTransaction tx(data_storage);

    std::vector<unsigned char> first_key_vec;
    marshal_uint256_t(index, first_key_vec);
    auto first_key_slice = vecToSlice(first_key_vec);
    auto it = tx.sequencerBatchItemGetIterator(&first_key_slice);
    it->Seek(first_key_slice);

    std::vector<std::vector<unsigned char>> ret;
    while (it->Valid()) {
        auto key_ptr = reinterpret_cast<const unsigned char*>(it->key().data());
        auto value_ptr =
            reinterpret_cast<const unsigned char*>(it->value().data());

        std::vector<unsigned char> bytes(key_ptr, key_ptr + it->key().size());
        bytes.insert(bytes.end(), value_ptr, value_ptr + it->value().size());
        ret.push_back(bytes);

        it->Next();
    }

    return {it->status(), ret};
}

ValueResult<uint256_t> ArbCore::getSequencerBlockNumberAt(
    uint256_t sequence_number) const {
    ReadTransaction tx(data_storage);

    std::vector<unsigned char> first_key_vec;
    marshal_uint256_t(sequence_number, first_key_vec);
    auto first_key_slice = vecToSlice(first_key_vec);
    auto it = tx.sequencerBatchItemGetIterator(&first_key_slice);
    it->Seek(first_key_slice);

    std::vector<std::vector<unsigned char>> ret;
    while (it->Valid()) {
        auto key_ptr = reinterpret_cast<const unsigned char*>(it->key().data());
        auto value_ptr =
            reinterpret_cast<const unsigned char*>(it->value().data());
        auto value_end_ptr = value_ptr + it->value().size();

        auto seq_batch_item = deserializeSequencerBatchItem(
            extractUint256(key_ptr), value_ptr, value_end_ptr);
        if (seq_batch_item.sequencer_message) {
            auto it = seq_batch_item.sequencer_message->begin();
            auto block_num = extractInboxMessageBlockNumber(it);
            return {rocksdb::Status::OK(), block_num};
        }

        it->Next();
    }

    if (it->status().ok()) {
        return {rocksdb::Status::NotFound(), 0};
    } else {
        return {it->status(), 0};
    }
}

ValueResult<std::vector<unsigned char>> ArbCore::genInboxProof(
    uint256_t seq_num,
    uint256_t batch_index,
    uint256_t batch_end_count) const {
    ReadSnapshotTransaction tx(data_storage);

    auto message_res = getMessagesImpl(tx, seq_num, 1, std::nullopt);
    if (!message_res.status.ok()) {
        return {message_res.status, std::vector<unsigned char>()};
    }
    auto message_data = message_res.data[0].message;
    auto message = extractInboxMessage(message_data);
    auto proof = message.serializeForProof();

    marshal_uint256_t(batch_index, proof);

    uint256_t start = seq_num;
    bool recording_prev = false;
    SequencerBatchItem prev_item;
    if (start > 0) {
        start -= 1;
        recording_prev = true;
    }

    std::vector<unsigned char> start_vec;
    marshal_uint256_t(start, start_vec);
    auto start_slice = vecToSlice(start_vec);
    auto it = tx.sequencerBatchItemGetIterator();
    it->SeekForPrev(start_slice);

    bool first_item = true;
    while (it->Valid()) {
        auto key_ptr = reinterpret_cast<const unsigned char*>(it->key().data());
        auto value_ptr =
            reinterpret_cast<const unsigned char*>(it->value().data());
        auto value_end_ptr = value_ptr + it->value().size();
        auto item = deserializeSequencerBatchItem(extractUint256(key_ptr),
                                                  value_ptr, value_end_ptr);

        if (item.last_sequence_number >= batch_end_count) {
            // We've somehow skipped past the end of the batch,
            // meaning we disagree on where it ends (probably a reorg)
            return {rocksdb::Status::NotFound(), std::vector<unsigned char>()};
        }

        if (recording_prev) {
            prev_item = item;
            recording_prev = false;
            it->Next();
            continue;
        }

        if (first_item) {
            first_item = false;
            bool is_delayed = !item.sequencer_message;
            proof.push_back(static_cast<uint8_t>(is_delayed));
            marshal_uint256_t(prev_item.accumulator, proof);

            if (is_delayed) {
                if (prev_item.accumulator == 0) {
                    marshal_uint256_t(0, proof);
                } else {
                    marshal_uint256_t(prev_item.last_sequence_number + 1,
                                      proof);
                }
                marshal_uint256_t(prev_item.total_delayed_count, proof);
                marshal_uint256_t(item.total_delayed_count, proof);
            }
        } else {
            if (item.sequencer_message) {
                proof.push_back(0);
                auto seq_msg = extractInboxMessage(*item.sequencer_message);
                marshal_uint256_t(seq_msg.prefixHash(), proof);
                marshal_uint256_t(::hash(seq_msg.data), proof);
            } else {
                proof.push_back(1);
                marshal_uint256_t(prev_item.total_delayed_count, proof);
                marshal_uint256_t(item.total_delayed_count, proof);
            }
        }

        if (item.last_sequence_number + 1 == batch_end_count) {
            proof.push_back(2);
            return {rocksdb::Status::OK(), proof};
        }

        prev_item = item;
        it->Next();
    }

    if (!it->status().ok()) {
        return {it->status(), std::vector<unsigned char>()};
    }

    // We should've found the last item by this point
    return {rocksdb::Status::NotFound(), std::vector<unsigned char>()};
}

ValueResult<std::vector<std::vector<unsigned char>>> ArbCore::getMessages(
    uint256_t index,
    uint256_t count) const {
    ReadSnapshotTransaction tx(data_storage);

    auto result = getMessagesImpl(tx, index, count, std::nullopt);
    if (!result.status.ok()) {
        return {result.status, {}};
    }

    std::vector<std::vector<unsigned char>> bytes_vec;
    bytes_vec.reserve(result.data.size());
    for (auto& message_and_acc : result.data) {
        bytes_vec.push_back(std::move(message_and_acc.message));
    }

    return {result.status, bytes_vec};
}

ValueResult<std::vector<RawMessageInfo>> ArbCore::getMessagesImpl(
    const ReadConsistentTransaction& tx,
    uint256_t index,
    uint256_t count,
    std::optional<uint256_t> start_acc) const {
    std::vector<RawMessageInfo> messages;

    uint256_t start = index;
    bool needs_consistency_check = false;
    if (start > 0) {
        // Check the previous item to ensure the inbox state is valid
        start -= 1;
        needs_consistency_check = true;
    }

    std::vector<unsigned char> tmp;
    tmp.reserve(32 * 2);
    rocksdb::Slice seq_batch_lower_bound;
    {
        auto ptr = reinterpret_cast<const char*>(tmp.data() + tmp.size());
        marshal_uint256_t(start, tmp);
        seq_batch_lower_bound = {ptr, 32};
    }
    auto seq_batch_it =
        tx.sequencerBatchItemGetIterator(&seq_batch_lower_bound);

    uint256_t prev_delayed_count = 0;
    rocksdb::Slice delayed_msg_lower_bound;
    std::unique_ptr<rocksdb::Iterator> delayed_msg_it;
    for (seq_batch_it->Seek(seq_batch_lower_bound); seq_batch_it->Valid();
         seq_batch_it->Next()) {
        auto item_key_ptr =
            reinterpret_cast<const unsigned char*>(seq_batch_it->key().data());
        auto item_value_ptr = reinterpret_cast<const unsigned char*>(
            seq_batch_it->value().data());
        auto item_value_end_ptr = item_value_ptr + seq_batch_it->value().size();
        auto last_sequence_number = extractUint256(item_key_ptr);
        auto item = deserializeSequencerBatchItem(
            last_sequence_number, item_value_ptr, item_value_end_ptr);

        if (needs_consistency_check) {
            if (start_acc && item.accumulator != *start_acc) {
                return {rocksdb::Status::NotFound(), {}};
            }
            needs_consistency_check = false;
            if (count == 0) {
                // Skip some possible work attempting to read delayed messages
                break;
            }
            prev_delayed_count = item.total_delayed_count;
            if (item.last_sequence_number >= index) {
                // We are in the middle of a delayed batch
                assert(!item.sequencer_message);
                // Offset prev_delayed_count by the distance to the end of the
                // batch
                prev_delayed_count -= item.last_sequence_number + 1 - index;
            } else {
                // We are just after this batch item
                assert(item.last_sequence_number + 1 == index);
                continue;
            }
        }

        if (item.sequencer_message) {
            messages.emplace_back(std::move(*item.sequencer_message),
                                  item.last_sequence_number, item.accumulator);
            if (prev_delayed_count != item.total_delayed_count) {
                throw std::runtime_error(
                    "Sequencer batch item included both sequencer message and "
                    "delayed messages");
            }
        } else if (item.total_delayed_count > prev_delayed_count) {
            if (!delayed_msg_it) {
                {
                    auto ptr =
                        reinterpret_cast<const char*>(tmp.data() + tmp.size());
                    marshal_uint256_t(prev_delayed_count, tmp);
                    delayed_msg_lower_bound = {ptr, 32};
                }
                delayed_msg_it =
                    tx.delayedMessageGetIterator(&delayed_msg_lower_bound);
                delayed_msg_it->Seek(delayed_msg_lower_bound);
            }

            while (delayed_msg_it->Valid() &&
                   prev_delayed_count < item.total_delayed_count &&
                   messages.size() < count) {
                auto delayed_key_ptr = reinterpret_cast<const unsigned char*>(
                    delayed_msg_it->key().data());
                auto delayed_value_ptr = reinterpret_cast<const unsigned char*>(
                    delayed_msg_it->value().data());
                auto delayed_value_end_ptr =
                    delayed_value_ptr + delayed_msg_it->value().size();
                if (extractUint256(delayed_key_ptr) != prev_delayed_count) {
                    throw std::runtime_error(
                        "Got wrong delayed message from database");
                }
                auto delayed_message = deserializeDelayedMessage(
                    prev_delayed_count, delayed_value_ptr,
                    delayed_value_end_ptr);
                auto new_seq_num = prev_delayed_count | (uint256_t(1) << 255);
                messages.emplace_back(std::move(delayed_message.message),
                                      new_seq_num, item.accumulator);
                prev_delayed_count += 1;
                delayed_msg_it->Next();
            }

            if (!delayed_msg_it->status().ok()) {
                return {delayed_msg_it->status(), {}};
            }
            if (messages.size() < count &&
                prev_delayed_count != item.total_delayed_count) {
                throw std::runtime_error(
                    "Sequencer batch item referenced nonexistent delayed "
                    "messages");
            }
        } else {
            // This batch item does nothing?
            assert(false);
        }
        if (messages.size() >= count) {
            break;
        }
        assert(item.last_sequence_number + 1 == index + messages.size());
    }

    if (!seq_batch_it->status().ok()) {
        return {seq_batch_it->status(), {}};
    }
    if (needs_consistency_check) {
        return {rocksdb::Status::NotFound(), {}};
    }

    return {rocksdb::Status::OK(), messages};
}

ValueResult<uint256_t> ArbCore::getNextSequencerBatchItemAccumulator(
    const ReadTransaction& tx,
    uint256_t sequence_number) const {
    std::vector<unsigned char> tmp;
    tmp.reserve(32);
    rocksdb::Slice seq_batch_lower_bound;
    {
        auto ptr = reinterpret_cast<const char*>(tmp.data());
        marshal_uint256_t(sequence_number, tmp);
        seq_batch_lower_bound = {ptr, 32};
    }
    auto seq_batch_it =
        tx.sequencerBatchItemGetIterator(&seq_batch_lower_bound);
    seq_batch_it->Seek(seq_batch_lower_bound);
    if (!seq_batch_it->Valid()) {
        if (seq_batch_it->status().ok()) {
            return {rocksdb::Status::NotFound(), 0};
        } else {
            return {seq_batch_it->status(), 0};
        }
    }

    auto value_ptr =
        reinterpret_cast<const unsigned char*>(seq_batch_it->value().data());
    auto accumulator = extractUint256(value_ptr);
    return {rocksdb::Status::OK(), accumulator};
}

ValueResult<std::vector<std::vector<unsigned char>>> ArbCore::getSends(
    uint256_t index,
    uint256_t count) const {
    ReadSnapshotTransaction tx(data_storage);

    if (count == 0) {
        return {rocksdb::Status::OK(), {}};
    }

    // Check if attempting to get entries past current valid sends
    auto send_count = sendInsertedCountImpl(tx);
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

    return tx.sendGetVectorVector(key_slice, intx::narrow_cast<size_t>(count));
}

ValueResult<uint256_t> ArbCore::getInboxAcc(uint256_t index) {
    ReadTransaction tx(data_storage);

    auto result = getNextSequencerBatchItemAccumulator(tx, index);
    if (!result.status.ok()) {
        return {result.status, 0};
    }

    return {rocksdb::Status::OK(), result.data};
}

ValueResult<uint256_t> ArbCore::getDelayedInboxAcc(uint256_t index) {
    ReadTransaction tx(data_storage);

    return getDelayedInboxAccImpl(tx, index);
}

ValueResult<uint256_t> ArbCore::getDelayedInboxAccImpl(
    const ReadTransaction& tx,
    uint256_t index) {
    std::vector<unsigned char> key_vec;
    marshal_uint256_t(index, key_vec);
    auto key_slice = vecToSlice(key_vec);
    auto result = tx.delayedMessageGetVector(key_slice);
    if (!result.status.ok()) {
        return {result.status, 0};
    }

    auto it = result.data.begin();
    uint256_t acc = deserializeDelayedMessageAccumulator(it);

    return {rocksdb::Status::OK(), acc};
}

ValueResult<std::pair<uint256_t, uint256_t>> ArbCore::getInboxAccPair(
    uint256_t index1,
    uint256_t index2) {
    ReadSnapshotTransaction tx(data_storage);

    auto result1 = getNextSequencerBatchItemAccumulator(tx, index1);
    if (!result1.status.ok()) {
        return {result1.status, {0, 0}};
    }

    auto result2 = getNextSequencerBatchItemAccumulator(tx, index2);
    if (!result2.status.ok()) {
        return {result2.status, {0, 0}};
    }

    return {rocksdb::Status::OK(), {result1.data, result2.data}};
}

ValueResult<size_t> ArbCore::countMatchingBatchAccs(
    std::vector<std::pair<uint256_t, uint256_t>> seq_nums_and_accs) const {
    // TODO: validate sequence numbers lies on batch boundaries
    if (seq_nums_and_accs.empty()) {
        return {rocksdb::Status::OK(), 0};
    }

    size_t matching = 0;
    std::vector<unsigned char> tmp;
    tmp.reserve(32 * 2);
    uint256_t first_seq = seq_nums_and_accs[0].first;
    rocksdb::Slice lower_bound;
    {
        auto ptr = reinterpret_cast<const char*>(tmp.data());
        marshal_uint256_t(first_seq, tmp);
        lower_bound = {ptr, 32};
    }

    ReadTransaction tx(data_storage);
    auto it = tx.sequencerBatchItemGetIterator(&lower_bound);
    for (auto& seq_and_acc : seq_nums_and_accs) {
        if (seq_and_acc.first < first_seq) {
            throw std::runtime_error(
                "countMatchingBatchAccs received unsorted parameters");
        }
        rocksdb::Slice key_slice;
        {
            tmp.resize(32);
            auto ptr = reinterpret_cast<const char*>(tmp.data() + tmp.size());
            marshal_uint256_t(seq_and_acc.first, tmp);
            key_slice = {ptr, 32};
        }
        it->Seek(key_slice);
        if (!it->Valid()) {
            if (!it->status().ok()) {
                return {it->status(), 0};
            }
            break;
        }
        auto value_ptr =
            reinterpret_cast<const unsigned char*>(it->value().data());
        auto have_acc = extractUint256(value_ptr);
        if (have_acc == seq_and_acc.second) {
            matching++;
        } else {
            break;
        }
    }

    return {rocksdb::Status::OK(), matching};
}

ValueResult<uint256_t> ArbCore::getDelayedMessagesToSequence(
    uint256_t max_block_number) const {
    ReadSnapshotTransaction tx(data_storage);

    auto total_delayed_seq_res = totalDelayedMessagesSequencedImpl(tx);
    if (!total_delayed_seq_res.status.ok()) {
        return {total_delayed_seq_res.status, 0};
    }

    auto total_delayed_res = delayedMessageEntryInsertedCountImpl(tx);
    if (!total_delayed_res.status.ok()) {
        return {total_delayed_res.status, 0};
    }

    // Perform a binary search to find the last matching delayed message
    // After the search, low should be just after the last matching message
    auto low = total_delayed_seq_res.data;
    auto high = total_delayed_res.data;
    while (low != high) {
        auto mid = (low + high) / 2;
        std::vector<unsigned char> mid_vec;
        marshal_uint256_t(mid, mid_vec);
        auto mid_res = tx.delayedMessageGetVector(vecToSlice(mid_vec));
        if (!mid_res.status.ok()) {
            return {mid_res.status, 0};
        }
        auto mid_data_it = mid_res.data.begin();
        auto mid_block = deserializeDelayedMessageBlockNumber(mid_data_it);

        if (mid_block > max_block_number) {
            high = mid;
        } else {
            low = mid + 1;
        }
    }

    return {rocksdb::Status::OK(), low};
}

std::unique_ptr<Machine> ArbCore::getLastMachine() {
    std::shared_lock<std::shared_mutex> guard(last_machine_mutex);
    return std::make_unique<Machine>(*last_machine);
}

MachineOutput ArbCore::getLastMachineOutput() {
    std::shared_lock<std::shared_mutex> guard(last_machine_mutex);
    return last_machine->machine_state.output;
}

uint256_t ArbCore::machineMessagesRead() {
    std::shared_lock<std::shared_mutex> guard(last_machine_mutex);
    assert(last_machine);
    if (last_machine) {
        return last_machine->machine_state.output.fully_processed_inbox.count;
    } else {
        return 0;
    }
}

ValueResult<std::unique_ptr<ExecutionCursor>> ArbCore::getExecutionCursor(
    uint256_t total_gas_used) {
    std::unique_ptr<ExecutionCursor> execution_cursor;
    {
        ReadSnapshotTransaction tx(data_storage);

        auto closest_checkpoint =
            getClosestExecutionMachine(tx, total_gas_used);
        if (std::holds_alternative<rocksdb::Status>(closest_checkpoint)) {
            std::cerr << "No execution machine available" << std::endl;
            return {std::get<rocksdb::Status>(closest_checkpoint), nullptr};
        }

        execution_cursor = std::make_unique<ExecutionCursor>(
            std::get<ExecutionCursor>(closest_checkpoint));
    }

    auto status =
        advanceExecutionCursorImpl(*execution_cursor, total_gas_used, false,
                                   coreConfig.message_process_count);

    if (!status.ok()) {
        std::cerr << "Couldn't advance execution machine" << std::endl;
    }

    return {status, std::move(execution_cursor)};
}

rocksdb::Status ArbCore::advanceExecutionCursor(
    ExecutionCursor& execution_cursor,
    uint256_t max_gas,
    bool go_over_gas) {
    auto gas_target = execution_cursor.getOutput().arb_gas_used + max_gas;
    {
        ReadSnapshotTransaction tx(data_storage);

        auto closest_checkpoint = getClosestExecutionMachine(tx, gas_target);
        if (std::holds_alternative<rocksdb::Status>(closest_checkpoint)) {
            return std::get<rocksdb::Status>(closest_checkpoint);
        }

        auto checkpoint_cursor = std::get<ExecutionCursor>(closest_checkpoint);
        bool already_newer = false;
        if (execution_cursor.getOutput().arb_gas_used +
                coreConfig.checkpoint_load_gas_cost >
            checkpoint_cursor.getOutput().arb_gas_used) {
            // The existing execution cursor is far enough ahead that running it
            // up to the target gas will be cheaper than loading the checkpoint
            // from disk and running it. We just need to check that the
            // execution cursor is still valid (a reorg hasn't occurred).
            if (isValid(tx,
                        execution_cursor.getOutput().fully_processed_inbox)) {
                // Execution cursor machine still valid, so use it
                already_newer = true;
            }
        }

        if (!already_newer) {
            execution_cursor = std::move(checkpoint_cursor);
        }
    }

    return advanceExecutionCursorImpl(execution_cursor, gas_target, go_over_gas,
                                      coreConfig.message_process_count);
}

MachineState& resolveExecutionVariant(std::unique_ptr<Machine>& mach) {
    return mach->machine_state;
}

MachineStateKeys& resolveExecutionVariant(MachineStateKeys& mach) {
    return mach;
}

std::unique_ptr<Machine>& ArbCore::resolveExecutionCursorMachine(
    const ReadTransaction& tx,
    ExecutionCursor& execution_cursor) {
    if (std::holds_alternative<MachineStateKeys>(execution_cursor.machine)) {
        auto machine_state_keys =
            std::get<MachineStateKeys>(execution_cursor.machine);
        // Cache isn't very relevant as we're lazy loading
        auto cache = ValueCache(1, 0);
        execution_cursor.machine =
            getMachineUsingStateKeys<Machine>(tx, machine_state_keys, cache);
    }
    return std::get<std::unique_ptr<Machine>>(execution_cursor.machine);
}

std::unique_ptr<Machine> ArbCore::takeExecutionCursorMachineImpl(
    const ReadTransaction& tx,
    ExecutionCursor& execution_cursor) {
    auto mach = std::move(resolveExecutionCursorMachine(tx, execution_cursor));
    execution_cursor.machine = MachineStateKeys{mach->machine_state};
    return mach;
}

std::unique_ptr<Machine> ArbCore::takeExecutionCursorMachine(
    ExecutionCursor& execution_cursor) {
    ReadSnapshotTransaction tx(data_storage);
    return takeExecutionCursorMachineImpl(tx, execution_cursor);
}

rocksdb::Status ArbCore::advanceExecutionCursorImpl(
    ExecutionCursor& execution_cursor,
    uint256_t total_gas_used,
    bool go_over_gas,
    size_t message_group_size) {
    auto handle_reorg = true;
    size_t reorg_attempts = 0;
    while (handle_reorg) {
        handle_reorg = false;
        if (reorg_attempts > 0) {
            if (reorg_attempts % 4 == 0) {
                std::cerr
                    << "Execution cursor has attempted to handle "
                    << reorg_attempts
                    << " reorgs. Checkpoints may be inconsistent with messages."
                    << std::endl;
            }
            assert(false);
            std::this_thread::sleep_for(std::chrono::milliseconds(250));
            if (reorg_attempts >= 16) {
                return rocksdb::Status::Busy();
            }
        }
        reorg_attempts++;

        while (true) {
            // Run machine until specified gas is reached
            MachineExecutionConfig execConfig;
            execConfig.max_gas = total_gas_used;
            execConfig.go_over_gas = go_over_gas;

            {
                ReadSnapshotTransaction tx(data_storage);

                auto& mach =
                    resolveExecutionCursorMachine(tx, execution_cursor);

                uint256_t gas_used = execution_cursor.getOutput().arb_gas_used;
                if (gas_used == total_gas_used) {
                    break;
                } else if (go_over_gas && gas_used > total_gas_used) {
                    break;
                } else if (!go_over_gas &&
                           gas_used + mach->machine_state.nextGasCost() >
                               total_gas_used) {
                    break;
                }

                auto get_messages_result = readNextMessages(
                    tx, execution_cursor.getOutput().fully_processed_inbox,
                    message_group_size);
                if (get_messages_result.status.IsNotFound()) {
                    // Reorg occurred, need to recreate machine
                    handle_reorg = true;
                    break;
                }
                if (!get_messages_result.status.ok()) {
                    std::cout << "Error getting messages for execution cursor"
                              << std::endl;
                    return get_messages_result.status;
                }
                execConfig.inbox_messages = std::move(get_messages_result.data);
            }

            auto& mach =
                std::get<std::unique_ptr<Machine>>(execution_cursor.machine);
            mach->machine_state.context = AssertionContext(execConfig);
            auto assertion = mach->run();
            if (assertion.gas_count == 0) {
                break;
            }
        }

        if (handle_reorg) {
            ReadSnapshotTransaction tx(data_storage);

            auto closest_checkpoint =
                getClosestExecutionMachine(tx, total_gas_used);
            if (std::holds_alternative<rocksdb::Status>(closest_checkpoint)) {
                std::cerr << "No execution machine available" << std::endl;
                return std::get<rocksdb::Status>(closest_checkpoint);
            }
            execution_cursor =
                std::move(std::get<ExecutionCursor>(closest_checkpoint));
        }
    }

    return rocksdb::Status::OK();
}

std::variant<rocksdb::Status, ExecutionCursor>
ArbCore::getClosestExecutionMachine(ReadTransaction& tx,
                                    const uint256_t& total_gas_used) {
    // Check if the cache contains a usable machine
    {
        std::shared_lock<std::shared_mutex> guard(old_machine_cache_mutex);
        auto cache_it = old_machine_cache.upper_bound(total_gas_used);
        if (cache_it != old_machine_cache.begin()) {
            cache_it--;
            return ExecutionCursor(
                std::make_unique<Machine>(*cache_it->second));
        }
    }

    // Find the nearest database checkpoint
    auto target_gas_used = total_gas_used;
    while (true) {
        const std::lock_guard<std::mutex> lock(core_reorg_mutex);
        auto checkpoint_result = getCheckpointUsingGas(tx, target_gas_used);

        if (std::holds_alternative<rocksdb::Status>(checkpoint_result)) {
            return std::get<rocksdb::Status>(checkpoint_result);
        }

        return ExecutionCursor(std::get<MachineStateKeys>(checkpoint_result));
    }
}

ValueResult<std::vector<MachineMessage>> ArbCore::readNextMessages(
    const ReadConsistentTransaction& tx,
    const InboxState& fully_processed_inbox,
    size_t count) const {
    std::vector<MachineMessage> messages;
    messages.reserve(count);

    auto raw_result = getMessagesImpl(tx, fully_processed_inbox.count, count,
                                      fully_processed_inbox.accumulator);
    if (!raw_result.status.ok()) {
        return {raw_result.status, messages};
    }

    for (auto& raw_message : raw_result.data) {
        auto message = extractInboxMessage(raw_message.message);
        message.inbox_sequence_number = raw_message.sequence_number;
        messages.emplace_back(message, raw_message.accumulator);
    }

    return {rocksdb::Status::OK(), messages};
}

bool ArbCore::isValid(const ReadTransaction& tx,
                      const InboxState& fully_processed_inbox) const {
    if (fully_processed_inbox.count == 0) {
        return true;
    }
    auto result = getNextSequencerBatchItemAccumulator(
        tx, fully_processed_inbox.count - 1);
    return result.status.ok() &&
           result.data == fully_processed_inbox.accumulator;
}

ValueResult<uint256_t> ArbCore::logInsertedCount() const {
    ReadTransaction tx(data_storage);

    return logInsertedCountImpl(tx);
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
    ReadTransaction tx(data_storage);

    return sendInsertedCountImpl(tx);
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

ValueResult<uint256_t> ArbCore::schemaVersion(ReadTransaction& tx) const {
    return tx.stateGetUint256(vecToSlice(schema_version_key));
}
rocksdb::Status ArbCore::updateSchemaVersion(ReadWriteTransaction& tx,
                                             const uint256_t& schema_version) {
    std::vector<unsigned char> value;
    marshal_uint256_t(schema_version, value);

    return tx.statePut(vecToSlice(schema_version_key), vecToSlice(value));
}

ValueResult<uint256_t> ArbCore::messageEntryInsertedCount() const {
    ReadTransaction tx(data_storage);

    return messageEntryInsertedCountImpl(tx);
}

ValueResult<uint256_t> ArbCore::delayedMessageEntryInsertedCount() const {
    ReadTransaction tx(data_storage);

    return delayedMessageEntryInsertedCountImpl(tx);
}

ValueResult<uint256_t> ArbCore::messageEntryInsertedCountImpl(
    const ReadTransaction& tx) const {
    auto it = tx.sequencerBatchItemGetIterator();
    it->SeekToLast();
    if (it->Valid()) {
        auto key_ptr = reinterpret_cast<const unsigned char*>(it->key().data());
        auto seq_num = extractUint256(key_ptr);
        return {rocksdb::Status::OK(), seq_num + 1};
    } else {
        return {it->status(), 0};
    }
}

ValueResult<uint256_t> ArbCore::delayedMessageEntryInsertedCountImpl(
    const ReadTransaction& tx) const {
    auto it = tx.delayedMessageGetIterator();
    it->SeekToLast();
    if (it->Valid()) {
        auto key_ptr = reinterpret_cast<const unsigned char*>(it->key().data());
        auto seq_num = extractUint256(key_ptr);
        return {rocksdb::Status::OK(), seq_num + 1};
    } else {
        return {it->status(), 0};
    }
}

ValueResult<uint256_t> ArbCore::totalDelayedMessagesSequenced() const {
    ReadTransaction tx(data_storage);

    return totalDelayedMessagesSequencedImpl(tx);
}

ValueResult<uint256_t> ArbCore::totalDelayedMessagesSequencedImpl(
    const ReadTransaction& tx) const {
    auto it = tx.sequencerBatchItemGetIterator();
    it->SeekToLast();
    if (it->Valid()) {
        auto key_ptr = reinterpret_cast<const unsigned char*>(it->key().data());
        auto value_ptr =
            reinterpret_cast<const unsigned char*>(it->value().data());
        auto value_end_ptr = value_ptr + it->value().size();
        auto item = deserializeSequencerBatchItem(extractUint256(key_ptr),
                                                  value_ptr, value_end_ptr);
        return {rocksdb::Status::OK(), item.total_delayed_count};
    } else {
        return {it->status(), 0};
    }
}

// addMessages adds the next batch of messages to the machine.  If there is
// a reorg, the amount of gas used by the last checkpoint is returned.
ValueResult<std::optional<uint256_t>> ArbCore::addMessages(
    const ArbCore::message_data_struct& data,
    ValueCache& cache) {
    std::optional<uint256_t> last_gas_used;
    std::vector<std::pair<SequencerBatchItem, rocksdb::Slice>> seq_batch_items;
    for (auto& bytes : data.sequencer_batch_items) {
        auto it = bytes.begin();
        auto last_seq_num = extractUint256(it);
        rocksdb::Slice value_slice = {reinterpret_cast<const char*>(&*it),
                                      static_cast<size_t>(bytes.end() - it)};
        auto seq_message =
            deserializeSequencerBatchItem(last_seq_num, it, bytes.end());
        seq_batch_items.emplace_back(seq_message, value_slice);
    }

    std::vector<std::pair<DelayedMessage, rocksdb::Slice>> delayed_messages;
    for (auto& bytes : data.delayed_messages) {
        auto it = bytes.begin();
        auto seq_num = extractUint256(it);
        rocksdb::Slice value_slice = {reinterpret_cast<const char*>(&*it),
                                      static_cast<size_t>(bytes.end() - it)};
        auto delayed_message =
            deserializeDelayedMessage(seq_num, it, bytes.end());
        delayed_messages.emplace_back(delayed_message, value_slice);
    }

    std::optional<uint256_t> reorging_to_count;
    SequencerBatchItem prev_item;
    {
        std::vector<unsigned char> tmp;
        tmp.reserve(32);
        ReadWriteTransaction tx(data_storage);
        size_t duplicate_seq_batch_items = 0;

        if (data.reorg_batch_items) {
            tmp.clear();
            marshal_uint256_t(*data.reorg_batch_items, tmp);
            rocksdb::Slice start_slice = vecToSlice(tmp);

            auto seq_batch_it = tx.sequencerBatchItemGetIterator(&start_slice);
            seq_batch_it->Seek(start_slice);
            bool deleted_any = false;
            while (seq_batch_it->Valid()) {
                auto status = tx.sequencerBatchItemDelete(seq_batch_it->key());
                if (!status.ok()) {
                    return {status, std::nullopt};
                }
                deleted_any = true;
                seq_batch_it->Next();
            }
            if (!seq_batch_it->status().ok()) {
                return {seq_batch_it->status(), std::nullopt};
            }
            if (deleted_any) {
                reorging_to_count = *data.reorg_batch_items;
            }
        }

        if (!seq_batch_items.empty()) {
            uint256_t start = message_data.previous_message_count;
            bool checking_prev = false;
            if (start > 0) {
                start -= 1;
                checking_prev = true;
            }
            tmp.clear();
            marshal_uint256_t(start, tmp);
            rocksdb::Slice start_slice = vecToSlice(tmp);
            auto seq_batch_it = tx.sequencerBatchItemGetIterator();

            if (checking_prev) {
                seq_batch_it->Seek(start_slice);
                if (!seq_batch_it->status().ok()) {
                    return {seq_batch_it->status(), std::nullopt};
                }
                if (!seq_batch_it->Valid()) {
                    std::cerr << "addMessages: previous batch item not found"
                              << std::endl;
                    return {rocksdb::Status::NotFound(), std::nullopt};
                }
                auto key_ptr = reinterpret_cast<const unsigned char*>(
                    seq_batch_it->key().data());
                auto value_ptr = reinterpret_cast<const unsigned char*>(
                    seq_batch_it->value().data());
                auto value_end_ptr = value_ptr + seq_batch_it->value().size();
                auto db_item = deserializeSequencerBatchItem(
                    extractUint256(key_ptr), value_ptr, value_end_ptr);

                if (db_item.last_sequence_number != start) {
                    throw std::runtime_error(
                        "previous_message_count didn't fall on batch item "
                        "boundary");
                }
                if (db_item.accumulator != data.previous_batch_acc) {
                    throw std::runtime_error("prev_batch_acc didn't match");
                }
                prev_item = db_item;
                seq_batch_it->Next();
            } else {
                seq_batch_it->Seek(start_slice);
            }

            for (auto& item_and_bytes : seq_batch_items) {
                if (!seq_batch_it->Valid()) {
                    if (!seq_batch_it->status().ok()) {
                        return {seq_batch_it->status(), std::nullopt};
                    }
                    break;
                }

                auto& item = item_and_bytes.first;
                auto value_ptr = reinterpret_cast<const unsigned char*>(
                    seq_batch_it->value().data());
                auto accumulator =
                    deserializeSequencerBatchItemAccumulator(value_ptr);

                if (accumulator != item.accumulator) {
                    std::cerr << "INBOX FORCED REORG at sequence number "
                              << item.last_sequence_number << std::endl
                              << "Previous accumulator: " << accumulator
                              << std::endl
                              << "New accumulator:      " << item.accumulator
                              << std::endl;
                    if (item.last_sequence_number == 0) {
                        reorging_to_count = 0;
                    } else {
                        reorging_to_count = prev_item.last_sequence_number + 1;
                    }

                    // Ideally we would use DeleteRange here, but as far as I
                    // can tell it isn't supported on transactions.
                    while (seq_batch_it->Valid()) {
                        auto status =
                            tx.sequencerBatchItemDelete(seq_batch_it->key());
                        if (!status.ok()) {
                            return {status, std::nullopt};
                        }
                        seq_batch_it->Next();
                        if (!seq_batch_it->status().ok()) {
                            return {seq_batch_it->status(), std::nullopt};
                        }
                    }
                    break;
                }
                prev_item = item;
                duplicate_seq_batch_items++;
                seq_batch_it->Next();
            }
        }

        if (!delayed_messages.empty()) {
            auto delayed_msg_seq_res = totalDelayedMessagesSequencedImpl(tx);
            if (!delayed_msg_seq_res.status.ok()) {
                return {delayed_msg_seq_res.status, std::nullopt};
            }
            uint256_t start = delayed_messages[0].first.delayed_sequence_number;
            bool checking_prev = false;
            if (start > 0) {
                start -= 1;
                checking_prev = true;
            }
            tmp.clear();
            marshal_uint256_t(start, tmp);
            rocksdb::Slice lower_bound = vecToSlice(tmp);

            auto delayed_it = tx.delayedMessageGetIterator(&lower_bound);
            delayed_it->Seek(lower_bound);
            uint256_t prev_acc = 0;
            auto inserting = false;
            for (auto& item : delayed_messages) {
                auto& message = item.first;
                auto& value_slice = item.second;

                if (!inserting && !delayed_it->Valid()) {
                    if (!delayed_it->status().ok()) {
                        return {delayed_it->status(), std::nullopt};
                    }
                    if (checking_prev) {
                        throw std::runtime_error(
                            "Previous delayed message not found");
                    }
                    inserting = true;
                }

                if (!inserting) {
                    auto value_ptr = reinterpret_cast<const unsigned char*>(
                        delayed_it->value().data());
                    auto db_accumulator =
                        deserializeDelayedMessageAccumulator(value_ptr);

                    if (checking_prev) {
                        prev_acc = db_accumulator;
                        checking_prev = false;

                        delayed_it->Next();
                        if (delayed_it->Valid()) {
                            value_ptr = reinterpret_cast<const unsigned char*>(
                                delayed_it->value().data());
                            db_accumulator =
                                deserializeDelayedMessageAccumulator(value_ptr);
                        } else {
                            if (!delayed_it->status().ok()) {
                                return {delayed_it->status(), std::nullopt};
                            }
                            inserting = true;
                        }
                    }

                    if (message.delayed_accumulator == db_accumulator) {
                        prev_acc = db_accumulator;
                        delayed_it->Next();
                    } else {
                        if (message.delayed_sequence_number <
                            delayed_msg_seq_res.data) {
                            throw std::runtime_error(
                                "Attempted to reorg sequenced delayed "
                                "messages");
                        }
                        // Ideally we would use DeleteRange here, but as far as
                        // I can tell it isn't supported on transactions.
                        while (delayed_it->Valid()) {
                            auto status =
                                tx.delayedMessageDelete(delayed_it->key());
                            if (!status.ok()) {
                                return {status, std::nullopt};
                            }
                            delayed_it->Next();
                            if (!delayed_it->status().ok()) {
                                return {delayed_it->status(), std::nullopt};
                            }
                        }
                        inserting = true;
                    }
                }

                if (inserting) {
                    auto expected_acc = hash_inbox(prev_acc, message.message);
                    if (expected_acc != message.delayed_accumulator) {
                        throw std::runtime_error(
                            "Unexpected delayed accumulator");
                    }
                    prev_acc = message.delayed_accumulator;

                    std::vector<unsigned char> key_vec;
                    marshal_uint256_t(message.delayed_sequence_number, key_vec);
                    auto key_slice = vecToSlice(key_vec);
                    auto status = tx.delayedMessagePut(key_slice, value_slice);
                    if (!status.ok()) {
                        return {status, std::nullopt};
                    }
                }
            }
        }

        for (size_t i = duplicate_seq_batch_items; i < seq_batch_items.size();
             i++) {
            auto& item_and_slice = seq_batch_items[i];
            auto& item = item_and_slice.first;
            uint256_t delayed_acc;
            uint256_t expected_last_seq_num = 0;
            if (prev_item.accumulator != 0) {
                expected_last_seq_num = prev_item.last_sequence_number + 1;
            }
            if (item.total_delayed_count > prev_item.total_delayed_count) {
                expected_last_seq_num += item.total_delayed_count -
                                         prev_item.total_delayed_count - 1;
                if (item.sequencer_message) {
                    throw std::runtime_error(
                        "Attempted to add sequencer batch item with both "
                        "sequencer message and delayed messages");
                }
                auto res =
                    getDelayedInboxAccImpl(tx, item.total_delayed_count - 1);
                if (!res.status.ok()) {
                    std::cerr << "ArbCore failed to lookup delayed message "
                                 "accumulator"
                              << std::endl;
                    return {res.status, std::nullopt};
                }
                delayed_acc = res.data;
            } else if (item.total_delayed_count <
                       prev_item.total_delayed_count) {
                throw std::runtime_error(
                    "Attempted to add sequencer batch item that decreases "
                    "total messages read");
            }
            if (item.last_sequence_number != expected_last_seq_num) {
                throw std::runtime_error(
                    "Attempted to add sequencer batch item with unexpected "
                    "sequence number");
            }
            auto expected_acc = item.computeAccumulator(
                prev_item.accumulator, prev_item.total_delayed_count,
                delayed_acc);
            if (item.accumulator == 0) {
                item.accumulator = expected_acc;
                std::vector<unsigned char> acc_bytes;
                marshal_uint256_t(item.accumulator, acc_bytes);
                std::copy(acc_bytes.begin(), acc_bytes.end(),
                          const_cast<char*>(item_and_slice.second.data()));
            } else if (item.accumulator != expected_acc) {
                throw std::runtime_error(
                    "Sequencer batch item accumulator didn't match recomputed "
                    "value");
            }
            prev_item = item;
            std::vector<unsigned char> key_vec;
            marshal_uint256_t(item.last_sequence_number, key_vec);
            auto status = tx.sequencerBatchItemPut(vecToSlice(key_vec),
                                                   item_and_slice.second);
            if (!status.ok()) {
                return {status, std::nullopt};
            }
        }
        auto status = tx.commit();
        if (!status.ok()) {
            return {status, std::nullopt};
        }
    }
    if (reorging_to_count) {
        auto status =
            reorgToMessageCountOrBefore(*reorging_to_count, false, cache);
        if (!status.ok()) {
            return {status, std::nullopt};
        }

        last_gas_used = maxCheckpointGas();
    }
    return {rocksdb::Status::OK(), last_gas_used};
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
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        throw std::runtime_error("Invalid logsCursor index");
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    logs_cursors[cursor_index].data.clear();

    // Provide requested logs
    auto log_inserted_count = logInsertedCountImpl(tx);
    if (!log_inserted_count.status.ok()) {
        logs_cursors[cursor_index].error_string =
            log_inserted_count.status.ToString();
        logs_cursors[cursor_index].status = DataCursor::ERROR;

        std::cerr << "logscursor index " << cursor_index
                  << " error getting inserted count: "
                  << log_inserted_count.status.ToString() << std::endl;
        return;
    }

    auto current_count_result =
        logsCursorGetCurrentTotalCount(tx, cursor_index);
    if (!current_count_result.status.ok()) {
        logs_cursors[cursor_index].error_string =
            current_count_result.status.ToString();
        logs_cursors[cursor_index].status = DataCursor::ERROR;

        std::cerr << "logscursor index" << cursor_index
                  << " error getting cursor current total count: "
                  << current_count_result.status.ToString() << std::endl;
        return;
    }

    if (current_count_result.data == log_inserted_count.data) {
        // No new messages

        if (!logs_cursors[cursor_index].deleted_data.empty()) {
            // No new messages, but there are deleted messages to process
            logs_cursors[cursor_index].status = DataCursor::READY;
        }

        return;
    }
    if (current_count_result.data > log_inserted_count.data) {
        logs_cursors[cursor_index].error_string =
            "current_count_result greater than log_inserted_count";
        logs_cursors[cursor_index].status = DataCursor::ERROR;

        std::cerr << "handleLogsCursor current count: "
                  << current_count_result.data << " > "
                  << " log inserted count: " << log_inserted_count.data
                  << std::endl;
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
}

// handleLogsCursorReorg must be called before logs are deleted.
// Note that cursor reorg never adds new messages, but might add deleted
// messages.
rocksdb::Status ArbCore::handleLogsCursorReorg(size_t cursor_index,
                                               uint256_t log_count,
                                               ValueCache& cache) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        throw std::runtime_error("Invalid logsCursor index");
    }

    ReadWriteTransaction tx(data_storage);

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

    if (logs_cursors[cursor_index].status == DataCursor::READY &&
        logs_cursors[cursor_index].data.empty() &&
        logs_cursors[cursor_index].deleted_data.empty()) {
        logs_cursors[cursor_index].status = DataCursor::REQUESTED;
    }

    return tx.commit();
}

bool ArbCore::logsCursorRequest(size_t cursor_index, uint256_t count) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        throw std::runtime_error("Invalid logsCursor index");
    }

    if (logs_cursors[cursor_index].status != DataCursor::EMPTY) {
        return false;
    }

    logs_cursors[cursor_index].number_requested = count;
    logs_cursors[cursor_index].status = DataCursor::REQUESTED;

    return true;
}

ValueResult<ArbCore::logscursor_logs> ArbCore::logsCursorGetLogs(
    size_t cursor_index) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        throw std::runtime_error("Invalid logsCursor index");
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    DataCursor::status_enum status = logs_cursors[cursor_index].status;
    if (status == DataCursor::REQUESTED) {
        // No new logs yet
        return {rocksdb::Status::TryAgain(), {}};
    } else if (status != DataCursor::READY) {
        throw std::runtime_error("Unexpected logsCursor status " + status);
    }

    ReadTransaction tx(data_storage);
    auto current_count_result =
        logsCursorGetCurrentTotalCount(tx, cursor_index);
    if (!current_count_result.status.ok()) {
        std::cerr << "logs cursor " << cursor_index
                  << " unable to get current total count: "
                  << current_count_result.status.ToString() << "\n";
        return {current_count_result.status, {}};
    }

    logs_cursors[cursor_index].pending_total_count =
        current_count_result.data + logs_cursors[cursor_index].data.size();

    ArbCore::logscursor_logs logs{};
    logs.first_log_index = current_count_result.data;
    logs.logs = std::move(logs_cursors[cursor_index].data);
    logs.deleted_logs = std::move(logs_cursors[cursor_index].deleted_data);
    logs_cursors[cursor_index].data.clear();
    logs_cursors[cursor_index].deleted_data.clear();
    logs_cursors[cursor_index].status = DataCursor::DELIVERED;

    return {rocksdb::Status::OK(), std::move(logs)};
}

bool ArbCore::logsCursorConfirmReceived(size_t cursor_index) {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        throw std::runtime_error("Invalid logsCursor index");
    }

    const std::lock_guard<std::mutex> lock(
        logs_cursors[cursor_index].reorg_mutex);

    if (logs_cursors[cursor_index].status != DataCursor::DELIVERED) {
        logs_cursors[cursor_index].error_string =
            "logsCursorConfirmReceived called at wrong state";
        std::cerr << "logsCursorConfirmReceived called at wrong state: "
                  << logs_cursors[cursor_index].status << "\n";
        logs_cursors[cursor_index].status = DataCursor::ERROR;
        return false;
    }

    if (!logs_cursors[cursor_index].deleted_data.empty()) {
        // Deleted logs were added since the last time logsCursorGetLogs was
        // called, so need to call logsCursorGetLogs again
        logs_cursors[cursor_index].status = DataCursor::READY;
        return false;
    }

    ReadWriteTransaction tx(data_storage);
    auto status = logsCursorSaveCurrentTotalCount(
        tx, cursor_index, logs_cursors[cursor_index].pending_total_count);
    tx.commit();

    logs_cursors[cursor_index].status = DataCursor::EMPTY;

    return true;
}

bool ArbCore::logsCursorCheckError(size_t cursor_index) const {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        throw std::runtime_error("Invalid logsCursor index");
    }

    return logs_cursors[cursor_index].status == DataCursor::ERROR;
}

ValueResult<uint256_t> ArbCore::logsCursorPosition(size_t cursor_index) const {
    if (cursor_index >= logs_cursors.size()) {
        std::cerr << "Invalid logsCursor index: " << cursor_index << "\n";
        throw std::runtime_error("Invalid logsCursor index");
    }

    ReadTransaction tx(data_storage);
    return logsCursorGetCurrentTotalCount(tx, cursor_index);
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
        return "logsCursorClearError called when status not ERROR";
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
    const ReadTransaction& tx,
    size_t cursor_index) const {
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

    if (!it->Valid()) {
        return {rocksdb::Status::NotFound(), 0};
    }

    auto value_slice = it->value();

    return {s, intx::be::unsafe::load<uint256_t>(
                   reinterpret_cast<const unsigned char*>(value_slice.data()))};
}

rocksdb::Status ArbCore::deleteSideloadsStartingAt(
    ReadWriteTransaction& tx,
    const uint256_t& block_number) {
    // Clear the cache
    timed_sideload_cache.reorg(block_number);
    {
        std::unique_lock<std::shared_mutex> guard(lru_sideload_cache_mutex);
        auto it = lru_sideload_cache.lower_bound(block_number);
        while (it != lru_sideload_cache.end()) {
            it = lru_sideload_cache.erase(it);
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
    bool allow_slow_lookup) {
    // Check the cache
    if (auto cached_machine = timed_sideload_cache.get(block_number)) {
        return {rocksdb::Status::OK(), std::move(cached_machine)};
    }
    {
        std::shared_lock<std::shared_mutex> lock(lru_sideload_cache_mutex);
        // Look for the first value after the value we want
        auto it = lru_sideload_cache.upper_bound(block_number);
        if (it != lru_sideload_cache.begin()) {
            // Go back a value to find the one we want
            it--;
            return {rocksdb::Status::OK(),
                    std::make_unique<Machine>(*it->second)};
        }
    }

    if (!allow_slow_lookup) {
        // Don't try to query database
        return {rocksdb::Status::OK(), nullptr};
    }

    uint256_t gas_target;
    std::unique_ptr<ExecutionCursor> execution_cursor;
    {
        // Not found in cache, try the DB
        ReadSnapshotTransaction tx(data_storage);
        auto position_res = getSideloadPosition(tx, block_number);
        if (!position_res.status.ok()) {
            return {position_res.status, std::unique_ptr<Machine>(nullptr)};
        }

        auto closest_checkpoint =
            getClosestExecutionMachine(tx, position_res.data);
        if (std::holds_alternative<rocksdb::Status>(closest_checkpoint)) {
            return {std::get<rocksdb::Status>(closest_checkpoint), nullptr};
        }

        gas_target = position_res.data;
        execution_cursor = std::make_unique<ExecutionCursor>(
            std::get<ExecutionCursor>(closest_checkpoint));
    }

    auto status = advanceExecutionCursorImpl(
        *execution_cursor, gas_target, false, coreConfig.message_process_count);

    ReadSnapshotTransaction tx(data_storage);
    return {status, takeExecutionCursorMachineImpl(tx, *execution_cursor)};
}
