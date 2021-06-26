// import { Bridge } from 'arb-ts/src'
import { ethers } from 'hardhat'

// import { ProxyAdmin__factory } from 'arb-ts/src/lib/abi/factories/ProxyAdmin__factory'
import { networks } from 'arb-ts/src/lib/networks'

import MainnetAddresses from '../deployment-42161.json'
// import RinkebyAddresses from '../deployment-421611.json'

const POST_UPGRADE_INIT = "0x95fcea78"

const infuraKey = process.env['INFURA_KEY']
if (!infuraKey) throw new Error('No INFURA_KEY')

const l1Prov = new ethers.providers.JsonRpcProvider('http://127.0.0.1:8545')
// const l1Prov = new ethers.providers.JsonRpcProvider('https://mainnet.infura.io/v3/' + infuraKey)
// const l2Prov = new ethers.providers.JsonRpcProvider('https://arb1.arbitrum.io/rpc')

const l1PrivKey = process.env['L1_PRIVKEY']
if (!l1PrivKey) throw new Error('No L1_PRIVKEY')
// const l2PrivKey = process.env['L2_PRIVKEY']
// if (!l2PrivKey) throw new Error('No L2_PRIVKEY')

const l1Wallet = ethers.Wallet.fromMnemonic(l1PrivKey)
// const l2Wallet = ethers.Wallet.fromMnemonic(l2PrivKey)

const l1ProxyAdminAddr = MainnetAddresses.l1ProxyAdmin
const l2ProxyAdminAddr = MainnetAddresses.l2ProxyAdmin

// const l1Prov = new ethers.providers.JsonRpcProvider('https://rinkeby.infura.io/v3/' + infuraKey)
// const l2Prov = new ethers.providers.JsonRpcProvider('https://rinkeby.arbitrum.io/rpc')

// const l1ProxyAdminAddr = RinkebyAddresses.l1ProxyAdmin
// const l2ProxyAdminAddr = RinkebyAddresses.l2ProxyAdmin

// const privKey = process.env['DEVNET_PRIVKEY']
// if (!privKey) throw new Error('No DEVNET_PRIVKEY')

// const l1Wallet = new ethers.Wallet(privKey)
// const l2Wallet = new ethers.Wallet(privKey)

const l1Signer = l1Wallet.connect(l1Prov)
// const l2Signer = l2Wallet.connect(l2Prov)

const wait = (ms: number) => new Promise(res => setTimeout(res, ms))

const main = async () => {
  // const l1Network = await l1Prov.getNetwork()
  // const tokenbridge = networks[l1Network.chainId].tokenBridge
  const tokenbridge = networks['1'].tokenBridge
  const l1ERC20GatewayProxy = tokenbridge.l1ERC20Gateway
  const l2ERC20GatewayProxy = tokenbridge.l2ERC20Gateway
  const l1CustomGatewayProxy = tokenbridge.l1CustomGateway
  const l2CustomGatewayProxy = tokenbridge.l2CustomGateway

  //   upgrade contracts
  // l1 validation
  const ProxyAdmin__factory = await ethers.getContractFactory("ProxyAdmin")

  const l1ProxyAdmin = ProxyAdmin__factory.attach(l1ProxyAdminAddr).connect(l1Signer)

  const l1Admin = await l1ProxyAdmin.owner()

  if (l1Admin.toLowerCase() !== l1Wallet.address.toLowerCase()) {
    throw new Error('Wrong account')
  }

  // l2 validation
  // const l2ProxyAdmin = ProxyAdmin__factory.attach(l2ProxyAdminAddr).connect(l2Signer)

  // const l2Admin = await l2ProxyAdmin.owner()

  // if (l2Admin.toLowerCase() !== l2Wallet.address.toLowerCase()) {
  //   throw new Error('Wrong account')
  // }

  console.log('Data validated.')

  // deploy new logics
  console.log('deploying new L1 logic std')
  const l1Erc20Factory = await (
    await ethers.getContractFactory('L1ERC20Gateway')
  ).connect(l1Signer)
  const newl1Erc20 = await l1Erc20Factory.deploy()
  await newl1Erc20.deployed()

  console.log('deploying new L1 logic custom')
  const l1CustomFactory = await (
    await ethers.getContractFactory('L1CustomGateway')
  ).connect(l1Signer)
  const newl1Custom = await l1CustomFactory.deploy()
  await newl1Custom.deployed()


  // console.log('deploying new L2 logic std')
  // const l2Erc20Factory = await (
  //   await ethers.getContractFactory('L2ERC20Gateway')
  // ).connect(l2Signer)
  // const newL2Erc20 = await l2Erc20Factory.deploy()
  // await newL2Erc20.deployed()


  // console.log('deploying new L2 logic custom')
  // const l2CustomFactory = await (
  //   await ethers.getContractFactory('L2CustomGateway')
  // ).connect(l2Signer)
  // const newl2Custom = await l2CustomFactory.deploy()
  // await newl2Custom.deployed()

  // console.log('deploying new L2 logic weth')
  // const l2WethFactory = await (
  //   await ethers.getContractFactory('L2WethGateway')
  // ).connect(l2Signer)
  // const newl2Weth = await l2WethFactory.deploy()
  // await newl2Weth.deployed()

  // perform upgrade

  console.log('upgrading L1 std')
  const l1tx = await l1ProxyAdmin.upgradeAndCall(
    l1ERC20GatewayProxy,
    newl1Erc20.address,
    POST_UPGRADE_INIT
  )
  console.log({ l1tx })
  const l1receipt = await l1tx.wait()
  console.log({ l1receipt })

  console.log('upgrading L1 custom')
  const l1txcustom = await l1ProxyAdmin.upgradeAndCall(
    l1CustomGatewayProxy,
    newl1Custom.address,
    POST_UPGRADE_INIT
  )
  console.log({ l1txcustom })
  const l1receiptcustom = await l1txcustom.wait()
  console.log({ l1receiptcustom })

  // // console.log('upgrading L2 std')
  // // const l2tx = await l2ProxyAdmin.upgradeAndCall(
  // //   l2ERC20GatewayProxy,
  // //   newL2Erc20.address,
  // //   POST_UPGRADE_INIT
  // // )
  // // console.log({ l2tx })
  // // const l2receipt = await l2tx.wait()
  // // console.log({ l2receipt })

  // // console.log('upgrading L2 custom')
  // // const l2txcustom = await l2ProxyAdmin.upgradeAndCall(
  // //   l2CustomGatewayProxy,
  // //   newl2Custom.address,
  // //   POST_UPGRADE_INIT
  // // )
  // console.log({ l2txcustom })
  // const l2receiptcustom = await l2txcustom.wait()
  // console.log({ l2receiptcustom })
}

main()
  .then(() => console.log('done'))
  .catch(err => console.log(err))
