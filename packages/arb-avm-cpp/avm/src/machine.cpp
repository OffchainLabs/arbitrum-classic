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

#include <sys/stat.h>
#include <fstream>
#include <iostream>

#include <avm/machine.hpp>
#include <avm_values/opcodes.hpp>
#include <avm_values/util.hpp>
#include <bigint_utils.hpp>
#include <data_storage/checkpoint/checkpointstorage.hpp>
#include <data_storage/checkpoint/machinestatedeleter.hpp>

std::ostream& operator<<(std::ostream& os, const MachineState& val) {
    os << "status " << static_cast<int>(val.state) << "\n";
    os << "codePointHash " << to_hex_str(hash(val.code->code[val.pc])) << "\n";
    os << "stackHash " << to_hex_str(val.stack.hash()) << "\n";
    os << "auxStackHash " << to_hex_str(val.auxstack.hash()) << "\n";
    os << "registerHash " << to_hex_str(hash(val.registerVal)) << "\n";
    os << "staticHash " << to_hex_str(hash(val.code->staticVal)) << "\n";
    os << "errHandlerHash " << to_hex_str(hash(val.errpc)) << "\n";
    return os;
}

std::ostream& operator<<(std::ostream& os, const Machine& val) {
    os << val.machine_state;
    return os;
}

bool Machine::initializeMachine(const std::string& filename) {
    return machine_state.initialize_machinestate(filename);
}

void Machine::initializeMachine(const MachineState& initial_state) {
    machine_state = initial_state;
}

Assertion Machine::run(uint64_t stepCount,
                       uint256_t timeBoundStart,
                       uint256_t timeBoundEnd,
                       Tuple messages,
                       std::chrono::seconds wallLimit) {
    bool has_time_limit = wallLimit.count() != 0;
    auto start_time = std::chrono::system_clock::now();
    machine_state.context = AssertionContext{
        TimeBounds{{timeBoundStart, timeBoundEnd}}, std::move(messages)};
    while (machine_state.context.numSteps < stepCount) {
        auto blockReason = runOne();
        if (!nonstd::get_if<NotBlocked>(&blockReason)) {
            break;
        }
        if (has_time_limit && machine_state.context.numSteps % 10000 == 0) {
            auto end_time = std::chrono::system_clock::now();
            auto run_time = end_time - start_time;
            if (run_time >= wallLimit) {
                break;
            }
        }
    }
    return {machine_state.context.numSteps, machine_state.context.numGas,
            std::move(machine_state.context.outMessage),
            std::move(machine_state.context.logs),
            machine_state.context.didInboxInsn};
}

SaveResults Machine::checkpoint(CheckpointStorage& storage) {
    return machine_state.checkpointState(storage);
}

bool Machine::restoreCheckpoint(
    const CheckpointStorage& storage,
    const std::vector<unsigned char>& checkpoint_key) {
    return machine_state.restoreCheckpoint(storage, checkpoint_key);
}

DeleteResults Machine::deleteCheckpoint(CheckpointStorage& storage) {
    std::vector<unsigned char> checkpoint_key;
    marshal_uint256_t(hash(), checkpoint_key);

    return ::deleteCheckpoint(storage, checkpoint_key);
}

struct TupleTree {
    bool is_read = false;
    std::shared_ptr<TupleTree> children[8] = {
        nullptr, nullptr, nullptr, nullptr, nullptr, nullptr, nullptr, nullptr};
};

MachineState Machine::trustlessCall(uint64_t steps,
                                    uint64_t& copy_start,
                                    uint64_t& aux_copy_start) {
    auto copyMachine = machine_state;
    copy_start = copyMachine.stack.stacksize();
    aux_copy_start = copyMachine.auxstack.stacksize();
    std::shared_ptr<TupleTree> original_register_contents =
        std::make_shared<TupleTree>(TupleTree());
    std::shared_ptr<TupleTree> current_register_contents =
        std::make_shared<TupleTree>(*original_register_contents);
    Operation current_op;

    std::vector<std::shared_ptr<TupleTree>> original_stack_contents(copy_start);
    for (auto& tree : original_stack_contents) {
        tree = std::make_shared<TupleTree>();
    }
    std::vector<std::shared_ptr<TupleTree>> current_stack_contents;

    std::vector<std::shared_ptr<TupleTree>> aux_stack_contents(aux_copy_start);
    for (auto& tree : aux_stack_contents) {
        tree = std::make_shared<TupleTree>();
    }
    std::vector<std::shared_ptr<TupleTree>> current_aux_contents;

    current_stack_contents.reserve(original_stack_contents.size());
    for (auto& tree : original_stack_contents) {
        current_stack_contents.push_back(std::make_shared<TupleTree>(*tree));
    }
    current_aux_contents.reserve(aux_stack_contents.size());
    for (auto& tree : aux_stack_contents) {
        current_aux_contents.push_back(std::make_shared<TupleTree>(*tree));
    }

    for (uint64_t i = steps; i > 0; i--) {
        current_op = copyMachine.code->code[copyMachine.pc].op;
        uint256_t* index;
        uint64_t read_depth = 1;
        switch (current_op.opcode) {
            case OpCode::TGET:
                if (current_op.immediate) {
                    index = nonstd::get_if<uint256_t>(&*current_op.immediate);
                } else {
                    current_stack_contents.back()->is_read = true;
                    index = nonstd::get_if<uint256_t>(&copyMachine.stack[0]);
                    current_stack_contents.pop_back();
                }
                if (index && *index < 8) {
                    if (current_stack_contents.back()
                            ->children[(uint64_t)*index] == nullptr) {
                        // current_stack_contents.back()->is_read = true;
                        current_stack_contents.back()
                            ->children[(uint64_t)*index] =
                            std::make_shared<TupleTree>();
                    }
                    current_stack_contents.back() =
                        current_stack_contents.back()
                            ->children[(uint64_t)*index];
                }
                read_depth = 1;
                break;
            case OpCode::TSET:
                if (current_op.immediate) {
                    index = nonstd::get_if<uint256_t>(&*current_op.immediate);
                } else {
                    current_stack_contents.back()->is_read = true;
                    index = nonstd::get_if<uint256_t>(&copyMachine.stack[0]);
                    current_stack_contents.pop_back();
                }
                current_stack_contents.back()->is_read = true;
                (**(current_stack_contents.end() - 1)).is_read = true;
                for (auto& subtree :
                     (**(current_stack_contents.end() - 1)).children) {
                    if (!subtree) {
                        *subtree = TupleTree();
                    }
                }
                if (index && *index < 8) {
                    TupleTree new_tuple = **(current_stack_contents.end() - 1);
                    new_tuple.children[(uint64_t)*index] =
                        current_stack_contents.back();
                    current_stack_contents.pop_back();
                    current_stack_contents.back() =
                        std::make_shared<TupleTree>(new_tuple);
                }
                read_depth = 1;
                break;
            case OpCode::EQ:
                if (!current_op.immediate) {
                    current_stack_contents.pop_back();
                }
                current_stack_contents.back() = std::make_shared<TupleTree>();
                read_depth = 1;
                break;
            case OpCode::TLEN:
            case OpCode::HASH:
            case OpCode::TYPE:
                if (!current_op.immediate) {
                    current_stack_contents.back() =
                        std::make_shared<TupleTree>();
                }
                read_depth = 1;
                break;
            case OpCode::AUXPUSH:
                if (!current_op.immediate) {
                    current_aux_contents.push_back(
                        std::move(current_stack_contents.back()));
                } else {
                    current_aux_contents.push_back(
                        std::make_shared<TupleTree>());
                }
                current_stack_contents.pop_back();
                read_depth = 0;
                break;
            case OpCode::AUXPOP:
                if (!current_op.immediate) {
                    current_stack_contents.push_back(
                        std::move(current_aux_contents.back()));
                } else {
                    current_stack_contents.push_back(
                        std::make_shared<TupleTree>());
                }
                read_depth = 0;
                current_aux_contents.pop_back();
                break;
            case OpCode::DUP0:
                if (!current_op.immediate) {
                    // current_stack_contents.back()->is_read = true;
                    current_stack_contents.push_back(
                        current_stack_contents.back());
                }
                read_depth = 2;
                break;
            case OpCode::DUP1:
                if (!current_op.immediate) {
                    (*(current_stack_contents.end() - 1))->is_read = true;
                    current_stack_contents.push_back(
                        *(current_stack_contents.end() - 1));
                }
                read_depth = 3;
                break;
            case OpCode::DUP2:
                if (!current_op.immediate) {
                    (*(current_stack_contents.end() - 2))->is_read = true;
                    current_stack_contents.push_back(
                        *(current_stack_contents.end() - 2));
                }
                read_depth = 4;
                break;
            case OpCode::SWAP1:
                if (!current_op.immediate) {
                    current_stack_contents.back().swap(
                        *(current_stack_contents.end() - 1));
                    current_stack_contents.back()->is_read = true;
                    (*(current_stack_contents.end() - 1))->is_read = true;
                } else {
                    current_stack_contents.back() =
                        std::make_shared<TupleTree>();
                }
                read_depth = 2;
                break;
            case OpCode::SWAP2:
                if (!current_op.immediate) {
                    std::swap(current_stack_contents.back(),
                              *(current_stack_contents.end() - 2));
                    current_stack_contents.back()->is_read = true;
                    (*(current_stack_contents.end() - 2))->is_read = true;
                } else {
                    *(current_stack_contents.end() - 1) =
                        std::make_shared<TupleTree>();
                }
                read_depth = 3;
                break;
            case OpCode::RPUSH:
                current_register_contents->is_read = true;
                current_stack_contents.push_back(current_register_contents);
                break;
            case OpCode::RSET:
                current_stack_contents.back()->is_read = true;
                current_register_contents = current_stack_contents.back();
                break;
            default:
                break;
        }
        copyMachine.runOne();
        if (copyMachine.stack.stacksize() != current_stack_contents.size()) {
            current_stack_contents.resize(copyMachine.stack.stacksize(),
                                          std::make_shared<TupleTree>());
        }
        if (copyMachine.stack.stacksize() - read_depth < copy_start) {
            copy_start = copyMachine.stack.stacksize() - read_depth;
        }
        if (copyMachine.auxstack.stacksize() < aux_copy_start) {
            aux_copy_start = copyMachine.auxstack.stacksize();
        }
    }
    auto outputMachine = machine_state;
    outputMachine.stack.values =
        std::vector<value>(machine_state.stack.values.begin() + copy_start,
                           machine_state.stack.values.end());
    for (uint64_t i = 0; i < outputMachine.stack.stacksize(); i++) {
        if (original_stack_contents[i + copy_start]->is_read) {
            uint256_t old_hash = ::hash(outputMachine.stack.values[i]);
            outputMachine.stack.values[i] = HashOnly();
            nonstd::get_if<HashOnly>(&outputMachine.stack.values[i])->value =
                old_hash;
        }
    }
    outputMachine.stack.hashes = std::vector<uint256_t>(
        machine_state.stack.hashes.begin(), machine_state.stack.hashes.end());
    outputMachine.auxstack.values =
        std::vector<value>(machine_state.auxstack.values.begin() + copy_start,
                           machine_state.auxstack.values.end());
    for (uint64_t i = 0; i < outputMachine.auxstack.stacksize(); i++) {
        if (aux_stack_contents[i + copy_start]->is_read) {
            uint256_t old_hash = ::hash(outputMachine.auxstack.values[i]);
            outputMachine.auxstack.values[i] = HashOnly();
            nonstd::get_if<HashOnly>(&outputMachine.auxstack.values[i])->value =
                old_hash;
        }
    }
    outputMachine.auxstack.hashes =
        std::vector<uint256_t>(machine_state.auxstack.hashes.begin(),
                               machine_state.auxstack.hashes.end());
    return outputMachine;
}

void Machine::glueIn(MachineState state,
                     uint64_t stack_start,
                     uint64_t aux_start) {
    machine_state.stack.values.resize(stack_start);
    machine_state.stack.values.insert(machine_state.stack.values.end(),
                                      state.stack.values.begin(),
                                      state.stack.values.end());
    machine_state.stack.hashes.insert(machine_state.stack.hashes.end(),
                                      state.stack.hashes.begin(),
                                      state.stack.hashes.end());
    machine_state.auxstack.values.resize(aux_start);
    machine_state.auxstack.values.insert(machine_state.auxstack.values.end(),
                                         state.auxstack.values.begin(),
                                         state.auxstack.values.end());
    machine_state.auxstack.hashes.insert(machine_state.auxstack.hashes.end(),
                                         state.auxstack.hashes.begin(),
                                         state.auxstack.hashes.end());
    std::swap(machine_state.stack, state.stack);
    std::swap(machine_state.auxstack, state.auxstack);
    machine_state = std::move(state);
}
