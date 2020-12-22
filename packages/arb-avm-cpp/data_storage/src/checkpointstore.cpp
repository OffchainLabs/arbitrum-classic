/*
 * Copyright 2020, Offchain Labs, Inc.
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

#include <data_storage/aggregator.hpp>
#include <data_storage/checkpoint.hpp>
#include <data_storage/checkpointstore.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>

void CheckpointStore::saveCheckpoint(Machine& machine) {
    auto tx = Transaction::makeTransaction(data_storage);

    auto machine_result = saveMachine(*tx, machine);
    if (!machine_result.status.ok()) {
        throw std::runtime_error("error saving machine:" +
                                 machine_result.status.ToString());
    }

    // TODO Still need to populate the following:
    // inbox_accumulator_hash
    // block_hash
    // block_height
    pending_checkpoint.machine_hash = machine.hash();

    auto checkpoint_result = Checkpoint::putCheckpoint(*tx, pending_checkpoint);
    if (!checkpoint_result.ok()) {
        throw std::runtime_error("error saving machine: " +
                                 checkpoint_result.ToString());
    }

    auto status = tx->commit();
    if (!status.ok()) {
        throw std::runtime_error("error saving assertion: " +
                                 status.ToString());
    }
}

void CheckpointStore::saveAssertion(const Assertion& assertion) {
    auto tx = Transaction::makeTransaction(data_storage);

    for (const auto& log : assertion.logs) {
        std::vector<unsigned char> logData;
        marshal_value(log, logData);
        AggregatorStore::saveLog(*tx->transaction, logData);
    }

    for (const auto& msg : assertion.outMessages) {
        std::vector<unsigned char> msgData;
        marshal_value(msg, msgData);
        AggregatorStore::saveMessage(*tx->transaction, msgData);
    }

    auto status = tx->commit();
    if (!status.ok()) {
        throw std::runtime_error("error saving assertion: " +
                                 status.ToString());
    }

    pending_checkpoint.step_count = assertion.stepCount;
    pending_checkpoint.messages_read_count += assertion.inbox_messages_consumed;
    pending_checkpoint.logs_output += assertion.logs.size();
    pending_checkpoint.messages_output += assertion.outMessages.size();
    pending_checkpoint.arb_gas_used += assertion.gasCount;
}

rocksdb::Status CheckpointStore::deleteCheckpoint(
    const uint64_t& message_number) {
    auto tx = Transaction::makeTransaction(data_storage);

    auto delete_status = Checkpoint::deleteCheckpoint(*tx, message_number);

    auto commit_status = tx->commit();
    if (!commit_status.ok()) {
        throw std::runtime_error("error saving assertion: " +
                                 commit_status.ToString());
    }

    return delete_status;
}

DbResult<Checkpoint> CheckpointStore::getCheckpoint(
    const uint64_t& message_number) const {
    auto tx = Transaction::makeTransaction(data_storage);
    return Checkpoint::getCheckpoint(*tx, message_number);
}

uint64_t CheckpointStore::maxMessageNumber() const {
    auto tx = Transaction::makeTransaction(data_storage);
    return Checkpoint::maxCheckpointMessageNumber(*tx);
}

bool CheckpointStore::isEmpty() const {
    auto tx = Transaction::makeTransaction(data_storage);
    return Checkpoint::isEmpty(*tx);
}
