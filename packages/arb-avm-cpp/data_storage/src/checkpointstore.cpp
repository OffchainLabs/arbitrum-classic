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

#include <boost/endian/conversion.hpp>
#include <data_storage/aggregator.hpp>
#include <data_storage/checkpoint.hpp>
#include <data_storage/checkpointstore.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>

namespace {
constexpr auto message_number_size = 32;

std::array<char, message_number_size> toKey(const uint64_t& message_number) {
    // TODO need to fix
    std::array<char, message_number_size> key{};

    auto big_message_number = boost::endian::native_to_big(message_number);
    to_big_endian(big_message_number, key.begin());

    return key;
}

uint64_t keyToMessageNumber(const rocksdb::Slice& key) {
    // TODO need to fix
    auto big_message_number = intx::be::unsafe::load<uint256_t>(
        reinterpret_cast<const unsigned char*>(key.data()));

    return intx::narrow_cast<uint64_t>(big_message_number);
}
}  // namespace

void CheckpointStore::saveCheckpoint(Machine& machine) {
    auto tx = Transaction::makeTransaction(data_storage);

    auto status =
        saveMachineState(*tx, machine, pending_checkpoint.machine_state_keys);
    if (!status.ok()) {
        throw std::runtime_error("error saving machine:" + status.ToString());
    }

    // TODO Still need to populate the following:
    // inbox_accumulator_hash
    // block_hash
    // block_height

    auto key = toKey(pending_checkpoint.messages_output);
    rocksdb::Slice key_slice(key.begin(), key.size());
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

    auto commit_status = tx->commit();
    if (!commit_status.ok()) {
        throw std::runtime_error("error saving assertion: " +
                                 commit_status.ToString());
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

    pending_checkpoint.step_count = +assertion.stepCount;
    pending_checkpoint.messages_read_count += assertion.inbox_messages_consumed;
    pending_checkpoint.logs_output += assertion.logs.size();
    pending_checkpoint.messages_output += assertion.outMessages.size();
    pending_checkpoint.arb_gas_used += assertion.gasCount;
}

rocksdb::Status CheckpointStore::deleteCheckpoint(
    const uint64_t& message_number) {
    auto tx = Transaction::makeTransaction(data_storage);

    auto key = toKey(message_number);
    rocksdb::Slice key_slice(key.begin(), key.size());
    auto checkpoint_result = getCheckpointWithKey(*tx, key_slice);
    if (!checkpoint_result.status.ok()) {
        throw std::runtime_error("error getting checkpoint to delete: " +
                                 checkpoint_result.status.ToString());
    }

    deleteMachineState(*tx, checkpoint_result.data.machine_state_keys);

    auto delete_status = tx->datastorage->txn_db->DB::Delete(
        rocksdb::WriteOptions(), tx->datastorage->checkpoint_column.get(),
        key_slice);

    auto commit_status = tx->commit();
    if (!commit_status.ok()) {
        throw std::runtime_error("error committing checkpoint delete: " +
                                 commit_status.ToString());
    }

    return delete_status;
}

DbResult<Checkpoint> CheckpointStore::getCheckpoint(
    const uint64_t& message_number) const {
    auto tx = Transaction::makeTransaction(data_storage);
    auto key = toKey(message_number);

    rocksdb::Slice key_slice(key.begin(), key.size());
    return getCheckpointWithKey(*tx, key_slice);
}

bool CheckpointStore::isEmpty() const {
    auto tx = Transaction::makeTransaction(data_storage);
    auto it =
        std::unique_ptr<rocksdb::Iterator>(tx->datastorage->txn_db->NewIterator(
            rocksdb::ReadOptions(), tx->datastorage->checkpoint_column.get()));
    it->SeekToLast();
    return !it->Valid();
}

uint64_t CheckpointStore::maxMessageNumber() {
    auto tx = Transaction::makeTransaction(data_storage);
    auto it =
        std::unique_ptr<rocksdb::Iterator>(tx->datastorage->txn_db->NewIterator(
            rocksdb::ReadOptions(), tx->datastorage->checkpoint_column.get()));
    it->SeekToLast();
    if (it->Valid()) {
        return keyToMessageNumber(it->key());
    } else {
        return 0;
    }
}

DbResult<Checkpoint> CheckpointStore::atMessageOrPrevious(
    const uint64_t& message_number) {
    auto tx = Transaction::makeTransaction(data_storage);
    auto it =
        std::unique_ptr<rocksdb::Iterator>(tx->datastorage->txn_db->NewIterator(
            rocksdb::ReadOptions(), tx->datastorage->checkpoint_column.get()));
    auto key = toKey(message_number);
    rocksdb::Slice key_slice(key.begin(), key.size());
    it->SeekForPrev(key_slice);
    if (it->Valid()) {
        std::vector<unsigned char> saved_value(
            it->value().data(), it->value().data() + it->value().size());
        auto parsed_state = extractCheckpoint(saved_value);
        return DbResult<Checkpoint>{rocksdb::Status::OK(), 1, parsed_state};
    } else {
        return DbResult<Checkpoint>{rocksdb::Status::NotFound(), 0, {}};
    }
}

DbResult<Checkpoint> getCheckpointWithKey(Transaction& transaction,
                                          rocksdb::Slice key_slice) {
    std::string returned_value;

    auto status = transaction.datastorage->txn_db->DB::Get(
        rocksdb::ReadOptions(),
        transaction.datastorage->checkpoint_column.get(), key_slice,
        &returned_value);

    std::vector<unsigned char> saved_value(returned_value.begin(),
                                           returned_value.end());
    auto parsed_state = extractCheckpoint(saved_value);

    return DbResult<Checkpoint>{status, 1, parsed_state};
}
