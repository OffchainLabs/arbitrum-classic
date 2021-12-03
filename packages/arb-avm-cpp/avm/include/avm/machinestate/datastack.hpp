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

#ifndef datastack_hpp
#define datastack_hpp

#include <avm_values/tuple.hpp>
#include <avm_values/value.hpp>

#include <vector>

class Transaction;
struct SaveResults;

struct DataStackProof {
    HashPreImage bottom;
    std::vector<unsigned char> data;
    uint8_t count;
};

class Datastack {
    // lazyCount defines how many unhashed items are allowed on the stack
    // This serves to bound the total time it can take to hash the machine
    // while removing the need to hash the stack while churn is occuring
    // near the top
    static constexpr int lazyCount = 100000;

    void addHash() const;
    void calculateAllHashes() const;

   public:
    std::vector<Value> values;
    mutable std::vector<HashPreImage> hashes;

    Datastack() {
        values.reserve(lazyCount);
        hashes.reserve(lazyCount);
    }

    explicit Datastack(Tuple tuple_rep);

    Tuple getTupleRepresentation() const;

    void push(Value&& newdata) {
        values.push_back(std::move(newdata));
        if (values.size() > hashes.size() + lazyCount) {
            addHash();
        }
    }

    const Value& operator[](size_t index) const {
        if (index >= values.size()) {
            throw stack_too_small();
        }
        return values[values.size() - 1 - index];
    }

    Value& operator[](size_t index) {
        if (index >= values.size()) {
            throw stack_too_small();
        }
        return values[values.size() - 1 - index];
    }

    Value pop() {
        auto stackEmpty = values.empty();
        if (stackEmpty) {
            throw std::runtime_error("Stack is empty");
        }

        auto val = std::move(values.back());
        popClear();
        return val;
    }

    void prepForMod(size_t count) {
        if (static_cast<size_t>(count) > values.size()) {
            throw stack_too_small();
        }
        while (!hashes.empty() && hashes.size() > values.size() - count) {
            hashes.pop_back();
        }
    }

    void popClear() {
        if (values.empty()) {
            throw stack_too_small();
        }
        values.pop_back();
        if (hashes.size() > values.size()) {
            hashes.pop_back();
        }
    }

    DataStackProof marshalForProof(const std::vector<size_t>& stackInfo,
                                   const Code& code) const;

    Value& peek() {
        if (values.empty()) {
            throw stack_too_small();
        }

        return values.back();
    }

    uint64_t stacksize() const { return values.size(); }

    uint256_t hash() const;

    HashPreImage getHashPreImage() const;

    uint256_t getTotalValueSize() const;
};

std::ostream& operator<<(std::ostream& os, const Datastack& val);

#endif /* datastack_hpp */
