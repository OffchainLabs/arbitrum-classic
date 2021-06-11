import { ethers } from 'hardhat'
// import deployments from '../deployment.json'
import { providers, Signer } from 'ethers'
import { L1ERC20Gateway__factory } from 'arb-ts/src/lib/abi/factories/L1ERC20Gateway__factory'
import { L2ERC20Gateway__factory } from 'arb-ts/src/lib/abi/factories/L2ERC20Gateway__factory'
import { L1GatewayRouter__factory } from 'arb-ts/src/lib/abi/factories/L1GatewayRouter__factory'
import { L2GatewayRouter__factory } from 'arb-ts/src/lib/abi/factories/L2GatewayRouter__factory'

import { writeFileSync } from 'fs'
// import { writeFileSync } from 'fs'
// import { spawnSync } from 'child_process'

const main = async () => {
  const accounts = await ethers.getSigners()

  // parse needed vars
  const inboxAddress =
    process.env.INBOX_ADDRESS || '0x578BAde599406A8fE3d24Fd7f7211c0911F5B29e'
  if (!inboxAddress) throw new Error('Please set inbox address! INBOX_ADDRESS')

  // set whitelistAddress to address(0) to disable whitelist
  const whitelistAddress = process.env.WHITELIST_ADDRESS
  if (!whitelistAddress)
    throw new Error('Please set whitelist address! WHITELIST_ADDRESS')

  const l2PrivKey = process.env['DEVNET_PRIVKEY']
  if (!l2PrivKey) throw new Error('Missing l2 priv key DEVNET_PRIVKEY')

  const l2ProviderRpc =
    process.env['DEVNET_RPC'] || 'https://rinkeby.arbitrum.io/rpc'
  if (!l2ProviderRpc) throw new Error('Missing l2 rpc DEVNET_RPC')

  // deploy L1 logic contracts
  const L1GatewayRouter = await ethers.getContractFactory('L1GatewayRouter')
  const l1GatewayRouter = await L1GatewayRouter.deploy()
  await l1GatewayRouter.deployed()
  console.log('L1 GatewayRouter logic deployed to:', l1GatewayRouter.address)

  const L1ERC20Gateway = await ethers.getContractFactory('L1ERC20Gateway')
  const l1ERC20Gateway = await L1ERC20Gateway.deploy()
  await l1ERC20Gateway.deployed()
  console.log('L1 ERC20Gateway logic deployed to:', l1ERC20Gateway.address)

  // deploy L1 proxy contracts
  const L1ProxyAdmin = await ethers.getContractFactory('ProxyAdmin')
  const l1ProxyAdmin = await L1ProxyAdmin.deploy()
  await l1ProxyAdmin.deployed()
  console.log('L1 proxy admin at', l1ProxyAdmin.address)

  const L1TransparentUpgradeableProxy = await ethers.getContractFactory(
    'TransparentUpgradeableProxy'
  )

  const l1GatewayRouterProxy = await L1TransparentUpgradeableProxy.deploy(
    l1GatewayRouter.address,
    l1ProxyAdmin.address,
    '0x'
  )
  await l1GatewayRouterProxy.deployed()
  console.log('L1 GatewayRouter Proxy at', l1GatewayRouterProxy.address)

  const l1ERC20GatewayProxy = await L1TransparentUpgradeableProxy.deploy(
    l1ERC20Gateway.address,
    l1ProxyAdmin.address,
    '0x'
  )
  await l1ERC20GatewayProxy.deployed()
  console.log('L1 ERC20Gateway Proxy at', l1ERC20GatewayProxy.address)

  // deploy L2 logic contracts
  const l2Provider = new providers.JsonRpcProvider(l2ProviderRpc)
  const l2Signer = new ethers.Wallet(l2PrivKey, l2Provider)

  const StandardArbERC20 = (
    await ethers.getContractFactory('StandardArbERC20')
  ).connect(l2Signer)
  const standardArbERC20 = await StandardArbERC20.deploy()
  await standardArbERC20.deployed()
  console.log(`erc20 logic at ${standardArbERC20.address}`)

  const UpgradeableBeacon = (
    await ethers.getContractFactory('UpgradeableBeacon')
  ).connect(l2Signer)
  const erc20Beacon = await UpgradeableBeacon.deploy(standardArbERC20.address)
  await erc20Beacon.deployed()
  console.log(`erc20 beacon at ${erc20Beacon.address}`)

  const L2ERC20Gateway = (
    await ethers.getContractFactory('L2ERC20Gateway')
  ).connect(l2Signer)
  const l2ERC20Gateway = await L2ERC20Gateway.deploy()
  await l2ERC20Gateway.deployed()
  console.log('L2 ERC20 gateway logic deployed to:', l2ERC20Gateway.address)

  const L2GatewayRouter = (
    await ethers.getContractFactory('L2GatewayRouter')
  ).connect(l2Signer)
  const l2GatewayRouter = await L2GatewayRouter.deploy()
  await l2GatewayRouter.deployed()
  console.log('L2 gateway router logic deployed to:', l2GatewayRouter.address)

  // deploy L2 proxy contracts
  const L2ProxyAdmin = (await ethers.getContractFactory('ProxyAdmin')).connect(
    l2Signer
  )
  const l2ProxyAdmin = await L2ProxyAdmin.deploy()
  await l2ProxyAdmin.deployed()
  console.log('L2 proxy admin at', l2ProxyAdmin.address)

  const L2TransparentUpgradeableProxy = (
    await ethers.getContractFactory('TransparentUpgradeableProxy')
  ).connect(l2Signer)

  const l2ERC20GatewayProxy = await L2TransparentUpgradeableProxy.deploy(
    l2ERC20Gateway.address,
    l2ProxyAdmin.address,
    '0x'
  )
  await l2ERC20GatewayProxy.deployed()
  console.log('L2 ERC20Gateway Proxy at', l2ERC20GatewayProxy.address)

  const l2GatewayRouterProxy = await L2TransparentUpgradeableProxy.deploy(
    l2GatewayRouter.address,
    l2ProxyAdmin.address,
    '0x'
  )
  await l2GatewayRouterProxy.deployed()
  console.log('L2 Router Proxy at', l2GatewayRouterProxy.address)

  // initialize proxies and setup txs

  const l1ERC20GatewayConnectedAsProxy = L1ERC20Gateway__factory.connect(
    l1ERC20GatewayProxy.address,
    accounts[0]
  )

  const initL1Bridge = await l1ERC20GatewayConnectedAsProxy.initialize(
    l2ERC20GatewayProxy.address,
    l1GatewayRouterProxy.address,
    inboxAddress
  )
  console.log('init L1 hash', initL1Bridge.hash)
  await initL1Bridge.wait()
  console.log('l1 bridge proxy initted')

  const l2ERC20GatewayConnectedAsProxy = L2ERC20Gateway__factory.connect(
    l2ERC20GatewayProxy.address,
    l2Signer
  )

  const initL2Bridge = await l2ERC20GatewayConnectedAsProxy.initialize(
    l1ERC20GatewayProxy.address,
    l2GatewayRouterProxy.address,
    erc20Beacon.address
  )
  console.log('init L2 hash', initL2Bridge.hash)
  await initL2Bridge.wait()
  console.log('l2 bridge proxy initted')

  // TODO: set default gateway to address(0) instead of standardERC20

  const l1DefaultGateway = l1ERC20GatewayProxy.address
  const l2DefaultGateway = l2ERC20GatewayProxy.address

  const l1GatewayRouterConnected = L1GatewayRouter__factory.connect(
    l1GatewayRouterProxy.address,
    accounts[0]
  )
  const initL1RouterTx = await l1GatewayRouterConnected.initialize(
    accounts[0].address,
    l1DefaultGateway,
    whitelistAddress,
    l2GatewayRouterProxy.address,
    inboxAddress
  )
  console.log('init L1 Router hash', initL1RouterTx.hash)
  await initL1RouterTx.wait()
  console.log('l1 router proxy initted')

  const l2GatewayRouterConnectedAtProxy = L2GatewayRouter__factory.connect(
    l2GatewayRouterProxy.address,
    l2Signer
  )

  const initL2Router = await l2GatewayRouterConnectedAtProxy.initialize(
    l1GatewayRouterProxy.address,
    l2DefaultGateway
  )
  console.log('init L2 Router hash', initL2Router.hash)
  await initL2Router.wait()
  console.log('l2 router proxy initted')

  console.log('Done.')

  const contracts = JSON.stringify(
    {
      l1GatewayRouter: l1GatewayRouterProxy.address,
      l2GatewayRouter: l2GatewayRouterProxy.address,
      l1ERC20GatewayProxy: l1ERC20GatewayProxy.address,
      l2ERC20GatewayProxy: l2ERC20GatewayProxy.address,
      l1ProxyAdmin: l1ProxyAdmin.address,
      l2ProxyAdmin: l2ProxyAdmin.address,
      l1Deployer: accounts[0].address,
      l2Deployer: l2Signer.address,
      inbox: inboxAddress,
    },
    null,
    4
  )

  const chainId = l2Provider.network.chainId
  const deployFilePath = `./deployment-${chainId}.json`
  console.log(`Writing to JSON at ${deployFilePath}`)
  writeFileSync(deployFilePath, contracts)
  console.log('Wrote to deployments.json')

  // TODO: transfer admin proxy ownership
  // TODO: transfer gateways and router ownership
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
