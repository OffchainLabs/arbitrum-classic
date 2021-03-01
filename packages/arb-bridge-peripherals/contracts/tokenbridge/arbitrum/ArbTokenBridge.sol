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

import "@openzeppelin/contracts/utils/Address.sol";
import "./StandardArbERC20.sol";
import "./StandardArbERC777.sol";
import "arb-bridge-eth/contracts/libraries/CloneFactory.sol";

import "./IArbToken.sol";
import "arb-bridge-eth/contracts/libraries/ICloneable.sol";
import "arbos-contracts/contracts/ArbSys.sol";

contract ArbTokenBridge is CloneFactory {
    using Address for address;

    /// @notice This mapping is from L1 address to L2 address
    mapping(address => address) public customToken;

    uint256 exitNum;

    ICloneable public immutable templateERC20;
    ICloneable public immutable templateERC777;
    address immutable public l1Pair;

    modifier onlyEthPair {
        // This ensures that this method can only be called from the L1 pair of this contract
        require(tx.origin == l1Pair, "ONLY_ETH_PAIR");
        _;
    }

    constructor(address _l1Pair) public {
        templateERC20 = new StandardArbERC20();
        templateERC777 = new StandardArbERC777();
        l1Pair = _l1Pair;
    }

    function mintERC777FromL1(
        address l1ERC20,
        address account,
        uint256 amount,
        uint8 decimals
    ) external onlyEthPair {
        IArbToken token = ensureERC777TokenExists(l1ERC20, decimals);
        token.bridgeMint(account, amount);
    }

    function mintERC20FromL1(
        address l1ERC20,
        address account,
        uint256 amount,
        uint8 decimals
    ) external onlyEthPair {
        IArbToken token = ensureERC20TokenExists(l1ERC20, decimals);
        token.bridgeMint(account, amount);
    }

    function updateERC777TokenInfo(
        address l1ERC20,
        string calldata name,
        string calldata symbol,
        uint8 decimals
    ) external onlyEthPair {
        IArbToken token = ensureERC777TokenExists(l1ERC20, decimals);
        token.updateInfo(name, symbol);
    }

    function updateERC20TokenInfo(
        address l1ERC20,
        string calldata name,
        string calldata symbol,
        uint8 decimals
    ) external onlyEthPair {
        IArbToken token = ensureERC20TokenExists(l1ERC20, decimals);
        token.updateInfo(name, symbol);
    }

    function customTokenRegistered(address l1Address, address l2Address) external onlyEthPair {
        customToken[l1Address] = l2Address;
    }

    function withdraw(address l1ERC20, address destination, uint256 amount) external {
        require(msg.sender == calculateBridgedERC777Address(l1ERC20)
            || msg.sender == calculateBridgedERC20Address(l1ERC20)
            || msg.sender == customToken[l1ERC20], "NOT_FROM_TOKEN");

        ArbSys(100).sendTxToL1(
            address(this),
            abi.encodeWithSignature(
                "withdrawFromL2(uint256,address,address,uint256)",
                exitNum,
                l1ERC20,
                destination,
                amount
            )
        );
        exitNum++;
    }

    function migrate(address l1ERC20, address target, address account, uint256 amount) external {
        address bridgedERC777 = calculateBridgedERC777Address(l1ERC20);
        address bridgedERC20 = calculateBridgedERC20Address(l1ERC20);

        require(msg.sender == bridgedERC777 || msg.sender == bridgedERC20, "NOT_FROM_TOKEN");
        require(target == bridgedERC777 || target == bridgedERC20 || target == customToken[l1ERC20], "NOT_TO_TOKEN");

        IArbToken(target).bridgeMint(account, amount);
    }

    function calculateBridgedERC777Address(address l1ERC20) public view returns (address) {
        return calculateCreate2CloneAddress(templateERC777, bytes32(uint256(l1ERC20)));
    }

    function calculateBridgedERC20Address(address l1ERC20) public view returns (address) {
        return calculateCreate2CloneAddress(templateERC20, bytes32(uint256(l1ERC20)));
    }

    function ensureERC777TokenExists(address l1ERC20, uint8 decimals) private returns (IArbToken) {
        address _customToken = customToken[l1ERC20];
        if (_customToken != address(0)) {
            return IArbToken(_customToken);
        }

        address l2Contract = calculateBridgedERC777Address(l1ERC20);

        if (!l2Contract.isContract()) {
            address createdContract = create2Clone(templateERC777, bytes32(uint256(l1ERC20)));
            assert(createdContract == l2Contract);
            IArbToken(l2Contract).initialize(address(this), l1ERC20, decimals);
        }
        return IArbToken(l2Contract);
    }

    function ensureERC20TokenExists(address l1ERC20, uint8 decimals) private returns (IArbToken) {
        address l2Contract = calculateBridgedERC20Address(l1ERC20);
        if (!l2Contract.isContract()) {
            address createdContract = create2Clone(templateERC20, bytes32(uint256(l1ERC20)));
            assert(createdContract == l2Contract);
            IArbToken(l2Contract).initialize(address(this), l1ERC20, decimals);
        }
        return IArbToken(l2Contract);
    }
}

contract ArbSymmetricTokenBridge is ArbTokenBridge {
    // assumes the L1 pair is deployed to the same address
    constructor() ArbTokenBridge(address(this)) public {}
}
