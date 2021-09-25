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

#include <data_storage/lrusideloadcache.hpp>

size_t LRUSideloadCache::size() {
    return cache.size();
}

void LRUSideloadCache::add(std::unique_ptr<Machine> machine) {
    auto gas_used = machine->machine_state.output.arb_gas_used;

    reorg(gas_used);

    if (cache.size() >= max_size) {
        // Cache is full, evict the least recently used item
        evict();
    }

    // Insert new item
    lru_list.push_front(gas_used);
    cache[gas_used] = std::make_pair(std::move(machine), lru_list.begin());
}

std::optional<LRUSideloadCache::map_type::iterator>
LRUSideloadCache::atOrBeforeGas(uint256_t gas_used) {
    // Lookup value in the cache
    auto cache_it = cache.upper_bound(gas_used);
    if (cache_it == cache.begin()) {
        // Nothing in cache
        return std::nullopt;
    }

    // Upper_bound returns the element after the one desired
    cache_it--;

    // Return the value, but first update its place in the most
    // recently used list
    auto list_it = cache_it->second.second;
    if (list_it == lru_list.begin()) {
        // The item is already at the front of the most recently
        // used list so just return iterator
        return cache_it;
    }

    // Move item to the front of the most recently used list
    lru_list.erase(list_it);
    lru_list.push_front(gas_used);

    // Update iterator in map
    cache_it->second.second = lru_list.begin();

    // Return the value
    return cache_it;
}

void LRUSideloadCache::reorg(uint256_t next_gas_used) {
    if (cache.empty()) {
        // Nothing to reorg
        return;
    }

    auto it = cache.lower_bound(next_gas_used);
    while (it != cache.end()) {
        lru_list.erase(it->second.second);
        it = cache.erase(it);
    }
}

void LRUSideloadCache::evict() {
    // Evict item from the end of most recently used list
    auto i = --lru_list.end();
    cache.erase(*i);
    lru_list.erase(i);
}