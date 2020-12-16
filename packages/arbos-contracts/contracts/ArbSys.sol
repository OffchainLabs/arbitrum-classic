/*
 * Copyright 2019-2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity >=0.4.21 <0.6.0;

interface ArbSys {
    // Send given amount of ERC-20 tokens to dest with token contract sender.
    // This is safe to freely call since the sender is authenticated and thus
    // you can only send fake tokens, not steal real ones
    function withdrawERC20(address dest, uint256 amount) external;

    // Send given ERC-721 token to dest with token contract sender.
    // This is safe by the above arguement
    function withdrawERC721(address dest, uint256 id) external;

    // Send given amount of Eth to dest with from sender.
    function withdrawEth(address dest) external payable;

    // Return the number of transactions issued by the given external account
    // or the account sequence number of the given contract
    function getTransactionCount(address account) external view returns (uint256);

    // Return the value of the storage slot for the given account at the given index
    // This function is only callable from address 0 to prevent contracts from being
    // able to call it
    function getStorageAt(address account, uint256 index) external view returns (uint256);

    // Register an address in the address table
    // Return index of the address (existing index, or newly created index if not already registered)
    function addressTable_register(address addr) external returns (uint256);

    // Return index of an address in the address table (revert if address isn't in the table)
    function addressTable_lookup(address addr) external view returns (uint256);

    // Check whether an address exists in the address table
    function addressTable_addressExists(address addr) external view returns (bool);

    // Get size of address table (= first unused index)
    function addressTable_size() external view returns (uint256);

    // Return address at a given index in address table (revert if index is beyond end of table)
    function addressTable_lookupIndex(uint256 index) external view returns (address);

    // Read a compressed address from a bytes buffer
    // Return resulting address and updated offset into the buffer (revert if buffer is too short)
    function addressTable_decompress(bytes calldata buf, uint256 offset)
        external
        view
        returns (address, uint256);

    // Compress an address and return the result
    function addressTable_compress(address addr) external view returns (bytes memory);

    // Associate a BLS public key with the caller's address
    function registerBlsKey(
        uint256 x0,
        uint256 x1,
        uint256 y0,
        uint256 y1
    ) external;

    // Get the BLS public key associated with an address (revert if there isn't one)
    function getBlsPublicKey(address addr)
        external
        view
        returns (
            uint256,
            uint256,
            uint256,
            uint256
        );

    // Upload a serialized function table and associate it with the caller's address
    // If caller already had a function table, this will overwrite the old one
    // Revert if buf is mal-formatted
    // (Caller will typically be an aggregator)
    function uploadFunctionTable(bytes calldata buf) external;

    // Get the size of addr's function table; revert if addr doesn't have a function table
    function functionTableSize(address addr) external view returns (uint256);

    // Get the entry from addr's function table, at index; revert if addr has no table or index out of bounds
    // Returns (functionCode, isPayable, gasLimit)
    function functionTableGet(address addr, uint256 index)
        external
        view
        returns (
            uint256,
            bool,
            uint256
        );

    event EthWithdrawal(address indexed destAddr, uint256 amount);
    event ERC20Withdrawal(address indexed destAddr, address indexed tokenAddr, uint256 amount);
    event ERC721Withdrawal(address indexed destAddr, address indexed tokenAddr, uint256 indexed id);
}
