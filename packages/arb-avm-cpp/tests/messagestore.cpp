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

#include "helper.hpp"

#include <data_storage/arbcore.hpp>

#include <catch2/catch.hpp>

TEST_CASE("MessageStore tests") {
    /* TODO
    DBDeleter deleter;
    auto storage = std::make_shared<DataStorage>(dbpath);
    auto store = std::make_unique<ArbCore>(storage);

    SECTION("ArbCore save, get and delete") {
        REQUIRE(!store->getNextMessage());
        REQUIRE(!store->getLastMessage());

        std::vector<uint256_t> inbox_hashes;
        inbox_hashes.emplace_back(20);
        inbox_hashes.emplace_back(21);
        inbox_hashes.emplace_back(22);

        std::vector<std::vector<unsigned char>> messages;
        messages.emplace_back(std::vector<unsigned char>(5, 1));
        messages.emplace_back(std::vector<unsigned char>(5, 2));
        messages.emplace_back(std::vector<unsigned char>(5, 3));

        auto tx = Transaction::makeTransaction(storage);
        REQUIRE(store->addMessages(*tx, 5, 12, messages, inbox_hashes).ok());
        tx->commit();
        tx = nullptr;

        auto entry5 = store->getNextMessage();
        REQUIRE(entry5);
        REQUIRE(entry5->sequence_number == 5);
        REQUIRE(entry5->message == messages[0]);
        REQUIRE(entry5->inbox_hash == inbox_hashes[0]);
        REQUIRE(entry5->last_message_in_block == false);

        auto entry7 = store->getLastMessage();
        REQUIRE(entry7);
        REQUIRE(entry7->sequence_number == 7);
        REQUIRE(entry7->message == messages[2]);
        REQUIRE(entry7->inbox_hash == inbox_hashes[2]);
        REQUIRE(entry7->last_message_in_block == true);

        tx = Transaction::makeTransaction(storage);
        REQUIRE(deleteMessagesStartingAt(*tx, 6)->ok());
        tx->commit();
        tx = nullptr;

        entry5 = store->getNextMessage();
        REQUIRE(entry5);
        REQUIRE(entry5->sequence_number == 5);
        REQUIRE(entry5->message == messages[0]);
        REQUIRE(entry5->inbox_hash == inbox_hashes[0]);
        REQUIRE(entry5->last_message_in_block == false);

        entry5 = store->getLastMessage();
        REQUIRE(entry5);
        REQUIRE(entry5->sequence_number == 5);
        REQUIRE(entry5->message == messages[0]);
        REQUIRE(entry5->inbox_hash == inbox_hashes[0]);
        REQUIRE(entry5->last_message_in_block == false);

        tx = Transaction::makeTransaction(storage);
        REQUIRE(deleteMessagesStartingAt(*tx, 5)->ok());
        tx->commit();
        tx = nullptr;

        REQUIRE(!store->getNextMessage());
        REQUIRE(!store->getLastMessage());
    }

    SECTION("ArbCore sequence check and inbox hash check") {
        std::vector<uint256_t> inbox_hashes;
        inbox_hashes.emplace_back(20);
        inbox_hashes.emplace_back(21);

        std::vector<std::vector<unsigned char>> messages;
        messages.emplace_back(std::vector<unsigned char>(5, 4));
        messages.emplace_back(std::vector<unsigned char>(5, 5));

        // Add initial messages
        auto tx = Transaction::makeTransaction(storage);
        REQUIRE(
            addMessagesWithoutCheck(*tx, 5, 12, messages, inbox_hashes).ok());
        tx->commit();
        tx = nullptr;

        std::vector<uint256_t> inbox_hashes2;
        inbox_hashes2.emplace_back(20);

        std::vector<std::vector<unsigned char>> messages2;
        messages2.emplace_back(std::vector<unsigned char>(5, 6));

        // Attempt to add message with non-contiguous sequence number
        REQUIRE(!store
                     ->addMessages(8, 13, messages2, inbox_hashes2,
                                   inbox_hashes[1] + 1)
                     ->ok());

        // Attempt to add message with incorrect previous hash
        REQUIRE(store->addMessages(7, 13, messages2, inbox_hashes2,
                                   inbox_hashes[1] + 1) == nonstd::nullopt);

        // Add message with correct previous hash
        REQUIRE(
            store->addMessages(7, 13, messages2, inbox_hashes2, inbox_hashes[1])
                ->ok());

        tx = Transaction::makeTransaction(storage);
        REQUIRE(deleteMessagesStartingAt(*tx, 5)->ok());
        tx->commit();
        tx = nullptr;
    }

    SECTION("ArbCore endian sort") {
        std::vector<uint256_t> inbox_hashes;
        inbox_hashes.emplace_back(20);

        uint256_t sequence5 = intx::from_string<uint256_t>(
            "0x000000000000000000000000000000007F7F7F7F7F7F7F7F7F7F7F7F7F7F7F7"
            "F");
        uint256_t sequence6 = intx::from_string<uint256_t>(
            "0x000000000000000000000000000000010000000000000000000000000000000"
            "0");

        std::vector<std::vector<unsigned char>> messages;
        messages.emplace_back(std::vector<unsigned char>(5, 7));
        std::vector<std::vector<unsigned char>> messages2;
        messages2.emplace_back(std::vector<unsigned char>(5, 8));
        auto tx = Transaction::makeTransaction(storage);
        REQUIRE(
            addMessagesWithoutCheck(*tx, sequence5, 12, messages, inbox_hashes)
                .ok());
        REQUIRE(
            addMessagesWithoutCheck(*tx, sequence6, 42, messages2, inbox_hashes)
                .ok());
        tx->commit();
        tx = nullptr;

        auto entry5 = store->getNextMessage();
        REQUIRE(entry5);
        REQUIRE(entry5->sequence_number == sequence5);
        REQUIRE(entry5->message == messages[0]);
        REQUIRE(entry5->inbox_hash == inbox_hashes[0]);
        REQUIRE(entry5->last_message_in_block == true);

        auto entry6 = store->getLastMessage();
        REQUIRE(entry6);
        REQUIRE(entry6->sequence_number == sequence6);
        REQUIRE(entry6->message == messages2[0]);
        REQUIRE(entry6->inbox_hash == inbox_hashes[0]);
        REQUIRE(entry6->last_message_in_block == true);

        REQUIRE(store->deleteMessage(*entry5));

        entry6 = store->getNextMessage();
        REQUIRE(entry6);
        REQUIRE(entry6->sequence_number == sequence6);
        REQUIRE(entry6->message == messages2[0]);
        REQUIRE(entry6->inbox_hash == inbox_hashes[0]);
        REQUIRE(entry6->last_message_in_block == true);

        REQUIRE(store->deleteMessage(*entry6));

        REQUIRE(!store->getNextMessage());
        REQUIRE(!store->getLastMessage());
    }
    */
}
