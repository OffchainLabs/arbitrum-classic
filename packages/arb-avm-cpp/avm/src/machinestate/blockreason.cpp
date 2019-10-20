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

#include "avm/machinestate/blockreason.hpp"
#include "avm/machinestate/statesaverutils.hpp"

struct CheckpointSerializer {
    std::vector<unsigned char> operator()(const NotBlocked& val) const {
        std::vector<unsigned char> return_value;
        return_value.push_back((unsigned char)val.type);
        return return_value;
    }
    std::vector<unsigned char> operator()(const HaltBlocked& val) const {
        std::vector<unsigned char> return_value;
        return_value.push_back((unsigned char)val.type);

        return return_value;
    }
    std::vector<unsigned char> operator()(const ErrorBlocked& val) const {
        std::vector<unsigned char> return_value;
        return_value.push_back((unsigned char)val.type);

        return return_value;
    }
    std::vector<unsigned char> operator()(const BreakpointBlocked& val) const {
        std::vector<unsigned char> return_value;
        return_value.push_back((unsigned char)val.type);

        return return_value;
    }

    std::vector<unsigned char> operator()(const InboxBlocked& val) const {
        std::vector<unsigned char> return_value;
        return_value.push_back((unsigned char)val.type);

        std::vector<unsigned char> inbox_char_vector;
        marshal_uint256_t(val.inbox, inbox_char_vector);

        return_value.insert(return_value.end(), inbox_char_vector.begin(),
                            inbox_char_vector.end());

        return return_value;
    }
    std::vector<unsigned char> operator()(const SendBlocked& val) const {
        std::vector<unsigned char> return_value;
        return_value.push_back((unsigned char)val.type);

        std::vector<unsigned char> data_vector;
        marshal_uint256_t(val.currency, data_vector);

        return_value.insert(return_value.end(), data_vector.begin(),
                            data_vector.end());

        return_value.insert(return_value.end(), std::begin(val.tokenType),
                            std::end(val.tokenType));

        return return_value;
    }
};

std::unordered_map<BlockType, int> blockreason_type_length = {
    {Not, 1}, {Halt, 1}, {Error, 1}, {Breakpoint, 1}, {Inbox, 34}, {Send, 55}};

std::vector<unsigned char> serializeForCheckpoint(const BlockReason& val) {
    return nonstd::visit(CheckpointSerializer{}, val);
}

BlockReason deserializeBlockReason(std::vector<unsigned char> data) {
    auto current_it = data.begin();
    auto blocktype = (BlockType)*current_it;
    current_it++;

    switch (blocktype) {
        case Inbox: {
            auto next_it = current_it + TOKEN_VAL_LENGTH;
            std::vector<unsigned char> inbox_vector(current_it, next_it);
            auto inbox = Checkpoint::deserializeUint256(inbox_vector);
            return InboxBlocked(inbox);
        }
        case Send: {
            auto next_it = current_it + TOKEN_VAL_LENGTH;
            std::vector<unsigned char> currency_vector(current_it, next_it);
            auto currency = Checkpoint::deserializeUint256(currency_vector);

            current_it = next_it;
            next_it = current_it + TOKEN_TYPE_LENGTH;

            std::array<unsigned char, TOKEN_TYPE_LENGTH> token_type;
            std::copy(current_it, next_it, token_type.begin());

            return SendBlocked(currency, token_type);
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
