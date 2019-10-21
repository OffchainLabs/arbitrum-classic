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

#include "avm/machinestate/checkpointdeleter.hpp"

CheckpointDeleter::CheckpointDeleter(CheckpointStorage* storage) {
    checkpoint_storage = storage;
}

DeleteResults CheckpointDeleter::deleteCheckpoint(
    const std::vector<unsigned char>& checkpoint_name) {
    auto results = checkpoint_storage->getValue(checkpoint_name);

    if (results.status.ok()) {
        auto delete_results = checkpoint_storage->deleteValue(checkpoint_name);

        if (delete_results.reference_count < 1) {
            auto parsed_state =
                Checkpoint::Utils::parseState(results.stored_value);

            auto delete_static_res = deleteValue(parsed_state.static_val_key);
            auto delete_register_res =
                deleteValue(parsed_state.register_val_key);
            auto delete_cp_key = deleteValue(parsed_state.pc_key);
            auto delete_datastack_res = deleteTuple(parsed_state.datastack_key);
            auto delete_auxstack_res = deleteTuple(parsed_state.auxstack_key);
            auto delete_inbox_res = deleteTuple(parsed_state.inbox_key);
            auto delete_inbox_count = deleteValue(parsed_state.inbox_count_key);
            auto delete_pendinginbox_res =
                deleteTuple(parsed_state.pending_key);
            auto delete_pending_count =
                deleteValue(parsed_state.pending_count_key);

            if (delete_static_res.status.ok() &&
                delete_register_res.status.ok() && delete_cp_key.status.ok() &&
                delete_datastack_res.status.ok() &&
                delete_auxstack_res.status.ok() &&
                delete_inbox_res.status.ok() &&
                delete_pendinginbox_res.status.ok() &&
                delete_inbox_count.status.ok() &&
                delete_pending_count.status.ok()) {
            }
        }

        return delete_results;
    } else {
        return DeleteResults{0, results.status};
    }
}

DeleteResults CheckpointDeleter::deleteValue(
    const std::vector<unsigned char>& hash_key) {
    auto results = checkpoint_storage->getValue(hash_key);

    if (results.status.ok()) {
        auto type = (valueTypes)results.stored_value[0];

        if (type == TUPLE_TYPE) {
            return deleteTuple(hash_key, results);
        } else {
            return checkpoint_storage->deleteValue(hash_key);
        }
    } else {
        return DeleteResults{0, results.status};
    }
}

DeleteResults CheckpointDeleter::deleteTuple(
    const std::vector<unsigned char>& hash_key) {
    auto results = checkpoint_storage->getValue(hash_key);
    return deleteTuple(hash_key, results);
}

DeleteResults CheckpointDeleter::deleteTuple(
    const std::vector<unsigned char>& hash_key,
    GetResults& results) {
    if (results.status.ok()) {
        if (results.reference_count == 1) {
            auto value_vectors =
                Checkpoint::Utils::parseSerializedTuple(results.stored_value);

            for (auto& vector : value_vectors) {
                if ((valueTypes)vector[0] == TUPLE_TYPE) {
                    vector.erase(vector.begin());
                    auto delete_stat = deleteTuple(vector);
                }
            }
        }
        return checkpoint_storage->deleteValue(hash_key);
    } else {
        return DeleteResults{0, results.status};
    }
}
