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

    // test vectors from https://homes.esat.kuleuven.be/~nsmart/MPC/sha-256-test.txt

    const initialHashState =
      '0x6a09e667bb67ae853c6ef372a54ff53a510e527f9b05688c1f83d9ab5be0cd19'

    const input1 = ['0x0', '0x0']

    const input2 = [
      '0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF',
      '0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF',
    ]

    const input3 = [
      '0x243F6A8885A308D313198A2E03707344A4093822299F31D0082EFA98EC4E6C89',
      '0x452821E638D01377BE5466CF34E90C6CC0AC29B7C97C50DD3f84D5B5b5470917',
    ]

    const input4 = [
      '0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f',
      '0x202122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f',
    ]

    const output1 =
      '98757204029056169846550522095972853119225293520952391190693233024408770554328'
    const output2 =
      '108124777405892723987183679135752579682699194348193935577055847415223704392402'
    const output3 =
      '93648008072072659063844430594620143031399579590857428012818982161960985627086'
    const output4 =
      '114254289553293425223588562168656701874573032962834543531534328118719287297191'
    expect(
      await precompilesTester.sha256Block(input1, initialHashState)
    ).to.equal(output1)
    expect(
      await precompilesTester.sha256Block(input2, initialHashState)
    ).to.equal(output2)
    expect(
      await precompilesTester.sha256Block(input3, initialHashState)
    ).to.equal(output3)
    expect(
      await precompilesTester.sha256Block(input4, initialHashState)
    ).to.equal(output4)
  })
})
