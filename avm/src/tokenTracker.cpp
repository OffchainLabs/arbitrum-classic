//
//  tokenTracker.cpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 6/19/19.
//

#include "avm/tokenTracker.hpp"

bool isToken(TokenType tok){
    return tok[20] == 0;
}

uint256_t fromTokenType(TokenType &tok){
    uint256_t val;
    val.backend().resize(3,3);
    memcpy(val.backend().limbs(), &tok[0], 21);
    val.backend().normalize();
    return val;
}

void toTokenType(uint256_t tokTypeVal, TokenType &tok){
    auto count = tokTypeVal.backend().size();
    auto tsize = sizeof(boost::serialization::mp::limb_type);
    if (count > 2){
        memcpy(&tok[0], tokTypeVal.backend().limbs(), 21);
    } else {
        memcpy(&tok[0], tokTypeVal.backend().limbs(), count*tsize);
        memset(&tok[0], 0, 21-count*tsize);
    }
}

bool BalanceTracker::CanSpend(const TokenType tokType, const uint256_t amount) const {
    // if token is fungible check that the spend amount <= the amount assigned to that token
    if(isToken(tokType)){
        return (amount <= tokenAmounts[tokenLookup.at(tokType)]);
    } else {
        // for non-fungible tokens, check that amount == amount assigned to that token
        nftKey key = {tokType, amount};
        if (NFTLookup.find(key) == NFTLookup.end()){
            return false;
        }
        return tokenAmounts[NFTLookup.at(key)] == amount;
    }
}

bool BalanceTracker::Spend(TokenType tokType, uint256_t amount){
    if (!CanSpend(tokType, amount)) {
        //        errors.New("not enough balance to spend")
        return false;
    }
    
    if(isToken(tokType)){
        tokenAmounts[tokenLookup[tokType]] -= amount;
        return true;
    } else {
        // for non-fungible tokens, check that amount == amount assigned to that token
        nftKey key = {tokType, amount};
        std::map<nftKey,int>::iterator it = NFTLookup.find(key);
        if (it == NFTLookup.end()){
            return false;
        }
        tokenAmounts[it->second] = 0;
        return true;
    }
}

void BalanceTracker::add(TokenType tokType, uint256_t amount){
    if(isToken(tokType)){
        std::map<TokenType,int>::iterator it = tokenLookup.find(tokType);
        if (it == tokenLookup.end()){
            //add token
            tokenAmounts.push_back(amount);
            tokenLookup.insert(std::pair<TokenType, int>(tokType, tokenAmounts.size()-1));
        } else {
            //add amount to token
            tokenAmounts[it->second] += amount;
        }
    } else {
        nftKey key = {tokType, amount};
        std::map<nftKey,int>::iterator it = NFTLookup.find(key);
        if (it == NFTLookup.end()){
            // add token
            tokenAmounts.push_back(amount);
            NFTLookup.insert(std::pair<nftKey, int>(key, tokenAmounts.size()-1));
        } else {
            //set amount
            tokenAmounts[it->second] = amount;
        }
    }
}

uint256_t BalanceTracker::tokenValue(const TokenType tokType) const{
    // if token is fungible check that the spend amount <= the amount assigned to that token
    if(isToken(tokType)){
        return tokenAmounts[tokenLookup.at(tokType)];
    } else {
        // for non-fungible tokens, check that amount == amount assigned to that token
        nftKey key = {tokType, 1};
        if (NFTLookup.find(key) == NFTLookup.end()){
            return 0;
        }
        return tokenAmounts[NFTLookup.at(key)];
    }
}
