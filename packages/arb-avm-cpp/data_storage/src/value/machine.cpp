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
#include <data_storage/value/code.hpp>
#include <data_storage/value/value.hpp>

#include <avm/machine.hpp>

#include <iostream>

namespace {
using iterator = std::vector<unsigned char>::const_iterator;

uint256_t extractUint256(iterator& iter) {
    auto ptr = reinterpret_cast<const char*>(&*iter);
    auto int_val = deserializeUint256t(ptr);
    iter += 32;
    return int_val;
}

CodePointRef extractCodePointRef(iterator& iter) {
    auto ptr = reinterpret_cast<const char*>(&*iter);
    auto segment_val = deserialize_uint64_t(ptr);
    auto pc_val = deserialize_uint64_t(ptr);
    iter += sizeof(pc_val) + sizeof(segment_val);
    return {segment_val, pc_val};
}

CodePointStub extractCodePointStub(iterator& iter) {
    auto ref = extractCodePointRef(iter);
    auto next_hash = extractUint256(iter);
    return {ref, next_hash};
}

MachineStateKeys extractStateKeys(
    const std::vector<unsigned char>& stored_state) {
    auto current_iter = stored_state.begin();
    auto status = static_cast<Status>(*current_iter);
    ++current_iter;
    auto static_hash = extractUint256(current_iter);
    auto register_hash = extractUint256(current_iter);
    auto datastack_hash = extractUint256(current_iter);
    auto auxstack_hash = extractUint256(current_iter);
    auto arb_gas_remaining = extractUint256(current_iter);
    auto pc = extractCodePointRef(current_iter);
    auto err_pc = extractCodePointStub(current_iter);
    auto staged_message_hash = extractUint256(current_iter);

    return MachineStateKeys{static_hash,   register_hash,       datastack_hash,
                            auxstack_hash, arb_gas_remaining,   pc,
                            err_pc,        staged_message_hash, status};
}

std::vector<unsigned char> serializeStateKeys(
    const MachineStateKeys& state_data) {
    std::vector<unsigned char> state_data_vector;
    state_data_vector.push_back(static_cast<unsigned char>(state_data.status));
    marshal_uint256_t(state_data.static_hash, state_data_vector);
    marshal_uint256_t(state_data.register_hash, state_data_vector);
    marshal_uint256_t(state_data.datastack_hash, state_data_vector);
    marshal_uint256_t(state_data.auxstack_hash, state_data_vector);
    marshal_uint256_t(state_data.arb_gas_remaining, state_data_vector);
    state_data.pc.marshal(state_data_vector);
    state_data.err_pc.marshal(state_data_vector);
    marshal_uint256_t(state_data.staged_message_hash, state_data_vector);
    return state_data_vector;
}
}  // namespace

DeleteResults deleteMachine(Transaction& transaction, uint256_t machine_hash) {
    std::map<uint64_t, uint64_t> segment_counts;
    std::vector<unsigned char> checkpoint_name;
    marshal_uint256_t(machine_hash, checkpoint_name);
    auto key = vecToSlice(checkpoint_name);
    auto results = getRefCountedData(*transaction.transaction, key);

    if (!results.status.ok()) {
        return DeleteResults{0, results.status,
                             std::move(results.stored_value)};
    }

    auto delete_results = deleteRefCountedData(*transaction.transaction, key);

    if (delete_results.reference_count < 1) {
        auto parsed_state = extractStateKeys(results.stored_value);
        auto delete_static_res = deleteValueImpl(
            transaction, parsed_state.static_hash, segment_counts);
        auto delete_register_res = deleteValueImpl(
            transaction, parsed_state.register_hash, segment_counts);
        auto delete_datastack_res = deleteValueImpl(
            transaction, parsed_state.datastack_hash, segment_counts);
        auto delete_auxstack_res = deleteValueImpl(
            transaction, parsed_state.auxstack_hash, segment_counts);
        auto delete_staged_message_res = deleteValueImpl(
            transaction, parsed_state.staged_message_hash, segment_counts);

        ++segment_counts[parsed_state.pc.segment];
        ++segment_counts[parsed_state.err_pc.pc.segment];

        deleteCode(transaction, segment_counts);

        if (!delete_static_res.status.ok()) {
            std::cout << "error deleting static in checkpoint" << std::endl;
        }

        if (!delete_register_res.status.ok()) {
            std::cout << "error deleting register in checkpoint" << std::endl;
        }

        if (!delete_datastack_res.status.ok()) {
            std::cout << "error deleting datastack in checkpoint" << std::endl;
        }

        if (!delete_auxstack_res.status.ok()) {
            std::cout << "error deleting auxstack in checkpoint" << std::endl;
        }

        if (!delete_staged_message_res.status.ok()) {
            std::cout << "error deleting staged message in checkpoint"
                      << std::endl;
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
    std::map<uint64_t, uint64_t> segment_counts;

    std::vector<unsigned char> checkpoint_name;
    marshal_uint256_t(machine.hash(), checkpoint_name);
    auto key = vecToSlice(checkpoint_name);

    auto currentResult = getRefCountedData(*transaction.transaction, key);
    if (currentResult.status.ok()) {
        // Already saved so just increment reference count
        return saveRefCountedData(*transaction.transaction, key,
                                  currentResult.stored_value);
    }

    auto& machinestate = machine.machine_state;
    auto static_val_results =
        saveValueImpl(transaction, machinestate.static_val, segment_counts);
    auto register_val_results =
        saveValueImpl(transaction, machinestate.registerVal, segment_counts);
    auto datastack_tup = machinestate.stack.getTupleRepresentation();
    auto datastack_results =
        saveValueImpl(transaction, datastack_tup, segment_counts);
    auto auxstack_tup = machinestate.auxstack.getTupleRepresentation();
    auto auxstack_results =
        saveValueImpl(transaction, auxstack_tup, segment_counts);
    auto staged_message_results =
        saveValueImpl(transaction, machinestate.staged_message, segment_counts);
    if (!datastack_results.status.ok() || !auxstack_results.status.ok() ||
        !register_val_results.status.ok() ||
        !staged_message_results.status.ok()) {
        return SaveResults{0, rocksdb::Status::Aborted()};
    }

    ++segment_counts[machinestate.pc.segment];
    ++segment_counts[machinestate.errpc.pc.segment];

    saveCode(transaction, *machinestate.code, segment_counts);

    auto machine_state_data =
        MachineStateKeys{hash_value(machinestate.static_val),
                         hash_value(machinestate.registerVal),
                         hash(datastack_tup),
                         hash(auxstack_tup),
                         machinestate.arb_gas_remaining,
                         machinestate.pc,
                         machinestate.errpc,
                         hash_value(machinestate.staged_message),
                         machinestate.state};
    auto serialized_state = serializeStateKeys(machine_state_data);
    return saveRefCountedData(*transaction.transaction, key, serialized_state);
}
