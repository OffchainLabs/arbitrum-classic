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

import MerkleTree, { checkProof, merkleRoot, checkProofSolidityFactory } from 'merkle-tree-solidity'

const utils = require('ethereumjs-util');
const abi = require('ethereumjs-abi');
import {BigNumber} from 'bignumber.js';

class TokenNumTracker {
  constructor(messages) {
    let tokens = {};
    let tokenTypes = [];
    let fullTokenTypes = [];
    var tokenCount = 0;
    messages.forEach(function(message) {
      if (message.isToken()) {
        if (!(message.tokenType in tokens)) {
          tokens[message.tokenType] = tokenCount;
          tokenCount++;
          tokenTypes.push(message.tokenType);
          fullTokenTypes.push([message.tokenType, 0]);
        }
      } else {
        if (!(message.tokenType in tokens)) {
          tokens[message.tokenType] = {};
        }
        tokens[message.tokenType][message.amount.toString(16)] = tokenCount;
        tokenTypes.push(message.tokenType);
        fullTokenTypes.push([message.tokenType, message.amount]);
      }
    });
    this.tokens = tokens;
    this.tokenTypes = tokenTypes;
    this.fullTokenTypes = fullTokenTypes;
  }

  getTokenNum(tokenType, amount) {
    let buf = Buffer.from(tokenType.slice(2), "hex");
    if (buf.readUInt8(20) == 0) {
      return this.tokens[tokenType];
    } else {
      return this.tokens[tokenType][amount.toString(16)];
    }
  }
}

function breakUpMessages(messages) {
  var messageData = "0x";
  var messageAmount = [];
  var messageTokenNum = [];
  var messageDestination = [];
  let tokenNumTracker = new TokenNumTracker(message);

  messages.forEach(function(message) {
    let buf = Buffer.from(message.tokenType.slice(2), "hex");
    messageData += message.data.slice(2);
    messageTokenNum.push(tokenNumTracker.getTokenNum(message.tokenType, message.amount))
    messageAmount.push(message.amount);
    messageDestination.push(message.destination);
  });
  return [tokenNumTracker, messageData, messageTokenNum, messageAmount, messageDestination];
}

class ArbMessage {
  constructor(data, tokenType, amount, destination) {
    this.data = data;
    this.tokenType = tokenType;
    this.amount = amount;
    this.destination = destination;
  }

  isToken() {
    let buf = Buffer.from(this.tokenType.slice(2), "hex");
    return buf.readUInt8(20) == 0;
  }

  hash(arbMachineLib, arbValueLib) {
    let self = this;
    return arbValueLib.deserialize_value_hash(this.data).then(function(valueHash) {
      return arbMachineLib.generateMessageStubHash(
        valueHash,
        self.tokenType,
        self.amount,
        self.destination
      );
    })
  }

  print() {
    console.log(`message(${this.data}, ${this.tokenType}, ${this.amount}, ${this.destination})`);
  }

  isEqualTo(other) {
    return this.data == other.data &&
    this.tokenType == other.tokenType &&
    this.amount.eq(other.amount) &&
    this.destination == other.destination
  }
}

function getTotalVals(tracker, messages) {
  var totalVals = [];
  for (var i = 0; i < tracker.tokenTypes.length; i++) {
    totalVals.push(web3.toBigNumber(0));
  }

  for (var i = 0; i < messages.length; i++) {
    let index = tracker.getTokenNum(messages[i].tokenType, messages[i].amount);
    totalVals[index] = totalVals[index].add(messages[i].amount);
  }
  return totalVals;
}

class FullPreconditionAssertion {

  static freshAssertion(beforeHash, timeBounds, beforeInbox, afterHash, numSteps, messages) {
    var newMessages = [];
    for (var i = 0; i < messages.length; i++) {
      newMessages.push(new ArbMessage(
        messages[i].data,
        messages[i].tokenType,
        messages[i].amount,
        messages[i].destination
      ));
    }
    let tracker = new TokenNumTracker(newMessages);
    let beforeValues = getTotalVals(tracker, newMessages);
    return new FullPreconditionAssertion(
      beforeHash,
      timeBounds,
      beforeInbox,
      afterHash,
      numSteps,
      messages,
      tracker,
      beforeValues,
      '0x00'
    );
  }

  constructor(
    beforeHash,
    timeBounds,
    beforeInbox,
    afterHash,
    numSteps,
    messages,
    tracker,
    beforeValues,
    firstMessageHash
  ) {
    this.beforeHash = beforeHash;
    this.startTime = timeBounds[0];
    this.endTime = timeBounds[1];
    this.beforeInbox = beforeInbox;
    this.afterHash = afterHash;
    this.numSteps = numSteps;
    this.tracker = tracker;
    this.beforeValues = beforeValues;
    this.firstMessageHash = firstMessageHash;

    this.messages = [];
    for (var i = 0; i < messages.length; i++) {
      this.messages.push(new ArbMessage(
        messages[i].data,
        messages[i].tokenType,
        messages[i].amount,
        messages[i].destination
      ));
    }
  }

  async lastMessageHash(arbMachineLib, arbValueLib) {
    let endHash = this.firstMessageHash;
    for (var i = 0; i < this.messages.length; i++) {
      let messageHash = await this.messages[i].hash(arbMachineLib, arbValueLib);
      endHash = '0x' + abi.soliditySHA3(['bytes32', 'bytes32'], [endHash, messageHash]).toString('hex');
    }
    return endHash;
  }

  preconditionHash(arbMachineLib) {
    let tokenTracker = new TokenNumTracker(this.messages);;
    return arbMachineLib.generatePreconditionHash(
      this.beforeHash,
      this.startTime,
      this.endTime,
      this.beforeInbox,
      tokenTracker.tokenTypes,
      this.beforeValues
    );
  }

  breakUpMessages(tokenNumTracker) {
    var messageData = "0x";
    var messageAmount = [];
    var messageTokenNum = [];
    var messageDestination = [];
    if (tokenNumTracker == null) {
      tokenNumTracker = new TokenNumTracker(this.messages);
    }

    this.messages.forEach(function(message) {
      let buf = Buffer.from(message.tokenType.slice(2), "hex");
      messageData += message.data.slice(2);
      messageTokenNum.push(tokenNumTracker.getTokenNum(message.tokenType, message.amount))
      messageAmount.push(message.amount);
      messageDestination.push(message.destination);
    });
    return [tokenNumTracker, messageData, messageTokenNum, messageAmount, messageDestination];
  }

  async breakUpMessageHashes(arbValue) {
    var messageData = [];
    var messageAmount = [];
    var messageTokens = [];
    var messageDestination = [];
    for (var i = 0; i < this.messages.length; i++) {
      messageData.push(await deserialize_value_hash(this.messages[i].data));
      messageTokens.push(this.messages[i].tokenType);
      messageAmount.push(this.messages[i].amount);
      messageDestination.push(this.messages[i].destination);
    }
    return [messageData, messageTokens, messageAmount, messageDestination];
  }

  getTotalVals() {
    return getTotalVals(this.tracker, this.messages);
  }

  generateBisectionStub(tokenNumTracker) {
    let tracker = new TokenNumTracker(this.messages);

    var totalVals = this.getTotalVals();
    var totalCoinNums = [];
    for (var i = 0; i < tracker.fullTokenTypes.length; i++) {
      totalCoinNums.push(tokenNumTracker.getTokenNum(tracker.fullTokenTypes[i][0], tracker.fullTokenTypes[i][1]));
    }
    return [totalVals, totalCoinNums];
  }

  async bisect(firstHalfAssertion, arbMachineLib, arbValueLib) {
    if (
      firstHalfAssertion.beforeHash != this.beforeHash ||
      firstHalfAssertion.beforeInbox != this.beforeInbox ||
      firstHalfAssertion.startTime != this.startTime ||
      firstHalfAssertion.endTime != this.endTime
    ) {
      throw("bisect: preconditions don't match");
    }

    if (Math.floor(this.numSteps / 2) != firstHalfAssertion.numSteps) {
      throw("bisectAssertion: firstHalfAssertion has wrong number of steps");
    }

    if (firstHalfAssertion.messages.length > this.messages.length) {
      throw("bisect: firstHalfAssertion has more messages than original");
    }

    for (var i = 0; i < firstHalfAssertion.messages.length; i++) {
      if (!firstHalfAssertion.messages[i].isEqualTo(this.messages[i])) {
        firstHalfAssertion.messages[i].print();
        this.messages[i].print();
        throw("bisect: firstHalfAssertion has different message than original");
      }
    }

    if (firstHalfAssertion.firstMessageHash != this.firstMessageHash) {
      throw("bisectAssertion: firstHalfAssertion message start doesn't match full assertion");
    }

    let secondHalfStart = await firstHalfAssertion.lastMessageHash(arbMachineLib, arbValueLib);

    let newBeforeValues = [];
    for (var i = 0; i < this.beforeValues.length; i++) {
      newBeforeValues.push(this.beforeValues[i]);
    }

    let firstHalfTotals = getTotalVals(this.tracker, firstHalfAssertion.messages);
    for (var i = 0; i < firstHalfTotals.length; i++) {
      newBeforeValues[i] -= firstHalfTotals[i]
    }

    let secondHalfAssertion = new FullPreconditionAssertion(
      firstHalfAssertion.afterHash,
      [this.startTime, this.endTime],
      this.beforeInbox,
      this.afterHash,
      this.numSteps - firstHalfAssertion.numSteps,
      this.messages.slice(firstHalfAssertion.messages.length),
      this.tracker,
      newBeforeValues,
      secondHalfStart
    );
    return [firstHalfAssertion, secondHalfAssertion];
  }
}

class StubPreconditionAssertion {
  constructor(
    beforeHash,
    timeBounds,
    beforeInbox,
    tokenTypes,
    beforeValues,
    afterHash,
    numSteps,
    firstMessageHash,
    lastMessageHash,
    totalMessageValues
  ) {
    this.beforeHash = beforeHash;
    this.startTime = timeBounds[0];
    this.endTime = timeBounds[1];
    this.beforeInbox = beforeInbox;
    this.beforeValues = beforeValues;
    this.afterHash = afterHash;
    this.numSteps = numSteps;
    this.firstMessageHash = firstMessageHash;
    this.lastMessageHash = lastMessageHash;
    this.totalMessageValues = totalMessageValues;
    this.tokenTypes = tokenTypes;
  }

  assertionHash(arbMachineLib) {
    return arbMachineLib.generateAssertionHash(
      this.afterHash,
      this.numSteps,
      this.firstMessageHash,
      this.lastMessageHash,
      this.totalMessageValues
    );
  }

  preconditionHash(arbMachineLib) {
    return arbMachineLib.generatePreconditionHash(
      this.beforeHash,
      [this.startTime, this.endTime],
      this.beforeInbox,
      this.tokenTypes,
      this.beforeValues
    );
  }

  preconditionAssertionHash(arbMachineLib) {
    return Promise.all([this.preconditionHash(arbMachineLib), this.assertionHash(arbMachineLib)]).then(function(values) {
      return abi.soliditySHA3(["bytes32", "bytes32"], values);
    })
  }

  generateChildren(
    afterHashes,
    numSteps,
    messageBisections,
    totalMessageAmounts,
    totalMessageCoinNums,
    totalMessageCoinCount
  ) {
    var children = [];
    var beforeHash = this.beforeHash;
    var count = 0;
    var newBeforeValues = [];
    for (var i = 0; i < this.beforeValues.length; i++) {
      newBeforeValues.push(web3.toBigNumber(this.beforeValues[i]));
    }

    for (var i = 0; i < afterHashes.length; i++) {
      var values = [];
      for (var j = 0; j < this.tokenTypes.length; j++) {
        values.push(web3.toBigNumber(0));
      }
      for (var j = 0; j < totalMessageCoinCount[i]; j++) {
        let coinNum = totalMessageCoinNums[count + i];
        values[coinNum] = values[coinNum].add(totalMessageAmounts[count]);
      }

      children.push(new StubPreconditionAssertion(
        beforeHash,
        [this.startTime, this.endTime],
        this.beforeInbox,
        this.tokenTypes,
        newBeforeValues,
        afterHashes[i],
        numSteps[i],
        messageBisections[i],
        messageBisections[i + 1],
        values
      ));

      for (var j = 0; j < totalMessageCoinCount[i]; j++) {
        let coinNum = totalMessageCoinNums[count + i];
        newBeforeValues[coinNum] = newBeforeValues[coinNum].sub(totalMessageAmounts[count]);
      }

      beforeHash = afterHashes[i];
      count += totalMessageCoinCount[i];

    }
    return children;
  }
}

export class ArbVM {
  constructor(vmTracker, arbMachineLib, arbValueLib, vmId, config, address) {
    this.vmTracker = vmTracker;
    this.arbMachineLib = arbMachineLib;
    this.arbValueLib = arbValueLib;
    this.vmId = vmId;
    this.config = config;
    this.managerNum = -1;
    this.address = address;
    this.pendingAssertion = null;
    this.bisectionPreconditions = null;
    this.challengeVerifier = config["challengeVerifier"];
    for (const [i, key] of config.assertKeys.entries()) {
      if (key == address) {
        this.managerNum = i;
        break;
      }
    }
    var self = this;
    const disputableAssertionWatcher = vmTracker.DisputableAssertion({vmId: vmId});
    disputableAssertionWatcher.watch(function(err, result) {
      self.pendingPreconditionAssertion = new StubPreconditionAssertion(
        result.args.beforeHash,
        [result.args.timeBounds[0].toNumber(), result.args.timeBounds[1].toNumber()],
        result.args.beforeInbox,
        result.args.tokenTypes,
        result.args.totalMessageValues,
        result.args.afterHash,
        result.args.numSteps.toNumber(),
        "0x00",
        result.args.lastMessageHash,
        result.args.totalMessageValues
      );

      self.pendingAsserter = result.args.asserterNum.toNumber();
      self.pendingAssertionDeadline = result.args.deadline.toNumber();
    });

    const bisectionAssertionWatcher = self.challengeVerifier.BisectedAssertion({vmId: vmId});
    bisectionAssertionWatcher.watch(function(err, result) {
      self.bisectionPreconditionAssertions = self.pendingPreconditionAssertion.generateChildren(
        result.args.afterHash,
        result.args.numSteps,
        result.args.messageBisections,
        result.args.totalMessageAmounts,
        result.args.totalMessageCoinNums,
        result.args.totalMessageCoinCount
      );
      let bisectionHashPromises = [];
      for (var i = 0; i < self.bisectionPreconditionAssertions.length; i++) {
        bisectionHashPromises.push(self.bisectionPreconditionAssertions[i].preconditionAssertionHash(self.arbMachineLib));
      }
      Promise.all(bisectionHashPromises).then(function(bisectionHashes) {
        self.bisectionHashes = bisectionHashes;
      })
    });

    const challengeWatcher = self.challengeVerifier.ContinuedChallenge({vmId: self.vmId});
    challengeWatcher.watch(function(err, result) {
      if (self.myBisections != null) {
        self.myAssertion = self.myBisections[result.args.assertionIndex];
      }
    });

    const timeoutWatcher = self.challengeVerifier.TimedOutChallenge({vmId: self.vmId});
    timeoutWatcher.watch(function(err, result) {
      self.pendingAssertion = null;
      self.pendingPrecondition = null;
    })
  }

  static async vmWithId(challengeManager, vmTracker, arbMachineLib, arbValueLib, vmId, address, assertKeys) {
    let result = await vmTracker.getConfig(vmId);

    return new ArbVM(vmTracker, arbMachineLib, arbValueLib, vmId, {
      "escrowRequired": result[0].toNumber(),
      "gracePeriod": result[1].toNumber(),
      "maxExecutionSteps":result[2].toNumber(),
      "assertKeys": assertKeys,
      "challengeVerifier": challengeManager
    }, address);
  }

  getVmInfo() {
    return this.vmTracker.getVm.call(this.vmId).then(function(result) {
      var tokens = {};
      var nfts = {};
      for (var i = 0; i < result[4].length; i++) {
        tokens[result[4][i]] = result[5][i];
      }
      for (var i = 0; i < result[6].length; i++) {
        if (!(result[6][i] in nfts)) {
          nfts[result[6][i]] = {};
        }
        nfts[result[6][i]][result[7][i]] = true;
      }
      return {
        "state": result[0].toNumber(),
        "stateHash": result[1],
        "inboxHash": result[2],
        "pendingMessages": result[3],
        "tokens": tokens,
        "nfts": nfts
      };
    });
  }

  getChallengeInfo() {
    return this.vmTracker.getChallenge.call(this.vmId).then(function(result) {
      return {
        "state": result[0].toNumber(),
        "asserter": result[1],
        "challenger": result[2],
        "asserterEscrow": result[3].toNumber(),
        "challengerEscrow": result[4].toNumber(),
        "challengePeriod": result[5].toNumber(),
        "deadline": result[6].toNumber(),
        "assertionPreconditionRoot": result[7]
      };
    });
  }

  generatePrecondition(timeRange) {
    return this.getVmInfo().then(function(vmState) {
      var blockNum = web3.eth.blockNumber;
      return {
        "inbox":vmState.inboxHash,
        "stateHash":vmState.stateHash,
        "beforeTime":blockNum + 2,
        "afterTime":blockNum
      }
    });
  }

  async unanimousAssertHash(precondition, assertion) {
    var messageHashes = []
    for (var i = 0; i < assertion.messages.length; i++) {
      messageHashes.push(await hashMessage(assertion.messages[i], this.arbMachineLib));
    }
    

    var [messageData, messageAmount, messageDestination] = breakUpMessages(assertion.messages);
    return this.vmTracker.unanimousAssertHash(
      this.vmId,
      [precondition.stateHash, precondition.inbox],
      [precondition.afterTime, precondition.beforeTime],
      precondition.balance,
      assertion.afterHash,
      assertion.numSteps,
      messageHashes
    );
  }

  unanimousAssert(precondition, assertion, signatures) {
    var self = this;
    var [messageData, messageAmount, messageDestination] = breakUpMessages(assertion.messages);
    var sigVs = [];
    var sigRs = [];
    var sigSs = [];

    signatures.forEach(function(signature) {
      sigVs.push(signature.v);
      sigRs.push('0x' + signature.r.toString('hex'));
      sigSs.push('0x' + signature.s.toString('hex'));
    });

    return this.vmTracker.unanimousAssert(
      [this.vmId, precondition.stateHash, precondition.inbox, assertion.afterHash],
      [precondition.afterTime, precondition.beforeTime],
      [precondition.balance, assertion.numSteps],
      messageData,
      messageAmount,
      messageDestination,
      sigVs,
      sigRs,
      sigSs,
      {from: this.address}
    );
  }

  disputableAssert(precondition, assertion) {
    return this._disputableAssert(precondition, assertion, this.config.assertEscrowRequired);
  }

  async _disputableAssert(vmInfo, assertion, escrow) {

    this.myAssertion = FullPreconditionAssertion.freshAssertion(
      vmInfo.stateHash,
      [web3.eth.blockNumber, web3.eth.blockNumber + 2],
      vmInfo.inboxHash,
      assertion.afterHash,
      assertion.numSteps,
      assertion.messages
    );

    var [
      tokenTracker, 
      messageData, 
      messageTokenNum, 
      messageAmount, 
      messageDestinations
    ] = this.myAssertion.breakUpMessages();

    let tree = this.managerMerkleTree();
    let proof = tree.getProofOrdered(tree.elements[this.managerNum], this.managerNum + 1, true);

    let result = await this.vmTracker.disputableAssert(
      [this.vmId, this.myAssertion.beforeHash, this.myAssertion.beforeInbox, this.myAssertion.afterHash],
      [this.managerNum, this.myAssertion.numSteps, this.myAssertion.startTime, this.myAssertion.endTime],
      proof,
      tokenTracker.tokenTypes,
      messageData,
      messageTokenNum,
      messageAmount,
      messageDestinations,
      {from:this.address}
      // {value: escrow, from:this.address}
    );
    this.myAssertionDeadline = result.logs[result.logs.length - 1].args.deadline.toNumber();
  }

  async confirmAsserted() {
    var self = this;
    if (this.myAssertion == null) {
      throw("Must have assertion pending");
    }
    var [
      tokenTracker, 
      messageData, 
      messageTokenNum, 
      messageAmount, 
      messageDestinations
    ] = this.myAssertion.breakUpMessages();
    var [messageData, messageTokens, messageAmount, messageDestination] = await this.myAssertion.breakUpMessageHashes(this.arbValueLib);
    let preconditionHash = await this.myAssertion.preconditionHash(this.arbMachineLib);
    let lastMessageHash = await this.myAssertion.lastMessageHash(this.arbMachineLib, this.arbValueLib);
    return this.vmTracker.confirmAsserted(
      this.config.assertKeys[this.pendingAsserter],
      this.myAssertionDeadline,
      [
        this.vmId,
        preconditionHash,
        this.myAssertion.afterHash,
        this.myAssertion.firstMessageHash,
        lastMessageHash
      ],
      this.myAssertion.numSteps,
      this.myAssertion.getTotalVals(),
      messageData,
      messageTokens,
      messageAmount,
      messageDestination,
      {from:this.address}
    ).then(function(result) {
      self.myAssertion = null;
    });
  }

  async initiateChallenge() {
    let tree = this.managerMerkleTree();
    let proof = tree.getProofOrdered(tree.elements[this.managerNum], this.managerNum + 1, true);
    let assertionHash = await this.pendingPreconditionAssertion.assertionHash(this.arbMachineLib);
    return this.vmTracker.initiateChallenge(
      this.vmId,
      this.config.assertKeys[this.pendingAsserter],
      this.managerNum,
      proof,
      this.pendingAssertionDeadline,
      this.pendingPreconditionAssertion.beforeHash,
      [this.pendingPreconditionAssertion.startTime, this.pendingPreconditionAssertion.endTime],
      this.pendingPreconditionAssertion.beforeInbox,
      this.pendingPreconditionAssertion.tokenTypes,
      this.pendingPreconditionAssertion.totalMessageValues,
      assertionHash,
      {from:this.address}
    );
  }

  async bisectAssertion(firstHalfAssertion) {
    if (this.myAssertion == null) {
      throw ("There must be a pending assertion to bisect");
    }
    let fullFirstHalfAssertion = new FullPreconditionAssertion(
      this.myAssertion.beforeHash,
      [this.myAssertion.startTime, this.myAssertion.endTime],
      this.myAssertion.beforeInbox,
      firstHalfAssertion.afterHash,
      firstHalfAssertion.numSteps,
      firstHalfAssertion.messages,
      this.myAssertion.tracker,
      this.myAssertion.beforeValues,
      this.myAssertion.firstMessageHash
    );
    let assertions = await this.myAssertion.bisect(fullFirstHalfAssertion, this.arbMachineLib, this.arbValueLib);
    return this._bisectAssertion(this.myAssertion, assertions);
  }
  
  async _bisectAssertion(fullAssertion, assertions) {
    this.myBisections = assertions;
    var self = this;
    let afterHashes = [];
    let numSteps = [];
    let messageBisections = [];
    let messageCoinCount = [];
    let messageTotals = [];
    let messageCoinNums = [];

    let prevHash = '0x00';
    messageBisections = [prevHash];
    for (var i = 0; i < assertions.length; i++) {
      messageBisections.push(await assertions[i].lastMessageHash(this.arbMachineLib, this.arbValueLib));
      prevHash = assertions[i].lastMessage;
      afterHashes.push(assertions[i].afterHash);
      numSteps.push(assertions[i].numSteps);
      var [totalVals, totalCoinNums] = assertions[i].generateBisectionStub(fullAssertion.tracker);
      messageCoinCount.push(totalVals.length);
      for (var j = 0; j < totalVals.length; j++) {
        messageTotals.push(totalVals[j]);
        messageCoinNums.push(totalCoinNums[j])
      }
    }

    return this.challengeVerifier.bisectAssertion(
      [this.vmId, fullAssertion.beforeHash, fullAssertion.beforeInbox],
      afterHashes,
      numSteps,
      messageBisections,
      messageTotals,
      messageCoinNums,
      messageCoinCount,
      [fullAssertion.startTime, fullAssertion.endTime],
      fullAssertion.tracker.tokenTypes,
      fullAssertion.beforeValues,
      {from:this.address}
    );
  }

  async bisectionMerkleTree() {
    var elements = [];
    for (var i = 0; i < this.bisectionHashes.length; i++) {
      elements.push(utils.toBuffer(this.bisectionHashes[i]));
    }
    return new MerkleTree(elements, true);
  }

  managerMerkleTree() {
    var elements = [];
    for (var i = 0; i < this.config.assertKeys.length; i++) {
      elements.push(abi.soliditySHA3(["address"], [this.config.assertKeys[i]]));
    }
    return new MerkleTree(elements, true);
  }

  async continueChallenge(assertionIndex) {
    if (this.bisectionPreconditionAssertions == null) {
      throw("Cannot continue challenge if bisection has not been seen");
    }
    let tree = await this.bisectionMerkleTree();
    let proof = tree.getProofOrdered(tree.elements[assertionIndex], assertionIndex + 1, true);
    this.pendingPreconditionAssertion = this.bisectionPreconditionAssertions[assertionIndex];
    return this.challengeVerifier.continueChallenge(
      this.vmId,
      assertionIndex,
      proof,
      utils.bufferToHex(tree.elements[assertionIndex]),
      {from:this.address}
    );
  }

  async oneStepProof(proof) {
    return this._oneStepProof(this.myAssertion, proof);
  }

  async _oneStepProof(assertion, proof) {
      
      var tokenTypes = [];
      var amounts = [];
      if (assertion.messages.length == 1) {
        tokenTypes.push(assertion.messages[0].tokenType);
        amounts.push(assertion.messages[0].amount);
      }

      let lastMessageHash = await assertion.lastMessageHash(this.arbMachineLib, this.arbValueLib);

      let proofPromise = this.challengeVerifier.oneStepProof(
        this.vmId,
        [assertion.beforeHash, assertion.beforeInbox],
        [assertion.startTime, assertion.endTime],
        tokenTypes,
        assertion.beforeValues,
        [assertion.afterHash, assertion.firstMessageHash, lastMessageHash],
        amounts,
        proof,
        {from: this.address}
      );
      this.pendingAssertion = null;
      this.pendingPrecondition = null;
      return proofPromise;
  }
}
