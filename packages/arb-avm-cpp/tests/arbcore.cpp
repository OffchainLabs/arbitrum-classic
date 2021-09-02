/*
 * Copyright 2020, Offchain Labs, Inc.
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
#include "helper.hpp"

#include <data_storage/arbstorage.hpp>
#include <data_storage/storageresult.hpp>

#include <avm/inboxmessage.hpp>

#include <avm_values/vmValueParser.hpp>

#include <catch2/catch.hpp>
#include <nlohmann/json.hpp>

std::vector<unsigned char> serializeForCore(
    const SequencerBatchItem& batch_item) {
    std::vector<unsigned char> full_batch_data;
    marshal_uint256_t(batch_item.last_sequence_number, full_batch_data);
    auto batch_data_part = serializeSequencerBatchItem(batch_item);
    full_batch_data.insert(full_batch_data.end(), batch_data_part.begin(),
                           batch_data_part.end());
    return full_batch_data;
}

std::vector<SequencerBatchItem> buildBatch(
    const std::vector<InboxMessage>& inbox_messages) {
    std::vector<SequencerBatchItem> seq_batch_items;
    seq_batch_items.reserve(inbox_messages.size());
    uint256_t last_item_acc = 0;
    for (const auto& msg : inbox_messages) {
        SequencerBatchItem batch_item = {msg.inbox_sequence_number, 0, 0,
                                         msg.serialize()};
        batch_item.accumulator =
            batch_item.computeAccumulator(last_item_acc, 0, 0);
        last_item_acc = batch_item.accumulator;
        seq_batch_items.push_back(batch_item);
    }
    return seq_batch_items;
}

void waitForDelivery(std::shared_ptr<ArbCore>& arbCore) {
    ArbCore::message_status_enum status;
    while (true) {
        status = arbCore->messagesStatus();
        if (status != ArbCore::MESSAGES_EMPTY &&
            status != ArbCore::MESSAGES_READY) {
            break;
        }
        std::this_thread::sleep_for(std::chrono::milliseconds(100));
    }
    if (status == ArbCore::MESSAGES_ERROR) {
        INFO(arbCore->messagesClearError());
    }
    REQUIRE(status == ArbCore::MESSAGES_SUCCESS);
}

void runCheckArbCore(std::shared_ptr<ArbCore>& arbCore,
                     const std::vector<SequencerBatchItem>& seq_batch_items,
                     uint256_t prev_message_count,
                     uint256_t prev_inbox_acc,
                     uint256_t target_message_count,
                     uint256_t send_count,
                     uint256_t log_count) {
    auto initial_count_res = arbCore->messageEntryInsertedCount();
    REQUIRE(initial_count_res.status.ok());

    std::vector<std::vector<unsigned char>> raw_seq_batch_items;
    raw_seq_batch_items.reserve(seq_batch_items.size());
    for (const auto& batch_item : seq_batch_items) {
        raw_seq_batch_items.push_back(serializeForCore(batch_item));
    }

    REQUIRE(arbCore->deliverMessages(
        prev_message_count, prev_inbox_acc, raw_seq_batch_items,
        std::vector<std::vector<unsigned char>>(), std::nullopt));

    waitForDelivery(arbCore);

    int tries = 0;
    while (true) {
        auto countRes = arbCore->messageEntryInsertedCount();
        REQUIRE(countRes.status.ok());
        if (countRes.data == target_message_count) {
            break;
        }
        std::this_thread::sleep_for(std::chrono::milliseconds(1000));
        tries++;
        REQUIRE(tries < 5);
    }

    auto accRes = arbCore->getInboxAcc(target_message_count - 1);
    REQUIRE(accRes.status.ok());
    REQUIRE(accRes.data != 0);

    while (!arbCore->machineIdle()) {
        auto err_str = arbCore->machineClearError();
        REQUIRE(!err_str.has_value());
        std::this_thread::sleep_for(std::chrono::milliseconds(1000));
    }

    auto producedLogCountRes = arbCore->logInsertedCount();
    REQUIRE(producedLogCountRes.status.ok());
    REQUIRE(producedLogCountRes.data == log_count);

    auto producedSendCountRes = arbCore->sendInsertedCount();
    REQUIRE(producedSendCountRes.status.ok());
    REQUIRE(producedSendCountRes.data == send_count);
}

TEST_CASE("ArbCore tests") {
    DBDeleter deleter;
    ValueCache value_cache{1, 0};

    std::vector<std::string> files = {
        "evm_direct_deploy_add", "evm_direct_deploy_and_call_add",
        "evm_test_arbsys", "evm_xcontract_call_with_constructors"};

    uint64_t logs_count = 0;
    ArbCoreConfig coreConfig{};

    for (const auto& filename : files) {
        INFO("Testing " << filename);

        ArbStorage storage(dbpath, coreConfig);
        REQUIRE(storage.initialize(arb_os_path).ok());
        auto arbCore = storage.getArbCore();
        REQUIRE(arbCore->startThread());

        auto test_file =
            std::string{arb_os_test_cases_path} + "/" + filename + ".aoslog";

        std::ifstream i(test_file);
        nlohmann::json j;
        i >> j;

        std::vector<Tuple> inbox_message_tuples;
        for (auto& json_message : j.at("inbox")) {
            auto tup = std::get<Tuple>(simple_value_from_json(json_message));
            inbox_message_tuples.push_back(std::move(tup));
        }

        std::vector<InboxMessage> inbox_messages;
        inbox_messages.reserve(inbox_message_tuples.size());
        for (const auto& msg : inbox_message_tuples) {
            inbox_messages.push_back(InboxMessage::fromTuple(msg));
        }

        auto logs_json = j.at("logs");
        std::vector<value> logs;
        for (auto& log_json : logs_json) {
            logs.push_back(simple_value_from_json(log_json));
        }

        auto sends_json = j.at("sends");
        std::vector<std::vector<uint8_t>> sends;
        for (auto& send_json : sends_json) {
            sends.push_back(send_from_json(send_json));
        }

        runCheckArbCore(arbCore, buildBatch(inbox_messages), 0, 0,
                        inbox_messages.size(), sends.size(), logs.size());

        auto logsRes = arbCore->getLogs(0, logs.size(), value_cache);
        REQUIRE(logsRes.status.ok());
        REQUIRE(logsRes.data.size() == logs.size());
        for (size_t k = 0; k < logs.size(); ++k) {
            REQUIRE(logsRes.data[k].val == logs[k]);
        }

        auto sendsRes = arbCore->getSends(0, sends.size());
        REQUIRE(sendsRes.status.ok());
        REQUIRE(sendsRes.data.size() == sends.size());
        for (size_t k = 0; k < sends.size(); ++k) {
            REQUIRE(sendsRes.data[k] == sends[k]);
        }

        int tries = 0;
        bool done = false;
        while (!done) {
            auto log_request_count = 3;
            REQUIRE(arbCore->logsCursorRequest(0, log_request_count));
            while (true) {
                auto result = arbCore->logsCursorGetLogs(0);
                REQUIRE((result.status.ok() || result.status.IsTryAgain()));
                REQUIRE(!arbCore->logsCursorCheckError(0));
                if (result.status.ok()) {
                    REQUIRE(result.data.deleted_logs.size() <= logs_count);
                    logs_count -= result.data.deleted_logs.size();
                    REQUIRE(result.data.first_log_index == logs_count);
                    REQUIRE(result.data.logs.size() <=
                            logs.size() - logs_count);
                    for (uint64_t k = 0; k < result.data.logs.size(); ++k) {
                        REQUIRE(result.data.logs[k].val ==
                                logs[logs_count + k]);
                    }
                    logs_count += result.data.logs.size();
                    REQUIRE(arbCore->logsCursorConfirmReceived(0));
                    if (logs_count == logs.size()) {
                        done = true;
                    }
                    break;
                }
                std::this_thread::sleep_for(std::chrono::milliseconds(100));
            }
            REQUIRE(tries < 20);
            tries++;
        }
        REQUIRE(logs_count == logs.size());

        auto cursor = arbCore->getExecutionCursor(0);
        REQUIRE(cursor.status.ok());
        REQUIRE(cursor.data->getOutput().arb_gas_used == 0);

        auto advanceStatus =
            arbCore->advanceExecutionCursor(*cursor.data, 100, false);
        REQUIRE(advanceStatus.ok());
        REQUIRE(cursor.data->getOutput().arb_gas_used > 0);

        //        auto before_sideload = arbCore->getMachineForSideload(
        //            inbox_messages.back().block_number, value_cache);
        //        REQUIRE(before_sideload.status.ok());
        //        REQUIRE(before_sideload.data->machine_state.loadCurrentInstruction()
        //                    .op.opcode == OpCode::SIDELOAD);

        auto final_output = arbCore->getLastMachineOutput();

        // Create a new arbCore and verify it gets to the same point
        storage.closeArbStorage();
        storage = ArbStorage(dbpath, coreConfig);
        REQUIRE(storage.initialize(arb_os_path).ok());
        arbCore = storage.getArbCore();
        logs_count = uint64_t(arbCore->getLastMachineOutput().log_count);
        if (!sends.empty()) {
            // If there were sends, the block must have ended, so there should
            // be a checkpoint present
            REQUIRE(
                arbCore->getLastMachineOutput().fully_processed_inbox.count >
                0);
        }
        REQUIRE(arbCore->startThread());

        int n = 0;
        while (arbCore->getLastMachineOutput().arb_gas_used <
               final_output.arb_gas_used) {
            std::this_thread::sleep_for(std::chrono::milliseconds(100));
            REQUIRE(n++ < 100);
        }
        REQUIRE(arbCore->getLastMachineOutput() == final_output);
    }
}

/*
 Test file in separate repo, but source code of mini program is included here
 for reference

 type IncomingRequestFromInbox = struct {
     kind: uint,               // type of message
     ethBlockNumber: uint,     // block number of the L1 block
     timestamp: uint,          // timestamp of the L1 block
     sender: address,          // address of the sender
     requestId: uint,
     gasPriceL1: uint,         // L1 gas price paid by this tx
     msgSize: uint,
     msgData: buffer,
 }

 impure func main() {
     let blockNum = 0;
     loop {
         let rawSideloadMsg = asm(blockNum,) any { sideload };
         if (rawSideloadMsg != ()) {
             panic;
         }
         let newMsg = asm() IncomingRequestFromInbox { inbox };
         blockNum = newMsg.ethBlockNumber;
         asm(blockNum,) { log };
     }
 }

 */
TEST_CASE("ArbCore inbox") {
    DBDeleter deleter;

    ArbCoreConfig coreConfig{};
    ArbStorage storage(dbpath, coreConfig);
    REQUIRE(
        storage.initialize(std::string{machine_test_cases_path} + "/inbox.mexe")
            .ok());
    auto arbCore = storage.getArbCore();
    REQUIRE(arbCore->startThread());

    std::vector<InboxMessage> inbox_messages;
    for (int i = 0; i < 5; i++) {
        auto message = InboxMessage(0, {}, i, 0, i, 0, {});
        inbox_messages.push_back(message);
    }
    auto items = buildBatch(inbox_messages);

    uint256_t inbox_acc = 0;
    for (int i = 0; i < 5; i++) {
        auto batch_item = items[i];
        INFO("RUN " << i);
        runCheckArbCore(arbCore, {batch_item}, i, inbox_acc, i + 1, 0, i + 1);
        inbox_acc = batch_item.accumulator;
    }
    auto tx = storage.makeReadTransaction();
    auto position = arbCore->getSideloadPosition(*tx, 1);
    REQUIRE(position.status.ok());

    auto cursor = arbCore->getExecutionCursor(position.data);
    REQUIRE(cursor.status.ok());
    REQUIRE(cursor.data->getOutput().arb_gas_used > 0);
    REQUIRE(cursor.data->getOutput().arb_gas_used <= position.data);

    auto cursor_machine_hash = cursor.data->machineHash();
    REQUIRE(cursor_machine_hash.has_value());

    auto cursor_machine = arbCore->takeExecutionCursorMachine(*cursor.data);
    REQUIRE(cursor_machine);
    REQUIRE(cursor_machine_hash.value() == cursor_machine->hash());

    auto machine = arbCore->getLastMachine();
    REQUIRE(machine);
}

TEST_CASE("ArbCore backwards reorg") {
    ArbCoreConfig coreConfig{};
    ArbStorage storage(dbpath, coreConfig);
    REQUIRE(
        storage.initialize(std::string{machine_test_cases_path} + "/inbox.mexe")
            .ok());
    auto arbCore = storage.getArbCore();
    REQUIRE(arbCore->startThread());

    REQUIRE(arbCore->deliverMessages(
        0, 0, std::vector<std::vector<unsigned char>>(),
        std::vector<std::vector<unsigned char>>(), 0));
    waitForDelivery(arbCore);
    REQUIRE(arbCore->messageEntryInsertedCount().data == 0);

    auto maxGas = std::numeric_limits<uint256_t>::max();
    auto initialState = arbCore->getExecutionCursor(maxGas);
    REQUIRE(initialState.status.ok());
    REQUIRE(initialState.data->getTotalMessagesRead() == 0);

    auto message = InboxMessage(0, {}, 0, 0, 0, 0, {});

    std::vector<std::vector<unsigned char>> rawSeqBatchItems;
    for (const auto& batch_item : buildBatch(std::vector(1, message))) {
        rawSeqBatchItems.push_back(serializeForCore(batch_item));
    }

    REQUIRE(arbCore->deliverMessages(0, 0, rawSeqBatchItems,
                                     std::vector<std::vector<unsigned char>>(),
                                     std::nullopt));
    waitForDelivery(arbCore);

    auto newState = arbCore->getExecutionCursor(maxGas);
    REQUIRE(newState.status.ok());
    REQUIRE(newState.data->getTotalMessagesRead() == 1);

    REQUIRE(arbCore->deliverMessages(
        0, 0, std::vector<std::vector<unsigned char>>(),
        std::vector<std::vector<unsigned char>>(), 0));
    waitForDelivery(arbCore);

    auto reorgState = arbCore->getExecutionCursor(maxGas);
    REQUIRE(reorgState.status.ok());
    REQUIRE(reorgState.data->getTotalMessagesRead() == 0);
    REQUIRE(reorgState.data->machineHash() == initialState.data->machineHash());
    REQUIRE(arbCore->getLastMachine()
                ->machine_state.output.fully_processed_inbox.count == 0);
}

TEST_CASE("ArbCore duplicate code segments") {
    DBDeleter deleter;

    ArbCoreConfig coreConfig{};
    coreConfig.min_gas_checkpoint_frequency = 1;
    ArbStorage storage(dbpath, coreConfig);
    REQUIRE(storage
                .initialize(std::string{machine_test_cases_path} +
                            "/dupsegments.mexe")
                .ok());
    auto arbCore = storage.getArbCore();
    REQUIRE(arbCore->startThread());

    constexpr int CHECKPOINTS = 2;

    std::vector<InboxMessage> messages;
    for (int i = 0; i < CHECKPOINTS; i++) {
        messages.push_back(InboxMessage(0, {}, 0, 0, i, 0, {}));
    }
    auto batch = buildBatch(messages);
    REQUIRE(batch.size() == CHECKPOINTS);

    uint256_t last_acc = 0;
    for (int i = 0; i < CHECKPOINTS; i++) {
        auto batch_item = batch[i];
        std::vector<std::vector<unsigned char>> rawSeqBatchItems(
            1, serializeForCore(batch_item));

        REQUIRE(arbCore->deliverMessages(
            i, last_acc, rawSeqBatchItems,
            std::vector<std::vector<unsigned char>>(), std::nullopt));
        waitForDelivery(arbCore);
        last_acc = batch_item.accumulator;

        int j = 0;
        while (arbCore->getLastMachineOutput().last_sideload != i ||
               !arbCore->machineIdle()) {
            std::this_thread::sleep_for(std::chrono::milliseconds(100));
            REQUIRE(j++ < 10);
        }

        if (i == 0) {
            // Restart ArbCore
            storage.closeArbStorage();
            storage = ArbStorage(dbpath, coreConfig);
            REQUIRE(storage
                        .initialize(std::string{machine_test_cases_path} +
                                    "/dupsegments.mexe")
                        .ok());
            arbCore = storage.getArbCore();
            REQUIRE(arbCore->startThread());
        }
    }

    auto cursor =
        arbCore->getExecutionCursor(std::numeric_limits<uint256_t>::max());
    REQUIRE(cursor.status.ok());
    REQUIRE(std::get<std::unique_ptr<Machine>>(cursor.data->machine)
                ->currentStatus() == Status::Halted);
}
