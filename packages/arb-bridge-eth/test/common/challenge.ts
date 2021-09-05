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
import { BigNumberish, BigNumber } from 'ethers'
import { BytesLike } from '@ethersproject/bytes'
import { ContractTransaction } from '@ethersproject/contracts'

import {
  BridgeMock,
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

export interface ChallengeSegment {
  Start: BigNumberish
  Length: BigNumberish
}

export interface Bisection {
  ChallengedSegment: ChallengeSegment
  Cuts: BytesLike[]
}

class SequencerBatchItem {
  constructor(
    public lastSeqNum: BigNumberish,
    public accumulator: BytesLike,
    public totalDelayedCount: BigNumberish,
    public sequencerMessage: BytesLike
  ) {}
}

interface ChainTime {
  BlockNum: BigNumberish
  Timestamp: BigNumberish
}

export interface Message {
  Kind: number
  Sender: string
  InboxSeqNum: BigNumberish
  GasPrice: BigNumberish
  Data: BytesLike
  ChainTime: ChainTime
}

function messageToBytes(msg: Message): BytesLike {
  return ethers.utils.solidityPack(
    ['uint8', 'address', 'uint256', 'uint256', 'uint256', 'uint256', 'bytes'],
    [
      msg.Kind,
      msg.Sender,
      msg.ChainTime.BlockNum,
      msg.ChainTime.Timestamp,
      msg.InboxSeqNum,
      msg.GasPrice,
      msg.Data,
    ]
  )
}

export function newSequencerItem(
  totalDelayedCount: BigNumberish,
  msg: Message,
  prevAcc: BytesLike
): SequencerBatchItem {
  const inner = ethers.utils.solidityKeccak256(
    ['address', 'uint256', 'uint256'],
    [msg.Sender, msg.ChainTime.BlockNum, msg.ChainTime.Timestamp]
  )
  const acc = ethers.utils.solidityKeccak256(
    ['bytes32', 'uint256', 'bytes32', 'bytes32'],
    [prevAcc, msg.InboxSeqNum, inner, ethers.utils.keccak256(msg.Data)]
  )
  return new SequencerBatchItem(
    msg.InboxSeqNum,
    acc,
    totalDelayedCount,
    messageToBytes(msg)
  )
}

export function newDelayedItem(
  lastSeqNum: BigNumberish,
  totalDelayedCount: BigNumberish,
  prevAcc: BytesLike,
  prevDelayedCount: BigNumberish,
  delayedAcc: BytesLike
): SequencerBatchItem {
  const firstSeqNum = BigNumber.from(1)
    .add(lastSeqNum)
    .add(prevDelayedCount)
    .sub(totalDelayedCount)
  const acc = ethers.utils.solidityKeccak256(
    ['string', 'bytes32', 'uint256', 'uint256', 'uint256', 'bytes32'],
    [
      'Delayed messages:',
      prevAcc,
      firstSeqNum,
      prevDelayedCount,
      totalDelayedCount,
      delayedAcc,
    ]
  )
  return new SequencerBatchItem(lastSeqNum, acc, totalDelayedCount, '0x')
}

function bisectionChunkCount(
  segmentIndex: BigNumberish,
  segmentCount: BigNumberish,
  totalLength: BigNumberish
): BigNumber {
  const total = BigNumber.from(totalLength)
  let size = total.div(segmentCount)
  if (BigNumber.from(segmentIndex).eq(0)) {
    size = size.add(total.mod(segmentCount))
  }
  return size
}

function bisectionChunkHash(
  segmentStart: BigNumberish,
  segmentLength: BigNumberish,
  startHash: BytesLike,
  endHash: BytesLike
): BytesLike {
  return ethers.utils.solidityKeccak256(
    ['uint256', 'uint256', 'bytes32', 'bytes32'],
    [segmentStart, segmentLength, startHash, endHash]
  )
}

function buildMerkleTree(chunks: BytesLike[]): BytesLike[][] {
  const layers: BytesLike[][] = []
  layers.push(chunks)
  while (layers[layers.length - 1].length > 1) {
    const elements = layers[layers.length - 1]
    const nextLayer: BytesLike[] = []
    for (let i = 0; i < elements.length; i++) {
      if (i % 2 == 1) {
        continue
      }
      if (i + 1 >= elements.length) {
        nextLayer.push(elements[i])
      } else {
        const data = ethers.utils.solidityKeccak256(
          ['bytes32', 'bytes32'],
          [elements[i], elements[i + 1]]
        )
        nextLayer.push(data)
      }
    }
    layers.push(nextLayer)
  }
  return layers
}

function pathToInt(path: boolean[]): BigNumber {
  let route = BigNumber.from(0)
  for (const entry of path) {
    route = route.mul(2)
    if (entry) {
      route = route.add(1)
    }
  }
  return route
}

function merkleProof(
  layers: BytesLike[][],
  index: number
): { proof: BytesLike[]; path: BigNumber } {
  if (index == 0 && layers.length == 1) {
    return { proof: [], path: BigNumber.from(0) }
  }
  const proof: BytesLike[] = []
  let path: boolean[] = []
  for (const layer of layers) {
    const pairIndex = index % 2 == 0 ? index + 1 : index - 1
    if (pairIndex < layer.length) {
      path.push(index % 2 == 0)
      proof.push(layer[pairIndex])
    }
    index = Math.floor(index / 2)
  }
  path = path.reverse()
  return { proof, path: pathToInt(path) }
}

function calculateBisectionTree(bisection: Bisection): BytesLike[][] {
  const segmentCount = bisection.Cuts.length - 1
  let segmentStart = BigNumber.from(bisection.ChallengedSegment.Start)
  const chunks: BytesLike[] = []
  for (let i = 0; i < segmentCount; i++) {
    const segmentLength = bisectionChunkCount(
      i,
      segmentCount,
      bisection.ChallengedSegment.Length
    )
    const chunkHash = bisectionChunkHash(
      segmentStart,
      segmentLength,
      bisection.Cuts[i],
      bisection.Cuts[i + 1]
    )
    chunks.push(chunkHash)
    segmentStart = segmentStart.add(segmentLength)
  }
  return buildMerkleTree(chunks)
}

export function executeBisectMove(
  challenge: Challenge,
  move: BisectMove
): Promise<ContractTransaction> {
  const prevTree = calculateBisectionTree(move.PrevBisection)
  const { proof, path } = merkleProof(prevTree, move.SegmentToChallenge)
  return challenge.bisectExecution(
    proof,
    path,
    move.InconsistentSegment.Start,
    move.InconsistentSegment.Length,
    move.PrevBisection.Cuts[move.SegmentToChallenge + 1],
    move.StartState.gasUsed,
    challengeRestHash(move.StartState),
    move.SubCuts
  )
}

function getProver(opcode: number): number {
  if ((opcode >= 0xa1 && opcode <= 0xa6) || opcode == 0x70) {
    return 1
  } else if (opcode >= 0x20 && opcode <= 0x24) {
    return 2
  } else {
    return 0
  }
}

export function executeProveContinuedMove(
  challenge: Challenge,
  move: ProveContinuedMove
): Promise<ContractTransaction> {
  const prevTree = calculateBisectionTree(move.PrevBisection)
  const { proof, path } = merkleProof(prevTree, move.SegmentToChallenge)
  return challenge.proveContinuedExecution(
    proof,
    path,
    move.ChallengedSegment.Start,
    move.ChallengedSegment.Length,
    move.PrevBisection.Cuts[move.SegmentToChallenge + 1],
    move.PreviousCut.gasUsed,
    challengeRestHash(move.PreviousCut)
  )
}

export function executeOneStepProofMove(
  challenge: Challenge,
  move: OneStepProofMove
): Promise<ContractTransaction> {
  const prevTree = calculateBisectionTree(move.PrevBisection)
  const { proof, path } = merkleProof(prevTree, move.SegmentToChallenge)
  const prover = getProver(ethers.utils.arrayify(move.ProofData)[0])
  return challenge.oneStepProveExecution(
    proof,
    path,
    move.ChallengedSegment.Start,
    move.ChallengedSegment.Length,
    move.PrevBisection.Cuts[move.SegmentToChallenge + 1],
    move.PreviousCut.inboxCount,
    [move.PreviousCut.sendAcc, move.PreviousCut.logAcc],
    [
      move.PreviousCut.gasUsed,
      move.PreviousCut.sendCount,
      move.PreviousCut.logCount,
    ],
    move.ProofData,
    move.BufferProofData,
    prover
  )
}

export async function executeMove(
  challenge: Challenge,
  move: Move
): Promise<ContractTransaction | undefined> {
  if (move === null) {
    await ethers.provider.send('evm_mine', [])
    return
  }
  switch (move.Kind) {
    case 'Bisect': {
      return executeBisectMove(challenge, move)
    }
    case 'OneStepProof': {
      return executeOneStepProofMove(challenge, move)
    }
    case 'ProveContinuedExecution': {
      return executeProveContinuedMove(challenge, move)
    }
    case 'Timeout': {
      return challenge.timeout()
    }
  }
}

export interface BisectMove {
  Kind: 'Bisect'
  PrevBisection: Bisection
  StartState: ExecutionState
  SegmentToChallenge: number
  InconsistentSegment: ChallengeSegment
  SubCuts: BytesLike[]
}

export interface ProveContinuedMove {
  Kind: 'ProveContinuedExecution'
  Assertion: Assertion
  PrevBisection: Bisection
  SegmentToChallenge: number
  ChallengedSegment: ChallengeSegment
  PreviousCut: ExecutionState
}

export interface OneStepProofMove {
  Kind: 'OneStepProof'
  Assertion: Assertion
  PrevBisection: Bisection
  SegmentToChallenge: number
  ChallengedSegment: ChallengeSegment
  PreviousCut: ExecutionState
  ProofData: BytesLike
  BufferProofData: BytesLike
}

export interface TimeoutMove {
  Kind: 'Timeout'
}

export type Move =
  | BisectMove
  | ProveContinuedMove
  | OneStepProofMove
  | TimeoutMove
  | null

export class ChallengeDeployment {
  constructor(
    public challengeTester: ChallengeTester,
    public bridge: BridgeMock,
    public sequencerInbox: SequencerInbox
  ) {}

  async startChallenge(
    challengedAssertion: Assertion,
    asserter: string,
    challenger: string,
    asserterTimeLeft: BigNumberish,
    challengerTimeLeft: BigNumberish
  ): Promise<Challenge> {
    await this.challengeTester.startChallenge(
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
  owner: string
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

  const Bridge = await ethers.getContractFactory('BridgeMock')
  const bridge = await Bridge.deploy()
  await bridge.deployed()
  await bridge.initialize()

  const SequencerInbox = await ethers.getContractFactory('SequencerInbox')
  const sequencerInbox = await SequencerInbox.deploy()
  await sequencerInbox.deployed()
  await sequencerInbox.initialize(bridge.address, owner, owner)
  await sequencerInbox.setMaxDelay(100000000, 10000000000)
  return new ChallengeDeployment(challengeTester, bridge, sequencerInbox)
}

export async function bisectExecution(
  challenge: Challenge,
  merkleNodes: BytesLike[],
  merkleRoute: BigNumberish,
  assertion: Assertion,
  chainHashes: BytesLike[]
): Promise<ContractTransaction> {
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
