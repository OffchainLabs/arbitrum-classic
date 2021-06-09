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

const l1ERC20GatewayProxy = RinkebyAddresses.l1ERC20GatewayProxy
const l2ERC20GatewayProxy = RinkebyAddresses.l2ERC20GatewayProxy
const l1ProxyAdminAddr = '0x0DbAF24efA2bc9Dd1a6c0530DD252BCcF883B89A'
const l2ProxyAdminAddr = '0x58816566EB91815Cc07f3Ad5230eE0820fe1A19a'

const main = async () => {
  const l1Erc20Bridge = (
    await ethers.getContractAt('L1ERC20Gateway', l1ERC20GatewayProxy)
  ).connect(l1Signer)
  const l2Erc20Bridge = (
    await ethers.getContractAt('L2ERC20Gateway', l2ERC20GatewayProxy)
  ).connect(l2Signer)

  const depositsTopic = await l1Erc20Bridge.filters.OutboundTransferInitiated(
    null,
    null,
    null
  )
  const deposits = (
    await l1Prov.getLogs({
      ...depositsTopic,
      fromBlock: '0x851976', // block erc20 bridge was deployed to rinkeby
      toBlock: 'latest',
    })
  ).map(log => l1Erc20Bridge.interface.parseLog(log))

  console.log('total deposits')
  console.log(deposits.length)

  const tokens = deposits.map(curr => curr.args['token'])

  console.log('unique tokens')
  console.log(new Set(tokens).size)

  const BeaconProxyFactory = (
    await ethers.getContractFactory('BeaconProxyFactory')
  ).connect(l2Signer)

  // query the old `beacon` address. this function still exists in the old l2ERC20Gateway
  const l2Erc20BridgeTemp = await BeaconProxyFactory.attach(l2ERC20GatewayProxy)
  const beacon = await l2Erc20BridgeTemp.beacon()
  console.log({ beacon })

  const beaconProxyFactory = await BeaconProxyFactory.deploy()
  await beaconProxyFactory.deployed()
  console.log('deployed beacon proxy factory')
  const cloneableProxyHash = await beaconProxyFactory.cloneableProxyHash()

  const factoryInitTx = await beaconProxyFactory.initialize(beacon)
  await factoryInitTx.wait()
  console.log('initted beacon proxy factory')

  //   upgrade contracts
  // l1 validation
  const l1ProxyAdmin = ProxyAdmin__factory.connect(l1ProxyAdminAddr, l1Signer)

  const l1Admin = await l1ProxyAdmin.owner()

  if (l1Admin.toLowerCase() !== signer.address.toLowerCase()) {
    throw new Error('Wrong account')
  }

  // l2 validation
  const l2ProxyAdmin = ProxyAdmin__factory.connect(l2ProxyAdminAddr, l2Signer)

  const l2Admin = await l2ProxyAdmin.owner()

  if (l2Admin.toLowerCase() !== signer.address.toLowerCase()) {
    throw new Error('Wrong account')
  }

  console.log('Data validated.')

  // deploy new logics
  console.log('deploying new L1 logic')
  const l1Erc20Factory = await (
    await ethers.getContractFactory('L1ERC20Gateway')
  ).connect(l1Signer)
  const newl1Erc20 = await l1Erc20Factory.deploy()
  await newl1Erc20.deployed()

  console.log('deploying new L2 logic')
  const l2Erc20Factory = await (
    await ethers.getContractFactory('L2ERC20Gateway')
  ).connect(l2Signer)
  const newL2Erc20 = await l2Erc20Factory.deploy()
  await newL2Erc20.deployed()

  // perform upgrade

  console.log('upgrading L1')
  const l1tx = await l1ProxyAdmin.upgrade(
    l1ERC20GatewayProxy,
    newl1Erc20.address
  )
  console.log({ l1tx })
  const l1receipt = await l1tx.wait()
  console.log({ l1receipt })

  console.log('upgrading L2')
  const l2tx = await l2ProxyAdmin.upgrade(
    l2ERC20GatewayProxy,
    newL2Erc20.address
  )
  console.log({ l2tx })
  const l2receipt = await l2tx.wait()
  console.log({ l2receipt })

  //   post upgrade init
  const l2PostUpgradeInit = await l2Erc20Bridge.postUpgradeInit(
    beaconProxyFactory.address
  )
  await l2PostUpgradeInit.wait()
  console.log('initted l1')

  const l1PostUpgradeInit = await l1Erc20Bridge.postUpgradeInit(
    cloneableProxyHash,
    beaconProxyFactory.address
  )
  await l1PostUpgradeInit.wait()
  console.log('initted l2')
}

main()
  .then(() => console.log('done'))
  .catch(err => console.log(err))
