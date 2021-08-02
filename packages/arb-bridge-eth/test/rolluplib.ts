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
import { hexDataLength } from '@ethersproject/bytes'

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

export function nodeHash(
  hasSibling: boolean,
  lastHash: BytesLike,
  assertionExecHash: BytesLike,
  inboxAcc: BytesLike
): BytesLike {
  return ethers.utils.solidityKeccak256(
    ['bool', 'bytes32', 'bytes32', 'bytes32'],
    [hasSibling, lastHash, assertionExecHash, inboxAcc]
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

export class ExecutionState {
  constructor(
    public gasUsed: BigNumberish,
    public machineHash: BytesLike,
    public inboxCount: BigNumberish,
    public sendCount: BigNumberish,
    public logCount: BigNumberish,
    public sendAcc: BytesLike,
    public logAcc: BytesLike
  ) {}

  challengeRestHash(): BytesLike {
    return assertionRestHash(
      this.inboxCount,
      this.machineHash,
      this.sendAcc,
      this.sendCount,
      this.logAcc,
      this.logCount
    )
  }

  challengeHash(): BytesLike {
    return assertionHash(this.gasUsed, this.challengeRestHash())
  }

  bytes32Fields(): [BytesLike, BytesLike, BytesLike] {
    return [this.machineHash, this.sendAcc, this.logAcc]
  }

  intFields(): [BigNumberish, BigNumberish, BigNumberish, BigNumberish] {
    return [this.gasUsed, this.inboxCount, this.sendCount, this.logCount]
  }
}

export class NodeState {
  constructor(
    public execState: ExecutionState,
    public proposedBlock: number,
    public inboxMaxCount: BigNumberish
  ) {}

  hash(): string {
    return ethers.utils.solidityKeccak256(
      [
        'uint256',
        'bytes32',
        'uint256',
        'uint256',
        'uint256',
        'bytes32',
        'bytes32',
        'uint256',
        'uint256',
      ],
      [
        this.execState.gasUsed,
        this.execState.machineHash,
        this.execState.inboxCount,
        this.execState.sendCount,
        this.execState.logCount,
        this.execState.sendAcc,
        this.execState.logAcc,
        this.proposedBlock,
        this.inboxMaxCount,
      ]
    )
  }

  equals(other: NodeState): boolean {
    return this.hash() == other.hash()
  }
}

function buildAccumulator(base: BytesLike, hashes: BytesLike[]): BytesLike {
  let acc = base
  for (const h of hashes) {
    const hash = ethers.utils.solidityKeccak256(['bytes'], [h])
    acc = ethers.utils.solidityKeccak256(['bytes32', 'bytes32'], [acc, hash])
  }
  return acc
}

export class Assertion {
  public afterState: ExecutionState

  constructor(
    public beforeState: NodeState,
    gasUsed: BigNumberish,
    afterMachineHash: BytesLike,
    messages: BytesLike[],
    sends: BytesLike[],
    logs: BytesLike[]
  ) {
    this.afterState = new ExecutionState(
      ethers.BigNumber.from(this.beforeState.execState.gasUsed).add(gasUsed),
      afterMachineHash,
      ethers.BigNumber.from(this.beforeState.execState.inboxCount).add(
        messages.length
      ),
      ethers.BigNumber.from(this.beforeState.execState.sendCount).add(
        sends.length
      ),
      ethers.BigNumber.from(this.beforeState.execState.logCount).add(
        logs.length
      ),
      buildAccumulator(this.beforeState.execState.sendAcc, sends),
      buildAccumulator(this.beforeState.execState.logAcc, logs)
    )
  }

  gasUsed(): BigNumber {
    return ethers.BigNumber.from(this.afterState.gasUsed).sub(
      this.beforeState.execState.gasUsed
    )
  }

  bytes32Fields(): [
    [BytesLike, BytesLike, BytesLike],
    [BytesLike, BytesLike, BytesLike]
  ] {
    return [
      this.beforeState.execState.bytes32Fields(),
      this.afterState.bytes32Fields(),
    ]
  }

  intFields(): [
    [BigNumberish, BigNumberish, BigNumberish, BigNumberish],
    [BigNumberish, BigNumberish, BigNumberish, BigNumberish]
  ] {
    return [this.beforeState.execState.intFields(), this.afterState.intFields()]
  }

  executionHash(): BytesLike {
    return bisectionChunkHash(
      this.beforeState.execState.gasUsed,
      this.gasUsed(),
      this.beforeState.execState.challengeHash(),
      this.afterState.challengeHash()
    )
  }
}

export class Node {
  public beforeState: NodeState
  public afterState: NodeState
  constructor(
    assertion: Assertion,
    blockCreated: number,
    inboxMaxCount: BigNumberish,
    public nodeHash: BytesLike
  ) {
    this.beforeState = assertion.beforeState
    this.afterState = new NodeState(
      assertion.afterState,
      blockCreated,
      inboxMaxCount
    )
  }

  gasUsed(): BigNumber {
    return ethers.BigNumber.from(
      ethers.BigNumber.from(this.afterState.execState.gasUsed).sub(
        this.beforeState.execState.gasUsed
      )
    )
  }

  checkTime(arbGasSpeedLimitPerBlock: BigNumberish): number {
    return this.gasUsed().div(arbGasSpeedLimitPerBlock).toNumber()
  }

  executionHash(): BytesLike {
    return bisectionChunkHash(
      this.beforeState.execState.gasUsed,
      this.gasUsed(),
      this.beforeState.execState.challengeHash(),
      this.afterState.execState.challengeHash()
    )
  }
}

export interface NodeCreatedEvent {
  nodeNum: BigNumberish
  parentNodeHash: BytesLike
  nodeHash: BytesLike
  executionHash: BytesLike
  inboxMaxCount: BigNumberish
  afterInboxAcc: BytesLike
  assertionBytes32Fields: [
    [BytesLike, BytesLike, BytesLike],
    [BytesLike, BytesLike, BytesLike]
  ]
  assertionIntFields: [
    [BigNumberish, BigNumberish, BigNumberish, BigNumberish],
    [BigNumberish, BigNumberish, BigNumberish, BigNumberish]
  ]
}

export class RollupContract {
  constructor(public rollup: Rollup) {}

  connect(signerOrProvider: Signer | Provider | string): RollupContract {
    return new RollupContract(this.rollup.connect(signerOrProvider))
  }

  newStake(overrides: PayableOverrides = {}): Promise<ContractTransaction> {
    return this.rollup.newStake(overrides)
  }

  async stakeOnNewNode(
    parentNode: Node,
    assertion: Assertion,
    afterInboxAcc: BytesLike,
    batchProof: BytesLike,
    prevNode?: Node
  ): Promise<{ tx: ContractTransaction; node: Node; event: NodeCreatedEvent }> {
    if (!prevNode) {
      prevNode = parentNode
    }
    const isChild = prevNode.afterState.equals(assertion.beforeState)
    const newNodeHash = nodeHash(
      !isChild,
      prevNode.nodeHash,
      assertion.executionHash(),
      afterInboxAcc
    )
    const tx = await this.rollup.stakeOnNewNode(
      newNodeHash,
      assertion.bytes32Fields(),
      assertion.intFields(),
      parentNode.afterState.proposedBlock,
      parentNode.afterState.inboxMaxCount,
      batchProof
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
      args: NodeCreatedEvent
    }
    const node = new Node(
      assertion,
      receipt.blockNumber!,
      parsedEv.args.inboxMaxCount,
      newNodeHash
    )
    const event = parsedEv.args
    return { tx, node, event }
  }

  stakeOnExistingNode(nodeNum: BigNumberish, nodeHash: BytesLike) {
    return this.rollup.stakeOnExistingNode(nodeNum, nodeHash)
  }

  confirmNextNode(
    prevSendAcc: BytesLike,
    prevSendCount: BigNumberish,
    sends: BytesLike[],
    afterlogAcc: BytesLike,
    afterLogCount: BigNumberish
  ): Promise<ContractTransaction> {
    const messageData = ethers.utils.concat(sends)
    const messageLengths = sends.map(msg => hexDataLength(msg))
    return this.rollup.confirmNextNode(
      prevSendAcc,
      messageData,
      messageLengths,
      BigNumber.from(prevSendCount).add(sends.length),
      afterlogAcc,
      afterLogCount
    )
  }

  rejectNextNode(stakerAddress: string): Promise<ContractTransaction> {
    return this.rollup.rejectNextNode(stakerAddress)
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
      [node1.executionHash(), node2.executionHash()],
      [node1.afterState.proposedBlock, node2.afterState.proposedBlock],
      [
        node1.afterState.execState.inboxCount,
        node2.afterState.execState.inboxCount,
      ]
    )
  }

  addToDeposit(
    staker: string,
    overrides: PayableOverrides = {}
  ): Promise<ContractTransaction> {
    return this.rollup.addToDeposit(staker, overrides)
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
