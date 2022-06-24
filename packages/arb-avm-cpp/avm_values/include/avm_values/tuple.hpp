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

#include <avm_values/buffer.hpp>
#include <avm_values/codepointstub.hpp>
#include <avm_values/exceptions.hpp>
#include <avm_values/pool.hpp>
#include <avm_values/unloadedvalue.hpp>
#include <avm_values/valuetype.hpp>
#include <utility>

#include <memory>

HashPreImage zeroPreimage();
struct BasicValChecker;
struct ValueBeingParsed;

class Value;

const static std::vector<Value> empty_value_vector;

class Tuple {
   private:
    std::shared_ptr<RawTuple> tpl;

    explicit Tuple(std::shared_ptr<RawTuple> tpl) : tpl(std::move(tpl)){};

    void calculateHashPreImage() const;

    void unsafe_set_element(uint64_t pos, Value&& newval);

    friend BasicValChecker;
    friend RawTuple;
    friend ValueBeingParsed;

   public:
    Tuple() : tpl(nullptr) {}

    [[nodiscard]] uint256_t getSize() const {
        return getHashPreImage().getSize();
    }

    static Tuple createSizedTuple(size_t size);

    static Tuple createTuple(std::vector<Value> values);

    static Tuple createTuple(Value val);

    Tuple(Value val1, Value val2);

    Tuple(Value val1, Value val2, Value val3);

    Tuple(Value val1, Value val2, Value val3, Value val4);

    Tuple(Value val1, Value val2, Value val3, Value val4, Value val5);

    Tuple(Value val1,
          Value val2,
          Value val3,
          Value val4,
          Value val5,
          Value val6);

    Tuple(Value val1,
          Value val2,
          Value val3,
          Value val4,
          Value val5,
          Value val6,
          Value val7);

    Tuple(Value val1,
          Value val2,
          Value val3,
          Value val4,
          Value val5,
          Value val6,
          Value val7,
          Value val8);

    [[nodiscard]] uint64_t tuple_size() const {
        if (tpl) {
            return tpl->data.size();
        } else {
            return 0;
        }
    }

    void set_element(const uint64_t pos, Value newval);

    [[nodiscard]] Value get_element(const uint64_t pos) const;

    [[nodiscard]] std::vector<Value>::const_iterator begin() const {
        if (tpl == nullptr) {
            return empty_value_vector.begin();
        }
        return tpl->data.begin();
    }

    [[nodiscard]] std::vector<Value>::const_iterator end() const {
        if (tpl == nullptr) {
            return empty_value_vector.end();
        }
        return tpl->data.end();
    }

    [[nodiscard]] std::reverse_iterator<std::vector<Value>::const_iterator>
    rbegin() const {
        return std::reverse_iterator(end());
    }

    [[nodiscard]] std::reverse_iterator<std::vector<Value>::const_iterator>
    rend() const {
        return std::reverse_iterator(begin());
    }

    [[nodiscard]] const Value& get_element_unsafe(const uint64_t pos) const;

    [[nodiscard]] Value& get_element_mutable_unsafe(const uint64_t pos) const;

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

struct ValueTypeVisitor {
    ValueTypes operator()(const uint256_t&) const { return NUM; }
    ValueTypes operator()(const CodePointStub&) const { return CODEPT; }
    ValueTypes operator()(const Tuple&) const { return TUPLE; }
    ValueTypes operator()(const std::shared_ptr<HashPreImage>&) const {
        return TUPLE;
    }
    ValueTypes operator()(const Buffer&) const { return BUFFER; }
    ValueTypes operator()(const UnloadedValue& val) const { return val.type(); }
};

#endif /* tuple_hpp */
