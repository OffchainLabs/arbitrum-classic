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

import { AddressZero } from '@ethersproject/constants'

import {
  AeWETH__factory,
  L1GatewayRouter__factory,
  L2GatewayRouter__factory,
  L2WethGateway__factory,
  L1WethGateway__factory,
  L2CustomGateway__factory,
  L1CustomGateway__factory,
  L2ERC20Gateway__factory,
  L1ERC20Gateway__factory,
} from '../src/lib/abi'

import { instantiateBridge } from '../scripts/instantiate_bridge'
import { existentTestERC20, existentTestCustomToken } from './testHelpers'

const expectIgnoreCase = (expected: string, actual: string) => {
  expect(expected.toLocaleLowerCase()).to.equal(actual.toLocaleLowerCase())
}

describe('sanity checks (read-only)', async () => {
  it('standard gateways public storage vars properly set', async () => {
    const { l1Signer, l2Signer, l2Network } = await instantiateBridge()
    const l1Gateway = await L1ERC20Gateway__factory.connect(
      l2Network.tokenBridge.l1ERC20Gateway,
      l1Signer
    )
    const l2Gateway = await L2ERC20Gateway__factory.connect(
      l2Network.tokenBridge.l2ERC20Gateway,
      l2Signer
    )

    const l1ClonableProxyHash = await l1Gateway.cloneableProxyHash()
    const l2ClonableProxyHash = await l2Gateway.cloneableProxyHash()
    expect(l1ClonableProxyHash).to.equal(l2ClonableProxyHash)

    const l1BeaconProxyHash = await l1Gateway.l2BeaconProxyFactory()
    const l2BeaconProxyHash = await l2Gateway.beaconProxyFactory()
    expect(l1BeaconProxyHash).to.equal(l2BeaconProxyHash)

    const l1GatewayCounterParty = await l1Gateway.counterpartGateway()
    expect(l1GatewayCounterParty).to.equal(l2Network.tokenBridge.l2ERC20Gateway)

    const l2GatewayCounterParty = await l2Gateway.counterpartGateway()
    expect(l2GatewayCounterParty).to.equal(l2Network.tokenBridge.l1ERC20Gateway)

    const l1Router = await l1Gateway.router()
    expect(l1Router).to.equal(l2Network.tokenBridge.l1GatewayRouter)

    const l2Router = await l2Gateway.router()
    expect(l2Router).to.equal(l2Network.tokenBridge.l2GatewayRouter)
  })

  it('custom gateways public storage vars properly set', async () => {
    const { l1Signer, l2Signer, l2Network } = await instantiateBridge()
    const l1Gateway = await L1CustomGateway__factory.connect(
      l2Network.tokenBridge.l1CustomGateway,
      l1Signer
    )
    const l2Gateway = await L2CustomGateway__factory.connect(
      l2Network.tokenBridge.l2CustomGateway,
      l2Signer
    )
    const l1GatewayCounterParty = await l1Gateway.counterpartGateway()
    expect(l1GatewayCounterParty).to.equal(
      l2Network.tokenBridge.l2CustomGateway
    )

    const l2GatewayCounterParty = await l2Gateway.counterpartGateway()
    expect(l2GatewayCounterParty).to.equal(
      l2Network.tokenBridge.l1CustomGateway
    )

    const l1Router = await l1Gateway.router()
    expect(l1Router).to.equal(l2Network.tokenBridge.l1GatewayRouter)

    const l2Router = await l2Gateway.router()
    expect(l2Router).to.equal(l2Network.tokenBridge.l2GatewayRouter)
  })

  it('customtoken gateway properly set', async () => {
    const { l2Network, tokenBridger, l1Signer, l2Signer } =
      await instantiateBridge()

    const l1customGatewayAddress = await tokenBridger.getL1GatewayAddress(
      existentTestCustomToken,
      l1Signer.provider!
    )

    expect(l1customGatewayAddress).to.equal(
      l2Network.tokenBridge.l1CustomGateway
    )

    const l2Address = await tokenBridger.getL2ERC20Address(
      existentTestCustomToken,
      l1Signer.provider!
    )

    const l2CustomGateway = await L2CustomGateway__factory.connect(
      l2Network.tokenBridge.l2CustomGateway,
      l2Signer
    )
    expect(l2Address === AddressZero).to.be.false

    const l2AddressOnGateway = await l2CustomGateway.l1ToL2Token(
      existentTestCustomToken
    )

    expect(l2AddressOnGateway).to.equal(l2Address)
  })

  it('tokens properly set on gateway routers', async () => {
    const { l1Signer, l2Signer, l2Network } = await instantiateBridge()

    const {
      l1ERC20Gateway,
      l1CustomGateway,
      l1WethGateway,
      l2ERC20Gateway,
      l2CustomGateway,
      l2WethGateway,
      l1Weth,
    } = l2Network.tokenBridge

    const l2GatewayRouter = L2GatewayRouter__factory.connect(
      l2Network.tokenBridge.l2GatewayRouter,
      l2Signer.provider!
    )
    const l1GatewayRouter = L1GatewayRouter__factory.connect(
      l2Network.tokenBridge.l1GatewayRouter,
      l1Signer.provider!
    )
    const tokens = [existentTestERC20, existentTestCustomToken, l1Weth]
    const l1Gateways = [l1ERC20Gateway, l1CustomGateway, l1WethGateway]
    const l2Gateways = [l2ERC20Gateway, l2CustomGateway, l2WethGateway]

    for (let i = 0; i < tokens.length; i++) {
      const token = tokens[i]
      const expectedL1Gateway = l1Gateways[i]
      const expectedL2Gateway = l2Gateways[i]
      const l1Gateway = await l1GatewayRouter.getGateway(token)
      const l2Gateway = await l2GatewayRouter.getGateway(token)

      expect(expectedL1Gateway).to.equal(l1Gateway)
      expect(expectedL2Gateway).to.equal(l2Gateway)
    }
  })

  it('weth gateways gateways public storage vars properly set', async () => {
    const { l1Signer, l2Signer, l2Network } = await instantiateBridge()

    const l1Gateway = await L1WethGateway__factory.connect(
      l2Network.tokenBridge.l1WethGateway,
      l1Signer
    )
    const l2Gateway = await L2WethGateway__factory.connect(
      l2Network.tokenBridge.l2WethGateway,
      l2Signer
    )

    const l1Weth = await l1Gateway.l1Weth()
    expectIgnoreCase(l1Weth, l2Network.tokenBridge.l1Weth)

    const l2Weth = await l2Gateway.l2Weth()
    expectIgnoreCase(l2Weth, l2Network.tokenBridge.l2Weth)

    const l1GatewayCounterParty = await l1Gateway.counterpartGateway()
    expectIgnoreCase(l1GatewayCounterParty, l2Network.tokenBridge.l2WethGateway)

    const l2GatewayCounterParty = await l2Gateway.counterpartGateway()
    expectIgnoreCase(l2GatewayCounterParty, l2Network.tokenBridge.l1WethGateway)

    const l1Router = await l1Gateway.router()
    expectIgnoreCase(l1Router, l2Network.tokenBridge.l1GatewayRouter)

    const l2Router = await l2Gateway.router()
    expectIgnoreCase(l2Router, l2Network.tokenBridge.l2GatewayRouter)
  })

  it('aeWETh public vars properly set', async () => {
    const { l2Signer, l2Network } = await instantiateBridge()

    const aeWeth = AeWETH__factory.connect(
      l2Network.tokenBridge.l2Weth,
      l2Signer
    )

    const l2GatewayOnAeWeth = await aeWeth.l2Gateway()
    expectIgnoreCase(l2GatewayOnAeWeth, l2Network.tokenBridge.l2WethGateway)

    const l1AddressOnAeWeth = await aeWeth.l1Address()
    expectIgnoreCase(l1AddressOnAeWeth, l2Network.tokenBridge.l1Weth)
  })

  it('l1 gateway router points to right gateways', async () => {
    const { tokenBridger, l1Signer, l2Network } = await instantiateBridge()

    const gateway = await tokenBridger.getL1GatewayAddress(
      l2Network.tokenBridge.l1Weth,
      l1Signer.provider!
    )

    expect(gateway).to.equal(l2Network.tokenBridge.l1WethGateway)
  })

  it('L1 and L2 implementations of calculateL2ERC20Address match', async () => {
    const { l1Signer, l2Signer, l2Network, tokenBridger } =
      await instantiateBridge()
    const erc20L2AddressAsPerL1 = await tokenBridger.getL2ERC20Address(
      existentTestERC20,
      l1Signer.provider!
    )
    const l2gr = L2GatewayRouter__factory.connect(
      l2Network.tokenBridge.l2GatewayRouter,
      l2Signer.provider!
    )
    const erc20L2AddressAsPerL2 = await l2gr.calculateL2TokenAddress(
      existentTestERC20
    )

    expect(erc20L2AddressAsPerL2).to.equal(erc20L2AddressAsPerL1)
  })
})
