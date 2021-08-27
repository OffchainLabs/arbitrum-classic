
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

#include "config.hpp"
#include "helper.hpp"

std::vector<uint8_t> getFile(std::string fname) {
    std::ifstream input(fname, std::ios::binary);
    std::vector<uint8_t> bytes((std::istreambuf_iterator<char>(input)), (std::istreambuf_iterator<char>()));
    input.close();
    return bytes;
}

MachineState mkWasmMachine(WasmResult res, std::string fname) {
    CodeResult cres = wasmAvmToCode(res);

    MachineState state(cres.code, 0);
    auto arg_buf = getFile(fname);
    state.stack.push(arg_buf.size());
    state.stack.push(vec2buf(arg_buf));
    state.stack.push(std::move(cres.table));
    state.stack.push(0);
    state.arb_gas_remaining = 1000000000000;
    state.output.arb_gas_used = 0;

    return state;

}

TEST_CASE("Wasm") {
    SECTION("Making compiler machine") {
        RunWasm runner("/home/sami/arbitrum/compiler.wasm");
        auto buf = getFile("/home/sami/arbitrum/compiler.wasm");
        auto res = runner.run_wasm(vec2buf(buf), buf.size());

        std::string test_file = "/home/sami/arbitrum/test.wasm";

        auto m = mkWasmMachine(res, test_file);
        auto start = std::chrono::system_clock::now();
        runWasmMachine(m);
        auto end = std::chrono::system_clock::now();

        std::cerr << "Result stack " << m.stack[0] << "\n";
        std::cerr << "Result stack " << m.stack[1] << "\n";
        std::cerr << "Result stack " << m.stack[2] << "\n";

        std::cerr << "Table " << hash_value(m.stack[4]) << " \n";
        std::cerr << "Codepoint " << hash_value(m.stack[3]) << " \n";

        std::chrono::duration<double> elapsed_seconds = end-start;

        std::cerr << "elapsed time: " << elapsed_seconds.count() << "s\n";

        auto buf2 = getFile(test_file);
        auto res2 = runner.run_wasm(vec2buf(buf2), buf2.size());

        CodeResult cres = wasmAvmToCode(res2);
        std::cerr << "Made table " << hash_value(cres.table) << " \n";
        std::cerr << "Codepoint " << hash_value(cres.stub) << " \n";

        REQUIRE(hash_value(m.stack[4]) == hash_value(cres.table));
        REQUIRE(hash_value(m.stack[3]) == hash_value(cres.stub));

    }

}

