//
//  machinestatedata.cpp
//  avm
//
//  Created by Minh Truong on 9/23/19.
//

#include "avm/machinestatedata.hpp"

struct BlockSerializer {
    SerializedBlockReason operator()(const NotBlocked& val) const {
        return SerializedBlockReason{val.type, std::vector<unsigned char>()};
    }
    SerializedBlockReason operator()(const HaltBlocked& val) const {
        return SerializedBlockReason{val.type, std::vector<unsigned char>()};
    }
    SerializedBlockReason operator()(const ErrorBlocked& val) const {
        return SerializedBlockReason{val.type, std::vector<unsigned char>()};
    }
    SerializedBlockReason operator()(const BreakpointBlocked& val) const {
        return SerializedBlockReason{val.type, std::vector<unsigned char>()};
    }

    SerializedBlockReason operator()(const InboxBlocked& val) const {
        std::vector<unsigned char> inbox_char_vector;
        marshal_uint256_t(val.inbox, inbox_char_vector);

        return SerializedBlockReason{val.type, inbox_char_vector};
    }
    SerializedBlockReason operator()(const SendBlocked& val) const {
        std::vector<unsigned char> data_vector;
        marshal_uint256_t(val.currency, data_vector);

        data_vector.insert(data_vector.end(), std::begin(val.tokenType),
                           std::end(val.tokenType));

        return SerializedBlockReason{val.type, data_vector};
    }
};

SerializedBlockReason SerializeBlockReason(const BlockReason& val) {
    return nonstd::visit(BlockSerializer{}, val);
}
