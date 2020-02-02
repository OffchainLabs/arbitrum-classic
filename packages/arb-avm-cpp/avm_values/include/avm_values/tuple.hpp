/*
 * Copyright 2019, Offchain Labs, Inc.
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

#include <avm_values/codepoint.hpp>
#include <avm_values/exceptions.hpp>
#include <avm_values/pool.hpp>
#include <avm_values/value.hpp>

#include <memory>

auto zeroHash() -> uint256_t;

class Tuple {
   private:
    TuplePool* tuplePool{nullptr};
    std::shared_ptr<RawTuple> tpl;

    friend auto hash(const Tuple&) -> uint256_t;

   public:
    Tuple() = default;
    auto calculateHash() const -> uint256_t;

    Tuple(TuplePool* pool, size_t size) {
        tuplePool = pool;
        if (size > 0) {
            tpl = pool->getResource(size);
            for (size_t i = 0; i < size; i++) {
                tpl->data.push_back(Tuple{});
            }
            tpl->deferredHashing = true;
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

    //    ~Tuple() {
    //        if (tpl.use_count() == 1) {
    //            tuplePool->returnResource(std::move(tpl));
    //        }
    //    }

    auto tuple_size() const -> uint64_t {
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
        // turned off because this optimization may not be thread safe
        //        if (tpl.use_count() > 1) {
        // make new copy tuple
        std::shared_ptr<RawTuple> tmp = tuplePool->getResource(tuple_size());

        std::copy(tpl->data.begin(), tpl->data.end(),
                  std::back_inserter(tmp->data));
        tpl = tmp;
        //        }
        tpl->data[pos] = std::move(newval);
        tpl->deferredHashing = true;
    }

    auto get_element(uint64_t pos) const -> value {
        if (pos >= tuple_size()) {
            throw bad_tuple_index{};
        }
        return tpl->data[pos];
    }

    void marshal(std::vector<unsigned char>& buf) const;
    auto clone_shallow() -> value;
};

inline auto hash(const Tuple& tup) -> uint256_t {
    if (tup.tpl) {
        if (tup.tpl->deferredHashing) {
            tup.tpl->cachedHash = tup.calculateHash();
            tup.tpl->deferredHashing = false;
        }
        return tup.tpl->cachedHash;
    } else {
        static uint256_t zeroHashVal = zeroHash();
        return zeroHashVal;
    }
}

inline auto operator==(const Tuple& val1, const Tuple& val2) -> bool {
    if (val1.tuple_size() != val2.tuple_size()) {
        return false;
    }
    return hash(val1) == hash(val2);
}

inline auto operator!=(const Tuple& val1, const Tuple& val2) -> bool {
    if (val1.tuple_size() == val2.tuple_size()) {
        return false;
    }
    return hash(val1) != hash(val2);
}

auto operator<<(std::ostream& os, const Tuple& val) -> std::ostream&;

#endif /* tuple_hpp */
