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

import "./IRollup.sol";

contract Insurance {
    uint private constant ActionDelayInBlocks = 4*60*24*14;  // two weeks
    uint private constant ENDMARKER = 1<<255;

    IRollup private rollup;

    // information about nodes (pending nodes, latestConfirmed, and older nodes that had challenges on their children)
    uint private lastNodeCreated;
    uint private lastNodeConfirmed;
    uint private lastChallengeNodeAsOfLatestResolved;
    uint private lastChallengeNodeAsOfLatestConfirmed;
    int private changeInAmount;
    mapping(uint => uint) numSharesAtNode;
    mapping(uint => uint) numChallengesOnNode;
    mapping(uint => uint) nextNodeWithChallenge;

    // information about insurance depositors
    mapping(address => uint) lastNodeCollected;
    mapping(address => uint) firstNodeEligible;
    mapping(address => uint) lastNodeEligible;
    mapping(address => uint) amountDeposited;

    modifier onlyRollup {
        require(msg.sender == address(rollup), "ROLLUP_ONLY");
        _;
    }

    constructor(IRollup rollup_) public {
        rollup = rollup_;
        lastNodeCreated = 1;
        lastNodeConfirmed = 0;
        lastChallengeNodeAsOfLatestResolved = 0;
        lastChallengeNodeAsOfLatestConfirmed = 0;
        changeInAmount = 0;
        nextNodeWithChallenge[0] = ENDMARKER;
    }

    function notifyNewNode(
        uint challengePeriod,
        uint nodeCostGasEstimate
    ) external onlyRollup returns(uint) {
        uint numShares = uint(int(numSharesAtNode[lastNodeCreated]) + changeInAmount);
        changeInAmount = 0;
        lastNodeCreated++;
        numSharesAtNode[lastNodeCreated] = numShares;
        return recommendNodeIntervalBlocks(challengePeriod, numShares, nodeCostGasEstimate);
    }

    function notifyNodeResolved(uint nodeNum, bool confirmed) external onlyRollup {
        if (confirmed) {
            lastNodeConfirmed = nodeNum;
            if (numChallengesOnNode[nodeNum] > 0) {
                nextNodeWithChallenge[lastChallengeNodeAsOfLatestResolved] = nodeNum;
                lastChallengeNodeAsOfLatestResolved = nodeNum;
            }
            lastChallengeNodeAsOfLatestConfirmed = lastChallengeNodeAsOfLatestResolved;
            numSharesAtNode[lastNodeConfirmed] = 0;   // zero unneeded storage
        } else {
            if (numChallengesOnNode[nodeNum] > 0) {
                nextNodeWithChallenge[lastChallengeNodeAsOfLatestResolved] = nodeNum;
                lastChallengeNodeAsOfLatestResolved = nodeNum;
            }
            numSharesAtNode[nodeNum] = 0;    // zero unneeded storage
        }
    }

    function notifyChallenge(uint parentNodeNum) external onlyRollup {
        require( (parentNodeNum >= lastNodeConfirmed) || (parentNodeNum <= lastNodeCreated), "bad node number");
        numChallengesOnNode[parentNodeNum]++;
        if (parentNodeNum == lastNodeConfirmed) {
            nextNodeWithChallenge[lastChallengeNodeAsOfLatestConfirmed] = parentNodeNum;
        }
    }

    function addStake() external payable {
        require(amountDeposited[msg.sender] == 0, "already a depositor");
        lastNodeCollected[msg.sender] = lastChallengeNodeAsOfLatestResolved;
        firstNodeEligible[msg.sender] = lastNodeCreated + 1;
        lastNodeEligible[msg.sender] = ENDMARKER;
        amountDeposited[msg.sender] = msg.value;
        changeInAmount = changeInAmount + int(msg.value);
    }

    function scheduleStakeRemoval() external {
        require(amountDeposited[msg.sender] > 0, "not a depositor");
        require(lastNodeEligible[msg.sender] < ENDMARKER, "already scheduled");
        lastNodeEligible[msg.sender] = lastNodeCreated;
        changeInAmount = changeInAmount - int(amountDeposited[msg.sender]);
    }

    function recoverStake(uint maxNodesToPay) external {
        require(amountDeposited[msg.sender] > 0, "not a depositor");
        require(lastNodeEligible[msg.sender] < lastNodeConfirmed);

        require(payOut(msg.sender, maxNodesToPay));   // if this fails, sender should call payOut, then try again

        lastNodeCollected[msg.sender] = 0;
        firstNodeEligible[msg.sender] = 0;
        lastNodeEligible[msg.sender] = 0;
        uint amount = amountDeposited[msg.sender];
        amountDeposited[msg.sender] = 0;
        msg.sender.transfer(amount);
    }

    function payOut(address payable account, uint maxNumNodes) public returns(bool) {  // returns true iff done paying for now
        require(amountDeposited[account] > 0, "not a depositor");
        uint numOwed = 0;
        uint nodeNum = lastNodeCollected[account];
        while(nodeNum < lastNodeConfirmed) {
            uint next = nextNodeWithChallenge[nodeNum];
            if (next < lastNodeConfirmed) {
                if (nodeNum >= firstNodeEligible[account]) {  // make sure node is after account made its deposit
                    numOwed = numOwed++;
                }
                nodeNum = next;
                if (maxNumNodes > 1) {
                    maxNumNodes--;
                } else {
                    lastNodeCollected[account] = nodeNum;
                    account.transfer(numOwed * amountDeposited[account]);
                    return false;
                }
            }
        }
        lastNodeCollected[account] = nodeNum;
        account.transfer(numOwed * amountDeposited[account]);
        return true;
    }

    function recommendNodeIntervalBlocks(
        uint challengePeriod,
        uint totalDeposited,
        uint nodeGasCostEstimate
    ) internal pure returns(uint) {
        if (challengePeriod < 80) {
            return 40;
        }
        if (totalDeposited < 16 * nodeGasCostEstimate + 1) {  // add 1 so this is true if totalDeposited==0
            return challengePeriod / 2;
        }
        uint ratio = 4 * challengePeriod * challengePeriod * nodeGasCostEstimate / totalDeposited;
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
