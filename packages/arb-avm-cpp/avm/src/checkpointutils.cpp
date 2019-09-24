//
//  checkpointutils.cpp
//  avm
//
//  Created by Minh Truong on 9/22/19.
//

#include "avm/checkpointutils.hpp"

std::vector<unsigned char> CombineData(
    std::vector<unsigned char> token_char_list,
    unsigned char status_value,
    SerializedBlockReason block_reason) {
    std::vector<unsigned char> return_value;
    return_value.push_back(status_value);
    return_value.push_back((unsigned char)block_reason.type);
    return_value.insert(return_value.end(), block_reason.data.begin(),
                        block_reason.data.end());
    return_value.insert(return_value.end(), token_char_list.begin(),
                        token_char_list.end());

    return return_value;
}

std::vector<unsigned char> SerializeData(BalanceTracker& balance,
                                         Status& state,
                                         BlockReason& blockReason) {
    auto list_of_token_pairs = balance.GetAllTokenPairs();
    std::vector<unsigned char> token_char_list;

    for (auto& pair : list_of_token_pairs) {
        auto tokentype = std::get<0>(pair);
        auto value = std::get<1>(pair);
        auto value_vector = ConvertToCharVector(value);

        token_char_list.insert(token_char_list.end(), std::begin(tokentype),
                               std::end(tokentype));
        token_char_list.insert(token_char_list.end(), value_vector.begin(),
                               value_vector.end());
    }

    auto status_value = (unsigned char)state;
    auto blockreasondata = SerializeBlockReason(blockReason);

    auto state_data =
        CombineData(token_char_list, status_value, blockreasondata);

    return state_data;
}
