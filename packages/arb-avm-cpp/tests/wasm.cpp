
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

MachineState mkWasmMachine(WasmResult res, std::vector<uint8_t> arg_buf) {
    CodeResult cres = wasmAvmToCode(res);

    MachineState state(cres.code, 0);
    state.stack.push(arg_buf.size());
    state.stack.push(vec2buf(arg_buf));
    state.stack.push(std::move(cres.table));
    std::vector<uint8_t> b1 = {1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1};
    std::vector<uint8_t> b2 = {2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2};
    auto tpl = Tuple(
        123,
        vec2buf(b1),
        Tuple(vec2buf(b2), 234, 234),
        12345678
    );
    state.stack.push(tpl);
    state.arb_gas_remaining = 1000000000000;
    state.output.arb_gas_used = 0;

    return state;

}

MachineState mkWasmMachine(WasmResult res, std::string fname) {
    return mkWasmMachine(res, getFile(fname));
}

TEST_CASE("Wasm") {
    SECTION("Making compiler machine") {
        RunWasm runner(wasm_compile_path);
        auto buf = getFile(wasm_compile_path);
        auto res = runner.run_wasm(vec2buf(buf), buf.size());

        std::string test_file = wasm_test_path;

        auto m = mkWasmMachine(res, test_file);
        std::cerr << "Running machine " << "\n";
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
    /* This test has problems because signal handling conflict with the test library.
    SECTION("Test env functions") {
        for (int i = 0; i < 11; i++) {
            RunWasm runner(wasm_compile_path);
            auto buf = getFile(wasm_env_path);
            auto res = runner.run_wasm(vec2buf(buf), buf.size());

            auto test_buf = std::vector<uint8_t>();
            test_buf.push_back(i);
            for (int j = 1; j < 128; j++) {
                test_buf.push_back(j);
            }

            auto m = mkWasmMachine(res, test_buf);
            // std::cerr << "Running machine " << "\n";
            auto start = std::chrono::system_clock::now();
            runWasmMachine(m);
            auto end = std::chrono::system_clock::now();

            uint256_t err_code = intx::from_string<uint256_t>("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff");

            bool err = false;
            // std::cerr << "Result stack " << m.stack[0] << "\n";
            if (m.stack[0] == value(err_code)) {
                // std::cerr << "Error at test " << i << "\n";
                err = true;
            }

            std::chrono::duration<double> elapsed_seconds = end-start;

            // std::cerr << "elapsed time: " << elapsed_seconds.count() << "s\n";

            std::vector<uint8_t> b1 = {1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1};
            std::vector<uint8_t> b2 = {2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2};
            auto tpl = Tuple(
                123,
                vec2buf(b1),
                Tuple(vec2buf(b2), 234, 234),
                12345678
            );
            RunWasm fail_runner(wasm_env_path);
            auto res2 = fail_runner.run_wasm(vec2buf(test_buf), test_buf.size(), tpl);
            REQUIRE(res2.error == err);
            if (!err) {
                REQUIRE(m.stack[0] == value(res2.buffer_len));
                REQUIRE(m.stack[1] == value(res2.buffer));
            }

        }
    }
        */

}

