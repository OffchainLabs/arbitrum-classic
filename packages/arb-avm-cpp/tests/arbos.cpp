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

#include <avm/machine.hpp>
#include <data_storage/checkpointstorage.hpp>

#include <avm_values/vmValueParser.hpp>

#include <catch2/catch.hpp>
#include <nlohmann/json.hpp>

#include <fstream>

TEST_CASE("ARBOS test vectors") {
    DBDeleter deleter;
    TuplePool pool;

    std::vector<std::string> files = {
        "evm_direct_deploy_add", "evm_direct_deploy_and_call_add",
        "evm_load_add_and_verify", "evm_load_fib_and_verify"
        //        "evm_test_arbsys",
        //        "evm_xcontract_call_and_verify",
        //        "evm_xcontract_call_with_constructors"
    };

    for (const auto& filename : files) {
        DYNAMIC_SECTION(filename) {
            auto test_file = std::string{arb_os_test_cases_path} + "/" +
                             filename + ".aoslog";

            std::ifstream i(test_file);
            nlohmann::json j;
            i >> j;

            auto inbox =
                simple_value_from_json(j.at("inbox"), pool).get<Tuple>();
            auto logs_json = j.at("logs");
            std::vector<value> logs;
            for (const auto& log_json : logs_json) {
                logs.push_back(simple_value_from_json(log_json, pool));
            }

            CheckpointStorage storage(dbpath, arb_os_path);
            auto mach = storage.getInitialMachine();
            mach.machine_state.stack.push(uint256_t{0});
            auto assertion = mach.run(1000000000, TimeBounds{}, inbox,
                                      std::chrono::seconds{0});
            std::cout << "test: Machine ran for " << assertion.stepCount
                      << " steps\n";
            INFO("Machine ran for " << assertion.stepCount << " steps");
            REQUIRE(assertion.logs.size() == logs.size());
            auto log = logs[0].get<Tuple>();
            for (size_t i = 0; i < assertion.logs.size(); ++i) {
                REQUIRE(assertion.logs[i] == logs[i]);
            }
        }
    }
}
