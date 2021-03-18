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

void runCheckArbCore(
    std::shared_ptr<ArbCore>& arbCore,
    const std::vector<std::vector<unsigned char>>& raw_messages,
    uint256_t prev_inbox_acc,
    uint256_t target_message_count,
    int send_count,
    int log_count,
    bool last_message) {
    auto initial_count_res = arbCore->messageEntryInsertedCount();
    REQUIRE(initial_count_res.status.ok());

    REQUIRE(arbCore->deliverMessages(raw_messages, prev_inbox_acc, last_message,
                                     std::nullopt));

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

    for (const auto& filename : files) {
        INFO("Testing " << filename);

        ArbStorage storage(dbpath);
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

        std::vector<std::vector<unsigned char>> raw_messages;
        raw_messages.reserve(inbox_messages.size());
        for (const auto& msg : inbox_messages) {
            raw_messages.push_back(msg.serialize());
        }

        for (size_t k = 0; k < raw_messages.size(); ++k) {
            auto msg = extractInboxMessage(raw_messages[k]);
            auto msg_tup = msg.toTuple();
            REQUIRE(hash(msg_tup) == hash(inbox_message_tuples[k]));
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

        runCheckArbCore(arbCore, raw_messages, 0, raw_messages.size(),
                        sends.size(), logs.size(), false);

        auto logsRes = arbCore->getLogs(0, logs.size(), value_cache);
        REQUIRE(logsRes.status.ok());
        REQUIRE(logsRes.data.size() == logs.size());
        for (size_t k = 0; k < logs.size(); ++k) {
            REQUIRE(logsRes.data[k] == logs[k]);
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
                        REQUIRE(result.data.logs[k] == logs[logs_count + k]);
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

        auto cursor = arbCore->getExecutionCursor(0, value_cache);
        REQUIRE(cursor.status.ok());
        REQUIRE(cursor.data->getOutput().arb_gas_used == 0);

        auto advanceStatus = arbCore->advanceExecutionCursor(
            *cursor.data, 100, false, value_cache);
        REQUIRE(advanceStatus.ok());
        REQUIRE(cursor.data->getOutput().arb_gas_used > 0);

        //        auto before_sideload = arbCore->getMachineForSideload(
        //            inbox_messages.back().block_number, value_cache);
        //        REQUIRE(before_sideload.status.ok());
        //        REQUIRE(before_sideload.data->machine_state.loadCurrentInstruction()
        //                    .op.opcode == OpCode::SIDELOAD);
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
         let sameBlockNum = asm(blockNum,) bool { inboxpeek };
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
    ValueCache value_cache{1, 0};

    ArbStorage storage(dbpath);
    REQUIRE(
        storage.initialize(std::string{machine_test_cases_path} + "/inbox.mexe")
            .ok());
    auto arbCore = storage.getArbCore();
    REQUIRE(arbCore->startThread());

    uint256_t inbox_acc = 0;
    for (int i = 0; i < 5; i++) {
        std::vector<std::vector<unsigned char>> raw_messages;
        auto message = InboxMessage(0, {}, i, 0, i, 0, {});
        raw_messages.push_back(message.serialize());
        INFO("RUN " << i);
        runCheckArbCore(arbCore, raw_messages, inbox_acc, i + 1, 0, i + 1,
                        true);
        inbox_acc = hash_inbox(inbox_acc, message.serialize());
    }
    auto tx = storage.makeReadTransaction();
    auto position = arbCore->getSideloadPosition(*tx, 1);
    REQUIRE(position.status.ok());

    auto cursor = arbCore->getExecutionCursor(position.data, value_cache);
    REQUIRE(cursor.status.ok());
    REQUIRE(cursor.data->getOutput().arb_gas_used > 0);
    REQUIRE(cursor.data->getOutput().arb_gas_used <= position.data);
    REQUIRE(cursor.data->machineHash().has_value());
}
