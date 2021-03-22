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
import { Signer, BigNumber, providers, constants, utils } from 'ethers'
import { EthERC20Bridge__factory } from './abi/factories/EthERC20Bridge__Factory'
import { EthERC20Bridge } from './abi/EthERC20Bridge'
import { Inbox } from './abi/Inbox'
import { Inbox__factory } from './abi/factories/Inbox__Factory'
import { ERC20 } from './abi/ERC20'

import { ERC20__factory } from './abi/factories/ERC20__Factory'
import { OZERC777__factory } from './abi/factories/OZERC777__Factory'
import { addressToSymbol } from './bridge_helpers'

utils.computeAddress
const MIN_APPROVAL = constants.MaxUint256
//TODO handle address update / lowercase

export interface L1TokenData {
  ERC20?: {
    contract: ERC20
    balance: BigNumber
    allowed: boolean
    symbol: string
    decimals: number
    name: string
  }
  ERC777?: {
    contract: OZERC777__factory
    balance: BigNumber
    allowed: boolean
    symbol: string
  }
  CUSTOM?: {
    contract: ERC20
    balance: BigNumber
    allowed: boolean
    symbol: string
  }
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
  l1EthBalance: BigNumber

  constructor(erc20BridgeAddress: string, l1Signer: Signer) {
    this.l1Signer = l1Signer
    this.l1Tokens = {}

    const l1Provider = l1Signer.provider

    if (l1Provider === undefined) {
      throw new Error('Signer must be connected to an Ethereum provider')
    }
    this.l1Provider = l1Provider
    this.ethERC20Bridge = EthERC20Bridge__factory.connect(
      erc20BridgeAddress,
      l1Signer
    )
    this.l1EthBalance = BigNumber.from(0)
  }

  public async updateAllL1Tokens() {
    for (const l1Address in this.l1Tokens) {
      await this.getAndUpdateL1TokenData(l1Address)
    }
    return this.l1Tokens
  }

  public async getAndUpdateL1TokenData(erc20L1Address: string) {
    const tokenData = this.l1Tokens[erc20L1Address] || {
      ERC20: undefined,
      ERC777: undefined,
      CUSTOM: undefined,
    }
    this.l1Tokens[erc20L1Address] = tokenData
    const walletAddress = await this.getWalletAddress()
    const indboxAddress = (await this.getInbox()).address

    if (!tokenData.ERC20) {
      if ((await this.l1Provider.getCode(erc20L1Address)).length > 2) {
        // If this will throw if not an ERC20, which is what we *want*.
        const ethERC20TokenContract = await ERC20__factory.connect(
          erc20L1Address,
          this.l1Signer
        )
        const balance = await ethERC20TokenContract.balanceOf(walletAddress)

        const allowance = await ethERC20TokenContract.allowance(
          walletAddress,
          indboxAddress
        )
        // non-standard
        let symbol, decimals, name
        try {
          symbol = await ethERC20TokenContract.symbol()
          decimals = await ethERC20TokenContract.decimals()
          name = await ethERC20TokenContract.name()
        } catch (e) {
          console.info(
            `Weird but technically standard ERC20! ah! ${erc20L1Address}`
          )
          symbol = addressToSymbol(erc20L1Address)
          decimals = 18 // 🤷‍♂️
          name = symbol + '_Token'
        }

        const allowed = await allowance.gte(MIN_APPROVAL.div(2))
        tokenData.ERC20 = {
          contract: ethERC20TokenContract,
          balance,
          allowed,
          symbol,
          decimals,
          name,
        }
        console.warn('DDDDDD erc token data added')
      } else {
        throw new Error(`No ERC20 at ${erc20L1Address} `)
      }
    } else {
      const ethERC20TokenContract = await ERC20__factory.connect(
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
    maxSubmissionCost: BigNumber,
    maxGas: BigNumber,
    gasPriceBid: BigNumber,
    destinationAddress?: string
  ) {
    const destination = destinationAddress || (await this.getWalletAddress())
    const tokenData = await this.getAndUpdateL1TokenData(erc20L1Address)
    if (!tokenData.ERC20) {
      throw new Error(`Can't deposit; No ERC20 at ${erc20L1Address}`)
    }
    return this.ethERC20Bridge.depositAsERC20(
      erc20L1Address,
      destination,
      amount,
      maxSubmissionCost,
      maxGas,
      gasPriceBid,
      ''
    )
  }
  public async depositAsERC777(
    erc20L1Address: string,
    amount: BigNumber,
    maxSubmissionCost: BigNumber,
    maxGas: BigNumber,
    gasPriceBid: BigNumber,
    destinationAddress?: string
  ) {
    const destination = destinationAddress || (await this.getWalletAddress())
    return this.ethERC20Bridge.depositAsERC777(
      erc20L1Address,
      destination,
      amount,
      maxSubmissionCost,
      maxGas,
      gasPriceBid,
      ''
    )
  }

  public async depositAsCustomToken(
    erc20L1Address: string,
    amount: BigNumber,
    maxSubmissionCost: BigNumber,
    maxGas: BigNumber,
    gasPriceBid: BigNumber,
    destinationAddress?: string
  ) {
    const destination = destinationAddress || (await this.getWalletAddress())
    const customTokenL2Address = await this.ethERC20Bridge.customL2Tokens(
      erc20L1Address
    )
    if (customTokenL2Address === constants.AddressZero) {
      throw new Error(`Custom token at ${erc20L1Address} not registtered on L2`)
    }
    return this.ethERC20Bridge.depositAsCustomToken(
      erc20L1Address,
      destination,
      amount,
      maxSubmissionCost,
      maxGas,
      gasPriceBid,
      ''
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
    this.inboxCached = Inbox__factory.connect(inboxAddress, this.l1Signer)
    return this.inboxCached
  }

  public async getAndUpdateL1EthBalance(): Promise<BigNumber> {
    const bal = await this.l1Signer.getBalance()
    this.l1EthBalance = bal
    return bal
  }
}
