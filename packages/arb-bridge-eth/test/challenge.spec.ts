/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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
import { ethers, Signer } from 'ethers'
import { Challenge } from '../build/types'
import { initializeAccounts } from './utils'
import { setupChallengeTest, ChallengeDeployment } from './common/challenge'
import { Assertion, makeAssertion, challengeHash } from './common/rolluplib'
import { bisectExecution } from './common/challenge'

const initialVmState =
  '0x9900000000000000000000000000000000000000000000000000000000000000'

let accounts: Signer[]
let challengeDeployment: ChallengeDeployment

describe('Challenge', () => {
  before(async () => {
    accounts = await initializeAccounts()
    challengeDeployment = await setupChallengeTest(
      await accounts[0].getAddress()
    )
  })

  let challenge: Challenge
  let challengedAssertion: Assertion
  it('should initiate challenge', async function () {
    challengedAssertion = makeAssertion(
      {
        gasUsed: 0,
        machineHash: initialVmState,
        inboxCount: 0,
        sendCount: 0,
        logCount: 0,
        sendAcc: ethers.constants.HashZero,
        logAcc: ethers.constants.HashZero,
      },
      10000000,
      ethers.constants.HashZero,
      [],
      [],
      []
    )

    challenge = await challengeDeployment.startChallenge(
      challengedAssertion,
      await accounts[0].getAddress(),
      await accounts[1].getAddress(),
      100,
      100
    )
  })

  it('should bisect execution', async function () {
    const chunks = Array(401).fill(
      challengeHash(challengedAssertion.beforeState)
    )
    const tx = await bisectExecution(
      challenge.connect(accounts[1]),
      [],
      0,
      challengedAssertion,
      chunks
    )
    const receipt = await tx.wait()
    console.log('Bisection gas used', receipt.gasUsed.toNumber())
  })
})
