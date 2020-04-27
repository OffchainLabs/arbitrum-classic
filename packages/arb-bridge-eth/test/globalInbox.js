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
const ethers = require('ethers')

function calcTxHash(chain, to, sequenceNum, value, messageData) {
  return web3.utils.soliditySha3(
    { t: 'address', v: chain },
    { t: 'address', v: to },
    { t: 'uint256', v: sequenceNum },
    { t: 'uint256', v: value },
    {
      t: 'bytes32',
      v: web3.utils.soliditySha3({ t: 'bytes', v: messageData }),
    }
  )
}
contract('GlobalInbox', accounts => {
  it('should make initial call', async () => {
    let global_inbox = await GlobalInbox.deployed()
    await global_inbox.sendTransactionMessage(
      '0xffffffffffffffffffffffffffffffffffffffff',
      '0xffffffffffffffffffffffffffffffffffffffff',
      2000,
      54254535454544,
      '0x'
    )
  })

  it('should make second call', async () => {
    let global_inbox = await GlobalInbox.deployed()
    await global_inbox.sendTransactionMessage(
      '0xffffffffffffffffffffffffffffffffffffffff',
      '0xffffffffffffffffffffffffffffffffffffffff',
      2000,
      54254535454544,
      '0x'
    )
  })

  it('should make bigger call', async () => {
    let global_inbox = await GlobalInbox.deployed()
    await global_inbox.sendTransactionMessage(
      '0xffffffffffffffffffffffffffffffffffffffff',
      '0xffffffffffffffffffffffffffffffffffffffff',
      2000,
      54254535454544,
      // 64 bytes of data
      '0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff'
    )
  })

  it('should make a batch call', async () => {
    let messageCount = 100
    let chain = '0xffffffffffffffffffffffffffffffffffffffff'
    let txDataTemplate = {
      to: '0xffffffffffffffffffffffffffffffffffffffff',
      sequenceNum: 2000,
      value: 54254535454544,
      messageData: '0x1254',
    }

    let transactionsData = []
    for (let i = 0; i < messageCount; i++) {
      transactionsData.push(txDataTemplate)
    }

    let tos = []
    let sequenceNums = []
    let values = []
    let messageData = '0x'
    let messageLengths = []
    let signatures = '0x'

    for (var i = 0; i < transactionsData.length; i++) {
      let txData = transactionsData[i]

      let txHash = calcTxHash(
        chain,
        txData['to'],
        txData['sequenceNum'],
        txData['value'],
        txData['messageData']
      )
      let signedTxHash = await web3.eth.sign(txHash, accounts[0])
      tos.push(txData['to'])
      sequenceNums.push(txData['sequenceNum'])
      values.push(txData['value'])
      messageLengths.push((txData['messageData'].length - 2) / 2)
      messageData += txData['messageData'].slice(2)
      signatures += signedTxHash.slice(2)
    }

    let global_inbox = await GlobalInbox.deployed()
    let tx = await global_inbox.deliverTransactionBatch(
      chain,
      tos,
      sequenceNums,
      values,
      messageLengths,
      messageData,
      signatures
    )

    let txObj = await web3.eth.getTransaction(tx.tx)

    console.log('Input data length', ethers.utils.hexDataLength(txObj.input))

    let txData = ethers.utils.defaultAbiCoder.decode(
      [
        'address',
        'address[]',
        'uint256[]',
        'uint256[]',
        'uint32[]',
        'bytes',
        'bytes',
      ],
      ethers.utils.hexDataSlice(txObj.input, 4)
    )

    //   console.log("txData", txData);

    //   console.log(await web3.eth.getTransaction(tx.tx));

    //   assert.equal(tx.logs.length, 1)

    //   let ev = tx.logs[0]

    //   assert.equal(
    //     ev.event,
    //     'TransactionMessageBatchDelivered',
    //     'Incorrect event type'
    //   )

    //   assert.equal(ev.args.tos.length, messageCount)
    //   assert.equal(ev.args.seqNumbers.length, messageCount)
    //   assert.equal(ev.args.values.length, messageCount)
    //   assert.equal(ev.args.dataLengths.length, messageCount)

    //   var dataOffset = 0

    //   for (var i = 0; i < messageCount; i++) {
    //     assert.equal(
    //       ev.args.tos[i].toLowerCase(),
    //       transactionsData[i]['to'],
    //       'Incorrect to address'
    //     )
    //     assert.equal(
    //       ev.args.seqNumbers[i],
    //       transactionsData[i]['sequenceNum'],
    //       'Incorrect sequence num'
    //     )
    //     assert.equal(
    //       ev.args.values[i],
    //       transactionsData[i]['value'],
    //       'Incorrect value'
    //     )
    //     assert.equal(
    //       ethers.utils.hexDataSlice(
    //         ev.args.data,
    //         dataOffset,
    //         dataOffset + ev.args.dataLengths[i].toNumber()
    //       ),
    //       transactionsData[i]['messageData'],
    //       'Incorrect message data'
    //     )

    //     let txHash = calcTxHash(
    //       chain,
    //       ev.args.tos[i],
    //       ev.args.seqNumbers[i],
    //       ev.args.values[i],
    //       ethers.utils.hexDataSlice(
    //         ev.args.data,
    //         dataOffset,
    //         dataOffset + ev.args.dataLengths[i].toNumber()
    //       )
    //     )

    //     const messageHashBytes = ethers.utils.arrayify(txHash)

    //     let sig = ethers.utils.hexDataSlice(
    //       ev.args.signatures,
    //       i*65,
    //       (i+1)*65
    //     );

    //     let recoveredAddress = ethers.utils.verifyMessage(messageHashBytes, ethers.utils.arrayify(sig));
    //     assert.equal(recoveredAddress, accounts[0], "Incorrect signature")
    //     dataOffset += ev.args.dataLengths[i].toNumber()
    //   }
  })
})
