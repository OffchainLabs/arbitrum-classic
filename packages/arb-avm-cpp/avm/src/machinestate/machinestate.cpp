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
#include <data_storage/checkpoint/checkpointutils.hpp>
#include <data_storage/checkpoint/machinestatefetcher.hpp>
#include <data_storage/checkpoint/machinestatesaver.hpp>
#include <data_storage/storageresult.hpp>

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
    auto initial_state = parseInitialVmValues(contract_filename, *pool.get());

    if (!initial_state.valid_state) {
        return std::make_pair(MachineState{}, false);
    }

    return std::make_pair(
        MachineState{initial_state.code, initial_state.staticVal,
                     std::move(pool)},
        true);
}

std::pair<MachineState, bool> MachineState::loadFromCheckpoint(
    const CheckpointStorage& storage,
    const std::vector<unsigned char>& checkpoint_key) {
    auto transaction = storage.makeConstTransaction();
    auto results = getMachineState(*transaction, checkpoint_key);

    auto initial_values = storage.getInitialVmValues();
    if (!initial_values.valid_state) {
        return std::make_pair(MachineState{}, false);
    }

    if (!results.status.ok()) {
        return std::make_pair(MachineState{}, false);
    }

    auto state_data = results.data;

    auto register_results =
        getValue(*transaction, state_data.register_val_key, storage.pool.get());
    if (!register_results.status.ok()) {
        return std::make_pair(MachineState{}, false);
    }

    Datastack stack;
    if (!stack.initializeDataStack(*transaction, state_data.datastack_key,
                                   storage.pool.get())) {
        return std::make_pair(MachineState{}, false);
    }

    Datastack auxstack;
    if (!auxstack.initializeDataStack(*transaction, state_data.auxstack_key,
                                      storage.pool.get())) {
        return std::make_pair(MachineState{}, false);
    }

    MachineState machine_state{storage.pool,
                               initial_values.code,
                               initial_values.staticVal,
                               std::move(register_results.data),
                               std::move(stack),
                               std::move(auxstack),
                               static_cast<Status>(state_data.status_char),
                               CodePointRef(state_data.pc),
                               CodePointRef(state_data.err_pc)};
    return std::make_pair(std::move(machine_state), true);
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
        auto val = ::hash_value(staticVal);
        oit = to_big_endian(val, oit);
    }
    {
        auto val = ::hash(code[errpc]);
        oit = to_big_endian(val, oit);
    }

    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(data.data(), data.size(), hashData.data());
    return from_big_endian(hashData.begin(), hashData.end());
}

uint256_t MachineState::getMachineSize() {
    uint256_t machine_size = 0;

    machine_size += getSize(staticVal);
    machine_size += getSize(registerVal);
    machine_size += stack.getTotalValueSize();
    machine_size += auxstack.getTotalValueSize();

    return machine_size;
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
    auto stackProof = stack.marshalForProof(stackPops, code);
    auto auxStackProof = auxstack.marshalForProof(auxStackPops, code);

    ::marshalStub(CodePointStub{code[pc + 1]}, buf, code);
    stackProof.first.marshal(buf);
    auxStackProof.first.marshal(buf);
    ::marshalStub(registerVal, buf, code);
    ::marshalStub(staticVal, buf, code);
    ::marshalStub(CodePointStub{code[errpc]}, buf, code);
    code[pc].op.marshalForProof(buf, includeImmediateVal, code);

    buf.insert(buf.end(), stackProof.second.begin(), stackProof.second.end());
    buf.insert(buf.end(), auxStackProof.second.begin(),
               auxStackProof.second.end());
    return buf;
}

SaveResults MachineState::checkpointState(CheckpointStorage& storage) {
    auto transaction = storage.makeTransaction();

    auto datastack_results = stack.checkpointState(*transaction, pool.get());
    auto auxstack_results = auxstack.checkpointState(*transaction, pool.get());

    auto register_val_results = saveValue(*transaction, registerVal);
    auto err_pc_stub = CodePointStub{code[errpc]};
    auto pc_stub = CodePointStub{code[pc]};

    auto status_str = static_cast<unsigned char>(state);

    std::vector<unsigned char> hash_key;
    marshal_uint256_t(hash(), hash_key);

    if (datastack_results.status.ok() && auxstack_results.status.ok() &&
        register_val_results.status.ok()) {
        auto machine_state_data =
            MachineStateKeys{register_val_results.storage_key,
                             datastack_results.storage_key,
                             auxstack_results.storage_key,
                             pc_stub,
                             err_pc_stub,
                             status_str};

        auto results =
            saveMachineState(*transaction, machine_state_data, hash_key);
        results.status = transaction->commit();
        return results;
    } else {
        return SaveResults{0, rocksdb::Status().Aborted(), hash_key};
    }
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
        default:
            std::cerr << "Unhandled opcode <" << InstructionNames.at(opcode)
                      << ">" << std::hex << static_cast<int>(opcode);
            state = Status::Error;
    }

    return NotBlocked{};
}
