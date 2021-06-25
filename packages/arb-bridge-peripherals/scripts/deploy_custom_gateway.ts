import { ethers } from 'hardhat'

import MainnetDeployments from '../deployment-42161.json'

const infuraKey = process.env['INFURA_KEY']
if (!infuraKey) throw new Error('No INFURA_KEY')

const l1Prov = new ethers.providers.JsonRpcProvider(
  'https://mainnet.infura.io/v3/' + infuraKey
)
const l2Prov = new ethers.providers.JsonRpcProvider(
  'https://arb1.arbitrum.io/rpc'
)

const l1privKey = process.env['L1_PRIVKEY']
if (!l1privKey) throw new Error('No L1_PRIVKEY')

const l2privKey = process.env['L2_PRIVKEY']
if (!l2privKey) throw new Error('No L2_PRIVKEY')

const L1Signer = ethers.Wallet.fromMnemonic(l1privKey)
const L2Signer = ethers.Wallet.fromMnemonic(l2privKey)

const l1Signer = L1Signer.connect(l1Prov)
const l2Signer = L2Signer.connect(l2Prov)

const wait = (ms: number) => new Promise(res => setTimeout(res, ms))

const l1InboxAddr = MainnetDeployments.inbox
const l1GatewayRouterAddr = MainnetDeployments.l1GatewayRouter
const l2GatewayRouterAddr = MainnetDeployments.l2GatewayRouter
const l1ProxyAdminAddr = '0x9aD46fac0Cf7f790E5be05A0F15223935A0c0aDa'
const l2ProxyAdminAddr = '0xd570aCE65C43af47101fC6250FD6fC63D1c22a86'

const main = async () => {
  const l1Router = (
    await ethers.getContractAt('L1GatewayRouter', l1GatewayRouterAddr)
  ).connect(l1Signer)
  const l2Router = (
    await ethers.getContractAt('L2GatewayRouter', l2GatewayRouterAddr)
  ).connect(l2Signer)

  //   check if user owns router
  const expectedOwner = await l1Router.owner()

  if (expectedOwner.toLowerCase() !== l1Signer.address.toLowerCase()) {
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
  let l1CustomGatewayProxy = await L1TransparentUpgradeableProxy.deploy(
    l1CustomGateway.address,
    l1ProxyAdminAddr,
    '0x'
  )
  await l1CustomGatewayProxy.deployed()
  l1CustomGatewayProxy = L1CustomGateway.attach(l1CustomGatewayProxy.address)

  console.log('deploying L2 proxy')
  let l2CustomGatewayProxy = await L2TransparentUpgradeableProxy.deploy(
    l2CustomGateway.address,
    l2ProxyAdminAddr,
    '0x'
  )
  await l2CustomGatewayProxy.deployed()
  l2CustomGatewayProxy = L2CustomGateway.attach(l2CustomGatewayProxy.address)

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
    l1Signer.address
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
