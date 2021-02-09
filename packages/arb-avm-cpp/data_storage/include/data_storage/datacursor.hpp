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

#ifndef data_storage_datacursor_hpp
#define data_storage_datacursor_hpp

#include <avm/machinestate/status.hpp>
#include <avm/machinethread.hpp>
#include <avm_values/bigint.hpp>
#include <avm_values/codepointstub.hpp>
#include <data_storage/checkpoint.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresultfwd.hpp>
#include <data_storage/value/machine.hpp>
#include <utility>

class DataCursor {
   public:
    typedef enum {
        EMPTY,      // Out: Ready to receive request for data
        REQUESTED,  // In: Data requested
        READY,      // Out: Data is ready to be picked up
        ERROR       // Out: Error getting data
    } status_enum;

   public:
    std::atomic<status_enum> status{EMPTY};

    // Mutex is acquired by core thread when reorg is occurring.
    // Other threads should acquire mutex whenever accessing below data.
    std::mutex reorg_mutex;
    uint256_t pending_total_count;
    uint256_t current_total_count;
    std::vector<value> data;
    std::vector<value> deleted_data;
    std::string error_string;

    // Input value
    uint256_t number_requested;

   public:
    DataCursor() = default;
};

#endif /* data_storage_datacursor_hpp */
