import { ArbTsError } from '../dataEntities/errors'
import { JsonRpcProvider, Formatter } from '@ethersproject/providers'
import {
  ArbTransactionReceipt,
  ArbTransactionResponse,
  BatchInfo,
  FeeStat,
  FeeStats,
  ReturnCode,
} from '../dataEntities/arbTransaction'
import { providers, BigNumber, logger } from 'ethers'
import { Formats, FormatFuncs } from '@ethersproject/providers/lib/formatter'
import { getL1Network, getL2Network } from '../..'
import { SignerProviderUtils } from '../dataEntities/signerOrProvider'
import { SequencerInbox, SequencerInbox__factory } from '../abi'
import {
  DelayedInboxForcedEvent,
  SequencerBatchDeliveredEvent,
  SequencerBatchDeliveredFromOriginEvent,
} from '../abi/ISequencerInbox'
import { EventFetcher } from './eventFetcher'

type ArbFormats = Formats & {
  feeStats: FormatFuncs
  feeStat: FormatFuncs
  batchInfo: FormatFuncs
}

class ArbFormatter extends Formatter {
  readonly formats!: ArbFormats

  public getDefaultFormats(): ArbFormats {
    // formats was already initialised in super, so we can just access here
    const superFormats = super.getDefaultFormats()

    const address = this.address.bind(this)
    const bigNumber = this.bigNumber.bind(this)
    const data = this.data.bind(this)
    const hash = this.hash.bind(this)
    const number = this.number.bind(this)
    const feeStats = this.feeStats.bind(this)
    const feeStat = this.feeStat.bind(this)
    const batchInfo = this.batchInfo.bind(this)
    const returnCode = this.returnCode.bind(this)

    const arbTransactionFormat = {
      ...superFormats.transaction,

      l1SequenceNumber: bigNumber,
      parentRequestId: hash,
      indexInParent: number,
      arbType: number,
      arbSubType: number,
      l1BlockNumber: number,
    }

    const arbReceiptFormat = {
      ...superFormats.receipt,
      returnData: data,
      returnCode: returnCode,
      feeStats: feeStats,
      batchInfo: Formatter.allowNull(batchInfo, null),
      l1BlockNumber: number,
    }

    const feeStatsFormat = {
      prices: feeStat,
      unitsUsed: feeStat,
      paid: feeStat,
    }

    const feeStatFormat = {
      l1Transaction: bigNumber,
      l1Calldata: bigNumber,
      l2Storage: bigNumber,
      l2Computation: bigNumber,
    }

    const batchInfoFormat = {
      confirmations: number,
      blockNumber: number,
      logAddress: address,
      logTopics: Formatter.arrayOf(hash),
      logData: data,
    }

    return {
      ...superFormats,
      transaction: arbTransactionFormat,
      receipt: arbReceiptFormat,
      feeStats: feeStatsFormat,
      feeStat: feeStatFormat,
      batchInfo: batchInfoFormat,
    }
  }

  public returnCode(value: any): ReturnCode {
    const bn = BigNumber.from(value)
    const returnNum = bn.toNumber()
    if (!Object.values(ReturnCode).includes(returnNum)) {
      return logger.throwArgumentError('invalid return code', 'value', value)
    }
    return returnNum
  }

  public feeStat(feeStat: any): FeeStat {
    return Formatter.check(this.formats.feeStat, feeStat)
  }

  public feeStats(feeStats: any): FeeStats {
    return Formatter.check(this.formats.feeStats, feeStats)
  }

  public batchInfo(batchInfo: any): BatchInfo {
    return Formatter.check(this.formats.batchInfo, batchInfo)
  }

  public transactionResponse(transaction: any): ArbTransactionResponse {
    return super.transactionResponse(transaction) as ArbTransactionResponse
  }

  public receipt(value: any): ArbTransactionReceipt {
    return super.receipt(value) as ArbTransactionReceipt
  }
}

const getBatch = async (
  seqNum: BigNumber,
  l1BlockNumber: number,
  l1Provider: providers.Provider
): Promise<Omit<BatchInfo, 'confirmations'> | null> => {
  const ethNetwork = await getL1Network(l1Provider)
  const arbNetwork = await getL2Network(ethNetwork.partnerChainIDs[0])
  const inbox = SequencerInbox__factory.connect(
    arbNetwork.ethBridge.sequencerInbox,
    l1Provider
  )

  const delta = 1000

  const delayBlocks = (await inbox.maxDelayBlocks()).toNumber()

  const startBlock = l1BlockNumber - delta
  const delayedBlockMax = l1BlockNumber + delayBlocks
  const currentBlock = await l1Provider.getBlockNumber()

  const endBlock =
    delayedBlockMax > currentBlock ? currentBlock : delayedBlockMax

  const batchEvents = new EventFetcher(l1Provider)
  const events = await batchEvents.getEvents<
    SequencerInbox,
    | DelayedInboxForcedEvent
    | SequencerBatchDeliveredEvent
    | SequencerBatchDeliveredFromOriginEvent
  >(
    inbox.address,
    SequencerInbox__factory,
    c => {
      const eventTopics = [
        c.interface.getEventTopic(c.interface.getEvent('DelayedInboxForced')),
        c.interface.getEventTopic(
          c.interface.getEvent('SequencerBatchDelivered')
        ),
        c.interface.getEventTopic(
          c.interface.getEvent('SequencerBatchDeliveredFromOrigin')
        ),
      ]

      return { topics: [eventTopics] }
    },
    { fromBlock: startBlock, toBlock: endBlock }
  )

  // find the batch containing the seq number
  const batch = events.filter(
    b => b.event.firstMessageNum <= seqNum && b.event.newMessageCount > seqNum
  )[0]

  if (!batch) return null

  return {
    blockNumber: batch.blockNumber,
    logAddress: batch.address,
    logData: batch.data,
    logTopics: batch.topics,
  }
}

/**
 * Fetch a transaction receipt from an l2Provider
 * If an l1Provider is also provided then info about the l1 data
 * availability of the transaction will also be returned in the l1InboxBatchInfo
 * field
 * @param l2Provider
 * @param txHash
 * @param l1ProviderForBatch
 * @returns
 */
export const getRawArbTransactionReceipt = async (
  l2Provider: JsonRpcProvider,
  txHash: string,
  l1ProviderForBatch?: JsonRpcProvider
): Promise<ArbTransactionReceipt | null> => {
  const rec = await l2Provider.send('eth_getTransactionReceipt', [txHash])
  if (rec == null) return null
  const arbFormatter = new ArbFormatter()
  const arbTxReceipt = arbFormatter.receipt(rec)

  // if we haven't already got batch info, and it has been requested
  // then we fetch it and append it
  if (!arbTxReceipt.l1InboxBatchInfo && l1ProviderForBatch) {
    const l2Network = await getL2Network(l2Provider)
    const l1Network = await getL1Network(l2Network.partnerChainID)
    SignerProviderUtils.checkNetworkMatches(
      l1ProviderForBatch,
      parseInt(l1Network.chainID)
    )

    const tx = await getTransaction(l2Provider, txHash)
    if (tx) {
      const batch = await getBatch(
        tx.l1SequenceNumber,
        tx.l1BlockNumber,
        l1ProviderForBatch
      )
      const currentBlock = await l1ProviderForBatch.getBlockNumber()

      arbTxReceipt.l1InboxBatchInfo = batch
        ? { ...batch, confirmations: currentBlock - batch.blockNumber }
        : null
    }
  }

  return arbTxReceipt
}

export const getTransaction = async (
  l2Provider: JsonRpcProvider,
  txHash: string
): Promise<ArbTransactionResponse | null> => {
  const tx = await l2Provider.send('eth_getTransactionByHash', [txHash])
  if (tx === null) return null
  const arbFormatter = new ArbFormatter()
  return arbFormatter.transactionResponse(tx)
}
