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
    mapping(address => uint256) public lastContractRefund;
    uint256 public maxRefundeeBalance;
    uint256 public extraGasMargin;
    uint256 public calldataCost;
    uint256 public maxGasTip;
    uint256 public maxSingleRefund;

    event RefundedGasCosts(
        address indexed refundee,
        address indexed contractAddress,
        uint256 gas,
        uint256 amountPaid,
        bool success
    );
    // Reason can currently be 0 for contract already refunded this block,
    // or 1 for the refundee is already over the max refundee balance.
    event RefundGasCostsDenied(
        address indexed refundee,
        address indexed contractAddress,
        uint256 gas,
        uint256 reason
    );
    event Deposited(address sender, uint256 amount);
    event Withdrawn(address initiator, address destination, uint256 amount);
    event ContractAllowedSet(address indexed contractAddress, bool allowed);
    // Current parameter IDs:
    // 0: maxRefundeeBalance
    // 1: extraGasMargin
    // 2: calldataCost
    // 3: maxGasTip
    // 4: maxSingleRefund
    event ParameterSet(uint256 indexed parameter, uint256 value);

    constructor() Ownable() {
        extraGasMargin = 4000;
        calldataCost = 12; // Between 4 for 0 bytes and 16 for non-zero bytes
        maxGasTip = 2e9; // 2 gwei
        maxSingleRefund = 2e17; // 0.2 ETH
    }

    function setContractAllowed(address contractAddress, bool allowed) external onlyOwner {
        allowedContracts[contractAddress] = allowed;
        emit ContractAllowedSet(contractAddress, allowed);
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

    function setMaxSingleRefund(uint256 newValue) external onlyOwner {
        maxSingleRefund = newValue;
        emit ParameterSet(4, newValue);
    }

    receive() external payable {
        emit Deposited(msg.sender, msg.value);
    }

    function withdraw(address payable destination, uint256 amount) external onlyOwner {
        destination.transfer(amount);
        emit Withdrawn(msg.sender, destination, amount);
    }

    function onGasSpent(
        address payable refundee,
        uint256 gasUsed,
        uint256 calldataSize
    ) external override {
        uint256 startGasLeft = gasleft();
        require(allowedContracts[msg.sender], "NOT_ALLOWED_CONTRACT");

        if (lastContractRefund[msg.sender] == block.number) {
            // There was already a refund this block, don't refund further
            emit RefundGasCostsDenied(refundee, msg.sender, gasUsed, 0);
            return;
        }
        lastContractRefund[msg.sender] = block.number;

        uint256 estGasPrice = block.basefee + maxGasTip;
        if (tx.gasprice < estGasPrice) {
            estGasPrice = tx.gasprice;
        }

        // Cache these variables and retrieve them before measuring gasleft()
        uint256 refundeeBalance = refundee.balance;
        uint256 maxRefundeeBalanceCache = maxRefundeeBalance;
        uint256 maxSingleRefundCached = maxSingleRefund;

        // Add in a bit of a buffer for the tx costs not measured with gasleft
        gasUsed += startGasLeft + extraGasMargin + (calldataSize * calldataCost);
        // Split this up into two statements so that gasleft() comes after the storage load of extraGasMargin
        gasUsed -= gasleft();

        uint256 refundAmount = estGasPrice * gasUsed;
        if (
            maxRefundeeBalanceCache != 0 && refundeeBalance + refundAmount > maxRefundeeBalanceCache
        ) {
            if (refundeeBalance > maxRefundeeBalanceCache) {
                // The refundee is already above their max balance
                emit RefundGasCostsDenied(refundee, msg.sender, gasUsed, 1);
                return;
            } else {
                refundAmount = maxRefundeeBalanceCache - refundeeBalance;
            }
        }
        if (maxSingleRefundCached != 0 && maxSingleRefundCached < refundAmount) {
            refundAmount = maxSingleRefundCached;
        }

        bool success = refundee.send(refundAmount);
        emit RefundedGasCosts(refundee, msg.sender, gasUsed, refundAmount, success);
    }
}
