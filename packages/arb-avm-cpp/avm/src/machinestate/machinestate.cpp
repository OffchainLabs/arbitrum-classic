/*
 * Copyright 2019-2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include <avm/machinestate/machinestate.hpp>

#include <avm/machinestate/machineoperation.hpp>
#include <avm_values/exceptions.hpp>
#include <avm_values/vmValueParser.hpp>

#include <avm_values/util.hpp>
#include <bigint_utils.hpp>

void uint256_t_to_buf(const uint256_t& val, std::vector<unsigned char>& buf) {
    std::array<unsigned char, 32> tmpbuf;
    to_big_endian(val, tmpbuf.begin());
    buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
}

std::pair<MachineState, bool> MachineState::loadFromFile(
    const std::string& contract_filename) {
    auto pool = std::make_shared<TuplePool>();
    auto ret = parseStaticVmValues(contract_filename, *pool.get());

    if (!ret.second) {
        return std::make_pair(MachineState{}, false);
    }

    return std::make_pair(MachineState{std::make_shared<const StaticVmValues>(
                                           std::move(ret.first)),
                                       std::move(pool)},
                          true);
}

uint256_t MachineState::hash() const {
    if (state == Status::Halted)
        return 0;
    if (state == Status::Error)
        return 1;

    std::array<unsigned char, 32 * 6> data;
    auto oit = data.begin();
    {
        auto val = ::hash(static_values->code[pc]);
        oit = to_big_endian(val, oit);
    }
    {
        auto val = stack.hash();
        oit = to_big_endian(val, oit);
    }
    {
        auto val = auxstack.hash();
        oit = to_big_endian(val, oit);
    }
    {
        auto val = ::hash_value(registerVal);
        oit = to_big_endian(val, oit);
    }
    {
        auto val = ::hash_value(static_values->staticVal);
        oit = to_big_endian(val, oit);
    }
    {
        auto val = ::hash(static_values->code[errpc]);
        oit = to_big_endian(val, oit);
    }

    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(data.data(), data.size(), hashData.data());
    return from_big_endian(hashData.begin(), hashData.end());
}

uint256_t MachineState::getMachineSize() {
    uint256_t machine_size = 0;

    machine_size += getSize(static_values->staticVal);
    machine_size += getSize(registerVal);
    machine_size += stack.getTotalValueSize();
    machine_size += auxstack.getTotalValueSize();

    return machine_size;
}

namespace {
std::vector<unsigned char> marshalState(const Code& code,
                                        uint256_t next_codepoint_hash,
                                        HashPreImage stackPreImage,
                                        HashPreImage auxStackPreImage,
                                        value registerVal,
                                        value staticVal,
                                        CodePointStub errpc) {
    auto buf = std::vector<unsigned char>();
    marshal_uint256_t(next_codepoint_hash, buf);

    stackPreImage.marshal(buf);
    auxStackPreImage.marshal(buf);

    ::marshalStub(registerVal, buf, code);
    ::marshalStub(staticVal, buf, code);
    marshal_uint256_t(::hash(errpc), buf);

    return buf;
}
}  // namespace

std::vector<unsigned char> MachineState::marshalState() const {
    auto stackPreImage = stack.getHashPreImage();
    auto auxStackPreImage = auxstack.getHashPreImage();

    return ::marshalState(static_values->code, ::hash(static_values->code[pc]),
                          stackPreImage, auxStackPreImage, registerVal,
                          static_values->staticVal,
                          CodePointStub{static_values->code[errpc]});
}

std::vector<unsigned char> MachineState::marshalForProof() {
    auto opcode = static_values->code[pc].op.opcode;
    std::vector<bool> stackPops = InstructionStackPops.at(opcode);
    bool includeImmediateVal = false;
    if (static_values->code[pc].op.immediate && !stackPops.empty()) {
        includeImmediateVal = stackPops[0] == true;
        stackPops.erase(stackPops.begin());
    }
    std::vector<bool> auxStackPops = InstructionAuxStackPops.at(opcode);

    auto stackProof = stack.marshalForProof(stackPops, static_values->code);
    auto auxStackProof =
        auxstack.marshalForProof(auxStackPops, static_values->code);

    auto buf = ::marshalState(
        static_values->code, static_values->code[pc].nextHash, stackProof.first,
        auxStackProof.first, registerVal, static_values->staticVal,
        CodePointStub{static_values->code[errpc]});

    static_values->code[pc].op.marshalForProof(buf, includeImmediateVal,
                                               static_values->code);

    buf.insert(buf.end(), stackProof.second.begin(), stackProof.second.end());
    buf.insert(buf.end(), auxStackProof.second.begin(),
               auxStackProof.second.end());
    return buf;
}

BlockReason MachineState::isBlocked(uint256_t currentTime,
                                    bool newMessages) const {
    if (state == Status::Error) {
        return ErrorBlocked();
    } else if (state == Status::Halted) {
        return HaltBlocked();
    }
    auto& instruction = static_values->code[pc];
    if (instruction.op.opcode == OpCode::INBOX) {
        if (newMessages) {
            return NotBlocked();
        }

        auto& immediate = instruction.op.immediate;
        const value* param;
        if (immediate) {
            param = immediate.get();
        } else {
            param = &stack[0];
        }
        auto paramNum = nonstd::get_if<uint256_t>(param);
        if (!paramNum) {
            return NotBlocked();
        }
        if (currentTime < *paramNum) {
            return InboxBlocked(*paramNum);
        } else {
            return NotBlocked();
        }
    } else {
        return NotBlocked();
    }
}

BlockReason MachineState::runOp(OpCode opcode) {
    switch (opcode) {
            /**************************/
            /*  Arithmetic Operations */
            /**************************/
        case OpCode::ADD:
            machineoperation::add(*this);
            break;
        case OpCode::MUL:
            machineoperation::mul(*this);
            break;
        case OpCode::SUB:
            machineoperation::sub(*this);
            break;
        case OpCode::DIV:
            machineoperation::div(*this);
            break;
        case OpCode::SDIV:
            machineoperation::sdiv(*this);
            break;
        case OpCode::MOD:
            machineoperation::mod(*this);
            break;
        case OpCode::SMOD:
            machineoperation::smod(*this);
            break;
        case OpCode::ADDMOD:
            machineoperation::addmod(*this);
            break;
        case OpCode::MULMOD:
            machineoperation::mulmod(*this);
            break;
        case OpCode::EXP:
            machineoperation::exp(*this);
            break;
            /******************************************/
            /*  Comparison & Bitwise Logic Operations */
            /******************************************/
        case OpCode::LT:
            machineoperation::lt(*this);
            break;
        case OpCode::GT:
            machineoperation::gt(*this);
            break;
        case OpCode::SLT:
            machineoperation::slt(*this);
            break;
        case OpCode::SGT:
            machineoperation::sgt(*this);
            break;
        case OpCode::EQ:
            machineoperation::eq(*this);
            break;
        case OpCode::ISZERO:
            machineoperation::iszero(*this);
            break;
        case OpCode::BITWISE_AND:
            machineoperation::bitwiseAnd(*this);
            break;
        case OpCode::BITWISE_OR:
            machineoperation::bitwiseOr(*this);
            break;
        case OpCode::BITWISE_XOR:
            machineoperation::bitwiseXor(*this);
            break;
        case OpCode::BITWISE_NOT:
            machineoperation::bitwiseNot(*this);
            break;
        case OpCode::BYTE:
            machineoperation::byte(*this);
            break;
        case OpCode::SIGNEXTEND:
            machineoperation::signExtend(*this);
            break;

            /***********************/
            /*  Hashing Operations */
            /***********************/
        case OpCode::HASH:
            machineoperation::hashOp(*this);
            break;
        case OpCode::TYPE:
            machineoperation::typeOp(*this);
            break;
        case OpCode::ETHHASH2:
            machineoperation::ethhash2Op(*this);
            break;

            /***********************************************/
            /*  Stack, Memory, Storage and Flow Operations */
            /***********************************************/
        case OpCode::POP:
            machineoperation::pop(*this);
            break;
        case OpCode::SPUSH:
            machineoperation::spush(*this);
            break;
        case OpCode::RPUSH:
            machineoperation::rpush(*this);
            break;
        case OpCode::RSET:
            machineoperation::rset(*this);
            break;
        case OpCode::JUMP:
            machineoperation::jump(*this);
            break;
        case OpCode::CJUMP:
            machineoperation::cjump(*this);
            break;
        case OpCode::STACKEMPTY:
            machineoperation::stackEmpty(*this);
            break;
        case OpCode::PCPUSH:
            machineoperation::pcPush(*this);
            break;
        case OpCode::AUXPUSH:
            machineoperation::auxPush(*this);
            break;
        case OpCode::AUXPOP:
            machineoperation::auxPop(*this);
            break;
        case OpCode::AUXSTACKEMPTY:
            machineoperation::auxStackEmpty(*this);
            break;
        case OpCode::NOP:
            ++pc;
            break;
        case OpCode::ERRPUSH:
            machineoperation::errPush(*this);
            break;
        case OpCode::ERRSET:
            machineoperation::errSet(*this);
            break;
            /****************************************/
            /*  Duplication and Exchange Operations */
            /****************************************/
        case OpCode::DUP0:
            machineoperation::dup0(*this);
            break;
        case OpCode::DUP1:
            machineoperation::dup1(*this);
            break;
        case OpCode::DUP2:
            machineoperation::dup2(*this);
            break;
        case OpCode::SWAP1:
            machineoperation::swap1(*this);
            break;
        case OpCode::SWAP2:
            machineoperation::swap2(*this);
            break;
            /*********************/
            /*  Tuple Operations */
            /*********************/
        case OpCode::TGET:
            machineoperation::tget(*this);
            break;
        case OpCode::TSET:
            machineoperation::tset(*this);
            break;
        case OpCode::TLEN:
            machineoperation::tlen(*this);
            break;
            /***********************/
            /*  Logging Operations */
            /***********************/
        case OpCode::BREAKPOINT:
            return machineoperation::breakpoint(*this);
        case OpCode::LOG:
            machineoperation::log(*this);
            break;
        case OpCode::DEBUG:
            machineoperation::debug(*this);
            break;
            /**********************/
            /*  System Operations */
            /**********************/
        case OpCode::SEND: {
            auto send_results = machineoperation::send(*this);

            if (send_results == false) {
                std::cerr << "Send failure: over size limit" << std::endl;
            }

            break;
        }
        case OpCode::GETTIME:
            machineoperation::getTime(*this);
            break;
        case OpCode::INBOX:
            return machineoperation::inboxOp(*this);
        case OpCode::ERROR:
            state = Status::Error;
            break;
        case OpCode::HALT:
            state = Status::Halted;
            break;
            /*****************/
            /*  Precompiles  */
            /*****************/
        case OpCode::ECRECOVER:
            machineoperation::ec_recover(*this);
            break;
        default:
            std::cerr << "Unhandled opcode <" << InstructionNames.at(opcode)
                      << ">" << std::hex << static_cast<int>(opcode);
            state = Status::Error;
    }

    return NotBlocked{};
}

std::ostream& operator<<(std::ostream& os, const MachineState& val) {
    os << "status " << static_cast<int>(val.state) << "\n";
    os << "codePointHash " << to_hex_str(hash(val.static_values->code[val.pc]))
       << "\n";
    os << "stackHash " << to_hex_str(val.stack.hash()) << "\n";
    os << "auxStackHash " << to_hex_str(val.auxstack.hash()) << "\n";
    os << "registerHash " << to_hex_str(hash_value(val.registerVal)) << "\n";
    os << "staticHash " << to_hex_str(hash_value(val.static_values->staticVal))
       << "\n";
    os << "errHandlerHash "
       << to_hex_str(hash(val.static_values->code[val.errpc])) << "\n";
    return os;
}
