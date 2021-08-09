/*
 * Copyright 2021, Offchain Labs, Inc.
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

#include "helper.hpp"

Machine getComplexMachine() {
    auto core_code = std::make_shared<CoreCode>();
    auto code = std::make_shared<RunningCode>(core_code);
    auto stub = code->addSegment();
    stub = code->addOperation(stub.pc, Operation(OpCode::ADD));
    stub = code->addOperation(stub.pc, Operation(OpCode::MUL));
    stub = code->addOperation(stub.pc, Operation(OpCode::SUB));
    uint256_t register_val = 100;
    auto static_val = Tuple(register_val, Tuple());

    CodePointStub code_point_stub({0, 1}, 654546);

    Datastack data_stack;
    data_stack.push(register_val);
    Datastack aux_stack;
    aux_stack.push(register_val);
    aux_stack.push(code_point_stub);

    uint256_t arb_gas_remaining = 534574678365;

    CodePointRef pc{0, 0};
    CodePointStub err_pc({0, 0}, 968769876);
    Status state = Status::Extensive;

    auto output =
        MachineOutput{{42, 54}, 23, 54, 12, 65, 76, 43, 65, 72, 73, 74};

    return Machine(MachineState(output, pc, std::move(code), register_val,
                                std::move(static_val), data_stack, aux_stack,
                                arb_gas_remaining, state, err_pc));
}

Machine getDefaultMachine() {
    auto core_code = std::make_shared<CoreCode>();
    auto code = std::make_shared<RunningCode>(core_code);
    code->addSegment();
    auto static_val = Tuple();
    auto register_val = Tuple();
    auto data_stack = Tuple();
    auto aux_stack = Tuple();
    uint256_t arb_gas_remaining = 534574678365;
    CodePointRef pc(0, 0);
    CodePointStub err_pc({0, 0}, 968769876);
    Status state = Status::Extensive;
    auto output =
        MachineOutput{{42, 54}, 23, 54, 12, 65, 76, 43, 34, 72, 73, 74};
    return Machine(MachineState(output, pc, std::move(code), register_val,
                                std::move(static_val), data_stack, aux_stack,
                                arb_gas_remaining, state, err_pc));
}