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
"use strict";

const ethers = require("ethers");
const jaysonBrowserClient = require("jayson/lib/client/browser");
const fetch = require("node-fetch");
const promisePoller = require("promise-poller").default;
const vmTrackerJson = require("./VMTracker.json");
const ArbValue = require("./arb-value");

// EthBridge event names
const EB_EVENT_VMC = "VMCreated";
const EB_EVENT_CUA = "ConfirmedUnanimousAssertion";
const EB_EVENT_FUA = "FinalUnanimousAssertion";
const EB_EVENT_CDA = "ConfirmedDisputableAssertion";

function _arbClient(managerAddress) {
  var callServer = function(request, callback) {
    var options = {
      method: "POST",
      body: request, // request is a string
      headers: {
        "Content-Type": "application/json"
      }
    };

    fetch(managerAddress, options)
      .then(function(res) {
        return res.text();
      })
      .then(function(text) {
        callback(null, text);
      })
      .catch(function(err) {
        callback(err);
      });
  };

  return jaysonBrowserClient(callServer, {});
}

const EVM_REVERT_CODE = 0;
const EVM_INVALID_CODE = 1;
const EVM_RETURN_CODE = 2;
const EVM_STOP_CODE = 3;
const EVM_BAD_SEQUENCE_CODE = 4;

function logValToLog(value) {
  return {
    contractId: value.get(0).bignum,
    data: ArbValue.sizedByteRangeToHex(value.get(1)),
    topics: value.contents.slice(2).map(val => val.bignum)
  };
}

function stackValueToList(value) {
  let values = [];
  while (value.contents.length != 0) {
    values.push(value.get(1));
    value = value.get(0);
  }
  return values;
}

function processMessage(value) {
  try {
    let wrappedData = value.get(0);
    let calldata = wrappedData.get(0);
    return {
      data: ArbValue.sizedByteRangeToHex(calldata.get(0)),
      contractID: ethers.utils.hexDataSlice(
        calldata.get(1).bignum.toHexString(),
        12
      ),
      sequenceNum: calldata.get(2).bignum.toHexString(),
      timestamp: wrappedData.get(1).bignum.toHexString(),
      blockHeight: wrappedData.get(2).bignum.toHexString(),
      txHash: wrappedData.get(3).bignum.toHexString(),
      tokenType: value.get(3).bignum.toHexString(),
      value: value.get(2).bignum.toHexString(),
      caller: ethers.utils.hexDataSlice(value.get(1).bignum.toHexString(), 12)
    };
  } catch (err) {
    console.log("processMessage got error", err);
    return 0;
  }
}

function processLog(value) {
  let origMessage = processMessage(value.get(0));
  let logs = value.get(1);
  let returnVal = value.get(2);
  let returnCode = value.get(3);

  switch (returnCode.bignum.toNumber()) {
    case EVM_RETURN_CODE:
      return {
        orig: origMessage,
        data: ArbValue.sizedByteRangeToHex(returnVal),
        logs: stackValueToList(logs).map(logValToLog),
        returnType: EVM_RETURN_CODE
      };
      break;
    case EVM_REVERT_CODE:
      return {
        orig: origMessage,
        data: ArbValue.sizedByteRangeToHex(returnVal),
        returnType: EVM_REVERT_CODE
      };
      break;
    case EVM_STOP_CODE:
      return {
        orig: origMessage,
        logs: stackValueToList(logs).map(logValToLog),
        returnType: EVM_STOP_CODE
      };
      break;
    case EVM_BAD_SEQUENCE_CODE:
      return {
        orig: origMessage,
        returnType: EVM_BAD_SEQUENCE_CODE
      };
      break;
    case EVM_INVALID_CODE:
      return {
        orig: origMessage,
        returnType: EVM_INVALID_CODE
      };
      break;
    default:
      throw "processLogs Invalid EVM return code";
  }
}

class ArbClient {
  constructor(managerUrl) {
    this.client = _arbClient(managerUrl);
  }

  async getMessageResult(txHash) {
    let self = this;
    let result = await new Promise((resolve, reject) => {
      self.client.request(
        "Validator.GetMessageResult",
        [
          {
            txHash: txHash
          }
        ],
        function(err, error, result) {
          if (err) {
            reject(err);
          } else if (error) {
            reject(error);
          } else {
            resolve(result);
          }
        }
      );
    });
    if (result["found"]) {
      let vmId = await self.getVmID();
      let val = ArbValue.unmarshal(result["rawVal"]);
      let evmVal = processLog(val);

      let data = {
        vmId: vmId,
        val: val,
        logPreHash: result["logPreHash"],
        logPostHash: result["logPostHash"],
        logValHashes: result["logValHashes"],
        validatorSigs: result["validatorSigs"],
        partialHash: result["partialHash"],
        onChainTxHash: result["onChainTxHash"]
      };

      return {
        data: data,
        evmVal: evmVal
      };
    } else {
      return null;
    }
  }

  sendMessage(value, sig, pubkey) {
    let self = this;
    return new Promise(function(resolve, reject) {
      self.client.request(
        "Validator.SendMessage",
        [
          {
            data: ArbValue.marshal(value),
            signature: sig,
            pubkey: pubkey
          }
        ],
        function(err, error, result) {
          if (err) {
            reject(err);
          } else if (error) {
            reject(error);
          } else {
            resolve(result["hash"]);
          }
        }
      );
    });
  }

  call(value, sender) {
    let self = this;
    return new Promise(function(resolve, reject) {
      self.client.request(
        "Validator.CallMessage",
        [
          {
            data: ArbValue.marshal(value),
            sender: sender
          }
        ],
        function(err, error, result) {
          if (err) {
            reject(err);
          } else if (error) {
            reject(error);
          } else {
            if (result["Success"]) {
              resolve(result["ReturnVal"]);
            } else {
              reject(new Error("Call was reverted"));
            }
          }
        }
      );
    });
  }

  findLogs(fromBlock, toBlock, address, topics) {
    let self = this;
    return new Promise(function(resolve, reject) {
      return self.client.request(
        "Validator.FindLogs",
        [
          {
            fromHeight: fromBlock,
            toHeight: toBlock,
            address: address,
            topics: topics
          }
        ],
        function(err, error, result) {
          if (err) {
            reject(err);
          } else if (error) {
            reject(error);
          } else {
            resolve(result["logs"]);
          }
        }
      );
    });
  }

  getVmID() {
    let self = this;
    return new Promise(function(resolve, reject) {
      self.client.request("Validator.GetVMInfo", [], function(
        err,
        error,
        result
      ) {
        if (err) {
          reject(err);
        } else if (error) {
          reject(error);
        } else {
          resolve(result["vmID"]);
        }
      });
    });
  }

  getAssertionCount() {
    let self = this;
    return new Promise(function(resolve, reject) {
      self.client.request("Validator.GetAssertionCount", [], function(
        err,
        error,
        result
      ) {
        if (err) {
          reject(err);
        } else if (error) {
          reject(error);
        } else {
          resolve(result["assertionCount"]);
        }
      });
    });
  }

  getVMCreatedTxHash() {
    let self = this;
    return new Promise(function(resolve, reject) {
      self.client.request("Validator.GetVMCreatedTxHash", [], function(
        err,
        error,
        result
      ) {
        if (err) {
          reject(err);
        } else if (error) {
          reject(error);
        } else {
          resolve(result["vmCreatedTxHash"]);
        }
      });
    });
  }
}

class ArbProvider extends ethers.providers.BaseProvider {
  constructor(managerUrl, contracts, provider) {
    super(123456789);
    this.chainId = 123456789;
    this.provider = provider;
    this.client = new ArbClient(managerUrl);
    let contractAddress = "0x5EBF59dBff8dCDa41610738634b396DfCB24A7c7";
    this.vmTracker = new ethers.Contract(
      contractAddress,
      vmTrackerJson["abi"],
      provider
    );
    this.contracts = {};
    for (var contract of contracts) {
      this.contracts[contract.address.toLowerCase()] = contract;
    }
  }

  async getSigner(index) {
    let wallet = new ArbWallet(
      this.client,
      this.contracts,
      this.provider.getSigner(index),
      this
    );
    await wallet.initialize();
    return wallet;
  }

  // value: *Value
  // logPreHash: hexString
  // logPostHash: hexString
  // logValHashes: []hexString
  // Returns true if the hash of value is in logPostHash and false otherwise
  processLogsProof(value, logPreHash, logPostHash, logValHashes) {
    const kh = (t, v) => ethers.utils.solidityKeccak256(t, v);
    const startHash = kh(["bytes32", "bytes32"], [logPreHash, value.hash()]);
    const checkHash = logValHashes.reduce(
      (acc, hash) => kh(["bytes32", "bytes32"], [acc, hash]),
      startHash
    );

    return logPostHash === checkHash;
  }

  // partialHash: hexString
  // logPostHash: hexString
  // validatorSigs: []hexString
  // Returns true if assertionHash is signed by all validators
  async processUnanimousAssertion(partialHash, logPostHash, validatorSigs) {
    const vmId = await this.getVmID();
    const validatorAddresses = await this.getValidatorAddresses();
    if (validatorAddresses.length !== validatorSigs.length) {
      console.error(
        "Expected:",
        validatorAddresses.length,
        "signatures.\n",
        "Received:",
        validatorSigs.length
      );
      return false;
    }

    let assertionHash = ethers.utils.solidityKeccak256(
      ["bytes32", "bytes32", "bytes32"],
      [vmId, partialHash, logPostHash]
    );

    let addresses = validatorSigs
      .map(sig =>
        ethers.utils
          .verifyMessage(ethers.utils.arrayify(assertionHash), sig)
          .toLowerCase()
          .slice(2)
      )
      .sort();

    for (let i = 0; i < validatorAddresses; i++) {
      if (validatorAddresses[i] !== addresses[i]) {
        console.error("Invalid signature");
        return false;
      }
    }
    return true;
  }

  // logPostHash: hexString
  // onChainTxHash: hexString
  // Returns true if assertionHash is logged by the onChainTxHash
  async processConfirmedDisputableAssertion(logPostHash, onChainTxHash) {
    let receipt = await this.provider.waitForTransaction(onChainTxHash);
    let events = receipt.logs.map(l => this.vmTracker.interface.parseLog(l));
    // DisputableAssertion Event
    let cda = events.find(event => event["name"] === EB_EVENT_CDA);
    if (cda) {
      const vmId = await this.getVmID();
      // Check correct VM
      if (cda.values.vmId !== vmId) {
        console.error(
          "DisputableAssertion Event is from a different VM:",
          cda.values.vmId,
          "\nExpected VM ID:",
          vmId
        );
        return false;
      }

      // Check correct logs hash
      if (cda.values.logsAccHash !== logPostHash) {
        console.error(
          "DisputableAssertion Event on-chain logPostHash is:",
          cda.values.logsAccHash,
          "\nExpected:",
          logPostHash
        );
        return false;
      }

      // DisputableAssertion is correct
      // TODO: must wait for finality (past the re-org period)
      return true;
    } else {
      console.error("DisputableAssertion", onChainTxHash, "not found on chain");
      return false;
    }
  }

  async getValidatorAddresses() {
    if (!this._validatorAddresses) {
      let eventTxHash = await this.client.getVMCreatedTxHash();
      let receipt = await this.provider.waitForTransaction(eventTxHash);
      let events = receipt.logs.map(l => this.vmTracker.interface.parseLog(l));
      let vmCreatedEvent = events.find(event => event["name"] === EB_EVENT_VMC);
      if (!vmCreatedEvent) {
        throw new Error("VMCreated Event not found");
      }

      // Get vmId
      const vmId = await this.getVmID();
      if (vmCreatedEvent.values.vmId !== vmId) {
        console.error(
          "VMCreated Event TxHash is from the wrong VM ID:",
          vmCreatedEvent.values.vmId,
          "\nExpected:",
          vmId
        );
        throw new Error("VMCreated Event vmId does not match");
      }

      // Cache the set of lowercase validator addresses (without "0x")
      this._validatorAddresses = vmCreatedEvent.values.validators
        .map(addr => addr.toLowerCase().slice(2))
        .sort();
      console.log("Validator Addresses are:", this._validatorAddresses);
    }
    return this._validatorAddresses;
  }

  async getVmID() {
    if (!this._vmId) {
      const vmId = await this.client.getVmID();
      // Guard against race condition
      if (!this._vmId) {
        this._vmId = vmId;
        console.log("VM ID is:", vmId);
      }
    }
    return this._vmId;
  }

  async getMessageResult(txHash) {
    let result = await this.client.getMessageResult(txHash);
    if (!result) {
      return null;
    }
    let { data, evmVal } = result;
    let {
      val,
      logPreHash,
      logPostHash,
      logValHashes,
      validatorSigs,
      partialHash,
      onChainTxHash
    } = data;

    const vmId = await this.getVmID();
    let txHashCheck = ethers.utils.solidityKeccak256(
      ["bytes32", "bytes32", "uint256", "bytes21"],
      [
        vmId,
        val
          .get(0)
          .get(0)
          .get(0)
          .hash(),
        evmVal.orig.value,
        ethers.utils.hexDataSlice(evmVal.orig.tokenType, 11)
      ]
    );

    // Check txHashCheck matches txHash
    if (txHash !== txHashCheck) {
      console.error("txHash did not match its pre-image", txHash, txHashCheck);
      return null;
    }

    // Step 1: prove that val is in logPostHash
    if (!this.processLogsProof(val, logPreHash, logPostHash, logValHashes)) {
      console.error("Failed to prove val is in logPostHash");
      return null;
    }

    // Step 2: prove that logPostHash is in assertion and assertion is valid
    if (validatorSigs && validatorSigs.length > 0) {
      if (
        !this.processUnanimousAssertion(partialHash, logPostHash, validatorSigs)
      ) {
        return null;
      }
    } else {
      // TODO: enable disputable assertion checks
      if (
        !this.processConfirmedDisputableAssertion(logPostHash, onChainTxHash)
      ) {
        return null;
      }
    }

    return {
      evmVal: evmVal,
      txHash: txHashCheck
    };
  }

  // This should return a Promise (and may throw errors)
  // method is the method name (e.g. getBalance) and params is an
  // object with normalized values passed in, depending on the method
  perform(method, params) {
    // console.log("perform", method, params)
    var self = this;
    switch (method) {
      case "getCode":
        if (self.contracts[params.address.toLowerCase()]) {
          return new Promise((resolve, reject) => {
            resolve(self.contracts[params.address.toLowerCase()].code);
          });
        }
        break;
      case "getBlockNumber":
        return this.client.getAssertionCount();
      case "getTransactionReceipt":
        return this.getMessageResult(params.transactionHash).then(result => {
          if (result) {
            let status = 0;
            if (
              result.evmVal.returnType == EVM_RETURN_CODE ||
              result.evmVal.returnType == EVM_STOP_CODE
            ) {
              status = 1;
            }
            return {
              to: result.evmVal.orig.contractID,
              from: result.evmVal.orig.caller,
              transactionIndex: 0,
              gasUsed: 1,
              blockHash: result.txHash,
              transactionHash: result.txHash,
              logs: [],
              blockNumber: result.evmVal.orig.blockHeight,
              confirmations: 1000,
              cumulativeGasUsed: 1,
              status: status
            };
          } else {
            return null;
          }
        });
      case "getTransaction":
        var getMessageRequest = () => {
          return self.getMessageResult(params.transactionHash).then(result => {
            if (result) {
              return {
                hash: result.txHash,
                blockHash: result.txHash,
                blockNumber: result.evmVal.orig.blockHeight,
                transactionIndex: 0,
                confirmations: 1000,
                from: result.evmVal.orig.caller,
                gasPrice: 1,
                gasLimit: 1,
                to: result.evmVal.orig.contractID,
                cumulativeGasUsed: 1,
                value: result.evmVal.orig.value,
                nonce: 0,
                data: result.evmVal.orig.data,
                status:
                  result.evmVal.returnType == EVM_RETURN_CODE ||
                  result.evmVal.returnType == EVM_STOP_CODE
              };
            } else {
              return null;
            }
          });
        };
        return promisePoller({
          taskFn: getMessageRequest,
          interval: 100,
          shouldContinue: (reason, value) => {
            if (reason) {
              return true;
            } else if (value) {
              return false;
            } else {
              return true;
            }
          }
        });
      case "getLogs":
        return this.client.findLogs(
          params.filter.fromBlock,
          params.filter.toBlock,
          params.filter.address,
          params.filter.topics
        );
    }
    let forwardResponse = self.provider.perform(method, params);
    console.log("Forwarding query to provider", method, forwardResponse);
    return forwardResponse;
  }

  async call(transaction) {
    let dest = await transaction.to;
    let contractData = this.contracts[dest.toLowerCase()];
    if (contractData) {
      var maxSeq = ethers.utils.bigNumberify(2);
      for (var i = 0; i < 255; i++) {
        maxSeq = maxSeq.mul(2);
      }
      maxSeq = maxSeq.sub(2);
      let arbMsg = new ArbValue.TupleValue([
        ArbValue.hexToSizedByteRange(transaction.data),
        new ArbValue.IntValue(dest),
        new ArbValue.IntValue(maxSeq)
      ]);
      let sender = await this.provider.getSigner(0).getAddress();
      return this.client.call(arbMsg, sender);
    } else {
      return self.provider.call(transaction);
    }
  }
}

class ArbWallet extends ethers.Signer {
  constructor(client, contracts, signer, provider) {
    super();
    this.contracts = contracts;
    this.signer = signer;
    this.provider = provider;
    this.client = client;
    this.vmTracker = provider.vmTracker.connect(signer);
    this.seq = 0;
  }

  async initialize() {
    if (this.seq > 0) {
      return;
    }

    return this.provider.provider.getBlockNumber().then(height => {
      var seq = ethers.utils.bigNumberify(height);
      for (var i = 0; i < 128; i++) {
        seq = seq.mul(2);
      }
      var timeStamp = Math.floor(Date.now());
      seq = seq.add(timeStamp);
      seq = seq.mul(2);
      this.seq = seq;
    });
  }

  getAddress() {
    return this.signer.getAddress();
  }

  async sendTransaction(transaction) {
    let self = this;
    let dest = await transaction.to;
    if (self.contracts[dest.toLowerCase()]) {
      self.seq = self.seq.add(2);
      let vmId = await self.provider.getVmID();
      let encodedData = ArbValue.hexToSizedByteRange(transaction.data);
      let arbMsg = new ArbValue.TupleValue([
        encodedData,
        new ArbValue.IntValue(dest),
        new ArbValue.IntValue(self.seq)
      ]);
      if (!transaction.value) {
        transaction.value = ethers.utils.bigNumberify(0);
      }
      let args = [
        vmId,
        arbMsg.hash(),
        transaction.value,
        ethers.utils.hexZeroPad("0x00", 21)
      ];
      let messageHash = ethers.utils.solidityKeccak256(
        ["bytes32", "bytes32", "uint256", "bytes21"],
        args
      );
      let tx = {
        hash: messageHash,
        from: await self.getAddress(),
        gasPrice: 1,
        gasLimit: 1,
        to: dest,
        value: transaction.value,
        nonce: self.seq,
        data: transaction.data
      };
      if (ethers.utils.bigNumberify(transaction.value).eq(0)) {
        let messageHashBytes = ethers.utils.arrayify(messageHash);
        let sig = await self.signer.signMessage(messageHashBytes);
        if (!self.pubkey) {
          self.pubkey = ethers.utils.recoverPublicKey(
            ethers.utils.arrayify(ethers.utils.hashMessage(messageHashBytes)),
            sig
          );
        }
        await self.client.sendMessage(arbMsg, sig, self.pubkey);
      } else {
        let tx = await self.vmTracker.sendEthMessage(
          vmId,
          ArbValue.marshal(arbMsg),
          {
            value: transaction.value
          }
        );

        await tx.wait();
      }
      return self.provider._wrapTransaction(tx, messageHash);
    } else {
      return self.signer.sendTransaction(transaction);
    }
  }
}

module.exports = ArbProvider;
