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

#ifndef transaction_hpp
#define transaction_hpp

#include <data_storage/readwritetransaction.hpp>
#include <data_storage/storageresultfwd.hpp>

#include <memory>
#include <vector>

namespace rocksdb {
class Transaction;
class Status;
struct Slice;
}  // namespace rocksdb

GetResults getRefCountedData(const ReadTransaction& tx,
                             const rocksdb::Slice& hash_key);
SaveResults saveRefCountedData(ReadWriteTransaction& tx,
                               const rocksdb::Slice& hash_key,
                               const std::vector<unsigned char>& value,
                               uint32_t new_references = 1,
                               bool allow_replacement = false);
SaveResults incrementReference(ReadWriteTransaction& tx,
                               const rocksdb::Slice& hash_key,
                               uint32_t new_references = 1);
DeleteResults deleteRefCountedData(ReadWriteTransaction& transaction,
                                   const rocksdb::Slice& hash_key,
                                   uint32_t deleted_references = 1);

#endif /* transaction_hpp */
