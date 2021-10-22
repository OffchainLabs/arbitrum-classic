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
import type { TypedEventFilter, TypedEvent, TypedListener } from './common'

interface AddressMappingTestInterface extends ethers.utils.Interface {
  functions: {
    'getL1AddressTest(address)': FunctionFragment
  }

  encodeFunctionData(
    functionFragment: 'getL1AddressTest',
    values: [string]
  ): string

  decodeFunctionResult(
    functionFragment: 'getL1AddressTest',
    data: BytesLike
  ): Result

  events: {
    'TxToL1(address,address,uint256,bytes)': EventFragment
  }

  getEvent(nameOrSignatureOrTopic: 'TxToL1'): EventFragment
}

export type TxToL1Event = TypedEvent<
  [string, string, BigNumber, string] & {
    _from: string
    _to: string
    _id: BigNumber
    _data: string
  }
>

export class AddressMappingTest extends BaseContract {
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

  interface: AddressMappingTestInterface

  functions: {
    getL1AddressTest(
      sender: string,
      overrides?: CallOverrides
    ): Promise<[string] & { l1Address: string }>
  }

  getL1AddressTest(sender: string, overrides?: CallOverrides): Promise<string>

  callStatic: {
    getL1AddressTest(sender: string, overrides?: CallOverrides): Promise<string>
  }

  filters: {
    'TxToL1(address,address,uint256,bytes)'(
      _from?: string | null,
      _to?: string | null,
      _id?: BigNumberish | null,
      _data?: null
    ): TypedEventFilter<
      [string, string, BigNumber, string],
      { _from: string; _to: string; _id: BigNumber; _data: string }
    >

    TxToL1(
      _from?: string | null,
      _to?: string | null,
      _id?: BigNumberish | null,
      _data?: null
    ): TypedEventFilter<
      [string, string, BigNumber, string],
      { _from: string; _to: string; _id: BigNumber; _data: string }
    >
  }

  estimateGas: {
    getL1AddressTest(
      sender: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>
  }

  populateTransaction: {
    getL1AddressTest(
      sender: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>
  }
}
