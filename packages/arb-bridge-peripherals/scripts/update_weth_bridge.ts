// import { Bridge } from 'arb-ts/src'
import { ethers } from 'hardhat'

import { ProxyAdmin__factory } from 'arb-ts/src/lib/abi/factories/ProxyAdmin__factory'
import { TransparentUpgradeableProxy__factory } from 'arb-ts/src/lib/abi/factories/TransparentUpgradeableProxy__factory'
import { L1WethGateway__factory } from 'arb-ts/src/lib/abi/factories/L1WethGateway__factory'
import { L2WethGateway__factory } from 'arb-ts/src/lib/abi/factories/L2WethGateway__factory'
import dotenv from 'dotenv'
dotenv.config()
// import RinkebyAddresses from '../deployment-421611.json'

const infuraKey = process.env['INFURA_KEY']
if (!infuraKey) throw new Error('No INFURA_KEY')

const l1Prov = new ethers.providers.JsonRpcProvider(
  'https://mainnet.infura.io/v3/8838d00c028a46449be87e666387c71a'
)
const l2Prov = new ethers.providers.JsonRpcProvider(
  'https://arb1.arbitrum.io/rpc'
)

let l1Signer = ethers.Wallet.fromMnemonic(process.env.DEV_MNEMONIC_1 as string)
let l2Signer = ethers.Wallet.fromMnemonic(process.env.DEV_MNEMONIC_2 as string)

l1Signer = l1Signer.connect(l1Prov)
l2Signer = l2Signer.connect(l2Prov)

const main = async () => {
  const accounts = await ethers.getSigners()

  // The standard arb erc20 uses a beacon proxy, so this update method won't work!
  const l1WethGWProxyAddress = '0xd92023E9d9911199a6711321D1277285e6d4e2db'
  const l2WethGWProxyAddress = '0x6c411aD3E74De3E7Bd422b94A27770f5B86C623B'

  const l1admin = '0x9aD46fac0Cf7f790E5be05A0F15223935A0c0aDa'
  const l2admin = '0xd570aCE65C43af47101fC6250FD6fC63D1c22a86'

  console.warn('admin', l1admin)

  const proxyAdminL1 = await ProxyAdmin__factory.connect(
    l1admin,
    l1Signer
  ).deployed()

  const proxyAdminOwnerL1 = await proxyAdminL1.owner()
  const currAdminl1 = await proxyAdminL1.getProxyAdmin(l1WethGWProxyAddress)

  if (proxyAdminOwnerL1.toLowerCase() !== l1Signer.address.toLowerCase()) {
    console.log('Current admin', currAdminl1)
    console.log('Current account', accounts[0].address)
    throw new Error('Current account does not control ProxyAdmin')
  }

  const proxyAdminL2 = await ProxyAdmin__factory.connect(
    l2admin,
    l2Signer
  ).deployed()

  const proxyAdminOwnerL2 = await proxyAdminL2.owner()
  const currAdminL2 = await proxyAdminL2.getProxyAdmin(l2WethGWProxyAddress)

  if (proxyAdminOwnerL2.toLowerCase() !== l2Signer.address.toLowerCase()) {
    console.log('Current admin', currAdminL2)
    console.log('Current account', accounts[0].address)
    throw new Error('Current account does not control ProxyAdmin')
  }
  console.log('all good')

  // // get factories and deploy logic

  const L1WethGateway = (
    await ethers.getContractFactory('L1WethGateway')
  ).connect(l1Signer)
  const L2WethGateway = (
    await ethers.getContractFactory('L2WethGateway')
  ).connect(l2Signer)

  console.log('deploy L1 weth gateway logic')
  const l1WethGatewayLogic = await L1WethGateway.deploy()
  await l1WethGatewayLogic.deployed()

  console.log('deploy L2 weth gateway logic')
  const l2WethGateway = await L2WethGateway.deploy()
  await l2WethGateway.deployed()

  const upgradeL1Tx = await proxyAdminL1.upgrade(
    l1WethGWProxyAddress,
    l1WethGatewayLogic.address
  )

  const upgradeL1Rec = await upgradeL1Tx.wait()

  const upgradeL2Tx = await proxyAdminL2.upgrade(
    l2WethGWProxyAddress,
    l1WethGatewayLogic.address
  )

  const upgradeL2Rec = await upgradeL2Tx.wait()

  const L1WethGatewayProxy = L1WethGateway__factory.connect(
    l1WethGWProxyAddress,
    l1Signer
  )

  const L2WethGatewayProxy = L2WethGateway__factory.connect(
    l2WethGWProxyAddress,
    l2Signer
  )

  let res = await L1WethGatewayProxy.postUpgradeInit()
  let rec = await res.wait()

  res = await L2WethGatewayProxy.postUpgradeInit2()
  rec = await res.wait()
}

main()
  .then(() => console.log('done'))
  .catch(err => console.log(err))
