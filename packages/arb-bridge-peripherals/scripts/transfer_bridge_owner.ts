// import { Bridge } from 'arb-ts/src'
import { ethers } from 'hardhat'
// const Confirm = require('prompt-confirm');

import { ProxyAdmin__factory } from 'arb-ts/src/lib/abi/factories/ProxyAdmin__factory'

import RinkebyAddresses from '../deployment-421611.json'
import MainnetAddresses from '../deployment-42161.json'

const NEW_OWNER = '0x1c7d91ccBdBf378bAC0F074678b09CB589184e4E'

const l1Prov = new ethers.providers.JsonRpcProvider(
  'https://mainnet.infura.io/v3/c13a0d6955b14bf181c924bf4c7797fc'
  //  'https://rinkeby.infura.io/v3/c13a0d6955b14bf181c924bf4c7797fc'
)
const l2Prov = new ethers.providers.JsonRpcProvider(
  'https://arb1.arbitrum.io/rpc'
  // 'https://rinkeby.arbitrum.io/rpc'
)

const privKey = process.env['DEVNET_PRIVKEY']
if (!privKey) throw new Error('No DEVNET_PRIVKEY')

const signer = new ethers.Wallet(privKey)

const l1Signer = signer.connect(l1Prov)
const l2Signer = signer.connect(l2Prov)

const wait = (ms: number) => new Promise(res => setTimeout(res, ms))

const l1ProxyAdminAddr = MainnetAddresses.l1ProxyAdmin
const l2ProxyAdminAddr = MainnetAddresses.l2ProxyAdmin

const main = async () => {
  const l1ProxyAdmin = ProxyAdmin__factory.connect(l1ProxyAdminAddr, l1Signer)

  const l1Admin = await l1ProxyAdmin.owner()
  console.log('current L1 proxy admin owner:', l1Admin)

  if (l1Admin.toLowerCase() !== signer.address.toLowerCase()) {
    throw new Error('Wrong account')
  }

  // l2 validation
  const l2ProxyAdmin = ProxyAdmin__factory.connect(l2ProxyAdminAddr, l2Signer)

  const l2Admin = await l2ProxyAdmin.owner()
  console.log('current L2 proxy admin owner:', l2Admin)

  if (l2Admin.toLowerCase() !== signer.address.toLowerCase()) {
    throw new Error('Wrong account')
  }

  console.log('transferring l1 owner:')
  let res = await l1ProxyAdmin.transferOwnership(NEW_OWNER)
  let rec = await res.wait()

  console.log('l1 ownership transferred to ', NEW_OWNER)

  console.log('transferring l2 owner:')
  res = await l2ProxyAdmin.transferOwnership(NEW_OWNER)
  rec = await res.wait()

  console.log('l2 ownership transferred to ', NEW_OWNER)
}

main()
  .then(() => console.log('done'))
  .catch(err => console.log(err))
