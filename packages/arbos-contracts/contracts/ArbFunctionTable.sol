// SPDX-License-Identifier: Apache-2.0

pragma solidity >=0.4.21 <0.7.0;

interface ArbFunctionTable {
    // Upload a serialized function table and associate it with the caller's address
    // If caller already had a function table, this will overwrite the old one
    // Revert if buf is mal-formatted
    // (Caller will typically be an aggregator)
    function upload(bytes calldata buf) external;

    // Get the size of addr's function table; revert if addr doesn't have a function table
    function size(address addr) external view returns (uint256);

    // Get the entry from addr's function table, at index; revert if addr has no table or index out of bounds
    // Returns (functionCode, isPayable, gasLimit)
    function get(address addr, uint256 index)
        external
        view
        returns (
            uint256,
            bool,
            uint256
        );
}
