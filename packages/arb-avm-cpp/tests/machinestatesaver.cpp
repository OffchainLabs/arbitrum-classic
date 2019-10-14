/*
 * Copyright 2019, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include "avm/machinestate/machinestatesaver.hpp"
#include <catch2/catch.hpp>

std::string path =
    "/Users/minhtruong/Dev/arbitrum/packages/arb-avm-cpp/build/tests/rocksDb";

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
    auto serialized_val = StateSaverUtils::serializeValue(results.val);
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

void getTupleValues(MachineStateSaver& saver,
                    std::vector<unsigned char>& hash_key,
                    std::vector<uint256_t> value_hashes) {
    auto results = saver.getTuple(hash_key);
    REQUIRE(results.tuple.tuple_size() == value_hashes.size());

    for (size_t i = 0; i < value_hashes.size(); i++) {
        REQUIRE(hash(results.tuple.get_element(i)) == value_hashes[i]);
    }
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

    SECTION("save empty tuple") {
        auto tuple = Tuple();
        auto hash_key = GetHashKey(tuple);
        auto tup_hash = tuple.calculateHash();

        saveValue(saver, tuple, 1, true);
        getValue(saver, hash_key, 1, tup_hash, TUPLE_TYPE, true);
    }
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

TEST_CASE("Save and get tuple values") {
    SECTION("save num tuple") {
        TuplePool pool;
        CheckpointStorage storage(path);
        auto saver = MachineStateSaver(&storage, &pool);

        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);

        saveTuple(saver, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(num)};
        auto hash_key = GetHashKey(tuple);

        getTupleValues(saver, hash_key, hashes);
    }
    SECTION("save codepoint tuple") {
        TuplePool pool;
        CheckpointStorage storage(path);
        auto saver = MachineStateSaver(&storage, &pool);

        auto code_point = CodePoint(3, Operation(), 0);
        auto tuple = Tuple(code_point, &pool);

        saveTuple(saver, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(code_point)};
        auto hash_key = GetHashKey(tuple);

        getTupleValues(saver, hash_key, hashes);
    }
    SECTION("save nested tuple") {
        TuplePool pool;
        CheckpointStorage storage(path);
        auto saver = MachineStateSaver(&storage, &pool);

        auto inner_tuple = Tuple();
        auto tuple = Tuple(inner_tuple, &pool);

        saveTuple(saver, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(inner_tuple)};
        auto hash_key = GetHashKey(tuple);

        getTupleValues(saver, hash_key, hashes);
    }
    SECTION("save multiple valued tuple") {
        TuplePool pool;
        CheckpointStorage storage(path);
        auto saver = MachineStateSaver(&storage, &pool);

        auto inner_tuple = Tuple();
        uint256_t num = 1;
        auto code_point = CodePoint(3, Operation(), 0);
        auto tuple = Tuple(inner_tuple, num, code_point, &pool);

        saveTuple(saver, tuple, 1, true);

        std::vector<uint256_t> hashes{hash(inner_tuple), hash(num),
                                      hash(code_point)};
        auto hash_key = GetHashKey(tuple);

        getTupleValues(saver, hash_key, hashes);
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
    SECTION("save tuple in tuple") {
        TuplePool pool;
        CheckpointStorage storage(path);
        auto saver = MachineStateSaver(&storage, &pool);

        uint256_t num = 1;
        auto inner_tuple = Tuple(num, &pool);
        auto tuple = Tuple(inner_tuple, &pool);
        saveTuple(saver, tuple, 1, true);

        auto inner_hash_key = GetHashKey(inner_tuple);
        auto inner_tup_hash = inner_tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);
        auto tup_hash = tuple.calculateHash();

        getTuple(saver, hash_key, 1, tup_hash, 1, true);
        getTuple(saver, inner_hash_key, 1, inner_tup_hash, 1, true);
    }
    SECTION("save 2 tuples in tuple") {
        TuplePool pool;
        CheckpointStorage storage(path);
        auto saver = MachineStateSaver(&storage, &pool);

        uint256_t num = 1;
        auto inner_tuple = Tuple(num, &pool);
        uint256_t num2 = 2;
        auto inner_tuple2 = Tuple(num2, &pool);
        auto tuple = Tuple(inner_tuple, inner_tuple2, &pool);
        saveTuple(saver, tuple, 1, true);

        auto inner_hash_key = GetHashKey(inner_tuple);
        auto inner_tup_hash = inner_tuple.calculateHash();
        auto inner_hash_key2 = GetHashKey(inner_tuple2);
        auto inner_tup_hash2 = inner_tuple2.calculateHash();
        auto hash_key = GetHashKey(tuple);
        auto tup_hash = tuple.calculateHash();

        getTuple(saver, hash_key, 1, tup_hash, 2, true);
        getTuple(saver, inner_hash_key, 1, inner_tup_hash, 1, true);
        getTuple(saver, inner_hash_key2, 1, inner_tup_hash2, 1, true);
    }
    SECTION("save saved tuple in tuple") {
        TuplePool pool;
        CheckpointStorage storage(path);
        auto saver = MachineStateSaver(&storage, &pool);

        uint256_t num = 1;
        auto inner_tuple = Tuple(num, &pool);
        auto tuple = Tuple(inner_tuple, &pool);
        auto inner_hash_key = GetHashKey(inner_tuple);
        auto inner_tup_hash = inner_tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);
        auto tup_hash = tuple.calculateHash();

        saveTuple(saver, inner_tuple, 1, true);
        getTuple(saver, inner_hash_key, 1, inner_tup_hash, 1, true);
        saveTuple(saver, tuple, 1, true);
        getTuple(saver, hash_key, 1, tup_hash, 1, true);
        getTuple(saver, inner_hash_key, 2, inner_tup_hash, 1, true);
    }
}

// MachineStateStorageData getStateStorageData(MachineStateSaver& saver){
//    TuplePool pool;
//
//    uint256_t static_val = 100;
//    auto register_val = Tuple(static_val, Tuple(), &pool);
//
//
//}

// TEST_CASE("Save Machinestate"){
//    TuplePool pool;
//    CheckpointStorage storage(path);
//    auto saver = MachineStateSaver(&storage, &pool);
//
//    auto code_point = CodePoint(3, Operation(), 0);
//    auto code_point2 = CodePoint(2, Operation(), 0);
//    uint256_t num_val = 5;
//
//    uint256_t static_val = 100;
//    auto register_val = Tuple(static_val, Tuple(), &pool);
//
//    auto data_stack = Tuple(code_point, static_val, Tuple(), &pool);
//    auto aux_stack = Tuple(num_val, data_stack, &pool);
//
//    auto inbox_msgs = Tuple(code_point2, static_val, data_stack, &pool);
//    auto pending_msgs = Tuple
//
//}
