pragma solidity ^0.6.1;

contract Merkle {
    
    bytes32[] zeros;
    
    constructor() public {
        zeros[0] = keccak1(0);
        for (uint i = 1; i < 64; i++) {
            zeros[i] = keccak2(zeros[i-1], zeros[i-1]);
        }
    }
    
    function keccak1(bytes32 b) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(b));
    }

    function keccak2(bytes32 a, bytes32 b) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(a, b));
    }

    // hashes are normalized
    function get(bytes32 buf, uint loc, bytes32[] memory proof) public pure returns (bytes32) {
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

    function set(bytes32 buf, uint loc, bytes32 v, bytes32[] memory proof, uint loc2, bytes32[] memory proof2) public view returns (bytes32) {
        // three possibilities, the tree depth stays same, it becomes lower or it's extended
        bytes32 acc = keccak1(v);
        // check that the proof matches original
        get(buf, loc, proof);
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
        require(proof2.length <= proof.length);
        require(proof2.length > 0);
        bytes32 acc2 = keccak1(proof2[0]);
        for (uint i = 1; i < proof2.length-1; i++) {
            if (loc2 & 1 == 1) acc2 = keccak2(acc2, proof2[i]);
            else acc2 = keccak2(proof2[i], acc2);
            loc2 = loc2 >> 1;
        }
        require(acc2 != zeros[proof2.length]);
        bytes32 res = keccak2(proof2[proof2.length-1], acc2);
        acc2 = res;
        for (uint i = proof.length-1; i < proof.length; i++) {
            acc2 = keccak2(res, zeros[i]);
        }
        require(acc2 == acc);
        return res;
    }
    
}
