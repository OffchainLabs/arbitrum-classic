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

#include <avm/machinestate/tokenTracker.hpp>
#include <bigint_utils.hpp>

#include <boost/algorithm/hex.hpp>
#include <boost/functional/hash.hpp>

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
    std::array<unsigned char, 32> val;
    to_big_endian(tokTypeVal, val.begin());
    std::copy(val.begin() + 11, val.end(), tok.begin());
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
