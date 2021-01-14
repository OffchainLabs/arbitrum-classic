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

#include <data_storage/checkpointedmachine.hpp>

#include "value/utils.hpp"

#include <data_storage/aggregator.hpp>
#include <data_storage/checkpoint.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>
#include <data_storage/value/value.hpp>
#include <data_storage/value/valuecache.hpp>
#include <set>

namespace {
constexpr auto initial_slice_label = "initial";

DbResult<Checkpoint> getCheckpointUsingKey(Transaction& transaction,
                                           uint256_t message_sequence_number,
                                           rocksdb::Slice key_slice) {
    std::string returned_value;

    auto status = transaction.datastorage->txn_db->DB::Get(
        rocksdb::ReadOptions(),
        transaction.datastorage->checkpoint_column.get(), key_slice,
        &returned_value);

    std::vector<unsigned char> saved_value(returned_value.begin(),
                                           returned_value.end());
    auto parsed_state = extractCheckpoint(message_sequence_number, saved_value);

    return DbResult<Checkpoint>{status, 1, parsed_state};
}

}  // namespace

std::unique_ptr<Transaction> CheckpointedMachine::makeTransaction() {
    return Transaction::makeTransaction(data_storage);
}

std::unique_ptr<const Transaction> CheckpointedMachine::makeConstTransaction()
    const {
    auto transaction =
        std::unique_ptr<rocksdb::Transaction>(data_storage->beginTransaction());
    return std::make_unique<Transaction>(data_storage, std::move(transaction));
}

void CheckpointedMachine::initialize(LoadedExecutable executable) {
    auto tx = makeTransaction();
    code->addSegment(std::move(executable.code));
    machine = std::make_unique<Machine>(
        MachineState{code, std::move(executable.static_val)});
    auto res = saveMachine(*tx, *machine);
    if (!res.status.ok()) {
        throw std::runtime_error("failed to save initial machine");
    }
    std::vector<unsigned char> value_data;
    marshal_uint256_t(machine->hash(), value_data);
    rocksdb::Slice value_slice{reinterpret_cast<const char*>(value_data.data()),
                               value_data.size()};
    auto s =
        tx->transaction->Put(rocksdb::Slice(initial_slice_label), value_slice);
    if (!s.ok()) {
        throw std::runtime_error(
            "failed to save initial machine values into db");
    }
    s = tx->commit();
    if (!s.ok()) {
        throw std::runtime_error("failed to commit initial machine into db");
    }
}

bool CheckpointedMachine::initialized() const {
    auto tx = makeConstTransaction();
    std::string initial_raw;
    auto s = tx->transaction->GetForUpdate(rocksdb::ReadOptions(),
                                           rocksdb::Slice(initial_slice_label),
                                           &initial_raw);
    return s.ok();
}

std::unique_ptr<Machine> CheckpointedMachine::getInitialMachine(
    ValueCache& value_cache) {
    auto tx = makeConstTransaction();
    std::string initial_raw;
    auto s = tx->transaction->GetForUpdate(rocksdb::ReadOptions(),
                                           rocksdb::Slice(initial_slice_label),
                                           &initial_raw);
    if (!s.ok()) {
        throw std::runtime_error("failed to load initial val");
    }

    auto machine_hash = intx::be::unsafe::load<uint256_t>(
        reinterpret_cast<const unsigned char*>(initial_raw.data()));
    return getMachine(machine_hash, value_cache);
}

std::unique_ptr<Machine> CheckpointedMachine::getMachine(
    uint256_t machineHash,
    ValueCache& value_cache) {
    auto transaction = makeTransaction();
    auto results = getMachineStateKeys(*transaction, machineHash);
    if (!results.status.ok()) {
        throw std::runtime_error("failed to load machine state");
    }

    return getMachineUsingStateKeys(*transaction, results.data, value_cache);
}

void CheckpointedMachine::saveCheckpoint() {
    auto tx = Transaction::makeTransaction(data_storage);

    auto status =
        saveMachineState(*tx, *machine, pending_checkpoint.machine_state_keys);
    if (!status.ok()) {
        throw std::runtime_error("error saving machine:" + status.ToString());
    }

    // TODO Still need to populate the following:
    // processed_message_accumulator_hash
    // reorg_index

    std::vector<unsigned char> key;
    marshal_uint256_t(pending_checkpoint.message_sequence_number_processed,
                      key);
    rocksdb::Slice key_slice(reinterpret_cast<const char*>(key.data()),
                             key.size());
    auto serialized_checkpoint = serializeCheckpoint(pending_checkpoint);
    std::string value_str(serialized_checkpoint.begin(),
                          serialized_checkpoint.end());
    auto put_status = tx->datastorage->txn_db->DB::Put(
        rocksdb::WriteOptions(), tx->datastorage->checkpoint_column.get(),
        key_slice, value_str);
    if (!put_status.ok()) {
        throw std::runtime_error("error saving machine: " +
                                 put_status.ToString());
    }

    rocksdb::Status commit_status = tx->commit();
    if (!commit_status.ok()) {
        throw std::runtime_error("error saving checkpoint: " +
                                 commit_status.ToString());
    }
}

void CheckpointedMachine::saveAssertion(uint256_t first_message_sequence_number,
                                        const Assertion& assertion) {
    auto tx = Transaction::makeTransaction(data_storage);

    for (const auto& log : assertion.logs) {
        std::vector<unsigned char> logData;
        marshal_value(log, logData);
        AggregatorStore::saveLog(*tx->transaction, logData);
    }

    for (const auto& send : assertion.sends) {
        std::vector<unsigned char> sendData;
        marshal_value(send, sendData);
        AggregatorStore::saveSend(*tx->transaction, sendData);
    }

    auto status = tx->commit();
    if (!status.ok()) {
        throw std::runtime_error("error saving assertion: " +
                                 status.ToString());
    }

    pending_checkpoint.arb_gas_used += assertion.gasCount;
    pending_checkpoint.message_sequence_number_processed =
        first_message_sequence_number + assertion.inbox_messages_consumed - 1;
    pending_checkpoint.step_count += assertion.stepCount;
    pending_checkpoint.send_count += assertion.sends.size();
    pending_checkpoint.log_count += assertion.logs.size();
}

uint256_t CheckpointedMachine::reorgToMessageOrBefore(
    const uint256_t& message_sequence_number) {
    auto tx = Transaction::makeTransaction(data_storage);

    auto result = getCheckpointAtOrBeforeMessage(message_sequence_number);
    if (!result.status.ok()) {
        throw std::runtime_error("error getting checkpoint for reorg: " +
                                 result.status.ToString());
    }

    // Truncate sends and logs
    pending_checkpoint = result.data;
    static_cast<AggregatorStore>(data_storage)
        .reorg(pending_checkpoint.block_height, pending_checkpoint.send_count,
               pending_checkpoint.log_count);

    ValueCache value_cache{};
    machine = getMachineUsingStateKeys(
        *tx, pending_checkpoint.machine_state_keys, value_cache);

    return pending_checkpoint.message_sequence_number_processed;
}

DbResult<Checkpoint> CheckpointedMachine::getCheckpoint(
    const uint256_t& message_sequence_number) const {
    auto tx = Transaction::makeTransaction(data_storage);

    std::vector<unsigned char> key;
    marshal_uint256_t(message_sequence_number, key);

    return getCheckpointUsingKey(*tx, uint256_t(), vecToSlice(key));
}

bool CheckpointedMachine::isEmpty() const {
    auto tx = Transaction::makeTransaction(data_storage);
    auto it =
        std::unique_ptr<rocksdb::Iterator>(tx->datastorage->txn_db->NewIterator(
            rocksdb::ReadOptions(), tx->datastorage->checkpoint_column.get()));
    it->SeekToLast();
    return !it->Valid();
}

uint256_t CheckpointedMachine::maxMessageSequenceNumber() {
    auto tx = Transaction::makeTransaction(data_storage);
    auto it =
        std::unique_ptr<rocksdb::Iterator>(tx->datastorage->txn_db->NewIterator(
            rocksdb::ReadOptions(), tx->datastorage->checkpoint_column.get()));
    it->SeekToLast();
    if (it->Valid()) {
        auto keyBuf = it->key().data();
        return deserializeUint256t(keyBuf);
    } else {
        return 0;
    }
}

DbResult<Checkpoint> CheckpointedMachine::getCheckpointAtOrBeforeMessage(
    const uint256_t& message_sequence_number) {
    auto tx = Transaction::makeTransaction(data_storage);
    auto it =
        std::unique_ptr<rocksdb::Iterator>(tx->datastorage->txn_db->NewIterator(
            rocksdb::ReadOptions(), tx->datastorage->checkpoint_column.get()));
    std::vector<unsigned char> key;
    marshal_uint256_t(message_sequence_number, key);
    auto key_slice = vecToSlice(key);
    it->SeekForPrev(key_slice);
    if (it->Valid()) {
        std::vector<unsigned char> saved_value(
            it->value().data(), it->value().data() + it->value().size());
        auto parsed_state =
            extractCheckpoint(message_sequence_number, saved_value);
        return DbResult<Checkpoint>{rocksdb::Status::OK(), 1, parsed_state};
    } else {
        return DbResult<Checkpoint>{rocksdb::Status::NotFound(), 0, {}};
    }
}

std::unique_ptr<Machine> CheckpointedMachine::getMachineUsingStateKeys(
    Transaction& transaction,
    MachineStateKeys state_data,
    ValueCache& value_cache) {
    std::set<uint64_t> segment_ids;

    auto static_results = ::getValueImpl(transaction, state_data.static_hash,
                                         segment_ids, value_cache);
    if (!static_results.status.ok()) {
        throw std::runtime_error("failed loaded core machine static");
    }

    auto register_results = ::getValueImpl(
        transaction, state_data.register_hash, segment_ids, value_cache);
    if (!register_results.status.ok()) {
        throw std::runtime_error("failed to load machine register");
    }

    auto stack_results = ::getValueImpl(transaction, state_data.datastack_hash,
                                        segment_ids, value_cache);
    if (!stack_results.status.ok() ||
        !nonstd::holds_alternative<Tuple>(stack_results.data)) {
        throw std::runtime_error("failed to load machine stack");
    }

    auto auxstack_results = ::getValueImpl(
        transaction, state_data.auxstack_hash, segment_ids, value_cache);
    if (!auxstack_results.status.ok() ||
        !nonstd::holds_alternative<Tuple>(auxstack_results.data)) {
        throw std::runtime_error("failed to load machine auxstack");
    }

    auto staged_message_results = ::getValueImpl(
        transaction, state_data.staged_message_hash, segment_ids, value_cache);
    if (!staged_message_results.status.ok()) {
        throw std::runtime_error("failed to load machine saved message");
    }

    segment_ids.insert(state_data.pc.segment);
    segment_ids.insert(state_data.err_pc.pc.segment);

    bool loaded_segment = true;
    while (loaded_segment) {
        loaded_segment = false;
        std::set<uint64_t> next_segment_ids;
        for (auto it = segment_ids.rbegin(); it != segment_ids.rend(); ++it) {
            if (code->containsSegment(*it)) {
                // If the segment is already loaded, no need to restore it
                continue;
            }
            auto segment =
                getCodeSegment(transaction, *it, next_segment_ids, value_cache);
            code->restoreExistingSegment(std::move(segment));
            loaded_segment = true;
        }
        segment_ids = std::move(next_segment_ids);
    }

    auto state =
        MachineState{code,
                     std::move(register_results.data),
                     std::move(static_results.data),
                     Datastack(nonstd::get<Tuple>(stack_results.data)),
                     Datastack(nonstd::get<Tuple>(auxstack_results.data)),
                     state_data.arb_gas_remaining,
                     state_data.status,
                     state_data.pc,
                     state_data.err_pc,
                     std::move(staged_message_results.data.get<Tuple>())};

    return std::make_unique<Machine>(state);
}

Assertion CheckpointedMachine::run(
    uint64_t gas_limit,
    bool hard_gas_limit,
    uint256_t first_message_sequence_number,
    const std::vector<rocksdb::Slice>& inbox_messages,
    nonstd::optional<uint256_t> final_block) {
    auto assertion =
        machine->run(gas_limit, hard_gas_limit, inbox_messages, final_block);

    if (assertion.inbox_messages_consumed != inbox_messages.size()) {
        throw std::runtime_error("Not all inbox messages were consumed");
    }

    saveAssertion(first_message_sequence_number, assertion);

    return assertion;
}
