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

import "@openzeppelin/contracts/proxy/UpgradeableBeacon.sol";
import "@openzeppelin/contracts/utils/Create2.sol";
import "./L2ArbitrumGateway.sol";
import "../StandardArbERC20.sol";
import "../../libraries/ClonableBeaconProxy.sol";

contract L2ERC20Gateway is L2ArbitrumGateway {
    address public beaconProxyFactory;

    function initialize(
        address _l1Counterpart,
        address _router,
        address _beaconProxyFactory
    ) public {
        L2ArbitrumGateway._initialize(_l1Counterpart, _router);
        require(_beaconProxyFactory != address(0), "INVALID_BEACON");
        beaconProxyFactory = _beaconProxyFactory;
    }

    /**
     * @notice Calculate the address used when bridging an ERC20 token
     * @dev the L1 and L2 address oracles may not always be in sync.
     * For example, a custom token may have been registered but not deploy or the contract self destructed.
     * @param l1ERC20 address of L1 token
     * @return L2 address of a bridged ERC20 token
     */
    function calculateL2TokenAddress(address l1ERC20)
        public
        view
        virtual
        override
        returns (address)
    {
        // this method is marked virtual to be overriden in subclasses used in testing
        return
            BeaconProxyFactory(beaconProxyFactory).calculateExpectedAddress(
                address(this),
                getUserSalt(l1ERC20)
            );
    }

    function cloneableProxyHash() public view returns (bytes32) {
        return BeaconProxyFactory(beaconProxyFactory).cloneableProxyHash();
    }

    function getUserSalt(address l1ERC20) public pure returns (bytes32) {
        return keccak256(abi.encode(l1ERC20));
    }

    /**
     * @notice internal utility function used to deploy ERC20 tokens with the beacon proxy pattern.
     * @dev the transparent proxy implementation by OpenZeppelin can't be used if we want to be able to
     * upgrade the token logic.
     * @param l1ERC20 L1 address of ERC20
     * @param expectedL2Address L2 address of ERC20
     * @param deployData encoded symbol/name/decimal data for initial deploy
     */
    function handleNoContract(
        address l1ERC20,
        address expectedL2Address,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory deployData
    ) internal override returns (bool shouldHalt) {
        bytes32 userSalt = getUserSalt(l1ERC20);
        address createdContract = BeaconProxyFactory(beaconProxyFactory).createProxy(userSalt);

        StandardArbERC20(createdContract).bridgeInit(l1ERC20, deployData);

        if (createdContract == expectedL2Address) {
            return false;
        } else {
            // trigger withdrawal then halt
            // this codepath should only be hit if the system is setup incorrectly
            // this withdrawal is for error recovery, not composing with L2 dapps, so we ignore the return value
            triggerWithdrawal(l1ERC20, address(this), _from, _amount, "");
            return true;
        }
    }
}
