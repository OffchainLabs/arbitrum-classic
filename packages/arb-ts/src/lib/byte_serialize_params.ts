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

/**
#### Byte Serializing Solidity Arguments Schema

Arb-ts includes methods for [serializing parameters](https://developer.offchainlabs.com/docs/special_features#parameter-byte-serialization) for a solidity method into a single byte array to minimize calldata. It uses the following schema:

#### address[]:

| field         | size (bytes)       | Description                                                             |
| ------------- | ------------------ | ----------------------------------------------------------------------- |
| length        | 1                  | Size of array                                                           |
| is-registered | 1                  | 1 = all registered, 0 = not all registered                              |
| addresses     | 4 or 20 (x length) | If is registered, left-padded 4-byte integers; otherwise, eth addresses |

#### non-address[]:

| field  | size (bytes) | Description              |
| ------ | ------------ | ------------------------ |
| length | 1            | Size of array            |
| items  | (variable)   | All items (concatenated) |

#### address:

| field         | size (bytes) | Description                                                       |
| ------------- | ------------ | ----------------------------------------------------------------- |
| is-registered | 1            | 1 = registered, 0 = not registered                                |
| address       | 4 or 20      | If registered, left-padded 4-byte integer; otherwise, eth address |

 * @module Byte-Serialization
 */

import { Provider } from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'
import { isAddress as _isAddress } from '@ethersproject/address'
import { concat, hexZeroPad } from '@ethersproject/bytes'
import { BigNumber } from '@ethersproject/bignumber'

import { ArbAddressTable__factory, ArbAddressTable } from './abi'

import { ARB_ADDRESS_TABLE_ADDRESS } from './constants'

type PrimativeType = string | number | boolean | BigNumber
type PrimativeOrPrimativeArray = PrimativeType | PrimativeType[]
type BytesNumber = 1 | 4 | 8 | 16 | 32

interface AddressIndexMemo {
  [address: string]: number
}

export const getAddressIndex = (() => {
  const addressToIndexMemo: AddressIndexMemo = {}
  let arbAddressTable: ArbAddressTable | undefined

  return async (address: string, signerOrProvider: Signer | Provider) => {
    if (addressToIndexMemo[address]) {
      return addressToIndexMemo[address]
    }
    arbAddressTable =
      arbAddressTable ||
      ArbAddressTable__factory.connect(
        ARB_ADDRESS_TABLE_ADDRESS,
        signerOrProvider
      )
    const isRegistered = await arbAddressTable.addressExists(address)
    if (isRegistered) {
      const index = (await arbAddressTable.lookup(address)).toNumber()
      addressToIndexMemo[address] = index
      return index
    } else {
      return -1
    }
  }
})()

/**
  // to use:
  ```js
  const mySerializeParamsFunction = argSerializerConstructor("rpcurl")
  mySerializeParamsFunction(["4","5", "6"])
  ```
*/
export const argSerializerConstructor = (
  arbProvider: Provider
): ((params: PrimativeOrPrimativeArray[]) => Promise<Uint8Array>) => {
  return async (params: PrimativeOrPrimativeArray[]) => {
    return await serializeParams(params, async (address: string) => {
      return await getAddressIndex(address, arbProvider)
    })
  }
}

const isAddress = (input: PrimativeType) =>
  typeof input === 'string' && _isAddress(input)

const toUint = (val: PrimativeType, bytes: BytesNumber) =>
  hexZeroPad(BigNumber.from(val).toHexString(), bytes)

//  outputs string suitable for formatting
const formatPrimative = (value: PrimativeType) => {
  if (isAddress(value)) {
    return value as string
  } else if (typeof value === 'boolean') {
    return toUint(value ? 1 : 0, 1)
  } else if (
    typeof value === 'number' ||
    +value ||
    BigNumber.isBigNumber(value)
  ) {
    return toUint(value, 32)
  } else {
    throw new Error('unsupported type')
  }
}
/**
 * @param params array of serializable types to
 * @param addressToIndex optional getter of address index registered in table
 */
export const serializeParams = async (
  params: PrimativeOrPrimativeArray[],
  addressToIndex: (address: string) => Promise<number> = () =>
    new Promise(exec => exec(-1))
): Promise<Uint8Array> => {
  const formattedParams: string[] = []

  for (const param of params) {
    // handle arrays
    if (Array.isArray(param)) {
      let paramArray: PrimativeType[] = param as PrimativeType[]
      formattedParams.push(toUint(paramArray.length, 1))

      if (isAddress(paramArray[0])) {
        const indices = await Promise.all(
          paramArray.map(
            async address => await addressToIndex(address as string)
          )
        )
        // If all addresses are registered, serialize as indices
        if (indices.every(i => i > -1)) {
          paramArray = indices as number[]
          formattedParams.push(toUint(1, 1))
          paramArray.forEach(value => {
            formattedParams.push(toUint(value, 4))
          })
          // otherwise serialize as address
        } else {
          formattedParams.push(toUint(0, 1))
          paramArray.forEach(value => {
            formattedParams.push(formatPrimative(value))
          })
        }
      } else {
        paramArray.forEach(value => {
          formattedParams.push(formatPrimative(value))
        })
      }
    } else {
      //  handle non-arrays
      if (isAddress(param)) {
        const index = await addressToIndex(param as string)
        if (index > -1) {
          formattedParams.push(toUint(1, 1), toUint(index, 4))
        } else {
          formattedParams.push(toUint(0, 1), formatPrimative(param))
        }
      } else {
        formattedParams.push(formatPrimative(param))
      }
    }
  }
  return concat(formattedParams)
}
