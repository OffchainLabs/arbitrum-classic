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

#include "ccheckpointstorage.h"
#include <avm/checkpointstorage.hpp>
#include <avm/machinestate/checkpointdeleter.hpp>
#include <string>

CCheckpointStorage* createCheckpointStorage(const char* filename) {
    auto string_filename = std::string(filename);

    CheckpointStorage storage(string_filename);
    CheckpointStorage* storage_ptr = &storage;
    return static_cast<void*>(storage_ptr);
}

void destroyCheckpointStorage(CCheckpointStorage* storage) {
    if (storage == NULL)
        return;
    delete static_cast<CheckpointStorage*>(storage);
}

int deleteCheckpoint(CCheckpointStorage* storage_ptr,
                     const char* checkpoint_name) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);

    auto deleter = CheckpointDeleter(storage);
    auto name_str = std::string(checkpoint_name);
    auto name_vector =
        std::vector<unsigned char>(name_str.begin(), name_str.end());
    auto result = deleter.deleteCheckpoint(name_vector);

    return result.status.ok();
}
