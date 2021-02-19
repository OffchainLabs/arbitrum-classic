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
  totalMessagesRead: BigNumberish,
  machineState: BytesLike,
  sendAcc: BytesLike,
  sendCount: BigNumberish,
  logAcc: BytesLike,
  logCount: BigNumberish
): BytesLike {
  return ethers.utils.solidityKeccak256(
    ['uint256', 'bytes32', 'bytes32', 'uint256', 'bytes32', 'uint256'],
    [totalMessagesRead, machineState, sendAcc, sendCount, logAcc, logCount]
  )
}

function challengeRootHash(
  execution: BytesLike,
  gasUsed: BigNumberish,
  maxMessageCount: BigNumberish
): BytesLike {
  return ethers.utils.solidityKeccak256(
    ['bytes32', 'uint256', 'uint256'],
    [execution, gasUsed, maxMessageCount]
  )
}

export class NodeState {
  constructor(
    public proposedBlock: number,
    public totalGasUsed: BigNumberish,
    public machineHash: BytesLike,
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
        'uint256',
        'uint256',
        'uint256',
        'uint256',
      ],
      [
        this.proposedBlock,
        this.totalGasUsed,
        this.machineHash,
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
  public inboxMessagesRead: BigNumberish
  public sendAcc: BytesLike
  public sendCount: BigNumberish
  public logAcc: BytesLike
  public logCount: BigNumberish

  constructor(
    public prevNodeState: NodeState,
    public gasUsed: BigNumberish,
    public afterMachineHash: BytesLike,
    messages: BytesLike[],
    sends: BytesLike[],
    logs: BytesLike[]
  ) {
    this.inboxMessagesRead = messages.length

    this.sendAcc = buildAccumulator(zerobytes32, sends)
    this.sendCount = sends.length

    this.logAcc = buildAccumulator(zerobytes32, logs)
    this.logCount = logs.length
  }

  afterMessageCount(): BigNumber {
    return ethers.BigNumber.from(this.prevNodeState.inboxCount).add(
      this.inboxMessagesRead
    )
  }

  createdNodeState(
    proposedBlock: number,
    inboxMaxCount: BigNumberish
  ): NodeState {
    return new NodeState(
      proposedBlock,
      ethers.BigNumber.from(this.prevNodeState.totalGasUsed).add(this.gasUsed),
      this.afterMachineHash,
      this.afterMessageCount(),
      ethers.BigNumber.from(this.prevNodeState.totalSendCount).add(
        this.sendCount
      ),
      ethers.BigNumber.from(this.prevNodeState.totalLogCount).add(
        this.logCount
      ),
      inboxMaxCount
    )
  }

  startAssertionRestHash(): BytesLike {
    return assertionRestHash(
      this.prevNodeState.inboxCount,
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
      this.afterMessageCount(),
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

  challengeRoot(inboxMaxCount: BigNumberish): BytesLike {
    return challengeRootHash(this.executionHash(), this.gasUsed, inboxMaxCount)
  }

  bytes32Fields(): [BytesLike, BytesLike, BytesLike, BytesLike] {
    return [
      this.prevNodeState.machineHash,
      this.sendAcc,
      this.logAcc,
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
    public inboxMaxCount: BigNumberish
  ) {}

  afterNodeState(): NodeState {
    return this.assertion.createdNodeState(
      this.blockCreated,
      this.inboxMaxCount
    )
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
      args: { inboxMaxCount: BigNumberish }
    }
    const node = new Node(
      assertion,
      receipt.blockNumber!,
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
    return this.rollup.createChallenge(
      [staker1Address, staker2Address],
      [nodeNum1, nodeNum2],
      [node1.assertion.executionHash(), node2.assertion.executionHash()],
      [node1.blockCreated, node2.blockCreated],
      [node1.assertion.afterMessageCount(), node2.assertion.afterMessageCount()]
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
