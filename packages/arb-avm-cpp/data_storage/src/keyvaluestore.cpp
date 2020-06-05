/*
 * Copyright 2020 Offchain Labs, Inc.
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

#include <data_storage/datastorage.hpp>
#include <data_storage/keyvaluestore.hpp>

#include <rocksdb/utilities/transaction_db.h>
#include <data_storage/storageresult.hpp>

rocksdb::Status KeyValueStore::saveData(
    const rocksdb::Slice& key,
    const std::vector<unsigned char>& value) {
    std::string value_str(value.begin(), value.end());
    rocksdb::Slice value_slice(reinterpret_cast<const char*>(value.data()),
                               value.size());
    return data_storage->txn_db->Put(rocksdb::WriteOptions(),
                                     data_storage->default_column.get(), key,
                                     value_slice);
}

rocksdb::Status KeyValueStore::deleteData(const rocksdb::Slice& key) {
    return data_storage->txn_db->Delete(
        rocksdb::WriteOptions(), data_storage->default_column.get(), key);
}

DataResults KeyValueStore::getData(const rocksdb::Slice& key) const {
    auto read_options = rocksdb::ReadOptions();
    std::string stored_value;
    auto status = data_storage->txn_db->Get(
        read_options, data_storage->default_column.get(), key, &stored_value);
    auto data =
        std::vector<unsigned char>(stored_value.begin(), stored_value.end());

    return DataResults{status, data};
}
