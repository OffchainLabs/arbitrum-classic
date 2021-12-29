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

import { Provider } from '@ethersproject/abstract-provider'

import { Multicall2__factory } from '../abi'

/**
 * Input to multicall aggregator
 */
type CallInput<T extends any> = {
  /**
   * Address of the target contract to be called
   */
  targetAddr: string
  /**
   * Function to produce encoded call data
   */
  encoder: () => string
  /**
   * Function to decode the result of the call
   */
  decoder: (returnData: string) => T
}

/**
 * For each item in T this DecoderReturnType<T> yields the return
 * type of the decoder property
 */
type DecoderReturnType<T extends CallInput<any>[]> = {
  [P in keyof T]: T[P] extends CallInput<any>
    ? ReturnType<T[P]['decoder']> | undefined
    : never
}

/**
 * Util for executing multi calls against the MultiCallV2 contract
 */
export class MultiCaller {
  constructor(
    public readonly provider: Provider,
    private readonly multicallerAddress: string
  ) {}

  /**
   * Executes a multicall for the given parameters
   * Return values are order the same as the inputs.
   * If a call failed undefined is returned instead of the value.
   *
   * To get better type inference create your inputs as a tuple and pass the tuple in
   * The return type will be a tuple of the decoded return types. eg.
   *
   * ```typescript
   *   const inputs: [
   *     CallInput<Awaited<ReturnType<ERC20['functions']['balanceOf']>>[0]>,
   *     CallInput<Awaited<ReturnType<ERC20['functions']['name']>>[0]>
   *   ] = [
   *     {
   *       targetAddr: token.address,
   *       encoder: () => token.interface.encodeFunctionData('balanceOf', ['']),
   *       decoder: (returnData: string) =>
   *         token.interface.decodeFunctionResult('balanceOf', returnData)[0],
   *     },
   *     {
   *       targetAddr: token.address,
   *       encoder: () => token.interface.encodeFunctionData('name'),
   *       decoder: (returnData: string) =>
   *         token.interface.decodeFunctionResult('name', returnData)[0],
   *     },
   *   ]
   *
   *   const res = await multiCaller.call(provider, inputs)
   * ```
   * @param provider
   * @param params
   * @param requireSuccess Fail the whole call if any internal call fails
   * @returns
   */
  public async call<T extends CallInput<any>[]>(
    params: T,
    requireSuccess = false
  ): Promise<DecoderReturnType<T>> {
    const multiCall = Multicall2__factory.connect(
      this.multicallerAddress,
      this.provider
    )
    const args = params.map(p => ({
      target: p.targetAddr,
      callData: p.encoder(),
    }))

    const outputs = await multiCall.callStatic.tryAggregate(
      requireSuccess,
      args
    )

    return outputs.map(({ success, returnData }, index) => {
      if (success) {
        return params[index].decoder(returnData)
      }
      return undefined
    }) as DecoderReturnType<T>
  }
}
