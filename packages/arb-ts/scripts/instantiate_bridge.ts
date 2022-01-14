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
import { EthBridger, TokenBridger } from '../src'
import {
  L1Network,
  l1Networks,
  L2Network,
  l2Networks,
} from '../src/lib/utils/networks'
import { Signer } from 'ethers'
import { AdminTokenBridger } from '../src/lib/assetBridger'

dotenv.config()

const pk = process.env['DEVNET_PRIVKEY'] as string
const mnemonic = process.env['DEV_MNEMONIC'] as string
const verbose = process.env['VERBOSE'] as string

const defaultNetworkId = 421611

export const instantiateBridge = async (
  l1pkParam?: string,
  l2PkParam?: string
): Promise<{
  l1Network: L1Network
  l2Network: L2Network
  l1Signer: Signer
  l2Signer: Signer
  tokenBridger: TokenBridger
  ethBridger: EthBridger
  adminTokenBridger: AdminTokenBridger
}> => {
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
  const isL1 = !!l1Networks[networkID]
  const isL2 = !!l2Networks[networkID]
  if (!isL1 && !isL2) {
    throw new Error(`Unrecognized network ID: ${networkID}`)
  }
  if (!isL2) {
    throw new Error(`Tests must specify an L2 network ID: ${networkID}`)
  }

  const l2Network = l2Networks[networkID]
  const l1Network = l1Networks[l2Network.partnerChainID]

  if (!l1Network) {
    throw new Error(
      `Unrecognised partner chain id: ${l2Network.partnerChainID}`
    )
  }

  if (!l1Network.rpcURL) {
    throw new Error('L1 rpc url not set (see .env.sample or networks.ts)')
  }
  if (!l2Network.rpcURL) {
    throw new Error('L2 rpc url not set (see .env.sample or utils/networks.ts)')
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

  if (verbose) {
    console.log('')
    console.log(
      '**** Bridger instantiated w/ address',
      l1Signer.address,
      '****'
    )
    console.log('')
  }

  const tokenBridger = new TokenBridger(l2Network)
  const adminTokenBridger = new AdminTokenBridger(l2Network)
  const ethBridger = new EthBridger(l2Network)

  return {
    l1Network,
    l2Network,
    l1Signer,
    l2Signer,
    tokenBridger,
    ethBridger,
    adminTokenBridger,
  }
}
