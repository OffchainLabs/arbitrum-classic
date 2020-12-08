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

uint256_t zero_hash(uint64_t sz) {
    if (sz == 32) {
        return hash(0);
    }
    auto h1 = zero_hash(sz/2);
    return hash(h1, h1);
}

uint256_t hash2(uint256_t a, uint256_t b) {
    return hash(a, b);
}

Packed normal(uint256_t hash, uint64_t sz) {
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
    uint64_t sz = packed.size;
    for (uint64_t i = 0; i < packed.packed; i++) {
        res = hash(res, zero_hash(sz));
        sz = sz*2;
    }
    return res;
}

Packed zero_packed(uint64_t sz) {
    if (sz == 32) {
        return normal(zero_hash(32), 32);
    }
    return pack(zero_packed(sz/2));
}

Packed hash_buf(uint8_t *buf, uint64_t offset, uint64_t sz) {
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

Packed hash_node(RawBuffer *buf, uint64_t offset, uint64_t len, uint64_t sz) {
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
    for (uint64_t i = 1; i < NODE_SIZE; i++) {

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

std::vector<RawBuffer> RawBuffer::serialize(std::vector<unsigned char>& value_vector) {
    // first check if it's empty
    std::cerr << "NSerializing " << size() << ":" << static_cast<uint64_t>(hash()) << " ? " << saved << std::endl;
    std::vector<RawBuffer> ret{};
    if (hash() == zero_hash(32)) {
        value_vector.push_back(0);
        return ret;
    }
    // save leaf (just save all the data)
    if (level == 0) {
        value_vector.push_back(1);
        for (uint64_t i = 0; i < LEAF_SIZE; i++) {
            if (leaf->size() < i) value_vector.push_back(0);
            else value_vector.push_back((*leaf)[i]);
        }
    }

    if (level > 0) {
        value_vector.push_back(1);
        for (uint64_t i = 0; i < NODE_SIZE; i++) {
            marshal_uint256_t((*node)[i].hash(), value_vector);
            // (*node)[i].serialize(value_vector);
            ret.push_back((*node)[i]);
        }
    }
    return ret;
}

uint256_t deserializeHash(const char*& bufptr) {
    auto ret = intx::be::unsafe::load<uint256_t>(
        reinterpret_cast<const unsigned char*>(bufptr));
    bufptr += 32;
    return ret;
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
        for (uint64_t i = 0; i < LEAF_SIZE; i++) {
            (*res)[i] = buf[i];
        }
        len += LEAF_SIZE;
        return RawBuffer(res);
    }
    // node
    auto res = std::vector<uint256_t>();
    for (uint64_t i = 0; i < NODE_SIZE; i++) {
        int nlen = 32;
        uint256_t hash = deserializeHash(buf);
        res.push_back(hash);
        /*
        int nlen = 0;
        res->push_back(RawBuffer::deserialize(buf, level-1, nlen));
        // std::cerr << "deserlen " << i << ": " << nlen << std::endl;
        */
        len += nlen;
        buf += nlen;
    }

    return RawBuffer(res, level);
}

uint64_t RawBuffer::sizePow2() const {
    uint64_t size = 0;
    if (level == 0 && leaf && leaf->size() > 0) {
        // std::cerr << "check size leaf" << std::endl;
        for (int i = LEAF_SIZE-1; i >= 0; i--) {
            if ((*leaf)[i] != 0) {
                size = i;
                break;
            }
        }
    }
    else if (node && node->size() > 0) {
        // std::cerr << "check size node" << std::endl;
        for (int i = NODE_SIZE-1; i >= 0; i--) {
            if ((*node)[i].hash() != zero_hash(32)) {
                size = (i+1)*calc_len(level-1);
                break;
            }
        }
    }
    uint64_t size_ext = calc_len(level);
    if (size_ext < 32) size_ext = 32;
    while (size_ext/2 >= size && size_ext > 32) {
        size_ext = size_ext/2;
    }
    return size_ext;
}

std::vector<unsigned char> RawBuffer::makeProof(uint64_t offset, uint64_t sz, uint64_t loc) {
    // std::cerr << "makeProof " << offset << " sz " << sz << " loc " << loc << std::endl;
    if (sz == 32) {
        if (!leaf || leaf->size() == 0) {
            return std::vector<unsigned char>(32, 0);
        }
        auto res = std::vector<unsigned char>(leaf->begin()+loc, leaf->begin()+loc+32);
        return res;
    } else if (level > 0 && sz == calc_len(level-1) && node) {
        return (*node)[offset/calc_len(level-1)].makeProof(offset%calc_len(level-1), sz, loc%calc_len(level-1));
    } else if (loc < offset + sz/2) {
        auto proof = makeProof(offset, sz/2, loc);
        marshal_uint256_t(merkleHash(offset+sz/2, sz/2), proof);
        return proof;
    } else {
        auto proof = makeProof(offset+sz/2, sz/2, loc);
        marshal_uint256_t(merkleHash(offset, sz/2), proof);
        return proof;
    }
}

uint256_t RawBuffer::merkleHash(uint64_t offset, uint64_t sz) {
    // std::cerr << "merkle hash " << offset << " sz " << sz << std::endl;
    if (hash() == zero_hash(32)) {
        return zero_hash(sz);
    }
    if (sz == 32) {
        auto hash_val = ethash::keccak256(leaf->data()+offset, 32);
        uint256_t res = intx::be::load<uint256_t>(hash_val);
        return res;
    } else if (level > 0 && sz == calc_len(level-1) && node) {
        return (*node)[offset/calc_len(level-1)].merkleHash(0, sz);
    }
    // std::cerr << "hashing " << offset << " to " << (offset+sz) << std::endl;
    auto h1 = merkleHash(offset, sz/2);
    auto h2 = merkleHash(offset+sz/2, sz/2);
    return hash2(h1, h2);
}

std::vector<unsigned char> RawBuffer::makeProof(uint64_t loc) {
    auto size = sizePow2();
    // std::cerr << "Got size " << size << std::endl;
    auto res = makeProof(0, size, ((loc/32) % (size/32))*32);
    // std::cerr << "Making " << size << " -- " << res.size()/32 << std::endl;
    return res;
}

std::vector<unsigned char> RawBuffer::makeNormalizationProof() {
    uint64_t sz = sizePow2();
    std::vector<unsigned char> res;
    for (int i = 0; i < 31; i++) {
        res.push_back(0);
    }

    if (sz == 32) {
        // std::cerr << "Simple normalization" << std::endl;
        res.push_back(0);
        marshal_uint256_t(merkleHash(0, sz), res);
        marshal_uint256_t(merkleHash(0, sz), res);
        return res;
    }

    res.push_back(makeProof(0, sz, 0).size()/32);
    marshal_uint256_t(merkleHash(0, sz/2), res);
    marshal_uint256_t(merkleHash(sz/2, sz/2), res);
    return res;
}

