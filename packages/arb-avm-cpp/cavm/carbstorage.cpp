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

#include "carbstorage.h"
#include "utils.hpp"

#include <data_storage/aggregator.hpp>
#include <data_storage/arbstorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>
#include <data_storage/value/value.hpp>

#include <avm/machine.hpp>
#include <avm_values/value.hpp>

#include <iostream>
#include <string>

CArbStorage* createArbStorage(const char* db_path) {
    auto string_filename = std::string(db_path);
    try {
        auto storage = new ArbStorage(string_filename);
        return static_cast<void*>(storage);
    } catch (const std::exception& e) {
        std::cerr << "Error creating storage " << e.what() << std::endl;
        return nullptr;
    }
}

int initializeArbStorage(CArbStorage* storage_ptr,
                         const char* executable_path) {
    auto storage = static_cast<ArbStorage*>(storage_ptr);
    try {
        auto status = storage->initialize(executable_path);
        if (!status.ok()) {
            std::cerr << "Error initializing storage" << status.ToString()
                      << std::endl;
            return false;
        }

        return true;
    } catch (const std::exception& e) {
        std::cerr << "Exception initializing storage" << e.what() << std::endl;
        return false;
    }
}

int arbStorageInitialized(CArbStorage* storage_ptr) {
    return static_cast<ArbStorage*>(storage_ptr)->initialized();
}

int closeArbStorage(CArbStorage* storage_ptr) {
    auto storage = static_cast<ArbStorage*>(storage_ptr);
    return storage->closeArbStorage();
}

void destroyArbStorage(CArbStorage* storage_ptr) {
    auto storage = static_cast<ArbStorage*>(storage_ptr);
    if (storage == nullptr) {
        return;
    }
    std::cerr << "closing ArbStorage" << std::endl;
    storage->closeArbStorage();
    std::cerr << "closed ArbStorage" << std::endl;
    delete static_cast<ArbStorage*>(storage);
}

CArbCore* createArbCore(CArbStorage* storage_ptr) {
    auto storage = static_cast<ArbStorage*>(storage_ptr);
    return storage->getArbCore().get();
}

CAggregatorStore* createAggregatorStore(CArbStorage* storage_ptr) {
    auto storage = static_cast<ArbStorage*>(storage_ptr);
    return storage->getAggregatorStore().release();
}
