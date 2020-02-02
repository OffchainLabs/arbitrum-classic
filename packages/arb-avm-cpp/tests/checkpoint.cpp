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

#include "config.hpp"

#include <avm/machinestate/machinestate.hpp>
#include <avm_values/vmValueParser.hpp>
#include <data_storage/checkpoint/checkpointstorage.hpp>
#include <data_storage/checkpoint/machinestatedeleter.hpp>
#include <data_storage/checkpoint/machinestatefetcher.hpp>
#include <data_storage/checkpoint/machinestatesaver.hpp>

#include <catch2/catch.hpp>

#include <boost/filesystem.hpp>

std::string path =
    boost::filesystem::current_path().generic_string() + "/machineDb";

void saveValue(MachineStateSaver& saver,
               const value& val,
               int expected_ref_count,
               bool expected_status) {
    auto results = saver.saveValue(val);
    saver.commitTransaction();
    REQUIRE(results.status.ok() == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
}

void getValue(MachineStateFetcher& fetcher,
              std::vector<unsigned char>& hash_key,
              int expected_ref_count,
              uint256_t& expected_hash,
              ValueTypes expected_value_type,
              bool expected_status) {
    auto results = fetcher.getValue(hash_key);
    auto serialized_val = checkpoint::utils::serializeValue(results.data);
    auto type = (ValueTypes)serialized_val[0];

    REQUIRE(results.status.ok() == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(type == expected_value_type);
    REQUIRE(hash(results.data) == expected_hash);
}

void saveTuple(MachineStateSaver& saver,
               const Tuple& tup,
               int expected_ref_count,
               bool expected_status) {
    auto results = saver.saveTuple(tup);
    saver.commitTransaction();
    REQUIRE(results.status.ok() == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
}

void getTuple(MachineStateFetcher& fetcher,
              std::vector<unsigned char>& hash_key,
              int expected_ref_count,
              uint256_t& expected_hash,
              int expected_tuple_size,
              bool expected_status) {
    auto results = fetcher.getTuple(hash_key);
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(results.data.calculateHash() == expected_hash);
    REQUIRE(results.data.tuple_size() == expected_tuple_size);
    REQUIRE(results.status.ok() == expected_status);
}

void getTupleValues(MachineStateFetcher& fetcher,
                    const std::vector<unsigned char>& hash_key,
                    const std::vector<uint256_t>& value_hashes) {
    auto results = fetcher.getTuple(hash_key);
    REQUIRE(results.data.tuple_size() == value_hashes.size());

    for (size_t i = 0; i < value_hashes.size(); i++) {
        REQUIRE(hash(results.data.get_element(i)) == value_hashes[i]);
    }
}

TEST_CASE("Save value") {
    CheckpointStorage storage(path, test_contract_path);
    auto saver = MachineStateSaver(storage.makeTransaction());

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
    boost::filesystem::remove_all(path);
}

TEST_CASE("Save tuple") {
    CheckpointStorage storage(path, test_contract_path);
    auto saver = MachineStateSaver(storage.makeTransaction());

    TuplePool pool;

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
    boost::filesystem::remove_all(path);
}

TEST_CASE("Save and get value") {
    SECTION("save empty tuple") {
        CheckpointStorage storage(path, test_contract_path);
        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        auto tuple = Tuple();
        auto hash_key = GetHashKey(tuple);
        auto tup_hash = tuple.calculateHash();

        saveValue(saver, tuple, 1, true);
        getValue(fetcher, hash_key, 1, tup_hash, TUPLE, true);
    }
    boost::filesystem::remove_all(path);
    SECTION("save tuple") {
        CheckpointStorage storage(path, test_contract_path);
        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        TuplePool pool;
        auto tuple = Tuple(num, &pool);
        auto hash_key = GetHashKey(tuple);
        auto tup_hash = tuple.calculateHash();

        saveValue(saver, tuple, 1, true);
        getValue(fetcher, hash_key, 1, tup_hash, TUPLE, true);
    }
    boost::filesystem::remove_all(path);
    SECTION("save num") {
        CheckpointStorage storage(path, test_contract_path);
        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        auto hash_key = GetHashKey(num);
        auto num_hash = hash(num);

        saveValue(saver, num, 1, true);
        getValue(fetcher, hash_key, 1, num_hash, NUM, true);
    }
    boost::filesystem::remove_all(path);
    SECTION("save codepoint") {
        CheckpointStorage storage(path, test_contract_path);
        CodePoint code_point = storage.getInitialVmValues().code[0];

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        auto hash_key = GetHashKey(code_point);
        auto cp_hash = hash(code_point);
        saveValue(saver, code_point, 1, true);
        getValue(fetcher, hash_key, 1, cp_hash, CODEPT, true);
    }
    SECTION("save err codepoint") {
        CheckpointStorage storage(path, test_contract_path);
        auto code_point = getErrCodePoint();

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        auto hash_key = GetHashKey(code_point);
        auto cp_hash = hash(code_point);
        saveValue(saver, code_point, 1, true);
        getValue(fetcher, hash_key, 1, cp_hash, CODEPT, true);
    }
    boost::filesystem::remove_all(path);
}

TEST_CASE("Save and get tuple values") {
    SECTION("save num tuple") {
        CheckpointStorage storage(path, test_contract_path);
        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        TuplePool pool;
        auto tuple = Tuple(num, &pool);

        saveTuple(saver, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(num)};
        auto hash_key = GetHashKey(tuple);

        getTupleValues(fetcher, hash_key, hashes);
    }
    boost::filesystem::remove_all(path);
    SECTION("save codepoint tuple") {
        CheckpointStorage storage(path, test_contract_path);
        CodePoint code_point = storage.getInitialVmValues().code[1];

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        TuplePool pool;
        auto tuple = Tuple(code_point, &pool);

        saveTuple(saver, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(code_point)};
        auto hash_key = GetHashKey(tuple);

        getTupleValues(fetcher, hash_key, hashes);
    }
    boost::filesystem::remove_all(path);
    SECTION("save codepoint tuple") {
        CheckpointStorage storage(path, test_contract_path);
        CodePoint code_point = storage.getInitialVmValues().code[0];

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        TuplePool pool;
        auto tuple = Tuple(code_point, &pool);

        saveValue(saver, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(code_point)};
        auto hash_key = GetHashKey(tuple);

        getTupleValues(fetcher, hash_key, hashes);
    }
    boost::filesystem::remove_all(path);
    SECTION("save nested tuple") {
        CheckpointStorage storage(path, test_contract_path);

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        auto inner_tuple = Tuple();
        TuplePool pool;
        auto tuple = Tuple(inner_tuple, &pool);

        saveTuple(saver, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(inner_tuple)};
        auto hash_key = GetHashKey(tuple);

        getTupleValues(fetcher, hash_key, hashes);
    }
    boost::filesystem::remove_all(path);
    SECTION("save multiple valued tuple") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);
        CodePoint code_point = storage.getInitialVmValues().code[0];

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        auto inner_tuple = Tuple();
        uint256_t num = 1;
        auto tuple = Tuple(inner_tuple, num, code_point, &pool);

        saveTuple(saver, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(inner_tuple), hash(num),
                                      hash(code_point)};
        auto hash_key = GetHashKey(tuple);

        getTupleValues(fetcher, hash_key, hashes);
    }
    boost::filesystem::remove_all(path);
    SECTION("save multiple valued tuple, saveValue()") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[2];

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        auto inner_tuple = Tuple();
        uint256_t num = 1;
        auto tuple = Tuple(inner_tuple, num, code_point, &pool);

        saveValue(saver, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(inner_tuple), hash(num),
                                      hash(code_point)};
        auto hash_key = GetHashKey(tuple);

        getTupleValues(fetcher, hash_key, hashes);
    }
    boost::filesystem::remove_all(path);
}

TEST_CASE("Save And Get Tuple") {
    SECTION("save 1 num tuple") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);
        auto tup_hash = tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);

        saveTuple(saver, tuple, 1, true);
        getTuple(fetcher, hash_key, 1, tup_hash, 1, true);
    }
    boost::filesystem::remove_all(path);
    SECTION("save codepoint in tuple") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);

        auto code_point = storage.getInitialVmValues().code[0];

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        auto tuple = Tuple(code_point, &pool);
        auto tup_hash = tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);

        saveTuple(saver, tuple, 1, true);
        getTuple(fetcher, hash_key, 1, tup_hash, 1, true);
    }
    boost::filesystem::remove_all(path);
    SECTION("save 1 num tuple twice") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto saver2 = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);
        auto tup_hash = tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);

        saveTuple(saver, tuple, 1, true);
        saveTuple(saver2, tuple, 2, true);
        getTuple(fetcher, hash_key, 2, tup_hash, 1, true);
    }
    boost::filesystem::remove_all(path);
    SECTION("save 2 num tuple") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);
        std::vector<CodePoint> code;
        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        uint256_t num2 = 2;
        auto tuple = Tuple(num, num2, &pool);
        auto tup_hash = tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);

        saveTuple(saver, tuple, 1, true);
        getTuple(fetcher, hash_key, 1, tup_hash, 2, true);
    }
    boost::filesystem::remove_all(path);
    SECTION("save tuple in tuple") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        auto inner_tuple = Tuple(num, &pool);
        auto tuple = Tuple(inner_tuple, &pool);
        saveTuple(saver, tuple, 1, true);

        auto inner_hash_key = GetHashKey(inner_tuple);
        auto inner_tup_hash = inner_tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);
        auto tup_hash = tuple.calculateHash();

        getTuple(fetcher, hash_key, 1, tup_hash, 1, true);
        getTuple(fetcher, inner_hash_key, 1, inner_tup_hash, 1, true);
    }
    boost::filesystem::remove_all(path);
    SECTION("save 2 tuples in tuple") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

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

        getTuple(fetcher, hash_key, 1, tup_hash, 2, true);
        getTuple(fetcher, inner_hash_key, 1, inner_tup_hash, 1, true);
        getTuple(fetcher, inner_hash_key2, 1, inner_tup_hash2, 1, true);
    }
    boost::filesystem::remove_all(path);
    SECTION("save saved tuple in tuple") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto saver2 = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        auto inner_tuple = Tuple(num, &pool);
        auto tuple = Tuple(inner_tuple, &pool);
        auto inner_hash_key = GetHashKey(inner_tuple);
        auto inner_tup_hash = inner_tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);
        auto tup_hash = tuple.calculateHash();

        saveTuple(saver, inner_tuple, 1, true);
        getTuple(fetcher, inner_hash_key, 1, inner_tup_hash, 1, true);
        saveTuple(saver2, tuple, 1, true);
        getTuple(fetcher, hash_key, 1, tup_hash, 1, true);
        getTuple(fetcher, inner_hash_key, 2, inner_tup_hash, 1, true);
    }
    boost::filesystem::remove_all(path);
}

void saveState(MachineStateSaver& saver,
               const MachineStateKeys& storage_data,
               const std::vector<unsigned char>& checkpoint_name) {
    auto results = saver.saveMachineState(storage_data, checkpoint_name);
    auto status = saver.commitTransaction();

    REQUIRE(results.reference_count == 1);
    REQUIRE(results.status.ok());
}

void getSavedState(MachineStateFetcher& fetcher,
                   const std::vector<unsigned char>& checkpoint_name,
                   const MachineStateKeys& expected_data,
                   int expected_ref_count,
                   const std::vector<std::vector<unsigned char>>& keys) {
    auto results = fetcher.getMachineState(checkpoint_name);

    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == expected_ref_count);

    auto data = results.data;

    REQUIRE(data.status_char == expected_data.status_char);
    REQUIRE(data.pc_key == expected_data.pc_key);
    REQUIRE(data.datastack_key == expected_data.datastack_key);
    REQUIRE(data.auxstack_key == expected_data.auxstack_key);
    REQUIRE(data.register_val_key == expected_data.register_val_key);

    for (auto& key : keys) {
        auto res = fetcher.getValue(key);
        REQUIRE(res.status.ok());
    }
}

void deleteCheckpoint(
    CheckpointStorage& storage,
    const MachineStateFetcher& fetcher,
    const std::vector<unsigned char>& checkpoint_name,
    const std::vector<std::vector<unsigned char>>& deleted_values) {
    auto res = deleteCheckpoint(storage, checkpoint_name);
    auto results = fetcher.getMachineState(checkpoint_name);
    REQUIRE(results.status.ok() == false);

    for (auto& hash_key : deleted_values) {
        auto res = fetcher.getValue(hash_key);
        REQUIRE(res.status.ok() == false);
    }
}

void deleteCheckpointSavedTwice(
    CheckpointStorage& storage,
    const MachineStateFetcher& fetcher,
    const std::vector<unsigned char>& checkpoint_name,
    const std::vector<std::vector<unsigned char>>& deleted_values) {
    auto res = deleteCheckpoint(storage, checkpoint_name);
    auto res2 = deleteCheckpoint(storage, checkpoint_name);
    auto results = fetcher.getMachineState(checkpoint_name);

    REQUIRE(results.status.ok() == false);

    for (auto& hash_key : deleted_values) {
        auto res = fetcher.getValue(hash_key);
        REQUIRE(res.status.ok() == false);
    }
}

void deleteCheckpointSavedTwiceReordered(
    CheckpointStorage& storage,
    const MachineStateFetcher& fetcher,
    const std::vector<unsigned char>& checkpoint_name,
    const std::vector<std::vector<unsigned char>>& deleted_values) {
    auto resultsx = fetcher.getMachineState(checkpoint_name);
    for (auto& hash_key : deleted_values) {
        auto res = fetcher.getValue(hash_key);
        REQUIRE(res.status.ok());
    }
    auto res = deleteCheckpoint(storage, checkpoint_name);
    auto results = fetcher.getMachineState(checkpoint_name);
    REQUIRE(results.status.ok() == true);

    for (auto& hash_key : deleted_values) {
        auto res = fetcher.getValue(hash_key);
        REQUIRE(res.status.ok());
    }
    auto res2 = deleteCheckpoint(storage, checkpoint_name);
    auto results2 = fetcher.getMachineState(checkpoint_name);
    REQUIRE(results2.status.ok() == false);

    for (auto& hash_key : deleted_values) {
        auto res = fetcher.getValue(hash_key);
        REQUIRE(res.status.ok() == false);
    }
}

auto makeStorageData(MachineStateSaver& stateSaver,
                     const value& registerVal,
                     const Datastack& stack,
                     const Datastack& auxstack,
                     Status state,
                     const CodePoint& pc,
                     const CodePoint& err_pc,
                     const BlockReason& blockReason) -> MachineStateKeys {
    TuplePool pool;

    auto datastack_results = stack.checkpointState(stateSaver, &pool);
    auto auxstack_results = auxstack.checkpointState(stateSaver, &pool);

    auto register_val_results = stateSaver.saveValue(registerVal);
    auto pc_results = stateSaver.saveValue(pc);
    auto err_pc_results = stateSaver.saveValue(err_pc);

    auto status_str = (unsigned char)state;

    return MachineStateKeys{
        register_val_results.storage_key, datastack_results.storage_key,
        auxstack_results.storage_key,     pc_results.storage_key,
        err_pc_results.storage_key,       status_str};
}

auto getStateValues(MachineStateSaver& saver) -> MachineStateKeys {
    TuplePool pool;
    uint256_t register_val = 100;
    auto static_val = Tuple(register_val, Tuple(), &pool);

    auto code_point = CodePoint(1, Operation(), 0);
    auto tup1 = Tuple(register_val, &pool);
    auto tup2 = Tuple(code_point, tup1, &pool);

    Datastack data_stack;
    data_stack.push(register_val);
    Datastack aux_stack;
    aux_stack.push(register_val);
    aux_stack.push(code_point);

    CodePoint pc_codepoint(0, Operation(), 0);
    CodePoint err_pc_codepoint(0, Operation(), 0);
    Status state = Status::Extensive;

    auto inbox_blocked = InboxBlocked(0);

    auto saved_data =
        makeStorageData(saver, register_val, data_stack, aux_stack, state,
                        pc_codepoint, err_pc_codepoint, inbox_blocked);

    return saved_data;
}

auto getDefaultValues(MachineStateSaver& saver) -> MachineStateKeys {
    TuplePool pool;
    auto register_val = Tuple();
    auto data_stack = Tuple();
    auto aux_stack = Tuple();

    Status state = Status::Extensive;
    CodePoint code_point(0, Operation(), 0);

    auto data = makeStorageData(saver, Tuple(), Datastack(), Datastack(), state,
                                code_point, code_point, NotBlocked());

    return data;
}

auto getHashKeys(const MachineStateKeys& data)
    -> std::vector<std::vector<unsigned char>> {
    std::vector<std::vector<unsigned char>> hash_keys;

    hash_keys.push_back(data.auxstack_key);
    hash_keys.push_back(data.datastack_key);
    hash_keys.push_back(data.pc_key);
    hash_keys.push_back(data.err_pc_key);
    hash_keys.push_back(data.register_val_key);

    return hash_keys;
}

TEST_CASE("Save Machinestatedata") {
    SECTION("default") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[0];

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto data_values = getDefaultValues(saver);
        std::vector<unsigned char> checkpoint_key = {'k', 'e', 'y'};

        saveState(saver, data_values, checkpoint_key);
    }
    boost::filesystem::remove_all(path);
    SECTION("with values") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[0];
        CodePoint code_point2 = storage.getInitialVmValues().code[1];

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto state_data = getStateValues(saver);

        std::vector<unsigned char> checkpoint_key = {'k', 'e', 'y'};

        saveState(saver, state_data, checkpoint_key);
    }
    boost::filesystem::remove_all(path);
}

TEST_CASE("Get Machinestate data") {
    SECTION("default") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[0];

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        auto data_values = getDefaultValues(saver);
        auto keys = getHashKeys(data_values);

        std::vector<unsigned char> checkpoint_key = {'k', 'e', 'y'};

        saver.saveMachineState(data_values, checkpoint_key);
        saver.commitTransaction();
        getSavedState(fetcher, checkpoint_key, data_values, 1, keys);
    }
    boost::filesystem::remove_all(path);
    SECTION("with values") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[0];
        CodePoint code_point2 = storage.getInitialVmValues().code[1];

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        auto state_data = getStateValues(saver);
        auto keys = getHashKeys(state_data);

        std::vector<unsigned char> checkpoint_key = {'k', 'e', 'y'};

        saveState(saver, state_data, checkpoint_key);
        getSavedState(fetcher, checkpoint_key, state_data, 1, keys);
    }
    boost::filesystem::remove_all(path);
}

TEST_CASE("Delete checkpoint") {
    SECTION("default") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[0];

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);
        auto data_values = getDefaultValues(saver);

        std::vector<unsigned char> checkpoint_key = {'k', 'e', 'y'};

        saver.saveMachineState(data_values, checkpoint_key);
        saver.commitTransaction();
        auto hash_keys = getHashKeys(data_values);

        deleteCheckpoint(storage, fetcher, checkpoint_key, hash_keys);
    }
    boost::filesystem::remove_all(path);
    SECTION("with actual state values") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[0];
        CodePoint code_point2 = storage.getInitialVmValues().code[2];

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);
        auto data_values = getStateValues(saver);

        std::vector<unsigned char> checkpoint_key = {'k', 'e', 'y'};

        saver.saveMachineState(data_values, checkpoint_key);
        saver.commitTransaction();
        auto hash_keys = getHashKeys(data_values);

        deleteCheckpoint(storage, fetcher, checkpoint_key, hash_keys);
    }
    boost::filesystem::remove_all(path);
    SECTION("delete checkpoint saved twice") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[0];
        CodePoint code_point2 = storage.getInitialVmValues().code[1];

        auto fetcher = MachineStateFetcher(storage);
        auto saver = MachineStateSaver(storage.makeTransaction());
        auto data_values = getStateValues(saver);

        std::vector<unsigned char> checkpoint_key = {'k', 'e', 'y'};

        saver.saveMachineState(data_values, checkpoint_key);
        saver.saveMachineState(data_values, checkpoint_key);
        saver.commitTransaction();
        auto hash_keys = getHashKeys(data_values);

        deleteCheckpointSavedTwice(storage, fetcher, checkpoint_key, hash_keys);
    }
    boost::filesystem::remove_all(path);
    SECTION("delete checkpoint saved twice, reordered") {
        TuplePool pool;
        CheckpointStorage storage(path, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[0];
        CodePoint code_point2 = storage.getInitialVmValues().code[1];

        auto fetcher = MachineStateFetcher(storage);
        auto saver = MachineStateSaver(storage.makeTransaction());
        auto saver2 = MachineStateSaver(storage.makeTransaction());
        auto data_values = getStateValues(saver);

        std::vector<unsigned char> checkpoint_key = {'k', 'e', 'y'};

        saver.saveMachineState(data_values, checkpoint_key);
        saver.commitTransaction();
        saver2.saveMachineState(data_values, checkpoint_key);
        saver2.commitTransaction();
        auto hash_keys = getHashKeys(data_values);

        deleteCheckpointSavedTwiceReordered(storage, fetcher, checkpoint_key,
                                            hash_keys);
    }
    boost::filesystem::remove_all(path);
}
