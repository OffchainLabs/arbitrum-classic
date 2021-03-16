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

#include <avm_values/bigint.hpp>
#include <avm_values/buffer.hpp>

#include <ethash/keccak.hpp>

uint256_t zero_hash(uint64_t sz) {
    if (sz == 5) {
        return hash(0);
    }
    auto h1 = zero_hash(sz - 1);
    return hash(h1, h1);
}

uint256_t hash2(uint256_t a, uint256_t b) {
    return hash(a, b);
}

Packed normal(uint256_t hash, uint64_t sz, uint64_t lastIndex) {
    return Packed{hash, sz, 0, lastIndex};
}

Packed pack(const Packed& packed) {
    return Packed{packed.hash, packed.size, packed.packed + 1,
                  packed.lastIndex};
}

bool is_zero_hash(const Packed& packed) {
    return packed.hash == hash(0);
}

uint256_t unpack(const Packed& packed) {
    uint256_t res = packed.hash;
    uint64_t sz = packed.size;
    for (int i = 0; i < packed.packed; i++) {
        res = hash(res, zero_hash(sz));
        sz = sz + 1;
    }
    return res;
}

Packed zero_packed(uint64_t sz) {
    if (sz == 5) {
        return normal(hash(0), 5, 0);
    }
    return pack(zero_packed(sz - 1));
}

Packed hash_buf(uint8_t* buf, uint64_t offset, uint64_t sz) {
    if (sz == 5) {
        auto hash_val = ethash::keccak256(buf + offset, 32);
        uint256_t res = intx::be::load<uint256_t>(hash_val);
        uint64_t lastIndex = 31;
        while (buf[offset + lastIndex] == 0 && lastIndex > 0)
            lastIndex--;
        return normal(res, 5, lastIndex);
    }
    auto h1 = hash_buf(buf, offset, sz - 1);
    auto h2 = hash_buf(buf, offset + (1 << (sz - 1)), sz - 1);
    if (is_zero_hash(h2)) {
        return pack(h1);
    }
    return normal(hash(unpack(h1), unpack(h2)), sz,
                  h2.lastIndex + (1 << (sz - 1)));
}

Packed hash_node(RawBuffer* buf, uint64_t offset, uint64_t len, uint64_t sz) {
    if (len == 1) {
        auto res = buf[offset].hash_aux();
        return res;
    }
    auto h1 = hash_node(buf, offset, len / 2, sz - 1);
    auto h2 = hash_node(buf, offset + len / 2, len / 2, sz - 1);
    if (is_zero_hash(h2)) {
        return pack(h1);
    }
    return normal(hash(unpack(h1), unpack(h2)), sz,
                  h2.lastIndex + (1 << (sz - 1)));
}

Packed RawBuffer::hash_aux() const {
    if (saved) {
        return savedHash;
    }
    Packed res;
    if (level == 0) {
        if (!leaf || leaf->size() == 0)
            res = zero_packed(LEAF_SIZE2);
        else
            res = hash_buf(leaf->data(), 0, LEAF_SIZE2);
    } else {
        if (!node)
            res = zero_packed(calc_height(level));
        else {
            res = hash_node(node->data(), 0, NODE_SIZE, calc_height(level));
        }
    }
    saved = true;
    savedHash = res;
    return res;
}

RawBuffer RawBuffer::normalize() const {
    if (hash() == zero_hash(5)) {
        return RawBuffer();
    }
    if (level == 0) {
        return *this;
    }
    // check if is a shrinkable node
    // cannot be null, otherwise the hash would have been zero
    bool shrinkable = true;
    for (uint64_t i = 1; i < NODE_SIZE; i++) {
        if ((*node)[i].hash() != zero_hash(5)) {
            shrinkable = false;
            break;
        }
    }
    if (shrinkable) {
        return (*node)[0].normalize();
    }
    return *this;
}

std::vector<RawBuffer> RawBuffer::serialize(
    std::vector<unsigned char>& value_vector) {
    // first check if it's empty
    std::vector<RawBuffer> ret{};
    if (hash() == zero_hash(5)) {
        value_vector.push_back(0);
        return ret;
    }
    // save leaf (just save all the data)
    if (level == 0) {
        value_vector.push_back(1);
        for (uint64_t i = 0; i < LEAF_SIZE; i++) {
            if (leaf->size() <= i)
                value_vector.push_back(0);
            else
                value_vector.push_back((*leaf)[i]);
        }
    }

    if (level > 0) {
        value_vector.push_back(1);
        for (uint64_t i = 0; i < NODE_SIZE; i++) {
            uint256_t hash_ = hash2(123, (*node)[i].hash());
            marshal_uint256_t(hash_, value_vector);
            ret.push_back((*node)[i]);
        }
    }
    return ret;
}

uint64_t RawBuffer::sizePow2() const {
    uint64_t size = 0;
    if (level == 0 && leaf && leaf->size() > 0) {
        for (int i = LEAF_SIZE - 1; i >= 0; i--) {
            if ((*leaf)[i] != 0) {
                size = i;
                break;
            }
        }
    } else if (node && node->size() > 0) {
        for (int i = NODE_SIZE - 1; i >= 0; i--) {
            if ((*node)[i].hash() != zero_hash(5)) {
                size = i * calc_len(level - 1) - 1 + calc_len(level - 1);
                break;
            }
        }
    }
    uint64_t size_ext = needed_height(size);
    if (size_ext < 5)
        size_ext = 5;
    return size_ext;
}

std::vector<unsigned char> RawBuffer::makeProof(uint64_t offset,
                                                uint64_t sz,
                                                uint64_t loc) {
    if (sz == 5) {
        if (!leaf || leaf->size() == 0) {
            return std::vector<unsigned char>(32, 0);
        }
        auto res = std::vector<unsigned char>(leaf->begin() + loc,
                                              leaf->begin() + loc + 32);
        return res;
    } else if (level > 0 && sz == calc_height(level - 1) && node) {
        return (*node)[offset / calc_len(level - 1)].makeProof(
            offset % calc_len(level - 1), sz, loc % calc_len(level - 1));
    } else if (loc < offset + (1L << (sz - 1))) {
        auto proof = makeProof(offset, sz - 1, loc);
        marshal_uint256_t(merkleHash(offset + (1L << (sz - 1)), sz - 1), proof);
        return proof;
    } else {
        auto proof = makeProof(offset + (1L << (sz - 1)), sz - 1, loc);
        marshal_uint256_t(merkleHash(offset, sz - 1), proof);
        return proof;
    }
}

uint256_t RawBuffer::merkleHash(uint64_t offset, uint64_t sz) {
    if (hash() == zero_hash(5)) {
        return zero_hash(sz);
    }
    if (sz == 5) {
        auto hash_val = ethash::keccak256(leaf->data() + offset, 32);
        uint256_t res = intx::be::load<uint256_t>(hash_val);
        return res;
    } else if (level > 0 && sz == calc_height(level - 1) && node) {
        return (*node)[offset / calc_len(level - 1)].merkleHash(0, sz);
    }
    auto h1 = merkleHash(offset, sz - 1);
    auto h2 = merkleHash(offset + (1L << (sz - 1)), sz - 1);
    return hash2(h1, h2);
}

std::vector<unsigned char> RawBuffer::makeProof(uint64_t loc) {
    auto size = sizePow2();
    auto res = makeProof(0, size, ((loc / 32) % (1L << (size - 5))) * 32);
    return res;
}

std::vector<unsigned char> RawBuffer::makeNormalizationProof() {
    uint64_t sz = sizePow2();
    std::vector<unsigned char> res;
    for (int i = 0; i < 31; i++) {
        res.push_back(0);
    }

    if (sz == 5) {
        res.push_back(0);
        marshal_uint256_t(merkleHash(0, sz), res);
        marshal_uint256_t(merkleHash(0, sz), res);
        return res;
    }

    res.push_back(makeProof(0, sz, 0).size() / 32);
    marshal_uint256_t(merkleHash(0, sz - 1), res);
    marshal_uint256_t(merkleHash(1L << (sz - 1), sz - 1), res);
    return res;
}
