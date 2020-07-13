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

#include <avm_values/tuple.hpp>

#include <avm_values/util.hpp>
#include <bigint_utils.hpp>

Tuple::Tuple(value val, TuplePool* pool)
    : tuplePool(pool), tpl(pool->getResource(1)) {
    tpl->data.push_back(std::move(val));
}

Tuple::Tuple(value val1, value val2, TuplePool* pool)
    : tuplePool(pool), tpl(pool->getResource(2)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
}

Tuple::Tuple(value val1, value val2, value val3, TuplePool* pool)
    : tuplePool(pool), tpl(pool->getResource(3)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->data.push_back(std::move(val3));
}

Tuple::Tuple(value val1, value val2, value val3, value val4, TuplePool* pool)
    : tuplePool(pool), tpl(pool->getResource(4)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->data.push_back(std::move(val3));
    tpl->data.push_back(std::move(val4));
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
}

Tuple::Tuple(std::vector<value> values, TuplePool* pool) : tuplePool(pool) {
    if (!values.empty() && values.size() < 9) {
        tpl = pool->getResource(values.size());
        for (auto& val : values) {
            tpl->data.push_back(std::move(val));
        }
    }
}

void Tuple::marshal(std::vector<unsigned char>& buf) const {
    buf.push_back(TUPLE + tuple_size());
    for (uint64_t i = 0; i < tuple_size(); i++) {
        marshal_value(get_element(i), buf);
    }
}

HashPreImage Tuple::calculateHashPreImage() const {
    std::array<unsigned char, 1 + 8 * 32> tupData;
    uint256_t size = 1;

    tupData[0] = tuple_size();
    auto oit = tupData.begin();
    ++oit;

    int val_length = 32;
    for (uint64_t i = 0; i < tuple_size(); i++) {
        const auto& element = get_element(i);
        auto valHash = hash_value(element);
        oit = to_big_endian(valHash, oit);
        size += ::getSize(element);
    }

    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(tupData.data(),
                    static_cast<unsigned int>(1 + val_length * (tuple_size())),
                    hashData.data());

    return HashPreImage{hashData, size};
}

HashPreImage zeroPreimage() {
    std::array<unsigned char, 1> tupData;
    tupData[0] = 0;

    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(tupData.data(), 1, hashData.data());
    return HashPreImage(hashData, 1);
}

std::ostream& operator<<(std::ostream& os, const Tuple& val) {
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
