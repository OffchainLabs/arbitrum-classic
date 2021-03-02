import { providers, utils, Wallet, BigNumber } from 'ethers'
import { Bridge } from '../src/lib/bridge'
import { expect } from 'chai'
import config from './config'
const {
  ethRPC,
  arbRPC,
  preFundedSignerPK,
  inboxAddress,
  erc20BridgeAddress,
} = config

const arbChainAddress = process.argv[process.argv.length - 1]

const ethProvider = new providers.JsonRpcProvider(ethRPC)
const arbProvider = new providers.JsonRpcProvider(arbRPC)

const preFundedWallet = new Wallet(preFundedSignerPK, ethProvider)

const l1TestWallet = Wallet.createRandom(ethProvider)
const l2TestWallet = Wallet.createRandom(arbProvider)

const bridge = new Bridge(
  inboxAddress,
  erc20BridgeAddress,
  erc20BridgeAddress,
  ethProvider,
  l1TestWallet,
  arbProvider,
  l2TestWallet
)

// **

//
