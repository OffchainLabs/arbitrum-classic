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

#include <data_storage/timedmachinecache.hpp>

#include <catch2/catch.hpp>

TEST_CASE("TimedMachineCache add") {
    auto expiration_seconds = 3;
    auto fake_time = 3600;
    TimedMachineCache cache(expiration_seconds);

    // Test that initial block with any time is added
    auto initial_machine = std::make_unique<Machine>(getComplexMachine());
    initial_machine->machine_state.output.last_inbox_timestamp = fake_time;
    initial_machine->machine_state.output.arb_gas_used = 1;
    cache.add(std::move(initial_machine));
    REQUIRE(cache.size() == 1);
    auto initial_machine2 = std::make_unique<Machine>(getComplexMachine());
    initial_machine2->machine_state.output.last_inbox_timestamp = fake_time + 1;
    initial_machine2->machine_state.output.arb_gas_used = 2;
    cache.add(std::move(initial_machine2));
    REQUIRE(cache.size() == 2);

    // Test that non-expired block is added and expires blocks that are too old
    auto valid_machine = std::make_unique<Machine>(getComplexMachine());
    valid_machine->machine_state.output.last_inbox_timestamp =
        std::time(nullptr);
    valid_machine->machine_state.output.arb_gas_used = 3;
    cache.add(std::move(valid_machine));
    REQUIRE(cache.size() == 1);
    auto valid_machine_b = cache.atOrBeforeGas(4);
    REQUIRE(valid_machine_b.has_value());
    REQUIRE(valid_machine_b.value()
                ->second.machine->machine_state.output.arb_gas_used == 3);

    // Test that expired block is not added
    auto expired_machine = std::make_unique<Machine>(getComplexMachine());
    expired_machine->machine_state.output.last_inbox_timestamp =
        std::time(nullptr) - expiration_seconds;
    expired_machine->machine_state.output.arb_gas_used = 4;
    cache.add(std::move(expired_machine));
    REQUIRE(cache.size() == 1);
    auto expired_machine_b = cache.atOrBeforeGas(4);
    REQUIRE(expired_machine_b.has_value());
    REQUIRE(expired_machine_b.value()
                ->second.machine->machine_state.output.arb_gas_used == 3);
}
TEST_CASE("TimedMachineCache get") {
    auto expiration_seconds = 3;
    TimedMachineCache cache(expiration_seconds);

    auto orig_machine = Machine::loadFromFile(
        std::string(machine_test_cases_path) + "/sideloadtest.mexe");

    auto machine1a = std::make_unique<Machine>(orig_machine);
    auto gas1a = machine1a->machine_state.output.arb_gas_used;

    auto machine2a = std::make_unique<Machine>(*machine1a);
    machine2a->machine_state.runOne();
    auto gas2a = machine2a->machine_state.output.arb_gas_used;
    REQUIRE(gas1a != gas2a);

    machine1a->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    machine2a->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    cache.add(std::move(machine1a));
    cache.add(std::move(machine2a));
    REQUIRE(cache.size() == 2);

    auto machine1b = cache.atOrBeforeGas(gas1a);
    REQUIRE(machine1b.has_value());
    REQUIRE(
        gas1a ==
        machine1b.value()->second.machine->machine_state.output.arb_gas_used);

    auto machine2b = cache.atOrBeforeGas(50);
    REQUIRE(machine2b.has_value());
    REQUIRE(
        gas2a ==
        machine2b.value()->second.machine->machine_state.output.arb_gas_used);
}

TEST_CASE("TimedMachineCache reorg") {
    auto expiration_seconds = 30;
    TimedMachineCache cache(expiration_seconds);

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
    machine0->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    machine1->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    machine2->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    machine3->machine_state.output.last_inbox_timestamp = std::time(nullptr);
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
    machine4->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    machine5->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    machine6->machine_state.output.last_inbox_timestamp = std::time(nullptr);
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
    machine7->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    machine8->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    cache.add(std::move(machine7));
    cache.add(std::move(machine8));
    REQUIRE(cache.size() == 2);

    // Older blocks are fine as long as the timestamp isn't too old
    machine9->machine_state.output.arb_gas_used = 30;
    machine9->machine_state.output.last_inbox_timestamp =
        std::time(nullptr) - 20;
    cache.add(std::move(machine9));
    REQUIRE(cache.size() == 3);
}

TEST_CASE("TimedMachineCache expire") {
    auto expiration_seconds = 2;
    TimedMachineCache cache(expiration_seconds);

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

    machine0->machine_state.output.arb_gas_used = 0;
    machine1->machine_state.output.arb_gas_used = 1;
    machine2->machine_state.output.arb_gas_used = 2;
    machine0->machine_state.output.last_inbox_timestamp =
        std::time(nullptr) - 1;
    machine1->machine_state.output.last_inbox_timestamp =
        std::time(nullptr) - 1;
    machine2->machine_state.output.last_inbox_timestamp =
        std::time(nullptr) - 1;
    cache.add(std::move(machine0));
    cache.add(std::move(machine1));
    cache.add(std::move(machine2));
    REQUIRE(cache.size() == 3);

    // Let cache expire
    sleep(expiration_seconds - 1);

    // Add one more record
    machine3->machine_state.output.arb_gas_used = 3;
    machine3->machine_state.output.last_inbox_timestamp = std::time(nullptr);
    cache.add(std::move(machine3));
    REQUIRE(cache.size() == 1);
}

TEST_CASE("TimedMachineCache expiredTimestamp") {
    auto timed_expire = 20;
    auto fake_time = 1000000;
    TimedMachineCache cache(timed_expire);

    auto expired = cache.expiredTimestamp();
    REQUIRE(expired == 0);

    auto orig_machine = Machine::loadFromFile(
        std::string(machine_test_cases_path) + "/sideloadtest.mexe");

    auto machine0 = std::make_unique<Machine>(orig_machine);
    machine0->machine_state.output.last_inbox_timestamp = fake_time;
    cache.add(std::move(machine0));
    REQUIRE(cache.size() == 1);

    auto expired2 = cache.expiredTimestamp();
    REQUIRE(expired2 == fake_time - timed_expire);
}

TEST_CASE("TimedMachineCache currentTimeExpired") {
    auto timed_expire = 20;
    auto expiration_fudge_factor = 10;
    TimedMachineCache cache(timed_expire);

    auto expired = cache.currentTimeExpired();
    REQUIRE(expired >=
            std::time(nullptr) - timed_expire - expiration_fudge_factor);
    REQUIRE(expired <=
            std::time(nullptr) - timed_expire + expiration_fudge_factor);
}
