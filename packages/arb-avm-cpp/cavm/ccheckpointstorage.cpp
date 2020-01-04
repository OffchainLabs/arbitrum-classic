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

#include <data_storage/checkpoint/checkpointstorage.hpp>
#include <data_storage/checkpoint/machinestatedeleter.hpp>
#include <data_storage/checkpoint/machinestatefetcher.hpp>
#include <data_storage/checkpoint/machinestatesaver.hpp>
#include <data_storage/storageresult.hpp>

#include <avm/machine.hpp>
#include <avm_values/value.hpp>

#include <string>

CCheckpointStorage* createCheckpointStorage(const char* db_path,
                                            const char* contract_path) {
    auto string_filename = std::string(db_path);
    auto string_contract_path = std::string(contract_path);

    TuplePool pool;

    auto state = getInitialVmValues(string_contract_path, &pool);

    if (state.valid_state) {
        auto storage = new CheckpointStorage(string_filename, state);
        return static_cast<void*>(storage);
    } else {
        return nullptr;
    }
}

void destroyCheckpointStorage(CCheckpointStorage* storage) {
    if (storage == NULL)
        return;
    delete static_cast<CheckpointStorage*>(storage);
}

CMachine* getInitialMachine(const CCheckpointStorage* storage_ptr) {
    auto storage = static_cast<const CheckpointStorage*>(storage_ptr);
    auto state = storage->getInitialVmValues();

    if (state.valid_state) {
        MachineState machine_state(state.code, state.staticVal);
        auto machine = new Machine();
        machine->initializeMachine(machine_state);

        return static_cast<void*>(machine);
    } else {
        return nullptr;
    }
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

int saveValue(CCheckpointStorage* storage_ptr, const void* value_data) {
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
        return false;
    }
}

ByteSlice getValue(const CCheckpointStorage* storage_ptr,
                   const void* hash_key) {
    auto storage = static_cast<const CheckpointStorage*>(storage_ptr);
    auto fetcher = MachineStateFetcher(*storage);

    auto key_ptr = reinterpret_cast<const char*>(hash_key);
    auto hash = deserializeUint256t(key_ptr);

    std::vector<unsigned char> hash_key_vector;
    marshal_value(hash, hash_key_vector);

    auto results = fetcher.getValue(hash_key_vector);

    std::vector<unsigned char> value;
    marshal_value(results.data, value);

    auto value_data = (unsigned char*)malloc(value.size());
    std::copy(value.begin(), value.end(), value_data);

    auto void_data = reinterpret_cast<void*>(value_data);
    return {void_data, static_cast<int>(value.size())};
}

int deleteValue(CCheckpointStorage* storage_ptr, const void* hash_key) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto deleter = MachineStateDeleter(storage->makeTransaction());

    auto key_ptr = reinterpret_cast<const char*>(hash_key);
    auto hash = deserializeUint256t(key_ptr);

    std::vector<unsigned char> hash_key_vector;
    marshal_value(hash, hash_key_vector);

    auto results = deleter.deleteValue(hash_key_vector);

    if (results.status.ok()) {
        auto status = deleter.commitTransaction();
        return status.ok();
    } else {
        return false;
    }
}

int saveData(CCheckpointStorage* storage_ptr,
             const void* key,
             const void* data) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto transaction = storage->makeTransaction();

    auto key_ptr = reinterpret_cast<const char*>(key);
    auto data_ptr = reinterpret_cast<const char*>(data);

    auto key_str = std::string(key_ptr);
    auto data_str = std::string(data_ptr);

    auto key_vector =
        std::vector<unsigned char>(key_str.begin(), key_str.end());
    auto data_vector =
        std::vector<unsigned char>(data_str.begin(), data_str.end());

    auto results = transaction->saveData(key_vector, data_vector);

    if (results.status.ok()) {
        auto status = transaction->commit();
        return status.ok();
    } else {
        return false;
    }
}

ByteSlice getData(const CCheckpointStorage* storage_ptr, const void* key) {
    auto storage = static_cast<const CheckpointStorage*>(storage_ptr);
    auto transaction = storage->makeConstTransaction();

    auto key_ptr = reinterpret_cast<const char*>(key);
    auto key_str = std::string(key_ptr);
    auto key_vector =
        std::vector<unsigned char>(key_str.begin(), key_str.end());

    auto results = transaction->getData(key_vector);

    auto value_data = (unsigned char*)malloc(results.stored_value.size());
    std::copy(results.stored_value.begin(), results.stored_value.end(),
              value_data);

    auto void_data = reinterpret_cast<void*>(value_data);
    return {void_data, static_cast<int>(results.stored_value.size())};
}

int deleteData(CCheckpointStorage* storage_ptr, const void* key) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto transaction = storage->makeTransaction();

    auto key_ptr = reinterpret_cast<const char*>(key);
    auto key_str = std::string(key_ptr);
    auto key_vector =
        std::vector<unsigned char>(key_str.begin(), key_str.end());

    auto results = transaction->deleteData(key_vector);

    if (results.status.ok()) {
        auto status = transaction->commit();
        return status.ok();
    } else {
        return false;
    }
}
