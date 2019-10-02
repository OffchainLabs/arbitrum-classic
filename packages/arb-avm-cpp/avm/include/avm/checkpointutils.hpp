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
#include "messagestack.hpp"
#include "rocksdb/db.h"

// struct StateData {
//    BalanceTracker& balance;
//    Status& state;
//    BlockReason& blockReason;
//};
//
// struct SerializedStateData {
//    std::vector<std::tuple<TokenType, uint256_t>> balance_data;
//    Status state;
//    BlockReason br;
//};
//
// struct MachineLoadData {
//    Tuple tuple_values;
//    SerializedStateData state_data;
//};

// struct CheckpointData {
//    value staticVal;
//    value registerVal;
//    Tuple stack;
//    Tuple auxstack;
//    MessageStack pendingInbox_messages;
//    MessageStack inbox_messages;
//    uint64_t pc;
//    CodePoint errpc;
//    BalanceTracker balance;
//    Status state;
//    BlockReason blockReason;
//};
//
// struct CHeckpointInfo {
//    value staticVal;
//    value registerVal;
//    Tuple stack;
//    Tuple auxstack;
//    MessageStack pendingInbox_messages;
//    MessageStack inbox_messages;
//    uint64_t pc;
//    CodePoint errpc;
//    std::vector<unsigned char> balance;
//    std::vector<unsigned char> state;
//    std::vector<unsigned char> blockReason;
//};
//
// class CheckpointParser {
//   public:
//    CheckpointData ParseData(std::string data);
//};

// std::vector<unsigned char> SerializeData(BalanceTracker& balance,
//                                         Status& state,
//                                         BlockReason& blockReason);
//
// SerializedStateData Deserialize(std::vector<unsigned char> state_data);

#endif /* checkpointutils_hpp */
