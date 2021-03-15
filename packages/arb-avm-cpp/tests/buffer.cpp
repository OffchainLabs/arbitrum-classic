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

#include "helper.hpp"

#include <data_storage/arbstorage.hpp>
#include <data_storage/storageresult.hpp>

#include <avm/machinestate/datastack.hpp>

#include <catch2/catch.hpp>

#include <ethash/keccak.hpp>

uint256_t hash_buffer_aux(uint8_t* buf,
                          int offset,
                          int sz,
                          bool pack,
                          bool& zero) {
    if (sz == 32) {
        auto hash_val = ethash::keccak256(buf + offset, 32);
        auto res = intx::be::load<uint256_t>(hash_val);
        zero = (res == hash(0));
        return res;
    }
    bool zero1 = false;
    bool zero2 = false;
    auto h2 = hash_buffer_aux(buf, offset + sz / 2, sz / 2, false, zero2);
    if (pack && zero2) {
        return hash_buffer_aux(buf, offset, sz / 2, pack, zero);
    }
    auto h1 = hash_buffer_aux(buf, offset, sz / 2, false, zero1);
    zero = zero1 && zero2;
    return hash(h1, h2);
}

uint256_t hash_buffer(uint8_t* buf, int offset, int sz) {
    bool zero = false;
    return hash_buffer_aux(buf, offset, sz, true, zero);
}

uint256_t hash_acc(uint8_t* buf, int sz) {
    Buffer acc;
    for (int i = 0; i < sz; i++) {
        acc = acc.set(i, buf[i]);
    }
    return acc.hash();
}

TEST_CASE("Buffer") {
    SECTION("calc len") {
        REQUIRE(calc_len(0) == 1024UL);
        REQUIRE(calc_len(1) == 8UL * 1024UL);
        REQUIRE(calc_len(16) == 8UL * 8UL * 8UL * 8UL * 8UL * 8UL * 8UL * 8UL *
                                    8UL * 8UL * 8UL * 8UL * 8UL * 8UL * 8UL *
                                    8UL * 1024UL);
        REQUIRE(calc_len(17) == 8UL * 8UL * 8UL * 8UL * 8UL * 8UL * 8UL * 8UL *
                                    8UL * 8UL * 8UL * 8UL * 8UL * 8UL * 8UL *
                                    8UL * 8UL * 1024UL);
        REQUIRE(calc_len(18) == 0);  // Overflow
    }

    SECTION("calc height") {
        REQUIRE(calc_height(0) == 10UL);
        REQUIRE(calc_height(1) == 1UL * 3UL + 10UL);
        REQUIRE(calc_height(10) == 10UL * 3UL + 10UL);
        REQUIRE(calc_height(1000) == 1000UL * 3UL + 10UL);
    }

    SECTION("needed height") {
        REQUIRE(needed_height(0) == 1);
        REQUIRE(needed_height(1) == 1);
        REQUIRE(needed_height(2) == 2);
        REQUIRE(needed_height(3) == 2);
        REQUIRE(needed_height(4) == 3);
        REQUIRE(needed_height(7) == 3);
        REQUIRE(needed_height(8) == 4);
        REQUIRE(needed_height(16) == 5);
        REQUIRE(needed_height(32) == 6);
        REQUIRE(needed_height(64) == 7);
        REQUIRE(needed_height(128) == 8);
        REQUIRE(needed_height(256) == 9);
        REQUIRE(needed_height(512) == 10);
        REQUIRE(needed_height(1023) == 10);
        REQUIRE(needed_height(1024) == 11);
    }

    SECTION("empty buffer") {
        Buffer buf;
        REQUIRE(buf.hash() == hash(0));
    }

    SECTION("setting") {
        Buffer buf;
        buf = buf.set(1000, 123);
        REQUIRE(buf.get(1000) == 123);
    }

    SECTION("hashing") {
        const int SIZE = 1048576;
        uint8_t arr[SIZE];
        for (int i = 0; i < SIZE; i++) {
            arr[i] = rand() % 256;
        }
        REQUIRE(hash_buffer(arr, 0, SIZE) == hash_acc(arr, SIZE));
    }

    SECTION("hashing with zeroes") {
        const int SIZE = 1048576;
        const int FILL = 100000;
        uint8_t arr[SIZE];
        for (int i = 0; i < SIZE; i++) {
            arr[i] = i < FILL ? rand() % 256 : 0;
        }
        REQUIRE(hash_buffer(arr, 0, 131072) == hash_acc(arr, SIZE));
    }

    SECTION("hashing with single zeroes") {
        const int SIZE = 1024 * 32;
        for (int j = 0; j < 1024; j++) {
            uint8_t arr[SIZE] = {};
            arr[j * 32] = 123;
            REQUIRE(hash_buffer(arr, 0, SIZE) == hash_acc(arr, SIZE));
        }
    }

    SECTION("last index") {
        Buffer buf;
        REQUIRE(buf.lastIndex() == 0);
        buf = buf.set(0, 123);
        REQUIRE(buf.lastIndex() == 0);
        buf = buf.set(1, 123);
        REQUIRE(buf.lastIndex() == 1);
        buf = buf.set(10, 123);
        REQUIRE(buf.lastIndex() == 10);
        buf = buf.set(31, 123);
        REQUIRE(buf.lastIndex() == 31);
        buf = buf.set(1000, 123);
        REQUIRE(buf.lastIndex() == 1000);
        buf = buf.set(2000, 123);
        REQUIRE(buf.lastIndex() == 2000);
        buf = buf.set(20000, 123);
        REQUIRE(buf.lastIndex() == 20000);
        buf = buf.set(300000, 123);
        REQUIRE(buf.lastIndex() == 300000);
    }
}

TEST_CASE("Buffer Serialization") {
    DBDeleter deleter;
    ArbStorage storage(dbpath);

    ValueCache value_cache{0, 0};

    Buffer buf;
    buf = buf.set(8191, 123);

    {
        auto transaction = storage.makeReadWriteTransaction();
        auto results = saveValue(*transaction, buf);
        transaction->commit();
        REQUIRE(results.status.ok());
    }

    {
        auto transaction = storage.makeReadTransaction();
        auto res = getValue(*transaction, hash_value(buf), value_cache);
        REQUIRE(std::holds_alternative<CountedData<value>>(res));
        REQUIRE(hash_value(std::get<CountedData<value>>(res).data) ==
                hash_value(buf));
    }
}
