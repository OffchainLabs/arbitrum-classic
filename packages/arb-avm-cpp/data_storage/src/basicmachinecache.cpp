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

#include <data_storage/basicmachinecache.hpp>

size_t BasicMachineCache::size() {
    return cache.size();
}

void BasicMachineCache::add(std::unique_ptr<Machine> machine) {
    auto gas_used = machine->machine_state.output.arb_gas_used;

    if (cache.size() >= max_size) {
        if (gas_used < cache.begin()->first) {
            // New machine older than oldest, so just don't add
            return;
        }

        // Cache is full, evict the item with the least gas used
        cache.erase(cache.begin());
    }

    // Add new entry
    cache[gas_used] = std::move(machine);
}

std::optional<BasicMachineCache::map_type::const_iterator>
BasicMachineCache::atOrBeforeGas(uint256_t gas_used) {
    auto it = cache.upper_bound(gas_used);
    if (it == cache.begin()) {
        return std::nullopt;
    }

    // Upper_bound returns the element after the one desired
    it--;
    return it;
}

std::optional<BasicMachineCache::map_type::const_iterator>
BasicMachineCache::findMatching(
    const std::function<bool(const Machine&)>& check_output) {
    for (auto rit = cache.crbegin(); rit != cache.crend(); rit++) {
        if (check_output(*rit->second)) {
            auto it = rit.base();

            // The reverse_iterator::base() method returns the element after
            // the current element, so need to go back one.
            it--;
            return it;
        }
    }

    return std::nullopt;
}

void BasicMachineCache::reorg(uint256_t next_gas_used) {
    auto it = cache.lower_bound(next_gas_used);
    cache.erase(it, cache.end());
}
