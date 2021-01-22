// SPDX-License-Identifier: MIT

pragma solidity ^0.6.11;

// Modified by Offchain Labs to update solidity version

/**
 * @author Wanseob Lim <email@wanseob.com>
 * @title Merkle Mountain Range solidity library
 *
 * @dev The index of this MMR implementation starts from 1 not 0.
 *      And it uses keccak256 for its hash function instead of blake2b
 */
library MMR {
    struct Tree {
        bytes32 root;
        uint256 size;
        uint256 width;
        mapping(uint256 => bytes32) hashes;
        mapping(bytes32 => bytes) data;
    }

    /**
     * @dev This only stores the hashed value of the leaf.
     *      If you need to retrieve the detail data later, use a map to store them.
     */
    function append(Tree storage tree, bytes memory data) public {
        // Hash the leaf node first
        bytes32 dataHash = keccak256(data);
        if (keccak256(tree.data[dataHash]) != dataHash) {
            tree.data[dataHash] = data;
        }
        bytes32 leaf = hashLeaf(tree.size + 1, dataHash);
        // Put the hashed leaf to the map
        tree.hashes[tree.size + 1] = leaf;
        tree.width += 1;
        // Find peaks for the enlarged tree
        uint256[] memory peakIndexes = getPeakIndexes(tree.width);
        // The right most peak's value is the new size of the updated tree
        tree.size = getSize(tree.width);
        // Starting from the left-most peak, get all peak hashes using _getOrCreateNode() function.
        bytes32[] memory peaks = new bytes32[](peakIndexes.length);
        for (uint256 i = 0; i < peakIndexes.length; i++) {
            peaks[i] = _getOrCreateNode(tree, peakIndexes[i]);
        }
        // Create the root hash and update the tree
        tree.root = peakBagging(tree.width, peaks);
    }

    function getPeaks(Tree storage tree) public view returns (bytes32[] memory peaks) {
        // Find peaks for the enlarged tree
        uint256[] memory peakNodeIndexes = getPeakIndexes(tree.width);
        // Starting from the left-most peak, get all peak hashes using _getOrCreateNode() function.
        peaks = new bytes32[](peakNodeIndexes.length);
        for (uint256 i = 0; i < peakNodeIndexes.length; i++) {
            peaks[i] = tree.hashes[peakNodeIndexes[i]];
        }
        return peaks;
    }

    function getLeafIndex(uint256 width) public pure returns (uint256) {
        if (width % 2 == 1) {
            return getSize(width);
        } else {
            return getSize(width - 1) + 1;
        }
    }

    function getSize(uint256 width) public pure returns (uint256) {
        return (width << 1) - numOfPeaks(width);
    }

    /**
     * @dev It returns the root value of the tree
     */
    function getRoot(Tree storage tree) public view returns (bytes32) {
        return tree.root;
    }

    /**
     * @dev It returns the size of the tree
     */
    function getSize(Tree storage tree) public view returns (uint256) {
        return tree.size;
    }

    /**
     * @dev It returns the hash value of a node for the given position. Note that the index starts from 1
     */
    function getNode(Tree storage tree, uint256 index) public view returns (bytes32) {
        return tree.hashes[index];
    }

    /**
     * @dev It returns a merkle proof for a leaf. Note that the index starts from 1
     */
    function getMerkleProof(Tree storage tree, uint256 index)
        public
        view
        returns (
            bytes32 root,
            uint256 width,
            bytes32[] memory peakBagging,
            bytes32[] memory siblings
        )
    {
        require(index < tree.size, "Out of range");
        require(isLeaf(index), "Not a leaf");

        root = tree.root;
        width = tree.width;
        // Find all peaks for bagging
        uint256[] memory peaks = getPeakIndexes(tree.width);

        peakBagging = new bytes32[](peaks.length);
        uint256 cursor;
        for (uint256 i = 0; i < peaks.length; i++) {
            // Collect the hash of all peaks
            peakBagging[i] = tree.hashes[peaks[i]];
            // Find the peak which includes the target index
            if (peaks[i] >= index && cursor == 0) {
                cursor = peaks[i];
            }
        }
        uint256 left;
        uint256 right;

        // Get hashes of the siblings in the mountain which the index belongs to.
        // It moves the cursor from the summit of the mountain down to the target index
        uint8 height = heightAt(cursor);
        siblings = new bytes32[](height - 1);
        while (cursor != index) {
            height--;
            (left, right) = getChildren(cursor);
            // Move the cursor down to the left side or right side
            cursor = index <= left ? left : right;
            // Remaining node is the sibling
            siblings[height - 1] = tree.hashes[index <= left ? right : left];
        }
    }

    function rollUp(
        bytes32 root,
        uint256 width,
        bytes32[] memory peaks,
        bytes32[] memory itemHashes
    ) public pure returns (bytes32 newRoot) {
        // Check the root equals the peak bagging hash
        require(root == peakBagging(width, peaks), "Invalid root hash from the peaks");
        uint256 tmpWidth = width;
        bytes32[255] memory tmpPeakMap = peaksToPeakMap(width, peaks);
        for (uint256 i = 0; i < itemHashes.length; i++) {
            tmpPeakMap = peakUpdate(tmpWidth, tmpPeakMap, itemHashes[i]);
            tmpWidth++;
        }
        return peakBagging(tmpWidth, peakMapToPeaks(tmpWidth, tmpPeakMap));
    }

    function peakBagging(uint256 width, bytes32[] memory peaks) public pure returns (bytes32) {
        uint256 size = getSize(width);
        require(numOfPeaks(width) == peaks.length, "Received invalid number of peaks");
        return keccak256(abi.encodePacked(size, keccak256(abi.encodePacked(size, peaks))));
    }

    function peaksToPeakMap(uint256 width, bytes32[] memory peaks)
        public
        pure
        returns (bytes32[255] memory peakMap)
    {
        uint256 bitIndex;
        uint256 peakRef;
        uint256 count = peaks.length;
        for (uint256 height = 1; height <= 255; height++) {
            // Index starts from the right most bit
            bitIndex = 255 - height;
            peakRef = 1 << (height - 1);
            if ((width & peakRef) != 0) {
                peakMap[bitIndex] = peaks[--count];
            } else {
                peakMap[bitIndex] = bytes32(0);
            }
        }
        require(count == 0, "Invalid number of peaks");
    }

    function peakMapToPeaks(uint256 width, bytes32[255] memory peakMap)
        public
        pure
        returns (bytes32[] memory peaks)
    {
        uint256 arrLength = numOfPeaks(width);
        peaks = new bytes32[](arrLength);
        uint256 count = 0;
        for (uint256 i = 0; i < 255; i++) {
            if (peakMap[i] != bytes32(0)) {
                peaks[count++] = peakMap[i];
            }
        }
        require(count == arrLength, "Invalid number of peaks");
    }

    function peakUpdate(
        uint256 width,
        bytes32[255] memory prevPeakMap,
        bytes32 itemHash
    ) public pure returns (bytes32[255] memory nextPeakMap) {
        uint256 newWidth = width + 1;
        uint256 cursorIndex = getLeafIndex(newWidth);
        bytes32 cursorNode = hashLeaf(cursorIndex, itemHash);
        uint256 bitIndex;
        uint256 peakRef;
        bool prevPeakExist;
        bool nextPeakExist;
        bool obtained;

        for (uint256 height = 1; height <= 255; height++) {
            // Index starts from the right most bit
            bitIndex = 255 - height;
            if (obtained) {
                nextPeakMap[bitIndex] = prevPeakMap[bitIndex];
            } else {
                peakRef = 1 << (height - 1);
                prevPeakExist = (width & peakRef) != 0;
                nextPeakExist = (newWidth & peakRef) != 0;

                // Get new cursor node with hashing the peak and the current cursor
                cursorIndex++;
                if (prevPeakExist) {
                    cursorNode = hashBranch(cursorIndex, prevPeakMap[bitIndex], cursorNode);
                }
                // If new peak exists for the bit index
                if (nextPeakExist) {
                    // If prev peak exists for the bit index
                    if (prevPeakExist) {
                        nextPeakMap[bitIndex] = prevPeakMap[bitIndex];
                    } else {
                        nextPeakMap[bitIndex] = cursorNode;
                    }
                    obtained = true;
                } else {
                    nextPeakMap[bitIndex] = bytes32(0);
                }
            }
        }
    }

    /** Pure functions */

    /**
     * @dev It returns true when the given params verifies that the given value exists in the tree or reverts the transaction.
     */
    function inclusionProof(
        bytes32 root,
        uint256 width,
        uint256 index,
        bytes memory value,
        bytes32[] memory peaks,
        bytes32[] memory siblings
    ) public pure returns (bool) {
        uint256 size = getSize(width);
        require(size >= index, "Index is out of range");
        // Check the root equals the peak bagging hash
        require(
            root == keccak256(abi.encodePacked(size, keccak256(abi.encodePacked(size, peaks)))),
            "Invalid root hash from the peaks"
        );

        // Find the mountain where the target index belongs to
        uint256 cursor;
        bytes32 targetPeak;
        uint256[] memory peakIndexes = getPeakIndexes(width);
        for (uint256 i = 0; i < peakIndexes.length; i++) {
            if (peakIndexes[i] >= index) {
                targetPeak = peaks[i];
                cursor = peakIndexes[i];
                break;
            }
        }
        require(targetPeak != bytes32(0), "Target is not found");

        // Find the path climbing down
        uint256[] memory path = new uint256[](siblings.length + 1);
        uint256 left;
        uint256 right;
        uint8 height = uint8(siblings.length) + 1;
        while (height > 0) {
            // Record the current cursor and climb down
            path[--height] = cursor;
            if (cursor == index) {
                // On the leaf node. Stop climbing down
                break;
            } else {
                // On the parent node. Go left or right
                (left, right) = getChildren(cursor);
                cursor = index > left ? right : left;
                continue;
            }
        }

        // Calculate the summit hash climbing up again
        bytes32 node;
        while (height < path.length) {
            // Move cursor
            cursor = path[height];
            if (height == 0) {
                // cursor is on the leaf
                node = hashLeaf(cursor, keccak256(value));
            } else if (cursor - 1 == path[height - 1]) {
                // cursor is on a parent and a sibling is on the left
                node = hashBranch(cursor, siblings[height - 1], node);
            } else {
                // cursor is on a parent and a sibling is on the right
                node = hashBranch(cursor, node, siblings[height - 1]);
            }
            // Climb up
            height++;
        }

        // Computed hash value of the summit should equal to the target peak hash
        require(node == targetPeak, "Hashed peak is invalid");
        return true;
    }

    /**
     * @dev It returns the hash a parent node with hash(M | Left child | Right child)
     *      M is the index of the node
     */
    function hashBranch(
        uint256 index,
        bytes32 left,
        bytes32 right
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(index, left, right));
    }

    /**
     * @dev it returns the hash of a leaf node with hash(M | DATA )
     *      M is the index of the node
     */
    function hashLeaf(uint256 index, bytes32 dataHash) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(index, dataHash));
    }

    /**
     * @dev It returns the height of the highest peak
     */
    function mountainHeight(uint256 size) public pure returns (uint8) {
        uint8 height = 1;
        while (uint256(1) << height <= size + height) {
            height++;
        }
        return height - 1;
    }

    /**
     * @dev It returns the height of the index
     */
    function heightAt(uint256 index) public pure returns (uint8 height) {
        uint256 reducedIndex = index;
        uint256 peakIndex;
        // If an index has a left mountain subtract the mountain
        while (reducedIndex > peakIndex) {
            reducedIndex -= (uint256(1) << height) - 1;
            height = mountainHeight(reducedIndex);
            peakIndex = (uint256(1) << height) - 1;
        }
        // Index is on the right slope
        height = height - uint8((peakIndex - reducedIndex));
    }

    /**
     * @dev It returns whether the index is the leaf node or not
     */
    function isLeaf(uint256 index) public pure returns (bool) {
        return heightAt(index) == 1;
    }

    /**
     * @dev It returns the children when it is a parent node
     */
    function getChildren(uint256 index) public pure returns (uint256 left, uint256 right) {
        left = index - (uint256(1) << (heightAt(index) - 1));
        right = index - 1;
        require(left != right, "Not a parent");
    }

    /**
     * @dev It returns all peaks of the smallest merkle mountain range tree which includes
     *      the given index(size)
     */
    function getPeakIndexes(uint256 width) public pure returns (uint256[] memory peakIndexes) {
        peakIndexes = new uint256[](numOfPeaks(width));
        uint256 count;
        uint256 size;
        for (uint256 i = 255; i > 0; i--) {
            if (width & (1 << (i - 1)) != 0) {
                // peak exists
                size = size + (1 << i) - 1;
                peakIndexes[count++] = size;
            }
        }
        require(count == peakIndexes.length, "Invalid bit calculation");
    }

    function numOfPeaks(uint256 width) public pure returns (uint256 num) {
        uint256 bits = width;
        while (bits > 0) {
            if (bits % 2 == 1) num++;
            bits = bits >> 1;
        }
        return num;
    }

    /**
     * @dev It returns the hash value of the node for the index.
     *      If the hash already exists it simply returns the stored value. On the other hand,
     *      it computes hashes recursively downward.
     *      Only appending an item calls this function
     */
    function _getOrCreateNode(Tree storage tree, uint256 index) private returns (bytes32) {
        require(index <= tree.size, "Out of range");
        if (tree.hashes[index] == bytes32(0)) {
            (uint256 leftIndex, uint256 rightIndex) = getChildren(index);
            bytes32 leftHash = _getOrCreateNode(tree, leftIndex);
            bytes32 rightHash = _getOrCreateNode(tree, rightIndex);
            tree.hashes[index] = hashBranch(index, leftHash, rightHash);
        }
        return tree.hashes[index];
    }
}
