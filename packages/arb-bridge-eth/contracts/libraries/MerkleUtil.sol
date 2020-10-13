// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.5.11;

library MerkleUtil {
    
    function makeZeros() internal pure returns (bytes32[] memory) {
        bytes32[] memory zeros = new bytes32[](64);
        zeros[0] = keccak1(0);
        for (uint i = 1; i < 64; i++) {
            zeros[i] = keccak2(zeros[i-1], zeros[i-1]);
        }
        return zeros;
    }

    function keccak1(bytes32 b) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(b));
    }

    function keccak2(bytes32 a, bytes32 b) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(a, b));
    }

    // hashes are normalized
    function get(bytes32 buf, uint loc, bytes32[] memory proof) internal pure returns (bytes32) {
        // empty tree is full of zeros
        if (proof.length == 0) {
            require(buf == bytes32(0));
            return 0;
        }
        bytes32 acc = keccak1(proof[0]);
        for (uint i = 1; i < proof.length; i++) {
            if (loc & 1 == 1) acc = keccak2(acc, proof[i]);
            else acc = keccak2(proof[i], acc);
            loc = loc >> 1;
        }
        require(acc == buf);
        // maybe it is a zero outside the actual tree
        if (loc > 0) return 0;
        return proof[0];
    }

    function calcHeight(uint loc) internal pure returns (uint) {
        if (loc == 0) return 0;
        else return 1+calcHeight(loc>>1);
    }

    function set(bytes32 buf, uint loc, bytes32 v, bytes32[] memory proof, uint nh, bytes32 normal1, bytes32 normal2) internal pure returns (bytes32) {
        // three possibilities, the tree depth stays same, it becomes lower or it's extended
        bytes32 acc = keccak1(v);
        // check that the proof matches original
        get(buf, loc, proof);
        bytes32[] memory zeros = makeZeros();
        // extended
        if (loc > (proof.length << 2)) {
            if (v == 0) return buf;
            uint height = calcHeight(loc);
            // build the left branch
            for (uint i = proof.length; i < height-1; i++) {
                buf = keccak2(buf, zeros[i]);
            }
            for (uint i = 1; i < height-1; i++) {
                if (loc & 1 == 1) acc = keccak2(acc, zeros[i]);
                else acc = keccak2(zeros[i], acc);
                loc = loc >> 1;
            }
            return keccak2(buf, acc);
        }
        for (uint i = 1; i < proof.length; i++) {
            bytes32 a = loc & 1 == 1 ? proof[i] : acc;
            bytes32 b = loc & 1 == 1 ? acc : proof[i];
            acc = keccak2(a, b);
            loc = loc >> 1;
        }
        if (v != bytes32(0)) return acc;
        require(normal2 != zeros[nh]);
        bytes32 res = keccak2(normal1, normal2);
        bytes32 acc2 = res;
        for (uint i = nh; i < proof.length; i++) {
            acc2 = keccak2(res, zeros[i]);
        }
        require(acc2 == acc);
        return res;
    }

    function getByte(bytes32 word, uint256 num) internal pure returns (uint256) {
        return (uint256(word) >> ((31-num)*8)) & 0xff;
    }

    function setByte(bytes32 word, uint256 num, uint256 b) internal pure returns (bytes32) {
        bytes memory arr = bytes32ToArray(word);
        arr[num] = bytes1(uint8(b));
        return bytes32(bytes32FromArray(arr));
    }

    function setByte(bytes32 word, uint256 num, bytes1 b) internal pure returns (bytes32) {
        bytes memory arr = bytes32ToArray(word);
        arr[num] = b;
        return bytes32(bytes32FromArray(arr));
    }

    function getOp(uint8 op, bytes32 buf, uint offset, bytes32[] memory proof1, bytes32[] memory proof2) public pure returns (uint256) {
        if (op == 0xa1) {
            return getByte(get(buf, offset/32, proof1), offset%32);
        } else if (op == 0xa2) {
            return getBuffer64(buf, offset, proof1, proof2);
        } else if (op == 0xa3) {
            return getBuffer256(buf, offset, proof1, proof2);
        }
    }

    function setOp(uint8 op, bytes32 buf, uint offset, uint256 b, bytes32[] memory proof1, bytes32[] memory nproof1, bytes32[] memory proof2, bytes32[] memory nproof2) public pure returns (bytes32) {
        if (op == 0xa4) {
            bytes32 word = get(buf, offset/32, proof1);
            bytes32 nword = setByte(word, offset%32, b);
            return set(buf, offset/32, nword, proof1, nproof1);
        } else if (op == 0xa5) {
            return setBuffer64(buf, offset, bytes32(b), proof1, nproof1, proof2, nproof2);
        } else if (op == 0xa6) {
        }
    }

    function bufferOp(uint8 op, bytes32 buf, uint offset, uint256 b, bytes32[] memory proof1, bytes32[] memory nproof1, bytes32[] memory proof2, bytes32[] memory nproof2) public pure returns (bytes32) {
        if (op == 0xa1) {
            return bytes32(getByte(get(buf, offset/32, proof1), offset%32));
        } else if (op == 0xa2) {
            return bytes32(getBuffer64(buf, offset, proof1, proof2));
        } else if (op == 0xa3) {
            return bytes32(getBuffer256(buf, offset, proof1, proof2));
        } else if (op == 0xa4) {
            bytes32 word = get(buf, offset/32, proof1);
            bytes32 nword = setByte(word, offset%32, b);
            return set(buf, offset/32, nword, proof1, nproof1);
        } else if (op == 0xa5) {
            return setBuffer64(buf, offset, bytes32(b), proof1, nproof1, proof2, nproof2);
        } else if (op == 0xa6) {
        }
    }

    function bytes32FromArray(bytes memory arr) internal pure returns (uint256) {
        uint256 res = 0;
        for (uint i = 0; i < arr.length; i++) {
            res = res << 8;
            res = res | uint256(uint8(arr[arr.length-1-i]));
        }
        return res;
    }

    function bytes32ToArray(bytes32 b) internal pure returns (bytes memory arr) {
        uint256 acc = uint256(b);
        bytes memory res = new bytes(32);
        for (uint i = 0; i < arr.length; i++) {
            res[31-i] = bytes1(uint8(acc));
            acc = acc >> 8;
        }
        return res;
    }

    function getBuffer64(bytes32 buf, uint256 offset, bytes32[] memory proof1, bytes32[] memory proof2) internal pure returns (uint256) {
        bytes memory res = new bytes(8);
        bytes32 word = get(buf, offset/32, proof1); 
        if (offset%32 + 8 >= 32) {
            bytes32 word2 = get(buf, offset/32 + 1, proof2); 
            for (uint i = 0; i < 8 - (offset%32 + 8 - 32); i++) {
                res[i] = bytes1(uint8(getByte(word, offset%32 + i)));
            }
            for (uint i = 8 - (offset%32 + 8 - 32); i < 8; i++) {
                res[i] = bytes1(uint8(getByte(word2, (offset + i) % 32)));
            }
        } else {
            for (uint i = 0; i < 8; i++) {
                res[i] = bytes1(uint8(getByte(word, offset%32 + i)));
            }
        }
        return bytes32FromArray(res);
    }

    function getBuffer256(bytes32 buf, uint256 offset, bytes32[] memory proof1, bytes32[] memory proof2) internal pure returns (uint256) {
        bytes memory res = new bytes(32);
        bytes32 word = get(buf, offset/32, proof1); 
        if (offset%32 + 32 >= 32) {
            bytes32 word2 = get(buf, offset/32 + 1, proof2); 
            for (uint i = 0; i < 32 - (offset%32 + 32 - 32); i++) {
                res[i] = bytes1(uint8(getByte(word, offset%32 + i)));
            }
            for (uint i = 8 - (offset%32 + 32 - 32); i < 32; i++) {
                res[i] = bytes1(uint8(getByte(word2, (offset + i) % 32)));
            }
        } else {
            for (uint i = 0; i < 32; i++) {
                res[i] = bytes1(uint8(getByte(word, offset%32 + i)));
            }
        }
        return bytes32FromArray(res);
    }

    function set(bytes32 buf, uint loc, bytes32 v, bytes32[] memory proof, bytes32[] memory nproof) internal pure returns (bytes32) {
        return set(buf, loc, v, proof, uint256(nproof[0]), nproof[1], nproof[2]);
    }

    function setBuffer64(bytes32 buf, uint256 offset, bytes32 val, bytes32[] memory proof1, bytes32[] memory nproof1, bytes32[] memory proof2, bytes32[] memory nproof2) internal pure returns (bytes32) {
        bytes memory arr = bytes32ToArray(val);
        bytes32 nword = get(buf, offset/32, proof1);
        if (offset%32 + 8 >= 32) {
            bytes32 nword2 = get(buf, offset/32 + 1, proof2); 
            for (uint i = 0; i < 8 - (offset%32 + 8 - 32); i++) {
                nword = setByte(nword, offset%32 + i, arr[i]);
            }
            for (uint i = 8 - (offset%32 + 8 - 32); i < 8; i++) {
                nword2 = setByte(nword2, (offset+i)%32, arr[i]);
                buf = set(buf, offset/32, nword, proof1, nproof1);
                buf = set(buf, offset/32 + 1, nword2, proof2, nproof2);
            }
        } else {
            for (uint i = 0; i < 8; i++) {
                nword = setByte(nword, offset%32 + i, arr[i]);
            }
            buf = set(buf, offset/32, nword, proof1, nproof1);
        }
        return buf;
    }

}
