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

#include <data_storage/checkpointstorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/value.hpp>

#include <avm/machinestate/datastack.hpp>

#include <catch2/catch.hpp>

#include <ethash/keccak.hpp>

uint256_t hash_buffer(uint8_t *buf, int offset, int sz, bool pack) {
    if (sz == 32) {
        auto hash_val = ethash::keccak256(buf+offset, 32);
        return intx::be::load<uint256_t>(hash_val);
    }
    auto h2 = hash_buffer(buf, offset+sz/2, sz/2, false);
    if (pack && hash(0) == h2) {
        return hash_buffer(buf, offset, sz/2, true);
    }
    auto h1 = hash_buffer(buf, offset, sz/2, false);
    return hash(h1, h2);
}

uint256_t hash_acc(uint8_t *buf, int sz) {
    Buffer acc;
    for (int i = 0; i < sz; i++) {
        acc = acc.set(i, buf[i]);
    }
    return acc.hash();
}

TEST_CASE("Buffer") {
    SECTION("empty buffer") {
        Buffer buf;
        assert(buf.hash() == hash(0));
    }

    SECTION("setting") {
        Buffer buf;
        buf = buf.set(1000, 123);
        assert(buf.get(1000) == 123);
    }

    SECTION("hashing") {
        const int SIZE = 1048576;
        uint8_t arr[SIZE];
        for (int i = 0; i < SIZE; i++) {
            arr[i] = rand() % 256;
        }
        assert(hash_buffer(arr, 0, SIZE, true) == hash_acc(arr, SIZE));
    }

    SECTION("hashing with zeroes") {
        const int SIZE = 1048576;
        const int FILL = 100000;
        uint8_t arr[SIZE];
        for (int i = 0; i < SIZE; i++) {
            arr[i] = i < FILL ? rand() % 256 : 0;
        }
        assert(hash_buffer(arr, 0, 131072, true) == hash_acc(arr, SIZE));
    }
}

