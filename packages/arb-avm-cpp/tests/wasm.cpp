
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

std::vector<uint8_t> getFile(std::string fname) {
    std::ifstream input(fname, std::ios::binary);
    std::vector<uint8_t> bytes((std::istreambuf_iterator<char>(input)), (std::istreambuf_iterator<char>()));
    input.close();
    return bytes;
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
    SECTION("JIT converter") {
        RunWasm runner0("/home/sami/arbitrum/compiler.wasm");
        auto runner = runner0;
        // RunWasm runner("/home/sami/wasm-hash/pkg/wasm_hash_bg.wasm");
        // auto m0 = MachineState();
        // auto m = m0;
        auto buf = getFile("/home/sami/arb-os/wasm-tests/test-buffer.wasm");
        auto res = runner.run_wasm(vec2buf(buf), buf.size());
        auto bytes = buf2vec(res.buffer, res.buffer_len);
        uint256_t hash1 = intx::be::unsafe::load<uint256_t>(bytes.data());
        uint256_t hash2 = intx::be::unsafe::load<uint256_t>(bytes.data()+32);
        std::cerr << "Result len " << bytes.size() << "\n";
        std::string hexstr;
        hexstr.resize(bytes.size()*2);
        boost::algorithm::hex(bytes.begin(), bytes.end(), hexstr.begin());
        std::cerr << "Result hash " << hexstr << "\n";
        std::cerr << "Result hash " << intx::to_string(hash1, 16) << ", " << intx::to_string(hash2, 16) << "\n";
        auto wasmcp = wasmAvmToCodePoint(res.extra, buf);
        if (hash_value(wasmcp.data->get_element(0)) != hash1 || hash_value(wasmcp.data->get_element(1)) != hash2) {
            std::cerr << "FAIL\n";
        }
    }

    /*

    SECTION("Testing") {
        std::ifstream input("/home/sami/extra.bin", std::ios::binary);

        std::vector<uint8_t> bytes((std::istreambuf_iterator<char>(input)), (std::istreambuf_iterator<char>()));

        input.close();
        std::vector<uint8_t> asd;
        wasmAvmToCodePoint(bytes, asd);
    }
    */
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
