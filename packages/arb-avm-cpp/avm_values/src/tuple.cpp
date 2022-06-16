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
#include <avm_values/value.hpp>

#include <ethash/keccak.hpp>

#include <iostream>

Tuple Tuple::createSizedTuple(const size_t size) {
    if (size == 0) {
        return {};
    }

    auto tpl = TuplePool::get_impl().getResource(size);
    tpl->data.resize(size);

    return Tuple(std::move(tpl));
}

Tuple Tuple::createTuple(std::vector<Value> values) {
    if (!values.empty() && values.size() > 8) {
        return {};
    }

    auto tpl = TuplePool::get_impl().getResource(values.size());
    tpl->data.insert(tpl->data.end(), values.begin(), values.end());

    return Tuple(std::move(tpl));
}

Tuple Tuple::createTuple(Value val) {
    auto tpl = TuplePool::get_impl().getResource(1);
    tpl->data.emplace_back(std::move(val));

    return Tuple(std::move(tpl));
}

Tuple::Tuple(Value val1, Value val2)
    : tpl(TuplePool::get_impl().getResource(2)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
}

Tuple::Tuple(Value val1, Value val2, Value val3)
    : tpl(TuplePool::get_impl().getResource(3)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->data.push_back(std::move(val3));
}

Tuple::Tuple(Value val1, Value val2, Value val3, Value val4)
    : tpl(TuplePool::get_impl().getResource(4)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->data.push_back(std::move(val3));
    tpl->data.push_back(std::move(val4));
}

Tuple::Tuple(Value val1, Value val2, Value val3, Value val4, Value val5)
    : tpl(TuplePool::get_impl().getResource(5)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->data.push_back(std::move(val3));
    tpl->data.push_back(std::move(val4));
    tpl->data.push_back(std::move(val5));
}

Tuple::Tuple(Value val1,
             Value val2,
             Value val3,
             Value val4,
             Value val5,
             Value val6)
    : tpl(TuplePool::get_impl().getResource(6)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->data.push_back(std::move(val3));
    tpl->data.push_back(std::move(val4));
    tpl->data.push_back(std::move(val5));
    tpl->data.push_back(std::move(val6));
}

Tuple::Tuple(Value val1,
             Value val2,
             Value val3,
             Value val4,
             Value val5,
             Value val6,
             Value val7)
    : tpl(TuplePool::get_impl().getResource(7)) {
    tpl->data.push_back(std::move(val1));
    tpl->data.push_back(std::move(val2));
    tpl->data.push_back(std::move(val3));
    tpl->data.push_back(std::move(val4));
    tpl->data.push_back(std::move(val5));
    tpl->data.push_back(std::move(val6));
    tpl->data.push_back(std::move(val7));
}

Tuple::Tuple(Value val1,
             Value val2,
             Value val3,
             Value val4,
             Value val5,
             Value val6,
             Value val7,
             Value val8)
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

void Tuple::unsafe_set_element(uint64_t pos, Value&& newval) {
    if (pos >= tuple_size()) {
        throw bad_tuple_index{};
    }
    tpl->data[pos] = std::move(newval);
    tpl->deferredHashing = true;
}

void Tuple::set_element(const uint64_t pos, Value newval) {
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

Value Tuple::get_element(const uint64_t pos) const {
    if (pos >= tuple_size()) {
        throw bad_tuple_index{};
    }
    return tpl->data[pos];
}

[[nodiscard]] const Value& Tuple::get_element_unsafe(const uint64_t pos) const {
    return tpl->data[pos];
}

[[nodiscard]] Value& Tuple::get_element_mutable_unsafe(
    const uint64_t pos) const {
    tpl->deferredHashing = true;
    return tpl->data[pos];
}

constexpr uint64_t hash_size = 32;

// BasicValChecker checks to see whether a value can be hashed without
// recursion. All non-tuple values or tuples with a cached hash are
// basic. Tuples that haven't been hashed yet are not
struct BasicValChecker {
    bool operator()(const Value& val) const { return visit(*this, val); }
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
                    tups.push_back(get<Tuple>(tup.get_element(i)));
                }
            }
            if (!found_complex) {
                // It's fine if multiple threads write the cached preimage
                // simultaneously, as they must be writing the same hash, and
                // this uses atomic writes.
                tup.tpl->cachedPreImage.writeAtomic(calcHashPreImage(tup));
                // The "release" ordering here ensures any other thread
                // with the default seqcst "acquire" ordering will see
                // the updated hash pre image if it sees this write.
                tup.tpl->deferredHashing.store(false,
                                               std::memory_order_release);
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
