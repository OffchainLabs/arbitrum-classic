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

#include <ethash/keccak.hpp>

#include <iostream>

Tuple Tuple::createSizedTuple(size_t size) {
    if (size == 0) {
        return {};
    }

    auto tpl = TuplePool::get_impl().getResource(size);
    tpl->data.resize(size);

    return Tuple(std::move(tpl));
}

Tuple Tuple::createTuple(std::vector<value> values) {
    if (!values.empty() && values.size() > 8) {
        return {};
    }

    auto tpl = TuplePool::get_impl().getResource(values.size());
    tpl->data.insert(tpl->data.end(), values.begin(), values.end());

    return Tuple(std::move(tpl));
}

Tuple Tuple::createTuple(value val) {
    auto tpl = TuplePool::get_impl().getResource(1);
    tpl->data.emplace_back(std::move(val));

    return Tuple(std::move(tpl));
}

Tuple::Tuple(value val1, value val2)
    : tpl(TuplePool::get_impl().getResource(2)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
}

Tuple::Tuple(value val1, value val2, value val3)
    : tpl(TuplePool::get_impl().getResource(3)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->data.push_back(std::move(val3));
}

Tuple::Tuple(value val1, value val2, value val3, value val4)
    : tpl(TuplePool::get_impl().getResource(4)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->data.push_back(std::move(val3));
    tpl->data.push_back(std::move(val4));
}

Tuple::Tuple(value val1, value val2, value val3, value val4, value val5)
    : tpl(TuplePool::get_impl().getResource(5)) {
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
             value val6)
    : tpl(TuplePool::get_impl().getResource(6)) {
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
             value val7)
    : tpl(TuplePool::get_impl().getResource(7)) {
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
             value val8)
    : tpl(TuplePool::get_impl().getResource(8)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->data.push_back(std::move(val3));
    tpl->data.push_back(std::move(val4));
    tpl->data.push_back(std::move(val5));
    tpl->data.push_back(std::move(val6));
    tpl->data.push_back(std::move(val7));
    tpl->data.push_back(std::move(val8));
}

constexpr uint64_t hash_size = 32;

// BasicValChecker checks to see whether a value can be hashed without
// recursion. All non-tuple values or tuples with a cached hash are
// basic. Tuples that haven't been hashed yet are not
struct BasicValChecker {
    bool operator()(const value& val) const { return std::visit(*this, val); }
    bool operator()(const Tuple& tup) const {
        return !tup.tpl || !tup.tpl->deferredHashing;
    }

    template <typename T>
    bool operator()(const T&) const {
        return true;
    }
};

HashPreImage calcHashPreImage(const Tuple& tup) {
    std::array<unsigned char, 1 + 8 * 32> tupData{};
    uint256_t size = 1;

    tupData[0] = tup.tuple_size();
    auto oit = tupData.begin();
    ++oit;

    for (uint64_t i = 0; i < tup.tuple_size(); i++) {
        const auto& element = tup.get_element(i);
        oit = to_big_endian(hash_value(element), oit);
        size += ::getSize(element);
    }

    auto hash_val = ethash::keccak256(
        tupData.data(),
        static_cast<unsigned int>(1 + hash_size * (tup.tuple_size())));
    std::array<unsigned char, 32> hashData{};
    std::copy(&hash_val.bytes[0], &hash_val.bytes[32], hashData.begin());

    return HashPreImage{hashData, size};
}

void Tuple::calculateHashPreImage() const {
    // Make sure children are already hashed
    std::vector<Tuple> tups{*this};
    while (!tups.empty()) {
        Tuple tup = tups.back();
        if (BasicValChecker{}(tup)) {
            tups.pop_back();
        } else {
            bool found_complex = false;
            for (uint64_t i = 0; i < tup.tuple_size(); ++i) {
                auto& elem = tup.get_element_unsafe(i);
                if (!BasicValChecker{}(elem)) {
                    found_complex = true;
                    tups.push_back(std::get<Tuple>(tup.get_element(i)));
                }
            }
            if (!found_complex) {
                tup.tpl->cachedPreImage = calcHashPreImage(tup);
                tup.tpl->deferredHashing = false;
                tups.pop_back();
            }
        }
    }
}

HashPreImage zeroPreimage() {
    std::array<unsigned char, 1> tupData{};
    tupData[0] = 0;

    auto hash_val = ethash::keccak256(tupData.data(), 1);

    std::array<unsigned char, 32> hashData{};
    std::copy(&hash_val.bytes[0], &hash_val.bytes[32], hashData.begin());

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
