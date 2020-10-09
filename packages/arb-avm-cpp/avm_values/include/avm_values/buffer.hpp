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

#include <memory>
#include <vector>
#include <cstdint>
#include <iostream>
#include <avm_values/bigint.hpp>

inline uint64_t calc_len(int h) {
    if (h == 0) {
        return 1024;
    }
    return 128*calc_len(h-1);
}

struct Packed {
    uint256_t hash;
    int size; // total size
    int packed; // packed levels
};

class Buffer {
   private:
    bool is_leaf;
    int level;

    std::shared_ptr<std::vector<uint8_t> > leaf;
    std::shared_ptr<std::vector<Buffer> > node;

    Buffer(std::shared_ptr<std::vector<uint8_t> > leaf_) : leaf(leaf_), node(nullptr) {
        is_leaf = true;
        level = 0;
    }

    Buffer(std::shared_ptr<std::vector<Buffer> > node_, int level_) : leaf(nullptr), node(node_) {
        is_leaf = false;
        level = level_;
    }

    Buffer(int level_, bool) : leaf(nullptr), node(nullptr) {
        is_leaf = (level_ == 0);
        level = level_;
    }

   public:
    Buffer() : leaf(nullptr), node(nullptr) {
        // std::cerr << "creating buffer\n";
        is_leaf = true;
        level = 0;
    }

    Buffer set(uint64_t offset, uint8_t v) {
        // std::cerr << "setting buffer " << level << " at " << offset << " to " << std::hex << int(v) << std::endl;
        if (is_leaf) {
            if (offset >= 1024) {
                std::shared_ptr<std::vector<uint8_t> > empty = std::make_shared<std::vector<uint8_t>>();
                std::shared_ptr<std::vector<Buffer> > vec = std::make_shared<std::vector<Buffer>>();
                vec->push_back(Buffer(leaf));
                for (int i = 1; i < 128; i++) {
                    vec->push_back(Buffer(empty));
                }
                Buffer buf = Buffer(vec, 1);
                return buf.set(offset, v);
            }
            auto buf = leaf ? std::make_shared<std::vector<uint8_t> >(*leaf) : std::make_shared<std::vector<uint8_t> >();
            if (buf->size() < 1024) {
                // std::cerr << "resize buf" << std::endl;
                buf->resize(1024, 0);
            }
            (*buf)[offset] = v;
            // std::cerr << "updated leaf " << level << " at " << offset << " to " << std::hex << int(v) << std::endl;
            return Buffer(buf);
        } else {
            if (offset >= calc_len(level)) {
                std::shared_ptr<std::vector<Buffer> > vec = std::make_shared<std::vector<Buffer>>();
                vec->push_back(Buffer(node, level));
                for (int i = 1; i < 128; i++) {
                    vec->push_back(Buffer(level, true));
                }
                Buffer buf = Buffer(vec, level+1);
                return buf.set(offset, v);
            }
            auto vec = std::make_shared<std::vector<Buffer> >(node ? *node : Buffer::make_empty(level-1));
            auto cell_len = calc_len(level-1);
            (*vec)[offset / cell_len] = (*vec)[offset / cell_len].set(offset % cell_len, v);
            return Buffer(vec, level);
        }
    }

    Buffer set_many(uint64_t offset, std::vector<uint8_t> arr) {
        // std::cerr << "setting buffer " << level << " at " << offset << " to " << std::hex << int(v) << std::endl;
        if (is_leaf) {
            if (offset >= 1024) {
                std::shared_ptr<std::vector<uint8_t> > empty = std::make_shared<std::vector<uint8_t>>();
                std::shared_ptr<std::vector<Buffer> > vec = std::make_shared<std::vector<Buffer>>();
                vec->push_back(Buffer(leaf));
                for (int i = 1; i < 128; i++) {
                    vec->push_back(Buffer(empty));
                }
                Buffer buf = Buffer(vec, 1);
                return buf.set_many(offset, arr);
            }
            auto buf = leaf ? std::make_shared<std::vector<uint8_t> >(*leaf) : std::make_shared<std::vector<uint8_t> >();
            if (buf->size() < 1024) {
                // std::cerr << "resize buf" << std::endl;
                buf->resize(1024, 0);
            }
            for (unsigned int i = 0; i < arr.size(); i++) {
                (*buf)[offset+i] = arr[i];
            }
            // std::cerr << "updated leaf " << level << " at " << offset << " to " << std::hex << int(v) << std::endl;
            return Buffer(buf);
        } else {
            if (offset >= calc_len(level)) {
                std::shared_ptr<std::vector<Buffer> > vec = std::make_shared<std::vector<Buffer>>();
                vec->push_back(Buffer(node, level));
                for (int i = 1; i < 128; i++) {
                    vec->push_back(Buffer(level, true));
                }
                Buffer buf = Buffer(vec, level+1);
                return buf.set_many(offset, arr);
            }
            auto vec = std::make_shared<std::vector<Buffer> >(node ? *node : Buffer::make_empty(level-1));
            auto cell_len = calc_len(level-1);
            (*vec)[offset / cell_len] = (*vec)[offset / cell_len].set_many(offset % cell_len, arr);
            return Buffer(vec, level);
        }
    }

    static std::vector<Buffer> make_empty(int level) {
        auto vec = std::vector<Buffer>();
        for (int i = 0; i < 128; i++) {
            vec.push_back(Buffer(level, true));
        }
        return vec;
    }

    uint8_t get(uint64_t pos) const {
        if (is_leaf) {
            if (!leaf) return 0;
            if (leaf->size() <= pos) return 0;
            return (*leaf)[pos];
        } else {
            uint64_t len = calc_len(level);
            uint64_t cell_len = calc_len(level-1);
            if (pos > len || !node || (pos / cell_len) >= node->size()) {
                return 0;
            }
            return (*node)[pos / cell_len].get(pos % cell_len);
        }
    }

    std::vector<uint8_t> get_many(uint64_t pos, int len) const {
        if (is_leaf) {
            auto res = std::vector<uint8_t>(len, 0);
            for (int i = 0; i < len; i++) {
                if (!leaf) res[i] = 0;
                else if (leaf->size() < pos+i) res[i] = 0;
                else res[i] = (*leaf)[pos+i];
            }
            return res;
        } else {
            uint64_t ln = calc_len(level);
            uint64_t cell_len = calc_len(level-1);
            if (pos > ln || !node) {
                return std::vector<uint8_t>(len, 0);
            }
            return (*node)[pos / cell_len].get_many(pos % cell_len, len);
        }
    }

    uint256_t hash() const;
    Packed hash_aux() const;

    uint64_t size() const {
        return calc_len(level);
    }

};

uint256_t hash(const Buffer&);

inline bool operator==(const Buffer& val1, const Buffer& val2) {
    return hash(val1) == hash(val2);
}

#endif /* buffer_hpp */
