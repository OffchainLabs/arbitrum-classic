//
// Created by Robert Gates on 5/5/20.
//

#include "avm_values/keccak.hpp"

extern "C" {
#include <keccak/KeccakHash.h>
}

#include <limits>

void keccak(const unsigned char* input,
            unsigned int inputByteLen,
            unsigned char* output) {
    // Ethereum started using Keccak and called it SHA3 before it was finalised.
    // Standard SHA3-256 (the FIPS accepted version) uses padding 0x06, but
    // Ethereum's "Keccak-256" uses padding 0x01.
    // All other constants are copied from Keccak_HashInitialize_SHA3_256 in
    // KeccakHash.h.
    Keccak_HashInstance hi;
    Keccak_HashInitialize(&hi, 1088, 512, 256, 0x01);
    Keccak_HashUpdate(
        &hi, input, inputByteLen * std::numeric_limits<unsigned char>::digits);
    Keccak_HashFinal(&hi, output);
}