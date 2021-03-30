
#include <catch2/catch.hpp>
#include <avm_values/vmValueParser.hpp>
#include <avm_values/code.hpp>
#include <avm_values/value.hpp>
#include <avm/machinestate/machinestate.hpp>
#include <fstream>
#include <iostream>

const int LEVEL = 5;

value table_to_tuple2(std::vector<value> tab, int prefix, int shift, int level) {
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
        v.push_back(table_to_tuple2(tab, prefix + (i << shift), shift + 3, level - 1));
    }
    return Tuple::createTuple(v);
}

value make_table(std::vector<value> tab) {
    return table_to_tuple2(tab, 0, 0, LEVEL-1);
}

uint256_t runMachine(MachineState &machine_state) {
    uint256_t start_steps = machine_state.output.total_steps;
    uint256_t start_gas = machine_state.output.arb_gas_used;

    bool has_gas_limit = machine_state.context.max_gas != 0;
    BlockReason block_reason = NotBlocked{};
    uint256_t initialConsumed = machine_state.getTotalMessagesRead();
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

        auto op = machine_state.loadCurrentInstruction();
        std::cerr << "op " << op << " state " << int(machine_state.state) << "\n";
        if (machine_state.stack.stacksize() > 0 && !std::get_if<Tuple>(&machine_state.stack[0])) {
            std::cerr << "stack top " << machine_state.stack[0] << "\n";
        }

        block_reason = machine_state.runOne();
        if (!std::get_if<NotBlocked>(&block_reason)) {
            break;
        }
    }    
    return start_gas - machine_state.arb_gas_remaining;
}

TEST_CASE("Wasm") {
    SECTION("Code to hash") {
        std::ifstream labels_input_stream("/home/sami/arb-os/labels.json");
        if (!labels_input_stream.is_open()) {
            throw std::runtime_error("doesn't exist");
        }
        nlohmann::json labels_json;
        labels_input_stream >> labels_json;
        std::vector<bool> has_labels;
        for (auto elem : labels_json) {
            has_labels.push_back(elem.get<int>() == 1);
        }
        // Load JSON
        std::ifstream executable_input_stream("/home/sami/arb-os/foo2.json");
        if (!executable_input_stream.is_open()) {
            throw std::runtime_error("doesn't exist");
        }
        nlohmann::json executable_json;
        executable_input_stream >> executable_json;
        auto& json_code = executable_json.at("code");
        if (!json_code.is_array()) {
            throw std::runtime_error("expected code to be array");
        }
        auto code = std::make_shared<Code>(0);
        CodePointStub stub = code->addSegment();
        std::vector<value> labels;
        int idx = json_code.size();
        for (auto it = json_code.rbegin(); it != json_code.rend(); ++it) {
            Operation op = simple_operation_from_json(*it);
            stub = code->addOperation(stub.pc, op);
            idx--;
            std::cerr << "Loaded op " << op << " idx " << idx << "\n";
            if (has_labels[idx]) {
                std::cerr << "Label " << stub << " at " << labels.size() << "\n";
                labels.push_back(stub);
            }
        }

        std::reverse(labels.begin(), labels.end());

        auto table = make_table(labels);
        std::cerr << "Here " << intx::to_string(stub.hash, 16) << " " << labels.size() << " \n";
        std::cerr << "Table " << table << " hash " << intx::to_string(hash_value(table), 16) << "\n";
        // std::cerr << "Table hash " << intx::to_string(hash_value(table), 16) << "\n";
        MachineState state(code, 0);
        state.stack.push(123);
        state.stack.push(Buffer());
        state.stack.push(std::move(table));

        std::cerr << "Starting " << intx::to_string(state.hash().value(), 16) << "\n";

        uint256_t gasUsed = runMachine(state);

        std::cerr << "Stopping " << intx::to_string(state.hash().value(), 16) << " gas used " << gasUsed << "\n";

        OneStepProof proof;
        state.marshalWasmProof(proof);
        std::cerr << "Made proof " << proof.buffer_proof.size() << "\n";
    }

}
