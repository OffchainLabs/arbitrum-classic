/*
 * Copyright 2021, Offchain Labs, Inc.
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

#ifndef ARB_AVM_CPP_UTIL_HPP
#define ARB_AVM_CPP_UTIL_HPP

#include "cpruningmode.h"

struct ArbCoreConfig {
    // Maximum number of messages to process at a time
    uint32_t message_process_count{10};

    // Time it takes to run checkpoint for given gas
    // is equivalent to the time it takes to load checkpoing from database
    uint256_t checkpoint_load_gas_cost{1'000'000};

    // When checkpoint is loaded from database with lazy loading,
    // the remaining gas needed to execute is more expensive
    // because it requires additional loads from database.
    uint256_t checkpoint_load_gas_factor{4};

    // Maximum amount of gas to spend executing machine forward
    uint256_t checkpoint_max_execution_gas{0};

    // Frequency to save checkpoint to database
    uint256_t checkpoint_gas_frequency{1'000'000};

    // Amount of gas between basic cache entries
    uint32_t basic_machine_cache_interval{1'000'000};

    // Number of machines to keep in basic cache
    uint32_t basic_machine_cache_size{100};

    // Number of machines to keep in LRU cache
    uint32_t lru_machine_cache_size{20};

    // How long to keep machines in memory cache
    uint32_t timed_cache_expiration_seconds{20 * 60};

    // Number of milliseconds to sleep when idle
    uint32_t idle_sleep_milliseconds{5};

    // Seed cache on startup by forcing re-execution from timed_cache_expiration
    bool seed_cache_on_startup{false};

    // Print extra debug messages to stderr
    bool debug{false};

    // Number of seconds to wait between saving rocksdb checkpoint, 0 to disable
    uint64_t database_save_interval{0};

    // Rocksdb checkpoints will be saved in database_save_path/timestamp/
    std::string database_save_path{};

    // If any profile_* parameters are non-zero, program will exit after
    // all profile conditions are satisfied.

    // Reorg database to l1 block
    uint64_t test_reorg_to_l1_block{0};

    // Reorg database to l2 block
    uint64_t test_reorg_to_l2_block{0};

    // Reorg database to log
    uint64_t test_reorg_to_log{0};

    // Reorg database to message
    uint64_t test_reorg_to_message{0};

    // Run until message reached
    uint64_t test_run_until{0};

    // Load specified number of machines backwards from test_run_until
    uint64_t test_load_count{0};

    // Delete all database entries except for inbox
    bool test_reset_db_except_inbox{false};

    // Exit after printing out metadata from database
    bool test_just_metadata{false};

    // Whether to lazy load the core machine
    bool lazy_load_core_machine{false};

    // Whether to lazy load archive queries
    bool lazy_load_archive_queries{false};

    // Do complete prune on startup
    bool checkpoint_prune_on_startup{false};

    // Perform database compaction
    bool database_compact{false};

    // Exit after manipulating database
    bool database_exit_after{false};

    // Number of seconds to keep checkpoints
    uint64_t checkpoint_pruning_age_seconds{0};

    // Number of seconds to keep checkpoints
    CPruningMode checkpoint_pruning_mode{PRUNING_MODE_DEFAULT};

    // Maximum number of checkpoints to prune at a time
    uint64_t checkpoint_max_to_prune{0};

    ArbCoreConfig() = default;
};

#endif  // ARB_AVM_CPP_UTIL_HPP
