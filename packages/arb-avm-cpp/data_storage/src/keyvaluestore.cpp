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

#include <data_storage/keyvaluestore.hpp>

#include <rocksdb/utilities/transaction_db.h>
#include <data_storage/storageresult.hpp>

KeyValueStore::KeyValueStore(rocksdb::Transaction* transaction_) {
    transaction = std::unique_ptr<rocksdb::Transaction>(transaction_);
}

auto KeyValueStore::saveData(const std::vector<unsigned char>& key,
                             const std::vector<unsigned char>& value)
    -> rocksdb::Status {
    std::string value_str(value.begin(), value.end());
    std::string key_str(key.begin(), key.end());

    auto save_status = transaction->Put(key_str, value_str);

    if (save_status.ok()) {
        return transaction->Commit();
    } else {
        return save_status;
    }
}

auto KeyValueStore::deleteData(const std::vector<unsigned char>& key)
    -> rocksdb::Status {
    std::string key_str(key.begin(), key.end());

    auto delete_status = transaction->Delete(key_str);

    if (delete_status.ok()) {
        return transaction->Commit();
    } else {
        return delete_status;
    }
}

auto KeyValueStore::getData(const std::vector<unsigned char>& key) const
    -> DataResults {
    auto read_options = rocksdb::ReadOptions();
    std::string stored_value;
    std::string key_str(key.begin(), key.end());

    auto status = transaction->Get(read_options, key_str, &stored_value);
    auto data =
        std::vector<unsigned char>(stored_value.begin(), stored_value.end());

    return DataResults{status, data};
}
