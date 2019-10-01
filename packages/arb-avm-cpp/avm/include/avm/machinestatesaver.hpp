//
//  MachineStateSaver.hpp
//  avm
//
//  Created by Minh Truong on 9/30/19.
//

#ifndef machinestatesaver_hpp
#define machinestatesaver_hpp

#include "avm/checkpointstorage.hpp"
#include "avm/machinestatedata.hpp"
#include "avm/value.hpp"

class MachineStateSaver {
   private:
    TuplePool* pool;
    std::vector<std::vector<unsigned char>> breakIntoValues(
        std::vector<unsigned char> data_vecgtor);

   public:
    CheckpointStorage storage;

    GetResults SaveStringValue(const std::string value,
                               const std::vector<unsigned char> key);
    GetResults SaveTuple(const Tuple& val);
    GetResults SaveValue(const value& val);
    value getValue(std::vector<unsigned char> hash_key);
    Tuple getTuple(std::vector<unsigned char> hash_key);

    //    rocksdb::Status SaveMachineState(std::string checkpoint_name, const
    //    Tuple& tuple, std::vector<unsigned char> state_data);
    // MachineLoadData GetMachineState(std::string checkpoint_name);
};

#endif /* machinestatesaver_hpp */
