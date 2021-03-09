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

#ifndef checkpoint_code_hpp
#define checkpoint_code_hpp

#include <data_storage/value/value.hpp>

#include <rocksdb/status.h>

#include <cstdint>
#include <map>
#include <memory>
#include <set>

class Transaction;
class CodeSegment;
class Code;

uint64_t getNextSegmentID(ReadTransaction& tx);

std::shared_ptr<CodeSegment> getCodeSegment(const ReadTransaction& tx,
                                            uint64_t segment_id,
                                            std::set<uint64_t>& segment_ids,
                                            ValueCache& value_cache);
rocksdb::Status saveCode(ReadWriteTransaction& tx,
                         const Code& code,
                         std::map<uint64_t, uint64_t>& segment_counts);
rocksdb::Status deleteCode(ReadWriteTransaction& tx,
                           std::map<uint64_t, uint64_t>& segment_counts);

#endif /* checkpoint_code_hpp */
