// import { Bridge } from 'arb-ts/src'
import { ethers } from 'hardhat'

import { ProxyAdmin__factory } from 'arb-ts/src/lib/abi/factories/ProxyAdmin__factory'

import RinkebyAddresses from '../deployment-421611.json'

const l1Prov = new ethers.providers.JsonRpcProvider(
  'https://rinkeby.infura.io/v3/c13a0d6955b14bf181c924bf4c7797fc'
)
const l2Prov = new ethers.providers.JsonRpcProvider(
  'https://rinkeby.arbitrum.io/rpc'
)

const privKey = process.env['DEVNET_PRIVKEY']
if (!privKey) throw new Error('No DEVNET_PRIVKEY')

const signer = new ethers.Wallet(privKey)

const l1Signer = signer.connect(l1Prov)
const l2Signer = signer.connect(l2Prov)

const wait = (ms: number) => new Promise(res => setTimeout(res, ms))

const l1InboxAddr = RinkebyAddresses.inbox
const l1GatewayRouterAddr = RinkebyAddresses.l1GatewayRouter
const l2GatewayRouterAddr = RinkebyAddresses.l2GatewayRouter
const l1ProxyAdminAddr = '0x0DbAF24efA2bc9Dd1a6c0530DD252BCcF883B89A'
const l2ProxyAdminAddr = '0x58816566EB91815Cc07f3Ad5230eE0820fe1A19a'

const l1WethAddr = '0xc778417e063141139fce010982780140aa0cd5ab'
const l2WethAddr = '0xB47e6A5f8b33b3F17603C83a0535A9dcD7E32681'

// L1 to L2 call parameters
const gasPriceBid = '0'
const maxGas = '800000000'
const maxSubmissionCost = '259829212830'
// TODO: actually calculate the needed value
const deposit = ethers.utils.parseEther('0.01')

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
  let l1WethGatewayProxy = await L1TransparentUpgradeableProxy.deploy(
    l1WethGateway.address,
    l1ProxyAdminAddr,
    '0x'
  )
  await l1WethGatewayProxy.deployed()
  l1WethGatewayProxy = L1WethGateway.attach(l1WethGatewayProxy.address)

  console.log('deploying L2 proxy')
  let l2WethGatewayProxy = await L2TransparentUpgradeableProxy.deploy(
    l2WethGateway.address,
    l2ProxyAdminAddr,
    '0x'
  )
  await l2WethGatewayProxy.deployed()
  l2WethGatewayProxy = L2WethGateway.attach(l2WethGatewayProxy.address)

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
