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

const GlobalInbox = artifacts.require('GlobalInbox')
const ArbValue = require('arb-provider-ethers').ArbValue
const ValueTester = artifacts.require('ValueTester')

const eutil = require('ethereumjs-util')

const {
  expectEvent,
  expectRevert,
  time,
} = require('@openzeppelin/test-helpers')

async function getMessageData(sender, receiver, value_tester) {
  let msgType = 1
  let value = 50

  const msg = new ArbValue.TupleValue([
    new ArbValue.IntValue(1),
    new ArbValue.IntValue(web3.utils.toBN(sender).toString()),
    new ArbValue.TupleValue([
      new ArbValue.IntValue(web3.utils.toBN(receiver).toString()),
      new ArbValue.IntValue(value),
    ]),
  ])

  const msg_data = ArbValue.marshal(msg)
  let res = await value_tester.deserializeMessageData(msg_data, 0)

  assert.isTrue(res['0'], 'did not deserialize first part corrctly')

  let offset = res['1'].toNumber()
  assert.equal(res['2'].toNumber(), 1, 'Incorrect message type, must be ethMsg')
  assert.equal(res['3'], sender, 'Incorrect sender')

  let res2 = await value_tester.getEthMsgData(msg_data, offset)
  assert.isTrue(res2['0'], "value didn't deserialize correctly")
  assert.equal(res2['2'], receiver, 'Incorrect receiver')

  assert.equal(res2['3'].toNumber(), value, 'Incorrect value sent')

  return msg_data
}

contract('GlobalInbox', accounts => {
  // it("should make initial call", async () => {
  //   let global_inbox = await GlobalInbox.deployed();
  //   await global_inbox.sendTransactionMessage(
  //     "0xffffffffffffffffffffffffffffffffffffffff",
  //     "0xffffffffffffffffffffffffffffffffffffffff",
  //     2000,
  //     54254535454544,
  //     "0x"
  //   );
  // });

  // it("should make second call", async () => {
  //   let global_inbox = await GlobalInbox.deployed();
  //   await global_inbox.sendTransactionMessage(
  //     "0xffffffffffffffffffffffffffffffffffffffff",
  //     "0xffffffffffffffffffffffffffffffffffffffff",
  //     2000,
  //     54254535454544,
  //     "0x"
  //   );
  // });

  // it("should make bigger call", async () => {
  //   let global_inbox = await GlobalInbox.deployed();
  //   await global_inbox.sendTransactionMessage(
  //     "0xffffffffffffffffffffffffffffffffffffffff",
  //     "0xffffffffffffffffffffffffffffffffffffffff",
  //     2000,
  //     54254535454544,
  //     // 64 bytes of data
  //     "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
  //   );
  // });

  // it("should make a batch call", async () => {
  //   let chain = "0xffffffffffffffffffffffffffffffffffffffff";
  //   let txDataTemplate = {
  //     to: "0xffffffffffffffffffffffffffffffffffffffff",
  //     sequenceNum: 2000,
  //     value: 54254535454544,
  //     messageData: "0x"
  //   };

  //   let transactionsData = [];
  //   for (let i = 0; i < 100; i++) {
  //     transactionsData.push(txDataTemplate);
  //   }

  //   let tos = [];
  //   let sequenceNums = [];
  //   let values = [];
  //   let messageData = "0x";
  //   let messageLengths = [];
  //   let signatures = "0x";

  //   for (var i = 0; i < transactionsData.length; i++) {
  //     let txData = transactionsData[i];

  //     let txHash = web3.utils.soliditySha3(
  //       { t: "address", v: chain },
  //       { t: "address", v: txData["to"] },
  //       { t: "uint256", v: txData["sequenceNum"] },
  //       { t: "uint256", v: txData["value"] },
  //       { t: "bytes", v: txData["messageData"] }
  //     );
  //     let signedTxHash = await web3.eth.sign(txHash, accounts[0]);
  //     tos.push(txData["to"]);
  //     sequenceNums.push(txData["sequenceNum"]);
  //     values.push(txData["value"]);
  //     messageLengths.push((txData["messageData"].length - 2) / 2);
  //     messageData += txData["messageData"].slice(2);
  //     signatures += signedTxHash.slice(2);
  //   }

  //   let global_inbox = await GlobalInbox.deployed();
  //   await global_inbox.deliverTransactionBatch(
  //     chain,
  //     tos,
  //     sequenceNums,
  //     values,
  //     messageLengths,
  //     messageData,
  //     signatures
  //   );
  // });

  let chain_address = accounts[6]
  let nodeHash =
    '0x10c9d77c3846591fdfc3f966935819eb7dd71ebb7e71d5d081b880868ca33e4d'
  let messageIndex = 0
  let originalOwner = accounts[0]
  let address2 = accounts[1]
  let address3 = accounts[2]
  let curr_owner = accounts[0]

  it('tradeable-exits: initial', async () => {
    console.log('inside tradeable-exits initial...')
    let global_inbox = await GlobalInbox.deployed()

    curr_owner = await global_inbox.getPaymentOwner(
      originalOwner,
      nodeHash,
      messageIndex
    )
    curr_owner = web3.utils.toChecksumAddress(curr_owner)
    assert.equal(
      originalOwner,
      curr_owner,
      'current owner must be original owner.'
    )

    let reciept = await global_inbox.transferPayment(
      originalOwner,
      address2,
      nodeHash,
      messageIndex,
      {
        from: originalOwner,
      }
    )
    await expectEvent(reciept, 'PaymentTransfer')

    curr_owner = await global_inbox.getPaymentOwner(
      originalOwner,
      nodeHash,
      messageIndex
    )
    curr_owner = web3.utils.toChecksumAddress(curr_owner)
    assert.isTrue(
      curr_owner == address2,
      'current owner must be new owner (address2).'
    )
    console.log('valid owner 2')
  })

  it('tradeable-exits: subsequent transfers', async () => {
    console.log('inside tradeable-exits: subsequent transfers...')
    let global_inbox = await GlobalInbox.deployed()
    curr_owner = await global_inbox.getPaymentOwner(
      originalOwner,
      nodeHash,
      messageIndex
    )
    curr_owner = web3.utils.toChecksumAddress(curr_owner)
    assert.isTrue(curr_owner == address2, 'current owner must be address2.')

    let reciept1 = global_inbox.transferPayment(
      originalOwner,
      address2,
      nodeHash,
      messageIndex,
      {
        from: originalOwner,
      }
    )

    await expectRevert(reciept1, 'Must be payment owner.')

    let reciept2 = await global_inbox.transferPayment(
      originalOwner,
      address3,
      nodeHash,
      messageIndex,
      {
        from: curr_owner,
      }
    )
    await expectEvent(reciept2, 'PaymentTransfer')

    curr_owner = await global_inbox.getPaymentOwner(
      originalOwner,
      nodeHash,
      messageIndex
    )
    curr_owner = web3.utils.toChecksumAddress(curr_owner)
    assert.isTrue(
      curr_owner == address3,
      'current owner must be new owner (address3).'
    )

    let recieptr = global_inbox.transferPayment(
      originalOwner,
      address2,
      nodeHash,
      messageIndex,
      {
        from: address2,
      }
    )
    await expectRevert(recieptr, 'Must be payment owner.')

    let value_tester = await ValueTester.new()

    console.log('inside tradeable-exits: commiting transfers...')
    // let global_inbox = await GlobalInbox.deployed();

    let chain_balance = await global_inbox.getEthBalance(chain_address)
    console.log('chain_address_balance: ' + chain_balance)
    let curr_owner_balance = await global_inbox.getEthBalance(curr_owner)
    console.log('curr_owner_balance: ' + curr_owner_balance)
    let originalOwner_balance = await global_inbox.getEthBalance(originalOwner)
    console.log('originalOwner_balance: ' + originalOwner_balance)

    let reciept3 = await global_inbox.depositEthMessage(
      chain_address,
      originalOwner,
      {
        from: originalOwner,
        value: 50,
      }
    )
    await expectEvent(reciept3, 'EthDepositMessageDelivered')

    chain_balance = await global_inbox.getEthBalance(chain_address)
    console.log('chain_address_balance: ' + chain_balance)
    curr_owner_balance = await global_inbox.getEthBalance(curr_owner)
    console.log('curr_owner_balance: ' + curr_owner_balance)
    originalOwner_balance = await global_inbox.getEthBalance(originalOwner)
    console.log('originalOwner_balance: ' + originalOwner_balance)

    let msg_data = await getMessageData(
      chain_address,
      originalOwner,
      value_tester
    )
    let msgCounts = [1]
    let bytes = web3.utils.hexToBytes(nodeHash)
    let nodeHashes = [bytes]

    let reciept4 = await global_inbox.sendMessages(
      msg_data,
      msgCounts,
      nodeHashes,
      {
        from: chain_address,
      }
    )

    chain_balance = await global_inbox.getEthBalance(chain_address)
    console.log('chain_address_balance: ' + chain_balance)
    curr_owner_balance = await global_inbox.getEthBalance(curr_owner)
    console.log('curr_owner_balance: ' + curr_owner_balance)
    originalOwner_balance = await global_inbox.getEthBalance(originalOwner)
    console.log('originalOwner_balance: ' + originalOwner_balance)
    // it("tradeable-exits: commiting transfers", async () => {
    //   let value_tester = await ValueTester.new();

    //   console.log("inside tradeable-exits: commiting transfers...");
    //   let global_inbox = await GlobalInbox.deployed();

    //   let reciept3 = await global_inbox.depositEthMessage(
    //     chain_address,
    //     originalOwner,
    //     {
    //       from: originalOwner,
    //       value: 50
    //     });
    //   await expectEvent(reciept3, "EthDepositMessageDelivered");

    //   let msg_data = await getMessageData(chain_address, originalOwner, value_tester);
    //   let msgCounts = [1];
    //   let bytes = web3.utils.hexToBytes(nodeHash);
    //   let nodeHashes = [bytes];

    //   let reciept4 = await global_inbox.sendMessages(msg_data, msgCounts, nodeHashes);
    //   // await time.advanceBlock();
    //   // await time.advanceBlock();

    //   // let balance = await global_inbox.getEthBalance(address3);
    //   // console.log("balance: "+ balance);
    //   // let balance2 = await global_inbox.getEthBalance(chain_address);
    //   // console.log("balance2: "+ balance2);
  })
})
