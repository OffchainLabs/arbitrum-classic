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

#include <data_storage/sideloadcache.hpp>

#include <catch2/catch.hpp>

TEST_CASE("SideloadCache add") {
    auto expiration_seconds = 3;
    SideloadCache cache(expiration_seconds);

    // Test that expired block is not added
    cache.add(0, std::time(nullptr) - expiration_seconds, nullptr);
    REQUIRE(cache.size() == 0);

    // Test that non-expired block is added
    cache.add(0, std::time(nullptr), nullptr);
    REQUIRE(cache.size() == 1);
}
TEST_CASE("SideloadCache get") {
    auto expiration_seconds = 3;
    SideloadCache cache(expiration_seconds);

    auto orig_machine = Machine::loadFromFile(
        std::string(machine_test_cases_path) + "/sideloadtest.mexe");

    auto machine1 = std::make_unique<Machine>(orig_machine);
    auto gas1 = machine1->machine_state.arb_gas_remaining;
    auto machine2 = std::make_unique<Machine>(*machine1);
    machine2->machine_state.runOne();
    auto gas2 = machine2->machine_state.arb_gas_remaining;
    REQUIRE(gas1 != gas2);

    cache.add(42, std::time(nullptr), std::move(machine1));
    cache.add(43, std::time(nullptr), std::move(machine2));
    REQUIRE(cache.size() == 2);

    auto machine3 = cache.get(42);
    REQUIRE(machine3);
    REQUIRE(gas1 == machine3->machine_state.arb_gas_remaining);

    // Check nothing returned when invalid height provided
    auto machine4 = cache.get(44);
    REQUIRE(!machine4);
}

TEST_CASE("SideloadCache reorg") {
    auto expiration_seconds = 30;
    SideloadCache cache(expiration_seconds);

    auto orig_machine = Machine::loadFromFile(
        std::string(machine_test_cases_path) + "/sideloadtest.mexe");

    auto machine0 = std::make_unique<Machine>(orig_machine);
    auto gas0 = machine0->machine_state.arb_gas_remaining;
    auto machine1 = std::make_unique<Machine>(*machine0);
    machine1->machine_state.runOne();
    auto gas1 = machine1->machine_state.arb_gas_remaining;
    REQUIRE(gas0 != gas1);
    auto machine2 = std::make_unique<Machine>(*machine1);
    machine2->machine_state.runOne();
    auto gas2 = machine2->machine_state.arb_gas_remaining;
    REQUIRE(gas1 != gas2);
    auto machine3 = std::make_unique<Machine>(*machine2);
    machine3->machine_state.runOne();
    auto gas3 = machine3->machine_state.arb_gas_remaining;
    REQUIRE(gas2 != gas3);

    cache.add(0, std::time(nullptr), std::move(machine0));
    cache.add(1, std::time(nullptr), std::move(machine1));
    cache.add(2, std::time(nullptr), std::move(machine2));
    cache.add(3, std::time(nullptr), std::move(machine3));
    REQUIRE(cache.size() == 4);
    REQUIRE(cache.oldestHeight() == 0);
    REQUIRE(cache.nextHeight() == 4);

    // Test reorg above current height
    cache.reorg(4);
    REQUIRE(cache.size() == 4);
    REQUIRE(cache.oldestHeight() == 0);
    REQUIRE(cache.nextHeight() == 4);

    // Test reorg single value
    cache.reorg(3);
    REQUIRE(cache.size() == 3);
    REQUIRE(cache.oldestHeight() == 0);
    REQUIRE(cache.nextHeight() == 3);

    // Test reorg entire cache
    cache.reorg(0);
    REQUIRE(cache.size() == 0);
    REQUIRE(cache.oldestHeight() == 0);
    REQUIRE(cache.nextHeight() == 0);

    cache.add(40, std::time(nullptr), std::move(machine0));
    cache.add(42, std::time(nullptr), std::move(machine2));
    cache.add(43, std::time(nullptr), std::move(machine3));
    REQUIRE(cache.size() == 3);
    REQUIRE(cache.oldestHeight() == 40);
    REQUIRE(cache.nextHeight() == 44);

    // Test reorg to single value
    cache.reorg(41);
    REQUIRE(cache.size() == 1);
    REQUIRE(cache.oldestHeight() == 40);
    REQUIRE(cache.nextHeight() == 41);

    // Test reorg below current stack
    cache.reorg(0);
    REQUIRE(cache.size() == 0);
    REQUIRE(cache.oldestHeight() == 0);
    REQUIRE(cache.nextHeight() == 0);

    // Test adding below current oldestHeight with empty cache
    cache.reorg(0);
    REQUIRE(cache.size() == 0);
    REQUIRE(cache.oldestHeight() == 0);
    REQUIRE(cache.nextHeight() == 0);

    cache.add(40, std::time(nullptr), std::move(machine0));
    cache.add(42, std::time(nullptr), std::move(machine2));

    // Test implicit reorg to value below current oldest
    cache.add(39, std::time(nullptr), std::move(machine3));

    REQUIRE(cache.size() == 1);
    REQUIRE(cache.oldestHeight() == 39);
    REQUIRE(cache.nextHeight() == 40);
}

TEST_CASE("SideloadCache expire") {
    auto expiration_seconds = 2;
    SideloadCache cache(expiration_seconds);

    auto orig_machine = Machine::loadFromFile(
        std::string(machine_test_cases_path) + "/sideloadtest.mexe");

    auto machine0 = std::make_unique<Machine>(orig_machine);
    auto gas0 = machine0->machine_state.arb_gas_remaining;
    auto machine1 = std::make_unique<Machine>(*machine0);
    machine1->machine_state.runOne();
    auto gas1 = machine1->machine_state.arb_gas_remaining;
    REQUIRE(gas0 != gas1);
    auto machine2 = std::make_unique<Machine>(*machine1);
    machine2->machine_state.runOne();
    auto gas2 = machine2->machine_state.arb_gas_remaining;
    REQUIRE(gas1 != gas2);
    auto machine3 = std::make_unique<Machine>(*machine2);
    machine3->machine_state.runOne();
    auto gas3 = machine3->machine_state.arb_gas_remaining;
    REQUIRE(gas2 != gas3);

    cache.add(0, std::time(nullptr), std::move(machine0));
    cache.add(1, std::time(nullptr), std::move(machine1));
    cache.add(2, std::time(nullptr), std::move(machine2));
    REQUIRE(cache.size() == 3);
    REQUIRE(cache.oldestHeight() == 0);
    REQUIRE(cache.nextHeight() == 3);

    // Let cache expire
    sleep(expiration_seconds);

    // Add one more record
    cache.add(3, std::time(nullptr), std::move(machine3));
    REQUIRE(cache.size() == 1);
    REQUIRE(cache.oldestHeight() == 3);
    REQUIRE(cache.nextHeight() == 4);
}