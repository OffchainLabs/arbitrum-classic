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

MachineCheckPoints::MachineCheckPoints() {
    // storage = new CheckpointStorage   no work??
    storage.Intialize();
};

MachineCheckPoints::~MachineCheckPoints() {
    storage.Close();
}

Checkpoint MachineCheckPoints::SaveMachine(std::string machine_state_name,
                                           Machine machine) {
    auto checkpoint_data = machine.getCheckPointData();
    auto pool = machine.getPool();
    //    auto status = storage.SaveValueAndMapToKey(checkpoint_tuple.value,
    //                                               machine_state_name);
    //
    //    if (status.ok()) {
    //        return Checkpoint();
    //    } else {
    //        // retry
    //        return Checkpoint();
    //    }
}

Machine MachineCheckPoints::RestoreMachine(std::string name,
                                           std::string contract_filename) {
    auto machine = Machine(contract_filename);
}
