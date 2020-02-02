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

#ifndef datastack_hpp
#define datastack_hpp

#include <avm_values/tuple.hpp>
#include <avm_values/value.hpp>

#include <vector>

class MachineStateSaver;
class MachineStateFetcher;
struct SaveResults;

class Datastack {
    static constexpr int lazyCount = 100;

    void addHash() const;
    void calculateAllHashes() const;
    void initializeDataStack(const Tuple& tuple);
    auto getTupleRepresentation(TuplePool* pool) const -> Tuple;

   public:
    std::vector<value> values;
    mutable std::vector<uint256_t> hashes;

    Datastack() {
        values.reserve(1000);
        hashes.reserve(1000);
    }

    void push(value&& newdata) {
        values.push_back(std::move(newdata));
        if (values.size() > hashes.size() + lazyCount) {
            addHash();
        }
    };

    auto operator[](size_t index) const -> const value& {
        return values[values.size() - 1 - index];
    }

    auto operator[](size_t index) -> value& {
        return values[values.size() - 1 - index];
    }

    auto pop() -> value {
        auto stackEmpty = values.empty();
        if (stackEmpty) {
            throw std::runtime_error("Stack is empty");
        }

        auto val = std::move(values.back());
        popClear();
        return val;
    }

    void prepForMod(int count) {
        while (hashes.size() > values.size() - count) {
            hashes.pop_back();
        }
    }

    void popClear() {
        values.pop_back();
        if (hashes.size() > values.size()) {
            hashes.pop_back();
        }
    }

    auto marshalForProof(const std::vector<bool>& stackInfo)
        -> std::pair<uint256_t, std::vector<unsigned char>>;

    auto peek() -> value& {
        if (values.size() == 0) {
            throw std::runtime_error("Stack is empty");
        }

        return values.back();
    }

    auto stacksize() -> uint64_t { return values.size(); }

    auto hash() const -> uint256_t;

    auto checkpointState(MachineStateSaver& saver, TuplePool* pool) const
        -> SaveResults;

    auto initializeDataStack(const MachineStateFetcher& fetcher,
                             const std::vector<unsigned char>& hash_key)
        -> bool;
};

auto operator<<(std::ostream& os, const Datastack& val) -> std::ostream&;

#endif /* datastack_hpp */
