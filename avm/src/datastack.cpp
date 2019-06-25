//
//  datastack.cpp
//  avm
//
//  Created by Harry Kalodner on 6/24/19.
//

#include <avm/datastack.hpp>
#include <avm/util.hpp>

void datastack::addHash() const {
    uint256_t prev;
    if (hashes.size() > 0) {
        prev = hashes.back();
    } else {
        prev = ::hash(Tuple());
    }
    std::array<unsigned char, 1 + 2 * 32> tupData;
    auto oit = tupData.begin();
    tupData[0] = TUPLE + 2;
    ++oit;
    auto valHash = ::hash(values[hashes.size()]);
    std::array<uint64_t, 4> valHashInts;
    to_big_endian(valHash, valHashInts.begin());
    std::copy(reinterpret_cast<unsigned char*>(valHashInts.data()),
              reinterpret_cast<unsigned char*>(valHashInts.data()) + 32, oit);
    oit += 32;
    std::array<uint64_t, 4> valHashInts2;
    to_big_endian(prev, valHashInts2.begin());
    std::copy(reinterpret_cast<unsigned char*>(valHashInts2.data()),
              reinterpret_cast<unsigned char*>(valHashInts2.data()) + 32, oit);
    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(tupData.data(), 1 + 32 * 2, hashData.data());
    hashes.emplace_back(from_big_endian(hashData.begin(), hashData.end()));
}

uint256_t datastack::hash() const {
    if (values.empty()) {
        return ::hash(Tuple());
    }
    while (hashes.size() < values.size()) {
        addHash();
    }
    return hashes.back();
}

std::ostream& operator<<(std::ostream& os, const datastack& val) {
    os << "[";
    for (uint64_t i = 0; i < val.values.size(); i++) {
        os << val.values[val.values.size() - 1 - i];
        if (i < val.values.size() - 1) {
            os << ", ";
        }
    }
    os << "]";
    return os;
}
