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
import { Signer, BigNumber, providers, constants } from 'ethers'
import { InboxFactory } from './abi/InboxFactory'
import { EthERC20BridgeFactory } from './abi/EthERC20BridgeFactory'
import { EthERC20Bridge } from './abi/EthERC20Bridge'
import { Inbox } from './abi/Inbox'
import { ERC20 } from './abi/ERC20'

import { ERC20Factory } from './abi/ERC20Factory'
import { ERC777Factory } from './abi/ERC777Factory'

const MIN_APPROVAL = constants.MaxUint256
//TODO handle address update / lowercase

interface L1TokenData {
  ERC20?: { contract: ERC20; balance: BigNumber; allowed: boolean }
  ERC777?: { contract: ERC777Factory; balance: BigNumber; allowed: boolean }
  CUSTOM?: { contract: ERC20; balance: BigNumber; allowed: boolean }
}

export interface Tokens {
  [contractAddress: string]: L1TokenData | undefined
}

export class L1Bridge {
  l1Signer: Signer
  ethERC20Bridge: EthERC20Bridge
  walletAddressCache?: string
  inboxCached?: Inbox
  l1Tokens: Tokens
  l1Provider: providers.Provider

  constructor(erc20BridgeAddress: string, l1Signer: Signer) {
    this.l1Signer = l1Signer
    this.l1Tokens = {}

    const l1Provider = l1Signer.provider

    if (l1Provider === undefined) {
      throw new Error('Signer must be connected to an Ethereum provider')
    }
    this.l1Provider = l1Provider
    this.ethERC20Bridge = EthERC20BridgeFactory.connect(
      erc20BridgeAddress,
      l1Signer
    )
  }

  public async updateAllL1Tokens() {
    for (const l1Address in this.l1Tokens) {
      await this.getAndUpdateL1TokenData(l1Address)
    }
    return this.l1Tokens
  }

  public async getAndUpdateL1TokenData(erc20L1Address: string) {
    if (!this.l1Tokens[erc20L1Address]) {
      this.l1Tokens[erc20L1Address] = {
        ERC20: undefined,
        ERC777: undefined,
        CUSTOM: undefined,
      }
    }
    const tokenData = this.l1Tokens[erc20L1Address] as L1TokenData // truthiness is ensured above
    const walletAddress = await this.getWalletAddress()
    const indboxAddress = (await this.getInbox()).address

    if (!tokenData.ERC20) {
      if ((await this.l1Provider.getCode(erc20L1Address)).length > 2) {
        const ethERC20TokenContract = await ERC20Factory.connect(
          erc20L1Address,
          this.l1Signer
        )
        const balance = await ethERC20TokenContract.balanceOf(walletAddress)

        const allowance = await ethERC20TokenContract.allowance(
          walletAddress,
          indboxAddress
        )
        const allowed = await allowance.gte(MIN_APPROVAL.div(2))
        tokenData.ERC20 = {
          contract: ethERC20TokenContract,
          balance,
          allowed,
        }
      } else {
        console.warn(`ERC20 at ${erc20L1Address} not found on L1`)
      }
    } else {
      const ethERC20TokenContract = await ERC20Factory.connect(
        erc20L1Address,
        this.l1Signer
      )
      const balance = await ethERC20TokenContract.balanceOf(walletAddress)
      tokenData.ERC20.balance = balance

      if (!tokenData.ERC20.allowed) {
        const allowance = await ethERC20TokenContract.allowance(
          walletAddress,
          indboxAddress
        )
        tokenData.ERC20.allowed = allowance.gte(MIN_APPROVAL.div(2))
      }
    }

    // TODO: erc777? meh
    return tokenData
  }

  public async depositETH(value: BigNumber, destinationAddress?: string) {
    const address = destinationAddress || (await this.getWalletAddress())
    const inbox = await this.getInbox()
    return inbox.depositEth(address, {
      value,
    })
  }

  public async approveToken(erc20L1Address: string) {
    const tokenData = await this.getAndUpdateL1TokenData(erc20L1Address)
    const inboxAddress = (await this.getInbox()).address
    if (!tokenData.ERC20) {
      throw new Error(`Can't approve; token ${erc20L1Address} not found`)
    }

    return tokenData.ERC20.contract.approve(inboxAddress, MIN_APPROVAL)
  }

  public async depositAsERC20(
    erc20L1Address: string,
    amount: BigNumber,
    maxGas: BigNumber,
    gasPriceBid: BigNumber,
    destinationAddress?: string
  ) {
    const destination = destinationAddress || (await this.getWalletAddress())
    this.ethERC20Bridge.depositAsERC20(
      erc20L1Address,
      destination,
      amount,
      maxGas,
      gasPriceBid
    )
  }
  public async depositAsERC777(
    erc20L1Address: string,
    amount: BigNumber,
    maxGas: BigNumber,
    gasPriceBid: BigNumber,
    destinationAddress?: string
  ) {
    const destination = destinationAddress || (await this.getWalletAddress())
    this.ethERC20Bridge.depositAsERC777(
      erc20L1Address,
      destination,
      amount,
      maxGas,
      gasPriceBid
    )
  }

  public async getWalletAddress() {
    if (this.walletAddressCache) {
      return this.walletAddressCache
    }
    this.walletAddressCache = await this.l1Signer.getAddress()
    return this.walletAddressCache
  }

  public async getInbox() {
    if (this.inboxCached) {
      return this.inboxCached
    }
    const inboxAddress = await this.ethERC20Bridge.inbox()
    this.inboxCached = InboxFactory.connect(inboxAddress, this.l1Signer)
    return this.inboxCached
  }
}
