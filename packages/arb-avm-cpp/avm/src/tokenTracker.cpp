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

#include "avm/tokenTracker.hpp"

#include "bigint_utils.hpp"

bool isToken(const TokenType& tok) {
    return tok[20] == 0;
}

uint256_t fromTokenType(const TokenType& tok) {
    std::array<unsigned char, 32> val;
    std::copy(tok.begin(), tok.end(), val.begin());
    return from_big_endian(val.begin(), val.end());
}

TokenType toTokenType(const uint256_t& tokTypeVal) {
    TokenType tok;
    std::vector<unsigned char> val;
    val.resize(32);
    to_big_endian(tokTypeVal, val.begin());
    std::copy(val.begin(), val.begin() + 21, tok.begin());
    return tok;
}

bool Message::deserialize(const value& val) {
    auto msgTup = nonstd::get_if<Tuple>(&val);
    if (!msgTup) {
        return false;
    }
    if (msgTup->tuple_size() != 4) {
        return false;
    }

    auto destVal = msgTup->get_element(1);
    auto destInt = nonstd::get_if<uint256_t>(&destVal);
    if (!destInt) {
        return false;
    }

    auto currencyAmountVal = msgTup->get_element(2);
    auto currencyAmountInt = nonstd::get_if<uint256_t>(&currencyAmountVal);
    if (!currencyAmountInt) {
        return false;
    }

    auto tokTypeVal = msgTup->get_element(3);
    auto tokTypeInt = nonstd::get_if<uint256_t>(&tokTypeVal);
    if (!tokTypeInt) {
        return false;
    }

    data = msgTup->get_element(0);
    destination = *destInt;
    currency = *currencyAmountInt;
    token = toTokenType(*tokTypeInt);
    return true;
}

value Message::toValue(TuplePool& pool) const {
    return Tuple{data, destination, currency, fromTokenType(token), &pool};
}

bool BalanceTracker::CanSpend(const TokenType& tokType,
                              const uint256_t& amount) const {
    // if token is fungible check that the spend amount <= the amount assigned
    // to that token
    if (isToken(tokType)) {
        return (amount <= tokenAmounts[tokenLookup.at(tokType)]);
    } else {
        // for non-fungible tokens, check that amount == amount assigned to that
        // token
        nftKey key = {tokType, amount};
        if (NFTLookup.find(key) == NFTLookup.end()) {
            return false;
        }
        return tokenAmounts[NFTLookup.at(key)] == amount;
    }
}

bool BalanceTracker::Spend(const TokenType& tokType, const uint256_t& amount) {
    if (!CanSpend(tokType, amount)) {
        //        errors.New("not enough balance to spend")
        return false;
    }

    if (isToken(tokType)) {
        tokenAmounts[tokenLookup[tokType]] -= amount;
        return true;
    } else {
        // for non-fungible tokens, check that amount == amount assigned to that
        // token
        nftKey key = {tokType, amount};
        std::map<nftKey, int>::iterator it = NFTLookup.find(key);
        if (it == NFTLookup.end()) {
            return false;
        }
        tokenAmounts[it->second] = 0;
        return true;
    }
}

void BalanceTracker::add(const TokenType& tokType, const uint256_t& amount) {
    if (isToken(tokType)) {
        std::map<TokenType, int>::iterator it = tokenLookup.find(tokType);
        if (it == tokenLookup.end()) {
            // add token
            tokenAmounts.push_back(amount);
            tokenLookup.insert(
                std::pair<TokenType, int>(tokType, tokenAmounts.size() - 1));
        } else {
            // add amount to token
            tokenAmounts[it->second] += amount;
        }
    } else {
        nftKey key = {tokType, amount};
        std::map<nftKey, int>::iterator it = NFTLookup.find(key);
        if (it == NFTLookup.end()) {
            // add token
            tokenAmounts.push_back(amount);
            NFTLookup.insert(
                std::pair<nftKey, int>(key, tokenAmounts.size() - 1));
        } else {
            // set amount
            tokenAmounts[it->second] = amount;
        }
    }
}

uint256_t BalanceTracker::tokenValue(const TokenType& tokType) const {
    // if token is fungible check that the spend amount <= the amount assigned
    // to that token
    if (isToken(tokType)) {
        return tokenAmounts[tokenLookup.at(tokType)];
    } else {
        // for non-fungible tokens, check that amount == amount assigned to that
        // token
        nftKey key = {tokType, 1};
        if (NFTLookup.find(key) == NFTLookup.end()) {
            return 0;
        }
        return tokenAmounts[NFTLookup.at(key)];
    }
}
