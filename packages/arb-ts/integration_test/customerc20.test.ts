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
import { OutgoingMessageState } from '../src/lib/bridge_helpers'

import chalk from 'chalk'
import {
  fundL1,
  fundL2,
  testRetryableTicket,
  prettyLog,
  warn,
  instantiateBridgeWithRandomWallet,
  fundL2Token,
  tokenFundAmount,
  skipIfMainnet,
  existentTestCustomToken,
} from './testHelpers'
const { Zero, AddressZero } = constants

describe('Custom ERC20', () => {
  beforeEach('skipIfMainnet', function () {
    skipIfMainnet(this)
  })

  it('deposits erc20 (no L2 Eth funding)', async () => {
    const { bridge } = await instantiateBridgeWithRandomWallet()
    await fundL1(bridge)
    await depositTokenTest(bridge)
  })
  it.skip('deposits erc20 (with L2 Eth funding)', async () => {
    const { bridge } = await instantiateBridgeWithRandomWallet()
    await fundL1(bridge)
    await fundL2(bridge)
    await depositTokenTest(bridge)
  })

  it('withdraws erc20', async function () {
    const tokenWithdrawAmount = BigNumber.from(1)
    const { bridge, l2Network } = await instantiateBridgeWithRandomWallet()
    await fundL2(bridge)
    const result = await fundL2Token(bridge, existentTestCustomToken)
    if (!result) {
      warn('Prefunded wallet not funded with tokens; skipping ERC20 withdraw')
      this.skip()
    }
    const withdrawRes = await bridge.withdrawERC20(
      existentTestCustomToken,
      tokenWithdrawAmount
    )
    const withdrawRec = await withdrawRes.wait()

    expect(withdrawRec.status).to.equal(1, 'initiate token withdraw txn failed')
    const withdrawEventData =
      bridge.getWithdrawalsInL2Transaction(withdrawRec)[0]

    expect(withdrawEventData, 'withdrawEventData not found').to.exist

    const outgoingMessageState = await bridge.getOutGoingMessageState(
      withdrawEventData.batchNumber,
      withdrawEventData.indexInBatch
    )
    expect(outgoingMessageState).to.equal(
      OutgoingMessageState.UNCONFIRMED,
      `custom token withdraw getOutGoingMessageState returned ${OutgoingMessageState.UNCONFIRMED}`
    )

    const l2Data = await bridge.getAndUpdateL2TokenData(existentTestCustomToken)
    const testWalletL2Balance = l2Data && l2Data.ERC20 && l2Data.ERC20.balance
    expect(
      testWalletL2Balance &&
        testWalletL2Balance.add(tokenWithdrawAmount).eq(tokenFundAmount),
      'wallet balance not properly deducted after withdraw'
    ).to.be.true
    const walletAddress = await bridge.l1Signer.getAddress()

    const gatewayWithdrawEvents = await bridge.getGatewayWithdrawEventData(
      l2Network.tokenBridge.l2CustomGateway,
      walletAddress,
      { fromBlock: withdrawRec.blockNumber }
    )
    expect(gatewayWithdrawEvents.length).to.equal(
      1,
      'token getGatewayWithdrawEventData query failed'
    )

    const tokenWithdrawEvents = await bridge.getTokenWithdrawEventData(
      existentTestCustomToken,
      walletAddress,
      { fromBlock: withdrawRec.blockNumber }
    )
    expect(tokenWithdrawEvents.length).to.equal(
      1,
      'token getTokenWithdrawEventData query failed'
    )
  })
})

const depositTokenTest = async (bridge: Bridge) => {
  const tokenDepositAmount = BigNumber.from(1)

  const testToken = TestERC20__factory.connect(
    existentTestCustomToken,
    bridge.l1Signer
  )
  const mintRes = await testToken.mint()
  const mintRec = await mintRes.wait()

  const approveRes = await bridge.approveToken(existentTestCustomToken)
  const approveRec = await approveRes.wait()

  const data = await bridge.getAndUpdateL1TokenData(existentTestCustomToken)
  const allowed = data.ERC20 && data.ERC20.allowed
  expect(allowed, 'set token allowance failed').to.be.true

  const expectedL1GatewayAddress = await bridge.l1Bridge.getGatewayAddress(
    testToken.address
  )
  const initialBridgeTokenBalance = await testToken.balanceOf(
    expectedL1GatewayAddress
  )

  const depositRes = await bridge.deposit(
    existentTestCustomToken,
    tokenDepositAmount
  )

  const depositRec = await depositRes.wait()

  const finalBridgeTokenBalance = await testToken.balanceOf(
    expectedL1GatewayAddress
  )

  expect(
    initialBridgeTokenBalance
      .add(tokenDepositAmount)
      .eq(finalBridgeTokenBalance),
    'bridge balance not properly updated after deposit'
  ).to.be.true
  await testRetryableTicket(bridge, depositRec)

  const l2Data = await bridge.getAndUpdateL2TokenData(existentTestCustomToken)

  const testWalletL2Balance = l2Data && l2Data.ERC20 && l2Data.ERC20.balance

  expect(
    testWalletL2Balance && testWalletL2Balance.eq(tokenDepositAmount),
    "l2 wallet balance not properly updated after deposit'"
  ).to.be.true
}
