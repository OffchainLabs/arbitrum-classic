import { ethers } from 'hardhat'
import { Provider, Block } from '@ethersproject/providers'
import { Signer, BigNumber, BigNumberish } from 'ethers'
import {
  Contract,
  ContractTransaction,
  PayableOverrides,
} from '@ethersproject/contracts'
import { BytesLike } from '@ethersproject/bytes'

import { Rollup } from '../build/types/Rollup'

const zerobytes32 =
  '0x0000000000000000000000000000000000000000000000000000000000000000'

function bisectionChunkHash(
  length: BigNumberish,
  startHash: BytesLike,
  endHash: BytesLike
): BytesLike {
  return ethers.utils.solidityKeccak256(
    ['uint256', 'bytes32', 'bytes32'],
    [length, startHash, endHash]
  )
}

function assertionHash(
  inboxDelta: BytesLike,
  arbGasUsed: BigNumberish,
  outputAcc: BytesLike,
  machineState: BytesLike
): BytesLike {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'uint256', 'bytes32', 'bytes32'],
    [inboxDelta, arbGasUsed, outputAcc, machineState]
  )
}

function outputAccHash(
  sendAcc: BytesLike,
  sendCount: BigNumberish,
  logAcc: BytesLike,
  logCount: BigNumberish
): BytesLike {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'uint256', 'bytes32', 'uint256'],
    [sendAcc, sendCount, logAcc, logCount]
  )
}

function inboxDeltaHash(inboxAcc: BytesLike, deltaAcc: BytesLike): BytesLike {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'bytes32'],
    [inboxAcc, deltaAcc]
  )
}

function challengeRootHash(
  inboxConsistency: BytesLike,
  inboxDelta: BytesLike,
  execution: BytesLike,
  executionCheckTime: BigNumberish
): BytesLike {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'bytes32', 'bytes32', 'uint256'],
    [inboxConsistency, inboxConsistency, execution, executionCheckTime]
  )
}

export class NodeState {
  constructor(
    public proposedBlock: number,
    public stepsRun: BigNumberish,
    public machineHash: BytesLike,
    public inboxTop: BytesLike,
    public inboxCount: BigNumberish,
    public sendCount: BigNumberish,
    public logCount: BigNumberish,
    public inboxMaxCount: BigNumberish
  ) {}

  hash(): BytesLike {
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

function buildAccumulator(base: BytesLike, hashes: BytesLike[]): BytesLike {
  let acc = base
  for (const h of hashes) {
    acc = ethers.utils.solidityKeccak256(['bytes32', 'bytes32'], [acc, h])
  }
  return acc
}

export class Assertion {
  public inboxDelta: BytesLike
  public inboxMessagesRead: BigNumberish
  public sendAcc: BytesLike
  public sendCount: BigNumberish
  public logAcc: BytesLike
  public logCount: BigNumberish
  public afterInboxHash: BytesLike

  constructor(
    public prevNodeState: NodeState,
    public stepsExecuted: BigNumberish,
    public gasUsed: BigNumberish,
    public afterMachineHash: BytesLike,
    messages: BytesLike[],
    sends: BytesLike[],
    logs: BytesLike[]
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
      ethers.BigNumber.from(this.prevNodeState.stepsRun).add(
        this.stepsExecuted
      ),
      this.afterMachineHash,
      this.afterInboxHash,
      ethers.BigNumber.from(this.prevNodeState.inboxCount).add(
        this.inboxMessagesRead
      ),
      ethers.BigNumber.from(this.prevNodeState.sendCount).add(this.sendCount),
      ethers.BigNumber.from(this.prevNodeState.logCount).add(this.logCount),
      inboxMaxCount
    )
  }

  inboxConsistencyHash(
    inboxMaxHash: BytesLike,
    inboxMaxCount: BigNumberish
  ): BytesLike {
    return bisectionChunkHash(
      ethers.BigNumber.from(inboxMaxCount)
        .sub(this.prevNodeState.inboxCount)
        .sub(this.inboxMessagesRead),
      inboxMaxHash,
      this.afterInboxHash
    )
  }

  inboxDeltaHash(): BytesLike {
    return bisectionChunkHash(
      this.inboxMessagesRead,
      inboxDeltaHash(this.afterInboxHash, zerobytes32),
      inboxDeltaHash(this.prevNodeState.inboxTop, this.inboxDelta)
    )
  }

  startAssertionHash(): BytesLike {
    return assertionHash(
      this.inboxDelta,
      0,
      outputAccHash(zerobytes32, 0, zerobytes32, 0),
      this.prevNodeState.machineHash
    )
  }

  endAssertionHash(): BytesLike {
    return assertionHash(
      zerobytes32,
      this.gasUsed,
      outputAccHash(this.sendAcc, this.sendCount, this.logAcc, this.logCount),
      this.afterMachineHash
    )
  }

  executionHash(): BytesLike {
    return bisectionChunkHash(
      this.stepsExecuted,
      this.startAssertionHash(),
      this.endAssertionHash()
    )
  }

  checkTime(arbGasSpeedLimitPerBlock: BigNumberish): number {
    return ethers.BigNumber.from(this.gasUsed)
      .div(arbGasSpeedLimitPerBlock)
      .toNumber()
  }

  challengeRoot(
    inboxMaxHash: BytesLike,
    inboxMaxCount: BigNumberish,
    arbGasSpeedLimitPerBlock: BigNumberish
  ): BytesLike {
    return challengeRootHash(
      this.inboxConsistencyHash(inboxMaxHash, inboxMaxCount),
      this.inboxDeltaHash(),
      this.executionHash(),
      this.checkTime(arbGasSpeedLimitPerBlock)
    )
  }

  bytes32Fields(): [
    BytesLike,
    BytesLike,
    BytesLike,
    BytesLike,
    BytesLike,
    BytesLike,
    BytesLike
  ] {
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

  intFields(): [
    BigNumberish,
    BigNumberish,
    BigNumberish,
    BigNumberish,
    BigNumberish,
    BigNumberish,
    BigNumberish,
    BigNumberish,
    BigNumberish,
    BigNumberish,
    BigNumberish
  ] {
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
    block: Block,
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
    block: Block,
    assertion: Assertion,
    newNodeNum: BigNumberish,
    prevNum: BigNumberish,
    overrides: PayableOverrides = {}
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
    block: Block,
    nodeNum: BigNumberish,
    overrides: PayableOverrides = {}
  ) {
    return this.rollup.newStakeOnExistingNode(
      block.hash,
      block.number,
      nodeNum,
      overrides
    )
  }

  addStakeOnExistingNode(block: Block, nodeNum: BigNumberish) {
    return this.rollup.addStakeOnExistingNode(block.hash, block.number, nodeNum)
  }

  confirmNextNode(
    logAcc: BytesLike,
    messages: BytesLike[]
  ): Promise<ContractTransaction> {
    const messageData = ethers.utils.concat(messages)
    const messageLengths = messages.map(msg => msg.length)
    return this.rollup.confirmNextNode(logAcc, messageData, messageLengths)
  }

  rejectNextNodeOutOfOrder(): Promise<ContractTransaction> {
    return this.rollup.rejectNextNodeOutOfOrder()
  }

  rejectNextNode(
    successorWithStake: BigNumberish,
    stakerAddress: string
  ): Promise<ContractTransaction> {
    return this.rollup.rejectNextNode(successorWithStake, stakerAddress)
  }

  async createChallenge(
    staker1Address: string,
    nodeNum1: BigNumberish,
    staker2Address: string,
    nodeNum2: BigNumberish,
    assertion: Assertion,
    inboxMaxHash: BytesLike,
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

  addToDeposit(overrides: PayableOverrides = {}): Promise<ContractTransaction> {
    return this.rollup.addToDeposit(overrides)
  }

  reduceDeposit(amount: BigNumberish): Promise<ContractTransaction> {
    return this.rollup.reduceDeposit(amount)
  }

  returnOldDeposit(stakerAddress: string): Promise<ContractTransaction> {
    return this.rollup.returnOldDeposit(stakerAddress)
  }

  removeZombieStaker(
    nodeNum: BigNumberish,
    stakerAddress: string
  ): Promise<ContractTransaction> {
    return this.rollup.removeZombieStaker(nodeNum, stakerAddress)
  }

  latestConfirmed(): Promise<BigNumber> {
    return this.rollup.latestConfirmed()
  }

  nodes(index: BigNumberish): Promise<string> {
    return this.rollup.nodes(index)
  }

  inboxMaxValue(): Promise<BytesLike> {
    return this.rollup.inboxMaxValue()
  }

  currentRequiredStake(): Promise<BigNumber> {
    return this.rollup.currentRequiredStake()
  }
}
