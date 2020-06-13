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
#include <data_storage/checkpoint/machinestatedeleter.hpp>

#include <data_storage/checkpoint/checkpointstorage.hpp>
#include <data_storage/checkpoint/checkpointutils.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/transaction.hpp>

namespace {
rocksdb::Slice vecToSlice(const std::vector<unsigned char>& vec) {
    return {reinterpret_cast<const char*>(vec.data()), vec.size()};
}
}  // namespace

class MachineStateDeleter {
   private:
    std::unique_ptr<Transaction> transaction;
    DeleteResults deleteTuple(const rocksdb::Slice& hash_key,
                              GetResults results);

   public:
    MachineStateDeleter(std::unique_ptr<Transaction> transaction_);
    DeleteResults deleteTuple(const rocksdb::Slice& hash_key);
    DeleteResults deleteValue(const rocksdb::Slice& hash_key);
    rocksdb::Status commitTransaction();
    rocksdb::Status rollBackTransaction();
};

MachineStateDeleter::MachineStateDeleter(
    std::unique_ptr<Transaction> transaction_) {
    transaction = std::move(transaction_);
}

DeleteResults MachineStateDeleter::deleteTuple(const rocksdb::Slice& hash_key,
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
                auto delete_status = deleteTuple(tupKey);
            }
        }
    }
    return transaction->deleteData(hash_key);
}

DeleteResults MachineStateDeleter::deleteTuple(const rocksdb::Slice& hash_key) {
    auto results = transaction->getData(hash_key);
    return deleteTuple(hash_key, results);
}

DeleteResults MachineStateDeleter::deleteValue(const rocksdb::Slice& hash_key) {
    auto results = transaction->getData(hash_key);

    if (!results.status.ok()) {
        return DeleteResults{0, results.status};
    }

    auto type = static_cast<ValueTypes>(results.stored_value[0]);

    if (type == TUPLE) {
        return deleteTuple(hash_key, results);
    } else {
        return transaction->deleteData(hash_key);
    }
}

rocksdb::Status MachineStateDeleter::commitTransaction() {
    return transaction->commit();
}

rocksdb::Status MachineStateDeleter::rollBackTransaction() {
    return transaction->rollBack();
}

DeleteResults deleteCheckpoint(
    CheckpointStorage& checkpoint_storage,
    const std::vector<unsigned char>& checkpoint_name) {
    auto results = checkpoint_storage.getValue(checkpoint_name);
    auto deleter = MachineStateDeleter(checkpoint_storage.makeTransaction());

    if (!results.status.ok()) {
        return DeleteResults{0, results.status};
    }

    auto key = vecToSlice(checkpoint_name);
    auto delete_results = deleter.deleteValue(key);

    if (delete_results.reference_count < 1) {
        auto parsed_state =
            checkpoint::utils::extractStateKeys(results.stored_value);

        auto register_key = vecToSlice(parsed_state.register_val_key);
        auto delete_register_res = deleter.deleteValue(register_key);
        auto datastack_key = vecToSlice(parsed_state.datastack_key);
        auto delete_datastack_res = deleter.deleteTuple(datastack_key);
        auto auxstack_key = vecToSlice(parsed_state.auxstack_key);
        auto delete_auxstack_res = deleter.deleteTuple(auxstack_key);

        if (!(delete_register_res.status.ok() &&
              delete_datastack_res.status.ok() &&
              delete_auxstack_res.status.ok())) {
            std::cout << "error deleting checkpoint" << std::endl;
        }
    }
    auto status = deleter.commitTransaction();
    delete_results.status = status;
    return delete_results;
}

DeleteResults deleteValue(CheckpointStorage& checkpoint_storage,
                          const std::vector<unsigned char>& hash_key) {
    auto deleter = MachineStateDeleter(checkpoint_storage.makeTransaction());
    auto key = vecToSlice(hash_key);
    return deleter.deleteValue(key);
}
