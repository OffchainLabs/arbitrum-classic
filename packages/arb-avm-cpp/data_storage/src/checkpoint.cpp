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

#include <data_storage/checkpoint.hpp>

#include "value/referencecount.hpp"
#include "value/utils.hpp"

#include <avm/machine.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/code.hpp>
#include <data_storage/value/value.hpp>

#include <boost/endian/conversion.hpp>

#include <rocksdb/status.h>
#include <rocksdb/utilities/transaction_db.h>

#include <iostream>

constexpr auto message_number_size = 32;

namespace {
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

using iterator = std::vector<unsigned char>::const_iterator;

uint256_t extractUint256(iterator& iter) {
    auto ptr = reinterpret_cast<const char*>(&*iter);
    auto int_val = deserializeUint256t(ptr);
    iter += 32;
    return int_val;
}

uint64_t extractUint64(iterator& it) {
    uint64_t big_height;
    auto big_height_ptr = reinterpret_cast<char*>(&big_height);
    std::copy(it, it + sizeof(big_height), big_height_ptr);
    it += sizeof(uint64_t);
    return boost::endian::big_to_native(big_height);
}

Checkpoint extractCheckpoint(const std::vector<unsigned char>& stored_state) {
    auto current_iter = stored_state.begin();

    auto step_count = extractUint64(current_iter);
    auto machine_hash = extractUint256(current_iter);
    auto messages_read_count = extractUint64(current_iter);
    auto inbox_accumulator_hash = extractUint256(current_iter);
    auto block_hash = extractUint256(current_iter);
    auto block_height = extractUint64(current_iter);
    auto logs_output = extractUint64(current_iter);
    auto messages_output = extractUint64(current_iter);
    auto arb_gas_used = extractUint256(current_iter);

    return Checkpoint{
        step_count,  machine_hash, messages_read_count, inbox_accumulator_hash,
        block_hash,  block_height, logs_output,         messages_output,
        arb_gas_used};
}

std::vector<unsigned char> serializeCheckpoint(const Checkpoint& state_data) {
    std::vector<unsigned char> state_data_vector;

    marshal_uint64_t(state_data.step_count, state_data_vector);
    marshal_uint256_t(state_data.machine_hash, state_data_vector);
    marshal_uint64_t(state_data.messages_read_count, state_data_vector);
    marshal_uint256_t(state_data.inbox_accumulator_hash, state_data_vector);
    marshal_uint256_t(state_data.block_hash, state_data_vector);
    marshal_uint64_t(state_data.block_height, state_data_vector);
    marshal_uint64_t(state_data.logs_output, state_data_vector);
    marshal_uint64_t(state_data.messages_output, state_data_vector);
    marshal_uint256_t(state_data.arb_gas_used, state_data_vector);

    return state_data_vector;
}
}  // namespace

SaveResults putCheckpoint(Transaction& transaction,
                          const Checkpoint& checkpoint) {
    auto key = toKey(checkpoint.messages_output);
    rocksdb::Slice key_slice(key.begin(), key.size());
    auto serialized_checkpoint = serializeCheckpoint(checkpoint);
    /*
    return transaction.datastorage->txn_db->DB::Put(rocksdb::WriteOptions(),
                                         transaction.datastorage->checkpoint_column.get(),
                                         key_slice, serialized_checkpoint);
    */
    return {};
}

rocksdb::Status deleteCheckpoint(Transaction& transaction,
                                 const uint64_t& message_count) {
    auto key = toKey(message_count);
    rocksdb::Slice key_slice(key.begin(), key.size());
    return transaction.datastorage->txn_db->DB::Delete(
        rocksdb::WriteOptions(),
        transaction.datastorage->checkpoint_column.get(), key_slice);
}

DbResult<Checkpoint> getKeyCheckpoint(Transaction& transaction,
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

DbResult<Checkpoint> getCheckpoint(Transaction& transaction,
                                   const uint64_t& message_number) {
    auto key = toKey(message_number);
    rocksdb::Slice key_slice(key.begin(), key.size());
    return getKeyCheckpoint(transaction, key_slice);
}

uint64_t maxCheckpointMessageNumber(Transaction& transaction) {
    auto it = std::unique_ptr<rocksdb::Iterator>(
        transaction.datastorage->txn_db->NewIterator(
            rocksdb::ReadOptions(),
            transaction.datastorage->checkpoint_column.get()));
    it->SeekToLast();
    if (it->Valid()) {
        return keyToMessageNumber(it->key());
    } else {
        return 0;
    }
}

DbResult<Checkpoint> checkpointAtMessageOrPrevious(Transaction& transaction,
                                                   uint64_t message_number) {
    auto it = std::unique_ptr<rocksdb::Iterator>(
        transaction.datastorage->txn_db->NewIterator(
            rocksdb::ReadOptions(),
            transaction.datastorage->checkpoint_column.get()));
    auto key = toKey(message_number);
    rocksdb::Slice key_slice(key.begin(), key.size());
    it->SeekForPrev(key_slice);
    if (it->Valid()) {
        return getKeyCheckpoint(transaction, it->key());
    } else {
        return DbResult<Checkpoint>{};
    }
}

bool isEmpty(Transaction& transaction) {
    auto it = std::unique_ptr<rocksdb::Iterator>(
        transaction.datastorage->txn_db->NewIterator(
            rocksdb::ReadOptions(),
            transaction.datastorage->checkpoint_column.get()));
    it->SeekToLast();
    return !it->Valid();
}
