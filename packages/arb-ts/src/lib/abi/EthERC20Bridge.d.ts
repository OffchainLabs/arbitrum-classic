/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import {
  ethers,
  EventFilter,
  Signer,
  BigNumber,
  BigNumberish,
  PopulatedTransaction,
} from 'ethers'
import {
  Contract,
  ContractTransaction,
  Overrides,
  PayableOverrides,
  CallOverrides,
} from '@ethersproject/contracts'
import { BytesLike } from '@ethersproject/bytes'
import { Listener, Provider } from '@ethersproject/providers'
import { FunctionFragment, EventFragment, Result } from '@ethersproject/abi'

interface EthERC20BridgeInterface extends ethers.utils.Interface {
  functions: {
    'codeHash()': FunctionFragment
    'customL2Tokens(address)': FunctionFragment
    'depositAsCustomToken(address,address,uint256,uint256,uint256)': FunctionFragment
    'depositAsERC20(address,address,uint256,uint256,uint256)': FunctionFragment
    'depositAsERC777(address,address,uint256,uint256,uint256)': FunctionFragment
    'fastWithdrawalFromL2(address,bytes,address,uint256,uint256)': FunctionFragment
    'finalizeBuddyDeploy(bool)': FunctionFragment
    'inbox()': FunctionFragment
    'initiateBuddyDeploy(uint256,uint256,bytes)': FunctionFragment
    'l2Buddy()': FunctionFragment
    'l2Connection()': FunctionFragment
    'l2Deployer()': FunctionFragment
    'notifyCustomToken(address,uint256,uint256)': FunctionFragment
    'registerCustomL2Token(address)': FunctionFragment
    'updateTokenInfo(address,uint256,uint256,bool)': FunctionFragment
    'withdrawFromL2(uint256,address,address,uint256)': FunctionFragment
  }

  encodeFunctionData(functionFragment: 'codeHash', values?: undefined): string
  encodeFunctionData(
    functionFragment: 'customL2Tokens',
    values: [string]
  ): string
  encodeFunctionData(
    functionFragment: 'depositAsCustomToken',
    values: [string, string, BigNumberish, BigNumberish, BigNumberish]
  ): string
  encodeFunctionData(
    functionFragment: 'depositAsERC20',
    values: [string, string, BigNumberish, BigNumberish, BigNumberish]
  ): string
  encodeFunctionData(
    functionFragment: 'depositAsERC777',
    values: [string, string, BigNumberish, BigNumberish, BigNumberish]
  ): string
  encodeFunctionData(
    functionFragment: 'fastWithdrawalFromL2',
    values: [string, BytesLike, string, BigNumberish, BigNumberish]
  ): string
  encodeFunctionData(
    functionFragment: 'finalizeBuddyDeploy',
    values: [boolean]
  ): string
  encodeFunctionData(functionFragment: 'inbox', values?: undefined): string
  encodeFunctionData(
    functionFragment: 'initiateBuddyDeploy',
    values: [BigNumberish, BigNumberish, BytesLike]
  ): string
  encodeFunctionData(functionFragment: 'l2Buddy', values?: undefined): string
  encodeFunctionData(
    functionFragment: 'l2Connection',
    values?: undefined
  ): string
  encodeFunctionData(functionFragment: 'l2Deployer', values?: undefined): string
  encodeFunctionData(
    functionFragment: 'notifyCustomToken',
    values: [string, BigNumberish, BigNumberish]
  ): string
  encodeFunctionData(
    functionFragment: 'registerCustomL2Token',
    values: [string]
  ): string
  encodeFunctionData(
    functionFragment: 'updateTokenInfo',
    values: [string, BigNumberish, BigNumberish, boolean]
  ): string
  encodeFunctionData(
    functionFragment: 'withdrawFromL2',
    values: [BigNumberish, string, string, BigNumberish]
  ): string

  decodeFunctionResult(functionFragment: 'codeHash', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'customL2Tokens',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'depositAsCustomToken',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'depositAsERC20',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'depositAsERC777',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'fastWithdrawalFromL2',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'finalizeBuddyDeploy',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'inbox', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'initiateBuddyDeploy',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'l2Buddy', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'l2Connection',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'l2Deployer', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'notifyCustomToken',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'registerCustomL2Token',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'updateTokenInfo',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'withdrawFromL2',
    data: BytesLike
  ): Result

  events: {}
}

export class EthERC20Bridge extends Contract {
  connect(signerOrProvider: Signer | Provider | string): this
  attach(addressOrName: string): this
  deployed(): Promise<this>

  on(event: EventFilter | string, listener: Listener): this
  once(event: EventFilter | string, listener: Listener): this
  addListener(eventName: EventFilter | string, listener: Listener): this
  removeAllListeners(eventName: EventFilter | string): this
  removeListener(eventName: any, listener: Listener): this

  interface: EthERC20BridgeInterface

  functions: {
    codeHash(overrides?: CallOverrides): Promise<[string]>

    'codeHash()'(overrides?: CallOverrides): Promise<[string]>

    customL2Tokens(arg0: string, overrides?: CallOverrides): Promise<[string]>

    'customL2Tokens(address)'(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<[string]>

    depositAsCustomToken(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<ContractTransaction>

    'depositAsCustomToken(address,address,uint256,uint256,uint256)'(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<ContractTransaction>

    depositAsERC20(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<ContractTransaction>

    'depositAsERC20(address,address,uint256,uint256,uint256)'(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<ContractTransaction>

    depositAsERC777(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<ContractTransaction>

    'depositAsERC777(address,address,uint256,uint256,uint256)'(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<ContractTransaction>

    fastWithdrawalFromL2(
      liquidityProvider: string,
      liquidityProof: BytesLike,
      erc20: string,
      amount: BigNumberish,
      exitNum: BigNumberish,
      overrides?: Overrides
    ): Promise<ContractTransaction>

    'fastWithdrawalFromL2(address,bytes,address,uint256,uint256)'(
      liquidityProvider: string,
      liquidityProof: BytesLike,
      erc20: string,
      amount: BigNumberish,
      exitNum: BigNumberish,
      overrides?: Overrides
    ): Promise<ContractTransaction>

    finalizeBuddyDeploy(
      success: boolean,
      overrides?: Overrides
    ): Promise<ContractTransaction>

    'finalizeBuddyDeploy(bool)'(
      success: boolean,
      overrides?: Overrides
    ): Promise<ContractTransaction>

    inbox(overrides?: CallOverrides): Promise<[string]>

    'inbox()'(overrides?: CallOverrides): Promise<[string]>

    initiateBuddyDeploy(
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      contractInitCode: BytesLike,
      overrides?: PayableOverrides
    ): Promise<ContractTransaction>

    'initiateBuddyDeploy(uint256,uint256,bytes)'(
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      contractInitCode: BytesLike,
      overrides?: PayableOverrides
    ): Promise<ContractTransaction>

    l2Buddy(overrides?: CallOverrides): Promise<[string]>

    'l2Buddy()'(overrides?: CallOverrides): Promise<[string]>

    l2Connection(overrides?: CallOverrides): Promise<[number]>

    'l2Connection()'(overrides?: CallOverrides): Promise<[number]>

    l2Deployer(overrides?: CallOverrides): Promise<[string]>

    'l2Deployer()'(overrides?: CallOverrides): Promise<[string]>

    notifyCustomToken(
      l1Address: string,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<ContractTransaction>

    'notifyCustomToken(address,uint256,uint256)'(
      l1Address: string,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<ContractTransaction>

    registerCustomL2Token(
      l2Address: string,
      overrides?: Overrides
    ): Promise<ContractTransaction>

    'registerCustomL2Token(address)'(
      l2Address: string,
      overrides?: Overrides
    ): Promise<ContractTransaction>

    updateTokenInfo(
      erc20: string,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      isERC20: boolean,
      overrides?: PayableOverrides
    ): Promise<ContractTransaction>

    'updateTokenInfo(address,uint256,uint256,bool)'(
      erc20: string,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      isERC20: boolean,
      overrides?: PayableOverrides
    ): Promise<ContractTransaction>

    withdrawFromL2(
      exitNum: BigNumberish,
      erc20: string,
      destination: string,
      amount: BigNumberish,
      overrides?: Overrides
    ): Promise<ContractTransaction>

    'withdrawFromL2(uint256,address,address,uint256)'(
      exitNum: BigNumberish,
      erc20: string,
      destination: string,
      amount: BigNumberish,
      overrides?: Overrides
    ): Promise<ContractTransaction>
  }

  codeHash(overrides?: CallOverrides): Promise<string>

  'codeHash()'(overrides?: CallOverrides): Promise<string>

  customL2Tokens(arg0: string, overrides?: CallOverrides): Promise<string>

  'customL2Tokens(address)'(
    arg0: string,
    overrides?: CallOverrides
  ): Promise<string>

  depositAsCustomToken(
    erc20: string,
    destination: string,
    amount: BigNumberish,
    maxGas: BigNumberish,
    gasPriceBid: BigNumberish,
    overrides?: PayableOverrides
  ): Promise<ContractTransaction>

  'depositAsCustomToken(address,address,uint256,uint256,uint256)'(
    erc20: string,
    destination: string,
    amount: BigNumberish,
    maxGas: BigNumberish,
    gasPriceBid: BigNumberish,
    overrides?: PayableOverrides
  ): Promise<ContractTransaction>

  depositAsERC20(
    erc20: string,
    destination: string,
    amount: BigNumberish,
    maxGas: BigNumberish,
    gasPriceBid: BigNumberish,
    overrides?: PayableOverrides
  ): Promise<ContractTransaction>

  'depositAsERC20(address,address,uint256,uint256,uint256)'(
    erc20: string,
    destination: string,
    amount: BigNumberish,
    maxGas: BigNumberish,
    gasPriceBid: BigNumberish,
    overrides?: PayableOverrides
  ): Promise<ContractTransaction>

  depositAsERC777(
    erc20: string,
    destination: string,
    amount: BigNumberish,
    maxGas: BigNumberish,
    gasPriceBid: BigNumberish,
    overrides?: PayableOverrides
  ): Promise<ContractTransaction>

  'depositAsERC777(address,address,uint256,uint256,uint256)'(
    erc20: string,
    destination: string,
    amount: BigNumberish,
    maxGas: BigNumberish,
    gasPriceBid: BigNumberish,
    overrides?: PayableOverrides
  ): Promise<ContractTransaction>

  fastWithdrawalFromL2(
    liquidityProvider: string,
    liquidityProof: BytesLike,
    erc20: string,
    amount: BigNumberish,
    exitNum: BigNumberish,
    overrides?: Overrides
  ): Promise<ContractTransaction>

  'fastWithdrawalFromL2(address,bytes,address,uint256,uint256)'(
    liquidityProvider: string,
    liquidityProof: BytesLike,
    erc20: string,
    amount: BigNumberish,
    exitNum: BigNumberish,
    overrides?: Overrides
  ): Promise<ContractTransaction>

  finalizeBuddyDeploy(
    success: boolean,
    overrides?: Overrides
  ): Promise<ContractTransaction>

  'finalizeBuddyDeploy(bool)'(
    success: boolean,
    overrides?: Overrides
  ): Promise<ContractTransaction>

  inbox(overrides?: CallOverrides): Promise<string>

  'inbox()'(overrides?: CallOverrides): Promise<string>

  initiateBuddyDeploy(
    maxGas: BigNumberish,
    gasPriceBid: BigNumberish,
    contractInitCode: BytesLike,
    overrides?: PayableOverrides
  ): Promise<ContractTransaction>

  'initiateBuddyDeploy(uint256,uint256,bytes)'(
    maxGas: BigNumberish,
    gasPriceBid: BigNumberish,
    contractInitCode: BytesLike,
    overrides?: PayableOverrides
  ): Promise<ContractTransaction>

  l2Buddy(overrides?: CallOverrides): Promise<string>

  'l2Buddy()'(overrides?: CallOverrides): Promise<string>

  l2Connection(overrides?: CallOverrides): Promise<number>

  'l2Connection()'(overrides?: CallOverrides): Promise<number>

  l2Deployer(overrides?: CallOverrides): Promise<string>

  'l2Deployer()'(overrides?: CallOverrides): Promise<string>

  notifyCustomToken(
    l1Address: string,
    maxGas: BigNumberish,
    gasPriceBid: BigNumberish,
    overrides?: PayableOverrides
  ): Promise<ContractTransaction>

  'notifyCustomToken(address,uint256,uint256)'(
    l1Address: string,
    maxGas: BigNumberish,
    gasPriceBid: BigNumberish,
    overrides?: PayableOverrides
  ): Promise<ContractTransaction>

  registerCustomL2Token(
    l2Address: string,
    overrides?: Overrides
  ): Promise<ContractTransaction>

  'registerCustomL2Token(address)'(
    l2Address: string,
    overrides?: Overrides
  ): Promise<ContractTransaction>

  updateTokenInfo(
    erc20: string,
    maxGas: BigNumberish,
    gasPriceBid: BigNumberish,
    isERC20: boolean,
    overrides?: PayableOverrides
  ): Promise<ContractTransaction>

  'updateTokenInfo(address,uint256,uint256,bool)'(
    erc20: string,
    maxGas: BigNumberish,
    gasPriceBid: BigNumberish,
    isERC20: boolean,
    overrides?: PayableOverrides
  ): Promise<ContractTransaction>

  withdrawFromL2(
    exitNum: BigNumberish,
    erc20: string,
    destination: string,
    amount: BigNumberish,
    overrides?: Overrides
  ): Promise<ContractTransaction>

  'withdrawFromL2(uint256,address,address,uint256)'(
    exitNum: BigNumberish,
    erc20: string,
    destination: string,
    amount: BigNumberish,
    overrides?: Overrides
  ): Promise<ContractTransaction>

  callStatic: {
    codeHash(overrides?: CallOverrides): Promise<string>

    'codeHash()'(overrides?: CallOverrides): Promise<string>

    customL2Tokens(arg0: string, overrides?: CallOverrides): Promise<string>

    'customL2Tokens(address)'(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<string>

    depositAsCustomToken(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: CallOverrides
    ): Promise<void>

    'depositAsCustomToken(address,address,uint256,uint256,uint256)'(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: CallOverrides
    ): Promise<void>

    depositAsERC20(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: CallOverrides
    ): Promise<void>

    'depositAsERC20(address,address,uint256,uint256,uint256)'(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: CallOverrides
    ): Promise<void>

    depositAsERC777(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: CallOverrides
    ): Promise<void>

    'depositAsERC777(address,address,uint256,uint256,uint256)'(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: CallOverrides
    ): Promise<void>

    fastWithdrawalFromL2(
      liquidityProvider: string,
      liquidityProof: BytesLike,
      erc20: string,
      amount: BigNumberish,
      exitNum: BigNumberish,
      overrides?: CallOverrides
    ): Promise<void>

    'fastWithdrawalFromL2(address,bytes,address,uint256,uint256)'(
      liquidityProvider: string,
      liquidityProof: BytesLike,
      erc20: string,
      amount: BigNumberish,
      exitNum: BigNumberish,
      overrides?: CallOverrides
    ): Promise<void>

    finalizeBuddyDeploy(
      success: boolean,
      overrides?: CallOverrides
    ): Promise<void>

    'finalizeBuddyDeploy(bool)'(
      success: boolean,
      overrides?: CallOverrides
    ): Promise<void>

    inbox(overrides?: CallOverrides): Promise<string>

    'inbox()'(overrides?: CallOverrides): Promise<string>

    initiateBuddyDeploy(
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      contractInitCode: BytesLike,
      overrides?: CallOverrides
    ): Promise<void>

    'initiateBuddyDeploy(uint256,uint256,bytes)'(
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      contractInitCode: BytesLike,
      overrides?: CallOverrides
    ): Promise<void>

    l2Buddy(overrides?: CallOverrides): Promise<string>

    'l2Buddy()'(overrides?: CallOverrides): Promise<string>

    l2Connection(overrides?: CallOverrides): Promise<number>

    'l2Connection()'(overrides?: CallOverrides): Promise<number>

    l2Deployer(overrides?: CallOverrides): Promise<string>

    'l2Deployer()'(overrides?: CallOverrides): Promise<string>

    notifyCustomToken(
      l1Address: string,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: CallOverrides
    ): Promise<void>

    'notifyCustomToken(address,uint256,uint256)'(
      l1Address: string,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: CallOverrides
    ): Promise<void>

    registerCustomL2Token(
      l2Address: string,
      overrides?: CallOverrides
    ): Promise<void>

    'registerCustomL2Token(address)'(
      l2Address: string,
      overrides?: CallOverrides
    ): Promise<void>

    updateTokenInfo(
      erc20: string,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      isERC20: boolean,
      overrides?: CallOverrides
    ): Promise<void>

    'updateTokenInfo(address,uint256,uint256,bool)'(
      erc20: string,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      isERC20: boolean,
      overrides?: CallOverrides
    ): Promise<void>

    withdrawFromL2(
      exitNum: BigNumberish,
      erc20: string,
      destination: string,
      amount: BigNumberish,
      overrides?: CallOverrides
    ): Promise<void>

    'withdrawFromL2(uint256,address,address,uint256)'(
      exitNum: BigNumberish,
      erc20: string,
      destination: string,
      amount: BigNumberish,
      overrides?: CallOverrides
    ): Promise<void>
  }

  filters: {}

  estimateGas: {
    codeHash(overrides?: CallOverrides): Promise<BigNumber>

    'codeHash()'(overrides?: CallOverrides): Promise<BigNumber>

    customL2Tokens(arg0: string, overrides?: CallOverrides): Promise<BigNumber>

    'customL2Tokens(address)'(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    depositAsCustomToken(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<BigNumber>

    'depositAsCustomToken(address,address,uint256,uint256,uint256)'(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<BigNumber>

    depositAsERC20(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<BigNumber>

    'depositAsERC20(address,address,uint256,uint256,uint256)'(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<BigNumber>

    depositAsERC777(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<BigNumber>

    'depositAsERC777(address,address,uint256,uint256,uint256)'(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<BigNumber>

    fastWithdrawalFromL2(
      liquidityProvider: string,
      liquidityProof: BytesLike,
      erc20: string,
      amount: BigNumberish,
      exitNum: BigNumberish,
      overrides?: Overrides
    ): Promise<BigNumber>

    'fastWithdrawalFromL2(address,bytes,address,uint256,uint256)'(
      liquidityProvider: string,
      liquidityProof: BytesLike,
      erc20: string,
      amount: BigNumberish,
      exitNum: BigNumberish,
      overrides?: Overrides
    ): Promise<BigNumber>

    finalizeBuddyDeploy(
      success: boolean,
      overrides?: Overrides
    ): Promise<BigNumber>

    'finalizeBuddyDeploy(bool)'(
      success: boolean,
      overrides?: Overrides
    ): Promise<BigNumber>

    inbox(overrides?: CallOverrides): Promise<BigNumber>

    'inbox()'(overrides?: CallOverrides): Promise<BigNumber>

    initiateBuddyDeploy(
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      contractInitCode: BytesLike,
      overrides?: PayableOverrides
    ): Promise<BigNumber>

    'initiateBuddyDeploy(uint256,uint256,bytes)'(
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      contractInitCode: BytesLike,
      overrides?: PayableOverrides
    ): Promise<BigNumber>

    l2Buddy(overrides?: CallOverrides): Promise<BigNumber>

    'l2Buddy()'(overrides?: CallOverrides): Promise<BigNumber>

    l2Connection(overrides?: CallOverrides): Promise<BigNumber>

    'l2Connection()'(overrides?: CallOverrides): Promise<BigNumber>

    l2Deployer(overrides?: CallOverrides): Promise<BigNumber>

    'l2Deployer()'(overrides?: CallOverrides): Promise<BigNumber>

    notifyCustomToken(
      l1Address: string,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<BigNumber>

    'notifyCustomToken(address,uint256,uint256)'(
      l1Address: string,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<BigNumber>

    registerCustomL2Token(
      l2Address: string,
      overrides?: Overrides
    ): Promise<BigNumber>

    'registerCustomL2Token(address)'(
      l2Address: string,
      overrides?: Overrides
    ): Promise<BigNumber>

    updateTokenInfo(
      erc20: string,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      isERC20: boolean,
      overrides?: PayableOverrides
    ): Promise<BigNumber>

    'updateTokenInfo(address,uint256,uint256,bool)'(
      erc20: string,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      isERC20: boolean,
      overrides?: PayableOverrides
    ): Promise<BigNumber>

    withdrawFromL2(
      exitNum: BigNumberish,
      erc20: string,
      destination: string,
      amount: BigNumberish,
      overrides?: Overrides
    ): Promise<BigNumber>

    'withdrawFromL2(uint256,address,address,uint256)'(
      exitNum: BigNumberish,
      erc20: string,
      destination: string,
      amount: BigNumberish,
      overrides?: Overrides
    ): Promise<BigNumber>
  }

  populateTransaction: {
    codeHash(overrides?: CallOverrides): Promise<PopulatedTransaction>

    'codeHash()'(overrides?: CallOverrides): Promise<PopulatedTransaction>

    customL2Tokens(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    'customL2Tokens(address)'(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    depositAsCustomToken(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<PopulatedTransaction>

    'depositAsCustomToken(address,address,uint256,uint256,uint256)'(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<PopulatedTransaction>

    depositAsERC20(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<PopulatedTransaction>

    'depositAsERC20(address,address,uint256,uint256,uint256)'(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<PopulatedTransaction>

    depositAsERC777(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<PopulatedTransaction>

    'depositAsERC777(address,address,uint256,uint256,uint256)'(
      erc20: string,
      destination: string,
      amount: BigNumberish,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<PopulatedTransaction>

    fastWithdrawalFromL2(
      liquidityProvider: string,
      liquidityProof: BytesLike,
      erc20: string,
      amount: BigNumberish,
      exitNum: BigNumberish,
      overrides?: Overrides
    ): Promise<PopulatedTransaction>

    'fastWithdrawalFromL2(address,bytes,address,uint256,uint256)'(
      liquidityProvider: string,
      liquidityProof: BytesLike,
      erc20: string,
      amount: BigNumberish,
      exitNum: BigNumberish,
      overrides?: Overrides
    ): Promise<PopulatedTransaction>

    finalizeBuddyDeploy(
      success: boolean,
      overrides?: Overrides
    ): Promise<PopulatedTransaction>

    'finalizeBuddyDeploy(bool)'(
      success: boolean,
      overrides?: Overrides
    ): Promise<PopulatedTransaction>

    inbox(overrides?: CallOverrides): Promise<PopulatedTransaction>

    'inbox()'(overrides?: CallOverrides): Promise<PopulatedTransaction>

    initiateBuddyDeploy(
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      contractInitCode: BytesLike,
      overrides?: PayableOverrides
    ): Promise<PopulatedTransaction>

    'initiateBuddyDeploy(uint256,uint256,bytes)'(
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      contractInitCode: BytesLike,
      overrides?: PayableOverrides
    ): Promise<PopulatedTransaction>

    l2Buddy(overrides?: CallOverrides): Promise<PopulatedTransaction>

    'l2Buddy()'(overrides?: CallOverrides): Promise<PopulatedTransaction>

    l2Connection(overrides?: CallOverrides): Promise<PopulatedTransaction>

    'l2Connection()'(overrides?: CallOverrides): Promise<PopulatedTransaction>

    l2Deployer(overrides?: CallOverrides): Promise<PopulatedTransaction>

    'l2Deployer()'(overrides?: CallOverrides): Promise<PopulatedTransaction>

    notifyCustomToken(
      l1Address: string,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<PopulatedTransaction>

    'notifyCustomToken(address,uint256,uint256)'(
      l1Address: string,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      overrides?: PayableOverrides
    ): Promise<PopulatedTransaction>

    registerCustomL2Token(
      l2Address: string,
      overrides?: Overrides
    ): Promise<PopulatedTransaction>

    'registerCustomL2Token(address)'(
      l2Address: string,
      overrides?: Overrides
    ): Promise<PopulatedTransaction>

    updateTokenInfo(
      erc20: string,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      isERC20: boolean,
      overrides?: PayableOverrides
    ): Promise<PopulatedTransaction>

    'updateTokenInfo(address,uint256,uint256,bool)'(
      erc20: string,
      maxGas: BigNumberish,
      gasPriceBid: BigNumberish,
      isERC20: boolean,
      overrides?: PayableOverrides
    ): Promise<PopulatedTransaction>

    withdrawFromL2(
      exitNum: BigNumberish,
      erc20: string,
      destination: string,
      amount: BigNumberish,
      overrides?: Overrides
    ): Promise<PopulatedTransaction>

    'withdrawFromL2(uint256,address,address,uint256)'(
      exitNum: BigNumberish,
      erc20: string,
      destination: string,
      amount: BigNumberish,
      overrides?: Overrides
    ): Promise<PopulatedTransaction>
  }
}
