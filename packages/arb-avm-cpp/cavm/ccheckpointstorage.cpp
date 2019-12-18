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

#include <data_storage/checkpointdeleter.hpp>
#include <data_storage/checkpointresult.hpp>
#include <data_storage/checkpointstorage.hpp>
#include <data_storage/machinestatesaver.hpp>

#include <avm_values/value.hpp>

#include <string>

CCheckpointStorage* createCheckpointStorage(const char* filename) {
    auto string_filename = std::string(filename);

    auto storage = new CheckpointStorage(string_filename);
    return static_cast<void*>(storage);
}

void destroyCheckpointStorage(CCheckpointStorage* storage) {
    if (storage == NULL)
        return;
    delete static_cast<CheckpointStorage*>(storage);
}

int deleteCheckpoint(CCheckpointStorage* storage_ptr,
                     const char* checkpoint_name) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);

    auto name_str = std::string(checkpoint_name);
    auto name_vector =
        std::vector<unsigned char>(name_str.begin(), name_str.end());
    auto result = deleteCheckpoint(*storage, name_vector);

    return result.status.ok();
}

int saveValue(CCheckpointStorage* storage_ptr, void* value_data) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto valueSaver = MachineStateSaver(storage->makeTransaction());

    auto data_ptr = reinterpret_cast<const char*>(value_data);

    TuplePool pool;
    auto val = deserialize_value(data_ptr, pool);
    auto results = valueSaver.saveValue(val);

    return results.status.ok();
}

int saveData(CCheckpointStorage* storage_ptr,
             const char* key,
             const char* value) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto transaction = storage->makeTransaction();

    auto key_str = std::string(key);
    auto key_vector =
        std::vector<unsigned char>(key_str.begin(), key_str.end());

    auto value_str = std::string(value);
    auto value_vector =
        std::vector<unsigned char>(value_str.begin(), value_str.end());

    auto results = transaction->saveValue(key_vector, value_vector);
    return results.status.ok();
}
