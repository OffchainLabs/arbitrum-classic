
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
    SECTION("JIT converter") {
        RunWasm runner0("/home/sami/arbitrum/compiler.wasm");
        auto runner = runner0;
        // RunWasm runner("/home/sami/wasm-hash/pkg/wasm_hash_bg.wasm");
        // auto m0 = MachineState();
        // auto m = m0;
        // auto buf = getFile("/home/sami/arb-os/wasm-tests/test-buffer.wasm");
        // auto buf = getFile("/home/sami/wasm-hash/pkg/wasm_hash_bg.wasm");
        auto buf = getFile("/home/sami/arbitrum/compiler.wasm");
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
        auto wasmcp = wasmAvmToCodePoint(res, buf);
        if (hash_value(wasmcp.data->get_element(0)) != hash1 || hash_value(wasmcp.data->get_element(1)) != hash2) {
            std::cerr << "FAIL\n";
        }
    }

    SECTION("Making machine") {
        RunWasm runner0("/home/sami/arbitrum/compiler.wasm");
        auto runner = runner0;
        // RunWasm runner("/home/sami/wasm-hash/pkg/wasm_hash_bg.wasm");
        // auto m0 = MachineState();
        // auto m = m0;
        // auto buf = getFile("/home/sami/arb-os/wasm-tests/test-buffer.wasm");
        auto buf = getFile("/home/sami/simple-wasm/pkg/simple_wasm_bg.wasm");
        // auto buf = getFile("/home/sami/arbitrum/compiler.wasm");
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
        
        auto m = makeWasmMachine(res, 0, Buffer(), 0);
        runWasmMachine(m);

        std::cerr << "Result stack " << m.stack[0] << "\n";
        std::cerr << "Result stack " << m.stack[1] << "\n";
        std::cerr << "Result stack " << m.stack[2] << "\n";

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

MachineState mkWasmMachine(WasmResult res) {
    auto code = std::make_shared<Code>(0);
    CodePointStub stub = code->addSegment();

    std::vector<CodePointStub> points;
    std::vector<value> tab_lst;

    for (int i = 0; i < res.insn->size(); i++) {
        points.push_back(stub);
        stub = code->addOperation(stub.pc, (*res.insn)[i]);
        // std::cerr << i << ": " << stub << " " << (*res.insn)[i] << "\n";
    }

    for (int i = 0; i < res.table.size(); i++) {
        auto offset = res.table[i].first;
        if (offset >= tab_lst.size()) {
            tab_lst.resize(offset+1);
        }
        tab_lst[offset] = points[res.table[i].second];
        // tab_lst[offset] = points[points.size() - res.table[i].second];
    }
    auto table = make_table(tab_lst);

    std::cerr << "Made table " << hash_value(table) << " \n";
    std::cerr << "Codepoint " << hash_value(stub) << " \n";

    MachineState state(code, 0);
    auto arg_buf = getFile("/home/sami/arb-os/wasm-tests/test-buffer.wasm");
    state.stack.push(arg_buf.size());
    state.stack.push(vec2buf(arg_buf));
    state.stack.push(std::move(table));
    state.stack.push(0);
    state.arb_gas_remaining = 1000000000000;
    state.output.arb_gas_used = 0;

    return state;

}

TEST_CASE("wasm_rlp") {
    SECTION("Making compiler machine") {
        RunWasm runner("/home/sami/rlp.wasm");
    }
}

TEST_CASE("wasm_3") {
    SECTION("Making compiler machine") {
        RunWasm runner("/home/sami/wasm2avm/pkg/wasm2avm_bg.wasm");
        auto buf = getFile("/home/sami/arb-os/wasm-tests/test-buffer.wasm");
        auto res = runner.run_wasm(vec2buf(buf), buf.size());

        std::cerr << "Here\n";

        auto code = std::make_shared<Code>(0);
        CodePointStub stub = code->addSegment();

        std::vector<CodePointStub> points;
        std::vector<value> tab_lst;

        for (int i = 0; i < res.insn->size(); i++) {
            points.push_back(stub);
            stub = code->addOperation(stub.pc, (*res.insn)[i]);
            // std::cerr << i << ": " << stub << " " << (*res.insn)[i] << "\n";
        }

        for (int i = 0; i < res.table.size(); i++) {
            auto offset = res.table[i].first;
            if (offset >= tab_lst.size()) {
                tab_lst.resize(offset+1);
            }
            tab_lst[offset] = points[res.table[i].second];
            // tab_lst[offset] = points[points.size() - res.table[i].second];
        }
        auto table = make_table(tab_lst);

        std::cerr << "Made table " << hash_value(table) << " \n";
        std::cerr << "Codepoint " << hash_value(stub) << " \n";

        /*
        MachineState state(code, 0);
        auto arg_buf = getFile("/home/sami/arb-os/wasm-tests/test-buffer.wasm");
        state.stack.push(arg_buf.size());
        state.stack.push(vec2buf(arg_buf));
        state.stack.push(std::move(table));
        state.arb_gas_remaining = 1000000000000;
        state.output.arb_gas_used = 0;

        runWasmMachine(state);
        */

        /*
        auto bytes = buf2vec(res.buffer, res.buffer_len);
        uint256_t hash1 = intx::be::unsafe::load<uint256_t>(bytes.data());
        uint256_t hash2 = intx::be::unsafe::load<uint256_t>(bytes.data()+32);
        std::cerr << "Result len " << bytes.size() << "\n";
        std::string hexstr;
        hexstr.resize(bytes.size()*2);
        boost::algorithm::hex(bytes.begin(), bytes.end(), hexstr.begin());
        std::cerr << "Result hash " << hexstr << "\n";
        std::cerr << "Result hash " << intx::to_string(hash1, 16) << ", " << intx::to_string(hash2, 16) << "\n";
        
        auto arg_buf = getFile("/home/sami/simple-wasm/pkg/simple_wasm_bg.wasm");
        // auto arg_buf = getFile("/home/sami/arb-os/wasm-tests/test-buffer2.wasm");
        auto m = makeWasmMachine(res.extra, arg_buf.size(), vec2buf(arg_buf));
        auto start = std::chrono::system_clock::now();
        runWasmMachine(m);
        auto end = std::chrono::system_clock::now();

        std::cerr << "Result stack " << m.stack[0] << "\n";
        std::cerr << "Result stack " << m.stack[1] << "\n";
        std::cerr << "Result stack " << m.stack[2] << "\n";

        std::chrono::duration<double> elapsed_seconds = end-start;
        std::time_t end_time = std::chrono::system_clock::to_time_t(end);

        std::cerr << "elapsed time: " << elapsed_seconds.count() << "s\n";        
        */
    }

}

TEST_CASE("wasm_4") {
    SECTION("Making compiler machine") {
        RunWasm runner("/home/sami/arbitrum/compiler.wasm");
        // RunWasm runner("/home/sami/wasm2avm/pkg/wasm2avm_bg.wasm");
        // RunWasm runner("/home/sami/complete.wasm");
        // auto buf = getFile("/home/sami/stripped.wasm");
        auto buf = getFile("/home/sami/wasm2avm/pkg/wasm2avm_bg.wasm");
        // auto buf = getFile("/home/sami/simple-wasm/pkg/simple_wasm_bg.wasm");
        auto res = runner.run_wasm(vec2buf(buf), buf.size());

        auto m = mkWasmMachine(res);
        auto start = std::chrono::system_clock::now();
        runWasmMachine(m);
        auto end = std::chrono::system_clock::now();

        std::cerr << "Result stack " << m.stack[0] << "\n";
        std::cerr << "Result stack " << m.stack[1] << "\n";
        std::cerr << "Result stack " << m.stack[2] << "\n";
        // std::cerr << "Result stack " << m.stack[3] << "\n";
        // std::cerr << "Result stack " << m.stack[4] << "\n";

        std::cerr << "Table " << hash_value(m.stack[4]) << " \n";
        std::cerr << "Codepoint " << hash_value(m.stack[3]) << " \n";

        std::chrono::duration<double> elapsed_seconds = end-start;
        std::time_t end_time = std::chrono::system_clock::to_time_t(end);

        std::cerr << "elapsed time: " << elapsed_seconds.count() << "s\n";        runWasmMachine(m);
    }

}

/*
TEST_CASE("wasm_2") {
    SECTION("Making compiler machine") {
        RunWasm runner("/home/sami/arbitrum/compiler.wasm");
        // auto buf = getFile("/home/sami/stripped.wasm");
        auto buf = getFile("/home/sami/wasm2avm/pkg/wasm2avm_bg.wasm");
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
        
        // auto arg_buf = getFile("/home/sami/simple-wasm/pkg/simple_wasm_bg.wasm");
        auto arg_buf = getFile("/home/sami/arb-os/wasm-tests/test-buffer.wasm");
        auto m = makeWasmMachine(res.extra, arg_buf.size(), vec2buf(arg_buf));
        auto start = std::chrono::system_clock::now();
        runWasmMachine(m);
        auto end = std::chrono::system_clock::now();

        std::cerr << "Result stack " << m.stack[0] << "\n";
        std::cerr << "Result stack " << m.stack[1] << "\n";
        std::cerr << "Result stack " << m.stack[2] << "\n";
        std::cerr << "Result stack " << m.stack[3] << "\n";
        std::cerr << "Result stack " << m.stack[4] << "\n";

        std::cerr << "Table " << hash_value(m.stack[3]) << " \n";
        std::cerr << "Codepoint " << hash_value(m.stack[4]) << " \n";

        std::chrono::duration<double> elapsed_seconds = end-start;
        std::time_t end_time = std::chrono::system_clock::to_time_t(end);

        std::cerr << "elapsed time: " << elapsed_seconds.count() << "s\n";        runWasmMachine(m);
    }

}

TEST_CASE("Wasm") {
    SECTION("Code to hash") {
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
*/
