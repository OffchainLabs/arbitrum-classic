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

#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/utils.hpp>
#include <data_storage/value/value.hpp>

#include <avm/machine.hpp>

#include <iostream>

void serializeMachineOutput(const MachineOutput& output_data,
                            std::vector<unsigned char>& state_data_vector) {
    marshal_uint256_t(output_data.fully_processed_inbox.count,
                      state_data_vector);
    marshal_uint256_t(output_data.fully_processed_inbox.accumulator,
                      state_data_vector);
    marshal_uint256_t(output_data.total_steps, state_data_vector);
    marshal_uint256_t(output_data.arb_gas_used, state_data_vector);
    marshal_uint256_t(output_data.send_acc, state_data_vector);
    marshal_uint256_t(output_data.log_acc, state_data_vector);
    marshal_uint256_t(output_data.send_count, state_data_vector);
    marshal_uint256_t(output_data.log_count, state_data_vector);
    marshal_uint256_t(output_data.l1_block_number, state_data_vector);
    marshal_uint256_t(output_data.l2_block_number, state_data_vector);
    marshal_uint256_t(output_data.last_inbox_timestamp, state_data_vector);

    auto last_sideload_raw = std::numeric_limits<uint256_t>::max();
    if (output_data.last_sideload.has_value()) {
        last_sideload_raw = *output_data.last_sideload;
    }
    marshal_uint256_t(last_sideload_raw, state_data_vector);
}

void serializeMachineStateKeys(const MachineStateKeys& state_data,
                               std::vector<unsigned char>& state_data_vector) {
    serializeMachineOutput(state_data.output, state_data_vector);

    marshal_uint256_t(state_data.pc_hash, state_data_vector);
    marshal_uint256_t(state_data.static_hash, state_data_vector);
    marshal_uint256_t(state_data.register_hash, state_data_vector);
    marshal_uint256_t(state_data.datastack_hash, state_data_vector);
    marshal_uint256_t(state_data.auxstack_hash, state_data_vector);
    marshal_uint256_t(state_data.arb_gas_remaining, state_data_vector);
    state_data_vector.push_back(static_cast<unsigned char>(state_data.state));
    marshal_uint256_t(state_data.err_pc_hash, state_data_vector);
}

MachineOutput extractMachineOutput(
    std::vector<unsigned char>::const_iterator& iter) {
    auto fully_processed_messages = extractUint256(iter);
    auto fully_processed_inbox_accumulator = extractUint256(iter);
    auto total_steps = extractUint256(iter);
    auto arb_gas_used = extractUint256(iter);
    auto send_acc = extractUint256(iter);
    auto log_acc = extractUint256(iter);
    auto send_count = extractUint256(iter);
    auto log_count = extractUint256(iter);
    auto l1_block_number = extractUint256(iter);
    auto l2_block_number = extractUint256(iter);
    auto last_inbox_timestamp = extractUint256(iter);
    auto last_sideload_raw = extractUint256(iter);

    std::optional<uint256_t> last_sideload;
    if (last_sideload_raw != std::numeric_limits<uint256_t>::max()) {
        last_sideload = last_sideload_raw;
    }

    return MachineOutput{
        {fully_processed_messages, fully_processed_inbox_accumulator},
        total_steps,
        arb_gas_used,
        send_acc,
        log_acc,
        send_count,
        log_count,
        l1_block_number,
        l2_block_number,
        last_inbox_timestamp,
        last_sideload};
}

MachineOutput getMachineOutput(const CheckpointVariant checkpoint_variant) {
    if (std::holds_alternative<MachineOutput>(checkpoint_variant)) {
        return std::get<MachineOutput>(checkpoint_variant);
    } else {
        return std::get<MachineStateKeys>(checkpoint_variant).output;
    }
}

CheckpointVariant extractMachineStateKeys(
    const std::vector<unsigned char>& data) {
    auto iter = data.cbegin();

    auto output = extractMachineOutput(iter);

    if (iter == data.cend()) {
        // Does not include machine
        return output;
    }

    auto pc = extractUint256(iter);
    auto static_hash = extractUint256(iter);
    auto register_hash = extractUint256(iter);
    auto datastack_hash = extractUint256(iter);
    auto auxstack_hash = extractUint256(iter);
    auto arb_gas_remaining = extractUint256(iter);
    auto state = static_cast<Status>(*iter);
    ++iter;
    auto err_pc = extractUint256(iter);

    return MachineStateKeys{
        output,
        pc,
        static_hash,
        register_hash,
        datastack_hash,
        auxstack_hash,
        arb_gas_remaining,
        state,
        err_pc,
    };
}

void deleteMachineState(ReadWriteTransaction& tx,
                        MachineStateKeys& parsed_state) {
    auto delete_pc_res = deleteValue(tx, parsed_state.pc_hash);
    auto delete_static_res = deleteValue(tx, parsed_state.static_hash);
    auto delete_register_res = deleteValue(tx, parsed_state.register_hash);
    auto delete_datastack_res = deleteValue(tx, parsed_state.datastack_hash);
    auto delete_auxstack_res = deleteValue(tx, parsed_state.auxstack_hash);
    auto delete_err_pc_res = deleteValue(tx, parsed_state.err_pc_hash);

    if (!delete_pc_res.status.ok()) {
        std::cerr << "error deleting pc in checkpoint: "
                  << delete_static_res.status.ToString() << std::endl;
    }

    if (!delete_static_res.status.ok()) {
        std::cerr << "error deleting static in checkpoint: "
                  << delete_static_res.status.ToString() << std::endl;
    }

    if (!delete_register_res.status.ok()) {
        std::cerr << "error deleting register in checkpoint: "
                  << delete_register_res.status.ToString() << std::endl;
    }

    if (!delete_datastack_res.status.ok()) {
        std::cerr << "error deleting datastack in checkpoint"
                  << delete_datastack_res.status.ToString() << std::endl;
    }

    if (!delete_auxstack_res.status.ok()) {
        std::cerr << "error deleting auxstack in checkpoint: "
                  << delete_auxstack_res.status.ToString() << std::endl;
    }

    if (!delete_err_pc_res.status.ok()) {
        std::cerr << "error deleting err_pc in checkpoint: "
                  << delete_static_res.status.ToString() << std::endl;
    }
}

DeleteResults deleteMachine(ReadWriteTransaction& tx, uint256_t machine_hash) {
    std::vector<unsigned char> checkpoint_name;
    marshal_uint256_t(machine_hash, checkpoint_name);
    auto key = vecToSlice(checkpoint_name);

    auto results = deleteRefCountedData(tx, key);
    if (!results.status.ok()) {
        return DeleteResults{0, results.status,
                             std::move(results.stored_value)};
    }
    if (results.reference_count < 1) {
        auto parsed_state = extractMachineStateKeys(results.stored_value);

        if (std::holds_alternative<MachineStateKeys>(parsed_state)) {
            deleteMachineState(tx, std::get<MachineStateKeys>(parsed_state));
        }
    }
    return results;
}

DbResult<CheckpointVariant> getMachineStateKeys(
    const ReadTransaction& transaction,
    uint256_t machineHash) {
    std::vector<unsigned char> checkpoint_name;
    marshal_uint256_t(machineHash, checkpoint_name);
    auto key = vecToSlice(checkpoint_name);
    auto results = getRefCountedData(transaction, key);

    if (!results.status.ok()) {
        return results.status;
    }
    auto parsed_state = extractMachineStateKeys(results.stored_value);

    return CountedData<CheckpointVariant>{results.reference_count,
                                          parsed_state};
}

rocksdb::Status saveMachineState(ReadWriteTransaction& tx,
                                 const Machine& machine) {
    auto& machinestate = machine.machine_state;
    auto pc_val_results = saveValue(
        tx, CodePointStub{machinestate.pc,
                          machinestate.loadCurrentInstructionConst()});
    if (!pc_val_results.status.ok()) {
        return {pc_val_results.status, {}};
    }

    auto static_val_results = saveValue(tx, machinestate.static_val);
    if (!static_val_results.status.ok()) {
        return {static_val_results.status, {}};
    }

    auto register_val_results = saveValue(tx, machinestate.registerVal);
    if (!register_val_results.status.ok()) {
        return {register_val_results.status, {}};
    }

    auto datastack_tup = machinestate.stack.getTupleRepresentation();
    auto datastack_results = saveValue(tx, datastack_tup);
    if (!datastack_results.status.ok()) {
        return {datastack_results.status, {}};
    }

    auto auxstack_tup = machinestate.auxstack.getTupleRepresentation();
    auto auxstack_results = saveValue(tx, auxstack_tup);
    if (!auxstack_results.status.ok()) {
        return {auxstack_results.status, {}};
    }

    auto err_pc_val_results = saveValue(tx, machinestate.errpc);
    if (!err_pc_val_results.status.ok()) {
        return {err_pc_val_results.status, {}};
    }

    return rocksdb::Status::OK();
}

SaveResults saveTestMachine(ReadWriteTransaction& transaction,
                            Machine& machine) {
    std::vector<unsigned char> checkpoint_name;
    auto machine_hash = machine.hash();
    if (!machine_hash) {
        return {0, rocksdb::Status::NotFound()};
    }
    marshal_uint256_t(machine.hash(), checkpoint_name);
    auto key = vecToSlice(checkpoint_name);

    auto save_res = incrementReference(transaction, key);
    if (save_res.status.ok()) {
        return save_res;
    }

    auto machine_save_res = saveMachineState(transaction, machine);
    if (!machine_save_res.ok()) {
        return {0, machine_save_res};
    }

    std::vector<unsigned char> serialized_state;
    serializeMachineStateKeys(MachineStateKeys(machine.machine_state),
                              serialized_state);
    return saveValueWithRefCount(transaction, 1, key, serialized_state);
}

uint256_t MachineStateKeys::getInboxAcc() const {
    return output.fully_processed_inbox.accumulator;
}

uint256_t MachineStateKeys::getTotalMessagesRead() const {
    return output.fully_processed_inbox.count;
}
