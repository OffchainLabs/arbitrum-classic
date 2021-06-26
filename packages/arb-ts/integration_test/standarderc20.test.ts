import {
  providers,
  utils,
  Wallet,
  BigNumber,
  constants,
  ethers,
  ContractReceipt,
} from 'ethers'
import { Bridge } from '../src/lib/bridge'
import { Network } from '../src/lib/networks'

import { expect } from 'chai'
import config from './config'
import { TestERC20__factory } from '../src/lib/abi/factories/TestERC20__factory'

import yargs from 'yargs/yargs'
import chalk from 'chalk'
import {
  fundL1,
  fundL2,
  testRetryableTicket,
  prettyLog,
  warn,
  instantiateRandomBridge,
  fundL2Token,
  tokenFundAmount,
} from './testHelpers'
const { Zero, AddressZero } = constants
import dotenv from 'dotenv'

dotenv.config()
const argv = yargs(process.argv.slice(2)).argv
let networkID = argv.networkID as string

networkID = networkID || '4'
if (!config[networkID]) {
  throw new Error('network not supported')
}
const { existentTestERC20 } = config[networkID]

describe('ERC20', () => {
  it('deposits erc20 (no L2 Eth funding)', async () => {
    const { bridge } = await instantiateRandomBridge()
    await fundL1(bridge)
    await depositTokenTest(bridge)
  })
  it.skip('deposits erc20 (with L2 Eth funding)', async () => {
    const { bridge } = await instantiateRandomBridge()
    await fundL1(bridge)
    await fundL2(bridge)
    await depositTokenTest(bridge)
  })

  it('withdraws erc20', async () => {
    const tokenWithdrawAmount = BigNumber.from(1)
    const { bridge } = await instantiateRandomBridge()
    await fundL2(bridge)
    const result = await fundL2Token(bridge)
    if (!result) {
      warn('Prefunded wallet not funded with tokens; skipping ERC20 withdraw')
      return
    }
    const withdrawRes = await bridge.withdrawERC20(
      existentTestERC20,
      tokenWithdrawAmount
    )
    const withdrawRec = await withdrawRes.wait()

    expect(withdrawRec.status).to.equal(1)
    const withdrawEventData = (
      await bridge.getWithdrawalsInL2Transaction(withdrawRec)
    )[0]

    expect(withdrawEventData).to.exist

    const l2Data = await bridge.getAndUpdateL2TokenData(existentTestERC20)
    const testWalletL2Balance = l2Data && l2Data.ERC20 && l2Data.ERC20.balance
    expect(
      testWalletL2Balance &&
        testWalletL2Balance.add(tokenWithdrawAmount).eq(tokenFundAmount)
    ).to.be.true
  })
})

const depositTokenTest = async (bridge: Bridge) => {
  const tokenDepositAmount = BigNumber.from(1)

  const testToken = TestERC20__factory.connect(
    existentTestERC20,
    bridge.l1Signer
  )
  const mintRes = await testToken.mint()
  const mintRec = await mintRes.wait()

  const approveRes = await bridge.approveToken(existentTestERC20)
  const approveRec = await approveRes.wait()

  const data = await bridge.getAndUpdateL1TokenData(existentTestERC20)
  const allowed = data.ERC20 && data.ERC20.allowed
  expect(allowed).to.be.true

  const expectedL1GatewayAddress = await bridge.l1Bridge.getGatewayAddress(
    testToken.address
  )
  const initialBridgeTokenBalance = await testToken.balanceOf(
    expectedL1GatewayAddress
  )

  const depositRes = await bridge.deposit(existentTestERC20, tokenDepositAmount)

  const depositRec = await depositRes.wait()

  const finalBridgeTokenBalance = await testToken.balanceOf(
    expectedL1GatewayAddress
  )

  expect(
    initialBridgeTokenBalance
      .add(tokenDepositAmount)
      .eq(finalBridgeTokenBalance)
  ).to.be.true
  await testRetryableTicket(bridge, depositRec)

  const l2Data = await bridge.getAndUpdateL2TokenData(existentTestERC20)

  const testWalletL2Balance = l2Data && l2Data.ERC20 && l2Data.ERC20.balance

  console.warn('')

  expect(testWalletL2Balance && testWalletL2Balance.eq(tokenDepositAmount)).to
    .be.true
}
