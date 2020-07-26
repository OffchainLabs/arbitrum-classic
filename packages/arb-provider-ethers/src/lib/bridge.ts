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

import { L2Transaction, L2Message } from './message'
import { GlobalInbox } from './abi/GlobalInbox'
import { GlobalInboxFactory } from './abi/GlobalInboxFactory'
import { ArbRollupFactory } from './abi/ArbRollupFactory'
import { TransactionOverrides } from './abi'

import * as ethers from 'ethers'

export class Bridge {
  public globalInboxCache?: GlobalInbox

  constructor(
    public signer: ethers.Signer,
    public provider: ethers.providers.Provider,
    public chainAddress: string | Promise<string>
  ) {}

  private async globalInboxConn(): Promise<GlobalInbox> {
    if (!this.globalInboxCache) {
      const arbRollup = ArbRollupFactory.connect(
        await this.chainAddress,
        this.provider
      )
      const globalInboxAddress = await arbRollup.globalInbox()
      const globalInbox = GlobalInboxFactory.connect(
        globalInboxAddress,
        this.provider
      ).connect(this.signer)
      this.globalInboxCache = globalInbox
      return globalInbox
    }
    return this.globalInboxCache
  }

  public async withdrawEthFromLockbox(): Promise<
    ethers.providers.TransactionResponse
  > {
    const globalInbox = await this.globalInboxConn()
    return globalInbox.withdrawEth()
  }

  public async withdrawERC20FromLockbox(
    erc20: string,
    overrides?: TransactionOverrides
  ): Promise<ethers.providers.TransactionResponse> {
    const globalInbox = await this.globalInboxConn()
    return globalInbox.withdrawERC20(erc20, overrides)
  }

  public async withdrawERC721FromLockbox(
    erc721: string,
    tokenId: ethers.utils.BigNumberish,
    overrides?: TransactionOverrides
  ): Promise<ethers.providers.TransactionResponse> {
    const globalInbox = await this.globalInboxConn()
    return globalInbox.withdrawERC721(erc721, tokenId, overrides)
  }

  public async depositERC20(
    to: string,
    erc20: string,
    value: ethers.utils.BigNumberish,
    overrides?: TransactionOverrides
  ): Promise<ethers.providers.TransactionResponse> {
    const sendValue = ethers.utils.bigNumberify(value)
    const globalInbox = await this.globalInboxConn()
    return globalInbox.depositERC20Message(
      await this.chainAddress,
      erc20,
      to,
      sendValue,
      overrides
    )
  }

  public async depositERC721(
    to: string,
    erc721: string,
    tokenId: ethers.utils.BigNumberish,
    overrides?: TransactionOverrides
  ): Promise<ethers.providers.TransactionResponse> {
    const globalInbox = await this.globalInboxConn()
    return globalInbox.depositERC721Message(
      await this.chainAddress,
      erc721,
      to,
      tokenId,
      overrides
    )
  }

  public async depositETH(
    to: string,
    value: ethers.utils.BigNumberish,
    overrides?: TransactionOverrides
  ): Promise<ethers.providers.TransactionResponse> {
    const globalInbox = await this.globalInboxConn()
    return globalInbox.depositEthMessage(await this.chainAddress, to, {
      ...overrides,
      value,
    })
  }

  public async transferPayment(
    originalOwner: string,
    newOwner: string,
    nodeHash: string,
    messageIndex: ethers.utils.BigNumberish,
    overrides?: TransactionOverrides
  ): Promise<ethers.providers.TransactionResponse> {
    const msgIndex = ethers.utils.bigNumberify(messageIndex)
    const globalInbox = await this.globalInboxConn()
    return globalInbox.transferPayment(
      originalOwner,
      newOwner,
      nodeHash,
      msgIndex,
      overrides
    )
  }

  public async sendL2Message(
    l2tx: L2Transaction,
    from: string,
    overrides?: TransactionOverrides
  ): Promise<ethers.providers.TransactionResponse> {
    const walletAddress = await this.signer.getAddress()
    if (from.toLowerCase() != walletAddress.toLowerCase()) {
      throw Error(
        `Can only send from wallet address ${from}, but tried to send from ${walletAddress}`
      )
    }
    const globalInbox = await this.globalInboxConn()
    return globalInbox.sendL2Message(
      await this.chainAddress,
      new L2Message(l2tx).asData()
    )
  }
}
