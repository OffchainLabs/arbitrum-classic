//
//  MachineState.hpp
//  avm
//
//  Created by Minh Truong on 9/19/19.
//

#ifndef machinestate_hpp
#define machinestate_hpp

#include <stdio.h>
#include <avm/datastack.hpp>
#include <avm/tokenTracker.hpp>
#include <avm/value.hpp>
#include <memory>
#include <vector>
#include "machinestatedata.hpp"
#include "messagestack.hpp"
//#include "avm/checkpointstorage.hpp"
#include "avm/machinestatesaver.hpp"

struct AssertionContext {
    uint32_t numSteps;
    TimeBounds timeBounds;
    std::vector<Message> outMessage;
    std::vector<value> logs;

    explicit AssertionContext(const TimeBounds& tb)
        : numSteps{0}, timeBounds(tb) {}
};

struct MachineState {
    MachineStateSaver msSaver;

    std::shared_ptr<TuplePool> pool;
    std::vector<CodePoint> code;
    value staticVal;
    value registerVal;
    Datastack stack;
    Datastack auxstack;
    Status state = Status::Extensive;
    uint64_t pc = 0;
    CodePoint errpc;
    MessageStack pendingInbox;
    AssertionContext context;
    MessageStack inbox;
    BalanceTracker balance;
    BlockReason blockReason;

    MachineState();

    bool deserialize(char* data);

    void readInbox(char* newInbox);
    std::vector<unsigned char> marshalForProof();
    uint64_t pendingMessageCount() const;
    void sendOnchainMessage(const Message& msg);
    void deliverOnchainMessages();
    void sendOffchainMessages(const std::vector<Message>& messages);
    BlockReason runOp(OpCode opcode);
    uint256_t hash() const;

    void setInbox(MessageStack ms);
    void setPendingInbox(MessageStack ms);

    //    int SetMachineState(CheckpointData data);
    //    int SaveMachine(std::string checkpoint_name);
};

#endif /* machinestate_hpp */
