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
import { ethers } from 'hardhat'
import { BigNumber, Signer } from 'ethers'
import { expect } from 'chai'
import { Challenge } from '../build/types'
import { initializeAccounts } from './utils'
import {
  setupChallengeTest,
  ChallengeDeployment,
  Message,
} from './common/challenge'
import { Assertion, makeAssertion, challengeHash } from './common/rolluplib'
import {
  bisectExecution,
  Move,
  executeMove,
  newDelayedItem,
  newSequencerItem,
} from './common/challenge'
import * as fs from 'fs'

const printGas = false

interface ChallengeSpec {
  ChallengedAssertion: Assertion
  Messages: Message[]
  Moves: Move[]
  AsserterError: string | undefined
}

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

describe('ReplayChallenges', () => {
  before(async () => {
    accounts = await initializeAccounts()
  })
  const files = fs.readdirSync('./test/challenges')
  for (const filename of files) {
    if (!filename.endsWith('json')) {
      continue
    }
    const file = fs.readFileSync('./test/challenges/' + filename)
    let challengeData: ChallengeSpec
    try {
      challengeData = JSON.parse(file.toString()) as ChallengeSpec
    } catch (e) {
      console.log(`Failed to load ${file}`)
      throw e
    }
    describe(`proofs from ${filename}`, function () {
      before(async () => {
        challengeDeployment = await setupChallengeTest(
          await accounts[0].getAddress()
        )
      })
      it('should setup messages', async function () {
        const bridge = challengeDeployment.bridge
        if (challengeData.Messages.length > 1) {
          throw Error('more than one message not supported')
        }
        for (const message of challengeData.Messages) {
          await bridge.deliverMessageToInboxTest(
            message.Kind,
            message.Sender,
            message.ChainTime.BlockNum,
            message.ChainTime.Timestamp,
            message.GasPrice,
            ethers.utils.keccak256(message.Data)
          )
          const count = await bridge.messageCount()
          const delayedAcc = await bridge.inboxAccs(count.sub(1))
          const delayedItem = newDelayedItem(
            0,
            1,
            ethers.constants.HashZero,
            0,
            delayedAcc
          )
          const endOfBlockMessage: Message = {
            Kind: 6,
            Sender: ethers.constants.AddressZero,
            InboxSeqNum: 1,
            GasPrice: 0,
            Data: '0x',
            ChainTime: message.ChainTime,
          }
          const endOfBlockItem = newSequencerItem(
            1,
            endOfBlockMessage,
            delayedItem.accumulator
          )
          const batchMetadata = [
            0,
            message.ChainTime.BlockNum,
            message.ChainTime.Timestamp,
            1,
            BigNumber.from(delayedAcc),
          ]
          await challengeDeployment.sequencerInbox.addSequencerL2Batch(
            '0x',
            [],
            batchMetadata,
            endOfBlockItem.accumulator
          )
        }
      })

      let challenge: Challenge
      let challengedAssertion: Assertion
      it('should initiate challenge', async function () {
        challengedAssertion = challengeData.ChallengedAssertion
        challenge = await challengeDeployment.startChallenge(
          challengedAssertion,
          await accounts[0].getAddress(),
          await accounts[1].getAddress(),
          10,
          10
        )
      })

      it('should make moves', async function () {
        let player = 1
        let i = 0
        for (const move of challengeData.Moves) {
          const chal = challenge.connect(accounts[player])
          // if (move) {
          //   console.log("Making move", move.Kind)
          // }
          const nextMove = executeMove(chal, move)
          if (
            i == challengeData.Moves.length - 1 &&
            challengeData.AsserterError
          ) {
            const prefix = 'execution reverted: '
            if (challengeData.AsserterError.startsWith(prefix)) {
              challengeData.AsserterError = challengeData.AsserterError.substr(
                prefix.length
              )
            }
            await expect(nextMove).to.be.revertedWith(
              challengeData.AsserterError
            )
          } else {
            const maybeTx = await nextMove
            if (maybeTx) {
              const receipt = await maybeTx.wait()
              if (printGas) {
                console.log(
                  `${move?.Kind} gas used ${receipt.gasUsed.toNumber()}`
                )
              }
            }
          }

          player += 1
          player %= 2
          i++
        }
      })
    })
  }
})
