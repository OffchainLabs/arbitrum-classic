import { ContractTransaction, ethers } from 'ethers'
import { L2ERC20Gateway__factory } from './abi/factories/L2ERC20Gateway__factory'
import { L1ERC20Gateway__factory } from './abi/factories/L1ERC20Gateway__factory'
import { L1GatewayRouter__factory } from './abi/factories/L1GatewayRouter__factory'
import { Outbox__factory } from './abi/factories/Outbox__factory'
import { IOutbox__factory } from './abi/factories/IOutbox__factory'

import { Bridge__factory } from './abi/factories/Bridge__factory'
import { Inbox__factory } from './abi/factories/Inbox__factory'
import { ArbSys__factory } from './abi/factories/ArbSys__factory'
import { Rollup__factory } from './abi/factories/Rollup__factory'
import { L2ArbitrumGateway__factory } from './abi/factories/L2ArbitrumGateway__factory'

import { providers, utils, constants } from 'ethers'
import { BigNumber, Contract, Signer } from 'ethers'

import { NODE_INTERFACE_ADDRESS, ARB_SYS_ADDRESS } from './precompile_addresses'

import { Whitelist__factory } from './abi/factories/Whitelist__factory'

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

export interface WithdrawalInitiated {
  l1Token: string
  _from: string
  _to: string
  _l2ToL1Id: BigNumber
  _exitNum: BigNumber
  _amount: BigNumber
  txHash: string
}

export interface DepositInitiated {
  l1Token: string
  _from: string
  _to: string
  _sequenceNumber: BigNumber
  amount: BigNumber
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

export interface GatewaySet {
  l1Token: string
  gateway: string
}

export enum OutgoingMessageState {
  /**
   * No corresponding {@link L2ToL1EventResult} emitted
   */
  NOT_FOUND,
  /**
   * ArbSys.sendTxToL1 called, but assertion not yet confirmed
   */
  UNCONFIRMED,
  /**
   * Assertion for outgoing message confirmed, but message not yet executed
   */
  CONFIRMED,
  /**
   * Outgoing message executed (terminal state)
   */
  EXECUTED,
}

export type ChainIdOrProvider = BigNumber | providers.Provider

/**
 * Stateless helper methods; most wrapped / accessible (and documented) via {@link Bridge}
 */
export class BridgeHelper {
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

  private static _calculateRetryableHashInternal = async (
    inboxSequenceNumber: BigNumber,
    chainIdOrL2Provider: ChainIdOrProvider,
    txnType: 0 | 1
  ): Promise<string> => {
    const requestID = await BridgeHelper.calculateL2TransactionHash(
      inboxSequenceNumber,
      chainIdOrL2Provider
    )
    return utils.keccak256(
      utils.concat([
        utils.zeroPad(requestID, 32),
        utils.zeroPad(BigNumber.from(txnType).toHexString(), 32),
      ])
    )
  }

  static calculateL2RetryableTransactionHash = async (
    inboxSequenceNumber: BigNumber,
    chainIdOrL2Provider: ChainIdOrProvider
  ): Promise<string> => {
    return BridgeHelper._calculateRetryableHashInternal(
      inboxSequenceNumber,
      chainIdOrL2Provider,
      0
    )
  }

  static calculateRetryableAutoRedeemTxnHash = async (
    inboxSequenceNumber: BigNumber,
    chainIdOrL2Provider: ChainIdOrProvider
  ): Promise<string> => {
    return BridgeHelper._calculateRetryableHashInternal(
      inboxSequenceNumber,
      chainIdOrL2Provider,
      1
    )
  }

  static waitForRetryableReceipt = async (
    seqNum: BigNumber,
    l2Provider: providers.Provider
  ) => {
    const l2RetryableHash =
      await BridgeHelper.calculateL2RetryableTransactionHash(seqNum, l2Provider)
    return l2Provider.waitForTransaction(l2RetryableHash)
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
      log => iface.parseLog(log).args as unknown as BuddyDeployEventResult
    )
  }

  static getDepositTokenEventData = async (
    l1Transaction: providers.TransactionReceipt,
    l1GatewayAddress: string
  ): Promise<Array<DepositInitiated>> => {
    const factory = new L1ERC20Gateway__factory()
    const contract = factory.attach(l1GatewayAddress)
    const iface = contract.interface
    const event = iface.getEvent('DepositInitiated')
    const eventTopic = iface.getEventTopic(event)
    const logs = l1Transaction.logs.filter(log => log.topics[0] === eventTopic)
    return logs.map(
      log => iface.parseLog(log).args as unknown as DepositInitiated
    )
  }

  /**
   * All withdrawals from given token
   */
  static async getTokenWithdrawEventData(
    l2Provider: ethers.providers.Provider,
    gatewayAddress: string,
    l1TokenAddress: string,
    fromAddress?: string,
    filter?: providers.Filter
  ) {
    const gatewayContract = L2ArbitrumGateway__factory.connect(
      gatewayAddress,
      l2Provider
    )
    const topics = [
      null,
      fromAddress ? utils.hexZeroPad(fromAddress, 32) : null,
    ]
    const logs = await BridgeHelper.getEventLogs(
      'WithdrawalInitiated',
      gatewayContract,
      // @ts-ignore
      topics,
      filter
    )

    return logs
      .map(log => {
        const data = {
          ...gatewayContract.interface.parseLog(log).args,
          txHash: log.transactionHash,
        }
        return data as unknown as WithdrawalInitiated
      })
      .filter(
        (log: WithdrawalInitiated) =>
          log.l1Token.toLocaleLowerCase() === l1TokenAddress.toLocaleLowerCase()
      )
  }

  static async getGatewayWithdrawEventData(
    l2Provider: ethers.providers.Provider,
    gatewayAddress: string,
    fromAddress?: string,
    filter?: providers.Filter
  ) {
    const gatewayContract = L2ArbitrumGateway__factory.connect(
      gatewayAddress,
      l2Provider
    )
    const topics = [
      null,
      fromAddress ? utils.hexZeroPad(fromAddress, 32) : null,
    ]
    const logs = await BridgeHelper.getEventLogs(
      'WithdrawalInitiated',
      gatewayContract,
      // @ts-ignore
      topics,
      filter
    )

    return logs.map(log => {
      const data = {
        ...gatewayContract.interface.parseLog(log).args,
        txHash: log.transactionHash,
      }
      return data as unknown as WithdrawalInitiated
    })
  }

  public static getEventLogs = (
    eventName: string,
    connectedContract: Contract,
    topics: string | string[] = [],
    filter: ethers.providers.Filter = {}
  ) => {
    const iface = connectedContract.interface
    const event = iface.getEvent(eventName)
    const eventTopic = iface.getEventTopic(event)

    return connectedContract.provider.getLogs({
      address: connectedContract.address,
      topics: [eventTopic, ...topics],
      fromBlock: filter.fromBlock || 0,
      toBlock: filter.toBlock || 'latest',
    })
  }

  static getGatewaySetEventData = async (
    gatewayRouterAddress: string,
    provider: providers.Provider
  ) => {
    const contract = L1GatewayRouter__factory.connect(
      gatewayRouterAddress,
      provider
    )
    const logs = await BridgeHelper.getEventLogs('GatewaySet', contract)
    return logs.map(
      log => contract.interface.parseLog(log).args as unknown as GatewaySet
    )
  }

  static getWithdrawalsInL2Transaction = (
    l2Transaction: providers.TransactionReceipt,
    l2Provider: providers.Provider
  ): Array<L2ToL1EventResult> => {
    const contract = ArbSys__factory.connect(ARB_SYS_ADDRESS, l2Provider)
    const iface = contract.interface
    const l2ToL1Event = iface.getEvent('L2ToL1Transaction')
    const eventTopic = iface.getEventTopic(l2ToL1Event)

    const logs = l2Transaction.logs.filter(log => log.topics[0] === eventTopic)

    return logs.map(
      log => iface.parseLog(log).args as unknown as L2ToL1EventResult
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

  /**
   * Attempt to retrieve data necessary to execute outbox message; available before outbox entry is created /confirmed
   */
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
      const actualError = e && (e.message || (e.error && e.error.message))
      if (actualError.includes(expectedError)) {
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
      const actualError = e && (e.message || (e.error && e.error.message))
      if (actualError.includes(expectedError)) {
        console.log('Withdrawal detected, but batch not created yet.')
      } else {
        console.log("Withdrawal proof didn't work. Not sure why")
        console.log(e)
      }
    }
    return null
  }

  static outboxEntryExists = (
    batchNumber: BigNumber,
    outboxAddress: string,
    l1Provider: providers.Provider
  ): Promise<boolean> => {
    const outbox = IOutbox__factory.connect(outboxAddress, l1Provider)
    return outbox.outboxEntryExists(batchNumber)
  }

  static waitUntilOutboxEntryCreated = async (
    batchNumber: BigNumber,
    outboxAddress: string,
    l1Provider: providers.Provider,
    retryDelay = 500
  ) => {
    const exists = await BridgeHelper.outboxEntryExists(
      batchNumber,
      outboxAddress,
      l1Provider
    )
    if (exists) {
      console.log('Found outbox entry!')
      return
    } else {
      console.log("can't find entry, lets wait a bit?")

      await BridgeHelper.wait(retryDelay)
      console.log('Starting new attempt')
      await BridgeHelper.waitUntilOutboxEntryCreated(
        batchNumber,
        outboxAddress,
        l1Provider,
        retryDelay
      )
    }
  }

  static getActiveOutbox = async (
    rollupAddress: string,
    l1Provider: providers.Provider
  ) => {
    return Rollup__factory.connect(rollupAddress, l1Provider).outbox()
  }

  static tryOutboxExecute = async (
    outboxProofData: OutboxProofData,
    outboxAddress: string,
    l1Signer: Signer
  ): Promise<ContractTransaction> => {
    if (!l1Signer.provider) throw new Error('No L1 provider in L1 signer')
    await BridgeHelper.waitUntilOutboxEntryCreated(
      outboxProofData.batchNumber,
      outboxAddress,
      l1Signer.provider
    )

    const outbox = Outbox__factory.connect(outboxAddress, l1Signer)
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
      return outboxExecute
    } catch (e) {
      console.log('failed to execute tx in layer 1')
      console.log(e)
      // TODO: should we just try again after delay instead of throwing?
      throw e
    }
  }

  static triggerL2ToL1Transaction = async (
    batchNumber: BigNumber,
    indexInBatch: BigNumber,
    outboxAddress: string,
    l2Provider: providers.Provider,
    l1Signer: Signer,
    singleAttempt = false
  ) => {
    const l1Provider = l1Signer.provider
    if (!l1Provider) throw new Error('Signer must be connected to L2 provider')

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
      const outGoingMessageState = await BridgeHelper.getOutGoingMessageState(
        batchNumber,
        indexInBatch,
        outboxAddress,
        l1Provider,
        l2Provider
      )

      const infoString = `batchNumber: ${batchNumber.toNumber()} indexInBatch: ${indexInBatch.toNumber()}`

      switch (outGoingMessageState) {
        case OutgoingMessageState.NOT_FOUND:
          throw new Error(`Outgoing message not found. ${infoString}`)
        case OutgoingMessageState.UNCONFIRMED:
          throw new Error(
            `Attempting to execute message that isn't yet confirmed. ${infoString}`
          )
        case OutgoingMessageState.EXECUTED:
          throw new Error(`Message already executed ${infoString}`)
        case OutgoingMessageState.CONFIRMED: {
          const _res = await BridgeHelper.tryGetProofOnce(
            batchNumber,
            indexInBatch,
            l2Provider
          )
          if (_res === null)
            throw new Error(
              `666: message is in a confirmed node but lookupMessageBatchProof returned null (!) ${infoString}`
            )
          res = _res
          break
        }
      }
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

    return BridgeHelper.tryOutboxExecute(proofData, outboxAddress, l1Signer)
  }

  static getL2ToL1EventData = async (
    fromAddress: string,
    l2Provider: providers.Provider,
    filter?: providers.Filter
  ) => {
    const contract = ArbSys__factory.connect(ARB_SYS_ADDRESS, l2Provider)

    const logs = await BridgeHelper.getEventLogs(
      'L2ToL1Transaction',
      contract,
      [ethers.utils.hexZeroPad(fromAddress, 32)],
      filter
    )

    return logs.map(
      log =>
        contract.interface.parseLog(log).args as unknown as L2ToL1EventResult
    )
  }

  /**
   * Check if given assertion has been confirmed
   */
  static assertionIsConfirmed = async (
    nodeNum: BigNumber,
    rollupAddress: string,
    l1Provider: providers.Provider
  ) => {
    const contract = Rollup__factory.connect(rollupAddress, l1Provider)
    const logs = await BridgeHelper.getEventLogs('NodeConfirmed', contract, [
      ethers.utils.hexZeroPad(nodeNum.toHexString(), 32),
    ])
    return logs.length === 1
  }

  static getNodeCreatedEvents = (
    rollupAddress: string,
    l1Provider: providers.Provider
  ) => {
    const contract = Rollup__factory.connect(rollupAddress, l1Provider)
    return BridgeHelper.getEventLogs('NodeCreated', contract)
  }

  static getOutgoingMessage = async (
    batchNumber: BigNumber,
    indexInBatch: BigNumber,
    l2Provider: providers.Provider
  ) => {
    const contract = ArbSys__factory.connect(ARB_SYS_ADDRESS, l2Provider)

    const topics = [
      null,
      null,
      ethers.utils.hexZeroPad(batchNumber.toHexString(), 32),
    ]

    const logs = await BridgeHelper.getEventLogs(
      'L2ToL1Transaction',
      contract,
      // @ts-ignore
      topics
    )

    const parsedData = logs.map(
      log =>
        contract.interface.parseLog(log).args as unknown as L2ToL1EventResult
    )

    return parsedData.filter(log => log.indexInBatch.eq(indexInBatch))
  }

  /**
   * Get outgoing message Id (key to in OutboxEntry.spentOutput)
   */
  static calculateOutgoingMessageId = (
    path: BigNumber,
    proofLength: BigNumber
  ) => {
    return utils.keccak256(
      utils.defaultAbiCoder.encode(['uint256', 'uint256'], [path, proofLength])
    )
  }
  /**
   * Check if given outbox message has already been executed
   */
  static messageHasExecuted = async (
    batchNumber: BigNumber,
    path: BigNumber,
    outboxAddress: string,
    l1Provider: providers.Provider
  ): Promise<boolean> => {
    const contract = Outbox__factory.connect(outboxAddress, l1Provider)
    const topics = [
      null,
      null,
      ethers.utils.hexZeroPad(batchNumber.toHexString(), 32),
    ]
    const logs = await BridgeHelper.getEventLogs(
      'OutBoxTransactionExecuted',
      contract,
      // @ts-ignore
      topics
    )
    const parsedData = logs.map(
      log =>
        contract.interface.parseLog(log)
          .args as unknown as OutBoxTransactionExecuted
    )
    return (
      parsedData.filter(executedEvent =>
        executedEvent.transactionIndex.eq(path)
      ).length === 1
    )
  }

  static getOutGoingMessageState = async (
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
        proofData.path,
        outBoxAddress,
        l1Provider
      )
      if (messageExecuted) {
        return OutgoingMessageState.EXECUTED
      }

      const outboxEntryExists = await BridgeHelper.outboxEntryExists(
        batchNumber,
        outBoxAddress,
        l1Provider
      )

      return outboxEntryExists
        ? OutgoingMessageState.CONFIRMED
        : OutgoingMessageState.UNCONFIRMED
    } catch (e) {
      console.warn('666: error in getOutGoingMessageState:', e)
      return OutgoingMessageState.NOT_FOUND
    }
  }
  static isWhiteListed(
    address: string,
    whiteListAddress: string,
    l1Provider: providers.Provider
  ) {
    const whiteList = Whitelist__factory.connect(whiteListAddress, l1Provider)
    return whiteList.isAllowed(address)
  }

  static percentIncrease(num: BigNumber, increase: BigNumber) {
    return num.add(num.mul(increase).div(100))
  }
}
