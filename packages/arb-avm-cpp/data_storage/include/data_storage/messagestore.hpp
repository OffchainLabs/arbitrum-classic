//
// Created by Joshua Colvin on 12/28/20.
//

#ifndef data_storage_messagestore_hpp
#define data_storage_messagestore_hpp

#include <data_storage/messagestore.hpp>

#include <data_storage/datastorage.hpp>
#include <data_storage/messageentry.hpp>

#include <cstdint>

class MessageStore {
   private:
    std::shared_ptr<DataStorage> data_storage;

   public:
    MessageStore() = delete;
    explicit MessageStore(std::shared_ptr<DataStorage> data_storage_)
        : data_storage(std::move(data_storage_)) {}

    nonstd::optional<rocksdb::Status> addMessages(
        uint256_t first_sequence_number,
        uint64_t block_height,
        const std::vector<std::vector<unsigned char>>& messages,
        const std::vector<uint256_t>& inbox_hashes,
        const uint256_t& previous_inbox_hash);
    nonstd::optional<MessageEntry> getNextMessage();
    nonstd::optional<MessageEntry> getLastMessage();
    bool deleteMessage(const MessageEntry& entry);
};

nonstd::optional<rocksdb::Status> deleteMessagesStartingAt(
    Transaction& tx,
    uint256_t sequence_number);

rocksdb::Status addMessagesWithoutCheck(
    Transaction& tx,
    uint256_t first_sequence_number,
    uint64_t block_height,
    const std::vector<std::vector<unsigned char>>& messages,
    const std::vector<uint256_t>& inbox_hashes);

#endif  // data_storage_messagestore_hpp
