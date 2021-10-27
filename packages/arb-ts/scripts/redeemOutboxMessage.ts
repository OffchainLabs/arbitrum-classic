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

import { ContractReceipt } from '@ethersproject/contracts'

import { instantiateBridge } from './instantiate_bridge'
import args from './getCLargs'
import { OutgoingMessageState } from '../src/lib/bridge_helpers'
import prompts from 'prompts'

if (!args.txid) {
  throw new Error('Include txid (--txid 0xmytxid)')
}

const l2Txn: string | ContractReceipt = args.txid as string

if (!l2Txn) {
  throw new Error('Need to set l1 txn hash')
}

;(async () => {
  const { bridge } = await instantiateBridge()
  const { outGoingMessageState, batchNumber, indexInBatch } =
    await bridge.getOutGoingMessageDataFromL2Transaction(l2Txn)
  switch (outGoingMessageState) {
    case OutgoingMessageState.NOT_FOUND:
      return console.log('Message not found')

    case OutgoingMessageState.EXECUTED:
      return console.log('Message already executed')
    case OutgoingMessageState.UNCONFIRMED:
      return console.log('Message not yet confirmed')
    case OutgoingMessageState.CONFIRMED: {
      console.log('Message is confirmed! Would you like to try to execute it?')
      const res = await prompts({
        type: 'confirm',
        name: 'value',
        message: 'Message is confirmed! Would you like to try to execute it?',
        initial: true,
      })
      if (!res.value) {
        return console.log('Okay, exiting')
      } else {
        console.log('Trying to redeem on L1...')
        const res = await bridge.triggerL2ToL1Transaction(
          batchNumber,
          indexInBatch,
          true
        )
        const rec = await res.wait(2)
        console.log('Success!', rec.transactionHash)
        return
      }
    }
    default:
      break
  }
})()
