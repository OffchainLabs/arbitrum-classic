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

#include "carbcore.h"

#include "utils.hpp"

#include <data_storage/arbcore.hpp>

int arbCoreStartThread(CArbCore* arbcore_ptr) {
    auto arb_core = static_cast<ArbCore*>(arbcore_ptr);
    auto status = arb_core->startThread();
    if (!status) {
        return 0;
    }

    return 1;
}

void arbCoreAbortThread(CArbCore* arbcore_ptr) {
    auto arb_core = static_cast<ArbCore*>(arbcore_ptr);
    arb_core->abortThread();
}

int arbCoreMachineIdle(CArbCore* arbcore_ptr) {
    auto arb_core = static_cast<ArbCore*>(arbcore_ptr);
    return arb_core->machineIdle();
}

int arbCoreDeliverMessages(CArbCore* arbcore_ptr,
                           ByteSliceArray inbox_messages,
                           void* previous_inbox_hash_ptr) {
    auto arb_core = static_cast<ArbCore*>(arbcore_ptr);
    auto messages = receiveByteSliceArray(inbox_messages);
    auto previous_inbox_hash = receiveUint256(previous_inbox_hash_ptr);

    try {
        arb_core->deliverMessages(messages, previous_inbox_hash);
    } catch (const std::exception& e) {
        return false;
    }

    return true;
}

Uint256Result arbCoreGetLogCount(CArbCore* arbcore_ptr) {
    auto arb_core = static_cast<ArbCore*>(arbcore_ptr);
    try {
        auto count_result = arb_core->logInsertedCount();
        if (!count_result.status.ok()) {
            return {{}, false};
        }
        return {returnUint256(count_result.data), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

ByteSliceArrayResult arbCoreGetLogs(CArbCore* arbcore_ptr,
                                    const void* start_index_ptr,
                                    const void* count_ptr) {
    try {
        ValueCache cache;
        auto logs = static_cast<ArbCore*>(arbcore_ptr)
                        ->getLogs(receiveUint256(start_index_ptr),
                                  receiveUint256(count_ptr), cache);
        if (!logs.status.ok()) {
            return {{}, false};
        }

        std::vector<std::vector<unsigned char>> data;
        for (const auto& val : logs.data) {
            std::vector<unsigned char> marshalled_value;
            marshal_value(val, marshalled_value);
            data.push_back(move(marshalled_value));
        }
        return {returnCharVectorVector(data), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

Uint256Result arbCoreGetSendCount(CArbCore* arbcore_ptr) {
    auto arb_core = static_cast<ArbCore*>(arbcore_ptr);
    try {
        auto count_result = arb_core->sendInsertedCount();
        if (!count_result.status.ok()) {
            return {{}, false};
        }
        return {returnUint256(count_result.data), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

ByteSliceArrayResult arbCoreGetSends(CArbCore* arbcore_ptr,
                                     const void* start_index_ptr,
                                     const void* count_ptr) {
    try {
        auto sends = static_cast<const ArbCore*>(arbcore_ptr)
                         ->getSends(receiveUint256(start_index_ptr),
                                    receiveUint256(count_ptr));
        if (!sends.status.ok()) {
            return {{}, false};
        }

        return {returnCharVectorVector(sends.data), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

Uint256Result arbCoreGetMessageCount(CArbCore* arbcore_ptr) {
    auto arb_core = static_cast<ArbCore*>(arbcore_ptr);
    try {
        auto count_result = arb_core->messageEntryInsertedCount();
        if (!count_result.status.ok()) {
            return {{}, false};
        }
        return {returnUint256(count_result.data), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

ByteSliceArrayResult arbCoreGetMessages(CArbCore* arbcore_ptr,
                                        const void* start_index_ptr,
                                        const void* count_ptr) {
    try {
        auto messages = static_cast<const ArbCore*>(arbcore_ptr)
                            ->getMessages(receiveUint256(start_index_ptr),
                                          receiveUint256(count_ptr));
        if (!messages.status.ok()) {
            return {{}, false};
        }

        return {returnCharVectorVector(messages.data), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

HashList arbCoreGetMessageHashes(CArbCore* arbcore_ptr,
                                 const void* start_index_ptr,
                                 const void* count_ptr) {
    try {
        auto hashes = static_cast<const ArbCore*>(arbcore_ptr)
                          ->getMessageHashes(receiveUint256(start_index_ptr),
                                             receiveUint256(count_ptr));
        if (!hashes.status.ok()) {
            return {{}, false};
        }

        std::vector<unsigned char> serializedHashes;
        for (const auto& hash : hashes.data) {
            marshal_uint256_t(hash, serializedHashes);
        }
        auto hashesData =
            reinterpret_cast<unsigned char*>(malloc(serializedHashes.size()));
        std::copy(serializedHashes.begin(), serializedHashes.end(), hashesData);
        return {hashesData, static_cast<int>(hashes.data.size())};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

int arbCoreGetInboxDelta(CArbCore* arbcore_ptr,
                         const void* start_index_ptr,
                         const void* count_ptr,
                         void* ret) {
    auto arb_core = static_cast<ArbCore*>(arbcore_ptr);
    try {
        auto index_result = arb_core->getInboxDelta(
            receiveUint256(start_index_ptr), receiveUint256(count_ptr));
        std::array<unsigned char, 32> val{};
        to_big_endian(index_result.data, val.begin());
        std::copy(val.begin(), val.end(), reinterpret_cast<char*>(ret));
        return true;
    } catch (const std::exception& e) {
        return false;
    }
}

int arbCoreGetInboxAcc(CArbCore* arbcore_ptr,
                       const void* index_ptr,
                       void* ret) {
    auto arb_core = static_cast<ArbCore*>(arbcore_ptr);
    try {
        auto index_result = arb_core->getInboxAcc(receiveUint256(index_ptr));
        std::array<unsigned char, 32> val{};
        to_big_endian(index_result.data, val.begin());
        std::copy(val.begin(), val.end(), reinterpret_cast<char*>(ret));
        return true;
    } catch (const std::exception& e) {
        return false;
    }
}

int arbCoreGetSendAcc(CArbCore* arbcore_ptr,
                      const void* start_acc_hash,
                      const void* start_index_ptr,
                      const void* count_ptr,
                      void* ret) {
    auto arb_core = static_cast<ArbCore*>(arbcore_ptr);
    try {
        auto index_result = arb_core->getSendAcc(
            receiveUint256(start_acc_hash), receiveUint256(start_index_ptr),
            receiveUint256(count_ptr));
        std::array<unsigned char, 32> val{};
        to_big_endian(index_result.data, val.begin());
        std::copy(val.begin(), val.end(), reinterpret_cast<char*>(ret));
        return true;
    } catch (const std::exception& e) {
        return false;
    }
}

int arbCoreGetLogAcc(CArbCore* arbcore_ptr,
                     const void* start_acc_hash,
                     const void* start_index_ptr,
                     const void* count_ptr,
                     void* ret) {
    auto arbcore = static_cast<ArbCore*>(arbcore_ptr);
    ValueCache cache;

    try {
        auto index_result = arbcore->getLogAcc(
            receiveUint256(start_acc_hash), receiveUint256(start_index_ptr),
            receiveUint256(count_ptr), cache);
        std::array<unsigned char, 32> val{};
        to_big_endian(index_result.data, val.begin());
        std::copy(val.begin(), val.end(), reinterpret_cast<char*>(ret));
        return true;
    } catch (const std::exception& e) {
        return false;
    }
}

int arbCoreLogsCursorRequest(CArbCore* arbcore_ptr, const void* count_ptr) {
    auto arbcore = static_cast<ArbCore*>(arbcore_ptr);
    auto count = receiveUint256(count_ptr);

    try {
        auto status = arbcore->logsCursorRequest(count);

        return status;
    } catch (const std::exception& e) {
        return false;
    }
}

ByteSliceArrayResult arbCoreLogsCursorGetLogs(CArbCore* arbcore_ptr) {
    auto arbcore = static_cast<ArbCore*>(arbcore_ptr);

    try {
        auto result = arbcore->logsCursorGetLogs();
        if (!result) {
            // Cursor not in the right state, may have deleted logs to process
            return {{}, false};
        }

        std::vector<std::vector<unsigned char>> data;
        for (const auto& val : *result) {
            std::vector<unsigned char> marshalled_value;
            marshal_value(val, marshalled_value);
            data.push_back(move(marshalled_value));
        }
        return {returnCharVectorVector(data), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

int arbCoreLogsCursorSetNextIndex(CArbCore* arbcore_ptr,
                                  const void* count_ptr) {
    auto arbcore = static_cast<ArbCore*>(arbcore_ptr);
    auto count = receiveUint256(count_ptr);

    try {
        auto status = arbcore->logsCursorSetNextIndex(count);

        return status;
    } catch (const std::exception& e) {
        return 0;
    }
}

int arbCoreLogsCursorCheckError(CArbCore* arbcore_ptr) {
    auto arbcore = static_cast<ArbCore*>(arbcore_ptr);

    try {
        return arbcore->logsCursorCheckError();
    } catch (const std::exception& e) {
        return 0;
    }
}

// Returned string must be freed
char* arbCoreLogsCursorClearError(CArbCore* arbcore_ptr) {
    auto arbcore = static_cast<ArbCore*>(arbcore_ptr);

    try {
        auto str = arbcore->logsCursorClearError();

        if (str.empty()) {
            return nullptr;
        }

        return strdup(str.c_str());
    } catch (const std::exception& e) {
        return strdup("exception occurred in logsCursorClearError");
    }
}

CExecutionCursor* arbCoreGetExecutionCursor(CArbCore* arbcore_ptr,
                                            const void* total_gas_used_ptr) {
    auto arbcore = static_cast<ArbCore*>(arbcore_ptr);
    ValueCache cache;
    auto total_gas_used = receiveUint256(total_gas_used_ptr);

    try {
        auto executionCursor =
            arbcore->getExecutionCursor(total_gas_used, cache);
        if (!executionCursor.status.ok()) {
            return nullptr;
        }
        return static_cast<void*>(executionCursor.data.release());
    } catch (const std::exception& e) {
        return nullptr;
    }
}

int arbCoreAdvanceExecutionCursor(CArbCore* arbcore_ptr,
                                  CExecutionCursor* execution_cursor_ptr,
                                  const void* max_gas_ptr,
                                  int go_over_gas) {
    auto arbCore = static_cast<ArbCore*>(arbcore_ptr);
    auto executionCursor = static_cast<ExecutionCursor*>(execution_cursor_ptr);
    auto max_gas = receiveUint256(max_gas_ptr);
    try {
        ValueCache cache;
        auto status =
            arbCore->Advance(*executionCursor, max_gas, go_over_gas, cache);
        if (!status.ok()) {
            return false;
        }

        return true;
    } catch (const std::exception& e) {
        return false;
    }
}
