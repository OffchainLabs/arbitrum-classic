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
#include "avm/machinestate/machinestate.hpp"

std::string path =
    "/Users/minhtruong/Dev/arbitrum/packages/arb-avm-cpp/build/tests/rocksDb";

void saveValue(MachineStateSaver& saver,
               const value& val,
               int expected_ref_count,
               bool expected_status) {
    auto results = saver.saveValue(val);
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
    auto type = (valueTypes)serialized_val[0];

    REQUIRE(results.status.ok() == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(type == expected_value_type);
    REQUIRE(hash(results.val) == expected_hash);
}

void saveTuple(MachineStateSaver& saver,
               const Tuple& tup,
               int expected_ref_count,
               bool expected_status) {
    auto results = saver.saveTuple(tup);
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

void saveState(MachineStateSaver& saver,
               MachineStateStorageData storage_data,
               std::string checkpoint_name) {
    auto results = saver.saveMachineState(storage_data, checkpoint_name);

    REQUIRE(results.reference_count == 1);
    REQUIRE(results.status.ok());
}

void saveAndGetState(MachineStateSaver& saver,
                     std::string checkpoint_name,
                     MachineStateFetchedData expected_data,
                     int expected_ref_count) {
    auto results = saver.getMachineStateData(checkpoint_name);

    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == expected_ref_count);

    auto data = results.state_data;

    REQUIRE(data.status_char == expected_data.status_char);
    REQUIRE(data.blockreason_str == expected_data.blockreason_str);
    REQUIRE(data.balancetracker_str == expected_data.balancetracker_str);
    REQUIRE(hash(data.static_val) == hash(expected_data.static_val));
    REQUIRE(hash(data.inbox_count) == hash(expected_data.inbox_count));
    REQUIRE(hash(data.pending_count) == hash(expected_data.pending_count));
    REQUIRE(hash(data.pc_codepoint) == hash(expected_data.pc_codepoint));
    REQUIRE(hash(data.pending_count) == hash(expected_data.pending_count));
    REQUIRE(data.inbox_tuple.calculateHash() ==
            expected_data.inbox_tuple.calculateHash());
    REQUIRE(data.pending_inbox_tuple.calculateHash() ==
            expected_data.pending_inbox_tuple.calculateHash());
    REQUIRE(data.datastack_tuple.calculateHash() ==
            expected_data.datastack_tuple.calculateHash());
    REQUIRE(data.auxstack_tuple.calculateHash() ==
            expected_data.auxstack_tuple.calculateHash());
    REQUIRE(hash(data.register_val) == hash(expected_data.register_val));
}

void deleteCheckpoint(MachineStateSaver& saver,
                      std::string checkpoint_name,
                      std::vector<std::vector<unsigned char>> deleted_values) {
    saver.deleteCheckpoint(checkpoint_name);
    auto results = saver.getMachineStateData(checkpoint_name);
    REQUIRE(results.status.ok() == false);

    for (auto& hash_key : deleted_values) {
        auto res = saver.getValue(hash_key);
        REQUIRE(res.status.ok() == false);
    }
}

void deleteCheckpointSavedTwice(
    MachineStateSaver& saver,
    std::string checkpoint_name,
    std::vector<std::vector<unsigned char>> deleted_values) {
    saver.deleteCheckpoint(checkpoint_name);
    auto results = saver.getMachineStateData(checkpoint_name);
    REQUIRE(results.status.ok() == true);

    for (auto& hash_key : deleted_values) {
        auto res = saver.getValue(hash_key);
        REQUIRE(res.status.ok());
    }

    saver.deleteCheckpoint(checkpoint_name);
    auto results2 = saver.getMachineStateData(checkpoint_name);
    REQUIRE(results2.status.ok() == false);

    for (auto& hash_key : deleted_values) {
        auto res = saver.getValue(hash_key);
        REQUIRE(res.status.ok() == false);
    }
}

MachineStateStorageData makeStorageData(MachineStateSaver& stateSaver,
                                        value staticVal,
                                        value registerVal,
                                        Datastack stack,
                                        Datastack auxstack,
                                        Status state,
                                        CodePoint pc,
                                        MessageStack inbox,
                                        MessageStack pendingInbox,
                                        BalanceTracker balance,
                                        BlockReason blockReason) {
    TuplePool pool;

    auto datastack_results = stack.checkpointState(stateSaver, &pool);
    auto auxstack_results = auxstack.checkpointState(stateSaver, &pool);
    auto inbox_results = inbox.checkpointState(stateSaver);
    auto pending_results = pendingInbox.checkpointState(stateSaver);

    auto static_val_results = stateSaver.saveValue(staticVal);
    auto register_val_results = stateSaver.saveValue(registerVal);
    auto pc_results = stateSaver.saveValue(pc);

    auto status_str = (unsigned char)state;
    auto blockreason_str = serializeForCheckpoint(blockReason);
    auto balancetracker_str = balance.serializeBalanceValues();

    return MachineStateStorageData{
        static_val_results,
        register_val_results,
        datastack_results,
        auxstack_results,
        inbox_results.msgs_tuple_results,
        inbox_results.msg_count_results,
        pending_results.msgs_tuple_results,
        pending_results.msg_count_results,
        pc_results,
        status_str,
        blockreason_str,
        balancetracker_str,
    };
}

MessageStack getMsgStack1() {
    TuplePool pool;

    auto inbox_stack = MessageStack(&pool);
    uint256_t val_data = 111;
    uint256_t destination = 2;
    uint256_t currency = 3;
    auto msg_token_type = std::array<unsigned char, 21>();
    msg_token_type[0] = 'a';
    auto msg = Message{val_data, destination, currency, msg_token_type};
    inbox_stack.addMessage(msg);

    return inbox_stack;
}

MessageStack getMsgStack2() {
    TuplePool pool;

    uint256_t val_data = 111;
    uint256_t destination = 2;
    uint256_t currency = 3;
    auto pending_stack = MessageStack(&pool);
    auto pending_token_type = std::array<unsigned char, 21>();
    pending_token_type[0] = 'b';
    auto pending_msg =
        Message{val_data, destination, currency, pending_token_type};
    pending_stack.addMessage(pending_msg);

    return pending_stack;
}

std::tuple<MachineStateStorageData, MachineStateFetchedData> getStateValues(
    MachineStateSaver& saver) {
    TuplePool pool;
    uint256_t register_val = 100;
    auto static_val = Tuple(register_val, Tuple(), &pool);

    auto code_point = CodePoint(3, Operation(), 0);
    auto tup1 = Tuple(register_val, &pool);
    auto tup2 = Tuple(code_point, tup1, &pool);

    Datastack data_stack;
    data_stack.push(register_val);
    Datastack aux_stack;
    aux_stack.push(register_val);
    aux_stack.push(code_point);

    auto inbox_stack = getMsgStack1();
    auto pending_stack = getMsgStack2();

    CodePoint pc_codepoint(1, Operation(), 0);
    Status state = Status::Extensive;

    std::array<unsigned char, 21> block_token_type = {10};
    auto send_blocked = SendBlocked(999, block_token_type);

    std::array<unsigned char, 21> token_type = {1, 99};
    uint256_t amount = 11;
    auto tracker = BalanceTracker();
    tracker.add(token_type, amount);

    auto saved_data = makeStorageData(
        saver, static_val, register_val, data_stack, aux_stack, state,
        pc_codepoint, inbox_stack, pending_stack, tracker, send_blocked);
    auto expected_data =
        MachineStateFetchedData{static_val,
                                register_val,
                                tup1,
                                tup2,
                                inbox_stack.messages,
                                (uint256_t)inbox_stack.messageCount,
                                pending_stack.messages,
                                (uint256_t)pending_stack.messageCount,
                                code_point,
                                (unsigned char)state,
                                serializeForCheckpoint(send_blocked),
                                tracker.serializeBalanceValues()};

    return std::make_tuple(saved_data, expected_data);
}

std::tuple<MachineStateStorageData, MachineStateFetchedData> getDefaultValues(
    MachineStateSaver& saver) {
    TuplePool pool;
    uint256_t static_val = 0;
    auto register_val = Tuple();
    auto data_stack = Tuple();
    auto aux_stack = Tuple();
    auto inbox_mssage = MessageStack(&pool);
    auto pending_mssage = MessageStack(&pool);
    auto tracker = BalanceTracker();
    auto block_reason = NotBlocked();
    Status state = Status::Extensive;
    CodePoint code_point(0, Operation(), 0);

    auto data =
        makeStorageData(saver, static_val, Tuple(), Datastack(), Datastack(),
                        state, code_point, MessageStack(&pool),
                        MessageStack(&pool), BalanceTracker(), NotBlocked());

    auto expected =
        MachineStateFetchedData{static_val,
                                register_val,
                                data_stack,
                                aux_stack,
                                inbox_mssage.messages,
                                (uint256_t)inbox_mssage.messageCount,
                                pending_mssage.messages,
                                (uint256_t)pending_mssage.messageCount,
                                code_point,
                                (unsigned char)state,
                                serializeForCheckpoint(block_reason),
                                tracker.serializeBalanceValues()};

    return std::make_tuple(data, expected);
}

TEST_CASE("Save Machinestatedata") {
    SECTION("default") {
        TuplePool pool;
        CheckpointStorage storage(path);
        auto saver = MachineStateSaver(&storage, &pool);

        auto data_values = getDefaultValues(saver);
        auto data = std::get<0>(data_values);

        saveState(saver, data, "checkpoint");
    }
    SECTION("with values") {
        TuplePool pool;
        CheckpointStorage storage(path);
        auto saver = MachineStateSaver(&storage, &pool);

        auto state_data = getStateValues(saver);
        auto data = std::get<0>(state_data);

        saveState(saver, data, "checkpoint");
    }
}

TEST_CASE("Get Machinestate data") {
    SECTION("default") {
        TuplePool pool;
        CheckpointStorage storage(path);
        auto saver = MachineStateSaver(&storage, &pool);

        auto data_values = getDefaultValues(saver);
        auto data = std::get<0>(data_values);
        auto expected_data = std::get<1>(data_values);

        saver.saveMachineState(data, "checkpoint");
        saveAndGetState(saver, "checkpoint", expected_data, 1);
    }
    SECTION("with values") {
        TuplePool pool;
        CheckpointStorage storage(path);
        auto saver = MachineStateSaver(&storage, &pool);

        auto state_data = getStateValues(saver);
        auto stored_data = std::get<0>(state_data);
        auto expected_data = std::get<1>(state_data);

        saveState(saver, stored_data, "checkpoint");
        saveAndGetState(saver, "checkpoint", expected_data, 1);
    }
}

TEST_CASE("Delete checkpoint") {
    SECTION("default") {
        TuplePool pool;
        CheckpointStorage storage(path);
        auto saver = MachineStateSaver(&storage, &pool);

        auto data_values = getDefaultValues(saver);
        auto data = std::get<0>(data_values);

        saver.saveMachineState(data, "checkpoint");
        std::vector<std::vector<unsigned char>> hash_keys;

        hash_keys.push_back(data.auxstack_results.storage_key);
        hash_keys.push_back(data.datastack_results.storage_key);
        hash_keys.push_back(data.inbox_count_results.storage_key);
        hash_keys.push_back(data.inbox_messages_results.storage_key);
        hash_keys.push_back(data.pc_results.storage_key);
        hash_keys.push_back(data.pending_count_results.storage_key);
        hash_keys.push_back(data.pending_messages_results.storage_key);
        hash_keys.push_back(data.register_val_results.storage_key);
        hash_keys.push_back(data.static_val_results.storage_key);

        deleteCheckpoint(saver, "checkpoint", hash_keys);
    }
    SECTION("with actual state values") {
        TuplePool pool;
        CheckpointStorage storage(path);
        auto saver = MachineStateSaver(&storage, &pool);

        auto data_values = getStateValues(saver);
        auto data = std::get<0>(data_values);

        saver.saveMachineState(data, "checkpoint");
        std::vector<std::vector<unsigned char>> hash_keys;

        hash_keys.push_back(data.auxstack_results.storage_key);
        hash_keys.push_back(data.datastack_results.storage_key);
        hash_keys.push_back(data.inbox_count_results.storage_key);
        hash_keys.push_back(data.inbox_messages_results.storage_key);
        hash_keys.push_back(data.pc_results.storage_key);
        hash_keys.push_back(data.pending_count_results.storage_key);
        hash_keys.push_back(data.pending_messages_results.storage_key);
        hash_keys.push_back(data.register_val_results.storage_key);
        hash_keys.push_back(data.static_val_results.storage_key);

        deleteCheckpoint(saver, "checkpoint", hash_keys);
    }
    //    SECTION("delete checkpoint saved twice") {
    //        TuplePool pool;
    //        CheckpointStorage storage(path);
    //        auto saver = MachineStateSaver(&storage, &pool);
    //
    //        auto data_values = getStateValues(saver);
    //        auto data = std::get<0>(data_values);
    //
    //        saver.saveMachineState(data, "checkpoint");
    //        saver.saveMachineState(data, "checkpoint");
    //        std::vector<std::vector<unsigned char>> hash_keys;
    //
    //        hash_keys.push_back(data.auxstack_results.storage_key);
    //        hash_keys.push_back(data.datastack_results.storage_key);
    //        hash_keys.push_back(data.inbox_count_results.storage_key);
    //        hash_keys.push_back(data.inbox_messages_results.storage_key);
    //        hash_keys.push_back(data.pc_results.storage_key);
    //        hash_keys.push_back(data.pending_count_results.storage_key);
    //        hash_keys.push_back(data.pending_messages_results.storage_key);
    //        hash_keys.push_back(data.register_val_results.storage_key);
    //        hash_keys.push_back(data.static_val_results.storage_key);
    //
    //        deleteCheckpointSavedTwice(saver, "checkpoint", hash_keys);
    //    }
}
