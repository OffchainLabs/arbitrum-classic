/*
 * Copyright 2020, Offchain Labs, Inc.
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

#include "config.hpp"
#include "helper.hpp"

#include <data_storage/arbstorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>

#include <avm_values/vmValueParser.hpp>

#include <catch2/catch.hpp>
#include <nlohmann/json.hpp>

TEST_CASE("ARBOS test vectors") {
    DBDeleter deleter;
    ValueCache value_cache{1, 0};

    std::vector<std::string> files = {
        "evm_direct_deploy_add", "evm_direct_deploy_and_call_add",
        "evm_test_arbsys", "evm_xcontract_call_with_constructors"};

    for (const auto& filename : files) {
        DYNAMIC_SECTION(filename) {
            auto test_file = std::string{arb_os_test_cases_path} + "/" +
                             filename + ".aoslog";

            std::ifstream i(test_file);
            nlohmann::json j;
            i >> j;

            std::vector<MachineMessage> messages;
            for (auto& json_message : j.at("inbox")) {
                messages.emplace_back(
                    InboxMessage::fromTuple(
                        std::get<Tuple>(simple_value_from_json(json_message))),
                    0);
            }

            auto logs_json = j.at("logs");
            std::vector<value> logs;
            for (auto& log_json : logs_json) {
                logs.push_back(simple_value_from_json(log_json));
            }

            auto sends_json = j.at("sends");
            std::vector<std::vector<uint8_t>> sends;
            for (auto& send_json : sends_json) {
                sends.push_back(send_from_json(send_json));
            }
            auto total_gas_target = j.at("total_gas").get<uint64_t>();

            ArbStorage storage(dbpath);
            REQUIRE(storage.initialize(arb_os_path).ok());
            auto mach = storage.getInitialMachine();
            MachineExecutionConfig config;
            config.inbox_messages = messages;
            mach->machine_state.context = AssertionContext(config);
            auto assertion = mach->run();
            INFO("Machine ran for " << assertion.gasCount << " gas with target "
                                    << total_gas_target);
            REQUIRE(assertion.logs.size() == logs.size());
            uint64_t block_log_count = 0;
            uint64_t tx_log_count = 0;
            for (size_t k = 0; k < assertion.logs.size(); ++k) {
                auto typecode = std::get<Tuple>(logs[k]).get_element(0);
                if (typecode == value{uint256_t{0}}) {
                    tx_log_count++;
                } else if (typecode == value{uint256_t{1}}) {
                    block_log_count++;
                }
            }
            INFO("Machine had " << tx_log_count << " tx logs and "
                                << block_log_count << " block logs");
            for (size_t k = 0; k < assertion.logs.size(); ++k) {
                INFO("Checking log " << k);
                CHECK(assertion.logs[k] == logs[k]);
                if (std::get<Tuple>(logs[k]).get_element(0) ==
                    value{uint256_t{1}}) {
                    block_log_count++;
                }
            }

            REQUIRE(assertion.sends.size() == sends.size());
            for (size_t k = 0; k < assertion.sends.size(); ++k) {
                INFO("Checking send " << k);
                CHECK(assertion.sends[k] == sends[k]);
            }
            CHECK(assertion.gasCount == total_gas_target);
            {
                auto tx = storage.makeReadWriteTransaction();
                saveMachine(*tx, *mach);
                tx->commit();
            }

            auto mach2 = storage.getMachine(mach->hash(), value_cache);
            REQUIRE(mach->hash() == mach2->hash());
            storage.closeArbStorage();

            ArbStorage storage2(dbpath);
            auto mach3 = storage2.getMachine(mach->hash(), value_cache);
            REQUIRE(mach->hash() == mach3->hash());

            {
                auto tx = storage2.makeReadWriteTransaction();
                deleteMachine(*tx, mach->hash());
                tx->commit();
            }
        }
    }
}
