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

#include <boost/algorithm/hex.hpp>
#include <boost/functional/hash.hpp>

namespace std {
std::size_t hash<nftKey>::operator()(const nftKey& k) const {
    using boost::hash_combine;

    std::array<unsigned char, 32> intData;
    to_big_endian(k.intVal, intData.begin());

    std::size_t seed = 3754345;
    hash_combine(seed, k.tokenType);
    hash_combine(seed, intData);
    return seed;
}

std::size_t hash<TokenType>::operator()(const TokenType& k) const {
    using boost::hash_combine;

    std::size_t seed = 9587356;
    hash_combine(seed, k);
    return seed;
}
}  // namespace std

std::ostream& operator<<(std::ostream& os, const Message& val) {
    std::string tokenType;
    boost::algorithm::hex(val.token.begin(), val.token.end(),
                          std::back_inserter(tokenType));
    return os << "Message(" << val.data << ", " << val.destination << ", "
              << val.currency << ", " << tokenType << ")";
}

bool isToken(const TokenType& tok) {
    return tok[20] == 0;
}

uint256_t fromTokenType(const TokenType& tok) {
    std::array<unsigned char, 32> val;
    val.fill(0);
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

uint256_t BalanceTracker::tokenValue(const TokenType& tokType) const {
    assert(isToken(tokType));
    return tokenLookup.at(tokType);
}

bool BalanceTracker::hasNFT(const TokenType& tokType,
                            const uint256_t& id) const {
    assert(!isToken(tokType));
    nftKey key = {tokType, id};
    return nftLookup.find(key) != nftLookup.end();
}

bool BalanceTracker::canSpend(const TokenType& tokType,
                              const uint256_t& amount) const {
    // if token is fungible check that the spend amount <= the amount assigned
    // to that token
    if (isToken(tokType)) {
        return amount <= tokenValue(tokType);
    } else {
        return hasNFT(tokType, amount);
    }
}

bool BalanceTracker::spend(const TokenType& tokType, const uint256_t& amount) {
    if (!canSpend(tokType, amount)) {
        return false;
    }

    if (isToken(tokType)) {
        tokenLookup[tokType] -= amount;
    } else {
        nftKey key = {tokType, amount};
        nftLookup.erase(key);
    }
    return true;
}

void BalanceTracker::add(const TokenType& tokType, const uint256_t& amount) {
    if (isToken(tokType)) {
        auto insertion =
            tokenLookup.insert(std::make_pair(tokType, uint256_t{0}));
        insertion.first->second += amount;
    } else {
        nftKey key = {tokType, amount};
        nftLookup.insert(key);
    }
}
