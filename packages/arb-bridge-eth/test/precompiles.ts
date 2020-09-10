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

/* eslint-env node, mocha */

import { ethers } from '@nomiclabs/buidler'
import * as chai from 'chai'
import chaiAsPromised from 'chai-as-promised'
import { PrecompilesTester } from '../build/types/PrecompilesTester'
import { solidity } from 'ethereum-waffle'

chai.use(chaiAsPromised)
chai.use(solidity)

const { expect } = chai

let precompilesTester: PrecompilesTester

describe('Precompiles', () => {
  before(async () => {
    const PrecompilesTester = await ethers.getContractFactory(
      'PrecompilesTester'
    )
    precompilesTester = (await PrecompilesTester.deploy()) as PrecompilesTester
    await precompilesTester.deployed()
  })

  it('calculates sha256 compression function correctly', async () => {
    // const msg = new Message.EthMessage(dest, value)
    // const inMsg = new Message.IncomingMessage(msg, 1000, 5345346, sender, 65465)

    const input = [
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
      '0x0',
    ]

    const output = '0x1'
    expect(await precompilesTester.sha256Block(input)).to.equal(output)
  })
})
