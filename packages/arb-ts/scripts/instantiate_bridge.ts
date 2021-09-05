import { Bridge, BridgeHelper, networks } from '../src'
import { providers, utils, Wallet, BigNumber, constants, ethers } from 'ethers'

import dotenv from 'dotenv'
import args from './getCLargs'
dotenv.config()

const pk = process.env['DEVNET_PRIVKEY'] as string
const mnemonic = process.env['DEV_MNEMONIC'] as string
const verbose = process.env['VERBOSE'] as string

const defaultNetworkId = 4

export const instantiateBridge = async (
  l1pkParam?: string,
  l2PkParam?: string
) => {
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

  const ethProvider = new providers.JsonRpcProvider(l1Network.rpcURL)
  const arbProvider = new providers.JsonRpcProvider(l2Network.rpcURL)

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
