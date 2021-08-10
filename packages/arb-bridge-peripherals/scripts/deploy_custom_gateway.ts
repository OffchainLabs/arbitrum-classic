import { ethers } from 'hardhat'
import { instantiateBridge } from 'arb-ts/scripts/instantiate_bridge'

const l1privKey = process.env['L1_PRIVKEY']
if (!l1privKey) throw new Error('No L1_PRIVKEY')

const l2privKey = process.env['L2_PRIVKEY']
if (!l2privKey) throw new Error('No L2_PRIVKEY')

const wait = (ms: number) => new Promise(res => setTimeout(res, ms))

const main = async () => {
  const { bridge, l1Network, l2Network } = await instantiateBridge(
    l1privKey,
    l2privKey
  )

  const { l1Signer, l2Signer } = bridge

  const l1SignerAddress = await l1Signer.getAddress()

  const l1ProxyAdminAddr = l1Network.tokenBridge.l1ProxyAdmin
  const l2ProxyAdminAddr = l2Network.tokenBridge.l2ProxyAdmin
  const l1GatewayRouterAddr = l1Network.tokenBridge.l1GatewayRouter
  const l2GatewayRouterAddr = l2Network.tokenBridge.l2GatewayRouter
  const l1InboxAddr = l1Network.tokenBridge.inbox

  const l1Router = (
    await ethers.getContractAt('L1GatewayRouter', l1GatewayRouterAddr)
  ).connect(l1Signer)
  const l2Router = (
    await ethers.getContractAt('L2GatewayRouter', l2GatewayRouterAddr)
  ).connect(l2Signer)

  //   check if user owns router
  const expectedOwner = await l1Router.owner()

  if (expectedOwner.toLowerCase() !== l1SignerAddress.toLowerCase()) {
    throw new Error('Not router owner')
  }

  // get factories and deploy logic

  const L1CustomGateway = (
    await ethers.getContractFactory('L1CustomGateway')
  ).connect(l1Signer)
  const L2CustomGateway = (
    await ethers.getContractFactory('L2CustomGateway')
  ).connect(l2Signer)

  console.log('deploy L1 Custom gateway logic')
  const l1CustomGateway = await L1CustomGateway.deploy()
  await l1CustomGateway.deployed()

  console.log('deploy L2 Custom gateway logic')
  const l2CustomGateway = await L2CustomGateway.deploy()
  await l2CustomGateway.deployed()

  // deploy proxies

  const L1TransparentUpgradeableProxy = (
    await ethers.getContractFactory('TransparentUpgradeableProxy')
  ).connect(l1Signer)
  const L2TransparentUpgradeableProxy = (
    await ethers.getContractFactory('TransparentUpgradeableProxy')
  ).connect(l2Signer)

  console.log('deploying L1 proxy')
  const l1CustomGatewayProxyDeployment =
    await L1TransparentUpgradeableProxy.deploy(
      l1CustomGateway.address,
      l1ProxyAdminAddr,
      '0x'
    )
  await l1CustomGatewayProxyDeployment.deployed()
  const l1CustomGatewayProxy = L1CustomGateway.attach(
    l1CustomGatewayProxyDeployment.address
  )

  console.log('deploying L2 proxy')
  const l2CustomGatewayProxyDeployment =
    await L2TransparentUpgradeableProxy.deploy(
      l2CustomGateway.address,
      l2ProxyAdminAddr,
      '0x'
    )
  await l2CustomGatewayProxyDeployment.deployed()
  const l2CustomGatewayProxy = L2CustomGateway.attach(
    l2CustomGatewayProxyDeployment.address
  )

  console.log({
    l1CustomGatewayProxy: l1CustomGatewayProxy.address,
    l2CustomGatewayProxy: l2CustomGatewayProxy.address,
  })

  // init
  console.log('init')
  const l1Init = await l1CustomGatewayProxy.initialize(
    l2CustomGatewayProxy.address,
    l1GatewayRouterAddr,
    l1InboxAddr,
    l1SignerAddress
  )

  await l1Init.wait()

  console.log('init L2')
  const l2Init = await l2CustomGatewayProxy.initialize(
    l1CustomGatewayProxy.address,
    l2GatewayRouterAddr
  )
  await l2Init.wait()
}

main()
  .then(() => console.log('done'))
  .catch(err => console.log(err))
