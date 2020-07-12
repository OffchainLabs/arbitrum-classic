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
