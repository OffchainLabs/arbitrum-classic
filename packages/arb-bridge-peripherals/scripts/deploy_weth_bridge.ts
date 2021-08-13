// import { Bridge } from 'arb-ts/src'
import { ethers } from 'hardhat'

import { instantiateBridge } from 'arb-ts/scripts/instantiate_bridge'

const infuraKey = process.env['INFURA_KEY']
if (!infuraKey) throw new Error('No INFURA_KEY')

const privKey = process.env['DEVNET_PRIVKEY']
if (!privKey) throw new Error('No DEVNET_PRIVKEY')

const signer = new ethers.Wallet(privKey)

const wait = (ms: number) => new Promise(res => setTimeout(res, ms))

// L1 to L2 call parameters
const gasPriceBid = '0'
const maxGas = '800000000'
const maxSubmissionCost = '259829212830'
// TODO: actually calculate the needed value
const deposit = ethers.utils.parseEther('0.01')

const main = async () => {
  const { bridge, l1Network } = await instantiateBridge(privKey, privKey)
  const { tokenBridge } = l1Network
  const { l1Signer, l2Signer } = bridge
  const l1SignerAddress = await l1Signer.getAddress()

  const l1InboxAddr = tokenBridge.inbox
  const l1GatewayRouterAddr = tokenBridge.l1GatewayRouter
  const l2GatewayRouterAddr = tokenBridge.l2GatewayRouter
  const l1ProxyAdminAddr = tokenBridge.l1ProxyAdmin
  const l2ProxyAdminAddr = tokenBridge.l2ProxyAdmin

  const l1WethAddr = tokenBridge.l1Weth
  const l2WethAddr = tokenBridge.l2Weth

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

  const L1WethGateway = (
    await ethers.getContractFactory('L1WethGateway')
  ).connect(l1Signer)
  const L2WethGateway = (
    await ethers.getContractFactory('L2WethGateway')
  ).connect(l2Signer)

  console.log('deploy L1 weth gateway logic')
  const l1WethGateway = await L1WethGateway.deploy()
  await l1WethGateway.deployed()

  console.log('deploy L2 weth gateway logic')
  const l2WethGateway = await L2WethGateway.deploy()
  await l2WethGateway.deployed()

  // deploy proxies

  const L1TransparentUpgradeableProxy = (
    await ethers.getContractFactory('TransparentUpgradeableProxy')
  ).connect(l1Signer)
  const L2TransparentUpgradeableProxy = (
    await ethers.getContractFactory('TransparentUpgradeableProxy')
  ).connect(l2Signer)

  console.log('deploying L1 proxy')
  const l1WethGatewayProxyDeployment =
    await L1TransparentUpgradeableProxy.deploy(
      l1WethGateway.address,
      l1ProxyAdminAddr,
      '0x'
    )
  await l1WethGatewayProxyDeployment.deployed()
  const l1WethGatewayProxy = L1WethGateway.attach(
    l1WethGatewayProxyDeployment.address
  )

  console.log('deploying L2 proxy')
  const l2WethGatewayProxyDeployment =
    await L2TransparentUpgradeableProxy.deploy(
      l2WethGateway.address,
      l2ProxyAdminAddr,
      '0x'
    )
  await l2WethGatewayProxyDeployment.deployed()
  const l2WethGatewayProxy = L2WethGateway.attach(
    l2WethGatewayProxyDeployment.address
  )

  console.log({
    l1WethGatewayProxy: l1WethGatewayProxy.address,
    l2WethGatewayProxy: l2WethGatewayProxy.address,
  })

  // init
  console.log('init proxies')
  console.log('init L1')
  const l1Init = await l1WethGatewayProxy.initialize(
    l2WethGatewayProxy.address,
    l1GatewayRouterAddr,
    l1InboxAddr,
    l1WethAddr,
    l2WethAddr
  )
  await l1Init.wait()

  console.log('init L2')
  const l2Init = await l2WethGatewayProxy.initialize(
    l1WethGatewayProxy.address,
    l2GatewayRouterAddr,
    l1WethAddr,
    l2WethAddr
  )
  await l2Init.wait()

  // set gateway in router

  const setGatewayTx = await l1Router.setGateways(
    [l1WethAddr],
    [l1WethGatewayProxy.address],
    maxGas,
    gasPriceBid,
    maxSubmissionCost,
    {
      value: deposit,
    }
  )

  console.log({ setGatewayTx })

  await setGatewayTx.wait()
}

main()
  .then(() => console.log('done'))
  .catch(err => console.log(err))
