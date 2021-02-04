// SPDX-License-Identifier: Apache-2.0

pragma solidity >=0.4.21 <0.7.0;

interface ArbSys {
    // Get ArbOS version number
    function arbOSVersion() external pure returns (uint256);

    // Send given amount of Eth to dest with from sender.
    function withdrawEth(address dest) external payable;

    // Send a transaction to L1
    function sendTxToL1(address destAddr, bytes calldata calldataForL1) external payable;

    // Return the number of transactions issued by the given external account
    // or the account sequence number of the given contract
    function getTransactionCount(address account) external view returns (uint256);

    // Return the value of the storage slot for the given account at the given index
    // This function is only callable from address 0 to prevent contracts from being
    // able to call it
    function getStorageAt(address account, uint256 index) external view returns (uint256);

    event EthWithdrawal(address indexed destAddr, uint256 amount);
}
