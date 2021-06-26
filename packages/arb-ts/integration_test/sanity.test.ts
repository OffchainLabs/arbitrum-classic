import { expect } from 'chai'
import { L1ERC20Gateway__factory } from '../src/lib/abi/factories/L1ERC20Gateway__factory'
import { L2ERC20Gateway__factory } from '../src/lib/abi/factories/L2ERC20Gateway__factory'
import { L1CustomGateway__factory } from '../src/lib/abi/factories/L1CustomGateway__factory'
import { L2CustomGateway__factory } from '../src/lib/abi/factories/L2CustomGateway__factory'
import { L1WethGateway__factory } from '../src/lib/abi/factories/L1WethGateway__factory'
import { L2WethGateway__factory } from '../src/lib/abi/factories/L2WethGateway__factory'
import { L2GatewayRouter__factory } from '../src/lib/abi/factories/L2GatewayRouter__factory'

import { AeWETH__factory } from '../src/lib/abi/factories/AeWETH__factory'

import { instantiateBridge } from '../scripts/instantiate_bridge'
import { existentTestERC20 } from './testHelpers'

describe('sanity checks', async () => {
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

    const l1clonableProxyHash = await l1Gateway.cloneableProxyHash()
    const l2clonableProxyHash = await l2Gateway.cloneableProxyHash()
    expect(l1clonableProxyHash).to.equal(l2clonableProxyHash)

    const l1beaconProxyHash = await l1Gateway.l2BeaconProxyFactory()
    const l2beaconProxyHash = await l2Gateway.beaconProxyFactory()
    expect(l1beaconProxyHash).to.equal(l2beaconProxyHash)

    const l1GatewayCoutnerParty = await l1Gateway.counterpartGateway()
    expect(l1GatewayCoutnerParty).to.equal(l2Network.tokenBridge.l2ERC20Gateway)

    const l2GatewayCoutnerParty = await l2Gateway.counterpartGateway()
    expect(l2GatewayCoutnerParty).to.equal(l2Network.tokenBridge.l1ERC20Gateway)

    const l1router = await l1Gateway.router()
    expect(l1router).to.equal(l1Network.tokenBridge.l1GatewayRouter)

    const l2router = await l2Gateway.router()
    expect(l2router).to.equal(l2Network.tokenBridge.l2GatewayRouter)
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
    const l1GatewayCoutnerParty = await l1Gateway.counterpartGateway()
    expect(l1GatewayCoutnerParty).to.equal(
      l2Network.tokenBridge.l2CustomGateway
    )

    const l2GatewayCoutnerParty = await l2Gateway.counterpartGateway()
    expect(l2GatewayCoutnerParty).to.equal(
      l2Network.tokenBridge.l1CustomGateway
    )

    const l1router = await l1Gateway.router()
    expect(l1router).to.equal(l1Network.tokenBridge.l1GatewayRouter)

    const l2router = await l2Gateway.router()
    expect(l2router).to.equal(l2Network.tokenBridge.l2GatewayRouter)
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
    expect(l1Weth).to.equal(l1Network.tokenBridge.l1Weth)

    const l2Weth = await l2Gateway.l2Weth()
    expect(l2Weth).to.equal(l2Network.tokenBridge.l2Weth)

    const l1GatewayCoutnerParty = await l1Gateway.counterpartGateway()
    expect(l1GatewayCoutnerParty).to.equal(l2Network.tokenBridge.l2WethGateway)

    const l2GatewayCoutnerParty = await l2Gateway.counterpartGateway()
    expect(l2GatewayCoutnerParty).to.equal(l2Network.tokenBridge.l1WethGateway)

    const l1router = await l1Gateway.router()
    expect(l1router).to.equal(l1Network.tokenBridge.l1GatewayRouter)

    const l2router = await l2Gateway.router()
    expect(l2router).to.equal(l2Network.tokenBridge.l2GatewayRouter)
  })

  it('aeWETh public vars properly set', async () => {
    const { bridge, l2Network } = await instantiateBridge()

    const aeWeth = AeWETH__factory.connect(
      l2Network.tokenBridge.l2Weth,
      bridge.l2Signer
    )

    const l2GatewayOnAeweth = await aeWeth.l2Gateway()
    expect(l2GatewayOnAeweth).to.equal(l2Network.tokenBridge.l2WethGateway)

    const l1AddressOnAeWeth = await aeWeth.l1Address()
    expect(l1AddressOnAeWeth).to.equal(l2Network.tokenBridge.l1Weth)
  })

  it('l1 gateway router points to right gateways', async () => {
    const { bridge, l1Network, l2Network } = await instantiateBridge()

    const gateway = await bridge.l1Bridge.getGatewayAddress(
      l1Network.tokenBridge.l1Weth
    )

    expect(gateway).to.equal(l1Network.tokenBridge.l1WethGateway)
  })

  it('L1 and L2 implementations of calculateL2ERC20Address match', async () => {
    const { bridge, l1Network, l2Network } = await instantiateBridge()
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
