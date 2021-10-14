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
import "arb-bridge-eth/contracts/libraries/Whitelist.sol";

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
    // whitelist not used anymore
    address public whitelist;

    // start of inline reentrancy guard
    // https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v3.4.2/contracts/utils/ReentrancyGuard.sol
    uint256 private constant _NOT_ENTERED = 1;
    uint256 private constant _ENTERED = 2;
    uint256 private _status;

    modifier nonReentrant() {
        // On the first call to nonReentrant, _notEntered will be true
        require(_status != _ENTERED, "ReentrancyGuard: reentrant call");
        // Any calls to nonReentrant after this point will fail
        _status = _ENTERED;
        _;
        _status = _NOT_ENTERED;
    }

    // end of inline reentrancy guard

    function postUpgradeInit() external override {
        _status = _NOT_ENTERED;
    }

    function outboundTransfer(
        address _l1Token,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes calldata _data
    ) public payable override nonReentrant returns (bytes memory res) {
        return super.outboundTransfer(_l1Token, _to, _amount, _maxGas, _gasPriceBid, _data);
    }

    function finalizeInboundTransfer(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) public payable virtual override nonReentrant {
        // the superclass checks onlyCounterpartGateway
        super.finalizeInboundTransfer(_token, _from, _to, _amount, _data);
    }

    function initialize(
        address _l2Counterpart,
        address _router,
        address _inbox,
        bytes32 _cloneableProxyHash,
        address _l2BeaconProxyFactory
    ) public {
        L1ArbitrumExtendedGateway._initialize(_l2Counterpart, _router, _inbox);
        require(_cloneableProxyHash != bytes32(0), "INVALID_PROXYHASH");
        require(_l2BeaconProxyFactory != address(0), "INVALID_BEACON");
        cloneableProxyHash = _cloneableProxyHash;
        l2BeaconProxyFactory = _l2BeaconProxyFactory;
        // disable whitelist by default
        whitelist = address(0);
        // reentrancy guard
        _status = _NOT_ENTERED;
    }

    /**
     * @notice utility function used to perform external read-only calls.
     * @dev the result is returned even if the call failed or was directed at an EOA,
     * it is cheaper to have the L2 consumer identify and deal with this.
     * @return result bytes, even if the call failed.
     */
    function callStatic(address targetContract, bytes4 targetFunction)
        internal
        view
        returns (bytes memory)
    {
        (
            ,
            /* bool success */
            bytes memory res
        ) = targetContract.staticcall(abi.encodeWithSelector(targetFunction));
        return res;
    }

    function getOutboundCalldata(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory _data
    ) public view override returns (bytes memory outboundCalldata) {
        // TODO: cheaper to make static calls or save isDeployed to storage?
        bytes memory deployData = abi.encode(
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
            GatewayMessageHandler.encodeToL2GatewayMsg(deployData, _data)
        );

        return outboundCalldata;
    }

    function calculateL2TokenAddress(address l1ERC20) public view override returns (address) {
        bytes32 salt = getSalt(l1ERC20);
        return Create2.computeAddress(salt, cloneableProxyHash, l2BeaconProxyFactory);
    }

    function getSalt(address l1ERC20) internal view returns (bytes32) {
        // TODO: use a library
        return keccak256(abi.encode(counterpartGateway, keccak256(abi.encode(l1ERC20))));
    }
}
