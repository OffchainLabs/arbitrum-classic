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
  Overrides,
  PayableOverrides,
  CallOverrides,
} from 'ethers'
import { BytesLike } from '@ethersproject/bytes'
import { Listener, Provider } from '@ethersproject/providers'
import { FunctionFragment, EventFragment, Result } from '@ethersproject/abi'
import { TypedEventFilter, TypedEvent, TypedListener } from './commons'

interface AddInterface extends ethers.utils.Interface {
  functions: {
    'add(uint256,uint256)': FunctionFragment
    'getSeqNum()': FunctionFragment
    'isNotTopLevel()': FunctionFragment
    'isTopLevel()': FunctionFragment
    'mult(uint256,uint256)': FunctionFragment
    'payTo(address)': FunctionFragment
    'pythag(uint256,uint256)': FunctionFragment
    'withdraw5000()': FunctionFragment
    'withdrawMyEth()': FunctionFragment
  }

  encodeFunctionData(
    functionFragment: 'add',
    values: [BigNumberish, BigNumberish]
  ): string
  encodeFunctionData(functionFragment: 'getSeqNum', values?: undefined): string
  encodeFunctionData(
    functionFragment: 'isNotTopLevel',
    values?: undefined
  ): string
  encodeFunctionData(functionFragment: 'isTopLevel', values?: undefined): string
  encodeFunctionData(
    functionFragment: 'mult',
    values: [BigNumberish, BigNumberish]
  ): string
  encodeFunctionData(functionFragment: 'payTo', values: [string]): string
  encodeFunctionData(
    functionFragment: 'pythag',
    values: [BigNumberish, BigNumberish]
  ): string
  encodeFunctionData(
    functionFragment: 'withdraw5000',
    values?: undefined
  ): string
  encodeFunctionData(
    functionFragment: 'withdrawMyEth',
    values?: undefined
  ): string

  decodeFunctionResult(functionFragment: 'add', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'getSeqNum', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'isNotTopLevel',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'isTopLevel', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'mult', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'payTo', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'pythag', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'withdraw5000',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'withdrawMyEth',
    data: BytesLike
  ): Result

  events: {}
}

export class Add extends BaseContract {
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

  interface: AddInterface

  functions: {
    add(
      x: BigNumberish,
      y: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>

    getSeqNum(overrides?: CallOverrides): Promise<[BigNumber]>

    isNotTopLevel(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>

    isTopLevel(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>

    mult(
      x: BigNumberish,
      y: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>

    payTo(
      addr: string,
      overrides?: PayableOverrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>

    pythag(
      x: BigNumberish,
      y: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>

    withdraw5000(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>

    withdrawMyEth(
      overrides?: PayableOverrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>
  }

  add(
    x: BigNumberish,
    y: BigNumberish,
    overrides?: CallOverrides
  ): Promise<BigNumber>

  getSeqNum(overrides?: CallOverrides): Promise<BigNumber>

  isNotTopLevel(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>

  isTopLevel(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>

  mult(
    x: BigNumberish,
    y: BigNumberish,
    overrides?: CallOverrides
  ): Promise<BigNumber>

  payTo(
    addr: string,
    overrides?: PayableOverrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>

  pythag(
    x: BigNumberish,
    y: BigNumberish,
    overrides?: CallOverrides
  ): Promise<BigNumber>

  withdraw5000(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>

  withdrawMyEth(
    overrides?: PayableOverrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>

  callStatic: {
    add(
      x: BigNumberish,
      y: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    getSeqNum(overrides?: CallOverrides): Promise<BigNumber>

    isNotTopLevel(overrides?: CallOverrides): Promise<boolean>

    isTopLevel(overrides?: CallOverrides): Promise<boolean>

    mult(
      x: BigNumberish,
      y: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    payTo(addr: string, overrides?: CallOverrides): Promise<void>

    pythag(
      x: BigNumberish,
      y: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    withdraw5000(overrides?: CallOverrides): Promise<void>

    withdrawMyEth(overrides?: CallOverrides): Promise<void>
  }

  filters: {}

  estimateGas: {
    add(
      x: BigNumberish,
      y: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    getSeqNum(overrides?: CallOverrides): Promise<BigNumber>

    isNotTopLevel(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>

    isTopLevel(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>

    mult(
      x: BigNumberish,
      y: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    payTo(
      addr: string,
      overrides?: PayableOverrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>

    pythag(
      x: BigNumberish,
      y: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    withdraw5000(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>

    withdrawMyEth(
      overrides?: PayableOverrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>
  }

  populateTransaction: {
    add(
      x: BigNumberish,
      y: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    getSeqNum(overrides?: CallOverrides): Promise<PopulatedTransaction>

    isNotTopLevel(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>

    isTopLevel(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>

    mult(
      x: BigNumberish,
      y: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    payTo(
      addr: string,
      overrides?: PayableOverrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>

    pythag(
      x: BigNumberish,
      y: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    withdraw5000(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>

    withdrawMyEth(
      overrides?: PayableOverrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>
  }
}
