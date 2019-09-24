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

// const machine?
Checkpoint MachineCheckPoints::SaveMachine(std::string machine_state_name,
                                           Machine& machine) {
    auto checkpoint_data = machine.getCheckPointData();
    auto pool = &machine.getPool();

    auto datastack_tuple =
        Tuple(checkpoint_data.stack.GetTupleRepresentation(pool),
              checkpoint_data.auxstack.GetTupleRepresentation(pool), pool);
    auto value_tuple =
        Tuple(checkpoint_data.staticVal, checkpoint_data.registerVal, pool);
    auto messagestack_tuple = Tuple(checkpoint_data.pendingInbox.messages,
                                    checkpoint_data.inbox.messages, pool);

    auto pc_codepoint = CodePoint();
    pc_codepoint.pc = checkpoint_data.pc;
    auto codepoint_tuple = Tuple(pc_codepoint, checkpoint_data.errpc, pool);

    auto all_tuple = Tuple(datastack_tuple, value_tuple, codepoint_tuple, pool);
    auto state_data =
        SerializeData(checkpoint_data.balance, checkpoint_data.state,
                      checkpoint_data.blockReason);

    auto save_status =
        storage.SaveMachineState(machine_state_name, all_tuple, state_data);

    return Checkpoint{machine_state_name, save_status};
}

Machine MachineCheckPoints::RestoreMachine(std::string name,
                                           std::string contract_filename) {
    auto machine = Machine(contract_filename);
}
