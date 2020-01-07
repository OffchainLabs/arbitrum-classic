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

MachineStateDeleter::MachineStateDeleter(
    std::unique_ptr<Transaction> transaction_) {
    transaction = std::move(transaction_);
}

DeleteResults MachineStateDeleter::deleteTuple(
    const std::vector<unsigned char>& hash_key,
    GetResults results) {
    if (results.status.ok()) {
        if (results.reference_count == 1) {
            auto value_vectors =
                checkpoint::utils::parseTuple(results.stored_value);

            for (auto& vector : value_vectors) {
                if (static_cast<ValueTypes>(vector[0]) == TUPLE) {
                    vector.erase(vector.begin());
                    auto delete_status = deleteTuple(vector);
                }
            }
        }
        return transaction->deleteData(hash_key);
    } else {
        return DeleteResults{0, results.status};
    }
}

DeleteResults MachineStateDeleter::deleteTuple(
    const std::vector<unsigned char>& hash_key) {
    auto results = transaction->getData(hash_key);
    return deleteTuple(hash_key, results);
}

DeleteResults MachineStateDeleter::deleteValue(
    const std::vector<unsigned char>& hash_key) {
    auto results = transaction->getData(hash_key);

    if (results.status.ok()) {
        auto type = static_cast<ValueTypes>(results.stored_value[0]);

        if (type == TUPLE) {
            return deleteTuple(hash_key, results);
        } else {
            return transaction->deleteData(hash_key);
        }
    } else {
        return DeleteResults{0, results.status};
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

    if (results.status.ok()) {
        auto delete_results = deleter.deleteValue(checkpoint_name);

        if (delete_results.reference_count < 1) {
            auto parsed_state =
                checkpoint::utils::extractStateKeys(results.stored_value);

            auto delete_static_res =
                deleter.deleteValue(parsed_state.static_val_key);
            auto delete_register_res =
                deleter.deleteValue(parsed_state.register_val_key);
            auto delete_cp_key = deleter.deleteValue(parsed_state.pc_key);
            auto delete_err_pc = deleter.deleteValue(parsed_state.err_pc_key);
            auto delete_datastack_res =
                deleter.deleteTuple(parsed_state.datastack_key);
            auto delete_auxstack_res =
                deleter.deleteTuple(parsed_state.auxstack_key);
            auto delete_inbox_res = deleter.deleteTuple(parsed_state.inbox_key);
            auto delete_inbox_count =
                deleter.deleteValue(parsed_state.inbox_count_key);
            auto delete_pendinginbox_res =
                deleter.deleteTuple(parsed_state.pending_key);
            auto delete_pending_count =
                deleter.deleteValue(parsed_state.pending_count_key);

            if (not(delete_static_res.status.ok() &&
                    delete_register_res.status.ok() &&
                    delete_cp_key.status.ok() &&
                    delete_datastack_res.status.ok() &&
                    delete_auxstack_res.status.ok() &&
                    delete_inbox_res.status.ok() &&
                    delete_pendinginbox_res.status.ok() &&
                    delete_inbox_count.status.ok() &&
                    delete_pending_count.status.ok() &&
                    delete_err_pc.status.ok())) {
                std::cout << "error deleting checkpoint" << std::endl;
            }
        }
        auto status = deleter.commitTransaction();
        delete_results.status = status;
        return delete_results;
    } else {
        return DeleteResults{0, results.status};
    }
}
