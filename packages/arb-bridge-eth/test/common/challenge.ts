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
import { BigNumberish } from 'ethers'
import { BytesLike } from '@ethersproject/bytes'

import {
  Bridge,
  Challenge,
  ChallengeTester,
  SequencerInbox,
} from '../../build/types'

import {
  Assertion,
  challengeRestHash,
  challengeHash,
  assertionExecutionHash,
  ExecutionState,
  assertionGasUsed,
} from './rolluplib'

interface ChallengeSegment {
  Start: BigNumberish
  Length: BigNumberish
}

interface Bisection {
  ChallengedSegment: ChallengeSegment
  Cuts: BytesLike[]
}

interface BisectMove {
  Kind: 'Bisect'
  PrevBisection: Bisection
  StartState: ExecutionState
  SegmentToChallenge: number
  InconsistentSegment: ChallengeSegment
  SubCuts: BytesLike[]
}

interface ProveContinuedMove {
  Kind: 'ProveContinued'
}

interface OneStepProofMove {
  Kind: 'OneStepProof'
}

type Move = BisectMove | ProveContinuedMove | OneStepProofMove

export class ChallengeDeployment {
  constructor(
    public challengeTester: ChallengeTester,
    public bridge: Bridge,
    public sequencerInbox: SequencerInbox
  ) {}

  async startChallenge(
    challengedAssertion: Assertion,
    asserter: string,
    challenger: string,
    asserterTimeLeft: BigNumberish,
    challengerTimeLeft: BigNumberish
  ): Promise<Challenge> {
    this.challengeTester.startChallenge(
      assertionExecutionHash(challengedAssertion),
      challengedAssertion.afterState.inboxCount,
      asserter,
      challenger,
      asserterTimeLeft,
      challengerTimeLeft,
      this.sequencerInbox.address,
      this.bridge.address
    )
    const challengeAddress = await this.challengeTester.challenge()
    const Challenge = await ethers.getContractFactory('Challenge')
    return Challenge.attach(challengeAddress)
  }
}

export async function setupChallengeTest(
  sequencer: string
): Promise<ChallengeDeployment> {
  const OneStepProof = await ethers.getContractFactory('OneStepProof')
  const osp = await OneStepProof.deploy()
  await osp.deployed()

  const OneStepProof2 = await ethers.getContractFactory('OneStepProof2')
  const osp2 = await OneStepProof2.deploy()
  await osp2.deployed()

  const OneStepProof3 = await ethers.getContractFactory('OneStepProofHash')
  const osp3 = await OneStepProof3.deploy()
  await osp3.deployed()

  const ChallengeTester = await ethers.getContractFactory('ChallengeTester')
  const challengeTester = await ChallengeTester.deploy([
    osp.address,
    osp2.address,
    osp3.address,
  ])
  await challengeTester.deployed()

  const Bridge = await ethers.getContractFactory('Bridge')
  const bridge = await Bridge.deploy()
  await bridge.deployed()
  await bridge.initialize()

  const RollupMock = await ethers.getContractFactory('RollupMock')
  const rollupMock = await RollupMock.deploy()
  await rollupMock.deployed()
  await rollupMock.setMock(15, 900)

  const SequencerInbox = await ethers.getContractFactory('SequencerInbox')
  const sequencerInbox = await SequencerInbox.deploy()
  await sequencerInbox.deployed()
  await sequencerInbox.initialize(bridge.address, sequencer, rollupMock.address)
  return new ChallengeDeployment(challengeTester, bridge, sequencerInbox)
}

// export async function bisectExecution2(challenge: Challenge, move: BisectMove) {
//   move.SubCuts
//   return await challenge.bisectExecution(
//     merkleNodes,
//     merkleRoute,
//     assertion.beforeState.gasUsed,
//     assertionGasUsed(assertion),
//     challengeHash(assertion.afterState),
//     assertion.beforeState.gasUsed,
//     challengeRestHash(assertion.beforeState),
//     chainHashes
//   )
// }

export async function bisectExecution(
  challenge: Challenge,
  merkleNodes: BytesLike[],
  merkleRoute: BigNumberish,
  assertion: Assertion,
  chainHashes: BytesLike[]
) {
  return await challenge.bisectExecution(
    merkleNodes,
    merkleRoute,
    assertion.beforeState.gasUsed,
    assertionGasUsed(assertion),
    challengeHash(assertion.afterState),
    assertion.beforeState.gasUsed,
    challengeRestHash(assertion.beforeState),
    chainHashes
  )
}
