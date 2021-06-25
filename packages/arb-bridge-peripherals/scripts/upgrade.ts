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
  const l1ERC20GatewayProxy = '0x91169Dbb45e6804743F94609De50D511C437572E'

  // l1 validation
  const l1ProxyAdmin = ProxyAdmin__factory.connect(
    l1ProxyAdminAddr,
    await signer.connect(l1Prov)
  )

  const l1Admin = await l1ProxyAdmin.owner()

  if (l1Admin.toLowerCase() !== signer.address.toLowerCase()) {
    throw new Error('Wrong account')
  }

  console.log('Data validated.')

  // deploy new logics
  console.log('deploying new L1 logic')
  const gatewayFactory = await (
    await ethers.getContractFactory('L1ERC20Gateway')
  ).connect(signer.connect(l1Prov))
  const newL1Gateway = await gatewayFactory.deploy()
  await newL1Gateway.deployed()

  // perform upgrade

  console.log('upgrading L1')
  const l1tx = await l1ProxyAdmin.upgrade(
    l1ERC20GatewayProxy,
    newL1Gateway.address
  )
  console.log({ l1tx })
  const l1receipt = await l1tx.wait()
  console.log({ l1receipt })
}

main()
  .then(() => console.log('done'))
  .catch(err => console.log(err))
