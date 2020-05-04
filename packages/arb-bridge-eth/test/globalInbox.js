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
const MessageTester = artifacts.require('MessageTester')
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

async function generateTxData(accounts, chain, messageCount) {
  let txDataTemplate = {
    to: '0xffffffffffffffffffffffffffffffffffffffff',
    sequenceNum: 2000,
    value: 54254535454544,
    messageData: '0x00',
  }

  let transactionsData = []
  for (let i = 0; i < messageCount; i++) {
    transactionsData.push(txDataTemplate)
  }

  let data = '0x'

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

    let packedTxData = ethers.utils.solidityPack(
      ['uint16', 'address', 'uint256', 'uint256', 'bytes', 'bytes'],
      [
        (txData['messageData'].length - 2) / 2,
        txData['to'],
        txData['sequenceNum'],
        txData['value'],
        signedTxHash,
        txData['messageData'],
      ]
    )
    data += packedTxData.slice(2)
  }
  return data
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
    let messageCount = 500
    let chain = '0xffffffffffffffffffffffffffffffffffffffff'

    // console.log(data);

    let data = await generateTxData(accounts, chain, messageCount)

    let globalInbox = await GlobalInbox.deployed()
    let tx = await globalInbox.deliverTransactionBatch(chain, data)

    assert.equal(tx.logs.length, 1)

    let ev = tx.logs[0]

    assert.equal(
      ev.event,
      'TransactionMessageBatchDelivered',
      'Incorrect event type'
    )

    assert.equal(
      ev.args.chain.toLowerCase(),
      chain.toLowerCase(),
      'incorrect chain in event'
    )

    let txObj = await web3.eth.getTransaction(tx.tx)
    let [chainInput, txDataInput] = ethers.utils.defaultAbiCoder.decode(
      ['address', 'bytes'],
      ethers.utils.hexDataSlice(txObj.input, 4)
    )

    assert.equal(
      chainInput.toLowerCase(),
      chain.toLowerCase(),
      'incorrect chain from input'
    )

    assert.equal(
      txDataInput.toLowerCase(),
      data.toLowerCase(),
      'incorrect tx data from input'
    )
  })
})
