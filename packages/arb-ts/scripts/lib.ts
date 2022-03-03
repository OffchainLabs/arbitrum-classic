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

import { ERC20__factory } from '../src/lib/abi/factories/ERC20__factory'
import { L1ToL2MessageStatus } from '../src/lib/message/L1ToL2Message'
import { L1TransactionReceipt } from '../src/lib/message/L1Transaction'
import { instantiateBridge } from './instantiate_bridge'

export const setStandardGateWays = async (
  tokens: string[]
): Promise<ContractReceipt> => {
  return setGateWays(tokens, 'standard')
}
export const setArbCustomGateways = async (
  tokens: string[]
): Promise<ContractReceipt> => {
  return setGateWays(tokens, 'arbCustom')
}

export const setGateWays = async (
  tokens: string[],
  type: 'standard' | 'arbCustom',
  overrideGateways: string[] = []
): Promise<ContractReceipt> => {
  const { adminErc20Bridger, l1Signer, l2Network, l2Signer } =
    await instantiateBridge()
  const l1Provider = l1Signer.provider!
  const l2Provider = l2Signer.provider!
  if (tokens.length === 0) {
    throw new Error('Include some tokens to set')
  }
  if (
    overrideGateways.length > 0 &&
    overrideGateways.length !== tokens.length
  ) {
    throw new Error('Token/Gateway arrays are different lengths')
  }
  console.log('Setting', tokens.length, 'tokens')

  for (const tokenAddress of tokens) {
    try {
      const token = await ERC20__factory.connect(tokenAddress, l1Provider)
      console.warn('calling name for ', tokenAddress)

      const symbol = await token.symbol()
      const name = await token.name()
      const decimals = await token.decimals()
      console.log(symbol, name, decimals)

      const looksGood =
        typeof symbol === 'string' &&
        typeof decimals === 'number' &&
        typeof name === 'string' &&
        decimals > 0 &&
        symbol.length > 0 &&
        name.length > 0

      if (!looksGood) {
        throw new Error(`${tokenAddress} doesn't look like an L1 erc20...`)
      }
    } catch (err) {
      console.warn('err', err)

      throw new Error(`${tokenAddress} doesn't look like an L1 erc20...`)
    }
  }
  console.log('L1 sanity checks passed...')
  const gateways = (() => {
    if (overrideGateways.length > 0) {
      return overrideGateways
    } else if (type === 'standard') {
      return tokens.map(() => l2Network.tokenBridge.l1ERC20Gateway)
    } else if (type === 'arbCustom') {
      return tokens.map(() => l2Network.tokenBridge.l1CustomGateway)
    } else {
      throw new Error('Unhandled else case')
    }
  })()

  const res = await adminErc20Bridger.setGateways(
    l1Signer,
    l2Provider,
    gateways.map((g, i) => ({
      tokenAddr: tokens[i],
      gatewayAddr: gateways[i],
    }))
  )
  console.log('Getting gateway(s)', res)
  const rec = await res.wait()
  console.log('Done', rec)

  if (rec.status !== 1) {
    throw new Error(`SetGateways failed on L1 ${rec.transactionHash}`)
  }

  console.log('redeeming retryable ticket:')
  const l2Tx = await rec.getL1ToL2Message(l2Signer)
  const messageRes = await l2Tx.wait(false)
  if (messageRes.status === L1ToL2MessageStatus.FUNDS_DEPOSITED_ON_L2) {
    const redeemRes = await l2Tx.redeem()
    const redeemRec = await redeemRes.wait()
    console.log('Done redeeming:', redeemRec)
    console.log(redeemRec.status === 1 ? ' success!' : 'failed...')
    return redeemRec
  } else console.log(`Unpexpected message status: ${messageRes.status}.`)
}

export const checkRetryableStatus = async (l1Hash: string): Promise<void> => {
  const { l1Signer, l2Signer } = await instantiateBridge()
  const l1Provider = l1Signer.provider!
  const l2Provider = l2Signer.provider!
  const rec = await l1Provider.getTransactionReceipt(l1Hash)
  const message = await new L1TransactionReceipt(rec).getL1ToL2Message(
    l2Provider
  )

  if (!message) throw new Error('no seq nums')

  const autoRedeemHash = message.autoRedeemId
  const autoRedeemRec = await l2Provider.getTransactionReceipt(autoRedeemHash)

  const redeemTxnHash = message.l2TxHash
  const redeemTxnRec = await l2Provider.getTransactionReceipt(redeemTxnHash)

  const retryableTicketHash = message.retryableCreationId

  const retryableTicketRec = await l2Provider.getTransactionReceipt(
    retryableTicketHash
  )

  console.log('*** autoRedeemHash', autoRedeemHash)
  console.log(
    '*** autoRedeem status',
    autoRedeemRec ? autoRedeemRec.status : autoRedeemRec
  )
  if (autoRedeemRec && autoRedeemRec.status !== 1) {
    console.log('**** autoredeem receipt', autoRedeemRec)
  }

  console.log('*** redeemTxnHash', redeemTxnHash)
  console.log(
    '*** redeemTxnHash status',
    redeemTxnRec ? redeemTxnRec.status : redeemTxnRec
  )
  if (redeemTxnRec && redeemTxnRec.status !== 1) {
    console.log('**** redeemTxnHash receipt', redeemTxnHash)
  }

  console.log('*** retryableTicketHash', retryableTicketHash)
  console.log(
    '*** retryableTicket status',
    retryableTicketRec ? retryableTicketRec.status : retryableTicketRec
  )
  if (retryableTicketRec && retryableTicketRec.status !== 1) {
    console.log('**** retryableTicket receipt', retryableTicketHash)
  }
}
