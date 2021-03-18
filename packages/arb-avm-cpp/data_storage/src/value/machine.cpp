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

void serializeStagedVariant(staged_variant message,
                            std::vector<unsigned char>& state_data_vector) {
    uint256_t next_block_height = 0;
    uint8_t inbox_message_present = 0;

    if (std::holds_alternative<uint256_t>(message)) {
        next_block_height = std::get<uint256_t>(message);
    } else if (std::holds_alternative<InboxMessage>(message)) {
        inbox_message_present = 1;
    }

    marshal_uint256_t(next_block_height, state_data_vector);
    state_data_vector.push_back(inbox_message_present);
    if (inbox_message_present == 1) {
        std::get<InboxMessage>(message).serializeImpl(state_data_vector);
    }
}

staged_variant extractStagedVariant(
    std::vector<unsigned char>::const_iterator& iter,
    const std::vector<unsigned char>::const_iterator& end) {
    staged_variant message;

    auto next_block_height = extractUint256(iter);
    uint8_t inbox_message_present = iter[0];
    iter++;
    if (inbox_message_present == 1) {
        message = extractInboxMessageImpl(iter, end);
    } else if (next_block_height != 0) {
        message = next_block_height;
    } else {
        message = std::monostate();
    }

    return message;
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

    marshal_uint256_t(state_data.output.fully_processed_inbox.count,
                      state_data_vector);
    marshal_uint256_t(state_data.output.fully_processed_inbox.accumulator,
                      state_data_vector);
    marshal_uint256_t(state_data.output.total_steps, state_data_vector);
    marshal_uint256_t(state_data.output.arb_gas_used, state_data_vector);
    marshal_uint256_t(state_data.output.send_acc, state_data_vector);
    marshal_uint256_t(state_data.output.log_acc, state_data_vector);
    marshal_uint256_t(state_data.output.send_count, state_data_vector);
    marshal_uint256_t(state_data.output.log_count, state_data_vector);

    auto last_sideload_raw = std::numeric_limits<uint256_t>::max();
    if (state_data.output.last_sideload.has_value()) {
        last_sideload_raw = *state_data.output.last_sideload;
    }
    marshal_uint256_t(last_sideload_raw, state_data_vector);

    serializeStagedVariant(state_data.staged_message, state_data_vector);
}

MachineStateKeys extractMachineStateKeys(
    std::vector<unsigned char>::const_iterator iter,
    const std::vector<unsigned char>::const_iterator end) {
    auto status = static_cast<Status>(*iter);
    ++iter;
    auto static_hash = extractUint256(iter);
    auto register_hash = extractUint256(iter);
    auto datastack_hash = extractUint256(iter);
    auto auxstack_hash = extractUint256(iter);
    auto arb_gas_remaining = extractUint256(iter);
    auto pc = extractCodePointStub(iter);
    auto err_pc = extractCodePointStub(iter);

    auto fully_processed_messages = extractUint256(iter);
    auto fully_processed_inbox_accumulator = extractUint256(iter);
    auto total_steps = extractUint256(iter);
    auto arb_gas_used = extractUint256(iter);
    auto send_acc = extractUint256(iter);
    auto log_acc = extractUint256(iter);
    auto send_count = extractUint256(iter);
    auto log_count = extractUint256(iter);
    auto last_sideload_raw = extractUint256(iter);

    std::optional<uint256_t> last_sideload;
    if (last_sideload_raw != std::numeric_limits<uint256_t>::max()) {
        last_sideload = last_sideload_raw;
    }

    auto staged_message = extractStagedVariant(iter, end);

    return MachineStateKeys{
        static_hash,
        register_hash,
        datastack_hash,
        auxstack_hash,
        arb_gas_remaining,
        pc,
        err_pc,
        std::move(staged_message),
        status,
        {{fully_processed_messages, fully_processed_inbox_accumulator},
         total_steps,
         arb_gas_used,
         send_acc,
         log_acc,
         send_count,
         log_count,
         last_sideload}};
}

void deleteMachineState(ReadWriteTransaction& tx,
                        MachineStateKeys& parsed_state) {
    std::map<uint64_t, uint64_t> segment_counts;
    auto delete_static_res =
        deleteValueImpl(tx, parsed_state.static_hash, segment_counts);
    auto delete_register_res =
        deleteValueImpl(tx, parsed_state.register_hash, segment_counts);
    auto delete_datastack_res =
        deleteValueImpl(tx, parsed_state.datastack_hash, segment_counts);
    auto delete_auxstack_res =
        deleteValueImpl(tx, parsed_state.auxstack_hash, segment_counts);

    ++segment_counts[parsed_state.pc.pc.segment];
    ++segment_counts[parsed_state.err_pc.pc.segment];

    deleteCode(tx, segment_counts);

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
}

DeleteResults deleteMachine(ReadWriteTransaction& tx, uint256_t machine_hash) {
    std::map<uint64_t, uint64_t> segment_counts;
    std::vector<unsigned char> checkpoint_name;
    marshal_uint256_t(machine_hash, checkpoint_name);
    auto key = vecToSlice(checkpoint_name);
    auto results = getRefCountedData(tx, key);

    if (!results.status.ok()) {
        return DeleteResults{0, results.status,
                             std::move(results.stored_value)};
    }

    auto delete_results = deleteRefCountedData(tx, key);

    if (delete_results.reference_count < 1) {
        auto iter = results.stored_value.cbegin();
        auto parsed_state =
            extractMachineStateKeys(iter, results.stored_value.cend());

        deleteMachineState(tx, parsed_state);
    }
    return delete_results;
}

bool MachineStateKeys::stagedMessageUnresolved() const {
    return std::holds_alternative<uint256_t>(staged_message);
}

DbResult<MachineStateKeys> getMachineStateKeys(
    const ReadTransaction& transaction,
    uint256_t machineHash) {
    std::vector<unsigned char> checkpoint_name;
    marshal_uint256_t(machineHash, checkpoint_name);
    auto key = vecToSlice(checkpoint_name);
    auto results = getRefCountedData(transaction, key);

    if (!results.status.ok()) {
        return results.status;
    }
    auto iter = results.stored_value.cbegin();
    auto parsed_state =
        extractMachineStateKeys(iter, results.stored_value.cend());

    return CountedData<MachineStateKeys>{results.reference_count, parsed_state};
}

rocksdb::Status saveMachineState(ReadWriteTransaction& tx,
                                 const Machine& machine) {
    std::map<uint64_t, uint64_t> segment_counts;

    auto& machinestate = machine.machine_state;
    auto static_val_results =
        saveValueImpl(tx, machinestate.static_val, segment_counts);
    if (!static_val_results.status.ok()) {
        return static_val_results.status;
    }

    auto register_val_results =
        saveValueImpl(tx, machinestate.registerVal, segment_counts);
    if (!register_val_results.status.ok()) {
        return register_val_results.status;
    }

    auto datastack_tup = machinestate.stack.getTupleRepresentation();
    auto datastack_results = saveValueImpl(tx, datastack_tup, segment_counts);
    if (!datastack_results.status.ok()) {
        return datastack_results.status;
    }

    auto auxstack_tup = machinestate.auxstack.getTupleRepresentation();
    auto auxstack_results = saveValueImpl(tx, auxstack_tup, segment_counts);
    if (!auxstack_results.status.ok()) {
        return auxstack_results.status;
    }

    ++segment_counts[machinestate.pc.segment];
    ++segment_counts[machinestate.errpc.pc.segment];

    auto code_status = saveCode(tx, *machinestate.code, segment_counts);
    if (!code_status.ok()) {
        return code_status;
    }

    return rocksdb::Status::OK();
}

SaveResults saveMachine(ReadWriteTransaction& transaction,
                        const Machine& machine) {
    std::vector<unsigned char> checkpoint_name;
    auto machine_hash = machine.hash();
    if (!machine_hash) {
        return {0, rocksdb::Status::NotFound()};
    }
    marshal_uint256_t(*machine.hash(), checkpoint_name);
    auto key = vecToSlice(checkpoint_name);

    auto transactionResult = getRefCountedData(transaction, key);
    if (transactionResult.status.ok()) {
        // Already saved so just increment reference count
        return saveRefCountedData(transaction, key,
                                  transactionResult.stored_value);
    }

    auto status = saveMachineState(transaction, machine);
    if (!status.ok()) {
        return {0, status};
    }
    std::vector<unsigned char> serialized_state;
    serializeMachineStateKeys(MachineStateKeys(machine.machine_state),
                              serialized_state);
    return saveRefCountedData(transaction, key, serialized_state);
}

std::optional<uint256_t> MachineStateKeys::getInboxAcc() const {
    return output.fully_processed_inbox.accWithStaged(staged_message);
}

uint256_t MachineStateKeys::getTotalMessagesRead() const {
    return output.fully_processed_inbox.countWithStaged(staged_message);
}
