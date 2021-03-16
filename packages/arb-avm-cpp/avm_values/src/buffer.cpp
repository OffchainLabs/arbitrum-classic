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
inline uint64_t length_of_depth(uint64_t depth) {
    return Buffer::leaf_size << depth;
}

// Returns the necessary depth of a buffer to hold a given number of bytes
inline uint64_t needed_depth(uint64_t size) {
    uint64_t depth = 0;
    while (size > Buffer::leaf_size) {
        // Divide rounding up
        size = (size + 1) / 2;
        depth += 1;
    }
    return depth;
}

const std::vector<std::shared_ptr<Buffer>> zero_buffers_of_depth =
    []() -> std::vector<std::shared_ptr<Buffer>> {
    std::vector<std::shared_ptr<Buffer>> buffers;
    auto last_buffer = std::make_shared<Buffer>();
    for (int i = 0; i < 50; i++) {
        buffers.push_back(last_buffer);
        last_buffer = std::make_shared<Buffer>(last_buffer, last_buffer);
    }
    return buffers;
}();

Buffer::Buffer(LeafData bytes) : components(bytes) {
    recompute();
}

Buffer::Buffer(std::shared_ptr<Buffer> left, std::shared_ptr<Buffer> right)
    : components(std::make_pair(left, right)) {
    recompute();
}

Buffer::NodeData* Buffer::get_children() {
    return std::get_if<NodeData>(&components);
}

const Buffer::NodeData* Buffer::get_children_const() const {
    return std::get_if<NodeData>(&components);
}

void Buffer::recompute() {
    if (auto children = get_children()) {
        if (children->first->depth != children->second->depth) {
            throw new std::runtime_error("Attempted to create uneven buffer");
        }
        depth = children->first->depth + 1;
        saved_hash = ::hash(children->first->hash(), children->second->hash());
        if (children->second->packed_size == 0) {
            // Ignore the second half as it's all zeroes
            packed_size = children->first->packed_size;
        } else {
            // Ignore any trailing zeroes in the left half of the buffer, as
            // they're followed by non-zero data in the right half of the buffer
            packed_size = length_of_depth(children->first->depth) +
                          children->second->packed_size;
        }
    } else {
        auto& bytes = std::get<LeafData>(components);
        saved_hash = ::hash(bytes);
        depth = 0;
        // Go backwards through the bytes to find the last non-zero byte
        packed_size = leaf_size;
        while (packed_size > 0) {
            if (bytes[packed_size - 1] != 0) {
                break;
            }
            packed_size--;
        }
    }
}

Buffer Buffer::grow(uint64_t new_depth) const {
    Buffer ret(*this);
    while (ret.depth < new_depth) {
        ret = Buffer(std::make_shared<Buffer>(ret),
                     zero_buffers_of_depth[ret.depth]);
    }
    return ret;
}

Buffer Buffer::trim() const {
    Buffer ret(*this);
    while (true) {
        if (auto children = ret.get_children()) {
            if (children->second->packed_size == 0) {
                ret = *children->first;
                continue;
            }
        }
        break;
    }
    return ret;
}

Buffer Buffer::set_many_without_resize(uint64_t offset,
                                       const std::vector<uint8_t>& arr,
                                       uint64_t arr_offset,
                                       uint64_t arr_length) const {
    Buffer ret(*this);
    std::vector<Buffer*> path;
    Buffer* target = &ret;
    while (true) {
        path.push_back(target);
        if (auto children = target->get_children()) {
            // Clone each buffer on our way down, and adjust the offset.
            auto child_size = children->first->size();
            if (offset >= child_size) {
                offset -= child_size;
                children->second = std::make_shared<Buffer>(*children->second);
                target = children->second.get();
            } else {
                children->first = std::make_shared<Buffer>(*children->first);
                target = children->first.get();
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
            for (auto it = path.rbegin(); it != path.rend(); it++) {
                (*it)->recompute();
            }
            return ret;
        }
    }
}

Buffer::Buffer() : Buffer(LeafData{}) {}

Buffer Buffer::fromData(const std::vector<uint8_t>& data) {
    // Grow the buffer to the necessary length
    Buffer buf;
    buf = buf.grow(needed_depth(data.size()));
    // Set each up to 32 byte chunk of the buffer
    for (uint64_t i = 0; i < data.size(); i += leaf_size) {
        uint64_t len = leaf_size;
        if (i + len > data.size()) {
            // The last chunk might be smaller than 32 bytes
            len = data.size() - i;
        }
        buf = buf.set_many_without_resize(i, data, i, len);
    }
    return buf.trim();
}

uint64_t Buffer::size() const {
    return length_of_depth(depth);
}

uint64_t Buffer::lastIndex() const {
    if (packed_size == 0) {
        return 0;
    } else {
        return packed_size - 1;
    }
}

uint64_t Buffer::data_length() const {
    return packed_size;
}

uint256_t Buffer::hash() const {
    return saved_hash;
}

Buffer Buffer::set_many(uint64_t offset, std::vector<uint8_t> arr) const {
    Buffer ret(*this);
    if (offset + arr.size() > ret.size()) {
        ret = ret.grow(needed_depth(offset + arr.size()));
    }
    ret = ret.set_many_without_resize(offset, arr, 0, arr.size());
    return ret.trim();
}

Buffer Buffer::set(uint64_t offset, uint8_t v) const {
    return set_many(offset, std::vector(1, v));
}

std::vector<uint8_t> Buffer::get_many(uint64_t offset, size_t len) const {
    const Buffer* target = this;
    while (true) {
        if (auto children = target->get_children_const()) {
            // Move downwards towards the target
            auto child_size = children->first->size();
            if (offset >= child_size) {
                offset -= child_size;
                target = children->second.get();
            } else {
                target = children->first.get();
            }
        } else {
            // We've found the target leaf
            auto& bytes = std::get<LeafData>(target->components);
            if (offset >= leaf_size) {
                return std::vector<uint8_t>(len, (unsigned char)0);
            } else if (offset + len > leaf_size) {
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
    ret.reserve(size());
    std::vector<const Buffer*> to_visit;
    to_visit.reserve(depth);
    const Buffer* current = this;
    // Visit the tree depth first
    while (true) {
        if (auto children = current->get_children_const()) {
            // Visit the left side of the buffer now, and save the right side
            // for later
            current = children->first.get();
            to_visit.push_back(children->second.get());
        } else {
            const auto& bytes = std::get<LeafData>(current->components);
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
    throw new std::runtime_error("TODO");
}

std::vector<unsigned char> Buffer::makeNormalizationProof() const {
    throw new std::runtime_error("TODO");
}

std::vector<Buffer> Buffer::serialize(
    std::vector<unsigned char>& value_vector) const {
    // first check if it's empty
    std::vector<Buffer> ret{};
    value_vector.push_back(depth);
    if (auto children = get_children_const()) {
        marshal_uint256_t(::hash(*children->first), value_vector);
        marshal_uint256_t(::hash(*children->second), value_vector);
        ret.push_back(*children->first);
        ret.push_back(*children->second);
    } else {
        auto& bytes = std::get<LeafData>(components);
        std::copy(bytes.begin(), bytes.end(), std::back_inserter(value_vector));
    }
    return ret;
}
