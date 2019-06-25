//
//  codepoint.cpp
//  avm
//
//  Created by Harry Kalodner on 6/23/19.
//

#include <avm/codepoint.hpp>

#include <avm/tuple.hpp>
#include <avm/util.hpp>

Operation::Operation(OpCode opcode_, value immediate_)
    : opcode(opcode_), immediate(std::make_unique<value>(immediate_)) {}
Operation::Operation(const Operation& op) {
    opcode = op.opcode;
    if (op.immediate) {
        immediate = std::make_unique<value>(*op.immediate);
    }
}
Operation::Operation(Operation&&) = default;
Operation& Operation::operator=(const Operation& cp) {
    opcode = cp.opcode;
    if (cp.immediate) {
        immediate = std::make_unique<value>(*cp.immediate);
    } else {
        immediate.reset();
    }
    return *this;
}

Operation& Operation::operator=(Operation&&) = default;
Operation::~Operation() = default;

uint256_t hash(const CodePoint& cp) {
    std::array<uint64_t, 4> nextHashInts;
    to_big_endian(cp.nextHash, nextHashInts.begin());
    if (cp.op.immediate) {
        std::array<unsigned char, 66> valData;
        valData[0] = CODEPT;
        valData[1] = static_cast<unsigned char>(cp.op.opcode);
        auto immHash = ::hash(*cp.op.immediate);
        std::array<uint64_t, 4> valHashInts;
        to_big_endian(immHash, valHashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(valHashInts.data()),
                  reinterpret_cast<unsigned char*>(valHashInts.data()) + 32,
                  valData.begin() + 2);
        std::copy(reinterpret_cast<unsigned char*>(nextHashInts.data()),
                  reinterpret_cast<unsigned char*>(nextHashInts.data()) + 32,
                  valData.end() - 32);
        std::array<unsigned char, 32> hashData;
        evm::Keccak_256(valData.data(), valData.size(), hashData.data());
        return from_big_endian(hashData.begin(), hashData.end());
    } else {
        std::array<unsigned char, 34> valData;
        valData[0] = CODEPT;
        valData[1] = static_cast<unsigned char>(cp.op.opcode);
        std::copy(reinterpret_cast<unsigned char*>(nextHashInts.data()),
                  reinterpret_cast<unsigned char*>(nextHashInts.data()) + 32,
                  valData.end() - 32);
        std::array<unsigned char, 32> hashData;
        evm::Keccak_256(valData.data(), valData.size(), hashData.data());
        return from_big_endian(hashData.begin(), hashData.end());
    }
}

std::ostream& operator<<(std::ostream& os, const Operation& val) {
    if (val.immediate) {
        os << "Immediate(" << InstructionNames.at(val.opcode) << ", "
        << *val.immediate << ")";
    } else {
        os << "Basic(" << InstructionNames.at(val.opcode) << ")";
    }
    return os;
}
