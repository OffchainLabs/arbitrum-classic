//
//  machineoperation.cpp
//  avm
//
//  Created by Minh Truong on 9/22/19.
//

#include "avm/machineoperation.hpp"

void MachineOperation::add(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum + bNum;
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::mul(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum * bNum;
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::sub(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum - bNum;
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::div(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (bNum == 0) {
        m.state = Status::Error;
    } else {
        m.stack[1] = aNum / bNum;
    }
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::sdiv(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    const auto min = (std::numeric_limits<uint256_t>::max() / 2) + 1;

    if (bNum == 0) {
        m.state = Status::Error;
    } else if (aNum == min && bNum == -1) {
        m.stack[1] = aNum;
    } else {
        const auto signA = get_sign(aNum);
        const auto signB = get_sign(bNum);
        if (signA == -1)
            aNum = 0 - aNum;
        if (signB == -1)
            bNum = 0 - bNum;
        m.stack[1] = (aNum / bNum) * signA * signB;
    }
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::mod(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (bNum != 0) {
        m.stack[1] = aNum % bNum;
    } else {
        m.state = Status::Error;
    }
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::smod(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    if (bNum == 0) {
        m.state = Status::Error;
    } else {
        const auto signA = get_sign(aNum);
        const auto signB = get_sign(bNum);
        if (signA == -1)
            aNum = 0 - aNum;
        if (signB == -1)
            bNum = 0 - bNum;
        m.stack[1] = (aNum % bNum) * signA;
    }
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::addmod(MachineState& m) {
    m.stack.prepForMod(3);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    auto& cNum = assumeInt(m.stack[2]);

    if (cNum == 0) {
        m.state = Status::Error;
    } else {
        uint512_t aBig = aNum;
        uint512_t bBig = bNum;
        m.stack[2] = static_cast<uint256_t>((aBig + bBig) % cNum);
    }
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::mulmod(MachineState& m) {
    m.stack.prepForMod(3);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    auto& cNum = assumeInt(m.stack[2]);

    if (cNum == 0) {
        m.state = Status::Error;
    } else {
        uint512_t aBig = aNum;
        uint512_t bBig = bNum;
        m.stack[2] = static_cast<uint256_t>((aBig * bBig) % cNum);
    }
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::exp(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    uint64_t bSmall = assumeInt64(bNum);
    m.stack[1] = power(aNum, bSmall);
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::lt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = (aNum < bNum) ? 1 : 0;
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::gt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = (aNum > bNum) ? 1 : 0;
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::slt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (aNum == bNum) {
        m.stack[1] = 0;
    } else {
        uint8_t signA = get_sign(aNum);
        uint8_t signB = get_sign(bNum);

        if (signA != signB) {
            m.stack[1] = signA == 1 ? 0 : 1;
        } else {
            m.stack[1] = aNum < bNum ? 1 : 0;
        }
    }
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::sgt(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (aNum == bNum) {
        m.stack[1] = 0;
    } else {
        uint8_t signA = get_sign(aNum);
        uint8_t signB = get_sign(bNum);

        if (signA != signB) {
            m.stack[1] = signA == 1 ? 1 : 0;
        } else {
            m.stack[1] = aNum > bNum ? 1 : 0;
        }
    }
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::eq(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aVal = m.stack[0];
    auto& bVal = m.stack[1];
    m.stack[1] = aVal == bVal ? 1 : 0;
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::iszero(MachineState& m) {
    m.stack.prepForMod(1);
    auto& aNum = assumeInt(m.stack[0]);
    m.stack[0] = aNum.is_zero() ? 1 : 0;
    ++m.pc;
}

void MachineOperation::bitwiseAnd(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum & bNum;
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::bitwiseOr(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum | bNum;
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::bitwiseXor(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    m.stack[1] = aNum ^ bNum;
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::bitwiseNot(MachineState& m) {
    m.stack.prepForMod(1);
    auto& aNum = assumeInt(m.stack[0]);
    m.stack[0] = ~aNum;
    ++m.pc;
}

void MachineOperation::byte(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    if (bNum >= 32) {
        m.stack[1] = 0;
    } else {
        const auto shift = 256 - 8 - 8 * shrink<uint8_t>(bNum);
        const auto mask = uint256_t(255) << shift;
        m.stack[1] = (aNum & mask) >> shift;
    }
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::signExtend(MachineState& m) {
    m.stack.prepForMod(2);
    auto& aNum = assumeInt(m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);

    if (bNum >= 32) {
        m.stack[1] = m.stack[0];
    } else {
        uint256_t t = 248 - 8 * bNum;
        int signBit = bit(aNum, (int)(255 - t));
        uint256_t mask = power(uint256_t(2), uint64_t(255 - t)) - 1;
        if (signBit == 0) {
            m.stack[1] = aNum & mask;
        } else {
            mask ^= -1;
            m.stack[1] = aNum | mask;
        }
    }
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::hashOp(MachineState& m) {
    m.stack.prepForMod(1);
    m.stack[0] = ::hash(m.stack[0]);
    ++m.pc;
}

void MachineOperation::typeOp(MachineState& m) {
    m.stack.prepForMod(1);
    if (nonstd::holds_alternative<uint256_t>(m.stack[0]))
        m.stack[0] = NUM;
    else if (nonstd::holds_alternative<CodePoint>(m.stack[0]))
        m.stack[0] = CODEPT;
    else if (nonstd::holds_alternative<Tuple>(m.stack[0]))
        m.stack[0] = TUPLE;
    ++m.pc;
}

void MachineOperation::pop(MachineState& m) {
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::spush(MachineState& m) {
    value copiedStatic = m.staticVal;
    m.stack.push(std::move(copiedStatic));
    ++m.pc;
}

void MachineOperation::rpush(MachineState& m) {
    value copiedRegister = m.registerVal;
    m.stack.push(std::move(copiedRegister));
    ++m.pc;
}

void MachineOperation::rset(MachineState& m) {
    m.stack.prepForMod(1);
    m.registerVal = m.stack[0];
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::jump(MachineState& m) {
    m.stack.prepForMod(1);
    auto target = nonstd::get_if<CodePoint>(&m.stack[0]);
    if (target) {
        m.pc = target->pc;
    } else {
        m.state = Status::Error;
    }
    m.stack.popClear();
}

void MachineOperation::cjump(MachineState& m) {
    m.stack.prepForMod(2);
    auto target = nonstd::get_if<CodePoint>(&m.stack[0]);
    auto& bNum = assumeInt(m.stack[1]);
    if (bNum != 0) {
        if (target) {
            m.pc = target->pc;
        } else {
            m.state = Status::Error;
        }
    } else {
        ++m.pc;
    }
    m.stack.popClear();
    m.stack.popClear();
}

void MachineOperation::stackEmpty(MachineState& m) {
    if (m.stack.stacksize() == 0) {
        m.stack.push(uint256_t{1});
    } else {
        m.stack.push(uint256_t{0});
    }
    ++m.pc;
}

void MachineOperation::pcPush(MachineState& m) {
    m.stack.push(m.code[m.pc]);
    ++m.pc;
}

void MachineOperation::auxPush(MachineState& m) {
    m.stack.prepForMod(1);
    m.auxstack.push(std::move(m.stack[0]));
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::auxPop(MachineState& m) {
    m.auxstack.prepForMod(1);
    m.stack.push(std::move(m.auxstack[0]));
    m.auxstack.popClear();
    ++m.pc;
}

void MachineOperation::auxStackEmpty(MachineState& m) {
    if (m.auxstack.stacksize() == 0) {
        m.stack.push(uint256_t{1});
    } else {
        m.stack.push(uint256_t{0});
    }
    ++m.pc;
}

void MachineOperation::errPush(MachineState& m) {
    m.stack.push(m.errpc);
    ++m.pc;
}

void MachineOperation::errSet(MachineState& m) {
    m.stack.prepForMod(1);
    auto codePointVal = nonstd::get_if<CodePoint>(&m.stack[0]);
    if (!codePointVal) {
        m.state = Status::Error;
    } else {
        m.errpc = *codePointVal;
    }
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::dup0(MachineState& m) {
    value valACopy = m.stack[0];
    m.stack.push(std::move(valACopy));
    ++m.pc;
}

void MachineOperation::dup1(MachineState& m) {
    value valBCopy = m.stack[1];
    m.stack.push(std::move(valBCopy));
    ++m.pc;
}

void MachineOperation::dup2(MachineState& m) {
    value valCCopy = m.stack[2];
    m.stack.push(std::move(valCCopy));
    ++m.pc;
}

void MachineOperation::swap1(MachineState& m) {
    m.stack.prepForMod(2);
    std::swap(m.stack[0], m.stack[1]);
    ++m.pc;
}

void MachineOperation::swap2(MachineState& m) {
    m.stack.prepForMod(3);
    std::swap(m.stack[0], m.stack[2]);
    ++m.pc;
}

void MachineOperation::tget(MachineState& m) {
    m.stack.prepForMod(2);
    auto& bigIndex = assumeInt(m.stack[0]);
    auto index = assumeInt64(bigIndex);
    auto& tup = assumeTuple(m.stack[1]);
    m.stack[1] = tup.get_element(index);
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::tset(MachineState& m) {
    m.stack.prepForMod(3);
    auto& bigIndex = assumeInt(m.stack[0]);
    auto index = assumeInt64(bigIndex);
    auto& tup = assumeTuple(m.stack[1]);
    tup.set_element(index, std::move(m.stack[2]));
    m.stack[2] = std::move(tup);
    m.stack.popClear();
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::tlen(MachineState& m) {
    m.stack.prepForMod(1);
    m.stack[0] = assumeTuple(m.stack[0]).tuple_size();
    ++m.pc;
}

BlockReason MachineOperation::breakpoint(MachineState&) {
    return BreakpointBlocked{};
}

void MachineOperation::log(MachineState& m) {
    m.stack.prepForMod(1);
    m.context.logs.push_back(std::move(m.stack[0]));
    m.stack.popClear();
    ++m.pc;
}

void MachineOperation::debug(MachineState& m) {
    datastack tmpstk;
    std::cout << std::endl;
    std::cout << "full stack - size=" << m.stack.stacksize() << std::endl;
    while (m.stack.stacksize() > 0) {
        std::cout << m.stack[0] << std::endl;
        tmpstk.push(std::move(m.stack[0]));
        m.stack.popClear();
    }
    while (tmpstk.stacksize() > 0) {
        m.stack.push(std::move(tmpstk[0]));
        tmpstk.popClear();
    }
    std::cout << "register val=" << m.registerVal << std::endl << std::endl;
    ++m.pc;
}

BlockReason MachineOperation::send(MachineState& m) {
    m.stack.prepForMod(1);
    Message outMsg;
    auto success = outMsg.deserialize(m.stack[0]);
    if (!success) {
        m.state = Status::Error;
        return NotBlocked();
    }
    if (!m.balance.spend(outMsg.token, outMsg.currency)) {
        return SendBlocked{outMsg.currency, outMsg.token};
    } else {
        m.stack.popClear();
        m.context.outMessage.push_back(outMsg);
        ++m.pc;
        return NotBlocked();
    }
}

void MachineOperation::nbsend(MachineState& m) {
    m.stack.prepForMod(1);

    Message outMsg;
    auto success = outMsg.deserialize(m.stack[0]);
    if (!success) {
        m.state = Status::Error;
        return;
    }

    bool spent = m.balance.spend(outMsg.token, outMsg.currency);
    if (!spent) {
        m.stack[0] = 0;
    } else {
        m.context.outMessage.push_back(outMsg);
        m.stack[0] = 1;
    }
    ++m.pc;
}

void MachineOperation::getTime(MachineState& m) {
    Tuple tup(m.pool.get(), 2);
    tup.set_element(0, m.context.timeBounds[0]);
    tup.set_element(1, m.context.timeBounds[1]);
    m.stack.push(std::move(tup));
    ++m.pc;
}

BlockReason MachineOperation::inboxOp(MachineState& m) {
    m.stack.prepForMod(1);
    auto stackTop = nonstd::get_if<Tuple>(&m.stack[0]);
    if (stackTop && m.inbox.messages == *stackTop) {
        return InboxBlocked{hash(m.inbox.messages)};
    } else {
        value inboxCopy = m.inbox.messages;
        m.stack[0] = std::move(inboxCopy);
        ++m.pc;
        return NotBlocked{};
    }
}
