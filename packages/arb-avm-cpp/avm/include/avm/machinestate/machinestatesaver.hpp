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

#ifndef machinestatesaver_hpp
#define machinestatesaver_hpp

#include "avm/checkpointstorage.hpp"
#include "value/tuple.hpp"
#include "value/value.hpp"

struct MachineStateStorageData {
    GetResults static_val_results;
    GetResults register_val_results;
    GetResults datastack_results;
    GetResults auxstack_results;
    GetResults inbox_results;
    GetResults pending_results;
    GetResults pc_results;
    unsigned char status_str;
    std::vector<unsigned char> blockreason_str;
    std::vector<unsigned char> balancetracker_str;
};

struct MachineStateFetchedData {
    value static_val_results;
    value register_val_results;
    Tuple datastack_results;
    Tuple auxstack_results;
    Tuple inbox_results;
    Tuple pending_results;
    CodePoint pc_results;
    unsigned char status_str;
    std::vector<unsigned char> blockreason_str;
    std::vector<unsigned char> balancetracker_str;
};

class MachineStateSaver {
   private:
    CheckpointStorage* checkpoint_storage;
    TuplePool* pool;
    std::vector<std::vector<unsigned char>> breakIntoValues(
        std::vector<unsigned char> data_vecgtor);
    std::vector<unsigned char> serializeState(
        MachineStateStorageData state_data);
    MachineStateFetchedData deserializeState(
        std::vector<unsigned char> stored_state);
    GetResults SaveStringValue(const std::string value,
                               const std::vector<unsigned char> key);
    GetResults GetStringValue(const std::vector<unsigned char> key);
    CodePoint getCodePoint(std::vector<unsigned char> hash_key);
    uint256_t getInt256(std::vector<unsigned char> hash_key);

   public:
    void setStorage(CheckpointStorage* storage, TuplePool* pool);
    GetResults SaveTuple(const Tuple& val);
    GetResults SaveValue(const value& val);
    value getValue(std::vector<unsigned char> hash_key);
    Tuple getTuple(std::vector<unsigned char> hash_key);

    GetResults SaveMachineState(MachineStateStorageData state_data,
                                std::string checkpoint_name);

    MachineStateFetchedData GetMachineStateData(std::string checkpoint_name);
};

#endif /* machinestatesaver_hpp */
