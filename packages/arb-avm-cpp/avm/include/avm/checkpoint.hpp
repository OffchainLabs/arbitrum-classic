//
//  checkpoint.hpp
//  avm
//
//  Created by Minh Truong on 9/22/19.
//

#ifndef checkpoint_hpp
#define checkpoint_hpp

#include <stdio.h>
#include <vector>
#include "avm/machine.hpp"
#include "checkpointstorage.hpp"

struct Checkpoint {
    std::string checkPointName;
    rocksdb::Status status;
};

class MachineCheckPoints {
   private:
    CheckpointStorage storage;

   public:
    MachineCheckPoints();
    ~MachineCheckPoints();
    bool Cleanup();
    Checkpoint SaveMachine(std::string name, Machine& machine);
    Machine RestoreMachine(std::string name, std::string contract_filename);
    std::vector<std::string> GetKeys();
};

#endif /* checkpoint_hpp */
