import { ethers } from 'hardhat'
import { providers } from 'ethers'
import { WETH9__factory } from 'arb-ts'

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

  const WETHProxyConnected = WETH9__factory.connect(
    '0x82aF49447D8a07e3bd95BD0d56f35241523fBab1',
    l2Signer
  )

  const rec = await WETHProxyConnected.upgradeTo(WETH9Logic.address)
  console.warn('upgraded:', rec)
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
