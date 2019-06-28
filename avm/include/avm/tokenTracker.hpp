//
//  tokenTracker.hpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 6/19/19.
//

#ifndef tokenTracker_hpp
#define tokenTracker_hpp

#include "value.hpp"
#include "datastack.hpp"

#include <stdio.h>

using TokenType = std::array<unsigned char, 21>;

bool isToken(TokenType tok);
void toTokenType(uint256_t tokTypeVal, TokenType &tok);
uint256_t fromTokenType(TokenType &tok);

struct Message {
    value data;
    TokenType token;
    uint256_t currency;
    uint256_t destination;
};

struct nftKey {
    TokenType tokenType;
    uint256_t intVal;
    bool operator<(const nftKey& n) const
    {
        return (this->tokenType < n.tokenType);
    }

};

class BalanceTracker {
    std::vector<TokenType> tokenTypes;
    std::vector<uint256_t> tokenAmounts;
    std::map<TokenType, int> tokenLookup;
    std::map<nftKey, int> NFTLookup;
    
public:
    bool CanSpend(const TokenType tokType, const uint256_t amount) const;
    bool Spend(TokenType tokType, uint256_t amount);
    void add(TokenType tokType, uint256_t amount);
    uint256_t tokenValue(const TokenType tokType) const;
};

#endif /* tokenTracker_hpp */
