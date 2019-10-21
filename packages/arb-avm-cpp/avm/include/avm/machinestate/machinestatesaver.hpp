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
#include "avm/machinestate/checkpointutils.hpp"
#include "avm/value/tuple.hpp"
#include "avm/value/value.hpp"
#include "rocksdb/db.h"

template <typename T>
struct DbResult {
    rocksdb::Status status;
    int reference_count;
    T data;
};

class MachineStateSaver {
   private:
    // when to use shared pointer for an object?
    CheckpointStorage* checkpoint_storage;
    TuplePool* pool;
    std::vector<CodePoint> code;

   public:
    MachineStateSaver(CheckpointStorage* checkpoint_storage,
                      TuplePool* pool,
                      std::vector<CodePoint> code);
    DbResult<CodePoint> getCodePoint(
        const std::vector<unsigned char>& hash_key);
    DbResult<uint256_t> getInt256(const std::vector<unsigned char>& hash_key);
    DbResult<value> getValue(const std::vector<unsigned char>& hash_key);
    DbResult<Tuple> getTuple(const std::vector<unsigned char>& hash_key);
    DbResult<ParsedState> getMachineState(
        const std::vector<unsigned char>& checkpoint_name);
    SaveResults saveTuple(const Tuple& val);
    SaveResults saveValue(const value& val);
    SaveResults saveMachineState(
        ParsedState state_data,
        const std::vector<unsigned char>& checkpoint_name);
};

#endif /* machinestatesaver_hpp */
