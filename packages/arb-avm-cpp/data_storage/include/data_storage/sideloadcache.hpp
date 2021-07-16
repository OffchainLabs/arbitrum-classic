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

#ifndef ARB_AVM_CPP_SIDELOADCACHE_H
#define ARB_AVM_CPP_SIDELOADCACHE_H

#include <avm/machine.hpp>
#include <avm_values/bigint.hpp>

#include <shared_mutex>

struct Record {
    uint256_t timestamp;

    std::unique_ptr<Machine> machine;
};

class SideloadCache {
   private:
    std::mutex mutex;
    std::map<uint256_t, Record> cache;

    const int expiration_seconds;

   public:
    explicit SideloadCache(int expiration_seconds)
        : expiration_seconds{expiration_seconds} {}

    size_t size();
    uint256_t oldestHeight();
    uint256_t nextHeight();
    void add(uint256_t height,
             uint256_t timestamp,
             std::unique_ptr<Machine> machine);
    std::unique_ptr<Machine> get(uint256_t height);
    void reorg(uint256_t next_height);

   private:
    void reorgNoLock(uint256_t new_next_height);
    void deleteExpiredNoLock();
    [[nodiscard]] uint256_t expiredTimestamp() const;
};

#endif  // ARB_AVM_CPP_SIDELOADCACHE_H
