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
#include <optional>
#include <utility>
#include <variant>
#include <vector>

constexpr uint64_t ALIGN = 32;

struct CachedCalculation {
    std::atomic<std::optional<uint256_t>> val;

    CachedCalculation() : val(std::nullopt) {}

    CachedCalculation(const CachedCalculation& o) : val(o.val.load()) {}

    CachedCalculation& operator=(const CachedCalculation& o) {
        val = o.val.load();
        return *this;
    }
};

class Buffer {
   public:
    static constexpr uint64_t leaf_size = 32;
    static constexpr uint64_t children_size = 2;
    using LeafData = std::array<unsigned char, leaf_size>;
    using NodeData =
        std::pair<std::shared_ptr<Buffer>, std::shared_ptr<Buffer>>;

   private:
    mutable CachedCalculation hash_cache;
    mutable CachedCalculation packed_size_cache;

    // The depth of this buffer as a tree. A leaf node (32 bytes) is depth 0.
    size_t depth;

    // The components of this buffer. If this is a leaf (at the bottom of the
    // tree, depth == 0), it'll contain raw bytes. If it's a branch (higher up
    // in the tree, depth > 0), it'll contain child buffers.
    std::variant<LeafData, NodeData> components;

    // Returns a pointer to this buffer's children, or null if this is a leaf
    NodeData* get_children();
    // Like get_children but const
    const NodeData* get_children_const() const;

    // The size of this buffer after trimming any zero bytes at the end
    uint256_t packed_size() const;

    // Returns a buffer with a depth of at least new_depth and the same data
    [[nodiscard]] Buffer grow(uint64_t new_depth) const;

    // Returns the smallest possible buffer representing the same data
    [[nodiscard]] Buffer trim() const;

    // Like the public method, but requires that the buffer must not need
    // growing or trimming, and specifying an offset and length to the passed
    // in array. The bytes set must be within a single 32 byte chunk.
    [[nodiscard]] Buffer set_many_without_resize(
        uint64_t offset,
        const std::vector<uint8_t>& arr,
        uint64_t arr_offset,
        uint64_t arr_length) const;

   public:
    // Creates an "empty" buffer (actually has 32 zero bytes)
    Buffer();

    // Create a leaf node
    explicit Buffer(LeafData bytes);

    // Create a branch node composed of two buffers with equal depths
    Buffer(std::shared_ptr<Buffer> left, std::shared_ptr<Buffer> right);

    // Creates a buffer representing the given bytes
    static Buffer fromData(const std::vector<uint8_t>& data);

    // Returns the size of the buffer, **including** any trailing zeroes
    [[nodiscard]] uint256_t size() const;

    // Returns the last non-zero index of the buffer, or 0 if the buffer is
    // entirely zeroes
    [[nodiscard]] uint64_t lastIndex() const;

    // Returns the size of the buffer, **not including** any trailing zeroes
    [[nodiscard]] uint256_t data_length() const;

    // Returns the hash of the buffer, "packing" any trailing zero segments
    [[nodiscard]] uint256_t hash() const;

    // Sets the byte at a given offset, growing or shrinking as needed
    [[nodiscard]] Buffer set(uint64_t offset, uint8_t v) const;

    // Sets bytes at a given offset, growing or shrinking as needed. The bytes
    // set must be within a single 32 byte chunk.
    [[nodiscard]] Buffer set_many(uint64_t offset,
                                  std::vector<uint8_t> arr) const;

    // Gets the byte at a given offset
    [[nodiscard]] uint8_t get(uint64_t pos) const;

    // Gets an array of bytes of a given length at a given position. The bytes
    // must be within a single 32 byte chunk.
    [[nodiscard]] std::vector<uint8_t> get_many(uint64_t pos, size_t len) const;

    // Converts the buffer to a single flat byte vector
    [[nodiscard]] std::vector<uint8_t> toFlatVector() const;

    [[nodiscard]] std::vector<unsigned char> makeProof(uint64_t loc) const;

    [[nodiscard]] std::vector<unsigned char> makeNormalizationProof() const;

    std::vector<Buffer> serialize(
        std::vector<unsigned char>& value_vector) const;
};

inline uint256_t hash(const Buffer& b) {
    return hash(123, b.hash());
}

inline bool operator==(const Buffer& val1, const Buffer& val2) {
    return val1.hash() == val2.hash();
}

inline bool operator!=(const Buffer& val1, const Buffer& val2) {
    return val1.hash() != val2.hash();
}

#endif /* buffer_hpp */
