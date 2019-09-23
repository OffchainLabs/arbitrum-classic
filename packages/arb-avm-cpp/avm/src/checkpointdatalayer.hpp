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
    //    ProcessStatus ProcessValue(CheckpointDataLayer& cp, const value&
    //    value);

   public:
    // static std::string ConvertMachineCode(std::vector<CodePoint> code);
    bool Intialize();
    void Close();
    rocksdb::Status SaveValue(const Tuple& val);
    rocksdb::Status SaveValueAndMapToKey(const Tuple& val,
                                         std::string hash_key);
    rocksdb::Status DeleteValue(std::string key);
    GetResults GetValueAndCount(std::string hash_key);
    std::string GetHashKey(const value& val);
};

// struct ValueProcessor{
//    CheckpointDataLayer cp;
//    ProcessStatus operator()(const Tuple& value);
//    ProcessStatus operator()(const uint256_t& value);
//    ProcessStatus operator()(const CodePoint& value);
//};

#endif /* checkpointdatalayer_hpp */
