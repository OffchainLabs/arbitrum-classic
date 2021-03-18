import { providers, utils, Wallet, BigNumber, constants } from 'ethers'
import { Bridge } from '../src/lib/bridge'
import { expect } from 'chai'
import config from './config'
import { TestERC20__factory } from '../src/lib/abi/factories/TestERC20__factory'
import { TestERC777__factory } from '../src/lib/abi/factories/TestERC777__factory'

import { EthERC20Bridge } from '../src/lib/abi/EthERC20Bridge'

const { parseEther } = utils
const {
  ethRPC,
  arbRPC,
  preFundedSignerPK,
  erc20BridgeAddress,
  arbTokenBridgeAddress,
} = config

const ethProvider = new providers.JsonRpcProvider(ethRPC)
const arbProvider = new providers.JsonRpcProvider(arbRPC)

const preFundedWallet = new Wallet(preFundedSignerPK, ethProvider)

const testPk = utils.formatBytes32String(Math.random().toString())

const l1TestWallet = new Wallet(testPk, ethProvider)
const l2TestWallet = new Wallet(testPk, arbProvider)

const depositAmount = '0.001'

console.info('preFundedWallet', preFundedWallet.address)

console.info('test wallet', l1TestWallet.address)

const bridge = new Bridge(
  erc20BridgeAddress,
  arbTokenBridgeAddress,
  l1TestWallet,
  l2TestWallet
)

describe('setup', () => {
  it('fund l1 test wallet with eth', async () => {
    const res = await preFundedWallet.sendTransaction({
      to: l1TestWallet.address,
      value: utils.parseEther(depositAmount),
    })
    const rec = await res.wait()
    const testWAlletBalance = await l1TestWallet.getBalance()
    expect(testWAlletBalance.eq(parseEther(depositAmount))).to.be.true
  })
})

describe('deposit ether', () => {
  let testWalletL1EthBalance: BigNumber
  let testWalletL2EthBalance: BigNumber

  it('has expected intial values', async () => {
    testWalletL1EthBalance = await bridge.getAndUpdateL1EthBalance()
    testWalletL2EthBalance = await bridge.getAndUpdateL2EthBalance()
    expect(testWalletL1EthBalance.eq(parseEther(depositAmount))).to.be.true
    expect(testWalletL2EthBalance.eq(constants.Zero)).to.be.true
  })

  it('deposits ether', async () => {
    const depositAmount = parseEther('0.0001')
    const res = await bridge.depositETH(depositAmount)
    const rec = await res.wait()

    expect(rec.status).to.equal(1)
    testWalletL2EthBalance = await bridge.getAndUpdateL2EthBalance()
    expect(testWalletL2EthBalance.eq(depositAmount)).to.be.true
  })
})

describe('ERC20', () => {
  let erc20Address: string
  it('create l1 erc20 w initial supply', async () => {
    const testTokenFactory = await new TestERC20__factory(preFundedWallet)
    const testToken = await testTokenFactory.deploy()

    const bal = await testToken.balanceOf(preFundedWallet.address)
    erc20Address = testToken.address
    expect(bal.eq(BigNumber.from(50000000))).to.be.true
  })

  it('initial erc20 deposit works', async () => {
    const res = await bridge.depositAsERC20(
      erc20Address,
      BigNumber.from(100),
      BigNumber.from(10000000000000),
      BigNumber.from(10000000000000)
    )
    const rec = await res.wait()

    expect(rec.status === 1).to.be.true
  })
})
