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

struct SaveResults {
    rocksdb::Status status;
    std::vector<unsigned char> storage_key;
    std::string stored_value;
};

struct DeleteResults {
    rocksdb::Status status;
    std::vector<unsigned char> storage_key;
    std::string stored_value;
};

class CheckpointStorage {
   private:
    rocksdb::TransactionDB* txn_db;
    rocksdb::Status SaveKeyValuePair(std::string value, std::string key);
    std::tuple<int, std::string> ParseCountAndValue(std::string string_value);
    std::string SerializeCountAndValue(int count, std::string value);
    rocksdb::Status DeleteValue(std::string key);

   public:
    CheckpointStorage();
    ~CheckpointStorage();

    GetResults saveValue(std::string val, std::vector<unsigned char> hash_key);
    GetResults getStoredValue(std::vector<unsigned char> hash_key);
    GetResults deleteStoredValue(std::vector<unsigned char> hash_key);
};

#endif /* checkpointstorage_hpp */
