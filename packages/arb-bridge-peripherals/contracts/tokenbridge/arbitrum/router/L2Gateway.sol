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

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "arbos-contracts/arbos/builtin/ArbSys.sol";

import "../IArbStandardToken.sol";

import "../../libraries/ClonableBeaconProxy.sol";
import "../../libraries/ITokenGateway.sol";
import "../../libraries/TokenGateway.sol";
import "../../libraries/IERC677.sol";

abstract contract L2ArbitrumGateway is TokenGateway {
    function initialize(address _target) public virtual override {
        super.initialize(_target);
    }

    function createOutboundTx(
        address _handler,
        address _user,
        uint256 _maxSubmissionCost,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes memory _data
    ) internal virtual override returns (bytes memory) {
        uint256 id = ArbSys(_handler).sendTxToL1(counterpartGateway, _data);
        return abi.encode(id);
    }
}

contract L2ERC20Gateway is L2ArbitrumGateway, ProxySetter {
    using Address for address;

    // amount of arbgas necessary to send user tokens in case
    // of the "onTokenTransfer" call consumes all available gas
    uint256 internal constant arbgasReserveIfCallRevert = 2500;
    bytes32 public constant cloneableProxyHash = keccak256(type(ClonableBeaconProxy).creationCode);

    /**
     * @notice utility function used in ClonableBeaconProxy.
     * @dev this method makes it possible to use ClonableBeaconProxy.creationCode without encoding constructor parameters
     * @return the beacon to be used by the proxy contract.
     */
    address public override beacon;

    function initialize(address _target, address _beacon) public virtual {
        super.initialize(_target);
        require(_beacon != address(0), "INVALID_BEACON");
        require(beacon == address(0), "ALREADY_INIT");
        beacon = _beacon;
    }

    function outboundTransfer(
        address _token,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes calldata _data
    ) external payable virtual override returns (bytes memory res) {
        require(msg.value == 0, "NO_VALUE");

        address _from = msg.sender;

        (address l1TokenAddr, bytes memory extraData) = abi.decode(_data, (address, bytes));

        address expectedTokenAddr = calculateL2TokenAddress(l1TokenAddr);
        require(_token == expectedTokenAddr, "WRONG_TOKEN_ADDR");

        handleEscrow(_token, _from, _amount);
        bytes memory outboundCalldata =
            getOutboundCalldata(l1TokenAddr, _from, _to, _amount, extraData);

        res = createOutboundTx(address(100), _from, 0, _maxGas, _gasPriceBid, outboundCalldata);

        emit OutboundTransferInitiated(l1TokenAddr, _from, _to, _amount, _data);

        return res;
    }

    function getSalt(address l1ERC20) internal pure virtual returns (bytes32) {
        return keccak256(abi.encode(l1ERC20));
    }

    /**
     * @notice Calculate the address used when bridging an ERC20 token
     * @dev this always returns the same as the L1 oracle, but may be out of date.
     * For example, a custom token may have been registered but not deploy or the contract self destructed.
     * @param l1ERC20 address of L1 token
     * @return L2 address of a bridged ERC20 token
     */
    function calculateL2TokenAddress(address l1ERC20) public view virtual returns (address) {
        bytes32 salt = getSalt(l1ERC20);
        return Create2.computeAddress(salt, cloneableProxyHash, address(this));
    }

    /**
     * @notice this function can only be callable by the bridge itself
     * @dev This method is inspired by EIP 677/1363 for calls to be executed after minting.
     * A reserve amount of gas is always kept in case this call reverts or uses up all gas.
     * The reserve is the amount of gas needed to catch the revert and do the necessary alternative logic.
     */
    function mintAndCall(
        IArbToken token,
        uint256 amount,
        address sender,
        address dest,
        bytes memory data
    ) external {
        require(msg.sender == address(this), "Mint can only be called by self");
        require(dest.isContract(), "Destination must be a contract");

        token.bridgeMint(dest, amount);

        // ~73 000 arbgas used to get here
        uint256 gasAvailable = gasleft() - arbgasReserveIfCallRevert;
        require(gasleft() > gasAvailable, "Mint and call gas left calculation undeflow");

        IERC677Receiver(dest).onTokenTransfer{ gas: gasAvailable }(sender, amount, data);
    }

    function finalizeInboundTransfer(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) external virtual override onlyCounterpartGateway returns (bytes memory) {
        (bytes memory deployData, bytes memory callHookData) = abi.decode(_data, (bytes, bytes));

        address expectedAddress = calculateL2TokenAddress(_token);

        if (!expectedAddress.isContract()) {
            address deployedToken = deployToken(_token, deployData);
            assert(deployedToken == expectedAddress);
        }
        // ignores deployData if token already deployed

        IArbToken token = IArbToken(expectedAddress);
        if (callHookData.length > 0) {
            bool success;
            try this.mintAndCall(token, _amount, _from, _to, callHookData) {
                success = true;
            } catch {
                // if reverted, then credit _from's account
                token.bridgeMint(_from, _amount);
                // success default value is false
            }
            // if success tokens got minted to _to, else to _from
            // emit TokenMinted(
            //     l1ERC20,
            //     expectedAddress,
            //     _from,
            //     success ? _to : _from,
            //     _amount,
            //     true
            // );
            // emit MintAndCallTriggered(success, _from, _to, _amount, callHookData);
        } else {
            token.bridgeMint(_to, _amount);
            // emit TokenMinted(l1ERC20, expectedAddress, _from, _to, _amount, false);
        }

        return bytes("");
    }

    /**
     * @notice internal utility function used to deploy ERC20 tokens with the beacon proxy pattern.
     * @dev the transparent proxy implementation by OpenZeppelin can't be used if we want to be able to
     * upgrade the token logic.
     * @param l1ERC20 L1 address of ERC20
     * @param deployData encoded symbol/name/decimal data for initial deploy
     */
    function deployToken(address l1ERC20, bytes memory deployData) internal returns (address) {
        bytes32 salt = getSalt(l1ERC20);
        address createdContract = address(new ClonableBeaconProxy{ salt: salt }());

        IArbStandardToken(createdContract).bridgeInit(l1ERC20, deployData);

        // emit TokenCreated(l1ERC20, createdContract);
        return createdContract;
    }

    // make it public so it can be used internally and externally for gas estimation
    function getOutboundCalldata(
        address _l1Token,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory _data
    ) public view virtual override returns (bytes memory) {
        return
            abi.encodeWithSelector(
                ITokenGateway.finalizeInboundTransfer.selector,
                _l1Token,
                _from,
                _to,
                _amount,
                _data
            );
    }

    function handleEscrow(
        address _token,
        address _from,
        uint256 _amount
    ) internal virtual override {
        IArbStandardToken(_token).bridgeBurn(_from, _amount);
    }

    // TODO: add transferAndCall affordance
}
