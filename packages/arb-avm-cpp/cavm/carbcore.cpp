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

void deleteArbCore(CArbCore* m) {
    delete static_cast<ArbCore*>(m);
}

int arbCoreDeliverMessages(CArbCore* arb_core_ptr,
                           const uint64_t first_message_sequence_number,
                           const uint64_t block_height,
                           void* inbox_messages,
                           void* previous_inbox_hash_ptr) {
    auto arb_core = static_cast<ArbCore*>(arb_core_ptr);
    auto messages = getInboxMessages(inbox_messages);
    auto previous_inbox_hash = receiveUint256(previous_inbox_hash_ptr);

    try {
        arb_core->deliverMessages(first_message_sequence_number, block_height,
                                  messages, previous_inbox_hash);
    } catch (const std::exception& e) {
        return false;
    }

    return true;
}

ByteSliceArrayResult arbCoreGetSends(CArbCore* arb_core_ptr,
                                     const void* start_index_ptr,
                                     const void* count_ptr) {
    try {
        auto sends = static_cast<const ArbCore*>(arb_core_ptr)
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

ByteSliceArrayResult arbCoreGetMessages(CArbCore* arb_core_ptr,
                                        const void* start_index_ptr,
                                        const void* count_ptr) {
    try {
        auto sends = static_cast<const ArbCore*>(arb_core_ptr)
                         ->getMessages(receiveUint256(start_index_ptr),
                                       receiveUint256(count_ptr));
        if (!sends.status.ok()) {
            return {{}, false};
        }

        return {returnCharVectorVector(sends.data), true};
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

Uint256Result arbCoreGetInboxDelta(CArbCore* arb_core_ptr,
                                   const void* start_index_ptr,
                                   const void* count_ptr) {
    try {
        auto index_result = static_cast<ArbCore*>(arb_core_ptr)
                                ->getInboxDelta(receiveUint256(start_index_ptr),
                                                receiveUint256(count_ptr));
        return returnUint256Result(index_result);
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

Uint256Result arbCoreGetInboxAcc(CArbCore* arb_core_ptr, const void* index) {
    try {
        auto hash = static_cast<ArbCore*>(arb_core_ptr)
                        ->getInboxAcc(receiveUint256(index));
        return returnUint256Result(hash);
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

Uint256Result arbCoreGetSendAcc(CArbCore* arb_core_ptr,
                                const void* start_acc_hash,
                                const void* start_index_ptr,
                                const void* count_ptr) {
    try {
        auto index_result = static_cast<ArbCore*>(arb_core_ptr)
                                ->getSendAcc(receiveUint256(start_acc_hash),
                                             receiveUint256(start_index_ptr),
                                             receiveUint256(count_ptr));
        return returnUint256Result(index_result);
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

Uint256Result arbCoreGetLogAcc(CArbCore* arb_core_ptr,
                               const void* start_acc_hash,
                               const void* start_index_ptr,
                               const void* count_ptr,
                               CValueCache* cache_ptr) {
    auto arbcore = static_cast<ArbCore*>(arb_core_ptr);
    auto cache = static_cast<ValueCache*>(cache_ptr);

    try {
        auto index_result = arbcore->getLogAcc(
            receiveUint256(start_acc_hash), receiveUint256(start_index_ptr),
            receiveUint256(count_ptr), *cache);
        return returnUint256Result(index_result);
    } catch (const std::exception& e) {
        return {{}, false};
    }
}

uint8_t arbCoreLogsCursorRequest(CArbCore* arb_core_ptr,
                                 const void* count_ptr) {
    auto arbcore = static_cast<ArbCore*>(arb_core_ptr);
    auto count = receiveUint256(count_ptr);

    try {
        auto status = arbcore->logsCursorRequest(count);

        return status;
    } catch (const std::exception& e) {
        return false;
    }
}

ByteSliceArrayResult arbCoreLogsCursorGetLogs(CArbCore* arb_core_ptr) {
    auto arbcore = static_cast<ArbCore*>(arb_core_ptr);

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

uint8_t arbCoreLogsCursorConfirmCount(CArbCore* arb_core_ptr,
                                      const void* count_ptr) {
    auto arbcore = static_cast<ArbCore*>(arb_core_ptr);
    auto count = receiveUint256(count_ptr);

    try {
        auto status = arbcore->logsCursorConfirmCount(count);

        return status;
    } catch (const std::exception& e) {
        return false;
    }
}

// Returned string must be freed
uint8_t arbCoreLogsCursorCheckError(CArbCore* arb_core_ptr) {
    auto arbcore = static_cast<ArbCore*>(arb_core_ptr);

    try {
        auto status = arbcore->logsCursorCheckError();

        return status;
    } catch (const std::exception& e) {
        return false;
    }
}

char* arbCoreLogsCursorClearError(CArbCore* arb_core_ptr) {
    auto arbcore = static_cast<ArbCore*>(arb_core_ptr);

    try {
        auto status = arbcore->logsCursorClearError();

        return strdup(status.c_str());
    } catch (const std::exception& e) {
        return strdup("exception occurred in logsCursorCheckError");
    }
}
