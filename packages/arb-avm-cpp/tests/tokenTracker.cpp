/*
 * Copyright 2019, Offchain Labs, Inc.
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

#include "avm/machinestate/tokenTracker.hpp"
#include <catch2/catch.hpp>

#define serialized_empty_balance_len 8

void serializedBalanceTracker(BalanceTracker& balance_tracker,
                              int expected_length,
                              unsigned int expected_pairs) {
    auto serialized = balance_tracker.serializeBalanceValues();

    unsigned int balance_tracker_length;
    auto current_iter = serialized.begin();
    memcpy(&balance_tracker_length, &(*current_iter), sizeof(expected_pairs));

    REQUIRE(serialized.size() == expected_length);
    REQUIRE(balance_tracker_length == expected_pairs);
}

TEST_CASE("serialize balance tracker") {
    SECTION("default") {
        auto tracker = BalanceTracker();

        serializedBalanceTracker(tracker, serialized_empty_balance_len, 0);
    }
    SECTION("default token pair, nft empty") {
        std::array<unsigned char, 21> token_type = {};
        uint256_t amount = 0;
        auto tracker = BalanceTracker();
        tracker.add(token_type, amount);

        auto token_pair_length = TOKEN_VAL_LENGTH + TOKEN_TYPE_LENGTH;
        auto total_len = token_pair_length + serialized_empty_balance_len;

        serializedBalanceTracker(tracker, total_len, 1);
    }
    SECTION("default token pairs") {
        std::array<unsigned char, 21> token_type = {};
        uint256_t amount = 0;

        std::array<unsigned char, 21> nft = {};
        nft[20] = 1;

        auto tracker = BalanceTracker();
        tracker.add(token_type, amount);
        tracker.add(nft, amount);

        auto token_pair_length = TOKEN_VAL_LENGTH + TOKEN_TYPE_LENGTH;
        auto total_len = token_pair_length * 2 + serialized_empty_balance_len;

        serializedBalanceTracker(tracker, total_len, 1);
    }
    SECTION("pair with values") {
        std::array<unsigned char, 21> token_type = {1, 99};
        uint256_t amount = 11;
        auto tracker = BalanceTracker();
        tracker.add(token_type, amount);

        auto token_pair_length = TOKEN_VAL_LENGTH + TOKEN_TYPE_LENGTH;
        auto total_len = token_pair_length + serialized_empty_balance_len;

        serializedBalanceTracker(tracker, total_len, 1);
    }
    SECTION("multiple pair with values, nft empty") {
        std::array<unsigned char, 21> token_type1 = {1, 99};
        uint256_t amount1 = 11;
        std::array<unsigned char, 21> token_type2 = {3, 9};
        uint256_t amount2 = 2;
        auto tracker = BalanceTracker();
        tracker.add(token_type1, amount1);
        tracker.add(token_type2, amount2);

        auto token_pair_length = TOKEN_VAL_LENGTH + TOKEN_TYPE_LENGTH;
        auto total_len = token_pair_length * 2 + serialized_empty_balance_len;

        serializedBalanceTracker(tracker, total_len, 2);
    }
    SECTION("multiple pair with values") {
        std::array<unsigned char, 21> token_type1 = {1, 99};
        uint256_t amount1 = 11;
        std::array<unsigned char, 21> token_type2 = {3, 9};
        uint256_t amount2 = 2;
        std::array<unsigned char, 21> nft = {};
        nft[20] = 2;

        auto tracker = BalanceTracker();
        tracker.add(token_type1, amount1);
        tracker.add(token_type2, amount2);
        tracker.add(nft, amount1);

        auto token_pair_length = TOKEN_VAL_LENGTH + TOKEN_TYPE_LENGTH;
        auto total_len = token_pair_length * 3 + serialized_empty_balance_len;

        serializedBalanceTracker(tracker, total_len, 2);
    }
}

void intializeTracker(std::vector<unsigned char> data,
                      std::unordered_map<TokenType, uint256_t> expected_lookup,
                      std::unordered_set<nftKey> expected_nft_lookup) {
    BalanceTracker tracker(data);

    for (auto& pair : expected_lookup) {
        auto val = tracker.tokenValue(pair.first);
        REQUIRE(val == pair.second);
    }

    for (auto& nft : expected_nft_lookup) {
        REQUIRE(tracker.hasNFT(nft.tokenType, nft.intVal));
    }
}

TEST_CASE("deserialize and initialize tracker") {
    SECTION("default") {
        auto tracker = BalanceTracker();
        std::unordered_map<TokenType, uint256_t> expected_lookup;
        std::unordered_set<nftKey> expected_nft_lookup;

        intializeTracker(tracker.serializeBalanceValues(), expected_lookup,
                         expected_nft_lookup);
    }
    SECTION("default pairs") {
        std::array<unsigned char, 21> token_type = {};
        uint256_t amount = 0;
        std::array<unsigned char, 21> nft_type = {};
        nft_type[20] = 1;

        auto tracker = BalanceTracker();
        tracker.add(token_type, amount);
        tracker.add(nft_type, amount);

        std::unordered_map<TokenType, uint256_t> expected_lookup;
        std::unordered_set<nftKey> expected_nft_lookup;

        nftKey key = {nft_type, amount};
        expected_nft_lookup.insert(key);
        expected_lookup[token_type] = amount;

        intializeTracker(tracker.serializeBalanceValues(), expected_lookup,
                         expected_nft_lookup);
    }
    SECTION("token with values, nft empty") {
        std::array<unsigned char, 21> token_type = {1, 99};
        uint256_t amount = 11;

        auto tracker = BalanceTracker();
        tracker.add(token_type, amount);

        std::unordered_map<TokenType, uint256_t> expected_lookup;
        expected_lookup[token_type] = amount;

        std::unordered_set<nftKey> expected_nft_lookup;

        intializeTracker(tracker.serializeBalanceValues(), expected_lookup,
                         expected_nft_lookup);
    }
    SECTION("nft with values, token empty") {
        std::array<unsigned char, 21> nft_type = {1, 7};
        nft_type[20] = 2;
        uint256_t amount = 13;

        auto tracker = BalanceTracker();
        tracker.add(nft_type, amount);

        std::unordered_map<TokenType, uint256_t> expected_lookup;

        std::unordered_set<nftKey> expected_nft_lookup;
        nftKey key = {nft_type, amount};
        expected_nft_lookup.insert(key);

        intializeTracker(tracker.serializeBalanceValues(), expected_lookup,
                         expected_nft_lookup);
    }
    SECTION("with values") {
        std::array<unsigned char, 21> token_type = {1, 99};
        uint256_t amount = 11;
        std::array<unsigned char, 21> nft_type = {1, 7};
        nft_type[20] = 2;
        uint256_t amount2 = 13;

        auto tracker = BalanceTracker();
        tracker.add(token_type, amount);
        tracker.add(nft_type, amount2);

        std::unordered_map<TokenType, uint256_t> expected_lookup;
        expected_lookup[token_type] = amount;

        std::unordered_set<nftKey> expected_nft_lookup;
        nftKey key = {nft_type, amount2};
        expected_nft_lookup.insert(key);

        intializeTracker(tracker.serializeBalanceValues(), expected_lookup,
                         expected_nft_lookup);
    }
    SECTION("multiple values") {
        std::array<unsigned char, 21> token_type1 = {1, 99};
        std::array<unsigned char, 21> token_type2 = {3, 9};
        std::array<unsigned char, 21> nft_type = {1, 7};
        nft_type[20] = 1;
        std::array<unsigned char, 21> nft_type2 = {4, 7};
        nft_type2[20] = 2;
        uint256_t amount2 = 2;
        uint256_t amount1 = 11;

        auto tracker = BalanceTracker();
        tracker.add(token_type1, amount1);
        tracker.add(token_type2, amount2);
        tracker.add(nft_type, amount2);
        tracker.add(nft_type2, amount1);

        std::unordered_map<TokenType, uint256_t> expected_lookup;
        expected_lookup[token_type1] = amount1;
        expected_lookup[token_type2] = amount2;

        std::unordered_set<nftKey> expected_nft_lookup;
        nftKey key1 = {nft_type, amount2};
        nftKey key2 = {nft_type2, amount1};
        expected_nft_lookup.insert(key1);
        expected_nft_lookup.insert(key2);

        intializeTracker(tracker.serializeBalanceValues(), expected_lookup,
                         expected_nft_lookup);
    }
}
