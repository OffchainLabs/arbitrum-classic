// SPDX-License-Identifier: Apache-2.0

pragma solidity >=0.4.21 <0.7.0;

/** @title Interface for providing Outbox proof data
 *  @notice This contract doesn't exist on-chain. Instead it is a virtual interface accessible at 0x00000000000000000000000000000000000000C8
 * This is a cute trick to allow an Arbitrum node to provide data without us having to implement an additional RPC )
 */

interface NodeInterface {
    /**
     * @notice Returns the proof necessary to redeem a message
     * @param batchNum index of outbox entry (i.e., outgoing messages Merkle root) in array of outbox entries
     * @param index index of outgoing message in outbox entry
     * @return proof Merkle proof of message inclusion in outbox entry
     * @return path Index of message in outbox entry
     * @return l2Sender sender if original message (i.e., caller of ArbSys.sendTxToL1)
     * @return l1Dest destination address for L1 contract call
     * @return l2Block l2 block number at which sendTxToL1 call was made
     * @return l1Block l1 block number at which sendTxToL1 call was made
     * @return timestamp l2 Timestamp at which sendTxToL1 call was made
     * @return amount value in L1 message in wei
     * @return calldataForL1 abi-encoded L1 message data
     */
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
