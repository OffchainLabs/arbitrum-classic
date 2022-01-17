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
import { L1TransactionReceipt } from '../src/lib/message/L1ToL2Message'

if (!args.txid) {
  throw new Error('Include txid (--txid 0xmytxid)')
}

const l1Txn: string | ContractReceipt = args.txid as string

;(async () => {
  const { l1Signer, l2Signer } = await instantiateBridge()
  const l1Receipt = new L1TransactionReceipt(
    await l1Signer.provider!.getTransactionReceipt(l1Txn)
  )
  const message = await l1Receipt.getL1ToL2Message(l2Signer)
  const res = await message.cancel()
  const rec = await res.wait()
  console.log('done:', rec)
  console.log(rec.status === 1 ? 'success!' : 'failed...')
})()
