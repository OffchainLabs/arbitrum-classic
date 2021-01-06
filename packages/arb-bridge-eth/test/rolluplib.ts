import { ethers } from '@nomiclabs/buidler'
import { Provider } from 'ethers/providers'
import { Signer, ContractTransaction, providers, utils } from 'ethers'
import { Arrayish, BigNumber, BigNumberish } from 'ethers/utils'

import { TransactionOverrides } from '../build/types'

import { Rollup } from '../build/types/Rollup'

const zerobytes32 =
  '0x0000000000000000000000000000000000000000000000000000000000000000'

function bisectionChunkHash(
  length: BigNumberish,
  startHash: string,
  endHash: string
): string {
  return ethers.utils.solidityKeccak256(
    ['uint256', 'bytes32', 'bytes32'],
    [length, startHash, endHash]
  )
}

function assertionHash(
  inboxDelta: string,
  arbGasUsed: BigNumberish,
  outputAcc: string,
  machineState: string
): string {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'uint256', 'bytes32', 'bytes32'],
    [inboxDelta, arbGasUsed, outputAcc, machineState]
  )
}

function outputAccHash(
  sendAcc: string,
  sendCount: BigNumberish,
  logAcc: string,
  logCount: BigNumberish
): string {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'uint256', 'bytes32', 'uint256'],
    [sendAcc, sendCount, logAcc, logCount]
  )
}

function inboxDeltaHash(inboxAcc: string, deltaAcc: string): string {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'bytes32'],
    [inboxAcc, deltaAcc]
  )
}

function challengeRootHash(
  inboxConsistency: string,
  inboxDelta: string,
  execution: string,
  executionCheckTime: BigNumberish
): string {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'bytes32', 'bytes32', 'uint256'],
    [inboxConsistency, inboxConsistency, execution, executionCheckTime]
  )
}

export class NodeState {
  constructor(
    public proposedBlock: number,
    public stepsRun: BigNumberish,
    public machineHash: string,
    public inboxTop: string,
    public inboxCount: BigNumberish,
    public sendCount: BigNumberish,
    public logCount: BigNumberish,
    public inboxMaxCount: BigNumberish
  ) {}

  hash(): string {
    return ethers.utils.solidityKeccak256(
      [
        'uint256',
        'uint256',
        'bytes32',
        'bytes32',
        'uint256',
        'uint256',
        'uint256',
        'uint256',
      ],
      [
        this.proposedBlock,
        this.stepsRun,
        this.machineHash,
        this.inboxTop,
        this.inboxCount,
        this.sendCount,
        this.logCount,
        this.inboxMaxCount,
      ]
    )
  }
}

function buildAccumulator(base: string, hashes: string[]): string {
  let acc = base
  for (const h of hashes) {
    acc = ethers.utils.solidityKeccak256(['bytes32', 'bytes32'], [acc, h])
  }
  return acc
}

export class Assertion {
  public inboxDelta: string
  public inboxMessagesRead: BigNumberish
  public sendAcc: string
  public sendCount: BigNumberish
  public logAcc: string
  public logCount: BigNumberish
  public afterInboxHash: string

  constructor(
    public prevNodeState: NodeState,
    public stepsExecuted: BigNumberish,
    public gasUsed: BigNumberish,
    public afterMachineHash: string,
    messages: string[],
    sends: string[],
    logs: string[]
  ) {
    this.inboxDelta = buildAccumulator(zerobytes32, messages.reverse())
    this.inboxMessagesRead = messages.length
    this.afterInboxHash = buildAccumulator(prevNodeState.inboxTop, messages)

    this.sendAcc = buildAccumulator(zerobytes32, sends)
    this.sendCount = sends.length

    this.logAcc = buildAccumulator(zerobytes32, logs)
    this.logCount = logs.length
  }

  createdNodeState(
    proposedBlock: number,
    inboxMaxCount: BigNumberish
  ): NodeState {
    return new NodeState(
      proposedBlock,
      utils.bigNumberify(this.prevNodeState.stepsRun).add(this.stepsExecuted),
      this.afterMachineHash,
      this.afterInboxHash,
      utils
        .bigNumberify(this.prevNodeState.inboxCount)
        .add(this.inboxMessagesRead),
      utils.bigNumberify(this.prevNodeState.sendCount).add(this.sendCount),
      utils.bigNumberify(this.prevNodeState.logCount).add(this.logCount),
      inboxMaxCount
    )
  }

  inboxConsistencyHash(
    inboxMaxHash: string,
    inboxMaxCount: BigNumberish
  ): string {
    return bisectionChunkHash(
      ethers.utils
        .bigNumberify(inboxMaxCount)
        .sub(this.prevNodeState.inboxCount)
        .sub(this.inboxMessagesRead),
      inboxMaxHash,
      this.afterInboxHash
    )
  }

  inboxDeltaHash(): string {
    return bisectionChunkHash(
      this.inboxMessagesRead,
      inboxDeltaHash(this.afterInboxHash, zerobytes32),
      inboxDeltaHash(this.prevNodeState.inboxTop, this.inboxDelta)
    )
  }

  executionHash(): string {
    return bisectionChunkHash(
      this.stepsExecuted,
      assertionHash(
        this.inboxDelta,
        0,
        outputAccHash(zerobytes32, 0, zerobytes32, 0),
        this.prevNodeState.machineHash
      ),
      assertionHash(
        zerobytes32,
        this.gasUsed,
        outputAccHash(this.sendAcc, this.sendCount, this.logAcc, this.logCount),
        this.afterMachineHash
      )
    )
  }

  checkTime(arbGasSpeedLimitPerBlock: BigNumberish): number {
    return ethers.utils
      .bigNumberify(this.gasUsed)
      .div(arbGasSpeedLimitPerBlock)
      .toNumber()
  }

  challengeRoot(
    inboxMaxHash: string,
    inboxMaxCount: BigNumberish,
    arbGasSpeedLimitPerBlock: BigNumberish
  ): string {
    return challengeRootHash(
      this.inboxConsistencyHash(inboxMaxHash, inboxMaxCount),
      this.inboxDeltaHash(),
      this.executionHash(),
      this.checkTime(arbGasSpeedLimitPerBlock)
    )
  }

  bytes32Fields(): string[] {
    return [
      this.prevNodeState.machineHash,
      this.prevNodeState.inboxTop,
      this.inboxDelta,
      this.sendAcc,
      this.logAcc,
      this.afterInboxHash,
      this.afterMachineHash,
    ]
  }

  intFields(): BigNumberish[] {
    return [
      this.prevNodeState.proposedBlock,
      this.prevNodeState.stepsRun,
      this.prevNodeState.inboxCount,
      this.prevNodeState.sendCount,
      this.prevNodeState.logCount,
      this.prevNodeState.inboxMaxCount,
      this.stepsExecuted,
      this.inboxMessagesRead,
      this.gasUsed,
      this.sendCount,
      this.logCount,
    ]
  }
}

export class RollupContract {
  constructor(public rollup: Rollup) {}

  connect(signerOrProvider: Signer | Provider | string): RollupContract {
    return new RollupContract(this.rollup.connect(signerOrProvider))
  }

  addStakeOnNewNode(
    block: providers.Block,
    assertion: Assertion,
    newNodeNum: BigNumberish
  ): Promise<ContractTransaction> {
    return this.rollup.addStakeOnNewNode(
      block.hash,
      block.number,
      newNodeNum,
      assertion.bytes32Fields(),
      assertion.intFields()
    )
  }

  newStakeOnNewNode(
    block: providers.Block,
    assertion: Assertion,
    newNodeNum: BigNumberish,
    prevNum: BigNumberish,
    overrides?: TransactionOverrides
  ): Promise<ContractTransaction> {
    return this.rollup.newStakeOnNewNode(
      block.hash,
      block.number,
      newNodeNum,
      prevNum,
      assertion.bytes32Fields(),
      assertion.intFields(),
      overrides
    )
  }

  newStakeOnExistingNode(
    block: providers.Block,
    nodeNum: BigNumberish,
    overrides?: TransactionOverrides
  ) {
    return this.rollup.newStakeOnExistingNode(
      block.hash,
      block.number,
      nodeNum,
      overrides
    )
  }

  addStakeOnExistingNode(block: providers.Block, nodeNum: BigNumberish) {
    return this.rollup.addStakeOnExistingNode(block.hash, block.number, nodeNum)
  }

  confirmNextNode(
    logAcc: Arrayish,
    messages: Arrayish[]
  ): Promise<ContractTransaction> {
    const messageData = utils.concat(messages)
    const messageLengths = messages.map(msg => msg.length)
    return this.rollup.confirmNextNode(logAcc, messageData, messageLengths)
  }

  rejectNextNodeOutOfOrder(): Promise<ContractTransaction> {
    return this.rollup.rejectNextNodeOutOfOrder()
  }

  async createChallenge(
    staker1Address: string,
    nodeNum1: BigNumberish,
    staker2Address: string,
    nodeNum2: BigNumberish,
    assertion: Assertion,
    inboxMaxHash: string,
    inboxMaxCount: BigNumberish
  ): Promise<ContractTransaction> {
    return this.rollup.createChallenge(
      staker1Address,
      nodeNum1,
      staker2Address,
      nodeNum2,
      assertion.inboxConsistencyHash(inboxMaxHash, inboxMaxCount),
      assertion.inboxDeltaHash(),
      assertion.executionHash(),
      assertion.checkTime(await this.rollup.arbGasSpeedLimitPerBlock())
    )
  }

  latestConfirmed(): Promise<BigNumber> {
    return this.rollup.latestConfirmed()
  }

  nodes(index: BigNumberish): Promise<string> {
    return this.rollup.nodes(index)
  }

  inboxMaxValue(): Promise<string> {
    return this.rollup.inboxMaxValue()
  }
}
