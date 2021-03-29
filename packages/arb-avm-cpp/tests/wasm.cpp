
#include <catch2/catch.hpp>
#include <avm_values/vmValueParser.hpp>
#include <fstream>
#include <iostream>

TEST_CASE("Wasm") {
    SECTION("Code to hash") {
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
        // auto op_count = json_code.size();
        // auto segment = std::make_shared<CodeSegment>(0);
        for (auto it = json_code.rbegin(); it != json_code.rend(); ++it) {
            Operation op = simple_operation_from_json(*it);
        }

        /*
        auto executable = loadExecutable("/home/sami/arb-os/foo2.json");
        auto code = std::make_shared<Code>(0);
        code->addSegment(std::move(executable.code));
        */
        /*
        std::ifstream i("/home/sami/arb-os/foo.json");
        nlohmann::json json_code;
        i >> json_code;
        std::vector<Operation> ops;
        auto op_count = json_code.size();
        auto segment = std::make_shared<CodeSegment>(0);
        for (auto it = json_code.rbegin(); it != json_code.rend(); ++it) {
            segment->addOperation(operation_from_json(*it, op_count, *segment));
        }
        */
        // Convert to hash
        // Make table
    }
}
