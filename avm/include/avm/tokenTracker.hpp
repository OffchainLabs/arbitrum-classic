//
//  tokenTracker.hpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 6/19/19.
//

#ifndef tokenTracker_hpp
#define tokenTracker_hpp

#include "datastack.hpp"
#include "value.hpp"

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
