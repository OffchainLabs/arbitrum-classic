/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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
#include "utils.hpp"

#include <data_storage/blockstore.hpp>
#include <data_storage/checkpoint/checkpointstorage.hpp>
#include <data_storage/checkpoint/machinestatedeleter.hpp>
#include <data_storage/checkpoint/machinestatefetcher.hpp>
#include <data_storage/checkpoint/machinestatesaver.hpp>
#include <data_storage/confirmednodestore.hpp>
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

CBlockStore* createBlockStore(CCheckpointStorage* storage_ptr) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    return storage->getBlockStore().release();
}

CConfirmedNodeStore* createConfirmedNodeStore(CCheckpointStorage* storage_ptr) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    return storage->getConfirmedNodeStore().release();
}

CMachine* getInitialMachine(const CCheckpointStorage* storage_ptr) {
    auto storage = static_cast<const CheckpointStorage*>(storage_ptr);
    auto state = storage->getInitialVmValues();

    if (!state.valid_state) {
        return nullptr;
    }

    MachineState machine_state(state.code, state.staticVal, storage->pool);
    auto machine = new Machine();
    machine->initializeMachine(machine_state);

    return static_cast<void*>(machine);
}

CMachine* getMachine(const CCheckpointStorage* storage_ptr,
                     const void* machine_hash) {
    auto storage = static_cast<const CheckpointStorage*>(storage_ptr);

    auto hash = receiveUint256(machine_hash);
    std::vector<unsigned char> machine_vector;
    marshal_uint256_t(hash, machine_vector);

    auto initial_state = storage->getInitialVmValues();

    if (!initial_state.valid_state) {
        return nullptr;
    }

    MachineState machine_state(initial_state.code, initial_state.staticVal,
                               storage->pool);
    auto machine = new Machine();
    machine->initializeMachine(machine_state);
    machine->restoreCheckpoint(*storage, machine_vector);

    return machine;
}

int deleteCheckpoint(CCheckpointStorage* storage_ptr,
                     const void* machine_hash) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto hash = receiveUint256(machine_hash);

    std::vector<unsigned char> hash_vector;
    marshal_uint256_t(hash, hash_vector);

    auto result = deleteCheckpoint(*storage, hash_vector);

    return result.status.ok();
}

int saveValue(CCheckpointStorage* storage_ptr, const void* value_data) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto transaction = storage->makeTransaction();

    auto data_ptr = reinterpret_cast<const char*>(value_data);

    TuplePool pool;
    auto val = deserialize_value(data_ptr, pool);
    auto results = saveValue(*transaction, val);

    if (!results.status.ok()) {
        return false;
    }
    return transaction->commit().ok();
}

ByteSlice getValue(const CCheckpointStorage* storage_ptr,
                   const void* hash_key) {
    auto storage = static_cast<const CheckpointStorage*>(storage_ptr);
    auto transaction = storage->makeConstTransaction();
    auto hash = receiveUint256(hash_key);

    std::vector<unsigned char> hash_key_vector;
    auto code = storage->getInitialVmValues().code;
    marshal_value(hash, hash_key_vector, code);

    return returnValueResult(
        getValue(*transaction, hash_key_vector, storage->pool.get()), code);
}

int deleteValue(CCheckpointStorage* storage_ptr, const void* hash_key) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto hash = receiveUint256(hash_key);

    std::vector<unsigned char> hash_key_vector;
    marshal_value(hash, hash_key_vector, storage->getInitialVmValues().code);

    auto results = deleteValue(*storage, hash_key_vector);
    return results.status.ok();
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

    auto key_slice = rocksdb::Slice(key_ptr, key_length);
    auto data_vector =
        std::vector<unsigned char>(data_ptr, data_ptr + data_length);

    return keyvalue_store->saveData(key_slice, data_vector).ok();
}

ByteSliceResult getData(CCheckpointStorage* storage_ptr,
                        const void* key,
                        int key_length) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto keyvalue_store = storage->makeKeyValueStore();

    auto key_ptr = reinterpret_cast<const char*>(key);
    auto key_slice = rocksdb::Slice(key_ptr, key_length);

    return returnDataResult(keyvalue_store->getData(key_slice));
}

int deleteData(CCheckpointStorage* storage_ptr,
               const void* key,
               int key_length) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto keyvalue_store = storage->makeKeyValueStore();

    auto key_ptr = reinterpret_cast<const char*>(key);
    auto key_slice = rocksdb::Slice(key_ptr, key_length);

    return keyvalue_store->deleteData(key_slice).ok();
}
