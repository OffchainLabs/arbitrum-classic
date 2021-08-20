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
#include <utility>

#include <memory>

HashPreImage zeroPreimage();
struct BasicValChecker;
struct ValueBeingParsed;

const static std::vector<value> empty_value_vector;

class Tuple {
   private:
    std::shared_ptr<RawTuple> tpl;

    explicit Tuple(std::shared_ptr<RawTuple> tpl) : tpl(std::move(tpl)){};

    void calculateHashPreImage() const;

    void unsafe_set_element(uint64_t pos, value&& newval) {
        if (pos >= tuple_size()) {
            throw bad_tuple_index{};
        }
        tpl->data[pos] = std::move(newval);
        tpl->deferredHashing = true;
    }

    friend BasicValChecker;
    friend RawTuple;
    friend ValueBeingParsed;

   public:
    Tuple() : tpl(nullptr) {}

    [[nodiscard]] uint256_t getSize() const {
        return getHashPreImage().getSize();
    }

    static Tuple createSizedTuple(const size_t size);

    static Tuple createTuple(std::vector<value> values);

    static Tuple createTuple(value val);

    Tuple(value val1, value val2);

    Tuple(value val1, value val2, value val3);

    Tuple(value val1, value val2, value val3, value val4);

    Tuple(value val1, value val2, value val3, value val4, value val5);

    Tuple(value val1,
          value val2,
          value val3,
          value val4,
          value val5,
          value val6);

    Tuple(value val1,
          value val2,
          value val3,
          value val4,
          value val5,
          value val6,
          value val7);

    Tuple(value val1,
          value val2,
          value val3,
          value val4,
          value val5,
          value val6,
          value val7,
          value val8);

    [[nodiscard]] uint64_t tuple_size() const {
        if (tpl) {
            return tpl->data.size();
        } else {
            return 0;
        }
    }

    void set_element(const uint64_t pos, value newval) {
        if (pos >= tuple_size()) {
            throw bad_tuple_index{};
        }
        std::shared_ptr<RawTuple> tmp =
            TuplePool::get_impl().getResource(tuple_size());
        for (uint64_t i = 0; i < tuple_size(); i++) {
            if (i == pos) {
                tmp->data.emplace_back(std::move(newval));
            } else {
                tmp->data.emplace_back(tpl->data[i]);
            }
        }
        tpl = std::move(tmp);
    }

    [[nodiscard]] value get_element(const uint64_t pos) const {
        if (pos >= tuple_size()) {
            throw bad_tuple_index{};
        }
        return tpl->data[pos];
    }

    [[nodiscard]] std::vector<value>::const_iterator begin() const {
        if (tpl == nullptr) {
            return empty_value_vector.begin();
        }
        return tpl->data.begin();
    }

    [[nodiscard]] std::vector<value>::const_iterator end() const {
        if (tpl == nullptr) {
            return empty_value_vector.end();
        }
        return tpl->data.end();
    }

    [[nodiscard]] std::reverse_iterator<std::vector<value>::const_iterator>
    rbegin() const {
        return std::reverse_iterator(end());
    }

    [[nodiscard]] std::reverse_iterator<std::vector<value>::const_iterator>
    rend() const {
        return std::reverse_iterator(begin());
    }

    [[nodiscard]] const value& get_element_unsafe(const uint64_t pos) const {
        return tpl->data[pos];
    }

    [[nodiscard]] value& get_element_mutable_unsafe(const uint64_t pos) const {
        tpl->deferredHashing = true;
        return tpl->data[pos];
    }

    [[nodiscard]] HashPreImage getHashPreImage() const {
        if (!tpl) {
            return zeroPreimage();
        }
        if (tpl->deferredHashing) {
            calculateHashPreImage();
        }
        return tpl->cachedPreImage;
    }
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
