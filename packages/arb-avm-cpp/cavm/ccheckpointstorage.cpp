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

#include <data_storage/aggregator.hpp>
#include <data_storage/blockstore.hpp>
#include <data_storage/checkpointstorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>
#include <data_storage/value/value.hpp>

#include <avm/machine.hpp>
#include <avm_values/value.hpp>

#include <iostream>
#include <string>

CCheckpointStorage* createCheckpointStorage(const char* db_path) {
    auto string_filename = std::string(db_path);
    try {
        auto storage = new CheckpointStorage(string_filename);
        return static_cast<void*>(storage);
    } catch (const std::exception&) {
        return nullptr;
    }
}

int initializeCheckpointStorage(CCheckpointStorage* storage_ptr,
                                const char* executable_path) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    try {
        storage->initialize(executable_path);
        return true;
    } catch (const std::exception&) {
        return false;
    }
}

int checkpointStorageInitialized(CCheckpointStorage* storage_ptr) {
    return static_cast<CheckpointStorage*>(storage_ptr)->initialized();
}

int closeCheckpointStorage(CCheckpointStorage* storage_ptr) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    return storage->closeCheckpointStorage();
}

void destroyCheckpointStorage(CCheckpointStorage* storage) {
    if (storage == nullptr) {
        return;
    }
    delete static_cast<CheckpointStorage*>(storage);
}

CBlockStore* createBlockStore(CCheckpointStorage* storage_ptr) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    return storage->getBlockStore().release();
}

CAggregatorStore* createAggregatorStore(CCheckpointStorage* storage_ptr) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    return storage->getAggregatorStore().release();
}

CMachine* getInitialMachine(const CCheckpointStorage* storage_ptr,
                            CValueCache* value_cache_ptr) {
    auto storage = static_cast<const CheckpointStorage*>(storage_ptr);
    auto value_cache = static_cast<ValueCache*>(value_cache_ptr);
    try {
        if (value_cache == nullptr) {
            ValueCache cache;
            return new Machine(storage->getInitialMachine(cache));
        }

        return new Machine(storage->getInitialMachine(*value_cache));
    } catch (const std::exception&) {
        return nullptr;
    }
}

CMachine* getMachine(const CCheckpointStorage* storage_ptr,
                     const void* machine_hash,
                     CValueCache* value_cache_ptr) {
    auto storage = static_cast<const CheckpointStorage*>(storage_ptr);
    auto hash = receiveUint256(machine_hash);
    auto value_cache = static_cast<ValueCache*>(value_cache_ptr);
    try {
        if (value_cache == nullptr) {
            ValueCache cache;
            return new Machine(storage->getMachine(hash, cache));
        }

        return new Machine(storage->getMachine(hash, *value_cache));
    } catch (const std::exception&) {
        return nullptr;
    }
}

int deleteCheckpoint(CCheckpointStorage* storage_ptr,
                     const void* machine_hash) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto hash = receiveUint256(machine_hash);
    auto transaction = storage->makeTransaction();
    auto results = deleteMachine(*transaction, hash);
    if (!results.status.ok()) {
        return false;
    }
    return transaction->commit().ok();
}

int saveValue(CCheckpointStorage* storage_ptr, const void* value_data) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto transaction = storage->makeTransaction();

    auto data_ptr = reinterpret_cast<const char*>(value_data);

    auto val = deserialize_value(data_ptr);
    auto results = saveValue(*transaction, val);

    if (!results.status.ok()) {
        return false;
    }
    return transaction->commit().ok();
}

ByteSlice getValue(const CCheckpointStorage* storage_ptr,
                   const void* hash_key,
                   CValueCache* value_cache_ptr) {
    auto storage = static_cast<const CheckpointStorage*>(storage_ptr);
    auto hash = receiveUint256(hash_key);
    auto value_cache = static_cast<ValueCache*>(value_cache_ptr);

    return returnValueResult(storage->getValue(hash, *value_cache));
}

int deleteValue(CCheckpointStorage* storage_ptr, const void* hash_key) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto hash = receiveUint256(hash_key);

    auto transaction = storage->makeTransaction();
    auto result = deleteValue(*transaction, hash);
    if (!result.status.ok()) {
        transaction->rollback();
        return false;
    }
    return transaction->commit().ok();
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

RawAssertion checkpointExecuteAssertion(CCheckpointStorage* storage_ptr,
                                        uint64_t maxSteps,
                                        void* inbox_messages,
                                        uint64_t message_count,
                                        uint64_t wallLimit) {
    auto storage = static_cast<CheckpointStorage*>(storage_ptr);
    auto messages = getInboxMessages(inbox_messages, message_count);

    try {
        auto assertion = storage->run(maxSteps, std::move(messages),
                                      std::chrono::seconds{wallLimit});

        return makeRawAssertion(assertion);
    } catch (const std::exception& e) {
        std::cerr << "Failed to make assertion, exception:" << e.what()
                  << std::endl;
        return makeEmptyAssertion();
    }
}
