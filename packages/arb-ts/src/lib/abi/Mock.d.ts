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
  CallOverrides,
} from '@ethersproject/contracts'
import { BytesLike } from '@ethersproject/bytes'
import { Listener, Provider } from '@ethersproject/providers'
import { FunctionFragment, EventFragment, Result } from '@ethersproject/abi'

interface MockInterface extends ethers.utils.Interface {
  functions: {
    'mocked()': FunctionFragment
    'sendContractTransaction(uint256,uint256,address,uint256,bytes)': FunctionFragment
  }

  encodeFunctionData(functionFragment: 'mocked', values?: undefined): string
  encodeFunctionData(
    functionFragment: 'sendContractTransaction',
    values: [BigNumberish, BigNumberish, string, BigNumberish, BytesLike]
  ): string

  decodeFunctionResult(functionFragment: 'mocked', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'sendContractTransaction',
    data: BytesLike
  ): Result

  events: {}
}

export class Mock extends Contract {
  connect(signerOrProvider: Signer | Provider | string): this
  attach(addressOrName: string): this
  deployed(): Promise<this>

  on(event: EventFilter | string, listener: Listener): this
  once(event: EventFilter | string, listener: Listener): this
  addListener(eventName: EventFilter | string, listener: Listener): this
  removeAllListeners(eventName: EventFilter | string): this
  removeListener(eventName: any, listener: Listener): this

  interface: MockInterface

  functions: {
    mocked(overrides?: CallOverrides): Promise<[string]>

    'mocked()'(overrides?: CallOverrides): Promise<[string]>

    sendContractTransaction(
      arg0: BigNumberish,
      arg1: BigNumberish,
      arg2: string,
      arg3: BigNumberish,
      arg4: BytesLike,
      overrides?: Overrides
    ): Promise<ContractTransaction>

    'sendContractTransaction(uint256,uint256,address,uint256,bytes)'(
      arg0: BigNumberish,
      arg1: BigNumberish,
      arg2: string,
      arg3: BigNumberish,
      arg4: BytesLike,
      overrides?: Overrides
    ): Promise<ContractTransaction>
  }

  mocked(overrides?: CallOverrides): Promise<string>

  'mocked()'(overrides?: CallOverrides): Promise<string>

  sendContractTransaction(
    arg0: BigNumberish,
    arg1: BigNumberish,
    arg2: string,
    arg3: BigNumberish,
    arg4: BytesLike,
    overrides?: Overrides
  ): Promise<ContractTransaction>

  'sendContractTransaction(uint256,uint256,address,uint256,bytes)'(
    arg0: BigNumberish,
    arg1: BigNumberish,
    arg2: string,
    arg3: BigNumberish,
    arg4: BytesLike,
    overrides?: Overrides
  ): Promise<ContractTransaction>

  callStatic: {
    mocked(overrides?: CallOverrides): Promise<string>

    'mocked()'(overrides?: CallOverrides): Promise<string>

    sendContractTransaction(
      arg0: BigNumberish,
      arg1: BigNumberish,
      arg2: string,
      arg3: BigNumberish,
      arg4: BytesLike,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    'sendContractTransaction(uint256,uint256,address,uint256,bytes)'(
      arg0: BigNumberish,
      arg1: BigNumberish,
      arg2: string,
      arg3: BigNumberish,
      arg4: BytesLike,
      overrides?: CallOverrides
    ): Promise<BigNumber>
  }

  filters: {}

  estimateGas: {
    mocked(overrides?: CallOverrides): Promise<BigNumber>

    'mocked()'(overrides?: CallOverrides): Promise<BigNumber>

    sendContractTransaction(
      arg0: BigNumberish,
      arg1: BigNumberish,
      arg2: string,
      arg3: BigNumberish,
      arg4: BytesLike,
      overrides?: Overrides
    ): Promise<BigNumber>

    'sendContractTransaction(uint256,uint256,address,uint256,bytes)'(
      arg0: BigNumberish,
      arg1: BigNumberish,
      arg2: string,
      arg3: BigNumberish,
      arg4: BytesLike,
      overrides?: Overrides
    ): Promise<BigNumber>
  }

  populateTransaction: {
    mocked(overrides?: CallOverrides): Promise<PopulatedTransaction>

    'mocked()'(overrides?: CallOverrides): Promise<PopulatedTransaction>

    sendContractTransaction(
      arg0: BigNumberish,
      arg1: BigNumberish,
      arg2: string,
      arg3: BigNumberish,
      arg4: BytesLike,
      overrides?: Overrides
    ): Promise<PopulatedTransaction>

    'sendContractTransaction(uint256,uint256,address,uint256,bytes)'(
      arg0: BigNumberish,
      arg1: BigNumberish,
      arg2: string,
      arg3: BigNumberish,
      arg4: BytesLike,
      overrides?: Overrides
    ): Promise<PopulatedTransaction>
  }
}
