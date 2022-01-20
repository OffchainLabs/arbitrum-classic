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
import { BigNumber } from 'ethers'

import { ERC20__factory, Multicall2__factory } from '../abi'
import { ArbTsError } from '../errors'
import {
  isL1Network,
  L1Network,
  l1Networks,
  L2Network,
  l2Networks,
} from './networks'

/**
 * Input to multicall aggregator
 */
export type CallInput<T extends unknown> = {
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
type DecoderReturnType<T extends CallInput<unknown>[]> = {
  [P in keyof T]: T[P] extends CallInput<unknown>
    ? ReturnType<T[P]['decoder']> | undefined
    : never
}

///////////////////////////////////////
/////// TOKEN CONDITIONAL TYPES ///////
///////////////////////////////////////
// these conditional types return check T, and if it matches
// the input type then they return a known output type
type AllowanceInputOutput<T> = T extends {
  allowance: { owner: string; spender: string }
}
  ? { allowance: BigNumber | undefined }
  : Record<string, never>
type BalanceInputOutput<T> = T extends { balanceOf: { account: string } }
  ? { balance: BigNumber | undefined }
  : Record<string, never>
type DecimalsInputOutput<T> = T extends { decimals: true }
  ? { decimals: number | undefined }
  : Record<string, never>
type NameInputOutput<T> = T extends { name: true }
  ? { name: string | undefined }
  : Record<string, never>
type SymbolInputOutput<T> = T extends { symbol: true }
  ? { symbol: string | undefined }
  : Record<string, never>
type TokenMultiInput = {
  balanceOf?: {
    account: string
  }
  allowance?: {
    owner: string
    spender: string
  }
  symbol?: true
  decimals?: true
  name?: true
}
// if we were given options at all then we convert
// those options to outputs
type TokenInputOutput<T> = T extends TokenMultiInput
  ? AllowanceInputOutput<T> &
      BalanceInputOutput<T> &
      DecimalsInputOutput<T> &
      NameInputOutput<T> &
      SymbolInputOutput<T>
  : { name: string }
//\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
//\\\\\ TOKEN CONDITIONAL TYPES \\\\\\\
//\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\

/**
 * Util for executing multi calls against the MultiCallV2 contract
 */
export class MultiCaller {
  constructor(
    public readonly provider: Provider,
    private readonly multicallerAddress: string
  ) {}

  /**
   * Finds the correct multicall address for the given provider and instantiates a multicaller
   * @param provider
   * @returns
   */
  public static async fromProvider(provider: Provider): Promise<MultiCaller> {
    const chainId = (await provider.getNetwork()).chainId
    const l2Network = l2Networks[chainId.toString()] as L2Network | undefined
    const l1Network = l1Networks[chainId.toString()] as L1Network | undefined

    const network = l2Network || l1Network
    if (!network) {
      throw new ArbTsError(
        `Unexpected network id: ${chainId}. Ensure that chain ${chainId} has been added as a network.`
      )
    }

    let multiCallAddr: string
    if (isL1Network(network)) {
      const firstL2 = l2Networks[network.partnerChainIDs[0]]
      if (!firstL2)
        throw new ArbTsError(
          `No partner chain found l1 network: ${network.chainID} : partner chain ids ${network.partnerChainIDs}`
        )
      multiCallAddr = firstL2.tokenBridge.l1MultiCall
    } else {
      multiCallAddr = network.tokenBridge.l2Multicall
    }

    return new MultiCaller(provider, multiCallAddr)
  }

  /**
   * Executes a multicall for the given parameters
   * Return values are order the same as the inputs.
   * If a call failed undefined is returned instead of the value.
   *
   * To get better type inference when the individual calls are of different types
   * create your inputs as a tuple and pass the tuple in. The return type will be
   * a tuple of the decoded return types. eg.
   *
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
   *   const res = await multiCaller.call(inputs)
   * ```
   * @param provider
   * @param params
   * @param requireSuccess Fail the whole call if any internal call fails
   * @returns
   */
  public async multiCall<T extends CallInput<unknown>[]>(
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
      if (success && returnData && returnData != '0x') {
        return params[index].decoder(returnData)
      }
      return undefined
    }) as DecoderReturnType<T>
  }

  /**
   * Multicall for token properties. Will collect all the requested properies for each of the
   * supplied token addresses.
   * @param erc20Addresses
   * @param options Defaults to just 'name'
   * @returns
   */
  public async getTokenData<T extends TokenMultiInput | undefined>(
    erc20Addresses: string[],
    options?: T
  ): // based on the type of options we return only the fields that were specified
  Promise<TokenInputOutput<T>[]>
  public async getTokenData<T extends TokenMultiInput | undefined>(
    erc20Addresses: string[],
    options?: T
  ): Promise<
    | { name: string }[]
    | {
        balance?: BigNumber
        allowance?: BigNumber
        symbol?: string
        decimals?: number
        name?: string
      }[]
  > {
    // if no options are supplied, then we just multicall for the names
    const defaultedOptions: TokenMultiInput = options || { name: true }
    const erc20Iface = ERC20__factory.createInterface()

    const input = []
    for (const t of erc20Addresses) {
      if (defaultedOptions.allowance) {
        input.push({
          targetAddr: t,
          encoder: () =>
            erc20Iface.encodeFunctionData('allowance', [
              defaultedOptions.allowance!.owner,
              defaultedOptions.allowance!.spender,
            ]),
          decoder: (returnData: string) =>
            erc20Iface.decodeFunctionResult(
              'allowance',
              returnData
            )[0] as BigNumber,
        })
      }

      if (defaultedOptions.balanceOf) {
        input.push({
          targetAddr: t,
          encoder: () =>
            erc20Iface.encodeFunctionData('balanceOf', [
              defaultedOptions.balanceOf!.account,
            ]),
          decoder: (returnData: string) =>
            erc20Iface.decodeFunctionResult(
              'balanceOf',
              returnData
            )[0] as BigNumber,
        })
      }

      if (defaultedOptions.decimals) {
        input.push({
          targetAddr: t,
          encoder: () => erc20Iface.encodeFunctionData('decimals'),
          decoder: (returnData: string) =>
            erc20Iface.decodeFunctionResult(
              'decimals',
              returnData
            )[0] as number,
        })
      }

      if (defaultedOptions.name) {
        input.push({
          targetAddr: t,
          encoder: () => erc20Iface.encodeFunctionData('name'),
          decoder: (returnData: string) =>
            erc20Iface.decodeFunctionResult('name', returnData)[0] as string,
        })
      }

      if (defaultedOptions.symbol) {
        input.push({
          targetAddr: t,
          encoder: () => erc20Iface.encodeFunctionData('symbol'),
          decoder: (returnData: string) =>
            erc20Iface.decodeFunctionResult('symbol', returnData)[0] as string,
        })
      }
    }

    const res = await this.multiCall(input)

    let i = 0
    const tokens = []
    while (i < res.length) {
      tokens.push({
        allowance: defaultedOptions.allowance
          ? (res[i++] as BigNumber)
          : undefined,
        balance: defaultedOptions.balanceOf
          ? (res[i++] as BigNumber)
          : undefined,
        decimals: defaultedOptions.decimals ? (res[i++] as number) : undefined,
        name: defaultedOptions.name ? (res[i++] as string) : undefined,
        symbol: defaultedOptions.symbol ? (res[i++] as string) : undefined,
      })
    }
    return tokens
  }
}
