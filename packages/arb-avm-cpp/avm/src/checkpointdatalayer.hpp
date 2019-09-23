//
//  checkpointdatalayer.hpp
//  avm
//
//  Created by Minh Truong on 9/22/19.
//

#ifndef checkpointdatalayer_hpp
#define checkpointdatalayer_hpp

#include <stdio.h>
#include <avm/codepoint.hpp>
#include <vector>
#include "avm/datastack.hpp"
#include "avm/machine.hpp"
#include "checkpointutils.hpp"
#include "rocksdb/db.h"
#include "rocksdb/utilities/transaction_db.h"
#include "serializedvalue.hpp"

class CheckpointDataLayer {
   private:
    rocksdb::TransactionDB* txn_db;
    rocksdb::Status SaveValue(std::string val, std::string key);
    std::string GetValue(std::string key);
    GetResults ParseCountAndValue(std::string string_value);
    std::string SerializeCountAndValue(int count, std::string value);

   public:
    bool Intialize();
    void Close();
    rocksdb::Status SaveValue(const Tuple& val);
    rocksdb::Status SaveValueAndMapToKey(const Tuple& val,
                                         std::string hash_key);
    rocksdb::Status GetMachineState(std::string machine_name);
    rocksdb::Status DeleteValue(std::string key);
    GetResults GetValueAndCount(std::string hash_key);
    std::string GetHashKey(const value& val);
};

#endif /* checkpointdatalayer_hpp */
