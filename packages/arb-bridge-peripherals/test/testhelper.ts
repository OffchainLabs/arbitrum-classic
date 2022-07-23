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
import { ContractTransaction, Wallet } from 'ethers'
import {
  InboxMock,
  InboxMock__factory,
  ArbSysMock__factory,
} from '../build/types'

export const processL1ToL2Tx = async (
  tx: Promise<ContractTransaction> | ContractTransaction
) => {
  const receipt = await (await tx).wait()
  const iface = InboxMock__factory.createInterface()
  const logs = receipt.logs.filter(
    log => log.topics[0] === iface.getEventTopic('InboxRetryableTicket')
  )
  if (logs.length === 0) throw new Error('No L1 to L2 txs')
  const l1ToL2Logs = logs.map(log => {
    const event = iface.parseLog(log)
    const to = event.args.to
    const data = event.args.data
    const from = event.args.from
    const value = event.args.value
    const fromAliased = applyAlias(from)

    return network.provider
      .request({
        // Fund fromAliased to send transaction
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
          value: value,
        })
      )
  })
  return Promise.all(l1ToL2Logs)
}

export const impersonateAccount = (address: string) =>
  network.provider
    .request({
      // Fund inboxMock to send transaction
      method: 'hardhat_setBalance',
      params: [address, '0xffffffffffffffffffff'],
    })
    .then(() =>
      network.provider.request({
        method: 'hardhat_impersonateAccount',
        params: [address],
      })
    )
    .then(() => ethers.getSigner(address))

export const applyAlias = (address: string) =>
  '0x' +
  BigInt.asUintN(
    160,
    BigInt(address) + BigInt('0x1111000000000000000000000000000000001111')
  )
    .toString(16)
    .padStart(40, '0')

export const processL2ToL1Tx = async (
  tx: Promise<ContractTransaction> | ContractTransaction,
  inboxMock: InboxMock
) => {
  const receipt = await (await tx).wait()
  const iface = ArbSysMock__factory.createInterface()
  const logs = receipt.logs.filter(
    log => log.topics[0] === iface.getEventTopic('ArbSysL2ToL1Tx')
  )
  if (logs.length === 0) throw new Error('No L2 to L1 txs')
  const l2ToL1Logs = logs.map(log => {
    const event = iface.parseLog(log)
    const to = event.args.to
    const data = event.args.data
    const from = event.args.from
    const value = event.args.value
    return inboxMock
      .setL2ToL1Sender(from, { gasLimit: 5000000 })
      .then(() => impersonateAccount(inboxMock.address))
      .then(signer =>
        signer.sendTransaction({
          to: to,
          data: data,
          value: value,
        })
      )
  })
  return Promise.all(l2ToL1Logs)
}

export async function getCorrectPermitSig(
  wallet: Wallet,
  token: any,
  spender: string,
  value: any,
  deadline: any,
  optional?: { nonce?: number; name?: string; chainId?: number; version?: string }
  ) { 
  const [nonce, name, version, chainId] = await Promise.all([
      optional?.nonce ?? token.nonces(wallet.address),
      optional?.name ?? token.name(),
      optional?.version ?? '1',
      optional?.chainId ?? network.config.chainId,
  ])
  
  const domain = {
      "name": name,
      "version": version,
      "chainId": chainId,
      "verifyingContract": token.address
  };
  
  const types = {
      Permit: [
      { name: 'owner', type: 'address' },
      { name: 'spender', type: 'address' },
      { name: 'value', type: 'uint256' },
      { name: 'nonce', type: 'uint256'},
      { name: 'deadline', type: 'uint256' },
  ],
  }
  
  const message = {
          owner: wallet.address,
          spender: spender,
          value: value,
          nonce: nonce,
          deadline: deadline
  };
  
  const sig = await wallet._signTypedData(domain, types, message);
  return sig;
}
