import { providers, utils, Wallet, BigNumber } from 'ethers'
import { Bridge } from '../src/lib/bridge'
import { expect } from 'chai'
import config from './config'
const { ethRPC, arbRPC, preFundedSignerPK, erc20BridgeAddress } = config

const ethProvider = new providers.JsonRpcProvider(ethRPC)
const arbProvider = new providers.JsonRpcProvider(arbRPC)

const preFundedWallet = new Wallet(preFundedSignerPK, ethProvider)

const l1TestWallet = Wallet.createRandom(ethProvider)
const l2TestWallet = Wallet.createRandom(arbProvider)

const bridge = new Bridge(
  erc20BridgeAddress,
  'arb20BridgeAddress',
  l1TestWallet,
  l2TestWallet
)

// **

//
