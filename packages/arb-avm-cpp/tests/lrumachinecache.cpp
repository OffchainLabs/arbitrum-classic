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

#include <data_storage/lrumachinecache.hpp>

#include <catch2/catch.hpp>

TEST_CASE("LRUMachineCache add") {
    auto cache_size = 2;
    LRUMachineCache cache(cache_size);

    // Test empty cache case
    REQUIRE_FALSE(cache.atOrBeforeGas(50).has_value());

    // Test empty findMatching
    auto check_machine_state = [&](const MachineState& mach) {
        return mach.output.arb_gas_used <= 10;
    };
    REQUIRE_FALSE(cache.findMatching(check_machine_state).has_value());

    // Basic add
    auto machine_zero = std::make_unique<Machine>(getComplexMachine());
    machine_zero->machine_state.output.arb_gas_used = 10;
    cache.add(std::move(machine_zero));
    REQUIRE(cache.size() == 1);

    // Adding machines with less gas is okay
    auto machine_one = std::make_unique<Machine>(getComplexMachine());
    machine_one->machine_state.output.arb_gas_used = 5;
    cache.add(std::move(machine_one));
    REQUIRE(cache.size() == 2);

    // Test that cache_size limit is not breached
    auto machine_two = std::make_unique<Machine>(getComplexMachine());
    machine_two->machine_state.output.arb_gas_used = 20;
    cache.add(std::move(machine_two));
    REQUIRE(cache.size() == 2);
    auto retrieved_machine2 = cache.atOrBeforeGas(10);
    REQUIRE(retrieved_machine2.has_value());
    REQUIRE(retrieved_machine2.value()
                ->second.first->machine_state.output.arb_gas_used == 5);

    // Test findMatching
    retrieved_machine2 = cache.findMatching(check_machine_state);
    REQUIRE(retrieved_machine2.has_value());
    REQUIRE(retrieved_machine2.value()
                ->second.first->machine_state.output.arb_gas_used == 5);
}
TEST_CASE("LRUMachineCache get") {
    auto cache_size = 3;
    LRUMachineCache cache(cache_size);

    auto orig_machine = Machine::loadFromFile(
        std::string(machine_test_cases_path) + "/sideloadtest.mexe");

    auto machine1a = std::make_unique<Machine>(orig_machine);
    auto gas1a = machine1a->machine_state.output.arb_gas_used;

    auto machine2a = std::make_unique<Machine>(*machine1a);
    machine2a->machine_state.runOne();
    auto gas2a = machine2a->machine_state.output.arb_gas_used;
    REQUIRE(gas1a != gas2a);

    auto machine3a = std::make_unique<Machine>(*machine2a);
    machine3a->machine_state.runOne();
    auto gas3a = machine3a->machine_state.output.arb_gas_used;
    REQUIRE(gas2a != gas3a);

    cache.add(std::move(machine1a));
    cache.add(std::move(machine2a));
    cache.add(std::move(machine3a));
    REQUIRE(cache.size() == 3);

    auto machine1b = cache.atOrBeforeGas(gas1a);
    REQUIRE(machine1b.has_value());
    REQUIRE(gas1a ==
            machine1b.value()->second.first->machine_state.output.arb_gas_used);

    auto machine2b = cache.atOrBeforeGas(gas2a);
    REQUIRE(machine2b.has_value());
    REQUIRE(gas2a ==
            machine2b.value()->second.first->machine_state.output.arb_gas_used);

    auto machine3b = cache.atOrBeforeGas(gas3a + 100);
    REQUIRE(machine3b.has_value());
    REQUIRE(gas3a ==
            machine3b.value()->second.first->machine_state.output.arb_gas_used);
}

TEST_CASE("LRUMachineCache reorg") {
    auto cache_size = 30;
    LRUMachineCache cache(cache_size);

    auto orig_machine = Machine::loadFromFile(
        std::string(machine_test_cases_path) + "/sideloadtest.mexe");

    auto machine0 = std::make_unique<Machine>(orig_machine);
    auto gas0 = machine0->machine_state.output.arb_gas_used;

    auto machine1 = std::make_unique<Machine>(*machine0);
    machine1->machine_state.runOne();
    auto gas1 = machine1->machine_state.output.arb_gas_used;
    REQUIRE(gas0 != gas1);

    auto machine2 = std::make_unique<Machine>(*machine1);
    machine2->machine_state.runOne();
    auto gas2 = machine2->machine_state.output.arb_gas_used;
    REQUIRE(gas1 != gas2);

    auto machine3 = std::make_unique<Machine>(*machine2);
    machine3->machine_state.runOne();
    auto gas3 = machine3->machine_state.output.arb_gas_used;
    REQUIRE(gas2 != gas3);

    auto machine4 = std::make_unique<Machine>(*machine3);
    auto machine5 = std::make_unique<Machine>(*machine4);
    auto machine6 = std::make_unique<Machine>(*machine5);
    auto machine7 = std::make_unique<Machine>(*machine6);
    auto machine8 = std::make_unique<Machine>(*machine7);
    auto machine9 = std::make_unique<Machine>(*machine8);

    machine0->machine_state.output.arb_gas_used = 0;
    machine1->machine_state.output.arb_gas_used = 1;
    machine2->machine_state.output.arb_gas_used = 2;
    machine3->machine_state.output.arb_gas_used = 3;
    cache.add(std::move(machine0));
    cache.add(std::move(machine1));
    cache.add(std::move(machine2));
    cache.add(std::move(machine3));
    REQUIRE(cache.size() == 4);

    // Test reorg above current height
    cache.reorg(4);
    REQUIRE(cache.size() == 4);

    // Test reorg single value
    cache.reorg(3);
    REQUIRE(cache.size() == 3);

    // Test reorg entire cache
    cache.reorg(0);
    REQUIRE(cache.size() == 0);

    machine4->machine_state.output.arb_gas_used = 40;
    machine5->machine_state.output.arb_gas_used = 42;
    machine6->machine_state.output.arb_gas_used = 43;
    cache.add(std::move(machine4));
    cache.add(std::move(machine5));
    cache.add(std::move(machine6));
    REQUIRE(cache.size() == 3);

    // Test reorg to single value
    cache.reorg(41);
    REQUIRE(cache.size() == 1);

    // Test reorg below current stack
    cache.reorg(0);
    REQUIRE(cache.size() == 0);

    machine7->machine_state.output.arb_gas_used = 39;
    machine8->machine_state.output.arb_gas_used = 40;
    cache.add(std::move(machine7));
    cache.add(std::move(machine8));
    REQUIRE(cache.size() == 2);

    // Older values are fine
    machine9->machine_state.output.arb_gas_used = 30;
    cache.add(std::move(machine9));
    REQUIRE(cache.size() == 3);
}
