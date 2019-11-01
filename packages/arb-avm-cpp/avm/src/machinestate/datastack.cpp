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
#include <bigint_utils.hpp>
#include <util.hpp>

uint256_t Datastack::hash() const {
    if (values.empty()) {
        return ::hash(Tuple());
    }
    calculateAllHashes();
    return hashes.back();
}

std::pair<uint256_t, std::vector<unsigned char>> Datastack::marshalForProof(
    const std::vector<bool>& stackInfo) {
    calculateAllHashes();
    Datastack c = *this;
    std::vector<unsigned char> buf;
    for (auto const& si : stackInfo) {
        value val = c.pop();
        if (si) {
            marshalShallow(val, buf);
        } else {
            buf.push_back(HASH_ONLY);
            std::array<unsigned char, 32> tmpbuf;
            to_big_endian(::hash(val), tmpbuf.begin());
            buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
        }
    }
    return std::make_pair(c.hash(), std::move(buf));
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
    uint256_t prev;
    if (hashes.size() > 0) {
        prev = hashes.back();
    } else {
        prev = ::hash(Tuple());
    }
    std::array<unsigned char, 1 + 2 * 32> tupData;
    auto oit = tupData.begin();
    tupData[0] = TUPLE + 2;
    ++oit;
    auto valHash = ::hash(values[hashes.size()]);
    std::array<uint64_t, 4> valHashInts;
    to_big_endian(valHash, valHashInts.begin());
    std::copy(reinterpret_cast<unsigned char*>(valHashInts.data()),
              reinterpret_cast<unsigned char*>(valHashInts.data()) + 32, oit);
    oit += 32;
    std::array<uint64_t, 4> valHashInts2;
    to_big_endian(prev, valHashInts2.begin());
    std::copy(reinterpret_cast<unsigned char*>(valHashInts2.data()),
              reinterpret_cast<unsigned char*>(valHashInts2.data()) + 32, oit);
    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(tupData.data(), 1 + 32 * 2, hashData.data());
    hashes.emplace_back(from_big_endian(hashData.begin(), hashData.end()));
}

void Datastack::calculateAllHashes() const {
    while (hashes.size() < values.size()) {
        addHash();
    }
}
