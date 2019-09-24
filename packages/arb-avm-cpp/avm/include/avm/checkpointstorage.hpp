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
    GetResults SaveValueToDb(std::string val, std::string key);
    std::string GetValueFromDb(std::string key);
    std::tuple<int, std::string> ParseCountAndValue(std::string string_value);
    std::string SerializeCountAndValue(int count, std::string value);

   public:
    bool Intialize();
    void Close();
    GetResults SaveValue(const Tuple& val);
    GetResults SaveValueAndMapToKey(const Tuple& val, std::string hash_key);
    GetResults GetMachineState(std::string machine_name);
    rocksdb::Status DeleteValue(std::string key);
    GetResults GetValue(std::string hash_key);
    std::string GetHashKey(const value& val);
};

#endif /* checkpointstorage_hpp */
