//
//  checkpointutils.cpp
//  avm
//
//  Created by Minh Truong on 9/22/19.
//

#include "avm/checkpointutils.hpp"

// std::vector<unsigned char> CombineData(
//    std::vector<unsigned char> balance_value,
//    std::vector<unsigned char> status_value,
//    std::vector<unsigned char> block_reason) {
//    status_value.insert(status_value.end(), balance_value.begin(),
//                        balance_value.end());
//
//    return status_value;
//}

// std::vector<unsigned char> SerializeData(BalanceTracker& balance,
//                                         Status& state,
//                                         BlockReason& blockReason) {
//    auto list_of_token_pairs = balance.GetAllTokenPairs();
//    std::vector<unsigned char> token_char_list;
//
//    for (auto& pair : list_of_token_pairs) {
//        auto tokentype = std::get<0>(pair);
//        auto value = std::get<1>(pair);
//
//        std::vector<unsigned char> value_vector;
//        marshal_uint256_t(value, value_vector);
//
//        token_char_list.insert(token_char_list.end(),
//        std::begin(tokentype),
//                               std::end(tokentype));
//        token_char_list.insert(token_char_list.end(),
//        value_vector.begin(),
//                               value_vector.end());
//    }
//
//    auto status_value = (unsigned char)state;
//    auto blockreasondata = SerializeBlockReason(blockReason);

//    auto br = balance.serializeBalanceValues();
//    auto stat = Serialize(state);
//    auto b = SerializeBlockReason(blockReason);
//
//    auto state_data = CombineData(br, stat, b);
//
//    return state_data;
//}

// SerializedStateData Deserialize(std::vector<unsigned char> state_data) {
//    auto return_data = SerializedStateData();
//
//    auto current_it = state_data.begin();
//    auto status_value = *current_it;
//
//    return_data.state = (Status)status_value;
//
//    current_it++;
//
//    auto blocktype = (BlockType)*current_it;
//    current_it++;
//
//    if (blocktype == Inbox) {
//        std::vector<unsigned char> inbox_vector(current_it, current_it + 33);
//
//        auto buff = reinterpret_cast<char*>(&inbox_vector[0]);
//        auto inbox = deserialize_int(buff);
//
//        current_it += 33;
//
//        InboxBlocked br{Inbox, inbox};
//
//        return_data.br = br;
//
//    } else if (blocktype == Send) {
//        std::vector<unsigned char> currency_vector(current_it, current_it +
//        33);
//
//        auto buff = reinterpret_cast<char*>(&currency_vector[0]);
//        auto currency = deserialize_int(buff);
//
//        current_it += 33;
//
//        std::array<unsigned char, 21> token_type;
//        std::copy(current_it, current_it + 21, token_type.begin());
//
//        current_it += 21;
//
//        SendBlocked sb{Send, currency, token_type};
//
//        return_data.br = sb;
//
//    } else if (blocktype == Not) {
//        return_data.br = NotBlocked();
//    } else if (blocktype == Halt) {
//        return_data.br = HaltBlocked();
//    } else if (blocktype == Error) {
//        return_data.br = ErrorBlocked();
//    } else if (blocktype == Breakpoint) {
//        return_data.br = BreakpointBlocked();
//    }
//
//    std::vector<unsigned char> balance(current_it, state_data.end());
//
//    std::vector<std::tuple<TokenType, uint256_t>> pairs;
//
//    while (current_it != state_data.end()) {
//        std::array<unsigned char, 21> token_type;
//
//        std::copy(current_it, current_it + 21, token_type.begin());
//        current_it += 21;
//
//        std::vector<unsigned char> value_vector(current_it, current_it + 33);
//        current_it += 33;
//
//        auto buff = reinterpret_cast<char*>(&value_vector[0]);
//        auto currency_val = deserialize_int(buff);
//
//        pairs.push_back(std::make_tuple(token_type, currency_val));
//    }
//
//    return_data.balance_data = pairs;
//
//    return return_data;
//}
