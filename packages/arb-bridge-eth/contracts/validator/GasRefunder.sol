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
    event MaxRefundeeBalanceSet(uint256 balance);
    event ExtraGasMarginSet(uint256 balance);

    function setContractAllowed(address contractAddress, bool allowed) external onlyOwner {
        allowedContracts[contractAddress] = allowed;
        emit ContractAllowedSet(contractAddress, allowed);
    }

    function setMaxRefundeeBalance(uint256 newMax) external onlyOwner {
        maxRefundeeBalance = newMax;
        emit MaxRefundeeBalanceSet(newMax);
    }

    function extraGasMarginSet(uint256 newMargin) external onlyOwner {
        extraGasMargin = newMargin;
        emit ExtraGasMarginSet(newMargin);
    }

    receive() external payable {
        emit Deposited(msg.sender, msg.value);
    }

    function withdraw(address payable destination, uint256 amount) external onlyOwner {
        destination.transfer(amount);
        emit Withdrawn(msg.sender, destination, amount);
    }

    function onGasSpent(address payable refundee, uint256 gasUsed) external override {
        uint256 startGasLeft = gasleft();
        require(allowedContracts[msg.sender], "NOT_ALLOWED_CONTRACT");

        if (lastContractRefund[msg.sender] == block.number) {
            // There was already a refund this block, don't refund further
            emit RefundGasCostsDenied(refundee, msg.sender, gasUsed, 0);
            return;
        }
        lastContractRefund[msg.sender] = block.number;

        // Allow for up to a 2 gwei tip
        uint256 estGasPrice = block.basefee + 2e9;
        if (tx.gasprice < estGasPrice) {
            estGasPrice = tx.gasprice;
        }

        // Cache these variables and retrieve them before measuring gasleft()
        uint256 spenderBalance = refundee.balance;
        uint256 maxRefundeeBalanceCache = maxRefundeeBalance;

        // Add in a bit of a buffer for the tx costs not measured with gasleft
        gasUsed += startGasLeft + extraGasMargin;
        // Split this up into two statements so that gasleft() comes after the storage load of extraGasMargin
        gasUsed -= gasleft();

        uint256 refundAmount = estGasPrice * gasUsed;
        if (
            maxRefundeeBalanceCache != 0 && spenderBalance + refundAmount > maxRefundeeBalanceCache
        ) {
            if (spenderBalance > maxRefundeeBalanceCache) {
                // The refundee is already above their max balance
                emit RefundGasCostsDenied(refundee, msg.sender, gasUsed, 1);
                return;
            } else {
                refundAmount = maxRefundeeBalanceCache - spenderBalance;
            }
        }

        bool success = refundee.send(refundAmount);
        emit RefundedGasCosts(refundee, msg.sender, gasUsed, refundAmount, success);
    }
}
