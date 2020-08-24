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

import { TransactionResponse } from 'ethers/providers'
import { BigNumberish, BigNumber } from 'ethers/utils'
import { Signer } from 'ethers'

export class L1Bridge {
  public globalInboxCache?: GlobalInbox

  constructor(
    public signer: Signer,
    public chainAddress: string | Promise<string>
  ) {}

  public async globalInbox(): Promise<GlobalInbox> {
    if (!this.globalInboxCache) {
      const arbRollup = ArbRollupFactory.connect(
        await this.chainAddress,
        this.signer
      )
      const globalInboxAddress = await arbRollup.globalInbox()
      const globalInbox = GlobalInboxFactory.connect(
        globalInboxAddress,
        this.signer
      ).connect(this.signer)
      this.globalInboxCache = globalInbox
      return globalInbox
    }
    return this.globalInboxCache
  }

  public async withdrawEthFromLockbox(): Promise<TransactionResponse> {
    const globalInbox = await this.globalInbox()
    return globalInbox.withdrawEth()
  }

  public async withdrawERC20FromLockbox(
    erc20: string,
    overrides: TransactionOverrides = {}
  ): Promise<TransactionResponse> {
    const globalInbox = await this.globalInbox()
    return globalInbox.withdrawERC20(erc20, overrides)
  }

  public async withdrawERC721FromLockbox(
    erc721: string,
    tokenId: BigNumberish,
    overrides: TransactionOverrides = {}
  ): Promise<TransactionResponse> {
    const globalInbox = await this.globalInbox()
    return globalInbox.withdrawERC721(erc721, tokenId, overrides)
  }

  public async depositERC20(
    to: string,
    erc20: string,
    value: BigNumberish,
    overrides: TransactionOverrides = {}
  ): Promise<TransactionResponse> {
    const globalInbox = await this.globalInbox()
    return globalInbox.depositERC20Message(
      await this.chainAddress,
      erc20,
      to,
      value,
      overrides
    )
  }

  public async depositERC721(
    to: string,
    erc721: string,
    tokenId: BigNumberish,
    overrides: TransactionOverrides = {}
  ): Promise<TransactionResponse> {
    const globalInbox = await this.globalInbox()
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
    value: BigNumberish,
    overrides: TransactionOverrides = {}
  ): Promise<TransactionResponse> {
    const globalInbox = await this.globalInbox()
    return globalInbox.depositEthMessage(await this.chainAddress, to, {
      ...overrides,
      value,
    })
  }

  public async transferPayment(
    originalOwner: string,
    newOwner: string,
    messageIndex: BigNumberish,
    overrides: TransactionOverrides = {}
  ): Promise<TransactionResponse> {
    const globalInbox = await this.globalInbox()
    return globalInbox.transferPayment(
      originalOwner,
      newOwner,
      messageIndex,
      overrides
    )
  }

  public async sendL2Message(
    l2tx: L2Transaction,
    from: string,
    overrides: TransactionOverrides = {}
  ): Promise<TransactionResponse> {
    const walletAddress = await this.signer.getAddress()
    if (from.toLowerCase() != walletAddress.toLowerCase()) {
      throw Error(
        `Can only send from wallet address ${from}, but tried to send from ${walletAddress}`
      )
    }
    const globalInbox = await this.globalInbox()
    return globalInbox.sendL2Message(
      await this.chainAddress,
      new L2Message(l2tx).asData(),
      overrides
    )
  }

  public async getEthLockBoxBalance(address: string): Promise<BigNumber> {
    const globalInbox = await this.globalInbox()
    return globalInbox.getEthBalance(address)
  }
  public async getERC20LockBoxBalance(
    contractAddress: string,
    ownerAddress: string
  ): Promise<BigNumber> {
    const globalInbox = await this.globalInbox()
    return globalInbox.getERC20Balance(contractAddress, ownerAddress)
  }
  public async getERC721LockBoxTokens(
    contractAddress: string,
    ownerAddress: string
  ): Promise<BigNumber[]> {
    const globalInbox = await this.globalInbox()
    return globalInbox.getERC721Tokens(contractAddress, ownerAddress)
  }
}
