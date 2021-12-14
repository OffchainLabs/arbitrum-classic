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

import { Filter } from '@ethersproject/abstract-provider'
import { ContractReceipt, PayableOverrides } from '@ethersproject/contracts'
import { Logger } from '@ethersproject/logger'
import { Zero } from '@ethersproject/constants'
import { parseEther } from '@ethersproject/units'
import { utils } from 'ethers'

import { NodeInterface__factory } from '../abi/factories/NodeInterface__factory'
import { L1ERC20Gateway__factory } from '../abi/factories/L1ERC20Gateway__factory'
import { L1WethGateway__factory } from '../abi/factories/L1WethGateway__factory'
import { Inbox__factory } from '../abi/factories/Inbox__factory'
import { Bridge__factory } from '../abi/factories/Bridge__factory'
import { OldOutbox__factory } from '../abi/factories/OldOutbox__factory'

import { Await, DepositParams, L1Bridge, L1TokenData } from '../l1Bridge'
import { L2Bridge, L2TokenData } from '../l2Bridge'
import {
  BridgeHelper,
  BuddyDeployEventResult,
  DepositInitiated,
  GatewaySet,
  L2ToL1EventResult,
  MessageBatchProofInfo,
  OutgoingMessageState,
  WithdrawalInitiated,
} from '../bridge_helpers'
import { NODE_INTERFACE_ADDRESS } from '../precompile_addresses'
import networks, { Network } from '../networks'
import {
  L1ERC20Gateway,
  L1GatewayRouter,
  Multicall2__factory,
  ArbMulticall2__factory,
  ERC20__factory,
  ERC20,
} from '../abi'
import { Result } from '@ethersproject/abi'
import { ArbRetryableTx__factory } from '../abi/factories/ArbRetryableTx__factory'
import { ARB_RETRYABLE_TX_ADDRESS } from '../precompile_addresses'
import { TransactionReceipt } from '@ethersproject/providers'
import { Provider } from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'
import { ContractTransaction } from '@ethersproject/contracts'
import { BigNumber } from '@ethersproject/bignumber'
import { constants } from 'ethers'
import { SignerOrProvider } from '../utils/signerOrProvider'
import {
  getMessageNumbersFromL1TxnReceipt,
  calculateRetryableTicketCreationHash,
  calculateL2MessageFromTicketTxnHash,
  L2TxnType,
} from './lib'
import { OutboxProofData } from '../bridge_helpers'

import { Interface, defaultAbiCoder } from '@ethersproject/abi'
import { Log } from '@ethersproject/abstract-provider'
import { concat, zeroPad, hexZeroPad } from '@ethersproject/bytes'

import { keccak256 } from '@ethersproject/keccak256'

import { L1GatewayRouter__factory } from '../abi/factories/L1GatewayRouter__factory'
import { Outbox__factory } from '../abi/factories/Outbox__factory'
import { IOutbox__factory } from '../abi/factories/IOutbox__factory'
import { ArbSys__factory } from '../abi/factories/ArbSys__factory'
import { Rollup__factory } from '../abi/factories/Rollup__factory'
import { L2ArbitrumGateway__factory } from '../abi/factories/L2ArbitrumGateway__factory'
import { Whitelist__factory } from '../abi/factories/Whitelist__factory'

import { ARB_SYS_ADDRESS } from '../precompile_addresses'
import { ArbMulticall2, Multicall2 } from '../abi'
import { FunctionFragment } from 'ethers/lib/utils'

class L2TransactionReceipt implements TransactionReceipt {
  public readonly to: string
  public readonly from: string
  public readonly contractAddress: string
  public readonly transactionIndex: number
  public readonly root?: string
  public readonly gasUsed: BigNumber
  public readonly logsBloom: string
  public readonly blockHash: string
  public readonly transactionHash: string
  public readonly logs: Array<Log>
  public readonly blockNumber: number
  public readonly confirmations: number
  public readonly cumulativeGasUsed: BigNumber
  public readonly effectiveGasPrice: BigNumber
  public readonly byzantium: boolean
  public readonly type: number
  public readonly status?: number

  constructor(tx: TransactionReceipt) {
    this.to = tx.to
    this.from = tx.from
    this.contractAddress = tx.contractAddress
    this.transactionIndex = tx.transactionIndex
    this.root = tx.root
    this.gasUsed = tx.gasUsed
    this.logsBloom = tx.logsBloom
    this.blockHash = tx.blockHash
    this.transactionHash = tx.transactionHash
    this.logs = tx.logs
    this.blockNumber = tx.blockNumber
    this.confirmations = tx.confirmations
    this.cumulativeGasUsed = tx.cumulativeGasUsed
    this.effectiveGasPrice = tx.effectiveGasPrice
    this.byzantium = tx.byzantium
    this.type = tx.type
    this.status = tx.status
  }

  public getL2ToL1Events() {
    const iface = ArbSys__factory.createInterface()
    const l2ToL1Event = iface.getEvent('L2ToL1Transaction')
    const eventTopic = iface.getEventTopic(l2ToL1Event)

    const logs = this.logs.filter(log => log.topics[0] === eventTopic)

    return logs.map(
      log => iface.parseLog(log).args as unknown as L2ToL1EventResult
    )
  }

  // CHRIS: why do we have both tryGetProofOnce and tryGetProof ? do we really expect it to fail so often?
  private async tryGetProofOnce(
    batchNumber: BigNumber,
    indexInBatch: BigNumber,
    l2Provider: Provider
  ): Promise<MessageBatchProofInfo | null> {
    const nodeInterface = NodeInterface__factory.connect(
      NODE_INTERFACE_ADDRESS,
      l2Provider
    )
    try {
      return nodeInterface.lookupMessageBatchProof(batchNumber, indexInBatch)
    } catch (e) {
      const expectedError = "batch doesn't exist"
      const err = e as any
      const actualError =
        err && (err.message || (err.error && err.error.message))
      if (actualError.includes(expectedError)) {
        console.log('Withdrawal detected, but batch not created yet.')
      } else {
        console.log("Withdrawal proof didn't work. Not sure why")
        console.log(e)
      }
    }
    return null
  }

  // CHRIS: where should this live, and what to do about that L1Bridge?
  private async getOutboxAddressByBatchNum(
    l1Bridge: L1Bridge,
    l1Provider: Provider,
    batchNum: BigNumber
  ): Promise<string> {
    const inbox = Inbox__factory.connect(
      (await l1Bridge.getInbox()).address,
      l1Provider
    )
    const bridge = await Bridge__factory.connect(
      await inbox.bridge(),
      l1Provider
    )
    const oldOutboxAddress = await bridge.allowedOutboxList(0)
    let newOutboxAddress: string
    try {
      newOutboxAddress = await bridge.allowedOutboxList(1)
    } catch {
      // new outbox not yet deployed; using old outbox
      return oldOutboxAddress
    }
    const oldOutbox = OldOutbox__factory.connect(oldOutboxAddress, l1Provider)
    const lastOldOutboxBatchNumber = await oldOutbox.outboxesLength()

    return batchNum.lt(lastOldOutboxBatchNumber)
      ? oldOutboxAddress
      : newOutboxAddress
  }

  public async getL2ToL1Messages(l2Provider: Provider, l1Signer: Signer) {
    const providerNetwork = await l1Signer.provider!.getNetwork()
    const arbNetwork = networks[providerNetwork.chainId]
    const l1Bridge = new L1Bridge(arbNetwork, l1Signer)
    const messages: L2ToL1MessageWriter[] = []

    for (const log of this.getL2ToL1Events()) {
      // CHRIS: this can fail for some, but not for others, what then?
      const proof = await this.tryGetProofOnce(
        log.batchNumber,
        log.indexInBatch,
        l2Provider
      )
      // CHRIS: do this properly
      if (!proof) throw Error()

      const outboxAddr = await this.getOutboxAddressByBatchNum(
        l1Bridge,
        // CHRIS: check provider and at the start of this func
        l1Signer.provider!,
        log.batchNumber
      )

      const l2ToL1Message = new L2ToL1MessageWriter(l1Signer, outboxAddr, {
        batchNumber: log.batchNumber,
        ...proof,
      })

      messages.push(l2ToL1Message)
    }

    return messages
  }
}

class L2ToL1MessageReader {
  constructor(
    protected readonly l1Provider: Provider,
    protected readonly outboxAddress: string,
    protected readonly outboxProofData: OutboxProofData
  ) {}

  private async outboxEntryExists() {
    const outbox = IOutbox__factory.connect(this.outboxAddress, this.l1Provider)
    return await outbox.outboxEntryExists(this.outboxProofData.batchNumber)
  }

  /**
   * Check if given outbox message has already been executed
   */
  private async hasExecuted(): Promise<boolean> {
    const outbox = Outbox__factory.connect(this.outboxAddress, this.l1Provider)
    try {
      await outbox.callStatic.executeTransaction(
        this.outboxProofData.batchNumber,
        this.outboxProofData.proof,
        this.outboxProofData.path,
        this.outboxProofData.l2Sender,
        this.outboxProofData.l1Dest,
        this.outboxProofData.l2Block,
        this.outboxProofData.l1Block,
        this.outboxProofData.timestamp,
        this.outboxProofData.amount,
        this.outboxProofData.calldataForL1
      )
      return false
    } catch (e: any) {
      if (e && e.message && e.message.toString().includes('ALREADY_SPENT')) {
        return true
      }
      if (e && e.message && e.message.toString().includes('NO_OUTBOX_ENTRY')) {
        return false
      }
      throw e
    }
  }

  public async status() {
    try {
      // CHRIS: this functionality still needs to exist somewhere
      // const proofData = await BridgeHelper.tryGetProofOnce(
      //   batchNumber,
      //   indexInBatch,
      //   l2Provider
      // )

      // if (!proofData) {
      //   return OutgoingMessageState.UNCONFIRMED
      // }

      const messageExecuted = await this.hasExecuted()
      if (messageExecuted) {
        return OutgoingMessageState.EXECUTED
      }

      const outboxEntryExists = await this.outboxEntryExists()

      return outboxEntryExists
        ? OutgoingMessageState.CONFIRMED
        : OutgoingMessageState.UNCONFIRMED
    } catch (e) {
      // CHRIS: this error needs updating. also 666?
      console.warn('666: error in getOutGoingMessageState:', e)
      return OutgoingMessageState.NOT_FOUND
    }
  }

  public async waitUntilOutboxEntryCreated(retryDelay = 500): Promise<void> {
    const exists = await this.outboxEntryExists()
    if (exists) {
      console.log('Found outbox entry!')
      return
    } else {
      console.log("can't find entry, lets wait a bit?")

      await BridgeHelper.wait(retryDelay)
      console.log('Starting new attempt')
      await this.waitUntilOutboxEntryCreated(retryDelay)
    }
  }
}

class L2ToL1MessageWriter extends L2ToL1MessageReader {
  constructor(
    private readonly l1Signer: Signer,
    outboxAddress: string,
    outboxProofData: OutboxProofData
  ) {
    // CHRIS: null check
    super(l1Signer.provider!, outboxAddress, outboxProofData)
  }

  /**
   * Executes the L2ToL1Message on L1.
   * @returns
   */
  public async execute() {
    await this.waitUntilOutboxEntryCreated()

    const outbox = Outbox__factory.connect(this.outboxAddress, this.l1Signer)
    try {
      // TODO: wait until assertion is confirmed before execute
      // We can predict and print number of missing blocks
      // if not challenged
      const outboxExecute = await outbox.functions.executeTransaction(
        this.outboxProofData.batchNumber,
        this.outboxProofData.proof,
        this.outboxProofData.path,
        this.outboxProofData.l2Sender,
        this.outboxProofData.l1Dest,
        this.outboxProofData.l2Block,
        this.outboxProofData.l1Block,
        this.outboxProofData.timestamp,
        this.outboxProofData.amount,
        this.outboxProofData.calldataForL1
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
}
