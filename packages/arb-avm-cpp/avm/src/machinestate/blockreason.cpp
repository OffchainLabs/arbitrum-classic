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

#include <avm/machinestate/blockreason.hpp>

#include <avm_values/value.hpp>

struct CheckpointSerializer {
    std::vector<unsigned char> operator()(const NotBlocked& val) const {
        auto block_type = static_cast<unsigned char>(val.type);
        std::vector<unsigned char> return_value{block_type};
        return return_value;
    }
    std::vector<unsigned char> operator()(const HaltBlocked& val) const {
        auto block_type = static_cast<unsigned char>(val.type);
        std::vector<unsigned char> return_value{block_type};

        return return_value;
    }
    std::vector<unsigned char> operator()(const ErrorBlocked& val) const {
        auto block_type = static_cast<unsigned char>(val.type);
        std::vector<unsigned char> return_value{block_type};

        return return_value;
    }
    std::vector<unsigned char> operator()(const BreakpointBlocked& val) const {
        auto block_type = static_cast<unsigned char>(val.type);
        std::vector<unsigned char> return_value{block_type};

        return return_value;
    }

    std::vector<unsigned char> operator()(const InboxBlocked& val) const {
        auto block_type = static_cast<unsigned char>(val.type);
        std::vector<unsigned char> return_value{block_type};

        std::vector<unsigned char> inbox_char_vector;
        marshal_uint256_t(val.inbox, inbox_char_vector);

        return_value.insert(return_value.end(), inbox_char_vector.begin(),
                            inbox_char_vector.end());

        return return_value;
    }
};

std::vector<unsigned char> serializeForCheckpoint(const BlockReason& val) {
    return nonstd::visit(CheckpointSerializer{}, val);
}

constexpr BlockType InboxBlocked::type;

BlockReason deserializeBlockReason(const std::vector<unsigned char>& data) {
    auto blocktype = static_cast<BlockType>(data[0]);
    switch (blocktype) {
        case Inbox: {
            auto buff = reinterpret_cast<const char*>(&data[2]);
            auto inbox = deserializeUint256t(buff);
            return InboxBlocked(inbox);
        }
        case Halt: {
            return HaltBlocked();
        }
        case Error: {
            return ErrorBlocked();
        }
        case Breakpoint: {
            return BreakpointBlocked();
        }
        default: {
            return NotBlocked();
        }
    }
}
