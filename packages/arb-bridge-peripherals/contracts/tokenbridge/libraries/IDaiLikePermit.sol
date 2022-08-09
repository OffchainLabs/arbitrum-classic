// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.6.11;

/**
 * @dev Dai Stablecoin ERC20 Interface
 *
 * For use when calling permit on tokens that implement permit similarly to Dai.
 *
 * https://github.com/makerdao/dss/blob/17187f7d47be2f4c71d218785e1155474bbafe8a/src/dai.sol
 */


interface IDaiLikePermit {
    
    function permit(
        address holder,
        address spender,
        uint256 nonce,
        uint256 expiry,
        bool allowed,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external;

}