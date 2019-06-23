//
//  code.cpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/30/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#include <avm/code.hpp>

#include <ostream>

std::ostream& operator<<(std::ostream& os, const instr& instruction) {
    if (instruction.immediate.has_value()) {
        os << "ImmediateInstruction(" << static_cast<int>(instruction.opcode)
           << ", " << *instruction.immediate << ")";
    } else {
        os << "BasicInstruction(" << static_cast<int>(instruction.opcode)
           << ")";
    }
    return os;
}
