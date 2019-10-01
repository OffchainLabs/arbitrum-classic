//
//  machinestatedata.cpp
//  avm
//
//  Created by Minh Truong on 9/23/19.
//

#include "avm/machinestatedata.hpp"

// struct BlockSerializer {
//    std::vector<unsigned char> operator()(const NotBlocked& val) const {
//        std::vector<unsigned char> return_value;
//        return_value.push_back((unsigned char)val.type);
//        return return_value;
//    }
//    std::vector<unsigned char> operator()(const HaltBlocked& val) const {
//        std::vector<unsigned char> return_value;
//        return_value.push_back((unsigned char)val.type);
//
//        return return_value;
//    }
//    std::vector<unsigned char> operator()(const ErrorBlocked& val) const {
//        std::vector<unsigned char> return_value;
//        return_value.push_back((unsigned char)val.type);
//
//        return return_value;
//    }
//    std::vector<unsigned char> operator()(const BreakpointBlocked& val) const
//    {
//        std::vector<unsigned char> return_value;
//        return_value.push_back((unsigned char)val.type);
//
//        return return_value;
//    }
//
//    std::vector<unsigned char> operator()(const InboxBlocked& val) const {
//        std::vector<unsigned char> return_value;
//        return_value.push_back((unsigned char)val.type);
//
//        std::vector<unsigned char> inbox_char_vector;
//        marshal_uint256_t(val.inbox, inbox_char_vector);
//
//        return_value.insert(return_value.end(), inbox_char_vector.begin(),
//                            inbox_char_vector.end());
//
//        return return_value;
//    }
//    std::vector<unsigned char> operator()(const SendBlocked& val) const {
//        std::vector<unsigned char> return_value;
//        return_value.push_back((unsigned char)val.type);
//
//        std::vector<unsigned char> data_vector;
//        marshal_uint256_t(val.currency, data_vector);
//
//        return_value.insert(return_value.end(), data_vector.begin(),
//                            data_vector.end());
//
//        return_value.insert(return_value.end(), std::begin(val.tokenType),
//                            std::end(val.tokenType));
//
//        return return_value;
//    }
//};
//
// std::vector<unsigned char> SerializeBlockReason(const BlockReason& val) {
//    return nonstd::visit(BlockSerializer{}, val);
//}
//
// struct SerializedBlockReason {
//    BlockType type;
//    std::vector<unsigned char> data;
//};
//
// BlockReason deserializeBlockReason(std::vector<unsigned char> data) {
//    BlockReason blockreason;
//
//    auto current_it = data.begin();
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
//        blockreason = br;
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
//        blockreason = sb;
//
//    } else if (blocktype == Not) {
//        blockreason = NotBlocked();
//    } else if (blocktype == Halt) {
//        blockreason = HaltBlocked();
//    } else if (blocktype == Error) {
//        blockreason = ErrorBlocked();
//    } else if (blocktype == Breakpoint) {
//        blockreason = BreakpointBlocked();
//    }
//
//    return blockreason;
//}
//
// std::vector<unsigned char> Serialize(Status status) {
//    return std::vector<unsigned char>((unsigned char)status);
//}
