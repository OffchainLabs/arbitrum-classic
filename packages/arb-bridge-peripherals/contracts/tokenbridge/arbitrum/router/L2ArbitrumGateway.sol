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

import "../../libraries/ITokenGateway.sol";
import "../../libraries/TokenGateway.sol";
import "../../libraries/IERC677.sol";

abstract contract L2ArbitrumGateway is TokenGateway {
    using Address for address;

    address internal constant arbsysAddr = address(100);

    uint256 public exitNum;

    function isCounterpartGateway() internal view virtual override returns (bool) {
        return msg.sender == counterpartGateway;
    }

    function _initialize(address _l1Counterpart, address _router) internal virtual override {
        // L2 gateway may have a router address(0)
        TokenGateway._initialize(_l1Counterpart, _router);
    }

    function arbgasReserveIfCallRevert() internal pure virtual returns (uint256) {
        // amount of arbgas necessary to send user tokens in case
        // of the "onTokenTransfer" call consumes all available gas
        return 2500;
    }

    function createOutboundTx(bytes memory _data) internal virtual returns (uint256) {
        uint256 id = ArbSys(arbsysAddr).sendTxToL1(counterpartGateway, _data);
        return id;
    }

    function getOutboundCalldata(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes memory _data
    ) public view virtual override returns (bytes memory outboundCalldata) {
        outboundCalldata = abi.encodeWithSelector(
            ITokenGateway.finalizeInboundTransfer.selector,
            _token,
            _from,
            _to,
            _amount,
            abi.encode(exitNum, _data)
        );

        return outboundCalldata;
    }

    function outboundTransfer(
        address _l1Token,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) public payable virtual returns (bytes memory) {
        return outboundTransfer(_l1Token, _to, _amount, 0, 0, _data);
    }

    /**
     * @notice Initiates a token withdrawal from Arbitrum to Ethereum
     * @param _l1Token l1 address of token
     * @param _to destination address
     * @param _amount amount of tokens withdrawn
     * @param _maxGas max gas provided for outbox execution market (todo)
     * @param _gasPriceBid provided for outbox execution market (todo)
     @ @return encoded unique identifier for withdrawal
     */

    function outboundTransfer(
        address _l1Token,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes calldata _data
    ) public payable virtual override returns (bytes memory) {
        // can be triggered directly or by router
        require(msg.value == 0, "NO_VALUE");

        (address _from, bytes memory _extraData) = parseOutboundData(_data);

        uint256 id;
        {
            address l2Token = _calculateL2TokenAddress(_l1Token);
            require(l2Token.isContract(), "TOKEN_NOT_DEPLOYED");
            // burns L2 tokens in order to release escrowed L1 tokens
            IArbStandardToken(l2Token).bridgeBurn(_from, _amount);

            bytes memory outboundCalldata =
                getOutboundCalldata(_l1Token, _from, _to, _amount, _extraData);

            id = createOutboundTx(outboundCalldata);
        }
        // exitNum incremented after being used in getOutboundCalldata
        exitNum++;
        emit OutboundTransferInitiated(_l1Token, _from, _to, id, _amount, _extraData);
        return abi.encode(id);
    }

    function parseOutboundData(bytes memory _data)
        internal
        view
        virtual
        returns (address _from, bytes memory _extraData)
    {
        if (isRouter()) {
            (_from, _extraData) = abi.decode(_data, (address, bytes));
        } else {
            _from = msg.sender;
            _extraData = _data;
        }
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
        uint256 gasAvailable = gasleft() - arbgasReserveIfCallRevert();
        require(gasleft() > gasAvailable, "Mint and call gas left calculation undeflow");

        IERC677Receiver(dest).onTokenTransfer{ gas: gasAvailable }(sender, amount, data);
    }

    /**
     * @notice Mint on L2 upon L1 deposit.
     * If token not yet deployed and symbol/name/decimal data is included, deploys StandardArbERC20
     * @dev Callable only by the L1ERC20Gateway.outboundTransfer method. For initial deployments of a token the L1 L1ERC20Gateway
     * is expected to include the deployData. If not a L1 withdrawal is automatically triggered for the user
     * @param _token L1 address of ERC20
     * @param _from account that initiated the deposit in the L1
     * @param _to account to be credited with the tokens in the L2 (can be the user's L2 account or a contract)
     * @param _amount token amount to be minted to the user
     * @param _data encoded symbol/name/decimal data for deploy, in addition to any additional callhook data
     */
    function finalizeInboundTransfer(
        address _token,
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) external virtual override onlyCounterpartGateway returns (bytes memory) {
        (bytes memory deployData, bytes memory callHookData) = abi.decode(_data, (bytes, bytes));

        address expectedAddress = _calculateL2TokenAddress(_token);

        if (!expectedAddress.isContract()) {
            bool shouldHalt = handleNoContract(_token, expectedAddress, deployData);
            if (shouldHalt) return bytes("");
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
            emit TransferAndCallTriggered(success, _from, _to, _amount, callHookData);
        } else {
            token.bridgeMint(_to, _amount);
        }

        emit InboundTransferFinalized(
            _token,
            _from,
            _to,
            uint256(uint160(expectedAddress)),
            _amount,
            _data
        );

        return bytes("");
    }

    function isRouter() internal view virtual override returns (bool) {
        return msg.sender == router;
    }

    // returns if function should halt after
    function handleNoContract(
        address l1ERC20,
        address expectedL2Address,
        bytes memory data
    ) internal virtual returns (bool shouldHalt);
}
