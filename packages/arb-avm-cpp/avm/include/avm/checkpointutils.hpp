//
//  checkpointutils.hpp
//  avm
//
//  Created by Minh Truong on 9/22/19.
//

#ifndef checkpointutils_hpp
#define checkpointutils_hpp

#include <string>
#include <vector>
#include "machinestatedata.hpp"
#include "messagestack.hpp"
#include "rocksdb/db.h"

struct GetResults {
    int reference_count = 0;
    rocksdb::Status status;
    std::vector<unsigned char> storage_key;
    std::string stored_value;
};

struct StateData {
    BalanceTracker& balance;
    Status& state;
    BlockReason& blockReason;
};

struct SerializedStateData {
    std::vector<std::tuple<TokenType, uint256_t>> balance_data;
    Status state;
    BlockReason br;
};

struct CheckpointData {
    value staticVal;
    value registerVal;
    datastack stack;
    datastack auxstack;
    MessageStack pendingInbox;
    MessageStack inbox;
    uint64_t pc;
    CodePoint errpc;
    BalanceTracker balance;
    Status state;
    BlockReason blockReason;
};

class CheckpointParser {
   public:
    CheckpointData ParseData(std::string data);
};

std::vector<unsigned char> SerializeData(BalanceTracker& balance,
                                         Status& state,
                                         BlockReason& blockReason);

SerializedStateData Deserialize(std::vector<unsigned char> state_data);

#endif /* checkpointutils_hpp */
