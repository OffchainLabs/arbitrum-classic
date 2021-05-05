// SPDX-License-Identifier: Apache-2.0

pragma solidity >=0.4.21 <0.7.0;

interface RetryableTicketCreator {
    /**
    @notice Put an message in the L2 inbox that can be reexecuted for some fixed amount of time if it reverts
    * @dev all msg.value will deposited to callValueRefundAddress on L2
    * @param destAddr destination L2 contract address
    * @param l2CallValue call value for retryable L2 message 
    * @param  maxSubmissionCost Max gas deducted from user's L2 balance to cover base submission fee
    * @param excessFeeRefundAddress maxgas x gasprice - execution cost gets credited here on L2 balance
    * @param callValueRefundAddress l2Callvalue gets credited here on L2 if retryable txn times out or gets cancelled
    * @param maxGas Max gas deducted from user's L2 balance to cover L2 execution
    * @param gasPriceBid price bid for L2 execution
    * @param data ABI encoded data of L2 message 
     */
    function createRetryableTicket(
        address destAddr,
        uint256 l2CallValue,
        uint256 maxSubmissionCost,
        address excessFeeRefundAddress,
        address callValueRefundAddress,
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes calldata data
    ) external payable;
}
