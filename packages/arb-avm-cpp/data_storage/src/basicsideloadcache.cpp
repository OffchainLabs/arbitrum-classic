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

#include <data_storage/basicsideloadcache.hpp>

size_t BasicSideloadCache::size() {
    return cache.size();
}

void BasicSideloadCache::add(std::unique_ptr<Machine> machine) {
    auto gas_used = machine->machine_state.output.arb_gas_used;

    reorg(gas_used);

    if (cache.size() >= max_size) {
        // Cache is full, evict the oldest item
        cache.erase(cache.begin());
    }

    // Add new entry
    cache[gas_used] = std::move(machine);
}

std::optional<BasicSideloadCache::map_type::iterator> BasicSideloadCache::atOrBeforeGas(uint256_t gas_used) {
    auto it = cache.upper_bound(gas_used);
    if (it == cache.begin()) {
        return std::nullopt;
    }

    // Upper_bound returns the element after the one desired
    it--;
    return it;
}

void BasicSideloadCache::reorg(uint256_t next_gas_used) {
    auto it = cache.lower_bound(next_gas_used);
    cache.erase(it, cache.end());
}

