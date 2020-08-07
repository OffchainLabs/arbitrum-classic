/*
 * Copyright 2019-2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#ifndef tuple_hpp
#define tuple_hpp

#include <avm_values/codepointstub.hpp>
#include <avm_values/exceptions.hpp>
#include <avm_values/pool.hpp>
#include <avm_values/value.hpp>

#include <memory>

HashPreImage zeroPreimage();

class Tuple {
   private:
    TuplePool* tuplePool;
    std::shared_ptr<RawTuple> tpl;

    HashPreImage calculateHashPreImage() const;

   public:
    Tuple() : tuplePool(nullptr) {}

    uint256_t getSize() const { return getHashPreImage().getSize(); }

    Tuple(TuplePool* pool, size_t size) {
        if (size > 0) {
            tuplePool = pool;
            tpl = pool->getResource(size);
            for (size_t i = 0; i < size; i++) {
                tpl->data.push_back(Tuple{});
            }
        }
    }

    Tuple(value val, TuplePool* pool);

    Tuple(value val1, value val2, TuplePool* pool);

    Tuple(value val1, value val2, value val3, TuplePool* pool);

    Tuple(value val1, value val2, value val3, value val4, TuplePool* pool);

    Tuple(value val1,
          value val2,
          value val3,
          value val4,
          value val5,
          TuplePool* pool);

    Tuple(value val1,
          value val2,
          value val3,
          value val4,
          value val5,
          value val6,
          TuplePool* pool);

    Tuple(value val1,
          value val2,
          value val3,
          value val4,
          value val5,
          value val6,
          value val7,
          TuplePool* pool);

    Tuple(value val1,
          value val2,
          value val3,
          value val4,
          value val5,
          value val6,
          value val7,
          value val8,
          TuplePool* pool);

    Tuple(std::vector<value> values, TuplePool* pool);

    uint64_t tuple_size() const {
        if (tpl) {
            return tpl->data.size();
        } else {
            return 0;
        }
    }

    void set_element(uint64_t pos, value newval) {
        if (pos >= tuple_size()) {
            throw bad_tuple_index{};
        }
        if (tpl.use_count() > 1) {
            std::shared_ptr<RawTuple> tmp =
                tuplePool->getResource(tuple_size());
            std::copy(tpl->data.begin(), tpl->data.end(),
                      std::back_inserter(tmp->data));
            tpl = tmp;
        }
        tpl->data[pos] = std::move(newval);
        tpl->deferredHashing = true;
    }

    value get_element(uint64_t pos) const {
        if (pos >= tuple_size()) {
            throw bad_tuple_index{};
        }
        return tpl->data[pos];
    }

    const value& get_element_unsafe(uint64_t pos) const {
        return tpl->data[pos];
    }

    HashPreImage getHashPreImage() const {
        if (!tpl) {
            return zeroPreimage();
        }
        if (tpl->deferredHashing) {
            tpl->cachedPreImage = calculateHashPreImage();
            tpl->deferredHashing = false;
        }
        return tpl->cachedPreImage;
    }

    void marshal(std::vector<unsigned char>& buf) const;
};

inline uint256_t hash(const Tuple& tup) {
    return hash(tup.getHashPreImage());
}

inline bool operator==(const Tuple& val1, const Tuple& val2) {
    if (val1.tuple_size() != val2.tuple_size())
        return false;
    return hash(val1) == hash(val2);
}

inline bool operator!=(const Tuple& val1, const Tuple& val2) {
    if (val1.tuple_size() == val2.tuple_size())
        return false;
    return hash(val1) != hash(val2);
}

std::ostream& operator<<(std::ostream& os, const Tuple& val);

#endif /* tuple_hpp */
