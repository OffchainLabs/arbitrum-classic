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
     * @return path Merkle path to message
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

    /**
     * @notice Estimate the cost of putting a message in the L2 inbox that is reexecuted
     * @param sender sender of the L1 and L2 transaction
     * @param deposit amount to deposit to sender in L2
     * @param destAddr destination L2 contract address
     * @param l2CallValue call value for retryable L2 message
     * @param maxSubmissionCost Max gas deducted from user's L2 balance to cover base submission fee
     * @param excessFeeRefundAddress maxgas x gasprice - execution cost gets credited here on L2 balance
     * @param callValueRefundAddress l2Callvalue gets credited here on L2 if retryable txn times out or gets cancelled
     * @param maxGas Max gas deducted from user's L2 balance to cover L2 execution
     * @param gasPriceBid price bid for L2 execution
     * @param data ABI encoded data of L2 message
     * @return gas used, and gas price to execute this transaction
     */
    function estimateRetryableTicket(
        address sender,
        uint256 deposit,
        address destAddr,
        uint256 l2CallValue,
        uint256 maxSubmissionCost,
        address excessFeeRefundAddress,
        address callValueRefundAddress,
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes calldata data
    ) external pure returns (uint256, uint256);
}
