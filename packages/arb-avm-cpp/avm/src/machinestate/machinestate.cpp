/*
 * Copyright 2019, Offchain Labs, Inc.
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
#include <data_storage/checkpoint/checkpointstorage.hpp>
#include <data_storage/checkpoint/machinestatefetcher.hpp>
#include <data_storage/checkpoint/machinestatesaver.hpp>

#include <avm_values/util.hpp>
#include <bigint_utils.hpp>

void uint256_t_to_buf(const uint256_t& val, std::vector<unsigned char>& buf) {
    std::array<unsigned char, 32> tmpbuf;
    to_big_endian(val, tmpbuf.begin());
    buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
}

MachineState::MachineState()
    : pool(std::make_unique<TuplePool>()), context({0, 0}, Tuple()) {}

MachineState::MachineState(const std::vector<CodePoint>& code_,
                           const value& static_val_,
                           std::shared_ptr<TuplePool> pool_)
    : pool(std::move(pool_)), context({0, 0}, Tuple()) {
    code = code_;
    staticVal = static_val_;

    errpc = getErrCodePoint();
    pc = 0;
}

bool MachineState::initialize_machinestate(
    const std::string& contract_filename) {
    auto initial_state = parseInitialVmValues(contract_filename, *pool.get());

    if (initial_state.valid_state) {
        code = initial_state.code;
        staticVal = initial_state.staticVal;

        errpc = getErrCodePoint();
        pc = 0;

        return true;
    } else {
        return false;
    }
}

uint256_t MachineState::hash() const {
    if (state == Status::Halted)
        return 0;
    if (state == Status::Error)
        return 1;

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

bool MachineState::verifyMachineValidity() {
    int machine_size = 0;

    machine_size += getSize(staticVal);
    machine_size += getSize(registerVal);

    machine_size += getSize(errpc);
    machine_size += stack.getTotalValueSize();
    machine_size += auxstack.getTotalValueSize();

    auto valid_machine = machine_size <= machine_size_limit;

    return valid_machine;
}

std::vector<unsigned char> MachineState::marshalForProof() {
    std::vector<unsigned char> buf;
    auto opcode = code[pc].op.opcode;
    std::vector<bool> stackPops = InstructionStackPops.at(opcode);
    bool includeImmediateVal = false;
    if (code[pc].op.immediate && !stackPops.empty()) {
        includeImmediateVal = stackPops[0] == true;
        stackPops.erase(stackPops.begin());
    }
    std::vector<bool> auxStackPops = InstructionAuxStackPops.at(opcode);
    auto stackProof = stack.marshalForProof(stackPops);
    auto auxStackProof = auxstack.marshalForProof(auxStackPops);

    //    HashOnly nextHash(code[pc].nextHash, 1);
    //    nextHash.marshal(buf);
    uint256_t_to_buf(code[pc].nextHash, buf);

    stackProof.first.marshal(buf);
    auxStackProof.first.marshal(buf);

    uint256_t_to_buf(::hash(registerVal), buf);
    uint256_t_to_buf(::hash(staticVal), buf);
    uint256_t_to_buf(::hash(errpc), buf);

    //    HashOnly registerValHash(::hash(registerVal), ::getSize(registerVal));
    //    registerValHash.marshal(buf);
    //
    //    HashOnly staticValHash(::hash(staticVal), ::getSize(staticVal));
    //    staticValHash.marshal(buf);
    //
    //    HashOnly errpcHash(::hash(errpc), ::getSize(errpc));
    //    errpcHash.marshal(buf);

    code[pc].op.marshalForProof(buf, includeImmediateVal);

    buf.insert(buf.end(), stackProof.second.begin(), stackProof.second.end());
    buf.insert(buf.end(), auxStackProof.second.begin(),
               auxStackProof.second.end());
    return buf;
}

SaveResults MachineState::checkpointState(CheckpointStorage& storage) {
    auto stateSaver = MachineStateSaver(storage.makeTransaction());

    auto datastack_results = stack.checkpointState(stateSaver, pool.get());
    auto auxstack_results = auxstack.checkpointState(stateSaver, pool.get());

    auto register_val_results = stateSaver.saveValue(registerVal);
    auto err_code_point = stateSaver.saveValue(errpc);
    auto pc_results = stateSaver.saveValue(code[pc]);

    auto status_str = static_cast<unsigned char>(state);

    std::vector<unsigned char> hash_key;
    marshal_uint256_t(hash(), hash_key);

    if (datastack_results.status.ok() && auxstack_results.status.ok() &&
        register_val_results.status.ok() && pc_results.status.ok() &&
        err_code_point.status.ok()) {
        auto machine_state_data = MachineStateKeys{
            register_val_results.storage_key, datastack_results.storage_key,
            auxstack_results.storage_key,     pc_results.storage_key,
            err_code_point.storage_key,       status_str};

        auto results =
            stateSaver.saveMachineState(machine_state_data, hash_key);
        results.status = stateSaver.commitTransaction();
        return results;
    } else {
        return SaveResults{0, rocksdb::Status().Aborted(), hash_key};
    }
}

bool MachineState::restoreCheckpoint(
    const CheckpointStorage& storage,
    const std::vector<unsigned char>& checkpoint_key) {
    auto stateFetcher = MachineStateFetcher(storage);
    auto results = stateFetcher.getMachineState(checkpoint_key);

    auto initial_values = storage.getInitialVmValues();
    if (!initial_values.valid_state) {
        return false;
    }

    code = initial_values.code;
    staticVal = initial_values.staticVal;
    pool = storage.pool;

    if (results.status.ok()) {
        auto state_data = results.data;

        auto register_results =
            stateFetcher.getValue(state_data.register_val_key);
        registerVal = register_results.data;

        auto pc_results = stateFetcher.getCodePoint(state_data.pc_key);
        pc = pc_results.data.pc;

        auto err_pc_results = stateFetcher.getCodePoint(state_data.err_pc_key);
        errpc = err_pc_results.data;

        if (!stack.initializeDataStack(stateFetcher,
                                       state_data.datastack_key)) {
            return false;
        }

        if (!auxstack.initializeDataStack(stateFetcher,
                                          state_data.auxstack_key)) {
            return false;
        }

        state = static_cast<Status>(state_data.status_char);
    }
    return results.status.ok();
}

BlockReason MachineState::isBlocked(uint256_t currentTime,
                                    bool newMessages) const {
    if (state == Status::Error) {
        return ErrorBlocked();
    } else if (state == Status::Halted) {
        return HaltBlocked();
    }
    auto& instruction = code[pc];
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

            if (send_results.success == false) {
                std::cerr << "Send failure: over size limit" << std::endl;
            }

            return send_results.block_reason;
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
        default:
            std::cerr << "Unhandled opcode <" << InstructionNames.at(opcode)
                      << ">" << std::hex << static_cast<int>(opcode);
            state = Status::Error;
    }

    auto valid_machine = verifyMachineValidity();

    if (!valid_machine) {
        state = Status::Error;
        std::cerr << "Machine size invalid : error state" << std::endl;
    }

    return NotBlocked{};
}
