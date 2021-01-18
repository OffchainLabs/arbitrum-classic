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
#include <vector>

namespace {
constexpr auto initial_slice_key = std::array<char, 1>{-10};
constexpr auto log_inserted_key = std::array<char, 1>{-60};
constexpr auto log_processed_key = std::array<char, 1>{-61};
constexpr auto send_inserted_key = std::array<char, 1>{-62};
constexpr auto send_processed_key = std::array<char, 1>{-63};
constexpr auto message_entry_inserted_key = std::array<char, 1>{-64};
constexpr auto message_entry_processed_key = std::array<char, 1>{-65};

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

ValueResult<MessageEntry> getMessageEntryUsingKey(
    Transaction& tx,
    uint256_t message_sequence_number,
    rocksdb::Slice key_slice) {
    auto result = getVectorUsingFamilyAndKey(
        *tx.transaction, tx.datastorage->messageentry_column.get(), key_slice);
    if (!result.status.ok()) {
        return {result.status, {}};
    }

    auto parsed_state =
        extractMessageEntry(message_sequence_number, vecToSlice(result.data));

    return {result.status, parsed_state};
}

ValueResult<MessageEntry> getMessageEntry(Transaction& tx,
                                          uint256_t message_sequence_number) {
    std::vector<unsigned char> previous_key;
    marshal_uint256_t(message_sequence_number, previous_key);
    return getMessageEntryUsingKey(tx, message_sequence_number,
                                   vecToSlice(previous_key));
}

}  // namespace

bool ArbCore::messagesEmpty() {
    return delivering_message_status == MESSAGES_EMPTY;
}

// deliverMessages sends messages to core thread.  Caller needs to verify that
// messagesEmpty() returns true before calling this function.
void ArbCore::deliverMessages(
    const uint256_t& first_sequence_number,
    const uint64_t block_height,
    const std::vector<std::vector<unsigned char>>& messages,
    const std::vector<uint256_t>& inbox_hashes,
    const uint256_t& previous_inbox_hash) {
    if (delivering_message_status != MESSAGES_EMPTY) {
        throw std::runtime_error("message_status != MESSAGES_EMPTY");
    }

    delivering_first_sequence_number = first_sequence_number;
    delivering_block_height = block_height;
    delivering_messages = messages;
    delivering_inbox_hashes = inbox_hashes;
    delivering_previous_inbox_hash = previous_inbox_hash;
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
    code->addSegment(executable.code);
    machine =
        std::make_unique<Machine>(MachineState{code, executable.static_val});
    auto res = saveMachine(*tx, *machine);
    if (!res.status.ok()) {
        throw std::runtime_error("failed to save initial machine");
    }
    std::vector<unsigned char> value_data;
    marshal_uint256_t(machine->hash(), value_data);
    rocksdb::Slice value_slice{reinterpret_cast<const char*>(value_data.data()),
                               value_data.size()};
    auto s = tx->transaction->Put(data_storage->state_column.get(),
                                  vecToSlice(message_entry_processed_key),
                                  value_slice);
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
    auto s = tx->transaction->GetForUpdate(
        rocksdb::ReadOptions(), data_storage->state_column.get(),
        vecToSlice(message_entry_processed_key), &initial_raw);
    return s.ok();
}

std::unique_ptr<Machine> ArbCore::getInitialMachine(ValueCache& value_cache) {
    auto tx = makeConstTransaction();
    std::string initial_raw;
    auto s = tx->transaction->GetForUpdate(
        rocksdb::ReadOptions(), data_storage->state_column.get(),
        vecToSlice(message_entry_processed_key), &initial_raw);
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

    std::vector<unsigned char> key;
    marshal_uint256_t(pending_checkpoint.arb_gas_used, key);
    auto key_slice = vecToSlice(key);
    auto serialized_checkpoint = serializeCheckpoint(pending_checkpoint);
    std::string value_str(serialized_checkpoint.begin(),
                          serialized_checkpoint.end());
    auto put_status = tx->transaction->Put(
        tx->datastorage->checkpoint_column.get(), key_slice, value_str);
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

rocksdb::Status ArbCore::reorgToMessageOrBefore(
    Transaction& tx,
    const uint256_t& message_sequence_number) {
    // Delete each checkpoint until at or below message_sequence_number
    auto it = std::unique_ptr<rocksdb::Iterator>(tx.transaction->GetIterator(
        rocksdb::ReadOptions(), tx.datastorage->checkpoint_column.get()));

    // Find first message to delete
    it->SeekToLast();
    if (!it->status().ok()) {
        return it->status();
    }

    bool good_checkpoint_found = false;
    Checkpoint last_deleted_checkpoint;
    while (it->Valid()) {
        auto keyBuf = it->key().data();
        auto gas_key = deserializeUint256t(keyBuf);

        std::vector<unsigned char> checkpoint_vector(
            it->value().data(), it->value().data() + it->value().size());
        auto checkpoint = extractCheckpoint(gas_key, checkpoint_vector);

        if (checkpoint.message_sequence_number_processed >
            message_sequence_number) {
            // Obsolete checkpoint, need to delete referenced machine
            last_deleted_checkpoint = checkpoint;
            deleteMachineState(tx, checkpoint.machine_state_keys);

            // Delete checkpoint to make sure it isn't used later
            tx.transaction->Delete(tx.datastorage->log_column.get(), it->key());
        } else {
            // Good checkpoint
            pending_checkpoint = checkpoint;
            good_checkpoint_found = true;
            break;
        }

        it->Next();
    }
    it = nullptr;

    if (!good_checkpoint_found) {
        // Nothing found, start database from scratch.
        // TODO
    }

    // Delete logs individually to handle reference counts
    auto optional_status =
        deleteLogsStartingAt(tx, pending_checkpoint.log_count + 1);
    if (optional_status && !optional_status->ok()) {
        return *optional_status;
    }
    std::vector<unsigned char> key;
    marshal_uint256_t(pending_checkpoint.log_count, key);
    auto status = updateLastLogInserted(*tx.transaction, vecToSlice(key));
    if (!status.ok()) {
        return status;
    }
    marshal_uint256_t(pending_checkpoint.send_count, key);
    status = updateLastSendInserted(*tx.transaction, vecToSlice(key));
    if (!status.ok()) {
        return status;
    }

    marshal_uint256_t(pending_checkpoint.message_sequence_number_processed,
                      key);
    status = updateLastMessageEntryProcessed(*tx.transaction, vecToSlice(key));
    if (!status.ok()) {
        return status;
    }

    return rocksdb::Status::OK();
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
    auto it = std::unique_ptr<rocksdb::Iterator>(tx->transaction->GetIterator(
        rocksdb::ReadOptions(), tx->datastorage->checkpoint_column.get()));
    it->SeekToLast();
    return !it->Valid();
}

uint256_t ArbCore::maxCheckpointGas() {
    auto tx = Transaction::makeTransaction(data_storage);
    auto it = std::unique_ptr<rocksdb::Iterator>(tx->transaction->GetIterator(
        rocksdb::ReadOptions(), tx->datastorage->checkpoint_column.get()));
    it->SeekToLast();
    if (it->Valid()) {
        auto keyBuf = it->key().data();
        return deserializeUint256t(keyBuf);
    } else {
        return 0;
    }
}

DbResult<Checkpoint> ArbCore::getCheckpointAtOrBeforeGas(
    const uint256_t& message_sequence_number) {
    auto tx = Transaction::makeTransaction(data_storage);
    auto it = std::unique_ptr<rocksdb::Iterator>(tx->transaction->GetIterator(
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

    if (delivering_message_status == MESSAGES_READY) {
        // Add messages
        auto add_status = addMessages(
            delivering_first_sequence_number, delivering_block_height,
            delivering_messages, delivering_inbox_hashes,
            delivering_previous_inbox_hash);
        if (!add_status) {
            // Messages from previous block invalid because of reorg so request
            // older messages
            delivering_message_status = MESSAGES_NEED_OLDER;
        } else if (!add_status->ok()) {
            delivering_error_string = add_status->ToString();
            delivering_message_status = MESSAGES_ERROR;
        }

        delivering_message_status = ArbCore::MESSAGES_SUCCESS;
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
                                       vecToSlice(log_inserted_key));
}
rocksdb::Status ArbCore::updateLastLogInserted(
    rocksdb::Transaction& transaction,
    rocksdb::Slice value_slice) {
    return transaction.Put(data_storage->state_column.get(),
                           vecToSlice(log_inserted_key), value_slice);
}

ValueResult<uint256_t> ArbCore::lastLogProcessed(
    rocksdb::Transaction& transaction) {
    return getUint256UsingFamilyAndKey(transaction,
                                       data_storage->state_column.get(),
                                       vecToSlice(log_processed_key));
}
rocksdb::Status ArbCore::updateLastLogProcessed(
    rocksdb::Transaction& transaction,
    rocksdb::Slice value_slice) {
    return transaction.Put(data_storage->state_column.get(),
                           vecToSlice(log_processed_key), value_slice);
}

ValueResult<uint256_t> ArbCore::lastSendInserted(
    rocksdb::Transaction& transaction) {
    return getUint256UsingFamilyAndKey(transaction,
                                       data_storage->state_column.get(),
                                       vecToSlice(send_inserted_key));
}
rocksdb::Status ArbCore::updateLastSendInserted(
    rocksdb::Transaction& transaction,
    rocksdb::Slice value_slice) {
    return transaction.Put(data_storage->state_column.get(),
                           vecToSlice(send_inserted_key), value_slice);
}

ValueResult<uint256_t> ArbCore::lastSendProcessed(
    rocksdb::Transaction& transaction) {
    return getUint256UsingFamilyAndKey(transaction,
                                       data_storage->state_column.get(),
                                       vecToSlice(send_processed_key));
}
rocksdb::Status ArbCore::updateLastSendProcessed(
    rocksdb::Transaction& transaction,
    rocksdb::Slice value_slice) {
    return transaction.Put(data_storage->state_column.get(),
                           vecToSlice(send_processed_key), value_slice);
}

ValueResult<uint256_t> ArbCore::lastMessageEntryInserted(
    rocksdb::Transaction& transaction) {
    return getUint256UsingFamilyAndKey(transaction,
                                       data_storage->state_column.get(),
                                       vecToSlice(message_entry_inserted_key));
}
rocksdb::Status ArbCore::updateLastMessageEntryInserted(
    rocksdb::Transaction& transaction,
    rocksdb::Slice value_slice) {
    return transaction.Put(data_storage->state_column.get(),
                           vecToSlice(message_entry_inserted_key), value_slice);
}

ValueResult<uint256_t> ArbCore::lastMessageEntryProcessed(
    rocksdb::Transaction& transaction) {
    return getUint256UsingFamilyAndKey(transaction,
                                       data_storage->state_column.get(),
                                       vecToSlice(message_entry_processed_key));
}
rocksdb::Status ArbCore::updateLastMessageEntryProcessed(
    rocksdb::Transaction& transaction,
    rocksdb::Slice value_slice) {
    return transaction.Put(data_storage->state_column.get(),
                           vecToSlice(message_entry_processed_key),
                           value_slice);
}

// addMessages stores all messages from given block into database.
// The last message in the list is flagged as the last message in the block.
// Returns nonstd::nullopt when caller needs to provide messages from earlier
// block.
nonstd::optional<rocksdb::Status> ArbCore::addMessages(
    const uint256_t first_sequence_number,
    const uint64_t block_height,
    const std::vector<std::vector<unsigned char>>& messages,
    const std::vector<uint256_t>& inbox_hashes,
    const uint256_t& previous_inbox_hash) {
    if (messages.size() != inbox_hashes.size()) {
        throw std::runtime_error(
            "Message and hash vector size mismatch in addMessages");
    }

    auto tx = Transaction::makeTransaction(data_storage);

    // Get the last message sequence number that was added to database
    auto last_result = lastMessageEntryInserted(*tx->transaction);
    if (last_result.status.ok()) {
        return last_result.status;
    }
    auto last_inserted_sequence_number = last_result.data;

    if (first_sequence_number > 0) {
        if (first_sequence_number > last_inserted_sequence_number + 1) {
            // Not allowed to skip message sequence numbers
            return nonstd::nullopt;
        }

        // Check that previous_inbox_hash matches hash from previous message
        auto previous_sequence_number = first_sequence_number - 1;
        auto previous_result = getMessageEntry(*tx, previous_sequence_number);
        if (!previous_result.status.ok()) {
            return previous_result.status;
        }

        if (previous_result.data.inbox_hash != previous_inbox_hash) {
            // Previous inbox doesn't match so reorg happened and
            // caller needs to try again with messages from earlier block
            return nonstd::nullopt;
        }
    }

    size_t current_message_index = 0;
    auto current_sequence_number = first_sequence_number;
    auto final_sequence_number = first_sequence_number + messages.size() - 1;

    if (messages.empty()) {
        // Truncating obsolete messages
        current_sequence_number = first_sequence_number - 1;
    }

    // Skip any valid messages that we already have in database
    while ((current_sequence_number <= last_inserted_sequence_number) &&
           (current_message_index < messages.size())) {
        auto existing_message_entry =
            getMessageEntry(*tx, current_sequence_number);
        if (!existing_message_entry.status.ok()) {
            return existing_message_entry.status;
        }
        if (existing_message_entry.data.inbox_hash !=
            inbox_hashes[current_message_index]) {
            // Entry doesn't match because of reorg
            break;
        }

        current_message_index++;
        current_sequence_number = first_sequence_number + current_message_index;
    }

    if (current_sequence_number <= last_inserted_sequence_number) {
        // Reorg occurred

        if (machine_last_sequence_number >= current_sequence_number) {
            // Machine may be running with out of date messages
            machine_abort = true;
        }

        auto last_valid_sequence_number = current_sequence_number - 1;
        std::vector<unsigned char> last_valid_key;
        marshal_uint256_t(last_valid_sequence_number, last_valid_key);

        // Truncate MessageEntries to last valid message
        updateLastMessageEntryInserted(*tx->transaction,
                                       vecToSlice(last_valid_key));

        // Reorg checkpoint and everything else
        auto reorg_status =
            reorgToMessageOrBefore(*tx, last_valid_sequence_number);
    }

    while (current_message_index < messages.size()) {
        // Encode key
        std::vector<unsigned char> key;
        marshal_uint256_t(current_sequence_number, key);

        // Encode message entry
        auto messageEntry = MessageEntry{
            current_sequence_number, inbox_hashes[current_message_index],
            block_height, current_sequence_number == final_sequence_number,
            (messages[current_message_index])};
        auto serialized_messageentry = serializeMessageEntry(messageEntry);

        // Save message entry into database
        auto put_status = tx->transaction->Put(
            tx->datastorage->messageentry_column.get(), vecToSlice(key),
            vecToSlice(serialized_messageentry));
        if (!put_status.ok()) {
            return put_status;
        }

        current_message_index++;
        current_sequence_number =
            first_sequence_number + current_sequence_number;
    }

    return tx->commit();
}

// deleteLogsStartingAt deletes the given index along with any
// newer logs. Returns nonstd::nullopt if nothing deleted.
nonstd::optional<rocksdb::Status> deleteLogsStartingAt(Transaction& tx,
                                                       uint256_t log_index) {
    auto it = std::unique_ptr<rocksdb::Iterator>(tx.transaction->GetIterator(
        rocksdb::ReadOptions(), tx.datastorage->log_column.get()));

    // Find first message to delete
    std::vector<unsigned char> key;
    marshal_uint256_t(log_index, key);
    it->Seek(vecToSlice(key));
    if (it->status().IsNotFound()) {
        // Nothing to delete
        return nonstd::nullopt;
    }
    if (!it->status().ok()) {
        return it->status();
    }

    while (it->Valid()) {
        // Remove reference to value
        auto value_hash_ptr = reinterpret_cast<const char*>(it->value().data());
        deleteValue(tx, deserializeUint256t(value_hash_ptr));

        it->Next();
    }

    return rocksdb::Status::OK();
}

// getNextMessage returns the next message to handle.
nonstd::optional<MessageEntry> ArbCore::getNextMessage() {
    auto tx = Transaction::makeTransaction(data_storage);
    auto it = std::unique_ptr<rocksdb::Iterator>(tx->transaction->GetIterator(
        rocksdb::ReadOptions(), tx->datastorage->messageentry_column.get()));

    it->SeekToFirst();
    if (!it->Valid()) {
        return nonstd::nullopt;
    }

    auto key = reinterpret_cast<const char*>(it->key().data());
    return extractMessageEntry(deserializeUint256t(key), it->value());
}

// deleteMessage deletes the provided message only if it has not changed in DB
bool ArbCore::deleteMessage(const MessageEntry& entry) {
    auto tx = Transaction::makeTransaction(data_storage);

    std::vector<unsigned char> key;
    marshal_uint256_t(entry.sequence_number, key);
    auto key_slice = vecToSlice(key);
    std::string value;
    auto get_status = tx->transaction->GetForUpdate(
        rocksdb::ReadOptions(), tx->datastorage->messageentry_column.get(),
        key_slice, &value);
    if (!get_status.ok()) {
        std::cerr << "In deleteMessage get: " << get_status.ToString()
                  << std::endl;
        return false;
    }

    auto db_entry =
        extractMessageEntry(entry.sequence_number, rocksdb::Slice(value));
    if (entry != db_entry) {
        // Entry changed, reorg probably occurred
        return false;
    }

    // Delete message entry
    auto delete_status = tx->transaction->Delete(
        tx->datastorage->messageentry_column.get(), key_slice);
    if (!delete_status.ok()) {
        std::cerr << "In deleteMessage delete: " << delete_status.ToString()
                  << std::endl;
        return false;
    }

    auto commit_status = tx->commit();
    if (!commit_status.ok()) {
        std::cerr << "In deleteMessage commit: " << commit_status.ToString()
                  << std::endl;
        return false;
    }

    return true;
}
