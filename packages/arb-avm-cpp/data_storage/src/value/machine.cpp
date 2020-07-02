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

#include <data_storage/value/machine.hpp>

#include "referencecount.hpp"
#include "utils.hpp"

#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/value.hpp>

#include <avm/machine.hpp>

namespace {
rocksdb::Slice vecToSlice(const std::vector<unsigned char>& vec) {
    return {reinterpret_cast<const char*>(vec.data()), vec.size()};
}

using iterator = std::vector<unsigned char>::const_iterator;

CodePointRef extractCodePointRef(iterator& iter) {
    auto ptr = reinterpret_cast<const char*>(&*iter);
    auto segment_val = checkpoint::utils::deserialize_uint64(ptr);
    auto pc_val = checkpoint::utils::deserialize_uint64(ptr);
    iter += sizeof(pc_val) + sizeof(segment_val);
    return {segment_val, pc_val};
}

uint256_t extractUint256(iterator& iter) {
    auto ptr = reinterpret_cast<const char*>(&*iter);
    auto int_val = deserializeUint256t(ptr);
    iter += 32;
    return int_val;
}

MachineStateKeys extractStateKeys(
    const std::vector<unsigned char>& stored_state) {
    auto current_iter = stored_state.begin();
    auto status = static_cast<Status>(*current_iter);
    ++current_iter;
    auto register_hash = extractUint256(current_iter);
    auto datastack_hash = extractUint256(current_iter);
    auto auxstack_hash = extractUint256(current_iter);
    auto arb_gas_remaining = extractUint256(current_iter);
    auto pc = extractCodePointRef(current_iter);
    auto err_pc = extractCodePointRef(current_iter);

    return MachineStateKeys{
        register_hash, datastack_hash, auxstack_hash, arb_gas_remaining, pc,
        err_pc,        status};
}

std::vector<unsigned char> serializeStateKeys(
    const MachineStateKeys& state_data) {
    std::vector<unsigned char> state_data_vector;
    state_data_vector.push_back(static_cast<unsigned char>(state_data.status));
    marshal_uint256_t(state_data.register_hash, state_data_vector);
    marshal_uint256_t(state_data.datastack_hash, state_data_vector);
    marshal_uint256_t(state_data.auxstack_hash, state_data_vector);
    marshal_uint256_t(state_data.arb_gas_remaining, state_data_vector);
    state_data.pc.marshal(state_data_vector);
    state_data.err_pc.marshal(state_data_vector);
    return state_data_vector;
}
}  // namespace

DeleteResults deleteMachine(Transaction& transaction, uint256_t machine_hash) {
    std::vector<unsigned char> checkpoint_name;
    marshal_uint256_t(machine_hash, checkpoint_name);
    auto key = vecToSlice(checkpoint_name);
    auto results = getRefCountedData(*transaction.transaction, key);

    if (!results.status.ok()) {
        return DeleteResults{0, results.status};
    }

    auto delete_results = deleteRefCountedData(*transaction.transaction, key);

    if (delete_results.reference_count < 1) {
        auto parsed_state = extractStateKeys(results.stored_value);

        auto delete_register_res =
            deleteValue(transaction, parsed_state.register_hash);
        auto delete_datastack_res =
            deleteValue(transaction, parsed_state.datastack_hash);
        auto delete_auxstack_res =
            deleteValue(transaction, parsed_state.auxstack_hash);

        if (!(delete_register_res.status.ok() &&
              delete_datastack_res.status.ok() &&
              delete_auxstack_res.status.ok())) {
            std::cout << "error deleting checkpoint" << std::endl;
        }
    }
    return delete_results;
}

DbResult<MachineStateKeys> getMachineState(const Transaction& transaction,
                                           uint256_t machineHash) {
    std::vector<unsigned char> checkpoint_name;
    marshal_uint256_t(machineHash, checkpoint_name);
    auto key = vecToSlice(checkpoint_name);
    auto results = getRefCountedData(*transaction.transaction, key);

    if (!results.status.ok()) {
        return DbResult<MachineStateKeys>{
            results.status, results.reference_count, MachineStateKeys()};
    }
    auto parsed_state = extractStateKeys(results.stored_value);

    return DbResult<MachineStateKeys>{results.status, results.reference_count,
                                      parsed_state};
}

SaveResults saveMachine(Transaction& transaction, const Machine& machine) {
    std::vector<unsigned char> checkpoint_name;
    marshal_uint256_t(machine.hash(), checkpoint_name);
    auto key = vecToSlice(checkpoint_name);

    auto currentResult = getRefCountedData(*transaction.transaction, key);
    if (currentResult.status.ok()) {
        // Already saved so just increment refence count
        return saveRefCountedData(*transaction.transaction, key,
                                  currentResult.stored_value);
    }

    auto& machinestate = machine.machine_state;
    auto pool = machinestate.pool.get();
    auto register_val_results =
        saveValue(transaction, machinestate.registerVal);
    auto datastack_tup = machinestate.stack.getTupleRepresentation(pool);
    auto datastack_results = saveValue(transaction, datastack_tup);
    auto auxstack_tup = machinestate.auxstack.getTupleRepresentation(pool);
    auto auxstack_results = saveValue(transaction, auxstack_tup);
    if (!datastack_results.status.ok() || !auxstack_results.status.ok() ||
        !register_val_results.status.ok()) {
        return SaveResults{0, rocksdb::Status().Aborted()};
    }
    auto machine_state_data =
        MachineStateKeys{hash_value(machinestate.registerVal),
                         hash(datastack_tup),
                         hash(auxstack_tup),
                         machinestate.arb_gas_remaining,
                         machinestate.pc,
                         machinestate.errpc,
                         machinestate.state};
    auto serialized_state = serializeStateKeys(machine_state_data);
    return saveRefCountedData(*transaction.transaction, key, serialized_state);
}
