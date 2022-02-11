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

#include <sys/stat.h>
#include <ethash/keccak.hpp>
#include <filesystem>
#include <iomanip>
#include <set>
#include <sstream>
#include <utility>
#include <vector>

#ifdef __linux__
#include <execinfo.h>
#include <sys/prctl.h>
#include <csignal>
#endif

namespace {
constexpr uint256_t arbcore_schema_version = 3;
constexpr auto pruning_mode_key = std::array<char, 1>{-59};
constexpr auto log_inserted_key = std::array<char, 1>{-60};
constexpr auto send_inserted_key = std::array<char, 1>{-62};
constexpr auto schema_version_key = std::array<char, 1>{-64};
constexpr auto logscursor_current_prefix = std::array<char, 1>{-120};

template <class T>
void printResult(const char* msg, const ValueResult<T>& result) {
    std::cout << msg;
    if (!result.status.ok()) {
        std::cout << result.status.ToString() << "\n";
    } else {
        std::cout << result.data << "\n";
    }
}

void printMachineOutputInfo(const std::string& msg,
                            MachineOutput& machine_output) {
    std::cout << msg << ",  total gas used: " << machine_output.arb_gas_used
              << ", L1 block: " << machine_output.l1_block_number
              << ", L2 block: " << machine_output.l2_block_number
              << ", log count: " << machine_output.log_count
              << ", messages count: "
              << machine_output.fully_processed_inbox.count << ", timestamp: "
              << std::put_time(
                     localtime((time_t*)&machine_output.last_inbox_timestamp),
                     "%c")
              << "\n";
}

void printCheckpointResult(
    const char* msg,
    const std::variant<rocksdb::Status, CheckpointVariant>& result) {
    if (std::holds_alternative<rocksdb::Status>(result)) {
        std::cout << std::get<rocksdb::Status>(result).ToString() << "\n";
    } else {
        auto output = getMachineOutput(std::get<CheckpointVariant>(result));
        printMachineOutputInfo(msg, output);
    }
}
}  // namespace

ArbCore::ArbCore(std::shared_ptr<DataStorage> data_storage_,
                 ArbCoreConfig coreConfig_)
    : coreConfig(std::move(coreConfig_)),
      data_storage(std::move(data_storage_)),
      core_code(std::make_shared<CoreCode>(getNextSegmentID(data_storage))),
      combined_machine_cache(coreConfig) {
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
#ifdef __linux__
        core_pthread = std::nullopt;
#endif
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
    return ValueLoader{std::make_unique<CoreValueLoader>(
        data_storage, core_code, ValueCache{1, 0})};
}

void ArbCore::printDatabaseMetadata() {
    ReadTransaction tx(data_storage);
    printResult("schema version:                   ", schemaVersion(tx));
    printResult("log inserted count:               ", logInsertedCountImpl(tx));
    for (size_t i = 0; i < logs_cursors.size(); i++) {
        std::cout << "log cursor " << i << " count:               ";
        printResult("", logsCursorGetCurrentTotalCount(tx, i));
    }
    printResult("send inserted count:              ",
                sendInsertedCountImpl(tx));
    printResult("message inserted count:           ",
                messageEntryInsertedCountImpl(tx));
    printResult("delayed message inserted count:   ",
                delayedMessageEntryInsertedCountImpl(tx));
    printResult("total delayed messages sequenced: ",
                totalDelayedMessagesSequencedImpl(tx));
    printResult("pruning mode:                     ", pruningMode(tx));
    std::cout << "Checkpoint count:                 " << countCheckpoints(tx)
              << "\n";

    uint256_t genesis_checkpoint = 1;
    printCheckpointResult("genesis checkpoint",
                          getCheckpointNumber(tx, genesis_checkpoint));

    uint256_t first_checkpoint = 2;
    printCheckpointResult("first   checkpoint",
                          getCheckpointNumber(tx, first_checkpoint));

    printCheckpointResult("last    checkpoint", getLastCheckpoint(tx));
}

InitializeResult ArbCore::applyConfig() {
    ValueCache cache{1, 0};

    bool database_exists = false;
    ValueResult<std::string> pruning_mode_result;
    {
        // Validate database schema version
        ReadTransaction tx(data_storage);

        auto logs_result = logInsertedCountImpl(tx);
        if (logs_result.status.ok()) {
            // Database exists
            database_exists = true;
        }

        auto schema_result = schemaVersion(tx);
        if (!schema_result.status.ok()) {
            if (database_exists) {
                // Old database that does not have schema version
                std::cerr << "Error getting schema version: "
                          << schema_result.status.ToString()
                          << ", delete database and try again" << std::endl;
                return {rocksdb::Status::Corruption(), false};
            }
        } else if (schema_result.data != arbcore_schema_version) {
            // Database has schema version that does not match
            std::cerr << "Database version " << schema_result.data
                      << " does not match expected version "
                      << arbcore_schema_version
                      << ", delete database and try again" << std::endl;
            return {rocksdb::Status::Corruption(), false};
        }

        pruning_mode_result = pruningMode(tx);
    }

    auto pruning_status =
        initializePruningMode(database_exists, pruning_mode_result);
    if (!pruning_status.ok()) {
        // Error message already output in initializePruningMode
        return {pruning_status, false};
    }

    if (coreConfig.test_reset_db_except_inbox) {
        // For testing, delete everything from database except inbox
        {
            ReadWriteTransaction tx(data_storage);
            saveNextSegmentID(tx, 0);
            auto s = tx.commit();
            if (!s.ok()) {
                std::cerr << "Error resetting segment: " << s.ToString()
                          << std::endl;
                return {s, false};
            }
        }

        auto s = data_storage->clearDBExceptInbox();
        if (!s.ok()) {
            std::cerr << "Error deleting columns: " << s.ToString()
                      << std::endl;
            return {s, false};
        }
    }

    if (coreConfig.checkpoint_prune_on_startup) {
        // Delete all out of date checkpoints before proceeding.
        // This could take a while if pruning hasn't been done before.
        std::cout << "Pruning database"
                  << "\n";
        uint256_t last_timestamp = 0;
        {
            ReadWriteTransaction tx(data_storage);
            auto checkpoint_result = getLastCheckpoint(tx);
            if (std::holds_alternative<CheckpointVariant>(checkpoint_result)) {
                auto output = getMachineOutput(
                    std::get<CheckpointVariant>(checkpoint_result));
                last_timestamp = output.last_inbox_timestamp;
            }

            auto status = updatePruningMode(tx, "on");
            tx.commit();
        }
        if (last_timestamp > coreConfig.checkpoint_pruning_age_seconds) {
            auto checkpoint_pruning_age_timestamp =
                last_timestamp - coreConfig.checkpoint_pruning_age_seconds;
            // Delete in batches to prevent too much RAM from being used
            while (true) {
                auto status = pruneToTimestampOrBefore(
                    checkpoint_pruning_age_timestamp, 100);
                if (status.IsNotFound()) {
                    // Nothing left to delete
                    break;
                }
                if (!status.ok()) {
                    std::cerr
                        << "Error pruning checkpoints: " << status.ToString()
                        << std::endl;
                    break;
                }
            }
        }

        if (coreConfig.database_compact) {
            // Optimize database
            // Second compaction only done if pruning was done on startup
            std::cout << "Compacting database"
                      << "\n";
            auto status = data_storage->compact(true);
            if (!status.ok()) {
                std::cerr << "Database failed compacting: " << status.ToString()
                          << "\n";
            }
            std::cout << "Database finished compacting"
                      << "\n";
        }
    }

    auto reorg_done = false;
    rocksdb::Status status = rocksdb::Status::OK();
    if (coreConfig.test_reorg_to_l1_block != 0) {
        // Reset database for profile testing
        status =
            reorgToL1Block(coreConfig.test_reorg_to_l1_block, false, cache);
        reorg_done = true;
    } else if (coreConfig.test_reorg_to_l2_block != 0) {
        // Reset database for profile testing
        status =
            reorgToL2Block(coreConfig.test_reorg_to_l2_block, false, cache);
        reorg_done = true;
    } else if (coreConfig.test_reorg_to_log != 0) {
        // Reset database for profile testing
        status =
            reorgToLogCountOrBefore(coreConfig.test_reorg_to_log, false, cache);
        reorg_done = true;
    } else if (coreConfig.test_reorg_to_message != 0) {
        // Reset database for profile testing
        status = reorgToMessageCountOrBefore(coreConfig.test_reorg_to_message,
                                             false, cache);
        reorg_done = true;
    } else if (coreConfig.seed_cache_on_startup) {
        status = reorgToTimestampOrBefore(
            combined_machine_cache.currentTimeExpired(), true, cache);
        reorg_done = true;
    }

    if (coreConfig.database_save_on_startup) {
        std::filesystem::path save_rocksdb_path(coreConfig.database_save_path);
        std::filesystem::create_directories(save_rocksdb_path);
        ReadTransaction tx(data_storage);
        saveRocksdbCheckpoint(save_rocksdb_path, tx);
    }

    if (coreConfig.database_exit_after) {
        // Exit program, nothing else to do.
        if (!status.ok()) {
            std::cerr << "Exiting early, but error occurred: "
                      << status.ToString() << std::endl;
        }

        return {rocksdb::Status::OK(), true};
    }

    if (!reorg_done) {
        status = reorgToLastMessage(cache);
    }

    return {status, false};
}

InitializeResult ArbCore::initialize(const LoadedExecutable& executable) {
    // Use latest existing checkpoint

    auto config_result = applyConfig();
    if (config_result.finished) {
        return config_result;
    } else if (!config_result.status.IsNotFound()) {
        std::cerr << "Error with initial reorg: "
                  << config_result.status.ToString() << std::endl;
        return config_result;
    }

    // Need to initialize database from scratch
    core_code->addSegment(executable.code);
    core_machine = std::make_unique<MachineThread>(
        MachineState{core_code, executable.static_val});
    core_machine->machine_state.value_loader = makeValueLoader();
    core_machine->machine_state.code = std::make_shared<RunningCode>(core_code);

    last_machine = std::make_unique<Machine>(*core_machine);

    ReadWriteTransaction tx(data_storage);

    auto status = updateSchemaVersion(tx, arbcore_schema_version);
    if (!status.ok()) {
        std::cerr << "failed to save schema version into db: "
                  << status.ToString() << std::endl;
        return {status, false};
    }

    status = saveCheckpoint(tx);
    if (!status.ok()) {
        std::cerr << "failed to save initial checkpoint into db: "
                  << status.ToString() << std::endl;
        return {status, false};
    }

    status = updateLogInsertedCount(tx, 0);
    if (!status.ok()) {
        std::cerr << "failed to initialize log inserted count: "
                  << status.ToString() << std::endl;
        return {status, false};
    }
    status = updateSendInsertedCount(tx, 0);
    if (!status.ok()) {
        std::cerr << "failed to initialize send inserted count: "
                  << status.ToString() << std::endl;
        return {status, false};
    }

    for (size_t i = 0; i < logs_cursors.size(); i++) {
        status = logsCursorSaveCurrentTotalCount(tx, i, 0);
        if (!status.ok()) {
            std::cerr << "failed to initialize logscursor counts: "
                      << status.ToString() << std::endl;
            return {status, false};
        }
    }

    status = tx.commit();
    if (!status.ok()) {
        std::cerr << "failed to commit initial state into db: "
                  << status.ToString() << std::endl;
        return {status, false};
    }

    // Save initial state to cache
    combined_machine_cache.basicAdd(std::make_unique<Machine>(*core_machine));

    return {rocksdb::Status::OK(), false};
}

bool ArbCore::initialized() const {
    ReadTransaction tx(data_storage);
    std::vector<unsigned char> key;
    marshal_uint256_t(0, key);
    return tx.checkpointGetVector(vecToSlice(key)).status.ok();
}

rocksdb::Status ArbCore::initializePruningMode(
    bool database_exists,
    ValueResult<std::string>& pruning_mode_result) {
    if (database_exists) {
        if (!pruning_mode_result.status.ok()) {
            // Old database, does not have pruning mode set
            std::string mode;
            if (coreConfig.checkpoint_pruning_mode == PRUNING_MODE_ON) {
                // Command line explicitly enabled pruning
                mode = "on";
            } else {
                // Default old database to archive mode, no pruning
                mode = "off";
                coreConfig.checkpoint_pruning_mode = PRUNING_MODE_OFF;
            }
            ReadWriteTransaction tx(data_storage);
            auto status = updatePruningMode(tx, mode);
            if (!status.ok()) {
                std::cerr << "error saving pruning mode to existing database: "
                          << status.ToString() << std::endl;
                return status;
            }
            tx.commit();
            pruning_mode_result.data = mode;
            pruning_mode_result.status = rocksdb::Status::OK();
        } else {
            // Use setting from database unless command line overrides
            if (pruning_mode_result.data != "on" &&
                pruning_mode_result.data != "off") {
                std::cerr << "unexpected database pruning mode: "
                          << pruning_mode_result.data << std::endl;
                return rocksdb::Status::Corruption();
            } else if (coreConfig.checkpoint_pruning_mode == PRUNING_MODE_ON &&
                       pruning_mode_result.data == "off") {
                // Command line forced pruning, so mark database as being pruned
                std::cerr << "enabling pruning on previously unpruned database"
                          << "\n";
                ReadWriteTransaction tx(data_storage);
                auto status = updatePruningMode(tx, "on");
                if (!status.ok()) {
                    std::cerr << "Error updating pruning mode to off: "
                              << status.ToString() << std::endl;
                    return status;
                }
                tx.commit();
                pruning_mode_result.data = "on";
            } else if (coreConfig.checkpoint_pruning_mode == PRUNING_MODE_OFF &&
                       pruning_mode_result.data == "on") {
                // Database set to pruned but command line forcing no pruning
                std::cerr
                    << "warning: disabling pruning on already pruned database"
                    << "\n";
                pruning_mode_result.data = "off";
            } else if (coreConfig.checkpoint_pruning_mode != PRUNING_MODE_ON &&
                       coreConfig.checkpoint_pruning_mode != PRUNING_MODE_OFF) {
                // Default to database settings
                if (pruning_mode_result.data == "on") {
                    coreConfig.checkpoint_pruning_mode = PRUNING_MODE_ON;
                } else {
                    coreConfig.checkpoint_pruning_mode = PRUNING_MODE_OFF;
                }
            }
        }
    }

    if (coreConfig.checkpoint_pruning_mode == PRUNING_MODE_OFF) {
        // Disable pruning
        coreConfig.checkpoint_pruning_age_seconds = 0;
    }

    return rocksdb::Status::OK();
}

template <class T>
std::unique_ptr<T> ArbCore::getMachineImpl(ReadTransaction& tx,
                                           uint256_t machineHash,
                                           ValueCache& value_cache,
                                           bool lazy_load) {
    auto results = getMachineStateKeys(tx, machineHash);
    if (std::holds_alternative<rocksdb::Status>(results)) {
        throw std::runtime_error("failed to load machine state");
    }

    auto res = std::get<CountedData<CheckpointVariant>>(results).data;
    if (std::holds_alternative<MachineStateKeys>(res)) {
        return getMachineUsingStateKeys<T>(tx, std::get<MachineStateKeys>(res),
                                           value_cache, lazy_load);
    }

    // Machine not found
    return nullptr;
}

template std::unique_ptr<Machine> ArbCore::getMachineImpl(
    ReadTransaction& tx,
    uint256_t machineHash,
    ValueCache& value_cache,
    bool lazy_load);
template std::unique_ptr<MachineThread> ArbCore::getMachineImpl(
    ReadTransaction& tx,
    uint256_t machineHash,
    ValueCache& value_cache,
    bool lazy_load);

template <class T>
std::unique_ptr<T> ArbCore::getMachine(uint256_t machineHash,
                                       ValueCache& value_cache) {
    ReadSnapshotTransaction tx(data_storage);
    return getMachineImpl<T>(tx, machineHash, value_cache,
                             coreConfig.lazy_load_archive_queries);
}

template std::unique_ptr<Machine> ArbCore::getMachine(uint256_t, ValueCache&);
template std::unique_ptr<MachineThread> ArbCore::getMachine(uint256_t,
                                                            ValueCache&);

// triggerSaveFullRocksdbCheckpointToDisk is used to save a copy of current
// database without having to stop program
void ArbCore::triggerSaveFullRocksdbCheckpointToDisk() {
    trigger_save_rocksdb_checkpoint = true;
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

    saveCodeToCore(*core_machine, save_res.second);
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

void ArbCore::deleteCheckpoint(ReadWriteTransaction& tx,
                               const CheckpointVariant& checkpoint_variant) {
    if (std::holds_alternative<MachineStateKeys>(checkpoint_variant)) {
        // Obsolete checkpoint, need to delete referenced machine
        auto checkpoint = std::get<MachineStateKeys>(checkpoint_variant);
        deleteMachineState(tx, checkpoint);
    }

    auto machine_output = getMachineOutput(checkpoint_variant);

    // Delete checkpoint to make sure it isn't used later
    std::vector<unsigned char> key;
    marshal_uint256_t(machine_output.arb_gas_used, key);
    auto key_slice = vecToSlice(key);
    tx.checkpointDelete(key_slice);
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

rocksdb::Status ArbCore::reorgToLastMessage(ValueCache& cache) {
    std::cerr << "Reloading chain to the last message saved"
              << "\n";

    return reorgCheckpoints([&](const MachineOutput&) { return true; }, true,
                            cache);
}

rocksdb::Status ArbCore::reorgToL1Block(const uint256_t& l1_block_number,
                                        bool initial_start,
                                        ValueCache& cache) {
    if (initial_start) {
        std::cerr << "Reloading chain starting L1 block " << l1_block_number
                  << "\n";
    } else {
        std::cerr << "Reorg'ing chain to gas " << l1_block_number << "\n";
    }

    return reorgCheckpoints(
        [&](const MachineOutput& output) {
            return output.l1_block_number <= l1_block_number;
        },
        initial_start, cache);
}

rocksdb::Status ArbCore::reorgToL2Block(const uint256_t& l2_block_number,
                                        bool initial_start,
                                        ValueCache& cache) {
    if (initial_start) {
        std::cerr << "Reloading chain starting L2 block " << l2_block_number
                  << "\n";
    } else {
        std::cerr << "Reorg'ing chain to L2 block " << l2_block_number << "\n";
    }

    return reorgCheckpoints(
        [&](const MachineOutput& output) {
            return output.l2_block_number <= l2_block_number;
        },
        initial_start, cache);
}

rocksdb::Status ArbCore::reorgToLogCountOrBefore(const uint256_t& log_count,
                                                 bool initial_start,
                                                 ValueCache& cache) {
    if (initial_start) {
        std::cerr << "Reloading chain starting with log " << log_count << "\n";
    } else {
        std::cerr << "Reorg'ing chain to log " << log_count << "\n";
    }

    return reorgCheckpoints(
        [&](const MachineOutput& output) {
            return output.log_count <= log_count;
        },
        initial_start, cache);
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
            return output.fully_processed_inbox.count <= message_count;
        },
        initial_start, cache);
}

rocksdb::Status ArbCore::reorgToTimestampOrBefore(const uint256_t& timestamp,
                                                  bool initial_start,
                                                  ValueCache& cache) {
    if (initial_start) {
        std::cerr << "Reloading chain starting with timestamp "
                  << std::put_time(std::localtime((time_t*)&timestamp), "%c")
                  << "\n";
    } else {
        std::cerr << "Reorg'ing chain to timestamp "
                  << std::put_time(std::localtime((time_t*)&timestamp), "%c")
                  << "\n";
    }

    return reorgCheckpoints(
        [&](const MachineOutput& output) {
            return output.last_inbox_timestamp <= timestamp;
        },
        initial_start, cache);
}

// reorgDatabaseToMachineOutput resets database entries to match the state when
// the given output was the latest machine
rocksdb::Status ArbCore::reorgDatabaseToMachineOutput(
    const MachineOutput& output,
    ValueCache& value_cache) {
    auto log_inserted_count = logInsertedCount();
    if (!log_inserted_count.status.ok()) {
        std::cerr << "Error getting inserted count in Cursor Reorg: "
                  << log_inserted_count.status.ToString() << "\n";
        return log_inserted_count.status;
    }

    if (output.log_count < log_inserted_count.data) {
        // Update log cursors, must be called before logs are deleted
        for (size_t i = 0; i < logs_cursors.size(); i++) {
            auto status =
                handleLogsCursorReorg(i, output.log_count, value_cache);
            if (!status.ok()) {
                return status;
            }
        }
    }

    ReadWriteTransaction tx(data_storage);
    uint256_t next_sideload_block_number = 0;
    if (output.last_sideload.has_value()) {
        next_sideload_block_number = *output.last_sideload + 1;
    }

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

    return tx.commit();
}

// advanceCoreToTarget synchronously advances the core machine to the given
// target
rocksdb::Status ArbCore::advanceCoreToTarget(const MachineOutput& target_output,
                                             bool cache_sideloads) {
    while (core_machine->machine_state.output.arb_gas_used <
           target_output.arb_gas_used) {
        // Need to run machine until caught up with current checkpoint
        auto last_gas_used = core_machine->machine_state.output.arb_gas_used;
        MachineExecutionConfig execConfig;
        execConfig.stop_on_sideload = cache_sideloads;
        execConfig.stop_on_breakpoint = false;
        execConfig.max_gas = target_output.arb_gas_used;

        // Add messages and run machine
        uint256_t target_messages = target_output.fully_processed_inbox.count;
        uint256_t have_messages =
            core_machine->machine_state.output.fully_processed_inbox.count;
        uint256_t num_messages = target_messages - have_messages;
        if (uint256_t(num_messages) > coreConfig.message_process_count) {
            num_messages = coreConfig.message_process_count;
        }
        auto success =
            runMachineWithMessages(execConfig, size_t(num_messages), false);
        if (!success) {
            std::cerr << "runMachineWithMessages failed" << core_error_string
                      << "\n";
            return rocksdb::Status::Aborted();
        }

        if (core_machine->status() == MachineThread::MACHINE_ERROR) {
            core_error_string = core_machine->getErrorString();
            std::cerr << "AVM machine stopped with error during reorg: "
                      << core_error_string << "\n";
            return rocksdb::Status::Aborted();
        }

        while (core_machine->nextAssertion().sideload_block_number) {
            combined_machine_cache.timedAdd(
                std::make_unique<Machine>(*core_machine));

            if (core_machine->machine_state.output.arb_gas_used >=
                target_output.arb_gas_used) {
                break;
            }

            // Machine was stopped to save sideload,
            // start machine back up where it stopped
            auto machine_success = core_machine->continueRunningMachine(false);
            if (!machine_success) {
                core_error_string =
                    "Error continuing machine thread while reorging";
                machine_error = true;
                std::cerr << "Error catching up: " << core_error_string << "\n";
                return rocksdb::Status::Aborted();
            }
        }

        if (last_gas_used == core_machine->machine_state.output.arb_gas_used) {
            // No forward progress made
            break;
        }
    }

    // Ensure future continueRunningMachine calls don't use the set gas limit.
    core_machine->machine_state.context.max_gas = 0;

    // We skip over sideloads while fast forwarding, so l2_block_number and
    // last_sideload may not be updated properly.  Update here for completeness
    core_machine->machine_state.output.l2_block_number =
        target_output.l2_block_number;
    core_machine->machine_state.output.last_sideload =
        target_output.last_sideload;

    if (core_machine->machine_state.output != target_output) {
        // Machine in unexpected state, data corruption might have occurred
        std::cerr
            << "Error catching up: machine in unexpected state, current, "
               "expected gas: "
            << core_machine->machine_state.output.arb_gas_used << ", "
            << target_output.arb_gas_used << ", current, expected inbox count: "
            << core_machine->machine_state.output.fully_processed_inbox.count
            << ", " << target_output.fully_processed_inbox.count
            << ", current, expected log_count: "
            << core_machine->machine_state.output.log_count << ", "
            << target_output.log_count << ", current, expected log_acc: "
            << core_machine->machine_state.output.log_acc << ", "
            << target_output.log_acc << ", current, expected send_count: "
            << core_machine->machine_state.output.send_count << ", "
            << target_output.send_count << ", current, expected send_acc: "
            << core_machine->machine_state.output.send_acc << ", "
            << target_output.send_acc << "\n";
        return rocksdb::Status::Aborted();
    }

    return rocksdb::Status::OK();
}

// reorgToLastMatchingCheckpoint resets the checkpoint
// such that machine state is at the first matching checkpoint. cleaning
// up old references as needed.
std::variant<CheckpointVariant, rocksdb::Status>
ArbCore::reorgToLastMatchingCheckpoint(
    const std::function<bool(const MachineOutput&)>& check_output,
    ReadWriteTransaction& tx,
    std::unique_ptr<rocksdb::Iterator>& checkpoint_it) {
    // Obtain checkpoint iterator
    checkpoint_it->SeekToLast();
    if (!checkpoint_it->status().ok()) {
        std::cerr << "Error: SeekToLast failed during reorg: "
                  << checkpoint_it->status().ToString() << std::endl;
        return checkpoint_it->status();
    }

    if (!checkpoint_it->Valid()) {
        return rocksdb::Status::NotFound();
    }

    bool output_checkpoint = true;

    // Delete each checkpoint until check_output() is satisfied
    while (checkpoint_it->Valid()) {
        // Deserialize checkpoint
        std::vector<unsigned char> checkpoint_vector(
            checkpoint_it->value().data(),
            checkpoint_it->value().data() + checkpoint_it->value().size());
        auto checkpoint_variant = extractMachineStateKeys(checkpoint_vector);
        auto machine_output = getMachineOutput(checkpoint_variant);

        if (output_checkpoint) {
            printMachineOutputInfo("Last        database checkpoint",
                                   machine_output);
            output_checkpoint = false;
        }

        // Looking for first valid checkpoint
        if (check_output(machine_output)) {
            if (isValid(tx, machine_output.fully_processed_inbox)) {
                // Save first valid checkpoint
                printMachineOutputInfo("Last  valid database checkpoint",
                                       machine_output);
                return checkpoint_variant;
            } else {
                std::cerr << "Unexpectedly invalid checkpoint inbox at "
                             "message count "
                          << machine_output.fully_processed_inbox.count
                          << std::endl;
                assert(false);
            }
        }

        deleteCheckpoint(tx, checkpoint_variant);
        checkpoint_it->Prev();
    }

    auto status = checkpoint_it->status();
    if (!status.ok()) {
        return status;
    }
    return rocksdb::Status::NotFound();
}

// loadLastMatchingMachine finds the first checkpoint
// that includes a machine and matches check_output, cleaning
// up invalid checkpoints as needed.
std::variant<std::unique_ptr<MachineThread>, rocksdb::Status>
ArbCore::loadLastMatchingMachine(
    const CheckpointVariant& last_matching_database_checkpoint,
    const std::function<bool(const MachineOutput&)>& check_output,
    ReadWriteTransaction& tx,
    std::unique_ptr<rocksdb::Iterator>& checkpoint_it,
    ValueCache& value_cache) {
    // First check if last saved machine matches
    auto mach = combined_machine_cache.checkSimpleMatching(check_output);
    auto last_matching_database_output =
        getMachineOutput(last_matching_database_checkpoint);
    if (mach.machine != nullptr &&
        mach.machine->machine_state.output.arb_gas_used >=
            last_matching_database_output.arb_gas_used) {
        // Found machine in cache
        return std::make_unique<MachineThread>(mach.machine->machine_state);
    }

    // Check last database entry since we already extracted machine state
    auto last_machine_thread = getMachineThreadFromCheckpoint(
        last_matching_database_checkpoint, check_output, tx, value_cache);
    if (last_machine_thread) {
        return last_machine_thread;
    }
    // Skip last entry since it was just checked
    checkpoint_it->Prev();

    while (checkpoint_it->Valid()) {
        // Load next entry from database
        std::vector<unsigned char> checkpoint_vector(
            checkpoint_it->value().data(),
            checkpoint_it->value().data() + checkpoint_it->value().size());
        auto current_database_checkpoint =
            extractMachineStateKeys(checkpoint_vector);

        auto current_machine_thread = getMachineThreadFromCheckpoint(
            current_database_checkpoint, check_output, tx, value_cache);
        if (current_machine_thread) {
            return current_machine_thread;
        }
        checkpoint_it->Prev();
    }
    auto status = checkpoint_it->status();
    if (!status.ok()) {
        return status;
    }
    return rocksdb::Status::NotFound();
}

std::unique_ptr<MachineThread> ArbCore::getMachineThreadFromCheckpoint(
    const CheckpointVariant& current_database_checkpoint,
    const std::function<bool(const MachineOutput&)>& check_output,
    ReadWriteTransaction& tx,
    ValueCache& value_cache) {
    if (!std::holds_alternative<MachineStateKeys>(
            current_database_checkpoint)) {
        // Just holds MachineOutput, so continue until full checkpoint found
        return nullptr;
    }
    auto checkpoint = std::get<MachineStateKeys>(current_database_checkpoint);
    if (!check_output(checkpoint.output)) {
        // Something is wrong, checkpoint should always be valid here,
        // so delete checkpoint and keep looking
        printMachineOutputInfo(
            "Unexpected checkpoint encountered in "
            "loadLastMatchingMachine",
            checkpoint.output);
        deleteCheckpoint(tx, current_database_checkpoint);
        assert(false);
        return nullptr;
    }

    auto check_machine_state = [&](const MachineState& mach) {
        return check_output(mach.output) &&
               // If lazy loading is off, make sure to only get non-lazy loaded
               // machines for the core thread
               (!coreConfig.lazy_load_core_machine || !mach.lazy_loaded);
    };

    auto mach = combined_machine_cache.findFirstMatching(
        check_machine_state, std::nullopt, checkpoint.output.arb_gas_used,
        false);
    if (mach.machine != nullptr) {
        // Found machine in cache
        auto& state = mach.machine->machine_state;
        state.code->cleanupAfterReorg();
        auto nextSegment = std::max(state.code->initialSegmentForChildCode(),
                                    core_code->initialSegmentForChildCode());
        state.code = std::make_shared<RunningCode>(state.code, nextSegment);
        return std::make_unique<MachineThread>(state);
    }

    try {
        // Load machine from database
        return getMachineUsingStateKeys<MachineThread>(
            tx, checkpoint, value_cache, coreConfig.lazy_load_core_machine);
    } catch (const std::exception& e) {
        printMachineOutputInfo(
            "Exception loading machine for checkpoint in "
            "loadLastMatchingMachine",
            checkpoint.output);
        std::cerr << "Exception message: " << e.what() << std::endl;
        assert(false);
    }

    deleteCheckpoint(tx, current_database_checkpoint);
    return nullptr;
}

// reorgCheckpoints resets the checkpoint and database entries
// such that machine state is at or before the requested message. cleaning
// up old references as needed.
// If initial_start is true, caches are seeded but no reorg is done.
rocksdb::Status ArbCore::reorgCheckpoints(
    const std::function<bool(const MachineOutput&)>& check_output,
    bool initial_start,
    ValueCache& value_cache) {
    if (initial_start) {
        if (coreConfig.seed_cache_on_startup) {
            std::cerr << "Seeding cache" << std::endl;
        } else {
            std::cerr << "Loading initial machine" << std::endl;
        }
    } else {
        std::cerr << "Reorg blockchain" << std::endl;
    }

    auto output = getLastMachineOutput();
    printMachineOutputInfo("Previous checkpoint", output);

    // Remove any stale machine
    if (core_machine != nullptr) {
        core_machine->abortMachine();
    }

    using checkpoint_pair =
        std::pair<MachineOutput, std::unique_ptr<MachineThread>>;

    auto found_checkpoint_or_status =
        [&]() -> std::variant<checkpoint_pair, rocksdb::Status> {
        ReadWriteTransaction tx(data_storage);
        auto checkpoint_it = tx.checkpointGetIterator();
        auto reorg_machine_begin_timepoint = std::chrono::steady_clock::now();

        // Get last valid checkpoint from database
        // (may or may not contain machine)
        auto output_or_status =
            reorgToLastMatchingCheckpoint(check_output, tx, checkpoint_it);
        if (std::holds_alternative<rocksdb::Status>(output_or_status)) {
            return std::get<rocksdb::Status>(output_or_status);
        }
        auto last_matching_database_checkpoint =
            std::get<CheckpointVariant>(output_or_status);

        // Get last valid checkpoint with machine from database
        auto last_matching_machine_checkpoint = loadLastMatchingMachine(
            last_matching_database_checkpoint, check_output, tx, checkpoint_it,
            value_cache);
        if (std::holds_alternative<rocksdb::Status>(
                last_matching_machine_checkpoint)) {
            return std::get<rocksdb::Status>(last_matching_machine_checkpoint);
        }
        auto last_matching_machine =
            std::move(std::get<std::unique_ptr<MachineThread>>(
                last_matching_machine_checkpoint));

        auto& last_machine_output = last_matching_machine->machine_state.output;
        printMachineOutputInfo("Loaded             full machine",
                               last_machine_output);

        checkpoint_it = nullptr;

        auto status = tx.commit();
        if (!status.ok()) {
            std::cerr << "Error: unable to commit during reorg"
                      << status.ToString() << std::endl;
            return status;
        }
        // Calculate time including database commit
        printElapsed(reorg_machine_begin_timepoint, "Reorg took ");

        MachineOutput target_machine_output;
        auto last_database_checkpoint_output =
            getMachineOutput(last_matching_database_checkpoint);
        if (last_database_checkpoint_output.arb_gas_used >=
            last_machine_output.arb_gas_used) {
            // Advance machine to last partial database checkpoint
            target_machine_output = last_database_checkpoint_output;
        } else {
            // Checkpoint newer than last database checkpoint loaded from cache
            target_machine_output = last_machine_output;
        }
        return std::make_pair(getMachineOutput(target_machine_output),
                              std::move(last_matching_machine));
    }();

    if (std::holds_alternative<rocksdb::Status>(found_checkpoint_or_status)) {
        return std::get<rocksdb::Status>(found_checkpoint_or_status);
    }
    auto [selected_machine_output, setup] =
        std::move(std::get<checkpoint_pair>(found_checkpoint_or_status));

    core_machine = std::move(setup);

    // Remove invalid cache entries after selected_machine_output
    combined_machine_cache.reorg(selected_machine_output.arb_gas_used + 1);

    // Remove invalid database entries
    auto status =
        reorgDatabaseToMachineOutput(selected_machine_output, value_cache);
    if (!status.ok()) {
        return status;
    }

    // Advance core_machine to same place as selected_machine_output.
    if (initial_start && (core_machine->machine_state.output.l2_block_number !=
                          selected_machine_output.l2_block_number)) {
        std::cerr << "Seeding cache between L2 blocks: "
                  << core_machine->machine_state.output.l2_block_number << " - "
                  << selected_machine_output.l2_block_number << std::endl;
    }

    status = advanceCoreToTarget(selected_machine_output, initial_start);
    if (!status.ok()) {
        return status;
    }

    // Update last machine
    {
        std::unique_lock<std::shared_mutex> guard(last_machine_mutex);
        last_machine = std::make_unique<Machine>(*core_machine);
    }

    // Checkpoint was saved at sideload, attempt to continue running
    core_machine->machine_state.context.clearInboxMessages();
    core_machine->continueRunningMachine(true);

    return rocksdb::Status::OK();
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

uint64_t ArbCore::countCheckpoints(ReadTransaction& tx) {
    auto it = tx.checkpointGetIterator();
    it->SeekToFirst();
    uint64_t count = 0;
    // Skip 0th checkpoint
    it->Next();
    // Skip genesis checkpoint
    it->Next();
    while (it->Valid()) {
        count++;
        it->Next();
    }

    return count;
}

std::variant<rocksdb::Status, CheckpointVariant> ArbCore::getCheckpointNumber(
    ReadTransaction& tx,
    uint256_t& number) {
    auto it = tx.checkpointGetIterator();
    it->SeekToFirst();

    for (uint256_t i = 0; i < number && it->Valid(); i += 1, it->Next()) {
        // Loop until desired checkpoint
    }

    if (it->Valid()) {
        std::vector<unsigned char> saved_value(
            it->value().data(), it->value().data() + it->value().size());
        return extractMachineStateKeys(saved_value);
    }

    return it->status();
}

std::variant<rocksdb::Status, CheckpointVariant> ArbCore::getLastCheckpoint(
    ReadTransaction& tx) {
    auto it = tx.checkpointGetIterator();
    it->SeekToLast();
    if (it->Valid()) {
        std::vector<unsigned char> saved_value(
            it->value().data(), it->value().data() + it->value().size());
        return extractMachineStateKeys(saved_value);
    }

    return it->status();
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
    ValueCache& value_cache,
    bool lazy_load) const {
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
                       value_cache, lazy_load);
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
        !holds_alternative<Tuple>(
            std::get<CountedData<Value>>(stack_results).data)) {
        std::cerr << "failed to load machine stack" << std::endl;
        throw std::runtime_error("failed to load machine stack");
    }

    auto auxstack_results = ::getValueImpl(
        transaction, state_data.auxstack_hash, segment_ids, value_cache, false);
    if (std::holds_alternative<rocksdb::Status>(auxstack_results)) {
        std::cerr << "failed to load machine auxstack" << std::endl;
        throw std::runtime_error("failed to load machine auxstack");
    }
    if (!holds_alternative<Tuple>(
            std::get<CountedData<Value>>(auxstack_results).data)) {
        std::cerr << "failed to load machine auxstack because of format error"
                  << std::endl;
        throw std::runtime_error(
            "failed to load machine auxstack because of format error");
    }

    segment_ids.insert(state_data.pc.pc.segment);
    segment_ids.insert(state_data.err_pc.pc.segment);

    restoreCodeSegments(transaction, core_code, value_cache, segment_ids,
                        lazy_load);

    auto state = MachineState{
        state_data.output,
        state_data.pc.pc,
        std::make_shared<RunningCode>(core_code),
        makeValueLoader(),
        std::move(std::get<CountedData<Value>>(register_results).data),
        std::move(std::get<CountedData<Value>>(static_results).data),
        Datastack(get<Tuple>(std::get<CountedData<Value>>(stack_results).data)),
        Datastack(
            get<Tuple>(std::get<CountedData<Value>>(auxstack_results).data)),
        state_data.arb_gas_remaining,
        state_data.state,
        state_data.err_pc,
        lazy_load};

    return std::make_unique<T>(state);
}

template std::unique_ptr<Machine> ArbCore::getMachineUsingStateKeys(
    const ReadTransaction& transaction,
    const MachineStateKeys& state_data,
    ValueCache& value_cache,
    bool lazy_load) const;
template std::unique_ptr<MachineThread> ArbCore::getMachineUsingStateKeys(
    const ReadTransaction& transaction,
    const MachineStateKeys& state_data,
    ValueCache& value_cache,
    bool lazy_load) const;

#ifdef __linux__
static void* backtrace_buffer[1024];
void sigUsr2Handler(int signal) {
    if (signal != SIGUSR2)
        return;
    int addrs =
        backtrace(backtrace_buffer, sizeof(backtrace_buffer) / sizeof(void*));
    backtrace_symbols_fd(backtrace_buffer, addrs, 2);
}
#endif

void ArbCore::printCoreThreadBacktrace() {
#ifdef __linux__
    auto pthread = core_pthread.load();
    if (pthread) {
        pthread_kill(*pthread, SIGUSR2);
        return;
    }
#endif
    std::cerr << "Core thread backtrace not available" << std::endl;
}

// operator() runs the main thread for ArbCore.  It is responsible for adding
// messages to the queue, starting machine thread when needed and collecting
// results of machine thread.
// This thread will update `delivering_messages` if and only if
// `delivering_messages` is set to MESSAGES_READY
void ArbCore::operator()() {
#ifdef __linux__
    prctl(PR_SET_NAME, "ArbCore", 0, 0, 0);
    signal(SIGUSR2, sigUsr2Handler);
    core_pthread = pthread_self();
#endif
    ValueCache cache{5, 0};
    MachineExecutionConfig execConfig;
    execConfig.stop_on_sideload = true;
    execConfig.stop_on_breakpoint = false;
    std::chrono::time_point<std::chrono::steady_clock>
        next_rocksdb_save_timepoint{};
    std::filesystem::path save_rocksdb_path(coreConfig.database_save_path);
    auto profiling_begin_timepoint = std::chrono::steady_clock::now();
    auto begin_message =
        core_machine->machine_state.output.fully_processed_inbox.count;
    auto last_messages_ready_check_timepoint = profiling_begin_timepoint;
    auto last_run_machine_check_timepoint = profiling_begin_timepoint;
    auto last_restart_machine_check_timepoint = profiling_begin_timepoint;
    auto perform_pruning = false;
    auto perform_save_rocksdb_checkpoint = false;

    if (coreConfig.database_save_interval > 0) {
        next_rocksdb_save_timepoint =
            std::chrono::steady_clock::now() +
            std::chrono::seconds(coreConfig.database_save_interval);
        std::filesystem::create_directories(save_rocksdb_path);
    }

    uint256_t next_checkpoint_gas =
        maxCheckpointGas() + coreConfig.checkpoint_gas_frequency;
    uint256_t next_basic_cache_gas =
        maxCheckpointGas() + coreConfig.basic_machine_cache_interval;
    try {
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
                next_checkpoint_gas = coreConfig.checkpoint_gas_frequency;
            }
            if (message_data_status == MESSAGES_READY) {
                std::chrono::time_point<std::chrono::steady_clock>
                    begin_messages_ready_timepoint;
                if (coreConfig.debug_timing) {
                    begin_messages_ready_timepoint =
                        std::chrono::steady_clock::now();
                    printElapsed(last_messages_ready_check_timepoint,
                                 "ArbCore addMessages check delay: ");
                }
                // Reorg might occur while adding messages
                try {
                    auto add_status = addMessages(message_data, cache);
                    if (!add_status.status.ok()) {
                        core_error_string = add_status.status.ToString();
                        message_data_status = MESSAGES_ERROR;
                        std::cerr << "ArbCore addMessages error: "
                                  << core_error_string << "\n";
                    } else {
                        machine_idle = false;
                        message_data_status = MESSAGES_SUCCESS;
                        if (add_status.data.has_value()) {
                            next_checkpoint_gas =
                                add_status.data.value() +
                                coreConfig.checkpoint_gas_frequency;
                        }
                    }
                } catch (DataStorage::shutting_down_exception& e) {
                    throw;
                } catch (const std::exception& e) {
                    core_error_string = e.what();
                    message_data_status = MESSAGES_ERROR;
                    std::cerr << "ArbCore addMessages exception: "
                              << core_error_string << "\n";
                }
                if (coreConfig.debug_timing) {
                    printElapsed(begin_messages_ready_timepoint,
                                 "ArbCore addMessages duration: ");
                }
            }
            if (coreConfig.debug_timing) {
                last_messages_ready_check_timepoint =
                    std::chrono::steady_clock::now();
            }

            // Check machine thread
            if (core_machine->status() == MachineThread::MACHINE_ERROR) {
                core_error_string = core_machine->getErrorString();
                std::cerr << "AVM machine stopped with error: "
                          << core_error_string << "\n";
                break;
            }

            if (core_machine->status() == MachineThread::MACHINE_SUCCESS) {
                std::chrono::time_point<std::chrono::steady_clock>
                    begin_machine_success_timepoint;
                std::chrono::time_point<std::chrono::steady_clock>
                    output_timepoint;
                std::chrono::time_point<std::chrono::steady_clock>
                    cache_timepoint;
                std::chrono::time_point<std::chrono::steady_clock>
                    logs_timepoint;
                if (coreConfig.debug_timing) {
                    begin_machine_success_timepoint =
                        std::chrono::steady_clock::now();
                }
                ReadWriteTransaction tx(data_storage);

                auto last_assertion = core_machine->nextAssertion();

                // Save last machine output
                {
                    std::unique_lock<std::shared_mutex> guard(
                        last_machine_mutex);
                    last_machine = std::make_unique<Machine>(*core_machine);
                }
                if (coreConfig.debug_timing) {
                    printElapsed(begin_machine_success_timepoint,
                                 "ArbCore machine output save time: ");
                }

                if (core_machine->machine_state.output.arb_gas_used >
                    next_basic_cache_gas) {
                    combined_machine_cache.basicAdd(
                        std::make_unique<Machine>(*core_machine));

                    next_basic_cache_gas =
                        core_machine->machine_state.output.arb_gas_used +
                        coreConfig.basic_machine_cache_interval;
                } else if (coreConfig.last_machine_cache) {
                    combined_machine_cache.lastAdd(
                        std::make_unique<Machine>(*core_machine));
                }
                if (coreConfig.debug_timing) {
                    printElapsed(output_timepoint, "ArbCore cache save time: ");
                }

                // Save logs and sends
                auto status = saveAssertion(
                    tx, last_assertion,
                    core_machine->machine_state.output.arb_gas_used);
                if (!status.ok()) {
                    core_error_string = status.ToString();
                    std::cerr << "ArbCore assertion saving failed: "
                              << core_error_string << "\n";
                    break;
                }
                if (coreConfig.debug_timing) {
                    printElapsed(cache_timepoint,
                                 "ArbCore logs and sends save time: ");
                }

                // Cache pre-sideload machines
                auto save_checkpoint_begin_timepoint =
                    std::chrono::steady_clock::now();
                auto checkpoint_was_saved = false;
                if (last_assertion.sideload_block_number) {
                    auto& output = core_machine->machine_state.output;

                    combined_machine_cache.timedAdd(
                        std::make_unique<Machine>(*core_machine));

                    auto current_timepoint = std::chrono::steady_clock::now();
                    if (coreConfig.database_save_interval != 0 &&
                        current_timepoint >= next_rocksdb_save_timepoint) {
                        trigger_save_rocksdb_checkpoint = true;

                        next_rocksdb_save_timepoint =
                            current_timepoint +
                            std::chrono::seconds(
                                coreConfig.database_save_interval);
                    }

                    if (trigger_save_rocksdb_checkpoint ||
                        output.arb_gas_used >= next_checkpoint_gas) {
                        // Save checkpoint after checkpoint_gas_frequency gas
                        // used
                        status = saveCheckpoint(tx);
                        if (!status.ok()) {
                            core_error_string = status.ToString();
                            std::cerr << "ArbCore checkpoint saving failed: "
                                      << core_error_string << "\n";
                            break;
                        }
                        printMachineOutputInfo("Saved checkpoint ", output);
                        checkpoint_was_saved = true;

                        // Clear oldest cache and start populating next cache
                        cache.nextCache();
                        next_checkpoint_gas =
                            output.arb_gas_used +
                            coreConfig.checkpoint_gas_frequency;

                        if (trigger_save_rocksdb_checkpoint) {
                            // database is ready to be copied
                            trigger_save_rocksdb_checkpoint = false;
                            perform_save_rocksdb_checkpoint = true;
                        }

                        // Perform pruning soon
                        if (coreConfig.checkpoint_pruning_mode !=
                            PRUNING_MODE_OFF) {
                            perform_pruning = true;
                        }
                    }

                    if (coreConfig.debug_timing) {
                        printElapsed(last_restart_machine_check_timepoint,
                                     "ArbCore machine restart delay: ");
                    }
                    // Machine was stopped to save sideload, update execConfig
                    // and start machine back up where it stopped
                    auto machine_success =
                        core_machine->continueRunningMachine(true);
                    if (!machine_success) {
                        core_error_string = "Error starting machine thread";
                        machine_error = true;
                        std::cerr << "ArbCore error: " << core_error_string
                                  << "\n";
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
                // Calculate time including database commit
                if (checkpoint_was_saved) {
                    printElapsed(save_checkpoint_begin_timepoint,
                                 "Saving checkpoint took ");
                }

                // Check if checkpoint of full database needs to be saved to
                // disk
                if (perform_save_rocksdb_checkpoint) {
                    perform_save_rocksdb_checkpoint = false;

                    saveRocksdbCheckpoint(save_rocksdb_path, tx);
                }

                auto output = getLastMachineOutput();
                if (coreConfig.test_run_until != 0 &&
                    output.fully_processed_inbox.count >=
                        coreConfig.test_run_until) {
                    // Reached stopping point for profiling
                    std::cout << "Done processing " << begin_message << " to "
                              << output.fully_processed_inbox.count << "\n";
                    printElapsed(profiling_begin_timepoint, "Profiling took ");

                    if (coreConfig.test_load_count > 0) {
                        auto load_begin_timepoint =
                            std::chrono::steady_clock::now();
                        auto target_gas = output.arb_gas_used;
                        for (uint64_t i = 0; i < coreConfig.test_load_count;
                             i++) {
                            std::cerr << "Loading machine " << i << std::endl;
                            auto current_execution = findCloserExecutionCursor(
                                tx, std::nullopt, target_gas, true);
                            if (std::holds_alternative<rocksdb::Status>(
                                    current_execution)) {
                                std::cerr
                                    << "Error loading profile machine number "
                                    << i << ": "
                                    << std::get<rocksdb::Status>(
                                           current_execution)
                                           .ToString()
                                    << std::endl;
                                break;
                            }
                        }

                        std::cerr << "Done loading "
                                  << coreConfig.test_load_count << " machines"
                                  << std::endl;
                        printElapsed(load_begin_timepoint,
                                     "Loading machines took ");
                    }

                    // Exit now that profiling is complete
                    break;
                }
                if (coreConfig.debug_timing) {
                    printElapsed(begin_machine_success_timepoint,
                                 "ArbCore machine success total save time: ");
                }
            }
            if (coreConfig.debug_timing) {
                last_restart_machine_check_timepoint =
                    std::chrono::steady_clock::now();
            }

            if (perform_pruning) {
                perform_pruning = false;

                uint256_t checkpoint_pruning_gas_used = 0;
                {
                    std::lock_guard<std::mutex> lock(checkpoint_pruning_mutex);
                    checkpoint_pruning_gas_used =
                        unsafe_checkpoint_pruning_gas_used;
                }

                if (checkpoint_pruning_gas_used > 0) {
                    // Prune checkpoints that have used less gas
                    // than specified
                    auto prune_status =
                        pruneToGasOrBefore(checkpoint_pruning_gas_used,
                                           coreConfig.checkpoint_max_to_prune);
                    if (!prune_status.ok() && !prune_status.IsNotFound()) {
                        // Non-fatal error
                        std::cerr << "Error pruning checkpoints: "
                                  << prune_status.ToString() << "\n";
                    }
                } else {
                    auto output = getLastMachineOutput();
                    auto last_inbox_timestamp = output.last_inbox_timestamp;

                    if (coreConfig.checkpoint_pruning_age_seconds > 0 &&
                        last_inbox_timestamp >
                            coreConfig.checkpoint_pruning_age_seconds) {
                        // Prune checkpoints that are too old
                        auto checkpoint_pruning_age_timestamp =
                            last_inbox_timestamp -
                            coreConfig.checkpoint_pruning_age_seconds;
                        auto prune_status = pruneToTimestampOrBefore(
                            checkpoint_pruning_age_timestamp,
                            coreConfig.checkpoint_max_to_prune);
                        if (!prune_status.ok() && !prune_status.IsNotFound()) {
                            // Non-fatal error
                            std::cerr << "Error pruning checkpoints: "
                                      << prune_status.ToString() << "\n";
                        }
                    }
                }
            }

            if (core_machine->status() == MachineThread::MACHINE_ABORTED) {
                // Just reset status so machine can be restarted
                core_machine->clearError();
            }

            if (core_machine->status() == MachineThread::MACHINE_NONE) {
                // Start execution of machine if new message available
                auto success = runMachineWithMessages(
                    execConfig, coreConfig.message_process_count, true);
                if (!success) {
                    break;
                }
                if (coreConfig.debug_timing && machine_idle == false) {
                    printElapsed(last_run_machine_check_timepoint,
                                 "ArbCore runMachineWithMessages delay: ");
                }
            }
            if (coreConfig.debug_timing) {
                last_run_machine_check_timepoint =
                    std::chrono::steady_clock::now();
            }

            for (size_t i = 0; i < logs_cursors.size(); i++) {
                ValueCache logs_cache{1, 0};
                if (logs_cursors[i].status == DataCursor::REQUESTED) {
                    ReadTransaction tx(data_storage);
                    handleLogsCursorRequested(tx, i, logs_cache);
                }
            }

            if (machine_idle && message_data_status != MESSAGES_READY) {
                // Machine blocked and no new messages, so sleep for a bit
                std::this_thread::sleep_for(std::chrono::milliseconds(
                    coreConfig.idle_sleep_milliseconds));
            }
        }
    } catch (DataStorage::shutting_down_exception& e) {
        std::cerr << "Caught shutting down exception in ArbCore thread"
                  << std::endl;
    } catch (const std::exception& e) {
        std::cerr << "ArbCore thread exception: " << e.what() << std::endl;
    }
    std::cerr << "Exiting main ArbCore thread" << std::endl;

    // Error occurred, make sure machine stops cleanly
    core_machine->abortMachine();
    for (auto& logs_cursor : logs_cursors) {
        logs_cursor.error_string = "arbcore thread aborted";
        logs_cursor.status = DataCursor::ERROR;
    }

#ifdef __linux__
    core_pthread = std::nullopt;
#endif
}
void ArbCore::saveRocksdbCheckpoint(
    const std::filesystem::path& save_rocksdb_path,
    ReadTransaction& tx) {
    struct stat info;
    if ((stat(save_rocksdb_path.c_str(), &info) != 0) &&
        (info.st_mode & S_IFDIR) == 0) {
        std::cerr << "Unable to save checkpoint into " << save_rocksdb_path
                  << " because directory doesn't exist" << std::endl;
    } else {
        auto save_rocksdb_checkpoint_begin_timepoint =
            std::chrono::steady_clock::now();
        auto timestamp_dir = std::to_string(seconds_since_epoch());
        auto checkpoint_dir = save_rocksdb_path / timestamp_dir;
        auto status = tx.createRocksdbCheckpoint(checkpoint_dir.string());
        if (!status.ok()) {
            std::cerr << "Unable to save checkpoint into " << checkpoint_dir
                      << ", error: " << status.ToString() << std::endl;
        } else {
            std::cerr << "Saving rocksdb checkpoint in " << checkpoint_dir
                      << std::endl;
            printElapsed(save_rocksdb_checkpoint_begin_timepoint,
                         "Saving rocksdb checkpoint took ");
        }
    }
}

void ArbCore::printElapsed(
    const std::chrono::time_point<std::chrono::steady_clock>& begin_timepoint,
    const std::string& message) const {
    auto machine_output_milliseconds =
        std::chrono::duration_cast<std::chrono::milliseconds>(
            std::chrono::steady_clock::now() - begin_timepoint)
            .count();
    if (machine_output_milliseconds > 0) {
        std::cout << message << machine_output_milliseconds << "ms"
                  << "\n";
    }
}

bool ArbCore::runMachineWithMessages(MachineExecutionConfig& execConfig,
                                     size_t max_message_batch_size,
                                     bool asynchronous) {
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

    if (std::holds_alternative<NotBlocked>(
            core_machine->isBlocked(!messages_result.data.empty()))) {
        execConfig.inbox_messages = messages_result.data;

        auto success = core_machine->runMachine(execConfig, asynchronous);
        if (!success) {
            core_error_string = "Error starting machine thread";
            machine_error = true;
            std::cerr << "ArbCore error: " << core_error_string << "\n";
            return false;
        }
    } else {
        // Machine all caught up, no messages to process
        machine_idle = true;
    }

    return true;
}

rocksdb::Status ArbCore::saveLogs(
    ReadWriteTransaction& tx,
    const std::vector<MachineEmission<Value>>& logs) {
    if (logs.empty()) {
        return rocksdb::Status::OK();
    }
    auto log_result = logInsertedCountImpl(tx);
    if (!log_result.status.ok()) {
        return log_result.status;
    }

    auto log_index = log_result.data;
    for (const auto& log : logs) {
        auto value_result = saveValue(tx, log.val);
        if (!value_result.status.ok()) {
            return value_result.status;
        }

        std::vector<unsigned char> key;
        marshal_uint256_t(log_index, key);
        auto key_slice = vecToSlice(key);

        std::vector<unsigned char> log_info;
        marshal_uint256_t(hash_value(log.val), log_info);
        marshal_uint256_t(log.inbox.count, log_info);
        marshal_uint256_t(log.inbox.accumulator, log_info);
        rocksdb::Slice log_info_slice(
            reinterpret_cast<const char*>(log_info.data()), log_info.size());

        auto status = tx.logPut(key_slice, log_info_slice);
        if (!status.ok()) {
            return status;
        }
        log_index += 1;
    }

    return updateLogInsertedCount(tx, log_index);
}

ValueResult<std::vector<MachineEmission<Value>>>
ArbCore::getLogs(uint256_t index, uint256_t count, ValueCache& valueCache) {
    ReadSnapshotTransaction tx(data_storage);

    return getLogsNoLock(tx, index, count, valueCache);
}

ValueResult<std::vector<MachineEmission<Value>>> ArbCore::getLogsNoLock(
    ReadTransaction& tx,
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

    std::vector<unsigned char> lower_key;
    marshal_uint256_t(index, lower_key);
    auto lower_slice = vecToSlice(lower_key);

    std::vector<unsigned char> upper_key;
    marshal_uint256_t(index + count, upper_key);
    auto upper_slice = vecToSlice(upper_key);

    auto it = tx.logGetIterator(&lower_slice, &upper_slice);
    it->SeekToFirst();
    std::vector<MachineEmission<Value>> logs;
    while (it->Valid()) {
        auto info = it->value().data();
        auto hash = extractUint256(info);
        auto val_result = getValue(tx, hash, valueCache, false);
        if (std::holds_alternative<rocksdb::Status>(val_result)) {
            return {std::get<rocksdb::Status>(val_result), {}};
        }
        auto inbox_count = extractUint256(info);
        auto inbox_accumulator = extractUint256(info);
        auto inbox = InboxState{inbox_count, inbox_accumulator};
        auto val = std::move(std::get<CountedData<Value>>(val_result).data);
        logs.push_back(MachineEmission<Value>{val, inbox});
        it->Next();
    }

    return {it->status(), std::move(logs)};
}

rocksdb::Status ArbCore::saveSends(
    ReadWriteTransaction& tx,
    const std::vector<MachineEmission<std::vector<unsigned char>>>& sends) {
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
        std::vector<unsigned char> val_data;
        marshal_uint256_t(send.inbox.count, val_data);
        marshal_uint256_t(send.inbox.accumulator, val_data);
        val_data.insert(val_data.end(), send.val.begin(), send.val.end());

        auto status = tx.sendPut(key_slice, vecToSlice(val_data));
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

    while (it->Valid()) {
        auto key_ptr = reinterpret_cast<const unsigned char*>(it->key().data());
        auto value_ptr =
            reinterpret_cast<const unsigned char*>(it->value().data());
        auto value_end_ptr = value_ptr + it->value().size();

        auto seq_batch_item = deserializeSequencerBatchItem(
            extractUint256(key_ptr), value_ptr, value_end_ptr);
        if (seq_batch_item.sequencer_message) {
            auto message_it = seq_batch_item.sequencer_message->begin();
            auto block_num = extractInboxMessageBlockNumber(message_it);
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
    auto current_message_data = message_res.data[0].message;
    auto message = extractInboxMessage(current_message_data);
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
                std::cerr << "Found reorg in getMessagesImpl, index: "
                          << intx::to_string(index, 10) << ", expected: "
                          << optionalUint256ToString(start_acc) << ", found: "
                          << intx::to_string(item.accumulator, 16) << std::endl;
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
        std::cerr << "Found reorg in getMessagesImpl, index: "
                  << intx::to_string(index, 10)
                  << ", expected: " << optionalUint256ToString(start_acc)
                  << ", unable to load message for consistency check"
                  << std::endl;
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

    std::vector<unsigned char> lower_key;
    marshal_uint256_t(index, lower_key);
    auto lower_slice = vecToSlice(lower_key);

    std::vector<unsigned char> upper_key;
    marshal_uint256_t(index + count, upper_key);
    auto upper_slice = vecToSlice(upper_key);

    auto it = tx.sendGetIterator(&lower_slice, &upper_slice);
    it->SeekToFirst();
    std::vector<std::vector<unsigned char>> send_data;
    while (it->Valid()) {
        auto value = it->value();
        // Skip inbox metadata
        auto data_start = value.data() + 64;
        auto data_end = value.data() + value.size();
        std::vector<unsigned char> data(data_start, data_end);
        send_data.push_back(data);
        it->Next();
    }

    return {it->status(), send_data};
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
    if (last_machine) {
        return last_machine->machine_state.output;
    }

    return MachineOutput{};
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
    uint256_t total_gas_used,
    bool allow_slow_lookup) {
    std::unique_ptr<ExecutionCursor> execution_cursor;
    {
        ReadSnapshotTransaction tx(data_storage);

        auto closest_checkpoint = findCloserExecutionCursor(
            tx, std::nullopt, total_gas_used, allow_slow_lookup);
        if (std::holds_alternative<rocksdb::Status>(closest_checkpoint)) {
            std::cerr << "No execution machine available close to gas: "
                      << total_gas_used << std::endl;
            return {std::get<rocksdb::Status>(closest_checkpoint), nullptr};
        }

        execution_cursor = std::make_unique<ExecutionCursor>(
            std::get<ExecutionCursor>(closest_checkpoint));
    }

    auto status = advanceExecutionCursorImpl(
        *execution_cursor, total_gas_used, false,
        coreConfig.message_process_count, allow_slow_lookup);

    if (!status.ok()) {
        std::cerr << "Couldn't advance execution machine" << std::endl;
    }

    return {status, std::move(execution_cursor)};
}

rocksdb::Status ArbCore::advanceExecutionCursor(
    ExecutionCursor& execution_cursor,
    uint256_t max_gas,
    bool go_over_gas,
    bool allow_slow_lookup) {
    auto current_gas = execution_cursor.getOutput().arb_gas_used;
    auto total_gas_used = current_gas + max_gas;
    {
        ReadSnapshotTransaction tx(data_storage);
        auto result = findCloserExecutionCursor(
            tx, std::move(execution_cursor), total_gas_used, allow_slow_lookup);
        if (std::holds_alternative<rocksdb::Status>(result)) {
            return std::get<rocksdb::Status>(result);
        }

        execution_cursor = std::get<ExecutionCursor>(result);
    }

    return advanceExecutionCursorImpl(
        execution_cursor, total_gas_used, go_over_gas,
        coreConfig.message_process_count, allow_slow_lookup);
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
        execution_cursor.machine = getMachineUsingStateKeys<Machine>(
            tx, machine_state_keys, cache,
            coreConfig.lazy_load_archive_queries);
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
    size_t message_group_size,
    bool allow_slow_lookup) {
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

        while (!execution_cursor.isAborted()) {
            // Run machine until specified gas is reached
            MachineExecutionConfig execConfig;
            execConfig.max_gas = total_gas_used;
            execConfig.go_over_gas = go_over_gas;

            {
                ReadSnapshotTransaction tx(data_storage);

                auto& mach =
                    resolveExecutionCursorMachine(tx, execution_cursor);

                if (execution_cursor.isAborted()) {
                    // Extra check in case machine was not resolved before abort
                    return rocksdb::Status::Aborted();
                }

                uint256_t gas_used = execution_cursor.getOutput().arb_gas_used;
                if ((gas_used >= total_gas_used) ||
                    (!go_over_gas &&
                     gas_used + mach->machine_state.nextGasCost() >
                         total_gas_used)) {
                    // Gas target reached
                    break;
                }

                if ((coreConfig.checkpoint_max_execution_gas != 0) &&
                    (total_gas_used - gas_used >
                     coreConfig.checkpoint_max_execution_gas)) {
                    // Execution will take too long
                    return rocksdb::Status::NotFound();
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
                    std::cerr << "Error getting messages for execution cursor"
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
                // Didn't make any progress
                break;
            }
        }

        if (execution_cursor.isAborted()) {
            return rocksdb::Status::Aborted();
        }

        if (handle_reorg) {
            ReadSnapshotTransaction tx(data_storage);

            auto closest_checkpoint = findCloserExecutionCursor(
                tx, std::nullopt, total_gas_used, allow_slow_lookup);
            if (std::holds_alternative<rocksdb::Status>(closest_checkpoint)) {
                std::cerr << "No execution machine available close to gas: "
                          << total_gas_used << std::endl;
                return std::get<rocksdb::Status>(closest_checkpoint);
            }
            execution_cursor =
                std::move(std::get<ExecutionCursor>(closest_checkpoint));
        }
    }

    if (std::holds_alternative<std::unique_ptr<Machine>>(
            execution_cursor.machine)) {
        auto& mach =
            std::get<std::unique_ptr<Machine>>(execution_cursor.machine);
        combined_machine_cache.lruAdd(std::make_unique<Machine>(*mach));
    }

    return rocksdb::Status::OK();
}

std::variant<rocksdb::Status, ExecutionCursor>
ArbCore::findCloserExecutionCursor(
    ReadTransaction& tx,
    std::optional<ExecutionCursor> execution_cursor,
    uint256_t& total_gas_used,
    bool allow_slow_lookup) {
    std::optional<uint256_t> existing_gas_used;
    if (execution_cursor.has_value()) {
        existing_gas_used = execution_cursor.value().getOutput().arb_gas_used;

        if (existing_gas_used.value() == total_gas_used) {
            // Nothing needs to be done
            return execution_cursor.value();
        }
    }

    std::optional<MachineStateKeys> database_machine_state_keys;
    std::optional<uint256_t> database_gas;
    if (allow_slow_lookup) {
        auto checkpoint_result = getCheckpointUsingGas(tx, total_gas_used);
        if (std::holds_alternative<rocksdb::Status>(checkpoint_result)) {
            return std::get<rocksdb::Status>(checkpoint_result);
        }

        database_machine_state_keys =
            std::get<MachineStateKeys>(checkpoint_result);

        // Guaranteed to always have a value at this point
        database_gas = database_machine_state_keys.value().output.arb_gas_used;
    }

    auto mach = combined_machine_cache.atOrBeforeGas(
        total_gas_used, existing_gas_used, database_gas, true);

    switch (mach.status) {
        case CombinedMachineCache::Success:
            mach.machine->machine_state.code =
                std::make_shared<RunningCode>(mach.machine->machine_state.code);
            return ExecutionCursor(std::move(mach.machine));
        case CombinedMachineCache::UseDatabase:
            return ExecutionCursor(database_machine_state_keys.value());
        case CombinedMachineCache::TooMuchExecution:
            return rocksdb::Status::NotFound();
        case CombinedMachineCache::UseExisting:
        case CombinedMachineCache::NotFound:
            break;
    }

    if (!execution_cursor.has_value()) {
        return rocksdb::Status::NotFound();
    }

    return execution_cursor.value();
}

std::variant<rocksdb::Status, ExecutionCursor>
ArbCore::getExecutionCursorAtBlock(const uint256_t& block_number,
                                   bool allow_slow_lookup) {
    uint256_t gas_target;
    std::unique_ptr<ExecutionCursor> execution_cursor;
    {
        ReadSnapshotTransaction tx(data_storage);
        auto gas_used_result = getSideloadPosition(tx, block_number);
        if (!gas_used_result.status.ok()) {
            return gas_used_result.status;
        }
        gas_target = gas_used_result.data;

        auto closest_checkpoint = findCloserExecutionCursor(
            tx, std::nullopt, gas_target, allow_slow_lookup);
        if (std::holds_alternative<rocksdb::Status>(closest_checkpoint)) {
            return std::get<rocksdb::Status>(closest_checkpoint);
        }

        execution_cursor = std::make_unique<ExecutionCursor>(
            std::get<ExecutionCursor>(closest_checkpoint));
    }

    auto status = advanceExecutionCursorImpl(
        *execution_cursor, gas_target, false, coreConfig.message_process_count,
        allow_slow_lookup);

    ReadSnapshotTransaction tx(data_storage);
    return *execution_cursor;
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
        messages.emplace_back(std::move(message), raw_message.accumulator);
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

ValueResult<uint256_t> ArbCore::schemaVersion(ReadTransaction& tx) const {
    return tx.stateGetUint256(vecToSlice(schema_version_key));
}
rocksdb::Status ArbCore::updateSchemaVersion(ReadWriteTransaction& tx,
                                             const uint256_t& schema_version) {
    std::vector<unsigned char> value;
    marshal_uint256_t(schema_version, value);

    return tx.statePut(vecToSlice(schema_version_key), vecToSlice(value));
}

ValueResult<std::string> ArbCore::pruningMode(ReadTransaction& tx) const {
    std::string mode;
    auto status = tx.stateGet(vecToSlice(pruning_mode_key), &mode);
    return {status, mode};
}
rocksdb::Status ArbCore::updatePruningMode(ReadWriteTransaction& tx,
                                           const std::string& pruning_mode) {
    return tx.statePut(vecToSlice(pruning_mode_key), pruning_mode);
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
    if (reorging_to_count.has_value()) {
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
            auto logs_to_keep =
                intx::narrow_cast<long>(log_count - current_count_result.data);
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
        throw std::runtime_error("Unexpected logsCursor status");
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

ValueResult<std::unique_ptr<Machine>> ArbCore::getMachineAtBlock(
    const uint256_t& block_number,
    bool allow_slow_lookup) {
    auto cursor = getExecutionCursorAtBlock(block_number, allow_slow_lookup);
    if (std::holds_alternative<rocksdb::Status>(cursor)) {
        return {std::get<rocksdb::Status>(cursor), nullptr};
    }

    ReadSnapshotTransaction tx(data_storage);
    return {rocksdb::Status::OK(), takeExecutionCursorMachineImpl(
                                       tx, std::get<ExecutionCursor>(cursor))};
}

uint64_t seconds_since_epoch() {
    return std::chrono::duration_cast<std::chrono::seconds>(
               std::chrono::system_clock::now().time_since_epoch())
        .count();
}

rocksdb::Status ArbCore::pruneCheckpoints(
    const std::function<bool(const MachineOutput&)>& check_output,
    uint64_t checkpoint_max_to_prune) {
    ReadWriteTransaction tx(data_storage);
    auto it = tx.checkpointGetIterator();
    if (!it->status().ok()) {
        std::cerr << "unable to prune old checkpoints, iterator error: "
                  << it->status().ToString() << std::endl;
        return it->status();
    }

    auto prune_begin_timepoint = std::chrono::steady_clock::now();
    uint64_t pruned_count = 0;
    it->SeekToFirst();
    if (!it->Valid()) {
        if (!it->status().ok()) {
            std::cerr << "unable to prune old checkpoints, SeekToFirst error: "
                      << it->status().ToString() << std::endl;
            return it->status();
        }

        std::cerr << "unable to prune old checkpoints, invalid SeekToFirst."
                  << std::endl;
        return rocksdb::Status::Corruption();
    }

    // Ignore 0th checkpoint
    it->Next();
    if (!it->Valid()) {
        // No data
        return rocksdb::Status::OK();
    }

    // Do not delete initial checkpoint
    it->Next();
    if (!it->Valid()) {
        // No data
        return rocksdb::Status::OK();
    }

    while (it->Valid() && (checkpoint_max_to_prune == 0 ||
                           pruned_count < checkpoint_max_to_prune)) {
        // Save current checkpoint entry
        std::vector<unsigned char> checkpoint_vector(
            it->value().data(), it->value().data() + it->value().size());
        auto checkpoint_variant = extractMachineStateKeys(checkpoint_vector);
        auto machine_output = getMachineOutput(checkpoint_variant);

        if (check_output(machine_output)) {
            // No more checkpoints to delete
            break;
        }

        // Check to see if current checkpoint is the last entry
        it->Next();
        if (!it->Valid()) {
            // Don't delete the last entry
            break;
        }

        deleteCheckpoint(tx, checkpoint_variant);
        printMachineOutputInfo("Pruned checkpoint", machine_output);
        pruned_count++;
    }
    if (!it->status().ok()) {
        std::cerr << "unable to delete old checkpoints, "
                  << "error iterating: " << it->status().ToString()
                  << std::endl;
        return it->status();
    }

    auto status = tx.commit();
    if (!status.ok()) {
        std::cerr << "unable to delete old checkpoints, "
                  << "error calling commit: " << status.ToString() << std::endl;
        return status;
    }

    if (pruned_count == 0) {
        return rocksdb::Status::NotFound();
    }

    // Calculate time including database commit
    std::cout << "Pruned " << pruned_count << " checkpoint(s)"
              << "\n";
    printElapsed(prune_begin_timepoint, "Pruning took: ");

    return rocksdb::Status::OK();
}

rocksdb::Status ArbCore::pruneToTimestampOrBefore(
    const uint256_t& timestamp,
    uint64_t checkpoint_max_to_prune) {
    return pruneCheckpoints(
        [&](const MachineOutput& output) {
            return output.last_inbox_timestamp >= timestamp;
        },
        checkpoint_max_to_prune);
}

rocksdb::Status ArbCore::pruneToGasOrBefore(const uint256_t& gas,
                                            uint64_t checkpoint_max_to_prune) {
    return pruneCheckpoints(
        [&](const MachineOutput& output) { return output.arb_gas_used >= gas; },
        checkpoint_max_to_prune);
}

void ArbCore::updateCheckpointPruningGas(uint256_t gas) {
    std::lock_guard<std::mutex> lock(checkpoint_pruning_mutex);
    unsafe_checkpoint_pruning_gas_used = gas;
}

std::string optionalUint256ToString(std::optional<uint256_t>& value) {
    if (!value.has_value()) {
        return "empty";
    }

    return intx::to_string(*value, 16);
}
