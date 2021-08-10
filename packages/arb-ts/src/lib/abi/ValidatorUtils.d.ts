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
  BaseContract,
  ContractTransaction,
  CallOverrides,
} from 'ethers'
import { BytesLike } from '@ethersproject/bytes'
import { Listener, Provider } from '@ethersproject/providers'
import { FunctionFragment, EventFragment, Result } from '@ethersproject/abi'
import { TypedEventFilter, TypedEvent, TypedListener } from './commons'

interface ValidatorUtilsInterface extends ethers.utils.Interface {
  functions: {
    'areUnresolvedNodesLinear(address)': FunctionFragment
    'checkDecidableNextNode(address)': FunctionFragment
    'findNodeConflict(address,uint256,uint256,uint256)': FunctionFragment
    'findStakerConflict(address,address,address,uint256)': FunctionFragment
    'getConfig(address)': FunctionFragment
    'getStakers(address,uint256,uint256)': FunctionFragment
    'latestStaked(address,address)': FunctionFragment
    'refundableStakers(address)': FunctionFragment
    'requireConfirmable(address)': FunctionFragment
    'requireRejectable(address)': FunctionFragment
    'stakedNodes(address,address)': FunctionFragment
    'stakerInfo(address,address)': FunctionFragment
    'timedOutChallenges(address,uint256,uint256)': FunctionFragment
  }

  encodeFunctionData(
    functionFragment: 'areUnresolvedNodesLinear',
    values: [string]
  ): string
  encodeFunctionData(
    functionFragment: 'checkDecidableNextNode',
    values: [string]
  ): string
  encodeFunctionData(
    functionFragment: 'findNodeConflict',
    values: [string, BigNumberish, BigNumberish, BigNumberish]
  ): string
  encodeFunctionData(
    functionFragment: 'findStakerConflict',
    values: [string, string, string, BigNumberish]
  ): string
  encodeFunctionData(functionFragment: 'getConfig', values: [string]): string
  encodeFunctionData(
    functionFragment: 'getStakers',
    values: [string, BigNumberish, BigNumberish]
  ): string
  encodeFunctionData(
    functionFragment: 'latestStaked',
    values: [string, string]
  ): string
  encodeFunctionData(
    functionFragment: 'refundableStakers',
    values: [string]
  ): string
  encodeFunctionData(
    functionFragment: 'requireConfirmable',
    values: [string]
  ): string
  encodeFunctionData(
    functionFragment: 'requireRejectable',
    values: [string]
  ): string
  encodeFunctionData(
    functionFragment: 'stakedNodes',
    values: [string, string]
  ): string
  encodeFunctionData(
    functionFragment: 'stakerInfo',
    values: [string, string]
  ): string
  encodeFunctionData(
    functionFragment: 'timedOutChallenges',
    values: [string, BigNumberish, BigNumberish]
  ): string

  decodeFunctionResult(
    functionFragment: 'areUnresolvedNodesLinear',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'checkDecidableNextNode',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'findNodeConflict',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'findStakerConflict',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'getConfig', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'getStakers', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'latestStaked',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'refundableStakers',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'requireConfirmable',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'requireRejectable',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'stakedNodes', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'stakerInfo', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'timedOutChallenges',
    data: BytesLike
  ): Result

  events: {}
}

export class ValidatorUtils extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this
  attach(addressOrName: string): this
  deployed(): Promise<this>

  listeners<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter?: TypedEventFilter<EventArgsArray, EventArgsObject>
  ): Array<TypedListener<EventArgsArray, EventArgsObject>>
  off<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>,
    listener: TypedListener<EventArgsArray, EventArgsObject>
  ): this
  on<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>,
    listener: TypedListener<EventArgsArray, EventArgsObject>
  ): this
  once<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>,
    listener: TypedListener<EventArgsArray, EventArgsObject>
  ): this
  removeListener<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>,
    listener: TypedListener<EventArgsArray, EventArgsObject>
  ): this
  removeAllListeners<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>
  ): this

  listeners(eventName?: string): Array<Listener>
  off(eventName: string, listener: Listener): this
  on(eventName: string, listener: Listener): this
  once(eventName: string, listener: Listener): this
  removeListener(eventName: string, listener: Listener): this
  removeAllListeners(eventName?: string): this

  queryFilter<EventArgsArray extends Array<any>, EventArgsObject>(
    event: TypedEventFilter<EventArgsArray, EventArgsObject>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEvent<EventArgsArray & EventArgsObject>>>

  interface: ValidatorUtilsInterface

  functions: {
    areUnresolvedNodesLinear(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<[boolean]>

    checkDecidableNextNode(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<[number]>

    findNodeConflict(
      rollup: string,
      node1: BigNumberish,
      node2: BigNumberish,
      maxDepth: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[number, BigNumber, BigNumber]>

    findStakerConflict(
      rollup: string,
      staker1: string,
      staker2: string,
      maxDepth: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[number, BigNumber, BigNumber]>

    getConfig(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<
      [BigNumber, BigNumber, BigNumber, BigNumber] & {
        confirmPeriodBlocks: BigNumber
        extraChallengeTimeBlocks: BigNumber
        arbGasSpeedLimitPerBlock: BigNumber
        baseStake: BigNumber
      }
    >

    getStakers(
      rollup: string,
      startIndex: BigNumberish,
      max: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[string[], boolean] & { hasMore: boolean }>

    latestStaked(
      rollup: string,
      staker: string,
      overrides?: CallOverrides
    ): Promise<[BigNumber, string]>

    refundableStakers(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<[string[]]>

    requireConfirmable(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<[void]>

    requireRejectable(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<[boolean]>

    stakedNodes(
      rollup: string,
      staker: string,
      overrides?: CallOverrides
    ): Promise<[BigNumber[]]>

    stakerInfo(
      rollup: string,
      stakerAddress: string,
      overrides?: CallOverrides
    ): Promise<
      [boolean, BigNumber, BigNumber, string] & {
        isStaked: boolean
        latestStakedNode: BigNumber
        amountStaked: BigNumber
        currentChallenge: string
      }
    >

    timedOutChallenges(
      rollup: string,
      startIndex: BigNumberish,
      max: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[string[], boolean] & { hasMore: boolean }>
  }

  areUnresolvedNodesLinear(
    rollup: string,
    overrides?: CallOverrides
  ): Promise<boolean>

  checkDecidableNextNode(
    rollup: string,
    overrides?: CallOverrides
  ): Promise<number>

  findNodeConflict(
    rollup: string,
    node1: BigNumberish,
    node2: BigNumberish,
    maxDepth: BigNumberish,
    overrides?: CallOverrides
  ): Promise<[number, BigNumber, BigNumber]>

  findStakerConflict(
    rollup: string,
    staker1: string,
    staker2: string,
    maxDepth: BigNumberish,
    overrides?: CallOverrides
  ): Promise<[number, BigNumber, BigNumber]>

  getConfig(
    rollup: string,
    overrides?: CallOverrides
  ): Promise<
    [BigNumber, BigNumber, BigNumber, BigNumber] & {
      confirmPeriodBlocks: BigNumber
      extraChallengeTimeBlocks: BigNumber
      arbGasSpeedLimitPerBlock: BigNumber
      baseStake: BigNumber
    }
  >

  getStakers(
    rollup: string,
    startIndex: BigNumberish,
    max: BigNumberish,
    overrides?: CallOverrides
  ): Promise<[string[], boolean] & { hasMore: boolean }>

  latestStaked(
    rollup: string,
    staker: string,
    overrides?: CallOverrides
  ): Promise<[BigNumber, string]>

  refundableStakers(
    rollup: string,
    overrides?: CallOverrides
  ): Promise<string[]>

  requireConfirmable(rollup: string, overrides?: CallOverrides): Promise<void>

  requireRejectable(rollup: string, overrides?: CallOverrides): Promise<boolean>

  stakedNodes(
    rollup: string,
    staker: string,
    overrides?: CallOverrides
  ): Promise<BigNumber[]>

  stakerInfo(
    rollup: string,
    stakerAddress: string,
    overrides?: CallOverrides
  ): Promise<
    [boolean, BigNumber, BigNumber, string] & {
      isStaked: boolean
      latestStakedNode: BigNumber
      amountStaked: BigNumber
      currentChallenge: string
    }
  >

  timedOutChallenges(
    rollup: string,
    startIndex: BigNumberish,
    max: BigNumberish,
    overrides?: CallOverrides
  ): Promise<[string[], boolean] & { hasMore: boolean }>

  callStatic: {
    areUnresolvedNodesLinear(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<boolean>

    checkDecidableNextNode(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<number>

    findNodeConflict(
      rollup: string,
      node1: BigNumberish,
      node2: BigNumberish,
      maxDepth: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[number, BigNumber, BigNumber]>

    findStakerConflict(
      rollup: string,
      staker1: string,
      staker2: string,
      maxDepth: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[number, BigNumber, BigNumber]>

    getConfig(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<
      [BigNumber, BigNumber, BigNumber, BigNumber] & {
        confirmPeriodBlocks: BigNumber
        extraChallengeTimeBlocks: BigNumber
        arbGasSpeedLimitPerBlock: BigNumber
        baseStake: BigNumber
      }
    >

    getStakers(
      rollup: string,
      startIndex: BigNumberish,
      max: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[string[], boolean] & { hasMore: boolean }>

    latestStaked(
      rollup: string,
      staker: string,
      overrides?: CallOverrides
    ): Promise<[BigNumber, string]>

    refundableStakers(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<string[]>

    requireConfirmable(rollup: string, overrides?: CallOverrides): Promise<void>

    requireRejectable(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<boolean>

    stakedNodes(
      rollup: string,
      staker: string,
      overrides?: CallOverrides
    ): Promise<BigNumber[]>

    stakerInfo(
      rollup: string,
      stakerAddress: string,
      overrides?: CallOverrides
    ): Promise<
      [boolean, BigNumber, BigNumber, string] & {
        isStaked: boolean
        latestStakedNode: BigNumber
        amountStaked: BigNumber
        currentChallenge: string
      }
    >

    timedOutChallenges(
      rollup: string,
      startIndex: BigNumberish,
      max: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[string[], boolean] & { hasMore: boolean }>
  }

  filters: {}

  estimateGas: {
    areUnresolvedNodesLinear(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    checkDecidableNextNode(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    findNodeConflict(
      rollup: string,
      node1: BigNumberish,
      node2: BigNumberish,
      maxDepth: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    findStakerConflict(
      rollup: string,
      staker1: string,
      staker2: string,
      maxDepth: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    getConfig(rollup: string, overrides?: CallOverrides): Promise<BigNumber>

    getStakers(
      rollup: string,
      startIndex: BigNumberish,
      max: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    latestStaked(
      rollup: string,
      staker: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    refundableStakers(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    requireConfirmable(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    requireRejectable(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    stakedNodes(
      rollup: string,
      staker: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    stakerInfo(
      rollup: string,
      stakerAddress: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    timedOutChallenges(
      rollup: string,
      startIndex: BigNumberish,
      max: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>
  }

  populateTransaction: {
    areUnresolvedNodesLinear(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    checkDecidableNextNode(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    findNodeConflict(
      rollup: string,
      node1: BigNumberish,
      node2: BigNumberish,
      maxDepth: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    findStakerConflict(
      rollup: string,
      staker1: string,
      staker2: string,
      maxDepth: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    getConfig(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    getStakers(
      rollup: string,
      startIndex: BigNumberish,
      max: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    latestStaked(
      rollup: string,
      staker: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    refundableStakers(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    requireConfirmable(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    requireRejectable(
      rollup: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    stakedNodes(
      rollup: string,
      staker: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    stakerInfo(
      rollup: string,
      stakerAddress: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    timedOutChallenges(
      rollup: string,
      startIndex: BigNumberish,
      max: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>
  }
}
