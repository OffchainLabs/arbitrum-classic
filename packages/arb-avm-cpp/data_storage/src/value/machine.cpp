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
}  // namespace

void serializeMachineStateKeys(const MachineStateKeys& state_data,
                               std::vector<unsigned char>& state_data_vector) {
    state_data_vector.push_back(static_cast<unsigned char>(state_data.status));
    marshal_uint256_t(state_data.static_hash, state_data_vector);
    marshal_uint256_t(state_data.register_hash, state_data_vector);
    marshal_uint256_t(state_data.datastack_hash, state_data_vector);
    marshal_uint256_t(state_data.auxstack_hash, state_data_vector);
    marshal_uint256_t(state_data.arb_gas_remaining, state_data_vector);
    state_data.pc.marshal(state_data_vector);
    state_data.err_pc.marshal(state_data_vector);
    marshal_uint256_t(state_data.staged_message_hash, state_data_vector);
}

MachineStateKeys extractMachineStateKeys(
    std::vector<unsigned char>::const_iterator& iter) {
    auto status = static_cast<Status>(*iter);
    ++iter;
    auto static_hash = extractUint256(iter);
    auto register_hash = extractUint256(iter);
    auto datastack_hash = extractUint256(iter);
    auto auxstack_hash = extractUint256(iter);
    auto arb_gas_remaining = extractUint256(iter);
    auto pc = extractCodePointRef(iter);
    auto err_pc = extractCodePointStub(iter);
    auto staged_message_hash = extractUint256(iter);

    return MachineStateKeys{static_hash,   register_hash,       datastack_hash,
                            auxstack_hash, arb_gas_remaining,   pc,
                            err_pc,        staged_message_hash, status};
}

void deleteMachineState(Transaction& transaction,
                        MachineStateKeys& parsed_state) {
    std::map<uint64_t, uint64_t> segment_counts;
    auto delete_static_res =
        deleteValueImpl(transaction, parsed_state.static_hash, segment_counts);
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
        std::cout << "error deleting staged message in checkpoint" << std::endl;
    }
}

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
        std::vector<unsigned char>::const_iterator iter =
            results.stored_value.begin();
        auto parsed_state = extractMachineStateKeys(iter);

        deleteMachineState(transaction, parsed_state);
    }
    return delete_results;
}

DbResult<MachineStateKeys> getMachineStateKeys(const Transaction& transaction,
                                               uint256_t machineHash) {
    std::vector<unsigned char> checkpoint_name;
    marshal_uint256_t(machineHash, checkpoint_name);
    auto key = vecToSlice(checkpoint_name);
    auto results = getRefCountedData(*transaction.transaction, key);

    if (!results.status.ok()) {
        return DbResult<MachineStateKeys>{
            results.status, results.reference_count, MachineStateKeys()};
    }
    iterator iter = results.stored_value.begin();
    auto parsed_state = extractMachineStateKeys(iter);

    return DbResult<MachineStateKeys>{results.status, results.reference_count,
                                      parsed_state};
}

rocksdb::Status saveMachineState(Transaction& transaction,
                                 const Machine& machine,
                                 MachineStateKeys& machine_state_keys) {
    std::map<uint64_t, uint64_t> segment_counts;

    auto& machinestate = machine.machine_state;
    auto static_val_results =
        saveValueImpl(transaction, machinestate.static_val, segment_counts);
    if (!static_val_results.status.ok()) {
        return static_val_results.status;
    }

    auto register_val_results =
        saveValueImpl(transaction, machinestate.registerVal, segment_counts);
    if (!register_val_results.status.ok()) {
        return register_val_results.status;
    }

    auto datastack_tup = machinestate.stack.getTupleRepresentation();
    auto datastack_results =
        saveValueImpl(transaction, datastack_tup, segment_counts);
    if (!datastack_results.status.ok()) {
        return datastack_results.status;
    }

    auto auxstack_tup = machinestate.auxstack.getTupleRepresentation();
    auto auxstack_results =
        saveValueImpl(transaction, auxstack_tup, segment_counts);
    if (!auxstack_results.status.ok()) {
        return auxstack_results.status;
    }

    auto staged_message_results =
        saveValueImpl(transaction, machinestate.staged_message, segment_counts);
    if (!staged_message_results.status.ok()) {
        return staged_message_results.status;
    }

    ++segment_counts[machinestate.pc.segment];
    ++segment_counts[machinestate.errpc.pc.segment];

    auto code_status =
        saveCode(transaction, *machinestate.code, segment_counts);
    if (!code_status.ok()) {
        return code_status;
    }

    machine_state_keys.static_hash = hash_value(machinestate.static_val);
    machine_state_keys.register_hash = hash_value(machinestate.registerVal);
    machine_state_keys.datastack_hash = hash(datastack_tup);
    machine_state_keys.auxstack_hash = hash(auxstack_tup);
    machine_state_keys.arb_gas_remaining = machinestate.arb_gas_remaining;
    machine_state_keys.pc = machinestate.pc;
    machine_state_keys.err_pc = machinestate.errpc;
    machine_state_keys.staged_message_hash =
        hash_value(machinestate.staged_message);
    machine_state_keys.status = machinestate.state;

    return rocksdb::Status::OK();
}

SaveResults saveMachine(Transaction& transaction, const Machine& machine) {
    std::vector<unsigned char> checkpoint_name;
    marshal_uint256_t(machine.hash(), checkpoint_name);
    auto key = vecToSlice(checkpoint_name);

    auto transactionResult = getRefCountedData(*transaction.transaction, key);
    if (transactionResult.status.ok()) {
        // Already saved so just increment reference count
        return saveRefCountedData(*transaction.transaction, key,
                                  transactionResult.stored_value);
    }

    MachineStateKeys state_keys{};
    auto status = saveMachineState(transaction, machine, state_keys);
    if (!status.ok()) {
        return SaveResults{0, status};
    }
    std::vector<unsigned char> serialized_state;
    serializeMachineStateKeys(state_keys, serialized_state);
    return saveRefCountedData(*transaction.transaction, key, serialized_state);
}
