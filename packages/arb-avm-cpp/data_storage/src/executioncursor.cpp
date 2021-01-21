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

#include <data_storage/executioncursor.hpp>

#include "value/referencecount.hpp"
#include "value/utils.hpp"

#include <data_storage/value/machine.hpp>

std::unique_ptr<Machine> ExecutionCursor::TakeMachine() {
    return std::move(machine);
}

bool ExecutionCursor::AdvanceExecutionCursor(uint256_t max_gas,
                                             bool go_over_gas) {
    if (!machine) {
        return false;
    }

    auto assertion = machine->run(max_gas, go_over_gas, messages,
                                  messages_to_skip, min_next_block_height);

    // TODO update checkpoint fields
}
