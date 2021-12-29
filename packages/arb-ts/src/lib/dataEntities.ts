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

import { BigNumber } from '@ethersproject/bignumber'
import { FunctionFragment } from 'ethers/lib/utils'

// CHRIS: TODO: break this file up and put the interfaces with their relevant functionality

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

export interface MessageBatchProofInfo {
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

// export type MulticallFunctionInput = Array<{
//   target: string
//   funcFragment: FunctionFragment
//   values?: Array<unknown>
// }>

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
