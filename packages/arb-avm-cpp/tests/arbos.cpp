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

            std::vector<Tuple> messages;
            for (auto& json_message : j.at("inbox")) {
                messages.push_back(
                    simple_value_from_json(json_message).get<Tuple>());
            }

            auto logs_json = j.at("logs");
            std::vector<value> logs;
            for (auto& log_json : logs_json) {
                logs.push_back(simple_value_from_json(log_json));
            }

            ArbStorage storage(dbpath);
            REQUIRE(storage.initialize(arb_os_path).ok());
            auto mach = storage.getInitialMachine(value_cache);
            mach->machine_state.stack.push(uint256_t{0});
            auto assertion = mach->run(MachineExecutionConfig());
            INFO("Machine ran for " << assertion.stepCount << " steps");
            REQUIRE(assertion.logs.size() == logs.size());
            auto log = logs[0].get<Tuple>();
            for (size_t k = 0; k < assertion.logs.size(); ++k) {
                REQUIRE(assertion.logs[k] == logs[k]);
            }
            {
                auto tx = storage.makeTransaction();
                saveMachine(*tx, *mach);
                tx->commit();
            }
            auto mach_hash = mach->hash();
            auto mach2 = storage.getMachine(mach_hash, value_cache);
            REQUIRE(mach_hash == mach2->hash());
            storage.closeArbStorage();

            ArbStorage storage2(dbpath);
            auto mach3 = storage2.getMachine(mach_hash, value_cache);
            REQUIRE(mach_hash == mach3->hash());

            {
                auto tx = storage2.makeTransaction();
                deleteMachine(*tx, mach_hash);
                tx->commit();
            }
        }
    }
}
