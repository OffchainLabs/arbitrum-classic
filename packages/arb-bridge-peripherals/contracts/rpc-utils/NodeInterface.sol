pragma solidity >=0.4.21 <0.7.0;

// This contract doesn't exist on-chain. Instead it is a virtual interface
// implemented by an Arbitrum node in order to provide additional data

interface NodeInterface {
    // Returns the proof necessary to redeem a message batch
    function lookupMessageBatchProof(uint256 batchNum, uint64 index)
        external
        view
        returns (
            bytes32[] memory proof,
            uint256 path,
            address l2Sender,
            address l1Dest,
            uint256 l2Block,
            uint256 l1Block,
            uint256 timestamp,
            uint256 amount,
            bytes memory calldataForL1
        );
}
