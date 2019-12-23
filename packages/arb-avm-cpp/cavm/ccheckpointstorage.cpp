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

#include <avm/checkpoint/checkpointstorage.hpp>
#include <avm/checkpoint/machinestatedeleter.hpp>
#include <avm/checkpoint/machinestatefetcher.hpp>
#include <avm/checkpoint/machinestatesaver.hpp>
#include <data_storage/storageresult.hpp>

#include <avm_values/value.hpp>

#include <string>

CCheckpointStorage* createCheckpointStorage(const char* db_path,
                                            const char* contract_path) {
    auto string_filename = std::string(db_path);
    auto string_contract_path = std::string(contract_path);

    auto initial_state = new MachineState(string_contract_path);
    auto storage = new CheckpointStorage(string_filename, *initial_state);
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

    if (results.status.ok()) {
        auto status = valueSaver.commitTransaction();
        return status.ok();
    } else {
        return results.status.ok();
    }
}

ByteSlice getValue(CCheckpointStorage* storage_ptr, void* key) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto fetcher = MachineStateFetcher(*storage);

    auto key_ptr = reinterpret_cast<const char*>(key);
    auto hash = deserializeUint256t(key_ptr);

    std::vector<unsigned char> hash_key_vector;
    marshal_value(hash, hash_key_vector);

    auto results = fetcher.getValue(hash_key_vector);

    // correct marshal?
    std::vector<unsigned char> value;
    marshal_value(results.data, value);

    auto value_data = (unsigned char*)malloc(value.size());
    std::copy(value.begin(), value.end(), value_data);

    auto void_data = reinterpret_cast<void*>(value_data);
    return {void_data, static_cast<int>(value.size())};
}

int deleteValue(CCheckpointStorage* storage_ptr, void* key) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto deleter = MachineStateDeleter(storage->makeTransaction());

    auto key_ptr = reinterpret_cast<const char*>(key);
    auto hash = deserializeUint256t(key_ptr);

    std::vector<unsigned char> hash_key_vector;
    marshal_value(hash, hash_key_vector);

    auto results = deleter.deleteValue(hash_key_vector);

    if (results.status.ok()) {
        auto status = deleter.commitTransaction();
        return status.ok();
    } else {
        return results.status.ok();
    }
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

    if (results.status.ok()) {
        auto status = transaction->commit();
        return status.ok();
    } else {
        return results.status.ok();
    }
}

char* getData(CCheckpointStorage* storage_ptr, const char* key) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto transaction = storage->makeTransaction();

    auto key_str = std::string(key);
    auto key_vector =
        std::vector<unsigned char>(key_str.begin(), key_str.end());

    auto results = transaction->getValue(key_vector);

    char* c_data = (char*)malloc(results.stored_value.size());
    std::copy(results.stored_value.begin(), results.stored_value.end(), c_data);

    return c_data;
}

int deleteData(CCheckpointStorage* storage_ptr, const char* key) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto transaction = storage->makeTransaction();

    auto key_str = std::string(key);
    auto key_vector =
        std::vector<unsigned char>(key_str.begin(), key_str.end());

    auto results = transaction->deleteValue(key_vector);

    if (results.status.ok()) {
        auto status = transaction->commit();
        return status.ok();
    } else {
        return results.status.ok();
    }
}
