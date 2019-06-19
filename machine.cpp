//
//  machine.cpp
//  AVMtest
//
//  Created by Harry Kalodner on 4/2/19.
//

#include "machine.hpp"
#include "code.hpp"
#include "opcodes.hpp"

#include <keccak/KeccakHash.h>

#include <iostream>

class bad_pop_type : public std::exception {
public:
    virtual const char *what() const noexcept override { return "bad_variant_access"; }
};

class int_out_of_bounds : public std::exception {
public:
    virtual const char *what() const noexcept override { return "int_out_of_bounds"; }
};

uint256_t &assumeInt(value &val) {
    auto aNum = mpark::get_if<uint256_t>(&val);
    if (!aNum) {
        throw bad_pop_type{};
    }
    return *aNum;
}

uint64_t assumeInt64(uint256_t &val) {
    if (val > std::numeric_limits<uint64_t>::max())
        throw int_out_of_bounds{};
    
    return static_cast<uint64_t>(val);
}

Tuple &assumeTuple(value &val) {
    auto tup = mpark::get_if<Tuple>(&val);
    if (!tup) {
        throw bad_pop_type{};
    }
    return *tup;
}

instr deserialize_opcode(uint64_t pc, char *&bufptr, TuplePool &pool){
    uint8_t immediateCount;
    memcpy(&immediateCount, bufptr, sizeof(immediateCount));
    bufptr+=sizeof(immediateCount);

    OpCode opcode;
    memcpy(&opcode, bufptr, sizeof(opcode));
    bufptr+=sizeof(opcode);

    if (immediateCount == 0x01) {
        return instr(pc, opcode, 0, deserialize_value(bufptr, pool));
    } else {
        return instr(pc, opcode, 0);
    }
}

Machine::Machine() : pool(std::make_unique<TuplePool>()) {}

Machine::Machine(char *&srccode, char *&inboxdata) : Machine() {
    char *bufptr=srccode;
    char *inboxbuf=inboxdata;

    uint32_t version;
    memcpy(&version, bufptr, sizeof(version));
    version = __builtin_bswap32(version);
    bufptr += sizeof(version);

    if (version != CURRENT_AO_VERSION){
        std::cout << "incorrect version of .ao file"<<std::endl;
        std::cout<<"expected version "<<CURRENT_AO_VERSION<<" found version "<< version<<std::endl;
        return;
    }

    uint32_t extentionId = 1;
    while (extentionId !=0) {
        memcpy(&extentionId, bufptr, sizeof(extentionId));
        extentionId = __builtin_bswap32(extentionId);
        bufptr += sizeof(extentionId);
        if (extentionId>0){
            std::cout<<"found exetention"<<std::endl;
        }
    }
    uint64_t codeCount;
    memcpy(&codeCount, bufptr, sizeof(codeCount));
    bufptr += sizeof(codeCount);
    codeCount = __builtin_bswap64(codeCount);
    code.reserve(codeCount);
    std::cout<<"codeCount="<<codeCount<<std::endl;

    for (uint64_t i = 0; i < codeCount; i++){
        code.push_back(deserialize_opcode(i, bufptr, *pool));
    }
    std::cout<<"code read"<<std::endl;
    staticVal = deserialize_value(bufptr, *pool);
    std::cout<<"static read"<<std::endl;
    pc=0;
    if (inboxbuf){
        std::cout<<"reading inbox"<<std::endl;
        inbox = deserialize_value(inboxbuf, *pool);
        std::cout<<"inbox value="<<inbox<<std::endl;
    }
}
Assertion Machine::run(uint64_t stepCount) {
    //testing opcodes
//    opcodeTests();
//    std::cout << "starting machine code size=" << code.size() << std::endl;
//    std::cout << "inbox=" << inbox << std::endl;
    uint64_t i;
    for (i = 0; i < stepCount; i++) {
//        std::cout << "Step #" << i << std::endl;
        auto ret = runOne();
        if (ret < 0) {
            break;
        } else if (state==ERROR) {
            break;
        } else if (state==HALTED) {
            break;
        }
    }
    
    if (state==ERROR){
        //set error return
        std::cout << "error state" << std::endl;
    }
    if (state==HALTED){
        //set error return
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
//    std::cout<<"in runOne"<<std::endl;
    if (state == ERROR){
        //set error return
        std::cout<<"error state"<<std::endl;
        return -1;
    }
    
    if (state==HALTED){
        //set error return
        std::cout<<"halted state"<<std::endl;
        std::cout<<"full stack - size="<< stack.stacksize()<<std::endl;
        while (stack.stacksize()>0){
            value A=stack.pop();
            std::cout << A << std::endl;
        }
        return -2;
    }
//    std::cout<<"pc="<<pc<<std::endl;
    auto &instruction = code[pc];
    auto immediateVal = instruction.immediate;
    if (immediateVal){
//        std::cout<<"immediateVal = "<<*immediateVal<<std::endl;
        stack.push(std::move(*immediateVal));
    }
    
    try {
//        std::cout<<"calling runInstruction"<<std::endl;
        runInstruction(instruction);
//        std::cout<<"after runInstruction stack size= "<<stack.stacksize()<< std::endl;
//        if (stack.stacksize()>0){
//            std::cout<<"top="<<stack.peek()<< std::endl;
//        }
    } catch (const bad_pop_type &e) {
        state = ERROR;
    } catch (const bad_tuple_index &e) {
        state = ERROR;
    }
    
    return 0;
}

template <typename T>
static T shrink(uint256_t i)
{
    return static_cast<T>(i & std::numeric_limits<T>::max());
}
void Machine::opcodeTests(){
    //TEST SUB
    {
        std::cout<<"TESTING sub "<<std::endl;
        instr inst(1, OpCode::SUB, 0);
        uint256_t val1 = 3;
        stack.push(std::move(val1));
        uint256_t val2= 4;
        stack.push(std::move(val2));
        runInstruction(inst);
        value res = stack.pop();
        auto &resNum = assumeInt(res);
        if (resNum != 1){
            std::cout<<"ERROR - sub failed - expected 1 received "<<resNum<<std::endl;
        }
    }
    //TEST SUB -1
    {
        std::cout<<"TESTING sub -1 "<<std::endl;
        instr inst(1, OpCode::SUB, 0);
        uint256_t val1 = 4;
        stack.push(std::move(val1));
        uint256_t val2= 3;
        stack.push(std::move(val2));
        runInstruction(inst);
        value res = stack.pop();
        auto &resNum = assumeInt(res);
        uint256_t expect=(uint256_t)(-1);
        if (resNum != expect){
            std::cout<<"ERROR - sub failed - expected "<<expect<<" received "<<resNum<<std::endl;
        }
    }
    //TEST DIV
    {
        std::cout<<"TESTING div"<<std::endl;
        instr inst(1, OpCode::DIV, 0);
        uint256_t val1 = 3;
        stack.push(std::move(val1));
        uint256_t val2= 12;
        stack.push(std::move(val2));
        runInstruction(inst);
        value res = stack.pop();
        auto &resNum = assumeInt(res);
        if (resNum != 4){
            std::cout<<"ERROR - div failed - expected 4 received "<<resNum<<std::endl;
        }
    }
    //TEST DIV divide by 0
    {
        std::cout<<"TESTING div"<<std::endl;
        instr inst(1, OpCode::DIV, 0);
        uint256_t val1 = 0;
        stack.push(std::move(val1));
        uint256_t val2= 3;
        stack.push(std::move(val2));
        runInstruction(inst);
        if (state!=ERROR){
            std::cout<<"ERROR - div failed - expected ERROR state"<<std::endl;
        }
    }
    //TEST SDIV
    {
        std::cout<<"TESTING sdiv"<<std::endl;
        instr inst(1, OpCode::SDIV, 0);
        uint256_t val1 = (uint256_t)-3;
        stack.push(std::move(val1));
        uint256_t val2= 12;
        stack.push(std::move(val2));
        runInstruction(inst);
        value res = stack.pop();
        auto &resNum = assumeInt(res);
        if (resNum != (uint256_t)-4){
            std::cout<<"ERROR - div failed - expected -4 received "<<resNum<<std::endl;
        }
    }
    //TEST SDIV
    {
        std::cout<<"TESTING sdiv"<<std::endl;
        instr inst(1, OpCode::SDIV, 0);
        uint256_t val1 = (uint256_t)-3;
        stack.push(std::move(val1));
        uint256_t val2= (uint256_t)-12;
        stack.push(std::move(val2));
        runInstruction(inst);
        value res = stack.pop();
        auto &resNum = assumeInt(res);
        if (resNum != 4){
            std::cout<<"ERROR - div failed - expected 4 received "<<resNum<<std::endl;
        }
    }
    //TEST MOD
    {
        std::cout<<"TESTING smod"<<std::endl;
        instr inst(1, OpCode::SMOD, 0);
        uint256_t val1 = 3;
        stack.push(std::move(val1));
        uint256_t val2 = (uint256_t)-8;
        stack.push(std::move(val2));
        runInstruction(inst);
        value res = stack.pop();
        auto &resNum = assumeInt(res);
        if (resNum != (uint256_t)-2){
            std::cout<<"ERROR - smod failed - expected -2 received "<<resNum<<std::endl;
        }
    }



}

void Machine::runInstruction( instr instruction ) {
//void Machine::runInstruction( ) {
//    auto &instruction = testInstr;
//    auto &instruction = code[pc];
//    std::stringstream ss;
//    ss << "in runInstruction, running " << std::hex << static_cast<int>(instruction.opcode);
//    std::cout << ss.str() <<", <"<< InstructionNames.at(instruction.opcode) <<">, stack size= "<<stack.stacksize()<< "\n";
//    if (stack.stacksize()>0){
//        std::cout<<"top="<<stack.peek()<< std::endl;
//    }
    bool shouldIncrement = true;
    switch (instruction.opcode) {
        /**************************/
        /*  Arithmetic Operations */
        /**************************/
        case OpCode::HALT:
            std::cout << "Hit halt opcode at instruction " << pc << "\n";
            state=HALTED;
            break;
        case OpCode::ADD: {
            value val1 = stack.pop();
            value &val2 = stack.peek();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            bNum += aNum;
            break;
        }
        case OpCode::MUL: {
            value val1 = stack.pop();
            value &val2 = stack.peek();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            bNum *= aNum;
            break;
        }
        case OpCode::SUB: {
            value val1 = stack.pop();
            value val2 = stack.pop();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            uint256_t ret;
            ret=aNum-bNum;
            stack.push(std::move(ret));
            break;
        }
        case OpCode::DIV: {
            value val1 = stack.pop();
            value val2 = stack.pop();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            uint256_t ret;
            if (bNum == 0){
                ret=0;
                state = ERROR;
            } else {
                ret=aNum/bNum;
            }
            stack.push(std::move(ret));
            break;
        }
        case OpCode::SDIV: {
            value val1 = stack.pop();
            value val2 = stack.pop();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            const auto min = (std::numeric_limits<uint256_t>::max() / 2) + 1;

            if (bNum == 0) {
                state=ERROR;
            } else if (aNum == min && bNum == -1) {
                bNum = aNum;
            } else {
                const auto signA = get_sign(aNum);
                const auto signB = get_sign(bNum);
                if (signA == -1)
                    aNum = 0 - aNum;
                if (signB == -1)
                    bNum = 0 - bNum;

                bNum = (aNum / bNum) * signA * signB;
            }
            stack.push(std::move(bNum));
            break;
        }
        case OpCode::MOD: {
            value val1 = stack.pop();
            value &val2 = stack.peek();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            if (bNum != 0) {
                bNum = aNum % bNum;
            } else {
                bNum = 0;
            }
            break;
        }
        case OpCode::SMOD: {
            value val1 = stack.pop();
            value val2 = stack.pop();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            
            if (bNum == 0) {
                state=ERROR;
            } else {
                const auto signA = get_sign(aNum);
                const auto signB = get_sign(bNum);
                if (signA == -1)
                    aNum = 0 - aNum;
                if (signB == -1)
                    bNum = 0 - bNum;
                
                bNum = (aNum % bNum) * signA;
            }
            stack.push(std::move(bNum));
            break;
        }
        case OpCode::ADDMOD: {
            value val1 = stack.pop();
            value val2 = stack.pop();
            value &val3 = stack.peek();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            auto &cNum = assumeInt(val3);
            
            if (cNum == 0) {
                cNum = 0;
            } else {
                uint512_t aBig = aNum;
                uint512_t bBig = bNum;
                
                cNum = static_cast<uint256_t>((aBig + bBig) % cNum);
            }
            break;
        }
        case OpCode::MULMOD: {
            value val1 = stack.pop();
            value val2 = stack.pop();
            value &val3 = stack.peek();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            auto &cNum = assumeInt(val3);
            
            if (cNum == 0) {
                cNum = 0;
            } else {
                uint512_t aBig = aNum;
                uint512_t bBig = bNum;
                
                cNum = static_cast<uint256_t>((aBig * bBig) % cNum);
            }
            break;
        }
        case OpCode::EXP: {
            value val1 = stack.pop();
            value &val2 = stack.peek();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            uint64_t bSmall = assumeInt64(bNum);
            
            bNum = power(aNum, bSmall);
           break;
        }
        /******************************************/
        /*  Comparison & Bitwise Logic Operations */
        /******************************************/
        case OpCode::LT: {
            value val1 = stack.pop();
            value &val2 = stack.peek();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            bNum = (aNum < bNum) ? 1 : 0;
           break;
        }
        case OpCode::GT: {
            value val1 = stack.pop();
            value &val2 = stack.peek();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            bNum = (aNum > bNum) ? 1 : 0;
            break;
        }
        case OpCode::SLT: {
            value val1 = stack.pop();
            value &val2 = stack.peek();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            if (aNum == bNum) {
                bNum = 0;
            } else {
                uint8_t signA = aNum.sign();
                uint8_t signB = bNum.sign();
                
                if (signA != signB) {
                    if (signA == 1) {
                        bNum = 1;
                    } else {
                        bNum = 0;
                    }
                } else {
                    bNum = (aNum < bNum) ? 1 : 0;
                }
            }
            break;
        }
        case OpCode::SGT: {
            value val1 = stack.pop();
            value &val2 = stack.peek();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            if (aNum == bNum) {
                bNum = 0;
            } else {
                uint8_t signA = aNum.sign();
                uint8_t signB = bNum.sign();
                
                if (signA != signB) {
                    if (signA == 1) {
                        bNum = 0;
                    } else {
                        bNum = 1;
                    }
                } else {
                    bNum = (aNum > bNum) ? 1 : 0;
                }
            }
            break;
        }
        case OpCode::EQ: {
            value val1 = stack.pop();
            value val2 = stack.pop();
            uint256_t ret;
            if (val1 == val2){
                ret=1;
            } else {
                ret=0;
            }
//            std::cout << "in eq val1="<<val1<<" val2="<<val2<<" ret="<<ret <<std::endl;
            stack.push(std::move(ret));
            break;
        }
        case OpCode::ISZERO: {
            value &val1 = stack.peek();
            auto &aNum = assumeInt(val1);
            aNum = (aNum == 0) ? 1 : 0;
            break;
        }
        case OpCode::BITWISE_AND: {
            value val1 = stack.pop();
            value &val2 = stack.peek();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            bNum &= aNum;
            break;
        }
        case OpCode::BITWISE_OR: {
            value val1 = stack.pop();
            value &val2 = stack.peek();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
//            std::cout << "in or aNum="<<aNum<<" bNum="<<bNum <<std::endl;
            bNum |= aNum;
            break;
        }
        case OpCode::BITWISE_XOR: {
            value val1 = stack.pop();
            value &val2 = stack.peek();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            bNum ^= aNum;
            break;
        }
        case OpCode::BITWISE_NOT: {
            value &val1 = stack.peek();
            auto &aNum = assumeInt(val1);
            aNum = ~aNum;
            break;
        }
        case OpCode::BYTE: {
            value val1 = stack.pop();
            value &val2 = stack.peek();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            
            if (aNum >= 32) {
                bNum = 0;
            } else {
                const auto shift = 256 - 8 - 8 * shrink<uint8_t>(aNum);
                const auto mask = uint256_t(255) << shift;
                bNum = (bNum & mask) >> shift;
            }
            break;
        }
        case OpCode::SIGNEXTEND: {
            value val1 = stack.pop();
            value &val2 = stack.peek();
            auto &aNum = assumeInt(val1);
            auto &bNum = assumeInt(val2);
            if (aNum >= 32)
            {
                bNum = bNum;
            } else {
                const uint8_t idx = 8 * shrink<uint8_t>(aNum) + 7;
                const auto sign = static_cast<uint8_t>((bNum >> idx) & 1);
                const auto mask = uint256_t(-1) >> (256 - idx);
                bNum = (uint256_t(-sign) << idx) | (bNum & mask);
            }
            break;
        }
            
        /***********************/
        /*  Hashing Operations */
        /***********************/
        case OpCode::HASH: {
            value &val1 = stack.peek();
            val1 = value_hash(val1);
            break;
        }
            
        /***********************************************/
        /*  Stack, Memory, Storage and Flow Operations */
        /***********************************************/
        case OpCode::POP:
//            std::cout << "in Pop" <<std::endl;
            stack.pop();
            break;
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
        case OpCode::RSET:
//            std::cout << "in RSET" << std::endl;
            stack.popSet(registerVal);
//            std::cout << "register set " << registerVal << std::endl;
            break;
        case OpCode::JUMP:{
            auto jumpDest = stack.pop();
            auto target = mpark::get_if<CodePoint>(&jumpDest);
            if (target) {
                pc = target->pc;
                shouldIncrement = false;
//                std::cout << "jumping to "<< pc <<std::endl;
            } else {
                state = ERROR;
            }
            break;
        }
        case OpCode::CJUMP:{
            value jumpDest = stack.pop();
            value val1 = stack.pop();
            auto target = mpark::get_if<CodePoint>(&jumpDest);
            auto &bNum = assumeInt(val1);
            if (bNum != 0){
                if (target) {
                    pc = target->pc;
                    shouldIncrement = false;
                } else {
                    state = ERROR;
                }
            } else{
                shouldIncrement=true;
            }
            break;
        }
        case OpCode::STACKEMPTY:{
            uint256_t ret;
            if (stack.stacksize() == 0){
                ret=1;
            } else {
                ret=0;
            }
            stack.push(std::move(ret));
            break;
        }
        case OpCode::PCPUSH:{
//            std::cout << "**** PCPUSH i=" << pc <<std::endl;
            stack.push(CodePoint{pc});
            break;
        }
        case OpCode::AUXPUSH:{
            value val1 = stack.pop();
            auxstack.push(std::move(val1));
            break;
        }
        case OpCode::AUXPOP:{
            value val1 = auxstack.pop();
            stack.push(std::move(val1));
            break;
        }
        case OpCode::AUXSTACKEMPTY:{
            uint256_t ret;
            if (auxstack.stacksize() == 0){
                ret=1;
            } else {
                ret=0;
            }
            stack.push(std::move(ret));
            break;
        }
        case OpCode::NOP:
            //nop
            break;
        case OpCode::ERRPUSH:
            stack.push(std::move(errpc));
            break;
        case OpCode::ERRSET:{
            value ret = stack.pop();
            auto errpc = mpark::get_if<CodePoint>(&ret);
            break;
        }
    /****************************************/
    /*  Duplication and Exchange Operations */
    /****************************************/
        case OpCode::DUP0:{
            value valA = stack.peek();
            stack.push(std::move(valA));
            break;
        }
        case OpCode::DUP1:{
            value valA = stack.pop();
            value valB = stack.peek();
            stack.push(std::move(valA));
            stack.push(std::move(valB));
            break;
        }
        case OpCode::DUP2:{
            value valA = stack.pop();
            value valB = stack.pop();
            value valC = stack.peek();
            stack.push(std::move(valB));
            stack.push(std::move(valA));
            stack.push(std::move(valC));
            break;
        }
        case OpCode::SWAP1:{
            value valA = stack.pop();
            value valB = stack.pop();
            stack.push(std::move(valA));
            stack.push(std::move(valB));
            break;
        }
        case OpCode::SWAP2:{
            value valA = stack.pop();
            value valB = stack.pop();
            value valC = stack.pop();
            stack.push(std::move(valA));
            stack.push(std::move(valB));
            stack.push(std::move(valC));
            
            break;
        }
    /*********************/
    /*  Tuple Operations */
    /*********************/
        case OpCode::TGET: {
            auto indexVal = stack.pop();
            auto tupVal = stack.pop();
            auto &index = assumeInt(indexVal);
            auto &tup = assumeTuple(tupVal);
            stack.push(tup.get_element(static_cast<uint32_t>(index)));
            break;
        }
        case OpCode::TSET: {
            auto indexVal = stack.pop(); // slot
            auto tupVal = stack.pop(); // tuple
            auto newVal = stack.pop(); // val
            auto &index = assumeInt(indexVal);
            auto &tup = assumeTuple(tupVal);
            tup.set_element(static_cast<uint32_t>(index), std::move(newVal));
            stack.push(std::move(tup));
            break;
        }
        case OpCode::TLEN: {
            auto tupVal = stack.pop(); // tuple
            auto &tup = assumeTuple(tupVal);
            uint256_t size = tup.tuple_size();
            stack.push(std::move(size));
            break;
        }
    /***********************/
    /*  Logging Operations */
    /***********************/
        case OpCode::BREAKPOINT:{
            state=HALTED;
            break;
        }
        case OpCode::LOG:{
            value val = stack.pop();
//            std::cout << "log val=" << val << std::endl << std::endl;
            break;
        }
        case OpCode::DEBUG: {
            datastack tmpstk;
            std::cout<<std::endl;
            std::cout<<"full stack - size="<<stack.stacksize()<<std::endl;
            while (stack.stacksize()>0){
                value A = stack.pop();
                std::cout << A << std::endl;
                tmpstk.push(std::move(A));
            }
            while (tmpstk.stacksize()>0){
                value A = tmpstk.pop();
                stack.push(std::move(A));
            }
            std::cout << "register val=" << registerVal << std::endl << std::endl;
            break;
        }
    /**********************/
    /*  System Operations */
    /**********************/
//        case OpCode::SEND:
//            break;
//        case OpCode::NBSEND:
//            break;
//        case OpCode::GETTIME:
//            break;
        case OpCode::INBOX:{
            value val=stack.pop();
//            std::cout<<"In inbox. val="<<val<<" inbox="<<inbox<<std::endl;
            if (inbox == val){
                state=HALTED;
            } else {
                stack.push(std::move(inbox));
            }
            break;
        }
//        case OpCode::ERROR:
//            break;
//        case OpCode::HALT:
//            break;
        default:
            std::stringstream ss;
            ss << "Unhandled opcode <"<< InstructionNames.at(instruction.opcode) <<">" << std::hex << static_cast<int>(instruction.opcode);
            throw std::runtime_error(ss.str());
    }
    if (shouldIncrement) {
        ++pc;
    }
}

/***********************************/
// test code
//void push_num(vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp, uint256_t num){
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
//void push_tuple(vector<instr> &code, unsigned long long &pc, int size, value *tpl, value *tmp){
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
//void print_stack(vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
//    instr *op;
//    pc++;
//    op = new instr(pc,PRTSTK,NULL);
//    code.push_back(*op);
//}
//
//void rset(vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
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
//void test_pop(vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
//    instr *op;
//    pc++;
//    op = new instr(pc,POP,NULL);
//    code.push_back(*op);
//}
//
//void test_tget( vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
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
//void test_add( vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
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
//void test_tset( vector<instr> &code, unsigned long long &pc, value *tpl, value *tpl2, value *tmp){
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
//void test_pcpush( vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
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
//void test_jump( vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
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
//void test_mul( vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
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
//void setupCode( vector<instr> &code){
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
