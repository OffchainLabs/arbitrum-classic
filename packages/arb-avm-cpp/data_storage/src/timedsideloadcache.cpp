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

#include <data_storage/timedsideloadcache.hpp>

size_t TimedSideloadCache::size() {
    return cache.size();
}

void TimedSideloadCache::add(std::unique_ptr<Machine> machine) {
    auto gas_used = machine->machine_state.output.arb_gas_used;
    auto timestamp = machine->machine_state.output.last_inbox_timestamp;

    reorg(gas_used);
    deleteExpired();

    if (timestamp <= expiredTimestamp()) {
        // Don't save expired machine to cache
        return;
    }

    // Add new entry
    cache[gas_used].timestamp = timestamp;
    cache[gas_used].machine = std::move(machine);
}

uint256_t TimedSideloadCache::peekAtOrBeforeGas(uint256_t gas_used) {
    auto it = cache.upper_bound(gas_used);
    if (it == cache.begin()) {
        return 0;
    }

    // Upper_bound returns the element after the one desired
    it--;
    return it->second.machine->machine_state.output.arb_gas_used;
}

std::optional<TimedSideloadCache::map_type::iterator> TimedSideloadCache::atOrBeforeGas(uint256_t gas_used) {
    auto it = cache.upper_bound(gas_used);
    if (it == cache.begin()) {
        return std::nullopt;
    }

    // Upper_bound returns the element after the one desired
    it--;

    return it;
}

void TimedSideloadCache::reorg(uint256_t next_gas_used) {
    if (next_gas_used <= cache.begin()->first) {
        // Remove everything
        cache.clear();

        return;
    }

    auto it = cache.lower_bound(next_gas_used);
    cache.erase(it, cache.end());
}

void TimedSideloadCache::deleteExpired() {
    auto expired = expiredTimestamp();
    for (auto it = cache.cbegin();
         it != cache.cend() && it->second.timestamp <= expired;) {
        it = cache.erase(it);
    }
}

uint256_t TimedSideloadCache::expiredTimestamp() const {
    return std::time(nullptr) - expiration_seconds;
}
