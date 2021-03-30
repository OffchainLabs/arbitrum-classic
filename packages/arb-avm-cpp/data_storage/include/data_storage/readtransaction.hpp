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

#ifndef data_storage_readtransaction_hpp
#define data_storage_readtransaction_hpp

#include <data_storage/datastorage.hpp>

class ReadSnapshotTransaction;
class ReadWriteTransaction;

class ReadTransaction {
   protected:
    std::unique_ptr<Transaction> transaction{};
    rocksdb::ReadOptions read_options{};

   public:
    ReadTransaction() = delete;
    explicit ReadTransaction(std::shared_ptr<DataStorage> store);

    rocksdb::Status defaultGet(const rocksdb::Slice& key,
                               std::string* value) const;
    rocksdb::Status stateGet(const rocksdb::Slice& key,
                             std::string* value) const;
    rocksdb::Status checkpointGet(const rocksdb::Slice& key,
                                  std::string* value) const;
    rocksdb::Status messageEntryGet(const rocksdb::Slice& key,
                                    std::string* value) const;
    rocksdb::Status logGet(const rocksdb::Slice& key, std::string* value) const;
    rocksdb::Status sendGet(const rocksdb::Slice& key,
                            std::string* value) const;
    rocksdb::Status sideloadGet(const rocksdb::Slice& key,
                                std::string* value) const;
    rocksdb::Status aggregatorGet(const rocksdb::Slice& key,
                                  std::string* value) const;
    rocksdb::Status refCountedGet(const rocksdb::Slice& key,
                                  std::string* value) const;

    [[nodiscard]] std::unique_ptr<rocksdb::Iterator> defaultGetIterator() const;
    [[nodiscard]] std::unique_ptr<rocksdb::Iterator> stateGetIterator() const;
    [[nodiscard]] std::unique_ptr<rocksdb::Iterator> checkpointGetIterator()
        const;
    [[nodiscard]] std::unique_ptr<rocksdb::Iterator> messageEntryGetIterator()
        const;
    [[nodiscard]] std::unique_ptr<rocksdb::Iterator> logGetIterator() const;
    [[nodiscard]] std::unique_ptr<rocksdb::Iterator> sendGetIterator() const;
    [[nodiscard]] std::unique_ptr<rocksdb::Iterator> sideloadGetIterator()
        const;
    [[nodiscard]] std::unique_ptr<rocksdb::Iterator> aggregatorGetIterator()
        const;
    [[nodiscard]] std::unique_ptr<rocksdb::Iterator> refCountedGetIterator()
        const;

    [[nodiscard]] ValueResult<uint256_t> defaultGetUint256(
        rocksdb::Slice key_slice) const;
    [[nodiscard]] ValueResult<uint256_t> stateGetUint256(
        rocksdb::Slice key_slice) const;
    [[nodiscard]] ValueResult<uint256_t> checkpointGetUint256(
        rocksdb::Slice key_slice) const;
    [[nodiscard]] ValueResult<uint256_t> messageEntryGetUint256(
        rocksdb::Slice key_slice) const;
    [[nodiscard]] ValueResult<uint256_t> logGetUint256(
        rocksdb::Slice key_slice) const;
    [[nodiscard]] ValueResult<uint256_t> sendGetUint256(
        rocksdb::Slice key_slice) const;
    [[nodiscard]] ValueResult<uint256_t> sideloadGetUint256(
        rocksdb::Slice key_slice) const;
    [[nodiscard]] ValueResult<std::vector<uint256_t>> logGetUint256Vector(
        rocksdb::Slice first_key_slice,
        size_t count) const;
    [[nodiscard]] ValueResult<std::vector<std::vector<unsigned char>>>
    messageEntryGetVectorVector(rocksdb::Slice first_key_slice,
                                size_t count) const;
    [[nodiscard]] ValueResult<std::vector<std::vector<unsigned char>>>
    sendGetVectorVector(rocksdb::Slice first_key_slice, size_t count) const;
    [[nodiscard]] ValueResult<std::vector<unsigned char>> messageEntryGetVector(
        rocksdb::Slice first_key_slice) const;
    [[nodiscard]] ValueResult<std::vector<unsigned char>> checkpointGetVector(
        rocksdb::Slice first_key_slice) const;
    [[nodiscard]] ValueResult<uint256_t> aggregatorGetUint256(
        rocksdb::Slice key_slice) const;
    [[nodiscard]] ValueResult<uint256_t> refCountedGetUint256(
        rocksdb::Slice key_slice) const;

   private:
    ValueResult<std::vector<std::vector<unsigned char>>>
    getVectorVectorUsingFamilyAndKey(rocksdb::ColumnFamilyHandle* family,
                                     rocksdb::Slice first_key_slice,
                                     size_t count) const;
    ValueResult<std::vector<unsigned char>> getVectorUsingFamilyAndKey(
        rocksdb::ColumnFamilyHandle* family,
        rocksdb::Slice key_slice) const;
    ValueResult<std::vector<uint256_t>> getUint256VectorUsingFamilyAndKey(
        rocksdb::ColumnFamilyHandle* family,
        rocksdb::Slice first_key_slice,
        size_t count) const;
    ValueResult<uint256_t> getUint256UsingFamilyAndKey(
        rocksdb::ColumnFamilyHandle* family,
        rocksdb::Slice key_slice) const;
};

#endif  // data_storage_readtransaction_hpp
