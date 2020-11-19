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

#include <avm_values/buffer.hpp>
#include <avm_values/bigint.hpp>

#include <ethash/keccak.hpp>

uint256_t zero_hash(int sz) {
    if (sz == 32) {
        return hash(0);
    }
    auto h1 = zero_hash(sz/2);
    return hash(h1, h1);
}

Packed normal(uint256_t hash, int sz) {
    return Packed{hash, sz, 0};
}

Packed pack(const Packed& packed) {
    return Packed{packed.hash, packed.size, packed.packed+1};
}

bool is_zero_hash(const Packed& packed) {
    return packed.hash == hash(0);
}

uint256_t unpack(const Packed &packed) {
    uint256_t res = packed.hash;
    int sz = packed.size;
    for (int i = 0; i < packed.packed; i++) {
        res = hash(res, zero_hash(sz));
        sz = sz*2;
    }
    return res;
}

Packed zero_packed(int sz) {
    if (sz == 32) {
        return normal(zero_hash(32), 32);
    }
    return pack(zero_packed(sz/2));
}

Packed hash_buf(uint8_t *buf, int offset, int sz) {
    if (sz == 32) {
        auto hash_val = ethash::keccak256(buf+offset, 32);
        uint256_t res = intx::be::load<uint256_t>(hash_val);
        return normal(res, 32);
    }
    // std::cerr << "hashing " << offset << " to " << (offset+sz) << std::endl;
    auto h1 = hash_buf(buf, offset, sz/2);
    auto h2 = hash_buf(buf, offset+sz/2, sz/2);
    if (is_zero_hash(h2)) {
        return pack(h1);
    }
    return normal(hash(unpack(h1), unpack(h2)), sz);
}

Packed hash_node(RawBuffer *buf, int offset, int len, int sz) {
    //    std::cerr << "hashing " << sz << " " << offset << " " << len << std::endl;
    if (len == 1) {
        return buf[offset].hash_aux();
    }
    auto h1 = hash_node(buf, offset, len/2, sz/2);
    auto h2 = hash_node(buf, offset + len/2, len/2, sz/2);
    //    std::cerr << "hashed " << sz << " " << offset << " " << len << std::endl;
    if (is_zero_hash(h2)) {
        return pack(h1);
    }
    return normal(hash(unpack(h1), unpack(h2)), sz);
}

uint256_t RawBuffer::hash() {
    uint256_t res = hash_aux().hash;
    // std::cerr << "Finished hashing " << size() << ":" << static_cast<uint64_t>(res) << std::endl;
    return res;
}

Packed RawBuffer::hash_aux() {
    if (saved) {
        // std::cerr << "found saved hash" << std::endl;
        return savedHash;
    }
    Packed res;
    if (level == 0) {
        // std::cerr << "Hashing buffer..." << std::endl;
        if (!leaf || leaf->size() == 0) res = zero_packed(LEAF_SIZE);
        else res = hash_buf(leaf->data(), 0, LEAF_SIZE);
    } else {
        if (!node) res = zero_packed(calc_len(level));
        else {
            // std::cerr << "Hashing node..." << static_cast<void*>(this) << std::endl;
            res = hash_node(node->data(), 0, NODE_SIZE, calc_len(level));
        }
    }
    saved = true;
    savedHash = res;
    // std::cerr << "Finished hashing " << size() << ":" << static_cast<uint64_t>(res.hash) << " ? " << saved << std::endl;
    return res;
}

RawBuffer RawBuffer::normalize() {
    if (hash() == zero_hash(32)) {
        return RawBuffer();
    }
    if (level == 0) {
        return *this;
    }
    // check if is a shrinkable node
    // cannot be null, otherwise the hash would have been zero
    // std::cerr << "Normalizing " << size() << ":" << static_cast<uint64_t>(hash()) << " ? " << node->size() << std::endl;
    bool shrinkable = true;
    for (int i = 1; i < NODE_SIZE; i++) {

        if ((*node)[i].hash() != zero_hash(32)) {
            shrinkable = false;
            break;
        }
    }
    if (shrinkable) {
        return (*node)[0].normalize();
    }
    return *this;
}

void RawBuffer::serialize(std::vector<unsigned char>& value_vector) {
    // first check if it's empty
    // std::cerr << "NSerializing " << size() << ":" << static_cast<uint64_t>(hash()) << " ? " << saved << std::endl;
    if (hash() == zero_hash(32)) {
        value_vector.push_back(0);
        return;
    }
    // save leaf (just save all the data)
    if (level == 0) {
        value_vector.push_back(1);
        for (int i = 0; i < LEAF_SIZE; i++) {
            if (leaf->size() < i) value_vector.push_back(0);
            else value_vector.push_back((*leaf)[i]);
        }
    }
    if (level > 0) {
        value_vector.push_back(1);
        for (int i = 0; i < NODE_SIZE; i++) {
            (*node)[i].serialize(value_vector);
        }
    }
}

RawBuffer RawBuffer::deserialize(const char *buf, int level, int &len) {
    // empty
    if (buf[0] == 0) {
        len++;
        return RawBuffer(level, true);
    }
    // otherwise buf[0] == 1
    len++;
    buf++;
    if (level == 0) {
        auto res = std::make_shared<std::vector<uint8_t> >();
        res->resize(LEAF_SIZE, 0);
        for (unsigned int i = 0; i < LEAF_SIZE; i++) {
            (*res)[i] = buf[i];
        }
        len += LEAF_SIZE;
        return RawBuffer(res);
    }
    // node
    auto res = std::make_shared<std::vector<RawBuffer> >();
    for (unsigned int i = 0; i < NODE_SIZE; i++) {
        int nlen = 0;
        res->push_back(RawBuffer::deserialize(buf, level-1, nlen));
        // std::cerr << "deserlen " << i << ": " << nlen << std::endl;
        len += nlen;
        buf += nlen;
    }

    return RawBuffer(res, level);
}
