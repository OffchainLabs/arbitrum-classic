import { ethers } from 'hardhat'
import { providers } from 'ethers'

const main = async () => {
  const l2PrivKey = process.env['DEVNET_PRIVKEY']
  if (!l2PrivKey) throw new Error('Missing l2 priv key')

  const l2Provider = new providers.JsonRpcProvider(
    'https://arb1.arbitrum.io/rpc'
  )
  const l2Signer = new ethers.Wallet(l2PrivKey, l2Provider)

  const WETH9 = (await ethers.getContractFactory('WETH9')).connect(l2Signer)

  const WETH9Logic = await WETH9.deploy()
  console.log('WETH9 logic deployed at', WETH9Logic.address)

  const L2TransparentUpgradeableProxy = (
    await ethers.getContractFactory('TransparentUpgradeableProxy')
  ).connect(l2Signer)

  const WETH9Proxy = await L2TransparentUpgradeableProxy.deploy(
    WETH9Logic.address,
    '0x18877d60d370893c41b3154ee107e4d520df22d4',
    '0x'
  )
  console.log('WETH9Proxy deployed at', WETH9Proxy.address)

  await WETH9Proxy.init()
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
