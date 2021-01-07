// SPDX-License-Identifier: Apache-2.0

pragma solidity >=0.4.21 <0.7.0;

interface ArbBLS {
    // Associate a BLS public key with the caller's address
    function register(
        uint256 x0,
        uint256 x1,
        uint256 y0,
        uint256 y1
    ) external;

    // Get the BLS public key associated with an address (revert if there isn't one)
    function getPublicKey(address addr)
        external
        view
        returns (
            uint256,
            uint256,
            uint256,
            uint256
        );
}
