// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2021, Offchain Labs, Inc.
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

import "@openzeppelin/contracts/proxy/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

import "./Validator.sol";

contract ValidatorWalletCreator is Ownable {
    event WalletCreated(
        address indexed walletAddress,
        address indexed userAddress,
        address adminProxy
    );
    event TemplateUpdated();

    address public template;

    constructor() public Ownable() {
        template = address(new Validator());
    }

    function setTemplate(address _template) external onlyOwner {
        template = _template;
        emit TemplateUpdated();
    }

    function createWallet() external returns (address) {
        ProxyAdmin admin = new ProxyAdmin();
        address proxy =
            address(new TransparentUpgradeableProxy(address(template), address(admin), ""));
        admin.transferOwnership(msg.sender);
        Validator(proxy).initialize();
        Validator(proxy).transferOwnership(msg.sender);
        emit WalletCreated(proxy, msg.sender, address(admin));
        return proxy;
    }
}
