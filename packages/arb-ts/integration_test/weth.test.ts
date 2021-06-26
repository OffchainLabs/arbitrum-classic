import { utils, constants } from 'ethers'

import { expect } from 'chai'
import { AeWETH__factory } from '../src/lib/abi/factories/AeWETH__factory'

import { testRetryableTicket, prettyLog, warn } from './testHelpers'
const { Zero, AddressZero } = constants
import {
  instantiateRandomBridge,
  fundL1,
  wait,
  fundL2,
  preFundAmount,
  skipIfMainnet,
} from './testHelpers'

describe('WETH', async () => {
  beforeEach('skipIfMainnet', function () {
    skipIfMainnet(this)
  })

  it('withdraws WETH', async () => {
    const wethToWrap = utils.parseEther('0.00001')
    const wethToWithdraw = utils.parseEther('0.00000001')

    const { bridge, l1Network, l2Network } = await instantiateRandomBridge()
    await fundL2(bridge)

    const l2Weth = AeWETH__factory.connect(
      l2Network.tokenBridge.l2Weth,
      bridge.l2Bridge.l2Signer
    )
    const res = await l2Weth.deposit({
      value: wethToWrap,
    })
    const rec = await res.wait()
    expect(rec.status).to.equal(1)

    const withdrawRes = await bridge.withdrawERC20(
      l1Network.tokenBridge.l1Weth,
      wethToWithdraw
    )
    const withdrawRec = await withdrawRes.wait()
    expect(withdrawRec.status).to.equal(1)
    const withdrawEventData = (
      await bridge.getWithdrawalsInL2Transaction(withdrawRec)
    )[0]

    expect(withdrawEventData).to.exist

    const _l2WethBalance = await bridge.getAndUpdateL2TokenData(
      l2Network.tokenBridge.l1Weth
    )
    const l2WethBalance =
      _l2WethBalance && _l2WethBalance.ERC20 && _l2WethBalance.ERC20.balance
    expect(l2WethBalance && l2WethBalance.add(wethToWithdraw).eq(wethToWrap)).to
      .be.true
  })

  it('deposits WETH', async () => {
    const { bridge, l1Network, l2Network } = await instantiateRandomBridge()
    const l1WethAddress = l1Network.tokenBridge.l1Weth

    const wethToWrap = utils.parseEther('0.0001')
    const wethToDeposit = utils.parseEther('0.00001')

    await fundL1(bridge)

    const l1WETH = AeWETH__factory.connect(
      l1Network.tokenBridge.l1Weth,
      bridge.l1Signer
    )
    const res = await l1WETH.deposit({
      value: wethToWrap,
    })
    const rec = await res.wait()
    prettyLog('wrapped some ether')

    const approveRes = await bridge.approveToken(l1WethAddress)
    const approveRec = await approveRes.wait()
    expect(approveRec.status).to.equal(1)

    const data = await bridge.getAndUpdateL1TokenData(l1WethAddress)
    const allowed = data.ERC20 && data.ERC20.allowed
    expect(allowed).to.be.true

    const depositRes = await bridge.deposit(l1WethAddress, wethToDeposit)
    const depositRec = await depositRes.wait()
    await testRetryableTicket(bridge, depositRec)

    const l2Data = await bridge.getAndUpdateL2TokenData(l1WethAddress)

    const testWalletL2Balance = l2Data && l2Data.ERC20 && l2Data.ERC20.balance

    expect(testWalletL2Balance && testWalletL2Balance.eq(wethToDeposit)).to.be
      .true
  })
})
