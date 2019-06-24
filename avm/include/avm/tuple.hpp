//
//  tuple.hpp
//  AVMtest
//
//  Created by Harry Kalodner on 6/19/19.
//

#ifndef tuple_hpp
#define tuple_hpp

#include "pool.hpp"
#include "value.hpp"
#include "codepoint.hpp"

#include <boost/smart_ptr/local_shared_ptr.hpp>

class Tuple {
   private:
    TuplePool* tuplePool;
    boost::local_shared_ptr<std::vector<value>> tpl;
    mutable uint256_t cachedHash;
    
    friend uint256_t hash(const Tuple &);
    
    uint256_t calculateHash() const;

   public:
    Tuple() {}
    Tuple(int size_, TuplePool* pool)
        : tuplePool(pool), tpl(pool->getResource(size_)) {}
    
    ~Tuple() {
        if (tpl.local_use_count() == 1) {
            tuplePool->returnResource(std::move(tpl));
        }
    }

    int tuple_size() const {
        if (tpl) {
            return tpl->size();
        } else {
            return 0;
        }
    }

    void set_element(int pos, value newval) {
        if (pos >= tuple_size()) {
            throw bad_tuple_index{};
        }

        if (tpl.local_use_count() > 1) {
            // make new copy tuple
            boost::local_shared_ptr<std::vector<value>> tmp =
                tuplePool->getResource(tuple_size());
            std::copy(tpl->begin(), tpl->end(), tmp->begin());
            tpl = tmp;
        }
        (*tpl)[pos] = std::move(newval);
    }

    value get_element(int pos) const {
        if (pos >= tuple_size()) {
            throw bad_tuple_index{};
        }
        return (*tpl)[pos];
    }
};

inline uint256_t hash(const Tuple &tup) {
    if (!tup.cachedHash) {
        tup.cachedHash = tup.calculateHash();
    }
    return tup.cachedHash;
}

inline bool operator==(const Tuple& val1, const Tuple& val2) {
    if (val1.tuple_size() != val2.tuple_size())
        return false;
    for (int i = 0; i < val1.tuple_size(); i++) {
        if (!(val1.get_element(i) == val2.get_element(i)))
            return false;
    }
    return true;
}

#endif /* tuple_hpp */
