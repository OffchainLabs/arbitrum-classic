import { providers, utils, Wallet, BigNumber, constants } from 'ethers'
import { Bridge } from '../src/lib/bridge'
import { expect } from 'chai'
import config from './config'
import { TestERC20__factory } from '../src/lib/abi/factories/TestERC20__factory'
import { TestERC777__factory } from '../src/lib/abi/factories/TestERC777__factory'
import { StandardArbERC20__factory } from '../src/lib/abi/factories/StandardArbERC20__factory'

import { EthERC20Bridge } from '../src/lib/abi/EthERC20Bridge'
const { parseEther } = utils
const {
  ethRPC,
  arbRPC,
  preFundedSignerPK,
  erc20BridgeAddress,
  arbTokenBridgeAddress,
  l1gasPrice,
  existantTestERC20,
} = config.kovan4

const ethProvider = new providers.JsonRpcProvider(ethRPC)
const arbProvider = new providers.JsonRpcProvider(arbRPC)

const preFundedWallet = new Wallet(preFundedSignerPK, ethProvider)

const testPk = utils.formatBytes32String(Math.random().toString())

const l1TestWallet = new Wallet(testPk, ethProvider)
const l2TestWallet = new Wallet(testPk, arbProvider)
const wait = (ms: number) => new Promise(res => setTimeout(res, ms))

const depositAmount = '0.01'
let erc20Address = existantTestERC20

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
    const ethToL2DepositAmount = parseEther('0.0001')
    const res = await bridge.depositETH(ethToL2DepositAmount)
    const rec = await res.wait()

    expect(rec.status).to.equal(1)
    await wait(2000)
    testWalletL2EthBalance = await bridge.getAndUpdateL2EthBalance()
    expect(testWalletL2EthBalance.eq(ethToL2DepositAmount)).to.be.true
  })
})

const tokenDepositAmmount = BigNumber.from(100)

describe('ERC20', () => {
  it('create l1 erc20 w initial supply', async () => {
    const testTokenFactory = await new TestERC20__factory(preFundedWallet)
    const testToken = erc20Address
      ? await testTokenFactory.attach(erc20Address)
      : await testTokenFactory.deploy()

    const bal = await testToken.balanceOf(preFundedWallet.address)
    expect(bal.gt(BigNumber.from(40000000))).to.be.true

    erc20Address = testToken.address
    console.info('deployed at l1 address', erc20Address)
    const res = await testToken.transfer(
      l1TestWallet.address,
      BigNumber.from(200)
    )
    const rec = await res.wait()
    const data = await bridge.getAndUpdateL1TokenData(erc20Address)
    const testWalletBal = data.ERC20 && data.ERC20.balance
    expect(testWalletBal && testWalletBal.eq(BigNumber.from(200))).to.be.true
  })

  it('initial erc20 deposit works', async () => {
    console.warn('approve tokekn?')

    const approveRes = await bridge.approveToken(erc20Address)
    const approveRec = await approveRes.wait()
    expect(approveRec.status).to.equal(1)

    console.warn('approve done')

    const data = await bridge.getAndUpdateL1TokenData(erc20Address)
    const allowed = data.ERC20 && data.ERC20.allowed
    expect(allowed).to.be.true

    const despositRes = await bridge.depositAsERC20(
      erc20Address,
      tokenDepositAmmount,
      BigNumber.from(10000000000000),
      BigNumber.from(0),
      undefined,
      { gasLimit: 210000, gasPrice: l1gasPrice }
    )

    const depositRec = await despositRes.wait()

    expect(depositRec.status === 1).to.be.true
    await wait(2000)
    const l2Data = await bridge.getAndUpdateL2TokenData(erc20Address)

    const testWalletL2Balance = l2Data && l2Data.ERC20 && l2Data.ERC20.balance

    expect(testWalletL2Balance && testWalletL2Balance.eq(tokenDepositAmmount))
      .to.be.true
  })

  it('erc20 is properly deployed in L2', async () => {
    const erc20L2Address = await bridge.getERC20L2Address(erc20Address)
    const arbERC20 = StandardArbERC20__factory.connect(
      erc20L2Address,
      arbProvider
    )
    const l2Code = await arbProvider.getCode(erc20L2Address)
    expect(l2Code.length > 2).to.be.true

    const balance = await arbERC20.balanceOf(l1TestWallet.address)
    expect(balance.eq(tokenDepositAmmount)).to.be.true
  })

  const withdrawAmount = BigNumber.from(2)
  it('withdraw erc20', async () => {
    const withdrawRes = await bridge.withdrawERC20(erc20Address, withdrawAmount)
    const withdrawRec = await withdrawRes.wait()
    expect(withdrawRec.status).to.equal(1)
    await wait(15000)
    const l1EventData = await bridge.getL2ToL1EventData(l1TestWallet.address)
    const withdrawTokenData = await bridge.getTokenWithdrawEventData(
      l1TestWallet.address
    )
    console.log(l1EventData, withdrawTokenData)
    expect(l1EventData.length).to.equal(withdrawTokenData.length)
  })
})
