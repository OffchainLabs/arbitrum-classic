//
//  checkpointmaker.hpp
//  avm
//
//  Created by Minh Truong on 9/12/19.
//

#ifndef checkpoint_hpp
#define checkpoint_hpp

#include <stdio.h>
#include <vector>
#include "CheckpointDataLayer.hpp"
#include "avm/machine.hpp"

class Checkpoint {
    std::string checkPointName;
};

class MachineCheckPoints {
   private:
    CheckpointDataLayer storage;

   public:
    MachineCheckPoints();
    ~MachineCheckPoints();
    bool Cleanup();
    Checkpoint SaveMachine(std::string name, Machine machine);
    Machine RestoreMachine(std::string name);
    std::vector<std::string> GetKeys();
};

#endif /* checkpoint_hpp */
