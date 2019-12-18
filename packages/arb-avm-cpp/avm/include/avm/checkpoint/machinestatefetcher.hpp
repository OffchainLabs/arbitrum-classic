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

#ifndef machinestatefetcher_hpp
#define machinestatefetcher_hpp

#include <avm_values/value.hpp>
#include <data_storage/checkpointutils.hpp>

class CheckpointStorage;

template <typename T>
struct DbResult;

class MachineStateFetcher {
   private:
    const CheckpointStorage& checkpoint_storage;
    TuplePool* pool;
    std::vector<CodePoint> code;

   public:
    MachineStateFetcher(const CheckpointStorage& checkpoint_storage,
                        TuplePool* pool,
                        const std::vector<CodePoint> code);
    DbResult<CodePoint> getCodePoint(
        const std::vector<unsigned char>& hash_key) const;
    DbResult<uint256_t> getUint256_t(
        const std::vector<unsigned char>& hash_key) const;
    DbResult<value> getValue(const std::vector<unsigned char>& hash_key) const;
    DbResult<Tuple> getTuple(const std::vector<unsigned char>& hash_key) const;
    DbResult<ParsedState> getMachineState(
        const std::vector<unsigned char>& checkpoint_name) const;
};

#endif /* machinestatefetcher_hpp */
