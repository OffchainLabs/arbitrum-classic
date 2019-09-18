//
//  CheckpointStorage.hpp
//  avm
//
//  Created by Minh Truong on 9/12/19.
//

#ifndef CheckpointStorage_hpp
#define CheckpointStorage_hpp

#include <stdio.h>
#include <avm/codepoint.hpp>
#include <vector>
#include "avm/pool.hpp"
#include "rocksdb/db.h"
#include "rocksdb/utilities/transaction_db.h"

class CheckpointStorage {
   private:
    rocksdb::TransactionDB* txn_db;
    //    rocksdb::Status SaveValue(std::string val, std::string key);
    //    std::string GetValue(rocksdb::TransactionDB* txn_db, std::string key);
    //    rocksdb::Status DeleteValue(rocksdb::TransactionDB* txn_db,
    //    std::string key);
    int GetReferenceCount(std::vector<unsigned char>& hash_key);
    rocksdb::Status UpdateReferenceCount(std::vector<unsigned char>& hash_key,
                                         int count);
    rocksdb::Status RemoveReference(std::string key);
    std::string MakeReferenceKey(std::vector<unsigned char>& key);
    std::string MakeKey(std::vector<unsigned char>& key);
    rocksdb::Status SaveValAndUpdateRef(std::vector<unsigned char>& hash_key,
                                        std::string value);
    rocksdb::Status DeleteValAndUpdateRef(std::vector<unsigned char>& hash_key,
                                          rocksdb::TransactionDB* txn_db);
    rocksdb::Status PutInTransaction(rocksdb::Transaction* transaction,
                                     std::string key,
                                     std::string value);

   public:
    rocksdb::Status StoreValue(Tuple& val);
    rocksdb::Status StoreValue(uint256_t& val);
    rocksdb::Status StoreValue(CodePoint& val);
    //    rocksdb::Status SaveMachineCode(std::vector<CodePoint> code);
    //    rocksdb::Status SaveValue(value& val, TuplePool& pool);
    //    rocksdb::Status SaveValue(Tuple& val, TuplePool& pool);
    //    rocksdb::Status SaveValue(uint256_t& val);
    //    rocksdb::Status SaveValue(CodePoint& val);
};

#endif /* CheckpointStorage_hpp */
