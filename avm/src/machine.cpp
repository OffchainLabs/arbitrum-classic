//
//  machine.cpp
//  AVMtest
//
//  Created by Harry Kalodner on 4/2/19.
//

#include <avm/machine.hpp>

#include <avm/opcodes.hpp>
#include <avm/util.hpp>

#include <iostream>

namespace {
std::vector<CodePoint> opsToCodePoints(const std::vector<Operation>& ops) {
    std::vector<CodePoint> cps;
    cps.reserve(ops.size());
    uint64_t pc = 0;
    for (auto& op : ops) {
        cps.emplace_back(pc, std::move(op), 0);
        pc++;
    }
    for (uint64_t i = 0; i < cps.size() - 1; i++) {
        cps[cps.size() - 2 - i].nextHash = hash(cps[cps.size() - 1 - i]);
    }
    return cps;
}
}  // namespace

class bad_pop_type : public std::exception {
   public:
    virtual const char* what() const noexcept override {
        return "bad_variant_access";
    }
};

class int_out_of_bounds : public std::exception {
   public:
    virtual const char* what() const noexcept override {
        return "int_out_of_bounds";
    }
};

uint256_t MachineState::hash() const {
    std::array<unsigned char, 32 * 6> data;
    auto oit = data.begin();
    {
        auto val = ::hash(code[pc]);
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }
    {
        auto val = stack.hash();
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }
    {
        auto val = auxstack.hash();
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }
    {
        auto val = ::hash(registerVal);
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }
    {
        auto val = ::hash(staticVal);
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }
    {
        auto val = ::hash(errpc);
        std::array<uint64_t, 4> hashInts;
        to_big_endian(val, hashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(hashInts.data()),
                  reinterpret_cast<unsigned char*>(hashInts.data()) + 32, oit);
        oit += 32;
    }

    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(data.data(), 32 * 6, hashData.data());
    return from_big_endian(hashData.begin(), hashData.end());
}

std::ostream& operator<<(std::ostream& os, const MachineState& val) {
    os << "codePointHash " << to_hex_str(hash(val.code[val.pc])) << "\n";
    os << "stackHash " << to_hex_str(val.stack.hash()) << "\n";
    os << "auxStackHash " << to_hex_str(val.auxstack.hash()) << "\n";
    os << "registerHash " << to_hex_str(hash(val.registerVal)) << "\n";
    os << "staticHash " << to_hex_str(hash(val.staticVal)) << "\n";
    os << "errHandlerHash " << to_hex_str(hash(val.errpc)) << "\n";
    return os;
}

std::ostream& operator<<(std::ostream& os, const Machine& val) {
    os << val.m;
    return os;
}

uint256_t& assumeInt(value& val) {
    auto aNum = mpark::get_if<uint256_t>(&val);
    if (!aNum) {
        throw bad_pop_type{};
    }
    return *aNum;
}

const uint256_t& assumeInt(const value& val) {
    auto aNum = mpark::get_if<uint256_t>(&val);
    if (!aNum) {
        throw bad_pop_type{};
    }
    return *aNum;
}

uint64_t assumeInt64(uint256_t& val) {
    if (val > std::numeric_limits<uint64_t>::max())
        throw int_out_of_bounds{};

    return static_cast<uint64_t>(val);
}

Tuple& assumeTuple(value& val) {
    auto tup = mpark::get_if<Tuple>(&val);
    if (!tup) {
        throw bad_pop_type{};
    }
    return *tup;
}

MachineState::MachineState() : pool(std::make_unique<TuplePool>()) {}

MachineState::MachineState(std::vector<CodePoint> code)
    : pool(std::make_unique<TuplePool>()), code(std::move(code)) {}

MachineState::MachineState(char*& srccode, char*& inboxdata, int inbox_sz) : MachineState() {
    char* bufptr = srccode;
    char* inboxbuf = inboxdata;

    uint32_t version;
    memcpy(&version, bufptr, sizeof(version));
    version = __builtin_bswap32(version);
    bufptr += sizeof(version);

    if (version != CURRENT_AO_VERSION) {
        std::cout << "incorrect version of .ao file" << std::endl;
        std::cout << "expected version " << CURRENT_AO_VERSION
                  << " found version " << version << std::endl;
        return;
    }

    uint32_t extentionId = 1;
    while (extentionId != 0) {
        memcpy(&extentionId, bufptr, sizeof(extentionId));
        extentionId = __builtin_bswap32(extentionId);
        bufptr += sizeof(extentionId);
        if (extentionId > 0) {
            std::cout << "found extention" << std::endl;
        }
    }
    uint64_t codeCount;
    memcpy(&codeCount, bufptr, sizeof(codeCount));
    bufptr += sizeof(codeCount);
    codeCount = boost::endian::big_to_native(codeCount);
    code.reserve(codeCount);

    std::vector<Operation> ops;
    for (uint64_t i = 0; i < codeCount; i++) {
        ops.emplace_back(deserializeOperation(bufptr, *pool));
    }
    code = opsToCodePoints(ops);

    staticVal = deserialize_value(bufptr, *pool);
    pc = 0;
    while (inboxbuf < inboxdata+inbox_sz) {
        if (*inboxbuf == 3){
            inbox = deserialize_value(inboxbuf, *pool);
        } else {
            addInboxMessage(inboxbuf);
        }
    }
    deliverMessages();
}

void MachineState::addInboxMessage(char*& newMsg){
    value msg=deserialize_value(newMsg, *pool);
    Tuple tup(pool.get(), 3);
    tup.set_element(0, uint256_t(0));
    tup.set_element(1, std::move(pendingInbox));
    tup.set_element(2, std::move(msg));
    pendingInbox = tup;
    auto &msgTup = assumeTuple(msg);
    TokenType tokType;
    value val=msgTup.get_element(3);
    auto &tokTypeVal=assumeInt(val);
    toTokenType(tokTypeVal, tokType);
    val=msgTup.get_element(2);
    auto &amt = assumeInt(val);
    context.afterBalance.add(tokType, amt);
}

void MachineState::addInboxMessage(value msg){
    Tuple tup(pool.get(), 3);
    tup.set_element(0, uint256_t(0));
    tup.set_element(1, std::move(pendingInbox));
    tup.set_element(2, std::move(msg));
    pendingInbox = tup;
    auto &msgTup = assumeTuple(msg);
    TokenType tokType;
    value val=msgTup.get_element(3);
    auto &tokTypeVal=assumeInt(val);
    toTokenType(tokTypeVal, tokType);
    val=msgTup.get_element(2);
    auto &amt = assumeInt(val);
    context.afterBalance.add(tokType, amt);
}

void MachineState::addInboxMessage(Message &msg){
    Tuple msgTup(pool.get(), 4);
    msgTup.set_element(0, msg.data);
    msgTup.set_element(1, msg.destination);
    msgTup.set_element(2, msg.currency);
    msgTup.set_element(3, fromTokenType(msg.token));
    addInboxMessage(msgTup);
}

void MachineState::deliverMessages(){
    inbox = pendingInbox;
    pendingInbox = Tuple();
}

void MachineState::setTimebounds(uint64_t timeBoundStart, uint64_t timeBoundEnd){
    context.precondition.timeBounds[0] = timeBoundStart;
    context.precondition.timeBounds[1] = timeBoundEnd;
}

void Machine::addInboxMessage(char *msg){
    m.addInboxMessage(msg);
}

void Machine::deliverMessages(){
    m.deliverMessages();
}

Assertion Machine::run(uint64_t stepCount, uint64_t timeBoundStart, uint64_t timeBoundEnd) {
    //    std::cout << "starting machine code size=" << code.size() <<
    //    std::endl; std::cout << "inbox=" << inbox << std::endl;
    m.setTimebounds(timeBoundStart, timeBoundEnd);
    uint64_t i;
    for (i = 0; i < stepCount; i++) {
        //        std::cout << "Step #" << i << std::endl;
        auto ret = runOne();
        if ((ret < 0) ||
            (m.state == ERROR)||
            (m.state == HALTED)||
            (m.state == BLOCKED))
        {
            break;
        }
    }

    if (m.state == ERROR) {
        //TODO: check if error handler set - jump there
        // set error return
        std::cout << "error state" << std::endl;
    }
    if (m.state == HALTED) {
        // set error return
        //        std::cout << "halted state" << std::endl;
    }
    //    std::cout << "full stack - size=" << stack.stacksize() << std::endl;
    //    while (stack.stacksize()>0){
    //        value A=stack.pop();
    //        std::cout << A << std::endl;
    //    }
    //    std::cout << "Total steps executed=" << i << std::endl;

    return {i};
}

int Machine::runOne() {
    //    std::cout << to_hex_str(hash()) << " " << m.code[m.pc].op <<
    //    std::endl; std::cout << *this << std::endl; std::cout<<"in
    //    runOne"<<std::endl;
    if (m.state == ERROR) {
        // set error return
        std::cout << "error state" << std::endl;
        return -1;
    }

    if (m.state == HALTED) {
        // set error return
        std::cout << "halted state" << std::endl;
        std::cout << "full stack - size=" << m.stack.stacksize() << std::endl;
        while (m.stack.stacksize() > 0) {
            std::cout << m.stack[0] << std::endl;
            m.stack.popClear();
        }
        return -2;
    }

    if (m.state == BLOCKED) {
        return -1;
    }
    //    std::cout<<"pc="<<pc<<std::endl;
    auto& instruction = m.code[m.pc];
    if (instruction.op.immediate) {
        //        std::cout<<"immediateVal = "<<*immediateVal<<std::endl;
        auto imm = *instruction.op.immediate;
        m.stack.push(std::move(imm));
    }

    try {
        //        std::cout<<"calling runInstruction"<<std::endl;
        m.runOp(instruction.op.opcode);
        //        std::cout<<"after runInstruction stack size=
        //        "<<stack.stacksize()<< std::endl; if (stack.stacksize()>0){
        //            std::cout<<"top="<<stack.peek()<< std::endl;
        //        }
    } catch (const bad_pop_type& e) {
        m.state = ERROR;
    } catch (const bad_tuple_index& e) {
        m.state = ERROR;
    }

    return 0;
}

template <typename T>
static T shrink(uint256_t i) {
    return static_cast<T>(i & std::numeric_limits<T>::max());
}

void MachineState::runOp(OpCode opcode) {
    // void Machine::runInstruction( ) {
    //    auto &instruction = testInstr;
    //    auto &instruction = code[pc];
    //    std::stringstream ss;
    //    ss << "in runInstruction, running " << std::hex <<
    //    static_cast<int>(instruction.opcode); std::cout << ss.str() <<", <"<<
    //    InstructionNames.at(instruction.opcode) <<">, stack size=
    //    "<<stack.stacksize()<< "\n"; if (stack.stacksize()>0){
    //        std::cout<<"top="<<stack.peek()<< std::endl;
    //    }
    bool shouldIncrement = true;
    switch (opcode) {
        /**************************/
        /*  Arithmetic Operations */
        /**************************/
        case OpCode::ADD: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            stack[1] = aNum + bNum;
            stack.popClear();
            break;
        }
        case OpCode::MUL: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            stack[1] = aNum * bNum;
            stack.popClear();
            break;
        }
        case OpCode::SUB: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            stack[1] = aNum - bNum;
            stack.popClear();
            break;
        }
        case OpCode::DIV: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            if (bNum == 0) {
                state = ERROR;
            } else {
                stack[1] = aNum / bNum;
            }
            stack.popClear();
            break;
        }
        case OpCode::SDIV: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            const auto min = (std::numeric_limits<uint256_t>::max() / 2) + 1;

            if (bNum == 0) {
                state = ERROR;
            } else if (aNum == min && bNum == -1) {
                stack[1] = aNum;
            } else {
                const auto signA = get_sign(aNum);
                const auto signB = get_sign(bNum);
                if (signA == -1)
                    aNum = 0 - aNum;
                if (signB == -1)
                    bNum = 0 - bNum;
                stack[1] = (aNum / bNum) * signA * signB;
            }
            stack.popClear();
            break;
        }
        case OpCode::MOD: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            if (bNum != 0) {
                stack[1] = aNum % bNum;
            } else {
                state = ERROR;
            }
            stack.popClear();
            break;
        }
        case OpCode::SMOD: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);

            if (bNum == 0) {
                state = ERROR;
            } else {
                const auto signA = get_sign(aNum);
                const auto signB = get_sign(bNum);
                if (signA == -1)
                    aNum = 0 - aNum;
                if (signB == -1)
                    bNum = 0 - bNum;
                stack[1] = (aNum % bNum) * signA;
            }
            break;
        }
        case OpCode::ADDMOD: {
            stack.prepForMod(3);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            auto& cNum = assumeInt(stack[2]);

            if (cNum == 0) {
                state = ERROR;
            } else {
                uint512_t aBig = aNum;
                uint512_t bBig = bNum;
                stack[2] = static_cast<uint256_t>((aBig + bBig) % cNum);
            }
            stack.popClear();
            stack.popClear();
            break;
        }
        case OpCode::MULMOD: {
            stack.prepForMod(3);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            auto& cNum = assumeInt(stack[2]);

            if (cNum == 0) {
                state = ERROR;
            } else {
                uint512_t aBig = aNum;
                uint512_t bBig = bNum;
                stack[2] = static_cast<uint256_t>((aBig * bBig) % cNum);
            }
            stack.popClear();
            stack.popClear();
            break;
        }
        case OpCode::EXP: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            uint64_t bSmall = assumeInt64(bNum);
            stack[1] = power(aNum, bSmall);
            stack.popClear();
            break;
        }
        /******************************************/
        /*  Comparison & Bitwise Logic Operations */
        /******************************************/
        case OpCode::LT: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            stack[1] = (aNum < bNum) ? 1 : 0;
            stack.popClear();
            break;
        }
        case OpCode::GT: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            stack[1] = (aNum > bNum) ? 1 : 0;
            stack.popClear();
            break;
        }
        case OpCode::SLT: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            if (aNum == bNum) {
                stack[1] = 0;
            } else {
                uint8_t signA = aNum.sign();
                uint8_t signB = bNum.sign();

                if (signA != signB) {
                    stack[1] = signA == 1 ? 1 : 0;
                } else {
                    stack[1] = aNum < bNum ? 1 : 0;
                }
            }
            stack.popClear();
            break;
        }
        case OpCode::SGT: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            if (aNum == bNum) {
                stack[1] = 0;
            } else {
                uint8_t signA = aNum.sign();
                uint8_t signB = bNum.sign();

                if (signA != signB) {
                    stack[1] = signA == 1 ? 0 : 1;
                } else {
                    stack[1] = aNum > bNum ? 1 : 0;
                }
            }
            stack.popClear();
            break;
        }
        case OpCode::EQ: {
            stack.prepForMod(2);
            auto& aVal = stack[0];
            auto& bVal = stack[1];
            stack[1] = aVal == bVal ? 1 : 0;
            stack.popClear();
            break;
        }
        case OpCode::ISZERO: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            stack[0] = (aNum == 0) ? 1 : 0;
            break;
        }
        case OpCode::BITWISE_AND: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            stack[1] = aNum & bNum;
            stack.popClear();
            break;
        }
        case OpCode::BITWISE_OR: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            stack[1] = aNum | bNum;
            stack.popClear();
            break;
        }
        case OpCode::BITWISE_XOR: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            stack[1] = aNum ^ bNum;
            stack.popClear();
            break;
        }
        case OpCode::BITWISE_NOT: {
            stack.prepForMod(1);
            auto& aNum = assumeInt(stack[0]);
            stack[0] = ~aNum;
            break;
        }
        case OpCode::BYTE: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);

            if (aNum >= 32) {
                stack[1] = 0;
            } else {
                const auto shift = 256 - 8 - 8 * shrink<uint8_t>(aNum);
                const auto mask = uint256_t(255) << shift;
                stack[1] = (bNum & mask) >> shift;
            }
            stack.popClear();
            break;
        }
        case OpCode::SIGNEXTEND: {
            stack.prepForMod(2);
            auto& aNum = assumeInt(stack[0]);
            auto& bNum = assumeInt(stack[1]);
            if (aNum >= 32) {
                stack[1] = stack[0];
            } else {
                const uint8_t idx = 8 * shrink<uint8_t>(aNum) + 7;
                const auto sign = static_cast<uint8_t>((bNum >> idx) & 1);
                const auto mask = uint256_t(-1) >> (256 - idx);
                stack[1] = uint256_t{-sign} << idx | (bNum & mask);
            }
            stack.popClear();
            break;
        }

        /***********************/
        /*  Hashing Operations */
        /***********************/
        case OpCode::HASH: {
            stack.prepForMod(1);
            stack[0] = ::hash(stack[0]);
            break;
        }

        /***********************************************/
        /*  Stack, Memory, Storage and Flow Operations */
        /***********************************************/
        case OpCode::POP: {
            stack.popClear();
            break;
        }
        case OpCode::SPUSH: {
            value copiedStatic = staticVal;
            stack.push(std::move(copiedStatic));
            break;
        }
        case OpCode::RPUSH: {
            value copiedRegister = registerVal;
            stack.push(std::move(copiedRegister));
            break;
        }
        case OpCode::RSET: {
            stack.prepForMod(1);
            registerVal = stack[0];
            stack.popClear();
            break;
        }
        case OpCode::JUMP: {
            stack.prepForMod(1);
            auto target = mpark::get_if<CodePoint>(&stack[0]);
            if (target) {
                pc = target->pc;
                shouldIncrement = false;
            } else {
                state = ERROR;
            }
            stack.popClear();
            break;
        }
        case OpCode::CJUMP: {
            stack.prepForMod(2);
            auto target = mpark::get_if<CodePoint>(&stack[0]);
            auto& bNum = assumeInt(stack[1]);
            if (bNum != 0) {
                if (target) {
                    pc = target->pc;
                    shouldIncrement = false;
                } else {
                    state = ERROR;
                }
            } else {
                shouldIncrement = true;
            }
            stack.popClear();
            stack.popClear();
            break;
        }
        case OpCode::STACKEMPTY: {
            if (stack.stacksize() == 0) {
                stack.push(1);
            } else {
                stack.push(0);
            }
            break;
        }
        case OpCode::PCPUSH: {
            stack.push(code[pc]);
            break;
        }
        case OpCode::AUXPUSH: {
            stack.prepForMod(1);
            auxstack.push(std::move(stack[0]));
            stack.popClear();
            break;
        }
        case OpCode::AUXPOP: {
            auxstack.prepForMod(1);
            stack.push(std::move(auxstack[0]));
            auxstack.popClear();
            break;
        }
        case OpCode::AUXSTACKEMPTY: {
            if (auxstack.stacksize() == 0) {
                stack.push(1);
            } else {
                stack.push(0);
            }
            break;
        }
        case OpCode::NOP: {
            break;
        }
        case OpCode::ERRPUSH: {
            stack.push(errpc);
            break;
        }
        case OpCode::ERRSET: {
            stack.prepForMod(1);
            auto codePointVal = mpark::get_if<CodePoint>(&stack[0]);
            if (!codePointVal) {
                state = ERROR;
            } else {
                errpc = *codePointVal;
            }
            stack.popClear();
            break;
        }
            /****************************************/
            /*  Duplication and Exchange Operations */
            /****************************************/
        case OpCode::DUP0: {
            value valACopy = stack[0];
            stack.push(std::move(valACopy));
            break;
        }
        case OpCode::DUP1: {
            value valBCopy = stack[1];
            stack.push(std::move(valBCopy));
            break;
        }
        case OpCode::DUP2: {
            value valCCopy = stack[2];
            stack.push(std::move(valCCopy));
            break;
        }
        case OpCode::SWAP1: {
            stack.prepForMod(2);
            value temp = stack[0];
            stack[0] = stack[1];
            stack[1] = temp;
            break;
        }
        case OpCode::SWAP2: {
            stack.prepForMod(3);
            value temp = stack[0];
            stack[0] = stack[2];
            stack[2] = temp;
            break;
        }
            /*********************/
            /*  Tuple Operations */
            /*********************/
        case OpCode::TGET: {
            stack.prepForMod(2);
            auto& index = assumeInt(stack[0]);
            auto& tup = assumeTuple(stack[1]);
            stack[1] = tup.get_element(static_cast<uint32_t>(index));
            stack.popClear();
            break;
        }
        case OpCode::TSET: {
            stack.prepForMod(3);
            auto& index = assumeInt(stack[0]);
            auto& tup = assumeTuple(stack[1]);
            tup.set_element(static_cast<uint32_t>(index), std::move(stack[2]));
            stack[2] = std::move(tup);
            stack.popClear();
            stack.popClear();
            break;
        }
        case OpCode::TLEN: {
            stack.prepForMod(1);
            stack[0] = assumeTuple(stack[0]).tuple_size();
            break;
        }
            /***********************/
            /*  Logging Operations */
            /***********************/
        case OpCode::BREAKPOINT: {
            state = HALTED;
            break;
        }
        case OpCode::LOG: {
            stack.prepForMod(1);
            value val = stack[0];
            stack.popClear();
            break;
        }
        case OpCode::DEBUG: {
            datastack tmpstk;
            std::cout << std::endl;
            std::cout << "full stack - size=" << stack.stacksize() << std::endl;
            while (stack.stacksize() > 0) {
                std::cout << stack[0] << std::endl;
                tmpstk.push(std::move(stack[0]));
                stack.popClear();
            }
            while (tmpstk.stacksize() > 0) {
                stack.push(std::move(tmpstk[0]));
                tmpstk.popClear();
            }
            std::cout << "register val=" << registerVal << std::endl
                      << std::endl;
            break;
        }
            /**********************/
            /*  System Operations */
            /**********************/
        case OpCode::SEND:{
            stack.prepForMod(1);
            auto &tup = assumeTuple(stack[0]);
            if (tup.tuple_size() != 4){
                    state=ERROR;
                    break;
                }
            Message outMsg;
            outMsg.data =tup.get_element(0);
            value dest =tup.get_element(1);
            outMsg.destination = assumeInt(dest);
            value amt =tup.get_element(2);
            outMsg.currency = assumeInt(amt);
            value tokVal =tup.get_element(3);
            auto &tokTypeVal = assumeInt(tokVal);

            toTokenType(tokTypeVal, outMsg.token);

            if (!context.afterBalance.Spend(outMsg.token, outMsg.currency)){
                state=BLOCKED;
            } else {
                stack.popClear();
                context.outMessage.push_back(outMsg);
            }
            break;
        }
        case OpCode::NBSEND:{
            stack.prepForMod(1);
            auto &tup = assumeTuple(stack[0]);
            if (tup.tuple_size() != 4){
                state=ERROR;
                break;
            }
            Message outMsg;
            outMsg.data =tup.get_element(0);
            value dest =tup.get_element(1);
            outMsg.destination = assumeInt(dest);
            value amt =tup.get_element(2);
            outMsg.currency = assumeInt(amt);
            value tokVal =tup.get_element(3);
            auto &tokTypeVal = assumeInt(tokVal);

            toTokenType(tokTypeVal, outMsg.token);

            if (!context.afterBalance.CanSpend(outMsg.token, outMsg.currency)){
                stack[0] = 0;
            } else {
                bool spent = context.afterBalance.Spend(outMsg.token, outMsg.currency);
                assert(spent);
                context.outMessage.push_back(outMsg);
                stack[0] = 1;
            }
            break;
        }
        case OpCode::GETTIME:{
            Tuple tup(2, pool.get());
            tup.set_element(0, context.precondition.timeBounds[0]);
            tup.set_element(1, context.precondition.timeBounds[1]);
            stack.push(std::move(tup));
            break;
        }
        case OpCode::INBOX: {
            stack.prepForMod(1);
            if (inbox == stack[0]) {
                state = BLOCKED;
                shouldIncrement = false;
            } else {
                value inboxCopy = inbox;
                stack[0] = std::move(inboxCopy);
            }
            break;
        }
        case OpCode::ERROR:
            //TODO: add error handler support
            state=ERROR;
            break;
        case OpCode::HALT:
            std::cout << "Hit halt opcode at instruction " << pc << "\n";
            state=HALTED;
            break;
        default:
            std::stringstream ss;
            ss << "Unhandled opcode <" << InstructionNames.at(opcode) << ">"
               << std::hex << static_cast<int>(opcode);
            throw std::runtime_error(ss.str());
    }
    if (shouldIncrement) {
        ++pc;
    }
}

/***********************************/
// test code
// void push_num(vector<instr> &code, unsigned long long &pc, value *tpl, value
// *tmp, uint256_t num){
//    instr *op;
//    //push(1)
//    pc++;
//    *tmp = num;
//    op = new instr(pc,NOP,tmp);
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//}
// void push_tuple(vector<instr> &code, unsigned long long &pc, int size, value
// *tpl, value *tmp){
//    instr *op;
//
//    if (size==5){
//        tpl->set_tuple_elem(0, (uint256_t)11);
//        tpl->set_tuple_elem(1, (uint256_t)12);
//        tpl->set_tuple_elem(2, (uint256_t)13);
//        tpl->set_tuple_elem(3, (uint256_t)14);
//        tpl->set_tuple_elem(4, (uint256_t)15);
//    } else {
//        tpl->set_tuple_elem(0, (uint256_t)21);
//        tpl->set_tuple_elem(1, (uint256_t)22);
//        tpl->set_tuple_elem(2, (uint256_t)23);
//    }
//    //push Tuple
//    pc++;
//    op = new instr(pc,NOP,tpl);
//    code.push_back(*op);
//    delete op;
//    // print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//}
// void print_stack(vector<instr> &code, unsigned long long &pc, value *tpl,
// value *tmp){
//    instr *op;
//    pc++;
//    op = new instr(pc,PRTSTK,NULL);
//    code.push_back(*op);
//}
//
// void rset(vector<instr> &code, unsigned long long &pc, value *tpl, value
// *tmp){
//    instr *op;
//    //rset
//    tmp->set_num((uint256_t)31);
//    op = new instr(pc,RSET,tmp);
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//    pc++;
//}
//
// void test_pop(vector<instr> &code, unsigned long long &pc, value *tpl, value
// *tmp){
//    instr *op;
//    pc++;
//    op = new instr(pc,POP,NULL);
//    code.push_back(*op);
//}
//
// void test_tget( vector<instr> &code, unsigned long long &pc, value *tpl,
// value *tmp){
//    instr *op;
//
//    //test tget
//    push_tuple( code, pc, 5, tpl, tmp);
//    push_num( code, pc, tpl, tmp, (uint256_t)2);
//
//    // tget()
//    pc++;
//    op = new instr(pc,TGET,NULL);
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//}
//
// void test_add( vector<instr> &code, unsigned long long &pc, value *tpl, value
// *tmp){
//    instr *op;
//    //test add
//    //push(10)
//    push_num( code, pc, tpl, tmp, (uint256_t)10);
//    push_num( code, pc, tpl, tmp, (uint256_t)20);
//    op = new instr(pc,ADD,NULL);
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//}
//
// void test_tset( vector<instr> &code, unsigned long long &pc, value *tpl,
// value *tpl2, value *tmp){
//    instr *op;
//    //test tset
//    //push(10)
//    //    push_num( code, pc, tpl, tmp, (uint256_t)10);
//    push_tuple( code, pc, 5, tpl, tmp);
//    push_tuple( code, pc, 3, tpl2, tmp);
//    push_num( code, pc, tpl, tmp, (uint256_t)1);
//    op = new instr(pc,TSET,NULL);
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//    pc++;
//    op = new instr(pc,PRTSTK,NULL);
//    code.push_back(*op);
//}
// void test_pcpush( vector<instr> &code, unsigned long long &pc, value *tpl,
// value *tmp){
//    instr *op;
//    //test pcpush
//    //pcpush
//    pc++;
//    op = new instr(pc,PCPUSH,NULL);
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//    //rset
//    pc++;
//    op = new instr(pc,RSET,NULL);
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//}
// void test_jump( vector<instr> &code, unsigned long long &pc, value *tpl,
// value *tmp){
//    instr *op;
//    //test jump
//    //    rset( code, pc, tpl, tmp);
//    op = new instr(pc,RPUSH,NULL); //rpush
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL); //print
//    code.push_back(*op);
//    pc++;
//    op = new instr(pc,JUMP,NULL); //jmp
//    //    op = new instr(pc,NOP,NULL); //jmp
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//}
//
// void test_mul( vector<instr> &code, unsigned long long &pc, value *tpl, value
// *tmp){
//    instr *op;
//    //test mul
//    push_num( code, pc, tpl, tmp, (uint256_t)10);
//    push_num( code, pc, tpl, tmp, (uint256_t)20);
//
//    op = new instr(pc,MUL,NULL);
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//
//}
//
// void setupCode( vector<instr> &code){
//    unsigned long long pc=0;
//    instr *op;
//
//    value *tpl=new value((int)5);
//    value *tpl3=new value((int)3);
//    //print stack
//    pc++;
//    op = new instr(pc,PRTSTK,NULL);
//    code.push_back(*op);
//    value *tmp=new value;
//
//    push_tuple(code, pc, 3, tpl3, tmp);
//    test_tget(code, pc, tpl, tmp);
//    print_stack(code, pc, tpl, tmp);
//    //    test_tget( code, pc, tpl, tmp);
//    //    print_stack(code, pc, tpl, tmp);
//    //    test_add( code, pc, tpl, tmp);
//    test_pcpush( code, pc, tpl, tmp);
//    test_tset( code, pc, tpl, tpl3, tmp);
//    print_stack(code, pc, tpl, tmp);
//    test_pop( code, pc, tpl, tmp);
//
//    test_jump( code, pc, tpl, tmp);
//    //    test_mul( code, pc, tpl, tmp);
//    //    test_pcpush( code, pc, tpl, tmp);
//
//}
/***********************************/
