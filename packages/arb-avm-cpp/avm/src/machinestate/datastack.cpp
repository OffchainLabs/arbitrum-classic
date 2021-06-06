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

#include <avm/machinestate/datastack.hpp>

#include <iostream>
#include <utility>

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

DataStackProof Datastack::marshalForProof(const std::vector<size_t>& stackInfo,
                                          const Code& code) const {
    calculateAllHashes();
    Datastack c = *this;
    std::vector<unsigned char> buf;
    std::vector<value> val;

    // If the stack is underflowing, just send what's left
    uint8_t items_to_pop = stackInfo.size();
    bool underflow = false;
    if (c.stacksize() < items_to_pop) {
        items_to_pop = c.stacksize();
        underflow = true;
    }

    for (size_t i = 0; i < items_to_pop; ++i) {
        val.push_back(c.pop());
    }

    // Marshal the values from deepest to most shallow in the stack
    for (size_t i = 0; i < val.size(); ++i) {
        auto index = val.size() - 1 - i;
        // Only marshal a stub if we are underflowing
        auto level = underflow ? 0 : stackInfo[index];
        ::marshalForProof(val[index], level, buf, code);
    }

    return {c.getHashPreImage(), std::move(buf), items_to_pop};
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

Tuple Datastack::getTupleRepresentation() const {
    Tuple rep;
    for (const auto& val : values) {
        rep = Tuple(val, rep);
    }
    return rep;
}

Datastack::Datastack(Tuple tuple_rep) : Datastack() {
    Tuple ret = std::move(tuple_rep);
    std::vector<value> vals;
    while (ret.tuple_size() == 2) {
        vals.push_back(ret.get_element(0));
        ret = std::get<Tuple>(ret.get_element(1));
    }

    for (size_t i = 0; i < vals.size(); i++) {
        push(std::move(vals[vals.size() - 1 - i]));
    }
}

void Datastack::addHash() const {
    HashPreImage prev = [&]() {
        if (!hashes.empty()) {
            return hashes.back();
        } else {
            return Tuple().getHashPreImage();
        }
    }();

    auto newVal = values[hashes.size()];
    auto tup = Tuple(newVal, std::make_shared<HashPreImage>(prev));
    hashes.emplace_back(tup.getHashPreImage());
}

void Datastack::calculateAllHashes() const {
    while (hashes.size() < values.size()) {
        addHash();
    }
}
