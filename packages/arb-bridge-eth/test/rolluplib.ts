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
import { Bridge } from '../build/types/Bridge'

const zerobytes32 =
  '0x0000000000000000000000000000000000000000000000000000000000000000'

function bisectionChunkHash(
  start: BigNumberish,
  length: BigNumberish,
  startHash: BytesLike,
  endHash: BytesLike
): BytesLike {
  return ethers.utils.solidityKeccak256(
    ['uint256', 'uint256', 'bytes32', 'bytes32'],
    [start, length, startHash, endHash]
  )
}

function assertionHash(
  arbGasUsed: BigNumberish,
  assertionRest: BytesLike
): BytesLike {
  return ethers.utils.solidityKeccak256(
    ['uint256', 'bytes32'],
    [arbGasUsed, assertionRest]
  )
}

function assertionRestHash(
  inboxDelta: BytesLike,
  machineState: BytesLike,
  sendAcc: BytesLike,
  sendCount: BigNumberish,
  logAcc: BytesLike,
  logCount: BigNumberish
): BytesLike {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'bytes32', 'bytes32', 'uint256', 'bytes32', 'uint256'],
    [inboxDelta, machineState, sendAcc, sendCount, logAcc, logCount]
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
  gasUsed: BigNumberish
): BytesLike {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'bytes32', 'bytes32', 'uint256'],
    [inboxConsistency, inboxConsistency, execution, gasUsed]
  )
}

export class NodeState {
  constructor(
    public proposedBlock: number,
    public totalGasUsed: BigNumberish,
    public machineHash: BytesLike,
    public inboxTop: BytesLike,
    public inboxCount: BigNumberish,
    public totalSendCount: BigNumberish,
    public totalLogCount: BigNumberish,
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
        this.totalGasUsed,
        this.machineHash,
        this.inboxTop,
        this.inboxCount,
        this.totalSendCount,
        this.totalLogCount,
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
      ethers.BigNumber.from(this.prevNodeState.totalGasUsed).add(this.gasUsed),
      this.afterMachineHash,
      this.afterInboxHash,
      ethers.BigNumber.from(this.prevNodeState.inboxCount).add(
        this.inboxMessagesRead
      ),
      ethers.BigNumber.from(this.prevNodeState.totalSendCount).add(
        this.sendCount
      ),
      ethers.BigNumber.from(this.prevNodeState.totalLogCount).add(
        this.logCount
      ),
      inboxMaxCount
    )
  }

  inboxConsistencyHash(
    inboxMaxHash: BytesLike,
    inboxMaxCount: BigNumberish
  ): BytesLike {
    const afterMessageCount = ethers.BigNumber.from(
      this.prevNodeState.inboxCount
    ).add(this.inboxMessagesRead)
    return bisectionChunkHash(
      afterMessageCount,
      ethers.BigNumber.from(inboxMaxCount).sub(afterMessageCount),
      inboxMaxHash,
      this.afterInboxHash
    )
  }

  inboxDeltaHash(): BytesLike {
    return bisectionChunkHash(
      0,
      this.inboxMessagesRead,
      inboxDeltaHash(this.afterInboxHash, zerobytes32),
      inboxDeltaHash(this.prevNodeState.inboxTop, this.inboxDelta)
    )
  }

  startAssertionRestHash(): BytesLike {
    return assertionRestHash(
      this.inboxDelta,
      this.prevNodeState.machineHash,
      zerobytes32,
      0,
      zerobytes32,
      0
    )
  }

  startAssertionHash(): BytesLike {
    return assertionHash(0, this.startAssertionRestHash())
  }

  endAssertionRestHash(): BytesLike {
    return assertionRestHash(
      zerobytes32,
      this.afterMachineHash,
      this.sendAcc,
      this.sendCount,
      this.logAcc,
      this.logCount
    )
  }

  endAssertionHash(): BytesLike {
    return assertionHash(this.gasUsed, this.endAssertionRestHash())
  }

  executionHash(): BytesLike {
    return bisectionChunkHash(
      0,
      this.gasUsed,
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
    inboxMaxCount: BigNumberish
  ): BytesLike {
    return challengeRootHash(
      this.inboxConsistencyHash(inboxMaxHash, inboxMaxCount),
      this.inboxDeltaHash(),
      this.executionHash(),
      this.gasUsed
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
    BigNumberish
  ] {
    return [
      this.prevNodeState.proposedBlock,
      this.prevNodeState.totalGasUsed,
      this.prevNodeState.inboxCount,
      this.prevNodeState.totalSendCount,
      this.prevNodeState.totalLogCount,
      this.prevNodeState.inboxMaxCount,
      this.inboxMessagesRead,
      this.gasUsed,
      this.sendCount,
      this.logCount,
    ]
  }
}

export class Node {
  constructor(
    public assertion: Assertion,
    public blockCreated: number,
    public inboxMaxHash: BytesLike,
    public inboxMaxCount: BigNumberish
  ) {}

  afterNodeState(): NodeState {
    return this.assertion.createdNodeState(
      this.blockCreated,
      this.inboxMaxCount
    )
  }

  challengeFields(): [BytesLike, BytesLike, BytesLike] {
    return [
      this.assertion.inboxConsistencyHash(
        this.inboxMaxHash,
        this.inboxMaxCount
      ),
      this.assertion.inboxDeltaHash(),
      this.assertion.executionHash(),
    ]
  }
}

export class RollupContract {
  constructor(public rollup: Rollup) {}

  connect(signerOrProvider: Signer | Provider | string): RollupContract {
    return new RollupContract(this.rollup.connect(signerOrProvider))
  }

  newStake(
    tokenAmount: BigNumberish,
    overrides: PayableOverrides = {}
  ): Promise<ContractTransaction> {
    return this.rollup.newStake(tokenAmount, overrides)
  }

  async stakeOnNewNode(
    block: Block,
    assertion: Assertion,
    newNodeNum: BigNumberish
  ): Promise<{ tx: ContractTransaction; node: Node }> {
    const tx = await this.rollup.stakeOnNewNode(
      block.hash,
      block.number,
      newNodeNum,
      assertion.bytes32Fields(),
      assertion.intFields()
    )
    const receipt = await tx.wait()
    if (receipt.logs == undefined) {
      throw Error('expected receipt to have logs')
    }

    const ev = this.rollup.interface.parseLog(
      receipt.logs[receipt.logs.length - 1]
    )
    if (ev.name != 'NodeCreated') {
      throw 'wrong event type'
    }
    const parsedEv = (ev as any) as {
      args: { inboxMaxCount: BigNumberish; inboxMaxHash: BytesLike }
    }
    const node = new Node(
      assertion,
      receipt.blockNumber!,
      parsedEv.args.inboxMaxHash,
      parsedEv.args.inboxMaxCount
    )
    return { tx, node }
  }

  stakeOnExistingNode(block: Block, nodeNum: BigNumberish) {
    return this.rollup.stakeOnExistingNode(block.hash, block.number, nodeNum)
  }

  confirmNextNode(
    logAcc: BytesLike,
    messages: BytesLike[]
  ): Promise<ContractTransaction> {
    const messageData = ethers.utils.concat(messages)
    const messageLengths = messages.map(msg => msg.length)
    return this.rollup.confirmNextNode(logAcc, messageData, messageLengths)
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
    node1: Node,
    node2: Node
  ): Promise<ContractTransaction> {
    const fields1 = node1.challengeFields()
    const fields2 = node2.challengeFields()
    return this.rollup.createChallenge(
      [staker1Address, staker2Address],
      [nodeNum1, nodeNum2],
      [fields1[0], fields1[1], fields1[2], fields2[0], fields2[1], fields2[2]],
      [node1.blockCreated, node2.blockCreated]
    )
  }

  addToDeposit(
    staker: string,
    tokenAmount: BigNumberish,
    overrides: PayableOverrides = {}
  ): Promise<ContractTransaction> {
    return this.rollup.addToDeposit(staker, tokenAmount, overrides)
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

  getNode(index: BigNumberish): Promise<string> {
    return this.rollup.getNode(index)
  }

  async inboxMaxValue(): Promise<BytesLike> {
    const bridgeAddress = await this.rollup.bridge()
    const Bridge = await ethers.getContractFactory('Bridge')
    const bridge = Bridge.attach(bridgeAddress) as Bridge
    const inboxInfo = await bridge.inboxInfo()
    return inboxInfo[1]
  }

  currentRequiredStake(): Promise<BigNumber> {
    return this.rollup.currentRequiredStake()
  }
}
