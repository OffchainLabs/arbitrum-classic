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

struct ValueResult {
    int reference_count = 0;
    value val;
};

struct TupleResult {
    int reference_count = 0;
    Tuple tuple;
};

struct MachineStateStorageData {
    SaveResults static_val_results;
    SaveResults register_val_results;
    SaveResults datastack_results;
    SaveResults auxstack_results;
    SaveResults inbox_results;
    SaveResults pending_results;
    SaveResults pc_results;
    unsigned char status_char;
    std::vector<unsigned char> blockreason_str;
    std::vector<unsigned char> balancetracker_str;
};

struct ParsedCheckpointState {
    std::vector<unsigned char> static_val_key;
    std::vector<unsigned char> register_val_key;
    std::vector<unsigned char> datastack_key;
    std::vector<unsigned char> auxstack_key;
    std::vector<unsigned char> inbox_key;
    std::vector<unsigned char> pending_key;
    std::vector<unsigned char> pc_key;
    unsigned char status_char;
    std::vector<unsigned char> blockreason_str;
    std::vector<unsigned char> balancetracker_str;
};

struct MachineStateFetchedData {
    value static_val;
    value register_val;
    Tuple datastack_tuple;
    Tuple auxstack_tuple;
    Tuple inbox_tuple;
    Tuple pending_inbox_tuple;
    CodePoint pc_codepoint;
    unsigned char status_char;
    std::vector<unsigned char> blockreason_str;
    std::vector<unsigned char> balancetracker_str;
};

class MachineStateSaver {
   private:
    // smart pointer?
    CheckpointStorage* checkpoint_storage;
    TuplePool* pool;
    std::vector<std::vector<unsigned char>> parseSerializedTuple(
        std::vector<unsigned char> data_vector);
    std::vector<unsigned char> serializeState(
        MachineStateStorageData state_data);
    MachineStateFetchedData deserializeCheckpointState(
        ParsedCheckpointState stored_state);
    // why not just use checkpointstorage directly
    SaveResults SaveStringValue(const std::string value,
                                const std::vector<unsigned char> key);
    GetResults GetStringValue(const std::vector<unsigned char> key);
    CodePoint getCodePoint(std::vector<unsigned char> hash_key);
    uint256_t getInt256(std::vector<unsigned char> hash_key);
    ParsedCheckpointState parseCheckpointState(
        std::vector<unsigned char> stored_state);

   public:
    void setStorage(CheckpointStorage* storage, TuplePool* pool);
    SaveResults SaveTuple(const Tuple& val);
    SaveResults SaveValue(const value& val);
    DeleteResults Delete(Tuple& tuple);

    ValueResult getValue(std::vector<unsigned char> hash_key);
    TupleResult getTuple(std::vector<unsigned char> hash_key);

    SaveResults SaveMachineState(MachineStateStorageData state_data,
                                 std::string checkpoint_name);
    MachineStateFetchedData GetMachineStateData(std::string checkpoint_name);
    DeleteResults DeleteCheckpoint(std::string checkpoint_name);
};

#endif /* machinestatesaver_hpp */
