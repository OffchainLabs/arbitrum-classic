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

#ifndef buffer_hpp
#define buffer_hpp

#include <avm_values/bigint.hpp>
#include <cstdint>
#include <iostream>
#include <memory>
#include <vector>

const uint64_t LEAF_SIZE2 = 10;
const uint64_t NODE_SIZE2 = 3;
const uint64_t LEAF_SIZE = 1 << LEAF_SIZE2;
const uint64_t NODE_SIZE = 1 << NODE_SIZE2;
const uint64_t ALIGN = LEAF_SIZE;

inline uint64_t calc_len(int h) {
    if (h == 0) {
        return LEAF_SIZE;
    }
    return NODE_SIZE * calc_len(h - 1);
}

inline uint64_t calc_height(int h) {
    if (h == 0) {
        return LEAF_SIZE2;
    }
    return NODE_SIZE2 + calc_height(h - 1);
}

inline uint64_t needed_height(uint64_t offset) {
    if (offset <= 1) {
        return 1;
    } else {
        return 1 + needed_height(offset / 2);
    }
}
struct Packed {
    uint256_t hash;
    uint64_t size;  // total height
    int packed;     // packed levels
    uint64_t lastIndex;
};

Packed zero_packed(uint64_t sz);

class RawBuffer {
   private:
    bool saved;
    Packed savedHash;

    std::shared_ptr<std::vector<uint8_t>> leaf;
    std::shared_ptr<std::vector<RawBuffer>> node;

   public:
    int level;

    RawBuffer(std::shared_ptr<std::vector<RawBuffer>> node_, int level_)
        : leaf(nullptr), node(node_) {
        level = level_;
        saved = false;
    }

    RawBuffer(std::shared_ptr<std::vector<uint8_t>> leaf_)
        : leaf(leaf_), node(nullptr) {
        level = 0;
        saved = false;
    }

    RawBuffer(int level_, bool) : leaf(nullptr), node(nullptr) {
        level = level_;
        saved = true;
        savedHash = zero_packed(calc_height(level));
    }

    RawBuffer() : leaf(nullptr), node(nullptr) {
        level = 0;
        saved = true;
        savedHash = zero_packed(LEAF_SIZE2);
    }

    RawBuffer set(uint64_t offset, uint8_t v) {
        std::vector<uint8_t> arr(1);
        arr[0] = v;
        return set_many(offset, arr);
    }

    // Note: pos and len must be aligned so that the data to be written is in
    // one leaf
    RawBuffer set_many(uint64_t offset, std::vector<uint8_t>& arr) {
        if (level == 0) {
            if (offset >= LEAF_SIZE) {
                std::shared_ptr<std::vector<uint8_t>> empty =
                    std::make_shared<std::vector<uint8_t>>();
                std::shared_ptr<std::vector<RawBuffer>> vec =
                    std::make_shared<std::vector<RawBuffer>>();
                vec->push_back(RawBuffer(leaf));
                for (uint64_t i = 1; i < NODE_SIZE; i++) {
                    vec->push_back(RawBuffer(empty));
                }
                RawBuffer buf = RawBuffer(vec, 1);
                return buf.set_many(offset, arr);
            }
            auto buf = leaf ? std::make_shared<std::vector<uint8_t>>(*leaf)
                            : std::make_shared<std::vector<uint8_t>>();
            if (buf->size() < LEAF_SIZE) {
                buf->resize(LEAF_SIZE, 0);
            }
            for (unsigned int i = 0; i < arr.size(); i++) {
                (*buf)[offset + i] = arr[i];
            }
            return RawBuffer(buf);
        } else {
            if (needed_height(offset) > calc_height(level)) {
                std::shared_ptr<std::vector<RawBuffer>> vec =
                    std::make_shared<std::vector<RawBuffer>>();
                vec->push_back(RawBuffer(node, level));
                for (uint64_t i = 1; i < NODE_SIZE; i++) {
                    vec->push_back(RawBuffer(level, true));
                }
                RawBuffer buf = RawBuffer(vec, level + 1);
                return buf.set_many(offset, arr);
            }
            auto vec = std::make_shared<std::vector<RawBuffer>>(
                node ? *node : RawBuffer::make_empty(level - 1));
            auto cell_len = calc_len(level - 1);
            (*vec)[offset / cell_len] =
                (*vec)[offset / cell_len].set_many(offset % cell_len, arr);
            return RawBuffer(vec, level);
        }
    }

    static std::vector<RawBuffer> make_empty(int level) {
        auto vec = std::vector<RawBuffer>();
        for (uint64_t i = 0; i < NODE_SIZE; i++) {
            vec.push_back(RawBuffer(level, true));
        }
        return vec;
    }

    uint8_t get(uint64_t pos) const {
        auto res = get_many(pos, 1);
        return res[0];
    }

    // Note: pos and len must be aligned so that the data to be read is from one
    // leaf
    std::vector<uint8_t> get_many(uint64_t pos, int len) const {
        if (level == 0) {
            auto res = std::vector<uint8_t>(len, 0);
            for (int i = 0; i < len; i++) {
                if (!leaf) {
                    res[i] = 0;
                } else if (leaf->size() <= pos + i) {
                    res[i] = 0;
                } else {
                    res[i] = (*leaf)[pos + i];
                }
            }
            return res;
        } else {
            uint64_t cell_len = calc_len(level - 1);
            if (needed_height(pos) > calc_height(level) || !node) {
                return std::vector<uint8_t>(len, 0);
            }
            auto next = (*node)[pos / cell_len];
            return next.get_many(pos % cell_len, len);
        }
    }

    Packed hash_aux();
    uint256_t hash() { return hash_aux().hash; }

    uint64_t lastIndex() { return hash_aux().lastIndex; }

    std::vector<RawBuffer> serialize(std::vector<unsigned char>& value_vector);

    RawBuffer normalize();

    uint64_t size() const { return calc_len(level); }

    uint64_t sizePow2() const;

    std::vector<unsigned char> makeProof(uint64_t offset,
                                         uint64_t sz,
                                         uint64_t loc);
    uint256_t merkleHash(uint64_t offset, uint64_t sz);

    std::vector<unsigned char> makeProof(uint64_t loc);
    std::vector<unsigned char> makeNormalizationProof();

    friend class Buffer;
};

class Buffer {
   public:
    std::shared_ptr<RawBuffer> buf;

    Buffer(const RawBuffer& buffer) {
        buf = std::make_shared<RawBuffer>(buffer);
    }

    Buffer() { buf = std::make_shared<RawBuffer>(); }

    Buffer(const std::vector<uint8_t>& data) : Buffer() {
        for (uint64_t i = 0; i < data.size(); i++) {
            buf = std::make_shared<RawBuffer>(buf->set(i, data[i]));
        }
    }

    Buffer set(uint64_t offset, uint8_t v) {
        return Buffer(buf->set(offset, v));
    }

    Buffer set_many(uint64_t offset, std::vector<uint8_t> arr) {
        return Buffer(buf->set_many(offset, arr));
    }

    uint8_t get(uint64_t pos) const { return buf->get(pos); }

    std::vector<uint8_t> get_many(uint64_t pos, int len) const {
        return buf->get_many(pos, len);
    }

    uint64_t size() const { return buf->size(); }

    uint64_t lastIndex() const { return buf->lastIndex(); }

    uint256_t hash() const { return buf->hash(); }

    std::vector<unsigned char> makeProof(uint64_t loc) {
        RawBuffer nbuf = buf->normalize();
        return nbuf.makeProof(loc);
    }

    std::vector<unsigned char> makeNormalizationProof() {
        RawBuffer nbuf = buf->normalize();
        return nbuf.makeNormalizationProof();
    }

    std::vector<RawBuffer> serialize(
        std::vector<unsigned char>& value_vector) const {
        RawBuffer nbuf = buf->normalize();
        value_vector.push_back(static_cast<uint8_t>(nbuf.level));
        return nbuf.serialize(value_vector);
    }

    std::vector<uint8_t> toFlatVector() const {
        std::vector<uint8_t> data;
        uint64_t last_index = lastIndex();
        uint64_t i = 0;
        while (true) {
            data.push_back(get(i));
            if (i == last_index) {
                return data;
            }
            i++;
        }
    }
};

inline uint256_t hash(const Buffer& b) {
    return hash(123, b.hash());
}

inline bool operator==(const Buffer& val1, const Buffer& val2) {
    return val1.hash() == val2.hash();
}

#endif /* buffer_hpp */
