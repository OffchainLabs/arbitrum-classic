//
//  checkpointstorage.hpp
//  avm
//
//  Created by Minh Truong on 9/23/19.
//

#ifndef checkpointstorage_hpp
#define checkpointstorage_hpp

#include <vector>
#include "rocksdb/db.h"
#include "rocksdb/utilities/transaction_db.h"

struct GetResults {
    int reference_count = 0;
    rocksdb::Status status;
    std::vector<unsigned char> storage_key;
    std::string stored_value;
};

class CheckpointStorage {
   private:
    rocksdb::TransactionDB* txn_db;
    rocksdb::Status SaveKeyValuePair(std::string value, std::string key);

   public:
    bool Intialize();
    void Close();

    GetResults SaveValueToDb(std::string val, std::vector<unsigned char> key);
    std::string GetValueFromDb(std::string key);
    std::tuple<int, std::string> ParseCountAndValue(std::string string_value);
    std::string SerializeCountAndValue(int count, std::string value);

    rocksdb::Status DeleteValue(std::string key);
    GetResults getStoredValue(std::vector<unsigned char> hash_key);
    // std::vector<unsigned char> GetHashKey(const value& val);
    // Tuple GetTuple(std::vector<unsigned char> hash_key);
};

#endif /* checkpointstorage_hpp */
