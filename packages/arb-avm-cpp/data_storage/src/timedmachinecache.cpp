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

#include <data_storage/timedmachinecache.hpp>

size_t TimedMachineCache::size() {
    return cache.size();
}

void TimedMachineCache::add(std::unique_ptr<Machine> machine) {
    auto gas_used = machine->machine_state.output.arb_gas_used;
    auto timestamp = machine->machine_state.output.last_inbox_timestamp;

    if (timestamp <= expiredTimestamp()) {
        // Don't save expired machine to cache
        return;
    }

    // Add new entry
    cache[gas_used].timestamp = timestamp;
    cache[gas_used].machine = std::move(machine);

    deleteExpired();
}

std::optional<TimedMachineCache::map_type::const_iterator>
TimedMachineCache::atOrBeforeGas(uint256_t gas_used) {
    auto it = cache.upper_bound(gas_used);
    if (it == cache.begin()) {
        return std::nullopt;
    }

    // Upper_bound returns the element after the one desired
    it--;

    return it;
}

std::optional<TimedMachineCache::map_type::const_iterator>
TimedMachineCache::findMatching(
    const std::function<bool(const MachineOutput&)>& check_output) {
    for (auto rit = cache.crbegin(); rit != cache.crend(); rit++) {
        if (check_output(rit->second.machine->machine_state.output)) {
            auto it = rit.base();
            it--;
            return it;
        }
    }

    return std::nullopt;
}

void TimedMachineCache::reorg(uint256_t next_gas_used) {
    if (cache.empty()) {
        return;
    }
    if (next_gas_used <= cache.begin()->first) {
        // Remove everything
        cache.clear();

        return;
    }

    auto it = cache.lower_bound(next_gas_used);
    cache.erase(it, cache.end());
}

void TimedMachineCache::deleteExpired() {
    auto expired = expiredTimestamp();

    for (auto it = cache.cbegin();
         it != cache.cend() && it->second.timestamp <= expired;) {
        it = cache.erase(it);
    }
}

uint256_t TimedMachineCache::expiredTimestamp() {
    if (cache.empty()) {
        return 0;
    }

    // Expire items based on the last timestamp added.  This is so the cache
    // is valid relative to the last block synced, even if syncing old machines.
    return cache.crbegin()->second.timestamp - expiration_seconds;
}

uint256_t TimedMachineCache::currentTimeExpired() const {
    return std::time(nullptr) - expiration_seconds;
}
