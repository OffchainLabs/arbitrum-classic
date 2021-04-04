/*
 * Copyright 2021, Offchain Labs, Inc.
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

#ifndef data_storage_readwritetransaction_hpp
#define data_storage_readwritetransaction_hpp

#include <data_storage/readtransaction.hpp>

class ReadWriteTransaction : public ReadConsistentTransaction {
   private:
    rocksdb::WriteOptions write_options{};

   public:
    explicit ReadWriteTransaction(std::shared_ptr<DataStorage> store);

    rocksdb::Status commit() { return transaction->transaction->Commit(); }
    rocksdb::Status rollback() { return transaction->transaction->Rollback(); }

    rocksdb::Status defaultPut(const rocksdb::Slice& key,
                               const rocksdb::Slice& value);
    rocksdb::Status statePut(const rocksdb::Slice& key,
                             const rocksdb::Slice& value);
    rocksdb::Status checkpointPut(const rocksdb::Slice& key,
                                  const rocksdb::Slice& value);
    rocksdb::Status messageEntryPut(const rocksdb::Slice& key,
                                    const rocksdb::Slice& value);
    rocksdb::Status logPut(const rocksdb::Slice& key,
                           const rocksdb::Slice& value);
    rocksdb::Status sendPut(const rocksdb::Slice& key,
                            const rocksdb::Slice& value);
    rocksdb::Status sideloadPut(const rocksdb::Slice& key,
                                const rocksdb::Slice& value);
    rocksdb::Status refCountedPut(const rocksdb::Slice& key,
                                  const rocksdb::Slice& value);

    rocksdb::Status defaultDelete(const rocksdb::Slice& key);
    rocksdb::Status stateDelete(const rocksdb::Slice& key);
    rocksdb::Status checkpointDelete(const rocksdb::Slice& key);
    rocksdb::Status messageEntryDelete(const rocksdb::Slice& key);
    rocksdb::Status logDelete(const rocksdb::Slice& key);
    rocksdb::Status sendDelete(const rocksdb::Slice& key);
    rocksdb::Status sideloadDelete(const rocksdb::Slice& key);
    rocksdb::Status aggregatorPut(const rocksdb::Slice& key,
                                  const rocksdb::Slice& value);
    rocksdb::Status aggregatorDelete(const rocksdb::Slice& key);
    rocksdb::Status refCountedDelete(const rocksdb::Slice& key);
};

#endif  // data_storage_readwritetransaction_hpp
