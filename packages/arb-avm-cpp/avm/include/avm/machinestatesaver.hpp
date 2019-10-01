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

struct MachineStateData {
    GetResults static_val_results;
    GetResults register_val_results;
    GetResults datastack_results;
    GetResults auxstack_results;
    GetResults inbox_results;
    GetResults pending_results;
    GetResults pc_results;
    std::vector<unsigned char> status_str;
    std::vector<unsigned char> blockreason_str;
    std::vector<unsigned char> balancetracker_str;
};

class MachineStateSaver {
   private:
    TuplePool* pool;
    std::vector<std::vector<unsigned char>> breakIntoValues(
        std::vector<unsigned char> data_vecgtor);
    std::vector<unsigned char> serializeState(MachineStateData state_data);
    MachineStateData deserializeState(std::vector<unsigned char> stored_state);
    GetResults SaveStringValue(const std::string value,
                               const std::vector<unsigned char> key);

   public:
    CheckpointStorage storage;

    GetResults SaveTuple(const Tuple& val);
    GetResults SaveValue(const value& val);
    value getValue(std::vector<unsigned char> hash_key);
    Tuple getTuple(std::vector<unsigned char> hash_key);

    GetResults SaveMachineState(MachineStateData state_data,
                                std::string checkpoint_name);
    //    rocksdb::Status SaveMachineState(std::string checkpoint_name, const
    //    Tuple& tuple, std::vector<unsigned char> state_data);
    // MachineLoadData GetMachineState(std::string checkpoint_name);
};

#endif /* machinestatesaver_hpp */
