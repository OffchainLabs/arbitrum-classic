// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2019-2021, Offchain Labs, Inc.
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

contract Insurance {
    uint private constant ActionDelayInBlocks = 4*60*24*14;  // two weeks

    uint private totalShares;
    uint private numPayments;
    mapping(address => uint) private numShares;
    mapping(address => uint) private numCollected;

    uint private delayedActionNonce;
    mapping(bytes32 => bool) private delayedActions;

    enum DelayedActionType { Add, Remove }

    event AddedShares(address account, uint numShares);
    event RemovedShares(address account, uint numShares);
    event PaidOut(address account, uint amount);
    event InsuredEvent();

    function getTotalShares() external view returns(uint) {
        return totalShares;
    }

    function getNumEvents() external view returns(uint) {
        return numPayments;
    }

    function getAccountState(address account) external view returns(uint, uint) {
        return (numShares[account], numPayments-numCollected[account]);
    }

    function updateAndPayOut(address payable account) public {
        uint numToPay = numShares[account] * (numPayments - numCollected[account]);
        numCollected[account] = numPayments;
        if (numToPay > 0) {
            account.transfer(numToPay);
            emit PaidOut(account, numToPay);
        }
    }

    function addShares(address payable account, uint value) internal {
        require(msg.value > 0);
        updateAndPayOut(account);
        numShares[account] = numShares[account] + value;
        totalShares = totalShares + value;
        emit AddedShares(account, value);
    }

    function removeShares(address payable account, uint numToRemove) internal {
        require((numToRemove > 0) && (numToRemove <= numShares[account]));
        updateAndPayOut(account);
        numShares[account] = numShares[account] - numToRemove;
        if (numShares[account] == 0) {
            numCollected[account] = 0;  // this isn't needed, so zero it to reduce storage
        }
        totalShares = totalShares - numToRemove;
        account.transfer(numToRemove);
        emit RemovedShares(account, numToRemove);
    }

    function fundInsuredEvent() external payable {  // no reason to restrict who can call this
        require(msg.value == totalShares);
        numPayments = numPayments + 1;
        emit InsuredEvent();
    }

    function prepareDelayedAdd(address account) external payable returns(bytes32, uint, uint) {
        uint daNonce = delayedActionNonce++;
        uint readyBlock = block.number + ActionDelayInBlocks;
        bytes32 hashedAction = keccak256(
            abi.encodePacked(
                DelayedActionType.Add,
                account,
                msg.value,
                daNonce,
                readyBlock
            )
        );
        delayedActions[hashedAction] = true;
        return (hashedAction, daNonce, readyBlock);
    }

    function redeemDelayedAdd(address payable account, uint value, uint nonce, uint blockNum) external {
        bytes32 hashedAction = keccak256(
            abi.encodePacked(
                DelayedActionType.Add,
                account,
                value,
                nonce,
                blockNum
            )
        );
        require(delayedActions[hashedAction]);
        delayedActions[hashedAction] = false;
        addShares(account, value);
    }

    function prepareDelayedRemove(address payable account, uint num) external returns(bytes32, uint, uint) {
        removeShares(account, num);
        uint daNonce = delayedActionNonce++;
        uint readyBlock = block.number + ActionDelayInBlocks;
        bytes32 hashedAction = keccak256(
            abi.encodePacked(
                DelayedActionType.Remove,
                account,
                num,
                daNonce,
                readyBlock
            )
        );
        delayedActions[hashedAction] = true;
        return (hashedAction, daNonce, readyBlock);
    }

    function redeemDelayedRemove(address payable account, uint num, uint nonce, uint blockNum) external {
        bytes32 hashedAction = keccak256(
            abi.encodePacked(
                DelayedActionType.Remove,
                account,
                num,
                nonce,
                blockNum
            )
        );
        require(delayedActions[hashedAction]);
        delayedActions[hashedAction] = false;
        account.transfer(num);
    }

    function recommendNodeIntervalBlocks(uint challengePeriod, uint nodeGasCostEstimate) external returns(uint) {
        if (challengePeriod < 80) {
            return 40;
        }
        if (totalShares < 16 * nodeGasCostEstimate + 1) {  // add 1 so this is true if totalShares==0
            return challengePeriod / 2;
        }
        uint ratio = 4 * challengePeriod * challengePeriod * nodeGasCostEstimate / totalShares;
        if (ratio <= 1600) {
            return 40;
        } else {
            return approxSqrt(ratio, challengePeriod/4);
        }
    }

    function approxSqrt(uint x, uint startPoint) internal pure returns(uint) {
        // approximate sqrt(x), using Newton's method
        // should be very close to true answer, assuming startPoint / 2**15 < trueAnswer < startPoint * 2**15
        uint ret = startPoint;
        for (uint i=0; i<20; i++) {
            ret = (ret + x/ret) / 2;
        }
        return ret;
    }
}
