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

#include <data_storage/arbcore.hpp>

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
constexpr auto log_inserted = std::array<char, 1>{-60};
constexpr auto log_processed = std::array<char, 1>{-61};
constexpr auto send_inserted = std::array<char, 1>{-62};
constexpr auto send_processed = std::array<char, 1>{-63};
constexpr auto message_entry_inserted = std::array<char, 1>{-64};
constexpr auto message_entry_processed = std::array<char, 1>{-65};

ValueResult<Checkpoint> getCheckpointUsingKey(Transaction& tx,
                                              uint256_t message_sequence_number,
                                              rocksdb::Slice key_slice) {
    auto result = getVectorUsingFamilyAndKey(
        *tx.transaction, tx.datastorage->checkpoint_column.get(), key_slice);
    if (!result.status.ok()) {
        return {result.status, {}};
    }

    auto parsed_state = extractCheckpoint(message_sequence_number, result.data);

    return {result.status, parsed_state};
}

}  // namespace

// deliverMessages sends messages to core thread.  Caller needs to verify that
// getMessageStatus return MESSAGES_NONE before calling this function.
void ArbCore::deliverMessages(
    const uint256_t& first_sequence_number_,
    const uint64_t block_height_,
    const std::vector<std::vector<unsigned char>>& messages_,
    const std::vector<uint256_t>& inbox_hashes_,
    const uint256_t& previous_inbox_hash_) {
    if (message_status != MESSAGES_EMPTY) {
        throw std::runtime_error("message_status != MESSAGES_EMPTY");
    }

    first_sequence_number = first_sequence_number_;
    block_height = block_height_;
    messages = messages_;
    inbox_hashes = inbox_hashes_;
    previous_inbox_hash = previous_inbox_hash_;
}

std::unique_ptr<Transaction> ArbCore::makeTransaction() {
    return Transaction::makeTransaction(data_storage);
}

std::unique_ptr<const Transaction> ArbCore::makeConstTransaction() const {
    auto transaction =
        std::unique_ptr<rocksdb::Transaction>(data_storage->beginTransaction());
    return std::make_unique<Transaction>(data_storage, std::move(transaction));
}

void ArbCore::initialize(const LoadedExecutable& executable) {
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

bool ArbCore::initialized() const {
    auto tx = makeConstTransaction();
    std::string initial_raw;
    auto s = tx->transaction->GetForUpdate(rocksdb::ReadOptions(),
                                           rocksdb::Slice(initial_slice_label),
                                           &initial_raw);
    return s.ok();
}

std::unique_ptr<Machine> ArbCore::getInitialMachine(ValueCache& value_cache) {
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

std::unique_ptr<Machine> ArbCore::getMachine(uint256_t machineHash,
                                             ValueCache& value_cache) {
    auto transaction = makeTransaction();
    auto results = getMachineStateKeys(*transaction, machineHash);
    if (!results.status.ok()) {
        throw std::runtime_error("failed to load machine state");
    }

    return getMachineUsingStateKeys(*transaction, results.data, value_cache);
}

void ArbCore::saveCheckpoint() {
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
    marshal_uint256_t(pending_checkpoint.arb_gas_used, key);
    auto key_slice = vecToSlice(key);
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

void ArbCore::saveAssertion(uint256_t first_message_sequence_number,
                            const Assertion& assertion) {
    auto tx = Transaction::makeTransaction(data_storage);

    for (const auto& log : assertion.logs) {
        saveLog(*tx, log);
    }

    for (const auto& send : assertion.sends) {
        saveSend(*tx, send);
    }

    auto status = tx->commit();
    if (!status.ok()) {
        throw std::runtime_error("error saving assertion: " +
                                 status.ToString());
    }

    pending_checkpoint.arb_gas_used += assertion.gasCount;
    pending_checkpoint.message_sequence_number_processed =
        first_message_sequence_number + assertion.inbox_messages_consumed - 1;
    pending_checkpoint.send_count += assertion.sends.size();
    pending_checkpoint.log_count += assertion.logs.size();
}

uint256_t ArbCore::reorgToMessageOrBefore(
    const uint256_t& message_sequence_number) {
    auto tx = Transaction::makeTransaction(data_storage);

    auto result = getCheckpointAtOrBeforeMessage(message_sequence_number);
    if (!result.status.ok()) {
        throw std::runtime_error("error getting checkpoint for reorg: " +
                                 result.status.ToString());
    }

    // TODO: truncate sends
    // TODO: cleanup checkpoints and logs to decrement references

    pending_checkpoint = result.data;
    static_cast<AggregatorStore>(data_storage)
        .reorg(pending_checkpoint.block_height);

    ValueCache value_cache{};
    machine = getMachineUsingStateKeys(
        *tx, pending_checkpoint.machine_state_keys, value_cache);

    return pending_checkpoint.message_sequence_number_processed;
}

ValueResult<Checkpoint> ArbCore::getCheckpoint(
    const uint256_t& message_sequence_number) const {
    auto tx = Transaction::makeTransaction(data_storage);

    std::vector<unsigned char> key;
    marshal_uint256_t(message_sequence_number, key);

    return getCheckpointUsingKey(*tx, uint256_t(), vecToSlice(key));
}

bool ArbCore::isCheckpointsEmpty() const {
    auto tx = Transaction::makeTransaction(data_storage);
    auto it =
        std::unique_ptr<rocksdb::Iterator>(tx->datastorage->txn_db->NewIterator(
            rocksdb::ReadOptions(), tx->datastorage->checkpoint_column.get()));
    it->SeekToLast();
    return !it->Valid();
}

uint256_t ArbCore::maxMessageSequenceNumber() {
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

DbResult<Checkpoint> ArbCore::getCheckpointAtOrBeforeMessage(
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

std::unique_ptr<Machine> ArbCore::getMachineUsingStateKeys(
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

Assertion ArbCore::run(
    uint64_t gas_limit,
    bool hard_gas_limit,
    uint256_t first_message_sequence_number,
    const std::vector<std::vector<unsigned char>>& inbox_messages,
    const nonstd::optional<uint256_t>& final_block) {
    /*
    auto assertion =
        machine->run(gas_limit, hard_gas_limit, inbox_messages, final_block);

    if (assertion.inbox_messages_consumed != inbox_messages.size()) {
        throw std::runtime_error("Not all inbox messages were consumed");
    }

    saveAssertion(first_message_sequence_number, assertion);

    return assertion;
    */

    return {};
}

void ArbCore::operator()() {
    std::unique_ptr<std::thread> machine_thread;
    std::atomic<ArbCore::machine_status_enum> machine_status{MACHINE_NONE};

    if (message_status == MESSAGES_READY) {
        // Add messages

        // Discard reorged checkpoints

        // Set ArbCore::MESSAGES_NEED_OLDER if needed

        message_status = ArbCore::MESSAGES_EMPTY;
    }

    // Check machine thread
    if (machine_status == MACHINE_FINISHED) {
        // Don't do anything if machine consumed reorged messages

        // Save logs and sends

        // Maybe save checkpoint

        machine_status = MACHINE_NONE;
    }

    if (machine_status == MACHINE_NONE &&
        false /* messages ready to be processed */) {
        // Start execution of machine with next block of messages
    }
}

void machineThread(std::mutex mutex, std::atomic<bool>&) {}

void ArbCore::stopThreads() {}

rocksdb::Status ArbCore::saveLog(Transaction& tx, const value& val) {
    auto last_result = lastLogInserted(*tx.transaction);
    if (!last_result.status.ok()) {
        return last_result.status;
    }

    auto value_result = saveValue(tx, val);
    if (!value_result.status.ok()) {
        return value_result.status;
    }

    auto next = last_result.data + 1;
    std::vector<unsigned char> key;
    marshal_uint256_t(next, key);
    auto key_slice = vecToSlice(key);

    std::vector<unsigned char> value_hash;
    marshal_uint256_t(hash_value(val), value_hash);
    rocksdb::Slice value_hash_slice(
        reinterpret_cast<const char*>(value_hash.data()), value_hash.size());

    auto status = tx.transaction->Put(data_storage->log_column.get(), key_slice,
                                      value_hash_slice);
    if (!status.ok()) {
        return status;
    }

    return updateLastLogInserted(*tx.transaction, key_slice);
}

DbResult<value> ArbCore::getLog(uint256_t index, ValueCache& valueCache) const {
    auto tx = Transaction::makeTransaction(data_storage);

    std::vector<unsigned char> key;
    marshal_uint256_t(index, key);

    auto hash_result = getUint256UsingFamilyAndKey(
        *tx->transaction, data_storage->log_column.get(), vecToSlice(key));
    if (!hash_result.status.ok()) {
        return {hash_result.status, 0, {}};
    }

    return getValue(*tx, hash_result.data, valueCache);
}

rocksdb::Status ArbCore::saveSend(Transaction& tx,
                                  const std::vector<unsigned char>& send) {
    auto last_result = lastSendInserted(*tx.transaction);
    if (!last_result.status.ok()) {
        return last_result.status;
    }

    auto next = last_result.data + 1;
    std::vector<unsigned char> key;
    marshal_uint256_t(next, key);
    auto key_slice = vecToSlice(key);

    auto status = tx.transaction->Put(tx.datastorage->send_column.get(),
                                      key_slice, vecToSlice(send));
    if (!status.ok()) {
        return status;
    }

    return updateLastSendInserted(*tx.transaction, key_slice);
}

ValueResult<std::vector<unsigned char>> ArbCore::getSend(
    uint256_t index) const {
    auto tx = Transaction::makeTransaction(data_storage);

    std::vector<unsigned char> key;
    marshal_uint256_t(index, key);
    auto key_slice = vecToSlice(key);

    return getVectorUsingFamilyAndKey(
        *tx->transaction, data_storage->send_column.get(), key_slice);
}

void checkMessages() {}

ValueResult<uint256_t> ArbCore::lastLogInserted(
    rocksdb::Transaction& transaction) {
    return getUint256UsingFamilyAndKey(transaction,
                                       data_storage->state_column.get(),
                                       vecToSlice(log_inserted));
}
rocksdb::Status ArbCore::updateLastLogInserted(
    rocksdb::Transaction& transaction,
    rocksdb::Slice value_slice) {
    return transaction.Put(data_storage->state_column.get(),
                           vecToSlice(log_inserted), value_slice);
}

ValueResult<uint256_t> ArbCore::lastLogProcessed(
    rocksdb::Transaction& transaction) {
    return getUint256UsingFamilyAndKey(transaction,
                                       data_storage->state_column.get(),
                                       vecToSlice(log_processed));
}
rocksdb::Status ArbCore::updateLastLogProcessed(
    rocksdb::Transaction& transaction,
    rocksdb::Slice value_slice) {
    return transaction.Put(data_storage->state_column.get(),
                           vecToSlice(log_processed), value_slice);
}

ValueResult<uint256_t> ArbCore::lastSendInserted(
    rocksdb::Transaction& transaction) {
    return getUint256UsingFamilyAndKey(transaction,
                                       data_storage->state_column.get(),
                                       vecToSlice(send_inserted));
}
rocksdb::Status ArbCore::updateLastSendInserted(
    rocksdb::Transaction& transaction,
    rocksdb::Slice value_slice) {
    return transaction.Put(data_storage->state_column.get(),
                           vecToSlice(send_inserted), value_slice);
}

ValueResult<uint256_t> ArbCore::lastSendProcessed(
    rocksdb::Transaction& transaction) {
    return getUint256UsingFamilyAndKey(transaction,
                                       data_storage->state_column.get(),
                                       vecToSlice(send_processed));
}
rocksdb::Status ArbCore::updateLastSendProcessed(
    rocksdb::Transaction& transaction,
    rocksdb::Slice value_slice) {
    return transaction.Put(data_storage->state_column.get(),
                           vecToSlice(send_processed), value_slice);
}

ValueResult<uint256_t> ArbCore::lastMessageEntryInserted(
    rocksdb::Transaction& transaction) {
    return getUint256UsingFamilyAndKey(transaction,
                                       data_storage->state_column.get(),
                                       vecToSlice(message_entry_inserted));
}
rocksdb::Status ArbCore::updateLastMessageEntryInserted(
    rocksdb::Transaction& transaction,
    rocksdb::Slice value_slice) {
    return transaction.Put(data_storage->state_column.get(),
                           vecToSlice(message_entry_inserted), value_slice);
}

ValueResult<uint256_t> ArbCore::lastMessageEntryProcessed(
    rocksdb::Transaction& transaction) {
    return getUint256UsingFamilyAndKey(transaction,
                                       data_storage->state_column.get(),
                                       vecToSlice(message_entry_processed));
}
rocksdb::Status ArbCore::updateLastMessageEntryProcessed(
    rocksdb::Transaction& transaction,
    rocksdb::Slice value_slice) {
    return transaction.Put(data_storage->state_column.get(),
                           vecToSlice(message_entry_processed), value_slice);
}
