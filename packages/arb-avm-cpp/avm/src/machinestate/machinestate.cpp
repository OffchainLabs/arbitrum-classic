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

#include <avm/machine.hpp>
#include <avm/machinestate/machineoperation.hpp>
#include <avm/machinestate/config.hpp>
#include <avm_values/exceptions.hpp>
#include <avm_values/vmValueParser.hpp>
#include <utility>

#include <ethash/keccak.hpp>

#include <iostream>
#include <fstream>

namespace {
uint256_t max_arb_gas_remaining = std::numeric_limits<uint256_t>::max();
}  // namespace

AssertionContext::AssertionContext(MachineExecutionConfig config)
    : inbox_messages(std::move(config.inbox_messages)),
      sideloads(std::move(config.sideloads)),
      stop_on_sideload(config.stop_on_sideload),
      max_gas(config.max_gas),
      go_over_gas(config.go_over_gas),
      inbox_messages_consumed(0) {}

MachineStateKeys::MachineStateKeys(const MachineState& machine)
    : static_hash(hash_value(machine.static_val)),
      register_hash(hash_value(machine.registerVal)),
      datastack_hash(machine.stack.hash()),
      auxstack_hash(machine.auxstack.hash()),
      arb_gas_remaining(machine.arb_gas_remaining),
      pc(machine.pc, machine.loadCurrentInstruction()),
      err_pc(machine.errpc),
      status(machine.state),
      output(machine.output) {}

uint256_t MachineStateKeys::machineHash() const {
    if (status == Status::Halted)
        return 0;
    if (status == Status::Error)
        return 1;

    std::array<unsigned char, 32 * 7> data{};
    auto oit = data.begin();
    {
        auto val = ::hash(pc);
        oit = to_big_endian(val, oit);
    }
    { oit = to_big_endian(datastack_hash, oit); }
    { oit = to_big_endian(auxstack_hash, oit); }
    { oit = to_big_endian(register_hash, oit); }
    { oit = to_big_endian(static_hash, oit); }
    { oit = to_big_endian(arb_gas_remaining, oit); }
    {
        auto val = ::hash_value(err_pc);
        oit = to_big_endian(val, oit);
    }
    assert(oit == data.end());

    auto hash_val = ethash::keccak256(data.data(), data.size());
    return intx::be::load<uint256_t>(hash_val);
}

void MachineState::addProcessedMessage(const MachineMessage& message) {
    output.fully_processed_inbox.addMessage(message);
}

void MachineState::addProcessedSend(std::vector<uint8_t> data) {
    output.send_count = output.send_count + 1;
    output.send_acc = ::hash(output.send_acc, ::hash(data));
    context.sends.push_back(std::move(data));
}

void MachineState::addProcessedLog(value log_val) {
    output.log_count = output.log_count + 1;
    output.log_acc = ::hash(output.log_acc, hash_value(log_val));
    context.logs.push_back(std::move(log_val));
}

MachineState::MachineState() : arb_gas_remaining(max_arb_gas_remaining) {}

MachineState::MachineState(std::shared_ptr<Code> code_, value static_val_)
    : code(std::move(code_)),
      static_val(std::move(static_val_)),
      arb_gas_remaining(max_arb_gas_remaining),
      pc(code->initialCodePointRef()) {}

MachineState::MachineState(std::shared_ptr<Code> code_,
                           value register_val_,
                           value static_val_,
                           Datastack stack_,
                           Datastack auxstack_,
                           uint256_t arb_gas_remaining_,
                           Status state_,
                           CodePointRef pc_,
                           CodePointStub errpc_,
                           MachineOutput output_)
    : code(std::move(code_)),
      registerVal(std::move(register_val_)),
      static_val(std::move(static_val_)),
      stack(std::move(stack_)),
      auxstack(std::move(auxstack_)),
      arb_gas_remaining(arb_gas_remaining_),
      state(state_),
      pc(pc_),
      errpc(errpc_),
      output(std::move(output_)) {}

MachineState MachineState::loadFromFile(
    const std::string& executable_filename) {
    auto executable = loadExecutable(executable_filename);
    auto code = std::make_shared<Code>(0);
    code->addSegment(std::move(executable.code));
    return MachineState{std::move(code), std::move(executable.static_val)};
}

uint256_t MachineState::getMachineSize() const {
    uint256_t machine_size = 0;

    machine_size += getSize(static_val);
    machine_size += getSize(registerVal);
    machine_size += stack.getTotalValueSize();
    machine_size += auxstack.getTotalValueSize();

    return machine_size;
}

namespace {
void marshalState(std::vector<unsigned char>& buf,
                  const Code& code,
                  uint256_t next_codepoint_hash,
                  HashPreImage stackPreImage,
                  HashPreImage auxStackPreImage,
                  const value& registerVal,
                  const value& staticVal,
                  uint256_t arb_gas_remaining,
                  CodePointStub errpc) {
    marshal_uint256_t(next_codepoint_hash, buf);

    stackPreImage.marshal(buf);
    auxStackPreImage.marshal(buf);

    ::marshalForProof(registerVal, 0, buf, code);
    ::marshalForProof(staticVal, 0, buf, code);
    marshal_uint256_t(arb_gas_remaining, buf);
    marshal_uint256_t(::hash(errpc), buf);
}
}  // namespace

std::vector<unsigned char> MachineState::marshalState() const {
    auto stackPreImage = stack.getHashPreImage();
    auto auxStackPreImage = auxstack.getHashPreImage();
    std::vector<unsigned char> buf;

    ::marshalState(buf, *code, ::hash(loadCurrentInstruction()), stackPreImage,
                   auxStackPreImage, registerVal, static_val, arb_gas_remaining,
                   errpc);
    return buf;
}

void insertSizes(std::vector<unsigned char>& buf,
                 int sz1,
                 int sz2,
                 int sz3,
                 int sz4) {
    int acc = 1;
    buf.push_back(static_cast<uint8_t>(acc));
    acc += sz1 / 32;
    buf.push_back(static_cast<uint8_t>(acc));
    acc += sz2 / 32;
    buf.push_back(static_cast<uint8_t>(acc));
    acc += sz3 / 32;
    buf.push_back(static_cast<uint8_t>(acc));
    acc += sz4 / 32;
    buf.push_back(static_cast<uint8_t>(acc));
    for (int i = 5; i < 32; i++) {
        buf.push_back(0);
    }
}

void makeSetBufferProof(std::vector<unsigned char>& buf,
                        uint64_t loc,
                        const Buffer& buffer,
                        uint256_t v,
                        int wordSize) {
    Buffer nbuffer = buffer;
    Buffer nbuffer1 = nbuffer;
    bool aligned = true;
    for (int i = 0; i < wordSize; i++) {
        if ((loc + i) % 32 == 0 && i > 0) {
            nbuffer1 = nbuffer;
            aligned = false;
        }
        nbuffer = nbuffer.set(
            loc + i,
            static_cast<uint8_t>((v >> ((wordSize - 1 - i) * 8)) & 0xff));
    }
    auto proof1 = buffer.makeProof(loc);

    if (aligned) {
        auto nproof1 = nbuffer.makeNormalizationProof();
        insertSizes(buf, proof1.size(), nproof1.size(), 0, 0);
        buf.insert(buf.end(), proof1.begin(), proof1.end());
        buf.insert(buf.end(), nproof1.begin(), nproof1.end());
    } else {
        auto nproof1 = nbuffer1.makeNormalizationProof();
        auto proof2 = nbuffer1.makeProof(loc + (wordSize - 1));
        auto nproof2 = nbuffer.makeNormalizationProof();
        insertSizes(buf, proof1.size(), nproof1.size(), proof2.size(),
                    nproof2.size());
        buf.insert(buf.end(), proof1.begin(), proof1.end());
        buf.insert(buf.end(), nproof1.begin(), nproof1.end());
        buf.insert(buf.end(), proof2.begin(), proof2.end());
        buf.insert(buf.end(), nproof2.begin(), nproof2.end());
    }
}

value MachineState::getStackOrImmed(uint64_t num, Operation op) const {
    if (num == 0 && op.immediate) return *op.immediate;
    if (op.immediate) num = num - 1;
    return stack[num];
}

void MachineState::marshalBufferProof(OneStepProof& proof) const {
    auto& op = loadCurrentOperation();
    auto opcode = op.opcode;
    if ((opcode < OpCode::GET_BUFFER8 || opcode > OpCode::SET_BUFFER256) &&
        opcode != OpCode::SEND) {
        return;
    }
    if (opcode == OpCode::SEND) {
        auto buffer = op.immediate ? std::get_if<Buffer>(&stack[0])
                                   : std::get_if<Buffer>(&stack[1]);
        if (!buffer) {
            return;
        }
        // Also need the offset
        auto size = op.immediate ? std::get_if<uint256_t>(&*op.immediate)
                                 : std::get_if<uint256_t>(&stack[0]);
        if (!size) {
            return;
        }
        auto loc = static_cast<uint64_t>(*size);
        if (loc > send_size_limit) {
            return;
        } else if (loc < buffer->data_length()) {
            // Loc must be at or past the last nonzero index in the buffer
            auto buf_proof = buffer->makeProof(loc);
            insertSizes(proof.buffer_proof, buf_proof.size(), 0, 0, 0);
            proof.buffer_proof.insert(proof.buffer_proof.end(),
                                      buf_proof.begin(), buf_proof.end());
        } else {
            auto data = buffer->toFlatVector();
            proof.standard_proof.insert(proof.standard_proof.end(),
                                        data.begin(), data.end());
            std::fill_n(std::back_inserter(proof.standard_proof),
                        loc - data.size(), 0);
        }
        return;
    }
    if (opcode == OpCode::GET_BUFFER8 || opcode == OpCode::GET_BUFFER64 ||
        opcode == OpCode::GET_BUFFER256) {
        // Find the buffer
        auto buffer = op.immediate ? std::get_if<Buffer>(&stack[0])
                                   : std::get_if<Buffer>(&stack[1]);
        if (!buffer) {
            insertSizes(proof.buffer_proof, 0, 0, 0, 0);
            return;
        }
        // Also need the offset
        auto offset = op.immediate ? std::get_if<uint256_t>(&*op.immediate)
                                   : std::get_if<uint256_t>(&stack[0]);
        if (!offset) {
            insertSizes(proof.buffer_proof, 0, 0, 0, 0);
            return;
        }
        if (*offset > std::numeric_limits<uint64_t>::max()) {
            insertSizes(proof.buffer_proof, 0, 0, 0, 0);
            return;
        }
        auto loc = static_cast<uint64_t>(*offset);
        if (opcode == OpCode::GET_BUFFER8) {
            auto buf_proof = buffer->makeProof(loc);
            insertSizes(proof.buffer_proof, buf_proof.size(), 0, 0, 0);
            proof.buffer_proof.insert(proof.buffer_proof.end(),
                                      buf_proof.begin(), buf_proof.end());
        } else if (opcode == OpCode::GET_BUFFER64) {
            auto buf_proof1 = buffer->makeProof(loc);
            auto buf_proof2 = buffer->makeProof(loc + 7);
            insertSizes(proof.buffer_proof, buf_proof1.size(), 0,
                        buf_proof2.size(), 0);
            proof.buffer_proof.insert(proof.buffer_proof.end(),
                                      buf_proof1.begin(), buf_proof1.end());
            proof.buffer_proof.insert(proof.buffer_proof.end(),
                                      buf_proof2.begin(), buf_proof2.end());
        } else if (opcode == OpCode::GET_BUFFER256) {
            auto buf_proof1 = buffer->makeProof(loc);
            auto buf_proof2 = buffer->makeProof(loc + 31);
            insertSizes(proof.buffer_proof, buf_proof1.size(), 0,
                        buf_proof2.size(), 0);
            proof.buffer_proof.insert(proof.buffer_proof.end(),
                                      buf_proof1.begin(), buf_proof1.end());
            proof.buffer_proof.insert(proof.buffer_proof.end(),
                                      buf_proof2.begin(), buf_proof2.end());
        }
    } else {
        auto buffer = op.immediate ? std::get_if<Buffer>(&stack[1])
                                   : std::get_if<Buffer>(&stack[2]);
        if (!buffer) {
            insertSizes(proof.buffer_proof, 0, 0, 0, 0);
            return;
        }
        // Also need the offset
        auto offset = op.immediate ? std::get_if<uint256_t>(&*op.immediate)
                                   : std::get_if<uint256_t>(&stack[0]);
        if (!offset) {
            insertSizes(proof.buffer_proof, 0, 0, 0, 0);
            return;
        }
        if (*offset > std::numeric_limits<uint64_t>::max()) {
            insertSizes(proof.buffer_proof, 0, 0, 0, 0);
            return;
        }
        auto val = op.immediate ? std::get_if<uint256_t>(&stack[0])
                                : std::get_if<uint256_t>(&stack[1]);
        if (!val) {
            insertSizes(proof.buffer_proof, 0, 0, 0, 0);
            return;
        }
        auto loc = static_cast<uint64_t>(*offset);
        if (opcode == OpCode::SET_BUFFER8) {
            Buffer nbuffer = buffer->set(loc, static_cast<uint8_t>(*val));
            auto buf_proof1 = buffer->makeProof(loc);
            auto buf_nproof1 = nbuffer.makeNormalizationProof();
            insertSizes(proof.buffer_proof, buf_proof1.size(),
                        buf_nproof1.size(), 0, 0);
            proof.buffer_proof.insert(proof.buffer_proof.end(),
                                      buf_proof1.begin(), buf_proof1.end());
            proof.buffer_proof.insert(proof.buffer_proof.end(),
                                      buf_nproof1.begin(), buf_nproof1.end());
        } else if (opcode == OpCode::SET_BUFFER64) {
            makeSetBufferProof(proof.buffer_proof, loc, *buffer, *val, 8);
        } else if (opcode == OpCode::SET_BUFFER256) {
            makeSetBufferProof(proof.buffer_proof, loc, *buffer, *val, 32);
        }
    }
}

const int LEVEL = 5;

value table_to_tuple2(std::vector<value> tab, int prefix, int shift, int level, int limit) {
    /*
    if (limit < prefix) {
        return 0;
    }*/
    if (level == 0) {
        std::vector<value> v;
        for (int i = 0; i < 8; i++) {
            uint64_t idx = prefix + (i << shift);
            if (idx < tab.size()) {
                v.push_back(tab[idx]);
            } else {
                v.push_back(0);
            }
        }
        return Tuple::createTuple(v);
    }
    std::vector<value> v;
    for (int i = 0; i < 8; i++) {
        v.push_back(table_to_tuple2(tab, prefix + (i << shift), shift + 3, level - 1, limit));
    }
    return Tuple::createTuple(v);
}

value make_table(std::vector<value> tab) {
    return table_to_tuple2(tab, 0, 0, LEVEL-1, tab.size());
}

value get_int_value(std::vector<uint8_t> &bytes, uint64_t offset) {
    uint256_t acc = 0;
    for (int i = 0; i < 32; i++) {
        acc = acc*256;
        acc += bytes[offset+i];
    }
    return acc;
}

CodeResult wasmAvmToCode(WasmResult &res) {
    auto code = std::make_shared<Code>(0);
    CodePointStub stub = code->addSegment();

    std::vector<CodePointStub> points;
    std::vector<value> tab_lst;

    for (uint64_t i = 0; i < res.insn->size(); i++) {
        points.push_back(stub);
        stub = code->addOperation(stub.pc, (*res.insn)[i]);
    }

    for (uint64_t i = 0; i < res.table.size(); i++) {
        auto offset = res.table[i].first;
        if (offset >= tab_lst.size()) {
            tab_lst.resize(offset+1);
        }
        tab_lst[offset] = points[res.table[i].second];
    }
    auto table = make_table(tab_lst);

    return {code, table, stub};
}

WasmCodePoint wasmAvmToCodePoint(WasmResult &wres, std::vector<uint8_t>& wasm_module) {
    auto res = wasmAvmToCode(wres);
    std::shared_ptr<Tuple> tpl = std::make_shared<Tuple>(res.stub, res.table, vec2buf(wasm_module), wasm_module.size());
    std::shared_ptr<WasmRunner> runner = std::make_shared<RunWasm>(wasm_module);
    return {std::move(tpl), std::move(runner)};
}

MachineState makeWasmMachine(WasmResult &wres, uint64_t len, Buffer buf, value arg) {
    auto res = wasmAvmToCode(wres);
    MachineState state(res.code, 0);
    state.stack.push(len);
    state.stack.push(buf);
    state.stack.push(std::move(res.table));
    state.stack.push(std::move(arg));
    state.arb_gas_remaining = 1000000000000;
    state.output.arb_gas_used = 0;

    return state;
}

uint256_t runWasmMachine(MachineState &machine_state) {
    uint256_t start_gas = machine_state.arb_gas_remaining;

    bool has_gas_limit = machine_state.context.max_gas != 0;
    BlockReason block_reason = NotBlocked{};
    // uint64_t counter = 0;
    while (true) {
        if (has_gas_limit) {
            if (!machine_state.context.go_over_gas) {
                if (machine_state.nextGasCost() +
                        machine_state.output.arb_gas_used >
                    machine_state.context.max_gas) {
                    // Next step would go over gas limit
                    break;
                }
            } else if (machine_state.output.arb_gas_used >=
                       machine_state.context.max_gas) {
                // Last step reached or went over gas limit
                break;
            }
        }
/*
        auto op = machine_state.loadCurrentInstruction();
        std::cerr << "op " << op << " state " << int(machine_state.state) << "\n";
        if (machine_state.stack.stacksize() > 0 && !std::get_if<Tuple>(&machine_state.stack[0])) {
            std::cerr << "stack top " << machine_state.stack[0] << "\n";
        }
       if (counter++ % 1000000 == 0) {
           std::cerr << "step " << counter << "\n";
       }
*/
        auto& op = machine_state.loadCurrentOperation();

        if (op.opcode == OpCode::HALT) {
            break;
        }
        block_reason = machine_state.runOne();
        

        if (!std::get_if<NotBlocked>(&block_reason)) {
            break;
        }
    }    
    return start_gas - machine_state.arb_gas_remaining;
}

void MachineState::marshalWasmProof(OneStepProof &proof) const {
    auto currentInstruction = loadCurrentInstruction();
    auto& current_op = currentInstruction.op;

    std::vector<size_t> stackPops = { 1, 1, 1, 1 };

    std::vector<size_t> auxStackPops;

    size_t immediateMarshalLevel = 0;
    if (current_op.immediate && !stackPops.empty()) {
        immediateMarshalLevel = stackPops[0];
        stackPops = { 1, 1, 1 };
    }

    auto stackProof = stack.marshalForProof(stackPops, *code);
    auto auxStackProof = auxstack.marshalForProof(auxStackPops, *code);

    proof.buffer_proof.push_back(static_cast<uint8_t>(current_op.opcode));
    proof.buffer_proof.push_back(stackProof.count +
                                   (current_op.immediate ? 1 : 0));
    proof.buffer_proof.push_back(auxStackProof.count);

    proof.buffer_proof.insert(proof.buffer_proof.cend(),
                                stackProof.data.begin(), stackProof.data.end());
    if (current_op.immediate) {
        ::marshalForProof(*current_op.immediate, immediateMarshalLevel,
                          proof.buffer_proof, *code);
    }
    proof.buffer_proof.insert(proof.buffer_proof.cend(),
                                auxStackProof.data.begin(),
                                auxStackProof.data.end());
    ::marshalState(proof.buffer_proof, *code, currentInstruction.nextHash,
                   stackProof.bottom, auxStackProof.bottom, registerVal,
                   static_val, arb_gas_remaining, errpc);

    proof.buffer_proof.push_back(current_op.immediate ? 1 : 0);

}

void MachineState::marshalWasmCompileProof(OneStepProof &proof) const {
    auto currentInstruction = loadCurrentInstruction();
    auto& current_op = currentInstruction.op;

    std::vector<size_t> stackPops = {
        1,
        1,
        1,
        1,
        1,
    };

    std::vector<size_t> auxStackPops;

    size_t immediateMarshalLevel = 0;
    if (current_op.immediate && !stackPops.empty()) {
        immediateMarshalLevel = stackPops[0];
        stackPops.erase(stackPops.cbegin());
    }

    auto stackProof = stack.marshalForProof(stackPops, *code);
    auto auxStackProof = auxstack.marshalForProof(auxStackPops, *code);

    proof.buffer_proof.push_back(static_cast<uint8_t>(current_op.opcode));
    proof.buffer_proof.push_back(stackProof.count +
                                   (current_op.immediate ? 1 : 0));
    proof.buffer_proof.push_back(auxStackProof.count);

    proof.buffer_proof.insert(proof.buffer_proof.cend(),
                                stackProof.data.begin(), stackProof.data.end());
    if (current_op.immediate) {
        ::marshalForProof(*current_op.immediate, immediateMarshalLevel,
                          proof.buffer_proof, *code);
    }
    proof.buffer_proof.insert(proof.buffer_proof.cend(),
                                auxStackProof.data.begin(),
                                auxStackProof.data.end());
    ::marshalState(proof.buffer_proof, *code, currentInstruction.nextHash,
                   stackProof.bottom, auxStackProof.bottom, registerVal,
                   static_val, arb_gas_remaining, errpc);

    proof.buffer_proof.push_back(current_op.immediate ? 1 : 0);

}

MachineState MachineState::initialWasmMachine() const {
    auto currentInstruction = loadCurrentInstruction();
    auto& op = currentInstruction.op;
    if (op.opcode == OpCode::WASM_RUN) {
        auto elem_arg = getStackOrImmed(0, op);
        auto elem0 = getStackOrImmed(1, op);
        auto elem1 = getStackOrImmed(2, op);
        auto elem2 = getStackOrImmed(3, op);
        uint64_t len = static_cast<uint64_t>(*std::get_if<uint256_t>(&elem0));
        Buffer buf = *std::get_if<Buffer>(&elem1);
        WasmCodePoint cp = *std::get_if<WasmCodePoint>(&elem2);

        auto tpl3 = cp.data->get_element(3);
        auto tpl2 = cp.data->get_element(2);
        // Looks like the table and jump point are useless, need to recreate them for the "submachine"
        uint64_t code_len = static_cast<uint64_t>(*std::get_if<uint256_t>(&tpl3));
        Buffer code_buffer = *std::get_if<Buffer>(&tpl2);
        auto res = compile->run_wasm(code_buffer, code_len);

        return makeWasmMachine(res, len, buf, elem_arg);
    } else if (op.opcode == OpCode::WASM_COMPILE) {
        auto elem0 = getStackOrImmed(0, op);
        auto elem1 = getStackOrImmed(1, op);
        uint64_t len = static_cast<uint64_t>(*std::get_if<uint256_t>(&elem0));
        Buffer buf = *std::get_if<Buffer>(&elem1);

        std::ifstream input(wasm_compiler_path, std::ios::binary);
        std::vector<uint8_t> bytes((std::istreambuf_iterator<char>(input)), (std::istreambuf_iterator<char>()));
        input.close();

        auto res = compile->run_wasm(vec2buf(bytes), bytes.size());

        return makeWasmMachine(res, len, buf, 0);
    } else {
        throw std::runtime_error("Not a wasm instruction");
    }
}

WasmCodePoint MachineState::compiledWasmCodePoint() const {
    auto currentInstruction = loadCurrentInstruction();
    auto& op = currentInstruction.op;
    auto elem0 = getStackOrImmed(0, op);
    auto elem1 = getStackOrImmed(1, op);
    uint64_t code_len = static_cast<uint64_t>(*std::get_if<uint256_t>(&elem0));
    Buffer code_buf = *std::get_if<Buffer>(&elem1);
    auto res = compile->run_wasm(code_buf, code_len);
    auto bytes = buf2vec(res.buffer, res.buffer_len);
    auto wasm_bytes = buf2vec(code_buf, code_len);
    return wasmAvmToCodePoint(res, wasm_bytes);
}

MachineState MachineState::finalWasmMachine() const {
    auto state = initialWasmMachine();
    runWasmMachine(state);
    return state;
}

OneStepProof MachineState::marshalForProof() const {
    auto currentInstruction = loadCurrentInstruction();
    auto& current_op = currentInstruction.op;
    auto opcode = current_op.opcode;

    std::vector<size_t> stackPops = [&]() {
        auto it = InstructionStackPops.find(opcode);
        if (it == InstructionStackPops.end()) {
            return InstructionStackPops.at(OpCode::ERROR);
        }
        return it->second;
    }();

    std::vector<size_t> auxStackPops = [&]() {
        auto it = InstructionAuxStackPops.find(opcode);
        if (it == InstructionAuxStackPops.end()) {
            return InstructionAuxStackPops.at(OpCode::ERROR);
        }
        return it->second;
    }();

    size_t immediateMarshalLevel = 0;
    if (current_op.immediate && !stackPops.empty()) {
        immediateMarshalLevel = stackPops[0];
        stackPops.erase(stackPops.cbegin());
    }

    OneStepProof proof;

    auto stackProof = stack.marshalForProof(stackPops, *code);
    auto auxStackProof = auxstack.marshalForProof(auxStackPops, *code);

    bool underflowed = stackProof.count < stackPops.size() ||
                       auxStackProof.count < auxStackPops.size();

    proof.standard_proof.push_back(static_cast<uint8_t>(current_op.opcode));
    proof.standard_proof.push_back(
        stackProof.count +
        static_cast<uint8_t>(static_cast<bool>(current_op.immediate)));
    proof.standard_proof.push_back(auxStackProof.count);

    proof.standard_proof.insert(proof.standard_proof.cend(),
                                stackProof.data.begin(), stackProof.data.end());
    if (current_op.immediate) {
        ::marshalForProof(*current_op.immediate, immediateMarshalLevel,
                          proof.standard_proof, *code);
    }
    proof.standard_proof.insert(proof.standard_proof.cend(),
                                auxStackProof.data.begin(),
                                auxStackProof.data.end());
    ::marshalState(proof.standard_proof, *code, currentInstruction.nextHash,
                   stackProof.bottom, auxStackProof.bottom, registerVal,
                   static_val, arb_gas_remaining, errpc);

    proof.standard_proof.push_back(current_op.immediate ? 1 : 0);

    if (current_op.opcode == OpCode::WASM_RUN) {

        // std::cerr << "machine state before wasm run " << *this;
        auto state = initialWasmMachine();

        uint256_t gasUsed = runWasmMachine(state);

        state.marshalWasmProof(proof);
        marshal_uint256_t(gasUsed, proof.buffer_proof);
    } else if (current_op.opcode == OpCode::WASM_COMPILE) {

        auto state = initialWasmMachine();

        uint256_t gasUsed = runWasmMachine(state);

        state.marshalWasmCompileProof(proof);
        marshal_uint256_t(gasUsed, proof.buffer_proof);
    } else if (!underflowed) {
        // Don't need a buffer proof if we're underflowing
        marshalBufferProof(proof);
    }
    return proof;
}

BlockReason MachineState::isBlocked(bool newMessages) const {
    if (state == Status::Error) {
        return ErrorBlocked();
    } else if (state == Status::Halted) {
        return HaltBlocked();
    }
    auto& op = loadCurrentOperation();
    if (op.opcode == OpCode::INBOX) {
        if (newMessages) {
            return NotBlocked();
        }
        return InboxBlocked();
    } else {
        return NotBlocked();
    }
}

CodePoint MachineState::loadCurrentInstruction() const {
    if (!loaded_segment || loaded_segment->segment->segmentID() != pc.segment) {
        loaded_segment = std::make_optional(code->loadCodeSegment(pc.segment));
    }
    return loaded_segment->segment->loadCodePoint(pc.pc);
}

const Operation& MachineState::loadCurrentOperation() const {
    if (!loaded_segment || loaded_segment->segment->segmentID() != pc.segment) {
        loaded_segment = std::make_optional(code->loadCodeSegment(pc.segment));
    }
    return loaded_segment->segment->loadOperation(pc.pc);
}

uint256_t MachineState::nextGasCost() const {
    auto& op = loadCurrentOperation();
    return gasCost(op);
}

uint256_t MachineState::gasCost(const Operation& op) const {
    auto base_gas = instructionGasCosts()[static_cast<size_t>(op.opcode)];
    if (op.opcode == OpCode::ECPAIRING) {
        base_gas += machineoperation::ec_pairing_variable_gas_cost(*this);
    }
    return base_gas;
}

BlockReason MachineState::runOne() {
    if (state == Status::Error) {
        return ErrorBlocked();
    }

    if (state == Status::Halted) {
        return HaltBlocked();
    }

    auto& op = loadCurrentOperation();

    static const auto error_gas_cost =
        instructionGasCosts()[static_cast<size_t>(OpCode::ERROR)];

    // Always push the immediate to the stack if we're not blocked
    if (op.immediate) {
        auto imm = *op.immediate;
        stack.push(std::move(imm));
    }

    // save stack size for stack cleanup in case of error
    uint64_t start_stack_size = stack.stacksize();
    uint64_t start_auxstack_size = auxstack.stacksize();

    // std::cerr << "Inst " << op.opcode << " gas used " << output.arb_gas_used << "\n";

    bool is_valid_instruction =
        instructionValidity()[static_cast<size_t>(op.opcode)];

    uint64_t stack_arg_count =
        is_valid_instruction ? InstructionStackPops.at(op.opcode).size() : 0;
    uint64_t auxstack_arg_count =
        is_valid_instruction ? InstructionAuxStackPops.at(op.opcode).size() : 0;

    // We're only blocked if we can't execute at all
    BlockReason blockReason = [&]() -> BlockReason {
        if (stack_arg_count > stack.stacksize() ||
            auxstack_arg_count > auxstack.stacksize()) {
            state = Status::Error;

            if (arb_gas_remaining < error_gas_cost) {
                arb_gas_remaining = max_arb_gas_remaining;
            } else {
                arb_gas_remaining -= error_gas_cost;
            }
            output.arb_gas_used += error_gas_cost;
            return NotBlocked();
        }

        uint256_t gas_cost =
            is_valid_instruction ? gasCost(op) : error_gas_cost;

        if (arb_gas_remaining < gas_cost) {
            // If there's insufficient gas remaining, execute by transitioning
            // to the error state with remaining gas set to max
            output.arb_gas_used += error_gas_cost;
            arb_gas_remaining = max_arb_gas_remaining;
            state = Status::Error;
            return NotBlocked();
        }
        arb_gas_remaining -= gas_cost;
        output.arb_gas_used += gas_cost;

        if (!is_valid_instruction) {
            // The opcode is invalid, execute by transitioning to the error
            // state
            state = Status::Error;
            return NotBlocked();
        }

        BlockReason blockReason = NotBlocked();
        try {
            blockReason = runOp(op.opcode);
        } catch (const stack_too_small&) {
            // Charge an error instruction instead
            arb_gas_remaining += gas_cost;
            output.arb_gas_used -= gas_cost;
        } catch (const std::exception&) {
            state = Status::Error;
        }

        if (!std::holds_alternative<NotBlocked>(blockReason)) {
            // Get rid of the immediate and reset the gas if the machine was
            // actually blocked
            arb_gas_remaining += gas_cost;
            output.arb_gas_used -= gas_cost;
            if (op.immediate) {
                stack.popClear();
            }
            return blockReason;
        }
        return NotBlocked();
    }();

    if (std::holds_alternative<NotBlocked>(blockReason)) {
        output.total_steps += 1;
    }

    if (state == Status::Error) {
        // if state is Error, clean up stack
        // Clear stack to base for instruction
        while (stack.stacksize() > 0 &&
               start_stack_size - stack.stacksize() < stack_arg_count) {
            stack.popClear();
        }

        while (auxstack.stacksize() > 0 &&
               start_auxstack_size - auxstack.stacksize() <
                   auxstack_arg_count) {
            auxstack.popClear();
        }
    }

    // If we're in the error state, jump to the error handler if one is set
    if (state == Status::Error && !errpc.is_error()) {
        pc = errpc.pc;
        state = Status::Extensive;
    }

    context.first_instruction = false;

    return blockReason;
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
        case OpCode::SIGNEXTEND:
            machineoperation::signExtend(*this);
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
        case OpCode::SHL:
            machineoperation::shl(*this);
            break;
        case OpCode::SHR:
            machineoperation::shr(*this);
            break;
        case OpCode::SAR:
            machineoperation::sar(*this);
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
        case OpCode::KECCAKF:
            machineoperation::keccakF(*this);
            break;
        case OpCode::SHA256F:
            machineoperation::sha256F(*this);
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
        case OpCode::XGET:
            machineoperation::xget(*this);
            break;
        case OpCode::XSET:
            machineoperation::xset(*this);
            break;
            /***********************/
            /*  Logging Operations */
            /***********************/
        case OpCode::BREAKPOINT:
            return machineoperation::breakpoint(*this);
        case OpCode::LOG:
            machineoperation::log(*this);
            break;
        case OpCode::DEBUG_PRINT:
            machineoperation::debug(*this);
            break;
            /**********************/
            /*  System Operations */
            /**********************/
        case OpCode::SEND: {
            machineoperation::send(*this);
            break;
        }
        case OpCode::INBOX:
            return machineoperation::inboxOp(*this);
        case OpCode::ERROR:
            state = Status::Error;
            break;
        case OpCode::HALT:
            state = Status::Halted;
            break;
        case OpCode::SET_GAS:
            machineoperation::setgas(*this);
            break;
        case OpCode::PUSH_GAS:
            machineoperation::pushgas(*this);
            break;
        case OpCode::ERR_CODE_POINT:
            machineoperation::errcodept(*this);
            break;
        case OpCode::PUSH_INSN:
            machineoperation::pushinsn(*this);
            break;
        case OpCode::PUSH_INSN_IMM:
            machineoperation::pushinsnimm(*this);
            break;
        case OpCode::SIDELOAD:
            return machineoperation::sideload(*this);
            break;
        case OpCode::NEW_BUFFER:
            machineoperation::newbuffer(*this);
            break;
        case OpCode::GET_BUFFER8:
            machineoperation::getbuffer8(*this);
            break;
        case OpCode::GET_BUFFER64:
            machineoperation::getbuffer64(*this);
            break;
        case OpCode::GET_BUFFER256:
            machineoperation::getbuffer256(*this);
            break;
        case OpCode::SET_BUFFER8:
            machineoperation::setbuffer8(*this);
            break;
        case OpCode::SET_BUFFER64:
            machineoperation::setbuffer64(*this);
            break;
        case OpCode::SET_BUFFER256:
            machineoperation::setbuffer256(*this);
            break;
        case OpCode::WASM_COMPILE:
            machineoperation::wasm_compile(*this);
            break;
        case OpCode::WASM_RUN:
            machineoperation::wasm_run(*this);
            break;
        /*****************/
        /*  Precompiles  */
        /*****************/
        case OpCode::ECRECOVER:
            machineoperation::ec_recover(*this);
            break;
        case OpCode::ECADD:
            machineoperation::ec_add(*this);
            break;
        case OpCode::ECMUL:
            machineoperation::ec_mul(*this);
            break;
        case OpCode::ECPAIRING:
            machineoperation::ec_pairing(*this);
            break;
        default:
            std::cerr << "Unhandled opcode <" << InstructionNames.at(opcode)
                      << ">" << std::hex << static_cast<int>(opcode);
            state = Status::Error;
    }

    return NotBlocked{};
}

std::ostream& operator<<(std::ostream& os, const MachineState& val) {
    os << "hash " << intx::to_string(val.hash(), 16) << "\n";
    os << "status " << static_cast<int>(val.state) << "\n";
    os << "pc " << val.pc << "\n";
    os << "data stack: " << val.stack << "\n";
    auto current_code_point = val.code->loadCodePoint(val.pc);
    os << "operation " << current_code_point.op << "\n";
    os << "codePointHash " << intx::to_string(hash(current_code_point), 16)
       << "\n";
    os << "stackHash " << intx::to_string(val.stack.hash(), 16) << "\n";
    os << "auxStackHash " << intx::to_string(val.auxstack.hash(), 16) << "\n";
    os << "registerHash " << intx::to_string(hash_value(val.registerVal), 16)
       << "\n";
    os << "staticHash " << intx::to_string(hash_value(val.static_val), 16)
       << "\n";
    os << "arb_gas_remaining " << val.arb_gas_remaining << "\n";
    os << "err handler " << val.errpc.pc << "\n";
    auto err_code_point = val.code->loadCodePoint(val.errpc.pc);
    os << "errHandlerHash " << intx::to_string(hash(err_code_point), 16)
       << "\n";
    return os;
}

uint256_t MachineState::getTotalMessagesRead() const {
    return output.fully_processed_inbox.count;
}
