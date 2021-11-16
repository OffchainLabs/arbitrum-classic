import { TransactionReceipt } from '@ethersproject/providers'
import { BigNumber } from '@ethersproject/bignumber'
import { Inbox__factory } from '../abi/factories/Inbox__factory'
import { keccak256 } from '@ethersproject/keccak256'
import { concat, zeroPad, hexZeroPad } from '@ethersproject/bytes'

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

const calculateL2MessageHashHelper = (
  messageNumber: BigNumber,
  l2ChainID: BigNumber,
  txnType: 0 | 1
): string => {
  const requestID = calculateL2TxnHash(messageNumber, l2ChainID)
  return keccak256(
    concat([
      zeroPad(requestID, 32),
      zeroPad(BigNumber.from(txnType).toHexString(), 32),
    ])
  )
}

export const calculateRetryableAutoRedeemTxnHash = (
  messageNumber: BigNumber,
  l2ChainID: BigNumber
): string => {
  return calculateL2MessageHashHelper(messageNumber, l2ChainID, 1)
}

export const calculateRetryableUserTxnHash = (
  messageNumber: BigNumber,
  l2ChainID: BigNumber
): string => {
  return calculateL2MessageHashHelper(messageNumber, l2ChainID, 0)
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
