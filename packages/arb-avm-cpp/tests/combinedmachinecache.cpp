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

#include <unistd.h>

#include "config.hpp"
#include "helper.hpp"

#include <data_storage/combinedmachinecache.hpp>

#include <catch2/catch.hpp>

TEST_CASE("CombinedMachineCache add and get") {
    ArbCoreConfig coreConfig;
    coreConfig.basic_machine_cache_size = 2;
    coreConfig.lru_machine_cache_size = 2;
    coreConfig.timed_cache_expiration_seconds = 20;
    coreConfig.checkpoint_load_gas_cost = 10;
    coreConfig.checkpoint_max_execution_gas = 100'000'000;
    CombinedMachineCache cache(coreConfig);

    // Test empty cache case
    REQUIRE(cache.atOrBeforeGas(50, std::nullopt, std::nullopt, true).machine ==
            nullptr);

    // Test empty findmatching
    auto check_machine_state = [&](const MachineState& mach) {
        return mach.output.arb_gas_used <= 50;
    };
    REQUIRE(cache
                .findFirstMatching(check_machine_state, std::nullopt,
                                   std::nullopt, true)
                .machine == nullptr);

    // Test that basic entry is added
    auto machine40 = std::make_unique<Machine>(getComplexMachine());
    machine40->machine_state.output.arb_gas_used = 40;
    cache.lastAdd(std::move(machine40));
    auto machine41 = std::make_unique<Machine>(getComplexMachine());
    machine41->machine_state.output.arb_gas_used = 41;
    cache.basicAdd(std::move(machine41));
    REQUIRE(cache.basicSize() == 1);
    auto machine41a = cache.atOrBeforeGas(50, std::nullopt, std::nullopt, true);
    REQUIRE(machine41a.machine != nullptr);
    REQUIRE(machine41a.machine->machine_state.output.arb_gas_used == 41);

    // Test that lru entry is added
    auto machine42 = std::make_unique<Machine>(getComplexMachine());
    machine42->machine_state.output.arb_gas_used = 42;
    cache.lruAdd(std::move(machine42));
    REQUIRE(cache.lruSize() == 1);
    auto machine42a = cache.atOrBeforeGas(50, std::nullopt, std::nullopt, true);
    REQUIRE(machine42a.machine != nullptr);
    REQUIRE(machine41a.machine->machine_state.output.arb_gas_used == 41);
    REQUIRE(machine42a.machine->machine_state.output.arb_gas_used == 42);

    // Test that timed entry is added
    auto machine43 = std::make_unique<Machine>(getComplexMachine());
    machine43->machine_state.output.arb_gas_used = 43;
    machine43->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    cache.timedAdd(std::move(machine43));
    REQUIRE(cache.timedSize() == 1);
    auto machine43a = cache.atOrBeforeGas(50, std::nullopt, std::nullopt, true);
    REQUIRE(machine43a.machine != nullptr);
    REQUIRE(machine41a.machine->machine_state.output.arb_gas_used == 41);
    REQUIRE(machine42a.machine->machine_state.output.arb_gas_used == 42);
    REQUIRE(machine43a.machine->machine_state.output.arb_gas_used == 43);

    // Test execution cheaper than database load
    auto machineDBa = cache.atOrBeforeGas(50, std::nullopt, 53, true);
    REQUIRE(machineDBa.machine != nullptr);
    REQUIRE(machineDBa.machine->machine_state.output.arb_gas_used == 43);

    // Test database load cheaper than execution
    auto machineDBb = cache.atOrBeforeGas(54, std::nullopt, 54, true);
    REQUIRE(machineDBb.machine == nullptr);
    REQUIRE(machineDBb.status == CombinedMachineCache::UseDatabase);

    // Test current machine closer than database load
    auto machineDBc = cache.atOrBeforeGas(53, 53, 52, true);
    REQUIRE(machineDBc.machine == nullptr);
    REQUIRE(machineDBc.status == CombinedMachineCache::UseExisting);

    // Test distance from last cache entry too far away
    auto machineDBd = cache.atOrBeforeGas(
        80 + coreConfig.checkpoint_max_execution_gas, 0, 53, true);
    REQUIRE(machineDBd.machine == nullptr);
    REQUIRE(machineDBd.status == CombinedMachineCache::TooMuchExecution);

    // Test distance from last database entry too far away
    auto machineDBe = cache.atOrBeforeGas(
        80 + coreConfig.checkpoint_max_execution_gas, 0, 60, true);
    REQUIRE(machineDBe.machine == nullptr);
    REQUIRE(machineDBe.status == CombinedMachineCache::TooMuchExecution);

    // Test distance from existing entry too far away
    auto machineDBf = cache.atOrBeforeGas(
        80 + coreConfig.checkpoint_max_execution_gas, 60, 53, true);
    REQUIRE(machineDBf.machine == nullptr);
    REQUIRE(machineDBf.status == CombinedMachineCache::TooMuchExecution);

    // Test that last machine entry is added
    cache.reorg(0);
    auto machine44 = std::make_unique<Machine>(getComplexMachine());
    machine44->machine_state.output.arb_gas_used = 44;
    machine44->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    cache.lastAdd(std::move(machine44));
    auto machine44a = cache.atOrBeforeGas(50, std::nullopt, std::nullopt, true);
    REQUIRE(machine44a.machine != nullptr);
    REQUIRE(machine44a.machine->machine_state.output.arb_gas_used == 44);

    // Test that last last machine entry is added
    auto machine45 = std::make_unique<Machine>(getComplexMachine());
    machine45->machine_state.output.arb_gas_used = 45;
    machine45->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    cache.lastAdd(std::move(machine45));
    auto machine45a = cache.atOrBeforeGas(50, std::nullopt, std::nullopt, true);
    REQUIRE(machine45a.machine != nullptr);
    REQUIRE(machine45a.machine->machine_state.output.arb_gas_used == 45);
    machine44a = cache.atOrBeforeGas(44, std::nullopt, std::nullopt, true);
    REQUIRE(machine44a.machine != nullptr);
    REQUIRE(machine44a.machine->machine_state.output.arb_gas_used == 44);

    // Test that last and last last machine entry is removed appropriately
    auto machine46 = std::make_unique<Machine>(getComplexMachine());
    machine46->machine_state.output.arb_gas_used = 46;
    machine46->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    cache.lruAdd(std::move(machine46));
    auto machine46a = cache.atOrBeforeGas(50, std::nullopt, std::nullopt, true);
    REQUIRE(machine46a.machine != nullptr);
    REQUIRE(machine46a.machine->machine_state.output.arb_gas_used == 46);

    // Test that last entries are removed on reorg
    machine44 = std::make_unique<Machine>(getComplexMachine());
    machine44->machine_state.output.arb_gas_used = 44;
    machine44->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    cache.lastAdd(std::move(machine44));
    machine45 = std::make_unique<Machine>(getComplexMachine());
    machine45->machine_state.output.arb_gas_used = 45;
    machine45->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    cache.lastAdd(std::move(machine45));
    cache.reorg(0);
    machine46a = cache.atOrBeforeGas(50, std::nullopt, std::nullopt, true);
    REQUIRE(machine46a.machine == nullptr);

    // Test only lru
    cache.reorg(0);
    machine41 = std::make_unique<Machine>(getComplexMachine());
    machine41->machine_state.output.arb_gas_used = 41;
    cache.lastAdd(std::move(machine41));
    machine42 = std::make_unique<Machine>(getComplexMachine());
    machine42->machine_state.output.arb_gas_used = 42;
    cache.lruAdd(std::move(machine42));
    machine42a = cache.atOrBeforeGas(50, std::nullopt, std::nullopt, true);
    REQUIRE(machine42a.machine != nullptr);
    REQUIRE(machine42a.machine->machine_state.output.arb_gas_used == 42);

    // Test match
    machine42a = cache.findFirstMatching(check_machine_state, std::nullopt,
                                         std::nullopt, true);
    REQUIRE(machine42a.machine != nullptr);
    REQUIRE(machine42a.machine->machine_state.output.arb_gas_used == 42);

    // Test only timed
    cache.reorg(0);
    machine41 = std::make_unique<Machine>(getComplexMachine());
    machine41->machine_state.output.arb_gas_used = 41;
    cache.lastAdd(std::move(machine41));
    machine42 = std::make_unique<Machine>(getComplexMachine());
    machine42->machine_state.output.arb_gas_used = 42;
    machine42->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    cache.timedAdd(std::move(machine42));
    machine42a = cache.atOrBeforeGas(50, std::nullopt, std::nullopt, true);
    REQUIRE(machine42a.machine != nullptr);
    REQUIRE(machine42a.machine->machine_state.output.arb_gas_used == 42);
}

TEST_CASE("CombinedMachineCache currentTimeExpired") {
    ArbCoreConfig coreConfig;
    coreConfig.basic_machine_cache_size = 2;
    coreConfig.lru_machine_cache_size = 2;
    coreConfig.timed_cache_expiration_seconds = 20;
    coreConfig.checkpoint_load_gas_cost = 10;
    coreConfig.checkpoint_max_execution_gas = 100'000'000;
    CombinedMachineCache cache(coreConfig);
    auto expiration_fudge_factor = 10;

    auto expired = cache.currentTimeExpired();
    REQUIRE(expired > std::time(nullptr) -
                          coreConfig.timed_cache_expiration_seconds -
                          expiration_fudge_factor);
    REQUIRE(expired < std::time(nullptr) -
                          coreConfig.timed_cache_expiration_seconds +
                          expiration_fudge_factor);
}
