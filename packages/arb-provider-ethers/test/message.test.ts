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
/* eslint-env node, jest */
'use strict'

import { expect, use } from 'chai'
import { solidity } from 'ethereum-waffle'

import * as Message from '../src/lib/message'
import * as ethers from 'ethers'

use(solidity)

describe('Serialization', function () {
  it('L2Transaction', function () {
    const tx = new Message.L2Transaction(
      5436365554,
      756747657564,
      543345234,
      '0x7654745845675674',
      74567468457564,
      '0x648576998870892435'
    )
    const l2Message = new Message.L2Message(tx)
    const data = l2Message.asData()
    expect(data.length).to.equal(
      1 + 5 * 32 + ethers.utils.hexDataLength(tx.calldata)
    )
    const parsedL2Message = Message.L2Message.fromData(data)
    expect(parsedL2Message.message.kind).to.equal(
      Message.L2MessageCode.Transaction
    )
    const tx2 = parsedL2Message.message as Message.L2Transaction
    expect(tx2.maxGas).to.equal(tx.maxGas)
    expect(tx2.gasPriceBid).to.equal(tx.gasPriceBid)
    expect(tx2.sequenceNum).to.equal(tx.sequenceNum)
    expect(tx2.destAddress).to.equal(tx.destAddress)
    expect(tx2.payment).to.equal(tx.payment)
    expect(tx2.calldata.toString()).to.equal(tx.calldata.toString())
  })
})

describe('Request IDs', function () {
  it('Should hash a subtype 0 tx', async () => {
    const tx = Message.L2Transaction.fromData(
      '0x000000000000000000000000000000000000000000000000000000174876e800000000000000000000000000000000000000000000000000000000000074aa31000000000000000000000000000000000000000000000000000000000052e4d40000000000000000000000007f6999eb9d18a44784045d87f3c67cf22746e99500000000000000000000000000000000000000000000000000000000051b5aa4af5a25367951baa2ff6cd471c483f15fb90badb37c5821b6d95526a41a9504680b4e7c8b763a1b1d49d4955c8486216325253fec738dd7a9e28bf921119c160f0702448615bbda08313f6a8eb668d20bf5059875921e668a5bdf2c7fc4844592d2572bcd'
    )
    const sender = '0xe91e00167939cb6694d2c422acd208a007293948'
    // const chain = "0x037c4d7bbb0407d1e2c64981855ad8681d0d86d1"
    const chainId = 237941675624145
    const targetHash =
      '0x00532596242ba0ded0a8a17d8897344282fa1b29de676aa41aad6f737898e4a2'
    expect(tx.messageID(sender, chainId)).to.equal(targetHash)
  })
})
