//
//  checkpointstorage.hpp
//  avm
//
//  Created by Minh Truong on 9/23/19.
//

#ifndef checkpointstorage_hpp
#define checkpointstorage_hpp

#include <stdio.h>
#include <avm/codepoint.hpp>
#include <vector>
#include "avm/checkpointutils.hpp"
#include "avm/datastack.hpp"
#include "avm/machine.hpp"
#include "rocksdb/db.h"
#include "rocksdb/utilities/transaction_db.h"

class CheckpointStorage {
   private:
    rocksdb::TransactionDB* txn_db;
    rocksdb::Status SaveKeyValuePair(std::string value, std::string key);
    GetResults SaveValueToDb(std::string val, std::vector<unsigned char> key);
    std::string GetValueFromDb(std::string key);
    std::tuple<int, std::string> ParseCountAndValue(std::string string_value);
    std::string SerializeCountAndValue(int count, std::string value);

   public:
    bool Intialize();
    void Close();
    GetResults SaveValue(const Tuple& val);
    rocksdb::Status SaveMachineState(std::string checkpoint_name,
                                     const Tuple& tuple,
                                     std::vector<unsigned char> state_data);
    GetResults GetMachineState(std::string checkpoint_name);
    rocksdb::Status DeleteValue(std::string key);
    GetResults GetValue(std::vector<unsigned char> hash_key);
    std::vector<unsigned char> GetHashKey(const value& val);
};

#endif /* checkpointstorage_hpp */
