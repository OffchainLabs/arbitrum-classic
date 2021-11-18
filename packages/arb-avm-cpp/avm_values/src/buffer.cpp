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

// Returns the length of a buffer with a given depth
inline uint256_t length_of_depth(uint64_t depth) {
    return uint256_t(RawBuffer::leaf_size) << depth;
}

// Returns the necessary depth of a buffer to hold a given number of bytes
inline uint64_t needed_depth(uint256_t size) {
    uint64_t depth = 0;
    while (size > RawBuffer::leaf_size) {
        // Divide rounding up
        size = (size + 1) / 2;
        depth += 1;
    }
    return depth;
}

const std::vector<Buffer> zero_buffers_of_depth = []() -> std::vector<Buffer> {
    std::vector<Buffer> buffers;
    auto last_buffer = Buffer();
    for (int i = 0; i < 64; i++) {
        buffers.push_back(last_buffer);
        last_buffer = Buffer(last_buffer, last_buffer);
    }
    return buffers;
}();

RawBuffer::RawBuffer(LeafData bytes) : depth(0), components(bytes) {}

RawBuffer::RawBuffer(Buffer left, Buffer right)
    : depth(left->depth + 1),
      components(std::make_pair(std::move(left), std::move(right))) {
    auto children = get_children_const();
    if (children->first->depth != children->second->depth) {
        throw new std::runtime_error("Attempted to create uneven buffer");
    }
}

RawBuffer::NodeData* RawBuffer::get_children() {
    return std::get_if<NodeData>(&components);
}

const RawBuffer::NodeData* RawBuffer::get_children_const() const {
    return std::get_if<NodeData>(&components);
}

uint256_t RawBuffer::hash() const {
    auto current_hash_cache = hash_cache.val.load();
    if (current_hash_cache) {
        return *current_hash_cache;
    }
    uint256_t calculated_hash;
    if (auto children = get_children_const()) {
        calculated_hash =
            ::hash(children->first->hash(), children->second->hash());
    } else {
        auto& bytes = std::get<LeafData>(components);
        calculated_hash = ::hash(bytes);
    }
    hash_cache.val.store(calculated_hash);
    return calculated_hash;
}

uint256_t RawBuffer::packed_size() const {
    auto current_packed = packed_size_cache.val.load();
    if (current_packed) {
        return *current_packed;
    }
    uint256_t calculated_packed_size = 0;
    if (auto children = get_children_const()) {
        auto first_packed_size = children->first->packed_size();
        auto second_packed_size = children->second->packed_size();
        if (second_packed_size == 0) {
            // Ignore the second half as it's all zeroes
            calculated_packed_size = first_packed_size;
        } else {
            // Ignore any trailing zeroes in the left half of the buffer, as
            // they're followed by non-zero data in the right half of the buffer
            calculated_packed_size =
                length_of_depth(children->first->depth) + second_packed_size;
        }
    } else {
        auto& bytes = std::get<LeafData>(components);
        // Go backwards through the bytes to find the last non-zero byte
        calculated_packed_size = leaf_size;
        while (calculated_packed_size > 0) {
            if (bytes[size_t(calculated_packed_size - 1)] != 0) {
                break;
            }
            calculated_packed_size -= 1;
        }
    }
    packed_size_cache.val.store(calculated_packed_size);
    return calculated_packed_size;
}

RawBuffer RawBuffer::grow(uint64_t new_depth) const {
    RawBuffer ret(*this);
    while (ret.depth < new_depth) {
        ret = RawBuffer(Buffer(ret), zero_buffers_of_depth[ret.depth]);
    }
    return ret;
}

RawBuffer RawBuffer::trim() const {
    RawBuffer ret(*this);
    while (true) {
        if (auto children = ret.get_children()) {
            if (children->second->packed_size() == 0) {
                ret = *children->first;
                continue;
            }
        }
        break;
    }
    return ret;
}

RawBuffer RawBuffer::set_many_without_resize(uint64_t offset,
                                             const std::vector<uint8_t>& arr,
                                             uint64_t arr_offset,
                                             uint64_t arr_length) const {
    RawBuffer ret(*this);
    RawBuffer* target = &ret;
    while (true) {
        target->hash_cache.val.store(std::nullopt);
        target->packed_size_cache.val.store(std::nullopt);
        if (auto children = target->get_children()) {
            // Clone each buffer on our way down, and adjust the offset.
            auto child_size = children->first->size();
            if (offset >= child_size) {
                offset -= uint64_t(child_size);
                children->second = Buffer(*children->second);
                target = &*children->second;
            } else {
                children->first = Buffer(*children->first);
                target = &*children->first;
            }
        } else {
            // We've found the target leaf
            auto& bytes = std::get<LeafData>(target->components);
            if (offset >= leaf_size) {
                throw new std::runtime_error(
                    "Buffer set_many_without_resize called but resize needed");
            } else if (offset + arr_length > leaf_size) {
                throw new std::runtime_error(
                    "Buffer set_many called with misaligned bytes");
            }
            auto output = bytes.begin() + offset;
            auto start = arr.begin() + arr_offset;
            auto end = start + arr_length;
            std::copy(start, end, output);
            return ret;
        }
    }
}

RawBuffer::RawBuffer() : RawBuffer(LeafData{}) {}

RawBuffer RawBuffer::fromData(const std::vector<uint8_t>& data) {
    std::vector<RawBuffer> leftwards_buffers;
    size_t i = 0;
    while (i < data.size()) {
        LeafData new_leaf{};
        size_t len = RawBuffer::leaf_size;
        if (i + len > data.size()) {
            len = data.size() - i;
        }
        std::copy(data.begin() + i, data.begin() + i + len, new_leaf.begin());
        i += len;
        // Add the new leaf to leftwards_buffers, collapsing where possible
        RawBuffer inserting(new_leaf);
        while (!leftwards_buffers.empty() &&
               leftwards_buffers.back().depth == inserting.depth) {
            inserting = RawBuffer(Buffer(std::move(leftwards_buffers.back())),
                                  Buffer(inserting));
            leftwards_buffers.pop_back();
        }
        leftwards_buffers.push_back(inserting);
    }
    while (leftwards_buffers.size() > 1) {
        // Forcibly collapse the last two buffers by adding padding to the
        // rightmost one. This is more efficient than inserting more zero leafs,
        // as this uses zero_buffers_of_depth.
        RawBuffer right = std::move(leftwards_buffers.back());
        leftwards_buffers.pop_back();
        RawBuffer& left = leftwards_buffers.back();
        right = right.grow(left.depth);
        left = RawBuffer(Buffer(std::move(left)), Buffer(std::move(right)));
    }
    if (leftwards_buffers.empty()) {
        return RawBuffer();
    }
    return std::move(leftwards_buffers[0]).trim();
}

uint256_t RawBuffer::size() const {
    return length_of_depth(depth);
}

uint256_t Buffer::size() const {
    return (*this)->size();
}

uint64_t Buffer::lastIndex() const {
    auto packed_size = (*this)->packed_size();
    if (packed_size == 0) {
        return 0;
    } else {
        auto ret = packed_size - 1;
        if (ret > std::numeric_limits<uint64_t>::max()) {
            throw new std::runtime_error("Buffer is too big");
        }
        return uint64_t(ret);
    }
}

uint256_t Buffer::data_length() const {
    return (*this)->packed_size();
}

RawBuffer RawBuffer::set_many(uint64_t offset,
                              const std::vector<uint8_t>& arr) const {
    RawBuffer ret(*this);
    if (uint256_t(offset) + arr.size() > ret.size()) {
        ret = ret.grow(needed_depth(uint256_t(offset) + arr.size()));
    }
    ret = ret.set_many_without_resize(offset, arr, 0, arr.size());
    return ret.trim();
}

Buffer Buffer::set(uint64_t offset, uint8_t v) const {
    return set_many(offset, std::vector(1, v));
}

std::vector<uint8_t> Buffer::get_many(uint64_t offset, size_t len) const {
    const RawBuffer* target = std::shared_ptr<RawBuffer>::get();
    while (true) {
        if (auto children = target->get_children_const()) {
            // Move downwards towards the target
            auto child_size = children->first->size();
            if (offset >= child_size) {
                offset -= uint64_t(child_size);
                target = &*children->second;
            } else {
                target = &*children->first;
            }
        } else {
            // We've found the target leaf
            auto& bytes = std::get<RawBuffer::LeafData>(target->components);
            if (offset >= RawBuffer::leaf_size) {
                return std::vector<uint8_t>(len, (unsigned char)0);
            } else if (offset + len > RawBuffer::leaf_size) {
                throw new std::runtime_error(
                    "Buffer get_many called with misaligned bytes");
            }
            auto start = bytes.begin() + offset;
            return std::vector<uint8_t>(start, start + len);
        }
    }
}

uint8_t Buffer::get(uint64_t offset) const {
    return get_many(offset, 1)[0];
}

std::vector<uint8_t> Buffer::toFlatVector() const {
    std::vector<uint8_t> ret;
    ret.reserve(uint64_t(size()));  // If this overflows we have bigger problems
    std::vector<const RawBuffer*> to_visit;
    to_visit.reserve((*this)->depth);
    const RawBuffer* current = std::shared_ptr<RawBuffer>::get();
    // Visit the tree depth first
    while (true) {
        if (auto children = current->get_children_const()) {
            // Visit the left side of the buffer now, and save the right side
            // for later
            current = &*children->first;
            to_visit.push_back(&*children->second);
        } else {
            const auto& bytes =
                std::get<RawBuffer::LeafData>(current->components);
            std::copy(bytes.begin(), bytes.end(), std::back_inserter(ret));
            if (to_visit.empty()) {
                // We've visited all the leaves
                // Trim the result and return it
                while (!ret.empty() && ret.back() == 0) {
                    ret.pop_back();
                }
                return ret;
            } else {
                current = to_visit.back();
                to_visit.pop_back();
            }
        }
    }
}

std::vector<unsigned char> Buffer::makeProof(uint64_t loc) const {
    if (loc >= size()) {
        // If we're trying to prove an element outside the buffer, we instead
        // need to prove the buffer's size, which we do by proving this element
        // instead. Proving this element specifically keeps compatiblity with
        // the Solidity visitor, as it looks at each bit of the location to
        // determine if it's on the left or right branch.
        loc %= uint64_t(size());
    }
    // Return a standard merkle proof
    const RawBuffer* target = std::shared_ptr<RawBuffer>::get();
    std::vector<uint256_t> proof;
    while (true) {
        if (auto children = target->get_children_const()) {
            // Move downwards towards the target
            // Add the sibling hash to the proof each step of the way
            auto child_size = children->first->size();
            if (loc >= child_size) {
                loc -= uint64_t(child_size);
                target = &*children->second;
                proof.push_back(children->first->hash());
            } else {
                target = &*children->first;
                proof.push_back(children->second->hash());
            }
        } else {
            // We've found the target leaf
            auto& bytes = std::get<RawBuffer::LeafData>(target->components);
            proof.push_back(intx::be::unsafe::load<uint256_t>(bytes.data()));
            break;
        }
    }
    std::vector<unsigned char> proof_bytes;
    for (auto it = proof.rbegin(); it != proof.rend(); it++) {
        unsigned char bytes[32]{0};
        intx::be::store(bytes, *it);
        std::copy(bytes, bytes + 32, std::back_inserter(proof_bytes));
    }
    return proof_bytes;
}

std::vector<unsigned char> Buffer::makeNormalizationProof() const {
    // Return the height, left subtree hash (or our hash if a leaf), and right
    // subtree hash (irrelevant if a leaf)
    std::vector<unsigned char> proof_bytes;
    proof_bytes.resize(32);
    intx::be::unsafe::store(proof_bytes.data(), uint256_t((*this)->depth));
    unsigned char left_hash[32]{0};
    unsigned char right_hash[32]{0};
    if (auto children = (*this)->get_children_const()) {
        intx::be::store(left_hash, children->first->hash());
        intx::be::store(right_hash, children->second->hash());
    } else {
        intx::be::store(left_hash, hash());
        // right_hash is ignored here
    }
    std::copy(left_hash, left_hash + 32, std::back_inserter(proof_bytes));
    std::copy(right_hash, right_hash + 32, std::back_inserter(proof_bytes));
    return proof_bytes;
}

std::vector<Buffer> Buffer::serialize(
    std::vector<unsigned char>& value_vector) const {
    // first check if it's empty
    std::vector<Buffer> ret{};
    value_vector.push_back((*this)->depth);
    if (auto children = (*this)->get_children_const()) {
        marshal_uint256_t(::hash(*children->first), value_vector);
        marshal_uint256_t(::hash(*children->second), value_vector);
        ret.push_back(children->first);
        ret.push_back(children->second);
    } else {
        auto& bytes = std::get<RawBuffer::LeafData>((*this)->components);
        std::copy(bytes.begin(), bytes.end(), std::back_inserter(value_vector));
    }
    return ret;
}

Buffer::Buffer(RawBuffer raw)
    : std::shared_ptr<RawBuffer>(std::make_shared<RawBuffer>(std::move(raw))) {}

Buffer::Buffer() : Buffer(RawBuffer()) {}

Buffer::Buffer(std::array<unsigned char, 32> bytes)
    : Buffer(RawBuffer(bytes)) {}

Buffer::Buffer(Buffer left, Buffer right)
    : Buffer(RawBuffer(std::move(left), std::move(right))) {}

Buffer Buffer::fromData(const std::vector<uint8_t>& data) {
    return Buffer(RawBuffer::fromData(data));
}

uint256_t Buffer::hash() const {
    return (*this)->hash();
}

Buffer Buffer::set_many(uint64_t offset,
                        const std::vector<uint8_t>& arr) const {
    return Buffer((*this)->set_many(offset, arr));
}
