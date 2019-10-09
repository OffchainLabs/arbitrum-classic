//
//  machinestatesaver.cpp
//  avm
//
//  Created by Minh Truong on 10/8/19.
//

#include "avm/machinestate/machinestatesaver.hpp"
#include <catch2/catch.hpp>

std::string path =
    "/Users/minhtruong/Dev/arbitrum/packages/arb-avm-cpp/build/tests/rocksDb";

// MachineStateSaver getStateSaver(){
//    TuplePool pool;
//    CheckpointStorage storage(path);
//    auto saver = MachineStateSaver(&storage, &pool);
//    return saver;
//}

void saveValue(MachineStateSaver& saver,
               const value& val,
               int expected_ref_count,
               bool expected_status) {
    auto results = saver.SaveValue(val);
    REQUIRE(results.status.ok() == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
}

void getValue(MachineStateSaver& saver,
              std::vector<unsigned char>& hash_key,
              int expected_ref_count,
              uint256_t& expected_hash,
              valueTypes expected_value_type,
              bool expected_status) {
    auto results = saver.getValue(hash_key);
    auto serialized_val = StateSaverUtils::SerializeValue(results.val);
    REQUIRE(results.status.ok() == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(serialized_val.type == expected_value_type);
    REQUIRE(hash(results.val) == expected_hash);
}

void saveTuple(MachineStateSaver& saver,
               const Tuple& tup,
               int expected_ref_count,
               bool expected_status) {
    auto results = saver.SaveTuple(tup);
    REQUIRE(results.status.ok() == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
}

void getTuple(MachineStateSaver& saver,
              std::vector<unsigned char>& hash_key,
              int expected_ref_count,
              uint256_t& expected_hash,
              int expected_tuple_size,
              bool expected_status) {
    auto results = saver.getTuple(hash_key);
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(results.tuple.calculateHash() == expected_hash);
    REQUIRE(results.tuple.tuple_size() == expected_tuple_size);
    REQUIRE(results.status.ok() == expected_status);
}

TEST_CASE("Save tuple") {
    TuplePool pool;
    CheckpointStorage storage(path);
    auto saver = MachineStateSaver(&storage, &pool);

    SECTION("save 1 num tuple") {
        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);
        saveTuple(saver, tuple, 1, true);
    }
    SECTION("save 2, 1 num tuples") {
        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);
        saveTuple(saver, tuple, 1, true);
        saveTuple(saver, tuple, 2, true);
    }
    SECTION("saved tuple in tuple") {
        uint256_t num = 1;
        auto inner_tuple = Tuple(num, &pool);
        auto tuple = Tuple(inner_tuple, &pool);
        saveTuple(saver, tuple, 1, true);
        saveTuple(saver, tuple, 2, true);
    }
}

TEST_CASE("Save value") {
    TuplePool pool;
    CheckpointStorage storage(path);
    auto saver = MachineStateSaver(&storage, &pool);

    SECTION("save 1 num tuple") {
        TuplePool pool;
        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);
        saveValue(saver, tuple, 1, true);
    }
    SECTION("save num") {
        uint256_t num = 1;
        saveValue(saver, num, 1, true);
    }
    SECTION("save codepoint") {
        auto code_point = CodePoint(1, Operation(), 0);
        saveValue(saver, code_point, 1, true);
    }
}

TEST_CASE("Save and get value") {
    TuplePool pool;
    CheckpointStorage storage(path);
    auto saver = MachineStateSaver(&storage, &pool);

    SECTION("save tuple") {
        TuplePool pool;
        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);
        auto hash_key = GetHashKey(tuple);
        auto tup_hash = tuple.calculateHash();

        saveValue(saver, tuple, 1, true);
        getValue(saver, hash_key, 1, tup_hash, TUPLE_TYPE, true);
    }
    SECTION("save num") {
        uint256_t num = 1;
        auto hash_key = GetHashKey(num);
        auto num_hash = hash(num);

        saveValue(saver, num, 1, true);
        getValue(saver, hash_key, 1, num_hash, NUM_TYPE, true);
    }
    SECTION("save codepoint") {
        CodePoint code_point(1, Operation(), 0);
        auto hash_key = GetHashKey(code_point);
        auto cp_hash = hash(code_point);
        saveValue(saver, code_point, 1, true);
        getValue(saver, hash_key, 1, cp_hash, CODEPT_TYPE, true);
    }
}

TEST_CASE("Save And Get Tuple") {
    SECTION("save 1 num tuple") {
        TuplePool pool;
        CheckpointStorage storage(path);
        auto saver = MachineStateSaver(&storage, &pool);

        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);
        auto tup_hash = tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);

        saveTuple(saver, tuple, 1, true);
        getTuple(saver, hash_key, 1, tup_hash, 1, true);
    }
    SECTION("save 1 num tuple twice") {
        TuplePool pool;
        CheckpointStorage storage(path);
        auto saver = MachineStateSaver(&storage, &pool);

        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);
        auto tup_hash = tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);

        saveTuple(saver, tuple, 1, true);
        saveTuple(saver, tuple, 2, true);
        getTuple(saver, hash_key, 2, tup_hash, 1, true);
    }
    SECTION("save 2 num tuple") {
        TuplePool pool;
        CheckpointStorage storage(path);
        auto saver = MachineStateSaver(&storage, &pool);

        uint256_t num = 1;
        uint256_t num2 = 2;
        auto tuple = Tuple(num, num2, &pool);
        auto tup_hash = tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);

        saveTuple(saver, tuple, 1, true);
        getTuple(saver, hash_key, 1, tup_hash, 2, true);
    }
}
