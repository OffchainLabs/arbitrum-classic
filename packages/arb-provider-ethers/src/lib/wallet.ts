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

import { ArbClient } from './client'
import { ArbProvider } from './provider'
import { GlobalInbox } from './abi/GlobalInbox'
import { ArbSysFactory } from './abi/ArbSysFactory'

import * as ethers from 'ethers'

const ARB_SYS_ADDRESS = '0x0000000000000000000000000000000000000064'

export class ArbWallet extends ethers.Signer {
  public client: ArbClient
  public signer: ethers.Signer
  public provider: ArbProvider
  public globalInboxCache?: GlobalInbox
  public seqCache?: number
  public pubkey?: string
  public channelMode: boolean

  constructor(
    client: ArbClient,
    signer: ethers.Signer,
    provider: ArbProvider,
    channelMode: boolean
  ) {
    super()
    this.signer = signer
    this.provider = provider
    this.client = client
    this.seqCache = undefined
    this.pubkey = undefined
    this.channelMode = channelMode
  }

  public async generateSeq(): Promise<number> {
    if (!this.seqCache) {
      const seq = await this.provider.getTransactionCount(
        await this.getAddress()
      )
      this.seqCache = seq
      return seq
    }
    return this.seqCache
  }

  public async incrementSeq(): Promise<void> {
    if (this.seqCache === undefined) {
      throw Error('Sequence number must have already been generated')
    }
    this.seqCache++
  }

  public async globalInboxConn(): Promise<GlobalInbox> {
    if (!this.globalInboxCache) {
      const globalInbox = await this.provider.globalInboxConn()
      const linkedGlobalInbox = globalInbox.connect(this.signer)
      this.globalInboxCache = linkedGlobalInbox
      return linkedGlobalInbox
    }
    return this.globalInboxCache
  }

  public getAddress(): Promise<string> {
    return this.signer.getAddress()
  }

  public async signMessage(
    message: ethers.utils.Arrayish | string
  ): Promise<string> {
    return this.signer.signMessage(message)
  }

  public async withdrawEthFromChain(
    value: ethers.utils.BigNumberish
  ): Promise<ethers.providers.TransactionResponse> {
    const valueNum = ethers.utils.bigNumberify(value)
    const arbsys = ArbSysFactory.connect(ARB_SYS_ADDRESS, this)
    return arbsys.withdrawEth(await this.getAddress(), valueNum)
  }

  public async withdrawEth(): Promise<ethers.providers.TransactionResponse> {
    const globalInbox = await this.globalInboxConn()
    return globalInbox.withdrawEth()
  }

  public async withdrawERC20(
    erc20: string
  ): Promise<ethers.providers.TransactionResponse> {
    const globalInbox = await this.globalInboxConn()
    return globalInbox.withdrawERC20(erc20)
  }

  public async withdrawERC721(
    erc721: string,
    tokenId: ethers.utils.BigNumberish
  ): Promise<ethers.providers.TransactionResponse> {
    const globalInbox = await this.globalInboxConn()
    return globalInbox.withdrawERC721(erc721, tokenId)
  }

  public async depositERC20(
    to: string,
    erc20: string,
    value: ethers.utils.BigNumberish
  ): Promise<ethers.providers.TransactionResponse> {
    const sendValue = ethers.utils.bigNumberify(value)
    const chain = await this.provider.getVmID()
    const globalInbox = await this.globalInboxConn()
    const tx = await globalInbox.depositERC20Message(
      chain,
      to,
      erc20,
      sendValue
    )
    return this.provider._wrapTransaction(tx, tx.hash)
  }

  public async depositERC721(
    to: string,
    erc721: string,
    tokenId: ethers.utils.BigNumberish
  ): Promise<ethers.providers.TransactionResponse> {
    const chain = await this.provider.getVmID()
    const globalInbox = await this.globalInboxConn()
    const tx = await globalInbox.depositERC721Message(
      chain,
      to,
      erc721,
      tokenId
    )
    return this.provider._wrapTransaction(tx, tx.hash)
  }

  public async depositETH(
    to: string,
    value: ethers.utils.BigNumberish
  ): Promise<ethers.providers.TransactionResponse> {
    const chain = await this.provider.getVmID()
    const globalInbox = await this.globalInboxConn()
    const tx = await globalInbox.depositEthMessage(chain, to, { value })
    return this.provider._wrapTransaction(tx, tx.hash)
  }

  public async sendTransactionMessage(
    to: string,
    value: ethers.utils.BigNumberish,
    data: string
  ): Promise<ethers.providers.TransactionResponse> {
    const vmId = await this.provider.getVmID()
    const valueNum = ethers.utils.bigNumberify(value)
    const globalInbox = await this.globalInboxConn()
    const seq = await this.generateSeq()
    const tx = await globalInbox.sendTransactionMessage(
      vmId,
      to,
      seq,
      valueNum,
      data
    )
    const tx2 = this.provider._wrapTransaction(tx, tx.hash)
    await this.incrementSeq()
    return tx2
  }

  public async sendTransaction(
    transaction: ethers.providers.TransactionRequest
  ): Promise<ethers.providers.TransactionResponse> {
    if (!transaction.to) {
      throw Error("Can't send transaction without destination")
    }
    const to = await transaction.to
    let data = '0x'
    if (transaction.data) {
      data = ethers.utils.hexlify(await transaction.data)
    }

    let value = ethers.utils.bigNumberify(0)
    if (transaction.value) {
      value = ethers.utils.bigNumberify(await transaction.value) // eslint-disable-line require-atomic-updates
    }
    return this.sendTransactionMessage(to, value, data)
  }
}
