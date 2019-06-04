/*
 * Copyright 2019, Offchain Labs, Inc.
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

import MerkleTree from 'merkle-tree-solidity'
var jayson = require('jayson/promise');
const utils = require('ethereumjs-util');
const abi = require('ethereumjs-abi');
var Web3Utils = require('web3-utils');

function breakUpMessages(messages) {
  var messageData = "0x";
  var messageAmount = [];
  var messageTokenNum = [];
  var messageDestination = [];

  messages.forEach(function(message) {
    let buf = Buffer.from(message.tokenType.slice(2), "hex");
    messageData += message.data.slice(2);
    messageTokenNum.push(message.tokenNum);
    messageAmount.push(message.amount);
    messageDestination.push(message.destination);
  });
  return [messageData, messageTokenNum, messageAmount, messageDestination];
}

function breakUpMessageHashes(messages) {
  var messageData = [];
  var messageAmount = [];
  var messageTokenNum = [];
  var messageDestination = [];

  messages.forEach(function(message) {
    messageData.push(message.dataHash);
    messageTokenNum.push(message.tokenNum);
    messageAmount.push(message.amount);
    messageDestination.push(message.destination);
  });
  return [messageData, messageTokenNum, messageAmount, messageDestination];
}

function hashPrecondition(beforeHash, startTime, endTime, beforeInbox, tokenTypes, beforeBalances) {
  return Web3Utils.soliditySha3(
    {t: 'bytes32', v: beforeHash},
    {t: 'uint256', v: startTime},
    {t: 'uint256', v: endTime},
    {t: 'bytes32', v: beforeInbox},
    {t: 'bytes21[]', v: tokenTypes},
    {t: 'uint256[]', v: beforeBalances}
  );
}

function hashAssertion(afterHash, numSteps, firstMessageHash, lastMessageHash, totalMessageValues) {
  return Web3Utils.soliditySha3(
    {t: 'bytes32', v: afterHash},
    {t: 'uint32', v: numSteps},
    {t: 'bytes32', v: firstMessageHash},
    {t: 'bytes32', v: lastMessageHash},
    {t: 'uint256[]', v: totalMessageValues}
  );
}

function bisection(client, vmId, challengeVerifier, assertionNum, address, managerNum, config, blockNumber) {
  let data = {
    "assertionNum": assertionNum,
    "vmId": vmId,
    "managerNum": managerNum,
    "slices": 5
  };
  client.request('RPCInterface.GenerateBisection', [data]).then(function(response) {
    if (response.result.assertion.numSteps > 1) {
      var afterHashes = [];
      var totalSteps = 0;
      var messageBisections = [response.result.bisections[0].firstMessageHash];
      var messageTotals = [];
      var messageCoinNums = [];
      var messageCoinCount = [];
      for (var i = 0; i < response.result.bisections.length; i++) {
        let bisection = response.result.bisections[i];
        afterHashes.push(bisection.afterHash);
        totalSteps += bisection.numSteps;
        messageBisections.push(bisection.lastMessageHash);
        messageTotals = messageTotals.concat(bisection.totalMessageAmounts);
        messageCoinNums = messageCoinNums.concat(bisection.totalMessageCoinNums);
        messageCoinCount.push(bisection.totalMessageAmounts.length);
      }
      let afterHashesAndMessageBisections = afterHashes.concat(messageBisections);
      console.log("Bisecting challenge");
      return challengeVerifier.bisectAssertion(
        [vmId, response.result.precondition.beforeHash, response.result.precondition.beforeInbox],
        afterHashesAndMessageBisections,
        messageTotals,
        messageCoinNums,
        messageCoinCount,
        totalSteps,
        response.result.precondition.timeBounds,
        response.result.precondition.tokenTypes,
        response.result.precondition.beforeValues,
        blockNumber + config.gracePeriod,
        {from:address}
      ).then(function(res) {
        console.log("Bisection gas used: ", res.receipt.gasUsed);
      });
    } else {
      console.log("One step prooving challenge");
      challengeVerifier.oneStepProof(
        vmId,
        [response.result.precondition.beforeHash, response.result.precondition.beforeInbox],
        response.result.precondition.timeBounds,
        response.result.precondition.tokenTypes,
        response.result.precondition.beforeValues,
        [
          response.result.assertion.afterHash,
          response.result.assertion.firstMessageHash,
          response.result.assertion.lastMessageHash
        ],
        response.result.assertion.totalMessageAmounts,
        response.result.assertion.totalMessageCoinNums,
        response.result.proof,
        blockNumber + config.gracePeriod,
        {from: address}
      );
    }
  });
}

const ManagerState = {
    WAITING: 0,
    PENDING: 1,
    BISECTING: 3,
    CHALLENGING: 4,
}

export class RPCManager {
  constructor(vmTracker, challengeVerifier, vmId, config, managerNum, forceChallenge) {
    var elements = [];
    console.log(config.assertKeys);
    for (var i = 0; i < config.assertKeys.length; i++) {
      let addressBytes = Buffer.concat([Buffer.alloc(12), Buffer.from(config.assertKeys[i].slice(2), 'hex')]);
      elements.push(addressBytes);
    }
    let tree = new MerkleTree(elements, true);
    this.config = config;
    this.vmTracker = vmTracker;
    this.managerProof = tree.getProofOrdered(tree.elements[managerNum], managerNum + 1, true);
    this.managerNum = managerNum;
    this.vmId = vmId;
    this.address = config.assertKeys[managerNum];
    this.managerState = ManagerState.WAITING;

    var client = jayson.client.http('http://localhost:1235/rpc');
    this.client = client;
    let self = this;

    let registerData = {
      "vmId": vmId,
      "managerNum": managerNum
    };
    client.request('RPCInterface.RegisterVM', [registerData]).then(function(reponse) {
      console.log("Registered VM", vmId, managerNum);
    });

    const disputableAssertionWatcher = vmTracker.DisputableAssertion({vmId: vmId});
    disputableAssertionWatcher.watch(function(err, result) {
      let data = {
        "vmId": vmId,
        "managerNum": managerNum,
        "beforeHash": result.args.beforeHash,
        "timeBounds": result.args.timeBounds.map(item => item.toString()),
        "tokenTypes": result.args.tokenTypes,
        "beforeInbox": result.args.beforeInbox,
        "afterHash":result.args.afterHash,
        "numSteps":result.args.numSteps.toNumber(),
        "lastMessageHash":result.args.lastMessageHash,
        "totalMessageValues":result.args.totalMessageValues.map(b => b.toString(10))
      };
      if (self.managerState != ManagerState.WAITING) {
        return;
      }
      console.log("Saw DisputableAssertion event");
      client.request('RPCInterface.SawDisputableAssertion', [data]).then(function(response) {
        if (!response.result.valid || forceChallenge) {
          console.log("Challenging assertion");
          self.managerState = ManagerState.CHALLENGING;
          let assertionHash = hashAssertion(
            result.args.afterHash,
            result.args.numSteps,
            '0x00',
            result.args.lastMessageHash,
            result.args.totalMessageValues
          );

          return vmTracker.initiateChallenge(
            [
              vmId,
              result.args.beforeHash,
              result.args.beforeInbox,
              assertionHash
            ],
            config.assertKeys[result.args.asserterNum.toNumber()],
            managerNum,
            self.managerProof,
            result.blockNumber + config.gracePeriod,
            result.args.timeBounds,
            result.args.tokenTypes,
            result.args.totalMessageValues,
            {from:self.address}
          );
        } else {
          console.log("Not challenging assertion");
          self.managerState = ManagerState.PENDING;
        }
      });
    });

    const confirmedAssertionWatcher = vmTracker.ConfirmedAssertion({vmId: vmId});
    confirmedAssertionWatcher.watch(function(err, result) {
      if (self.managerState == ManagerState.PENDING) {
        console.log("Saw new finalized assertion");
        let data = {
          "vmId": vmId,
          "managerNum": managerNum
        }
        client.request('RPCInterface.SawFinalizedAssertion', [data]).then(function(response) {
          console.log("Updated to new assertion");
          self.managerState = ManagerState.WAITING;
        });
      }
    });

    const bisectionAssertionWatcher = challengeVerifier.BisectedAssertion({vmId: vmId});
    bisectionAssertionWatcher.watch(function(err, result) {
      if (self.managerState != ManagerState.CHALLENGING) {
        return;
      }
      console.log("Saw BisectedAssertion event");
      let bisectionCount = (result.args.afterHashAndMessageBisections.length - 1) / 2;
      let totalSteps = result.args.totalSteps.toNumber();
      let extraSteps = totalSteps % bisectionCount;
      let stepCount = (totalSteps - extraSteps) / bisectionCount;

      let amounts = result.args.totalMessageAmounts.map(b => b.toString(10));
      var bisections = [];
      
      var count = 0;
      for (var i = 0; i < bisectionCount; i++) {
        let numSteps = stepCount;
        if (i < extraSteps) {
          numSteps++;
        }
        let coinCount = result.args.totalMessageCoinCount[i];
        bisections.push({
          "afterHash": result.args.afterHashAndMessageBisections[i],
          "numSteps": numSteps,
          "firstMessageHash": result.args.afterHashAndMessageBisections[bisectionCount + i],
          "lastMessageHash": result.args.afterHashAndMessageBisections[bisectionCount + i + 1],
          "totalMessageAmounts":amounts.slice(count, count + coinCount),
          "totalMessageCoinNums":result.args.totalMessageCoinNums.slice(count, count + coinCount)
        });
        count += coinCount
      }
      let data = {
        "vmId": vmId,
        "managerNum": managerNum,
        "bisections":bisections
      }
      client.request('RPCInterface.SawBisection', [data]).then(function(response) {
        if (response.error != null) {
          throw response.error;
        }
        var hashes = [];
        for (var i = 0; i < response.result.bisections.length; i++) {
          let bisection = response.result.bisections[i];
          let preconditionHash = hashPrecondition(
            bisection.beforeHash,
            bisection.startTime,
            bisection.endTime,
            bisection.beforeInbox,
            bisection.tokenTypes,
            bisection.beforeValues
          );
          let assertionHash = hashAssertion(
            bisection.afterHash,
            bisection.numSteps,
            bisection.firstMessageHash,
            bisection.lastMessageHash,
            bisection.totalMessageValues
          );
          hashes.push(utils.toBuffer(abi.soliditySHA3(["bytes32", "bytes32"], [preconditionHash, assertionHash])));
        }
        let tree = new MerkleTree(hashes, true);
        let proof = tree.getProofOrdered(tree.elements[response.result.invalidBisection], response.result.invalidBisection + 1, true);

        console.log("Continuing challenge");
        return challengeVerifier.continueChallenge(
          vmId,
          response.result.invalidBisection,
          proof,
          result.blockNumber + config.gracePeriod,
          utils.bufferToHex(tree.getRoot()),
          utils.bufferToHex(tree.elements[response.result.invalidBisection]),
          {from:self.address}
        ).then(function(res) {
          console.log("Challenge gas used: ", res.receipt.gasUsed);
        });
      });
    });

    const initiatedChallengeWatcher = challengeVerifier.InitiatedChallenge({vmId: vmId});
    initiatedChallengeWatcher.watch(function(err, result) {
      if (self.managerState != ManagerState.PENDING) {
        return;
      }
      self.managerState = ManagerState.BISECTING;
      console.log("Saw InitiatedChallenge event");
      return bisection(client, vmId, challengeVerifier, 0, self.address, managerNum, config, result.blockNumber);
    });

    const challengeWatcher = challengeVerifier.ContinuedChallenge({vmId: vmId});
    challengeWatcher.watch(function(err, result) {
      if (self.managerState != ManagerState.BISECTING) {
        return;
      }
      console.log("Saw ContinuedChallenge event");
      return bisection(client, vmId, challengeVerifier, result.args.assertionIndex.toNumber(), self.address, managerNum, config, result.blockNumber);
    });
  }

  async disputableAssert(numSteps) {
    let self = this;
    self.managerState = ManagerState.PENDING;
    let data = {
      "vmId": self.vmId,
      "managerNum": self.managerNum,
      "numSteps": numSteps,
      "timeBounds": [web3.eth.blockNumber, web3.eth.blockNumber + 2].map(val => val.toString())
    };
    let response = await this.client.request('RPCInterface.GenerateDisputableAssertion', [data]);
    var [ 
      messageData, 
      messageTokenNum, 
      messageAmount, 
      messageDestinations
    ] = breakUpMessages(response.result.messages);
    let disputableResult = await self.vmTracker.disputableAssert(
      [
        self.vmId, 
        response.result.beforeHash, 
        response.result.beforeInbox, 
        response.result.afterHash
      ],
      [self.managerNum, response.result.numSteps, data.timeBounds[0], data.timeBounds[1]],
      self.managerProof,
      response.result.tokenTypes,
      messageData,
      messageTokenNum,
      messageAmount,
      messageDestinations,
      {from:self.address}
    );
    let blockNumber = disputableResult.receipt.blockNumber;
    let preconditionHash = hashPrecondition(
      response.result.beforeHash,
      data.timeBounds[0],
      data.timeBounds[1],
      response.result.beforeInbox,
      response.result.tokenTypes,
      response.result.beforeValues
    );
    var [ 
      messageDataHashes, 
      messageTokenNum, 
      messageAmount, 
      messageDestinations
    ] = breakUpMessageHashes(response.result.messages);
    self.waitForBlock(blockNumber + self.config.gracePeriod).then(function() {
      console.log("Ready to confirm assertion");
      if (self.managerState == ManagerState.PENDING) {
        console.log("Confirmed assertion");
        return self.vmTracker.confirmAsserted(
          self.address,
          blockNumber + self.config.gracePeriod,
          [
            self.vmId,
            preHash,
            response.result.afterHash,
            response.result.lastMessageHash
          ],
          response.result.numSteps,
          response.result.tokenTypes,
          messageDataHashes,
          messageTokenNum,
          messageAmount,
          messageDestinations,
          {from:self.address}
        );
      }
    });
  }

  getVMState(callback) {
    let data = {
      "vmId": this.vmId,
      "managerNum": this.managerNum
    };
    this.client.request('RPCInterface.GetVMState', [data]).then(function(response) {
      callback(response);
    })
  }

  waitForBlock(height) {
    return new Promise (function (resolve, reject) {
      var filter = web3.eth.filter('latest');
      filter.watch(function(error, result){
        if (error) {
          filter.stopWatching();
          reject(error);
        } else {
          let block = web3.eth.getBlock(result, true);
          if (block.number >= height) {
            filter.stopWatching();
            resolve(result);
          }
        }
      })
    });
  }
}
