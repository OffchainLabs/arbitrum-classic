//
//  MachineStateSaver.hpp
//  avm
//
//  Created by Minh Truong on 9/30/19.
//

#ifndef machinestatesaver_hpp
#define machinestatesaver_hpp

#include "avm/checkpointstorage.hpp"
#include "avm/value.hpp"

struct MachineStateStorageData {
    GetResults static_val_results;
    GetResults register_val_results;
    GetResults datastack_results;
    GetResults auxstack_results;
    GetResults inbox_results;
    GetResults pending_results;
    GetResults pc_results;
    unsigned char status_str;
    std::vector<unsigned char> blockreason_str;
    std::vector<unsigned char> balancetracker_str;
};

struct MachineStateFetchedData {
    value static_val_results;
    value register_val_results;
    Tuple datastack_results;
    Tuple auxstack_results;
    Tuple inbox_results;
    Tuple pending_results;
    CodePoint pc_results;
    unsigned char status_str;
    std::vector<unsigned char> blockreason_str;
    std::vector<unsigned char> balancetracker_str;
};

class MachineStateSaver {
   private:
    CheckpointStorage* checkpoint_storage;
    TuplePool* pool;
    std::vector<std::vector<unsigned char>> breakIntoValues(
        std::vector<unsigned char> data_vecgtor);
    std::vector<unsigned char> serializeState(
        MachineStateStorageData state_data);
    MachineStateFetchedData deserializeState(
        std::vector<unsigned char> stored_state);
    GetResults SaveStringValue(const std::string value,
                               const std::vector<unsigned char> key);
    GetResults GetStringValue(const std::vector<unsigned char> key);
    CodePoint getCodePoint(std::vector<unsigned char> hash_key);
    uint256_t getInt256(std::vector<unsigned char> hash_key);

   public:
    void setStorage(CheckpointStorage* storage);
    GetResults SaveTuple(const Tuple& val);
    GetResults SaveValue(const value& val);
    value getValue(std::vector<unsigned char> hash_key);
    Tuple getTuple(std::vector<unsigned char> hash_key);

    GetResults SaveMachineState(MachineStateStorageData state_data,
                                std::string checkpoint_name);

    MachineStateFetchedData GetMachineStateData(std::string checkpoint_name);
};

#endif /* machinestatesaver_hpp */
