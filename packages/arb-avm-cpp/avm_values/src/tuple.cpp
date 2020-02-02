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

#include <avm_values/tuple.hpp>

#include <avm_values/util.hpp>
#include <bigint_utils.hpp>

Tuple::Tuple(value val, TuplePool* pool)
    : tuplePool(pool), tpl(pool->getResource(1)) {
    tpl->data.push_back(std::move(val));
    tpl->deferredHashing = true;
}

Tuple::Tuple(value val1, value val2, TuplePool* pool)
    : tuplePool(pool), tpl(pool->getResource(2)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->deferredHashing = true;
}

Tuple::Tuple(value val1, value val2, value val3, TuplePool* pool)
    : tuplePool(pool), tpl(pool->getResource(3)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->data.push_back(std::move(val3));
    tpl->deferredHashing = true;
}

Tuple::Tuple(value val1, value val2, value val3, value val4, TuplePool* pool)
    : tuplePool(pool), tpl(pool->getResource(4)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->data.push_back(std::move(val3));
    tpl->data.push_back(std::move(val4));
    tpl->deferredHashing = true;
}

Tuple::Tuple(value val1,
             value val2,
             value val3,
             value val4,
             value val5,
             TuplePool* pool)
    : tuplePool(pool), tpl(pool->getResource(5)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->data.push_back(std::move(val3));
    tpl->data.push_back(std::move(val4));
    tpl->data.push_back(std::move(val5));
    tpl->deferredHashing = true;
}

Tuple::Tuple(value val1,
             value val2,
             value val3,
             value val4,
             value val5,
             value val6,
             TuplePool* pool)
    : tuplePool(pool), tpl(pool->getResource(6)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->data.push_back(std::move(val3));
    tpl->data.push_back(std::move(val4));
    tpl->data.push_back(std::move(val5));
    tpl->data.push_back(std::move(val6));
    tpl->deferredHashing = true;
}

Tuple::Tuple(value val1,
             value val2,
             value val3,
             value val4,
             value val5,
             value val6,
             value val7,
             TuplePool* pool)
    : tuplePool(pool), tpl(pool->getResource(7)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->data.push_back(std::move(val3));
    tpl->data.push_back(std::move(val4));
    tpl->data.push_back(std::move(val5));
    tpl->data.push_back(std::move(val6));
    tpl->data.push_back(std::move(val7));
    tpl->deferredHashing = true;
}

Tuple::Tuple(value val1,
             value val2,
             value val3,
             value val4,
             value val5,
             value val6,
             value val7,
             value val8,
             TuplePool* pool)
    : tuplePool(pool), tpl(pool->getResource(8)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->data.push_back(std::move(val3));
    tpl->data.push_back(std::move(val4));
    tpl->data.push_back(std::move(val5));
    tpl->data.push_back(std::move(val6));
    tpl->data.push_back(std::move(val7));
    tpl->data.push_back(std::move(val8));
    tpl->deferredHashing = true;
}

Tuple::Tuple(std::vector<value> values, TuplePool* pool)
    : tuplePool(pool), tpl(pool->getResource(values.size())) {
    if (!values.empty()) {
        assert(tpl.get() != nullptr);
        for (auto& val : values) {
            tpl->data.push_back(std::move(val));
        }

        tpl->cachedHash = calculateHash();
    }
}

void Tuple::marshal(std::vector<unsigned char>& buf) const {
    buf.push_back(TUPLE + tuple_size());
    for (uint64_t i = 0; i < tuple_size(); i++) {
        marshal_value(get_element(i), buf);
    }
}

// marshalForProof does not use this
// see similar functionality in value.marshalShallow
auto Tuple::clone_shallow() -> value {
    Tuple tup(tuplePool, tuple_size());
    for (uint64_t i = 0; i < tuple_size(); i++) {
        auto val = get_element(i);
        if (nonstd::holds_alternative<uint256_t>(val)) {
            tup.set_element(i, val);
        } else {
            auto valHash = hash(get_element(i));
            tup.set_element(i, valHash);
        }
    }
    return tup;
}

auto Tuple::calculateHash() const -> uint256_t {
    std::array<unsigned char, 1 + 8 * 32> tupData;
    auto oit = tupData.begin();
    tupData[0] = TUPLE + tuple_size();
    ++oit;
    for (uint64_t i = 0; i < tuple_size(); i++) {
        auto valHash = hash(get_element(i));
        std::array<uint64_t, 4> valHashInts;
        to_big_endian(valHash, valHashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(valHashInts.data()),
                  reinterpret_cast<unsigned char*>(valHashInts.data()) + 32,
                  oit);
        oit += 32;
    }

    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(tupData.data(), 1 + 32 * tuple_size(), hashData.data());
    return from_big_endian(hashData.begin(), hashData.end());
}

auto zeroHash() -> uint256_t {
    std::array<unsigned char, 1> tupData;
    tupData[0] = TUPLE;
    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(tupData.data(), 1, hashData.data());
    return from_big_endian(hashData.begin(), hashData.end());
}

auto operator<<(std::ostream& os, const Tuple& val) -> std::ostream& {
    os << "Tuple(";
    for (uint64_t i = 0; i < val.tuple_size(); i++) {
        os << val.get_element(i);
        if (i < val.tuple_size() - 1) {
            os << ", ";
        }
    }
    os << ")";
    return os;
}
