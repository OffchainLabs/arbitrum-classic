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

pragma solidity ^0.8.7;

import "./IGasRefunder.sol";

import "@openzeppelin/contracts-0.8/access/Ownable.sol";

contract GasRefunder is IGasRefunder, Ownable {
    mapping(address => bool) public allowedContracts;
    mapping(address => bool) public allowedRefundees;
    mapping(address => uint256) public lastContractRefund;
    address public disallower;
    uint256 public maxRefundeeBalance;
    uint256 public extraGasMargin;
    uint256 public calldataCost;
    uint256 public maxGasTip;
    uint256 public maxGasCost;
    uint256 public maxSingleGasUsage;

    event RefundedGasCosts(
        address indexed refundee,
        address indexed contractAddress,
        bool indexed success,
        uint256 gas,
        uint256 gasPrice,
        uint256 amountPaid
    );
    // Current reason IDs:
    // 0: Contract not allowed
    // 1: Refundee not allowed
    // 2: Contract already processed refund this block
    // 3: Refundee is already above max balance
    event RefundGasCostsDenied(
        address indexed refundee,
        address indexed contractAddress,
        uint256 gas,
        uint256 reason
    );
    event Deposited(address sender, uint256 amount);
    event Withdrawn(address initiator, address destination, uint256 amount);
    event ContractAllowedSet(address indexed addr, bool indexed allowed);
    event RefundeeAllowedSet(address indexed addr, bool indexed allowed);
    event DisallowerSet(address indexed addr);
    // Current parameter IDs:
    // 0: maxRefundeeBalance
    // 1: extraGasMargin
    // 2: calldataCost
    // 3: maxGasTip
    // 4: maxGasCost
    // 5: maxSingleGasUsage
    event ParameterSet(uint256 indexed parameter, uint256 value);

    constructor() Ownable() {
        extraGasMargin = 4000; // 4k gas
        calldataCost = 12; // Between 4 for 0 bytes and 16 for non-zero bytes
        maxGasTip = 2e9; // 2 gwei
        maxGasCost = 120e9; // 120 gwei
        maxSingleGasUsage = 2e6; // 2 million gas
    }

    function setDisallower(address addr) external onlyOwner {
        disallower = addr;
        emit DisallowerSet(addr);
    }

    function allowContracts(address[] calldata addresses) external onlyOwner {
        setContractsAllowedImpl(addresses, true);
    }

    function disallowContracts(address[] calldata addresses) external {
        require(msg.sender == owner() || msg.sender == disallower, "NOT_AUTHORIZED");
        setContractsAllowedImpl(addresses, false);
    }

    function setContractsAllowedImpl(address[] calldata addresses, bool allow) internal {
        for (uint256 i = 0; i < addresses.length; i++) {
            address addr = addresses[i];
            allowedContracts[addr] = allow;
            emit ContractAllowedSet(addr, allow);
        }
    }

    function allowRefundees(address[] calldata addresses) external onlyOwner {
        setRefundeesAllowedImpl(addresses, true);
    }

    function disallowRefundees(address[] calldata addresses) external {
        require(msg.sender == owner() || msg.sender == disallower, "NOT_AUTHORIZED");
        setRefundeesAllowedImpl(addresses, false);
    }

    function setRefundeesAllowedImpl(address[] calldata addresses, bool allow) internal {
        for (uint256 i = 0; i < addresses.length; i++) {
            address addr = addresses[i];
            allowedRefundees[addr] = allow;
            emit RefundeeAllowedSet(addr, allow);
        }
    }

    function setMaxRefundeeBalance(uint256 newValue) external onlyOwner {
        maxRefundeeBalance = newValue;
        emit ParameterSet(0, newValue);
    }

    function setExtraGasMargin(uint256 newValue) external onlyOwner {
        extraGasMargin = newValue;
        emit ParameterSet(1, newValue);
    }

    function setCalldataCost(uint256 newValue) external onlyOwner {
        calldataCost = newValue;
        emit ParameterSet(2, newValue);
    }

    function setMaxGasTip(uint256 newValue) external onlyOwner {
        maxGasTip = newValue;
        emit ParameterSet(3, newValue);
    }

    function setMaxGasCost(uint256 newValue) external onlyOwner {
        maxGasCost = newValue;
        emit ParameterSet(4, newValue);
    }

    function setMaxSingleGasUsage(uint256 newValue) external onlyOwner {
        maxSingleGasUsage = newValue;
        emit ParameterSet(5, newValue);
    }

    receive() external payable {
        emit Deposited(msg.sender, msg.value);
    }

    function withdraw(address payable destination, uint256 amount) external onlyOwner {
        (bool success, ) = destination.call{ value: amount }("");
        require(success, "WITHDRAW_FAILED");
        emit Withdrawn(msg.sender, destination, amount);
    }

    function onGasSpent(
        address payable refundee,
        uint256 gasUsed,
        uint256 calldataSize
    ) external override returns (bool success) {
        uint256 startGasLeft = gasleft();

        if (!allowedContracts[msg.sender]) {
            emit RefundGasCostsDenied(refundee, msg.sender, gasUsed, 0);
            return false;
        }
        if (!allowedRefundees[refundee]) {
            emit RefundGasCostsDenied(refundee, msg.sender, gasUsed, 1);
            return false;
        }

        if (lastContractRefund[msg.sender] == block.number) {
            // There was already a refund this block, don't refund further
            emit RefundGasCostsDenied(refundee, msg.sender, gasUsed, 2);
            return false;
        }
        lastContractRefund[msg.sender] = block.number;

        uint256 estGasPrice = block.basefee + maxGasTip;
        if (tx.gasprice < estGasPrice) {
            estGasPrice = tx.gasprice;
        }
        if (maxGasCost != 0 && estGasPrice > maxGasCost) {
            estGasPrice = maxGasCost;
        }

        // Cache these variables and retrieve them before measuring gasleft()
        uint256 refundeeBalance = refundee.balance;
        uint256 maxRefundeeBalanceCache = maxRefundeeBalance;
        uint256 maxSingleGasUsageCached = maxSingleGasUsage;

        // Add in a bit of a buffer for the tx costs not measured with gasleft
        gasUsed += startGasLeft + extraGasMargin + (calldataSize * calldataCost);
        // Split this up into two statements so that gasleft() comes after the storage loads
        gasUsed -= gasleft();

        if (maxSingleGasUsageCached != 0 && gasUsed > maxSingleGasUsageCached) {
            gasUsed = maxSingleGasUsageCached;
        }

        uint256 refundAmount = estGasPrice * gasUsed;
        if (
            maxRefundeeBalanceCache != 0 && refundeeBalance + refundAmount > maxRefundeeBalanceCache
        ) {
            if (refundeeBalance > maxRefundeeBalanceCache) {
                // The refundee is already above their max balance
                emit RefundGasCostsDenied(refundee, msg.sender, gasUsed, 3);
                return false;
            } else {
                refundAmount = maxRefundeeBalanceCache - refundeeBalance;
            }
        }

        (success, ) = refundee.call{ value: refundAmount }("");
        emit RefundedGasCosts(refundee, msg.sender, success, gasUsed, estGasPrice, refundAmount);
    }
}
