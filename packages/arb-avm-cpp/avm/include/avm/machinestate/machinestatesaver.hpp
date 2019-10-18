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
#include "avm/machinestate/statesaverutils.hpp"
#include "avm/value/tuple.hpp"
#include "avm/value/value.hpp"
#include "rocksdb/db.h"

struct ValueResult {
    rocksdb::Status status;
    int reference_count;
    value val;
};

struct TupleResult {
    rocksdb::Status status;
    int reference_count;
    Tuple tuple;
};

struct NumResult {
    rocksdb::Status status;
    int reference_count;
    uint256_t num;
};

struct CodepointResult {
    rocksdb::Status status;
    int reference_count;
    CodePoint tuple;
};

struct StateResult {
    rocksdb::Status status;
    int reference_count;
    ParsedState state_data;
};

class MachineStateSaver {
   private:
    // unique pointer
    CheckpointStorage* checkpoint_storage;
    TuplePool* pool;
    std::vector<CodePoint> code;
    DeleteResults deleteTuple(std::vector<unsigned char> hash_key,
                              GetResults& results);
    DeleteResults deleteTuple(std::vector<unsigned char> hash_key);
    DeleteResults deleteValue(std::vector<unsigned char> hash_key);

   public:
    MachineStateSaver(CheckpointStorage* checkpoint_storage,
                      TuplePool* pool,
                      std::vector<CodePoint> code);
    CodepointResult getCodePoint(std::vector<unsigned char> hash_key);
    NumResult getInt256(std::vector<unsigned char> hash_key);
    SaveResults saveTuple(const Tuple& val);
    SaveResults saveValue(const value& val);
    ValueResult getValue(std::vector<unsigned char> hash_key);
    TupleResult getTuple(std::vector<unsigned char> hash_key);
    StateResult getMachineStateData(std::string checkpoint_name);
    DeleteResults deleteCheckpoint(std::string checkpoint_name);
    SaveResults saveMachineState(ParsedState state_data,
                                 std::string checkpoint_name);
};

#endif /* machinestatesaver_hpp */
