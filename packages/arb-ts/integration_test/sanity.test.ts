import { expect } from 'chai'
import { L1ERC20Gateway__factory } from '../src/lib/abi/factories/L1ERC20Gateway__factory'
import { L2ERC20Gateway__factory } from '../src/lib/abi/factories/L2ERC20Gateway__factory'
import { L1CustomGateway__factory } from '../src/lib/abi/factories/L1CustomGateway__factory'
import { L2CustomGateway__factory } from '../src/lib/abi/factories/L2CustomGateway__factory'
import { L1WethGateway__factory } from '../src/lib/abi/factories/L1WethGateway__factory'
import { L2WethGateway__factory } from '../src/lib/abi/factories/L2WethGateway__factory'
import { L2GatewayRouter__factory } from '../src/lib/abi/factories/L2GatewayRouter__factory'
import { L1GatewayRouter__factory } from '../src/lib/abi/factories/L1GatewayRouter__factory'

import { TestCustomTokenL1__factory } from '../src/lib/abi/factories/TestCustomTokenL1__factory'
import { IArbToken__factory } from '../src/lib/abi/factories/IArbToken__factory'

import { AeWETH__factory } from '../src/lib/abi/factories/AeWETH__factory'

import { instantiateBridge } from '../scripts/instantiate_bridge'
import { existentTestERC20, existentTestCustomToken } from './testHelpers'
import { constants } from 'ethers'

const expectIgnoreCase = (expected: string, actual: string) => {
  expect(expected.toLocaleLowerCase()).to.equal(actual.toLocaleLowerCase())
}

describe('sanity checks (read-only)', async () => {
  it('standard gateways public storage vars properly set', async () => {
    const { bridge, l1Network, l2Network } = await instantiateBridge()
    const l1Gateway = await L1ERC20Gateway__factory.connect(
      l1Network.tokenBridge.l1ERC20Gateway,
      bridge.l1Bridge.l1Signer
    )
    const l2Gateway = await L2ERC20Gateway__factory.connect(
      l2Network.tokenBridge.l2ERC20Gateway,
      bridge.l2Bridge.l2Signer
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
    expect(l1Router).to.equal(l1Network.tokenBridge.l1GatewayRouter)

    const l2Router = await l2Gateway.router()
    expect(l2Router).to.equal(l2Network.tokenBridge.l2GatewayRouter)
  })

  it('custom gateways public storage vars properly set', async () => {
    const { bridge, l1Network, l2Network } = await instantiateBridge()
    const l1Gateway = await L1CustomGateway__factory.connect(
      l1Network.tokenBridge.l1CustomGateway,
      bridge.l1Bridge.l1Signer
    )
    const l2Gateway = await L2CustomGateway__factory.connect(
      l2Network.tokenBridge.l2CustomGateway,
      bridge.l2Bridge.l2Signer
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
    expect(l1Router).to.equal(l1Network.tokenBridge.l1GatewayRouter)

    const l2Router = await l2Gateway.router()
    expect(l2Router).to.equal(l2Network.tokenBridge.l2GatewayRouter)
  })

  it('customtoken gateway properly set', async () => {
    // @ts-ignore
    const { bridge, l1Network, l2Network } = await instantiateBridge()

    const l1customGatewayAddress = await bridge.l1Bridge.getGatewayAddress(
      existentTestCustomToken
    )

    expect(l1customGatewayAddress).to.equal(
      l1Network.tokenBridge.l1CustomGateway
    )

    const l2Address = await bridge.l1Bridge.getERC20L2Address(
      existentTestCustomToken
    )

    const l2CustomGateway = await L2CustomGateway__factory.connect(
      l1Network.tokenBridge.l2CustomGateway,
      bridge.l2Signer
    )
    expect(l2Address === constants.AddressZero).to.be.false

    const l2AddressOnGateway = await l2CustomGateway.l1ToL2Token(
      existentTestCustomToken
    )

    expect(l2AddressOnGateway).to.equal(l2Address)
  })

  it('tokens properly set on gateway routers', async () => {
    const { bridge, l1Network, l2Network } = await instantiateBridge()

    const {
      l1ERC20Gateway,
      l1CustomGateway,
      l1WethGateway,
      l2ERC20Gateway,
      l2CustomGateway,
      l2WethGateway,
      l1Weth,
    } = l1Network.tokenBridge

    const l2GatewayRouter = L2GatewayRouter__factory.connect(
      l2Network.tokenBridge.l2GatewayRouter,
      bridge.l2Bridge.l2Provider
    )
    const l1GatewayRouter = L1GatewayRouter__factory.connect(
      l2Network.tokenBridge.l1GatewayRouter,
      bridge.l1Bridge.l1Provider
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
    const { bridge, l1Network, l2Network } = await instantiateBridge()

    const l1Gateway = await L1WethGateway__factory.connect(
      l1Network.tokenBridge.l1WethGateway,
      bridge.l1Bridge.l1Signer
    )
    const l2Gateway = await L2WethGateway__factory.connect(
      l2Network.tokenBridge.l2WethGateway,
      bridge.l2Bridge.l2Signer
    )

    const l1Weth = await l1Gateway.l1Weth()
    expectIgnoreCase(l1Weth, l1Network.tokenBridge.l1Weth)

    const l2Weth = await l2Gateway.l2Weth()
    expectIgnoreCase(l2Weth, l1Network.tokenBridge.l2Weth)

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
    const { bridge, l2Network } = await instantiateBridge()

    const aeWeth = AeWETH__factory.connect(
      l2Network.tokenBridge.l2Weth,
      bridge.l2Signer
    )

    const l2GatewayOnAeWeth = await aeWeth.l2Gateway()
    expectIgnoreCase(l2GatewayOnAeWeth, l2Network.tokenBridge.l2WethGateway)

    const l1AddressOnAeWeth = await aeWeth.l1Address()
    expectIgnoreCase(l1AddressOnAeWeth, l2Network.tokenBridge.l1Weth)
  })

  it('l1 gateway router points to right gateways', async () => {
    const { bridge, l1Network, l2Network } = await instantiateBridge()

    const gateway = await bridge.l1Bridge.getGatewayAddress(
      l1Network.tokenBridge.l1Weth
    )

    expect(gateway).to.equal(l1Network.tokenBridge.l1WethGateway)
  })

  it('L1 and L2 implementations of calculateL2ERC20Address match', async () => {
    const { bridge, l2Network } = await instantiateBridge()
    const erc20L2AddressAsPerL1 = await bridge.getERC20L2Address(
      existentTestERC20
    )
    const l2gr = L2GatewayRouter__factory.connect(
      l2Network.tokenBridge.l2GatewayRouter,
      bridge.l2Bridge.l2Provider
    )
    const erc20L2AddressAsPerL2 = await l2gr.calculateL2TokenAddress(
      existentTestERC20
    )

    expect(erc20L2AddressAsPerL2).to.equal(erc20L2AddressAsPerL1)
  })
})
