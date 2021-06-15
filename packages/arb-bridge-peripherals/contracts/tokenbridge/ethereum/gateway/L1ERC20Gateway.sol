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

import "./L1ArbitrumExtendedGateway.sol";
import "@openzeppelin/contracts/utils/Create2.sol";

/**
 * @title Layer 1 Gateway contract for bridging standard ERC20s
 * @notice This contract handles token deposits, holds the escrowed tokens on layer 1, and (ultimately) finalizes withdrawals.
 * @dev Any ERC20 that requires non-standard functionality should use a separate gateway.
 * Messages to layer 2 use the inbox's createRetryableTicket method.
 */
contract L1ERC20Gateway is L1ArbitrumExtendedGateway {
    // used for create2 address calculation
    bytes32 public cloneableProxyHash;
    // We don't use the solidity creationCode as it breaks when upgrading contracts
    // keccak256(type(ClonableBeaconProxy).creationCode);
    address public l2BeaconProxyFactory;

    function initialize(
        address _l2Counterpart,
        address _router,
        address _inbox,
        bytes32 _cloneableProxyHash,
        address _l2BeaconProxyFactory
    ) public virtual {
        L1ArbitrumExtendedGateway._initialize(_l2Counterpart, _router, _inbox);
        require(_cloneableProxyHash != bytes32(0), "INVALID_PROXYHASH");
        require(_l2BeaconProxyFactory != address(0), "INVALID_BEACON");
        cloneableProxyHash = _cloneableProxyHash;
        l2BeaconProxyFactory = _l2BeaconProxyFactory;
    }

    function postUpgradeInit(bytes32 _cloneableProxyHash, address _l2BeaconProxyFactory) external {
        // one time use method to initialize new fields from upgrade
        require(_cloneableProxyHash != bytes32(0), "INVALID_PROXYHASH");
        require(_l2BeaconProxyFactory != address(0), "INVALID_BEACON");
        require(cloneableProxyHash == bytes32(0), "ALREADY_INIT");
        require(l2BeaconProxyFactory == address(0), "ALREADY_INIT");
        cloneableProxyHash = _cloneableProxyHash;
        l2BeaconProxyFactory = _l2BeaconProxyFactory;
    }

    /**
     * @notice utility function used to perform external read-only calls.
     * @dev the result is returned even if the call failed, the L2 is expected to
     * identify and deal with this.
     * @return result bytes, even if the call failed.
     */
    function callStatic(address targetContract, bytes4 targetFunction)
        internal
        view
        returns (bytes memory)
    {
        (bool success, bytes memory res) =
            targetContract.staticcall(abi.encodeWithSelector(targetFunction));
        return res;
    }

    function getOutboundCalldata(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory _data
    ) public view virtual override returns (bytes memory outboundCalldata) {
        // TODO: cheaper to make static calls or save isDeployed to storage?
        bytes memory deployData =
            abi.encode(
                callStatic(_token, ERC20.name.selector),
                callStatic(_token, ERC20.symbol.selector),
                callStatic(_token, ERC20.decimals.selector)
            );

        outboundCalldata = abi.encodeWithSelector(
            ITokenGateway.finalizeInboundTransfer.selector,
            _token,
            _from,
            _to,
            _amount,
            abi.encode(deployData, _data)
        );

        return outboundCalldata;
    }

    /**
     * @notice Calculate the address used when bridging an ERC20 token
     * @dev this always returns the same as the L1 oracle, but may be out of date.
     * For example, a custom token may have been registered but not deploy or the contract self destructed.
     * @param l1ERC20 address of L1 token
     * @return L2 address of a bridged ERC20 token
     */
    function calculateL2TokenAddress(address l1ERC20)
        external
        view
        virtual
        override
        returns (address)
    {
        // will revert if not called by router
        return _calculateL2TokenAddress(l1ERC20);
    }

    function _calculateL2TokenAddress(address l1ERC20)
        internal
        view
        virtual
        override
        returns (address)
    {
        bytes32 salt = getSalt(l1ERC20);
        return Create2.computeAddress(salt, cloneableProxyHash, l2BeaconProxyFactory);
    }

    function getSalt(address l1ERC20) internal view virtual returns (bytes32) {
        // TODO: use a library
        return keccak256(abi.encode(counterpartGateway, keccak256(abi.encode(l1ERC20))));
    }
}
