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

#include <data_storage/combinedsideloadcache.hpp>

void CombinedSideloadCache::basic_add(std::unique_ptr<Machine> machine) {
    std::unique_lock lock(mutex);

    basic.add(std::move(machine));
}

void CombinedSideloadCache::lru_add(std::unique_ptr<Machine> machine) {
    std::unique_lock lock(mutex);

    lru.add(std::move(machine));
}

void CombinedSideloadCache::timed_add(std::unique_ptr<Machine> machine) {
    std::unique_lock lock(mutex);

    timed.add(std::move(machine));
}

size_t CombinedSideloadCache::basic_size() {
    std::shared_lock lock(mutex);

    return basic.size();
}

size_t CombinedSideloadCache::lru_size() {
    std::shared_lock lock(mutex);

    return lru.size();
}

size_t CombinedSideloadCache::timed_size() {
    std::shared_lock lock(mutex);

    return timed.size();
}

std::unique_ptr<Machine> CombinedSideloadCache::atOrBeforeGas(
    uint256_t gas_used,
    uint256_t existing_gas_used,
    uint256_t database_gas,
    uint256_t database_load_gas_cost,
    uint256_t max_execution_gas) {
    // Unique lock required to update LRU cache
    std::unique_lock lock(mutex);

    uint256_t basic_gas;
    uint256_t lru_gas;
    uint256_t timed_gas;

    auto basic_it = basic.atOrBeforeGas(gas_used);
    auto lru_it = lru.atOrBeforeGas(gas_used);
    auto timed_it = timed.atOrBeforeGas(gas_used);

    if (basic_it.has_value()) {
        basic_gas = basic_it.value()->second->machine_state.output.arb_gas_used;
    } else {
        basic_gas = 0;
    }
    if (lru_it.has_value()) {
        lru_gas =
            lru_it.value()->second.first->machine_state.output.arb_gas_used;
    } else {
        lru_gas = 0;
    }
    if (timed_it.has_value()) {
        timed_gas =
            timed_it.value()->second.machine->machine_state.output.arb_gas_used;
    } else {
        timed_gas = 0;
    }

    uint256_t best_non_db_gas;
    if (basic_gas > lru_gas && basic_gas > timed_gas) {
        best_non_db_gas = basic_gas;
    } else if (lru_gas > basic_gas && lru_gas > timed_gas) {
        best_non_db_gas = lru_gas;
    } else {
        best_non_db_gas = timed_gas;
    }

    auto load_from_database =
        (database_gas > best_non_db_gas) &&
        ((database_gas - best_non_db_gas) > database_load_gas_cost);
    if (load_from_database) {
        // Loading from database is quicker than executing last cache entry
        return nullptr;
    }

    if (existing_gas_used != 0 && existing_gas_used > best_non_db_gas) {
        // Use existing
        return nullptr;
    }

    if (gas_used - best_non_db_gas > max_execution_gas) {
        // Distance from last cache entry too far to execute
        return nullptr;
    }

    if (best_non_db_gas == basic_gas && basic_it.has_value()) {
        return std::make_unique<Machine>(*basic_it.value()->second);
    }

    if (best_non_db_gas == lru_gas && lru_it.has_value()) {
        // Update LRU order since we are actually using LRU value
        lru.updateUsed(*lru_it);
        return std::make_unique<Machine>(*lru_it.value()->second.first);
    }

    if (timed_it.has_value()) {
        return std::make_unique<Machine>(*timed_it.value()->second.machine);
    }

    return nullptr;
}

void CombinedSideloadCache::reorg(uint256_t next_gas_used) {
    std::unique_lock lock(mutex);

    basic.reorg(next_gas_used);
    lru.reorg(next_gas_used);
    timed.reorg(next_gas_used);
}

uint256_t CombinedSideloadCache::currentTimeExpired() {
    std::shared_lock lock(mutex);

    return timed.currentTimeExpired();
}
