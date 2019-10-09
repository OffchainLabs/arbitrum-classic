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
              uint256_t& expected_hash) {
    auto results = saver.getTuple(hash_key);
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(results.tuple.calculateHash() == expected_hash);
}

TEST_CASE("Save tuple") {
    //    TuplePool pool;
    //    CheckpointStorage storage(path);
    //    auto saver = MachineStateSaver(&storage, &pool);

    //    SECTION("save 1 num tuple"){
    //        TuplePool pool;
    //        uint256_t num = 1;
    //        auto tuple = Tuple(num, &pool);
    //        saveTuple(saver, tuple, 1, true);
    //    }
    //    SECTION("save 2, 1 num tuples"){
    //        TuplePool pool;
    //        uint256_t num = 1;
    //        auto tuple = Tuple(num, &pool);
    //        saveTuple(saver, tuple, 1, true);
    //        saveTuple(saver, tuple, 2, true);
    //    }
}

// TEST_CASE("Save And Get Tuple"){
//
//    SECTION("save 1 num tuple"){
//        TuplePool pool;
//        CheckpointStorage storage(path);
//        auto saver = MachineStateSaver(&storage, &pool);
//
//        uint256_t num = 1;
//        auto tuple = Tuple(num, &pool);
//        auto tup_hash = tuple.calculateHash();
//        auto hash_key = GetHashKey(tuple);
//
////        saveTuple(saver, tuple, 1, true);
//        getTuple(saver, hash_key, 1, tup_hash);
//    }
//}
