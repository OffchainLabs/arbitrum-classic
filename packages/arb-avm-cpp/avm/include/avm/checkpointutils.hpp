//
//  checkpointutils.hpp
//  avm
//
//  Created by Minh Truong on 9/22/19.
//

#ifndef checkpointutils_hpp
#define checkpointutils_hpp

#include <string>
#include "machinestatedata.hpp"
#include "messagestack.hpp"

struct GetResults {
    int reference_count = 0;
    std::string result_value;
};

struct CheckpointData {
    value staticVal;
    value registerVal;
    datastack stack;
    datastack auxstack;
    Status state;
    uint64_t pc;
    CodePoint errpc;
    MessageStack pendingInbox;
    MessageStack inbox;
    BalanceTracker balance;
    BlockReason blockReason;
};

class CheckpointParser {
   public:
    CheckpointData ParseData(std::string data);
};

#endif /* checkpointutils_hpp */
