// import { Bridge } from 'arb-ts/src'
import { ethers } from 'hardhat'

import { ProxyAdmin__factory } from 'arb-ts/src/lib/abi/factories/ProxyAdmin__factory'

const l1Prov = new ethers.providers.JsonRpcProvider(
  'https://rinkeby.infura.io/v3/c13a0d6955b14bf181c924bf4c7797fc'
)
const l2Prov = new ethers.providers.JsonRpcProvider(
  'https://rinkeby.arbitrum.io/rpc'
)

const privKey = process.env['DEVNET_PRIVKEY']
if (!privKey) throw new Error('No DEVNET_PRIVKEY')

const signer = new ethers.Wallet(privKey)

const wait = (ms: number) => new Promise(res => setTimeout(res, ms))

const main = async () => {
  const l1ProxyAdminAddr = '0x0DbAF24efA2bc9Dd1a6c0530DD252BCcF883B89A'
  const l2ProxyAdminAddr = '0x58816566EB91815Cc07f3Ad5230eE0820fe1A19a'

  const l1GatewayRouterProxy = '0x70C143928eCfFaf9F5b406f7f4fC28Dc43d68380'
  const l2ERC20GatewayProxy = '0x195C107F3F75c4C93Eba7d9a1312F19305d6375f'

  const oldL1GatewayRouterLogic = '0xca47f3A38526dB18818DD4FEeC36Aa02C6BB454E'
  const oldL2ERC20GatewayLogic = '0x4488FDeC54314f7A6EAD0A0058f2BFc698Af70B1'

  // l1 validation
  const l1ProxyAdmin = ProxyAdmin__factory.connect(
    l1ProxyAdminAddr,
    await signer.connect(l1Prov)
  )

  const l1Admin = await l1ProxyAdmin.owner()

  if (l1Admin.toLowerCase() !== signer.address.toLowerCase()) {
    throw new Error('Wrong account')
  }

  const gatewayLogicExpected = await l1ProxyAdmin.getProxyImplementation(
    l1GatewayRouterProxy
  )
  if (
    gatewayLogicExpected.toLowerCase() !== oldL1GatewayRouterLogic.toLowerCase()
  ) {
    throw new Error('Wrong logic l1')
  }

  // l2 validation
  const l2ProxyAdmin = ProxyAdmin__factory.connect(
    l2ProxyAdminAddr,
    await signer.connect(l2Prov)
  )

  const l2Admin = await l2ProxyAdmin.owner()

  if (l2Admin.toLowerCase() !== signer.address.toLowerCase()) {
    throw new Error('Wrong account')
  }

  const erc20LogicExpected = await l2ProxyAdmin.getProxyImplementation(
    l2ERC20GatewayProxy
  )
  if (
    erc20LogicExpected.toLowerCase() !== oldL2ERC20GatewayLogic.toLowerCase()
  ) {
    throw new Error('Wrong logic l2')
  }
  console.log('Data validated.')

  // deploy new logics
  console.log('deploying new L1 logic')
  const gatewayFactory = await (
    await ethers.getContractFactory('L1GatewayRouter')
  ).connect(signer.connect(l1Prov))
  const newL1Gateway = await gatewayFactory.deploy()
  await newL1Gateway.deployed()

  console.log('deploying new L2 logic')
  const l2Erc20Factory = await (
    await ethers.getContractFactory('L2ERC20Gateway')
  ).connect(signer.connect(l2Prov))
  const newL2Erc20 = await l2Erc20Factory.deploy()
  await newL2Erc20.deployed()

  // perform upgrade

  console.log('upgrading L1')
  const l1tx = await l1ProxyAdmin.upgrade(
    l1GatewayRouterProxy,
    newL1Gateway.address
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
}

main()
  .then(() => console.log('done'))
  .catch(err => console.log(err))
