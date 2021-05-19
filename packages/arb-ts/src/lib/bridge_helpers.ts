import { ContractReceipt, ethers } from 'ethers'
import { ArbTokenBridge__factory } from './abi/factories/ArbTokenBridge__factory'
import { EthERC20Bridge__factory } from './abi/factories/EthERC20Bridge__factory'
import { Outbox__factory } from './abi/factories/Outbox__factory'
import { OutboxEntry__factory } from './abi/factories/OutboxEntry__factory'

import { Bridge__factory } from './abi/factories/Bridge__factory'
import { Inbox__factory } from './abi/factories/Inbox__factory'
import { ArbSys__factory } from './abi/factories/ArbSys__factory'
import { Rollup__factory } from './abi/factories/Rollup__factory'
import { OutboxEntry } from './abi/OutboxEntry'

import { providers, utils, constants } from 'ethers'
import { BigNumber, Contract, Signer } from 'ethers'
import { ARB_SYS_ADDRESS } from './l2Bridge'
import { Bridge } from './bridge'

export const addressToSymbol = (erc20L1Address: string) => {
  return erc20L1Address.substr(erc20L1Address.length - 3).toUpperCase() + '?'
}

// TODO: can we import these interfaces directly from typechain?
export interface L2ToL1EventResult {
  caller: string
  destination: string
  uniqueId: BigNumber
  batchNumber: BigNumber
  indexInBatch: BigNumber
  arbBlockNum: BigNumber
  ethBlockNum: BigNumber
  timestamp: string
  callvalue: BigNumber
  data: string
}

export interface WithdrawTokenEventResult {
  id: BigNumber
  l1Address: string
  amount: BigNumber
  destination: string
  exitNum: BigNumber
  txHash: string
}

export interface DepositTokenEventResult {
  destination: string
  sender: string
  seqNum: BigNumber
  tokenType: 0 | 1 | 2
  amount: BigNumber
  tokenAddress: string
}

export interface BuddyDeployEventResult {
  _sender: string
  _contract: string
  withdrawalId: BigNumber
  success: boolean
}

export interface OutboxProofData {
  batchNumber: BigNumber
  proof: string[]
  path: BigNumber
  l2Sender: string
  l1Dest: string
  l2Block: BigNumber
  l1Block: BigNumber
  timestamp: BigNumber
  amount: BigNumber
  calldataForL1: string
}

export interface ActivateCustomTokenResult {
  seqNum: BigNumber
  l1Addresss: string
  l2Address: string
}

export interface OutBoxTransactionExecuted {
  destAddr: string
  l2Sender: string
  outboxIndex: BigNumber
  transactionIndex: BigNumber
}

export enum OutgoingMessageState {
  NOT_FOUND,
  UNCONFIRMED,
  CONFIRMED,
  EXECUTED,
}

export type ChainIdOrProvider = BigNumber | providers.Provider

const NODE_INTERFACE_ADDRESS = '0x00000000000000000000000000000000000000C8'

export class BridgeHelper {
  static getTokenWithdrawEventData = async (
    destinationAddress: string,
    l2BridgeAddr: string,
    l2Provider: providers.Provider
  ) => {
    const contract = ArbTokenBridge__factory.connect(l2BridgeAddr, l2Provider)
    const iface = contract.interface
    const tokenWithdrawEvent = iface.getEvent('WithdrawToken')
    const tokenWithdrawTopic = iface.getEventTopic(tokenWithdrawEvent)

    const topics = [
      tokenWithdrawTopic,
      null,
      utils.hexZeroPad(destinationAddress, 32),
    ]

    const logs = await l2Provider.getLogs({
      address: l2BridgeAddr,
      // @ts-ignore
      topics,
      fromBlock: 0,
      toBlock: 'latest',
    })

    return logs.map(log => {
      const data = { ...iface.parseLog(log).args, txHash: log.transactionHash }
      return (data as unknown) as WithdrawTokenEventResult
    })
  }

  static calculateL2TransactionHash = async (
    inboxSequenceNumber: BigNumber,
    chainIdOrL2Provider: ChainIdOrProvider
  ) => {
    const l2ChainId = BigNumber.isBigNumber(chainIdOrL2Provider)
      ? chainIdOrL2Provider
      : BigNumber.from((await chainIdOrL2Provider.getNetwork()).chainId)

    return utils.keccak256(
      utils.concat([
        utils.zeroPad(l2ChainId.toHexString(), 32),
        utils.zeroPad(
          BridgeHelper.bitFlipSeqNum(inboxSequenceNumber).toHexString(),
          32
        ),
      ])
    )
  }

  static bitFlipSeqNum = (seqNum: BigNumber) => {
    return seqNum.or(BigNumber.from(1).shl(255))
  }
  /**
   * Calculates hash of L2 side of a "retryable" transaction (L1 to L2 message, message type 9)
   */
  static calculateL2RetryableTransactionHash = async (
    inboxSequenceNumber: BigNumber,
    chainIdOrL2Provider: ChainIdOrProvider
  ): Promise<string> => {
    const requestID = await BridgeHelper.calculateL2TransactionHash(
      inboxSequenceNumber,
      chainIdOrL2Provider
    )
    return utils.keccak256(
      utils.concat([
        utils.zeroPad(requestID, 32),
        utils.zeroPad(BigNumber.from(1).toHexString(), 32),
      ])
    )
  }

  static calculateRetryableAutoReedemTxnHash = async (
    inboxSequenceNumber: BigNumber,
    chainIdOrL2Provider: ChainIdOrProvider
  ): Promise<string> => {
    const requestID = await BridgeHelper.calculateL2TransactionHash(
      inboxSequenceNumber,
      chainIdOrL2Provider
    )
    return utils.keccak256(
      utils.concat([
        utils.zeroPad(requestID, 32),
        utils.zeroPad(BigNumber.from(0).toHexString(), 32),
      ])
    )
  }

  /**
   * Return receipt of retryable transaction after execution
   */
  static waitForRetriableReceipt = async (
    seqNum: BigNumber,
    l2Provider: providers.Provider
  ) => {
    const l2RetriableHash = await BridgeHelper.calculateL2RetryableTransactionHash(
      seqNum,
      l2Provider
    )
    return l2Provider.waitForTransaction(l2RetriableHash)
  }

  static getL2Transaction = async (
    l2TransactionHash: string,
    l2Provider: providers.Provider
  ) => {
    const txReceipt = await l2Provider.getTransactionReceipt(l2TransactionHash)
    if (!txReceipt) throw new Error("Can't find L2 transaction receipt?")
    return txReceipt
  }

  static getL1Transaction = async (
    l1TransactionHash: string,
    l1Provider: providers.Provider
  ) => {
    const txReceipt = await l1Provider.getTransactionReceipt(l1TransactionHash)
    if (!txReceipt) throw new Error("Can't find L1 transaction receipt?")
    return txReceipt
  }

  static getBuddyDeployInL2Transaction = async (
    l2Transaction: providers.TransactionReceipt
  ) => {
    const iface = new utils.Interface([
      `event Deployed(address indexed _sender, address indexed _contract, uint256 indexed withdrawalId, bool _success)`,
    ])
    const DeployedEvent = iface.getEvent('Deployed')
    const eventTopic = iface.getEventTopic(DeployedEvent)
    const logs = l2Transaction.logs.filter(log => log.topics[0] === eventTopic)
    return logs.map(
      log => (iface.parseLog(log).args as unknown) as BuddyDeployEventResult
    )
  }

  static getDepositTokenEventData = async (
    l1Transaction: providers.TransactionReceipt,
    l2BridgeAddress: string
  ): Promise<Array<DepositTokenEventResult>> => {
    const factory = new EthERC20Bridge__factory()
    // TODO: does this work?
    const contract = factory.attach(l2BridgeAddress)
    const iface = contract.interface
    const event = iface.getEvent('DepositToken')
    const eventTopic = iface.getEventTopic(event)
    const logs = l1Transaction.logs.filter(log => log.topics[0] === eventTopic)
    return logs.map(
      log => (iface.parseLog(log).args as unknown) as DepositTokenEventResult
    )
  }

  static getActivateCustomTokenEventResult = async (
    l1Transaction: providers.TransactionReceipt,
    l1BridgeAddress: string
  ): Promise<Array<ActivateCustomTokenResult>> => {
    const factory = new EthERC20Bridge__factory()
    const contract = factory.attach(l1BridgeAddress)
    const iface = contract.interface
    const event = iface.getEvent('ActivateCustomToken')
    const eventTopic = iface.getEventTopic(event)

    const logs = l1Transaction.logs.filter(log => {
      return log.topics[0] === eventTopic
    })
    return logs.map(
      log => (iface.parseLog(log).args as unknown) as ActivateCustomTokenResult
    )
  }

  static getWithdrawalsInL2Transaction = async (
    l2Transaction: providers.TransactionReceipt,
    l2Provider: providers.Provider
  ): Promise<Array<L2ToL1EventResult>> => {
    // TODO: can we use dummies to get interface?
    const contract = ArbSys__factory.connect(ARB_SYS_ADDRESS, l2Provider)
    const iface = contract.interface
    const l2ToL1Event = iface.getEvent('L2ToL1Transaction')
    const eventTopic = iface.getEventTopic(l2ToL1Event)

    const logs = l2Transaction.logs.filter(log => log.topics[0] === eventTopic)

    return logs.map(
      log => (iface.parseLog(log).args as unknown) as L2ToL1EventResult
    )
  }

  static getCoreBridgeFromInbox = (
    inboxAddress: string,
    l1Provider: providers.Provider
  ) => {
    const contract = Inbox__factory.connect(inboxAddress, l1Provider)
    return contract.functions.bridge().then(([res]) => res)
  }

  static getInboxSeqNumFromContractTransaction = async (
    l1Transaction: providers.TransactionReceipt,
    inboxAddress: string
  ) => {
    const factory = new Inbox__factory()
    const contract = factory.attach(inboxAddress)
    const iface = contract.interface
    const messageDelivered = iface.getEvent('InboxMessageDelivered')
    const messageDeliveredFromOrigin = iface.getEvent(
      'InboxMessageDeliveredFromOrigin'
    )

    const eventTopics = {
      InboxMessageDelivered: iface.getEventTopic(messageDelivered),
      InboxMessageDeliveredFromOrigin: iface.getEventTopic(
        messageDeliveredFromOrigin
      ),
    }

    const logs = l1Transaction.logs.filter(
      log =>
        log.topics[0] === eventTopics.InboxMessageDelivered ||
        log.topics[0] === eventTopics.InboxMessageDeliveredFromOrigin
    )

    if (logs.length === 0) return undefined
    return logs.map(log => BigNumber.from(log.topics[1]))
  }

  static tryGetProof = async (
    batchNumber: BigNumber,
    indexInBatch: BigNumber,
    l2Provider: providers.Provider,
    retryDelay = 500
  ): Promise<{
    proof: Array<string>
    path: BigNumber
    l2Sender: string
    l1Dest: string
    l2Block: BigNumber
    l1Block: BigNumber
    timestamp: BigNumber
    amount: BigNumber
    calldataForL1: string
  }> => {
    const contractInterface = new utils.Interface([
      `function lookupMessageBatchProof(uint256 batchNum, uint64 index)
          external
          view
          returns (
              bytes32[] memory proof,
              uint256 path,
              address l2Sender,
              address l1Dest,
              uint256 l2Block,
              uint256 l1Block,
              uint256 timestamp,
              uint256 amount,
              bytes memory calldataForL1
          )`,
    ])
    const nodeInterface = new Contract(
      NODE_INTERFACE_ADDRESS,
      contractInterface
    ).connect(l2Provider)

    try {
      const res = await nodeInterface.callStatic.lookupMessageBatchProof(
        batchNumber,
        indexInBatch
      )
      return res
    } catch (e) {
      const expectedError = "batch doesn't exist"
      if (
        e &&
        e.error &&
        e.error.message &&
        e.error.message === expectedError
      ) {
        console.log(
          'Withdrawal detected, but batch not created yet. Going to wait a bit.'
        )
      } else {
        console.log("Withdrawal proof didn't work. Not sure why")
        console.log(e)
        console.log('Going to try again after waiting')
      }
      await BridgeHelper.wait(retryDelay)
      console.log('New attempt starting')
      // TODO: should exponential backoff?
      return BridgeHelper.tryGetProof(
        batchNumber,
        indexInBatch,
        l2Provider,
        retryDelay
      )
    }
  }

  static wait = (ms: number) => new Promise(res => setTimeout(res, ms))

  static tryGetProofOnce = async (
    batchNumber: BigNumber,
    indexInBatch: BigNumber,
    l2Provider: providers.Provider
  ): Promise<{
    proof: Array<string>
    path: BigNumber
    l2Sender: string
    l1Dest: string
    l2Block: BigNumber
    l1Block: BigNumber
    timestamp: BigNumber
    amount: BigNumber
    calldataForL1: string
  } | null> => {
    const nodeInterfaceAddress = '0x00000000000000000000000000000000000000C8'

    const contractInterface = new utils.Interface([
      `function lookupMessageBatchProof(uint256 batchNum, uint64 index)
          external
          view
          returns (
              bytes32[] memory proof,
              uint256 path,
              address l2Sender,
              address l1Dest,
              uint256 l2Block,
              uint256 l1Block,
              uint256 timestamp,
              uint256 amount,
              bytes memory calldataForL1
          )`,
    ])
    const nodeInterface = new Contract(
      nodeInterfaceAddress,
      contractInterface
    ).connect(l2Provider)

    try {
      const res = await nodeInterface.callStatic.lookupMessageBatchProof(
        batchNumber,
        indexInBatch
      )
      return res
    } catch (e) {
      const expectedError = "batch doesn't exist"
      if (
        e &&
        e.error &&
        e.error.message &&
        e.error.message === expectedError
      ) {
        console.log('Withdrawal detected, but batch not created yet.')
      } else {
        console.log("Withdrawal proof didn't work. Not sure why")
        console.log(e)
      }
    }
    return null
  }

  static getOutboxEntry = async (
    batchNumber: BigNumber,
    outboxAddress: string,
    l1Provider: providers.Provider
  ): Promise<string> => {
    const iface = new ethers.utils.Interface([
      'function outboxes(uint256) public view returns (address)',
      'function outboxesLength() public view returns (uint256)',
    ])
    const outbox = new ethers.Contract(outboxAddress, iface).connect(l1Provider)
    const len: BigNumber = await outbox.outboxesLength()
    if (batchNumber.gte(len)) {
      return constants.AddressZero
    }
    return outbox.outboxes(batchNumber)
  }

  static waitUntilOutboxEntryCreated = async (
    batchNumber: BigNumber,
    activeOutboxAddress: string,
    l1Provider: providers.Provider,
    retryDelay = 500
  ): Promise<string> => {
    try {
      // if outbox entry not created yet, this reads from array out of bounds
      const expectedEntry = await BridgeHelper.getOutboxEntry(
        batchNumber,
        activeOutboxAddress,
        l1Provider
      )
      console.log('Found entry index!')
      return expectedEntry
    } catch (e) {
      console.log("can't find entry, lets wait a bit?")
      if (e.message === 'invalid opcode: opcode 0xfe not defined') {
        console.log('Array out of bounds, wait until the entry is posted')
      } else {
        console.log(e)
        console.log(e.message)
      }
      await BridgeHelper.wait(retryDelay)
      console.log('Starting new attempt')
      return BridgeHelper.waitUntilOutboxEntryCreated(
        batchNumber,
        activeOutboxAddress,
        l1Provider,
        retryDelay
      )
    }
  }

  static getActiveOutbox = async (
    l1CoreBridgeAddress: string,
    l1Provider: providers.Provider
  ) => {
    const bridge = await Bridge__factory.connect(
      l1CoreBridgeAddress,
      l1Provider
    )

    const [activeOutboxAddress] = await bridge.functions.allowedOutboxList(0)
    try {
      // index 1 should not exist
      await bridge.functions.allowedOutboxList(1)
      console.error('There is more than 1 outbox registered with the bridge?!')
    } catch (e) {
      // this should fail!
      console.log('All is good')
    }
    return activeOutboxAddress
  }

  static tryOutboxExecute = async (
    outboxProofData: OutboxProofData,
    l1CoreBridgeAddress: string,
    l1Signer: Signer
  ): Promise<ContractReceipt> => {
    if (!l1Signer.provider) throw new Error('No L1 provider in L1 signer')

    const activeOutboxAddress = await BridgeHelper.getActiveOutbox(
      l1CoreBridgeAddress,
      l1Signer.provider
    )

    await BridgeHelper.waitUntilOutboxEntryCreated(
      outboxProofData.batchNumber,
      activeOutboxAddress,
      l1Signer.provider
    )

    const outbox = Outbox__factory.connect(activeOutboxAddress, l1Signer)

    try {
      // TODO: wait until assertion is confirmed before execute
      // We can predict and print number of missing blocks
      // if not challenged
      const outboxExecute = await outbox.functions.executeTransaction(
        outboxProofData.batchNumber,
        outboxProofData.proof,
        outboxProofData.path,
        outboxProofData.l2Sender,
        outboxProofData.l1Dest,
        outboxProofData.l2Block,
        outboxProofData.l1Block,
        outboxProofData.timestamp,
        outboxProofData.amount,
        outboxProofData.calldataForL1
      )
      console.log(`Transaction hash: ${outboxExecute.hash}`)
      console.log('Waiting for receipt')
      const receipt = await outboxExecute.wait()
      console.log('Receipt emitted')
      return receipt
    } catch (e) {
      console.log('failed to execute tx in layer 1')
      console.log(e)
      // TODO: should we just try again after delay instead of throwing?
      throw e
    }
  }
  /**
   * Attempt to execute an outbox message; must be confirmed to succeed (i.e., confirmation delay must have passed)
   */
  static triggerL2ToL1Transaction = async (
    batchNumber: BigNumber,
    indexInBatch: BigNumber,
    l1CoreBridgeAddress: string,
    l2Provider: providers.Provider,
    l1Signer: Signer,
    singleAttempt = false
  ) => {
    if (!l1Signer.provider)
      throw new Error('Signer must be connected to L2 provider')

    console.log('going to get proof')
    let res: {
      proof: Array<string>
      path: BigNumber
      l2Sender: string
      l1Dest: string
      l2Block: BigNumber
      l1Block: BigNumber
      timestamp: BigNumber
      amount: BigNumber
      calldataForL1: string
    }

    if (singleAttempt) {
      const _res = await BridgeHelper.tryGetProofOnce(
        batchNumber,
        indexInBatch,
        l2Provider
      )
      if (_res === null) {
        throw new Error('Proof not found')
      }
      res = _res
    } else {
      res = await BridgeHelper.tryGetProof(
        batchNumber,
        indexInBatch,
        l2Provider
      )
    }

    const proofData: OutboxProofData = {
      ...res,
      batchNumber,
    }

    console.log('got proof')

    const outboxExecuteTransactionReceipt = await BridgeHelper.tryOutboxExecute(
      proofData,
      l1CoreBridgeAddress,
      l1Signer
    )
    return outboxExecuteTransactionReceipt
  }

  static getL2ToL1EventData = async (
    destinationAddress: string,
    l2Provider: providers.Provider
  ) => {
    const contract = ArbSys__factory.connect(ARB_SYS_ADDRESS, l2Provider)
    const iface = contract.interface
    const l2ToL1TransactionEvent = iface.getEvent('L2ToL1Transaction')
    const l2ToL1TransactionTopic = iface.getEventTopic(l2ToL1TransactionEvent)

    const topics = [
      l2ToL1TransactionTopic,
      ethers.utils.hexZeroPad(destinationAddress, 32),
    ]

    const logs = await l2Provider.getLogs({
      address: ARB_SYS_ADDRESS,
      topics,
      fromBlock: 0,
      toBlock: 'latest',
    })

    return logs.map(
      log => (iface.parseLog(log).args as unknown) as L2ToL1EventResult
    )
  }

  static assertionIsConfirmed = async (
    nodeNum: BigNumber,
    rollupAddress: string,
    l1Provider: providers.Provider
  ) => {
    const contract = Rollup__factory.connect(rollupAddress, l1Provider)
      .interface
    const iface = contract
    const nodeConfirmedEvent = iface.getEvent('NodeConfirmed')
    const nodeConfirmedEventTopic = iface.getEventTopic(nodeConfirmedEvent)

    const logs = await l1Provider.getLogs({
      address: rollupAddress,
      topics: [
        nodeConfirmedEventTopic,
        ethers.utils.hexZeroPad(nodeNum.toHexString(), 32),
      ],
      fromBlock: 0,
      toBlock: 'latest',
    })

    return logs.length === 1
  }

  static getNodeCreatedEvents = async (
    rollupAddress: string,
    l1Provider: providers.Provider
  ) => {
    const contract = Rollup__factory.connect(rollupAddress, l1Provider)
      .interface
    const iface = contract
    const nodeCreatedEvent = iface.getEvent('NodeCreated')
    const nodeCreatedEventTopic = iface.getEventTopic(nodeCreatedEvent)

    const logs = await l1Provider.getLogs({
      address: rollupAddress,
      topics: [nodeCreatedEventTopic],
      fromBlock: 0,
      toBlock: 'latest',
    })

    return logs
  }
  static getOutgoingMessage = async (
    batchNumber: BigNumber,
    indexInBatch: BigNumber,
    l2Provider: providers.Provider
  ) => {
    const contract = ArbSys__factory.connect(ARB_SYS_ADDRESS, l2Provider)
    const iface = contract.interface
    const l2ToL1TransactionEvent = iface.getEvent('L2ToL1Transaction')
    const l2ToL1TransactionTopic = iface.getEventTopic(l2ToL1TransactionEvent)

    const topics = [
      l2ToL1TransactionTopic,
      null,
      null,
      ethers.utils.hexZeroPad(batchNumber.toHexString(), 32),
    ]

    const logs = await l2Provider.getLogs({
      address: ARB_SYS_ADDRESS,
      // @ts-ignore
      topics,
      fromBlock: 0,
      toBlock: 'latest',
    })

    const parsedData = logs.map(
      log => (iface.parseLog(log).args as unknown) as L2ToL1EventResult
    )

    return parsedData.filter(log => log.indexInBatch.eq(indexInBatch))
  }

  static calculateOutgoingMessageId = (
    path: BigNumber,
    proofLength: BigNumber
  ) => {
    return utils.keccak256(
      utils.defaultAbiCoder.encode(['uint256', 'uint256'], [path, proofLength])
    )
  }

  static messageHasExecuted = async (
    outboxIndex: BigNumber,
    messageIndex: BigNumber,
    outboxAddress: string,
    l1Provider: providers.Provider
  ): Promise<boolean> => {
    const contract = Outbox__factory.connect(outboxAddress, l1Provider)
      .interface
    const iface = contract
    const executedEvent = iface.getEvent('OutBoxTransactionExecuted')
    const executedTopic = iface.getEventTopic(executedEvent)
    const logs = await l1Provider.getLogs({
      address: outboxAddress,

      topics: [
        executedTopic,
        // @ts-ignore
        null,
        // @ts-ignore
        null,
        ethers.utils.hexZeroPad(outboxIndex.toHexString(), 32),
      ],
      fromBlock: 0,
      toBlock: 'latest',
    })
    const parsedData = logs.map(
      log => (iface.parseLog(log).args as unknown) as OutBoxTransactionExecuted
    )
    return (
      parsedData.filter(executedEvent =>
        executedEvent.transactionIndex.eq(messageIndex)
      ).length === 1
    )
  }

  static getOutgoingMessageState = async (
    batchNumber: BigNumber,
    indexInBatch: BigNumber,
    outBoxAddress: string,
    l1Provider: providers.Provider,
    l2Provider: providers.Provider
  ): Promise<OutgoingMessageState> => {
    try {
      const proofData = await BridgeHelper.tryGetProofOnce(
        batchNumber,
        indexInBatch,
        l2Provider
      )

      if (!proofData) {
        return OutgoingMessageState.UNCONFIRMED
      }

      const messageExecuted = await BridgeHelper.messageHasExecuted(
        batchNumber,
        indexInBatch,
        outBoxAddress,
        l1Provider
      )
      if (messageExecuted) {
        return OutgoingMessageState.EXECUTED
      }

      const outboxEntry = await BridgeHelper.getOutboxEntry(
        batchNumber,
        outBoxAddress,
        l1Provider
      )

      if (outboxEntry === constants.AddressZero) {
        return OutgoingMessageState.UNCONFIRMED
      } else {
        return OutgoingMessageState.CONFIRMED
      }
    } catch (e) {
      console.warn('666: error in getOutgoingMessageState:', e)
      return OutgoingMessageState.NOT_FOUND
    }
  }
}
