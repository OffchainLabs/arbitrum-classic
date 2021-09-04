import { ethers } from 'hardhat'
import { Provider } from '@ethersproject/providers'
import { Signer, BigNumber, BigNumberish } from 'ethers'
import { ContractTransaction, PayableOverrides } from '@ethersproject/contracts'
import { BytesLike } from '@ethersproject/bytes'

import { RollupUserFacet, RollupAdminFacet } from '../../build/types'
import { hexDataLength } from '@ethersproject/bytes'

const zerobytes32 =
  '0x0000000000000000000000000000000000000000000000000000000000000000'

export interface ExecutionState {
  machineHash: BytesLike
  inboxCount: BigNumberish
  gasUsed: BigNumberish
  sendCount: BigNumberish
  logCount: BigNumberish
  sendAcc: BytesLike
  logAcc: BytesLike
}

export interface Assertion {
  beforeState: ExecutionState
  afterState: ExecutionState
}

export interface Node {
  proposedBlock: number
  assertion: Assertion

  inboxMaxCount: BigNumberish
  nodeHash: BytesLike
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

type AssertionBytes32Fields = [BytesLike, BytesLike, BytesLike]
type AssertionIntFields = [
  BigNumberish,
  BigNumberish,
  BigNumberish,
  BigNumberish
]

export interface NodeCreatedEvent {
  nodeNum: BigNumberish
  parentNodeHash: BytesLike
  nodeHash: BytesLike
  executionHash: BytesLike
  inboxMaxCount: BigNumberish
  afterInboxAcc: BytesLike
  assertionBytes32Fields: [AssertionBytes32Fields, AssertionBytes32Fields]
  assertionIntFields: [AssertionIntFields, AssertionIntFields]
}

export function challengeRestHash(e: ExecutionState): BytesLike {
  return ethers.utils.solidityKeccak256(
    ['uint256', 'bytes32', 'bytes32', 'uint256', 'bytes32', 'uint256'],
    [e.inboxCount, e.machineHash, e.sendAcc, e.sendCount, e.logAcc, e.logCount]
  )
}

export function challengeHash(e: ExecutionState): BytesLike {
  return ethers.utils.solidityKeccak256(
    ['uint256', 'bytes32'],
    [e.gasUsed, challengeRestHash(e)]
  )
}

function executionStateBytes32Fields(
  e: ExecutionState
): AssertionBytes32Fields {
  return [e.machineHash, e.sendAcc, e.logAcc]
}

function executionStateIntFields(e: ExecutionState): AssertionIntFields {
  return [e.gasUsed, e.inboxCount, e.sendCount, e.logCount]
}

export function nodeStateHash(n: Node): string {
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
      n.assertion.afterState.gasUsed,
      n.assertion.afterState.machineHash,
      n.assertion.afterState.inboxCount,
      n.assertion.afterState.sendCount,
      n.assertion.afterState.logCount,
      n.assertion.afterState.sendAcc,
      n.assertion.afterState.logAcc,
      n.proposedBlock,
      n.inboxMaxCount,
    ]
  )
}

export function assertionGasUsed(a: Assertion): BigNumber {
  return ethers.BigNumber.from(a.afterState.gasUsed).sub(a.beforeState.gasUsed)
}

export function assertionExecutionHash(a: Assertion): BytesLike {
  return ethers.utils.solidityKeccak256(
    ['uint256', 'uint256', 'bytes32', 'bytes32'],
    [
      a.beforeState.gasUsed,
      assertionGasUsed(a),
      challengeHash(a.beforeState),
      challengeHash(a.afterState),
    ]
  )
}

export function makeAssertion(
  beforeState: ExecutionState,
  gasUsed: BigNumberish,
  afterMachineHash: BytesLike,
  messages: BytesLike[],
  sends: BytesLike[],
  logs: BytesLike[]
): Assertion {
  function buildAccumulator(base: BytesLike, vals: BytesLike[]): BytesLike {
    let acc = base
    for (const h of vals.map(val => ethers.utils.keccak256(val))) {
      acc = ethers.utils.solidityKeccak256(['bytes32', 'bytes32'], [acc, h])
    }
    return acc
  }
  return {
    beforeState: beforeState,
    afterState: {
      machineHash: afterMachineHash,
      inboxCount: ethers.BigNumber.from(beforeState.inboxCount).add(
        messages.length
      ),
      gasUsed: ethers.BigNumber.from(beforeState.gasUsed).add(gasUsed),
      sendCount: ethers.BigNumber.from(beforeState.sendCount).add(sends.length),
      logCount: ethers.BigNumber.from(beforeState.logCount).add(logs.length),
      sendAcc: buildAccumulator(beforeState.sendAcc, sends),
      logAcc: buildAccumulator(beforeState.logAcc, logs),
    },
  }
}

export class RollupContract {
  constructor(public rollup: RollupUserFacet) {}

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
    const isChild =
      challengeHash(prevNode.assertion.afterState) ==
      challengeHash(assertion.beforeState)
    const newNodeHash = ethers.utils.solidityKeccak256(
      ['bool', 'bytes32', 'bytes32', 'bytes32'],
      [
        !isChild,
        prevNode.nodeHash,
        assertionExecutionHash(assertion),
        afterInboxAcc,
      ]
    )
    const tx = await this.rollup.stakeOnNewNode(
      newNodeHash,
      [
        executionStateBytes32Fields(assertion.beforeState),
        executionStateBytes32Fields(assertion.afterState),
      ],
      [
        executionStateIntFields(assertion.beforeState),
        executionStateIntFields(assertion.afterState),
      ],
      parentNode.proposedBlock,
      parentNode.inboxMaxCount,
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
    const parsedEv = ev as any as {
      args: NodeCreatedEvent
    }
    const node = {
      assertion: assertion,
      proposedBlock: receipt.blockNumber!,
      inboxMaxCount: parsedEv.args.inboxMaxCount,
      nodeHash: newNodeHash,
    }
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
      [
        assertionExecutionHash(node1.assertion),
        assertionExecutionHash(node2.assertion),
      ],
      [node1.proposedBlock, node2.proposedBlock],
      [
        node1.assertion.afterState.inboxCount,
        node2.assertion.afterState.inboxCount,
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

  // removeZombieStaker(
  //   nodeNum: BigNumberish,
  //   stakerAddress: string
  // ): Promise<ContractTransaction> {
  //   return this.rollup.removeZombieStaker(nodeNum, stakerAddress)
  // }

  latestConfirmed(): Promise<BigNumber> {
    return this.rollup.latestConfirmed()
  }

  getNode(index: BigNumberish): Promise<string> {
    return this.rollup.getNode(index)
  }

  // async inboxMaxValue(): Promise<BytesLike> {
  //   const bridgeAddress = await this.rollup.delayedBridge()
  //   const Bridge = await ethers.getContractFactory('Bridge')
  //   const bridge = Bridge.attach(bridgeAddress) as Bridge
  //   const inboxInfo = await bridge.inboxInfo()
  //   return inboxInfo[1]
  // }

  currentRequiredStake(): Promise<BigNumber> {
    return this.rollup.currentRequiredStake()
  }
}

export async function forceCreateNode(
  rollupAdmin: RollupAdminFacet,
  newNodeHash: BytesLike,
  assertion: Assertion,
  batchProof: BytesLike,
  prevNode: Node,
  prevNodeIndex: BigNumberish
): Promise<{ tx: ContractTransaction; node: Node; event: NodeCreatedEvent }> {
  const tx = await rollupAdmin.forceCreateNode(
    newNodeHash,
    [
      executionStateBytes32Fields(assertion.beforeState),
      executionStateBytes32Fields(assertion.afterState),
    ],
    [
      executionStateIntFields(assertion.beforeState),
      executionStateIntFields(assertion.afterState),
    ],
    batchProof,
    prevNode.proposedBlock,
    prevNode.inboxMaxCount,
    prevNodeIndex
  )
  const receipt = await tx.wait()
  if (receipt.logs == undefined) {
    throw Error('expected receipt to have logs')
  }

  const evs = receipt.logs
    .map(log => {
      try {
        return rollupAdmin.interface.parseLog(log)
      } catch (e) {
        return undefined
      }
    })
    .filter(ev => ev && ev.name == 'NodeCreated')
  if (evs.length != 1) {
    throw Error('unique event not found')
  }
  const parsedEv = evs[0] as any as {
    args: NodeCreatedEvent
  }
  const node = {
    assertion: assertion,
    proposedBlock: receipt.blockNumber!,
    inboxMaxCount: parsedEv.args.inboxMaxCount,
    nodeHash: newNodeHash,
  }
  const event = parsedEv.args
  return { tx, node, event }
}
