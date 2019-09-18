//
//  checkpointmaker.cpp
//  avm
//
//  Created by Minh Truong on 9/12/19.
//

#include "checkpoint.hpp"
#include "CheckpointDataLayer.hpp"
#include "avm/machine.hpp"

MachineCheckPoints::MachineCheckPoints() {
    // storage = new CheckpointStorage   no work??
    storage.Intialize();
};

MachineCheckPoints::~MachineCheckPoints() {
    storage.Close();
}

Checkpoint MachineCheckPoints::SaveMachine(std::string name, Machine machine) {
    // should be a tuple?
    auto machine_state_info_tuple = machine.getMachineStateData();
    auto pc = std::get<0>(machine_state_info_tuple);
    auto data_stack = std::get<1>(machine_state_info_tuple);
    auto aux_stack = std::get<2>(machine_state_info_tuple);
    auto register_val = std::get<3>(machine_state_info_tuple);
    auto static_val = std::get<4>(machine_state_info_tuple);
}

// private functions
// ------------------------------------------------------------
void SaveMachineCode(){

};
