/*
 * Copyright 2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/* eslint-env node */
'use strict'

import { TransactionReceipt } from '@ethersproject/providers'
import { BigNumber } from '@ethersproject/bignumber'
import { Inbox__factory } from '../abi/factories/Inbox__factory'
import { keccak256 } from '@ethersproject/keccak256'
import { concat, zeroPad, hexZeroPad } from '@ethersproject/bytes'

export enum L2TxnType {
  USER_TXN = 0,
  AUTO_REDEEM = 1,
}

export interface MessageId {
  messageNumber: BigNumber
  l2ChainId: BigNumber
}

const bitFlip = (num: BigNumber): BigNumber => {
  return num.or(BigNumber.from(1).shl(255))
}

export const calculateL2TxnHash = (messageId: MessageId): string => {
  return keccak256(
    concat([
      zeroPad(messageId.l2ChainId.toHexString(), 32),
      zeroPad(bitFlip(messageId.messageNumber).toHexString(), 32),
    ])
  )
}

export const calculateRetryableTicketCreationHash = (
  messageId: MessageId
): string => {
  return calculateL2TxnHash(messageId)
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
  messageId: MessageId
): string => {
  const ticketCreationHash = calculateL2TxnHash(messageId)
  return calculateL2MessageFromTicketTxnHash(
    ticketCreationHash,
    L2TxnType.AUTO_REDEEM
  )
}

export const calculateRetryableUserTxnHash = (messageId: MessageId): string => {
  const ticketCreationHash = calculateL2TxnHash(messageId)
  return calculateL2MessageFromTicketTxnHash(
    ticketCreationHash,
    L2TxnType.USER_TXN
  )
}