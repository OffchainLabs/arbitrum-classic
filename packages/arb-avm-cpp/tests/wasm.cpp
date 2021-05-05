
#include <catch2/catch.hpp>
#include <avm_values/vmValueParser.hpp>
#include <avm_values/code.hpp>
#include <avm_values/value.hpp>
#include <avm/machinestate/machinestate.hpp>
#include <fstream>
#include <iostream>

#include <data_storage/arbstorage.hpp>
#include <avm/machinestate/runwasm.hpp>
#include <boost/algorithm/hex.hpp>

/*
value get_immed_value(uint8_t *a) {
    return 0;
}*/

value get_int_value(std::vector<uint8_t> bytes, uint64_t offset) {
    uint256_t acc = 0;
    for (int i = 0; i < 32; i++) {
        acc = acc*256;
        acc += bytes[offset+i];
    }
    return acc;
}

struct WasmCodepoint {
    value codept;
    value jump_table;
};

WasmCodepoint compile(std::vector<uint8_t>& bytes) {
    // code to hash
    auto code = std::make_shared<Code>(0);
    CodePointStub stub = code->addSegment();
    std::vector<value> labels;
    int i = 0;
    int num = 0;
    while (bytes[i] != 255) {
        OpCode opcode = static_cast<OpCode>(bytes[i]);
        i++;
        Operation op = {opcode};
        auto immed = bytes[i];
        i++;
        if (immed == 1) {
            op = {opcode, get_int_value(bytes, i)};
            i += 32;
        } else if (immed == 2) {
            std::vector<value> v;
            v.push_back(Buffer());
            v.push_back(0);
            v.push_back(Buffer());
            v.push_back(0);
            v.push_back(100000);  // check that these are the same
            op = {opcode, Tuple::createTuple(v)};
        } else if (immed == 3) {
            std::vector<value> v;
            v.push_back(Buffer());
            v.push_back(get_int_value(bytes, i));
            i += 32;
            op = {opcode, Tuple::createTuple(v)};
        }
        stub = code->addOperation(stub.pc, op);
        if (++num % 1000 == 0) {
            std::cerr << "Loaded " << num << " ops at " << i << "\n";
        }
        /*
        if (op.immediate) {
            std::cerr << "Immed hash " << op << " hash "
                      << intx::to_string(hash_value(*op.immediate), 16) << "\n";
        }
        std::cerr << "Loaded op " << op << " hash "
                  << intx::to_string(stub.hash, 16) << "\n";
        */
        if (bytes[i]) {
            // std::cerr << "Label " << stub << " at " << labels.size() <<
            // "\n";
            labels.push_back(stub);
        }
        i++;
    }

    std::reverse(labels.begin(), labels.end());
    auto table = make_table(labels);
    std::cerr << "Here " << intx::to_string(stub.hash, 16) << " "
              << labels.size() << " \n";
    // std::cerr << "Table " << table << " hash " <<
    // intx::to_string(hash_value(table), 16) << "\n";
    std::cerr << "Table hash " << intx::to_string(hash_value(table), 16)
              << " size " << getSize(table) << "\n";
    // convert table
    std::cerr << "Buffer hash " << intx::to_string(hash_value(Buffer()), 16)
              << "\n";
    return {stub, table};
}

TEST_CASE("wasm_compile") {
    /*
    SECTION("Compiler") {
        std::ifstream input("/home/sami/arbitrum/compiler.bin", std::ios::binary);

        std::vector<uint8_t> bytes((std::istreambuf_iterator<char>(input)), (std::istreambuf_iterator<char>()));

        input.close();
        compile(bytes);
    }
    */
    SECTION("Testing") {
        std::ifstream input("/home/sami/extra.bin", std::ios::binary);

        std::vector<uint8_t> bytes((std::istreambuf_iterator<char>(input)), (std::istreambuf_iterator<char>()));

        input.close();
        compile(bytes);
    }
}

TEST_CASE("Wasm") {
    SECTION("Code to hash") {
        /*

        auto res = run_wasm(Buffer(), 123);

        auto storage = ArbStorage("/home/sami/tmpstorage");
        // auto state = makeWasmMachine(123, Buffer());
        storage.initialize("/home/sami/arb-os/wasm-inst.json");

        auto arbcore = storage.getArbCore();
        arbcore->startThread();

        ValueCache value_cache{1, 0};
        auto cursor = arbcore->getExecutionCursor(10000000, value_cache);
        std::cerr << "Status: " << cursor.status.code() << "\n";
        std::cerr << "gas used: " << cursor.data->getOutput().arb_gas_used << "\n";
        std::cerr << "steps: " << cursor.data->getOutput().total_steps << "\n";
*/

        auto storage = ArbStorage("/home/sami/tmpstorage");
        auto state = makeWasmMachine(123, Buffer());
        storage.initialize(state);


        std::cerr << "Starting " << intx::to_string(state.hash().value(), 16) << "\n";

        uint256_t gasUsed = runWasmMachine(state);

        std::cerr << "Stopping " << intx::to_string(state.hash().value(), 16) << " gas used " << gasUsed << "\n";

        OneStepProof proof;
        state.marshalWasmProof(proof);
        std::cerr << "Made proof " << proof.buffer_proof.size() << "\n";
        marshal_uint256_t(gasUsed, proof.buffer_proof);
    }

}
