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

#include <data_storage/sideloadcache.hpp>

size_t SideloadCache::size() {
    return cache.size();
}

uint256_t SideloadCache::oldestHeight() {
    if (cache.empty()) {
        return 0;
    }

    return cache.begin()->first;
}

uint256_t SideloadCache::nextHeight() {
    if (cache.empty()) {
        return 0;
    }

    auto it = cache.end();
    it--;
    return it->first + 1;
}

void SideloadCache::add(uint256_t height,
                        uint256_t timestamp,
                        std::unique_ptr<Machine> machine) {
    std::lock_guard<std::mutex> guard(mutex);

    reorgNoLock(height);
    deleteExpiredNoLock();

    if (timestamp <= expiredTimestamp()) {
        // Don't save expired machine to cache
        return;
    }

    // Add new entry
    cache[height].timestamp = timestamp;
    cache[height].machine = std::move(machine);
}

std::unique_ptr<Machine> SideloadCache::get(uint256_t height) {
    std::lock_guard<std::mutex> guard(mutex);

    auto it = cache.find(height);
    if (it == cache.end()) {
        return nullptr;
    }

    return std::make_unique<Machine>(*it->second.machine);
}

void SideloadCache::SideloadCache::reorg(uint256_t next_height) {
    std::lock_guard<std::mutex> guard(mutex);

    reorgNoLock(next_height);
}

void SideloadCache::reorgNoLock(uint256_t next_height) {
    if (next_height <= cache.begin()->first) {
        // Remove everything
        cache.clear();
    }

    for (auto rit = cache.crbegin();
         rit != cache.crend() && rit->first >= next_height;) {
        rit = decltype(rit){cache.erase(std::next(rit).base())};
    }
}

void SideloadCache::deleteExpiredNoLock() {
    auto expired = expiredTimestamp();
    for (auto it = cache.cbegin();
         it != cache.cend() && it->second.timestamp <= expired;) {
        it = cache.erase(it);
    }
}

uint256_t SideloadCache::expiredTimestamp() const {
    return std::time(nullptr) - expiration_seconds;
}