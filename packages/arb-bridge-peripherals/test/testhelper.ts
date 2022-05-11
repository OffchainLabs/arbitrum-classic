/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

/* eslint-env node, mocha */
import { ethers, network } from 'hardhat'
import { ContractTransaction } from 'ethers'
import {
  InboxMock,
  L1ArbitrumMessenger__factory,
  L2ArbitrumMessenger__factory,
} from '../build/types'
import { TxToL1Event } from '../build/types/L2ArbitrumMessenger'

export const processL1ToL2Tx = async (
  tx: Promise<ContractTransaction> | ContractTransaction
) => {
  const receipt = await (await tx).wait()
  const iface = L1ArbitrumMessenger__factory.createInterface()
  const logs = receipt.logs.filter(
    log => log.topics[0] === iface.getEventTopic('TxToL2')
  )
  if (logs.length === 0) throw new Error('No L1 to L2 txs')
  const l1ToL2Logs = logs.map(log => {
    const event = iface.parseLog(log)
    const to = event.args!._to
    const data = event.args!._data
    const from = log.address
    const fromAliased =
      '0x' +
      BigInt.asUintN(
        160,
        BigInt(from) + BigInt('0x1111000000000000000000000000000000001111')
      )
        .toString(16)
        .padStart(40, '0')
    return network.provider
      .request({
        method: 'hardhat_setBalance',
        params: [fromAliased, '0xffffffffffffffffffff'],
      })
      .then(() =>
        network.provider.request({
          method: 'hardhat_impersonateAccount',
          params: [fromAliased],
        })
      )
      .then(() => ethers.getSigner(fromAliased))
      .then(signer =>
        signer.sendTransaction({
          to: to,
          data: data,
        })
      )
  })
  return Promise.all(l1ToL2Logs)
}

export const processL2ToL1Tx = async (
  tx: Promise<ContractTransaction> | ContractTransaction,
  inboxMock: InboxMock
) => {
  const receipt = await (await tx).wait()
  const iface = L2ArbitrumMessenger__factory.createInterface()
  const logs = receipt.logs.filter(
    log => log.topics[0] === iface.getEventTopic('TxToL1')
  )
  if (logs.length === 0) throw new Error('No L2 to L1 txs')
  const l2ToL1Logs = logs.map(log => {
    const event = iface.parseLog(log)
    const to = event.args._to
    const data = event.args._data
    const from = log.address
    return network.provider
      .request({
        method: 'hardhat_setBalance',
        params: [inboxMock.address, '0xffffffffffffffffffff'],
      })
      .then(() =>
        // Also fund to address (which can be wethgateway)
        network.provider.request({
          method: 'hardhat_setBalance',
          params: [to, '0xffffffffffffffffffff'],
        })
      )
      .then(() =>
        network.provider.request({
          method: 'hardhat_impersonateAccount',
          params: [inboxMock.address],
        })
      )
      .then(() => inboxMock.setL2ToL1Sender(from, { gasLimit: 5000000 }))
      .then(() => ethers.getSigner(inboxMock.address))
      .then(signer =>
        signer.sendTransaction({
          to: to,
          data: data,
        })
      )
  })
  return Promise.all(l2ToL1Logs)
}
