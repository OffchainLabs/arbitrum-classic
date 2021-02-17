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

#include <data_storage/arbcore.hpp>

#include <catch2/catch.hpp>

TEST_CASE("CheckpointedMachine tests") {
    DBDeleter deleter;
    auto storage = std::make_shared<DataStorage>(dbpath);
    auto arbcore = std::make_unique<ArbCore>(storage);
    auto executable = loadExecutable(test_contract_path);
    arbcore->initialize(executable);

    SECTION("CheckpointedMachine basic") {
        auto tx = arbcore->makeTransaction();
        REQUIRE(arbcore->initialized());
        REQUIRE(arbcore->maxCheckpointGas() == 0);

        arbcore->saveCheckpoint(*tx);
        REQUIRE(!arbcore->isCheckpointsEmpty(*tx));
        REQUIRE(arbcore->maxCheckpointGas() == 0);
    }
}
