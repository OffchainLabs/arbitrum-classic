
#include <catch2/catch.hpp>
#include <avm_values/vmValueParser.hpp>
#include <avm_values/code.hpp>
#include <avm_values/value.hpp>
#include <avm/machinestate/machinestate.hpp>
#include <fstream>
#include <iostream>

#include <data_storage/arbstorage.hpp>

TEST_CASE("Wasm") {
    SECTION("Code to hash") {

        /*

        auto storage = ArbStorage("/home/sami/tmpstorage");
        auto state = makeWasmMachine(123, Buffer());
        storage.initialize(state);
        */

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

        /*

        std::cerr << "Starting " << intx::to_string(state.hash().value(), 16) << "\n";

        uint256_t gasUsed = runWasmMachine(state);

        std::cerr << "Stopping " << intx::to_string(state.hash().value(), 16) << " gas used " << gasUsed << "\n";

        OneStepProof proof;
        state.marshalWasmProof(proof);
        std::cerr << "Made proof " << proof.buffer_proof.size() << "\n";
        marshal_uint256_t(gasUsed, proof.buffer_proof);
        */
    }

}
