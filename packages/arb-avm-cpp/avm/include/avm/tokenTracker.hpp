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

#include <avm/datastack.hpp>
#include <avm/value.hpp>

#include <stdio.h>

using TokenType = std::array<unsigned char, 21>;

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

struct nftKey {
    TokenType tokenType;
    uint256_t intVal;
    bool operator<(const nftKey& n) const {
        return (this->tokenType < n.tokenType);
    }
};

class BalanceTracker {
    std::vector<TokenType> tokenTypes;
    std::vector<uint256_t> tokenAmounts;
    std::map<TokenType, int> tokenLookup;
    std::map<nftKey, int> NFTLookup;

   public:
    bool CanSpend(const TokenType& tokType, const uint256_t& amount) const;
    bool Spend(const TokenType& tokType, const uint256_t& amount);
    void add(const TokenType& tokType, const uint256_t& amount);
    uint256_t tokenValue(const TokenType& tokType) const;
};

#endif /* tokenTracker_hpp */
