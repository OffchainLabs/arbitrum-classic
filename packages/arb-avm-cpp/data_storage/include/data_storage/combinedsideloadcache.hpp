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

#ifndef ARB_AVM_CPP_COMBINEDSIDELOADCACHE_HPP
#define ARB_AVM_CPP_COMBINEDSIDELOADCACHE_HPP

#include <avm/machine.hpp>
#include <avm_values/bigint.hpp>
#include <data_storage/basicsideloadcache.hpp>
#include <data_storage/lrusideloadcache.hpp>
#include <data_storage/timedsideloadcache.hpp>

class CombinedSideloadCache {
   public:
    typedef std::map<uint256_t, std::unique_ptr<Machine>> map_type;

   private:
    std::shared_mutex mutex;
    BasicSideloadCache basic;
    LRUSideloadCache lru;
    TimedSideloadCache timed;

   public:
    explicit CombinedSideloadCache(size_t basic_size,
                                   size_t lru_size,
                                   uint32_t timed_expiration_seconds)
        : basic{basic_size}, lru{lru_size}, timed{timed_expiration_seconds} {}

    void basic_add(std::unique_ptr<Machine> machine);
    void lru_add(std::unique_ptr<Machine> machine);
    void timed_add(std::unique_ptr<Machine> machine);
    size_t basic_size();
    size_t lru_size();
    size_t timed_size();
    std::unique_ptr<Machine> atOrBeforeGas(uint256_t gas_used,
                                           uint256_t existing_gas_used,
                                           uint256_t database_gas,
                                           uint256_t database_load_gas_cost,
                                           uint256_t max_execution_gas);
    void reorg(uint256_t next_gas_used);
    [[nodiscard]] uint256_t currentTimeExpired();

   private:
    std::optional<std::reference_wrapper<const Machine>> atOrBeforeGasImpl(
        uint256_t& gas_used);
};

#endif  // ARB_AVM_CPP_COMBINEDSIDELOADCACHE_HPP
