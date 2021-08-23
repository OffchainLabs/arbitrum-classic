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

#include <avm_values/buffer.hpp>

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
        zero = (res == hash(uint256_t(0)));
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
    auto acc = Buffer::fromData(std::vector<uint8_t>(buf, buf + sz));
    return acc.hash();
}

TEST_CASE("Buffer") {
    SECTION("empty buffer") {
        Buffer buf;
        REQUIRE(buf.hash() == hash(uint256_t(0)));
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
            arr[i] = static_cast<uint8_t>(rand() % 256);
        }
        REQUIRE(hash_buffer(arr, 0, SIZE) == hash_acc(arr, SIZE));
    }

    SECTION("hashing with zeroes") {
        const int SIZE = 1048576;
        const int FILL = 100000;
        uint8_t arr[SIZE];
        for (int i = 0; i < SIZE; i++) {
            arr[i] = i < FILL ? static_cast<uint8_t>(rand() % 256) : 0;
        }
        REQUIRE(hash_buffer(arr, 0, 131072) == hash_acc(arr, SIZE));
    }

    SECTION("hashing with single items") {
        const int SIZE = 1024 * 32;
        for (int j = 0; j < 1024; j++) {
            uint8_t arr[SIZE] = {};
            arr[j * 32] = 123;
            REQUIRE(hash_buffer(arr, 0, SIZE) == hash_acc(arr, SIZE));
        }
    }

    SECTION("set and read clone") {
        Buffer buf;
        Buffer buf2 = buf.set(10, 100);
        Buffer buf3 = buf.set(10, 150);
        REQUIRE(buf2.get(10) == 100);
        REQUIRE(buf3.get(10) == 150);
    }

    SECTION("flat construction") {
        std::vector<uint8_t> data;
        data.resize(32);
        std::fill(data.begin(), data.end(), 1);
        auto buf = Buffer::fromData(data);
        for (uint64_t i = 0; i < data.size(); i++) {
            REQUIRE(buf.get(i) == data[i]);
        }
    }

    SECTION("hashing random buffer") {
        std::random_device rd;
        std::mt19937 gen(rd());
        const int SIZE = 1024 * 32;
        std::uniform_int_distribution<uint64_t> distrib(0, SIZE - 1);
        for (int i = 0; i < 100; i++) {
            Buffer buf;
            uint8_t arr[SIZE] = {};
            for (int j = 0; j < 3; j++) {
                auto index = distrib(gen);
                buf = buf.set(index, 100);
                arr[index] = 100;
                REQUIRE(hash_buffer(arr, 0, SIZE) == buf.hash());
            }
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
        uint64_t idx = 300000L * 300000L;
        buf = buf.set(idx, 123);
        REQUIRE(buf.lastIndex() == idx);
        uint64_t idx2 = (1UL << 63UL);
        buf = buf.set(idx2, 123);
        REQUIRE(buf.lastIndex() == idx2);
    }
}

Buffer checkBuffer(ArbStorage& storage, Buffer& buf) {
    ValueCache value_cache{1, 0};
    {
        auto transaction = storage.makeReadWriteTransaction();
        auto results = saveValue(*transaction, buf);
        transaction->commit();
        REQUIRE(results.status.ok());
    }

    auto transaction = storage.makeReadTransaction();
    auto res = getValue(*transaction, hash_value(buf), value_cache, false);
    REQUIRE(std::holds_alternative<CountedData<value>>(res));
    REQUIRE(hash_value(std::get<CountedData<value>>(res).data) ==
            hash_value(buf));
    return std::get<Buffer>(std::get<CountedData<value>>(res).data);
}

TEST_CASE("Buffer Serialization") {
    DBDeleter deleter;
    ArbCoreConfig coreConfig{};
    ArbStorage storage(dbpath, coreConfig);

    ValueCache value_cache{1, 0};

    std::random_device
        rd;  // Will be used to obtain a seed for the random number engine
    std::mt19937 gen(rd());  // Standard mersenne_twister_engine seeded with
                             // rd()
    std::uniform_int_distribution<uint64_t> distrib(0, 20000);

    for (int i = 0; i < 1000; i++) {
        Buffer buf;
        for (int j = 0; j < 10; j++) {
            auto index = distrib(gen);
            buf = buf.set(index, 100);
        }
        auto buf2 = checkBuffer(storage, buf);
        for (int j = 0; j < 10; j++) {
            auto index = distrib(gen);
            buf = buf.set(index, 100);
            buf2 = buf2.set(index, 100);
            REQUIRE(buf.hash() == buf2.hash());
        }
        checkBuffer(storage, buf);
        checkBuffer(storage, buf2);
    }
    Buffer buf;
    buf = buf.set(300000L * 300000L, 12);
    checkBuffer(storage, buf);
    buf = buf.set(1UL << 63UL, 13);
    checkBuffer(storage, buf);
}

TEST_CASE("Buffer Hash Failure") {
    DBDeleter deleter;
    ArbCoreConfig coreConfig{};
    ArbStorage storage(dbpath, coreConfig);

    ValueCache value_cache{1, 0};

    Buffer buf;
    buf = buf.set(17750, 100);
    auto buf2 = checkBuffer(storage, buf);
    buf = buf.set(14721, 100);
    buf2 = buf2.set(14721, 100);
    REQUIRE(buf.hash() == buf2.hash());
}

std::vector<uint256_t> splitProof(std::vector<unsigned char> data) {
    auto buf = reinterpret_cast<const char*>(data.data());
    std::vector<uint256_t> res;
    for (uint64_t i = 0; i < data.size() / 32; i++) {
        auto a = deserializeUint256t(buf);
        res.push_back(a);
    }
    return res;
}

// From OneStepProof2.sol
uint256_t getProof(uint256_t buf, uint64_t loc, std::vector<uint256_t> proof) {
    // empty tree is full of zeros
    if (proof.size() == 0) {
        REQUIRE(buf == hash(uint256_t(0)));
        return 0;
    }
    uint256_t acc = hash(proof[0]);
    for (uint64_t i = 1; i < proof.size(); i++) {
        if (loc & 1)
            acc = hash(proof[i], acc);
        else
            acc = hash(acc, proof[i]);
        loc = loc >> 1;
    }
    REQUIRE(acc == buf);
    // maybe it is a zero outside the actual tree
    if (loc > 0)
        return 0;
    return proof[0];
}

uint256_t getByte(Buffer buf, uint64_t loc) {
    return buf.get(loc * 32 + 31);
}

TEST_CASE("Buffer get proofs") {
    SECTION("Empty buffer") {
        Buffer buf;
        for (uint64_t i = 0; i < 33; i++) {
            auto proof = buf.makeProof(i * 32);
            auto proof2 = splitProof(proof);
            REQUIRE(getByte(buf, i) == getProof(buf.hash(), i, proof2));
        }
        for (uint64_t i = 0; i < 32 * 33; i += 32) {
            auto proof = buf.makeProof(i * 32);
            auto proof2 = splitProof(proof);
            REQUIRE(getByte(buf, i) == getProof(buf.hash(), i, proof2));
        }
        for (uint64_t i = 32 * 32; i < 32 * 32 + 33; i++) {
            auto proof = buf.makeProof(i * 32);
            auto proof2 = splitProof(proof);
            REQUIRE(getByte(buf, i) == getProof(buf.hash(), i, proof2));
        }
    }

    SECTION("Buffer with one element") {
        Buffer buf;
        buf = buf.set(10000 * 32 + 31, 123);
        for (uint64_t i = 0; i < 33; i++) {
            auto proof = buf.makeProof(i * 32);
            auto proof2 = splitProof(proof);
            REQUIRE(getByte(buf, i) == getProof(buf.hash(), i, proof2));
        }
        for (uint64_t i = 0; i < 32 * 33; i += 32) {
            auto proof = buf.makeProof(i * 32);
            auto proof2 = splitProof(proof);
            REQUIRE(getByte(buf, i) == getProof(buf.hash(), i, proof2));
        }
        for (uint64_t i = 32 * 32; i < 32 * 32 + 33; i++) {
            auto proof = buf.makeProof(i * 32);
            auto proof2 = splitProof(proof);
            REQUIRE(getByte(buf, i) == getProof(buf.hash(), i, proof2));
        }
    }

    SECTION("Full buffer") {
        Buffer buf;
        for (uint64_t i = 0; i < 10000; i++) {
            buf = buf.set(i * 32 + 31, static_cast<uint8_t>(i % 256));
        }
        for (uint64_t i = 0; i < 100; i++) {
            auto proof = buf.makeProof(i * 32);
            auto proof2 = splitProof(proof);
            REQUIRE(getByte(buf, i) == getProof(buf.hash(), i, proof2));
        }
        for (uint64_t i = 10000; i < 10100; i++) {
            auto proof = buf.makeProof(i * 32);
            auto proof2 = splitProof(proof);
            REQUIRE(getByte(buf, i) == getProof(buf.hash(), i, proof2));
        }
    }

    SECTION("Random buffer") {
        std::random_device rd;
        std::mt19937 gen(rd());
        std::uniform_int_distribution<uint64_t> distrib(0, 20000);
        for (uint64_t i = 0; i < 100; i++) {
            Buffer buf;
            for (int j = 0; j < 3; j++) {
                auto index = distrib(gen);
                buf = buf.set(index * 32 + 31, 100);
            }
            for (int j = 0; j < 10; j++) {
                auto index = distrib(gen);
                auto proof = buf.makeProof(index * 32);
                auto proof2 = splitProof(proof);
                REQUIRE(getByte(buf, index) ==
                        getProof(buf.hash(), index, proof2));
            }
        }
    }
}

std::vector<uint256_t> makeZeros() {
    std::vector<uint256_t> zeros;
    zeros.resize(64);
    zeros[0] = hash(uint256_t(0));
    for (size_t i = 1; i < 64; i++) {
        zeros[i] = hash(zeros[i - 1], zeros[i - 1]);
    }
    return zeros;
}

size_t calcHeight(uint64_t loc) {
    if (loc == 0)
        return 1;
    else
        return 1 + calcHeight(loc >> 1);
}

// From OneStepProof2.sol
uint256_t setProof(uint256_t buf,
                   uint64_t loc,
                   uint256_t v,
                   std::vector<uint256_t> proof,
                   uint64_t nh,
                   uint256_t normal1,
                   uint256_t normal2) {
    // three possibilities, the tree depth stays same, it becomes lower or it's
    // extended
    uint256_t acc = hash(v);
    // check that the proof matches original
    getProof(buf, loc, proof);
    std::vector<uint256_t> zeros = makeZeros();
    // extended
    if (loc >= uint64_t(1 << (proof.size() - 1))) {
        if (v == 0)
            return buf;
        size_t height = calcHeight(loc);
        // build the left branch
        for (size_t i = proof.size(); i < height - 1; i++) {
            buf = hash(buf, zeros[i - 1]);
        }
        for (size_t i = 1; i < height - 1; i++) {
            if (loc & 1)
                acc = hash(zeros[i - 1], acc);
            else
                acc = hash(acc, zeros[i - 1]);
            loc = loc >> 1;
        }
        return hash(buf, acc);
    }
    for (uint64_t i = 1; i < proof.size(); i++) {
        uint256_t a = (loc & 1) ? proof[i] : acc;
        uint256_t b = (loc & 1) ? acc : proof[i];
        acc = hash(a, b);
        loc = loc >> 1;
    }
    if (v != 0)
        return acc;
    if (!(normal2 != zeros[nh] || nh == 0))
        throw std::runtime_error("fail");
    uint256_t res = nh == 0 ? normal1 : hash(normal1, normal2);
    uint256_t acc2 = res;
    for (uint64_t i = nh; i < proof.size() - 1; i++) {
        acc2 = hash(acc2, zeros[i]);
    }
    REQUIRE(acc2 == acc);
    return res;
}

void testSetProof(Buffer buf, uint64_t loc, uint8_t val) {
    auto prevHash = buf.hash();
    auto proof = buf.makeProof(loc * 32);
    auto proof2 = splitProof(proof);
    Buffer nbuf = buf.set(loc * 32 + 31, val);
    auto nproof = nbuf.makeNormalizationProof();
    auto nproof2 = splitProof(nproof);
    uint64_t nh = nproof[31];
    auto proofHash =
        setProof(prevHash, loc, val, proof2, nh, nproof2[1], nproof2[2]);
    REQUIRE(nbuf.hash() == proofHash);
}

TEST_CASE("Buffer set proofs") {
    SECTION("Empty buffer") {
        Buffer buf;
        for (uint64_t i = 0; i < 33; i++) {
            testSetProof(buf, i, 0);
            testSetProof(buf, i, 123);
            Buffer nbuf = buf.set(i * 32 + 31, 123);
            testSetProof(nbuf, i, 0);
        }
        for (uint64_t i = 0; i < 32 * 33; i += 32) {
            testSetProof(buf, i, 0);
            testSetProof(buf, i, 123);
            Buffer nbuf = buf.set(i * 32 + 31, 123);
            testSetProof(nbuf, i, 0);
        }
        for (uint64_t i = 32 * 32; i < 32 * 32 + 33; i++) {
            testSetProof(buf, i, 0);
            testSetProof(buf, i, 123);
            Buffer nbuf = buf.set(i * 32 + 31, 123);
            testSetProof(nbuf, i, 0);
        }
    }

    SECTION("Buffer with one elem") {
        Buffer buf;
        buf = buf.set(500 * 32 + 31, 123);
        for (uint64_t i = 0; i < 33; i++) {
            testSetProof(buf, i, 0);
            testSetProof(buf, i, 123);
            Buffer nbuf = buf.set(i * 32 + 31, 123);
            testSetProof(nbuf, i, 0);
        }
        for (uint64_t i = 0; i < 32 * 33; i += 32) {
            testSetProof(buf, i, 0);
            testSetProof(buf, i, 123);
            Buffer nbuf = buf.set(i * 32 + 31, 123);
            testSetProof(nbuf, i, 0);
        }
        for (uint64_t i = 32 * 32; i < 32 * 32 + 33; i++) {
            testSetProof(buf, i, 0);
            testSetProof(buf, i, 123);
            Buffer nbuf = buf.set(i * 32 + 31, 123);
            testSetProof(nbuf, i, 0);
        }
        uint64_t idx = 300000L * 300000L;
        testSetProof(buf, idx, 12);
        uint64_t idx2 = (1UL << 58UL);
        testSetProof(buf, idx2, 12);
    }

    SECTION("Full buffer") {
        Buffer buf;
        for (uint64_t i = 0; i < 10000; i++) {
            buf = buf.set(i * 32 + 31, static_cast<uint8_t>(i % 256));
        }
        for (uint64_t i = 0; i < 100; i += 10) {
            testSetProof(buf, i, 0);
            testSetProof(buf, i, 123);
        }
        for (uint64_t i = 10000; i < 10100; i += 10) {
            testSetProof(buf, i, 0);
            testSetProof(buf, i, 123);
        }
    }

    SECTION("Random buffer") {
        std::random_device rd;
        std::mt19937 gen(rd());
        std::uniform_int_distribution<uint64_t> distrib(0, 20000);
        for (uint64_t i = 0; i < 100; i++) {
            Buffer buf;
            for (uint64_t j = 0; j < 3; j++) {
                auto index = distrib(gen);
                buf = buf.set(index * 32 + 31, 100);
                testSetProof(buf, index, 0);
            }
            for (uint64_t j = 0; j < 10; j++) {
                auto index = distrib(gen);
                testSetProof(buf, index, 123);
            }
        }
    }
}
