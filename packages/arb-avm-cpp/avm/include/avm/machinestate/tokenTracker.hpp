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

#ifndef tokenTracker_hpp
#define tokenTracker_hpp

#include <avm/value/codepoint.hpp>
#include <avm/value/value.hpp>

#include <avm/value/tuple.hpp>
#include <unordered_map>
#include <unordered_set>

#define TOKEN_TYPE_LENGTH 21
#define TOKEN_VAL_LENGTH 33

using TokenType = std::array<unsigned char, TOKEN_TYPE_LENGTH>;

bool isToken(const TokenType& tok);
TokenType toTokenType(const uint256_t& tokTypeVal);
uint256_t fromTokenType(const TokenType& tok);

struct Message {
    value data;
    uint256_t destination;
    uint256_t currency;
    TokenType token;

    bool deserialize(const value& val);
    value toValue(TuplePool& pool) const;
};

std::ostream& operator<<(std::ostream& os, const Message& val);

struct nftKey {
    TokenType tokenType;
    uint256_t intVal;
    bool operator<(const nftKey& n) const {
        return std::tie(tokenType, intVal) < std::tie(n.tokenType, n.intVal);
    }

    bool operator==(const nftKey& n) const {
        return std::tie(tokenType, intVal) == std::tie(n.tokenType, n.intVal);
    }
};

namespace std {
template <>
struct hash<TokenType> {
    std::size_t operator()(const TokenType& k) const;
};
}  // namespace std

namespace std {
template <>
struct hash<nftKey> {
    std::size_t operator()(const nftKey& k) const;
};
}  // namespace std

class BalanceTracker {
    std::unordered_map<TokenType, uint256_t> tokenLookup;
    std::unordered_set<nftKey> nftLookup;
    void insertTokenLookup(std::vector<unsigned char>& return_vector);
    void insertNftLookup(std::vector<unsigned char>& return_vector);
    void initializeTokenLookup(std::vector<unsigned char>& token_lookup);
    void initializeNftLookup(std::vector<unsigned char>& nftkey_lookup);

   public:
    BalanceTracker() {}
    BalanceTracker(const std::vector<unsigned char>& checkpoint_data);
    bool canSpend(const TokenType& tokType, const uint256_t& amount) const;
    bool spend(const TokenType& tokType, const uint256_t& amount);
    void add(const TokenType& tokType, const uint256_t& amount);
    uint256_t tokenValue(const TokenType& tokType) const;
    bool hasNFT(const TokenType& tokType, const uint256_t& id) const;
    std::vector<unsigned char> serializeBalanceValues();
};

#endif /* tokenTracker_hpp */
