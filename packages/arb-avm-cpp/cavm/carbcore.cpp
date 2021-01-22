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

int deliverMessages(CArbCore* arb_core_ptr,
                    const uint64_t first_message_sequence_number,
                    const uint64_t block_height,
                    void* inbox_messages,
                    void* inbox_hashes_ptr,
                    void* previous_inbox_hash_ptr) {
    auto arb_core = static_cast<ArbCore*>(arb_core_ptr);
    auto messages = getInboxMessages(inbox_messages);
    auto inbox_hashes = receiveUint256Vector(inbox_hashes_ptr, messages.size());
    auto previous_inbox_hash = receiveUint256(previous_inbox_hash_ptr);

    try {
        arb_core->deliverMessages(first_message_sequence_number, block_height,
                                  messages, inbox_hashes, previous_inbox_hash);
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
                               const void* count_ptr) {
    try {
        auto index_result = static_cast<ArbCore*>(arb_core_ptr)
                                ->getLogAcc(receiveUint256(start_acc_hash),
                                            receiveUint256(start_index_ptr),
                                            receiveUint256(count_ptr));
        return returnUint256Result(index_result);
    } catch (const std::exception& e) {
        return {{}, false};
    }
}
