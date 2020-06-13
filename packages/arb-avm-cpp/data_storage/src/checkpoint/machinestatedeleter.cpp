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
#include <data_storage/checkpoint/checkpointstorage.hpp>
#include <data_storage/checkpoint/checkpointutils.hpp>
#include <data_storage/checkpoint/machinestatedeleter.hpp>
#include <data_storage/checkpoint/machinestatefetcher.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/transaction.hpp>

namespace {
rocksdb::Slice vecToSlice(const std::vector<unsigned char>& vec) {
    return {reinterpret_cast<const char*>(vec.data()), vec.size()};
}

DeleteResults deleteTuple(Transaction& transaction,
                          const rocksdb::Slice& hash_key);

DeleteResults deleteTuple(Transaction& transaction,
                          const rocksdb::Slice& hash_key,
                          GetResults results) {
    if (!results.status.ok()) {
        return DeleteResults{0, results.status};
    }

    if (results.reference_count == 1) {
        auto value_vectors =
            checkpoint::utils::parseTuple(results.stored_value);

        for (const auto& vec : value_vectors) {
            if (static_cast<ValueTypes>(vec[0]) == TUPLE) {
                rocksdb::Slice tupKey{
                    reinterpret_cast<const char*>(vec.data()) + 1,
                    vec.size() - 1};
                auto delete_status = deleteTuple(transaction, tupKey);
            }
        }
    }
    return transaction.deleteData(hash_key);
}

DeleteResults deleteTuple(Transaction& transaction,
                          const rocksdb::Slice& hash_key) {
    auto results = transaction.getData(hash_key);
    return deleteTuple(transaction, hash_key, results);
}
}  // namespace

DeleteResults deleteValue(Transaction& transaction,
                          const rocksdb::Slice& hash_key) {
    auto results = transaction.getData(hash_key);

    if (!results.status.ok()) {
        return DeleteResults{0, results.status};
    }

    auto type = static_cast<ValueTypes>(results.stored_value[0]);

    if (type == TUPLE) {
        return deleteTuple(transaction, hash_key, results);
    } else {
        return transaction.deleteData(hash_key);
    }
}

DeleteResults deleteValue(Transaction& transaction, uint256_t value_hash) {
    std::vector<unsigned char> hash_key;
    marshal_uint256_t(value_hash, hash_key);
    auto key = vecToSlice(hash_key);
    return deleteValue(transaction, key);
}
