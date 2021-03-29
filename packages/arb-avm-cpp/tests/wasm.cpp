
#include <catch2/catch.hpp>
#include <avm_values/vmValueParser.hpp>
#include <avm_values/code.hpp>
#include <avm_values/value.hpp>
#include <fstream>
#include <iostream>

const int LEVEL = 3;

value table_to_tuple2(std::vector<value> tab, int prefix, int shift, int level) {
    if (level == 0) {
        std::vector<value> v;
        for (int i = 0; i < 8; i++) {
            int idx = prefix + (i << shift);
            // std::cerr << "idx " << idx << " prefix " << prefix << " shift " << shift << "\n";
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
    return table_to_tuple2(tab, 0, 0, 1);
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
        Code code;
        CodePointStub stub = code.addSegment();
        std::vector<value> labels;
        int idx = json_code.size();
        for (auto it = json_code.rbegin(); it != json_code.rend(); ++it) {
            Operation op = simple_operation_from_json(*it);
            stub = code.addOperation(stub.pc, op);
            idx--;
            if (has_labels[idx]) {
                labels.push_back(stub);
            }
        }
        auto table = make_table(labels);
        std::cerr << "Here " << intx::to_string(stub.hash, 16) << " " << labels.size() << " \n";
        std::cerr << "Table " << table << " hash " << intx::to_string(hash_value(table), 16) << "\n";

}
