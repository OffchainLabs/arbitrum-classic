//
//  blockstore.hpp
//  avm
//
//  Created by Harry Kalodner on 5/17/20.
//

#ifndef blockstore_hpp
#define blockstore_hpp

#include <avm_values/bigint.hpp>

#include <memory>
#include <vector>

struct DataResults;

namespace rocksdb {
class TransactionDB;
class Status;
struct Slice;
class ColumnFamilyHandle;
}  // namespace rocksdb

class BlockStore {
   private:
    std::shared_ptr<rocksdb::TransactionDB> txn_db;
    std::shared_ptr<rocksdb::ColumnFamilyHandle> blocks_column;

   public:
    BlockStore(std::shared_ptr<rocksdb::TransactionDB> txn_db_,
               std::shared_ptr<rocksdb::ColumnFamilyHandle> blocks_column_)
        : txn_db(std::move(txn_db_)),
          blocks_column(std::move(blocks_column_)) {}
    rocksdb::Status putBlock(const uint256_t& height,
                             const uint256_t& hash,
                             const std::vector<char>& value);
    rocksdb::Status deleteBlock(const uint256_t& height, const uint256_t& hash);
    DataResults getBlock(const uint256_t& height, const uint256_t& hash) const;

    std::vector<uint256_t> blockHashesAtHeight(const uint256_t& height) const;
    bool isEmpty() const;
    uint256_t maxHeight() const;
    uint256_t minHeight() const;
};

#endif /* blockstore_hpp */
