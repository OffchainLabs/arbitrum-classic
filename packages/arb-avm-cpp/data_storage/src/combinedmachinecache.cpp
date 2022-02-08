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

#include <data_storage/combinedmachinecache.hpp>

void CombinedMachineCache::checkLastMachineNoLock(uint256_t& arb_gas_used) {
    if (last_machine &&
        last_machine->machine_state.output.arb_gas_used < arb_gas_used) {
        last_machine = nullptr;
    }
    if (last_last_machine &&
        last_last_machine->machine_state.output.arb_gas_used < arb_gas_used) {
        last_last_machine = nullptr;
    }
}

void CombinedMachineCache::lastAdd(std::unique_ptr<Machine> machine) {
    std::unique_lock lock(mutex);

    if (last_machine) {
        last_last_machine = std::move(last_machine);
    }
    last_machine = std::move(machine);
}

void CombinedMachineCache::basicAdd(std::unique_ptr<Machine> machine) {
    std::unique_lock lock(mutex);

    checkLastMachineNoLock(machine->machine_state.output.arb_gas_used);
    basic.add(std::move(machine));
}

void CombinedMachineCache::lruAdd(std::unique_ptr<Machine> machine) {
    std::unique_lock lock(mutex);

    checkLastMachineNoLock(machine->machine_state.output.arb_gas_used);
    lru.add(std::move(machine));
}

void CombinedMachineCache::timedAdd(std::unique_ptr<Machine> machine) {
    std::unique_lock lock(mutex);

    checkLastMachineNoLock(machine->machine_state.output.arb_gas_used);
    timed.add(std::move(machine));
}

size_t CombinedMachineCache::basicSize() {
    std::shared_lock lock(mutex);

    return basic.size();
}

size_t CombinedMachineCache::lruSize() {
    std::shared_lock lock(mutex);

    return lru.size();
}

size_t CombinedMachineCache::timedSize() {
    std::shared_lock lock(mutex);

    return timed.size();
}

std::optional<std::reference_wrapper<const Machine>>
CombinedMachineCache::getFirstMatchNoLock(
    const std::function<bool(const MachineState&)>& check_machine_state,
    const std::optional<BasicMachineCache::map_type::const_iterator>& basic_it,
    const std::optional<LRUMachineCache::map_type::const_iterator>& lru_it,
    const std::optional<TimedMachineCache::map_type::const_iterator>&
        timed_it) {
    uint256_t basic_gas;
    uint256_t lru_gas;
    uint256_t timed_gas;

    if (last_machine && check_machine_state(last_machine->machine_state)) {
        // Last machine will always have the greatest amount of gas used
        return std::cref(*last_machine);
    }

    if (last_last_machine &&
        check_machine_state(last_last_machine->machine_state)) {
        // Last last machine will always have the next greatest amount of gas
        // used
        return std::cref(*last_last_machine);
    }

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

    if (basic_gas >= lru_gas && basic_gas >= timed_gas &&
        basic_it.has_value()) {
        return std::cref(*basic_it.value()->second);
    }

    if (lru_gas >= timed_gas && lru_it.has_value()) {
        return std::cref(*lru_it.value()->second.first);
    }

    if (timed_it.has_value()) {
        return std::cref(*timed_it.value()->second.machine);
    }

    return std::nullopt;
}

CombinedMachineCache::CacheResultStruct CombinedMachineCache::atOrBeforeGas(
    uint256_t gas_used,
    std::optional<uint256_t> existing_gas_used,
    std::optional<uint256_t> database_gas,
    bool use_max_execution) {
    // Unique lock required to update LRU cache
    std::unique_lock lock(mutex);

    auto basic_it = basic.atOrBeforeGas(gas_used);
    auto lru_it = lru.atOrBeforeGas(gas_used);
    auto timed_it = timed.atOrBeforeGas(gas_used);

    auto check_machine_state = [&](const MachineState& mach) {
        return mach.output.arb_gas_used <= gas_used;
    };

    auto cache_machine =
        getFirstMatchNoLock(check_machine_state, basic_it, lru_it, timed_it);

    return findBestMachineNoLock(gas_used, cache_machine, existing_gas_used,
                                 database_gas, use_max_execution);
}

// checkSimpleMatching just checks the last two items added to cache
CombinedMachineCache::CacheResultStruct
CombinedMachineCache::checkSimpleMatching(
    const std::function<bool(const MachineOutput&)>& check_output) {
    if (last_machine && check_output(last_machine->machine_state.output)) {
        return {std::make_unique<Machine>(*last_machine), Success};
    }

    if (last_last_machine &&
        check_output(last_last_machine->machine_state.output)) {
        return {std::make_unique<Machine>(*last_machine), Success};
    }

    return {nullptr, NotFound};
}

CombinedMachineCache::CacheResultStruct CombinedMachineCache::findFirstMatching(
    const std::function<bool(const MachineState&)>& check_machine_state,
    std::optional<uint256_t> existing_gas_used,
    std::optional<uint256_t> database_gas,
    bool use_max_execution) {
    // Unique lock required to update LRU cache
    std::shared_lock lock(mutex);

    auto basic_it = basic.findMatching(check_machine_state);
    // Ignore the LRU cache since
    // A) it could contain lazy loaded machines
    // B) This function is just used to load core thread machines where the LRU
    // cache is unlikely to be correct C) The LRU cache requires a unique lock
    auto timed_it = timed.findMatching(check_machine_state);

    auto cache_machine = getFirstMatchNoLock(check_machine_state, basic_it,
                                             std::nullopt, timed_it);

    return findBestMachineNoLock(std::nullopt, cache_machine, existing_gas_used,
                                 database_gas, use_max_execution);
}

CombinedMachineCache::CacheResultStruct
CombinedMachineCache::findBestMachineNoLock(
    std::optional<uint256_t> current_gas_used,
    std::optional<std::reference_wrapper<const Machine>> cache_machine,
    std::optional<uint256_t> existing_gas_used,
    std::optional<uint256_t> database_gas,
    bool use_max_execution) {
    std::optional<uint256_t> best_non_database_gas;
    best_non_database_gas = existing_gas_used;

    std::optional<uint256_t> cache_gas;
    if (cache_machine.has_value()) {
        cache_gas =
            cache_machine.value().get().machine_state.output.arb_gas_used;

        if (!best_non_database_gas.has_value() ||
            cache_gas.value() > best_non_database_gas.value()) {
            best_non_database_gas = cache_gas.value();
        }
    }

    auto load_from_database =
        (database_gas.has_value() &&
         (!best_non_database_gas.has_value() ||
          ((database_gas.value() > best_non_database_gas.value()) &&
           ((database_gas.value() - best_non_database_gas.value()) >
            database_load_gas_cost))));
    if (load_from_database) {
        if (current_gas_used.has_value() && use_max_execution &&
            (max_execution_gas != 0) &&
            (current_gas_used.value() - database_gas.value() >
             max_execution_gas)) {
            // Distance from last database entry too far to execute
            return {nullptr, TooMuchExecution};
        }

        // Loading from database is quicker than executing last cache entry
        return {nullptr, UseDatabase};
    }

    if (existing_gas_used.has_value() && existing_gas_used > cache_gas) {
        if (current_gas_used.has_value() && use_max_execution &&
            (max_execution_gas != 0) &&
            (current_gas_used.value() - existing_gas_used.value() >
             max_execution_gas)) {
            // Distance from existing entry too far to execute
            return {nullptr, TooMuchExecution};
        }

        // Use existing
        return {nullptr, UseExisting};
    }

    if (cache_machine.has_value()) {
        if (current_gas_used.has_value() && use_max_execution &&
            (max_execution_gas != 0) &&
            (current_gas_used.value() - cache_gas.value() >
             max_execution_gas)) {
            // Distance from last cache entry too far to execute
            return {nullptr, TooMuchExecution};
        }

        return {std::make_unique<Machine>(cache_machine.value().get()),
                Success};
    }

    return {nullptr, NotFound};
}

void CombinedMachineCache::reorg(uint256_t next_gas_used) {
    std::unique_lock lock(mutex);

    if (last_machine &&
        last_machine->machine_state.output.arb_gas_used >= next_gas_used) {
        last_machine = nullptr;
    }
    if (last_last_machine &&
        last_last_machine->machine_state.output.arb_gas_used >= next_gas_used) {
        last_last_machine = nullptr;
    }
    basic.reorg(next_gas_used);
    lru.reorg(next_gas_used);
    timed.reorg(next_gas_used);
}

uint256_t CombinedMachineCache::currentTimeExpired() {
    std::shared_lock lock(mutex);

    return timed.currentTimeExpired();
}
