/*
 * Copyright 2019, Offchain Labs, Inc.
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

#include <avm/machinestate/messagestack.hpp>
#include <avm/machinestate/tokenTracker.hpp>

#include <data_storage/checkpoint/machinestatefetcher.hpp>
#include <data_storage/checkpoint/machinestatesaver.hpp>

MessageStackSaveResults MessageStack::checkpointState(
    MachineStateSaver& msSaver) {
    auto saved_msgs = msSaver.saveTuple(messages);

    return MessageStackSaveResults{saved_msgs};
}

bool MessageStack::initializeMessageStack(
    const MachineStateFetcher& fetcher,
    const std::vector<unsigned char>& msgs_key) {
    auto msgs_res = fetcher.getTuple(msgs_key);

    if (msgs_res.status.ok()) {
        messages = msgs_res.data;
        return true;
    } else {
        return false;
    }
}
