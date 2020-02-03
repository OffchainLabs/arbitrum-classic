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
pragma experimental ABIEncoderV2;

import "./IERC1155.sol";


/// @dev Mintable form of ERC1155
/// Shows how easy it is to mint new items
contract IERC1155Mintable is
    IERC1155
{

    /// @dev creates a new token
    /// @param uri URI of token
    /// @param isNF is non-fungible token
    /// @return _type of token (a unique identifier)
    function create(
        string calldata uri,
        bool isNF
    )
        external
        returns (uint256 type_);

    /// @dev mints fungible tokens
    /// @param id token type
    /// @param to beneficiaries of minted tokens
    /// @param quantities amounts of minted tokens
    function mintFungible(
        uint256 id,
        address[] calldata to,
        uint256[] calldata quantities
    )
        external;

    /// @dev mints a non-fungible token
    /// @param type_ token type
    /// @param to beneficiaries of minted tokens
    function mintNonFungible(
        uint256 type_,
        address[] calldata to
    )
        external;
}
