//
//  checkpoint.cpp
//  avm
//
//  Created by Minh Truong on 9/22/19.
//

#include "avm/checkpoint.hpp"
#include <avm/tuple.hpp>
#include "avm/checkpointstorage.hpp"
#include "avm/machine.hpp"
#include "avm/machinestatedata.hpp"

MachineCheckPoints::MachineCheckPoints() {
    // storage = new CheckpointStorage   no work??
    storage.Intialize();
};

MachineCheckPoints::~MachineCheckPoints() {
    storage.Close();
}

// Tuple makeCheckpointTuple(CheckpointData checkpoint_data, TuplePool* pool){
//
//    auto pc_codepoint = CodePoint();
//    pc_codepoint.pc = checkpoint_data.pc;
//
//    return Tuple(checkpoint_data.staticVal,
//                   checkpoint_data.registerVal,
//                   checkpoint_data.stack,
//                   checkpoint_data.auxstack,
//                   checkpoint_data.pendingInbox_messages,
//                   checkpoint_data.inbox_messages,
//                   pc_codepoint,
//                   checkpoint_data.errpc,
//                   pool);
//}

// CheckpointData extractMachineStateData(MachineLoadData machine_data){
//
//    auto tuple = machine_data.tuple_values;
//    auto state_data = machine_data.state_data;
//
//    BalanceTracker balancetracker;
//
//    for(auto& pair: state_data.balance_data){
//        pair.
//        balancetracker.add(pair, <#const uint256_t &amount#>)
//    }
//
//    if(tuple.tuple_size() != 8){
//        //error
//    }

//    return CheckpointData{
//        tuple.get_element(0),
//        tuple.get_element(1),
//        nonstd::get<Tuple>(tuple.get_element(2)),
//        nonstd::get<Tuple>(tuple.get_element(3)),
//        nonstd::get<Tuple>(tuple.get_element(4)),
//        nonstd::get<Tuple>(tuple.get_element(5)),
//        nonstd::get<CodePoint>(tuple.get_element(6)).pc,
//        nonstd::get<CodePoint>(tuple.get_element(7)),
//
//    };
//}

// const machine?
// Checkpoint MachineCheckPoints::SaveMachine(std::string machine_state_name,
// Machine& machine){
//
//    auto checkpoint_data = machine.getCheckPointData();
//    auto all_tuple = makeCheckpointTuple(checkpoint_data, &machine.getPool());
//    auto state_data = SerializeData(checkpoint_data.balance,
//    checkpoint_data.state, checkpoint_data.blockReason); auto save_status =
//    storage.SaveMachineState(machine_state_name, all_tuple, state_data);
//
//    return Checkpoint{machine_state_name, save_status};
//}
//
// Machine MachineCheckPoints::RestoreMachine(std::string name, std::string
// contract_filename){
//
//    auto machine = Machine(contract_filename);
//    auto machine_data = storage.GetMachineState(name);
//
//}
