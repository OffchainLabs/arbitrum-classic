import { TransactionReceipt } from '@ethersproject/providers'
import { BigNumber } from '@ethersproject/bignumber'
import { Inbox__factory } from '../abi/factories/Inbox__factory'
import { keccak256 } from '@ethersproject/keccak256'
import { concat, zeroPad, hexZeroPad } from '@ethersproject/bytes'

export enum L2TxnType {
  USER_TXN = 0,
  AUTO_REDEEM = 1,
}

const bitFlip = (num: BigNumber): BigNumber => {
  return num.or(BigNumber.from(1).shl(255))
}

export const calculateL2TxnHash = (
  messageNumber: BigNumber,
  l2ChainId: BigNumber
): string => {
  return keccak256(
    concat([
      zeroPad(l2ChainId.toHexString(), 32),
      zeroPad(bitFlip(messageNumber).toHexString(), 32),
    ])
  )
}

export const calculateRetryableTicketCreationHash = (
  messageNumber: BigNumber,
  l2ChainId: BigNumber
): string => {
  return calculateL2TxnHash(messageNumber, l2ChainId)
}

export const calculateL2MessageFromTicketTxnHash = (
  ticketCreationHash: string,
  l2TxnType: L2TxnType
): string => {
  return keccak256(
    concat([
      zeroPad(ticketCreationHash, 32),
      zeroPad(BigNumber.from(l2TxnType).toHexString(), 32),
    ])
  )
}

export const calculateRetryableAutoRedeemTxnHash = (
  messageNumber: BigNumber,
  l2ChainID: BigNumber
): string => {
  const ticketCreationHash = calculateL2TxnHash(messageNumber, l2ChainID)
  return calculateL2MessageFromTicketTxnHash(
    ticketCreationHash,
    L2TxnType.AUTO_REDEEM
  )
}

export const calculateRetryableUserTxnHash = (
  messageNumber: BigNumber,
  l2ChainID: BigNumber
): string => {
  const ticketCreationHash = calculateL2TxnHash(messageNumber, l2ChainID)
  return calculateL2MessageFromTicketTxnHash(
    ticketCreationHash,
    L2TxnType.USER_TXN
  )
}

export const getMessageNumbers = (
  l1Transaction: TransactionReceipt
): BigNumber[] | undefined => {
  const iface = Inbox__factory.createInterface()
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
