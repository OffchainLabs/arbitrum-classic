/*

  Copyright 2019 ZeroEx Intl.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

pragma solidity 0.5.15;

import "../../../utils/contracts/src/Ownable.sol";
import "../../../exchange-libs/contracts/src/LibExchangeRichErrors.sol";
import "../../../staking/contracts/src/interfaces/IStaking.sol";
import "./interfaces/IProtocolFees.sol";


contract MixinProtocolFees is
    IProtocolFees,
    Ownable
{
    /// @dev The protocol fee multiplier -- the owner can update this field.
    /// @return 0 Gas multplier.
    uint256 public protocolFeeMultiplier;

    /// @dev The address of the registered protocolFeeCollector contract -- the owner can update this field.
    /// @return 0 Contract to forward protocol fees to.
    address public protocolFeeCollector;

    /// @dev Allows the owner to update the protocol fee multiplier.
    /// @param updatedProtocolFeeMultiplier The updated protocol fee multiplier.
    function setProtocolFeeMultiplier(uint256 updatedProtocolFeeMultiplier)
        external
        onlyOwner
    {
        emit ProtocolFeeMultiplier(protocolFeeMultiplier, updatedProtocolFeeMultiplier);
        protocolFeeMultiplier = updatedProtocolFeeMultiplier;
    }

    /// @dev Allows the owner to update the protocolFeeCollector address.
    /// @param updatedProtocolFeeCollector The updated protocolFeeCollector contract address.
    function setProtocolFeeCollectorAddress(address updatedProtocolFeeCollector)
        external
        onlyOwner
    {
        _setProtocolFeeCollectorAddress(updatedProtocolFeeCollector);
    }

    /// @dev Sets the protocolFeeCollector contract address to 0.
    ///      Only callable by owner.
    function detachProtocolFeeCollector()
        external
        onlyOwner
    {
        _setProtocolFeeCollectorAddress(address(0));
    }

    /// @dev Sets the protocolFeeCollector address and emits an event.
    /// @param updatedProtocolFeeCollector The updated protocolFeeCollector contract address.
    function _setProtocolFeeCollectorAddress(address updatedProtocolFeeCollector)
        internal
    {
        emit ProtocolFeeCollectorAddress(protocolFeeCollector, updatedProtocolFeeCollector);
        protocolFeeCollector = updatedProtocolFeeCollector;
    }

    /// @dev Pays a protocol fee for a single fill.
    /// @param orderHash Hash of the order being filled.
    /// @param protocolFee Value of the fee being paid (equal to protocolFeeMultiplier * tx.gasPrice).
    /// @param makerAddress Address of maker of order being filled.
    /// @param takerAddress Address filling order.
    function _paySingleProtocolFee(
        bytes32 orderHash,
        uint256 protocolFee,
        address makerAddress,
        address takerAddress
    )
        internal
        returns (bool)
    {
        address feeCollector = protocolFeeCollector;
        if (feeCollector != address(0)) {
            _payProtocolFeeToFeeCollector(
                orderHash,
                feeCollector,
                address(this).balance,
                protocolFee,
                makerAddress,
                takerAddress
            );
            return true;
        } else {
            return false;
        }
    }

    /// @dev Pays a protocol fee for two orders (used when settling functions in MixinMatchOrders)
    /// @param orderHash1 Hash of the first order being filled.
    /// @param orderHash2 Hash of the second order being filled.
    /// @param protocolFee Value of the fee being paid (equal to protocolFeeMultiplier * tx.gasPrice).
    /// @param makerAddress1 Address of maker of first order being filled.
    /// @param makerAddress2 Address of maker of second order being filled.
    /// @param takerAddress Address filling orders.
    function _payTwoProtocolFees(
        bytes32 orderHash1,
        bytes32 orderHash2,
        uint256 protocolFee,
        address makerAddress1,
        address makerAddress2,
        address takerAddress
    )
        internal
        returns (bool)
    {
        address feeCollector = protocolFeeCollector;
        if (feeCollector != address(0)) {
            // Since the `BALANCE` opcode costs 400 gas, we choose to calculate this value by hand rather than calling it twice.
            uint256 exchangeBalance = address(this).balance;

            // Pay protocol fee and attribute to first maker
            uint256 valuePaid = _payProtocolFeeToFeeCollector(
                orderHash1,
                feeCollector,
                exchangeBalance,
                protocolFee,
                makerAddress1,
                takerAddress
            );

            // Pay protocol fee and attribute to second maker
            _payProtocolFeeToFeeCollector(
                orderHash2,
                feeCollector,
                exchangeBalance - valuePaid,
                protocolFee,
                makerAddress2,
                takerAddress
            );
            return true;
        } else {
            return false;
        }
    }

    /// @dev Pays a single protocol fee.
    /// @param orderHash Hash of the order being filled.
    /// @param feeCollector Address of protocolFeeCollector contract.
    /// @param exchangeBalance Assumed ETH balance of Exchange contract (in wei).
    /// @param protocolFee Value of the fee being paid (equal to protocolFeeMultiplier * tx.gasPrice).
    /// @param makerAddress Address of maker of order being filled.
    /// @param takerAddress Address filling order.
    function _payProtocolFeeToFeeCollector(
        bytes32 orderHash,
        address feeCollector,
        uint256 exchangeBalance,
        uint256 protocolFee,
        address makerAddress,
        address takerAddress
    )
        internal
        returns (uint256 valuePaid)
    {
        // Do not send a value with the call if the exchange has an insufficient balance
        // The protocolFeeCollector contract will fallback to charging WETH
        if (exchangeBalance >= protocolFee) {
            valuePaid = protocolFee;
        }
        bytes memory payProtocolFeeData = abi.encodeWithSelector(
            IStaking(address(0)).payProtocolFee.selector,
            makerAddress,
            takerAddress,
            protocolFee
        );
        // solhint-disable-next-line avoid-call-value
        (bool didSucceed, bytes memory returnData) = feeCollector.call.value(valuePaid)(payProtocolFeeData);
        if (!didSucceed) {
            revert();
        }
        return valuePaid;
    }
}
