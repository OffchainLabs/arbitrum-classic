//
// Created by Robert Gates on 5/5/20.
//

#include "avm_values/keccak.hpp"

extern "C" {
#include <keccak/KeccakHash.h>
}

#include <limits>
#include "avm_values/util.hpp"

void keccak(const unsigned char* input,
            unsigned int inputByteLen,
            unsigned char* output) {
    evm::Keccak_256(input, inputByteLen, output);
}