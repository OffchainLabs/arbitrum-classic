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

#include <data_storage/value/valuecache.hpp>

void ValueCache::maybeSave(value val) {
    if (caches.empty()) {
        return;
    }

    auto value_hash = hash_value(val);
    caches[saving_cache_index].emplace(value_hash, std::move(val));

    if (max_cache_size > 0 &&
        caches[saving_cache_index].size() >= max_cache_size) {
        nextCache();
    }
}

std::optional<value> ValueCache::loadIfExists(const uint256_t& hash) {
    if (caches.empty()) {
        return std::nullopt;
    }

    // Check cache currently saving to
    auto saving_iter = caches[saving_cache_index].find(hash);
    if (saving_iter != caches[saving_cache_index].end()) {
        // Exists in cache currently saving to
        return saving_iter->second;
    }

    // Check all the older caches
    for (size_t i = 1,
                current_cache_index = (saving_cache_index + 1) % caches.size();
         i < caches.size();
         i++, current_cache_index = (current_cache_index + 1) % caches.size()) {
        auto iter = caches[current_cache_index].find(hash);
        if (iter != caches[current_cache_index].end()) {
            // Exists in one of the older caches
            if (caches[saving_cache_index].find(hash) ==
                caches[saving_cache_index].end()) {
                // Need to add it to cache cache currently saving to
                caches[saving_cache_index].insert(
                    std::make_pair(hash, iter->second));
            }
            return iter->second;
        }
    }

    if (max_cache_size > 0 &&
        caches[saving_cache_index].size() >= max_cache_size) {
        nextCache();
    }

    return std::nullopt;
}

void ValueCache::nextCache() {
    if (caches.empty()) {
        return;
    }

    saving_cache_index = (saving_cache_index + 1) % caches.size();
    caches[saving_cache_index].clear();
}
