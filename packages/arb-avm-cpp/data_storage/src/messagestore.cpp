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

#include "data_storage/messagestore.hpp"

#include "value/utils.hpp"

#include <data_storage/aggregator.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/value.hpp>
#include <iostream>

// addMessages stores all messages from given block into database.
// The last message in the list is flagged as the last message in the block.
// Returns nonstd::nullopt when caller needs to provide messages from earlier
// block.
nonstd::optional<rocksdb::Status> MessageStore::addMessages(
    const uint256_t first_sequence_number,
    const uint64_t block_height,
    const std::vector<rocksdb::Slice>& messages,
    const std::vector<uint256_t>& inbox_hashes,
    const uint256_t& previous_inbox_hash) {
    if (first_sequence_number == 0) {
        throw std::runtime_error(
            "MessageStore::addMessages should never be called with "
            "first_sequence_number == 0");
    }

    auto tx = Transaction::makeTransaction(data_storage);

    // Check that previous_inbox_hash matches hash from previous message
    std::vector<unsigned char> previous_key;
    marshal_uint256_t(first_sequence_number - 1, previous_key);
    auto previous_key_slice = vecToSlice(previous_key);
    std::string previous_value;
    auto get_previous_status = tx->transaction->GetForUpdate(
        rocksdb::ReadOptions(), tx->datastorage->messageentry_column.get(),
        previous_key_slice, &previous_value);
    if (!get_previous_status.ok()) {
        return get_previous_status;
    }

    auto previous_entry =
        extractMessageEntry(previous_key_slice, rocksdb::Slice(previous_value));
    if (previous_entry.inbox_hash != previous_inbox_hash) {
        // Previous inbox doesn't match so reorg happened and
        // caller needs to try again with messages from earlier block
        return nonstd::nullopt;
    }

    auto add_status = addMessagesWithoutCheck(
        *tx, first_sequence_number, block_height, messages, inbox_hashes);
    if (!add_status.ok()) {
        return add_status;
    }

    return tx->commit();
}

rocksdb::Status addMessagesWithoutCheck(
    Transaction& tx,
    const uint256_t first_sequence_number,
    const uint64_t block_height,
    const std::vector<rocksdb::Slice>& messages,
    const std::vector<uint256_t>& inbox_hashes) {
    if (messages.size() != inbox_hashes.size()) {
        throw std::runtime_error(
            "Message and hash vector size mismatch in addMessagesWithoutCheck");
    }

    // If reorg occurred need to delete any obsolete messages
    auto delete_status = deleteMessagesStartingAt(tx, first_sequence_number);
    if (delete_status.has_value()) {
        if (!delete_status->ok()) {
            return *delete_status;
        }

        // Reorg occurred
        // TODO: Add entry into reorg table so checkpointedmachine knows to
        // update.
    }

    auto final_sequence_number = first_sequence_number + messages.size() - 1;
    auto current_sequence_number = first_sequence_number;
    for (size_t i = 0; i < messages.size(); i++) {
        // Encode key
        std::vector<unsigned char> key;
        marshal_uint256_t(current_sequence_number, key);

        // Encode message entry
        auto messageEntry = MessageEntry{
            current_sequence_number, inbox_hashes[i], block_height,
            current_sequence_number == final_sequence_number, (messages[i])};
        auto serialized_messageentry = serializeMessageEntry(messageEntry);

        // Save message entry into database
        auto put_status = tx.datastorage->txn_db->DB::Put(
            rocksdb::WriteOptions(), tx.datastorage->messageentry_column.get(),
            vecToSlice(key), vecToSlice(serialized_messageentry));
        if (!put_status.ok()) {
            return put_status;
        }

        current_sequence_number += 1;
    }

    return rocksdb::Status::OK();
}

// deleteMessagesStartingAt deletes the given sequence number along with any
// newer messages. Returns nonstd::nullopt if nothing deleted.
nonstd::optional<rocksdb::Status> deleteMessagesStartingAt(
    Transaction& tx,
    uint256_t sequence_number) {
    auto it =
        std::unique_ptr<rocksdb::Iterator>(tx.datastorage->txn_db->NewIterator(
            rocksdb::ReadOptions(), tx.datastorage->messageentry_column.get()));

    // Find first message to delete
    std::vector<unsigned char> key;
    marshal_uint256_t(sequence_number, key);
    it->Seek(vecToSlice(key));
    if (it->status().IsNotFound()) {
        // Nothing to delete
        return nonstd::nullopt;
    }
    if (!it->status().ok()) {
        return it->status();
    }

    while (it->Valid()) {
        // Delete message entry
        tx.transaction->Delete(tx.datastorage->messageentry_column.get(),
                               it->key());

        it->Next();
    }

    return rocksdb::Status::OK();
}

// getNextMessage returns the next message to handle.
nonstd::optional<MessageEntry> MessageStore::getNextMessage() {
    auto tx = Transaction::makeTransaction(data_storage);
    auto it =
        std::unique_ptr<rocksdb::Iterator>(tx->datastorage->txn_db->NewIterator(
            rocksdb::ReadOptions(),
            tx->datastorage->messageentry_column.get()));

    it->SeekToFirst();
    if (!it->Valid()) {
        return nonstd::nullopt;
    }

    return extractMessageEntry(it->key(), it->value());
}

// getLastMessage returns the last message added to DB.
nonstd::optional<MessageEntry> MessageStore::getLastMessage() {
    auto tx = Transaction::makeTransaction(data_storage);
    auto it =
        std::unique_ptr<rocksdb::Iterator>(tx->datastorage->txn_db->NewIterator(
            rocksdb::ReadOptions(),
            tx->datastorage->messageentry_column.get()));

    it->SeekToLast();
    if (!it->Valid()) {
        return nonstd::nullopt;
    }

    return extractMessageEntry(it->key(), it->value());
}

// deleteMessage deletes the provided message only if it has not changed in DB
bool MessageStore::deleteMessage(const MessageEntry& entry) {
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

    auto db_entry = extractMessageEntry(key_slice, rocksdb::Slice(value));
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
