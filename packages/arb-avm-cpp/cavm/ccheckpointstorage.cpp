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

    try {
        auto storage =
            new CheckpointStorage(string_filename, string_contract_path);

        if (storage->getInitialVmValues().valid_state) {
            return static_cast<void*>(storage);
        } else {
            return nullptr;
        }
    } catch (const std::exception& exp) {
        return nullptr;
    }
}

int closeCheckpointStorage(CCheckpointStorage* storage_ptr) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    return storage->closeCheckpointStorage();
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
        MachineState machine_state(state.code, state.staticVal, storage->pool);
        auto machine = new Machine();
        machine->initializeMachine(machine_state);

        return static_cast<void*>(machine);
    } else {
        return nullptr;
    }
}

int deleteCheckpoint(CCheckpointStorage* storage_ptr,
                     const void* machine_hash) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);

    auto machine_hash_ptr = reinterpret_cast<const char*>(machine_hash);
    auto hash = deserializeUint256t(machine_hash_ptr);

    std::vector<unsigned char> hash_vector;
    marshal_value(hash, hash_vector);

    auto result = deleteCheckpoint(*storage, hash_vector);

    return result.status.ok();
}

int saveValue(CCheckpointStorage* storage_ptr, const void* value_data) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto valueSaver = MachineStateSaver(storage->makeTransaction());

    auto data_ptr = reinterpret_cast<const char*>(value_data);

    TuplePool pool;
    auto val = deserialize_value(data_ptr, pool);
    auto results = valueSaver.saveValue(val);

    if (!results.status.ok()) {
        return false;
    }

    auto status = valueSaver.commitTransaction();
    return status.ok();
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

    if (!results.status.ok()) {
        return false;
    }

    auto status = deleter.commitTransaction();
    return status.ok();
}

int saveData(CCheckpointStorage* storage_ptr,
             const void* key,
             int key_length,
             const void* data,
             int data_length) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto keyvalue_store = storage->makeKeyValueStore();

    auto key_ptr = reinterpret_cast<const char*>(key);
    auto data_ptr = reinterpret_cast<const char*>(data);

    auto key_vector = std::vector<unsigned char>(key_ptr, key_ptr + key_length);
    auto data_vector =
        std::vector<unsigned char>(data_ptr, data_ptr + data_length);

    auto status = keyvalue_store->saveData(key_vector, data_vector);
    return status.ok();
}

ByteSlice getData(CCheckpointStorage* storage_ptr,
                  const void* key,
                  int key_length) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto keyvalue_store = storage->makeKeyValueStore();

    auto key_ptr = reinterpret_cast<const char*>(key);
    auto key_vector = std::vector<unsigned char>(key_ptr, key_ptr + key_length);

    auto results = keyvalue_store->getData(key_vector);

    auto value_data = (unsigned char*)malloc(results.data.size());
    std::copy(results.data.begin(), results.data.end(), value_data);

    auto void_data = reinterpret_cast<void*>(value_data);
    return {void_data, static_cast<int>(results.data.size())};
}

int deleteData(CCheckpointStorage* storage_ptr,
               const void* key,
               int key_length) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto keyvalue_store = storage->makeKeyValueStore();

    auto key_ptr = reinterpret_cast<const char*>(key);
    auto key_vector = std::vector<unsigned char>(key_ptr, key_ptr + key_length);

    auto status = keyvalue_store->deleteData(key_vector);
    return status.ok();
}
