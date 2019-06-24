//
//  bigint.cpp
//  avm
//
//  Created by Harry Kalodner on 6/24/19.
//

#include <avm/bigint.hpp>

#include <avm/util.hpp>

uint256_t hash(const uint256_t& val) {
    std::array<unsigned char, 32> intData;
    to_big_endian(val, intData.begin());

    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(intData.data(), 32, hashData.data());
    return from_big_endian(hashData.begin(), hashData.end());
}
