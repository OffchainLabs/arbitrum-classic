import { providers, utils, Wallet, BigNumber, constants } from 'ethers'
import { Bridge } from '../src/lib/bridge'
import { expect } from 'chai'
import config from './config'
import { EthERC20BridgeFactory } from '../src/lib/abi/EthERC20BridgeFactory'

const { parseEther } = utils
const { ethRPC, arbRPC, preFundedSignerPK, erc20BridgeAddress } = config

const ethProvider = new providers.JsonRpcProvider(ethRPC)
const arbProvider = new providers.JsonRpcProvider(arbRPC)

const preFundedWallet = new Wallet(preFundedSignerPK, ethProvider)

const l1TestWallet = Wallet.createRandom(ethProvider)
const l2TestWallet = Wallet.createRandom(arbProvider)

let arb20BridgeAddress: string

let bridge: Bridge | undefined

describe('setup bridge', () => {
  it('instantiate bridge ', async () => {
    const ethBridge = EthERC20BridgeFactory.connect(
      erc20BridgeAddress,
      l1TestWallet
    )
    arb20BridgeAddress = await ethBridge.l2Buddy()
    bridge = new Bridge(
      erc20BridgeAddress,
      arb20BridgeAddress,
      l1TestWallet,
      l2TestWallet
    )
    expect(true).to.be.true
  })

  it('fund l1 test wallet with eth', async () => {
    const res = await preFundedWallet.sendTransaction({
      to: l1TestWallet.address,
      value: utils.parseEther('0.01'),
    })
    await res.wait()
    const testWAlletBalance = await l1TestWallet.getBalance()
    expect(testWAlletBalance.eq(parseEther('0.01'))).to.be.true
  })
})

describe('deposit ether', () => {
  let testWalletL1EthBalance: BigNumber
  let testWalletL2EthBalance: BigNumber

  it('has expected intial values', async () => {
    testWalletL1EthBalance = await bridge.getAndUpdateL1EthBalance()
    testWalletL2EthBalance = await bridge.getAndUpdateL2EthBalance()
    expect(testWalletL1EthBalance.eq(parseEther('0.01'))).to.be.true
    expect(testWalletL2EthBalance.eq(constants.Zero)).to.be.true
  })

  it('deposits ether', async () => {
    const depositAmount = parseEther('0.001')
    const res = await bridge.depositETH(depositAmount)
    const rec = await res.wait()

    expect(rec.status).to.equal(1)
    testWalletL2EthBalance = await bridge.getAndUpdateL2EthBalance()
    expect(testWalletL2EthBalance.eq(depositAmount)).to.be.true
  })
})
