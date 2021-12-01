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

import { JsonRpcProvider } from '@ethersproject/providers'
import { Wallet } from '@ethersproject/wallet'

import dotenv from 'dotenv'
import args from './getCLargs'
import { Bridge, networks } from '../src'
import { Network } from '../src/lib/networks'

dotenv.config()

const pk = process.env['DEVNET_PRIVKEY'] as string
const mnemonic = process.env['DEV_MNEMONIC'] as string
const verbose = process.env['VERBOSE'] as string

const defaultNetworkId = 4

export const instantiateBridge = async (
  l1pkParam?: string,
  l2PkParam?: string
): Promise<{ bridge: Bridge; l1Network: Network; l2Network: Network }> => {
  if (!l1pkParam) {
    if (!pk && !mnemonic)
      throw new Error('need DEVNET_PRIVKEY or DEV_MNEMONIC env var')

    if (pk && mnemonic)
      throw new Error(
        'You have both a DEVNET_PRIVKEY and DEV_MNEMONIC var set; pick one! '
      )
  }

  let networkID = args.networkID
  if (!networkID) {
    verbose &&
      console.log(
        'No networkID command line arg provided; using network',
        defaultNetworkId
      )

    networkID = defaultNetworkId
  }
  const network = networks[networkID]
  if (!network) {
    throw new Error(`Unrecognized network ID: ${networkID}`)
  }

  const l1Network = network.isArbitrum
    ? networks[network.partnerChainID]
    : network
  const l2Network = networks[l1Network.partnerChainID]

  if (!l1Network.rpcURL) {
    throw new Error('L1 rpc url not set (see .env.sample or networks.ts)')
  }
  if (!l2Network.rpcURL) {
    throw new Error('L2 rpc url not set (see .env.sample or networks.ts)')
  }
  const ethProvider = new JsonRpcProvider(l1Network.rpcURL)

  const arbProvider = new JsonRpcProvider(l2Network.rpcURL)

  const l1Signer = (() => {
    if (l1pkParam) {
      return new Wallet(l1pkParam, ethProvider)
    } else if (mnemonic) {
      return Wallet.fromMnemonic(mnemonic).connect(ethProvider)
    } else if (pk) {
      return new Wallet(pk, ethProvider)
    } else {
      throw new Error('impossible path')
    }
  })()

  const l2Signer = (() => {
    if (l2PkParam) {
      return new Wallet(l2PkParam, arbProvider)
    } else if (mnemonic) {
      return Wallet.fromMnemonic(mnemonic).connect(arbProvider)
    } else if (pk) {
      return new Wallet(pk, arbProvider)
    } else {
      throw new Error('impossible path')
    }
  })()

  const bridge = await Bridge.init(l1Signer, l2Signer)
  if (verbose) {
    console.log('')
    console.log('**** Bridge instantiated w/ address', l1Signer.address, '****')
    console.log('')
  }

  return { bridge, l1Network, l2Network }
}
