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

import { Block, Provider } from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'
import { BigNumber } from '@ethersproject/bignumber'
import { ContractTransaction, PayableOverrides } from '@ethersproject/contracts'

import { ArbSys__factory } from './abi/factories/ArbSys__factory'
import { ArbSys } from './abi/ArbSys'
import { StandardArbERC20__factory } from './abi/factories/StandardArbERC20__factory'
import { StandardArbERC20 } from './abi/StandardArbERC20'
import { L2GatewayRouter__factory } from './abi/factories/L2GatewayRouter__factory'
import { L2GatewayRouter } from './abi/L2GatewayRouter'
import { ArbRetryableTx__factory } from './abi/factories/ArbRetryableTx__factory'
import { ArbRetryableTx } from './abi/ArbRetryableTx'

import {
  ARB_SYS_ADDRESS,
  ARB_RETRYABLE_TX_ADDRESS,
} from './precompile_addresses'
import { Network } from './networks'
import { ArbMulticall2__factory } from './abi/factories/ArbMulticall2__factory'
import { BridgeHelper } from './bridge_helpers'

export interface L2TokenData {
  contract: StandardArbERC20
  balance: BigNumber
}

/**
 * L2 side only of {@link Bridge}
 */
export class L2Bridge {
  l2Signer: Signer
  arbSys: ArbSys
  l2GatewayRouter: L2GatewayRouter
  l2Provider: Provider
  arbRetryableTx: ArbRetryableTx
  walletAddressCache?: string
  network: Network

  constructor(network: Network, l2Signer: Signer) {
    this.l2Signer = l2Signer
    this.network = network

    const l2Provider = l2Signer.provider

    if (l2Provider === undefined) {
      throw new Error('Signer must be connected to an (Arbitrum) provider')
    }
    this.l2Provider = l2Provider

    this.arbSys = ArbSys__factory.connect(ARB_SYS_ADDRESS, l2Signer)

    this.l2GatewayRouter = L2GatewayRouter__factory.connect(
      network.tokenBridge.l2GatewayRouter,
      l2Signer
    )

    this.arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      l2Signer
    )
  }

  public async setSigner(newSigner: Signer) {
    const newL2Provider = newSigner.provider
    if (newL2Provider === undefined) {
      throw new Error('Signer must be connected to an (Arbitrum) provider')
    }
    // check chainId to ensure its still in the same network.
    const [prevNetwork, newNetwork] = await Promise.all([
      this.l2Provider.getNetwork(),
      newL2Provider.getNetwork(),
    ])
    if (prevNetwork.chainId !== newNetwork.chainId)
      throw new Error('Error. New signer in L2 is a different network.')

    this.l2Provider = newL2Provider
    this.l2Signer = newSigner
    // we need to update the cache
    // TODO: remove this cache. can we memoize based on the signer? useCallback style
    this.walletAddressCache = await this.l2Signer.getAddress()

    // TODO: is it worth keeping contracts instantiated? we can instead have a util function
    this.arbSys = this.arbSys.connect(newSigner)
    this.l2GatewayRouter = this.l2GatewayRouter.connect(newSigner)
    this.arbRetryableTx = this.arbRetryableTx.connect(newSigner)
  }

  // CHRIS: on the bridger now
  /**
   * Initiate Ether withdrawal (via ArbSys)
   */
  // public async withdrawETH(
  //   value: BigNumber,
  //   destinationAddress?: string,
  //   overrides: PayableOverrides = {}
  // ): Promise<ContractTransaction> {
  //   const address = destinationAddress || (await this.getWalletAddress())
  //   return this.arbSys.functions.withdrawEth(address, {
  //     value,
  //     ...overrides,
  //   })
  // }

  // CHRIS: now on ethbridger
  // public async estimateGasWithdrawETH(
  //   value: BigNumber,
  //   destinationAddress?: string,
  //   overrides: PayableOverrides = {}
  // ): Promise<BigNumber> {
  //   const address = destinationAddress || (await this.getWalletAddress())
  //   return this.arbSys.estimateGas.withdrawEth(address, {
  //     value,
  //     ...overrides,
  //   })
  // }

  // public getLatestBlock(): Promise<Block> {
  //   return this.l2Provider.getBlock('latest')
  // }

  // CHRIS: now on bridger
  /**
   * Initiate token withdrawal (via l2ERC20Gateway)
   */
  // public async withdrawERC20(
  //   erc20l1Address: string,
  //   amount: BigNumber,
  //   destinationAddress?: string,
  //   overrides: PayableOverrides = {}
  // ): Promise<ContractTransaction> {
  //   const to = destinationAddress || (await this.getWalletAddress())

  //   return this.l2GatewayRouter.functions[
  //     'outboundTransfer(address,address,uint256,bytes)'
  //   ](erc20l1Address, to, amount, '0x', overrides)
  // }

  // CHRIS: now on bridger
  // public async estimateGasWithdrawERC20(
  //   erc20l1Address: string,
  //   amount: BigNumber,
  //   destinationAddress?: string,
  //   overrides: PayableOverrides = {}
  // ): Promise<BigNumber> {
  //   const to = destinationAddress || (await this.getWalletAddress())

  //   return this.l2GatewayRouter.estimateGas[
  //     'outboundTransfer(address,address,uint256,bytes)'
  //   ](erc20l1Address, to, amount, '0x', overrides)
  // }

  // public async getL2TokenData(l2ERC20Address: string): Promise<L2TokenData> {
  //   const walletAddress = await this.getWalletAddress()

  //   const arbERC20TokenContract = StandardArbERC20__factory.connect(
  //     l2ERC20Address,
  //     this.l2Signer
  //   )
  //   // this will throw if not a contract / ERC20
  //   const [balance] = await arbERC20TokenContract.functions.balanceOf(
  //     walletAddress
  //   )
  //   // TODO: should we include extra data? ie: `l2ERC20Address.l1Address()` and `l2GatewayRouter.getGateway(erc20L1Address)`
  //   return {
  //     contract: arbERC20TokenContract,
  //     balance,
  //   }
  // }

  // CHRIS: now on token bridger
  // public async getGatewayAddress(erc20L1Address: string): Promise<string> {
  //   return (await this.l2GatewayRouter.functions.getGateway(erc20L1Address))
  //     .gateway
  // }

  // CHRIS: now on token bridger
  // public getERC20L1Address(erc20L2Address: string): Promise<string | null> {
  //   const arbERC20 = StandardArbERC20__factory.connect(
  //     erc20L2Address,
  //     this.l2Signer
  //   )
  //   return arbERC20.functions
  //     .l1Address()
  //     .then(([res]) => res)
  //     .catch(e => {
  //       return null
  //     })
  // }

  // public getTxnSubmissionPrice(
  //   dataSize: BigNumber | number
  // ): Promise<[BigNumber, BigNumber]> {
  //   return this.arbRetryableTx.functions.getSubmissionPrice(dataSize)
  // }

  // public async getWalletAddress(): Promise<string> {
  //   if (this.walletAddressCache) {
  //     return this.walletAddressCache
  //   }
  //   this.walletAddressCache = await this.l2Signer.getAddress()
  //   return this.walletAddressCache
  // }

  // public getL2EthBalance(): Promise<BigNumber> {
  //   return this.l2Signer.getBalance()
  // }

  // public async getMulticallAggregate(functionCalls: MulticallFunctionInput) {
  //   const multicall = ArbMulticall2__factory.connect(
  //     this.network.tokenBridge.l2Multicall,
  //     this.l2Provider
  //   )

  //   return BridgeHelper.getMulticallTryAggregate(functionCalls, multicall)
  // }
}
