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
    std::array<unsigned char, 32> val;
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

std::vector<unsigned char> BalanceTracker::serializeBalanceValues() {
    std::vector<unsigned char> return_vector;

    for (const auto& pair : tokenLookup) {
        std::vector<unsigned char> value_vector;
        marshal_uint256_t(pair.second, value_vector);

        return_vector.insert(return_vector.end(), std::begin(pair.first),
                             std::end(pair.first));

        return_vector.insert(return_vector.end(), value_vector.begin(),
                             value_vector.end());
    }

    return return_vector;
}

BalanceTracker::BalanceTracker() {}

BalanceTracker::BalanceTracker(std::vector<unsigned char> data) {
    auto current_it = data.begin();

    while (current_it != data.end()) {
        std::array<unsigned char, 21> token_type;

        std::copy(current_it, current_it + 21, token_type.begin());
        current_it += 21;

        std::vector<unsigned char> value_vector(current_it, current_it + 33);
        current_it += 33;

        auto buff = reinterpret_cast<char*>(&value_vector[0]);
        auto currency_val = deserialize_int(buff);

        add(token_type, currency_val);
    }
}

struct BlockSerializer {
    std::vector<unsigned char> operator()(const NotBlocked& val) const {
        std::vector<unsigned char> return_value;
        return_value.push_back((unsigned char)val.type);
        return return_value;
    }
    std::vector<unsigned char> operator()(const HaltBlocked& val) const {
        std::vector<unsigned char> return_value;
        return_value.push_back((unsigned char)val.type);

        return return_value;
    }
    std::vector<unsigned char> operator()(const ErrorBlocked& val) const {
        std::vector<unsigned char> return_value;
        return_value.push_back((unsigned char)val.type);

        return return_value;
    }
    std::vector<unsigned char> operator()(const BreakpointBlocked& val) const {
        std::vector<unsigned char> return_value;
        return_value.push_back((unsigned char)val.type);

        return return_value;
    }

    std::vector<unsigned char> operator()(const InboxBlocked& val) const {
        std::vector<unsigned char> return_value;
        return_value.push_back((unsigned char)val.type);

        std::vector<unsigned char> inbox_char_vector;
        marshal_uint256_t(val.inbox, inbox_char_vector);

        return_value.insert(return_value.end(), inbox_char_vector.begin(),
                            inbox_char_vector.end());

        return return_value;
    }
    std::vector<unsigned char> operator()(const SendBlocked& val) const {
        std::vector<unsigned char> return_value;
        return_value.push_back((unsigned char)val.type);

        std::vector<unsigned char> data_vector;
        marshal_uint256_t(val.currency, data_vector);

        return_value.insert(return_value.end(), data_vector.begin(),
                            data_vector.end());

        return_value.insert(return_value.end(), std::begin(val.tokenType),
                            std::end(val.tokenType));

        return return_value;
    }
};

std::vector<unsigned char> SerializeBlockReason(const BlockReason& val) {
    return nonstd::visit(BlockSerializer{}, val);
}

struct SerializedBlockReason {
    BlockType type;
    std::vector<unsigned char> data;
};

BlockReason deserializeBlockReason(std::vector<unsigned char> data) {
    BlockReason blockreason;

    auto current_it = data.begin();

    auto blocktype = (BlockType)*current_it;
    current_it++;

    if (blocktype == Inbox) {
        std::vector<unsigned char> inbox_vector(current_it, current_it + 33);

        auto buff = reinterpret_cast<char*>(&inbox_vector[0]);
        auto inbox = deserialize_int(buff);

        current_it += 33;

        InboxBlocked br{Inbox, inbox};

        blockreason = br;

    } else if (blocktype == Send) {
        std::vector<unsigned char> currency_vector(current_it, current_it + 33);

        auto buff = reinterpret_cast<char*>(&currency_vector[0]);
        auto currency = deserialize_int(buff);

        current_it += 33;

        std::array<unsigned char, 21> token_type;
        std::copy(current_it, current_it + 21, token_type.begin());

        current_it += 21;

        SendBlocked sb{Send, currency, token_type};

        blockreason = sb;

    } else if (blocktype == Not) {
        blockreason = NotBlocked();
    } else if (blocktype == Halt) {
        blockreason = HaltBlocked();
    } else if (blocktype == Error) {
        blockreason = ErrorBlocked();
    } else if (blocktype == Breakpoint) {
        blockreason = BreakpointBlocked();
    }

    return blockreason;
}

std::vector<unsigned char> Serialize(Status status) {
    return std::vector<unsigned char>((unsigned char)status);
}
