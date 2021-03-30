/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

import { ArbSysFactory } from './abi/ArbSysFactory'

import { ArbAddressTable } from './abi/ArbAddressTable'
import { ArbAddressTableFactory } from './abi/ArbAddressTableFactory'

import { TransactionOverrides } from './abi'

import * as ethers from 'ethers'
import { JsonRpcProvider } from 'ethers/providers'
const ARB_SYS_ADDRESS = '0x0000000000000000000000000000000000000064'
const ARB_ADDRESS_TABLE_ADDRESS = '0x0000000000000000000000000000000000000066'

export async function withdrawEth(
  l2signer: ethers.Signer,
  value: ethers.utils.BigNumberish,
  overrides?: TransactionOverrides
): Promise<ethers.providers.TransactionResponse> {
  const arbsys = ArbSysFactory.connect(ARB_SYS_ADDRESS, l2signer)
  return arbsys.withdrawEth(await l2signer.getAddress(), {
    ...overrides,
    value,
  })
}

interface AddressIndexMemo {
  [address: string]: number
}

export const getAddressIndex = (() => {
  const addressToIndexMemo: AddressIndexMemo = {}
  let arbAddressTable: ArbAddressTable | undefined

  return async (
    address: string,
    signerOrProvider: ethers.Signer | JsonRpcProvider
  ) => {
    if (addressToIndexMemo[address]) {
      return addressToIndexMemo[address]
    }
    arbAddressTable =
      arbAddressTable ||
      ArbAddressTableFactory.connect(
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
