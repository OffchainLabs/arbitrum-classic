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

#include <avm/machinestate/datastack.hpp>

#include <data_storage/checkpoint/machinestatefetcher.hpp>
#include <data_storage/checkpoint/machinestatesaver.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/transaction.hpp>

#include <avm_values/util.hpp>
#include <bigint_utils.hpp>

uint256_t Datastack::hash() const {
    auto h_value = getHashPreImage();
    return h_value.hash();
}

uint256_t Datastack::getTotalValueSize() const {
    auto h_value = getHashPreImage();
    return h_value.getSize();
}

HashPreImage Datastack::getHashPreImage() const {
    if (values.empty()) {
        return Tuple().getHashPreImage();
    } else {
        calculateAllHashes();
        return hashes.back();
    }
}

std::pair<HashPreImage, std::vector<unsigned char>> Datastack::marshalForProof(
    const std::vector<bool>& stackInfo) {
    calculateAllHashes();
    Datastack c = *this;
    std::vector<unsigned char> buf;
    for (auto const& si : stackInfo) {
        value val = c.pop();
        if (si) {
            marshalShallow(val, buf);
        } else {
            marshalStub(val, buf);
        }
    }
    return std::make_pair(c.getHashPreImage(), std::move(buf));
}

std::ostream& operator<<(std::ostream& os, const Datastack& val) {
    os << "[";
    for (uint64_t i = 0; i < val.values.size(); i++) {
        os << val.values[val.values.size() - 1 - i];
        if (i < val.values.size() - 1) {
            os << ", ";
        }
    }
    os << "]";
    return os;
}

// can speed up by not creating tuple/save directly
SaveResults Datastack::checkpointState(MachineStateSaver& saver,
                                       TuplePool* pool) {
    auto tuple = getTupleRepresentation(pool);
    return saver.saveTuple(tuple);
}

bool Datastack::initializeDataStack(
    const MachineStateFetcher& fetcher,
    const std::vector<unsigned char>& hash_key) {
    auto results = fetcher.getTuple(hash_key);
    initializeDataStack(results.data);
    return results.status.ok();
}

Tuple Datastack::getTupleRepresentation(TuplePool* pool) {
    if (values.empty()) {
        return Tuple();
    } else {
        auto current_tuple = Tuple(values[0], pool);

        for (size_t i = 1; i < values.size(); i++) {
            auto new_tuple = Tuple(values[i], current_tuple, pool);
            current_tuple = new_tuple;
        }
        return current_tuple;
    }
}

void Datastack::initializeDataStack(const Tuple& tuple) {
    if (tuple.tuple_size() == 1) {
        push(tuple.get_element(0));
    } else if (tuple.tuple_size() == 2) {
        // catch exception if not tuple?
        auto inner_tuple = nonstd::get<Tuple>(tuple.get_element(1));
        initializeDataStack(inner_tuple);

        auto current_val = tuple.get_element(0);
        push(tuple.get_element(0));
    }
}

void Datastack::addHash() const {
    HashPreImage prev;
    if (hashes.size() > 0) {
        prev = hashes.back();
    } else {
        prev = Tuple().getHashPreImage();
    }

    auto newVal = values[hashes.size()];
    TuplePool pool;
    auto tup = Tuple(newVal, prev, &pool);
    hashes.emplace_back(tup.getHashPreImage());
}

void Datastack::calculateAllHashes() const {
    while (hashes.size() < values.size()) {
        addHash();
    }
}
