//
//  checkpointmaker.cpp
//  avm
//
//  Created by Minh Truong on 9/12/19.
//

#include "checkpoint.hpp"
#include <avm/tuple.hpp>
#include "CheckpointDataLayer.hpp"
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
    // should contain more data? like errpc?
    auto checkpoint_tuple = machine.getCheckPointTuple();

    auto status =
        storage.SaveValueAndMapToKey(checkpoint_tuple, machine_state_name);

    if (status.ok()) {
        return Checkpoint();
    } else {
        // retry
        return Checkpoint();
    }
}

Machine MachineCheckPoints::RestoreMachine(std::string name,
                                           std::string contract_filename) {
    Machine machine;
    machine
}

// private functions
// ------------------------------------------------------------
void SaveMachineCode(){

};
