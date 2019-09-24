//
//  machinestatedata.cpp
//  avm
//
//  Created by Minh Truong on 9/23/19.
//

#include "avm/machinestatedata.hpp"

struct BlockSerializer {
    SerializedBlockReason operator()(const NotBlocked& val) const {
        SerializedBlockReason x{val.type, std::vector<unsigned char>()};

        return x;
    }
    SerializedBlockReason operator()(const HaltBlocked& val) const {
        SerializedBlockReason x{val.type, std::vector<unsigned char>()};

        return x;
    }
    SerializedBlockReason operator()(const ErrorBlocked& val) const {
        SerializedBlockReason x{val.type, std::vector<unsigned char>()};

        return x;
    }
    SerializedBlockReason operator()(const BreakpointBlocked& val) const {
        SerializedBlockReason x{val.type, std::vector<unsigned char>()};

        return x;
    }
    SerializedBlockReason operator()(const InboxBlocked& val) const {
        auto inbox = ConvertToCharVector(val.inbox);

        SerializedBlockReason x{val.type, inbox};

        return x;
    }
    SerializedBlockReason operator()(const SendBlocked& val) const {
        auto data = ConvertToCharVector(val.currency);
        data.insert(data.end(), std::begin(val.tokenType),
                    std::end(val.tokenType));

        SerializedBlockReason x{val.type, data};

        return x;
    }
};

SerializedBlockReason SerializeBlockReason(const BlockReason& val) {
    return nonstd::visit(BlockSerializer{}, val);
}
