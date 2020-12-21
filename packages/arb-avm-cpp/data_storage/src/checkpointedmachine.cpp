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

#include "data_storage/checkpointedmachine.hpp"

#include "avm_values/opcodes.hpp"
#include "data_storage/storageresult.hpp"
#include "data_storage/value/machine.hpp"

#include <iostream>
#include <utility>

Assertion CheckpointedMachine::run(uint64_t stepCount,
                                   std::vector<Tuple> inbox_messages,
                                   std::chrono::seconds wallLimit) {
    Assertion assertion =
        mach->run(stepCount, std::move(inbox_messages), wallLimit);

    storage.saveAssertion(assertion);

    return assertion;
}

Assertion CheckpointedMachine::runCallServer(uint64_t stepCount,
                                             std::vector<Tuple> inbox_messages,
                                             std::chrono::seconds wallLimit,
                                             value fake_inbox_peek_value) {
    return mach->runCallServer(stepCount, std::move(inbox_messages), wallLimit,
                               std::move(fake_inbox_peek_value));
}

Assertion CheckpointedMachine::runSideloaded(uint64_t stepCount,
                                             std::vector<Tuple> inbox_messages,
                                             std::chrono::seconds wallLimit,
                                             Tuple sideload_value) {
    return mach->runSideloaded(stepCount, std::move(inbox_messages), wallLimit,
                               std::move(sideload_value));
}
