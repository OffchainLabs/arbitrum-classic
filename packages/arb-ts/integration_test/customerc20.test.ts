/*
 * Copyright 2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/* eslint-env node */
'use strict'

import { expect } from 'chai'

import { BigNumber } from '@ethersproject/bignumber'
import { Logger, LogLevel } from '@ethersproject/logger'
Logger.setLogLevel(LogLevel.ERROR)

import { L1CustomGateway__factory } from '../src/lib/abi/factories/L1CustomGateway__factory'
import { L1GatewayRouter__factory } from '../src/lib/abi/factories/L1GatewayRouter__factory'
import { L2GatewayRouter__factory } from '../src/lib/abi/factories/L2GatewayRouter__factory'
import { TestArbCustomToken__factory } from '../src/lib/abi/factories/TestArbCustomToken__factory'
import { TestCustomTokenL1__factory } from '../src/lib/abi/factories/TestCustomTokenL1__factory'
import { TestERC20__factory } from '../src/lib/abi/factories/TestERC20__factory'

import { L2ToL1MessageStatus } from '../src/lib/message/L2ToL1Message'

import {
  fundL1,
  fundL2,
  testRetryableTicket,
  warn,
  instantiateBridgeWithRandomWallet,
  fundL2Token,
  tokenFundAmount,
  skipIfMainnet,
  existentTestCustomToken,
} from './testHelpers'
import { L1ToL2MessageStatus, Erc20Bridger } from '../src'
import { Signer, constants } from 'ethers'
import { parseEther } from 'ethers/lib/utils'
import { SignerProviderUtils } from '../src/lib/dataEntities/signerOrProvider'

describe('Custom ERC20', () => {
  beforeEach('skipIfMainnet', function () {
    skipIfMainnet(this)
  })

  it('deposits erc20 (no L2 Eth funding)', async () => {
    const { l1Signer, l2Signer, erc20Bridger } =
      instantiateBridgeWithRandomWallet()
    await fundL1(l1Signer)
    await depositTokenTest(erc20Bridger, l1Signer, l2Signer)
  })
  it.skip('deposits erc20 (with L2 Eth funding)', async () => {
    const { l1Signer, l2Signer, erc20Bridger } =
      instantiateBridgeWithRandomWallet()
    await fundL1(l1Signer)
    await fundL2(l2Signer)
    await depositTokenTest(erc20Bridger, l1Signer, l2Signer)
  })

  it('withdraws erc20', async function () {
    const tokenWithdrawAmount = BigNumber.from(1)
    const { l2Network, l2Signer, l1Signer, erc20Bridger } =
      instantiateBridgeWithRandomWallet()

    await fundL2(l2Signer)
    const result = await fundL2Token(
      l1Signer.provider!,
      l2Signer,
      erc20Bridger,
      existentTestCustomToken
    )
    if (!result) {
      warn('Prefunded wallet not funded with tokens; skipping ERC20 withdraw')
      this.skip()
    }

    const withdrawRes = await erc20Bridger.withdraw({
      amount: tokenWithdrawAmount,
      erc20l1Address: existentTestCustomToken,
      l2Signer: l2Signer,
    })
    const withdrawRec = await withdrawRes.wait()

    expect(withdrawRec.status).to.equal(1, 'initiate token withdraw txn failed')

    const message = (
      await withdrawRec.getL2ToL1Messages(l1Signer.provider!, l2Network)
    )[0]
    expect(message, 'withdrawEventData not found').to.exist

    const messageStatus = await message.status(null)
    expect(
      messageStatus === L2ToL1MessageStatus.UNCONFIRMED ||
        messageStatus === L2ToL1MessageStatus.NOT_FOUND,
      `custom token withdraw status returned ${messageStatus}`
    ).to.be.true

    const l2Token = erc20Bridger.getL2TokenContract(
      l2Signer.provider!,
      await erc20Bridger.getL2ERC20Address(
        existentTestCustomToken,
        l1Signer.provider!
      )
    )
    const testWalletL2Balance = (
      await l2Token.functions.balanceOf(await l2Signer.getAddress())
    )[0]
    expect(
      testWalletL2Balance.add(tokenWithdrawAmount).eq(tokenFundAmount),
      'wallet balance not properly deducted after withdraw'
    ).to.be.true
    const walletAddress = await l1Signer.getAddress()

    const gatewayWithdrawEvents = await erc20Bridger.getL2WithdrawalEvents(
      l2Signer.provider!,
      erc20Bridger.l2Network.tokenBridge.l2CustomGateway,
      { fromBlock: withdrawRec.blockNumber, toBlock: 'latest' },
      undefined,
      walletAddress
    )
    expect(gatewayWithdrawEvents.length).to.equal(
      1,
      'token standard gateway query failed'
    )

    const gatewayAddress = await erc20Bridger.getL2GatewayAddress(
      existentTestCustomToken,
      l2Signer.provider!
    )
    const tokenWithdrawEvents = await erc20Bridger.getL2WithdrawalEvents(
      l2Signer.provider!,
      gatewayAddress,
      { fromBlock: withdrawRec.blockNumber, toBlock: 'latest' },
      existentTestCustomToken,
      walletAddress
    )
    expect(tokenWithdrawEvents.length).to.equal(
      1,
      'token filtered query failed'
    )
  })

  it('register custom token', async () => {
    const { l2Network, l2Signer, l1Signer, adminErc20Bridger } =
      instantiateBridgeWithRandomWallet()

    await fundL1(l1Signer, parseEther('0.01'))
    await fundL2(l2Signer, parseEther('0.01'))
    const l1SignerAddr = await l1Signer.getAddress()
    const l2SignerAddr = await l2Signer.getAddress()
    const sendAmount = 137
    const startTokenBalance = 50000000
    const l2Provider = SignerProviderUtils.getProviderOrThrow(l2Signer)

    // create a custom token on L1 and L2
    const l1CustomTokenFac = new TestCustomTokenL1__factory(l1Signer)
    const l1CustomToken = await l1CustomTokenFac.deploy(
      l2Network.tokenBridge.l1CustomGateway,
      l2Network.tokenBridge.l1GatewayRouter
    )
    // mint ourselves some tokens and approve the custom gateway
    await l1CustomToken.mint()
    await l1CustomToken.approve(
      l2Network.tokenBridge.l1CustomGateway,
      sendAmount
    )
    const l2CustomTokenFac = new TestArbCustomToken__factory(l2Signer)
    const l2CustomToken = await l2CustomTokenFac.deploy(
      l2Network.tokenBridge.l2CustomGateway,
      l1CustomToken.address
    )

    // check starting conditions - should initially use the default gateway
    const l1GatewayRouter = new L1GatewayRouter__factory(l1Signer).attach(
      l2Network.tokenBridge.l1GatewayRouter
    )
    const l2GatewayRouter = new L2GatewayRouter__factory(l2Signer).attach(
      l2Network.tokenBridge.l2GatewayRouter
    )
    const l1CustomGateway = new L1CustomGateway__factory(l1Signer).attach(
      l2Network.tokenBridge.l1CustomGateway
    )
    const l2CustomGateway = new L1CustomGateway__factory(l2Signer).attach(
      l2Network.tokenBridge.l2CustomGateway
    )
    const startL1GatewayAddress = await l1GatewayRouter.l1TokenToGateway(
      l1CustomToken.address
    )
    expect(
      startL1GatewayAddress,
      'Start l1GatewayAddress not equal empty address'
    ).to.eq(constants.AddressZero)
    const startL2GatewayAddress = await l2GatewayRouter.l1TokenToGateway(
      l2CustomToken.address
    )
    expect(
      startL2GatewayAddress,
      'Start l2GatewayAddress not equal empty address'
    ).to.eq(constants.AddressZero)
    const startL1Erc20Address = await l1CustomGateway.l1ToL2Token(
      l1CustomToken.address
    )
    expect(
      startL1Erc20Address,
      'Start l1Erc20Address not equal empty address'
    ).to.eq(constants.AddressZero)
    const startL2Erc20Address = await l2CustomGateway.l1ToL2Token(
      l1CustomToken.address
    )
    expect(
      startL2Erc20Address,
      'Start l2Erc20Address not equal empty address'
    ).to.eq(constants.AddressZero)
    const l1StartBalance = await l1CustomToken.balanceOf(l1SignerAddr)
    expect(l1StartBalance.toNumber(), 'Wrong L1 start balance').to.eq(
      startTokenBalance
    )
    const l2StartBalance = await l2CustomToken.balanceOf(l2SignerAddr)
    expect(l2StartBalance.toNumber(), 'Wrong L2 start balance').to.eq(
      constants.Zero.toNumber()
    )

    // send the messages
    const regTx = await adminErc20Bridger.registerCustomToken(
      l1CustomToken.address,
      l2CustomToken.address,
      l1Signer,
      l2Provider
    )
    const regRec = await regTx.wait()

    // wait on messages
    const l1ToL2Messages = await regRec.getL1ToL2Messages(l2Provider)
    expect(l1ToL2Messages.length, 'Should be 2 messages.').to.eq(2)

    const setTokenTx = await l1ToL2Messages[0].waitForStatus()
    expect(setTokenTx.status, 'Set token not redeemed.').to.eq(
      L1ToL2MessageStatus.REDEEMED
    )
    const setGateways = await l1ToL2Messages[1].waitForStatus()
    expect(setGateways.status, 'Set gateways not redeemed.').to.eq(
      L1ToL2MessageStatus.REDEEMED
    )

    // send a deposit to follow
    const depositTx = await adminErc20Bridger.deposit({
      amount: BigNumber.from(sendAmount),
      erc20L1Address: l1CustomToken.address,
      l1Signer: l1Signer,
      l2Provider: l2Provider,
    })
    const depositRec = await depositTx.wait()

    const depositStatus = await depositRec.waitForL2(l2Provider)
    expect(depositStatus.complete, 'Deposit is not complete').to.eq(true)
    expect(depositStatus.status, 'Deposit is not redeemed').to.eq(
      L1ToL2MessageStatus.REDEEMED
    )

    // check end conditions
    const endL1GatewayAddress = await l1GatewayRouter.l1TokenToGateway(
      l1CustomToken.address
    )
    expect(
      endL1GatewayAddress,
      'End l1GatewayAddress not equal to l1 custom gateway'
    ).to.eq(l2Network.tokenBridge.l1CustomGateway)
    const endL2GatewayAddress = await l1GatewayRouter.l1TokenToGateway(
      l2CustomToken.address
    )
    expect(
      endL2GatewayAddress,
      'End l2GatewayAddress not equal to l1 custom gateway'
    ).to.eq(l2Network.tokenBridge.l1CustomGateway)

    const endL1Erc20Address = await l1CustomGateway.l1ToL2Token(
      l1CustomToken.address
    )
    expect(
      endL1Erc20Address,
      'End l1Erc20Address not equal l1CustomToken address'
    ).to.eq(l2CustomToken.address)
    const endL2Erc20Address = await l2CustomGateway.l1ToL2Token(
      l1CustomToken.address
    )
    expect(
      endL2Erc20Address,
      'End l2Erc20Address not equal l2CustomToken address'
    ).to.eq(l2CustomToken.address)
    const l1EndBalance = await l1CustomToken.balanceOf(l1SignerAddr)
    expect(l1EndBalance.toNumber(), 'Wrong L1 end balance').to.eq(
      startTokenBalance - sendAmount
    )
    const l2EndBalance = await l2CustomToken.balanceOf(l2SignerAddr)
    expect(l2EndBalance.toNumber(), 'Wrong L2 end balance').to.eq(sendAmount)
  })
})

const depositTokenTest = async (
  erc20Bridger: Erc20Bridger,
  l1Signer: Signer,
  l2Signer: Signer
) => {
  const tokenDepositAmount = BigNumber.from(1)
  const testToken = TestERC20__factory.connect(
    existentTestCustomToken,
    l1Signer
  )
  const mintRes = await testToken.mint()
  await mintRes.wait()

  const approveRes = await erc20Bridger.approveToken({
    erc20L1Address: existentTestCustomToken,
    l1Signer: l1Signer,
  })
  await approveRes.wait()

  const expectedL1GatewayAddress = await erc20Bridger.getL1GatewayAddress(
    testToken.address,
    l1Signer.provider!
  )
  const l1Token = erc20Bridger.getL1TokenContract(
    l1Signer.provider!,
    existentTestCustomToken
  )
  const allowance = (
    await l1Token.functions.allowance(
      await l1Signer.getAddress(),
      expectedL1GatewayAddress
    )
  )[0]

  expect(allowance.eq(Erc20Bridger.MAX_APPROVAL), 'set token allowance failed')
    .to.be.true

  const initialBridgeTokenBalance = await testToken.balanceOf(
    expectedL1GatewayAddress
  )

  const depositRes = await erc20Bridger.deposit({
    erc20L1Address: existentTestCustomToken,
    amount: tokenDepositAmount,
    l1Signer: l1Signer,
    l2Provider: l2Signer.provider!,
  })

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
  await testRetryableTicket(l2Signer.provider!, depositRec)

  const l2Token = erc20Bridger.getL2TokenContract(
    l2Signer.provider!,
    await erc20Bridger.getL2ERC20Address(
      existentTestCustomToken,
      l1Signer.provider!
    )
  )
  const testWalletL2Balance = (
    await l2Token.functions.balanceOf(await l2Signer.getAddress())
  )[0]

  expect(
    testWalletL2Balance.eq(tokenDepositAmount),
    "l2 wallet balance not properly updated after deposit'"
  ).to.be.true
}
