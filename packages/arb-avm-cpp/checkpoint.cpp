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
    auto checkpoint_info = machine.getCheckPointInfo();
    auto hash_key = machine.hash();
}

// private functions
// ------------------------------------------------------------
void SaveMachineCode(){

};
