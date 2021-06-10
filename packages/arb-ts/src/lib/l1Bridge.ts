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
import { L1GatewayRouter__factory } from './abi/factories/L1GatewayRouter__factory'
import { L1GatewayRouter } from './abi/L1GatewayRouter'
import { L1ERC20Gateway } from './abi/L1ERC20Gateway'
import { L1ERC20Gateway__factory } from './abi/factories/L1ERC20Gateway__factory'

import { Inbox } from './abi/Inbox'
import { Inbox__factory } from './abi/factories/Inbox__factory'
import { ERC20 } from './abi/ERC20'
import networks from './networks'

import { ERC20__factory } from './abi/factories/ERC20__factory'
import { addressToSymbol } from './bridge_helpers'
import { PayableOverrides } from '@ethersproject/contracts'

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

/**
 * L1 side only of {@link Bridge}
 */
export class L1Bridge {
  l1Signer: Signer
  l1GatewayRouter: L1GatewayRouter
  walletAddressCache?: string
  inboxCached?: Inbox
  l1Tokens: Tokens
  l1Provider: providers.Provider
  l1EthBalance: BigNumber

  constructor(l1GatewayRouterAddress: string, l1Signer: Signer) {
    this.l1Signer = l1Signer
    this.l1Tokens = {}

    const l1Provider = l1Signer.provider

    if (l1Provider === undefined) {
      throw new Error('Signer must be connected to an Ethereum provider')
    }
    this.l1Provider = l1Provider
    this.l1GatewayRouter = L1GatewayRouter__factory.connect(
      l1GatewayRouterAddress,
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

  public getERC20L2Address(erc20L1Address: string) {
    return this.l1GatewayRouter.functions
      .calculateL2TokenAddress(erc20L1Address)
      .then(([res]) => res)
  }

  public async getAndUpdateL1TokenData(erc20L1Address: string) {
    const tokenData = this.l1Tokens[erc20L1Address] || {
      ERC20: undefined,
      CUSTOM: undefined,
    }
    this.l1Tokens[erc20L1Address] = tokenData
    const walletAddress = await this.getWalletAddress()
    const gatewayAddress = await this.getGatewayAddress(erc20L1Address)

    if (!tokenData.ERC20) {
      if ((await this.l1Provider.getCode(erc20L1Address)).length > 2) {
        // If this will throw if not an ERC20, which is what we *want*.
        const ethERC20TokenContract = await ERC20__factory.connect(
          erc20L1Address,
          this.l1Signer
        )
        const [balance] = await ethERC20TokenContract.functions.balanceOf(
          walletAddress
        )

        const [allowance] = await ethERC20TokenContract.functions.allowance(
          walletAddress,
          gatewayAddress
        )
        // non-standard
        const symbol = await ethERC20TokenContract.functions
          .symbol()
          .then(([res]) => res)
          .catch(_ => addressToSymbol(erc20L1Address))

        const decimals = await ethERC20TokenContract.functions
          .decimals()
          .then(([res]) => res)
          .catch(_ => 18)

        const name = await ethERC20TokenContract.functions
          .name()
          .then(([res]) => res)
          .catch(_ => symbol + '_Token')

        const allowed = await allowance.gte(MIN_APPROVAL.div(2))
        tokenData.ERC20 = {
          contract: ethERC20TokenContract,
          balance,
          allowed,
          symbol,
          decimals,
          name,
        }
      } else {
        throw new Error(`No ERC20 at ${erc20L1Address} `)
      }
    } else {
      const ethERC20TokenContract = await ERC20__factory.connect(
        erc20L1Address,
        this.l1Signer
      )
      const [balance] = await ethERC20TokenContract.functions.balanceOf(
        walletAddress
      )
      tokenData.ERC20.balance = balance

      if (!tokenData.ERC20.allowed) {
        const [allowance] = await ethERC20TokenContract.functions.allowance(
          walletAddress,
          gatewayAddress
        )
        tokenData.ERC20.allowed = allowance.gte(MIN_APPROVAL.div(2))
      }
    }

    return tokenData
  }

  public async depositETH(
    value: BigNumber,
    maxSubmissionPrice: BigNumber,
    overrides: PayableOverrides = {}
  ) {
    const inbox = await this.getInbox()
    return inbox.functions.depositEth(maxSubmissionPrice, {
      value,
      ...overrides,
    })
  }

  public async getGatewayAddress(erc20L1Address: string) {
    return (await this.l1GatewayRouter.functions.getGateway(erc20L1Address))
      .gateway
  }
  public async getDefaultL1Gateway() {
    const defaultGatewayAddress = await this.l1GatewayRouter.defaultGateway()

    if (defaultGatewayAddress === constants.AddressZero) {
      throw new Error('No default gateway set')
    }
    return L1ERC20Gateway__factory.connect(
      defaultGatewayAddress,
      this.l1Provider
    )
  }

  public async approveToken(
    erc20L1Address: string,
    overrides: PayableOverrides = {}
  ) {
    const tokenData = await this.getAndUpdateL1TokenData(erc20L1Address)
    if (!tokenData.ERC20) {
      throw new Error(`Can't approve; token ${erc20L1Address} not found`)
    }
    // you approve tokens to the gateway that the router will use
    const gatewayAddress = await this.getGatewayAddress(erc20L1Address)
    return tokenData.ERC20.contract.functions.approve(
      gatewayAddress,
      MIN_APPROVAL,
      overrides
    )
  }

  public async deposit(
    erc20L1Address: string,
    amount: BigNumber,
    maxSubmissionCost: BigNumber,
    maxGas: BigNumber,
    gasPriceBid: BigNumber,
    destinationAddress?: string,
    overrides: PayableOverrides = {}
  ) {
    const destination = destinationAddress || (await this.getWalletAddress())
    const tokenData = await this.getAndUpdateL1TokenData(erc20L1Address)
    if (!tokenData.ERC20) {
      throw new Error(`Can't deposit; No ERC20 at ${erc20L1Address}`)
    }
    const data = utils.defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )
    return this.l1GatewayRouter.functions.outboundTransfer(
      erc20L1Address,
      destination,
      amount,
      maxGas,
      gasPriceBid,
      data,
      overrides
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
    const { chainId } = await this.l1Provider.getNetwork()
    if (chainId === 1) {
      // patch: mainnet has no default gateway set
      this.inboxCached = Inbox__factory.connect(
        '0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f',
        this.l1Signer
      )
      return this.inboxCached
    }
    const gateway = await this.getDefaultL1Gateway()

    const inboxAddress = await gateway.inbox()
    this.inboxCached = Inbox__factory.connect(inboxAddress, this.l1Signer)
    return this.inboxCached
  }

  public async getAndUpdateL1EthBalance(): Promise<BigNumber> {
    const bal = await this.l1Signer.getBalance()
    this.l1EthBalance = bal
    return bal
  }
}
