---
title: MMR.sol Spec
---

Merkle Mountain Range solidity library

The index of this MMR implementation starts from 1 not 0.
And it uses keccak256 for its hash function instead of blake2b

### `append(struct MMR.Tree tree, bytes data)` (public)

This only stores the hashed value of the leaf.
If you need to retrieve the detail data later, use a map to store them.

### `getPeaks(struct MMR.Tree tree) → bytes32[] peaks` (public)

### `getLeafIndex(uint256 width) → uint256` (public)

### `getSize(uint256 width) → uint256` (public)

### `getRoot(struct MMR.Tree tree) → bytes32` (public)

It returns the root value of the tree

### `getSize(struct MMR.Tree tree) → uint256` (public)

It returns the size of the tree

### `getNode(struct MMR.Tree tree, uint256 index) → bytes32` (public)

It returns the hash value of a node for the given position. Note that the index starts from 1

### `getMerkleProof(struct MMR.Tree tree, uint256 index) → bytes32 root, uint256 width, bytes32[] peakBagging, bytes32[] siblings` (public)

It returns a merkle proof for a leaf. Note that the index starts from 1

### `rollUp(bytes32 root, uint256 width, bytes32[] peaks, bytes32[] itemHashes) → bytes32 newRoot` (public)

### `peakBagging(uint256 width, bytes32[] peaks) → bytes32` (public)

### `peaksToPeakMap(uint256 width, bytes32[] peaks) → bytes32[255] peakMap` (public)

### `peakMapToPeaks(uint256 width, bytes32[255] peakMap) → bytes32[] peaks` (public)

### `peakUpdate(uint256 width, bytes32[255] prevPeakMap, bytes32 itemHash) → bytes32[255] nextPeakMap` (public)

### `inclusionProof(bytes32 root, uint256 width, uint256 index, bytes value, bytes32[] peaks, bytes32[] siblings) → bool` (public)

It returns true when the given params verifies that the given value exists in the tree or reverts the transaction.

### `hashBranch(uint256 index, bytes32 left, bytes32 right) → bytes32` (public)

It returns the hash a parent node with hash(M | Left child | Right child)
M is the index of the node

### `hashLeaf(uint256 index, bytes32 dataHash) → bytes32` (public)

it returns the hash of a leaf node with hash(M | DATA )
M is the index of the node

### `mountainHeight(uint256 size) → uint8` (public)

It returns the height of the highest peak

### `heightAt(uint256 index) → uint8 height` (public)

It returns the height of the index

### `isLeaf(uint256 index) → bool` (public)

It returns whether the index is the leaf node or not

### `getChildren(uint256 index) → uint256 left, uint256 right` (public)

It returns the children when it is a parent node

### `getPeakIndexes(uint256 width) → uint256[] peakIndexes` (public)

It returns all peaks of the smallest merkle mountain range tree which includes
the given index(size)

### `numOfPeaks(uint256 width) → uint256 num` (public)
