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

#include <data_storage/basicmachinecache.hpp>

#include <catch2/catch.hpp>

TEST_CASE("BasicMachineCache add") {
    auto max_size = 2;
    BasicMachineCache cache(max_size);

    // Test empty cache case
    REQUIRE_FALSE(cache.atOrBeforeGas(50).has_value());

    // Test empty findMatching
    auto check_output = [&](const Machine& output) {
        return output.machine_state.output.arb_gas_used <= 42;
    };
    REQUIRE_FALSE(cache.findMatching(check_output).has_value());

    // Test that block is added
    auto machine42 = std::make_unique<Machine>(getComplexMachine());
    machine42->machine_state.output.arb_gas_used = 42;
    cache.add(std::move(machine42));
    REQUIRE(cache.size() == 1);
    auto machine42a = cache.atOrBeforeGas(50);
    REQUIRE(machine42a.has_value());
    REQUIRE(machine42a.value()->first == 42);

    // Test that block is added
    auto machine41 = std::make_unique<Machine>(getComplexMachine());
    machine41->machine_state.output.arb_gas_used = 41;
    cache.add(std::move(machine41));
    REQUIRE(cache.size() == 2);
    auto machine41a = cache.atOrBeforeGas(41);
    REQUIRE(machine41a.has_value());
    REQUIRE(machine41a.value()->first == 41);
    machine42a = cache.atOrBeforeGas(50);
    REQUIRE(machine42a.has_value());
    REQUIRE(machine42a.value()->first == 42);

    // Test that block is added and old block deleted
    auto machine43 = std::make_unique<Machine>(getComplexMachine());
    machine43->machine_state.output.arb_gas_used = 43;
    cache.add(std::move(machine43));
    REQUIRE(cache.size() == 2);
    auto machine43a = cache.atOrBeforeGas(50);
    REQUIRE(machine43a.has_value());
    REQUIRE(machine43a.value()->first == 43);
    machine42a = cache.atOrBeforeGas(42);
    REQUIRE(machine42a.has_value());
    REQUIRE(machine42a.value()->first == 42);

    // Adding block older than oldest when cache full should change nothing
    auto machine30 = std::make_unique<Machine>(getComplexMachine());
    machine30->machine_state.output.arb_gas_used = 30;
    cache.add(std::move(machine30));
    REQUIRE(cache.size() == 2);
    machine43a = cache.atOrBeforeGas(50);
    REQUIRE(machine43a.has_value());
    REQUIRE(machine43a.value()->first == 43);
    machine42a = cache.atOrBeforeGas(42);
    REQUIRE(machine42a.has_value());
    REQUIRE(machine42a.value()->first == 42);

    // Test findMatching
    machine42a = cache.findMatching(check_output);
    REQUIRE(machine42a.has_value());
    REQUIRE(machine42a.value()->first == 42);
}

TEST_CASE("BasicMachineCache reorg") {
    auto expiration_seconds = 30;
    BasicMachineCache cache(expiration_seconds);

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

    // Test adding below current oldestHeight with empty cache
    cache.reorg(0);
    REQUIRE(cache.size() == 0);

    machine7->machine_state.output.arb_gas_used = 39;
    machine8->machine_state.output.arb_gas_used = 40;
    cache.add(std::move(machine7));
    cache.add(std::move(machine8));
    REQUIRE(cache.size() == 2);

    // Older blocks are fine
    machine9->machine_state.output.arb_gas_used = 1;
    cache.add(std::move(machine9));
    REQUIRE(cache.size() == 3);
}
