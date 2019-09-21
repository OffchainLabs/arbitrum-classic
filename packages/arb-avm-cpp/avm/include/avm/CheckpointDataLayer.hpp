//
//  DataMapper.hpp
//  avm
//
//  Created by Minh Truong on 9/16/19.
//

#ifndef CheckpointDataLayer_hpp
#define CheckpointDataLayer_hpp

#include <stdio.h>
#include <avm/codepoint.hpp>
#include <vector>
#include "avm/datastack.hpp"
#include "avm/machine.hpp"
#include "checkpoint/checkpointutils.hpp"
#include "checkpoint/processstatus.hpp"
#include "rocksdb/db.h"
#include "rocksdb/utilities/transaction_db.h"

class CheckpointDataLayer {
   private:
    rocksdb::TransactionDB* txn_db;
    ProcessStatus ProcessValue(const value& value);
    std::string GetHashKey(const value& val);
    rocksdb::Status SaveValue(std::string val, std::string key);
    std::string GetValue(std::string key);
    GetResults ParseCountAndValue(std::string string_value);
    std::string SerializeCountAndValue(int count, std::string value);

   public:
    // static std::string ConvertMachineCode(std::vector<CodePoint> code);
    bool Intialize();
    void Close();
    rocksdb::Status SaveValue(const Tuple& val);
    rocksdb::Status SaveValueAndMapToKey(const Tuple& val,
                                         std::string hash_key);
    rocksdb::Status DeleteValue(std::string key);
    GetResults GetValueAndCount(std::string hash_key);
};

#endif /* DataMapper_hpp */
