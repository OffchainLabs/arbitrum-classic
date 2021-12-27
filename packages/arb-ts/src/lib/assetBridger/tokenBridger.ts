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
import { Signer } from '@ethersproject/abstract-signer'
import { Provider } from '@ethersproject/abstract-provider'
import { PayableOverrides } from '@ethersproject/contracts'
import { Zero, MaxUint256 } from '@ethersproject/constants'
import { Logger } from '@ethersproject/logger'
import { BigNumber, ethers } from 'ethers'

import {
  L1GatewayRouter__factory,
  L2GatewayRouter__factory,
  L1ERC20Gateway__factory,
  L1WethGateway__factory,
  StandardArbERC20__factory,
  ERC20__factory,
  ERC20,
  Multicall2__factory,
} from '../abi'
import { BridgeHelper } from '../bridge_helpers'
import { DepositParams } from '../l1Bridge'
import {
  GasOverrides,
  L1ToL2MessageGasEstimator,
} from '../message/L1ToL2MessageGasEstimator'
import { SignerProviderUtils } from '../utils/signerOrProvider'
import { L2Network } from '../utils/networks'
import { isError } from '../utils/lib'
import { ArbTsError } from '../errors'

import { EthDepositBase, EthWithdrawParams } from './ethBridger'
import { AssetBridger } from './assetBridger'

// CHRIS: do something better with these? we have different defaults
// const DEFAULT_SUBMISSION_PERCENT_INCREASE = BigNumber.from(400)
// const DEFAULT_MAX_GAS_PERCENT_INCREASE = BigNumber.from(50)

export interface TokenApproveParams {
  /**
   * L1 signer whose tokens are being approved
   */
  l1Signer: Signer

  /**
   * L1 address of the ERC20 token contract
   */
  erc20L1Address: string

  /**
   * Amount to approve. Defaults to max int.
   */
  amount?: BigNumber

  /**
   * Transaction overrides
   */
  overrides?: PayableOverrides
}

export interface TokenDepositParams extends EthDepositBase {
  /**
   * L1 address of the token ERC20 contract
   */
  erc20L1Address: string

  /**
   * L2 address of the entity receiving the funds. Defaults to the l1FromAddress
   */
  destinationAddress?: string

  /**
   * Overrides for the retryable ticket parameters
   */
  retryableGasOverrides?: GasOverrides
}

export interface TokenWithdrawParams extends EthWithdrawParams {
  /**
   * L1 address of the token ERC20 contract
   */
  erc20l1Address: string
}

/**
 * Bridger for moving ERC20 tokens back and forth betwen L1 to L2
 */
export class TokenBridger extends AssetBridger<
  TokenDepositParams,
  TokenWithdrawParams
> {
  public static MAX_APPROVAL = MaxUint256
  public static MIN_CUSTOM_DEPOSIT_MAXGAS = BigNumber.from(275000)

  public constructor(
    l2Network: L2Network,
    public readonly isCustomNetwork: boolean = false
  ) {
    super(l2Network)
  }

  /**
   * Get the address of the l1 gateway for this token
   * @param erc20L1Address
   * @param l1Provider
   * @returns
   */
  public async getL1GatewayAddress(
    erc20L1Address: string,
    l1Provider: Provider
  ): Promise<string> {
    const l1GatewayRouter = L1GatewayRouter__factory.connect(
      this.l2Network.tokenBridge.l1GatewayRouter,
      l1Provider
    )

    return (await l1GatewayRouter.functions.getGateway(erc20L1Address)).gateway
  }

  /**
   * Get the address of the l2 gateway for this token
   * @param erc20L1Address
   * @param l1Provider
   * @returns
   */
  public async getL2GatewayAddress(
    // CHRIS: was this a mistake in the old code? should this be l2 erc20 addr?
    erc20L1Address: string,
    l2Provider: Provider
  ): Promise<string> {
    const l2GatewayRouter = L1GatewayRouter__factory.connect(
      this.l2Network.tokenBridge.l2GatewayRouter,
      l2Provider
    )

    return (await l2GatewayRouter.functions.getGateway(erc20L1Address)).gateway
  }

  // CHRIS: aren't we missing an approve on L2?
  /**
   * Approve tokens for deposit to the bridge. The tokens will be approved for the relevant gateway.
   * @param params
   * @returns
   */
  public async approveToken(params: TokenApproveParams) {
    if (!SignerProviderUtils.signerHasProvider(params.l1Signer)) {
      throw new ArbTsError(
        'l1Signer does not have a connected provider and one is required.'
      )
    }

    // you approve tokens to the gateway that the router will use
    const gatewayAddress = await this.getL1GatewayAddress(
      params.erc20L1Address,
      params.l1Signer.provider
    )
    const contract = await ERC20__factory.connect(
      params.erc20L1Address,
      params.l1Signer
    )
    return contract.functions.approve(
      gatewayAddress,
      params.amount || TokenBridger.MAX_APPROVAL,
      params.overrides || {}
    )
  }

  // public async getTokenWithdrawEventData(
  //   l2Provider: Provider,
  //   l1TokenAddress: string,
  //   fromAddress?: string,
  //   filter?: Filter
  // ): Promise<WithdrawalInitiated[]> {
  //   // CHRIS: network
  //   const tokenBridge = networks['1'].tokenBridge
  //   const l2GatewayRouter = L2GatewayRouter__factory.connect(
  //     tokenBridge.l2GatewayRouter,
  //     l2Provider
  //   )

  //   const gatewayAddress = await l2GatewayRouter.getGateway(l1TokenAddress)

  //   return this.getTWithdrawalEventData(
  //     l2Provider,
  //     gatewayAddress,
  //     l1TokenAddress,
  //     fromAddress,
  //     filter
  //   )
  // }

  // CHRIS: reconcile this method with getTokenWithdrawEventData
  /**
   * All withdrawals from given gateway
   */
  // public async getGatewayWithdrawEventData(
  //   l2Provider: Provider,
  //   gatewayAddress: string,
  //   fromAddress?: string,
  //   filter?: Filter
  // ): Promise<WithdrawalInitiated[]> {
  //   return this.getTWithdrawalEventData(
  //     l2Provider,
  //     gatewayAddress,
  //     undefined,
  //     fromAddress,
  //     filter
  //   )
  // }

  // // CHRIS: do we need functions for this? shoud we have an event log getter?
  // private async getEventLogs<T extends Interface>(
  //   eventName: string,
  //   provider: Provider,
  //   iface: T,
  //   contractAddress: string,
  //   topics: (string | string[] | null)[] = [],
  //   filter: Filter = {}
  // ): Promise<Log[]> {
  //   // TODO: can we make eventName typesafe?
  //   const event = iface.getEvent(eventName)
  //   const eventTopic = iface.getEventTopic(event)

  //   if (!filter.fromBlock && !filter.toBlock)
  //     console.warn('Attempting to query from 0 to block latest')
  //   return await provider.getLogs({
  //     address: contractAddress,
  //     topics: [eventTopic, ...topics],
  //     fromBlock: filter.fromBlock || 0,
  //     toBlock: filter.toBlock || 'latest',
  //   })
  // }

  // private async getLogs<TInterface extends ethers.utils.Interface, TEvent extends Readonly<any[]>>(
  //   provider: Provider,
  //   interfaceGenerator: () => TInterface,
  //   eventSelector: (i: TInterface) => ethers.utils.EventFragment,
  //   ,
  //   topics: TEvent,
  //   filter: Omit<Filter, "topics">,
  // ) {
  //   // static createInterface(): L2ArbitrumGatewayInterface {
  //   //   return new utils.Interface(_abi) as L2ArbitrumGatewayInterface
  //   // }
  //   // static connect(
  //   //   address: string,
  //   //   signerOrProvider: Signer | Provider
  //   // ): L2ArbitrumGateway {
  //   //   return new Contract(address, _abi, signerOrProvider) as L2ArbitrumGateway
  //   // }

  //   const iFace = interfaceGenerator();
  //   const eventFrag = eventSelector(iFace);

  //   const encodedTopics = iFace.encodeFilterTopics(eventFrag, topics);
  //   const res = await provider.getLogs({...filter, topics: encodedTopics })
  //   return res.map(l => {
  //     const parsed = iFace.parseLog(l).args
  //     return (parsed as unknown) as TEvent
  //   })
  // }

  // CHRIS: needs a better name, was called getTokenWithdrawEventData
  // CHRIS: these vent log functions are pretty yuck?
  // private async getTWithdrawalEventData(
  //   l2Provider: Provider,
  //   gatewayAddress: string,
  //   l1TokenAddress?: string,
  //   fromAddress?: string,
  //   filter?: Filter
  // ): Promise<WithdrawalInitiated[]> {
  //   // this.getLogs<L2ArbitrumGatewayInterface, WithdrawalInitiatedEvent>(prvider,
  //   //   () => L2ArbitrumGateway__factory.createInterface(),
  //   //   i => i.events['WithdrawalInitiated(address,address,address,uint256,uint256,uint256)'],
  //   //   {
  //   //     topics
  //   //   }

  //   // )

  //   const iface = L2ArbitrumGateway__factory.createInterface()

  //   const gatway = L2ArbitrumGateway__factory.connect(
  //     gatewayAddress,
  //     l2Provider
  //   )
  //   const a = gatway.filters.WithdrawalInitiated(
  //     null,
  //     fromAddress ? hexZeroPad(fromAddress, 32) : null
  //   )

  //   const frag =
  //     iface.events[
  //       'WithdrawalInitiated(address,address,address,uint256,uint256,uint256)'
  //     ]
  //     const topics = [null, fromAddress ? hexZeroPad(fromAddress, 32) : null]
  //   const f2 = iface.encodeFilterTopics(frag, topics)
  //   const r2 = await l2Provider.getLogs({...filter, ...f2})

  //   const res = await l2Provider.getLogs({ ...filter, ...a }) // should be "as something?"

  //   const logs = await this.getEventLogs(
  //     'WithdrawalInitiated',
  //     l2Provider,
  //     iface,
  //     gatewayAddress,
  //     topics,
  //     filter
  //   )
  //   res.map(l => {
  //     const parsed = gatway.interface.parseLog(l).args
  //     const q = (parsed as unknown) as WithdrawalInitiatedEvent
  //     return q
  //   })

  //   const parsedLogs = logs.map(log => {
  //     const data = {
  //       ...iface.parseLog(log).args,
  //       txHash: log.transactionHash,
  //     }
  //     return (data as unknown) as WithdrawalInitiated
  //   })
  //   // TODO: use l1TokenAddress as filter in topics instead of here
  //   return l1TokenAddress
  //     ? parsedLogs.filter(
  //         (log: WithdrawalInitiated) =>
  //           log.l1Token.toLocaleLowerCase() ===
  //           l1TokenAddress.toLocaleLowerCase()
  //       )
  //     : parsedLogs
  // }

  /**
   * Does the provided address look like a weth gateway
   * @param potentialWethGatewayAddress
   * @param l1Provider
   * @returns
   */
  private async isUnkownWethGateway(
    potentialWethGatewayAddress: string,
    l1Provider: Provider
  ) {
    try {
      const potentialWethGateway = L1WethGateway__factory.connect(
        potentialWethGatewayAddress,
        l1Provider
      )
      await potentialWethGateway.l1Weth()
      return true
    } catch (err) {
      if (
        err instanceof Error &&
        isError(err) &&
        err.code === Logger.errors.CALL_EXCEPTION
      ) {
        return false
      } else {
        throw err
      }
    }
  }

  /**
   * Is this a known or unknown WETH gateway
   * @param gatewayAddress
   * @param l1Provider
   * @returns
   */
  private async isWethGateway(
    gatewayAddress: string,
    l1Provider: Provider
  ): Promise<boolean> {
    const wethAddress = this.l2Network.tokenBridge.l1WethGateway
    if (this.isCustomNetwork) {
      // For custom network, we do an ad-hoc check to see if it's a WETH gateway
      if (await this.isUnkownWethGateway(gatewayAddress, l1Provider)) {
        return true
      }
      // ...otherwise we directly check it against the config file
    } else if (wethAddress === gatewayAddress) {
      return true
    }
    return false
  }

  /**
   * Get the corresponding L2 for the provided L1 token
   * @param erc20L1Address
   * @param l1Provider
   * @returns
   */
  public getERC20L2Address(erc20L1Address: string, l1Provider: Provider) {
    const l1GatewayRouter = L1GatewayRouter__factory.connect(
      this.l2Network.tokenBridge.l1GatewayRouter,
      l1Provider
    )

    return l1GatewayRouter.functions
      .calculateL2TokenAddress(erc20L1Address)
      .then(([res]) => res)
  }

  /**
   * Get the corresponding L1 for the provided L2 token
   * @param erc20L1Address
   * @param l1Provider
   * @returns
   */
  public getERC20L1Address(erc20L2Address: string, l2Provider: Provider) {
    const arbERC20 = StandardArbERC20__factory.connect(
      erc20L2Address,
      l2Provider
    )
    return arbERC20.functions.l1Address().then(([res]) => res)
  }

  private async getDepositParams(
    params: TokenDepositParams
  ): Promise<DepositParams> {
    // // CHRIS: '1' vs '0.05'? why? should we parameterise the gas estimator?
    // parseEther('0.05').add(
    //   estimateGasCallValue
    // ) /** we add a 0.05 "deposit" buffer to pay for execution in the gas estimation  */,

    const {
      erc20L1Address,
      amount,
      l2Provider,
      l1Signer,
      destinationAddress,
      retryableGasOverrides,
      overrides,
    } = params

    if (!SignerProviderUtils.signerHasProvider(l1Signer)) {
      throw new ArbTsError(
        'l1Signer does not have a connected provider and one is required.'
      )
    }

    // 1. get the params for a gas estimate
    const l1GatewayAddress = await this.getL1GatewayAddress(
      erc20L1Address,
      l1Signer.provider
    )
    const l1Gateway = L1ERC20Gateway__factory.connect(
      l1GatewayAddress,
      l1Signer.provider
    )
    const sender = await l1Signer.getAddress()
    const to = destinationAddress ? destinationAddress : sender
    const depositCalldata = await l1Gateway.getOutboundCalldata(
      erc20L1Address,
      sender,
      to,
      amount,
      '0x'
    )

    // The WETH gateway is the only deposit that requires callvalue in the L2 user-tx (i.e., the recently un-wrapped ETH)
    // Here we check if this is a WETH deposit, and include the callvalue for the gas estimate query if so
    const estimateGasCallValue = (await this.isWethGateway(
      l1GatewayAddress,
      l1Signer.provider
    ))
      ? amount
      : Zero

    const l2Dest = await l1Gateway.counterpartGateway()
    const gasEstimator = new L1ToL2MessageGasEstimator(l2Provider)

    // 2. get the gas estimates
    const estimates = await gasEstimator.estimateGasValuesL1ToL2Creation(
      l1GatewayAddress,
      l2Dest,
      depositCalldata,
      estimateGasCallValue,
      retryableGasOverrides
    )

    // 3. Some special token deposit defaults and overrides
    let maxGas = estimates.maxGasBid
    if (
      l1GatewayAddress === this.l2Network.tokenBridge.l1CustomGateway &&
      estimates.maxGasBid.lt(TokenBridger.MIN_CUSTOM_DEPOSIT_MAXGAS)
    ) {
      // For insurance, we set a sane minimum max gas for the custom gateway
      maxGas = TokenBridger.MIN_CUSTOM_DEPOSIT_MAXGAS
    }

    let totalEthCallvalueToSend = (overrides && (await overrides.value)) || Zero
    if (
      // CHRIS: I dont think this is correct for weth where we need to send value? ppl may populate the value param
      !totalEthCallvalueToSend ||
      BigNumber.from(totalEthCallvalueToSend).isZero()
    ) {
      totalEthCallvalueToSend = estimates.totalDepositValue
    }

    return {
      maxGas,
      maxSubmissionCost: estimates.maxSubmissionPriceBid,
      maxGasPrice: estimates.maxGasPriceBid,
      l1CallValue: BigNumber.from(totalEthCallvalueToSend),
      destinationAddress: to,
      amount,
      erc20L1Address,
    }
  }

  private async depositTxOrGas<T extends boolean>(
    params: TokenDepositParams,
    estimate: T
  ): Promise<T extends true ? BigNumber : ethers.ContractTransaction>
  private async depositTxOrGas<T extends boolean>(
    params: TokenDepositParams,
    estimate: T
  ): Promise<BigNumber | ethers.ContractTransaction> {
    if (!SignerProviderUtils.signerHasProvider(params.l1Signer)) {
      throw new ArbTsError(
        'l1Signer does not have a connected provider and one is required.'
      )
    }

    const depositParams = await this.getDepositParams(params)

    const data = defaultAbiCoder.encode(
      ['uint256', 'bytes'],
      [depositParams.maxSubmissionCost, '0x']
    )
    if (params.overrides?.value)
      throw new Error('L1 call value should be set through l1CallValue param')
    if (depositParams.l1CallValue.eq(0))
      throw new Error('L1 call value should not be zero')
    if (depositParams.maxSubmissionCost.eq(0))
      throw new Error('Max submission cost should not be zero')

    const l1GatewayRouter = L1GatewayRouter__factory.connect(
      this.l2Network.tokenBridge.l1GatewayRouter,
      params.l1Signer
    )

    return (estimate
      ? l1GatewayRouter.estimateGas
      : l1GatewayRouter.functions
    ).outboundTransfer(
      depositParams.erc20L1Address,
      depositParams.destinationAddress,
      depositParams.amount,
      depositParams.maxGas,
      depositParams.maxGasPrice,
      data,
      { ...(params.overrides || {}), value: depositParams.l1CallValue }
    )
  }

  /**
   * Estimate the gas required for a token deposit
   * @param params
   * @returns
   */
  public async depositEstimateGas(params: TokenDepositParams) {
    return await this.depositTxOrGas(params, true)
  }

  /**
   * Execute a token deposit from L1 to L2
   * @param params
   * @returns
   */
  public async deposit(params: TokenDepositParams) {
    return await this.depositTxOrGas(params, false)
  }

  private async withdrawTxOrGas<T extends boolean>(
    params: TokenWithdrawParams,
    estimate: T
  ): Promise<T extends true ? BigNumber : ethers.ContractTransaction>
  private async withdrawTxOrGas<T extends boolean>(
    params: TokenWithdrawParams,
    estimate: T
  ): Promise<BigNumber | ethers.ContractTransaction> {
    if (!SignerProviderUtils.signerHasProvider(params.l2Signer)) {
      throw new ArbTsError(
        'l2Signer does not have a connected provider and one is required.'
      )
    }

    const to = params.destinationAddress || (await params.l2Signer.getAddress())

    const l2GatewayRouter = L2GatewayRouter__factory.connect(
      this.l2Network.tokenBridge.l2GatewayRouter,
      params.l2Signer
    )

    return (estimate ? l2GatewayRouter.estimateGas : l2GatewayRouter.functions)[
      'outboundTransfer(address,address,uint256,bytes)'
    ](params.erc20l1Address, to, params.amount, '0x', params.overrides || {})
  }

  /**
   * Estimate gas for withdrawing tokens from L2 to L1
   * @param params
   * @returns
   */
  public async withdrawEstimateGas(params: TokenWithdrawParams) {
    return this.withdrawTxOrGas(params, true)
  }

  /**
   * Withdraw tokens from L2 to L1
   * @param params
   * @returns
   */
  public async withdraw(params: TokenWithdrawParams) {
    return this.withdrawTxOrGas(params, false)
  }

  // CHRIS: move this off here
  public async getTokenBalanceBatch(
    l1OrL2Provider: Provider,
    userAddr: string,
    tokenAddrs: Array<string>,
    targetNetwork: 'L1' | 'L2'
  ): Promise<Array<{ tokenAddr: string; balance: BigNumber | undefined }>> {
    const iface = ERC20__factory.createInterface()

    const balanceCalls = tokenAddrs.map(token => ({
      target: token,
      funcFragment: iface.functions['balanceOf(address)'],
      values: [userAddr],
    }))

    type Await<T> = T extends {
      then(onfulfilled?: (value: infer U) => unknown): unknown
    }
      ? U
      : T
    type ExpectedReturnType = Await<ReturnType<ERC20['functions']['balanceOf']>>

    const tokenBridge = this.l2Network.tokenBridge
    const multicall = Multicall2__factory.connect(
      targetNetwork === 'L1'
        ? tokenBridge.l1MultiCall
        : tokenBridge.l2Multicall,
      l1OrL2Provider
    )

    const res = await BridgeHelper.getMulticallTryAggregate(
      balanceCalls,
      multicall
    )
    return res.map((bal, index) => ({
      tokenAddr: tokenAddrs[index],
      balance: bal ? (bal as ExpectedReturnType)[0] : undefined,
    }))
  }
}
