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

import "../../../utils/contracts/src/LibSafeMath.sol";
import "./ERC1155.sol";
import "./interfaces/IERC1155Mintable.sol";


/// @dev Mintable form of ERC1155
/// Shows how easy it is to mint new items
contract ERC1155Mintable is
    IERC1155Mintable,
    ERC1155
{
    using LibSafeMath for uint256;

    /// token nonce
    uint256 internal nonce;

    /// mapping from token to creator
    mapping (uint256 => address) public creators;

    /// mapping from token to max index
    mapping (uint256 => uint256) public maxIndex;

    /// asserts token is owned by msg.sender
    modifier creatorOnly(uint256 _id) {
        require(creators[_id] == msg.sender);
        _;
    }

    /// @dev creates a new token
    /// @param uri URI of token
    /// @param isNF is non-fungible token
    /// @return type_ of token (a unique identifier)
    function create(
        string calldata uri,
        bool isNF
    )
        external
        returns (uint256 type_)
    {
        // Store the type in the upper 128 bits
        type_ = (++nonce << 128);

        // Set a flag if this is an NFI.
        if (isNF) {
            type_ = type_ | TYPE_NF_BIT;
        }

        // This will allow restricted access to creators.
        creators[type_] = msg.sender;

        // emit a Transfer event with Create semantic to help with discovery.
        emit TransferSingle(
            msg.sender,
            address(0x0),
            address(0x0),
            type_,
            0
        );

        if (bytes(uri).length > 0) {
            emit URI(uri, type_);
        }
    }

    /// @dev creates a new token
    /// @param type_ of token
    /// @param uri URI of token
    function createWithType(
        uint256 type_,
        string calldata uri
    )
        external
    {
        // This will allow restricted access to creators.
        creators[type_] = msg.sender;

        // emit a Transfer event with Create semantic to help with discovery.
        emit TransferSingle(
            msg.sender,
            address(0x0),
            address(0x0),
            type_,
            0
        );

        if (bytes(uri).length > 0) {
            emit URI(uri, type_);
        }
    }

    /// @dev mints fungible tokens
    /// @param id token type
    /// @param to beneficiaries of minted tokens
    /// @param quantities amounts of minted tokens
    function mintFungible(
        uint256 id,
        address[] calldata to,
        uint256[] calldata quantities
    )
        external
        creatorOnly(id)
    {
        // sanity checks
        require(
            isFungible(id),
            "TRIED_TO_MINT_FUNGIBLE_FOR_NON_FUNGIBLE_TOKEN"
        );

        // mint tokens
        for (uint256 i = 0; i < to.length; ++i) {
            // cache to reduce number of loads
            address dst = to[i];
            uint256 quantity = quantities[i];

            // Grant the items to the caller
            balances[id][dst] = quantity.safeAdd(balances[id][dst]);

            // Emit the Transfer/Mint event.
            // the 0x0 source address implies a mint
            // It will also provide the circulating supply info.
            emit TransferSingle(
                msg.sender,
                address(0x0),
                dst,
                id,
                quantity
            );

            // if `to` is a contract then trigger its callback
            if (dst.isContract()) {
                bytes4 callbackReturnValue = IERC1155Receiver(dst).onERC1155Received(
                    msg.sender,
                    msg.sender,
                    id,
                    quantity,
                    ""
                );
                require(
                    callbackReturnValue == ERC1155_RECEIVED,
                    "BAD_RECEIVER_RETURN_VALUE"
                );
            }
        }
    }

    /// @dev mints a non-fungible token
    /// @param type_ token type
    /// @param to beneficiaries of minted tokens
    function mintNonFungible(
        uint256 type_,
        address[] calldata to
    )
        external
        creatorOnly(type_)
    {
        // No need to check this is a nf type rather than an id since
        // creatorOnly() will only let a type pass through.
        require(
            isNonFungible(type_),
            "TRIED_TO_MINT_NON_FUNGIBLE_FOR_FUNGIBLE_TOKEN"
        );

        // Index are 1-based.
        uint256 index = maxIndex[type_] + 1;

        for (uint256 i = 0; i < to.length; ++i) {
            // cache to reduce number of loads
            address dst = to[i];
            uint256 id  = type_ | index + i;

            nfOwners[id] = dst;

            // You could use base-type id to store NF type balances if you wish.
            // balances[_type][dst] = quantity.safeAdd(balances[_type][dst]);

            emit TransferSingle(msg.sender, address(0x0), dst, id, 1);

            // if `to` is a contract then trigger its callback
            if (dst.isContract()) {
                bytes4 callbackReturnValue = IERC1155Receiver(dst).onERC1155Received(
                    msg.sender,
                    msg.sender,
                    id,
                    1,
                    ""
                );
                require(
                    callbackReturnValue == ERC1155_RECEIVED,
                    "BAD_RECEIVER_RETURN_VALUE"
                );
            }
        }

        // record the `maxIndex` of this nft type
        // this allows us to mint more nft's of this type in a subsequent call.
        maxIndex[type_] = to.length.safeAdd(maxIndex[type_]);
    }
}
