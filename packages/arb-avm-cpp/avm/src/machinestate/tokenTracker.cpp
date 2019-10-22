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

#include <avm/machinestate/checkpointutils.hpp>
#include <avm/machinestate/tokenTracker.hpp>
#include <bigint_utils.hpp>

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

BalanceTracker::BalanceTracker(
    const std::vector<unsigned char>& checkpoint_data) {
    auto token_pair_length = TOKEN_VAL_LENGTH + TOKEN_TYPE_LENGTH;
    auto current_it = checkpoint_data.begin();

    unsigned int token_lookup_length;
    memcpy(&token_lookup_length, &(*current_it), sizeof(token_lookup_length));
    current_it += sizeof(token_lookup_length);

    auto total_lookup_len = token_pair_length * token_lookup_length;
    auto end_token_lookup = current_it + total_lookup_len;
    std::vector<unsigned char> token_lookup(current_it, end_token_lookup);
    current_it = end_token_lookup;

    initializeTokenLookup(token_lookup);

    unsigned int nftkey_lookup_length;
    memcpy(&nftkey_lookup_length, &(*current_it), sizeof(nftkey_lookup_length));
    current_it += sizeof(nftkey_lookup_length);

    auto total_nftlookup_len = token_pair_length * nftkey_lookup_length;
    auto end_nft_lookup = current_it + total_nftlookup_len;
    std::vector<unsigned char> nftkey_lookup(current_it, end_nft_lookup);

    initializeNftLookup(nftkey_lookup);
}

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
    insertTokenLookup(return_vector);
    insertNftLookup(return_vector);
    return return_vector;
}

// private -------------------------------------------------------------------

void BalanceTracker::initializeTokenLookup(
    std::vector<unsigned char>& token_lookup) {
    auto current_it = token_lookup.begin();

    while (current_it != token_lookup.end()) {
        std::array<unsigned char, TOKEN_TYPE_LENGTH> token_type;
        auto tok_type_end = current_it + TOKEN_TYPE_LENGTH;
        std::copy(current_it, tok_type_end, token_type.begin());
        current_it = tok_type_end;

        auto tok_val_end = current_it + TOKEN_VAL_LENGTH;
        std::vector<unsigned char> value_vector(current_it, tok_val_end);
        auto currency_val =
            Checkpoint::Utils::deserializeUint256_t(value_vector);
        current_it = tok_val_end;

        add(token_type, currency_val);
    }
}

void BalanceTracker::initializeNftLookup(
    std::vector<unsigned char>& nftkey_lookup) {
    auto nftkey_it = nftkey_lookup.begin();

    while (nftkey_it != nftkey_lookup.end()) {
        std::array<unsigned char, TOKEN_TYPE_LENGTH> token_type;
        auto tok_type_end = nftkey_it + TOKEN_TYPE_LENGTH;
        std::copy(nftkey_it, tok_type_end, token_type.begin());
        nftkey_it = tok_type_end;

        auto tok_val_end = nftkey_it + TOKEN_VAL_LENGTH;
        std::vector<unsigned char> value_vector(nftkey_it, tok_val_end);
        auto currency_val =
            Checkpoint::Utils::deserializeUint256_t(value_vector);
        nftkey_it = tok_val_end;

        nftKey key = {token_type, currency_val};
        nftLookup.insert(key);
    }
}

void BalanceTracker::insertTokenLookup(
    std::vector<unsigned char>& return_vector) {
    auto length = (unsigned int)tokenLookup.size();
    std::vector<unsigned char> length_vector(sizeof(unsigned int));
    memcpy(&length_vector[0], &length, sizeof(length));

    return_vector.insert(return_vector.end(), length_vector.begin(),
                         length_vector.end());

    for (const auto& pair : tokenLookup) {
        return_vector.insert(return_vector.end(), std::begin(pair.first),
                             std::end(pair.first));

        std::vector<unsigned char> value_vector;
        marshal_uint256_t(pair.second, value_vector);

        return_vector.insert(return_vector.end(), value_vector.begin(),
                             value_vector.end());
    }
}

void BalanceTracker::insertNftLookup(
    std::vector<unsigned char>& return_vector) {
    auto nft_length = nftLookup.size();
    std::vector<unsigned char> nft_length_vector(sizeof(unsigned int));
    memcpy(&nft_length_vector[0], &nft_length, sizeof(nft_length));

    return_vector.insert(return_vector.end(), nft_length_vector.begin(),
                         nft_length_vector.end());

    for (auto& nft_key : nftLookup) {
        return_vector.insert(return_vector.end(), std::begin(nft_key.tokenType),
                             std::end(nft_key.tokenType));

        std::vector<unsigned char> value_vector;
        marshal_uint256_t(nft_key.intVal, value_vector);

        return_vector.insert(return_vector.end(), value_vector.begin(),
                             value_vector.end());
    }
}
