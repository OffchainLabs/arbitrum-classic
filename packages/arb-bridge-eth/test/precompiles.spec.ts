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

import { ethers } from 'hardhat'
import { BytesLike } from '@ethersproject/bytes'
import { expect } from 'chai'
import { PrecompilesTester } from '../build/types'

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
    // test vectors from https://homes.esat.kuleuven.be/~nsmart/MPC/sha-256-test.txt

    const initialHashState =
      '0x6a09e667bb67ae853c6ef372a54ff53a510e527f9b05688c1f83d9ab5be0cd19'

    const input1: [BytesLike, BytesLike] = [
      ethers.constants.HashZero,
      ethers.constants.HashZero,
    ]

    const input2: [BytesLike, BytesLike] = [
      '0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f',
      '0x202122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f',
    ]

    const input3: [BytesLike, BytesLike] = [
      '0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF',
      '0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF',
    ]

    const input4: [BytesLike, BytesLike] = [
      '0x243F6A8885A308D313198A2E03707344A4093822299F31D0082EFA98EC4E6C89',
      '0x452821E638D01377BE5466CF34E90C6CC0AC29B7C97C50DD3f84D5B5b5470917',
    ]

    const output1 =
      '0xda5698be17b9b46962335799779fbeca8ce5d491c0d26243bafef9ea1837a9d8'
    const output2 =
      '0xfc99a2df88f42a7a7bb9d18033cdc6a20256755f9d5b9a5044a9cc315abe84a7'
    const output3 =
      '0xef0c748df4da50a8d6c43c013edc3ce76c9d9fa9a1458ade56eb86c0a64492d2'
    const output4 =
      '0xcf0ae4eb67d38ffeb94068984b22abde4e92bc548d14585e48dca8882d7b09ce'
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
