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

const GlobalInbox = artifacts.require("GlobalInbox");

const eutil = require("ethereumjs-util");

contract("GlobalInbox", accounts => {
  it("should make initial call", async () => {
    let global_inbox = await GlobalInbox.deployed();
    await global_inbox.sendTransactionMessage(
      "0xffffffffffffffffffffffffffffffffffffffff",
      "0xffffffffffffffffffffffffffffffffffffffff",
      2000,
      54254535454544,
      "0x"
    );
  });

  it("should make second call", async () => {
    let global_inbox = await GlobalInbox.deployed();
    await global_inbox.sendTransactionMessage(
      "0xffffffffffffffffffffffffffffffffffffffff",
      "0xffffffffffffffffffffffffffffffffffffffff",
      2000,
      54254535454544,
      "0x"
    );
  });

  it("should make bigger call", async () => {
    let global_inbox = await GlobalInbox.deployed();
    await global_inbox.sendTransactionMessage(
      "0xffffffffffffffffffffffffffffffffffffffff",
      "0xffffffffffffffffffffffffffffffffffffffff",
      2000,
      54254535454544,
      // 64 bytes of data
      "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
    );
  });

  it("should make a batch call", async () => {
    let chain = "0xffffffffffffffffffffffffffffffffffffffff";
    let txDataTemplate = {
      to: "0xffffffffffffffffffffffffffffffffffffffff",
      sequenceNum: 2000,
      value: 54254535454544,
      messageData: "0x"
    };

    let transactionsData = [];
    for (let i = 0; i < 100; i++) {
      transactionsData.push(txDataTemplate);
    }

    let tos = [];
    let sequenceNums = [];
    let values = [];
    let messageData = "0x";
    let messageLengths = [];
    let signatures = "0x";

    for (var i = 0; i < transactionsData.length; i++) {
      let txData = transactionsData[i];

      let txHash = web3.utils.soliditySha3(
        { t: "address", v: chain },
        { t: "address", v: txData["to"] },
        { t: "uint256", v: txData["sequenceNum"] },
        { t: "uint256", v: txData["value"] },
        { t: "bytes", v: txData["messageData"] }
      );
      let signedTxHash = await web3.eth.sign(txHash, accounts[0]);
      tos.push(txData["to"]);
      sequenceNums.push(txData["sequenceNum"]);
      values.push(txData["value"]);
      messageLengths.push((txData["messageData"].length - 2) / 2);
      messageData += txData["messageData"].slice(2);
      signatures += signedTxHash.slice(2);
    }

    let global_inbox = await GlobalInbox.deployed();
    await global_inbox.deliverTransactionBatch(
      chain,
      tos,
      sequenceNums,
      values,
      messageLengths,
      messageData,
      signatures
    );
  });
});
