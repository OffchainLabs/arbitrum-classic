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

import { defaultAbiCoder } from '@ethersproject/abi'
import { Provider } from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'
import { BigNumber } from '@ethersproject/bignumber'
import { ContractTransaction, PayableOverrides } from '@ethersproject/contracts'
import { MaxUint256, AddressZero } from '@ethersproject/constants'

import { L1GatewayRouter__factory } from './abi/factories/L1GatewayRouter__factory'
import { L1GatewayRouter } from './abi/L1GatewayRouter'
import { L1ERC20Gateway__factory } from './abi/factories/L1ERC20Gateway__factory'
import { Inbox__factory } from './abi/factories/Inbox__factory'
import { Inbox } from './abi/Inbox'
import { ERC20__factory } from './abi/factories/ERC20__factory'
import { ERC20 } from './abi/ERC20'
import { Multicall2__factory } from './abi/factories/Multicall2__factory'
import { L1ERC20Gateway } from './abi/L1ERC20Gateway'

import { Network } from './networks'
import {
  addressToSymbol,
  MulticallFunctionInput,
  BridgeHelper,
} from './bridge_helpers'

const MIN_APPROVAL = MaxUint256
//TODO handle address update / lowercase

// typing magic from https://stackoverflow.com/a/57364353
export type Await<T> = T extends {
  then(onfulfilled?: (value: infer U) => unknown): unknown
}
  ? U
  : T

export interface L1TokenData {
  contract: ERC20
  balance: BigNumber
  allowed: boolean
  symbol: string
  decimals: number
  name: string
}

export interface DepositParams {
  erc20L1Address: string
  amount: BigNumber
  l1CallValue: BigNumber
  maxSubmissionCost: BigNumber
  maxGas: BigNumber
  gasPriceBid: BigNumber
  destinationAddress?: string
}

/**
 * L1 side only of {@link Bridge}
 */
export class L1Bridge {
  l1Signer: Signer
  l1GatewayRouter: L1GatewayRouter
  walletAddressCache?: string
  inboxCached?: Inbox
  l1Provider: Provider
  network: Network

  constructor(network: Network, l1Signer: Signer) {
    this.l1Signer = l1Signer
    this.network = network

    const l1Provider = l1Signer.provider

    if (l1Provider === undefined) {
      throw new Error('Signer must be connected to an Ethereum provider')
    }
    this.l1Provider = l1Provider
    this.l1GatewayRouter = L1GatewayRouter__factory.connect(
      network.tokenBridge.l1GatewayRouter,
      l1Signer
    )
  }

  public async setSigner(newSigner: Signer) {
    const newL1Provider = newSigner.provider
    if (newL1Provider === undefined) {
      throw new Error('Signer must be connected to an Ethereum provider')
    }
    // check chainId to ensure its still in the same network.
    const [prevNetwork, newNetwork] = await Promise.all([
      this.l1Provider.getNetwork(),
      newL1Provider.getNetwork(),
    ])
    if (prevNetwork.chainId !== newNetwork.chainId)
      throw new Error('Error. New signer in L1 is a different network.')

    this.l1Provider = newL1Provider
    this.l1Signer = newSigner
    // we need to update the cache
    // TODO: remove this cache. can we memoize based on the signer? useCallback style
    this.walletAddressCache = await this.l1Signer.getAddress()
    // TODO: is it worth keeping contracts instantiated?
    this.inboxCached = this.inboxCached?.connect(newSigner)
    this.l1GatewayRouter = this.l1GatewayRouter.connect(newSigner)
  }

  public getERC20L2Address(erc20L1Address: string): Promise<string> {
    return this.l1GatewayRouter.functions
      .calculateL2TokenAddress(erc20L1Address)
      .then(([res]) => res)
  }

  public async contractExists(contractAddress: string): Promise<boolean> {
    const contractCode = await this.l1Provider.getCode(contractAddress)
    return !(contractCode.length > 2)
  }

  public async getL1TokenData(erc20L1Address: string): Promise<L1TokenData> {
    const walletAddress = await this.getWalletAddress()
    const gatewayAddress = await this.getGatewayAddress(erc20L1Address)

    const ethERC20TokenContract = await ERC20__factory.connect(
      erc20L1Address,
      this.l1Signer
    )
    // If this will throw if not a contract / ERC20, which is what we *want*.
    const iface = ERC20__factory.createInterface()
    const functionCalls: MulticallFunctionInput = [
      {
        target: erc20L1Address,
        funcFragment: iface.functions['balanceOf(address)'],
        values: [walletAddress],
      },
      {
        target: erc20L1Address,
        funcFragment: iface.functions['allowance(address,address)'],
        values: [walletAddress, gatewayAddress],
      },
      {
        target: erc20L1Address,
        funcFragment: iface.functions['symbol()'],
      },
      {
        target: erc20L1Address,
        funcFragment: iface.functions['decimals()'],
      },
      {
        target: erc20L1Address,
        funcFragment: iface.functions['name()'],
      },
    ]
    const [
      balanceResult,
      allowanceResult,
      symbolResult,
      decimalsResult,
      nameResult,
    ] = await this.getMulticallAggregate(functionCalls)

    if (balanceResult === undefined || allowanceResult === undefined)
      throw new Error('ERC20 Token does not support balance/allowance')

    type ERC20_FUNCTIONS = ERC20['functions']
    const balance = (
      balanceResult as Await<ReturnType<ERC20_FUNCTIONS['balanceOf']>>
    )[0]
    const allowance = (
      allowanceResult as Await<ReturnType<ERC20_FUNCTIONS['allowance']>>
    )[0]

    const symbol =
      symbolResult === undefined
        ? addressToSymbol(erc20L1Address)
        : (symbolResult as Await<ReturnType<ERC20_FUNCTIONS['symbol']>>)[0]

    const decimals =
      decimalsResult === undefined
        ? 18
        : (decimalsResult as Await<ReturnType<ERC20_FUNCTIONS['decimals']>>)[0]

    const name =
      nameResult === undefined
        ? symbol + '_Token'
        : (nameResult as Await<ReturnType<ERC20_FUNCTIONS['name']>>)[0]

    const allowanceLimit = BigNumber.from(
      '0xffffffffffffffffffffffff'
    ) /** for ERC20s that cap approve at 96 bits  */
    const allowed = allowance.gte(allowanceLimit.div(2))
    return {
      contract: ethERC20TokenContract,
      balance,
      allowed,
      symbol,
      decimals,
      name,
    }
  }

  public async depositETH(
    value: BigNumber,
    maxSubmissionPrice: BigNumber,
    overrides: PayableOverrides = {}
  ): Promise<ContractTransaction> {
    const inbox = await this.getInbox()
    return inbox.functions.depositEth(maxSubmissionPrice, {
      value,
      ...overrides,
    })
  }

  public async getGatewayAddress(erc20L1Address: string): Promise<string> {
    return (await this.l1GatewayRouter.functions.getGateway(erc20L1Address))
      .gateway
  }
  public async getDefaultL1Gateway(): Promise<L1ERC20Gateway> {
    const defaultGatewayAddress = await this.l1GatewayRouter.defaultGateway()

    if (defaultGatewayAddress === AddressZero) {
      console.log(
        'No default network assigned in contract, using standard l1ERC20Gateway:'
      )

      return L1ERC20Gateway__factory.connect(
        this.network.tokenBridge.l1ERC20Gateway,
        this.l1Provider
      )
    }

    return L1ERC20Gateway__factory.connect(
      defaultGatewayAddress,
      this.l1Provider
    )
  }

  public async approveToken(
    erc20L1Address: string,
    amount?: BigNumber,
    overrides: PayableOverrides = {}
  ): Promise<ContractTransaction> {
    // you approve tokens to the gateway that the router will use
    const gatewayAddress = await this.getGatewayAddress(erc20L1Address)
    const contract = await ERC20__factory.connect(erc20L1Address, this.l1Signer)
    return contract.functions.approve(
      gatewayAddress,
      amount || MIN_APPROVAL,
      overrides
    )
  }

  public async deposit(
    {
      erc20L1Address,
      amount,
      maxSubmissionCost,
      maxGas,
      gasPriceBid,
      destinationAddress,
      l1CallValue,
    }: DepositParams,
    overrides: PayableOverrides = {}
  ): Promise<ContractTransaction> {
    const destination = destinationAddress || (await this.getWalletAddress())
    const data = defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [maxSubmissionCost, '0x']
    )

    if (overrides.value)
      throw new Error('L1 call value should be set through l1CallValue param')
    if (l1CallValue.eq(0)) throw new Error('L1 call value should not be zero')
    if (maxSubmissionCost.eq(0))
      throw new Error('Max submission cost should not be zero')

    return this.l1GatewayRouter.functions.outboundTransfer(
      erc20L1Address,
      destination,
      amount,
      maxGas,
      gasPriceBid,
      data,
      { ...overrides, value: l1CallValue }
    )
  }

  public async getWalletAddress(): Promise<string> {
    if (this.walletAddressCache) {
      return this.walletAddressCache
    }
    this.walletAddressCache = await this.l1Signer.getAddress()
    return this.walletAddressCache
  }

  public async getInbox(): Promise<Inbox> {
    if (this.inboxCached) {
      return this.inboxCached
    }
    const gateway = await this.getDefaultL1Gateway()

    const inboxAddress = await gateway.inbox()
    this.inboxCached = Inbox__factory.connect(inboxAddress, this.l1Signer)
    return this.inboxCached
  }

  public getL1EthBalance(): Promise<BigNumber> {
    return this.l1Signer.getBalance()
  }

  public async getMulticallAggregate(functionCalls: MulticallFunctionInput) {
    const multicall = Multicall2__factory.connect(
      this.network.tokenBridge.l1MultiCall,
      this.l1Provider
    )

    return BridgeHelper.getMulticallTryAggregate(functionCalls, multicall)
  }
}
