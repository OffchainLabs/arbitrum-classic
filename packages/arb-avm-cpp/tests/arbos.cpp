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
    ValueCache value_cache{};

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

            std::vector<InboxMessage> messages;
            for (auto& json_message : j.at("inbox")) {
                messages.push_back(InboxMessage::fromTuple(
                    std::get<Tuple>(simple_value_from_json(json_message))));
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
            auto mach = storage.getInitialMachine(value_cache);
            MachineExecutionConfig config;
            config.inbox_messages = messages;
            auto assertion = mach->run(config);
            INFO("Machine ran for " << assertion.stepCount << " steps");
            REQUIRE(assertion.gasCount == total_gas_target);
            REQUIRE(assertion.logs.size() == logs.size());
            for (size_t k = 0; k < assertion.logs.size(); ++k) {
                REQUIRE(assertion.logs[k] == logs[k]);
            }
            REQUIRE(assertion.sends.size() == sends.size());
            for (size_t k = 0; k < assertion.sends.size(); ++k) {
                REQUIRE(assertion.sends[k] == sends[k]);
            }
            {
                auto tx = storage.getReadWriteTransaction();
                saveMachine(*tx, *mach);
                tx->commit();
            }
            auto mach_hash = mach->hash();
            REQUIRE(mach_hash);
            auto mach2 = storage.getMachine(*mach_hash, value_cache);
            auto mach2_hash = mach2->hash();
            REQUIRE(mach2_hash);
            REQUIRE(*mach_hash == *mach2_hash);
            storage.closeArbStorage();

            ArbStorage storage2(dbpath);
            auto mach3 = storage2.getMachine(*mach_hash, value_cache);
            auto mach3_hash = mach3->hash();
            REQUIRE(mach3_hash);
            REQUIRE(*mach_hash == *mach3_hash);

            {
                auto tx = storage2.getReadWriteTransaction();
                deleteMachine(*tx, *mach_hash);
                tx->commit();
            }
        }
    }
}
