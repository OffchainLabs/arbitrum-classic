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
/* eslint-env node */
'use strict'

import { Bridge } from './bridge'

import { L2Transaction, L2Message } from './message'
import { ArbProvider } from './provider'
import { GlobalInbox } from './abi/GlobalInbox'
import { ArbSysFactory } from './abi/ArbSysFactory'
import { TransactionOverrides } from './abi'

import * as ethers from 'ethers'

const ARB_SYS_ADDRESS = '0x0000000000000000000000000000000000000064'

export class ArbWallet extends Bridge implements ethers.Signer {
  public signer: ethers.Signer
  public provider: ArbProvider

  constructor(signer: ethers.Signer, provider: ArbProvider) {
    super(signer, provider.ethProvider, provider.chainAddress)
    this.signer = signer
    this.provider = provider
  }

  public getAddress(): Promise<string> {
    return this.signer.getAddress()
  }

  public signMessage(message: ethers.utils.Arrayish | string): Promise<string> {
    return this.signer.signMessage(message)
  }

  public async withdrawEthFromChain(
    value: ethers.utils.BigNumberish
  ): Promise<ethers.providers.TransactionResponse> {
    const valueNum = ethers.utils.bigNumberify(value)
    const arbsys = ArbSysFactory.connect(ARB_SYS_ADDRESS, this)
    return arbsys.withdrawEth(await this.getAddress(), { value: valueNum })
  }

  public async depositERC20(
    to: string,
    erc20: string,
    value: ethers.utils.BigNumberish,
    overrides?: TransactionOverrides
  ): Promise<ethers.providers.TransactionResponse> {
    const tx = await super.depositERC20(to, erc20, value, overrides)
    return this.provider._wrapTransaction(tx, tx.hash)
  }

  public async depositERC721(
    to: string,
    erc721: string,
    tokenId: ethers.utils.BigNumberish,
    overrides?: TransactionOverrides
  ): Promise<ethers.providers.TransactionResponse> {
    const tx = await super.depositERC20(to, erc721, tokenId, overrides)
    return this.provider._wrapTransaction(tx, tx.hash)
  }

  public async depositETH(
    to: string,
    value: ethers.utils.BigNumberish,
    overrides?: TransactionOverrides
  ): Promise<ethers.providers.TransactionResponse> {
    const tx = await super.depositETH(to, value, overrides)
    return this.provider._wrapTransaction(tx, tx.hash)
  }

  public async sendTransactionMessage(
    l2tx: L2Transaction,
    from: string,
    overrides?: TransactionOverrides
  ): Promise<ethers.providers.TransactionResponse> {
    this.sendL2Message(l2tx, from, overrides)
    const network = await this.provider.getNetwork()
    const tx: ethers.utils.Transaction = {
      data: ethers.utils.hexlify(l2tx.calldata),
      from: from,
      gasLimit: l2tx.maxGas,
      gasPrice: l2tx.gasPriceBid,
      hash: l2tx.messageID(from, network.chainId),
      nonce: l2tx.sequenceNum.toNumber(),
      to: l2tx.destAddress,
      value: l2tx.payment,
      chainId: network.chainId,
    }
    return this.provider._wrapTransaction(tx, tx.hash)
  }

  public async sendTransaction(
    transaction: ethers.providers.TransactionRequest
  ): Promise<ethers.providers.TransactionResponse> {
    try {
      return this.signer.sendTransaction(transaction)
    } catch (e) {
      // Fallback to sending at L1 if arbitrum provider fails
      return this.sendTransactionAtL1(transaction)
    }
  }

  private async sendTransactionAtL1(
    transaction: ethers.providers.TransactionRequest
  ): Promise<ethers.providers.TransactionResponse> {
    let gasLimit = await transaction.gasLimit
    if (!gasLimit) {
      // default to 90000 based on spec
      gasLimit = 90000
    }

    let gasPrice = await transaction.gasPrice
    if (!gasPrice) {
      // What do we want to make the default for this
      gasPrice = 0
    }

    let from = await transaction.from
    if (!from) {
      from = await this.getAddress()
    }

    let nonce = await transaction.nonce
    if (!nonce) {
      nonce = await this.provider.getTransactionCount(from)
    }

    const tx = new L2Transaction(
      gasLimit,
      gasPrice,
      nonce,
      await transaction.to,
      await transaction.value,
      await transaction.data
    )
    return this.sendTransactionMessage(tx, from)
  }
}
