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

struct ArbCoreConfig {
    // Maximum number of messages to process at a time
    uint32_t message_process_count{10};

    // Time it takes to run checkpoint for given gas
    // is equivalent to the time it takes to load checkpoing from database
    uint256_t checkpoint_load_gas_cost{1'000'000};

    // Maximum amount of gas to spend executing machine forward
    uint256_t checkpoint_max_execution_gas{0};

    // Frequency to save checkpoint to database
    uint256_t min_gas_checkpoint_frequency{1'000'000};

    // Amount of gas between basic cache entries
    uint32_t basic_machine_cache_interval{1'000'000};

    // Number of machines to keep in basic cache
    uint32_t basic_machine_cache_size{100};

    // Number of machines to keep in LRU cache
    uint32_t lru_machine_cache_size{20};

    // How long to keep machines in memory cache
    uint32_t timed_cache_expiration_seconds{20 * 60};

    // Seed cache on startup by forcing re-execution from timed_cache_expiration
    bool seed_cache_on_startup{false};

    // Print extra debug messages to stderr
    bool debug{false};

    // Number of seconds to wait between saving rocksdb checkpoint, 0 to disable
    uint64_t save_rocksdb_interval{0};

    // Rocksdb checkpoints will be saved in save_rocksdb_path/timestamp/
    std::string save_rocksdb_path{};

    // If any profile_* parameters are non-zero, program will exit after
    // all profile conditions are satisfied.

    // Reorg database to message
    uint64_t profile_reorg_to{0};

    // Run until message reached
    uint64_t profile_run_until{0};

    // Load specified number of machines backwards from profile_run_until
    uint64_t profile_load_count{0};

    // Delete all database entries except for inbox
    bool profile_reset_db_except_inbox{false};

    // Exit after printing out metadata from database
    bool profile_just_metadata{false};

    // Whether to lazy load the core machine
    bool lazy_load_core_machine{false};

    // Whether to lazy load archive queries
    bool lazy_load_archive_queries{false};

    ArbCoreConfig() = default;
};

#endif  // ARB_AVM_CPP_UTIL_HPP
