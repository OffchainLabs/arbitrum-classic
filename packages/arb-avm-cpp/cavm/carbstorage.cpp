/*
 * Copyright 2019-2021, Offchain Labs, Inc.
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

#include "carbstorage.h"

#include <data_storage/aggregator.hpp>
#include <data_storage/arbstorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>
#include <data_storage/value/value.hpp>

#include <avm/machine.hpp>
#include <avm_values/value.hpp>

#include <iostream>
#include <string>

CArbStorage* createArbStorage(const char* db_path,
                              CArbCoreConfig arb_core_config) {
    auto string_filename = std::string(db_path);
    auto string_save_rocksdb_path =
        std::string(arb_core_config.database_save_path);
    ArbCoreConfig coreConfig{};
    coreConfig.message_process_count = arb_core_config.message_process_count;
    coreConfig.checkpoint_load_gas_cost =
        arb_core_config.checkpoint_load_gas_cost;
    coreConfig.checkpoint_load_gas_factor =
        arb_core_config.checkpoint_load_gas_factor;
    coreConfig.checkpoint_max_execution_gas =
        arb_core_config.checkpoint_max_execution_gas;
    coreConfig.checkpoint_gas_frequency =
        arb_core_config.checkpoint_gas_frequency;
    coreConfig.last_machine_cache = arb_core_config.last_cache;
    coreConfig.basic_machine_cache_interval =
        arb_core_config.basic_cache_interval;
    coreConfig.basic_machine_cache_size = arb_core_config.basic_cache_size;
    coreConfig.lru_machine_cache_size = arb_core_config.lru_cache_size;
    coreConfig.timed_cache_expiration_seconds =
        arb_core_config.cache_expiration_seconds;
    coreConfig.idle_sleep_milliseconds =
        arb_core_config.idle_sleep_milliseconds;
    coreConfig.seed_cache_on_startup = arb_core_config.seed_cache_on_startup;
    coreConfig.debug = arb_core_config.debug;
    coreConfig.debug_timing = arb_core_config.debug_timing;
    coreConfig.database_save_interval = arb_core_config.database_save_interval;
    coreConfig.database_save_path = string_save_rocksdb_path;
    coreConfig.lazy_load_core_machine = arb_core_config.lazy_load_core_machine;
    coreConfig.lazy_load_archive_queries =
        arb_core_config.lazy_load_archive_queries;
    coreConfig.checkpoint_prune_on_startup =
        arb_core_config.checkpoint_prune_on_startup;
    coreConfig.checkpoint_pruning_age_seconds =
        arb_core_config.checkpoint_pruning_age_seconds;
    coreConfig.checkpoint_pruning_mode =
        arb_core_config.checkpoint_pruning_mode;
    coreConfig.checkpoint_max_to_prune =
        arb_core_config.checkpoint_max_to_prune;
    coreConfig.database_compact = arb_core_config.database_compact;
    coreConfig.database_exit_after = arb_core_config.database_exit_after;
    coreConfig.test_reorg_to_l1_block = arb_core_config.test_reorg_to_l1_block;
    coreConfig.test_reorg_to_l2_block = arb_core_config.test_reorg_to_l2_block;
    coreConfig.test_reorg_to_log = arb_core_config.test_reorg_to_log;
    coreConfig.test_reorg_to_message = arb_core_config.test_reorg_to_message;
    coreConfig.test_run_until = arb_core_config.test_run_until;
    coreConfig.test_load_count = arb_core_config.test_load_count;
    coreConfig.test_reset_db_except_inbox =
        arb_core_config.test_reset_db_except_inbox;
    coreConfig.test_just_metadata = arb_core_config.test_just_metadata;

    try {
        auto storage = new ArbStorage(string_filename, coreConfig);
        return static_cast<void*>(storage);
    } catch (const std::exception& e) {
        std::cerr << "Error creating storage: " << e.what() << std::endl;
        return nullptr;
    }
}

void printDatabaseMetadata(CArbStorage* storage_ptr) {
    auto storage = static_cast<ArbStorage*>(storage_ptr);
    storage->printDatabaseMetadata();
}

int initializeArbStorage(CArbStorage* storage_ptr,
                         const char* executable_path) {
    auto storage = static_cast<ArbStorage*>(storage_ptr);
    try {
        auto result = storage->initialize(executable_path);
        if (result.finished) {
            // Gracefully shutdown
            return false;
        }
        if (!result.status.ok()) {
            std::cerr << "Error initializing storage: "
                      << result.status.ToString() << std::endl;
            return false;
        }

        return true;
    } catch (const std::exception& e) {
        std::cerr << "Exception initializing storage:" << e.what() << std::endl;
        return false;
    }
}

int arbStorageInitialized(CArbStorage* storage_ptr) {
    return static_cast<ArbStorage*>(storage_ptr)->initialized();
}

int closeArbStorage(CArbStorage* storage_ptr) {
    auto storage = static_cast<ArbStorage*>(storage_ptr);
    return storage->closeArbStorage();
}

void destroyArbStorage(CArbStorage* storage_ptr) {
    auto storage = static_cast<ArbStorage*>(storage_ptr);
    if (storage == nullptr) {
        return;
    }
    storage->closeArbStorage();
    delete static_cast<ArbStorage*>(storage);
}

CArbCore* createArbCore(CArbStorage* storage_ptr) {
    auto storage = static_cast<ArbStorage*>(storage_ptr);
    return storage->getArbCore().get();
}

CAggregatorStore* createAggregatorStore(CArbStorage* storage_ptr) {
    auto storage = static_cast<ArbStorage*>(storage_ptr);
    return storage->getAggregatorStore().release();
}
