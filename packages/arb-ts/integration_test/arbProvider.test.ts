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

import { expect } from 'chai'

import { instantiateBridgeWithRandomWallet, skipIfMainnet } from './testHelpers'
import { getRawArbTransactionReceipt } from '../src'
import { JsonRpcProvider } from '@ethersproject/providers'

describe('ArbProvider', () => {
  beforeEach('skipIfMainnet', function () {
    skipIfMainnet(this)
  })

  it('does find l1 batch info', async () => {
    const { l2Signer, l1Signer } = await instantiateBridgeWithRandomWallet()
    const l1Provider = l1Signer.provider! as JsonRpcProvider
    const l2Provider = l2Signer.provider! as JsonRpcProvider

    const testTxHash =
      '0x4ab7b21ebb243a4eec8a5e08f9190e0c1c116363b165f15cc60aa9c247e530ca'

    const arbReceipt = await getRawArbTransactionReceipt(
      l2Provider,
      testTxHash,
      l1Provider
    )

    const expected = {
      blockNumber: 10052061,
      logAddress: '0xE1Ae39E91C5505f7F0ffC9e2bbF1f6E1122DCfA8',
      logData:
        '0x000000000000000000000000000000000000000000000000000000000114a7ff31b2dee865892a4558d530c5540bb19b05f9dc9ae6d4e677228b2d9f445094f0000000000000000000000000000000000000000000000000000000000000b33f',
      logTopics: [
        '0x10e0571aafaf282151fd5b0215b5495521c549509cb0de3a3f8310bd2e344682',
        '0x000000000000000000000000000000000000000000000000000000000114a63b',
        '0x0206cda4867cf7030ff225a3d8573e66cd0a9b40541ee63b6c1f0fc4c1e58d4c',
      ],
      confirmations: 238,
    }
    expect(
      arbReceipt?.l1InboxBatchInfo?.blockNumber,
      'Invalid block number'
    ).to.eq(expected.blockNumber)
    expect(arbReceipt?.l1InboxBatchInfo?.logAddress, 'Invalid address').to.eq(
      expected.logAddress
    )
    expect(arbReceipt?.l1InboxBatchInfo?.logData, 'Invalid data').to.eq(
      expected.logData
    )
    expect(
      arbReceipt?.l1InboxBatchInfo?.logTopics,
      'Invalid topics'
    ).to.deep.eq(expected.logTopics)
    expect(
      arbReceipt?.l1InboxBatchInfo?.confirmations,
      'Invalid confirmations'
    ).to.be.gt(expected.confirmations)
  })
})
