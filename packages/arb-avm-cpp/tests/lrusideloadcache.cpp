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

#include <data_storage/lrusideloadcache.hpp>

#include <catch2/catch.hpp>

TEST_CASE("LRUSideloadCache add") {
    auto cache_size = 1;
    LRUSideloadCache cache(cache_size);

    // Basic add
    auto machine_zero = std::make_unique<Machine>(getComplexMachine());
    machine_zero->machine_state.output.arb_gas_used = 0;
    cache.add(std::move(machine_zero));
    REQUIRE(cache.size() == 1);

    // Test that cache_size limit is not breached
    auto machine_one = std::make_unique<Machine>(getComplexMachine());
    machine_one->machine_state.output.arb_gas_used = 1;
    cache.add(std::move(machine_one));
    REQUIRE(cache.size() == 1);
    auto retrieved_machine = cache.atOrBeforeGas(0);
    REQUIRE(!retrieved_machine.has_value());
    auto retrieved_machine2 = cache.atOrBeforeGas(1);
    REQUIRE(retrieved_machine2.has_value());
    REQUIRE(retrieved_machine2.value()
                ->second.first->machine_state.output.arb_gas_used == 1);
}
TEST_CASE("LRUSideloadCache get") {
    auto cache_size = 3;
    LRUSideloadCache cache(cache_size);

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

TEST_CASE("LRUSideloadCache reorg") {
    auto cache_size = 30;
    LRUSideloadCache cache(cache_size);

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

    // Test implicit reorg to value below current oldest
    machine9->machine_state.output.arb_gas_used = 30;
    cache.add(std::move(machine9));

    REQUIRE(cache.size() == 1);
}
