/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

void ValueCache::clear() {
    cache.clear();
}

void ValueCache::maybeSave(value val) {
    auto value_hash = hash_value(val);
    cache.emplace(value_hash, std::move(val));
}

std::optional<value> ValueCache::loadIfExists(const uint256_t& hash) {
    auto iter = cache.find(hash);
    if (iter == cache.end()) {
        return std::nullopt;
    }

    return iter->second;
}
