//
//  tuple.cpp
//  avm
//
//  Created by Harry Kalodner on 6/24/19.
//

#include <avm/tuple.hpp>

#include <avm/util.hpp>

uint256_t Tuple::calculateHash() const {
    std::array<unsigned char, 1 + 8 * 32> tupData;
    auto oit = tupData.begin();
    tupData[0] = TUPLE + tuple_size();
    ++oit;
    for (int i = 0; i < tuple_size(); i++) {
        auto valHash = hash(get_element(i));
        std::array<uint64_t, 4> valHashInts;
        to_big_endian(valHash, valHashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(valHashInts.data()),
                  reinterpret_cast<unsigned char*>(valHashInts.data()) + 32,
                  oit);
        oit += 32;
    }

    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(tupData.data(), 1 + 32 * tuple_size(), hashData.data());
    return from_big_endian(hashData.begin(), hashData.end());
}

uint256_t zeroHash() {
    std::array<unsigned char, 1> tupData;
    tupData[0] = TUPLE;
    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(tupData.data(), 1, hashData.data());
    return from_big_endian(hashData.begin(), hashData.end());
}

std::ostream& operator<<(std::ostream& os, const Tuple& val) {
    os << "Tuple(";
    for (int i = 0; i < val.tuple_size(); i++) {
        std::cout << val.get_element(i);
        if (i < val.tuple_size() - 1) {
            os << ", ";
        }
    }
    os << ")";
    return os;
}
