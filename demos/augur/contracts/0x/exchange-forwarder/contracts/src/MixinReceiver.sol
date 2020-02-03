/*

  Copyright 2019 ZeroEx Intl.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

pragma solidity 0.5.15;


contract MixinReceiver {

    bytes4 constant public ERC1155_RECEIVED       = 0xf23a6e61;
    bytes4 constant public ERC1155_BATCH_RECEIVED = 0xbc197c81;

    /// @notice Handle the receipt of a single ERC1155 token type
    /// @dev The smart contract calls this function on the recipient
    /// after a `safeTransferFrom`. This function MAY throw to revert and reject the
    /// transfer. Return of other than the magic value MUST result in the
    ///transaction being reverted
    /// Note: the contract address is always the message sender
    /// @param operator  The address which called `safeTransferFrom` function
    /// @param from      The address which previously owned the token
    /// @param id        An array containing the ids of the token being transferred
    /// @param value     An array containing the amount of tokens being transferred
    /// @param data      Additional data with no specified format
    /// @return          `bytes4(keccak256("onERC1155Received(address,address,uint256,uint256,bytes)"))`
    function onERC1155Received(
        address operator,
        address from,
        uint256 id,
        uint256 value,
        bytes calldata data
    )
        external
        returns (bytes4)
    {
        return ERC1155_RECEIVED;
    }

    /// @notice Handle the receipt of multiple ERC1155 token types
    /// @dev The smart contract calls this function on the recipient
    /// after a `safeTransferFrom`. This function MAY throw to revert and reject the
    /// transfer. Return of other than the magic value MUST result in the
    /// transaction being reverted
    /// Note: the contract address is always the message sender
    /// @param operator  The address which called `safeTransferFrom` function
    /// @param from      The address which previously owned the token
    /// @param ids       An array containing ids of each token being transferred
    /// @param values    An array containing amounts of each token being transferred
    /// @param data      Additional data with no specified format
    /// @return           `bytes4(keccak256("onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)"))`
    function onERC1155BatchReceived(
        address operator,
        address from,
        uint256[] calldata ids,
        uint256[] calldata values,
        bytes calldata data
    )
        external
        returns (bytes4)
    {
        return ERC1155_BATCH_RECEIVED;
    }
}