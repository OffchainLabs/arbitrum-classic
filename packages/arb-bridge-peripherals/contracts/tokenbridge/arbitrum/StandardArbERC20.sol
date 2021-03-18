// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2020, Offchain Labs, Inc.
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

pragma solidity ^0.6.11;

import "./open-zeppelin/OZERC20.sol";
import "arb-bridge-eth/contracts/libraries/Cloneable.sol";
import "./IArbToken.sol";
import "./ArbTokenBridge.sol";

contract StandardArbERC20 is OZERC20, Cloneable, IArbToken {
    ArbTokenBridge public bridge;
    address public l1Address;
    uint256 private immutable deploymentChainId;
    bytes32 private immutable _DOMAIN_SEPARATOR;

    constructor() {
        uint256 chainId;
        assembly {chainId := chainid()}
        deploymentChainId = chainId;
        _DOMAIN_SEPARATOR = _calculateDomainSeparator(chainId);
    }

    modifier onlyBridge {
        require(msg.sender == address(bridge), "ONLY_BRIDGE");
        _;
    }

    function initialize(
        address _bridge,
        address _l1Address,
        uint8 decimals_
    ) external override {
        require(address(bridge) == address(0), "ALREADY_INIT");
        bridge = ArbTokenBridge(_bridge);
        l1Address = _l1Address;
        _decimals = decimals_;
    }

    function DOMAIN_SEPARATOR() external view returns (bytes32) {
        uint256 chainId;
        assembly {chainId := chainid()}
        return chainId == deploymentChainId ? _DOMAIN_SEPARATOR : _calculateDomainSeparator(chainId);
    }

    function updateInfo(string memory newName, string memory newSymbol) public override onlyBridge {
        if (bytes(newName).length != 0) {
            _name = newName;
        }
        if (bytes(newSymbol).length != 0) {
            _symbol = newSymbol;
        }
    }

    function bridgeMint(address account, uint256 amount) external override onlyBridge {
        _mint(account, amount);
    }

    function withdraw(address destination, uint256 amount) external override {
        _burn(msg.sender, amount);
        bridge.withdraw(l1Address, destination, amount);
    }

    function migrate(uint256 amount, address target) external {
        _burn(msg.sender, amount);
        bridge.migrate(l1Address, target, msg.sender, amount);
    }

    /// @dev Sets `value` as allowance of `spender` account over `owner` account's token, given `owner` account's signed approval.
    /// Emits {Approval} event.
    /// Requirements:
    ///   - `deadline` must be timestamp in future.
    ///   - `v`, `r` and `s` must be valid `secp256k1` signature from `owner` account over EIP712-formatted function arguments.
    ///   - the signature must use `owner` account's current nonce (see {nonces}).
    ///   - the signer cannot be zero address and must be `owner` account.
    /// For more information on signature format, see https://eips.ethereum.org/EIPS/eip-2612#specification[relevant EIP section].
    function permit(
        address owner,
        address spender,
        uint256 value,
        uint256 deadline,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external {
        require(block.timestamp <= deadline, "WETH: Expired permit");

        uint256 chainId;
        assembly {chainId := chainid()}
        
        bytes32 domainSeparator = chainId == deploymentChainId
            ? _DOMAIN_SEPARATOR
            : _calculateDomainSeparator(chainId);

        bytes32 hashStruct = keccak256(
            abi.encode(
                PERMIT_TYPEHASH,
                owner,
                spender,
                value,
                nonces[owner]++,
                deadline));

        bytes32 hash = keccak256(
            abi.encodePacked(
                "\x19\x01",
                domainSeparator,
                hashStruct));

        address signer = ecrecover(hash, v, r, s);
        require(signer != address(0) && signer == owner, "Invalid permit");

        _approve(owner, spender, value);
        emit Approval(owner, spender, value);
    }

    /// @dev Moves `value` tokens from caller's account to account (`to`), 
    /// after which a call is executed to an ERC677-compliant contract with the `data` parameter.
    /// Emits {Transfer} event.
    /// Returns boolean value indicating whether operation succeeded.
    /// Requirements:
    ///   - caller account must have at least `value` tokens.
    /// For more information on transferAndCall format, see https://github.com/ethereum/EIPs/issues/677.
    function transferAndCall(address to, uint value, bytes calldata data) external returns (bool) {
        _transfer(msg.sender, to, value);
        emit Transfer(msg.sender, to, value);

        return ITransferReceiver(to).onTokenTransfer(msg.sender, value, data);
    }
}
