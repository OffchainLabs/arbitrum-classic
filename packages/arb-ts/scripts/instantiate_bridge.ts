import { Bridge, BridgeHelper, networks } from '../src'
import { providers, utils, Wallet, BigNumber, constants, ethers } from 'ethers'

import yargs from 'yargs/yargs'

const pk = process.env['DEVNET_PRIVKEY']
if (!pk) throw new Error('need DEVNET_PRIVKEY')

export const instantiateBridge = async () => {
  const argv = yargs(process.argv.slice(2)).argv
  let networkID = argv.networkID as number
  if (!networkID) {
    networkID = 4
  }
  const network = networks[networkID]
  if (network) {
    throw new Error(`Unrecognized network ID: ${networkID}`)
  }

  const l1Network = network.isArbitrum
    ? networks[network.partnerChainID]
    : network
  const l2Network = networks[l1Network.partnerChainID]

  const ethProvider = new providers.JsonRpcProvider(l1Network.rpcURL)
  const arbProvider = new providers.JsonRpcProvider(l2Network.rpcURL)

  const l1Signer = new Wallet(pk, ethProvider)
  const l2Signer = new Wallet(pk, arbProvider)

  const bridge = await Bridge.init(l1Signer, l2Signer)

  return { bridge, l1Network, l2Network }
}
